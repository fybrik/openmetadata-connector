# \PipelineServiceApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdatePipelineService**](PipelineServiceApi.md#CreateOrUpdatePipelineService) | **Put** /v1/services/pipelineServices | Update pipeline service
[**CreatePipelineService**](PipelineServiceApi.md#CreatePipelineService) | **Post** /v1/services/pipelineServices | Create a pipeline service
[**DeletePipelineService**](PipelineServiceApi.md#DeletePipelineService) | **Delete** /v1/services/pipelineServices/{id} | Delete a pipeline service
[**GetPipelineServiceByFQN**](PipelineServiceApi.md#GetPipelineServiceByFQN) | **Get** /v1/services/pipelineServices/name/{name} | Get pipeline service by name
[**GetPipelineServiceByID**](PipelineServiceApi.md#GetPipelineServiceByID) | **Get** /v1/services/pipelineServices/{id} | Get a pipeline service
[**GetSpecificPipelineService**](PipelineServiceApi.md#GetSpecificPipelineService) | **Get** /v1/services/pipelineServices/{id}/versions/{version} | Get a version of the pipeline service
[**ListAllPipelineServiceVersion**](PipelineServiceApi.md#ListAllPipelineServiceVersion) | **Get** /v1/services/pipelineServices/{id}/versions | List pipeline service versions
[**ListPipelineService**](PipelineServiceApi.md#ListPipelineService) | **Get** /v1/services/pipelineServices | List pipeline services



## CreateOrUpdatePipelineService

> PipelineService CreateOrUpdatePipelineService(ctx).CreatePipelineService(createPipelineService).Execute()

Update pipeline service



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
    createPipelineService := *openapiclient.NewCreatePipelineService(*openapiclient.NewPipelineConnection(), "Name_example", "ServiceType_example") // CreatePipelineService |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineServiceApi.CreateOrUpdatePipelineService(context.Background()).CreatePipelineService(createPipelineService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineServiceApi.CreateOrUpdatePipelineService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdatePipelineService`: PipelineService
    fmt.Fprintf(os.Stdout, "Response from `PipelineServiceApi.CreateOrUpdatePipelineService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdatePipelineServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPipelineService** | [**CreatePipelineService**](CreatePipelineService.md) |  | 

### Return type

[**PipelineService**](PipelineService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePipelineService

> PipelineService CreatePipelineService(ctx).CreatePipelineService(createPipelineService).Execute()

Create a pipeline service



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
    createPipelineService := *openapiclient.NewCreatePipelineService(*openapiclient.NewPipelineConnection(), "Name_example", "ServiceType_example") // CreatePipelineService |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineServiceApi.CreatePipelineService(context.Background()).CreatePipelineService(createPipelineService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineServiceApi.CreatePipelineService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePipelineService`: PipelineService
    fmt.Fprintf(os.Stdout, "Response from `PipelineServiceApi.CreatePipelineService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePipelineServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPipelineService** | [**CreatePipelineService**](CreatePipelineService.md) |  | 

### Return type

[**PipelineService**](PipelineService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePipelineService

> DeletePipelineService(ctx, id).Recursive(recursive).HardDelete(hardDelete).Execute()

Delete a pipeline service



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
    id := "id_example" // string | Id of the pipeline service
    recursive := true // bool | Recursively delete this entity and it's children. (Default `false`) (optional) (default to false)
    hardDelete := true // bool | Hard delete the entity. (Default = `false`) (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineServiceApi.DeletePipelineService(context.Background(), id).Recursive(recursive).HardDelete(hardDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineServiceApi.DeletePipelineService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the pipeline service | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePipelineServiceRequest struct via the builder pattern


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


## GetPipelineServiceByFQN

> PipelineService GetPipelineServiceByFQN(ctx, name).Fields(fields).Include(include).Execute()

Get pipeline service by name



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
    fields := "pipelines,owner" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineServiceApi.GetPipelineServiceByFQN(context.Background(), name).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineServiceApi.GetPipelineServiceByFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPipelineServiceByFQN`: PipelineService
    fmt.Fprintf(os.Stdout, "Response from `PipelineServiceApi.GetPipelineServiceByFQN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineServiceByFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**PipelineService**](PipelineService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineServiceByID

> PipelineService GetPipelineServiceByID(ctx, id).Fields(fields).Include(include).Execute()

Get a pipeline service



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
    fields := "pipelines,owner" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineServiceApi.GetPipelineServiceByID(context.Background(), id).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineServiceApi.GetPipelineServiceByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPipelineServiceByID`: PipelineService
    fmt.Fprintf(os.Stdout, "Response from `PipelineServiceApi.GetPipelineServiceByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineServiceByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**PipelineService**](PipelineService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificPipelineService

> PipelineService GetSpecificPipelineService(ctx, id, version).Execute()

Get a version of the pipeline service



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
    id := "id_example" // string | pipeline service Id
    version := "0.1 or 1.1" // string | pipeline service version number in the form `major`.`minor`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineServiceApi.GetSpecificPipelineService(context.Background(), id, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineServiceApi.GetSpecificPipelineService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificPipelineService`: PipelineService
    fmt.Fprintf(os.Stdout, "Response from `PipelineServiceApi.GetSpecificPipelineService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | pipeline service Id | 
**version** | **string** | pipeline service version number in the form &#x60;major&#x60;.&#x60;minor&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificPipelineServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PipelineService**](PipelineService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllPipelineServiceVersion

> EntityHistory ListAllPipelineServiceVersion(ctx, id).Execute()

List pipeline service versions



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
    id := "id_example" // string | pipeline service Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineServiceApi.ListAllPipelineServiceVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineServiceApi.ListAllPipelineServiceVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllPipelineServiceVersion`: EntityHistory
    fmt.Fprintf(os.Stdout, "Response from `PipelineServiceApi.ListAllPipelineServiceVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | pipeline service Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllPipelineServiceVersionRequest struct via the builder pattern


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


## ListPipelineService

> PipelineServiceList ListPipelineService(ctx).Fields(fields).Limit(limit).Before(before).After(after).Include(include).Execute()

List pipeline services



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
    fields := "pipelines,owner" // string | Fields requested in the returned resource (optional)
    limit := int32(56) // int32 | Limit number services returned. (1 to 1000000, default 10) (optional) (default to 10)
    before := "before_example" // string | Returns list of services before this cursor (optional)
    after := "after_example" // string | Returns list of services after this cursor (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineServiceApi.ListPipelineService(context.Background()).Fields(fields).Limit(limit).Before(before).After(after).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineServiceApi.ListPipelineService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPipelineService`: PipelineServiceList
    fmt.Fprintf(os.Stdout, "Response from `PipelineServiceApi.ListPipelineService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPipelineServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Fields requested in the returned resource | 
 **limit** | **int32** | Limit number services returned. (1 to 1000000, default 10) | [default to 10]
 **before** | **string** | Returns list of services before this cursor | 
 **after** | **string** | Returns list of services after this cursor | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**PipelineServiceList**](PipelineServiceList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

