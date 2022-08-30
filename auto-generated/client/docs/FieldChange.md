# FieldChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**NewValue** | Pointer to **interface{}** |  | [optional] 
**OldValue** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewFieldChange

`func NewFieldChange() *FieldChange`

NewFieldChange instantiates a new FieldChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldChangeWithDefaults

`func NewFieldChangeWithDefaults() *FieldChange`

NewFieldChangeWithDefaults instantiates a new FieldChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FieldChange) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldChange) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldChange) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FieldChange) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewValue

`func (o *FieldChange) GetNewValue() interface{}`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *FieldChange) GetNewValueOk() (*interface{}, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *FieldChange) SetNewValue(v interface{})`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *FieldChange) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### GetOldValue

`func (o *FieldChange) GetOldValue() interface{}`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *FieldChange) GetOldValueOk() (*interface{}, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *FieldChange) SetOldValue(v interface{})`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *FieldChange) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


