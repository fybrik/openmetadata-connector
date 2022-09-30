// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package openapiconnectorcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/rs/zerolog"

	client "fybrik.io/openmetadata-connector/datacatalog-go-client"
	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
	api "fybrik.io/openmetadata-connector/datacatalog-go/go"
	dbtypes "fybrik.io/openmetadata-connector/pkg/database-types"
	"fybrik.io/openmetadata-connector/pkg/utils"
)

func getTag(ctx context.Context, c *client.APIClient, tagFQN string) client.TagLabel {
	if strings.Count(tagFQN, ".") == 0 {
		// Since this is not a 'category.primary' or 'category.primary.secondary' format,
		// we will translate it to 'Fybrik.tagFQN'. We try to create it
		// (whether it exists or not)
		createTag := *client.NewCreateTag(tagFQN, tagFQN)
		_, r, err := c.TagsApi.CreatePrimaryTag(ctx, Fybrik).CreateTag(createTag).Execute()
		if err != nil {
			logger.Trace().Err(err).Msg("could not create primary tag. it may already exist")
		} else {
			defer r.Body.Close()
		}
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

// transform variable `tags` into an array.
// traverse `tags` and create OM tags in the `categoryName` category
func (s *OpenMetadataAPIService) addTags(ctx context.Context, c *client.APIClient,
	categoryName string, tags interface{}) {
	tagsArr, ok := utils.InterfaceToArray(tags, s.logger)
	if !ok {
		s.logger.Warn().Msg("Malformed tag information")
		return
	}
	for i := range tagsArr {
		func() {
			tagMap, ok := tagsArr[i].(map[interface{}]interface{})
			if ok {
				name, ok1 := tagMap[Name]
				descriptionStr := ""
				description := tagMap[Description]
				if description != nil {
					descriptionStr = description.(string)
				}
				if ok1 {
					_, r, err := c.TagsApi.CreatePrimaryTag(ctx, categoryName).
						CreateTag(*client.NewCreateTag(descriptionStr, name.(string))).Execute()
					if err != nil {
						s.logger.Trace().Msg("Failed to create Tag. Maybe it already exists.")
					} else {
						r.Body.Close()
						return
					}
				} else {
					s.logger.Warn().Msg(fmt.Sprintf("malformed tag information. cannot cast %T to map[interface{}]interface{}", tagsArr[i]))
				}
			}
		}()
	}
}

func (s *OpenMetadataAPIService) PrepareOpenMetadataForFybrik() bool { //nolint
	ctx := context.Background()
	c := s.getOpenMetadataClient()

	// traverse tag categories. create categories as needed.
	// within each tag category, create the specified tags
	if tagCategories, ok := s.customization[TagCategories]; ok {
		tagCategoriesArr, ok := tagCategories.([]interface{})
		if ok {
			for i := range tagCategoriesArr {
				func() {
					tagCategory := tagCategoriesArr[i]
					tagCategoryMap, ok := tagCategory.(map[interface{}]interface{})
					if ok {
						name, ok1 := tagCategoryMap[Name]
						descriptionStr := ""
						description := tagCategoryMap[Description]
						if description != nil {
							descriptionStr = description.(string)
						}
						if ok1 {
							// Create Tag Category for Fybrik
							_, r, err := c.TagsApi.CreateTagCategory(ctx).CreateTagCategory(*client.NewCreateTagCategory("Classification",
								descriptionStr, name.(string))).Execute()
							if err != nil {
								s.logger.Trace().Msg("Failed to create the Tag category. Maybe it already exists.")
							} else {
								defer r.Body.Close()
							}
						}
						if tags, ok := tagCategoryMap[Tags]; ok {
							s.addTags(ctx, c, name.(string), tags)
						}
					} else {
						s.logger.Warn().Msg(fmt.Sprintf("malformed tag category. cannot cast %T to map[interface{}]interface{}", tagCategory))
					}
				}()
			}
		}
	}

	typeNameToID := make(map[string]string)

	typeList, r, err := c.MetadataApi.ListTypes(ctx).Limit(TypeListLengthLimit).Execute()
	if err != nil {
		s.logger.Warn().Msg(ErrorInPrepareOpenMetadataForFybrik + "Most likely OpenMetadata is not up yet")
		return false
	}
	defer r.Body.Close()

	for i := range typeList.Data {
		typeNameToID[*typeList.Data[i].FullyQualifiedName] = *typeList.Data[i].Id
	}

	var tableID string
	tableID, ok := typeNameToID[Table]
	if !ok {
		s.logger.Error().Msg(ErrorInPrepareOpenMetadataForFybrik + "Failed to find the ID for entity 'table'")
		return false
	}

	// traverse the table custom properties, and create each
	tableProperties, ok := s.customization[TableProperties]
	if ok {
		tablePropertiesArr, ok := tableProperties.([]interface{})
		if ok {
			for i := range tablePropertiesArr {
				func() {
					propertyMap, ok := tablePropertiesArr[i].(map[interface{}]interface{})
					if ok {
						typeName, ok1 := propertyMap[Type]
						// default property type is "string"
						if !ok1 {
							typeName = String
						}

						name, ok2 := propertyMap[Name]
						descriptionStr := ""
						description := propertyMap[Description]
						if description != nil {
							descriptionStr = description.(string)
						}

						if ok2 {
							typeID, ok4 := typeNameToID[typeName.(string)]
							if !ok4 {
								s.logger.Warn().Msg("unrecognized property type: " + typeName.(string))
							} else {
								r, err := c.MetadataApi.AddProperty(ctx, tableID).CustomProperty(*client.NewCustomProperty(
									descriptionStr, name.(string),
									*client.NewEntityReference(typeID, String))).Execute()
								if err != nil {
									defer r.Body.Close()
								}
							}
						} else {
							s.logger.Warn().Msg("missing fields for table property")
						}
					} else {
						s.logger.Warn().Msg(fmt.Sprintf("malformed table properties. cannot cast %T to map[interface{}]interface{}", tablePropertiesArr[i]))
					}
				}()
			}
		}
	}

	return true
}

// NewOpenMetadataApiService creates a new api service.
// It is initialized base on the configuration
func NewOpenMetadataAPIService(conf map[string]interface{}, customization map[string]interface{},
	logger *zerolog.Logger) OpenMetadataAPIServicer {
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

	nameToDatabaseStruct := make(map[string]dbtypes.DatabaseType)
	nameToDatabaseStruct[MysqlLowercase] = dbtypes.NewMysql(logger)
	nameToDatabaseStruct[S3] = dbtypes.NewS3(vaultConf, logger)

	s := &OpenMetadataAPIService{Endpoint: conf["openmetadata_endpoint"].(string),
		SleepIntervalMS:      SleepIntervalMS,
		NumRetries:           NumRetries,
		NameToDatabaseStruct: nameToDatabaseStruct,
		logger:               logger,
		NumRenameRetries:     DefaultNumRenameRetries,
		customization:        customization}

	s.initialized = s.PrepareOpenMetadataForFybrik()

	return s
}

func (s *OpenMetadataAPIService) getOpenMetadataClient() *client.APIClient {
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
func (s *OpenMetadataAPIService) findService(ctx context.Context,
	c *client.APIClient,
	dt dbtypes.DatabaseType,
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

func (s *OpenMetadataAPIService) createDatabaseService(ctx context.Context,
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
		s.logger.Warn().Msg(FailedToCreateDatabaseService + databaseServiceName + ". Let us try again with a different name")

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
			s.logger.Warn().Msg(FailedToCreateDatabaseService + newName)
		}

		s.logger.Error().Msg("Failed to create Database Service. Giving up.")
		return "", "", err
	}
	defer r.Body.Close()

	s.logger.Info().Msg(SucceededInCreatingDatabaseService + databaseServiceName)
	return databaseService.Id, *databaseService.FullyQualifiedName, nil
}

func (s *OpenMetadataAPIService) findAsset(ctx context.Context, c *client.APIClient, assetID string) (bool, *client.Table) {
	fields := Tags
	include := "all"
	table, r, err := c.TablesApi.GetTableByFQN(ctx, assetID).Fields(fields).Include(include).Execute()
	if err != nil {
		s.logger.Warn().Msg("Could not find asset: " + assetID)
	} else {
		defer r.Body.Close()
		s.logger.Info().Msg("Asset found: " + assetID)
	}
	return err == nil, table
}

func (s *OpenMetadataAPIService) findLatestAsset(ctx context.Context, c *client.APIClient, assetID string) (bool, *client.Table) {
	found, table := s.findAsset(ctx, c, assetID)
	if !found {
		return false, nil
	}
	version := fmt.Sprintf("%f", *table.Version)
	s.logger.Trace().Msg("Latest version of the asset: " + version)

	table, r, err := c.TablesApi.GetSpecificDatabaseVersion1(ctx, table.Id, version).Execute()
	if err != nil {
		s.logger.Error().Msg("Could not retrieve latest version of the asset")
		return false, nil
	}
	defer r.Body.Close()

	s.logger.Info().Msg("Succeeded in retrieving latest version of the asset")
	return true, table
}

func (s *OpenMetadataAPIService) findIngestionPipeline(ctx context.Context, c *client.APIClient,
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

func (s *OpenMetadataAPIService) createIngestionPipeline(ctx context.Context,
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
		s.logger.Error().Msg("Failed to create Ingestion Pipeline: " + ingestionPipelineName + ForDatabaseServiceID + databaseServiceID)
		return "", err
	}
	defer r.Body.Close()
	s.logger.Info().Msg("Succeeded in creating Ingestion Pipeline: " + ingestionPipelineName +
		ForDatabaseServiceID + databaseServiceID)
	return *ingestionPipeline.Id, nil
}

func (s *OpenMetadataAPIService) findOrCreateDatabase(ctx context.Context,
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
		return "", err
	}
	defer r1.Body.Close()
	return *database.Id, nil
}

func (s *OpenMetadataAPIService) findOrCreateDatabaseSchema(ctx context.Context,
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
		return "", err
	}
	defer r.Body.Close()
	return *databaseSchema.Id, nil
}

