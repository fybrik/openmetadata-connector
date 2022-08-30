# \MetricsApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMetric**](MetricsApi.md#CreateMetric) | **Post** /v1/metrics | Create a metric
[**CreateOrUpdateMetric**](MetricsApi.md#CreateOrUpdateMetric) | **Put** /v1/metrics | Create or update a metric
[**GetMetricByID**](MetricsApi.md#GetMetricByID) | **Get** /v1/metrics/{id} | Get a metric
[**ListMetrics**](MetricsApi.md#ListMetrics) | **Get** /v1/metrics | List metrics



## CreateMetric

> Metrics CreateMetric(ctx).Metrics(metrics).Execute()

Create a metric



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
    metrics := *openapiclient.NewMetrics("Id_example", "Name_example", *openapiclient.NewEntityReference("Id_example", "Type_example")) // Metrics |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.CreateMetric(context.Background()).Metrics(metrics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.CreateMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMetric`: Metrics
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.CreateMetric`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metrics** | [**Metrics**](Metrics.md) |  | 

### Return type

[**Metrics**](Metrics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateMetric

> Metrics CreateOrUpdateMetric(ctx).Metrics(metrics).Execute()

Create or update a metric



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
    metrics := *openapiclient.NewMetrics("Id_example", "Name_example", *openapiclient.NewEntityReference("Id_example", "Type_example")) // Metrics |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.CreateOrUpdateMetric(context.Background()).Metrics(metrics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.CreateOrUpdateMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateMetric`: Metrics
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.CreateOrUpdateMetric`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metrics** | [**Metrics**](Metrics.md) |  | 

### Return type

[**Metrics**](Metrics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricByID

> Metrics GetMetricByID(ctx, id).Fields(fields).Include(include).Execute()

Get a metric



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
    resp, r, err := apiClient.MetricsApi.GetMetricByID(context.Background(), id).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetMetricByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricByID`: Metrics
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.GetMetricByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**Metrics**](Metrics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetrics

> MetricsList ListMetrics(ctx).Fields(fields).Limit(limit).Before(before).After(after).Execute()

List metrics



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
    limit := int32(56) // int32 |  (optional) (default to 10)
    before := "before_example" // string | Returns list of metrics before this cursor (optional)
    after := "after_example" // string | Returns list of metrics after this cursor (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.ListMetrics(context.Background()).Fields(fields).Limit(limit).Before(before).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMetrics`: MetricsList
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.ListMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Fields requested in the returned resource | 
 **limit** | **int32** |  | [default to 10]
 **before** | **string** | Returns list of metrics before this cursor | 
 **after** | **string** | Returns list of metrics after this cursor | 

### Return type

[**MetricsList**](MetricsList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

