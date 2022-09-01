# ProfileResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DebugInfo** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Fragment** | Pointer to **bool** |  | [optional] 
**LuceneDescription** | Pointer to **string** |  | [optional] 
**QueryName** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**TimeBreakdown** | Pointer to **map[string]int64** |  | [optional] 

## Methods

### NewProfileResult

`func NewProfileResult() *ProfileResult`

NewProfileResult instantiates a new ProfileResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileResultWithDefaults

`func NewProfileResultWithDefaults() *ProfileResult`

NewProfileResultWithDefaults instantiates a new ProfileResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebugInfo

`func (o *ProfileResult) GetDebugInfo() map[string]map[string]interface{}`

GetDebugInfo returns the DebugInfo field if non-nil, zero value otherwise.

### GetDebugInfoOk

`func (o *ProfileResult) GetDebugInfoOk() (*map[string]map[string]interface{}, bool)`

GetDebugInfoOk returns a tuple with the DebugInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugInfo

`func (o *ProfileResult) SetDebugInfo(v map[string]map[string]interface{})`

SetDebugInfo sets DebugInfo field to given value.

### HasDebugInfo

`func (o *ProfileResult) HasDebugInfo() bool`

HasDebugInfo returns a boolean if a field has been set.

### GetFragment

`func (o *ProfileResult) GetFragment() bool`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *ProfileResult) GetFragmentOk() (*bool, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *ProfileResult) SetFragment(v bool)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *ProfileResult) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetLuceneDescription

`func (o *ProfileResult) GetLuceneDescription() string`

GetLuceneDescription returns the LuceneDescription field if non-nil, zero value otherwise.

### GetLuceneDescriptionOk

`func (o *ProfileResult) GetLuceneDescriptionOk() (*string, bool)`

GetLuceneDescriptionOk returns a tuple with the LuceneDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLuceneDescription

`func (o *ProfileResult) SetLuceneDescription(v string)`

SetLuceneDescription sets LuceneDescription field to given value.

### HasLuceneDescription

`func (o *ProfileResult) HasLuceneDescription() bool`

HasLuceneDescription returns a boolean if a field has been set.

### GetQueryName

`func (o *ProfileResult) GetQueryName() string`

GetQueryName returns the QueryName field if non-nil, zero value otherwise.

### GetQueryNameOk

`func (o *ProfileResult) GetQueryNameOk() (*string, bool)`

GetQueryNameOk returns a tuple with the QueryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryName

`func (o *ProfileResult) SetQueryName(v string)`

SetQueryName sets QueryName field to given value.

### HasQueryName

`func (o *ProfileResult) HasQueryName() bool`

HasQueryName returns a boolean if a field has been set.

### GetTime

`func (o *ProfileResult) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ProfileResult) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ProfileResult) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *ProfileResult) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimeBreakdown

`func (o *ProfileResult) GetTimeBreakdown() map[string]int64`

GetTimeBreakdown returns the TimeBreakdown field if non-nil, zero value otherwise.

### GetTimeBreakdownOk

`func (o *ProfileResult) GetTimeBreakdownOk() (*map[string]int64, bool)`

GetTimeBreakdownOk returns a tuple with the TimeBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBreakdown

`func (o *ProfileResult) SetTimeBreakdown(v map[string]int64)`

SetTimeBreakdown sets TimeBreakdown field to given value.

### HasTimeBreakdown

`func (o *ProfileResult) HasTimeBreakdown() bool`

HasTimeBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


