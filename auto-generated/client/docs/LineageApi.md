# \LineageApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLineageEdge**](LineageApi.md#AddLineageEdge) | **Put** /v1/lineage | Add a lineage edge
[**DeleteLineageEdge**](LineageApi.md#DeleteLineageEdge) | **Delete** /v1/lineage/{fromEntity}/{fromId}/{toEntity}/{toId} | Delete a lineage edge
[**GetLineage**](LineageApi.md#GetLineage) | **Get** /v1/lineage/{entity}/{id} | Get lineage
[**GetLineageByFQN**](LineageApi.md#GetLineageByFQN) | **Get** /v1/lineage/{entity}/name/{fqn} | Get lineage by name



## AddLineageEdge

> AddLineageEdge(ctx).AddLineage(addLineage).Execute()

Add a lineage edge



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
    addLineage := *openapiclient.NewAddLineage(*openapiclient.NewEntitiesEdge(*openapiclient.NewEntityReference("Id_example", "Type_example"), *openapiclient.NewEntityReference("Id_example", "Type_example"))) // AddLineage |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LineageApi.AddLineageEdge(context.Background()).AddLineage(addLineage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineageApi.AddLineageEdge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLineageEdgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addLineage** | [**AddLineage**](AddLineage.md) |  | 

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


## DeleteLineageEdge

> DeleteLineageEdge(ctx, fromEntity, fromId, toEntity, toId).Execute()

Delete a lineage edge



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
    fromEntity := "table, report, metrics, or dashboard" // string | Entity type of upstream entity of the edge
    fromId := "fromId_example" // string | Entity id
    toEntity := "table, report, metrics, or dashboard" // string | Entity type for downstream entity of the edge
    toId := "toId_example" // string | Entity id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LineageApi.DeleteLineageEdge(context.Background(), fromEntity, fromId, toEntity, toId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineageApi.DeleteLineageEdge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromEntity** | **string** | Entity type of upstream entity of the edge | 
**fromId** | **string** | Entity id | 
**toEntity** | **string** | Entity type for downstream entity of the edge | 
**toId** | **string** | Entity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLineageEdgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





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


## GetLineage

> EntityLineage GetLineage(ctx, entity, id).UpstreamDepth(upstreamDepth).DownstreamDepth(downstreamDepth).Execute()

Get lineage



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
    entity := "table, report, metrics, or dashboard" // string | Entity type for which lineage is requested
    id := "id_example" // string | Entity id
    upstreamDepth := int32(56) // int32 | Upstream depth of lineage (default=1, min=0, max=3) (optional) (default to 1)
    downstreamDepth := int32(56) // int32 | Upstream depth of lineage (default=1, min=0, max=3) (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LineageApi.GetLineage(context.Background(), entity, id).UpstreamDepth(upstreamDepth).DownstreamDepth(downstreamDepth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineageApi.GetLineage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLineage`: EntityLineage
    fmt.Fprintf(os.Stdout, "Response from `LineageApi.GetLineage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entity** | **string** | Entity type for which lineage is requested | 
**id** | **string** | Entity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLineageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upstreamDepth** | **int32** | Upstream depth of lineage (default&#x3D;1, min&#x3D;0, max&#x3D;3) | [default to 1]
 **downstreamDepth** | **int32** | Upstream depth of lineage (default&#x3D;1, min&#x3D;0, max&#x3D;3) | [default to 1]

### Return type

[**EntityLineage**](EntityLineage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLineageByFQN

> EntityLineage GetLineageByFQN(ctx, entity, fqn).UpstreamDepth(upstreamDepth).DownstreamDepth(downstreamDepth).Execute()

Get lineage by name



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
    entity := "table, report, metrics, or dashboard" // string | Entity type for which lineage is requested
    fqn := "fqn_example" // string | Fully qualified name of the entity that uniquely identifies an entity
    upstreamDepth := int32(56) // int32 | Upstream depth of lineage (default=1, min=0, max=3) (optional) (default to 1)
    downstreamDepth := int32(56) // int32 | Upstream depth of lineage (default=1, min=0, max=3) (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LineageApi.GetLineageByFQN(context.Background(), entity, fqn).UpstreamDepth(upstreamDepth).DownstreamDepth(downstreamDepth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineageApi.GetLineageByFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLineageByFQN`: EntityLineage
    fmt.Fprintf(os.Stdout, "Response from `LineageApi.GetLineageByFQN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entity** | **string** | Entity type for which lineage is requested | 
**fqn** | **string** | Fully qualified name of the entity that uniquely identifies an entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLineageByFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upstreamDepth** | **int32** | Upstream depth of lineage (default&#x3D;1, min&#x3D;0, max&#x3D;3) | [default to 1]
 **downstreamDepth** | **int32** | Upstream depth of lineage (default&#x3D;1, min&#x3D;0, max&#x3D;3) | [default to 1]

### Return type

[**EntityLineage**](EntityLineage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

