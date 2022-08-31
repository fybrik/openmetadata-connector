package openapiconnectorcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	zerolog "github.com/rs/zerolog"

	database_types "fybrik.io/openmetadata-connector/database-types"
	client "fybrik.io/openmetadata-connector/datacatalog-go-client"
	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
	api "fybrik.io/openmetadata-connector/datacatalog-go/go"
	utils "fybrik.io/openmetadata-connector/utils"
)

func getTag(ctx context.Context, c *client.APIClient, tagFQN string) client.TagLabel {
	if strings.Count(tagFQN, ".") == 0 {
		// Since this is not a 'category.primary' or 'category.primary.secondary' format,
		// we will translate it to 'Fybrik.tagFQN'. We try to create it
		// (whether it exists or not)
		createTag := *client.NewCreateTag(tagFQN, tagFQN)
		_, r, _ := c.TagsApi.CreatePrimaryTag(ctx, Fybrik).CreateTag(createTag).Execute()
		defer r.Body.Close()
		tagFQN = Fybrik + "." + tagFQN
	}
	return client.TagLabel{
		LabelType: "Manual",
		Source:    "Tag",
		State:     "Confirmed",
		TagFQN:    tagFQN,
	}
}

func tagColumn(ctx context.Context, c *client.APIClient, columns []client.Column,
	colName string, colTags map[string]interface{}) []client.Column {
	// traverse columns
	for i := range columns {
		// search for colName
		if columns[i].Name == colName {
			for tag := range colTags {
				columns[i].Tags = append(columns[i].Tags, getTag(ctx, c, tag))
			}
			return columns
		}
	}
	return columns
}

func (s *OpenMetadataApiService) prepareOpenMetadataForFybrik() bool {
	ctx := context.Background()
	c := s.getOpenMetadataClient()

	// Create Tag Category for Fybrik
	_, r, err := c.TagsApi.CreateTagCategory(ctx).CreateTagCategory(*client.NewCreateTagCategory("Classification",
		"Parent Category for all Fybrik labels", Fybrik)).Execute()
	if err != nil {
		s.logger.Trace().Msg("Failed to create the Fybrik Tag category. Maybe it already exists.")
	} else {
		defer r.Body.Close()
	}

	// Find the ID for the 'table' entity
	var tableID string

	typeList, r, err := c.MetadataApi.ListTypes(ctx).Category("entity").Limit(TypeListLengthLimit).Execute()
	if err != nil {
		s.logger.Warn().Msg(ErrorInPrepareOpenMetadataForFybrik)
		s.logger.Warn().Msg("Most likely OpenMetadata is not up yet")
		return false
	}
	defer r.Body.Close()

	for i := range typeList.Data {
		if *typeList.Data[i].FullyQualifiedName == "table" {
			tableID = *typeList.Data[i].Id
			break
		}
	}

	if tableID == "" {
		s.logger.Error().Msg(ErrorInPrepareOpenMetadataForFybrik)
		s.logger.Error().Msg("Failed to find the ID for entity 'table'")
		return false
	}

	// Find the ID for the 'string' type
	var stringID string
	typeList, r, err = c.MetadataApi.ListTypes(ctx).Category("field").Limit(TypeListLengthLimit).Execute()
	if err != nil {
		s.logger.Error().Msg(fmt.Sprintf("Error when calling `MetadataApi.ListTypes``: %v\n", err))
		s.logger.Error().Msg(fmt.Sprintf(FullHTTPResponse, r))
		return false
	}
	defer r.Body.Close()

	for i := range typeList.Data {
		if *typeList.Data[i].FullyQualifiedName == "string" {
			stringID = *typeList.Data[i].Id
			break
		}
	}

	if stringID == "" {
		s.logger.Error().Msg(ErrorInPrepareOpenMetadataForFybrik)
		s.logger.Error().Msg("Failed to find the ID for entity 'string'")
		return false
	}

	// Add custom properties for tables
	r1, _ := c.MetadataApi.AddProperty(ctx, tableID).CustomProperty(*client.NewCustomProperty(
		"The vault plugin path where the destination data credentials will be stored as kubernetes secrets", Credentials,
		*client.NewEntityReference(stringID, String))).Execute()
	defer r1.Body.Close()
	r2, _ := c.MetadataApi.AddProperty(ctx, tableID).CustomProperty(*client.NewCustomProperty(
		"Connection type, e.g.: s3 or mysql", ConnectionType,
		*client.NewEntityReference(stringID, String))).Execute()
	defer r2.Body.Close()
	r3, _ := c.MetadataApi.AddProperty(ctx, tableID).CustomProperty(*client.NewCustomProperty(
		"Name of the resource", Name,
		*client.NewEntityReference(stringID, String))).Execute()
	defer r3.Body.Close()
	r4, _ := c.MetadataApi.AddProperty(ctx, tableID).CustomProperty(*client.NewCustomProperty(
		"Geography of the resource", Geography,
		*client.NewEntityReference(stringID, String))).Execute()
	defer r4.Body.Close()
	r5, _ := c.MetadataApi.AddProperty(ctx, tableID).CustomProperty(*client.NewCustomProperty(
		"Owner of the resource", Owner,
		*client.NewEntityReference(stringID, String))).Execute()
	defer r5.Body.Close()
	r6, _ := c.MetadataApi.AddProperty(ctx, tableID).CustomProperty(*client.NewCustomProperty(
		"Data format", DataFormat,
		*client.NewEntityReference(stringID, String))).Execute()
	r6.Body.Close()

	return true
}

