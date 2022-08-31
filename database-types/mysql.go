package databasetypes

import (
	"fmt"
	"reflect"

	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
	"fybrik.io/openmetadata-connector/utils"

	zerolog "github.com/rs/zerolog"
)

const MYSQL = "Mysql"
const Other = "other"

type mysql struct {
	dataBase
	standardFields map[string]bool
}

func NewMysql(logger *zerolog.Logger) *mysql {
	standardFields := map[string]bool{
		DatabaseSchema: true,
		"hostPort":     true,
		"password":     true,
		"scheme":       true,
		"username":     true,
	}
	return &mysql{standardFields: standardFields, dataBase: dataBase{name: MYSQL, logger: logger}}
}

func (m *mysql) TranslateFybrikConfigToOpenMetadataConfig(config map[string]interface{}, credentials *string) map[string]interface{} {
	return config
}

func (m *mysql) TranslateOpenMetadataConfigToFybrikConfig(tableName string, credentials string,
	config map[string]interface{}) map[string]interface{} {
	other := make(map[string]interface{})
	ret := make(map[string]interface{})
	for key, value := range config {
		if _, ok := m.standardFields[key]; ok {
			ret[key] = value
		} else {
			other[key] = value
		}
	}

	ret[Other] = other
	return ret
}

func (m *mysql) TableFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string {
	connectionProperties, ok := utils.InterfaceToMap(createAssetRequest.Details.GetConnection().AdditionalProperties["mysql"])
	if !ok {
		m.logger.Warn().Msg(fmt.Sprintf(FailedToConvert, AdditionalProperties))
		return ""
	}
	assetName := *createAssetRequest.DestinationAssetID
	databaseSchema, found := connectionProperties[DatabaseSchema]
	if found {
		return serviceName + "." + Default + "." + databaseSchema.(string) + "." + assetName
	}
	return serviceName + "." + Default + "." + assetName
}

func (m *mysql) EquivalentServiceConfigurations(requestConfig, serviceConfig map[string]interface{}) bool {
	for property, value := range requestConfig {
		if !reflect.DeepEqual(serviceConfig[property], value) {
			return false
		}
	}
	return true
}

func (m *mysql) DatabaseName(createAssetRequest *models.CreateAssetRequest) string {
	return Default
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
