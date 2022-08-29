package databasetypes

import (
	"errors"
	"reflect"
	"strings"

	zerolog "github.com/rs/zerolog"

	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
	utils "fybrik.io/openmetadata-connector/utils"
	vault "fybrik.io/openmetadata-connector/vault"
)

type s3 struct {
	dataBase
	vaultClientConfiguration map[interface{}]interface{}
	logger                   *zerolog.Logger
}

const REGION = "region"
const ENDPOINT = "endpoint"
const AccessKeyID = "access_key_id"
const SecretAccessID = "secret_access_id"
const AwsRegion = "awsRegion"
const EndPointURL = "endPointURL"
const AwsAccessKeyID = "awsAccessKeyId"
const AwsSecretAccessKey = "awsSecretAccessKey"
const DATALAKE = "Datalake"
const ConfigSource = "configSource"
const BucketName = "bucketName"
const Bucket = "bucket"
const SecurityConfig = "securityConfig"
const S3 = "s3"

var translate = map[string]string{
	REGION:         AwsRegion,
	ENDPOINT:       EndPointURL,
	AccessKeyID:    AwsAccessKeyID,
	SecretAccessID: AwsSecretAccessKey,
}
var translateInv = map[string]string{
	AwsRegion:          REGION,
	EndPointURL:        ENDPOINT,
	AwsAccessKeyID:     AccessKeyID,
	AwsSecretAccessKey: SecretAccessID,
}

func NewS3(vaultClientConfiguration map[interface{}]interface{}, logger *zerolog.Logger) *s3 {
	return &s3{dataBase: dataBase{name: DATALAKE},
		vaultClientConfiguration: vaultClientConfiguration,
		logger:                   logger}
}

const GetTokenFailed = "GetToken failed"
const GetSecretFailed = "GetSecret failed"

func (s *s3) getS3Credentials(vaultClientConfiguration map[interface{}]interface{},
	credentialsPath *string) (string, string, error) {
	client := vault.NewVaultClient(vaultClientConfiguration, s.logger)
	token, err := client.GetToken()
	if err != nil {
		s.logger.Warn().Msg(GetTokenFailed)
		return "", "", errors.New(GetTokenFailed)
	}
	secret, err := client.GetSecret(token, *credentialsPath)
	if err != nil {
		s.logger.Warn().Msg(GetSecretFailed)
		return "", "", errors.New(GetSecretFailed)
	}
	accessKey, secretKey, err := client.ExtractS3CredentialsFromSecret(secret)
	if err != nil {
		s.logger.Warn().Msg("S3 credentials extraction failed")
		return "", "", err
	}
	return accessKey, secretKey, nil
}

func (s *s3) TranslateFybrikConfigToOpenMetadataConfig(config map[string]interface{}, credentialsPath *string) map[string]interface{} {
	ret := make(map[string]interface{})
	configSourceMap := make(map[string]interface{})
	ret["type"] = DATALAKE
	bucketName, found := config[Bucket]
	if found {
		ret[BucketName] = bucketName
	}

	securityMap := make(map[string]interface{})
	securityMap["awsRegion"] = "eu-de" // awsRegion field is mandatory, although it is persumably ignored if endpoint is provided
	for key, value := range config {
		translation, found := translate[key]
		if found {
			securityMap[translation] = value
		}
	}

	if s.vaultClientConfiguration != nil && credentialsPath != nil {
		awsAccessKeyID, awsSecretAccessKey, err := s.getS3Credentials(s.vaultClientConfiguration, credentialsPath)
		if err == nil && awsAccessKeyID != "" && awsSecretAccessKey != "" {
			securityMap["awsAccessKeyId"] = awsAccessKeyID
			securityMap["awsSecretAccessKey"] = awsSecretAccessKey
		}
	}

	configSourceMap[SecurityConfig] = securityMap
	ret[ConfigSource] = configSourceMap
	return ret
}

func (s *s3) TranslateOpenMetadataConfigToFybrikConfig(config map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})

	securityConfig := config[ConfigSource].(map[string]interface{})[SecurityConfig].(map[string]interface{})

	for key, value := range securityConfig {
		if translation, found := translateInv[key]; found {
			ret[translation] = value
		}
	}
	if value, found := config[BucketName]; found {
		ret[Bucket] = value
	}

	return ret
}

func (s *s3) compareConfigSource(fromService, fromRequest map[string]interface{}) bool {
	// ignore some fields, such as 'aws_token' which would appear only serviceSecurityConfig
	serviceSecurityConfig := fromService[SecurityConfig].(map[string]interface{})
	requestSecurityConfig := fromRequest[SecurityConfig].(map[string]interface{})
	for property, value := range requestSecurityConfig {
		if !reflect.DeepEqual(serviceSecurityConfig[property], value) {
			return false
		}
	}
	return true
}

func (s *s3) CompareServiceConfigurations(requestConfig, serviceConfig map[string]interface{}) bool {
	for property, value := range requestConfig {
		if property == ConfigSource {
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
func (s *s3) DatabaseName(createAssetRequest *models.CreateAssetRequest) string {
	return DEFAULT
}

func (s *s3) DatabaseFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string {
	return utils.AppendStrings(serviceName, s.DatabaseName(createAssetRequest))
}

func (s *s3) DatabaseSchemaName(createAssetRequest *models.CreateAssetRequest) string {
	connectionProperties := createAssetRequest.Details.GetConnection().AdditionalProperties[S3].(map[string]interface{})
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
	return ""
}

func (s *s3) DatabaseSchemaFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string {
	return utils.AppendStrings(s.DatabaseFQN(serviceName, createAssetRequest),
		s.DatabaseSchemaName(createAssetRequest))
}

func (s *s3) TableName(createAssetRequest *models.CreateAssetRequest) string {
	connectionProperties := createAssetRequest.Details.GetConnection().AdditionalProperties[S3].(map[string]interface{})
	objectKey, found := connectionProperties["object_key"]
	if found {
		return objectKey.(string)
	}
	split := strings.Split(*createAssetRequest.DestinationAssetID, ".")
	return split[len(split)-1]
}

func (s *s3) TableFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string {
	return utils.AppendStrings(s.DatabaseSchemaFQN(serviceName, createAssetRequest),
		s.TableName(createAssetRequest))
}
