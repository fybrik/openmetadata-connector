# Aggregations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsMap** | Pointer to [**map[string]Aggregation**](Aggregation.md) |  | [optional] 
**Fragment** | Pointer to **bool** |  | [optional] 

## Methods

### NewAggregations

`func NewAggregations() *Aggregations`

NewAggregations instantiates a new Aggregations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationsWithDefaults

`func NewAggregationsWithDefaults() *Aggregations`

NewAggregationsWithDefaults instantiates a new Aggregations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsMap

`func (o *Aggregations) GetAsMap() map[string]Aggregation`

GetAsMap returns the AsMap field if non-nil, zero value otherwise.

### GetAsMapOk

`func (o *Aggregations) GetAsMapOk() (*map[string]Aggregation, bool)`

GetAsMapOk returns a tuple with the AsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsMap

`func (o *Aggregations) SetAsMap(v map[string]Aggregation)`

SetAsMap sets AsMap field to given value.

### HasAsMap

`func (o *Aggregations) HasAsMap() bool`

HasAsMap returns a boolean if a field has been set.

### GetFragment

`func (o *Aggregations) GetFragment() bool`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *Aggregations) GetFragmentOk() (*bool, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *Aggregations) SetFragment(v bool)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *Aggregations) HasFragment() bool`

HasFragment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