// NewOpenMetadataApiService creates a new api service.
// It is initialized base on the configuration
func NewOpenMetadataAPIService(conf map[interface{}]interface{}, logger *zerolog.Logger) OpenMetadataApiServicer {
	var SleepIntervalMS int
	var NumRetries int

	var vaultConf map[interface{}]interface{} = nil
	if vaultConfMap, ok := conf["vault"]; ok {
		vaultConf = vaultConfMap.(map[interface{}]interface{})
	}

	value, ok := conf["openmetadata_sleep_interval"]
	if ok {
		SleepIntervalMS = value.(int)
	} else {
		SleepIntervalMS = DefaultSleepIntervalMS
	}

	value, ok = conf["openmetadata_num_retries"]
	if ok {
		NumRetries = value.(int)
	} else {
		NumRetries = DefaultNumRetries
	}

	nameToDatabaseStruct := make(map[string]database_types.DatabaseType)
	nameToDatabaseStruct[MysqlLowercase] = database_types.NewMysql(logger)
	nameToDatabaseStruct[S3] = database_types.NewS3(vaultConf, logger)

	s := &OpenMetadataApiService{Endpoint: conf["openmetadata_endpoint"].(string),
		SleepIntervalMS:      SleepIntervalMS,
		NumRetries:           NumRetries,
		NameToDatabaseStruct: nameToDatabaseStruct,
		logger:               logger,
		NumRenameRetries:     DefaultNumRenameRetries}

	s.initialized = s.prepareOpenMetadataForFybrik()

	return s
}

func (s *OpenMetadataApiService) getOpenMetadataClient() *client.APIClient {
	conf := client.Configuration{Servers: client.ServerConfigurations{
		client.ServerConfiguration{
			URL:         s.Endpoint,
			Description: "Endpoint URL",
		},
	},
	}
	return client.NewAPIClient(&conf)
}

// traverse database services looking for a service with identical configuration
func (s *OpenMetadataApiService) findService(ctx context.Context,
	c *client.APIClient,
	dt database_types.DatabaseType,
	connectionProperties map[string]interface{}) (string, string, bool) {
	connectionType := dt.OMTypeName()

	serviceList, r, err := c.DatabaseServiceApi.ListDatabaseServices(ctx).Execute()
	if err != nil {
		s.logger.Warn().Msg("Could not list database services. Let us assume that we need to create a new service")
		return "", "", false
	}
	defer r.Body.Close()

	// traverse existing services
	for i := range serviceList.Data {
		found := true

		// first let us compare connection types (e.g. s3, mysql)
		if connectionType != serviceList.Data[i].ServiceType {
			found = false
		} else if !dt.EquivalentServiceConfigurations(connectionProperties, serviceList.Data[i].Connection.Config) {
			found = false
		}
		if found {
			s.logger.Trace().Msg("Found an identical database service")
			return serviceList.Data[i].Id, *serviceList.Data[i].FullyQualifiedName, true
		}
	}
	s.logger.Trace().Msg("Identical database service not found")
	return "", "", false
}

