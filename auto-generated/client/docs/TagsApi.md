# \TagsApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdatePrimaryTag**](TagsApi.md#CreateOrUpdatePrimaryTag) | **Put** /v1/tags/{category}/{primaryTag} | Update a primaryTag
[**CreateOrUpdateSecondaryTag**](TagsApi.md#CreateOrUpdateSecondaryTag) | **Put** /v1/tags/{category}/{primaryTag}/{secondaryTag} | Update a secondaryTag
[**CreateOrUpdateTagCategory**](TagsApi.md#CreateOrUpdateTagCategory) | **Put** /v1/tags/{category} | Update a tag category
[**CreatePrimaryTag**](TagsApi.md#CreatePrimaryTag) | **Post** /v1/tags/{category} | Create a primary tag
[**CreateSecondaryTag**](TagsApi.md#CreateSecondaryTag) | **Post** /v1/tags/{category}/{primaryTag} | Create a secondary tag
[**CreateTagCategory**](TagsApi.md#CreateTagCategory) | **Post** /v1/tags | Create a tag category
[**DeleteTagCategory**](TagsApi.md#DeleteTagCategory) | **Delete** /v1/tags/{id} | Delete tag category
[**DeleteTags**](TagsApi.md#DeleteTags) | **Delete** /v1/tags/{category}/{id} | Delete tag
[**GetPrimaryTag**](TagsApi.md#GetPrimaryTag) | **Get** /v1/tags/{category}/{primaryTag} | Get a primary tag
[**GetSecondaryTag**](TagsApi.md#GetSecondaryTag) | **Get** /v1/tags/{category}/{primaryTag}/{secondaryTag} | Get a secondary tag
[**GetTagCategoryByName**](TagsApi.md#GetTagCategoryByName) | **Get** /v1/tags/{category} | Get a tag category
[**ListTagCategories**](TagsApi.md#ListTagCategories) | **Get** /v1/tags | List tag categories



## CreateOrUpdatePrimaryTag

> CreateOrUpdatePrimaryTag(ctx, category, primaryTag).CreateTag(createTag).Execute()

Update a primaryTag



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
    category := "category_example" // string | Tag category name
    primaryTag := "<primaryTag> fully qualified name <categoryName>.<primaryTag>" // string | Primary tag name
    createTag := *openapiclient.NewCreateTag("Description_example", "Name_example") // CreateTag |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.CreateOrUpdatePrimaryTag(context.Background(), category, primaryTag).CreateTag(createTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.CreateOrUpdatePrimaryTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** | Tag category name | 
**primaryTag** | **string** | Primary tag name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdatePrimaryTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createTag** | [**CreateTag**](CreateTag.md) |  | 

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


## CreateOrUpdateSecondaryTag

> CreateOrUpdateSecondaryTag(ctx, category, primaryTag, secondaryTag).CreateTag(createTag).Execute()

Update a secondaryTag



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
    category := "category_example" // string | Tag category name
    primaryTag := "<primaryTag> fully qualified name <categoryName>.<primaryTag>" // string | Primary tag name
    secondaryTag := "<secondaryTag> fully qualified name <categoryName>.<primaryTag>.<secondaryTag>" // string | SecondaryTag tag name
    createTag := *openapiclient.NewCreateTag("Description_example", "Name_example") // CreateTag |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.CreateOrUpdateSecondaryTag(context.Background(), category, primaryTag, secondaryTag).CreateTag(createTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.CreateOrUpdateSecondaryTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** | Tag category name | 
**primaryTag** | **string** | Primary tag name | 
**secondaryTag** | **string** | SecondaryTag tag name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateSecondaryTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createTag** | [**CreateTag**](CreateTag.md) |  | 

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


## CreateOrUpdateTagCategory

> CreateOrUpdateTagCategory(ctx, category).CreateTagCategory(createTagCategory).Execute()

Update a tag category



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
    category := "category_example" // string | Tag category name
    createTagCategory := *openapiclient.NewCreateTagCategory("CategoryType_example", "Description_example", "Name_example") // CreateTagCategory |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.CreateOrUpdateTagCategory(context.Background(), category).CreateTagCategory(createTagCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.CreateOrUpdateTagCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** | Tag category name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateTagCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTagCategory** | [**CreateTagCategory**](CreateTagCategory.md) |  | 

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


## CreatePrimaryTag

> Tag CreatePrimaryTag(ctx, category).CreateTag(createTag).Execute()

Create a primary tag



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
    category := "category_example" // string | Tag category name
    createTag := *openapiclient.NewCreateTag("Description_example", "Name_example") // CreateTag |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.CreatePrimaryTag(context.Background(), category).CreateTag(createTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.CreatePrimaryTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePrimaryTag`: Tag
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.CreatePrimaryTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** | Tag category name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrimaryTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTag** | [**CreateTag**](CreateTag.md) |  | 

### Return type

[**Tag**](Tag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecondaryTag

> Tag CreateSecondaryTag(ctx, category, primaryTag).CreateTag(createTag).Execute()

Create a secondary tag



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
    category := "category_example" // string | Tag category name
    primaryTag := "<primaryTag> fully qualified name <categoryName>.<primaryTag>" // string | Primary tag name
    createTag := *openapiclient.NewCreateTag("Description_example", "Name_example") // CreateTag |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.CreateSecondaryTag(context.Background(), category, primaryTag).CreateTag(createTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.CreateSecondaryTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecondaryTag`: Tag
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.CreateSecondaryTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** | Tag category name | 
**primaryTag** | **string** | Primary tag name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecondaryTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createTag** | [**CreateTag**](CreateTag.md) |  | 

### Return type

[**Tag**](Tag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTagCategory

> TagCategory CreateTagCategory(ctx).CreateTagCategory(createTagCategory).Execute()

Create a tag category



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
    createTagCategory := *openapiclient.NewCreateTagCategory("CategoryType_example", "Description_example", "Name_example") // CreateTagCategory |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.CreateTagCategory(context.Background()).CreateTagCategory(createTagCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.CreateTagCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTagCategory`: TagCategory
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.CreateTagCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTagCategory** | [**CreateTagCategory**](CreateTagCategory.md) |  | 

### Return type

[**TagCategory**](TagCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTagCategory

> DeleteTagCategory(ctx, id).Execute()

Delete tag category



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
    id := "id_example" // string | Tag category id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.DeleteTagCategory(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.DeleteTagCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tag category id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagCategoryRequest struct via the builder pattern


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


## DeleteTags

> DeleteTags(ctx, category, id).Execute()

Delete tag



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
    category := "category_example" // string | Tag id
    id := "id_example" // string | Tag id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.DeleteTags(context.Background(), category, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.DeleteTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** | Tag id | 
**id** | **string** | Tag id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagsRequest struct via the builder pattern


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


## GetPrimaryTag

> Tag GetPrimaryTag(ctx, category, primaryTag).Fields(fields).Execute()

Get a primary tag



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
    category := "category_example" // string | Tag category name
    primaryTag := "<primaryTag> fully qualified name <categoryName>.<primaryTag>" // string | Primary tag name
    fields := "usageCount" // string | Fields requested in the returned resource (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.GetPrimaryTag(context.Background(), category, primaryTag).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GetPrimaryTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrimaryTag`: Tag
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.GetPrimaryTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** | Tag category name | 
**primaryTag** | **string** | Primary tag name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrimaryTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Fields requested in the returned resource | 

### Return type

[**Tag**](Tag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecondaryTag

> Tag GetSecondaryTag(ctx, category, primaryTag, secondaryTag).Fields(fields).Execute()

Get a secondary tag



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
    category := "category_example" // string | Tag category name
    primaryTag := "<primaryTag> fully qualified name <categoryName>.<primaryTag>" // string | Primary tag name
    secondaryTag := "<secondaryTag> fully qualified name <categoryName>.<primaryTag>.<SecondaryTag>" // string | Secondary tag name
    fields := "usageCount" // string | Fields requested in the returned resource (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.GetSecondaryTag(context.Background(), category, primaryTag, secondaryTag).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GetSecondaryTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecondaryTag`: Tag
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.GetSecondaryTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** | Tag category name | 
**primaryTag** | **string** | Primary tag name | 
**secondaryTag** | **string** | Secondary tag name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecondaryTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **string** | Fields requested in the returned resource | 

### Return type

[**Tag**](Tag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagCategoryByName

> TagCategory GetTagCategoryByName(ctx, category).Fields(fields).Execute()

Get a tag category



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
    category := "category_example" // string | Tag category name
    fields := "usageCount" // string | Fields requested in the returned resource (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.GetTagCategoryByName(context.Background(), category).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GetTagCategoryByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagCategoryByName`: TagCategory
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.GetTagCategoryByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** | Tag category name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagCategoryByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 

### Return type

[**TagCategory**](TagCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTagCategories

> CategoryList ListTagCategories(ctx).Fields(fields).Execute()

List tag categories



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
    fields := "usageCount" // string | Fields requested in the returned resource (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.ListTagCategories(context.Background()).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.ListTagCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTagCategories`: CategoryList
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.ListTagCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTagCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Fields requested in the returned resource | 

### Return type

[**CategoryList**](CategoryList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

