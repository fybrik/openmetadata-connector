package databasetypes

import models "fybrik.io/openmetadata-connector/datacatalog-go-models"

const Default = "default"
const DatabaseSchema = "databaseSchema"

type DatabaseType interface {
	// OMTypeName() returns the OpenMetadata name for different connections types. For instance, it
	// return "Mysql" for MYSQL and "Deltalake" for s3
	OMTypeName() string

	// translate the connection information from the Fybrik format to the OM format
	TranslateFybrikConfigToOpenMetadataConfig(map[string]interface{}, *string) map[string]interface{}

	// translate the connection information from the OM format to the Fybrik format
	TranslateOpenMetadataConfigToFybrikConfig(string, string, map[string]interface{}) map[string]interface{}

	// In checking whether a certain databaseService already exists, compare whether two
	// OM configuration informations are equivalent.Return 'true' if they are
	EquivalentServiceConfigurations(map[string]interface{}, map[string]interface{}) bool

	// The hierarchy in OM is:
	// - Database Service
	// - Database
	// - Database Schema
	// - Table
	// Consider an asset called: openmetadata-s3.default.fake-csv-bucket."fake.csv"
	// For this asset:
	// 'openmetadata-s3' is the name of the Database Service
	// DatabaseName() will return 'default'
	// DatabaseFQN() will return 'openmetadata-s3.default'
	// DatabaseSchemaName() will return 'fake-csv-bucket'
	// DatabaseSchemaFQN() will return 'openmetadata-s3.default.fake-csv-bucket'
	// TableName() will return '"fake.csv"'
	// TableFQN() will return 'openmetadata-s3.default.fake-csv-bucket."fake.csv"'
	DatabaseName(createAssetRequest *models.CreateAssetRequest) string
	DatabaseFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string
	DatabaseSchemaName(createAssetRequest *models.CreateAssetRequest) string
	DatabaseSchemaFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string
	TableName(createAssetRequest *models.CreateAssetRequest) string
	TableFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string
}

type dataBase struct {
	name string
}

func (db dataBase) OMTypeName() string {
	return db.name
}
