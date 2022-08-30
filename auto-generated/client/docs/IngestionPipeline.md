# IngestionPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AirflowConfig** | [**AirflowConfig**](AirflowConfig.md) |  | 
**ChangeDescription** | Pointer to [**ChangeDescription**](ChangeDescription.md) |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Deployed** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Extension** | Pointer to **map[string]interface{}** |  | [optional] 
**Followers** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**FullyQualifiedName** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LoggerLevel** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**OpenMetadataServerConnection** | [**OpenMetadataServerConnection**](OpenMetadataServerConnection.md) |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**PipelineStatuses** | Pointer to [**[]PipelineStatus**](PipelineStatus.md) |  | [optional] 
**PipelineType** | **string** |  | 
**Service** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**SourceConfig** | [**SourceConfig**](SourceConfig.md) |  | 
**Tags** | Pointer to [**[]TagLabel**](TagLabel.md) |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **float64** |  | [optional] 

## Methods

### NewIngestionPipeline

`func NewIngestionPipeline(airflowConfig AirflowConfig, name string, openMetadataServerConnection OpenMetadataServerConnection, pipelineType string, sourceConfig SourceConfig, ) *IngestionPipeline`

NewIngestionPipeline instantiates a new IngestionPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionPipelineWithDefaults

`func NewIngestionPipelineWithDefaults() *IngestionPipeline`

NewIngestionPipelineWithDefaults instantiates a new IngestionPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAirflowConfig

`func (o *IngestionPipeline) GetAirflowConfig() AirflowConfig`

GetAirflowConfig returns the AirflowConfig field if non-nil, zero value otherwise.

### GetAirflowConfigOk

`func (o *IngestionPipeline) GetAirflowConfigOk() (*AirflowConfig, bool)`

GetAirflowConfigOk returns a tuple with the AirflowConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflowConfig

`func (o *IngestionPipeline) SetAirflowConfig(v AirflowConfig)`

SetAirflowConfig sets AirflowConfig field to given value.


### GetChangeDescription

`func (o *IngestionPipeline) GetChangeDescription() ChangeDescription`

GetChangeDescription returns the ChangeDescription field if non-nil, zero value otherwise.

### GetChangeDescriptionOk

`func (o *IngestionPipeline) GetChangeDescriptionOk() (*ChangeDescription, bool)`

GetChangeDescriptionOk returns a tuple with the ChangeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDescription

`func (o *IngestionPipeline) SetChangeDescription(v ChangeDescription)`

SetChangeDescription sets ChangeDescription field to given value.

### HasChangeDescription

`func (o *IngestionPipeline) HasChangeDescription() bool`

HasChangeDescription returns a boolean if a field has been set.

### GetDeleted

`func (o *IngestionPipeline) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IngestionPipeline) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IngestionPipeline) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IngestionPipeline) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeployed

`func (o *IngestionPipeline) GetDeployed() bool`

GetDeployed returns the Deployed field if non-nil, zero value otherwise.

### GetDeployedOk

`func (o *IngestionPipeline) GetDeployedOk() (*bool, bool)`

GetDeployedOk returns a tuple with the Deployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployed

`func (o *IngestionPipeline) SetDeployed(v bool)`

SetDeployed sets Deployed field to given value.

### HasDeployed

`func (o *IngestionPipeline) HasDeployed() bool`

HasDeployed returns a boolean if a field has been set.

### GetDescription

`func (o *IngestionPipeline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IngestionPipeline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IngestionPipeline) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IngestionPipeline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *IngestionPipeline) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IngestionPipeline) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IngestionPipeline) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IngestionPipeline) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnabled

`func (o *IngestionPipeline) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IngestionPipeline) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IngestionPipeline) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IngestionPipeline) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExtension

`func (o *IngestionPipeline) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *IngestionPipeline) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *IngestionPipeline) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *IngestionPipeline) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetFollowers

`func (o *IngestionPipeline) GetFollowers() []EntityReference`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *IngestionPipeline) GetFollowersOk() (*[]EntityReference, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *IngestionPipeline) SetFollowers(v []EntityReference)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *IngestionPipeline) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetFullyQualifiedName

`func (o *IngestionPipeline) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *IngestionPipeline) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *IngestionPipeline) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.

### HasFullyQualifiedName

`func (o *IngestionPipeline) HasFullyQualifiedName() bool`

HasFullyQualifiedName returns a boolean if a field has been set.

### GetHref

`func (o *IngestionPipeline) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *IngestionPipeline) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *IngestionPipeline) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *IngestionPipeline) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *IngestionPipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IngestionPipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IngestionPipeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IngestionPipeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoggerLevel

