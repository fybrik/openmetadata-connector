# CreateTopic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CleanupPolicies** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **map[string]interface{}** |  | [optional] 
**MaximumMessageSize** | Pointer to **int32** |  | [optional] 
**MinimumInSyncReplicas** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Partitions** | **int32** |  | 
**ReplicationFactor** | Pointer to **int32** |  | [optional] 
**RetentionSize** | Pointer to **float64** |  | [optional] 
**RetentionTime** | Pointer to **float64** |  | [optional] 
**SchemaText** | Pointer to **string** |  | [optional] 
**SchemaType** | Pointer to **string** |  | [optional] 
**Service** | [**EntityReference**](EntityReference.md) |  | 
**Tags** | Pointer to [**[]TagLabel**](TagLabel.md) |  | [optional] 
**TopicConfig** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateTopic

`func NewCreateTopic(name string, partitions int32, service EntityReference, ) *CreateTopic`

NewCreateTopic instantiates a new CreateTopic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTopicWithDefaults

`func NewCreateTopicWithDefaults() *CreateTopic`

NewCreateTopicWithDefaults instantiates a new CreateTopic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanupPolicies

`func (o *CreateTopic) GetCleanupPolicies() []string`

GetCleanupPolicies returns the CleanupPolicies field if non-nil, zero value otherwise.

### GetCleanupPoliciesOk

`func (o *CreateTopic) GetCleanupPoliciesOk() (*[]string, bool)`

GetCleanupPoliciesOk returns a tuple with the CleanupPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupPolicies

`func (o *CreateTopic) SetCleanupPolicies(v []string)`

SetCleanupPolicies sets CleanupPolicies field to given value.

### HasCleanupPolicies

`func (o *CreateTopic) HasCleanupPolicies() bool`

HasCleanupPolicies returns a boolean if a field has been set.

### GetDescription

`func (o *CreateTopic) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTopic) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTopic) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTopic) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateTopic) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateTopic) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateTopic) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateTopic) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtension

`func (o *CreateTopic) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *CreateTopic) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *CreateTopic) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *CreateTopic) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetMaximumMessageSize

`func (o *CreateTopic) GetMaximumMessageSize() int32`

GetMaximumMessageSize returns the MaximumMessageSize field if non-nil, zero value otherwise.

### GetMaximumMessageSizeOk

`func (o *CreateTopic) GetMaximumMessageSizeOk() (*int32, bool)`

GetMaximumMessageSizeOk returns a tuple with the MaximumMessageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMessageSize

`func (o *CreateTopic) SetMaximumMessageSize(v int32)`

SetMaximumMessageSize sets MaximumMessageSize field to given value.

### HasMaximumMessageSize

`func (o *CreateTopic) HasMaximumMessageSize() bool`

HasMaximumMessageSize returns a boolean if a field has been set.

### GetMinimumInSyncReplicas

`func (o *CreateTopic) GetMinimumInSyncReplicas() int32`

GetMinimumInSyncReplicas returns the MinimumInSyncReplicas field if non-nil, zero value otherwise.

### GetMinimumInSyncReplicasOk

`func (o *CreateTopic) GetMinimumInSyncReplicasOk() (*int32, bool)`

GetMinimumInSyncReplicasOk returns a tuple with the MinimumInSyncReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumInSyncReplicas

`func (o *CreateTopic) SetMinimumInSyncReplicas(v int32)`

SetMinimumInSyncReplicas sets MinimumInSyncReplicas field to given value.

### HasMinimumInSyncReplicas

`func (o *CreateTopic) HasMinimumInSyncReplicas() bool`

HasMinimumInSyncReplicas returns a boolean if a field has been set.

### GetName

`func (o *CreateTopic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTopic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTopic) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *CreateTopic) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateTopic) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateTopic) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateTopic) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPartitions

`func (o *CreateTopic) GetPartitions() int32`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *CreateTopic) GetPartitionsOk() (*int32, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *CreateTopic) SetPartitions(v int32)`

SetPartitions sets Partitions field to given value.


### GetReplicationFactor

`func (o *CreateTopic) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *CreateTopic) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *CreateTopic) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *CreateTopic) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetRetentionSize

`func (o *CreateTopic) GetRetentionSize() float64`

GetRetentionSize returns the RetentionSize field if non-nil, zero value otherwise.

### GetRetentionSizeOk

`func (o *CreateTopic) GetRetentionSizeOk() (*float64, bool)`

GetRetentionSizeOk returns a tuple with the RetentionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionSize

`func (o *CreateTopic) SetRetentionSize(v float64)`

SetRetentionSize sets RetentionSize field to given value.

### HasRetentionSize

`func (o *CreateTopic) HasRetentionSize() bool`

HasRetentionSize returns a boolean if a field has been set.

### GetRetentionTime

`func (o *CreateTopic) GetRetentionTime() float64`

GetRetentionTime returns the RetentionTime field if non-nil, zero value otherwise.

### GetRetentionTimeOk

`func (o *CreateTopic) GetRetentionTimeOk() (*float64, bool)`

GetRetentionTimeOk returns a tuple with the RetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionTime

`func (o *CreateTopic) SetRetentionTime(v float64)`

SetRetentionTime sets RetentionTime field to given value.

### HasRetentionTime

`func (o *CreateTopic) HasRetentionTime() bool`

HasRetentionTime returns a boolean if a field has been set.

### GetSchemaText

`func (o *CreateTopic) GetSchemaText() string`

GetSchemaText returns the SchemaText field if non-nil, zero value otherwise.

### GetSchemaTextOk

`func (o *CreateTopic) GetSchemaTextOk() (*string, bool)`

GetSchemaTextOk returns a tuple with the SchemaText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaText

`func (o *CreateTopic) SetSchemaText(v string)`

SetSchemaText sets SchemaText field to given value.

### HasSchemaText

`func (o *CreateTopic) HasSchemaText() bool`

HasSchemaText returns a boolean if a field has been set.

### GetSchemaType

`func (o *CreateTopic) GetSchemaType() string`

GetSchemaType returns the SchemaType field if non-nil, zero value otherwise.

### GetSchemaTypeOk

`func (o *CreateTopic) GetSchemaTypeOk() (*string, bool)`

GetSchemaTypeOk returns a tuple with the SchemaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaType

`func (o *CreateTopic) SetSchemaType(v string)`

SetSchemaType sets SchemaType field to given value.

### HasSchemaType

`func (o *CreateTopic) HasSchemaType() bool`

HasSchemaType returns a boolean if a field has been set.

### GetService

`func (o *CreateTopic) GetService() EntityReference`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CreateTopic) GetServiceOk() (*EntityReference, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CreateTopic) SetService(v EntityReference)`

SetService sets Service field to given value.


### GetTags

`func (o *CreateTopic) GetTags() []TagLabel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateTopic) GetTagsOk() (*[]TagLabel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateTopic) SetTags(v []TagLabel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateTopic) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTopicConfig

`func (o *CreateTopic) GetTopicConfig() map[string]interface{}`

GetTopicConfig returns the TopicConfig field if non-nil, zero value otherwise.

### GetTopicConfigOk

`func (o *CreateTopic) GetTopicConfigOk() (*map[string]interface{}, bool)`

GetTopicConfigOk returns a tuple with the TopicConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicConfig

`func (o *CreateTopic) SetTopicConfig(v map[string]interface{})`

SetTopicConfig sets TopicConfig field to given value.

### HasTopicConfig

`func (o *CreateTopic) HasTopicConfig() bool`

HasTopicConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


