# DataModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to [**[]Column**](Column.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**GeneratedAt** | Pointer to **time.Time** |  | [optional] 
**ModelType** | **string** |  | 
**Path** | Pointer to **string** |  | [optional] 
**RawSql** | Pointer to **string** |  | [optional] 
**Sql** | **string** |  | 
**Upstream** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDataModel

`func NewDataModel(modelType string, sql string, ) *DataModel`

NewDataModel instantiates a new DataModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataModelWithDefaults

`func NewDataModelWithDefaults() *DataModel`

NewDataModelWithDefaults instantiates a new DataModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *DataModel) GetColumns() []Column`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *DataModel) GetColumnsOk() (*[]Column, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *DataModel) SetColumns(v []Column)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *DataModel) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetDescription

`func (o *DataModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *DataModel) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *DataModel) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *DataModel) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *DataModel) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetModelType

`func (o *DataModel) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *DataModel) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *DataModel) SetModelType(v string)`

SetModelType sets ModelType field to given value.


### GetPath

`func (o *DataModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DataModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DataModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DataModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRawSql

`func (o *DataModel) GetRawSql() string`

GetRawSql returns the RawSql field if non-nil, zero value otherwise.

### GetRawSqlOk

`func (o *DataModel) GetRawSqlOk() (*string, bool)`

GetRawSqlOk returns a tuple with the RawSql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSql

`func (o *DataModel) SetRawSql(v string)`

SetRawSql sets RawSql field to given value.

### HasRawSql

`func (o *DataModel) HasRawSql() bool`

HasRawSql returns a boolean if a field has been set.

### GetSql

`func (o *DataModel) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *DataModel) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *DataModel) SetSql(v string)`

SetSql sets Sql field to given value.


### GetUpstream

`func (o *DataModel) GetUpstream() []string`

GetUpstream returns the Upstream field if non-nil, zero value otherwise.

### GetUpstreamOk

`func (o *DataModel) GetUpstreamOk() (*[]string, bool)`

GetUpstreamOk returns a tuple with the Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstream

`func (o *DataModel) SetUpstream(v []string)`

SetUpstream sets Upstream field to given value.

### HasUpstream

`func (o *DataModel) HasUpstream() bool`

HasUpstream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


