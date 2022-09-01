# EntityLineage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownstreamEdges** | Pointer to [**[]Edge**](Edge.md) |  | [optional] 
**Entity** | [**EntityReference**](EntityReference.md) |  | 
**Nodes** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**UpstreamEdges** | Pointer to [**[]Edge**](Edge.md) |  | [optional] 

## Methods

### NewEntityLineage

`func NewEntityLineage(entity EntityReference, ) *EntityLineage`

NewEntityLineage instantiates a new EntityLineage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityLineageWithDefaults

`func NewEntityLineageWithDefaults() *EntityLineage`

NewEntityLineageWithDefaults instantiates a new EntityLineage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownstreamEdges

`func (o *EntityLineage) GetDownstreamEdges() []Edge`

GetDownstreamEdges returns the DownstreamEdges field if non-nil, zero value otherwise.

### GetDownstreamEdgesOk

`func (o *EntityLineage) GetDownstreamEdgesOk() (*[]Edge, bool)`

GetDownstreamEdgesOk returns a tuple with the DownstreamEdges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreamEdges

`func (o *EntityLineage) SetDownstreamEdges(v []Edge)`

SetDownstreamEdges sets DownstreamEdges field to given value.

### HasDownstreamEdges

`func (o *EntityLineage) HasDownstreamEdges() bool`

HasDownstreamEdges returns a boolean if a field has been set.

### GetEntity

`func (o *EntityLineage) GetEntity() EntityReference`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *EntityLineage) GetEntityOk() (*EntityReference, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *EntityLineage) SetEntity(v EntityReference)`

SetEntity sets Entity field to given value.


### GetNodes

`func (o *EntityLineage) GetNodes() []EntityReference`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *EntityLineage) GetNodesOk() (*[]EntityReference, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *EntityLineage) SetNodes(v []EntityReference)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *EntityLineage) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetUpstreamEdges

`func (o *EntityLineage) GetUpstreamEdges() []Edge`

GetUpstreamEdges returns the UpstreamEdges field if non-nil, zero value otherwise.

### GetUpstreamEdgesOk

`func (o *EntityLineage) GetUpstreamEdgesOk() (*[]Edge, bool)`

GetUpstreamEdgesOk returns a tuple with the UpstreamEdges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamEdges

`func (o *EntityLineage) SetUpstreamEdges(v []Edge)`

SetUpstreamEdges sets UpstreamEdges field to given value.

### HasUpstreamEdges

`func (o *EntityLineage) HasUpstreamEdges() bool`

HasUpstreamEdges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


