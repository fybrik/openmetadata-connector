# MlModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** |  | 
**ChangeDescription** | Pointer to [**ChangeDescription**](ChangeDescription.md) |  | [optional] 
**Dashboard** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **map[string]interface{}** |  | [optional] 
**Followers** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**FullyQualifiedName** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**MlFeatures** | Pointer to [**[]MlFeature**](MlFeature.md) |  | [optional] 
**MlHyperParameters** | Pointer to [**[]MlHyperParameter**](MlHyperParameter.md) |  | [optional] 
**MlStore** | Pointer to [**MlStore**](MlStore.md) |  | [optional] 
**Name** | **string** |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**Service** | [**EntityReference**](EntityReference.md) |  | 
**ServiceType** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagLabel**](TagLabel.md) |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**UsageSummary** | Pointer to [**UsageDetails**](UsageDetails.md) |  | [optional] 
**Version** | Pointer to **float64** |  | [optional] 

## Methods

### NewMlModel

`func NewMlModel(algorithm string, id string, name string, service EntityReference, ) *MlModel`

NewMlModel instantiates a new MlModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlModelWithDefaults

`func NewMlModelWithDefaults() *MlModel`

NewMlModelWithDefaults instantiates a new MlModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *MlModel) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *MlModel) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *MlModel) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetChangeDescription

`func (o *MlModel) GetChangeDescription() ChangeDescription`

GetChangeDescription returns the ChangeDescription field if non-nil, zero value otherwise.

### GetChangeDescriptionOk

`func (o *MlModel) GetChangeDescriptionOk() (*ChangeDescription, bool)`

GetChangeDescriptionOk returns a tuple with the ChangeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDescription

`func (o *MlModel) SetChangeDescription(v ChangeDescription)`

SetChangeDescription sets ChangeDescription field to given value.

### HasChangeDescription

`func (o *MlModel) HasChangeDescription() bool`

HasChangeDescription returns a boolean if a field has been set.

### GetDashboard

`func (o *MlModel) GetDashboard() EntityReference`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *MlModel) GetDashboardOk() (*EntityReference, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *MlModel) SetDashboard(v EntityReference)`

SetDashboard sets Dashboard field to given value.

### HasDashboard

`func (o *MlModel) HasDashboard() bool`

HasDashboard returns a boolean if a field has been set.

### GetDeleted

`func (o *MlModel) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *MlModel) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *MlModel) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *MlModel) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDescription

`func (o *MlModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MlModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MlModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MlModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *MlModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MlModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MlModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MlModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtension

`func (o *MlModel) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *MlModel) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *MlModel) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *MlModel) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetFollowers

`func (o *MlModel) GetFollowers() []EntityReference`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *MlModel) GetFollowersOk() (*[]EntityReference, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *MlModel) SetFollowers(v []EntityReference)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *MlModel) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetFullyQualifiedName

`func (o *MlModel) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *MlModel) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *MlModel) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.

### HasFullyQualifiedName

`func (o *MlModel) HasFullyQualifiedName() bool`

HasFullyQualifiedName returns a boolean if a field has been set.

### GetHref

`func (o *MlModel) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *MlModel) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *MlModel) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *MlModel) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *MlModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MlModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MlModel) SetId(v string)`

SetId sets Id field to given value.


### GetMlFeatures

`func (o *MlModel) GetMlFeatures() []MlFeature`

GetMlFeatures returns the MlFeatures field if non-nil, zero value otherwise.

### GetMlFeaturesOk

`func (o *MlModel) GetMlFeaturesOk() (*[]MlFeature, bool)`

GetMlFeaturesOk returns a tuple with the MlFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlFeatures

`func (o *MlModel) SetMlFeatures(v []MlFeature)`

SetMlFeatures sets MlFeatures field to given value.

### HasMlFeatures

`func (o *MlModel) HasMlFeatures() bool`

HasMlFeatures returns a boolean if a field has been set.

### GetMlHyperParameters

`func (o *MlModel) GetMlHyperParameters() []MlHyperParameter`

GetMlHyperParameters returns the MlHyperParameters field if non-nil, zero value otherwise.

### GetMlHyperParametersOk

`func (o *MlModel) GetMlHyperParametersOk() (*[]MlHyperParameter, bool)`

GetMlHyperParametersOk returns a tuple with the MlHyperParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlHyperParameters

`func (o *MlModel) SetMlHyperParameters(v []MlHyperParameter)`

SetMlHyperParameters sets MlHyperParameters field to given value.

### HasMlHyperParameters

`func (o *MlModel) HasMlHyperParameters() bool`

HasMlHyperParameters returns a boolean if a field has been set.

### GetMlStore

`func (o *MlModel) GetMlStore() MlStore`

GetMlStore returns the MlStore field if non-nil, zero value otherwise.

### GetMlStoreOk

`func (o *MlModel) GetMlStoreOk() (*MlStore, bool)`

GetMlStoreOk returns a tuple with the MlStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlStore

`func (o *MlModel) SetMlStore(v MlStore)`

SetMlStore sets MlStore field to given value.

### HasMlStore

`func (o *MlModel) HasMlStore() bool`

HasMlStore returns a boolean if a field has been set.

### GetName

`func (o *MlModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MlModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MlModel) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *MlModel) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MlModel) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MlModel) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *MlModel) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetServer

`func (o *MlModel) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *MlModel) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *MlModel) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *MlModel) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetService

`func (o *MlModel) GetService() EntityReference`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *MlModel) GetServiceOk() (*EntityReference, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *MlModel) SetService(v EntityReference)`

SetService sets Service field to given value.


### GetServiceType

`func (o *MlModel) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *MlModel) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *MlModel) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *MlModel) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetTags

`func (o *MlModel) GetTags() []TagLabel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MlModel) GetTagsOk() (*[]TagLabel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MlModel) SetTags(v []TagLabel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MlModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTarget

`func (o *MlModel) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MlModel) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *MlModel) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *MlModel) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MlModel) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MlModel) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MlModel) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MlModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *MlModel) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *MlModel) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *MlModel) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *MlModel) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUsageSummary

`func (o *MlModel) GetUsageSummary() UsageDetails`

GetUsageSummary returns the UsageSummary field if non-nil, zero value otherwise.

### GetUsageSummaryOk

`func (o *MlModel) GetUsageSummaryOk() (*UsageDetails, bool)`

GetUsageSummaryOk returns a tuple with the UsageSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageSummary

`func (o *MlModel) SetUsageSummary(v UsageDetails)`

SetUsageSummary sets UsageSummary field to given value.

### HasUsageSummary

`func (o *MlModel) HasUsageSummary() bool`

HasUsageSummary returns a boolean if a field has been set.

### GetVersion

`func (o *MlModel) GetVersion() float64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MlModel) GetVersionOk() (*float64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MlModel) SetVersion(v float64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MlModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


