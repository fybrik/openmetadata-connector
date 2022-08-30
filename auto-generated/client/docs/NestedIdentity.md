# NestedIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Child** | Pointer to [**NestedIdentity**](NestedIdentity.md) |  | [optional] 
**Field** | Pointer to [**Text**](Text.md) |  | [optional] 
**Fragment** | Pointer to **bool** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 

## Methods

### NewNestedIdentity

`func NewNestedIdentity() *NestedIdentity`

NewNestedIdentity instantiates a new NestedIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedIdentityWithDefaults

`func NewNestedIdentityWithDefaults() *NestedIdentity`

NewNestedIdentityWithDefaults instantiates a new NestedIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChild

`func (o *NestedIdentity) GetChild() NestedIdentity`

GetChild returns the Child field if non-nil, zero value otherwise.

### GetChildOk

`func (o *NestedIdentity) GetChildOk() (*NestedIdentity, bool)`

GetChildOk returns a tuple with the Child field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChild

`func (o *NestedIdentity) SetChild(v NestedIdentity)`

SetChild sets Child field to given value.

### HasChild

`func (o *NestedIdentity) HasChild() bool`

HasChild returns a boolean if a field has been set.

### GetField

`func (o *NestedIdentity) GetField() Text`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *NestedIdentity) GetFieldOk() (*Text, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *NestedIdentity) SetField(v Text)`

SetField sets Field field to given value.

### HasField

`func (o *NestedIdentity) HasField() bool`

HasField returns a boolean if a field has been set.

### GetFragment

`func (o *NestedIdentity) GetFragment() bool`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *NestedIdentity) GetFragmentOk() (*bool, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *NestedIdentity) SetFragment(v bool)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *NestedIdentity) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetOffset

`func (o *NestedIdentity) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *NestedIdentity) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *NestedIdentity) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *NestedIdentity) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


