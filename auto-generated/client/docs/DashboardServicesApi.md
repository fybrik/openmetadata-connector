# \DashboardServicesApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboardService**](DashboardServicesApi.md#CreateDashboardService) | **Post** /v1/services/dashboardServices | Create a dashboard service
[**CreateOrUpdateDashboardService**](DashboardServicesApi.md#CreateOrUpdateDashboardService) | **Put** /v1/services/dashboardServices | Update a Dashboard service
[**DeleteDashboardService**](DashboardServicesApi.md#DeleteDashboardService) | **Delete** /v1/services/dashboardServices/{id} | Delete a Dashboard service
[**GetDashboardServiceByFQN**](DashboardServicesApi.md#GetDashboardServiceByFQN) | **Get** /v1/services/dashboardServices/name/{name} | Get dashboard service by name
[**GetDashboardServiceByID**](DashboardServicesApi.md#GetDashboardServiceByID) | **Get** /v1/services/dashboardServices/{id} | Get a dashboard service
[**GetSpecificDashboardServiceVersion**](DashboardServicesApi.md#GetSpecificDashboardServiceVersion) | **Get** /v1/services/dashboardServices/{id}/versions/{version} | Get a version of the dashboard service
[**ListAllDashboardServiceVersion**](DashboardServicesApi.md#ListAllDashboardServiceVersion) | **Get** /v1/services/dashboardServices/{id}/versions | List dashboard service versions
[**ListDashboardsService**](DashboardServicesApi.md#ListDashboardsService) | **Get** /v1/services/dashboardServices | List dashboard services



## CreateDashboardService

> DashboardService CreateDashboardService(ctx).CreateDashboardService(createDashboardService).Execute()

Create a dashboard service



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
    createDashboardService := *openapiclient.NewCreateDashboardService(*openapiclient.NewDashboardConnection(), "Name_example", "ServiceType_example") // CreateDashboardService |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardServicesApi.CreateDashboardService(context.Background()).CreateDashboardService(createDashboardService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardServicesApi.CreateDashboardService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDashboardService`: DashboardService
    fmt.Fprintf(os.Stdout, "Response from `DashboardServicesApi.CreateDashboardService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDashboardService** | [**CreateDashboardService**](CreateDashboardService.md) |  | 

### Return type

[**DashboardService**](DashboardService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateDashboardService

> DashboardService CreateOrUpdateDashboardService(ctx).CreateDashboardService(createDashboardService).Execute()

Update a Dashboard service



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
    createDashboardService := *openapiclient.NewCreateDashboardService(*openapiclient.NewDashboardConnection(), "Name_example", "ServiceType_example") // CreateDashboardService |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardServicesApi.CreateOrUpdateDashboardService(context.Background()).CreateDashboardService(createDashboardService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardServicesApi.CreateOrUpdateDashboardService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateDashboardService`: DashboardService
    fmt.Fprintf(os.Stdout, "Response from `DashboardServicesApi.CreateOrUpdateDashboardService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateDashboardServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDashboardService** | [**CreateDashboardService**](CreateDashboardService.md) |  | 

### Return type

[**DashboardService**](DashboardService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboardService

> DeleteDashboardService(ctx, id).Recursive(recursive).HardDelete(hardDelete).Execute()

Delete a Dashboard service



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
    id := "id_example" // string | Id of the dashboard service
    recursive := true // bool | Recursively delete this entity and it's children. (Default `false`) (optional) (default to false)
    hardDelete := true // bool | Hard delete the entity. (Default = `false`) (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardServicesApi.DeleteDashboardService(context.Background(), id).Recursive(recursive).HardDelete(hardDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardServicesApi.DeleteDashboardService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the dashboard service | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardServiceRequest struct via the builder pattern


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


## GetDashboardServiceByFQN

> DashboardService GetDashboardServiceByFQN(ctx, name).Fields(fields).Include(include).Execute()

Get dashboard service by name



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
    resp, r, err := apiClient.DashboardServicesApi.GetDashboardServiceByFQN(context.Background(), name).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardServicesApi.GetDashboardServiceByFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboardServiceByFQN`: DashboardService
    fmt.Fprintf(os.Stdout, "Response from `DashboardServicesApi.GetDashboardServiceByFQN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardServiceByFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**DashboardService**](DashboardService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardServiceByID

> DashboardService GetDashboardServiceByID(ctx, id).Fields(fields).Include(include).Execute()

Get a dashboard service



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
    resp, r, err := apiClient.DashboardServicesApi.GetDashboardServiceByID(context.Background(), id).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardServicesApi.GetDashboardServiceByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboardServiceByID`: DashboardService
    fmt.Fprintf(os.Stdout, "Response from `DashboardServicesApi.GetDashboardServiceByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardServiceByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**DashboardService**](DashboardService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificDashboardServiceVersion

> DashboardService GetSpecificDashboardServiceVersion(ctx, id, version).Execute()

Get a version of the dashboard service



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
    id := "id_example" // string | dashboard service Id
    version := "0.1 or 1.1" // string | dashboard service version number in the form `major`.`minor`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardServicesApi.GetSpecificDashboardServiceVersion(context.Background(), id, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardServicesApi.GetSpecificDashboardServiceVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificDashboardServiceVersion`: DashboardService
    fmt.Fprintf(os.Stdout, "Response from `DashboardServicesApi.GetSpecificDashboardServiceVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | dashboard service Id | 
**version** | **string** | dashboard service version number in the form &#x60;major&#x60;.&#x60;minor&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificDashboardServiceVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DashboardService**](DashboardService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllDashboardServiceVersion

> EntityHistory ListAllDashboardServiceVersion(ctx, id).Execute()

List dashboard service versions



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
    id := "id_example" // string | dashboard service Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardServicesApi.ListAllDashboardServiceVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardServicesApi.ListAllDashboardServiceVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllDashboardServiceVersion`: EntityHistory
    fmt.Fprintf(os.Stdout, "Response from `DashboardServicesApi.ListAllDashboardServiceVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | dashboard service Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllDashboardServiceVersionRequest struct via the builder pattern


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


## ListDashboardsService

> DashboardServiceList ListDashboardsService(ctx).Name(name).Fields(fields).Limit(limit).Before(before).After(after).Include(include).Execute()

List dashboard services



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
    name := "name_example" // string |  (optional)
    fields := "owner" // string | Fields requested in the returned resource (optional)
    limit := int32(56) // int32 |  (optional) (default to 10)
    before := "before_example" // string | Returns list of dashboard services before this cursor (optional)
    after := "after_example" // string | Returns list of dashboard services after this cursor (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardServicesApi.ListDashboardsService(context.Background()).Name(name).Fields(fields).Limit(limit).Before(before).After(after).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardServicesApi.ListDashboardsService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDashboardsService`: DashboardServiceList
    fmt.Fprintf(os.Stdout, "Response from `DashboardServicesApi.ListDashboardsService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDashboardsServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **fields** | **string** | Fields requested in the returned resource | 
 **limit** | **int32** |  | [default to 10]
 **before** | **string** | Returns list of dashboard services before this cursor | 
 **after** | **string** | Returns list of dashboard services after this cursor | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**DashboardServiceList**](DashboardServiceList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

