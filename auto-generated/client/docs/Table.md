# Table

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeDescription** | Pointer to [**ChangeDescription**](ChangeDescription.md) |  | [optional] 
**Columns** | [**[]Column**](Column.md) |  | 
**DataModel** | Pointer to [**DataModel**](DataModel.md) |  | [optional] 
**Database** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**DatabaseSchema** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **map[string]interface{}** |  | [optional] 
**Followers** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**FullyQualifiedName** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Joins** | Pointer to [**TableJoins**](TableJoins.md) |  | [optional] 
**Location** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Name** | **string** |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**ProfileQuery** | Pointer to **string** |  | [optional] 
**ProfileSample** | Pointer to **float64** |  | [optional] 
**SampleData** | Pointer to [**TableData**](TableData.md) |  | [optional] 
**Service** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**TableConstraints** | Pointer to [**[]TableConstraint**](TableConstraint.md) |  | [optional] 
**TablePartition** | Pointer to [**TablePartition**](TablePartition.md) |  | [optional] 
**TableProfile** | Pointer to [**[]TableProfile**](TableProfile.md) |  | [optional] 
**TableQueries** | Pointer to [**[]SQLQuery**](SQLQuery.md) |  | [optional] 
**TableTests** | Pointer to [**[]TableTest**](TableTest.md) |  | [optional] 
**TableType** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagLabel**](TagLabel.md) |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**UsageSummary** | Pointer to [**UsageDetails**](UsageDetails.md) |  | [optional] 
**Version** | Pointer to **float64** |  | [optional] 
**ViewDefinition** | Pointer to **string** |  | [optional] 

## Methods

### NewTable

`func NewTable(columns []Column, id string, name string, ) *Table`

NewTable instantiates a new Table object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableWithDefaults

`func NewTableWithDefaults() *Table`

NewTableWithDefaults instantiates a new Table object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeDescription

`func (o *Table) GetChangeDescription() ChangeDescription`

GetChangeDescription returns the ChangeDescription field if non-nil, zero value otherwise.

### GetChangeDescriptionOk

`func (o *Table) GetChangeDescriptionOk() (*ChangeDescription, bool)`

GetChangeDescriptionOk returns a tuple with the ChangeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDescription

`func (o *Table) SetChangeDescription(v ChangeDescription)`

SetChangeDescription sets ChangeDescription field to given value.

### HasChangeDescription

`func (o *Table) HasChangeDescription() bool`

HasChangeDescription returns a boolean if a field has been set.

### GetColumns

`func (o *Table) GetColumns() []Column`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *Table) GetColumnsOk() (*[]Column, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *Table) SetColumns(v []Column)`

SetColumns sets Columns field to given value.


### GetDataModel

`func (o *Table) GetDataModel() DataModel`

GetDataModel returns the DataModel field if non-nil, zero value otherwise.

### GetDataModelOk

`func (o *Table) GetDataModelOk() (*DataModel, bool)`

GetDataModelOk returns a tuple with the DataModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataModel

`func (o *Table) SetDataModel(v DataModel)`

SetDataModel sets DataModel field to given value.

### HasDataModel

`func (o *Table) HasDataModel() bool`

HasDataModel returns a boolean if a field has been set.

### GetDatabase

`func (o *Table) GetDatabase() EntityReference`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *Table) GetDatabaseOk() (*EntityReference, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *Table) SetDatabase(v EntityReference)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *Table) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetDatabaseSchema

`func (o *Table) GetDatabaseSchema() EntityReference`

GetDatabaseSchema returns the DatabaseSchema field if non-nil, zero value otherwise.

### GetDatabaseSchemaOk

`func (o *Table) GetDatabaseSchemaOk() (*EntityReference, bool)`

GetDatabaseSchemaOk returns a tuple with the DatabaseSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseSchema

`func (o *Table) SetDatabaseSchema(v EntityReference)`

SetDatabaseSchema sets DatabaseSchema field to given value.

### HasDatabaseSchema

`func (o *Table) HasDatabaseSchema() bool`

HasDatabaseSchema returns a boolean if a field has been set.

### GetDeleted

`func (o *Table) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Table) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Table) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Table) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDescription

