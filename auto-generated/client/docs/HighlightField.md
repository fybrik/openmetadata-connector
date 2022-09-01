# HighlightField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fragment** | Pointer to **bool** |  | [optional] 
**Fragments** | Pointer to [**[]Text**](Text.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewHighlightField

`func NewHighlightField() *HighlightField`

NewHighlightField instantiates a new HighlightField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighlightFieldWithDefaults

`func NewHighlightFieldWithDefaults() *HighlightField`

NewHighlightFieldWithDefaults instantiates a new HighlightField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFragment

`func (o *HighlightField) GetFragment() bool`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *HighlightField) GetFragmentOk() (*bool, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *HighlightField) SetFragment(v bool)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *HighlightField) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetFragments

`func (o *HighlightField) GetFragments() []Text`

GetFragments returns the Fragments field if non-nil, zero value otherwise.

### GetFragmentsOk

`func (o *HighlightField) GetFragmentsOk() (*[]Text, bool)`

GetFragmentsOk returns a tuple with the Fragments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragments

`func (o *HighlightField) SetFragments(v []Text)`

SetFragments sets Fragments field to given value.

### HasFragments

`func (o *HighlightField) HasFragments() bool`

HasFragments returns a boolean if a field has been set.

### GetName

`func (o *HighlightField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HighlightField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HighlightField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HighlightField) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


