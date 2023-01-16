// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package databasetypes

import (
	zerolog "github.com/rs/zerolog"

	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
	"fybrik.io/openmetadata-connector/pkg/utils"
)

type DatabaseType interface {
	// OMTypeName() returns the OpenMetadata name for different connections types. For instance, it
	// return "Mysql" for MYSQL and "Deltalake" for s3
	OMTypeName() string

	// translate the connection information from the Fybrik format to the OM format.
	// 'credentials' is the URI of a vault secret
	TranslateFybrikConfigToOpenMetadataConfig(config map[string]interface{},
		connectionType string, credentials *string) map[string]interface{}

	// translate the connection information from the OM format to the Fybrik format.
	// also returns the connection type
	TranslateOpenMetadataConfigToFybrikConfig(tableName string,
		config map[string]interface{}) (map[string]interface{}, string, error)

	// In checking whether a certain databaseService already exists, compare whether two
	// OM configuration informations are equivalent.Return 'true' if they are
	EquivalentServiceConfigurations(map[string]interface{}, map[string]interface{}) bool

	// The hierarchy in OM is:
	// - Database Service
	// - Database
	// - Database Schema
	// - Table

	// DatabaseName returns the name of the asset Database, e.g. 'default'
	DatabaseName(createAssetRequest *models.CreateAssetRequest) string

	// DatabaseSchemaName returns the name of the asset DatabaseSchema, e.g. 'fake-csv-bucket'
	DatabaseSchemaName(createAssetRequest *models.CreateAssetRequest) string

	// TableName returns the name of the asset Table, e.g. '"fake.csv"'
	TableName(createAssetRequest *models.CreateAssetRequest) (string, error)
}

type dataBase struct {
	name   string
	logger *zerolog.Logger
}

func (db dataBase) OMTypeName() string {
	return db.name
}

// DatabaseFQN returns the Fully Qualified Name of the asset Database, e.g. 'openmetadata-s3.default'
func DatabaseFQN(p DatabaseType, serviceName string, createAssetRequest *models.CreateAssetRequest) string {
	return utils.AppendStrings(serviceName, p.DatabaseName(createAssetRequest))
}

// DatabaseSchemaFQN returns the Fully Qualified Name of the asset DatabaseSchema,
// e.g. 'openmetadata-s3.default.fake-csv-bucket'
func DatabaseSchemaFQN(p DatabaseType, serviceName string, createAssetRequest *models.CreateAssetRequest) string {
	return utils.AppendStrings(DatabaseFQN(p, serviceName, createAssetRequest),
		p.DatabaseSchemaName(createAssetRequest))
}

// TableFQN returns the Fully Qualified Name of the asset DatabaseSchema,
// e.g. 'openmetadata-s3.default.fake-csv-bucket."fake.csv"'
func TableFQN(p DatabaseType, serviceName string, createAssetRequest *models.CreateAssetRequest) (string, error) {
	tableName, err := p.TableName(createAssetRequest)
	if err != nil {
		return EmptyString, err
	}
	return utils.AppendStrings(DatabaseSchemaFQN(p, serviceName, createAssetRequest), tableName), nil
}
