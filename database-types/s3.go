package database_types

import (
	"reflect"
	"strings"

	models "github.com/fybrik/datacatalog-go-models"
	utils "github.com/fybrik/openmetadata-connector/utils"
	vault "github.com/fybrik/openmetadata-connector/vault"
	"github.com/rs/zerolog"
)

type s3 struct {
	Translate                map[string]string
	TranslateInv             map[string]string
	VaultClientConfiguration map[interface{}]interface{}
	logger                   zerolog.Logger
}

func NewS3(vaultClientConfiguration map[interface{}]interface{}, logger zerolog.Logger) *s3 {
	translate := map[string]string{
		"region":           "awsRegion",
		"endpoint":         "endPointURL",
		"access_key_id":    "awsAccessKeyId",
		"secret_access_id": "awsSecretAccessKey",
	}
	translateInv := map[string]string{
		"awsRegion":          "region",
		"endPointURL":        "endpoint",
		"awsAccessKeyId":     "access_key_id",
		"awsSecretAccessKey": "secret_access_id",
	}
	return &s3{Translate: translate,
		TranslateInv:             translateInv,
		VaultClientConfiguration: vaultClientConfiguration,
		logger:                   logger}
}

func (s *s3) getS3Credentials(vaultClientConfiguration map[interface{}]interface{}, credentialsPath *string) (string, string) {
	client := vault.NewVaultClient(vaultClientConfiguration, s.logger)
	token, err := client.GetToken()
	if err != nil {
		s.logger.Error().Msg("GetToken failed")
		return "", ""
	}
	secret, err := client.GetSecret(token, *credentialsPath)
	if err != nil {
		s.logger.Error().Msg("GetSecret failed")
		return "", ""
	}
	return client.ExtractS3CredentialsFromSecret(secret)
}

func (s *s3) TranslateFybrikConfigToOpenMetadataConfig(config map[string]interface{}, credentialsPath *string) map[string]interface{} {
	ret := make(map[string]interface{})
	configSourceMap := make(map[string]interface{})
	ret["type"] = "Datalake"
	bucketName, found := config["bucket"]
	if found {
		ret["bucketName"] = bucketName
	}

	securityMap := make(map[string]interface{})
	securityMap["awsRegion"] = "eu-de" // awsRegion field is mandatory, although it is persumably ignored if endpoint is provided
	for key, value := range config {
		translation, found := s.Translate[key]
		if found {
			securityMap[translation] = value
		}
	}

	if s.VaultClientConfiguration != nil && credentialsPath != nil {
		awsAccessKeyId, awsSecretAccessKey := s.getS3Credentials(s.VaultClientConfiguration, credentialsPath)
		if awsAccessKeyId != "" && awsSecretAccessKey != "" {
			securityMap["awsAccessKeyId"] = awsAccessKeyId
			securityMap["awsSecretAccessKey"] = awsSecretAccessKey
		}
	}

	configSourceMap["securityConfig"] = securityMap
	ret["configSource"] = configSourceMap
	return ret
}

func (s *s3) TranslateOpenMetadataConfigToFybrikConfig(config map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})

	securityConfig := config["configSource"].(map[string]interface{})["securityConfig"].(map[string]interface{})

	for key, value := range securityConfig {
		if translation, found := s.TranslateInv[key]; found {
			ret[translation] = value
		}
	}
	if value, found := config["bucketName"]; found {
		ret["bucket"] = value
	}

	return ret
}

func (m *s3) OMTypeName() string {
	return "Datalake"
}

func (s *s3) compareConfigSource(fromService map[string]interface{}, fromRequest map[string]interface{}) bool {
	// ignore some fields, such as 'aws_token' which would appear only serviceSecurityConfig
	serviceSecurityConfig := fromService["securityConfig"].(map[string]interface{})
	requestSecurityConfig := fromRequest["securityConfig"].(map[string]interface{})
	for property, value := range requestSecurityConfig {
		if !reflect.DeepEqual(serviceSecurityConfig[property], value) {
			return false
		}
	}
	return true
}

func (s *s3) CompareServiceConfigurations(requestConfig map[string]interface{}, serviceConfig map[string]interface{}) bool {
	for property, value := range requestConfig {
		if property == "configSource" {
			if !s.compareConfigSource(serviceConfig[property].(map[string]interface{}), value.(map[string]interface{})) {
				return false
			}
		} else {
			if !reflect.DeepEqual(serviceConfig[property], value) {
				return false
			}
		}
	}
	return true
}
func (s *s3) DatabaseName(createAssetRequest models.CreateAssetRequest) string {
	return "default"
}

func (s *s3) DatabaseFQN(serviceName string, createAssetRequest models.CreateAssetRequest) string {
	return utils.AppendStrings(serviceName, s.DatabaseSchemaName(createAssetRequest))
}

func (s *s3) DatabaseSchemaName(createAssetRequest models.CreateAssetRequest) string {
	connectionProperties := createAssetRequest.Details.GetConnection().AdditionalProperties["s3"].(map[string]interface{})
	bucket, found := connectionProperties["bucket"]
	if found {
		return bucket.(string)
	} else {
		assetID := *createAssetRequest.DestinationAssetID
		split := strings.Split(assetID, ".")
		if len(split) > 1 {
			return split[len(split)-2]
		}
	}
	s.logger.Warn().Msg("Could not determine the name of the DatabaseSchema (bucket)")
	return ""
}

func (s *s3) DatabaseSchemaFQN(serviceName string, createAssetRequest models.CreateAssetRequest) string {

	return utils.AppendStrings(s.DatabaseFQN(serviceName, createAssetRequest),
		s.DatabaseSchemaName(createAssetRequest))
}

func (s *s3) TableName(createAssetRequest models.CreateAssetRequest) string {
	connectionProperties := createAssetRequest.Details.GetConnection().AdditionalProperties["s3"].(map[string]interface{})
	objectKey, found := connectionProperties["object_key"]
	if found {
		return objectKey.(string)
	} else {
		split := strings.Split(*createAssetRequest.DestinationAssetID, ".")
		return split[len(split)-1]
	}
}

func (s *s3) TableFQN(serviceName string, createAssetRequest models.CreateAssetRequest) string {
	return utils.AppendStrings(s.DatabaseSchemaFQN(serviceName, createAssetRequest),
		s.TableName(createAssetRequest))
}
