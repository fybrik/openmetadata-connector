# \IngestionPipelinesApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckRestAirflowStatus**](IngestionPipelinesApi.md#CheckRestAirflowStatus) | **Get** /v1/services/ingestionPipelines/status | Check the Airflow REST status
[**CreateIngestionPipeline**](IngestionPipelinesApi.md#CreateIngestionPipeline) | **Post** /v1/services/ingestionPipelines | Create a Ingestion Pipeline
[**CreateOrUpdateIngestionPipeline**](IngestionPipelinesApi.md#CreateOrUpdateIngestionPipeline) | **Put** /v1/services/ingestionPipelines | Create or update a IngestionPipeline
[**DeleteIngestionPipeline**](IngestionPipelinesApi.md#DeleteIngestionPipeline) | **Delete** /v1/services/ingestionPipelines/{id} | Delete a Ingestion
[**DeployIngestion**](IngestionPipelinesApi.md#DeployIngestion) | **Post** /v1/services/ingestionPipelines/deploy/{id} | Deploy a ingestion pipeline run
[**GetIngestionPipelineByID**](IngestionPipelinesApi.md#GetIngestionPipelineByID) | **Get** /v1/services/ingestionPipelines/{id} | Get a IngestionPipeline
[**GetLastIngestionLogs**](IngestionPipelinesApi.md#GetLastIngestionLogs) | **Get** /v1/services/ingestionPipelines/logs/{id}/last | Retrieve all logs from last ingestion pipeline run
[**GetSpecificIngestionPipelineByFQN**](IngestionPipelinesApi.md#GetSpecificIngestionPipelineByFQN) | **Get** /v1/services/ingestionPipelines/name/{fqn} | Get a IngestionPipeline by name
[**GetSpecificIngestionPipelineVersion**](IngestionPipelinesApi.md#GetSpecificIngestionPipelineVersion) | **Get** /v1/services/ingestionPipelines/{id}/versions/{version} | Get a version of the IngestionPipeline
[**ListAllIngestionPipelineVersion**](IngestionPipelinesApi.md#ListAllIngestionPipelineVersion) | **Get** /v1/services/ingestionPipelines/{id}/versions | List ingestion workflow versions
[**ListIngestionPipelines**](IngestionPipelinesApi.md#ListIngestionPipelines) | **Get** /v1/services/ingestionPipelines | List Ingestion Pipelines for Metadata Operations
[**PatchIngestionPipeline**](IngestionPipelinesApi.md#PatchIngestionPipeline) | **Patch** /v1/services/ingestionPipelines/{id} | Update a IngestionPipeline
[**TestConnection**](IngestionPipelinesApi.md#TestConnection) | **Post** /v1/services/ingestionPipelines/testConnection | Test Connection of a Service
[**ToggleIngestionPipelineEnabled**](IngestionPipelinesApi.md#ToggleIngestionPipelineEnabled) | **Post** /v1/services/ingestionPipelines/toggleIngestion/{id} | Set an Ingestion pipeline either as Enabled or Disabled
[**TriggerIngestionPipelineRun**](IngestionPipelinesApi.md#TriggerIngestionPipelineRun) | **Post** /v1/services/ingestionPipelines/trigger/{id} | Trigger a ingestion pipeline run



## CheckRestAirflowStatus

> CheckRestAirflowStatus(ctx).Execute()

Check the Airflow REST status



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IngestionPipelinesApi.CheckRestAirflowStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestionPipelinesApi.CheckRestAirflowStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckRestAirflowStatusRequest struct via the builder pattern


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


## CreateIngestionPipeline

> IngestionPipeline CreateIngestionPipeline(ctx).CreateIngestionPipeline(createIngestionPipeline).Execute()

Create a Ingestion Pipeline



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
    createIngestionPipeline := *openapiclient.NewCreateIngestionPipeline(*openapiclient.NewAirflowConfig(), "Name_example", "PipelineType_example", *openapiclient.NewEntityReference("Id_example", "Type_example"), *openapiclient.NewSourceConfig()) // CreateIngestionPipeline |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IngestionPipelinesApi.CreateIngestionPipeline(context.Background()).CreateIngestionPipeline(createIngestionPipeline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestionPipelinesApi.CreateIngestionPipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIngestionPipeline`: IngestionPipeline
    fmt.Fprintf(os.Stdout, "Response from `IngestionPipelinesApi.CreateIngestionPipeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIngestionPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIngestionPipeline** | [**CreateIngestionPipeline**](CreateIngestionPipeline.md) |  | 

### Return type

[**IngestionPipeline**](IngestionPipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateIngestionPipeline

> IngestionPipeline CreateOrUpdateIngestionPipeline(ctx).CreateIngestionPipeline(createIngestionPipeline).Execute()

Create or update a IngestionPipeline



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
    createIngestionPipeline := *openapiclient.NewCreateIngestionPipeline(*openapiclient.NewAirflowConfig(), "Name_example", "PipelineType_example", *openapiclient.NewEntityReference("Id_example", "Type_example"), *openapiclient.NewSourceConfig()) // CreateIngestionPipeline |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IngestionPipelinesApi.CreateOrUpdateIngestionPipeline(context.Background()).CreateIngestionPipeline(createIngestionPipeline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestionPipelinesApi.CreateOrUpdateIngestionPipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateIngestionPipeline`: IngestionPipeline
    fmt.Fprintf(os.Stdout, "Response from `IngestionPipelinesApi.CreateOrUpdateIngestionPipeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateIngestionPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIngestionPipeline** | [**CreateIngestionPipeline**](CreateIngestionPipeline.md) |  | 

### Return type

[**IngestionPipeline**](IngestionPipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIngestionPipeline

> DeleteIngestionPipeline(ctx, id).HardDelete(hardDelete).Execute()

Delete a Ingestion



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
    id := "id_example" // string | Pipeline Id
    hardDelete := true // bool | Hard delete the entity. (Default = `false`) (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IngestionPipelinesApi.DeleteIngestionPipeline(context.Background(), id).HardDelete(hardDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestionPipelinesApi.DeleteIngestionPipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Pipeline Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIngestionPipelineRequest struct via the builder pattern


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


## DeployIngestion

> IngestionPipeline DeployIngestion(ctx, id).Execute()

Deploy a ingestion pipeline run



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IngestionPipelinesApi.DeployIngestion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestionPipelinesApi.DeployIngestion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployIngestion`: IngestionPipeline
    fmt.Fprintf(os.Stdout, "Response from `IngestionPipelinesApi.DeployIngestion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployIngestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IngestionPipeline**](IngestionPipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIngestionPipelineByID

> IngestionPipeline GetIngestionPipelineByID(ctx, id).Fields(fields).Include(include).Execute()

Get a IngestionPipeline



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
    resp, r, err := apiClient.IngestionPipelinesApi.GetIngestionPipelineByID(context.Background(), id).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestionPipelinesApi.GetIngestionPipelineByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIngestionPipelineByID`: IngestionPipeline
    fmt.Fprintf(os.Stdout, "Response from `IngestionPipelinesApi.GetIngestionPipelineByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIngestionPipelineByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**IngestionPipeline**](IngestionPipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastIngestionLogs

> GetLastIngestionLogs(ctx, id).Execute()

Retrieve all logs from last ingestion pipeline run



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
    id := "id_example" // string | Pipeline Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IngestionPipelinesApi.GetLastIngestionLogs(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestionPipelinesApi.GetLastIngestionLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Pipeline Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastIngestionLogsRequest struct via the builder pattern


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


## GetSpecificIngestionPipelineByFQN

> IngestionPipeline GetSpecificIngestionPipelineByFQN(ctx, fqn).Fields(fields).Include(include).Execute()

Get a IngestionPipeline by name



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
    fields := "owner" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IngestionPipelinesApi.GetSpecificIngestionPipelineByFQN(context.Background(), fqn).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestionPipelinesApi.GetSpecificIngestionPipelineByFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificIngestionPipelineByFQN`: IngestionPipeline
    fmt.Fprintf(os.Stdout, "Response from `IngestionPipelinesApi.GetSpecificIngestionPipelineByFQN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqn** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificIngestionPipelineByFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**IngestionPipeline**](IngestionPipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificIngestionPipelineVersion

> IngestionPipeline GetSpecificIngestionPipelineVersion(ctx, id, version).Execute()

Get a version of the IngestionPipeline



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
    id := "id_example" // string | Ingestion Id
    version := "0.1 or 1.1" // string | Ingestion version number in the form `major`.`minor`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IngestionPipelinesApi.GetSpecificIngestionPipelineVersion(context.Background(), id, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestionPipelinesApi.GetSpecificIngestionPipelineVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificIngestionPipelineVersion`: IngestionPipeline
    fmt.Fprintf(os.Stdout, "Response from `IngestionPipelinesApi.GetSpecificIngestionPipelineVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Ingestion Id | 
**version** | **string** | Ingestion version number in the form &#x60;major&#x60;.&#x60;minor&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificIngestionPipelineVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IngestionPipeline**](IngestionPipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllIngestionPipelineVersion

> EntityHistory ListAllIngestionPipelineVersion(ctx, id).Execute()

List ingestion workflow versions



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
    id := "id_example" // string | IngestionPipeline Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IngestionPipelinesApi.ListAllIngestionPipelineVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestionPipelinesApi.ListAllIngestionPipelineVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllIngestionPipelineVersion`: EntityHistory
    fmt.Fprintf(os.Stdout, "Response from `IngestionPipelinesApi.ListAllIngestionPipelineVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | IngestionPipeline Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllIngestionPipelineVersionRequest struct via the builder pattern


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


## ListIngestionPipelines

> IngestionPipeline ListIngestionPipelines(ctx).Fields(fields).Service(service).Limit(limit).Before(before).After(after).Include(include).Execute()

List Ingestion Pipelines for Metadata Operations



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
    service := "snowflakeWestCoast" // string | Filter airflow pipelines by service fully qualified name (optional)
    limit := int32(56) // int32 | Limit the number ingestion returned. (1 to 1000000, default = 10) (optional) (default to 10)
    before := "before_example" // string | Returns list of ingestion before this cursor (optional)
    after := "after_example" // string | Returns list of ingestion after this cursor (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IngestionPipelinesApi.ListIngestionPipelines(context.Background()).Fields(fields).Service(service).Limit(limit).Before(before).After(after).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestionPipelinesApi.ListIngestionPipelines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIngestionPipelines`: IngestionPipeline
    fmt.Fprintf(os.Stdout, "Response from `IngestionPipelinesApi.ListIngestionPipelines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIngestionPipelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Fields requested in the returned resource | 
 **service** | **string** | Filter airflow pipelines by service fully qualified name | 
 **limit** | **int32** | Limit the number ingestion returned. (1 to 1000000, default &#x3D; 10) | [default to 10]
 **before** | **string** | Returns list of ingestion before this cursor | 
 **after** | **string** | Returns list of ingestion after this cursor | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**IngestionPipeline**](IngestionPipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIngestionPipeline

> PatchIngestionPipeline(ctx, id).Body(body).Execute()

Update a IngestionPipeline



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
    resp, r, err := apiClient.IngestionPipelinesApi.PatchIngestionPipeline(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestionPipelinesApi.PatchIngestionPipeline``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchIngestionPipelineRequest struct via the builder pattern


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


## TestConnection

> TestConnection(ctx).TestServiceConnection(testServiceConnection).Execute()

Test Connection of a Service



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
    testServiceConnection := *openapiclient.NewTestServiceConnection() // TestServiceConnection |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IngestionPipelinesApi.TestConnection(context.Background()).TestServiceConnection(testServiceConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestionPipelinesApi.TestConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testServiceConnection** | [**TestServiceConnection**](TestServiceConnection.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleIngestionPipelineEnabled

> IngestionPipeline ToggleIngestionPipelineEnabled(ctx, id).Execute()

Set an Ingestion pipeline either as Enabled or Disabled



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IngestionPipelinesApi.ToggleIngestionPipelineEnabled(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestionPipelinesApi.ToggleIngestionPipelineEnabled``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ToggleIngestionPipelineEnabled`: IngestionPipeline
    fmt.Fprintf(os.Stdout, "Response from `IngestionPipelinesApi.ToggleIngestionPipelineEnabled`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleIngestionPipelineEnabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IngestionPipeline**](IngestionPipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerIngestionPipelineRun

> IngestionPipeline TriggerIngestionPipelineRun(ctx, id).Execute()

Trigger a ingestion pipeline run



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IngestionPipelinesApi.TriggerIngestionPipelineRun(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestionPipelinesApi.TriggerIngestionPipelineRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TriggerIngestionPipelineRun`: IngestionPipeline
    fmt.Fprintf(os.Stdout, "Response from `IngestionPipelinesApi.TriggerIngestionPipelineRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerIngestionPipelineRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IngestionPipeline**](IngestionPipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

