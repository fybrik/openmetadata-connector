// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package openapiconnectorcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rs/zerolog"

	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
	api "fybrik.io/openmetadata-connector/datacatalog-go/go"
	dbtypes "fybrik.io/openmetadata-connector/pkg/database-types"
	"fybrik.io/openmetadata-connector/pkg/utils"
)

type OpenMetadataAPIService struct {
	Endpoint             string
	SleepIntervalMS      int
	NumRetries           int
	NameToDatabaseStruct map[string]dbtypes.DatabaseType
	logger               *zerolog.Logger
	NumRenameRetries     int
	initialized          bool
	customization        map[string]interface{}
}

// CreateAsset - This REST API writes data asset information to the data catalog configured in fybrik
func (s *OpenMetadataAPIService) CreateAsset(ctx context.Context, //nolint
	xRequestDatacatalogWriteCred string,
	createAssetRequest *models.CreateAssetRequest) (api.ImplResponse, error) {
	if !s.initialized {
		s.initialized = s.PrepareOpenMetadataForFybrik()
	}

	connectionType := createAssetRequest.Details.Connection.Name

	// check whether connectionType is one of the connection types supported by the OM connector
	dt, found := s.NameToDatabaseStruct[connectionType]
	if !found {
		s.logger.Error().Msg(fmt.Sprintf(ConnectionTypeNotSupported, connectionType))
		return api.Response(http.StatusBadRequest, nil),
			fmt.Errorf(ConnectionTypeNotSupported, connectionType)
	}

	c := s.getOpenMetadataClient()

	var databaseServiceID string
	var databaseServiceName string
	var err error

	// Let us begin by checking whether the database service already exists.
	// step 1: Translate the fybrik connection information to the OM connection information.
	//         This configuration information will later be used to create an OM connection
	//         (if it does not already exist)
	config, ok := utils.InterfaceToMap(createAssetRequest.Details.GetConnection().AdditionalProperties[connectionType], s.logger)
	if !ok {
		s.logger.Error().Msg(FailedToCovert)
		return api.Response(http.StatusBadRequest, nil), errors.New(FailedToCovert)
	}
	omConfig := dt.TranslateFybrikConfigToOpenMetadataConfig(config, createAssetRequest.Credentials)
	// step 2: compare the transformed connection information to that of all existing services
	databaseServiceID, databaseServiceName, found = s.findService(ctx, c, dt, omConfig)

	if !found {
		// If does not exist, let us create database service
		databaseServiceID, databaseServiceName, err =
			s.createDatabaseService(ctx, c, createAssetRequest, connectionType, omConfig, dt.OMTypeName())
		if err != nil {
			s.logger.Error().Msg("unable to create Database Service for " + dt.OMTypeName() + " connection")
			return api.Response(http.StatusBadRequest, nil), err
		}
	}

	// now that we know the of the database service, we can determine the asset name in OpenMetadata
	assetID, err := dt.TableFQN(databaseServiceName, createAssetRequest)
	if err != nil {
		s.logger.Error().Err(err).Msg("cannot determine table FQN")
		return api.Response(http.StatusBadRequest, nil), err
	}

	// Let's check whether OM already has this asset
	found, _ = s.findAsset(ctx, c, assetID)
	if found {
		s.logger.Error().Msg("Could not create asset, as asset already exists")
		return api.Response(http.StatusBadRequest, nil), errors.New("asset already exists")
	}

	// Asset not discovered yet
	// Let's check whether there is an ingestion pipeline we can trigger
	ingestionPipelineName := "pipeline-" + createAssetRequest.DestinationCatalogID + "." + *createAssetRequest.DestinationAssetID
	ingestionPipelineNameFull := utils.AppendStrings(databaseServiceName, ingestionPipelineName)

	// var ingestionPipelineID string nolint
	// ingestionPipelineID, found = s.findIngestionPipeline(ctx, c, ingestionPipelineNameFull)
	_, found = s.findIngestionPipeline(ctx, c, ingestionPipelineNameFull)

	if !found {
		// Let us create an ingestion pipeline
		s.logger.Info().Msg("Ingestion Pipeline not found. Creating.")
		_, _ = s.createIngestionPipeline(ctx, c, databaseServiceID, ingestionPipelineName)
	}

	databaseID, err := s.findOrCreateDatabase(ctx, c, databaseServiceID,
		dt.DatabaseFQN(databaseServiceName, createAssetRequest),
		dt.DatabaseName(createAssetRequest))
	if err != nil {
		return api.Response(http.StatusBadRequest, nil), err
	}

	databaseSchemaID, _ := s.findOrCreateDatabaseSchema(ctx, c, databaseID,
		dt.DatabaseSchemaFQN(databaseServiceName, createAssetRequest),
		dt.DatabaseSchemaName(createAssetRequest))

	columns := utils.ExtractColumns(createAssetRequest.ResourceMetadata.Columns)

	tableName, err := dt.TableName(createAssetRequest)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to determine table name")
		return api.Response(http.StatusBadRequest, nil), err
	}
	table, err := s.createTable(ctx, c, databaseSchemaID, tableName, columns)
	if err != nil {
		return api.Response(http.StatusBadRequest, nil), err
	}

	s.logger.Info().Msg("Enriching asset with additional information (e.g. tags)")
	// Now that OM is aware of the asset, we need to enrich it --
	// add tags to asset and to columns, and populate the custom properties
	err = s.enrichAsset(ctx, table, c,
		createAssetRequest.Credentials, createAssetRequest.ResourceMetadata.Geography,
		createAssetRequest.ResourceMetadata.Name, createAssetRequest.ResourceMetadata.Owner,
		createAssetRequest.Details.DataFormat,
		createAssetRequest.ResourceMetadata.Tags,
		createAssetRequest.ResourceMetadata.Columns, nil, connectionType)

	if err != nil {
		s.logger.Error().Msg("Asset enrichment failed")
		return api.Response(http.StatusBadRequest, nil), err
	}

	s.logger.Info().Msg("Asset creation and enrichment successful")

	return api.Response(http.StatusCreated, api.CreateAssetResponse{AssetID: assetID}), nil
}

