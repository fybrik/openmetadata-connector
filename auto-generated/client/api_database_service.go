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

// DatabaseServiceApiService DatabaseServiceApi service
type DatabaseServiceApiService service

type ApiCreateDatabaseServiceRequest struct {
	ctx                   context.Context
	ApiService            *DatabaseServiceApiService
	createDatabaseService *CreateDatabaseService
}

func (r ApiCreateDatabaseServiceRequest) CreateDatabaseService(createDatabaseService CreateDatabaseService) ApiCreateDatabaseServiceRequest {
	r.createDatabaseService = &createDatabaseService
	return r
}

func (r ApiCreateDatabaseServiceRequest) Execute() (*DatabaseService, *http.Response, error) {
	return r.ApiService.CreateDatabaseServiceExecute(r)
}

/*
CreateDatabaseService Create database service

Create a new database service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDatabaseServiceRequest
*/
func (a *DatabaseServiceApiService) CreateDatabaseService(ctx context.Context) ApiCreateDatabaseServiceRequest {
	return ApiCreateDatabaseServiceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return DatabaseService
func (a *DatabaseServiceApiService) CreateDatabaseServiceExecute(r ApiCreateDatabaseServiceRequest) (*DatabaseService, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseService
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseServiceApiService.CreateDatabaseService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/databaseServices"

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
	localVarPostBody = r.createDatabaseService
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

type ApiCreateOrUpdateDatabaseServiceRequest struct {
	ctx                   context.Context
	ApiService            *DatabaseServiceApiService
	createDatabaseService *CreateDatabaseService
}

func (r ApiCreateOrUpdateDatabaseServiceRequest) CreateDatabaseService(createDatabaseService CreateDatabaseService) ApiCreateOrUpdateDatabaseServiceRequest {
	r.createDatabaseService = &createDatabaseService
	return r
}

func (r ApiCreateOrUpdateDatabaseServiceRequest) Execute() (*DatabaseService, *http.Response, error) {
	return r.ApiService.CreateOrUpdateDatabaseServiceExecute(r)
}

/*
CreateOrUpdateDatabaseService Update database service

Update an existing or create a new database service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateOrUpdateDatabaseServiceRequest
*/
func (a *DatabaseServiceApiService) CreateOrUpdateDatabaseService(ctx context.Context) ApiCreateOrUpdateDatabaseServiceRequest {
	return ApiCreateOrUpdateDatabaseServiceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return DatabaseService
func (a *DatabaseServiceApiService) CreateOrUpdateDatabaseServiceExecute(r ApiCreateOrUpdateDatabaseServiceRequest) (*DatabaseService, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseService
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseServiceApiService.CreateOrUpdateDatabaseService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/databaseServices"

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
	localVarPostBody = r.createDatabaseService
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

type ApiDeleteDatabaseServiceRequest struct {
	ctx        context.Context
	ApiService *DatabaseServiceApiService
	id         string
	recursive  *bool
	hardDelete *bool
}

// Recursively delete this entity and it&#39;s children. (Default &#x60;false&#x60;)
func (r ApiDeleteDatabaseServiceRequest) Recursive(recursive bool) ApiDeleteDatabaseServiceRequest {
	r.recursive = &recursive
	return r
}

// Hard delete the entity. (Default &#x3D; &#x60;false&#x60;)
func (r ApiDeleteDatabaseServiceRequest) HardDelete(hardDelete bool) ApiDeleteDatabaseServiceRequest {
	r.hardDelete = &hardDelete
	return r
}

func (r ApiDeleteDatabaseServiceRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteDatabaseServiceExecute(r)
}

/*
DeleteDatabaseService Delete a database service

Delete a database services. If databases (and tables) belong the service, it can't be deleted.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Id of the database service
 @return ApiDeleteDatabaseServiceRequest
*/
func (a *DatabaseServiceApiService) DeleteDatabaseService(ctx context.Context, id string) ApiDeleteDatabaseServiceRequest {
	return ApiDeleteDatabaseServiceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *DatabaseServiceApiService) DeleteDatabaseServiceExecute(r ApiDeleteDatabaseServiceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseServiceApiService.DeleteDatabaseService")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/databaseServices/{id}"
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

type ApiGetDatabaseServiceByFQNRequest struct {
	ctx        context.Context
	ApiService *DatabaseServiceApiService
	name       string
	fields     *string
	include    *string
}

// Fields requested in the returned resource
func (r ApiGetDatabaseServiceByFQNRequest) Fields(fields string) ApiGetDatabaseServiceByFQNRequest {
	r.fields = &fields
	return r
}

// Include all, deleted, or non-deleted entities.
func (r ApiGetDatabaseServiceByFQNRequest) Include(include string) ApiGetDatabaseServiceByFQNRequest {
	r.include = &include
	return r
}

func (r ApiGetDatabaseServiceByFQNRequest) Execute() (*DatabaseService, *http.Response, error) {
	return r.ApiService.GetDatabaseServiceByFQNExecute(r)
}

/*
GetDatabaseServiceByFQN Get database service by name

Get a database service by the service `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @return ApiGetDatabaseServiceByFQNRequest
*/
func (a *DatabaseServiceApiService) GetDatabaseServiceByFQN(ctx context.Context, name string) ApiGetDatabaseServiceByFQNRequest {
	return ApiGetDatabaseServiceByFQNRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//  @return DatabaseService
func (a *DatabaseServiceApiService) GetDatabaseServiceByFQNExecute(r ApiGetDatabaseServiceByFQNRequest) (*DatabaseService, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseService
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseServiceApiService.GetDatabaseServiceByFQN")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/databaseServices/name/{name}"
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

type ApiGetDatabaseServiceByIDRequest struct {
	ctx        context.Context
	ApiService *DatabaseServiceApiService
	id         string
	fields     *string
	include    *string
}

// Fields requested in the returned resource
func (r ApiGetDatabaseServiceByIDRequest) Fields(fields string) ApiGetDatabaseServiceByIDRequest {
	r.fields = &fields
	return r
}

// Include all, deleted, or non-deleted entities.
func (r ApiGetDatabaseServiceByIDRequest) Include(include string) ApiGetDatabaseServiceByIDRequest {
	r.include = &include
	return r
}

func (r ApiGetDatabaseServiceByIDRequest) Execute() (*DatabaseService, *http.Response, error) {
	return r.ApiService.GetDatabaseServiceByIDExecute(r)
}

/*
GetDatabaseServiceByID Get a database service

Get a database service by `id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetDatabaseServiceByIDRequest
*/
func (a *DatabaseServiceApiService) GetDatabaseServiceByID(ctx context.Context, id string) ApiGetDatabaseServiceByIDRequest {
	return ApiGetDatabaseServiceByIDRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return DatabaseService
func (a *DatabaseServiceApiService) GetDatabaseServiceByIDExecute(r ApiGetDatabaseServiceByIDRequest) (*DatabaseService, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseService
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseServiceApiService.GetDatabaseServiceByID")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/databaseServices/{id}"
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

type ApiGetSpecificDatabaseServiceVersionRequest struct {
	ctx        context.Context
	ApiService *DatabaseServiceApiService
	id         string
	version    string
}

func (r ApiGetSpecificDatabaseServiceVersionRequest) Execute() (*DatabaseService, *http.Response, error) {
	return r.ApiService.GetSpecificDatabaseServiceVersionExecute(r)
}

/*
GetSpecificDatabaseServiceVersion Get a version of the database service

Get a version of the database service by given `id`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id database service Id
 @param version database service version number in the form `major`.`minor`
 @return ApiGetSpecificDatabaseServiceVersionRequest
*/
func (a *DatabaseServiceApiService) GetSpecificDatabaseServiceVersion(ctx context.Context, id string, version string) ApiGetSpecificDatabaseServiceVersionRequest {
	return ApiGetSpecificDatabaseServiceVersionRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		version:    version,
	}
}

// Execute executes the request
//  @return DatabaseService
func (a *DatabaseServiceApiService) GetSpecificDatabaseServiceVersionExecute(r ApiGetSpecificDatabaseServiceVersionRequest) (*DatabaseService, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseService
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseServiceApiService.GetSpecificDatabaseServiceVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/databaseServices/{id}/versions/{version}"
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

type ApiListAllDatabaseServiceVersionRequest struct {
	ctx        context.Context
	ApiService *DatabaseServiceApiService
	id         string
}

func (r ApiListAllDatabaseServiceVersionRequest) Execute() (*EntityHistory, *http.Response, error) {
	return r.ApiService.ListAllDatabaseServiceVersionExecute(r)
}

/*
ListAllDatabaseServiceVersion List database service versions

Get a list of all the versions of a database service identified by `id`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id database service Id
 @return ApiListAllDatabaseServiceVersionRequest
*/
func (a *DatabaseServiceApiService) ListAllDatabaseServiceVersion(ctx context.Context, id string) ApiListAllDatabaseServiceVersionRequest {
	return ApiListAllDatabaseServiceVersionRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return EntityHistory
func (a *DatabaseServiceApiService) ListAllDatabaseServiceVersionExecute(r ApiListAllDatabaseServiceVersionRequest) (*EntityHistory, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EntityHistory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseServiceApiService.ListAllDatabaseServiceVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/databaseServices/{id}/versions"
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

type ApiListDatabaseServicesRequest struct {
	ctx        context.Context
	ApiService *DatabaseServiceApiService
	fields     *string
	limit      *int32
	before     *string
	after      *string
	include    *string
}

// Fields requested in the returned resource
func (r ApiListDatabaseServicesRequest) Fields(fields string) ApiListDatabaseServicesRequest {
	r.fields = &fields
	return r
}

func (r ApiListDatabaseServicesRequest) Limit(limit int32) ApiListDatabaseServicesRequest {
	r.limit = &limit
	return r
}

// Returns list of database services before this cursor
func (r ApiListDatabaseServicesRequest) Before(before string) ApiListDatabaseServicesRequest {
	r.before = &before
	return r
}

// Returns list of database services after this cursor
func (r ApiListDatabaseServicesRequest) After(after string) ApiListDatabaseServicesRequest {
	r.after = &after
	return r
}

// Include all, deleted, or non-deleted entities.
func (r ApiListDatabaseServicesRequest) Include(include string) ApiListDatabaseServicesRequest {
	r.include = &include
	return r
}

func (r ApiListDatabaseServicesRequest) Execute() (*DatabaseServiceList, *http.Response, error) {
	return r.ApiService.ListDatabaseServicesExecute(r)
}

/*
ListDatabaseServices List database services

Get a list of database services.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListDatabaseServicesRequest
*/
func (a *DatabaseServiceApiService) ListDatabaseServices(ctx context.Context) ApiListDatabaseServicesRequest {
	return ApiListDatabaseServicesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return DatabaseServiceList
func (a *DatabaseServiceApiService) ListDatabaseServicesExecute(r ApiListDatabaseServicesRequest) (*DatabaseServiceList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseServiceList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseServiceApiService.ListDatabaseServices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/databaseServices"

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