func (s *OpenMetadataApiService) createDatabaseService(ctx context.Context,
	c *client.APIClient,
	createAssetRequest *models.CreateAssetRequest,
	connectionName string,
	omConfig map[string]interface{},
	omTypeName string) (string, string, error) {
	connection := client.DatabaseConnection{Config: omConfig}

	databaseServiceName := createAssetRequest.DestinationCatalogID + "-" + connectionName
	createDatabaseService := client.NewCreateDatabaseService(connection, databaseServiceName, omTypeName)

	databaseService, r, err := c.DatabaseServiceApi.CreateDatabaseService(ctx).CreateDatabaseService(*createDatabaseService).Execute()
	if err != nil {
		s.logger.Trace().Msg(fmt.Sprintf(ErrorCallingCreateDatabaseService, err))
		s.logger.Trace().Msg(fmt.Sprintf(FullHTTPResponse, r))
		s.logger.Warn().Msg(FailedToCreateDatabaseService + databaseServiceName)
		s.logger.Warn().Msg("Let us try again with a different name")

		// let's try creating the service with different names
		for i := 0; i < s.NumRenameRetries; i++ {
			newName := databaseServiceName + "-" + utils.RandSeq(RandomStringLength)
			createDatabaseService.SetName(newName)
			s.logger.Info().Msg("Trying to create a Database Service: " + newName)
			databaseService2, r1, err1 := c.DatabaseServiceApi.CreateDatabaseService(ctx).CreateDatabaseService(*createDatabaseService).Execute()
			if err1 == nil {
				s.logger.Info().Msg(SucceededInCreatingDatabaseService + newName)
				return databaseService2.Id, *databaseService2.FullyQualifiedName, nil
			}
			r1.Body.Close()
			s.logger.Trace().Msg(fmt.Sprintf(ErrorCallingCreateDatabaseService, err1))
			s.logger.Trace().Msg(fmt.Sprintf(FullHTTPResponse, r))
			s.logger.Warn().Msg(FailedToCreateDatabaseService + newName)
		}

		s.logger.Error().Msg("Failed to create Database Service. Giving up.")
		return "", "", err
	}
	defer r.Body.Close()

	s.logger.Info().Msg(SucceededInCreatingDatabaseService + databaseServiceName)
	return databaseService.Id, *databaseService.FullyQualifiedName, nil
}

func (s *OpenMetadataApiService) findAsset(ctx context.Context, c *client.APIClient, assetID string) (bool, *client.Table) {
	fields := "tags"
	include := "all"
	table, r, err := c.TablesApi.GetTableByFQN(ctx, assetID).Fields(fields).Include(include).Execute()
	if err != nil {
		s.logger.Trace().Msg(fmt.Sprintf("Error when calling `IngestionPipelinesApi.GetTableByFQN``: %v\n", err))
		s.logger.Trace().Msg(fmt.Sprintf(FullHTTPResponse, r))
		s.logger.Warn().Msg("Could not find asset: " + assetID)
	} else {
		defer r.Body.Close()
		s.logger.Info().Msg("Asset found: " + assetID)
	}
	return err == nil, table
}

func (s *OpenMetadataApiService) findLatestAsset(ctx context.Context, c *client.APIClient, assetID string) (bool, *client.Table) {
	found, table := s.findAsset(ctx, c, assetID)
	if !found {
		return false, nil
	}
	version := fmt.Sprintf("%f", *table.Version)
	s.logger.Trace().Msg("Latest version of the asset: " + version)

	table, r, err := c.TablesApi.GetSpecificDatabaseVersion1(ctx, table.Id, version).Execute()
	if err != nil {
		s.logger.Trace().Msg(fmt.Sprintf("Error when calling `TablesApi.GetSpecificDatabaseVersion1``: %v\n", err))
		s.logger.Trace().Msg(fmt.Sprintf(FullHTTPResponse, r))
		s.logger.Error().Msg("Could not retrieve latest version of the asset")
		return false, nil
	}
	defer r.Body.Close()

	s.logger.Info().Msg("Succeeded in retrieving latest version of the asset")
	return true, table
}

