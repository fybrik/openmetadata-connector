# \DatabaseServiceApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatabaseService**](DatabaseServiceApi.md#CreateDatabaseService) | **Post** /v1/services/databaseServices | Create database service
[**CreateOrUpdateDatabaseService**](DatabaseServiceApi.md#CreateOrUpdateDatabaseService) | **Put** /v1/services/databaseServices | Update database service
[**DeleteDatabaseService**](DatabaseServiceApi.md#DeleteDatabaseService) | **Delete** /v1/services/databaseServices/{id} | Delete a database service
[**GetDatabaseServiceByFQN**](DatabaseServiceApi.md#GetDatabaseServiceByFQN) | **Get** /v1/services/databaseServices/name/{name} | Get database service by name
[**GetDatabaseServiceByID**](DatabaseServiceApi.md#GetDatabaseServiceByID) | **Get** /v1/services/databaseServices/{id} | Get a database service
[**GetSpecificDatabaseServiceVersion**](DatabaseServiceApi.md#GetSpecificDatabaseServiceVersion) | **Get** /v1/services/databaseServices/{id}/versions/{version} | Get a version of the database service
[**ListAllDatabaseServiceVersion**](DatabaseServiceApi.md#ListAllDatabaseServiceVersion) | **Get** /v1/services/databaseServices/{id}/versions | List database service versions
[**ListDatabaseServices**](DatabaseServiceApi.md#ListDatabaseServices) | **Get** /v1/services/databaseServices | List database services



## CreateDatabaseService

> DatabaseService CreateDatabaseService(ctx).CreateDatabaseService(createDatabaseService).Execute()

Create database service



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
    createDatabaseService := *openapiclient.NewCreateDatabaseService(*openapiclient.NewDatabaseConnection(), "Name_example", "ServiceType_example") // CreateDatabaseService |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServiceApi.CreateDatabaseService(context.Background()).CreateDatabaseService(createDatabaseService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServiceApi.CreateDatabaseService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatabaseService`: DatabaseService
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServiceApi.CreateDatabaseService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDatabaseService** | [**CreateDatabaseService**](CreateDatabaseService.md) |  | 

### Return type

[**DatabaseService**](DatabaseService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateDatabaseService

> DatabaseService CreateOrUpdateDatabaseService(ctx).CreateDatabaseService(createDatabaseService).Execute()

Update database service



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
    createDatabaseService := *openapiclient.NewCreateDatabaseService(*openapiclient.NewDatabaseConnection(), "Name_example", "ServiceType_example") // CreateDatabaseService |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServiceApi.CreateOrUpdateDatabaseService(context.Background()).CreateDatabaseService(createDatabaseService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServiceApi.CreateOrUpdateDatabaseService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateDatabaseService`: DatabaseService
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServiceApi.CreateOrUpdateDatabaseService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateDatabaseServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDatabaseService** | [**CreateDatabaseService**](CreateDatabaseService.md) |  | 

### Return type

[**DatabaseService**](DatabaseService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabaseService

> DeleteDatabaseService(ctx, id).Recursive(recursive).HardDelete(hardDelete).Execute()

Delete a database service



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
    id := "id_example" // string | Id of the database service
    recursive := true // bool | Recursively delete this entity and it's children. (Default `false`) (optional) (default to false)
    hardDelete := true // bool | Hard delete the entity. (Default = `false`) (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServiceApi.DeleteDatabaseService(context.Background(), id).Recursive(recursive).HardDelete(hardDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServiceApi.DeleteDatabaseService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the database service | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseServiceRequest struct via the builder pattern


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


## GetDatabaseServiceByFQN

> DatabaseService GetDatabaseServiceByFQN(ctx, name).Fields(fields).Include(include).Execute()

Get database service by name



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
    resp, r, err := apiClient.DatabaseServiceApi.GetDatabaseServiceByFQN(context.Background(), name).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServiceApi.GetDatabaseServiceByFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseServiceByFQN`: DatabaseService
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServiceApi.GetDatabaseServiceByFQN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseServiceByFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**DatabaseService**](DatabaseService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseServiceByID

> DatabaseService GetDatabaseServiceByID(ctx, id).Fields(fields).Include(include).Execute()

Get a database service



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
    resp, r, err := apiClient.DatabaseServiceApi.GetDatabaseServiceByID(context.Background(), id).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServiceApi.GetDatabaseServiceByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseServiceByID`: DatabaseService
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServiceApi.GetDatabaseServiceByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseServiceByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**DatabaseService**](DatabaseService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificDatabaseServiceVersion

> DatabaseService GetSpecificDatabaseServiceVersion(ctx, id, version).Execute()

Get a version of the database service



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
    id := "id_example" // string | database service Id
    version := "0.1 or 1.1" // string | database service version number in the form `major`.`minor`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServiceApi.GetSpecificDatabaseServiceVersion(context.Background(), id, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServiceApi.GetSpecificDatabaseServiceVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificDatabaseServiceVersion`: DatabaseService
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServiceApi.GetSpecificDatabaseServiceVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | database service Id | 
**version** | **string** | database service version number in the form &#x60;major&#x60;.&#x60;minor&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificDatabaseServiceVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DatabaseService**](DatabaseService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllDatabaseServiceVersion

> EntityHistory ListAllDatabaseServiceVersion(ctx, id).Execute()

List database service versions



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
    id := "id_example" // string | database service Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServiceApi.ListAllDatabaseServiceVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServiceApi.ListAllDatabaseServiceVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllDatabaseServiceVersion`: EntityHistory
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServiceApi.ListAllDatabaseServiceVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | database service Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllDatabaseServiceVersionRequest struct via the builder pattern


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


## ListDatabaseServices

> DatabaseServiceList ListDatabaseServices(ctx).Fields(fields).Limit(limit).Before(before).After(after).Include(include).Execute()

List database services



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
    limit := int32(56) // int32 |  (optional) (default to 10)
    before := "before_example" // string | Returns list of database services before this cursor (optional)
    after := "after_example" // string | Returns list of database services after this cursor (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServiceApi.ListDatabaseServices(context.Background()).Fields(fields).Limit(limit).Before(before).After(after).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServiceApi.ListDatabaseServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatabaseServices`: DatabaseServiceList
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServiceApi.ListDatabaseServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDatabaseServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Fields requested in the returned resource | 
 **limit** | **int32** |  | [default to 10]
 **before** | **string** | Returns list of database services before this cursor | 
 **after** | **string** | Returns list of database services after this cursor | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**DatabaseServiceList**](DatabaseServiceList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

