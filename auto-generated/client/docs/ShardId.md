# ShardId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fragment** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Index** | Pointer to [**Index**](Index.md) |  | [optional] 
**IndexName** | Pointer to **string** |  | [optional] 

## Methods

### NewShardId

`func NewShardId() *ShardId`

NewShardId instantiates a new ShardId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardIdWithDefaults

`func NewShardIdWithDefaults() *ShardId`

NewShardIdWithDefaults instantiates a new ShardId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFragment

`func (o *ShardId) GetFragment() bool`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *ShardId) GetFragmentOk() (*bool, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *ShardId) SetFragment(v bool)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *ShardId) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetId

`func (o *ShardId) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShardId) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShardId) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShardId) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndex

`func (o *ShardId) GetIndex() Index`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ShardId) GetIndexOk() (*Index, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ShardId) SetIndex(v Index)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ShardId) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetIndexName

`func (o *ShardId) GetIndexName() string`

GetIndexName returns the IndexName field if non-nil, zero value otherwise.

### GetIndexNameOk

`func (o *ShardId) GetIndexNameOk() (*string, bool)`

GetIndexNameOk returns a tuple with the IndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexName

`func (o *ShardId) SetIndexName(v string)`

SetIndexName sets IndexName field to given value.

### HasIndexName

`func (o *ShardId) HasIndexName() bool`

HasIndexName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


