// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package openapiconnectorcore

import (
	"context"

	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
	api "fybrik.io/openmetadata-connector/datacatalog-go/go"
)

type OpenMetadataAPIServicer interface {
	// CHANGE-FROM-GENERATED-CODE: replaced api.CreateAssetRequest with models.CreateAssetRequest
	CreateAsset(context.Context, string, *models.CreateAssetRequest) (api.ImplResponse, error)

	DeleteAsset(context.Context, string, *api.DeleteAssetRequest) (api.ImplResponse, error)
	GetAssetInfo(context.Context, string, *api.GetAssetRequest) (api.ImplResponse, error)
	UpdateAsset(context.Context, string, *api.UpdateAssetRequest) (api.ImplResponse, error)
}
