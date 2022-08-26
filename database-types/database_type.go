package databasetypes

import models "github.com/fybrik/datacatalog-go-models"

const DEFAULT = "default"
const DatabaseSchema = "databaseSchema"

type DatabaseType interface {
	OMTypeName() string
	TranslateFybrikConfigToOpenMetadataConfig(map[string]interface{}, *string) map[string]interface{}
	TranslateOpenMetadataConfigToFybrikConfig(map[string]interface{}) map[string]interface{}
	CompareServiceConfigurations(map[string]interface{}, map[string]interface{}) bool
	DatabaseName(createAssetRequest *models.CreateAssetRequest) string
	DatabaseFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string
	DatabaseSchemaName(createAssetRequest *models.CreateAssetRequest) string
	DatabaseSchemaFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string
	TableName(createAssetRequest *models.CreateAssetRequest) string
	TableFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string
}