func (s *OpenMetadataApiService) findIngestionPipeline(ctx context.Context, c *client.APIClient,
	ingestionPipelineName string) (string, bool) {
	pipeline, r, err := c.IngestionPipelinesApi.GetSpecificIngestionPipelineByFQN(ctx, ingestionPipelineName).Execute()
	if err != nil {
		s.logger.Info().Msg("Ingestion Pipeline not found: " + ingestionPipelineName)
		return "", false
	}
	defer r.Body.Close()
	s.logger.Info().Msg("Ingestion Pipeline found: " + ingestionPipelineName)
	return *pipeline.Id, true
}

func (s *OpenMetadataApiService) createIngestionPipeline(ctx context.Context,
	c *client.APIClient,
	databaseServiceID string,
	ingestionPipelineName string) (string, error) {
	sourceConfig := client.SourceConfig{Config: map[string]interface{}{"type": "DatabaseMetadata"}}
	newCreateIngestionPipeline := *client.NewCreateIngestionPipeline(client.AirflowConfig{},
		ingestionPipelineName,
		"metadata", *client.NewEntityReference(databaseServiceID, DatabaseService),
		sourceConfig)

	ingestionPipeline, r, err :=
		c.IngestionPipelinesApi.CreateIngestionPipeline(ctx).
			CreateIngestionPipeline(newCreateIngestionPipeline).Execute()
	if err != nil {
		s.logger.Trace().Msg(fmt.Sprintf("Error when calling `IngestionPipelinesApi.CreateIngestionPipeline``: %v\n", err))
		s.logger.Trace().Msg(fmt.Sprintf(FullHTTPResponse, r))
		s.logger.Error().Msg("Failed to create Ingestion Pipeline: " + ingestionPipelineName + ForDatabaseServiceID + databaseServiceID)
		return "", err
	}
	defer r.Body.Close()
	s.logger.Info().Msg("Succeeded in creating Ingestion Pipeline: " + ingestionPipelineName +
		ForDatabaseServiceID + databaseServiceID)
	return *ingestionPipeline.Id, nil
}

func (s *OpenMetadataApiService) findOrCreateDatabase(ctx context.Context,
	c *client.APIClient,
	databaseServiceID string,
	databaseFQN string,
	databaseName string) (string, error) {
	// we begin by checking whether the database exists
	database, r, err := c.DatabasesApi.GetDatabaseByFQN(ctx, databaseFQN).Execute()
	if err == nil {
		defer r.Body.Close()
		s.logger.Trace().Msg("Database already exists: " + databaseFQN)
		return *database.Id, nil
	}

	s.logger.Trace().Msg("Database " + databaseFQN + " does not exist. Creating")
	database, r1, err := c.DatabasesApi.CreateDatabase(ctx).CreateDatabase(*client.NewCreateDatabase(databaseName,
		*client.NewEntityReference(databaseServiceID, DatabaseService))).Execute()
	if err != nil {
		s.logger.Trace().Msg(fmt.Sprintf("Error when calling `DatabasesApi.CreateDatabase``: %v\n", err))
		s.logger.Trace().Msg(fmt.Sprintf(FullHTTPResponse, r1))
		return "", err
	}
	defer r1.Body.Close()
	return *database.Id, nil
}

func (s *OpenMetadataApiService) findOrCreateDatabaseSchema(ctx context.Context,
	c *client.APIClient,
	databaseID string,
	databaseSchemaFQN string,
	databaseSchemaName string) (string, error) {
	// first let us check whether this Database Schema already exists
	databaseSchema, r, err := c.DatabaseSchemasApi.GetDBSchemaByFQN(ctx, databaseSchemaFQN).Execute()
	if err == nil {
		defer r.Body.Close()
		s.logger.Trace().Msg("Database Schema already exists: " + databaseSchemaName)
		return *databaseSchema.Id, nil
	}

	databaseSchema, r, err = c.DatabaseSchemasApi.CreateDBSchema(ctx).CreateDatabaseSchema(
		*client.NewCreateDatabaseSchema(*client.NewEntityReference(databaseID, "database"), databaseSchemaName)).Execute()
	if err != nil {
		s.logger.Trace().Msg(fmt.Sprintf("Error when calling `DatabaseSchemasApi.CreateDBSchema``: %v\n", err))
		s.logger.Trace().Msg(fmt.Sprintf(FullHTTPResponse, r))
		return "", err
	}
	defer r.Body.Close()
	return *databaseSchema.Id, nil
}

