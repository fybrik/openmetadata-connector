// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package databasetypes

import (
	zerolog "github.com/rs/zerolog"

	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
)

type DatabaseType interface {
	// OMTypeName() returns the OpenMetadata name for different connections types. For instance, it
	// return "Mysql" for MYSQL and "Deltalake" for s3
	OMTypeName() string

	// translate the connection information from the Fybrik format to the OM format
	TranslateFybrikConfigToOpenMetadataConfig(map[string]interface{}, *string) map[string]interface{}

	// translate the connection information from the OM format to the Fybrik format
	TranslateOpenMetadataConfigToFybrikConfig(string, string, map[string]interface{}) (map[string]interface{}, error)

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

	// DatabaseFQN returns the Fully Qualified Name of the asset Database, e.g. 'openmetadata-s3.default'
	DatabaseFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string

	// DatabaseSchemaName returns the name of the asset DatabaseSchema, e.g. 'fake-csv-bucket'
	DatabaseSchemaName(createAssetRequest *models.CreateAssetRequest) string

	// DatabaseSchemaFQN returns the Fully Qualified Name of the asset DatabaseSchema,
	// e.g. 'openmetadata-s3.default.fake-csv-bucket'
	DatabaseSchemaFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string

	// TableName returns the name of the asset Table, e.g. '"fake.csv"'
	TableName(createAssetRequest *models.CreateAssetRequest) string

	// TableFQN returns the Fully Qualified Name of the asset DatabaseSchema,
	// e.g. 'openmetadata-s3.default.fake-csv-bucket."fake.csv"'
	TableFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) (string, error)
}

type dataBase struct {
	name   string
	logger *zerolog.Logger
}

func (db dataBase) OMTypeName() string {
	return db.name
}