`func (o *Table) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Table) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Table) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Table) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *Table) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Table) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Table) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Table) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtension

`func (o *Table) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *Table) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *Table) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *Table) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetFollowers

`func (o *Table) GetFollowers() []EntityReference`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *Table) GetFollowersOk() (*[]EntityReference, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *Table) SetFollowers(v []EntityReference)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *Table) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetFullyQualifiedName

`func (o *Table) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *Table) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *Table) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.

### HasFullyQualifiedName

`func (o *Table) HasFullyQualifiedName() bool`

HasFullyQualifiedName returns a boolean if a field has been set.

### GetHref

`func (o *Table) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Table) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Table) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Table) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *Table) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Table) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Table) SetId(v string)`

SetId sets Id field to given value.


### GetJoins

`func (o *Table) GetJoins() TableJoins`

GetJoins returns the Joins field if non-nil, zero value otherwise.

### GetJoinsOk

`func (o *Table) GetJoinsOk() (*TableJoins, bool)`

GetJoinsOk returns a tuple with the Joins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoins

`func (o *Table) SetJoins(v TableJoins)`

SetJoins sets Joins field to given value.

### HasJoins

`func (o *Table) HasJoins() bool`

HasJoins returns a boolean if a field has been set.

### GetLocation

`func (o *Table) GetLocation() EntityReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Table) GetLocationOk() (*EntityReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Table) SetLocation(v EntityReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Table) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *Table) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Table) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Table) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *Table) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Table) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Table) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Table) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProfileQuery

`func (o *Table) GetProfileQuery() string`

GetProfileQuery returns the ProfileQuery field if non-nil, zero value otherwise.

### GetProfileQueryOk

`func (o *Table) GetProfileQueryOk() (*string, bool)`

GetProfileQueryOk returns a tuple with the ProfileQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileQuery

`func (o *Table) SetProfileQuery(v string)`

SetProfileQuery sets ProfileQuery field to given value.

### HasProfileQuery

`func (o *Table) HasProfileQuery() bool`

HasProfileQuery returns a boolean if a field has been set.

### GetProfileSample

`func (o *Table) GetProfileSample() float64`

GetProfileSample returns the ProfileSample field if non-nil, zero value otherwise.

### GetProfileSampleOk

`func (o *Table) GetProfileSampleOk() (*float64, bool)`

GetProfileSampleOk returns a tuple with the ProfileSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileSample

`func (o *Table) SetProfileSample(v float64)`

SetProfileSample sets ProfileSample field to given value.

### HasProfileSample

`func (o *Table) HasProfileSample() bool`

HasProfileSample returns a boolean if a field has been set.

### GetSampleData

`func (o *Table) GetSampleData() TableData`

GetSampleData returns the SampleData field if non-nil, zero value otherwise.

### GetSampleDataOk

`func (o *Table) GetSampleDataOk() (*TableData, bool)`

GetSampleDataOk returns a tuple with the SampleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleData

`func (o *Table) SetSampleData(v TableData)`

SetSampleData sets SampleData field to given value.

### HasSampleData

`func (o *Table) HasSampleData() bool`

HasSampleData returns a boolean if a field has been set.

### GetService

`func (o *Table) GetService() EntityReference`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Table) GetServiceOk() (*EntityReference, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Table) SetService(v EntityReference)`

SetService sets Service field to given value.

### HasService

`func (o *Table) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceType

