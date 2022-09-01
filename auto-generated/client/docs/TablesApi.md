# \TablesApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddColumnTest**](TablesApi.md#AddColumnTest) | **Put** /v1/tables/{id}/columnTest | Add column test cases
[**AddCustomMetric**](TablesApi.md#AddCustomMetric) | **Put** /v1/tables/{id}/customMetric | Add column custom metrics
[**AddDataModel**](TablesApi.md#AddDataModel) | **Put** /v1/tables/{id}/dataModel | Add data modeling information to a table
[**AddDataProfiler**](TablesApi.md#AddDataProfiler) | **Put** /v1/tables/{id}/tableProfile | Add table profile data
[**AddFollowerToTable**](TablesApi.md#AddFollowerToTable) | **Put** /v1/tables/{id}/followers | Add a follower
[**AddLocationToTable**](TablesApi.md#AddLocationToTable) | **Put** /v1/tables/{id}/location | Add a location
[**AddSampleData**](TablesApi.md#AddSampleData) | **Put** /v1/tables/{id}/sampleData | Add sample data
[**AddTableJoinInfo**](TablesApi.md#AddTableJoinInfo) | **Put** /v1/tables/{id}/joins | Add table join information
[**AddTableQuery**](TablesApi.md#AddTableQuery) | **Put** /v1/tables/{id}/tableQuery | Add table query data
[**AddTableTest**](TablesApi.md#AddTableTest) | **Put** /v1/tables/{id}/tableTest | Add table test cases
[**CreateOrUpdateTable**](TablesApi.md#CreateOrUpdateTable) | **Put** /v1/tables | Create or update a table
[**CreateTable**](TablesApi.md#CreateTable) | **Post** /v1/tables | Create a table
[**DeleteColumnTest**](TablesApi.md#DeleteColumnTest) | **Delete** /v1/tables/{id}/columnTest/{columnName}/{columnTestType} | delete column test case
[**DeleteCustomMetric**](TablesApi.md#DeleteCustomMetric) | **Delete** /v1/tables/{id}/customMetric/{columnName}/{customMetricName} | delete custom metric from a column
[**DeleteFollower**](TablesApi.md#DeleteFollower) | **Delete** /v1/tables/{id}/followers/{userId} | Remove a follower
[**DeleteLocation1**](TablesApi.md#DeleteLocation1) | **Delete** /v1/tables/{id}/location | Remove the location
[**DeleteTable**](TablesApi.md#DeleteTable) | **Delete** /v1/tables/{id} | Delete a table
[**DeleteTableTest**](TablesApi.md#DeleteTableTest) | **Delete** /v1/tables/{id}/tableTest/{tableTestType} | delete table test case
[**GetSpecificDatabaseVersion1**](TablesApi.md#GetSpecificDatabaseVersion1) | **Get** /v1/tables/{id}/versions/{version} | Get a version of the table
[**GetTableByFQN**](TablesApi.md#GetTableByFQN) | **Get** /v1/tables/name/{fqn} | Get a table by name
[**GetTableByID**](TablesApi.md#GetTableByID) | **Get** /v1/tables/{id} | Get a table
[**ListAllTableVersion**](TablesApi.md#ListAllTableVersion) | **Get** /v1/tables/{id}/versions | List table versions
[**ListTables**](TablesApi.md#ListTables) | **Get** /v1/tables | List tables
[**PatchTable**](TablesApi.md#PatchTable) | **Patch** /v1/tables/{id} | Update a table



## AddColumnTest

> Table AddColumnTest(ctx, id).CreateColumnTest(createColumnTest).Execute()

Add column test cases



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
    id := "id_example" // string | Id of the table
    createColumnTest := *openapiclient.NewCreateColumnTest("ColumnName_example", *openapiclient.NewColumnTestCase()) // CreateColumnTest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.AddColumnTest(context.Background(), id).CreateColumnTest(createColumnTest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.AddColumnTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddColumnTest`: Table
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.AddColumnTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddColumnTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createColumnTest** | [**CreateColumnTest**](CreateColumnTest.md) |  | 

### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddCustomMetric

> Table AddCustomMetric(ctx, id).CreateCustomMetric(createCustomMetric).Execute()

Add column custom metrics



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
    id := "id_example" // string | Id of the table
    createCustomMetric := *openapiclient.NewCreateCustomMetric("ColumnName_example", "Expression_example", "Name_example") // CreateCustomMetric |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.AddCustomMetric(context.Background(), id).CreateCustomMetric(createCustomMetric).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.AddCustomMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCustomMetric`: Table
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.AddCustomMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCustomMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCustomMetric** | [**CreateCustomMetric**](CreateCustomMetric.md) |  | 

### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddDataModel

> Table AddDataModel(ctx, id).DataModel(dataModel).Execute()

Add data modeling information to a table



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
    id := "id_example" // string | Id of the table
    dataModel := *openapiclient.NewDataModel("ModelType_example", "Sql_example") // DataModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.AddDataModel(context.Background(), id).DataModel(dataModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.AddDataModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDataModel`: Table
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.AddDataModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDataModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataModel** | [**DataModel**](DataModel.md) |  | 

### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddDataProfiler

> Table AddDataProfiler(ctx, id).TableProfile(tableProfile).Execute()

Add table profile data



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
    id := "id_example" // string | Id of the table
    tableProfile := *openapiclient.NewTableProfile() // TableProfile |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.AddDataProfiler(context.Background(), id).TableProfile(tableProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.AddDataProfiler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDataProfiler`: Table
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.AddDataProfiler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDataProfilerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tableProfile** | [**TableProfile**](TableProfile.md) |  | 

### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddFollowerToTable

> ChangeEvent AddFollowerToTable(ctx, id).Body(body).Execute()

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
    id := "id_example" // string | Id of the table
    body := "body_example" // string | Id of the user to be added as follower (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.AddFollowerToTable(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.AddFollowerToTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddFollowerToTable`: ChangeEvent
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.AddFollowerToTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFollowerToTableRequest struct via the builder pattern


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


## AddLocationToTable

> Table AddLocationToTable(ctx, id).Body(body).Execute()

Add a location



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
    id := "id_example" // string | Id of the table
    body := "body_example" // string | Id of the location to be added (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.AddLocationToTable(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.AddLocationToTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLocationToTable`: Table
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.AddLocationToTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLocationToTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Id of the location to be added | 

### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSampleData

> Table AddSampleData(ctx, id).TableData(tableData).Execute()

Add sample data



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
    id := "id_example" // string | Id of the table
    tableData := *openapiclient.NewTableData() // TableData |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.AddSampleData(context.Background(), id).TableData(tableData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.AddSampleData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSampleData`: Table
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.AddSampleData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSampleDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tableData** | [**TableData**](TableData.md) |  | 

### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTableJoinInfo

> Table AddTableJoinInfo(ctx, id).TableJoins(tableJoins).Execute()

Add table join information



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
    id := "id_example" // string | Id of the table
    tableJoins := *openapiclient.NewTableJoins() // TableJoins |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.AddTableJoinInfo(context.Background(), id).TableJoins(tableJoins).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.AddTableJoinInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTableJoinInfo`: Table
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.AddTableJoinInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTableJoinInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tableJoins** | [**TableJoins**](TableJoins.md) |  | 

### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTableQuery

> Table AddTableQuery(ctx, id).SQLQuery(sQLQuery).Execute()

Add table query data



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
    id := "id_example" // string | Id of the table
    sQLQuery := *openapiclient.NewSQLQuery() // SQLQuery |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.AddTableQuery(context.Background(), id).SQLQuery(sQLQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.AddTableQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTableQuery`: Table
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.AddTableQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTableQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sQLQuery** | [**SQLQuery**](SQLQuery.md) |  | 

### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTableTest

> Table AddTableTest(ctx, id).CreateTableTest(createTableTest).Execute()

Add table test cases



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
    id := "id_example" // string | Id of the table
    createTableTest := *openapiclient.NewCreateTableTest(*openapiclient.NewTableTestCase()) // CreateTableTest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.AddTableTest(context.Background(), id).CreateTableTest(createTableTest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.AddTableTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTableTest`: Table
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.AddTableTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTableTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTableTest** | [**CreateTableTest**](CreateTableTest.md) |  | 

### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateTable

> Table CreateOrUpdateTable(ctx).CreateTable(createTable).Execute()

Create or update a table



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
    createTable := *openapiclient.NewCreateTable([]openapiclient.Column{*openapiclient.NewColumn("DataType_example", "Name_example")}, *openapiclient.NewEntityReference("Id_example", "Type_example"), "Name_example") // CreateTable |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.CreateOrUpdateTable(context.Background()).CreateTable(createTable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.CreateOrUpdateTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateTable`: Table
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.CreateOrUpdateTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTable** | [**CreateTable**](CreateTable.md) |  | 

### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTable

> Table CreateTable(ctx).CreateTable(createTable).Execute()

Create a table



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
    createTable := *openapiclient.NewCreateTable([]openapiclient.Column{*openapiclient.NewColumn("DataType_example", "Name_example")}, *openapiclient.NewEntityReference("Id_example", "Type_example"), "Name_example") // CreateTable |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.CreateTable(context.Background()).CreateTable(createTable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.CreateTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTable`: Table
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.CreateTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTable** | [**CreateTable**](CreateTable.md) |  | 

### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteColumnTest

> Table DeleteColumnTest(ctx, id, columnName, columnTestType).Execute()

delete column test case



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
    id := "id_example" // string | Id of the table
    columnName := "columnName_example" // string | column of the table
    columnTestType := "columnTestType_example" // string | column Test Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.DeleteColumnTest(context.Background(), id, columnName, columnTestType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.DeleteColumnTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteColumnTest`: Table
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.DeleteColumnTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the table | 
**columnName** | **string** | column of the table | 
**columnTestType** | **string** | column Test Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteColumnTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomMetric

> Table DeleteCustomMetric(ctx, id, columnName, customMetricName).Execute()

delete custom metric from a column



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
    id := "id_example" // string | Id of the table
    columnName := "columnName_example" // string | column of the table
    customMetricName := "customMetricName_example" // string | column Test Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.DeleteCustomMetric(context.Background(), id, columnName, customMetricName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.DeleteCustomMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCustomMetric`: Table
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.DeleteCustomMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the table | 
**columnName** | **string** | column of the table | 
**customMetricName** | **string** | column Test Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFollower

> ChangeEvent DeleteFollower(ctx, id, userId).Execute()

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
    id := "id_example" // string | Id of the table
    userId := "userId_example" // string | Id of the user being removed as follower

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.DeleteFollower(context.Background(), id, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.DeleteFollower``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFollower`: ChangeEvent
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.DeleteFollower`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the table | 
**userId** | **string** | Id of the user being removed as follower | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFollowerRequest struct via the builder pattern


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


## DeleteLocation1

> Table DeleteLocation1(ctx, id).Execute()

Remove the location



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
    id := "id_example" // string | Id of the table

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.DeleteLocation1(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.DeleteLocation1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLocation1`: Table
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.DeleteLocation1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocation1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTable

> DeleteTable(ctx, id).HardDelete(hardDelete).Execute()

Delete a table



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
    id := "id_example" // string | Id of the table
    hardDelete := true // bool | Hard delete the entity. (Default = `false`) (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.DeleteTable(context.Background(), id).HardDelete(hardDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.DeleteTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTableRequest struct via the builder pattern


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


## DeleteTableTest

> Table DeleteTableTest(ctx, id, tableTestType).Execute()

delete table test case



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
    id := "id_example" // string | Id of the table
    tableTestType := "tableTestType_example" // string | Table Test Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.DeleteTableTest(context.Background(), id, tableTestType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.DeleteTableTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTableTest`: Table
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.DeleteTableTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the table | 
**tableTestType** | **string** | Table Test Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTableTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificDatabaseVersion1

> Table GetSpecificDatabaseVersion1(ctx, id, version).Execute()

Get a version of the table



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
    id := "id_example" // string | table Id
    version := "0.1 or 1.1" // string | table version number in the form `major`.`minor`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.GetSpecificDatabaseVersion1(context.Background(), id, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.GetSpecificDatabaseVersion1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificDatabaseVersion1`: Table
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.GetSpecificDatabaseVersion1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | table Id | 
**version** | **string** | table version number in the form &#x60;major&#x60;.&#x60;minor&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificDatabaseVersion1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTableByFQN

> Table GetTableByFQN(ctx, fqn).Fields(fields).Include(include).Execute()

Get a table by name



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
    fqn := "fqn_example" // string | Fully qualified name of the table
    fields := "tableConstraints,tablePartition,usageSummary,owner,profileSample,profileQuery,customMetrics,tags,followers,joins,sampleData,viewDefinition,tableProfile,location,tableQueries,dataModel,tests,extension" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.GetTableByFQN(context.Background(), fqn).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.GetTableByFQN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTableByFQN`: Table
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.GetTableByFQN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqn** | **string** | Fully qualified name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTableByFQNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTableByID

> Table GetTableByID(ctx, id).Fields(fields).Include(include).Execute()

Get a table



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
    id := "id_example" // string | table Id
    fields := "tableConstraints,tablePartition,usageSummary,owner,profileSample,profileQuery,customMetrics,tags,followers,joins,sampleData,viewDefinition,tableProfile,location,tableQueries,dataModel,tests,extension" // string | Fields requested in the returned resource (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.GetTableByID(context.Background(), id).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.GetTableByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTableByID`: Table
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.GetTableByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | table Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTableByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Fields requested in the returned resource | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllTableVersion

> EntityHistory ListAllTableVersion(ctx, id).Execute()

List table versions



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
    id := "id_example" // string | table Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.ListAllTableVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.ListAllTableVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllTableVersion`: EntityHistory
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.ListAllTableVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | table Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllTableVersionRequest struct via the builder pattern


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


## ListTables

> TableList ListTables(ctx).Fields(fields).Database(database).Limit(limit).Before(before).After(after).Include(include).Execute()

List tables



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
    fields := "tableConstraints,tablePartition,usageSummary,owner,profileSample,profileQuery,customMetrics,tags,followers,joins,sampleData,viewDefinition,tableProfile,location,tableQueries,dataModel,tests,extension" // string | Fields requested in the returned resource (optional)
    database := "snowflakeWestCoast.financeDB" // string | Filter tables by database fully qualified name (optional)
    limit := int32(56) // int32 | Limit the number tables returned. (1 to 1000000, default = 10)  (optional) (default to 10)
    before := "before_example" // string | Returns list of tables before this cursor (optional)
    after := "after_example" // string | Returns list of tables after this cursor (optional)
    include := "include_example" // string | Include all, deleted, or non-deleted entities. (optional) (default to "non-deleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.ListTables(context.Background()).Fields(fields).Database(database).Limit(limit).Before(before).After(after).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.ListTables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTables`: TableList
    fmt.Fprintf(os.Stdout, "Response from `TablesApi.ListTables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Fields requested in the returned resource | 
 **database** | **string** | Filter tables by database fully qualified name | 
 **limit** | **int32** | Limit the number tables returned. (1 to 1000000, default &#x3D; 10)  | [default to 10]
 **before** | **string** | Returns list of tables before this cursor | 
 **after** | **string** | Returns list of tables after this cursor | 
 **include** | **string** | Include all, deleted, or non-deleted entities. | [default to &quot;non-deleted&quot;]

### Return type

[**TableList**](TableList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTable

> PatchTable(ctx, id).RequestBody(requestBody).Execute()

Update a table



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
    id := "id_example" // string | Id of the table
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | JsonPatch with array of operations (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablesApi.PatchTable(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablesApi.PatchTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]map[string]interface{}** | JsonPatch with array of operations | 

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

