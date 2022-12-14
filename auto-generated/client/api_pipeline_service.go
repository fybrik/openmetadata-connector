/*
OpenMetadata Apis

# Overview  OpenMetadata supports REST APIs for getting metadata in and out of metadata store. The API resources are grouped under following categories: - **Data assets** - includes resources for data entities, such as `databases`, `tables`, and `topics`. Resources for data assets created from data, such as `dashboards`, `reports`, `metrics`, and `ML Features` are part of this collection. `pipelines`, `notebooks`, etc. that are used for creating data assets are also available as resources as of this collection. - **Teams and Users** - includes `users`, `teams`, a special type of user called `bots` that performs many automated tasks such as ingestion. - **Services** - are services that OpenMetadata integrates with. Currently `databaseService` is the only service under this collection that represents data sources. In the future, services related to Dashboards, Reports, ETL pipelines will be added under this collection. - **Glossary** - OpenMetadata supports hierarchical tags that can be used to build business vocabulary to describe and classify data available under `tags` resource.

API version: 0.11.0
Contact: openmetadata-dev@googlegroups.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// PipelineServiceApiService PipelineServiceApi service
type PipelineServiceApiService service

type ApiCreateOrUpdatePipelineServiceRequest struct {
	ctx                   context.Context
	ApiService            *PipelineServiceApiService
	createPipelineService *CreatePipelineService
}

func (r ApiCreateOrUpdatePipelineServiceRequest) CreatePipelineService(createPipelineService CreatePipelineService) ApiCreateOrUpdatePipelineServiceRequest {
	r.createPipelineService = &createPipelineService
	return r
}

func (r ApiCreateOrUpdatePipelineServiceRequest) Execute() (*PipelineService, *http.Response, error) {
	return r.ApiService.CreateOrUpdatePipelineServiceExecute(r)
}

/*
CreateOrUpdatePipelineService Update pipeline service

Create a new pipeline service or update an existing pipeline service identified by `id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateOrUpdatePipelineServiceRequest
*/
func (a *PipelineServiceApiService) CreateOrUpdatePipelineService(ctx context.Context) ApiCreateOrUpdatePipelineServiceRequest {
	return ApiCreateOrUpdatePipelineServiceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return PipelineService
func (a *PipelineServiceApiService) CreateOrUpdatePipelineServiceExecute(r ApiCreateOrUpdatePipelineServiceRequest) (*PipelineService, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PipelineService
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelineServiceApiService.CreateOrUpdatePipelineService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/pipelineServices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createPipelineService
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreatePipelineServiceRequest struct {
	ctx                   context.Context
	ApiService            *PipelineServiceApiService
	createPipelineService *CreatePipelineService
}

func (r ApiCreatePipelineServiceRequest) CreatePipelineService(createPipelineService CreatePipelineService) ApiCreatePipelineServiceRequest {
	r.createPipelineService = &createPipelineService
	return r
}

func (r ApiCreatePipelineServiceRequest) Execute() (*PipelineService, *http.Response, error) {
	return r.ApiService.CreatePipelineServiceExecute(r)
}

/*
CreatePipelineService Create a pipeline service

Create a new pipeline service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreatePipelineServiceRequest
*/
func (a *PipelineServiceApiService) CreatePipelineService(ctx context.Context) ApiCreatePipelineServiceRequest {
	return ApiCreatePipelineServiceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return PipelineService
func (a *PipelineServiceApiService) CreatePipelineServiceExecute(r ApiCreatePipelineServiceRequest) (*PipelineService, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PipelineService
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelineServiceApiService.CreatePipelineService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/pipelineServices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createPipelineService
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeletePipelineServiceRequest struct {
	ctx        context.Context
	ApiService *PipelineServiceApiService
	id         string
	recursive  *bool
	hardDelete *bool
}

// Recursively delete this entity and it&#39;s children. (Default &#x60;false&#x60;)
func (r ApiDeletePipelineServiceRequest) Recursive(recursive bool) ApiDeletePipelineServiceRequest {
	r.recursive = &recursive
	return r
}

// Hard delete the entity. (Default &#x3D; &#x60;false&#x60;)
func (r ApiDeletePipelineServiceRequest) HardDelete(hardDelete bool) ApiDeletePipelineServiceRequest {
	r.hardDelete = &hardDelete
	return r
}

func (r ApiDeletePipelineServiceRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePipelineServiceExecute(r)
}

/*
DeletePipelineService Delete a pipeline service

Delete a pipeline services. If pipelines (and tasks) belong to the service, it can't be deleted.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Id of the pipeline service
 @return ApiDeletePipelineServiceRequest
*/
func (a *PipelineServiceApiService) DeletePipelineService(ctx context.Context, id string) ApiDeletePipelineServiceRequest {
	return ApiDeletePipelineServiceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *PipelineServiceApiService) DeletePipelineServiceExecute(r ApiDeletePipelineServiceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelineServiceApiService.DeletePipelineService")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/pipelineServices/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.recursive != nil {
		localVarQueryParams.Add("recursive", parameterToString(*r.recursive, ""))
	}
	if r.hardDelete != nil {
		localVarQueryParams.Add("hardDelete", parameterToString(*r.hardDelete, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetPipelineServiceByFQNRequest struct {
	ctx        context.Context
	ApiService *PipelineServiceApiService
	name       string
	fields     *string
	include    *string
}

// Fields requested in the returned resource
func (r ApiGetPipelineServiceByFQNRequest) Fields(fields string) ApiGetPipelineServiceByFQNRequest {
	r.fields = &fields
	return r
}

// Include all, deleted, or non-deleted entities.
func (r ApiGetPipelineServiceByFQNRequest) Include(include string) ApiGetPipelineServiceByFQNRequest {
	r.include = &include
	return r
}

func (r ApiGetPipelineServiceByFQNRequest) Execute() (*PipelineService, *http.Response, error) {
	return r.ApiService.GetPipelineServiceByFQNExecute(r)
}

/*
GetPipelineServiceByFQN Get pipeline service by name

Get a pipeline service by the service `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @return ApiGetPipelineServiceByFQNRequest
*/
func (a *PipelineServiceApiService) GetPipelineServiceByFQN(ctx context.Context, name string) ApiGetPipelineServiceByFQNRequest {
	return ApiGetPipelineServiceByFQNRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//  @return PipelineService
func (a *PipelineServiceApiService) GetPipelineServiceByFQNExecute(r ApiGetPipelineServiceByFQNRequest) (*PipelineService, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PipelineService
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelineServiceApiService.GetPipelineServiceByFQN")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/pipelineServices/name/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterToString(r.name, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPipelineServiceByIDRequest struct {
	ctx        context.Context
	ApiService *PipelineServiceApiService
	id         string
	fields     *string
	include    *string
}

// Fields requested in the returned resource
func (r ApiGetPipelineServiceByIDRequest) Fields(fields string) ApiGetPipelineServiceByIDRequest {
	r.fields = &fields
	return r
}

// Include all, deleted, or non-deleted entities.
func (r ApiGetPipelineServiceByIDRequest) Include(include string) ApiGetPipelineServiceByIDRequest {
	r.include = &include
	return r
}

func (r ApiGetPipelineServiceByIDRequest) Execute() (*PipelineService, *http.Response, error) {
	return r.ApiService.GetPipelineServiceByIDExecute(r)
}

/*
GetPipelineServiceByID Get a pipeline service

Get a pipeline service by `id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetPipelineServiceByIDRequest
*/
func (a *PipelineServiceApiService) GetPipelineServiceByID(ctx context.Context, id string) ApiGetPipelineServiceByIDRequest {
	return ApiGetPipelineServiceByIDRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return PipelineService
func (a *PipelineServiceApiService) GetPipelineServiceByIDExecute(r ApiGetPipelineServiceByIDRequest) (*PipelineService, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PipelineService
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelineServiceApiService.GetPipelineServiceByID")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/pipelineServices/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetSpecificPipelineServiceRequest struct {
	ctx        context.Context
	ApiService *PipelineServiceApiService
	id         string
	version    string
}

func (r ApiGetSpecificPipelineServiceRequest) Execute() (*PipelineService, *http.Response, error) {
	return r.ApiService.GetSpecificPipelineServiceExecute(r)
}

/*
GetSpecificPipelineService Get a version of the pipeline service

Get a version of the pipeline service by given `id`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id pipeline service Id
 @param version pipeline service version number in the form `major`.`minor`
 @return ApiGetSpecificPipelineServiceRequest
*/
func (a *PipelineServiceApiService) GetSpecificPipelineService(ctx context.Context, id string, version string) ApiGetSpecificPipelineServiceRequest {
	return ApiGetSpecificPipelineServiceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		version:    version,
	}
}

// Execute executes the request
//  @return PipelineService
func (a *PipelineServiceApiService) GetSpecificPipelineServiceExecute(r ApiGetSpecificPipelineServiceRequest) (*PipelineService, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PipelineService
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelineServiceApiService.GetSpecificPipelineService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/pipelineServices/{id}/versions/{version}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", url.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListAllPipelineServiceVersionRequest struct {
	ctx        context.Context
	ApiService *PipelineServiceApiService
	id         string
}

func (r ApiListAllPipelineServiceVersionRequest) Execute() (*EntityHistory, *http.Response, error) {
	return r.ApiService.ListAllPipelineServiceVersionExecute(r)
}

/*
ListAllPipelineServiceVersion List pipeline service versions

Get a list of all the versions of a pipeline service identified by `id`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id pipeline service Id
 @return ApiListAllPipelineServiceVersionRequest
*/
func (a *PipelineServiceApiService) ListAllPipelineServiceVersion(ctx context.Context, id string) ApiListAllPipelineServiceVersionRequest {
	return ApiListAllPipelineServiceVersionRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return EntityHistory
func (a *PipelineServiceApiService) ListAllPipelineServiceVersionExecute(r ApiListAllPipelineServiceVersionRequest) (*EntityHistory, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EntityHistory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelineServiceApiService.ListAllPipelineServiceVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/pipelineServices/{id}/versions"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListPipelineServiceRequest struct {
	ctx        context.Context
	ApiService *PipelineServiceApiService
	fields     *string
	limit      *int32
	before     *string
	after      *string
	include    *string
}

// Fields requested in the returned resource
func (r ApiListPipelineServiceRequest) Fields(fields string) ApiListPipelineServiceRequest {
	r.fields = &fields
	return r
}

// Limit number services returned. (1 to 1000000, default 10)
func (r ApiListPipelineServiceRequest) Limit(limit int32) ApiListPipelineServiceRequest {
	r.limit = &limit
	return r
}

// Returns list of services before this cursor
func (r ApiListPipelineServiceRequest) Before(before string) ApiListPipelineServiceRequest {
	r.before = &before
	return r
}

// Returns list of services after this cursor
func (r ApiListPipelineServiceRequest) After(after string) ApiListPipelineServiceRequest {
	r.after = &after
	return r
}

// Include all, deleted, or non-deleted entities.
func (r ApiListPipelineServiceRequest) Include(include string) ApiListPipelineServiceRequest {
	r.include = &include
	return r
}

func (r ApiListPipelineServiceRequest) Execute() (*PipelineServiceList, *http.Response, error) {
	return r.ApiService.ListPipelineServiceExecute(r)
}

/*
ListPipelineService List pipeline services

Get a list of pipeline services. Use cursor-based pagination to limit the number entries in the list using `limit` and `before` or `after` query params.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListPipelineServiceRequest
*/
func (a *PipelineServiceApiService) ListPipelineService(ctx context.Context) ApiListPipelineServiceRequest {
	return ApiListPipelineServiceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return PipelineServiceList
func (a *PipelineServiceApiService) ListPipelineServiceExecute(r ApiListPipelineServiceRequest) (*PipelineServiceList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PipelineServiceList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelineServiceApiService.ListPipelineService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/pipelineServices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.before != nil {
		localVarQueryParams.Add("before", parameterToString(*r.before, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