func (s *OpenMetadataApiService) createTable(ctx context.Context,
	c *client.APIClient,
	databaseSchemaID string,
	tableName string,
	columns []client.Column) (*client.Table, error) {
	createTable := client.NewCreateTable(columns,
		*client.NewEntityReference(databaseSchemaID, "databaseSchema"),
		tableName)
	table, r, err := c.TablesApi.CreateTable(ctx).CreateTable(*createTable).Execute()
	if err != nil {
		s.logger.Trace().Msg(fmt.Sprintf("Error when calling `TablesApi.CreateTable``: %v\n", err))
		s.logger.Trace().Msg(fmt.Sprintf(FullHTTPResponse, r))
		s.logger.Error().Msg("createTable failed: " + tableName)
		return nil, err
	}
	defer r.Body.Close()
	return table, nil
}

// enrichAsset is called after asset is created, or during an updateAsset request
// OM uses the JsonPatch format for updates
func (s *OpenMetadataApiService) enrichAsset(ctx context.Context, table *client.Table, c *client.APIClient,
	credentials *string, geography *string, name *string, owner *string,
	dataFormat *string,
	requestTags map[string]interface{},
	requestColumnsModels []models.ResourceColumn,
	requestColumnsAPI []api.ResourceColumn, connectionType string) error {
	var requestBody []map[string]interface{}

	customProperties := make(map[string]interface{})
	utils.UpdateCustomProperty(customProperties, table.Extension, Credentials, credentials)
	utils.UpdateCustomProperty(customProperties, table.Extension, Geography, geography)
	utils.UpdateCustomProperty(customProperties, table.Extension, Name, name)
	utils.UpdateCustomProperty(customProperties, table.Extension, Owner, owner)
	utils.UpdateCustomProperty(customProperties, table.Extension, DataFormat, dataFormat)
	utils.UpdateCustomProperty(customProperties, table.Extension, ConnectionType, &connectionType)

	propertiesUpdate := make(map[string]interface{})
	propertiesUpdate["op"] = "add"
	propertiesUpdate["path"] = "/extension"
	propertiesUpdate["value"] = customProperties
	requestBody = append(requestBody, propertiesUpdate)

	if requestTags != nil {
		var tags []client.TagLabel
		// traverse createAssetRequest.ResourceMetadata.Tags
		// use only the key, ignore the value (assume value is 'true')
		for tagFQN := range requestTags {
			tags = append(tags, getTag(ctx, c, tagFQN))
		}

		tagsUpdate := make(map[string]interface{})
		tagsUpdate[Op] = Add
		tagsUpdate[Path] = "/tags"
		tagsUpdate[Value] = tags
		requestBody = append(requestBody, tagsUpdate)
	}

	if requestColumnsModels != nil || requestColumnsAPI != nil {
		columns := table.Columns

		for _, col := range requestColumnsModels {
			if len(col.Tags) > 0 {
				columns = tagColumn(ctx, c, columns, col.Name, col.Tags)
			}
		}

		for _, col := range requestColumnsAPI {
			if len(col.Tags) > 0 {
				columns = tagColumn(ctx, c, columns, col.Name, col.Tags)
			}
		}

		columnUpdate := make(map[string]interface{})
		columnUpdate[Op] = Add
		columnUpdate[Path] = "/columns"
		columnUpdate[Value] = columns
		requestBody = append(requestBody, columnUpdate)
	}

	resp, err := c.TablesApi.PatchTable(ctx, table.Id).RequestBody(requestBody).Execute()
	if err != nil {
		s.logger.Trace().Msg(fmt.Sprintf("Error when calling `TablesApi.PatchTable``: %v\n", err))
		s.logger.Trace().Msg(fmt.Sprintf(FullHTTPResponse, resp))
		s.logger.Error().Msg("Asset Enrichment failed (using PatchTable)")
		return err
	}
	defer resp.Body.Close()

	s.logger.Info().Msg("Asset Enrichment succeeded")
	return nil
}

