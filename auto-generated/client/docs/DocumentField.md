# DocumentField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fragment** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 
**Values** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewDocumentField

`func NewDocumentField() *DocumentField`

NewDocumentField instantiates a new DocumentField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentFieldWithDefaults

`func NewDocumentFieldWithDefaults() *DocumentField`

NewDocumentFieldWithDefaults instantiates a new DocumentField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFragment

`func (o *DocumentField) GetFragment() bool`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *DocumentField) GetFragmentOk() (*bool, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *DocumentField) SetFragment(v bool)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *DocumentField) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetName

`func (o *DocumentField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocumentField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *DocumentField) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DocumentField) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DocumentField) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *DocumentField) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValues

`func (o *DocumentField) GetValues() []map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *DocumentField) GetValuesOk() (*[]map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *DocumentField) SetValues(v []map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *DocumentField) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


