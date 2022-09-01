# \DatabaseSchemasApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDBSchema**](DatabaseSchemasApi.md#CreateDBSchema) | **Post** /v1/databaseSchemas | Create a schema
[**CreateOrUpdateDBSchema**](DatabaseSchemasApi.md#CreateOrUpdateDBSchema) | **Put** /v1/databaseSchemas | Create or update schema
[**DeleteDBSchema**](DatabaseSchemasApi.md#DeleteDBSchema) | **Delete** /v1/databaseSchemas/{id} | Delete a schema
[**GetDBSchemaByFQN**](DatabaseSchemasApi.md#GetDBSchemaByFQN) | **Get** /v1/databaseSchemas/name/{fqn} | Get a schema by name
[**GetDBSchemaByID**](DatabaseSchemasApi.md#GetDBSchemaByID) | **Get** /v1/databaseSchemas/{id} | Get a schema
[**GetSpecificDBSchemaVersion**](DatabaseSchemasApi.md#GetSpecificDBSchemaVersion) | **Get** /v1/databaseSchemas/{id}/versions/{version} | Get a version of the schema
[**ListAllDBSchemaVersion**](DatabaseSchemasApi.md#ListAllDBSchemaVersion) | **Get** /v1/databaseSchemas/{id}/versions | List schema versions
[**ListDBSchemas**](DatabaseSchemasApi.md#ListDBSchemas) | **Get** /v1/databaseSchemas | List database schemas
[**PatchDBSchema**](DatabaseSchemasApi.md#PatchDBSchema) | **Patch** /v1/databaseSchemas/{id} | Update a database schema



## CreateDBSchema

> DatabaseSchema CreateDBSchema(ctx).CreateDatabaseSchema(createDatabaseSchema).Execute()

Create a schema



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
    createDatabaseSchema := *openapiclient.NewCreateDatabaseSchema(*openapiclient.NewEntityReference("Id_example", "Type_example"), "Name_example") // CreateDatabaseSchema |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseSchemasApi.CreateDBSchema(context.Background()).CreateDatabaseSchema(createDatabaseSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseSchemasApi.CreateDBSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDBSchema`: DatabaseSchema
    fmt.Fprintf(os.Stdout, "Response from `DatabaseSchemasApi.CreateDBSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDBSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDatabaseSchema** | [**CreateDatabaseSchema**](CreateDatabaseSchema.md) |  | 

### Return type

[**DatabaseSchema**](DatabaseSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateDBSchema

> DatabaseSchema CreateOrUpdateDBSchema(ctx).CreateDatabaseSchema(createDatabaseSchema).Execute()

Create or update schema



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
    createDatabaseSchema := *openapiclient.NewCreateDatabaseSchema(*openapiclient.NewEntityReference("Id_example", "Type_example"), "Name_example") // CreateDatabaseSchema |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseSchemasApi.CreateOrUpdateDBSchema(context.Background()).CreateDatabaseSchema(createDatabaseSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseSchemasApi.CreateOrUpdateDBSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateDBSchema`: DatabaseSchema
    fmt.Fprintf(os.Stdout, "Response from `DatabaseSchemasApi.CreateOrUpdateDBSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateDBSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDatabaseSchema** | [**CreateDatabaseSchema**](CreateDatabaseSchema.md) |  | 

### Return type

[**DatabaseSchema**](DatabaseSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDBSchema

> DeleteDBSchema(ctx, id).Recursive(recursive).HardDelete(hardDelete).Execute()

Delete a schema



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
    recursive := true // bool | Recursively delete this entity and it's children. (Default `false`) (optional) (default to false)
    hardDelete := true // bool | Hard delete the entity. (Default = `false`) (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseSchemasApi.DeleteDBSchema(context.Background(), id).Recursive(recursive).HardDelete(hardDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseSchemasApi.DeleteDBSchema``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDBSchemaRequest struct via the builder pattern


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


## GetDBSchemaByFQN

> DatabaseSchema GetDBSchemaByFQN(ctx, fqn).Fields(fields).Include(include).Execute()

Get a schema by name



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
    fields := "owner,tables,usageSummary" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseSchemasApi.GetDBSchemaByFQN(context.Background(), fqn).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseSchemasApi.GetDBSchemaByFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDBSchemaByFQN`: DatabaseSchema
    fmt.Fprintf(os.Stdout, "Response from `DatabaseSchemasApi.GetDBSchemaByFQN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqn** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDBSchemaByFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**DatabaseSchema**](DatabaseSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDBSchemaByID

> DatabaseSchema GetDBSchemaByID(ctx, id).Fields(fields).Include(include).Execute()

Get a schema



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
    fields := "owner,tables,usageSummary" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseSchemasApi.GetDBSchemaByID(context.Background(), id).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseSchemasApi.GetDBSchemaByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDBSchemaByID`: DatabaseSchema
    fmt.Fprintf(os.Stdout, "Response from `DatabaseSchemasApi.GetDBSchemaByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDBSchemaByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**DatabaseSchema**](DatabaseSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificDBSchemaVersion

> DatabaseSchema GetSpecificDBSchemaVersion(ctx, id, version).Execute()

Get a version of the schema



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
    id := "id_example" // string | Database schema Id
    version := "0.1 or 1.1" // string | Database schema version number in the form `major`.`minor`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseSchemasApi.GetSpecificDBSchemaVersion(context.Background(), id, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseSchemasApi.GetSpecificDBSchemaVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificDBSchemaVersion`: DatabaseSchema
    fmt.Fprintf(os.Stdout, "Response from `DatabaseSchemasApi.GetSpecificDBSchemaVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database schema Id | 
**version** | **string** | Database schema version number in the form &#x60;major&#x60;.&#x60;minor&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificDBSchemaVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DatabaseSchema**](DatabaseSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllDBSchemaVersion

> EntityHistory ListAllDBSchemaVersion(ctx, id).Execute()

List schema versions



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
    id := "id_example" // string | Database schema Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseSchemasApi.ListAllDBSchemaVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseSchemasApi.ListAllDBSchemaVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllDBSchemaVersion`: EntityHistory
    fmt.Fprintf(os.Stdout, "Response from `DatabaseSchemasApi.ListAllDBSchemaVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database schema Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllDBSchemaVersionRequest struct via the builder pattern


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


## ListDBSchemas

> DatabaseSchemaList ListDBSchemas(ctx).Fields(fields).Database(database).Limit(limit).Before(before).After(after).Include(include).Execute()

List database schemas



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
    fields := "owner,tables,usageSummary" // string | Fields requested in the returned resource (optional)
    database := "customerDatabase" // string | Filter schemas by database name (optional)
    limit := int32(56) // int32 | Limit the number schemas returned. (1 to 1000000, default = 10) (optional) (default to 10)
    before := "before_example" // string | Returns list of schemas before this cursor (optional)
    after := "after_example" // string | Returns list of schemas after this cursor (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseSchemasApi.ListDBSchemas(context.Background()).Fields(fields).Database(database).Limit(limit).Before(before).After(after).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseSchemasApi.ListDBSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDBSchemas`: DatabaseSchemaList
    fmt.Fprintf(os.Stdout, "Response from `DatabaseSchemasApi.ListDBSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDBSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Fields requested in the returned resource | 
 **database** | **string** | Filter schemas by database name | 
 **limit** | **int32** | Limit the number schemas returned. (1 to 1000000, default &#x3D; 10) | [default to 10]
 **before** | **string** | Returns list of schemas before this cursor | 
 **after** | **string** | Returns list of schemas after this cursor | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**DatabaseSchemaList**](DatabaseSchemaList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDBSchema

> PatchDBSchema(ctx, id).Body(body).Execute()

Update a database schema



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
    resp, r, err := apiClient.DatabaseSchemasApi.PatchDBSchema(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseSchemasApi.PatchDBSchema``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchDBSchemaRequest struct via the builder pattern


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