func (s *OpenMetadataApiService) deleteAsset(ctx context.Context, c *client.APIClient, assetID string) (int, error) {
	table, r, err := c.TablesApi.GetTableByFQN(ctx, assetID).Execute()
	if err != nil {
		s.logger.Trace().Msg(fmt.Sprintf("Error when calling `TablesApi.GetTableByFQN``: %v\n", err))
		s.logger.Trace().Msg(fmt.Sprintf(FullHTTPResponse, r))
		s.logger.Error().Msg("Asset deletion failed -- asset not found")
		return http.StatusNotFound, err
	}
	defer r.Body.Close()

	s.logger.Trace().Msg("deleteAsset -- asset found")
	r, err = c.TablesApi.DeleteTable(ctx, table.Id).Execute()
	if err != nil {
		s.logger.Trace().Msg(fmt.Sprintf("Error when calling `TablesApi.DeleteTable``: %v\n", err))
		s.logger.Trace().Msg(fmt.Sprintf(FullHTTPResponse, r))
		s.logger.Error().Msg("Asset deletion failed")
		return http.StatusBadRequest, err
	}
	defer r.Body.Close()

	s.logger.Error().Msg("Asset deletion successful")
	return http.StatusOK, nil
}

// populate the values in a GetAssetResponse structure to include everything:
// credentials, name, owner, geography, dataFormat, connection informations,
// tags, and columns
func (s *OpenMetadataApiService) constructAssetResponse(ctx context.Context,
	c *client.APIClient,
	table *client.Table) (*models.GetAssetResponse, error) {
	// Let's begin by finding the Database Service.
	// We need it for the connection information.
	respService, r, err := c.DatabaseServiceApi.GetDatabaseServiceByID(ctx, table.Service.Id).Execute()
	if err != nil {
		s.logger.Trace().Msg(fmt.Sprintf("Error when calling `ServicesApi.GetDatabaseServiceByID``: %v\n", err))
		s.logger.Error().Msg("Could not find Database Service: " + table.Service.Id)
		s.logger.Error().Msg("Therefore, unable to get connection information for asset: " + *table.FullyQualifiedName)
		return nil, err
	}
	defer r.Body.Close()

	ret := &models.GetAssetResponse{}
	customProperties := table.GetExtension()

	credentials := customProperties[Credentials]
	if credentials != nil {
		ret.Credentials = credentials.(string)
	}
	name := customProperties[Name]
	if name != nil {
		nameStr := name.(string)
		ret.ResourceMetadata.Name = &nameStr
	}

	owner := customProperties[Owner]
	if owner != nil {
		ownerStr := owner.(string)
		ret.ResourceMetadata.Owner = &ownerStr
	}

	geography := customProperties[Geography]
	if geography != nil {
		geographyStr := geography.(string)
		ret.ResourceMetadata.Geography = &geographyStr
	}

	dataFormat := customProperties[DataFormat]
	if dataFormat != nil {
		dataFormatStr := dataFormat.(string)
		ret.Details.DataFormat = &dataFormatStr
	}

	connectionType := customProperties[ConnectionType].(string)

	dt, found := s.NameToDatabaseStruct[connectionType]
	if !found {
		return nil, errors.New("Unrecognized connection type: " + connectionType)
	}

	config := dt.TranslateOpenMetadataConfigToFybrikConfig(table.Name, ret.Credentials,
		respService.Connection.GetConfig())

	additionalProperties := make(map[string]interface{})
	ret.Details.Connection.Name = connectionType
	additionalProperties[connectionType] = config
	ret.Details.Connection.AdditionalProperties = additionalProperties

	for i := range table.Columns {
		if len(table.Columns[i].Tags) > 0 {
			tags := make(map[string]interface{})
			for _, t := range table.Columns[i].Tags {
				tags[utils.StripTag(t.TagFQN)] = "true"
			}
			ret.ResourceMetadata.Columns = append(ret.ResourceMetadata.Columns, models.ResourceColumn{Name: table.Columns[i].Name, Tags: tags})
		} else {
			ret.ResourceMetadata.Columns = append(ret.ResourceMetadata.Columns, models.ResourceColumn{Name: table.Columns[i].Name})
		}
	}

	if len(table.Tags) > 0 {
		tags := make(map[string]interface{})
		for _, s := range table.Tags {
			tags[utils.StripTag(s.TagFQN)] = "true"
		}
		ret.ResourceMetadata.Tags = tags
	}

	s.logger.Info().Msg("Succesfully constructed asset response")
	return ret, nil
}
