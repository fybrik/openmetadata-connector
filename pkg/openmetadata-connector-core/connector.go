// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package openapiconnectorcore

import (
	"encoding/json"
	"net/http"
	"strings"

	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
	api "fybrik.io/openmetadata-connector/datacatalog-go/go"
)

// DefaultAPIController binds http requests to an api service and writes the service results to the http response
type DefaultAPIController struct {
	// replace 'DefaultAPIServicer' with 'OpenMetadataAPIServicer'
	service      OpenMetadataAPIServicer
	errorHandler api.ErrorHandler
}

// DefaultApiOption for how the controller is set up.
type DefaultAPIOption func(*DefaultAPIController)

// WithDefaultApiErrorHandler inject ErrorHandler into controller
func WithDefaultAPIErrorHandler(h api.ErrorHandler) DefaultAPIOption {
	return func(c *DefaultAPIController) {
		c.errorHandler = h
	}
}

// NewDefaultApiController creates a default api controller
func NewOpenMetadataAPIController(s OpenMetadataAPIServicer, opts ...DefaultAPIOption) api.Router {
	controller := &DefaultAPIController{
		service:      s,
		errorHandler: api.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultApiController
func (c *DefaultAPIController) Routes() api.Routes {
	return api.Routes{
		api.Route{
			Name:        "CreateAsset",
			Method:      strings.ToUpper(http.MethodPost),
			Pattern:     "/createAsset",
			HandlerFunc: c.CreateAsset,
		},
		api.Route{
			Name:        "DeleteAsset",
			Method:      strings.ToUpper(http.MethodDelete),
			Pattern:     "/deleteAsset",
			HandlerFunc: c.DeleteAsset,
		},
		api.Route{
			Name:        "GetAssetInfo",
			Method:      strings.ToUpper(http.MethodPost),
			Pattern:     "/getAssetInfo",
			HandlerFunc: c.GetAssetInfo,
		},
		api.Route{
			Name:        "UpdateAsset",
			Method:      strings.ToUpper(http.MethodPatch),
			Pattern:     "/updateAsset",
			HandlerFunc: c.UpdateAsset,
		},
	}
}

// CreateAsset - This REST API writes data asset information to the data catalog configured in fybrik
func (c *DefaultAPIController) CreateAsset(w http.ResponseWriter, r *http.Request) {
	xRequestDatacatalogWriteCredParam := r.Header.Get("X-Request-Datacatalog-Write-Cred")

	// CHANGE-FROM-GENERATED-CODE: we translate the body to be a CreateAssetRequest object
	// taken from client-code models instead of from server-code models
	createAssetRequestParam := models.CreateAssetRequest{}

	d := json.NewDecoder(r.Body)

	d.DisallowUnknownFields()

	if err := d.Decode(&createAssetRequestParam); err != nil {
		c.errorHandler(w, r, &api.ParsingError{Err: err}, nil)
		return
	}

	// CHANGE-FROM-GENERATED-CODE: commented out AssertCreateAssetRequestRequired() because
	// this method does not exist for CreateAssetRequest struct defined in Fybrik code
	// if err := api.AssertCreateAssetRequestRequired(createAssetRequestParam); err != nil {
	//	c.errorHandler(w, r, err, nil)
	//	return
	//}

	result, err := c.service.CreateAsset(r.Context(), xRequestDatacatalogWriteCredParam, &createAssetRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = api.EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteAsset - This REST API deletes data asset
func (c *DefaultAPIController) DeleteAsset(w http.ResponseWriter, r *http.Request) { //nolint
	xRequestDatacatalogCredParam := r.Header.Get(XRequestDatacatalogCred)
	deleteAssetRequestParam := api.DeleteAssetRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&deleteAssetRequestParam); err != nil {
		c.errorHandler(w, r, &api.ParsingError{Err: err}, nil)
		return
	}
	if err := api.AssertDeleteAssetRequestRequired(deleteAssetRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.DeleteAsset(r.Context(), xRequestDatacatalogCredParam, &deleteAssetRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = api.EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetAssetInfo - This REST API gets data asset information from the data catalog configured in fybrik
// for the data sets indicated in FybrikApplication yaml
func (c *DefaultAPIController) GetAssetInfo(w http.ResponseWriter, r *http.Request) { //nolint
	xRequestDatacatalogCredParam := r.Header.Get(XRequestDatacatalogCred)
	getAssetRequestParam := api.GetAssetRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&getAssetRequestParam); err != nil {
		c.errorHandler(w, r, &api.ParsingError{Err: err}, nil)
		return
	}
	if err := api.AssertGetAssetRequestRequired(getAssetRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.GetAssetInfo(r.Context(), xRequestDatacatalogCredParam, &getAssetRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = api.EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateAsset - This REST API updates data asset information in the data catalog configured in fybrik
func (c *DefaultAPIController) UpdateAsset(w http.ResponseWriter, r *http.Request) {
	xRequestDatacatalogUpdateCredParam := r.Header.Get("X-Request-Datacatalog-Update-Cred")
	updateAssetRequestParam := api.UpdateAssetRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&updateAssetRequestParam); err != nil {
		c.errorHandler(w, r, &api.ParsingError{Err: err}, nil)
		return
	}
	if err := api.AssertUpdateAssetRequestRequired(updateAssetRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateAsset(r.Context(), xRequestDatacatalogUpdateCredParam, &updateAssetRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = api.EncodeJSONResponse(result.Body, &result.Code, w)
}
