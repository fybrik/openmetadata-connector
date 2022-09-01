# \MlModelServiceApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMlModelService**](MlModelServiceApi.md#CreateMlModelService) | **Post** /v1/services/mlmodelServices | Create a mlModel service
[**CreateOrUpdateMlModelService**](MlModelServiceApi.md#CreateOrUpdateMlModelService) | **Put** /v1/services/mlmodelServices | Update mlModel service
[**DeleteMlModelService**](MlModelServiceApi.md#DeleteMlModelService) | **Delete** /v1/services/mlmodelServices/{id} | Delete a mlModel service
[**GetMlModelServiceByFQN**](MlModelServiceApi.md#GetMlModelServiceByFQN) | **Get** /v1/services/mlmodelServices/name/{name} | Get mlModel service by name
[**GetMlModelServiceByID**](MlModelServiceApi.md#GetMlModelServiceByID) | **Get** /v1/services/mlmodelServices/{id} | Get a mlModel service
[**GetSpecificMlModelService**](MlModelServiceApi.md#GetSpecificMlModelService) | **Get** /v1/services/mlmodelServices/{id}/versions/{version} | Get a version of the mlModel service
[**ListAllMlModelServiceVersion**](MlModelServiceApi.md#ListAllMlModelServiceVersion) | **Get** /v1/services/mlmodelServices/{id}/versions | List mlModel service versions
[**ListMlModelService**](MlModelServiceApi.md#ListMlModelService) | **Get** /v1/services/mlmodelServices | List mlModel services



## CreateMlModelService

> MlModelService CreateMlModelService(ctx).CreateMlModelService(createMlModelService).Execute()

Create a mlModel service



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
    createMlModelService := *openapiclient.NewCreateMlModelService(*openapiclient.NewMlModelConnection(), "Name_example", "ServiceType_example") // CreateMlModelService |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MlModelServiceApi.CreateMlModelService(context.Background()).CreateMlModelService(createMlModelService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelServiceApi.CreateMlModelService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMlModelService`: MlModelService
    fmt.Fprintf(os.Stdout, "Response from `MlModelServiceApi.CreateMlModelService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMlModelServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMlModelService** | [**CreateMlModelService**](CreateMlModelService.md) |  | 

### Return type

[**MlModelService**](MlModelService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateMlModelService

> MlModelService CreateOrUpdateMlModelService(ctx).CreateMlModelService(createMlModelService).Execute()

Update mlModel service



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
    createMlModelService := *openapiclient.NewCreateMlModelService(*openapiclient.NewMlModelConnection(), "Name_example", "ServiceType_example") // CreateMlModelService |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MlModelServiceApi.CreateOrUpdateMlModelService(context.Background()).CreateMlModelService(createMlModelService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelServiceApi.CreateOrUpdateMlModelService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateMlModelService`: MlModelService
    fmt.Fprintf(os.Stdout, "Response from `MlModelServiceApi.CreateOrUpdateMlModelService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateMlModelServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMlModelService** | [**CreateMlModelService**](CreateMlModelService.md) |  | 

### Return type

[**MlModelService**](MlModelService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMlModelService

> DeleteMlModelService(ctx, id).Recursive(recursive).HardDelete(hardDelete).Execute()

Delete a mlModel service



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
    id := "id_example" // string | Id of the mlModel service
    recursive := true // bool | Recursively delete this entity and it's children. (Default `false`) (optional) (default to false)
    hardDelete := true // bool | Hard delete the entity. (Default = `false`) (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MlModelServiceApi.DeleteMlModelService(context.Background(), id).Recursive(recursive).HardDelete(hardDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelServiceApi.DeleteMlModelService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the mlModel service | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMlModelServiceRequest struct via the builder pattern


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


## GetMlModelServiceByFQN

> MlModelService GetMlModelServiceByFQN(ctx, name).Fields(fields).Include(include).Execute()

Get mlModel service by name



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
    resp, r, err := apiClient.MlModelServiceApi.GetMlModelServiceByFQN(context.Background(), name).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelServiceApi.GetMlModelServiceByFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMlModelServiceByFQN`: MlModelService
    fmt.Fprintf(os.Stdout, "Response from `MlModelServiceApi.GetMlModelServiceByFQN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMlModelServiceByFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**MlModelService**](MlModelService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMlModelServiceByID

> MlModelService GetMlModelServiceByID(ctx, id).Fields(fields).Include(include).Execute()

Get a mlModel service



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
    resp, r, err := apiClient.MlModelServiceApi.GetMlModelServiceByID(context.Background(), id).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelServiceApi.GetMlModelServiceByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMlModelServiceByID`: MlModelService
    fmt.Fprintf(os.Stdout, "Response from `MlModelServiceApi.GetMlModelServiceByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMlModelServiceByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**MlModelService**](MlModelService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificMlModelService

> MlModelService GetSpecificMlModelService(ctx, id, version).Execute()

Get a version of the mlModel service



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
    id := "id_example" // string | mlModel service Id
    version := "0.1 or 1.1" // string | mlModel service version number in the form `major`.`minor`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MlModelServiceApi.GetSpecificMlModelService(context.Background(), id, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelServiceApi.GetSpecificMlModelService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificMlModelService`: MlModelService
    fmt.Fprintf(os.Stdout, "Response from `MlModelServiceApi.GetSpecificMlModelService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | mlModel service Id | 
**version** | **string** | mlModel service version number in the form &#x60;major&#x60;.&#x60;minor&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificMlModelServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MlModelService**](MlModelService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllMlModelServiceVersion

> EntityHistory ListAllMlModelServiceVersion(ctx, id).Execute()

List mlModel service versions



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
    id := "id_example" // string | mlModel service Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MlModelServiceApi.ListAllMlModelServiceVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelServiceApi.ListAllMlModelServiceVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllMlModelServiceVersion`: EntityHistory
    fmt.Fprintf(os.Stdout, "Response from `MlModelServiceApi.ListAllMlModelServiceVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | mlModel service Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllMlModelServiceVersionRequest struct via the builder pattern


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


## ListMlModelService

> MlModelServiceList ListMlModelService(ctx).Fields(fields).Limit(limit).Before(before).After(after).Include(include).Execute()

List mlModel services



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
    resp, r, err := apiClient.MlModelServiceApi.ListMlModelService(context.Background()).Fields(fields).Limit(limit).Before(before).After(after).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelServiceApi.ListMlModelService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMlModelService`: MlModelServiceList
    fmt.Fprintf(os.Stdout, "Response from `MlModelServiceApi.ListMlModelService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMlModelServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Fields requested in the returned resource | 
 **limit** | **int32** | Limit number services returned. (1 to 1000000, default 10) | [default to 10]
 **before** | **string** | Returns list of services before this cursor | 
 **after** | **string** | Returns list of services after this cursor | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**MlModelServiceList**](MlModelServiceList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

