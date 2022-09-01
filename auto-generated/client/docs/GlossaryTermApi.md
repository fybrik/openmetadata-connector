# \GlossaryTermApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGlossaryTerm**](GlossaryTermApi.md#CreateGlossaryTerm) | **Post** /v1/glossaryTerms | Create a glossary term
[**CreateOrUpdateGlossaryTerm**](GlossaryTermApi.md#CreateOrUpdateGlossaryTerm) | **Put** /v1/glossaryTerms | Create or update a glossary term
[**Delete**](GlossaryTermApi.md#Delete) | **Delete** /v1/glossaryTerms/{id} | Delete a glossary term
[**GetGlossaryTermByFQN**](GlossaryTermApi.md#GetGlossaryTermByFQN) | **Get** /v1/glossaryTerms/name/{name} | Get a glossary term by name
[**GetGlossaryTermByID**](GlossaryTermApi.md#GetGlossaryTermByID) | **Get** /v1/glossaryTerms/{id} | Get a glossary term
[**GetSpecificGlossaryTermVersion**](GlossaryTermApi.md#GetSpecificGlossaryTermVersion) | **Get** /v1/glossaryTerms/{id}/versions/{version} | Get a version of the glossary term
[**ListAllGlossaryTermVersion**](GlossaryTermApi.md#ListAllGlossaryTermVersion) | **Get** /v1/glossaryTerms/{id}/versions | List glossary term versions
[**ListGlossaryTerm**](GlossaryTermApi.md#ListGlossaryTerm) | **Get** /v1/glossaryTerms | List glossary terms
[**PatchGlossaryTerm**](GlossaryTermApi.md#PatchGlossaryTerm) | **Patch** /v1/glossaryTerms/{id} | Update a glossary term



## CreateGlossaryTerm

> GlossaryTerm CreateGlossaryTerm(ctx).CreateGlossaryTerm(createGlossaryTerm).Execute()

Create a glossary term



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createGlossaryTerm := *openapiclient.NewCreateGlossaryTerm("Description_example", *openapiclient.NewEntityReference("Id_example", "Type_example"), "Name_example") // CreateGlossaryTerm |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlossaryTermApi.CreateGlossaryTerm(context.Background()).CreateGlossaryTerm(createGlossaryTerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlossaryTermApi.CreateGlossaryTerm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGlossaryTerm`: GlossaryTerm
    fmt.Fprintf(os.Stdout, "Response from `GlossaryTermApi.CreateGlossaryTerm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlossaryTermRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGlossaryTerm** | [**CreateGlossaryTerm**](CreateGlossaryTerm.md) |  | 

### Return type

[**GlossaryTerm**](GlossaryTerm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateGlossaryTerm

> GlossaryTerm CreateOrUpdateGlossaryTerm(ctx).CreateGlossaryTerm(createGlossaryTerm).Execute()

Create or update a glossary term



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createGlossaryTerm := *openapiclient.NewCreateGlossaryTerm("Description_example", *openapiclient.NewEntityReference("Id_example", "Type_example"), "Name_example") // CreateGlossaryTerm |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlossaryTermApi.CreateOrUpdateGlossaryTerm(context.Background()).CreateGlossaryTerm(createGlossaryTerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlossaryTermApi.CreateOrUpdateGlossaryTerm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateGlossaryTerm`: GlossaryTerm
    fmt.Fprintf(os.Stdout, "Response from `GlossaryTermApi.CreateOrUpdateGlossaryTerm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateGlossaryTermRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGlossaryTerm** | [**CreateGlossaryTerm**](CreateGlossaryTerm.md) |  | 

### Return type

[**GlossaryTerm**](GlossaryTerm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, id).Recursive(recursive).HardDelete(hardDelete).Execute()

Delete a glossary term



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Glossary Term Id
    recursive := true // bool | Recursively delete this entity and it's children. (Default `false`) (optional) (default to false)
    hardDelete := true // bool | Hard delete the entity. (Default = `false`) (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlossaryTermApi.Delete(context.Background(), id).Recursive(recursive).HardDelete(hardDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlossaryTermApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Glossary Term Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recursive** | **bool** | Recursively delete this entity and it&#39;s children. (Default &#x60;false&#x60;) | [default to false]
 **hardDelete** | **bool** | Hard delete the entity. (Default &#x3D; &#x60;false&#x60;) | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlossaryTermByFQN

> Glossary GetGlossaryTermByFQN(ctx, name).Fields(fields).Include(include).Execute()

Get a glossary term by name



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | 
    fields := "children,relatedTerms,reviewers,tags,usageCount" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlossaryTermApi.GetGlossaryTermByFQN(context.Background(), name).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlossaryTermApi.GetGlossaryTermByFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlossaryTermByFQN`: Glossary
    fmt.Fprintf(os.Stdout, "Response from `GlossaryTermApi.GetGlossaryTermByFQN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlossaryTermByFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**Glossary**](Glossary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlossaryTermByID

> Glossary GetGlossaryTermByID(ctx, id).Fields(fields).Include(include).Execute()

Get a glossary term



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    fields := "children,relatedTerms,reviewers,tags,usageCount" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlossaryTermApi.GetGlossaryTermByID(context.Background(), id).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlossaryTermApi.GetGlossaryTermByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlossaryTermByID`: Glossary
    fmt.Fprintf(os.Stdout, "Response from `GlossaryTermApi.GetGlossaryTermByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlossaryTermByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**Glossary**](Glossary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificGlossaryTermVersion

> Glossary GetSpecificGlossaryTermVersion(ctx, id, version).Execute()

Get a version of the glossary term



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | glossary Id
    version := "0.1 or 1.1" // string | glossary term version number in the form `major`.`minor`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlossaryTermApi.GetSpecificGlossaryTermVersion(context.Background(), id, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlossaryTermApi.GetSpecificGlossaryTermVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificGlossaryTermVersion`: Glossary
    fmt.Fprintf(os.Stdout, "Response from `GlossaryTermApi.GetSpecificGlossaryTermVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | glossary Id | 
**version** | **string** | glossary term version number in the form &#x60;major&#x60;.&#x60;minor&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificGlossaryTermVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Glossary**](Glossary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllGlossaryTermVersion

> EntityHistory ListAllGlossaryTermVersion(ctx, id).Execute()

List glossary term versions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | glossary Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlossaryTermApi.ListAllGlossaryTermVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlossaryTermApi.ListAllGlossaryTermVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllGlossaryTermVersion`: EntityHistory
    fmt.Fprintf(os.Stdout, "Response from `GlossaryTermApi.ListAllGlossaryTermVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | glossary Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllGlossaryTermVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntityHistory**](EntityHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlossaryTerm

> GlossaryTermList ListGlossaryTerm(ctx).Glossary(glossary).Parent(parent).Fields(fields).Limit(limit).Before(before).After(after).Include(include).Execute()

List glossary terms



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    glossary := "children,relatedTerms,reviewers,tags,usageCount" // string | List glossary terms filtered by glossary identified by Id given in `glossary` parameter. (optional)
    parent := "children,relatedTerms,reviewers,tags,usageCount" // string | List glossary terms filtered by children of glossary term identified by Id given in `parent` parameter. (optional)
    fields := "children,relatedTerms,reviewers,tags,usageCount" // string | Fields requested in the returned resource (optional)
    limit := int32(56) // int32 | Limit the number glossary terms returned. (1 to 1000000, default = 10) (optional) (default to 10)
    before := "before_example" // string | Returns list of glossary terms before this cursor (optional)
    after := "after_example" // string | Returns list of glossary terms after this cursor (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlossaryTermApi.ListGlossaryTerm(context.Background()).Glossary(glossary).Parent(parent).Fields(fields).Limit(limit).Before(before).After(after).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlossaryTermApi.ListGlossaryTerm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGlossaryTerm`: GlossaryTermList
    fmt.Fprintf(os.Stdout, "Response from `GlossaryTermApi.ListGlossaryTerm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGlossaryTermRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glossary** | **string** | List glossary terms filtered by glossary identified by Id given in &#x60;glossary&#x60; parameter. | 
 **parent** | **string** | List glossary terms filtered by children of glossary term identified by Id given in &#x60;parent&#x60; parameter. | 
 **fields** | **string** | Fields requested in the returned resource | 
 **limit** | **int32** | Limit the number glossary terms returned. (1 to 1000000, default &#x3D; 10) | [default to 10]
 **before** | **string** | Returns list of glossary terms before this cursor | 
 **after** | **string** | Returns list of glossary terms after this cursor | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**GlossaryTermList**](GlossaryTermList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchGlossaryTerm

> PatchGlossaryTerm(ctx, id).Body(body).Execute()

Update a glossary term



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} | JsonPatch with array of operations (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlossaryTermApi.PatchGlossaryTerm(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlossaryTermApi.PatchGlossaryTerm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchGlossaryTermRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | JsonPatch with array of operations | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

