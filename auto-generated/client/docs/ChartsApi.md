# \ChartsApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFollowerToChart**](ChartsApi.md#AddFollowerToChart) | **Put** /v1/charts/{id}/followers | Add a follower
[**CreateChart**](ChartsApi.md#CreateChart) | **Post** /v1/charts | Create a chart
[**CreateOrUpdateChart**](ChartsApi.md#CreateOrUpdateChart) | **Put** /v1/charts | Create or update chart
[**DeleteChart**](ChartsApi.md#DeleteChart) | **Delete** /v1/charts/{id} | Delete a Chart
[**DeleteFollowerFromChart**](ChartsApi.md#DeleteFollowerFromChart) | **Delete** /v1/charts/{id}/followers/{userId} | Remove a follower
[**GetChartByFQN**](ChartsApi.md#GetChartByFQN) | **Get** /v1/charts/name/{fqn} | Get a chart by name
[**GetChartByID**](ChartsApi.md#GetChartByID) | **Get** /v1/charts/{id} | Get a Chart
[**GetSpecificChartVersion**](ChartsApi.md#GetSpecificChartVersion) | **Get** /v1/charts/{id}/versions/{version} | Get a version of the chart
[**ListAllChartVersions**](ChartsApi.md#ListAllChartVersions) | **Get** /v1/charts/{id}/versions | List chart versions
[**ListCharts**](ChartsApi.md#ListCharts) | **Get** /v1/charts | List charts
[**PatchChart**](ChartsApi.md#PatchChart) | **Patch** /v1/charts/{id} | Update a chart



## AddFollowerToChart

> AddFollowerToChart(ctx, id).Body(body).Execute()

Add a follower



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
    id := "id_example" // string | Id of the chart
    body := "body_example" // string | Id of the user to be added as follower (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartsApi.AddFollowerToChart(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartsApi.AddFollowerToChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the chart | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFollowerToChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Id of the user to be added as follower | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChart

> Chart CreateChart(ctx).CreateChart(createChart).Execute()

Create a chart



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
    createChart := *openapiclient.NewCreateChart("Name_example", *openapiclient.NewEntityReference("Id_example", "Type_example")) // CreateChart |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartsApi.CreateChart(context.Background()).CreateChart(createChart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartsApi.CreateChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateChart`: Chart
    fmt.Fprintf(os.Stdout, "Response from `ChartsApi.CreateChart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createChart** | [**CreateChart**](CreateChart.md) |  | 

### Return type

[**Chart**](Chart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateChart

> Chart CreateOrUpdateChart(ctx).CreateChart(createChart).Execute()

Create or update chart



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
    createChart := *openapiclient.NewCreateChart("Name_example", *openapiclient.NewEntityReference("Id_example", "Type_example")) // CreateChart |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartsApi.CreateOrUpdateChart(context.Background()).CreateChart(createChart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartsApi.CreateOrUpdateChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateChart`: Chart
    fmt.Fprintf(os.Stdout, "Response from `ChartsApi.CreateOrUpdateChart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createChart** | [**CreateChart**](CreateChart.md) |  | 

### Return type

[**Chart**](Chart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChart

> DeleteChart(ctx, id).HardDelete(hardDelete).Execute()

Delete a Chart



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
    id := "id_example" // string | Chart Id
    hardDelete := true // bool | Hard delete the entity. (Default = `false`) (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartsApi.DeleteChart(context.Background(), id).HardDelete(hardDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartsApi.DeleteChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Chart Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## DeleteFollowerFromChart

> DeleteFollowerFromChart(ctx, id, userId).Execute()

Remove a follower



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
    id := "id_example" // string | Id of the chart
    userId := "userId_example" // string | Id of the user being removed as follower

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartsApi.DeleteFollowerFromChart(context.Background(), id, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartsApi.DeleteFollowerFromChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the chart | 
**userId** | **string** | Id of the user being removed as follower | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFollowerFromChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChartByFQN

> Chart GetChartByFQN(ctx, fqn).Fields(fields).Include(include).Execute()

Get a chart by name



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
    fqn := "fqn_example" // string | 
    fields := "owner,followers,tags" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartsApi.GetChartByFQN(context.Background(), fqn).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartsApi.GetChartByFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChartByFQN`: Chart
    fmt.Fprintf(os.Stdout, "Response from `ChartsApi.GetChartByFQN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqn** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChartByFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**Chart**](Chart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChartByID

> Chart GetChartByID(ctx, id).Fields(fields).Include(include).Execute()

Get a Chart



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
    fields := "owner,followers,tags" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartsApi.GetChartByID(context.Background(), id).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartsApi.GetChartByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChartByID`: Chart
    fmt.Fprintf(os.Stdout, "Response from `ChartsApi.GetChartByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChartByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**Chart**](Chart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificChartVersion

> Chart GetSpecificChartVersion(ctx, id, version).Execute()

Get a version of the chart



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
    id := "id_example" // string | Chart Id
    version := "0.1 or 1.1" // string | Chart version number in the form `major`.`minor`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartsApi.GetSpecificChartVersion(context.Background(), id, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartsApi.GetSpecificChartVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificChartVersion`: Chart
    fmt.Fprintf(os.Stdout, "Response from `ChartsApi.GetSpecificChartVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Chart Id | 
**version** | **string** | Chart version number in the form &#x60;major&#x60;.&#x60;minor&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificChartVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Chart**](Chart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllChartVersions

> EntityHistory ListAllChartVersions(ctx, id).Execute()

List chart versions



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
    id := "id_example" // string | Chart Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartsApi.ListAllChartVersions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartsApi.ListAllChartVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllChartVersions`: EntityHistory
    fmt.Fprintf(os.Stdout, "Response from `ChartsApi.ListAllChartVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Chart Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllChartVersionsRequest struct via the builder pattern


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


## ListCharts

> ChartList ListCharts(ctx).Fields(fields).Service(service).Limit(limit).Before(before).After(after).Include(include).Execute()

List charts



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
    fields := "owner,followers,tags" // string | Fields requested in the returned resource (optional)
    service := "superset" // string | Filter charts by service name (optional)
    limit := int32(56) // int32 | Limit the number charts returned. (1 to 1000000, default = 10) (optional) (default to 10)
    before := "before_example" // string | Returns list of charts before this cursor (optional)
    after := "after_example" // string | Returns list of charts after this cursor (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartsApi.ListCharts(context.Background()).Fields(fields).Service(service).Limit(limit).Before(before).After(after).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartsApi.ListCharts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCharts`: ChartList
    fmt.Fprintf(os.Stdout, "Response from `ChartsApi.ListCharts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListChartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Fields requested in the returned resource | 
 **service** | **string** | Filter charts by service name | 
 **limit** | **int32** | Limit the number charts returned. (1 to 1000000, default &#x3D; 10) | [default to 10]
 **before** | **string** | Returns list of charts before this cursor | 
 **after** | **string** | Returns list of charts after this cursor | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**ChartList**](ChartList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchChart

> PatchChart(ctx, id).Body(body).Execute()

Update a chart



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
    body := map[string]interface{}{ ... } // map[string]interface{} | JsonPatch with array of operations (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartsApi.PatchChart(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartsApi.PatchChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | JsonPatch with array of operations | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

