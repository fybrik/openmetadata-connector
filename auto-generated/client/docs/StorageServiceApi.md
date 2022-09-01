# \StorageServiceApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateStorageService**](StorageServiceApi.md#CreateOrUpdateStorageService) | **Put** /v1/services/storageServices | Update storage service
[**CreateStorageService**](StorageServiceApi.md#CreateStorageService) | **Post** /v1/services/storageServices | Create storage service
[**DeleteStorageService**](StorageServiceApi.md#DeleteStorageService) | **Delete** /v1/services/storageServices/{id} | Delete a storage service
[**GetSpecificStorageServiceVersion**](StorageServiceApi.md#GetSpecificStorageServiceVersion) | **Get** /v1/services/storageServices/{id}/versions/{version} | Get a version of the storage service
[**GetStorageServiceByFQN**](StorageServiceApi.md#GetStorageServiceByFQN) | **Get** /v1/services/storageServices/name/{name} | Get storage service by name
[**GetStorageServiceByID**](StorageServiceApi.md#GetStorageServiceByID) | **Get** /v1/services/storageServices/{id} | Get a storage service
[**ListAllStorageServiceVersion**](StorageServiceApi.md#ListAllStorageServiceVersion) | **Get** /v1/services/storageServices/{id}/versions | List storage service versions
[**ListStorageService**](StorageServiceApi.md#ListStorageService) | **Get** /v1/services/storageServices | List storage services



## CreateOrUpdateStorageService

> StorageService CreateOrUpdateStorageService(ctx).CreateStorageService(createStorageService).Execute()

Update storage service



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
    createStorageService := *openapiclient.NewCreateStorageService("Name_example") // CreateStorageService |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageServiceApi.CreateOrUpdateStorageService(context.Background()).CreateStorageService(createStorageService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageServiceApi.CreateOrUpdateStorageService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateStorageService`: StorageService
    fmt.Fprintf(os.Stdout, "Response from `StorageServiceApi.CreateOrUpdateStorageService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateStorageServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStorageService** | [**CreateStorageService**](CreateStorageService.md) |  | 

### Return type

[**StorageService**](StorageService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStorageService

> StorageService CreateStorageService(ctx).CreateStorageService(createStorageService).Execute()

Create storage service



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
    createStorageService := *openapiclient.NewCreateStorageService("Name_example") // CreateStorageService |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageServiceApi.CreateStorageService(context.Background()).CreateStorageService(createStorageService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageServiceApi.CreateStorageService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorageService`: StorageService
    fmt.Fprintf(os.Stdout, "Response from `StorageServiceApi.CreateStorageService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorageServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStorageService** | [**CreateStorageService**](CreateStorageService.md) |  | 

### Return type

[**StorageService**](StorageService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStorageService

> DeleteStorageService(ctx, id).Recursive(recursive).HardDelete(hardDelete).Execute()

Delete a storage service



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
    id := "id_example" // string | Id of the storage service
    recursive := true // bool | Recursively delete this entity and it's children. (Default `false`) (optional) (default to false)
    hardDelete := true // bool | Hard delete the entity. (Default = `false`) (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageServiceApi.DeleteStorageService(context.Background(), id).Recursive(recursive).HardDelete(hardDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageServiceApi.DeleteStorageService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the storage service | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorageServiceRequest struct via the builder pattern


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


## GetSpecificStorageServiceVersion

> StorageService GetSpecificStorageServiceVersion(ctx, id, version).Execute()

Get a version of the storage service



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
    id := "id_example" // string | storage service Id
    version := "0.1 or 1.1" // string | storage service version number in the form `major`.`minor`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageServiceApi.GetSpecificStorageServiceVersion(context.Background(), id, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageServiceApi.GetSpecificStorageServiceVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificStorageServiceVersion`: StorageService
    fmt.Fprintf(os.Stdout, "Response from `StorageServiceApi.GetSpecificStorageServiceVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | storage service Id | 
**version** | **string** | storage service version number in the form &#x60;major&#x60;.&#x60;minor&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificStorageServiceVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StorageService**](StorageService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageServiceByFQN

> StorageService GetStorageServiceByFQN(ctx, name).Fields(fields).Include(include).Execute()

Get storage service by name



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
    fields := "owner" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageServiceApi.GetStorageServiceByFQN(context.Background(), name).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageServiceApi.GetStorageServiceByFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorageServiceByFQN`: StorageService
    fmt.Fprintf(os.Stdout, "Response from `StorageServiceApi.GetStorageServiceByFQN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageServiceByFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**StorageService**](StorageService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageServiceByID

> StorageService GetStorageServiceByID(ctx, id).Fields(fields).Include(include).Execute()

Get a storage service



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
    fields := "owner" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageServiceApi.GetStorageServiceByID(context.Background(), id).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageServiceApi.GetStorageServiceByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorageServiceByID`: StorageService
    fmt.Fprintf(os.Stdout, "Response from `StorageServiceApi.GetStorageServiceByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageServiceByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**StorageService**](StorageService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllStorageServiceVersion

> EntityHistory ListAllStorageServiceVersion(ctx, id).Execute()

List storage service versions



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
    id := "id_example" // string | storage service Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageServiceApi.ListAllStorageServiceVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageServiceApi.ListAllStorageServiceVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllStorageServiceVersion`: EntityHistory
    fmt.Fprintf(os.Stdout, "Response from `StorageServiceApi.ListAllStorageServiceVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | storage service Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllStorageServiceVersionRequest struct via the builder pattern


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


## ListStorageService

> StorageServiceList ListStorageService(ctx).Fields(fields).Limit(limit).Before(before).After(after).Include(include).Execute()

List storage services



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
    fields := "owner" // string | Fields requested in the returned resource (optional)
    limit := int32(56) // int32 | Limit number of services returned. (1 to 1000000, default 10) (optional) (default to 10)
    before := "before_example" // string | Returns list of services before this cursor (optional)
    after := "after_example" // string | Returns list of services after this cursor (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageServiceApi.ListStorageService(context.Background()).Fields(fields).Limit(limit).Before(before).After(after).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageServiceApi.ListStorageService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStorageService`: StorageServiceList
    fmt.Fprintf(os.Stdout, "Response from `StorageServiceApi.ListStorageService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStorageServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Fields requested in the returned resource | 
 **limit** | **int32** | Limit number of services returned. (1 to 1000000, default 10) | [default to 10]
 **before** | **string** | Returns list of services before this cursor | 
 **after** | **string** | Returns list of services after this cursor | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**StorageServiceList**](StorageServiceList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

