# EntitiesEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**FromEntity** | [**EntityReference**](EntityReference.md) |  | 
**LineageDetails** | Pointer to [**LineageDetails**](LineageDetails.md) |  | [optional] 
**ToEntity** | [**EntityReference**](EntityReference.md) |  | 

## Methods

### NewEntitiesEdge

`func NewEntitiesEdge(fromEntity EntityReference, toEntity EntityReference, ) *EntitiesEdge`

NewEntitiesEdge instantiates a new EntitiesEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitiesEdgeWithDefaults

`func NewEntitiesEdgeWithDefaults() *EntitiesEdge`

NewEntitiesEdgeWithDefaults instantiates a new EntitiesEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EntitiesEdge) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitiesEdge) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitiesEdge) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitiesEdge) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFromEntity

`func (o *EntitiesEdge) GetFromEntity() EntityReference`

GetFromEntity returns the FromEntity field if non-nil, zero value otherwise.

### GetFromEntityOk

`func (o *EntitiesEdge) GetFromEntityOk() (*EntityReference, bool)`

GetFromEntityOk returns a tuple with the FromEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEntity

`func (o *EntitiesEdge) SetFromEntity(v EntityReference)`

SetFromEntity sets FromEntity field to given value.


### GetLineageDetails

`func (o *EntitiesEdge) GetLineageDetails() LineageDetails`

GetLineageDetails returns the LineageDetails field if non-nil, zero value otherwise.

### GetLineageDetailsOk

`func (o *EntitiesEdge) GetLineageDetailsOk() (*LineageDetails, bool)`

GetLineageDetailsOk returns a tuple with the LineageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineageDetails

`func (o *EntitiesEdge) SetLineageDetails(v LineageDetails)`

SetLineageDetails sets LineageDetails field to given value.

### HasLineageDetails

`func (o *EntitiesEdge) HasLineageDetails() bool`

HasLineageDetails returns a boolean if a field has been set.

### GetToEntity

`func (o *EntitiesEdge) GetToEntity() EntityReference`

GetToEntity returns the ToEntity field if non-nil, zero value otherwise.

### GetToEntityOk

`func (o *EntitiesEdge) GetToEntityOk() (*EntityReference, bool)`

GetToEntityOk returns a tuple with the ToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToEntity

`func (o *EntitiesEdge) SetToEntity(v EntityReference)`

SetToEntity sets ToEntity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


