# ProfileShardResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationProfileResults** | Pointer to [**AggregationProfileShardResult**](AggregationProfileShardResult.md) |  | [optional] 
**QueryProfileResults** | Pointer to [**[]QueryProfileShardResult**](QueryProfileShardResult.md) |  | [optional] 

## Methods

### NewProfileShardResult

`func NewProfileShardResult() *ProfileShardResult`

NewProfileShardResult instantiates a new ProfileShardResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileShardResultWithDefaults

`func NewProfileShardResultWithDefaults() *ProfileShardResult`

NewProfileShardResultWithDefaults instantiates a new ProfileShardResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationProfileResults

`func (o *ProfileShardResult) GetAggregationProfileResults() AggregationProfileShardResult`

GetAggregationProfileResults returns the AggregationProfileResults field if non-nil, zero value otherwise.

### GetAggregationProfileResultsOk

`func (o *ProfileShardResult) GetAggregationProfileResultsOk() (*AggregationProfileShardResult, bool)`

GetAggregationProfileResultsOk returns a tuple with the AggregationProfileResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationProfileResults

`func (o *ProfileShardResult) SetAggregationProfileResults(v AggregationProfileShardResult)`

SetAggregationProfileResults sets AggregationProfileResults field to given value.

### HasAggregationProfileResults

`func (o *ProfileShardResult) HasAggregationProfileResults() bool`

HasAggregationProfileResults returns a boolean if a field has been set.

### GetQueryProfileResults

`func (o *ProfileShardResult) GetQueryProfileResults() []QueryProfileShardResult`

GetQueryProfileResults returns the QueryProfileResults field if non-nil, zero value otherwise.

### GetQueryProfileResultsOk

`func (o *ProfileShardResult) GetQueryProfileResultsOk() (*[]QueryProfileShardResult, bool)`

GetQueryProfileResultsOk returns a tuple with the QueryProfileResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryProfileResults

`func (o *ProfileShardResult) SetQueryProfileResults(v []QueryProfileShardResult)`

SetQueryProfileResults sets QueryProfileResults field to given value.

### HasQueryProfileResults

`func (o *ProfileShardResult) HasQueryProfileResults() bool`

HasQueryProfileResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


