# \WebhookApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateWebhook**](WebhookApi.md#CreateOrUpdateWebhook) | **Put** /v1/webhook | Updated an existing or create a new webhook
[**CreateWebHook**](WebhookApi.md#CreateWebHook) | **Post** /v1/webhook | Subscribe to a new webhook
[**DeleteWebHook**](WebhookApi.md#DeleteWebHook) | **Delete** /v1/webhook/{id} | Delete a webhook
[**GetSpecificWebhookVersion**](WebhookApi.md#GetSpecificWebhookVersion) | **Get** /v1/webhook/{id}/versions/{version} | Get a version of the webhook
[**GetWebHookByFQN**](WebhookApi.md#GetWebHookByFQN) | **Get** /v1/webhook/name/{name} | Get a webhook by name
[**GetWebHookByID**](WebhookApi.md#GetWebHookByID) | **Get** /v1/webhook/{id} | Get a webhook
[**ListAllWebHookVersion**](WebhookApi.md#ListAllWebHookVersion) | **Get** /v1/webhook/{id}/versions | List webhook versions
[**ListWebHooks**](WebhookApi.md#ListWebHooks) | **Get** /v1/webhook | List webhooks
[**PatchWebHook**](WebhookApi.md#PatchWebHook) | **Patch** /v1/webhook/{id} | Update a webhook



## CreateOrUpdateWebhook

> Webhook CreateOrUpdateWebhook(ctx).CreateWebhook(createWebhook).Execute()

Updated an existing or create a new webhook



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
    createWebhook := *openapiclient.NewCreateWebhook("Endpoint_example", []openapiclient.EventFilter{*openapiclient.NewEventFilter("EventType_example")}, "Name_example") // CreateWebhook |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.CreateOrUpdateWebhook(context.Background()).CreateWebhook(createWebhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.CreateOrUpdateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateWebhook`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.CreateOrUpdateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWebhook** | [**CreateWebhook**](CreateWebhook.md) |  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebHook

> Webhook CreateWebHook(ctx).CreateWebhook(createWebhook).Execute()

Subscribe to a new webhook



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
    createWebhook := *openapiclient.NewCreateWebhook("Endpoint_example", []openapiclient.EventFilter{*openapiclient.NewEventFilter("EventType_example")}, "Name_example") // CreateWebhook |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.CreateWebHook(context.Background()).CreateWebhook(createWebhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.CreateWebHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebHook`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.CreateWebHook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWebhook** | [**CreateWebhook**](CreateWebhook.md) |  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebHook

> Webhook DeleteWebHook(ctx, id).Execute()

Delete a webhook



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
    id := "id_example" // string | webhook Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.DeleteWebHook(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.DeleteWebHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWebHook`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.DeleteWebHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | webhook Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Webhook**](Webhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificWebhookVersion

> Webhook GetSpecificWebhookVersion(ctx, id, version).Execute()

Get a version of the webhook



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
    id := "id_example" // string | webhook Id
    version := "0.1 or 1.1" // string | webhook version number in the form `major`.`minor`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.GetSpecificWebhookVersion(context.Background(), id, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.GetSpecificWebhookVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificWebhookVersion`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.GetSpecificWebhookVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | webhook Id | 
**version** | **string** | webhook version number in the form &#x60;major&#x60;.&#x60;minor&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificWebhookVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Webhook**](Webhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebHookByFQN

> Webhook GetWebHookByFQN(ctx, name).Include(include).Execute()

Get a webhook by name



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
    name := "name_example" // string | Name of the webhook
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.GetWebHookByFQN(context.Background(), name).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.GetWebHookByFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebHookByFQN`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.GetWebHookByFQN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the webhook | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebHookByFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**Webhook**](Webhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebHookByID

> Webhook GetWebHookByID(ctx, id).Include(include).Execute()

Get a webhook



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
    id := "id_example" // string | webhook Id
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.GetWebHookByID(context.Background(), id).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.GetWebHookByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebHookByID`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.GetWebHookByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | webhook Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebHookByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**Webhook**](Webhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllWebHookVersion

> EntityHistory ListAllWebHookVersion(ctx, id).Execute()

List webhook versions



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
    id := "id_example" // string | webhook Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.ListAllWebHookVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.ListAllWebHookVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllWebHookVersion`: EntityHistory
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.ListAllWebHookVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | webhook Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllWebHookVersionRequest struct via the builder pattern


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


## ListWebHooks

> WebhookList ListWebHooks(ctx).Status(status).Limit(limit).Before(before).After(after).Include(include).Execute()

List webhooks



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
    status := "active" // string | Filter webhooks by status (optional)
    limit := int32(56) // int32 | Limit the number webhooks returned. (1 to 1000000, default = 10)  (optional) (default to 10)
    before := "before_example" // string | Returns list of webhooks before this cursor (optional)
    after := "after_example" // string | Returns list of webhooks after this cursor (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.ListWebHooks(context.Background()).Status(status).Limit(limit).Before(before).After(after).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.ListWebHooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebHooks`: WebhookList
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.ListWebHooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWebHooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Filter webhooks by status | 
 **limit** | **int32** | Limit the number webhooks returned. (1 to 1000000, default &#x3D; 10)  | [default to 10]
 **before** | **string** | Returns list of webhooks before this cursor | 
 **after** | **string** | Returns list of webhooks after this cursor | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**WebhookList**](WebhookList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchWebHook

> PatchWebHook(ctx, id).Body(body).Execute()

Update a webhook



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
    resp, r, err := apiClient.WebhookApi.PatchWebHook(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.PatchWebHook``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchWebHookRequest struct via the builder pattern


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

