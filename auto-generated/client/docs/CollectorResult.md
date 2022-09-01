# CollectorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fragment** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProfiledChildren** | Pointer to [**[]CollectorResult**](CollectorResult.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewCollectorResult

`func NewCollectorResult() *CollectorResult`

NewCollectorResult instantiates a new CollectorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectorResultWithDefaults

`func NewCollectorResultWithDefaults() *CollectorResult`

NewCollectorResultWithDefaults instantiates a new CollectorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFragment

`func (o *CollectorResult) GetFragment() bool`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *CollectorResult) GetFragmentOk() (*bool, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *CollectorResult) SetFragment(v bool)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *CollectorResult) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetName

`func (o *CollectorResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CollectorResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CollectorResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CollectorResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProfiledChildren

`func (o *CollectorResult) GetProfiledChildren() []CollectorResult`

GetProfiledChildren returns the ProfiledChildren field if non-nil, zero value otherwise.

### GetProfiledChildrenOk

`func (o *CollectorResult) GetProfiledChildrenOk() (*[]CollectorResult, bool)`

GetProfiledChildrenOk returns a tuple with the ProfiledChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiledChildren

`func (o *CollectorResult) SetProfiledChildren(v []CollectorResult)`

SetProfiledChildren sets ProfiledChildren field to given value.

### HasProfiledChildren

`func (o *CollectorResult) HasProfiledChildren() bool`

HasProfiledChildren returns a boolean if a field has been set.

### GetReason

`func (o *CollectorResult) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CollectorResult) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CollectorResult) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CollectorResult) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetTime

`func (o *CollectorResult) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CollectorResult) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CollectorResult) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *CollectorResult) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


