# Column

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArrayDataType** | Pointer to **string** |  | [optional] 
**Children** | Pointer to [**[]Column**](Column.md) |  | [optional] 
**ColumnTests** | Pointer to [**[]ColumnTest**](ColumnTest.md) |  | [optional] 
**Constraint** | Pointer to **string** |  | [optional] 
**CustomMetrics** | Pointer to [**[]CustomMetric**](CustomMetric.md) |  | [optional] 
**DataLength** | Pointer to **int32** |  | [optional] 
**DataType** | **string** |  | 
**DataTypeDisplay** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**FullyQualifiedName** | Pointer to **string** |  | [optional] 
**JsonSchema** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**OrdinalPosition** | Pointer to **int32** |  | [optional] 
**Precision** | Pointer to **int32** |  | [optional] 
**Scale** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to [**[]TagLabel**](TagLabel.md) |  | [optional] 

## Methods

### NewColumn

`func NewColumn(dataType string, name string, ) *Column`

NewColumn instantiates a new Column object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColumnWithDefaults

`func NewColumnWithDefaults() *Column`

NewColumnWithDefaults instantiates a new Column object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArrayDataType

`func (o *Column) GetArrayDataType() string`

GetArrayDataType returns the ArrayDataType field if non-nil, zero value otherwise.

### GetArrayDataTypeOk

`func (o *Column) GetArrayDataTypeOk() (*string, bool)`

GetArrayDataTypeOk returns a tuple with the ArrayDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayDataType

`func (o *Column) SetArrayDataType(v string)`

SetArrayDataType sets ArrayDataType field to given value.

### HasArrayDataType

`func (o *Column) HasArrayDataType() bool`

HasArrayDataType returns a boolean if a field has been set.

### GetChildren

`func (o *Column) GetChildren() []Column`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Column) GetChildrenOk() (*[]Column, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Column) SetChildren(v []Column)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Column) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetColumnTests

`func (o *Column) GetColumnTests() []ColumnTest`

GetColumnTests returns the ColumnTests field if non-nil, zero value otherwise.

### GetColumnTestsOk

`func (o *Column) GetColumnTestsOk() (*[]ColumnTest, bool)`

GetColumnTestsOk returns a tuple with the ColumnTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnTests

`func (o *Column) SetColumnTests(v []ColumnTest)`

SetColumnTests sets ColumnTests field to given value.

### HasColumnTests

`func (o *Column) HasColumnTests() bool`

HasColumnTests returns a boolean if a field has been set.

### GetConstraint

`func (o *Column) GetConstraint() string`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *Column) GetConstraintOk() (*string, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *Column) SetConstraint(v string)`

SetConstraint sets Constraint field to given value.

### HasConstraint

`func (o *Column) HasConstraint() bool`

HasConstraint returns a boolean if a field has been set.

### GetCustomMetrics

`func (o *Column) GetCustomMetrics() []CustomMetric`

GetCustomMetrics returns the CustomMetrics field if non-nil, zero value otherwise.

### GetCustomMetricsOk

`func (o *Column) GetCustomMetricsOk() (*[]CustomMetric, bool)`

GetCustomMetricsOk returns a tuple with the CustomMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetrics

`func (o *Column) SetCustomMetrics(v []CustomMetric)`

SetCustomMetrics sets CustomMetrics field to given value.

### HasCustomMetrics

`func (o *Column) HasCustomMetrics() bool`

HasCustomMetrics returns a boolean if a field has been set.

### GetDataLength

`func (o *Column) GetDataLength() int32`

GetDataLength returns the DataLength field if non-nil, zero value otherwise.

### GetDataLengthOk

`func (o *Column) GetDataLengthOk() (*int32, bool)`

GetDataLengthOk returns a tuple with the DataLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLength

`func (o *Column) SetDataLength(v int32)`

SetDataLength sets DataLength field to given value.

### HasDataLength

`func (o *Column) HasDataLength() bool`

HasDataLength returns a boolean if a field has been set.

### GetDataType

`func (o *Column) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *Column) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *Column) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetDataTypeDisplay

`func (o *Column) GetDataTypeDisplay() string`

GetDataTypeDisplay returns the DataTypeDisplay field if non-nil, zero value otherwise.

### GetDataTypeDisplayOk

`func (o *Column) GetDataTypeDisplayOk() (*string, bool)`

GetDataTypeDisplayOk returns a tuple with the DataTypeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeDisplay

`func (o *Column) SetDataTypeDisplay(v string)`

SetDataTypeDisplay sets DataTypeDisplay field to given value.

### HasDataTypeDisplay

`func (o *Column) HasDataTypeDisplay() bool`

HasDataTypeDisplay returns a boolean if a field has been set.

### GetDescription

`func (o *Column) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Column) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Column) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Column) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *Column) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Column) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Column) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Column) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFullyQualifiedName

`func (o *Column) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *Column) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *Column) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.

### HasFullyQualifiedName

`func (o *Column) HasFullyQualifiedName() bool`

HasFullyQualifiedName returns a boolean if a field has been set.

### GetJsonSchema

`func (o *Column) GetJsonSchema() string`

GetJsonSchema returns the JsonSchema field if non-nil, zero value otherwise.

### GetJsonSchemaOk

`func (o *Column) GetJsonSchemaOk() (*string, bool)`

GetJsonSchemaOk returns a tuple with the JsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonSchema

`func (o *Column) SetJsonSchema(v string)`

SetJsonSchema sets JsonSchema field to given value.

### HasJsonSchema

`func (o *Column) HasJsonSchema() bool`

HasJsonSchema returns a boolean if a field has been set.

### GetName

`func (o *Column) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Column) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Column) SetName(v string)`

SetName sets Name field to given value.


### GetOrdinalPosition

`func (o *Column) GetOrdinalPosition() int32`

GetOrdinalPosition returns the OrdinalPosition field if non-nil, zero value otherwise.

### GetOrdinalPositionOk

`func (o *Column) GetOrdinalPositionOk() (*int32, bool)`

GetOrdinalPositionOk returns a tuple with the OrdinalPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdinalPosition

`func (o *Column) SetOrdinalPosition(v int32)`

SetOrdinalPosition sets OrdinalPosition field to given value.

### HasOrdinalPosition

`func (o *Column) HasOrdinalPosition() bool`

HasOrdinalPosition returns a boolean if a field has been set.

### GetPrecision

`func (o *Column) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *Column) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *Column) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *Column) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetScale

`func (o *Column) GetScale() int32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *Column) GetScaleOk() (*int32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *Column) SetScale(v int32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *Column) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetTags

`func (o *Column) GetTags() []TagLabel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Column) GetTagsOk() (*[]TagLabel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Column) SetTags(v []TagLabel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Column) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


