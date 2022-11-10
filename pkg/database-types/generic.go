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
}

func NewGeneric(logger *zerolog.Logger) *generic {
	return &generic{dataBase: dataBase{name: CustomDatabase, logger: logger}}
}

func (m *generic) TranslateFybrikConfigToOpenMetadataConfig(config map[string]interface{}, credentials *string) map[string]interface{} {
	ret := make(map[string]string)
	for key, value := range config {
		valueStr, ok := value.(string)
		if ok {
			ret[key] = valueStr
		} else {
			jsonStr, _ := json.Marshal(value)
			ret[key] = string(jsonStr)
		}
	}
	return map[string]interface{}{ConnectionOptions: ret}
}

func (m *generic) TranslateOpenMetadataConfigToFybrikConfig(tableName string,
	config map[string]interface{}) (map[string]interface{}, error) {
	configSource, ok := utils.InterfaceToMap(config[ConnectionOptions], m.logger)
	if !ok {
		return nil, fmt.Errorf(FailedToConvert, ConnectionOptions)
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
	return ret, nil
}

func (m *generic) TableFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) (string, error) {
	tableName, err := m.TableName(createAssetRequest)
	if err != nil {
		return EmptyString, err
	}
	return utils.AppendStrings(m.DatabaseSchemaFQN(serviceName, createAssetRequest), tableName), nil
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
	return utils.AppendStrings(serviceName, m.DatabaseName(createAssetRequest))
}

func (m *generic) DatabaseSchemaFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string {
	return utils.AppendStrings(m.DatabaseFQN(serviceName, createAssetRequest),
		m.DatabaseSchemaName(createAssetRequest))
}

func (m *generic) DatabaseSchemaName(createAssetRequest *models.CreateAssetRequest) string {
	return createAssetRequest.DestinationCatalogID
}

func (m *generic) TableName(createAssetRequest *models.CreateAssetRequest) (string, error) {
	return *createAssetRequest.DestinationAssetID, nil
}
