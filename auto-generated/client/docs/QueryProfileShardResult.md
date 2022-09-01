# QueryProfileShardResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectorResult** | Pointer to [**CollectorResult**](CollectorResult.md) |  | [optional] 
**Fragment** | Pointer to **bool** |  | [optional] 
**QueryResults** | Pointer to [**[]ProfileResult**](ProfileResult.md) |  | [optional] 
**RewriteTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewQueryProfileShardResult

`func NewQueryProfileShardResult() *QueryProfileShardResult`

NewQueryProfileShardResult instantiates a new QueryProfileShardResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryProfileShardResultWithDefaults

`func NewQueryProfileShardResultWithDefaults() *QueryProfileShardResult`

NewQueryProfileShardResultWithDefaults instantiates a new QueryProfileShardResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectorResult

`func (o *QueryProfileShardResult) GetCollectorResult() CollectorResult`

GetCollectorResult returns the CollectorResult field if non-nil, zero value otherwise.

### GetCollectorResultOk

`func (o *QueryProfileShardResult) GetCollectorResultOk() (*CollectorResult, bool)`

GetCollectorResultOk returns a tuple with the CollectorResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectorResult

`func (o *QueryProfileShardResult) SetCollectorResult(v CollectorResult)`

SetCollectorResult sets CollectorResult field to given value.

### HasCollectorResult

`func (o *QueryProfileShardResult) HasCollectorResult() bool`

HasCollectorResult returns a boolean if a field has been set.

### GetFragment

`func (o *QueryProfileShardResult) GetFragment() bool`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *QueryProfileShardResult) GetFragmentOk() (*bool, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *QueryProfileShardResult) SetFragment(v bool)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *QueryProfileShardResult) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetQueryResults

`func (o *QueryProfileShardResult) GetQueryResults() []ProfileResult`

GetQueryResults returns the QueryResults field if non-nil, zero value otherwise.

### GetQueryResultsOk

`func (o *QueryProfileShardResult) GetQueryResultsOk() (*[]ProfileResult, bool)`

GetQueryResultsOk returns a tuple with the QueryResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryResults

`func (o *QueryProfileShardResult) SetQueryResults(v []ProfileResult)`

SetQueryResults sets QueryResults field to given value.

### HasQueryResults

`func (o *QueryProfileShardResult) HasQueryResults() bool`

HasQueryResults returns a boolean if a field has been set.

### GetRewriteTime

`func (o *QueryProfileShardResult) GetRewriteTime() int64`

GetRewriteTime returns the RewriteTime field if non-nil, zero value otherwise.

### GetRewriteTimeOk

`func (o *QueryProfileShardResult) GetRewriteTimeOk() (*int64, bool)`

GetRewriteTimeOk returns a tuple with the RewriteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewriteTime

`func (o *QueryProfileShardResult) SetRewriteTime(v int64)`

SetRewriteTime sets RewriteTime field to given value.

### HasRewriteTime

`func (o *QueryProfileShardResult) HasRewriteTime() bool`

HasRewriteTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


