# ShardSearchFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cause** | Pointer to [**ShardSearchFailureCause**](ShardSearchFailureCause.md) |  | [optional] 
**Fragment** | Pointer to **bool** |  | [optional] 

## Methods

### NewShardSearchFailure

`func NewShardSearchFailure() *ShardSearchFailure`

NewShardSearchFailure instantiates a new ShardSearchFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardSearchFailureWithDefaults

`func NewShardSearchFailureWithDefaults() *ShardSearchFailure`

NewShardSearchFailureWithDefaults instantiates a new ShardSearchFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCause

`func (o *ShardSearchFailure) GetCause() ShardSearchFailureCause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *ShardSearchFailure) GetCauseOk() (*ShardSearchFailureCause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *ShardSearchFailure) SetCause(v ShardSearchFailureCause)`

SetCause sets Cause field to given value.

### HasCause

`func (o *ShardSearchFailure) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetFragment

`func (o *ShardSearchFailure) GetFragment() bool`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *ShardSearchFailure) GetFragmentOk() (*bool, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *ShardSearchFailure) SetFragment(v bool)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *ShardSearchFailure) HasFragment() bool`

HasFragment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