`func (o *IngestionPipeline) GetLoggerLevel() string`

GetLoggerLevel returns the LoggerLevel field if non-nil, zero value otherwise.

### GetLoggerLevelOk

`func (o *IngestionPipeline) GetLoggerLevelOk() (*string, bool)`

GetLoggerLevelOk returns a tuple with the LoggerLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggerLevel

`func (o *IngestionPipeline) SetLoggerLevel(v string)`

SetLoggerLevel sets LoggerLevel field to given value.

### HasLoggerLevel

`func (o *IngestionPipeline) HasLoggerLevel() bool`

HasLoggerLevel returns a boolean if a field has been set.

### GetName

`func (o *IngestionPipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IngestionPipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IngestionPipeline) SetName(v string)`

SetName sets Name field to given value.


### GetOpenMetadataServerConnection

`func (o *IngestionPipeline) GetOpenMetadataServerConnection() OpenMetadataServerConnection`

GetOpenMetadataServerConnection returns the OpenMetadataServerConnection field if non-nil, zero value otherwise.

### GetOpenMetadataServerConnectionOk

`func (o *IngestionPipeline) GetOpenMetadataServerConnectionOk() (*OpenMetadataServerConnection, bool)`

GetOpenMetadataServerConnectionOk returns a tuple with the OpenMetadataServerConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenMetadataServerConnection

`func (o *IngestionPipeline) SetOpenMetadataServerConnection(v OpenMetadataServerConnection)`

SetOpenMetadataServerConnection sets OpenMetadataServerConnection field to given value.


### GetOwner

`func (o *IngestionPipeline) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IngestionPipeline) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IngestionPipeline) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IngestionPipeline) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPipelineStatuses

`func (o *IngestionPipeline) GetPipelineStatuses() []PipelineStatus`

GetPipelineStatuses returns the PipelineStatuses field if non-nil, zero value otherwise.

### GetPipelineStatusesOk

`func (o *IngestionPipeline) GetPipelineStatusesOk() (*[]PipelineStatus, bool)`

GetPipelineStatusesOk returns a tuple with the PipelineStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineStatuses

`func (o *IngestionPipeline) SetPipelineStatuses(v []PipelineStatus)`

SetPipelineStatuses sets PipelineStatuses field to given value.

### HasPipelineStatuses

`func (o *IngestionPipeline) HasPipelineStatuses() bool`

HasPipelineStatuses returns a boolean if a field has been set.

### GetPipelineType

`func (o *IngestionPipeline) GetPipelineType() string`

GetPipelineType returns the PipelineType field if non-nil, zero value otherwise.

### GetPipelineTypeOk

`func (o *IngestionPipeline) GetPipelineTypeOk() (*string, bool)`

GetPipelineTypeOk returns a tuple with the PipelineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineType

`func (o *IngestionPipeline) SetPipelineType(v string)`

SetPipelineType sets PipelineType field to given value.


### GetService

`func (o *IngestionPipeline) GetService() EntityReference`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *IngestionPipeline) GetServiceOk() (*EntityReference, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *IngestionPipeline) SetService(v EntityReference)`

SetService sets Service field to given value.

### HasService

`func (o *IngestionPipeline) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSourceConfig

`func (o *IngestionPipeline) GetSourceConfig() SourceConfig`

GetSourceConfig returns the SourceConfig field if non-nil, zero value otherwise.

### GetSourceConfigOk

`func (o *IngestionPipeline) GetSourceConfigOk() (*SourceConfig, bool)`

GetSourceConfigOk returns a tuple with the SourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConfig

`func (o *IngestionPipeline) SetSourceConfig(v SourceConfig)`

SetSourceConfig sets SourceConfig field to given value.


### GetTags

`func (o *IngestionPipeline) GetTags() []TagLabel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IngestionPipeline) GetTagsOk() (*[]TagLabel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IngestionPipeline) SetTags(v []TagLabel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IngestionPipeline) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IngestionPipeline) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IngestionPipeline) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IngestionPipeline) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IngestionPipeline) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *IngestionPipeline) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *IngestionPipeline) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *IngestionPipeline) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *IngestionPipeline) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetVersion

`func (o *IngestionPipeline) GetVersion() float64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IngestionPipeline) GetVersionOk() (*float64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IngestionPipeline) SetVersion(v float64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IngestionPipeline) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