`func (o *Table) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *Table) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *Table) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *Table) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetTableConstraints

`func (o *Table) GetTableConstraints() []TableConstraint`

GetTableConstraints returns the TableConstraints field if non-nil, zero value otherwise.

### GetTableConstraintsOk

`func (o *Table) GetTableConstraintsOk() (*[]TableConstraint, bool)`

GetTableConstraintsOk returns a tuple with the TableConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableConstraints

`func (o *Table) SetTableConstraints(v []TableConstraint)`

SetTableConstraints sets TableConstraints field to given value.

### HasTableConstraints

`func (o *Table) HasTableConstraints() bool`

HasTableConstraints returns a boolean if a field has been set.

### GetTablePartition

`func (o *Table) GetTablePartition() TablePartition`

GetTablePartition returns the TablePartition field if non-nil, zero value otherwise.

### GetTablePartitionOk

`func (o *Table) GetTablePartitionOk() (*TablePartition, bool)`

GetTablePartitionOk returns a tuple with the TablePartition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTablePartition

`func (o *Table) SetTablePartition(v TablePartition)`

SetTablePartition sets TablePartition field to given value.

### HasTablePartition

`func (o *Table) HasTablePartition() bool`

HasTablePartition returns a boolean if a field has been set.

### GetTableProfile

`func (o *Table) GetTableProfile() []TableProfile`

GetTableProfile returns the TableProfile field if non-nil, zero value otherwise.

### GetTableProfileOk

`func (o *Table) GetTableProfileOk() (*[]TableProfile, bool)`

GetTableProfileOk returns a tuple with the TableProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableProfile

`func (o *Table) SetTableProfile(v []TableProfile)`

SetTableProfile sets TableProfile field to given value.

### HasTableProfile

`func (o *Table) HasTableProfile() bool`

HasTableProfile returns a boolean if a field has been set.

### GetTableQueries

`func (o *Table) GetTableQueries() []SQLQuery`

GetTableQueries returns the TableQueries field if non-nil, zero value otherwise.

### GetTableQueriesOk

`func (o *Table) GetTableQueriesOk() (*[]SQLQuery, bool)`

GetTableQueriesOk returns a tuple with the TableQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableQueries

`func (o *Table) SetTableQueries(v []SQLQuery)`

SetTableQueries sets TableQueries field to given value.

### HasTableQueries

`func (o *Table) HasTableQueries() bool`

HasTableQueries returns a boolean if a field has been set.

### GetTableTests

`func (o *Table) GetTableTests() []TableTest`

GetTableTests returns the TableTests field if non-nil, zero value otherwise.

### GetTableTestsOk

`func (o *Table) GetTableTestsOk() (*[]TableTest, bool)`

GetTableTestsOk returns a tuple with the TableTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableTests

`func (o *Table) SetTableTests(v []TableTest)`

SetTableTests sets TableTests field to given value.

### HasTableTests

`func (o *Table) HasTableTests() bool`

HasTableTests returns a boolean if a field has been set.

### GetTableType

`func (o *Table) GetTableType() string`

GetTableType returns the TableType field if non-nil, zero value otherwise.

### GetTableTypeOk

`func (o *Table) GetTableTypeOk() (*string, bool)`

GetTableTypeOk returns a tuple with the TableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableType

`func (o *Table) SetTableType(v string)`

SetTableType sets TableType field to given value.

### HasTableType

`func (o *Table) HasTableType() bool`

HasTableType returns a boolean if a field has been set.

### GetTags

`func (o *Table) GetTags() []TagLabel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Table) GetTagsOk() (*[]TagLabel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Table) SetTags(v []TagLabel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Table) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Table) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Table) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Table) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Table) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Table) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Table) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Table) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Table) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUsageSummary

`func (o *Table) GetUsageSummary() UsageDetails`

GetUsageSummary returns the UsageSummary field if non-nil, zero value otherwise.

### GetUsageSummaryOk

`func (o *Table) GetUsageSummaryOk() (*UsageDetails, bool)`

GetUsageSummaryOk returns a tuple with the UsageSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageSummary

`func (o *Table) SetUsageSummary(v UsageDetails)`

SetUsageSummary sets UsageSummary field to given value.

### HasUsageSummary

`func (o *Table) HasUsageSummary() bool`

HasUsageSummary returns a boolean if a field has been set.

### GetVersion

`func (o *Table) GetVersion() float64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Table) GetVersionOk() (*float64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Table) SetVersion(v float64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Table) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetViewDefinition

`func (o *Table) GetViewDefinition() string`

GetViewDefinition returns the ViewDefinition field if non-nil, zero value otherwise.

### GetViewDefinitionOk

`func (o *Table) GetViewDefinitionOk() (*string, bool)`

GetViewDefinitionOk returns a tuple with the ViewDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewDefinition

`func (o *Table) SetViewDefinition(v string)`

SetViewDefinition sets ViewDefinition field to given value.

### HasViewDefinition

`func (o *Table) HasViewDefinition() bool`

HasViewDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


