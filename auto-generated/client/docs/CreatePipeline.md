# CreatePipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Concurrency** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**PipelineLocation** | Pointer to **string** |  | [optional] 
**PipelineUrl** | Pointer to **string** |  | [optional] 
**Service** | [**EntityReference**](EntityReference.md) |  | 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**Tags** | Pointer to [**[]TagLabel**](TagLabel.md) |  | [optional] 
**Tasks** | Pointer to [**[]Task**](Task.md) |  | [optional] 

## Methods

### NewCreatePipeline

`func NewCreatePipeline(name string, service EntityReference, ) *CreatePipeline`

NewCreatePipeline instantiates a new CreatePipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePipelineWithDefaults

`func NewCreatePipelineWithDefaults() *CreatePipeline`

NewCreatePipelineWithDefaults instantiates a new CreatePipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConcurrency

`func (o *CreatePipeline) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *CreatePipeline) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *CreatePipeline) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *CreatePipeline) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### GetDescription

`func (o *CreatePipeline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePipeline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePipeline) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePipeline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreatePipeline) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreatePipeline) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreatePipeline) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreatePipeline) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtension

`func (o *CreatePipeline) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *CreatePipeline) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *CreatePipeline) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *CreatePipeline) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetName

`func (o *CreatePipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePipeline) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *CreatePipeline) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreatePipeline) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreatePipeline) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreatePipeline) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPipelineLocation

`func (o *CreatePipeline) GetPipelineLocation() string`

GetPipelineLocation returns the PipelineLocation field if non-nil, zero value otherwise.

### GetPipelineLocationOk

`func (o *CreatePipeline) GetPipelineLocationOk() (*string, bool)`

GetPipelineLocationOk returns a tuple with the PipelineLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineLocation

`func (o *CreatePipeline) SetPipelineLocation(v string)`

SetPipelineLocation sets PipelineLocation field to given value.

### HasPipelineLocation

`func (o *CreatePipeline) HasPipelineLocation() bool`

HasPipelineLocation returns a boolean if a field has been set.

### GetPipelineUrl

`func (o *CreatePipeline) GetPipelineUrl() string`

GetPipelineUrl returns the PipelineUrl field if non-nil, zero value otherwise.

### GetPipelineUrlOk

`func (o *CreatePipeline) GetPipelineUrlOk() (*string, bool)`

GetPipelineUrlOk returns a tuple with the PipelineUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineUrl

`func (o *CreatePipeline) SetPipelineUrl(v string)`

SetPipelineUrl sets PipelineUrl field to given value.

### HasPipelineUrl

`func (o *CreatePipeline) HasPipelineUrl() bool`

HasPipelineUrl returns a boolean if a field has been set.

### GetService

`func (o *CreatePipeline) GetService() EntityReference`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CreatePipeline) GetServiceOk() (*EntityReference, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CreatePipeline) SetService(v EntityReference)`

SetService sets Service field to given value.


### GetStartDate

`func (o *CreatePipeline) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreatePipeline) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreatePipeline) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreatePipeline) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTags

`func (o *CreatePipeline) GetTags() []TagLabel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreatePipeline) GetTagsOk() (*[]TagLabel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreatePipeline) SetTags(v []TagLabel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreatePipeline) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTasks

`func (o *CreatePipeline) GetTasks() []Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *CreatePipeline) GetTasksOk() (*[]Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *CreatePipeline) SetTasks(v []Task)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *CreatePipeline) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