func (s *OpenMetadataAPIService) createTable(ctx context.Context,
	c *client.APIClient,
	databaseSchemaID string,
	tableName string,
	columns []client.Column) (*client.Table, error) {
	createTable := client.NewCreateTable(columns,
		*client.NewEntityReference(databaseSchemaID, "databaseSchema"),
		tableName)
	table, r, err := c.TablesApi.CreateTable(ctx).CreateTable(*createTable).Execute()
	if err != nil {
		s.logger.Error().Msg("createTable failed: " + tableName)
		return nil, err
	}
	defer r.Body.Close()
	return table, nil
}

// enrichAsset is called after asset is created, or during an updateAsset request
// OM uses the JsonPatch format for updates
func (s *OpenMetadataAPIService) enrichAsset(ctx context.Context, table *client.Table, c *client.APIClient,
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
	propertiesUpdate["path"] = ExtensionPath
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
		tagsUpdate[Path] = TagsPath
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
		columnUpdate[Path] = ColumnsPath
		columnUpdate[Value] = columns
		requestBody = append(requestBody, columnUpdate)
	}

	resp, err := c.TablesApi.PatchTable(ctx, table.Id).RequestBody(requestBody).Execute()
	if err != nil {
		s.logger.Error().Msg("Asset Enrichment failed (using PatchTable)")
		return err
	}
	defer resp.Body.Close()

	s.logger.Info().Msg("Asset Enrichment succeeded")
	return nil
}

