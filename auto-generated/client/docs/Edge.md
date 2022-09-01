# Edge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**FromEntity** | **string** |  | 
**LineageDetails** | Pointer to [**LineageDetails**](LineageDetails.md) |  | [optional] 
**ToEntity** | **string** |  | 

## Methods

### NewEdge

`func NewEdge(fromEntity string, toEntity string, ) *Edge`

NewEdge instantiates a new Edge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeWithDefaults

`func NewEdgeWithDefaults() *Edge`

NewEdgeWithDefaults instantiates a new Edge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Edge) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Edge) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Edge) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Edge) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFromEntity

`func (o *Edge) GetFromEntity() string`

GetFromEntity returns the FromEntity field if non-nil, zero value otherwise.

### GetFromEntityOk

`func (o *Edge) GetFromEntityOk() (*string, bool)`

GetFromEntityOk returns a tuple with the FromEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEntity

`func (o *Edge) SetFromEntity(v string)`

SetFromEntity sets FromEntity field to given value.


### GetLineageDetails

`func (o *Edge) GetLineageDetails() LineageDetails`

GetLineageDetails returns the LineageDetails field if non-nil, zero value otherwise.

### GetLineageDetailsOk

`func (o *Edge) GetLineageDetailsOk() (*LineageDetails, bool)`

GetLineageDetailsOk returns a tuple with the LineageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineageDetails

`func (o *Edge) SetLineageDetails(v LineageDetails)`

SetLineageDetails sets LineageDetails field to given value.

### HasLineageDetails

`func (o *Edge) HasLineageDetails() bool`

HasLineageDetails returns a boolean if a field has been set.

### GetToEntity

`func (o *Edge) GetToEntity() string`

GetToEntity returns the ToEntity field if non-nil, zero value otherwise.

### GetToEntityOk

`func (o *Edge) GetToEntityOk() (*string, bool)`

GetToEntityOk returns a tuple with the ToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToEntity

`func (o *Edge) SetToEntity(v string)`

SetToEntity sets ToEntity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


