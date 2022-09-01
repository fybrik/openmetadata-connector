# \FeedsApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPostToThread**](FeedsApi.md#AddPostToThread) | **Post** /v1/feed/{id}/posts | Add post to a thread
[**CloseTask**](FeedsApi.md#CloseTask) | **Put** /v1/feed/tasks/{id}/close | Close a task
[**CountThreads**](FeedsApi.md#CountThreads) | **Get** /v1/feed/count | count of threads
[**CreateThread**](FeedsApi.md#CreateThread) | **Post** /v1/feed | Create a thread
[**DeletePostFromThread**](FeedsApi.md#DeletePostFromThread) | **Delete** /v1/feed/{threadId}/posts/{postId} | Delete a post from its thread
[**GetAllPostOfThread**](FeedsApi.md#GetAllPostOfThread) | **Get** /v1/feed/{id}/posts | Get all the posts of a thread
[**GetTaskByID**](FeedsApi.md#GetTaskByID) | **Get** /v1/feed/tasks/{id} | Get a task thread by task id
[**GetThreadByID**](FeedsApi.md#GetThreadByID) | **Get** /v1/feed/{id} | Get a thread
[**ListThreads**](FeedsApi.md#ListThreads) | **Get** /v1/feed | List threads
[**PatchPostOfThread**](FeedsApi.md#PatchPostOfThread) | **Patch** /v1/feed/{threadId}/posts/{postId} | Update post of a thread by &#x60;id&#x60;.
[**PatchThread**](FeedsApi.md#PatchThread) | **Patch** /v1/feed/{id} | Update a thread by &#x60;id&#x60;.
[**ResolveTask**](FeedsApi.md#ResolveTask) | **Put** /v1/feed/tasks/{id}/resolve | Resolve a task



## AddPostToThread

> Thread AddPostToThread(ctx, id).CreatePost(createPost).Execute()

Add post to a thread



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
    createPost := *openapiclient.NewCreatePost("From_example", "Message_example") // CreatePost |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeedsApi.AddPostToThread(context.Background(), id).CreatePost(createPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedsApi.AddPostToThread``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPostToThread`: Thread
    fmt.Fprintf(os.Stdout, "Response from `FeedsApi.AddPostToThread`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPostToThreadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPost** | [**CreatePost**](CreatePost.md) |  | 

### Return type

[**Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloseTask

> Thread CloseTask(ctx, id).CloseTask(closeTask).Execute()

Close a task



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
    closeTask := *openapiclient.NewCloseTask("Comment_example") // CloseTask |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeedsApi.CloseTask(context.Background(), id).CloseTask(closeTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedsApi.CloseTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloseTask`: Thread
    fmt.Fprintf(os.Stdout, "Response from `FeedsApi.CloseTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **closeTask** | [**CloseTask**](CloseTask.md) |  | 

### Return type

[**Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountThreads

> ThreadCount CountThreads(ctx).EntityLink(entityLink).Type_(type_).IsResolved(isResolved).Execute()

count of threads



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
    entityLink := "<E#/{entityType}/{entityFQN}/{fieldName}>" // string | Filter threads by entity link (optional)
    type_ := "type__example" // string | The type of thread to filter the results. It can take one of 'Conversation', 'Task', 'Announcement' (optional)
    isResolved := true // bool | Filter threads by whether it is active or resolved (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeedsApi.CountThreads(context.Background()).EntityLink(entityLink).Type_(type_).IsResolved(isResolved).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedsApi.CountThreads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountThreads`: ThreadCount
    fmt.Fprintf(os.Stdout, "Response from `FeedsApi.CountThreads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountThreadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityLink** | **string** | Filter threads by entity link | 
 **type_** | **string** | The type of thread to filter the results. It can take one of &#39;Conversation&#39;, &#39;Task&#39;, &#39;Announcement&#39; | 
 **isResolved** | **bool** | Filter threads by whether it is active or resolved | [default to false]

### Return type

[**ThreadCount**](ThreadCount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateThread

> Thread CreateThread(ctx).CreateThread(createThread).Execute()

Create a thread



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
    createThread := *openapiclient.NewCreateThread("About_example", "From_example", "Message_example") // CreateThread |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeedsApi.CreateThread(context.Background()).CreateThread(createThread).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedsApi.CreateThread``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateThread`: Thread
    fmt.Fprintf(os.Stdout, "Response from `FeedsApi.CreateThread`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateThreadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createThread** | [**CreateThread**](CreateThread.md) |  | 

### Return type

[**Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePostFromThread

> DeletePostFromThread(ctx, threadId, postId).Execute()

Delete a post from its thread



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
    threadId := "threadId_example" // string | ThreadId of the post to be deleted
    postId := "postId_example" // string | PostId of the post to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeedsApi.DeletePostFromThread(context.Background(), threadId, postId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedsApi.DeletePostFromThread``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** | ThreadId of the post to be deleted | 
**postId** | **string** | PostId of the post to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostFromThreadRequest struct via the builder pattern


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


## GetAllPostOfThread

> PostList GetAllPostOfThread(ctx, id).Execute()

Get all the posts of a thread



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
    resp, r, err := apiClient.FeedsApi.GetAllPostOfThread(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedsApi.GetAllPostOfThread``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllPostOfThread`: PostList
    fmt.Fprintf(os.Stdout, "Response from `FeedsApi.GetAllPostOfThread`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPostOfThreadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostList**](PostList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskByID

> Thread GetTaskByID(ctx, id).Execute()

Get a task thread by task id



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
    resp, r, err := apiClient.FeedsApi.GetTaskByID(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedsApi.GetTaskByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskByID`: Thread
    fmt.Fprintf(os.Stdout, "Response from `FeedsApi.GetTaskByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThreadByID

> Thread GetThreadByID(ctx, id).Execute()

Get a thread



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
    resp, r, err := apiClient.FeedsApi.GetThreadByID(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedsApi.GetThreadByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetThreadByID`: Thread
    fmt.Fprintf(os.Stdout, "Response from `FeedsApi.GetThreadByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThreadByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListThreads

> ThreadList ListThreads(ctx).LimitPosts(limitPosts).Limit(limit).Before(before).After(after).EntityLink(entityLink).UserId(userId).FilterType(filterType).Resolved(resolved).Type_(type_).TaskStatus(taskStatus).Execute()

List threads



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
    limitPosts := int64(789) // int64 | Limit the number of posts sorted by chronological order (1 to 1000000, default = 3) (optional) (default to 3)
    limit := int32(56) // int32 | Limit the number of threads returned. (1 to 1000000, default = 10) (optional) (default to 10)
    before := "before_example" // string | Returns list of threads before this cursor (optional)
    after := "after_example" // string | Returns list of threads after this cursor (optional)
    entityLink := "<E#/{entityType}/{entityFQN}/{fieldName}>" // string | Filter threads by entity link (optional)
    userId := "userId_example" // string | Filter threads by user id. This filter requires a 'filterType' query param. The default filter type is 'OWNER'. This filter cannot be combined with the entityLink filter. (optional)
    filterType := "filterType_example" // string | Filter type definition for the user filter. It can take one of 'OWNER', 'FOLLOWS', 'MENTIONS'. This must be used with the 'user' query param (optional)
    resolved := true // bool | Filter threads by whether they are resolved or not. By default resolved is false (optional) (default to false)
    type_ := "type__example" // string | The type of thread to filter the results. It can take one of 'Conversation', 'Task', 'Announcement' (optional)
    taskStatus := "taskStatus_example" // string | The status of tasks to filter the results. It can take one of 'Open', 'Closed'. This filter will take effect only when threadType is set to Task (optional) (default to "Open")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeedsApi.ListThreads(context.Background()).LimitPosts(limitPosts).Limit(limit).Before(before).After(after).EntityLink(entityLink).UserId(userId).FilterType(filterType).Resolved(resolved).Type_(type_).TaskStatus(taskStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedsApi.ListThreads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListThreads`: ThreadList
    fmt.Fprintf(os.Stdout, "Response from `FeedsApi.ListThreads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListThreadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limitPosts** | **int64** | Limit the number of posts sorted by chronological order (1 to 1000000, default &#x3D; 3) | [default to 3]
 **limit** | **int32** | Limit the number of threads returned. (1 to 1000000, default &#x3D; 10) | [default to 10]
 **before** | **string** | Returns list of threads before this cursor | 
 **after** | **string** | Returns list of threads after this cursor | 
 **entityLink** | **string** | Filter threads by entity link | 
 **userId** | **string** | Filter threads by user id. This filter requires a &#39;filterType&#39; query param. The default filter type is &#39;OWNER&#39;. This filter cannot be combined with the entityLink filter. | 
 **filterType** | **string** | Filter type definition for the user filter. It can take one of &#39;OWNER&#39;, &#39;FOLLOWS&#39;, &#39;MENTIONS&#39;. This must be used with the &#39;user&#39; query param | 
 **resolved** | **bool** | Filter threads by whether they are resolved or not. By default resolved is false | [default to false]
 **type_** | **string** | The type of thread to filter the results. It can take one of &#39;Conversation&#39;, &#39;Task&#39;, &#39;Announcement&#39; | 
 **taskStatus** | **string** | The status of tasks to filter the results. It can take one of &#39;Open&#39;, &#39;Closed&#39;. This filter will take effect only when threadType is set to Task | [default to &quot;Open&quot;]

### Return type

[**ThreadList**](ThreadList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPostOfThread

> PatchPostOfThread(ctx, threadId, postId).Body(body).Execute()

Update post of a thread by `id`.



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
    threadId := "threadId_example" // string | 
    postId := "postId_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} | JsonPatch with array of operations (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeedsApi.PatchPostOfThread(context.Background(), threadId, postId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedsApi.PatchPostOfThread``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** |  | 
**postId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPostOfThreadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** | JsonPatch with array of operations | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchThread

> PatchThread(ctx, id).Body(body).Execute()

Update a thread by `id`.



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
    resp, r, err := apiClient.FeedsApi.PatchThread(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedsApi.PatchThread``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchThreadRequest struct via the builder pattern


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


## ResolveTask

> Thread ResolveTask(ctx, id).ResolveTask(resolveTask).Execute()

Resolve a task



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
    resolveTask := *openapiclient.NewResolveTask("NewValue_example") // ResolveTask |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeedsApi.ResolveTask(context.Background(), id).ResolveTask(resolveTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedsApi.ResolveTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResolveTask`: Thread
    fmt.Fprintf(os.Stdout, "Response from `FeedsApi.ResolveTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resolveTask** | [**ResolveTask**](ResolveTask.md) |  | 

### Return type

[**Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

