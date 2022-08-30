# SearchShardTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterAlias** | Pointer to **string** |  | [optional] 
**FullyQualifiedIndexName** | Pointer to **string** |  | [optional] 
**Index** | Pointer to **string** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**NodeIdText** | Pointer to [**Text**](Text.md) |  | [optional] 
**OriginalIndices** | Pointer to **map[string]interface{}** |  | [optional] 
**ShardId** | Pointer to [**ShardId**](ShardId.md) |  | [optional] 

## Methods

### NewSearchShardTarget

`func NewSearchShardTarget() *SearchShardTarget`

NewSearchShardTarget instantiates a new SearchShardTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchShardTargetWithDefaults

`func NewSearchShardTargetWithDefaults() *SearchShardTarget`

NewSearchShardTargetWithDefaults instantiates a new SearchShardTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterAlias

`func (o *SearchShardTarget) GetClusterAlias() string`

GetClusterAlias returns the ClusterAlias field if non-nil, zero value otherwise.

### GetClusterAliasOk

`func (o *SearchShardTarget) GetClusterAliasOk() (*string, bool)`

GetClusterAliasOk returns a tuple with the ClusterAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterAlias

`func (o *SearchShardTarget) SetClusterAlias(v string)`

SetClusterAlias sets ClusterAlias field to given value.

### HasClusterAlias

`func (o *SearchShardTarget) HasClusterAlias() bool`

HasClusterAlias returns a boolean if a field has been set.

### GetFullyQualifiedIndexName

`func (o *SearchShardTarget) GetFullyQualifiedIndexName() string`

GetFullyQualifiedIndexName returns the FullyQualifiedIndexName field if non-nil, zero value otherwise.

### GetFullyQualifiedIndexNameOk

`func (o *SearchShardTarget) GetFullyQualifiedIndexNameOk() (*string, bool)`

GetFullyQualifiedIndexNameOk returns a tuple with the FullyQualifiedIndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedIndexName

`func (o *SearchShardTarget) SetFullyQualifiedIndexName(v string)`

SetFullyQualifiedIndexName sets FullyQualifiedIndexName field to given value.

### HasFullyQualifiedIndexName

`func (o *SearchShardTarget) HasFullyQualifiedIndexName() bool`

HasFullyQualifiedIndexName returns a boolean if a field has been set.

### GetIndex

`func (o *SearchShardTarget) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SearchShardTarget) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SearchShardTarget) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *SearchShardTarget) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetNodeId

`func (o *SearchShardTarget) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *SearchShardTarget) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *SearchShardTarget) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *SearchShardTarget) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeIdText

`func (o *SearchShardTarget) GetNodeIdText() Text`

GetNodeIdText returns the NodeIdText field if non-nil, zero value otherwise.

### GetNodeIdTextOk

`func (o *SearchShardTarget) GetNodeIdTextOk() (*Text, bool)`

GetNodeIdTextOk returns a tuple with the NodeIdText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIdText

`func (o *SearchShardTarget) SetNodeIdText(v Text)`

SetNodeIdText sets NodeIdText field to given value.

### HasNodeIdText

`func (o *SearchShardTarget) HasNodeIdText() bool`

HasNodeIdText returns a boolean if a field has been set.

### GetOriginalIndices

`func (o *SearchShardTarget) GetOriginalIndices() map[string]interface{}`

GetOriginalIndices returns the OriginalIndices field if non-nil, zero value otherwise.

### GetOriginalIndicesOk

`func (o *SearchShardTarget) GetOriginalIndicesOk() (*map[string]interface{}, bool)`

GetOriginalIndicesOk returns a tuple with the OriginalIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalIndices

`func (o *SearchShardTarget) SetOriginalIndices(v map[string]interface{})`

SetOriginalIndices sets OriginalIndices field to given value.

### HasOriginalIndices

`func (o *SearchShardTarget) HasOriginalIndices() bool`

HasOriginalIndices returns a boolean if a field has been set.

### GetShardId

`func (o *SearchShardTarget) GetShardId() ShardId`

GetShardId returns the ShardId field if non-nil, zero value otherwise.

### GetShardIdOk

`func (o *SearchShardTarget) GetShardIdOk() (*ShardId, bool)`

GetShardIdOk returns a tuple with the ShardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardId

`func (o *SearchShardTarget) SetShardId(v ShardId)`

SetShardId sets ShardId field to given value.

### HasShardId

`func (o *SearchShardTarget) HasShardId() bool`

HasShardId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


