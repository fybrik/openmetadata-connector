# ChangeDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldsAdded** | Pointer to [**[]FieldChange**](FieldChange.md) |  | [optional] 
**FieldsDeleted** | Pointer to [**[]FieldChange**](FieldChange.md) |  | [optional] 
**FieldsUpdated** | Pointer to [**[]FieldChange**](FieldChange.md) |  | [optional] 
**PreviousVersion** | Pointer to **float64** |  | [optional] 

## Methods

### NewChangeDescription

`func NewChangeDescription() *ChangeDescription`

NewChangeDescription instantiates a new ChangeDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeDescriptionWithDefaults

`func NewChangeDescriptionWithDefaults() *ChangeDescription`

NewChangeDescriptionWithDefaults instantiates a new ChangeDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldsAdded

`func (o *ChangeDescription) GetFieldsAdded() []FieldChange`

GetFieldsAdded returns the FieldsAdded field if non-nil, zero value otherwise.

### GetFieldsAddedOk

`func (o *ChangeDescription) GetFieldsAddedOk() (*[]FieldChange, bool)`

GetFieldsAddedOk returns a tuple with the FieldsAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsAdded

`func (o *ChangeDescription) SetFieldsAdded(v []FieldChange)`

SetFieldsAdded sets FieldsAdded field to given value.

### HasFieldsAdded

`func (o *ChangeDescription) HasFieldsAdded() bool`

HasFieldsAdded returns a boolean if a field has been set.

### GetFieldsDeleted

`func (o *ChangeDescription) GetFieldsDeleted() []FieldChange`

GetFieldsDeleted returns the FieldsDeleted field if non-nil, zero value otherwise.

### GetFieldsDeletedOk

`func (o *ChangeDescription) GetFieldsDeletedOk() (*[]FieldChange, bool)`

GetFieldsDeletedOk returns a tuple with the FieldsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsDeleted

`func (o *ChangeDescription) SetFieldsDeleted(v []FieldChange)`

SetFieldsDeleted sets FieldsDeleted field to given value.

### HasFieldsDeleted

`func (o *ChangeDescription) HasFieldsDeleted() bool`

HasFieldsDeleted returns a boolean if a field has been set.

### GetFieldsUpdated

`func (o *ChangeDescription) GetFieldsUpdated() []FieldChange`

GetFieldsUpdated returns the FieldsUpdated field if non-nil, zero value otherwise.

### GetFieldsUpdatedOk

`func (o *ChangeDescription) GetFieldsUpdatedOk() (*[]FieldChange, bool)`

GetFieldsUpdatedOk returns a tuple with the FieldsUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsUpdated

`func (o *ChangeDescription) SetFieldsUpdated(v []FieldChange)`

SetFieldsUpdated sets FieldsUpdated field to given value.

### HasFieldsUpdated

`func (o *ChangeDescription) HasFieldsUpdated() bool`

HasFieldsUpdated returns a boolean if a field has been set.

### GetPreviousVersion

`func (o *ChangeDescription) GetPreviousVersion() float64`

GetPreviousVersion returns the PreviousVersion field if non-nil, zero value otherwise.

### GetPreviousVersionOk

`func (o *ChangeDescription) GetPreviousVersionOk() (*float64, bool)`

GetPreviousVersionOk returns a tuple with the PreviousVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersion

`func (o *ChangeDescription) SetPreviousVersion(v float64)`

SetPreviousVersion sets PreviousVersion field to given value.

### HasPreviousVersion

`func (o *ChangeDescription) HasPreviousVersion() bool`

HasPreviousVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


