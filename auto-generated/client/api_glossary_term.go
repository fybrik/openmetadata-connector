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

// GlossaryTermApiService GlossaryTermApi service
type GlossaryTermApiService service

type ApiCreateGlossaryTermRequest struct {
	ctx                context.Context
	ApiService         *GlossaryTermApiService
	createGlossaryTerm *CreateGlossaryTerm
}

func (r ApiCreateGlossaryTermRequest) CreateGlossaryTerm(createGlossaryTerm CreateGlossaryTerm) ApiCreateGlossaryTermRequest {
	r.createGlossaryTerm = &createGlossaryTerm
	return r
}

func (r ApiCreateGlossaryTermRequest) Execute() (*GlossaryTerm, *http.Response, error) {
	return r.ApiService.CreateGlossaryTermExecute(r)
}

/*
CreateGlossaryTerm Create a glossary term

Create a new glossary term.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateGlossaryTermRequest
*/
func (a *GlossaryTermApiService) CreateGlossaryTerm(ctx context.Context) ApiCreateGlossaryTermRequest {
	return ApiCreateGlossaryTermRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GlossaryTerm
func (a *GlossaryTermApiService) CreateGlossaryTermExecute(r ApiCreateGlossaryTermRequest) (*GlossaryTerm, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GlossaryTerm
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlossaryTermApiService.CreateGlossaryTerm")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/glossaryTerms"

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
	localVarPostBody = r.createGlossaryTerm
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

type ApiCreateOrUpdateGlossaryTermRequest struct {
	ctx                context.Context
	ApiService         *GlossaryTermApiService
	createGlossaryTerm *CreateGlossaryTerm
}

func (r ApiCreateOrUpdateGlossaryTermRequest) CreateGlossaryTerm(createGlossaryTerm CreateGlossaryTerm) ApiCreateOrUpdateGlossaryTermRequest {
	r.createGlossaryTerm = &createGlossaryTerm
	return r
}

func (r ApiCreateOrUpdateGlossaryTermRequest) Execute() (*GlossaryTerm, *http.Response, error) {
	return r.ApiService.CreateOrUpdateGlossaryTermExecute(r)
}

/*
CreateOrUpdateGlossaryTerm Create or update a glossary term

Create a new glossary term, if it does not exist or update an existing glossary term.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateOrUpdateGlossaryTermRequest
*/
func (a *GlossaryTermApiService) CreateOrUpdateGlossaryTerm(ctx context.Context) ApiCreateOrUpdateGlossaryTermRequest {
	return ApiCreateOrUpdateGlossaryTermRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GlossaryTerm
func (a *GlossaryTermApiService) CreateOrUpdateGlossaryTermExecute(r ApiCreateOrUpdateGlossaryTermRequest) (*GlossaryTerm, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GlossaryTerm
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlossaryTermApiService.CreateOrUpdateGlossaryTerm")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/glossaryTerms"

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
	localVarPostBody = r.createGlossaryTerm
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

type ApiDeleteRequest struct {
	ctx        context.Context
	ApiService *GlossaryTermApiService
	id         string
	recursive  *bool
	hardDelete *bool
}

// Recursively delete this entity and it&#39;s children. (Default &#x60;false&#x60;)
func (r ApiDeleteRequest) Recursive(recursive bool) ApiDeleteRequest {
	r.recursive = &recursive
	return r
}

// Hard delete the entity. (Default &#x3D; &#x60;false&#x60;)
func (r ApiDeleteRequest) HardDelete(hardDelete bool) ApiDeleteRequest {
	r.hardDelete = &hardDelete
	return r
}

func (r ApiDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Delete a glossary term

Delete a glossary term by `id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Glossary Term Id
 @return ApiDeleteRequest
*/
func (a *GlossaryTermApiService) Delete(ctx context.Context, id string) ApiDeleteRequest {
	return ApiDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *GlossaryTermApiService) DeleteExecute(r ApiDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlossaryTermApiService.Delete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/glossaryTerms/{id}"
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

type ApiGetGlossaryTermByFQNRequest struct {
	ctx        context.Context
	ApiService *GlossaryTermApiService
	name       string
	fields     *string
	include    *string
}

// Fields requested in the returned resource
func (r ApiGetGlossaryTermByFQNRequest) Fields(fields string) ApiGetGlossaryTermByFQNRequest {
	r.fields = &fields
	return r
}

// Include all, deleted, or non-deleted entities.
func (r ApiGetGlossaryTermByFQNRequest) Include(include string) ApiGetGlossaryTermByFQNRequest {
	r.include = &include
	return r
}

func (r ApiGetGlossaryTermByFQNRequest) Execute() (*Glossary, *http.Response, error) {
	return r.ApiService.GetGlossaryTermByFQNExecute(r)
}

/*
GetGlossaryTermByFQN Get a glossary term by name

Get a glossary term by name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @return ApiGetGlossaryTermByFQNRequest
*/
func (a *GlossaryTermApiService) GetGlossaryTermByFQN(ctx context.Context, name string) ApiGetGlossaryTermByFQNRequest {
	return ApiGetGlossaryTermByFQNRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//  @return Glossary
func (a *GlossaryTermApiService) GetGlossaryTermByFQNExecute(r ApiGetGlossaryTermByFQNRequest) (*Glossary, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Glossary
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlossaryTermApiService.GetGlossaryTermByFQN")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/glossaryTerms/name/{name}"
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

type ApiGetGlossaryTermByIDRequest struct {
	ctx        context.Context
	ApiService *GlossaryTermApiService
	id         string
	fields     *string
	include    *string
}

// Fields requested in the returned resource
func (r ApiGetGlossaryTermByIDRequest) Fields(fields string) ApiGetGlossaryTermByIDRequest {
	r.fields = &fields
	return r
}

// Include all, deleted, or non-deleted entities.
func (r ApiGetGlossaryTermByIDRequest) Include(include string) ApiGetGlossaryTermByIDRequest {
	r.include = &include
	return r
}

func (r ApiGetGlossaryTermByIDRequest) Execute() (*Glossary, *http.Response, error) {
	return r.ApiService.GetGlossaryTermByIDExecute(r)
}

/*
GetGlossaryTermByID Get a glossary term

Get a glossary term by `id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetGlossaryTermByIDRequest
*/
func (a *GlossaryTermApiService) GetGlossaryTermByID(ctx context.Context, id string) ApiGetGlossaryTermByIDRequest {
	return ApiGetGlossaryTermByIDRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return Glossary
func (a *GlossaryTermApiService) GetGlossaryTermByIDExecute(r ApiGetGlossaryTermByIDRequest) (*Glossary, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Glossary
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlossaryTermApiService.GetGlossaryTermByID")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/glossaryTerms/{id}"
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

type ApiGetSpecificGlossaryTermVersionRequest struct {
	ctx        context.Context
	ApiService *GlossaryTermApiService
	id         string
	version    string
}

func (r ApiGetSpecificGlossaryTermVersionRequest) Execute() (*Glossary, *http.Response, error) {
	return r.ApiService.GetSpecificGlossaryTermVersionExecute(r)
}

/*
GetSpecificGlossaryTermVersion Get a version of the glossary term

Get a version of the glossary term by given `id`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id glossary Id
 @param version glossary term version number in the form `major`.`minor`
 @return ApiGetSpecificGlossaryTermVersionRequest
*/
func (a *GlossaryTermApiService) GetSpecificGlossaryTermVersion(ctx context.Context, id string, version string) ApiGetSpecificGlossaryTermVersionRequest {
	return ApiGetSpecificGlossaryTermVersionRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		version:    version,
	}
}

// Execute executes the request
//  @return Glossary
func (a *GlossaryTermApiService) GetSpecificGlossaryTermVersionExecute(r ApiGetSpecificGlossaryTermVersionRequest) (*Glossary, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Glossary
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlossaryTermApiService.GetSpecificGlossaryTermVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/glossaryTerms/{id}/versions/{version}"
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

type ApiListAllGlossaryTermVersionRequest struct {
	ctx        context.Context
	ApiService *GlossaryTermApiService
	id         string
}

func (r ApiListAllGlossaryTermVersionRequest) Execute() (*EntityHistory, *http.Response, error) {
	return r.ApiService.ListAllGlossaryTermVersionExecute(r)
}

/*
ListAllGlossaryTermVersion List glossary term versions

Get a list of all the versions of a glossary terms identified by `id`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id glossary Id
 @return ApiListAllGlossaryTermVersionRequest
*/
func (a *GlossaryTermApiService) ListAllGlossaryTermVersion(ctx context.Context, id string) ApiListAllGlossaryTermVersionRequest {
	return ApiListAllGlossaryTermVersionRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return EntityHistory
func (a *GlossaryTermApiService) ListAllGlossaryTermVersionExecute(r ApiListAllGlossaryTermVersionRequest) (*EntityHistory, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EntityHistory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlossaryTermApiService.ListAllGlossaryTermVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/glossaryTerms/{id}/versions"
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

type ApiListGlossaryTermRequest struct {
	ctx        context.Context
	ApiService *GlossaryTermApiService
	glossary   *string
	parent     *string
	fields     *string
	limit      *int32
	before     *string
	after      *string
	include    *string
}

// List glossary terms filtered by glossary identified by Id given in &#x60;glossary&#x60; parameter.
func (r ApiListGlossaryTermRequest) Glossary(glossary string) ApiListGlossaryTermRequest {
	r.glossary = &glossary
	return r
}

// List glossary terms filtered by children of glossary term identified by Id given in &#x60;parent&#x60; parameter.
func (r ApiListGlossaryTermRequest) Parent(parent string) ApiListGlossaryTermRequest {
	r.parent = &parent
	return r
}

// Fields requested in the returned resource
func (r ApiListGlossaryTermRequest) Fields(fields string) ApiListGlossaryTermRequest {
	r.fields = &fields
	return r
}

// Limit the number glossary terms returned. (1 to 1000000, default &#x3D; 10)
func (r ApiListGlossaryTermRequest) Limit(limit int32) ApiListGlossaryTermRequest {
	r.limit = &limit
	return r
}

// Returns list of glossary terms before this cursor
func (r ApiListGlossaryTermRequest) Before(before string) ApiListGlossaryTermRequest {
	r.before = &before
	return r
}

// Returns list of glossary terms after this cursor
func (r ApiListGlossaryTermRequest) After(after string) ApiListGlossaryTermRequest {
	r.after = &after
	return r
}

// Include all, deleted, or non-deleted entities.
func (r ApiListGlossaryTermRequest) Include(include string) ApiListGlossaryTermRequest {
	r.include = &include
	return r
}

func (r ApiListGlossaryTermRequest) Execute() (*GlossaryTermList, *http.Response, error) {
	return r.ApiService.ListGlossaryTermExecute(r)
}

/*
ListGlossaryTerm List glossary terms

Get a list of glossary terms. Use `fields` parameter to get only necessary fields.  Use cursor-based pagination to limit the number entries in the list using `limit` and `before` or `after` query params.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListGlossaryTermRequest
*/
func (a *GlossaryTermApiService) ListGlossaryTerm(ctx context.Context) ApiListGlossaryTermRequest {
	return ApiListGlossaryTermRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GlossaryTermList
func (a *GlossaryTermApiService) ListGlossaryTermExecute(r ApiListGlossaryTermRequest) (*GlossaryTermList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GlossaryTermList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlossaryTermApiService.ListGlossaryTerm")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/glossaryTerms"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.glossary != nil {
		localVarQueryParams.Add("glossary", parameterToString(*r.glossary, ""))
	}
	if r.parent != nil {
		localVarQueryParams.Add("parent", parameterToString(*r.parent, ""))
	}
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

type ApiPatchGlossaryTermRequest struct {
	ctx        context.Context
	ApiService *GlossaryTermApiService
	id         string
	body       *map[string]interface{}
}

// JsonPatch with array of operations
func (r ApiPatchGlossaryTermRequest) Body(body map[string]interface{}) ApiPatchGlossaryTermRequest {
	r.body = &body
	return r
}

func (r ApiPatchGlossaryTermRequest) Execute() (*http.Response, error) {
	return r.ApiService.PatchGlossaryTermExecute(r)
}

/*
PatchGlossaryTerm Update a glossary term

Update an existing glossary term using JsonPatch.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiPatchGlossaryTermRequest
*/
func (a *GlossaryTermApiService) PatchGlossaryTerm(ctx context.Context, id string) ApiPatchGlossaryTermRequest {
	return ApiPatchGlossaryTermRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *GlossaryTermApiService) PatchGlossaryTermExecute(r ApiPatchGlossaryTermRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlossaryTermApiService.PatchGlossaryTerm")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/glossaryTerms/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json"}

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
	localVarPostBody = r.body
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
