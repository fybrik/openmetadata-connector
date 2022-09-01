# \MessagingServiceApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessagingService**](MessagingServiceApi.md#CreateMessagingService) | **Post** /v1/services/messagingServices | Create a messaging service
[**CreateOrUpdateMessagingService**](MessagingServiceApi.md#CreateOrUpdateMessagingService) | **Put** /v1/services/messagingServices | Update messaging service
[**DeleteMessagingService**](MessagingServiceApi.md#DeleteMessagingService) | **Delete** /v1/services/messagingServices/{id} | Delete a messaging service
[**GetMessagingServiceByFQN**](MessagingServiceApi.md#GetMessagingServiceByFQN) | **Get** /v1/services/messagingServices/name/{name} | Get messaging service by name
[**GetMessagingServiceByID**](MessagingServiceApi.md#GetMessagingServiceByID) | **Get** /v1/services/messagingServices/{id} | Get a messaging service
[**GetSpecificMessagingServiceVersion**](MessagingServiceApi.md#GetSpecificMessagingServiceVersion) | **Get** /v1/services/messagingServices/{id}/versions/{version} | Get a version of the messaging service
[**ListAllMessagingServiceVersion**](MessagingServiceApi.md#ListAllMessagingServiceVersion) | **Get** /v1/services/messagingServices/{id}/versions | List messaging service versions
[**ListMessagingService**](MessagingServiceApi.md#ListMessagingService) | **Get** /v1/services/messagingServices | List messaging services



## CreateMessagingService

> MessagingService CreateMessagingService(ctx).CreateMessagingService(createMessagingService).Execute()

Create a messaging service



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
    createMessagingService := *openapiclient.NewCreateMessagingService(*openapiclient.NewMessagingConnection(), "Name_example", "ServiceType_example") // CreateMessagingService |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagingServiceApi.CreateMessagingService(context.Background()).CreateMessagingService(createMessagingService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagingServiceApi.CreateMessagingService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMessagingService`: MessagingService
    fmt.Fprintf(os.Stdout, "Response from `MessagingServiceApi.CreateMessagingService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessagingServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMessagingService** | [**CreateMessagingService**](CreateMessagingService.md) |  | 

### Return type

[**MessagingService**](MessagingService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateMessagingService

> MessagingService CreateOrUpdateMessagingService(ctx, id).CreateMessagingService(createMessagingService).Execute()

Update messaging service



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
    id := "id_example" // string | Id of the messaging service
    createMessagingService := *openapiclient.NewCreateMessagingService(*openapiclient.NewMessagingConnection(), "Name_example", "ServiceType_example") // CreateMessagingService |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagingServiceApi.CreateOrUpdateMessagingService(context.Background(), id).CreateMessagingService(createMessagingService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagingServiceApi.CreateOrUpdateMessagingService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateMessagingService`: MessagingService
    fmt.Fprintf(os.Stdout, "Response from `MessagingServiceApi.CreateOrUpdateMessagingService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the messaging service | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateMessagingServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMessagingService** | [**CreateMessagingService**](CreateMessagingService.md) |  | 

### Return type

[**MessagingService**](MessagingService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMessagingService

> DeleteMessagingService(ctx, id).Recursive(recursive).HardDelete(hardDelete).Execute()

Delete a messaging service



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
    id := "id_example" // string | Id of the messaging service
    recursive := true // bool | Recursively delete this entity and it's children. (Default `false`) (optional) (default to false)
    hardDelete := true // bool | Hard delete the entity. (Default = `false`) (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagingServiceApi.DeleteMessagingService(context.Background(), id).Recursive(recursive).HardDelete(hardDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagingServiceApi.DeleteMessagingService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the messaging service | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMessagingServiceRequest struct via the builder pattern


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


## GetMessagingServiceByFQN

> MessagingService GetMessagingServiceByFQN(ctx, name).Fields(fields).Include(include).Execute()

Get messaging service by name



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
    resp, r, err := apiClient.MessagingServiceApi.GetMessagingServiceByFQN(context.Background(), name).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagingServiceApi.GetMessagingServiceByFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessagingServiceByFQN`: MessagingService
    fmt.Fprintf(os.Stdout, "Response from `MessagingServiceApi.GetMessagingServiceByFQN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessagingServiceByFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**MessagingService**](MessagingService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessagingServiceByID

> MessagingService GetMessagingServiceByID(ctx, id).Fields(fields).Include(include).Execute()

Get a messaging service



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
    resp, r, err := apiClient.MessagingServiceApi.GetMessagingServiceByID(context.Background(), id).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagingServiceApi.GetMessagingServiceByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessagingServiceByID`: MessagingService
    fmt.Fprintf(os.Stdout, "Response from `MessagingServiceApi.GetMessagingServiceByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessagingServiceByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**MessagingService**](MessagingService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificMessagingServiceVersion

> MessagingService GetSpecificMessagingServiceVersion(ctx, id, version).Execute()

Get a version of the messaging service



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
    id := "id_example" // string | messaging service Id
    version := "0.1 or 1.1" // string | messaging service version number in the form `major`.`minor`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagingServiceApi.GetSpecificMessagingServiceVersion(context.Background(), id, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagingServiceApi.GetSpecificMessagingServiceVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificMessagingServiceVersion`: MessagingService
    fmt.Fprintf(os.Stdout, "Response from `MessagingServiceApi.GetSpecificMessagingServiceVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | messaging service Id | 
**version** | **string** | messaging service version number in the form &#x60;major&#x60;.&#x60;minor&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificMessagingServiceVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MessagingService**](MessagingService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllMessagingServiceVersion

> EntityHistory ListAllMessagingServiceVersion(ctx, id).Execute()

List messaging service versions



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
    id := "id_example" // string | messaging service Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagingServiceApi.ListAllMessagingServiceVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagingServiceApi.ListAllMessagingServiceVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllMessagingServiceVersion`: EntityHistory
    fmt.Fprintf(os.Stdout, "Response from `MessagingServiceApi.ListAllMessagingServiceVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | messaging service Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllMessagingServiceVersionRequest struct via the builder pattern


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


## ListMessagingService

> MessagingServiceList ListMessagingService(ctx).Fields(fields).Limit(limit).Before(before).After(after).Include(include).Execute()

List messaging services



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
    limit := int32(56) // int32 | Limit number services returned. (1 to 1000000, default 10) (optional) (default to 10)
    before := "before_example" // string | Returns list of services before this cursor (optional)
    after := "after_example" // string | Returns list of services after this cursor (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagingServiceApi.ListMessagingService(context.Background()).Fields(fields).Limit(limit).Before(before).After(after).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagingServiceApi.ListMessagingService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMessagingService`: MessagingServiceList
    fmt.Fprintf(os.Stdout, "Response from `MessagingServiceApi.ListMessagingService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMessagingServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Fields requested in the returned resource | 
 **limit** | **int32** | Limit number services returned. (1 to 1000000, default 10) | [default to 10]
 **before** | **string** | Returns list of services before this cursor | 
 **after** | **string** | Returns list of services after this cursor | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**MessagingServiceList**](MessagingServiceList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

