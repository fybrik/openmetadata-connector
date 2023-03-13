// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package databasetypes

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/rs/zerolog"

	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
	"fybrik.io/openmetadata-connector/pkg/utils"
	"fybrik.io/openmetadata-connector/pkg/vault"
)

type mysql struct {
	dataBase
	vaultClientConfiguration map[interface{}]interface{}
	standardFields           map[string]bool
}

func NewMysql(vaultClientConfiguration map[interface{}]interface{}, logger *zerolog.Logger) *mysql {
	standardFields := map[string]bool{
		DatabaseSchema: true,
		HostPort:       true,
		Password:       true,
		Scheme:         true,
		Username:       true,
	}
	return &mysql{
		standardFields:           standardFields,
		dataBase:                 dataBase{name: Mysql, logger: logger},
		vaultClientConfiguration: vaultClientConfiguration,
	}
}

func (m *mysql) getCredentials(vaultClientConfiguration map[interface{}]interface{}, //nolint:dupl
	credentialsPath *string) (string, string, error) {
	client := vault.NewVaultClient(vaultClientConfiguration, m.logger, utils.HTTPClient)
	secrets, err := client.GetSecretMap(credentialsPath)
	if err != nil {
		m.logger.Warn().Msg("MySQL credentials extraction failed")
		return EmptyString, EmptyString, err
	}
	requiredFields := []string{Username, Password}
	secretStrings := utils.InterfaceMapToStringMap(secrets, requiredFields, m.logger)
	if secretStrings == nil {
		m.logger.Warn().Msg(fmt.Sprintf(SomeRequiredFieldsMissing, requiredFields))
		return EmptyString, EmptyString, fmt.Errorf(SomeRequiredFieldsMissing, requiredFields)
	}
	return secretStrings[Username], secretStrings[Password], nil
}

func (m *mysql) TranslateFybrikConfigToOpenMetadataConfig(config map[string]interface{},
	connectionType string, credentials *string) map[string]interface{} {
	if m.vaultClientConfiguration != nil && credentials != nil {
		username, password, err := m.getCredentials(m.vaultClientConfiguration, credentials)
		if err == nil && username != EmptyString && password != EmptyString {
			config[Username] = username
			config[Password] = password
		}
	}
	host, ok1 := config[Host]
	port, ok2 := config[Port]
	if ok1 && ok2 {
		delete(config, Host)
		delete(config, Port)
		config[HostPort] = fmt.Sprintf("%s:%.0f", host, port)
	}

	if database, ok := config[Database]; ok {
		delete(config, Database)
		config[DatabaseSchema] = database
	}

	config[Scheme] = "mysql+pymysql"
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

	// remove sensitive information
	delete(ret, Username)
	delete(ret, Password)

	if databaseSchema, ok := config[DatabaseSchema]; ok {
		delete(ret, DatabaseSchema)
		ret[Database] = databaseSchema
	}

	if hostPort, ok := config[HostPort]; ok {
		delete(ret, HostPort)
		s := strings.Split(hostPort.(string), ":")
		ret[Host] = s[0]
		if len(s) > 1 {
			if port, err := strconv.Atoi(s[1]); err == nil {
				ret[Port] = port
			}
		}
	}

	return ret, MysqlLowercase, nil
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

func (m *mysql) DatabaseSchemaName(createAssetRequest *models.CreateAssetRequest) string {
	connectionProperties, ok := utils.InterfaceToMap(createAssetRequest.Details.GetConnection().AdditionalProperties[MysqlLowercase], m.logger)
	if !ok {
		return EmptyString
	}
	databaseSchema, found := connectionProperties[DatabaseSchema]
	if found {
		return databaseSchema.(string)
	}
	return EmptyString
}

func (m *mysql) TableName(createAssetRequest *models.CreateAssetRequest) (string, error) {
	return *createAssetRequest.DestinationAssetID, nil
}
