# \RolesApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateRole**](RolesApi.md#CreateOrUpdateRole) | **Put** /v1/roles | Update role
[**CreateRole**](RolesApi.md#CreateRole) | **Post** /v1/roles | Create a role
[**DeleteRole**](RolesApi.md#DeleteRole) | **Delete** /v1/roles/{id} | Delete a role
[**GetRoleByFQN**](RolesApi.md#GetRoleByFQN) | **Get** /v1/roles/name/{name} | Get a role by name
[**GetRoleByID**](RolesApi.md#GetRoleByID) | **Get** /v1/roles/{id} | Get a role
[**GetSpecificRoleVersion**](RolesApi.md#GetSpecificRoleVersion) | **Get** /v1/roles/{id}/versions/{version} | Get a version of the role
[**ListAllRoleVersion**](RolesApi.md#ListAllRoleVersion) | **Get** /v1/roles/{id}/versions | List role versions
[**ListRoles**](RolesApi.md#ListRoles) | **Get** /v1/roles | List roles
[**PatchRole**](RolesApi.md#PatchRole) | **Patch** /v1/roles/{id} | Update a role



## CreateOrUpdateRole

> Role CreateOrUpdateRole(ctx).CreateRole(createRole).Execute()

Update role



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
    createRole := *openapiclient.NewCreateRole("Name_example", []openapiclient.EntityReference{*openapiclient.NewEntityReference("Id_example", "Type_example")}) // CreateRole |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.CreateOrUpdateRole(context.Background()).CreateRole(createRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.CreateOrUpdateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.CreateOrUpdateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRole** | [**CreateRole**](CreateRole.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> Role CreateRole(ctx).CreateRole(createRole).Execute()

Create a role



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
    createRole := *openapiclient.NewCreateRole("Name_example", []openapiclient.EntityReference{*openapiclient.NewEntityReference("Id_example", "Type_example")}) // CreateRole |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.CreateRole(context.Background()).CreateRole(createRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.CreateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.CreateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRole** | [**CreateRole**](CreateRole.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, id).HardDelete(hardDelete).Execute()

Delete a role



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
    hardDelete := true // bool | Hard delete the entity. (Default = `false`) (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.DeleteRole(context.Background(), id).HardDelete(hardDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.DeleteRole``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


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


## GetRoleByFQN

> Role GetRoleByFQN(ctx, name).Fields(fields).Include(include).Execute()

Get a role by name



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
    name := "name_example" // string | 
    fields := "policies,teams,users" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.GetRoleByFQN(context.Background(), name).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetRoleByFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleByFQN`: Role
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetRoleByFQN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleByFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleByID

> Role GetRoleByID(ctx, id).Fields(fields).Include(include).Execute()

Get a role



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
    fields := "policies,teams,users" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.GetRoleByID(context.Background(), id).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetRoleByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleByID`: Role
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetRoleByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificRoleVersion

> Role GetSpecificRoleVersion(ctx, id, version).Execute()

Get a version of the role



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
    id := "id_example" // string | Role Id
    version := "0.1 or 1.1" // string | Role version number in the form `major`.`minor`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.GetSpecificRoleVersion(context.Background(), id, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetSpecificRoleVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificRoleVersion`: Role
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetSpecificRoleVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Role Id | 
**version** | **string** | Role version number in the form &#x60;major&#x60;.&#x60;minor&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificRoleVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllRoleVersion

> EntityHistory ListAllRoleVersion(ctx, id).Execute()

List role versions



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
    id := "id_example" // string | role Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.ListAllRoleVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ListAllRoleVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllRoleVersion`: EntityHistory
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.ListAllRoleVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllRoleVersionRequest struct via the builder pattern


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


## ListRoles

> RoleList ListRoles(ctx).Default_(default_).Fields(fields).Limit(limit).Before(before).After(after).Include(include).Execute()

List roles



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
    default_ := true // bool | List only default role(s) (optional)
    fields := "policies,teams,users" // string | Fields requested in the returned resource (optional)
    limit := int32(56) // int32 | Limit the number tables returned. (1 to 1000000, default = 10) (optional) (default to 10)
    before := "before_example" // string | Returns list of tables before this cursor (optional)
    after := "after_example" // string | Returns list of tables after this cursor (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.ListRoles(context.Background()).Default_(default_).Fields(fields).Limit(limit).Before(before).After(after).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoles`: RoleList
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.ListRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **default_** | **bool** | List only default role(s) | 
 **fields** | **string** | Fields requested in the returned resource | 
 **limit** | **int32** | Limit the number tables returned. (1 to 1000000, default &#x3D; 10) | [default to 10]
 **before** | **string** | Returns list of tables before this cursor | 
 **after** | **string** | Returns list of tables after this cursor | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**RoleList**](RoleList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRole

> PatchRole(ctx, id).Body(body).Execute()

Update a role



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
    resp, r, err := apiClient.RolesApi.PatchRole(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.PatchRole``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchRoleRequest struct via the builder pattern


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

