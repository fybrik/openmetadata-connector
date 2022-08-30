# \LocationsApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFollower**](LocationsApi.md#AddFollower) | **Put** /v1/locations/{id}/followers | Add a follower
[**CreateLocation**](LocationsApi.md#CreateLocation) | **Post** /v1/locations | Create a location
[**CreateOrUpdateLocation**](LocationsApi.md#CreateOrUpdateLocation) | **Put** /v1/locations | Create or update location
[**DeleteFollower1**](LocationsApi.md#DeleteFollower1) | **Delete** /v1/locations/{id}/followers/{userId} | Remove a follower
[**DeleteLocation2**](LocationsApi.md#DeleteLocation2) | **Delete** /v1/locations/{id} | Delete a location
[**GetLocationByFQN**](LocationsApi.md#GetLocationByFQN) | **Get** /v1/locations/name/{fqn} | Get a location by name
[**GetLocationByID**](LocationsApi.md#GetLocationByID) | **Get** /v1/locations/{id} | Get a location
[**GetSpecificLocationVersion**](LocationsApi.md#GetSpecificLocationVersion) | **Get** /v1/locations/{id}/versions/{version} | Get a version of the location
[**ListAllLocationVersion**](LocationsApi.md#ListAllLocationVersion) | **Get** /v1/locations/{id}/versions | List location versions
[**ListLocationPrefixes**](LocationsApi.md#ListLocationPrefixes) | **Get** /v1/locations/prefixes/{fqn} | List locations that are prefixes
[**ListLocations**](LocationsApi.md#ListLocations) | **Get** /v1/locations | List locations
[**PatchLocation**](LocationsApi.md#PatchLocation) | **Patch** /v1/locations/{id} | Update a location



## AddFollower

> ChangeEvent AddFollower(ctx, id).Body(body).Execute()

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
    id := "id_example" // string | Id of the location
    body := "body_example" // string | Id of the user to be added as follower (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.AddFollower(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.AddFollower``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddFollower`: ChangeEvent
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.AddFollower`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the location | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFollowerRequest struct via the builder pattern


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


## CreateLocation

> Location CreateLocation(ctx).CreateLocation(createLocation).Execute()

Create a location



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
    createLocation := *openapiclient.NewCreateLocation("Name_example", *openapiclient.NewEntityReference("Id_example", "Type_example")) // CreateLocation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.CreateLocation(context.Background()).CreateLocation(createLocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.CreateLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocation`: Location
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.CreateLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLocation** | [**CreateLocation**](CreateLocation.md) |  | 

### Return type

[**Location**](Location.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateLocation

> Location CreateOrUpdateLocation(ctx).CreateLocation(createLocation).Execute()

Create or update location



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
    createLocation := *openapiclient.NewCreateLocation("Name_example", *openapiclient.NewEntityReference("Id_example", "Type_example")) // CreateLocation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.CreateOrUpdateLocation(context.Background()).CreateLocation(createLocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.CreateOrUpdateLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateLocation`: Location
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.CreateOrUpdateLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLocation** | [**CreateLocation**](CreateLocation.md) |  | 

### Return type

[**Location**](Location.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFollower1

> ChangeEvent DeleteFollower1(ctx, id, userId).Execute()

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
    id := "id_example" // string | Id of the location
    userId := "userId_example" // string | Id of the user being removed as follower

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.DeleteFollower1(context.Background(), id, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.DeleteFollower1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFollower1`: ChangeEvent
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.DeleteFollower1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the location | 
**userId** | **string** | Id of the user being removed as follower | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFollower1Request struct via the builder pattern


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


## DeleteLocation2

> DeleteLocation2(ctx, id).HardDelete(hardDelete).Execute()

Delete a location



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
    id := "id_example" // string | Location Id
    hardDelete := true // bool | Hard delete the entity. (Default = `false`) (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.DeleteLocation2(context.Background(), id).HardDelete(hardDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.DeleteLocation2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocation2Request struct via the builder pattern


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


## GetLocationByFQN

> Location GetLocationByFQN(ctx, fqn).Fields(fields).Include(include).Execute()

Get a location by name



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
    fqn := "fqn_example" // string | Fully qualified name of the location urlencoded if needed
    fields := "owner,followers,tags,path" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.GetLocationByFQN(context.Background(), fqn).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.GetLocationByFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocationByFQN`: Location
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.GetLocationByFQN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqn** | **string** | Fully qualified name of the location urlencoded if needed | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationByFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**Location**](Location.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocationByID

> Location GetLocationByID(ctx, id).Fields(fields).Include(include).Execute()

Get a location



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
    id := "id_example" // string | location Id
    fields := "owner,followers,tags,path" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.GetLocationByID(context.Background(), id).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.GetLocationByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocationByID`: Location
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.GetLocationByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**Location**](Location.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificLocationVersion

> Location GetSpecificLocationVersion(ctx, id, version).Execute()

Get a version of the location



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
    id := "id_example" // string | location Id
    version := "0.1 or 1.1" // string | location version number in the form `major`.`minor`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.GetSpecificLocationVersion(context.Background(), id, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.GetSpecificLocationVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificLocationVersion`: Location
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.GetSpecificLocationVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | location Id | 
**version** | **string** | location version number in the form &#x60;major&#x60;.&#x60;minor&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificLocationVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Location**](Location.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllLocationVersion

> EntityHistory ListAllLocationVersion(ctx, id).Execute()

List location versions



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
    id := "id_example" // string | location Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.ListAllLocationVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.ListAllLocationVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllLocationVersion`: EntityHistory
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.ListAllLocationVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllLocationVersionRequest struct via the builder pattern


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


## ListLocationPrefixes

> LocationList ListLocationPrefixes(ctx, fqn).Fields(fields).Limit(limit).Before(before).After(after).Execute()

List locations that are prefixes



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
    fqn := "fqn_example" // string | Fully qualified name of the location urlencoded if needed
    fields := "owner,followers,tags,path" // string | Fields requested in the returned resource (optional)
    limit := int32(56) // int32 | Limit the number locations returned. (1 to 1000000, default = 10) (optional) (default to 10)
    before := "before_example" // string | Returns list of locations before this cursor (optional)
    after := "after_example" // string | Returns list of locations after this cursor (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.ListLocationPrefixes(context.Background(), fqn).Fields(fields).Limit(limit).Before(before).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.ListLocationPrefixes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocationPrefixes`: LocationList
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.ListLocationPrefixes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqn** | **string** | Fully qualified name of the location urlencoded if needed | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLocationPrefixesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **limit** | **int32** | Limit the number locations returned. (1 to 1000000, default &#x3D; 10) | [default to 10]
 **before** | **string** | Returns list of locations before this cursor | 
 **after** | **string** | Returns list of locations after this cursor | 

### Return type

[**LocationList**](LocationList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocations

> LocationList ListLocations(ctx).Fields(fields).Service(service).Limit(limit).Before(before).After(after).Include(include).Execute()

List locations



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
    fields := "owner,followers,tags,path" // string | Fields requested in the returned resource (optional)
    service := "s3://bucket/folder1" // string | Filter locations by prefix of the FQN (optional)
    limit := int32(56) // int32 | Limit the number locations returned. (1 to 1000000, default = 10) (optional) (default to 10)
    before := "before_example" // string | Returns list of locations before this cursor (optional)
    after := "after_example" // string | Returns list of locations after this cursor (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.ListLocations(context.Background()).Fields(fields).Service(service).Limit(limit).Before(before).After(after).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.ListLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocations`: LocationList
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.ListLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Fields requested in the returned resource | 
 **service** | **string** | Filter locations by prefix of the FQN | 
 **limit** | **int32** | Limit the number locations returned. (1 to 1000000, default &#x3D; 10) | [default to 10]
 **before** | **string** | Returns list of locations before this cursor | 
 **after** | **string** | Returns list of locations after this cursor | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**LocationList**](LocationList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchLocation

> PatchLocation(ctx, id).Body(body).Execute()

Update a location



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
    resp, r, err := apiClient.LocationsApi.PatchLocation(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.PatchLocation``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchLocationRequest struct via the builder pattern


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

