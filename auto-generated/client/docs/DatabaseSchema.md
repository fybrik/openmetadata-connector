# DatabaseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeDescription** | Pointer to [**ChangeDescription**](ChangeDescription.md) |  | [optional] 
**Database** | [**EntityReference**](EntityReference.md) |  | 
**Deleted** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **map[string]interface{}** |  | [optional] 
**Followers** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**FullyQualifiedName** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Service** | [**EntityReference**](EntityReference.md) |  | 
**ServiceType** | Pointer to **string** |  | [optional] 
**Tables** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**Tags** | Pointer to [**[]TagLabel**](TagLabel.md) |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**UsageSummary** | Pointer to [**UsageDetails**](UsageDetails.md) |  | [optional] 
**Version** | Pointer to **float64** |  | [optional] 

## Methods

### NewDatabaseSchema

`func NewDatabaseSchema(database EntityReference, name string, service EntityReference, ) *DatabaseSchema`

NewDatabaseSchema instantiates a new DatabaseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseSchemaWithDefaults

`func NewDatabaseSchemaWithDefaults() *DatabaseSchema`

NewDatabaseSchemaWithDefaults instantiates a new DatabaseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeDescription

`func (o *DatabaseSchema) GetChangeDescription() ChangeDescription`

GetChangeDescription returns the ChangeDescription field if non-nil, zero value otherwise.

### GetChangeDescriptionOk

`func (o *DatabaseSchema) GetChangeDescriptionOk() (*ChangeDescription, bool)`

GetChangeDescriptionOk returns a tuple with the ChangeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDescription

`func (o *DatabaseSchema) SetChangeDescription(v ChangeDescription)`

SetChangeDescription sets ChangeDescription field to given value.

### HasChangeDescription

`func (o *DatabaseSchema) HasChangeDescription() bool`

HasChangeDescription returns a boolean if a field has been set.

### GetDatabase

`func (o *DatabaseSchema) GetDatabase() EntityReference`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *DatabaseSchema) GetDatabaseOk() (*EntityReference, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *DatabaseSchema) SetDatabase(v EntityReference)`

SetDatabase sets Database field to given value.


### GetDeleted

`func (o *DatabaseSchema) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DatabaseSchema) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DatabaseSchema) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DatabaseSchema) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDescription

`func (o *DatabaseSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatabaseSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatabaseSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatabaseSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *DatabaseSchema) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DatabaseSchema) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DatabaseSchema) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DatabaseSchema) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtension

`func (o *DatabaseSchema) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *DatabaseSchema) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *DatabaseSchema) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *DatabaseSchema) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetFollowers

`func (o *DatabaseSchema) GetFollowers() []EntityReference`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *DatabaseSchema) GetFollowersOk() (*[]EntityReference, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *DatabaseSchema) SetFollowers(v []EntityReference)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *DatabaseSchema) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetFullyQualifiedName

`func (o *DatabaseSchema) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *DatabaseSchema) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *DatabaseSchema) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.

### HasFullyQualifiedName

`func (o *DatabaseSchema) HasFullyQualifiedName() bool`

HasFullyQualifiedName returns a boolean if a field has been set.

### GetHref

`func (o *DatabaseSchema) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DatabaseSchema) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DatabaseSchema) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DatabaseSchema) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *DatabaseSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DatabaseSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DatabaseSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseSchema) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *DatabaseSchema) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DatabaseSchema) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DatabaseSchema) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DatabaseSchema) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetService

`func (o *DatabaseSchema) GetService() EntityReference`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DatabaseSchema) GetServiceOk() (*EntityReference, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DatabaseSchema) SetService(v EntityReference)`

SetService sets Service field to given value.


### GetServiceType

`func (o *DatabaseSchema) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *DatabaseSchema) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *DatabaseSchema) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *DatabaseSchema) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetTables

`func (o *DatabaseSchema) GetTables() []EntityReference`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *DatabaseSchema) GetTablesOk() (*[]EntityReference, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *DatabaseSchema) SetTables(v []EntityReference)`

SetTables sets Tables field to given value.

### HasTables

`func (o *DatabaseSchema) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetTags

`func (o *DatabaseSchema) GetTags() []TagLabel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DatabaseSchema) GetTagsOk() (*[]TagLabel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DatabaseSchema) SetTags(v []TagLabel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DatabaseSchema) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DatabaseSchema) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DatabaseSchema) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DatabaseSchema) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DatabaseSchema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DatabaseSchema) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DatabaseSchema) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DatabaseSchema) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DatabaseSchema) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUsageSummary

`func (o *DatabaseSchema) GetUsageSummary() UsageDetails`

GetUsageSummary returns the UsageSummary field if non-nil, zero value otherwise.

### GetUsageSummaryOk

`func (o *DatabaseSchema) GetUsageSummaryOk() (*UsageDetails, bool)`

GetUsageSummaryOk returns a tuple with the UsageSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageSummary

`func (o *DatabaseSchema) SetUsageSummary(v UsageDetails)`

SetUsageSummary sets UsageSummary field to given value.

### HasUsageSummary

`func (o *DatabaseSchema) HasUsageSummary() bool`

HasUsageSummary returns a boolean if a field has been set.

### GetVersion

`func (o *DatabaseSchema) GetVersion() float64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DatabaseSchema) GetVersionOk() (*float64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DatabaseSchema) SetVersion(v float64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DatabaseSchema) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


