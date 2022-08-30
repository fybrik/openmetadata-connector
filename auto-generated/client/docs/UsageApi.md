# \UsageApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEntityUsageByFQN**](UsageApi.md#GetEntityUsageByFQN) | **Get** /v1/usage/{entity}/name/{fqn} | Get usage by name
[**GetEntityUsageByID**](UsageApi.md#GetEntityUsageByID) | **Get** /v1/usage/{entity}/{id} | Get usage
[**ReportEntityUsageWithFQN**](UsageApi.md#ReportEntityUsageWithFQN) | **Post** /v1/usage/{entity}/name/{fqn} | Report usage by name
[**ReportEntityUsageWithID**](UsageApi.md#ReportEntityUsageWithID) | **Post** /v1/usage/{entity}/{id} | Report usage



## GetEntityUsageByFQN

> EntityUsage GetEntityUsageByFQN(ctx, entity, fqn).Days(days).Date(date).Execute()

Get usage by name



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
    entity := "table, report, metrics, or dashboard" // string | Entity type for which usage is requested
    fqn := "fqn_example" // string | Fully qualified name of the entity that uniquely identifies an entity
    days := int32(56) // int32 | Usage for number of days going back from the given date (default=1, min=1, max=30) (optional)
    date := "date_example" // string | Usage for number of days going back from this date in ISO 8601 format (default = currentDate) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetEntityUsageByFQN(context.Background(), entity, fqn).Days(days).Date(date).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetEntityUsageByFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntityUsageByFQN`: EntityUsage
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetEntityUsageByFQN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entity** | **string** | Entity type for which usage is requested | 
**fqn** | **string** | Fully qualified name of the entity that uniquely identifies an entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityUsageByFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **days** | **int32** | Usage for number of days going back from the given date (default&#x3D;1, min&#x3D;1, max&#x3D;30) | 
 **date** | **string** | Usage for number of days going back from this date in ISO 8601 format (default &#x3D; currentDate) | 

### Return type

[**EntityUsage**](EntityUsage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntityUsageByID

> EntityUsage GetEntityUsageByID(ctx, entity, id).Days(days).Date(date).Execute()

Get usage



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
    entity := "table, report, metrics, or dashboard" // string | Entity type for which usage is requested
    id := "id_example" // string | Entity id
    days := int32(56) // int32 | Usage for number of days going back from the given date (default=1, min=1, max=30) (optional)
    date := "date_example" // string | Usage for number of days going back from this date in ISO 8601 format. (default = currentDate) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetEntityUsageByID(context.Background(), entity, id).Days(days).Date(date).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetEntityUsageByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntityUsageByID`: EntityUsage
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetEntityUsageByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entity** | **string** | Entity type for which usage is requested | 
**id** | **string** | Entity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityUsageByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **days** | **int32** | Usage for number of days going back from the given date (default&#x3D;1, min&#x3D;1, max&#x3D;30) | 
 **date** | **string** | Usage for number of days going back from this date in ISO 8601 format. (default &#x3D; currentDate) | 

### Return type

[**EntityUsage**](EntityUsage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportEntityUsageWithFQN

> EntityUsage ReportEntityUsageWithFQN(ctx, entity, fqn).DailyCount(dailyCount).Execute()

Report usage by name



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
    entity := "table, report, metrics, or dashboard" // string | Entity type for which usage is reported
    fqn := "fqn_example" // string | Fully qualified name of the entity that uniquely identifies an entity
    dailyCount := *openapiclient.NewDailyCount(int32(123), "Date_example") // DailyCount | Usage information a given date (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.ReportEntityUsageWithFQN(context.Background(), entity, fqn).DailyCount(dailyCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.ReportEntityUsageWithFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportEntityUsageWithFQN`: EntityUsage
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.ReportEntityUsageWithFQN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entity** | **string** | Entity type for which usage is reported | 
**fqn** | **string** | Fully qualified name of the entity that uniquely identifies an entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportEntityUsageWithFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dailyCount** | [**DailyCount**](DailyCount.md) | Usage information a given date | 

### Return type

[**EntityUsage**](EntityUsage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportEntityUsageWithID

> EntityUsage ReportEntityUsageWithID(ctx, entity, id).DailyCount(dailyCount).Execute()

Report usage



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
    entity := "table, report, metrics, or dashboard" // string | Entity type for which usage is reported
    id := "id_example" // string | Entity id
    dailyCount := *openapiclient.NewDailyCount(int32(123), "Date_example") // DailyCount | Usage information a given date (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.ReportEntityUsageWithID(context.Background(), entity, id).DailyCount(dailyCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.ReportEntityUsageWithID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportEntityUsageWithID`: EntityUsage
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.ReportEntityUsageWithID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entity** | **string** | Entity type for which usage is reported | 
**id** | **string** | Entity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportEntityUsageWithIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dailyCount** | [**DailyCount**](DailyCount.md) | Usage information a given date | 

### Return type

[**EntityUsage**](EntityUsage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

