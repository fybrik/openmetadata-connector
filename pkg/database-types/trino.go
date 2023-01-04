// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package databasetypes //nolint:dupl

import (
	"fmt"
	"reflect"

	"github.com/rs/zerolog"

	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
	"fybrik.io/openmetadata-connector/pkg/utils"
)

type trino struct {
	dataBase
	standardFields map[string]bool
}

func NewTrino(logger *zerolog.Logger) *trino {
	standardFields := map[string]bool{
		DatabaseSchema: true,
		HostPort:       true,
		Password:       true,
		Scheme:         true,
		Username:       true,
	}
	return &trino{standardFields: standardFields, dataBase: dataBase{name: Trino, logger: logger}}
}

func (t *trino) TranslateFybrikConfigToOpenMetadataConfig(config map[string]interface{},
	connectionType string, credentials *string) map[string]interface{} {
	return config
}

func (t *trino) TranslateOpenMetadataConfigToFybrikConfig(tableName string,
	config map[string]interface{}) (map[string]interface{}, string, error) {
	other := make(map[string]interface{})
	ret := make(map[string]interface{})
	for key, value := range config {
		if _, ok := t.standardFields[key]; ok {
			ret[key] = value
		} else {
			other[key] = value
		}
	}

	ret[Other] = other
	return ret, Trino, nil
}

func (t *trino) TableFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) (string, error) {
	connectionProperties, ok := utils.InterfaceToMap(createAssetRequest.Details.GetConnection().AdditionalProperties["trino"], t.logger)
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

func (t *trino) EquivalentServiceConfigurations(requestConfig, serviceConfig map[string]interface{}) bool {
	for property, value := range requestConfig {
		if !reflect.DeepEqual(serviceConfig[property], value) {
			return false
		}
	}
	return true
}

func (t *trino) DatabaseName(createAssetRequest *models.CreateAssetRequest) string {
	return Default
}

func (t *trino) DatabaseFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string {
	return EmptyString
}

func (t *trino) DatabaseSchemaName(createAssetRequest *models.CreateAssetRequest) string {
	return EmptyString
}

func (t *trino) DatabaseSchemaFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string {
	return EmptyString
}

func (t *trino) TableName(createAssetRequest *models.CreateAssetRequest) (string, error) {
	return EmptyString, nil
}
