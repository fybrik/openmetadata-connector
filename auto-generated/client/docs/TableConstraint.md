# TableConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to **[]string** |  | [optional] 
**ConstraintType** | Pointer to **string** |  | [optional] 

## Methods

### NewTableConstraint

`func NewTableConstraint() *TableConstraint`

NewTableConstraint instantiates a new TableConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableConstraintWithDefaults

`func NewTableConstraintWithDefaults() *TableConstraint`

NewTableConstraintWithDefaults instantiates a new TableConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *TableConstraint) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TableConstraint) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TableConstraint) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *TableConstraint) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetConstraintType

`func (o *TableConstraint) GetConstraintType() string`

GetConstraintType returns the ConstraintType field if non-nil, zero value otherwise.

### GetConstraintTypeOk

`func (o *TableConstraint) GetConstraintTypeOk() (*string, bool)`

GetConstraintTypeOk returns a tuple with the ConstraintType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintType

`func (o *TableConstraint) SetConstraintType(v string)`

SetConstraintType sets ConstraintType field to given value.

### HasConstraintType

`func (o *TableConstraint) HasConstraintType() bool`

HasConstraintType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


