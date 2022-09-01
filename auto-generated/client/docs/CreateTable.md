# CreateTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | [**[]Column**](Column.md) |  | 
**DatabaseSchema** | [**EntityReference**](EntityReference.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**ProfileQuery** | Pointer to **string** |  | [optional] 
**ProfileSample** | Pointer to **float64** |  | [optional] 
**TableConstraints** | Pointer to [**[]TableConstraint**](TableConstraint.md) |  | [optional] 
**TablePartition** | Pointer to [**TablePartition**](TablePartition.md) |  | [optional] 
**TableType** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagLabel**](TagLabel.md) |  | [optional] 
**ViewDefinition** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateTable

`func NewCreateTable(columns []Column, databaseSchema EntityReference, name string, ) *CreateTable`

NewCreateTable instantiates a new CreateTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTableWithDefaults

`func NewCreateTableWithDefaults() *CreateTable`

NewCreateTableWithDefaults instantiates a new CreateTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *CreateTable) GetColumns() []Column`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *CreateTable) GetColumnsOk() (*[]Column, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *CreateTable) SetColumns(v []Column)`

SetColumns sets Columns field to given value.


### GetDatabaseSchema

`func (o *CreateTable) GetDatabaseSchema() EntityReference`

GetDatabaseSchema returns the DatabaseSchema field if non-nil, zero value otherwise.

### GetDatabaseSchemaOk

`func (o *CreateTable) GetDatabaseSchemaOk() (*EntityReference, bool)`

GetDatabaseSchemaOk returns a tuple with the DatabaseSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseSchema

`func (o *CreateTable) SetDatabaseSchema(v EntityReference)`

SetDatabaseSchema sets DatabaseSchema field to given value.


### GetDescription

`func (o *CreateTable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateTable) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateTable) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateTable) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateTable) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtension

`func (o *CreateTable) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *CreateTable) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *CreateTable) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *CreateTable) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetName

`func (o *CreateTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTable) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *CreateTable) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateTable) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateTable) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateTable) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProfileQuery

`func (o *CreateTable) GetProfileQuery() string`

GetProfileQuery returns the ProfileQuery field if non-nil, zero value otherwise.

### GetProfileQueryOk

`func (o *CreateTable) GetProfileQueryOk() (*string, bool)`

GetProfileQueryOk returns a tuple with the ProfileQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileQuery

`func (o *CreateTable) SetProfileQuery(v string)`

SetProfileQuery sets ProfileQuery field to given value.

### HasProfileQuery

`func (o *CreateTable) HasProfileQuery() bool`

HasProfileQuery returns a boolean if a field has been set.

### GetProfileSample

`func (o *CreateTable) GetProfileSample() float64`

GetProfileSample returns the ProfileSample field if non-nil, zero value otherwise.

### GetProfileSampleOk

`func (o *CreateTable) GetProfileSampleOk() (*float64, bool)`

GetProfileSampleOk returns a tuple with the ProfileSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileSample

`func (o *CreateTable) SetProfileSample(v float64)`

SetProfileSample sets ProfileSample field to given value.

### HasProfileSample

`func (o *CreateTable) HasProfileSample() bool`

HasProfileSample returns a boolean if a field has been set.

### GetTableConstraints

`func (o *CreateTable) GetTableConstraints() []TableConstraint`

GetTableConstraints returns the TableConstraints field if non-nil, zero value otherwise.

### GetTableConstraintsOk

`func (o *CreateTable) GetTableConstraintsOk() (*[]TableConstraint, bool)`

GetTableConstraintsOk returns a tuple with the TableConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableConstraints

`func (o *CreateTable) SetTableConstraints(v []TableConstraint)`

SetTableConstraints sets TableConstraints field to given value.

### HasTableConstraints

`func (o *CreateTable) HasTableConstraints() bool`

HasTableConstraints returns a boolean if a field has been set.

### GetTablePartition

`func (o *CreateTable) GetTablePartition() TablePartition`

GetTablePartition returns the TablePartition field if non-nil, zero value otherwise.

### GetTablePartitionOk

`func (o *CreateTable) GetTablePartitionOk() (*TablePartition, bool)`

GetTablePartitionOk returns a tuple with the TablePartition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTablePartition

`func (o *CreateTable) SetTablePartition(v TablePartition)`

SetTablePartition sets TablePartition field to given value.

### HasTablePartition

`func (o *CreateTable) HasTablePartition() bool`

HasTablePartition returns a boolean if a field has been set.

### GetTableType

`func (o *CreateTable) GetTableType() string`

GetTableType returns the TableType field if non-nil, zero value otherwise.

### GetTableTypeOk

`func (o *CreateTable) GetTableTypeOk() (*string, bool)`

GetTableTypeOk returns a tuple with the TableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableType

`func (o *CreateTable) SetTableType(v string)`

SetTableType sets TableType field to given value.

### HasTableType

`func (o *CreateTable) HasTableType() bool`

HasTableType returns a boolean if a field has been set.

### GetTags

`func (o *CreateTable) GetTags() []TagLabel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateTable) GetTagsOk() (*[]TagLabel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateTable) SetTags(v []TagLabel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateTable) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetViewDefinition

`func (o *CreateTable) GetViewDefinition() string`

GetViewDefinition returns the ViewDefinition field if non-nil, zero value otherwise.

### GetViewDefinitionOk

`func (o *CreateTable) GetViewDefinitionOk() (*string, bool)`

GetViewDefinitionOk returns a tuple with the ViewDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewDefinition

`func (o *CreateTable) SetViewDefinition(v string)`

SetViewDefinition sets ViewDefinition field to given value.

### HasViewDefinition

`func (o *CreateTable) HasViewDefinition() bool`

HasViewDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


