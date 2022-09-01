# AggregationProfileShardResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fragment** | Pointer to **bool** |  | [optional] 
**ProfileResults** | Pointer to [**[]ProfileResult**](ProfileResult.md) |  | [optional] 

## Methods

### NewAggregationProfileShardResult

`func NewAggregationProfileShardResult() *AggregationProfileShardResult`

NewAggregationProfileShardResult instantiates a new AggregationProfileShardResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationProfileShardResultWithDefaults

`func NewAggregationProfileShardResultWithDefaults() *AggregationProfileShardResult`

NewAggregationProfileShardResultWithDefaults instantiates a new AggregationProfileShardResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFragment

`func (o *AggregationProfileShardResult) GetFragment() bool`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *AggregationProfileShardResult) GetFragmentOk() (*bool, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *AggregationProfileShardResult) SetFragment(v bool)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *AggregationProfileShardResult) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetProfileResults

`func (o *AggregationProfileShardResult) GetProfileResults() []ProfileResult`

GetProfileResults returns the ProfileResults field if non-nil, zero value otherwise.

### GetProfileResultsOk

`func (o *AggregationProfileShardResult) GetProfileResultsOk() (*[]ProfileResult, bool)`

GetProfileResultsOk returns a tuple with the ProfileResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileResults

`func (o *AggregationProfileShardResult) SetProfileResults(v []ProfileResult)`

SetProfileResults sets ProfileResults field to given value.

### HasProfileResults

`func (o *AggregationProfileShardResult) HasProfileResults() bool`

HasProfileResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


