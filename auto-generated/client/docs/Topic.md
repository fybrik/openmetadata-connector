# Topic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeDescription** | Pointer to [**ChangeDescription**](ChangeDescription.md) |  | [optional] 
**CleanupPolicies** | Pointer to **[]string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **map[string]interface{}** |  | [optional] 
**Followers** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**FullyQualifiedName** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**MaximumMessageSize** | Pointer to **int32** |  | [optional] 
**MinimumInSyncReplicas** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Partitions** | **int32** |  | 
**ReplicationFactor** | Pointer to **int32** |  | [optional] 
**RetentionSize** | Pointer to **float64** |  | [optional] 
**RetentionTime** | Pointer to **float64** |  | [optional] 
**SampleData** | Pointer to [**TopicSampleData**](TopicSampleData.md) |  | [optional] 
**SchemaText** | Pointer to **string** |  | [optional] 
**SchemaType** | Pointer to **string** |  | [optional] 
**Service** | [**EntityReference**](EntityReference.md) |  | 
**ServiceType** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagLabel**](TagLabel.md) |  | [optional] 
**TopicConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **float64** |  | [optional] 

## Methods

### NewTopic

`func NewTopic(id string, name string, partitions int32, service EntityReference, ) *Topic`

NewTopic instantiates a new Topic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicWithDefaults

`func NewTopicWithDefaults() *Topic`

NewTopicWithDefaults instantiates a new Topic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeDescription

`func (o *Topic) GetChangeDescription() ChangeDescription`

GetChangeDescription returns the ChangeDescription field if non-nil, zero value otherwise.

### GetChangeDescriptionOk

`func (o *Topic) GetChangeDescriptionOk() (*ChangeDescription, bool)`

GetChangeDescriptionOk returns a tuple with the ChangeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDescription

`func (o *Topic) SetChangeDescription(v ChangeDescription)`

SetChangeDescription sets ChangeDescription field to given value.

### HasChangeDescription

`func (o *Topic) HasChangeDescription() bool`

HasChangeDescription returns a boolean if a field has been set.

### GetCleanupPolicies

`func (o *Topic) GetCleanupPolicies() []string`

GetCleanupPolicies returns the CleanupPolicies field if non-nil, zero value otherwise.

### GetCleanupPoliciesOk

`func (o *Topic) GetCleanupPoliciesOk() (*[]string, bool)`

GetCleanupPoliciesOk returns a tuple with the CleanupPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupPolicies

`func (o *Topic) SetCleanupPolicies(v []string)`

SetCleanupPolicies sets CleanupPolicies field to given value.

### HasCleanupPolicies

`func (o *Topic) HasCleanupPolicies() bool`

HasCleanupPolicies returns a boolean if a field has been set.

### GetDeleted

`func (o *Topic) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Topic) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Topic) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Topic) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDescription

`func (o *Topic) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Topic) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Topic) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Topic) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *Topic) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Topic) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Topic) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Topic) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtension

`func (o *Topic) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *Topic) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *Topic) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *Topic) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetFollowers

`func (o *Topic) GetFollowers() []EntityReference`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *Topic) GetFollowersOk() (*[]EntityReference, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *Topic) SetFollowers(v []EntityReference)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *Topic) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetFullyQualifiedName

`func (o *Topic) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *Topic) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *Topic) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.

### HasFullyQualifiedName

`func (o *Topic) HasFullyQualifiedName() bool`

HasFullyQualifiedName returns a boolean if a field has been set.

### GetHref

`func (o *Topic) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Topic) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Topic) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Topic) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *Topic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Topic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Topic) SetId(v string)`

SetId sets Id field to given value.


### GetMaximumMessageSize

`func (o *Topic) GetMaximumMessageSize() int32`

GetMaximumMessageSize returns the MaximumMessageSize field if non-nil, zero value otherwise.

### GetMaximumMessageSizeOk

`func (o *Topic) GetMaximumMessageSizeOk() (*int32, bool)`

GetMaximumMessageSizeOk returns a tuple with the MaximumMessageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMessageSize

`func (o *Topic) SetMaximumMessageSize(v int32)`

SetMaximumMessageSize sets MaximumMessageSize field to given value.

### HasMaximumMessageSize

`func (o *Topic) HasMaximumMessageSize() bool`

HasMaximumMessageSize returns a boolean if a field has been set.

### GetMinimumInSyncReplicas

`func (o *Topic) GetMinimumInSyncReplicas() int32`

GetMinimumInSyncReplicas returns the MinimumInSyncReplicas field if non-nil, zero value otherwise.

### GetMinimumInSyncReplicasOk

`func (o *Topic) GetMinimumInSyncReplicasOk() (*int32, bool)`

GetMinimumInSyncReplicasOk returns a tuple with the MinimumInSyncReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumInSyncReplicas

`func (o *Topic) SetMinimumInSyncReplicas(v int32)`

