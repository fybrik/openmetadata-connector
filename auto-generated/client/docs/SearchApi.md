# \SearchApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSuggestedEntities**](SearchApi.md#GetSuggestedEntities) | **Get** /v1/search/suggest | Suggest Entities
[**SearchEntitiesWithQuery**](SearchApi.md#SearchEntitiesWithQuery) | **Get** /v1/search/query | Search entities



## GetSuggestedEntities

> Suggest GetSuggestedEntities(ctx).Q(q).Index(index).Field(field).Execute()

Suggest Entities



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
    q := "q_example" // string | Suggest API can be used to auto-fill the entities name while use is typing search text <br/> 1. To get suggest results pass q=us or q=user etc.. <br/> 2. Do not add any wild-cards such as * like in search api <br/> 3. suggest api is a prefix suggestion <br/>
    index := "index_example" // string |  (optional) (default to "table_search_index")
    field := "field_example" // string |  (optional) (default to "suggest")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetSuggestedEntities(context.Background()).Q(q).Index(index).Field(field).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetSuggestedEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSuggestedEntities`: Suggest
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetSuggestedEntities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSuggestedEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Suggest API can be used to auto-fill the entities name while use is typing search text &lt;br/&gt; 1. To get suggest results pass q&#x3D;us or q&#x3D;user etc.. &lt;br/&gt; 2. Do not add any wild-cards such as * like in search api &lt;br/&gt; 3. suggest api is a prefix suggestion &lt;br/&gt; | 
 **index** | **string** |  | [default to &quot;table_search_index&quot;]
 **field** | **string** |  | [default to &quot;suggest&quot;]

### Return type

[**Suggest**](Suggest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchEntitiesWithQuery

> SearchResponse SearchEntitiesWithQuery(ctx).Q(q).Index(index).Deleted(deleted).From(from).Size(size).SortField(sortField).SortOrder(sortOrder).TrackTotalHits(trackTotalHits).Execute()

Search entities



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
    q := "q_example" // string | Search Query Text, Pass *text* for substring match; Pass without wildcards for exact match. <br/> 1. For listing all tables or topics pass q=* <br/>2. For search tables or topics pass q=*search_term* <br/>3. For searching field names such as search by column_name pass q=column_names:address <br/>4. For searching by tag names pass q=tags:user.email <br/>5. When user selects a filter pass q=query_text AND tags:user.email AND platform:MYSQL <br/>6. Search with multiple values of same filter q=tags:user.email AND tags:user.address <br/> logic operators such as AND and OR must be in uppercase 
    index := "index_example" // string | ElasticSearch Index name, defaults to table_search_index (optional) (default to "table_search_index")
    deleted := true // bool | Filter documents by deleted param. By default deleted is false (optional) (default to false)
    from := int32(56) // int32 | From field to paginate the results, defaults to 0 (optional) (default to 0)
    size := int32(56) // int32 | Size field to limit the no.of results returned, defaults to 10 (optional) (default to 10)
    sortField := "sortField_example" // string | Sort the search results by field, available fields to sort weekly_stats , daily_stats, monthly_stats, last_updated_timestamp (optional)
    sortOrder := "sortOrder_example" // string | Sort order asc for ascending or desc for descending, defaults to desc (optional) (default to "desc")
    trackTotalHits := true // bool | Track Total Hits (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchEntitiesWithQuery(context.Background()).Q(q).Index(index).Deleted(deleted).From(from).Size(size).SortField(sortField).SortOrder(sortOrder).TrackTotalHits(trackTotalHits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchEntitiesWithQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchEntitiesWithQuery`: SearchResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchEntitiesWithQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchEntitiesWithQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Search Query Text, Pass *text* for substring match; Pass without wildcards for exact match. &lt;br/&gt; 1. For listing all tables or topics pass q&#x3D;* &lt;br/&gt;2. For search tables or topics pass q&#x3D;*search_term* &lt;br/&gt;3. For searching field names such as search by column_name pass q&#x3D;column_names:address &lt;br/&gt;4. For searching by tag names pass q&#x3D;tags:user.email &lt;br/&gt;5. When user selects a filter pass q&#x3D;query_text AND tags:user.email AND platform:MYSQL &lt;br/&gt;6. Search with multiple values of same filter q&#x3D;tags:user.email AND tags:user.address &lt;br/&gt; logic operators such as AND and OR must be in uppercase  | 
 **index** | **string** | ElasticSearch Index name, defaults to table_search_index | [default to &quot;table_search_index&quot;]
 **deleted** | **bool** | Filter documents by deleted param. By default deleted is false | [default to false]
 **from** | **int32** | From field to paginate the results, defaults to 0 | [default to 0]
 **size** | **int32** | Size field to limit the no.of results returned, defaults to 10 | [default to 10]
 **sortField** | **string** | Sort the search results by field, available fields to sort weekly_stats , daily_stats, monthly_stats, last_updated_timestamp | 
 **sortOrder** | **string** | Sort order asc for ascending or desc for descending, defaults to desc | [default to &quot;desc&quot;]
 **trackTotalHits** | **bool** | Track Total Hits | [default to false]

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

