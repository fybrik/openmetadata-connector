# \EventsApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListChangeEvents**](EventsApi.md#ListChangeEvents) | **Get** /v1/events | Get change events



## ListChangeEvents

> ChangeEventList ListChangeEvents(ctx).Timestamp(timestamp).EntityCreated(entityCreated).EntityUpdated(entityUpdated).EntityDeleted(entityDeleted).Execute()

Get change events



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
    timestamp := int64(1426349294842) // int64 | Events starting from this unix timestamp in milliseconds
    entityCreated := "table,dashboard,..." // string | List of comma separated entities requested for `entityCreated` event. When set to `*` all entities will be returned (optional)
    entityUpdated := "table,dashboard,..." // string | List of comma separated entities requested for `entityCreated` event. When set to `*` all entities will be returned (optional)
    entityDeleted := "table,dashboard,..." // string | List of comma separated entities requested for `entityCreated` event. When set to `*` all entities will be returned (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.ListChangeEvents(context.Background()).Timestamp(timestamp).EntityCreated(entityCreated).EntityUpdated(entityUpdated).EntityDeleted(entityDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListChangeEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListChangeEvents`: ChangeEventList
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ListChangeEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListChangeEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | Events starting from this unix timestamp in milliseconds | 
 **entityCreated** | **string** | List of comma separated entities requested for &#x60;entityCreated&#x60; event. When set to &#x60;*&#x60; all entities will be returned | 
 **entityUpdated** | **string** | List of comma separated entities requested for &#x60;entityCreated&#x60; event. When set to &#x60;*&#x60; all entities will be returned | 
 **entityDeleted** | **string** | List of comma separated entities requested for &#x60;entityCreated&#x60; event. When set to &#x60;*&#x60; all entities will be returned | 

### Return type

[**ChangeEventList**](ChangeEventList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