SetMinimumInSyncReplicas sets MinimumInSyncReplicas field to given value.

### HasMinimumInSyncReplicas

`func (o *Topic) HasMinimumInSyncReplicas() bool`

HasMinimumInSyncReplicas returns a boolean if a field has been set.

### GetName

`func (o *Topic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Topic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Topic) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *Topic) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Topic) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Topic) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Topic) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPartitions

`func (o *Topic) GetPartitions() int32`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *Topic) GetPartitionsOk() (*int32, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *Topic) SetPartitions(v int32)`

SetPartitions sets Partitions field to given value.


### GetReplicationFactor

`func (o *Topic) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *Topic) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *Topic) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *Topic) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetRetentionSize

`func (o *Topic) GetRetentionSize() float64`

GetRetentionSize returns the RetentionSize field if non-nil, zero value otherwise.

### GetRetentionSizeOk

`func (o *Topic) GetRetentionSizeOk() (*float64, bool)`

GetRetentionSizeOk returns a tuple with the RetentionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionSize

`func (o *Topic) SetRetentionSize(v float64)`

SetRetentionSize sets RetentionSize field to given value.

### HasRetentionSize

`func (o *Topic) HasRetentionSize() bool`

HasRetentionSize returns a boolean if a field has been set.

### GetRetentionTime

`func (o *Topic) GetRetentionTime() float64`

GetRetentionTime returns the RetentionTime field if non-nil, zero value otherwise.

### GetRetentionTimeOk

`func (o *Topic) GetRetentionTimeOk() (*float64, bool)`

GetRetentionTimeOk returns a tuple with the RetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionTime

`func (o *Topic) SetRetentionTime(v float64)`

SetRetentionTime sets RetentionTime field to given value.

### HasRetentionTime

`func (o *Topic) HasRetentionTime() bool`

HasRetentionTime returns a boolean if a field has been set.

### GetSampleData

`func (o *Topic) GetSampleData() TopicSampleData`

GetSampleData returns the SampleData field if non-nil, zero value otherwise.

### GetSampleDataOk

`func (o *Topic) GetSampleDataOk() (*TopicSampleData, bool)`

GetSampleDataOk returns a tuple with the SampleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleData

`func (o *Topic) SetSampleData(v TopicSampleData)`

SetSampleData sets SampleData field to given value.

### HasSampleData

`func (o *Topic) HasSampleData() bool`

HasSampleData returns a boolean if a field has been set.

### GetSchemaText

`func (o *Topic) GetSchemaText() string`

GetSchemaText returns the SchemaText field if non-nil, zero value otherwise.

### GetSchemaTextOk

`func (o *Topic) GetSchemaTextOk() (*string, bool)`

GetSchemaTextOk returns a tuple with the SchemaText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaText

`func (o *Topic) SetSchemaText(v string)`

SetSchemaText sets SchemaText field to given value.

### HasSchemaText

`func (o *Topic) HasSchemaText() bool`

HasSchemaText returns a boolean if a field has been set.

### GetSchemaType

`func (o *Topic) GetSchemaType() string`

GetSchemaType returns the SchemaType field if non-nil, zero value otherwise.

### GetSchemaTypeOk

`func (o *Topic) GetSchemaTypeOk() (*string, bool)`

GetSchemaTypeOk returns a tuple with the SchemaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaType

`func (o *Topic) SetSchemaType(v string)`

SetSchemaType sets SchemaType field to given value.

### HasSchemaType

`func (o *Topic) HasSchemaType() bool`

HasSchemaType returns a boolean if a field has been set.

### GetService

`func (o *Topic) GetService() EntityReference`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Topic) GetServiceOk() (*EntityReference, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Topic) SetService(v EntityReference)`

SetService sets Service field to given value.


### GetServiceType

`func (o *Topic) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *Topic) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *Topic) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *Topic) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetTags

`func (o *Topic) GetTags() []TagLabel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Topic) GetTagsOk() (*[]TagLabel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Topic) SetTags(v []TagLabel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Topic) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTopicConfig

`func (o *Topic) GetTopicConfig() map[string]interface{}`

GetTopicConfig returns the TopicConfig field if non-nil, zero value otherwise.

### GetTopicConfigOk

`func (o *Topic) GetTopicConfigOk() (*map[string]interface{}, bool)`

GetTopicConfigOk returns a tuple with the TopicConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicConfig

`func (o *Topic) SetTopicConfig(v map[string]interface{})`

SetTopicConfig sets TopicConfig field to given value.

### HasTopicConfig

`func (o *Topic) HasTopicConfig() bool`

HasTopicConfig returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Topic) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Topic) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Topic) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Topic) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Topic) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Topic) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Topic) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Topic) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetVersion

`func (o *Topic) GetVersion() float64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Topic) GetVersionOk() (*float64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Topic) SetVersion(v float64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Topic) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


