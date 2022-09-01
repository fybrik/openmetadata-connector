# ColumnLineage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromColumns** | Pointer to **[]string** |  | [optional] 
**Function** | Pointer to **string** |  | [optional] 
**ToColumn** | Pointer to **string** |  | [optional] 

## Methods

### NewColumnLineage

`func NewColumnLineage() *ColumnLineage`

NewColumnLineage instantiates a new ColumnLineage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColumnLineageWithDefaults

`func NewColumnLineageWithDefaults() *ColumnLineage`

NewColumnLineageWithDefaults instantiates a new ColumnLineage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromColumns

`func (o *ColumnLineage) GetFromColumns() []string`

GetFromColumns returns the FromColumns field if non-nil, zero value otherwise.

### GetFromColumnsOk

`func (o *ColumnLineage) GetFromColumnsOk() (*[]string, bool)`

GetFromColumnsOk returns a tuple with the FromColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromColumns

`func (o *ColumnLineage) SetFromColumns(v []string)`

SetFromColumns sets FromColumns field to given value.

### HasFromColumns

`func (o *ColumnLineage) HasFromColumns() bool`

HasFromColumns returns a boolean if a field has been set.

### GetFunction

`func (o *ColumnLineage) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ColumnLineage) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ColumnLineage) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *ColumnLineage) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetToColumn

`func (o *ColumnLineage) GetToColumn() string`

GetToColumn returns the ToColumn field if non-nil, zero value otherwise.

### GetToColumnOk

`func (o *ColumnLineage) GetToColumnOk() (*string, bool)`

GetToColumnOk returns a tuple with the ToColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToColumn

`func (o *ColumnLineage) SetToColumn(v string)`

SetToColumn sets ToColumn field to given value.

### HasToColumn

`func (o *ColumnLineage) HasToColumn() bool`

HasToColumn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


