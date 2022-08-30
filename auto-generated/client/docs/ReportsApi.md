# \ReportsApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateReport**](ReportsApi.md#CreateOrUpdateReport) | **Put** /v1/reports | Create or update a report
[**GetReportByFQN**](ReportsApi.md#GetReportByFQN) | **Post** /v1/reports | Create a report
[**GetReportByID**](ReportsApi.md#GetReportByID) | **Get** /v1/reports/{id} | Get a report
[**ListReports**](ReportsApi.md#ListReports) | **Get** /v1/reports | List reports



## CreateOrUpdateReport

> Report CreateOrUpdateReport(ctx).Report(report).Execute()

Create or update a report



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
    report := *openapiclient.NewReport("Id_example", "Name_example", *openapiclient.NewEntityReference("Id_example", "Type_example")) // Report |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.CreateOrUpdateReport(context.Background()).Report(report).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.CreateOrUpdateReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateReport`: Report
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.CreateOrUpdateReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **report** | [**Report**](Report.md) |  | 

### Return type

[**Report**](Report.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportByFQN

> Report GetReportByFQN(ctx).Report(report).Execute()

Create a report



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
    report := *openapiclient.NewReport("Id_example", "Name_example", *openapiclient.NewEntityReference("Id_example", "Type_example")) // Report |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.GetReportByFQN(context.Background()).Report(report).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetReportByFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportByFQN`: Report
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetReportByFQN`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportByFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **report** | [**Report**](Report.md) |  | 

### Return type

[**Report**](Report.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportByID

> Report GetReportByID(ctx, id).Fields(fields).Include(include).Execute()

Get a report



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
    fields := "owner,usageSummary" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.GetReportByID(context.Background(), id).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetReportByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportByID`: Report
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetReportByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**Report**](Report.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReports

> ReportList ListReports(ctx).Fields(fields).Execute()

List reports



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
    fields := "owner,usageSummary" // string | Fields requested in the returned resource (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ListReports(context.Background()).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ListReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListReports`: ReportList
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ListReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Fields requested in the returned resource | 

### Return type

[**ReportList**](ReportList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

