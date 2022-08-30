# CreateIngestionPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AirflowConfig** | [**AirflowConfig**](AirflowConfig.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **map[string]interface{}** |  | [optional] 
**LoggerLevel** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**PipelineType** | **string** |  | 
**Service** | [**EntityReference**](EntityReference.md) |  | 
**SourceConfig** | [**SourceConfig**](SourceConfig.md) |  | 

## Methods

### NewCreateIngestionPipeline

`func NewCreateIngestionPipeline(airflowConfig AirflowConfig, name string, pipelineType string, service EntityReference, sourceConfig SourceConfig, ) *CreateIngestionPipeline`

NewCreateIngestionPipeline instantiates a new CreateIngestionPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIngestionPipelineWithDefaults

`func NewCreateIngestionPipelineWithDefaults() *CreateIngestionPipeline`

NewCreateIngestionPipelineWithDefaults instantiates a new CreateIngestionPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAirflowConfig

`func (o *CreateIngestionPipeline) GetAirflowConfig() AirflowConfig`

GetAirflowConfig returns the AirflowConfig field if non-nil, zero value otherwise.

### GetAirflowConfigOk

`func (o *CreateIngestionPipeline) GetAirflowConfigOk() (*AirflowConfig, bool)`

GetAirflowConfigOk returns a tuple with the AirflowConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflowConfig

`func (o *CreateIngestionPipeline) SetAirflowConfig(v AirflowConfig)`

SetAirflowConfig sets AirflowConfig field to given value.


### GetDescription

`func (o *CreateIngestionPipeline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateIngestionPipeline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateIngestionPipeline) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateIngestionPipeline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateIngestionPipeline) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateIngestionPipeline) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateIngestionPipeline) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateIngestionPipeline) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtension

`func (o *CreateIngestionPipeline) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *CreateIngestionPipeline) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *CreateIngestionPipeline) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *CreateIngestionPipeline) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetLoggerLevel

`func (o *CreateIngestionPipeline) GetLoggerLevel() string`

GetLoggerLevel returns the LoggerLevel field if non-nil, zero value otherwise.

### GetLoggerLevelOk

`func (o *CreateIngestionPipeline) GetLoggerLevelOk() (*string, bool)`

GetLoggerLevelOk returns a tuple with the LoggerLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggerLevel

`func (o *CreateIngestionPipeline) SetLoggerLevel(v string)`

SetLoggerLevel sets LoggerLevel field to given value.

### HasLoggerLevel

`func (o *CreateIngestionPipeline) HasLoggerLevel() bool`

HasLoggerLevel returns a boolean if a field has been set.

### GetName

`func (o *CreateIngestionPipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIngestionPipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIngestionPipeline) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *CreateIngestionPipeline) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateIngestionPipeline) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateIngestionPipeline) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateIngestionPipeline) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPipelineType

`func (o *CreateIngestionPipeline) GetPipelineType() string`

GetPipelineType returns the PipelineType field if non-nil, zero value otherwise.

### GetPipelineTypeOk

`func (o *CreateIngestionPipeline) GetPipelineTypeOk() (*string, bool)`

GetPipelineTypeOk returns a tuple with the PipelineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineType

`func (o *CreateIngestionPipeline) SetPipelineType(v string)`

SetPipelineType sets PipelineType field to given value.


### GetService

`func (o *CreateIngestionPipeline) GetService() EntityReference`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CreateIngestionPipeline) GetServiceOk() (*EntityReference, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CreateIngestionPipeline) SetService(v EntityReference)`

SetService sets Service field to given value.


### GetSourceConfig

`func (o *CreateIngestionPipeline) GetSourceConfig() SourceConfig`

GetSourceConfig returns the SourceConfig field if non-nil, zero value otherwise.

### GetSourceConfigOk

`func (o *CreateIngestionPipeline) GetSourceConfigOk() (*SourceConfig, bool)`

GetSourceConfigOk returns a tuple with the SourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConfig

`func (o *CreateIngestionPipeline) SetSourceConfig(v SourceConfig)`

SetSourceConfig sets SourceConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


