# AddLineage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Edge** | [**EntitiesEdge**](EntitiesEdge.md) |  | 

## Methods

### NewAddLineage

`func NewAddLineage(edge EntitiesEdge, ) *AddLineage`

NewAddLineage instantiates a new AddLineage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLineageWithDefaults

`func NewAddLineageWithDefaults() *AddLineage`

NewAddLineageWithDefaults instantiates a new AddLineage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AddLineage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLineage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLineage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLineage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEdge

`func (o *AddLineage) GetEdge() EntitiesEdge`

GetEdge returns the Edge field if non-nil, zero value otherwise.

### GetEdgeOk

`func (o *AddLineage) GetEdgeOk() (*EntitiesEdge, bool)`

GetEdgeOk returns a tuple with the Edge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdge

`func (o *AddLineage) SetEdge(v EntitiesEdge)`

SetEdge sets Edge field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


