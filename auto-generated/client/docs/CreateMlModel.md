# CreateMlModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** |  | 
**Dashboard** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **map[string]interface{}** |  | [optional] 
**MlFeatures** | Pointer to [**[]MlFeature**](MlFeature.md) |  | [optional] 
**MlHyperParameters** | Pointer to [**[]MlHyperParameter**](MlHyperParameter.md) |  | [optional] 
**MlStore** | Pointer to [**MlStore**](MlStore.md) |  | [optional] 
**Name** | **string** |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**Service** | [**EntityReference**](EntityReference.md) |  | 
**Tags** | Pointer to [**[]TagLabel**](TagLabel.md) |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateMlModel

`func NewCreateMlModel(algorithm string, name string, service EntityReference, ) *CreateMlModel`

NewCreateMlModel instantiates a new CreateMlModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMlModelWithDefaults

`func NewCreateMlModelWithDefaults() *CreateMlModel`

NewCreateMlModelWithDefaults instantiates a new CreateMlModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *CreateMlModel) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *CreateMlModel) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *CreateMlModel) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetDashboard

`func (o *CreateMlModel) GetDashboard() EntityReference`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *CreateMlModel) GetDashboardOk() (*EntityReference, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *CreateMlModel) SetDashboard(v EntityReference)`

SetDashboard sets Dashboard field to given value.

### HasDashboard

`func (o *CreateMlModel) HasDashboard() bool`

HasDashboard returns a boolean if a field has been set.

### GetDescription

`func (o *CreateMlModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMlModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMlModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateMlModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateMlModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateMlModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateMlModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateMlModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtension

`func (o *CreateMlModel) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *CreateMlModel) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *CreateMlModel) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *CreateMlModel) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetMlFeatures

`func (o *CreateMlModel) GetMlFeatures() []MlFeature`

GetMlFeatures returns the MlFeatures field if non-nil, zero value otherwise.

### GetMlFeaturesOk

`func (o *CreateMlModel) GetMlFeaturesOk() (*[]MlFeature, bool)`

GetMlFeaturesOk returns a tuple with the MlFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlFeatures

`func (o *CreateMlModel) SetMlFeatures(v []MlFeature)`

SetMlFeatures sets MlFeatures field to given value.

### HasMlFeatures

`func (o *CreateMlModel) HasMlFeatures() bool`

HasMlFeatures returns a boolean if a field has been set.

### GetMlHyperParameters

`func (o *CreateMlModel) GetMlHyperParameters() []MlHyperParameter`

GetMlHyperParameters returns the MlHyperParameters field if non-nil, zero value otherwise.

### GetMlHyperParametersOk

`func (o *CreateMlModel) GetMlHyperParametersOk() (*[]MlHyperParameter, bool)`

GetMlHyperParametersOk returns a tuple with the MlHyperParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlHyperParameters

`func (o *CreateMlModel) SetMlHyperParameters(v []MlHyperParameter)`

SetMlHyperParameters sets MlHyperParameters field to given value.

### HasMlHyperParameters

`func (o *CreateMlModel) HasMlHyperParameters() bool`

HasMlHyperParameters returns a boolean if a field has been set.

### GetMlStore

`func (o *CreateMlModel) GetMlStore() MlStore`

GetMlStore returns the MlStore field if non-nil, zero value otherwise.

### GetMlStoreOk

`func (o *CreateMlModel) GetMlStoreOk() (*MlStore, bool)`

GetMlStoreOk returns a tuple with the MlStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlStore

`func (o *CreateMlModel) SetMlStore(v MlStore)`

SetMlStore sets MlStore field to given value.

### HasMlStore

`func (o *CreateMlModel) HasMlStore() bool`

HasMlStore returns a boolean if a field has been set.

### GetName

`func (o *CreateMlModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMlModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMlModel) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *CreateMlModel) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateMlModel) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateMlModel) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateMlModel) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetServer

`func (o *CreateMlModel) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *CreateMlModel) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *CreateMlModel) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *CreateMlModel) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetService

`func (o *CreateMlModel) GetService() EntityReference`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CreateMlModel) GetServiceOk() (*EntityReference, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CreateMlModel) SetService(v EntityReference)`

SetService sets Service field to given value.


### GetTags

`func (o *CreateMlModel) GetTags() []TagLabel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateMlModel) GetTagsOk() (*[]TagLabel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateMlModel) SetTags(v []TagLabel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateMlModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTarget

`func (o *CreateMlModel) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CreateMlModel) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CreateMlModel) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CreateMlModel) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


