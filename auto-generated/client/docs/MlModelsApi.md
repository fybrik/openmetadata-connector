# \MlModelsApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFollower1**](MlModelsApi.md#AddFollower1) | **Put** /v1/mlmodels/{id}/followers | Add a follower
[**CreateMlModel**](MlModelsApi.md#CreateMlModel) | **Post** /v1/mlmodels | Create an ML Model
[**CreateOrUpdateMlModel**](MlModelsApi.md#CreateOrUpdateMlModel) | **Put** /v1/mlmodels | Create or update an ML Model
[**DeleteFollower2**](MlModelsApi.md#DeleteFollower2) | **Delete** /v1/mlmodels/{id}/followers/{userId} | Remove a follower
[**DeleteMlModel**](MlModelsApi.md#DeleteMlModel) | **Delete** /v1/mlmodels/{id} | Delete an ML Model
[**GetMlModelByFQN**](MlModelsApi.md#GetMlModelByFQN) | **Get** /v1/mlmodels/name/{fqn} | Get an ML Model by name
[**GetMlModelByID**](MlModelsApi.md#GetMlModelByID) | **Get** /v1/mlmodels/{id} | Get an ML Model
[**GetSpecificMlModelVersion**](MlModelsApi.md#GetSpecificMlModelVersion) | **Get** /v1/mlmodels/{id}/versions/{version} | Get a version of the ML Model
[**ListAllMlModelVersion**](MlModelsApi.md#ListAllMlModelVersion) | **Get** /v1/mlmodels/{id}/versions | List Ml Model versions
[**ListMlModels**](MlModelsApi.md#ListMlModels) | **Get** /v1/mlmodels | List ML Models
[**PatchMlModel**](MlModelsApi.md#PatchMlModel) | **Patch** /v1/mlmodels/{id} | Update an ML Model



## AddFollower1

> ChangeEvent AddFollower1(ctx, id).Body(body).Execute()

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
    id := "id_example" // string | Id of the model
    body := "body_example" // string | Id of the user to be added as follower (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MlModelsApi.AddFollower1(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelsApi.AddFollower1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddFollower1`: ChangeEvent
    fmt.Fprintf(os.Stdout, "Response from `MlModelsApi.AddFollower1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the model | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFollower1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Id of the user to be added as follower | 

### Return type

[**ChangeEvent**](ChangeEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMlModel

> MlModel CreateMlModel(ctx).CreateMlModel(createMlModel).Execute()

Create an ML Model



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
    createMlModel := *openapiclient.NewCreateMlModel("Algorithm_example", "Name_example", *openapiclient.NewEntityReference("Id_example", "Type_example")) // CreateMlModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MlModelsApi.CreateMlModel(context.Background()).CreateMlModel(createMlModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelsApi.CreateMlModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMlModel`: MlModel
    fmt.Fprintf(os.Stdout, "Response from `MlModelsApi.CreateMlModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMlModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMlModel** | [**CreateMlModel**](CreateMlModel.md) |  | 

### Return type

[**MlModel**](MlModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateMlModel

> MlModel CreateOrUpdateMlModel(ctx).CreateMlModel(createMlModel).Execute()

Create or update an ML Model



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
    createMlModel := *openapiclient.NewCreateMlModel("Algorithm_example", "Name_example", *openapiclient.NewEntityReference("Id_example", "Type_example")) // CreateMlModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MlModelsApi.CreateOrUpdateMlModel(context.Background()).CreateMlModel(createMlModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelsApi.CreateOrUpdateMlModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateMlModel`: MlModel
    fmt.Fprintf(os.Stdout, "Response from `MlModelsApi.CreateOrUpdateMlModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateMlModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMlModel** | [**CreateMlModel**](CreateMlModel.md) |  | 

### Return type

[**MlModel**](MlModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFollower2

> ChangeEvent DeleteFollower2(ctx, id, userId).Execute()

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
    id := "id_example" // string | Id of the model
    userId := "userId_example" // string | Id of the user being removed as follower

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MlModelsApi.DeleteFollower2(context.Background(), id, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelsApi.DeleteFollower2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFollower2`: ChangeEvent
    fmt.Fprintf(os.Stdout, "Response from `MlModelsApi.DeleteFollower2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the model | 
**userId** | **string** | Id of the user being removed as follower | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFollower2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ChangeEvent**](ChangeEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMlModel

> DeleteMlModel(ctx, id).HardDelete(hardDelete).Execute()

Delete an ML Model



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
    id := "id_example" // string | ML Model Id
    hardDelete := true // bool | Hard delete the entity. (Default = `false`) (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MlModelsApi.DeleteMlModel(context.Background(), id).HardDelete(hardDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelsApi.DeleteMlModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ML Model Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMlModelRequest struct via the builder pattern


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


## GetMlModelByFQN

> MlModel GetMlModelByFQN(ctx, fqn).Fields(fields).Include(include).Execute()

Get an ML Model by name



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
    fields := "owner,dashboard,followers,tags,usageSummary" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MlModelsApi.GetMlModelByFQN(context.Background(), fqn).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelsApi.GetMlModelByFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMlModelByFQN`: MlModel
    fmt.Fprintf(os.Stdout, "Response from `MlModelsApi.GetMlModelByFQN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqn** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMlModelByFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**MlModel**](MlModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMlModelByID

> MlModel GetMlModelByID(ctx, id).Fields(fields).Include(include).Execute()

Get an ML Model



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
    fields := "owner,dashboard,followers,tags,usageSummary" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MlModelsApi.GetMlModelByID(context.Background(), id).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelsApi.GetMlModelByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMlModelByID`: MlModel
    fmt.Fprintf(os.Stdout, "Response from `MlModelsApi.GetMlModelByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMlModelByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**MlModel**](MlModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificMlModelVersion

> MlModel GetSpecificMlModelVersion(ctx, id, version).Execute()

Get a version of the ML Model



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
    id := "id_example" // string | ML Model Id
    version := "0.1 or 1.1" // string | ML Model version number in the form `major`.`minor`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MlModelsApi.GetSpecificMlModelVersion(context.Background(), id, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelsApi.GetSpecificMlModelVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificMlModelVersion`: MlModel
    fmt.Fprintf(os.Stdout, "Response from `MlModelsApi.GetSpecificMlModelVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ML Model Id | 
**version** | **string** | ML Model version number in the form &#x60;major&#x60;.&#x60;minor&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificMlModelVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MlModel**](MlModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllMlModelVersion

> EntityHistory ListAllMlModelVersion(ctx, id).Execute()

List Ml Model versions



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
    id := "id_example" // string | ML Model Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MlModelsApi.ListAllMlModelVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelsApi.ListAllMlModelVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllMlModelVersion`: EntityHistory
    fmt.Fprintf(os.Stdout, "Response from `MlModelsApi.ListAllMlModelVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ML Model Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllMlModelVersionRequest struct via the builder pattern


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


## ListMlModels

> MlModelList ListMlModels(ctx).Fields(fields).Service(service).Limit(limit).Before(before).After(after).Include(include).Execute()

List ML Models



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
    fields := "owner,dashboard,followers,tags,usageSummary" // string | Fields requested in the returned resource (optional)
    service := "airflow" // string | Filter MlModels by service name (optional)
    limit := int32(56) // int32 | Limit the number models returned. (1 to 1000000, default = 10) (optional) (default to 10)
    before := "before_example" // string | Returns list of models before this cursor (optional)
    after := "after_example" // string | Returns list of models after this cursor (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MlModelsApi.ListMlModels(context.Background()).Fields(fields).Service(service).Limit(limit).Before(before).After(after).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelsApi.ListMlModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMlModels`: MlModelList
    fmt.Fprintf(os.Stdout, "Response from `MlModelsApi.ListMlModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMlModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Fields requested in the returned resource | 
 **service** | **string** | Filter MlModels by service name | 
 **limit** | **int32** | Limit the number models returned. (1 to 1000000, default &#x3D; 10) | [default to 10]
 **before** | **string** | Returns list of models before this cursor | 
 **after** | **string** | Returns list of models after this cursor | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**MlModelList**](MlModelList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMlModel

> PatchMlModel(ctx, id).Body(body).Execute()

Update an ML Model



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
    id := "id_example" // string | Id of the ML Model
    body := map[string]interface{}{ ... } // map[string]interface{} | JsonPatch with array of operations (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MlModelsApi.PatchMlModel(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MlModelsApi.PatchMlModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the ML Model | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMlModelRequest struct via the builder pattern


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