func (s *OpenMetadataAPIService) deleteAsset(ctx context.Context, c *client.APIClient, assetID string) (int, error) {
	table, r, err := c.TablesApi.GetTableByFQN(ctx, assetID).Execute()
	if err != nil {
		s.logger.Error().Msg("Asset deletion failed -- asset not found")
		return http.StatusNotFound, err
	}
	defer r.Body.Close()

	s.logger.Trace().Msg("deleteAsset -- asset found")
	r, err = c.TablesApi.DeleteTable(ctx, table.Id).Execute()
	if err != nil {
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
func (s *OpenMetadataAPIService) constructAssetResponse(ctx context.Context, //nolint
	c *client.APIClient,
	table *client.Table) (*models.GetAssetResponse, error) {
	// Let's begin by finding the Database Service.
	// We need it for the connection information.
	respService, r, err := c.DatabaseServiceApi.GetDatabaseServiceByID(ctx, table.Service.Id).Execute()
	if err != nil {
		s.logger.Error().Msg("Could not find Database Service: " + table.Service.Id +
			". Therefore, unable to get connection information for asset: " + *table.FullyQualifiedName)
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

	connectionType := customProperties[ConnectionType]

	if connectionType == nil {
		return nil, errors.New(ConnectionType + " value missing from table custom properties")
	}
	connectionTypeStr := connectionType.(string)
	dt, found := s.NameToDatabaseStruct[connectionTypeStr]
	if !found {
		return nil, errors.New("Unrecognized connection type: " + connectionTypeStr)
	}

	config, err := dt.TranslateOpenMetadataConfigToFybrikConfig(table.Name, ret.Credentials,
		respService.Connection.GetConfig())
	if err != nil {
		s.logger.Error().Err(err).Msg("error in translating openmetadata config to fybrik format")
		return nil, err
	}

	additionalProperties := make(map[string]interface{})
	ret.Details.Connection.Name = connectionTypeStr
	additionalProperties[connectionTypeStr] = config
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
