// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package databasetypes

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/rs/zerolog"

	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
	"fybrik.io/openmetadata-connector/pkg/utils"
)

type generic struct {
	dataBase
	databaseTypeParent
}

func NewGeneric(logger *zerolog.Logger) *generic {
	return &generic{dataBase: dataBase{name: CustomDatabase, logger: logger}}
}

// OM databaseService-s of type CustomDatabase expect the configuration key-value pairs to be placed
// within the 'connectionOptions' field. The values in the key-value pairs must be strings
func (m *generic) TranslateFybrikConfigToOpenMetadataConfig(config map[string]interface{},
	connectionType string, credentials *string) map[string]interface{} {
	ret := make(map[string]string)
	for key, value := range config {
		valueStr, ok := value.(string)
		if ok {
			ret[key] = valueStr
		} else {
			jsonStr, err := json.Marshal(value)
			if err != nil {
				m.logger.Warn().Err(err).Msg(fmt.Sprintf("failed to transfrom value %s for key %s", value, key))
			} else {
				ret[key] = string(jsonStr)
			}
		}
	}
	ret[ConnectionType] = connectionType
	return map[string]interface{}{ConnectionOptions: ret}
}

// take the configuration key-value pairs from the `connectionOptions` field, and convert the
// JSON fields to maps
func (m *generic) TranslateOpenMetadataConfigToFybrikConfig(tableName string,
	config map[string]interface{}) (map[string]interface{}, string, error) {
	connectionType := Generic
	connectionOption, ok := config[ConnectionOptions]
	if !ok {
		return nil, connectionType, fmt.Errorf("%s missing from configuration", ConnectionOptions)
	}
	configSource, ok := utils.InterfaceToMap(connectionOption, m.logger)
	if !ok {
		return nil, connectionType, fmt.Errorf(FailedToConvert, ConnectionOptions)
	}
	ret := make(map[string]interface{})
	for key, value := range configSource {
		valueMap := make(map[string]interface{})
		err := json.Unmarshal([]byte(fmt.Sprint(value)), &valueMap)
		if err == nil {
			ret[key] = valueMap
		} else {
			ret[key] = value.(string)
		}
	}
	if connectionTypeVal, ok := ret[ConnectionType]; ok {
		delete(ret, ConnectionType)
		connectionType = connectionTypeVal.(string)
	}
	return ret, connectionType, nil
}

func (m *generic) EquivalentServiceConfigurations(requestConfig, serviceConfig map[string]interface{}) bool {
	return reflect.DeepEqual(requestConfig, serviceConfig)
}

func (m *generic) DatabaseName(createAssetRequest *models.CreateAssetRequest) string {
	return Default
}

func (m *generic) DatabaseSchemaName(createAssetRequest *models.CreateAssetRequest) string {
	return createAssetRequest.DestinationCatalogID
}

func (m *generic) TableName(createAssetRequest *models.CreateAssetRequest) (string, error) {
	return *createAssetRequest.DestinationAssetID, nil
}
