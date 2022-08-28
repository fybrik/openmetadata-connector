package databasetypes

import (
	"reflect"

	models "github.com/fybrik/datacatalog-go-models"
)

const MYSQL = "Mysql"

type mysql struct {
	dataBase
	standardFields map[string]bool
}

func NewMysql() *mysql {
	standardFields := map[string]bool{
		DatabaseSchema: true,
		"hostPort":     true,
		"password":     true,
		"scheme":       true,
		"username":     true,
	}
	return &mysql{standardFields: standardFields, dataBase: dataBase{name: MYSQL}}
}

func (m *mysql) TranslateFybrikConfigToOpenMetadataConfig(config map[string]interface{}, credentials *string) map[string]interface{} {
	return config
}

func (m *mysql) TranslateOpenMetadataConfigToFybrikConfig(config map[string]interface{}) map[string]interface{} {
	other := make(map[string]interface{})
	ret := make(map[string]interface{})
	for key, value := range config {
		if _, ok := m.standardFields[key]; ok {
			ret[key] = value
		} else {
			other[key] = value
		}
	}

	ret["other"] = other
	return ret
}

func (m *mysql) TableFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string {
	connectionProperties := createAssetRequest.Details.GetConnection().AdditionalProperties["mysql"].(map[string]interface{})
	assetName := *createAssetRequest.DestinationAssetID
	databaseSchema, found := connectionProperties[DatabaseSchema]
	if found {
		return serviceName + "." + DEFAULT + "." + databaseSchema.(string) + "." + assetName
	}
	return serviceName + "." + DEFAULT + "." + assetName
}

func (m *mysql) CompareServiceConfigurations(requestConfig, serviceConfig map[string]interface{}) bool {
	for property, value := range requestConfig {
		if !reflect.DeepEqual(serviceConfig[property], value) {
			return false
		}
	}
	return true
}

func (m *mysql) DatabaseName(createAssetRequest *models.CreateAssetRequest) string {
	return DEFAULT
}

func (m *mysql) DatabaseFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string {
	return ""
}

func (m *mysql) DatabaseSchemaName(createAssetRequest *models.CreateAssetRequest) string {
	return ""
}

func (m *mysql) DatabaseSchemaFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string {
	return ""
}

func (m *mysql) TableName(createAssetRequest *models.CreateAssetRequest) string {
	return ""
}
