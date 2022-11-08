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

type generic struct {
	dataBase
}

func NewGeneric(logger *zerolog.Logger) *generic {
	return &generic{dataBase: dataBase{name: CustomDatabase, logger: logger}}
}

func (m *generic) TranslateFybrikConfigToOpenMetadataConfig(config map[string]interface{}, credentials *string) map[string]interface{} {
	ret := make(map[string]interface{})
	if raw, ok := config["raw"]; ok {
		rawArray, ok1 := utils.InterfaceToArray(raw, m.logger)
		if !ok1 {
			return nil
		}

		for index, line := range rawArray {
			ret[fmt.Sprintf("line-%d", index)] = line
		}
	}
	return map[string]interface{}{"connectionOptions": ret}
}

func (m *generic) TranslateOpenMetadataConfigToFybrikConfig(tableName string,
	config map[string]interface{}) (map[string]interface{}, error) {
	return config, nil
}

func (m *generic) TableFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) (string, error) {
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

func (m *generic) EquivalentServiceConfigurations(requestConfig, serviceConfig map[string]interface{}) bool {
	for property, value := range requestConfig {
		if !reflect.DeepEqual(serviceConfig[property], value) {
			return false
		}
	}
	return true
}

func (m *generic) DatabaseName(createAssetRequest *models.CreateAssetRequest) string {
	return Default
}

func (m *generic) DatabaseFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string {
	return EmptyString
}

func (m *generic) DatabaseSchemaName(createAssetRequest *models.CreateAssetRequest) string {
	return EmptyString
}

func (m *generic) DatabaseSchemaFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string {
	return EmptyString
}

func (m *generic) TableName(createAssetRequest *models.CreateAssetRequest) (string, error) {
	return EmptyString, nil
}