// DeleteAsset - This REST API deletes data asset
func (s *OpenMetadataAPIService) DeleteAsset(ctx context.Context, xRequestDatacatalogCred string,
	deleteAssetRequest *api.DeleteAssetRequest) (api.ImplResponse, error) {
	if !s.initialized {
		s.initialized = s.PrepareOpenMetadataForFybrik()
	}

	c := s.getOpenMetadataClient()
	errorCode, err := s.deleteAsset(ctx, c, deleteAssetRequest.AssetID)

	if err != nil {
		s.logger.Info().Msg("Asset deletion failed")
		return api.Response(errorCode, nil), err
	}

	s.logger.Info().Msg("Asset deletion successful")
	return api.Response(http.StatusOK, api.DeleteAssetResponse{}), nil
}

// GetAssetInfo - This REST API gets data asset information from the data catalog configured in
// fybrik for the data sets indicated in FybrikApplication yaml
func (s *OpenMetadataAPIService) GetAssetInfo(ctx context.Context, xRequestDatacatalogCred string,
	getAssetRequest *api.GetAssetRequest) (api.ImplResponse, error) {
	if !s.initialized {
		s.initialized = s.PrepareOpenMetadataForFybrik()
	}

	c := s.getOpenMetadataClient()

	assetID := getAssetRequest.AssetID

	found, table := s.findLatestAsset(ctx, c, assetID)
	if !found {
		s.logger.Error().Msg("Asset not found")
		return api.Response(http.StatusNotFound, nil), errors.New("asset not found")
	}

	assetResponse, err := s.constructAssetResponse(ctx, c, table)
	if err != nil {
		s.logger.Error().Msg("Construction of Asset Reponse failed")
		return api.Response(http.StatusBadRequest, nil), err
	}

	s.logger.Info().Msg("GetAssetInfo successful")
	return api.Response(http.StatusOK, assetResponse), nil
}

// UpdateAsset - This REST API updates data asset information in the data catalog configured in fybrik
func (s *OpenMetadataAPIService) UpdateAsset(ctx context.Context, xRequestDatacatalogUpdateCred string,
	updateAssetRequest *api.UpdateAssetRequest) (api.ImplResponse, error) {
	if !s.initialized {
		s.initialized = s.PrepareOpenMetadataForFybrik()
	}

	c := s.getOpenMetadataClient()
	assetID := updateAssetRequest.AssetID

	found, table := s.findLatestAsset(ctx, c, assetID)
	if !found {
		s.logger.Error().Msg(AssetNotFound)
		return api.Response(http.StatusNotFound, nil), errors.New(AssetNotFound)
	}

	err := s.enrichAsset(ctx, table, c, nil, nil, &updateAssetRequest.Name, &updateAssetRequest.Owner, nil,
		updateAssetRequest.Tags, nil, updateAssetRequest.Columns, "")
	if err != nil {
		s.logger.Error().Msg("asset enrichment failed")
		return api.Response(http.StatusBadRequest, nil), err
	}

	s.logger.Info().Msg("UpdateAsset successful")
	return api.Response(http.StatusOK, api.UpdateAssetResponse{Status: "Asset update operation successful"}), nil
}
