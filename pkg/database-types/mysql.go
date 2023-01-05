// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package databasetypes

import (
	"fmt"
	"reflect"

	"github.com/rs/zerolog"

	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
	"fybrik.io/openmetadata-connector/pkg/utils"
)

type mysql struct {
	dataBase
	standardFields map[string]bool
}

func NewMysql(logger *zerolog.Logger) *mysql {
	standardFields := map[string]bool{
		DatabaseSchema: true,
		HostPort:       true,
		Password:       true,
		Scheme:         true,
		Username:       true,
	}
	return &mysql{standardFields: standardFields, dataBase: dataBase{name: Mysql, logger: logger}}
}

func (m *mysql) TranslateFybrikConfigToOpenMetadataConfig(config map[string]interface{},
	connectionType string, credentials *string) map[string]interface{} {
	return config
}

func (m *mysql) TranslateOpenMetadataConfigToFybrikConfig(tableName string,
	config map[string]interface{}) (map[string]interface{}, string, error) {
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
	return ret, Mysql, nil
}

func (m *mysql) TableFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) (string, error) {
	connectionProperties, ok := utils.InterfaceToMap(createAssetRequest.Details.GetConnection().AdditionalProperties["mysql"], m.logger)
	if !ok {
		return EmptyString, fmt.Errorf(FailedToConvert, AdditionalProperties)
	}
	assetName := *createAssetRequest.DestinationAssetID
	databaseSchema, found := connectionProperties[DatabaseSchema]
	if found {
		return fmt.Sprintf("%s.%s.%s.%s", serviceName, Default, databaseSchema.(string), assetName), nil
	}
	return fmt.Sprintf("%s.%s.%s", serviceName, Default, assetName), nil
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
	return EmptyString
}

func (m *mysql) DatabaseSchemaName(createAssetRequest *models.CreateAssetRequest) string {
	return EmptyString
}

func (m *mysql) DatabaseSchemaFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string {
	return EmptyString
}

func (m *mysql) TableName(createAssetRequest *models.CreateAssetRequest) (string, error) {
	return EmptyString, nil
}
