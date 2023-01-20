// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package databasetypes

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/rs/zerolog"

	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
	"fybrik.io/openmetadata-connector/pkg/utils"
	"fybrik.io/openmetadata-connector/pkg/vault"
)

type s3 struct {
	dataBase
	vaultClientConfiguration map[interface{}]interface{}
}

var translate = map[string]string{
	Region:         AwsRegion,
	Endpoint:       EndPointURL,
	AccessKeyID:    AwsAccessKeyID,
	SecretAccessID: AwsSecretAccessKey,
}
var translateInv = map[string]string{
	AwsRegion:          Region,
	EndPointURL:        Endpoint,
	AwsAccessKeyID:     AccessKeyID,
	AwsSecretAccessKey: SecretAccessID,
}

func NewS3(vaultClientConfiguration map[interface{}]interface{}, logger *zerolog.Logger) *s3 {
	return &s3{dataBase: dataBase{name: Datalake, logger: logger},
		vaultClientConfiguration: vaultClientConfiguration}
}

func (s *s3) getS3Credentials(vaultClientConfiguration map[interface{}]interface{}, //nolint:dupl
	credentialsPath *string) (string, string, error) {
	client := vault.NewVaultClient(vaultClientConfiguration, s.logger, utils.HTTPClient)
	secrets, err := client.GetSecretMap(credentialsPath)
	if err != nil {
		s.logger.Warn().Msg("S3 credentials extraction failed")
		return EmptyString, EmptyString, err
	}
	requiredFields := []string{AccessKey, SecretKey}
	secretStrings := utils.InterfaceMapToStringMap(secrets, requiredFields)
	if secretStrings == nil {
		s.logger.Warn().Msg(fmt.Sprintf(SomeRequiredFieldsMissing, requiredFields))
		return EmptyString, EmptyString, fmt.Errorf(SomeRequiredFieldsMissing, requiredFields)
	}
	return secretStrings[AccessKey], secretStrings[SecretKey], nil
}

func (s *s3) TranslateFybrikConfigToOpenMetadataConfig(config map[string]interface{},
	connectionType string, credentialsPath *string) map[string]interface{} {
	ret := make(map[string]interface{})
	configSourceMap := make(map[string]interface{})
	ret[Type] = s.OMTypeName()
	bucketName, found := config[Bucket]
	if found {
		ret[BucketName] = bucketName
	}

	securityMap := make(map[string]interface{})
	securityMap[AwsRegion] = "eu-de" // awsRegion field is mandatory, although it is presumably ignored if endpoint is provided
	for key, value := range config {
		translation, found := translate[key]
		if found {
			securityMap[translation] = value
		}
	}

	if s.vaultClientConfiguration != nil && credentialsPath != nil {
		awsAccessKeyID, awsSecretAccessKey, err := s.getS3Credentials(s.vaultClientConfiguration, credentialsPath)
		if err == nil && awsAccessKeyID != EmptyString && awsSecretAccessKey != EmptyString {
			securityMap[AwsAccessKeyID] = awsAccessKeyID
			securityMap[AwsSecretAccessKey] = awsSecretAccessKey
		}
	}

	configSourceMap[SecurityConfig] = securityMap
	ret[ConfigSource] = configSourceMap
	return ret
}

func (s *s3) TranslateOpenMetadataConfigToFybrikConfig(tableName string,
	config map[string]interface{}) (map[string]interface{}, string, error) {
	ret := make(map[string]interface{})
	ret[ObjectKey] = tableName

	configSource, ok := utils.InterfaceToMap(config[ConfigSource], s.logger)
	if !ok {
		return nil, S3, fmt.Errorf(FailedToConvert, ConfigSource)
	}
	securityConfig, ok := utils.InterfaceToMap(configSource[SecurityConfig], s.logger)
	if !ok {
		return nil, S3, fmt.Errorf(FailedToConvert, SecurityConfig)
	}

	for key, value := range securityConfig {
		if translation, found := translateInv[key]; found {
			ret[translation] = value
		}
	}
	if value, found := config[BucketName]; found {
		ret[Bucket] = value
	}

	delete(ret, AccessKeyID)
	delete(ret, SecretAccessID)

	return ret, S3, nil
}

// Check whether the fields in the 'configSource' section are equivalent.
// They don't have to be identical. We allow additional fields (e.g. 'aws_token')
// in the OM configuration, but we require that all fields in the request configuration
// also appear in the OM configuration, and that the values be identical
func (s *s3) equivalentConfigSource(fromService, fromRequest map[string]interface{}) bool {
	// ignore some fields, such as 'aws_token' which would appear only serviceSecurityConfig
	serviceSecurityConfig, ok := utils.InterfaceToMap(fromService[SecurityConfig], s.logger)
	if !ok {
		s.logger.Warn().Msg(fmt.Sprintf(FailedToConvert, "OM "+SecurityConfig))
		return false
	}
	requestSecurityConfig, ok := utils.InterfaceToMap(fromRequest[SecurityConfig], s.logger)
	if !ok {
		s.logger.Warn().Msg(fmt.Sprintf(FailedToConvert, "Request "+SecurityConfig))
		return false
	}
	for property, value := range requestSecurityConfig {
		if !reflect.DeepEqual(serviceSecurityConfig[property], value) {
			return false
		}
	}
	return true
}

func (s *s3) EquivalentServiceConfigurations(requestConfig, serviceConfig map[string]interface{}) bool {
	for property, value := range requestConfig {
		if property == ConfigSource {
			servicePropertyMap, ok := utils.InterfaceToMap(serviceConfig[property], s.logger)
			if !ok {
				s.logger.Warn().Msg(fmt.Sprintf(FailedToConvert, property))
				return false
			}
			valueMap, ok := utils.InterfaceToMap(value, s.logger)
			if !ok {
				s.logger.Warn().Msg(fmt.Sprintf(FailedToConvert, Value))
				return false
			}
			if !s.equivalentConfigSource(servicePropertyMap, valueMap) {
				return false
			}
		} else if !reflect.DeepEqual(serviceConfig[property], value) {
			return false
		}
	}
	return true
}
func (s *s3) DatabaseName(createAssetRequest *models.CreateAssetRequest) string {
	return Default
}

func (s *s3) DatabaseSchemaName(createAssetRequest *models.CreateAssetRequest) string {
	connectionProperties, ok := utils.InterfaceToMap(createAssetRequest.Details.GetConnection().
		AdditionalProperties[S3], s.logger)
	if !ok {
		s.logger.Warn().Msg(fmt.Sprintf(FailedToConvert, AdditionalProperties))
		return EmptyString
	}
	bucket, found := connectionProperties[Bucket]
	if found {
		return bucket.(string)
	}

	assetID := *createAssetRequest.DestinationAssetID
	split := strings.Split(assetID, ".")
	if len(split) > 1 {
		return split[len(split)-2]
	}

	s.logger.Warn().Msg("Could not determine the name of the DatabaseSchema (bucket)")
	return EmptyString
}

func (s *s3) TableName(createAssetRequest *models.CreateAssetRequest) (string, error) {
	connectionProperties, ok := utils.InterfaceToMap(createAssetRequest.Details.GetConnection().AdditionalProperties[S3], s.logger)
	if !ok {
		return EmptyString, fmt.Errorf(FailedToConvert, AdditionalProperties)
	}
	objectKey, found := connectionProperties[ObjectKey]
	if found {
		return objectKey.(string), nil
	}
	split := strings.Split(*createAssetRequest.DestinationAssetID, ".")
	return split[len(split)-1], nil
}
