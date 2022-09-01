# ShardSearchFailureCause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalizedMessage** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**StackTrace** | Pointer to [**[]ShardSearchFailureCauseStackTraceInner**](ShardSearchFailureCauseStackTraceInner.md) |  | [optional] 
**Suppressed** | Pointer to [**[]ShardSearchFailureCauseSuppressedInner**](ShardSearchFailureCauseSuppressedInner.md) |  | [optional] 

## Methods

### NewShardSearchFailureCause

`func NewShardSearchFailureCause() *ShardSearchFailureCause`

NewShardSearchFailureCause instantiates a new ShardSearchFailureCause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardSearchFailureCauseWithDefaults

`func NewShardSearchFailureCauseWithDefaults() *ShardSearchFailureCause`

NewShardSearchFailureCauseWithDefaults instantiates a new ShardSearchFailureCause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalizedMessage

`func (o *ShardSearchFailureCause) GetLocalizedMessage() string`

GetLocalizedMessage returns the LocalizedMessage field if non-nil, zero value otherwise.

### GetLocalizedMessageOk

`func (o *ShardSearchFailureCause) GetLocalizedMessageOk() (*string, bool)`

GetLocalizedMessageOk returns a tuple with the LocalizedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedMessage

`func (o *ShardSearchFailureCause) SetLocalizedMessage(v string)`

SetLocalizedMessage sets LocalizedMessage field to given value.

### HasLocalizedMessage

`func (o *ShardSearchFailureCause) HasLocalizedMessage() bool`

HasLocalizedMessage returns a boolean if a field has been set.

### GetMessage

`func (o *ShardSearchFailureCause) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ShardSearchFailureCause) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ShardSearchFailureCause) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ShardSearchFailureCause) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStackTrace

`func (o *ShardSearchFailureCause) GetStackTrace() []ShardSearchFailureCauseStackTraceInner`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *ShardSearchFailureCause) GetStackTraceOk() (*[]ShardSearchFailureCauseStackTraceInner, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *ShardSearchFailureCause) SetStackTrace(v []ShardSearchFailureCauseStackTraceInner)`

SetStackTrace sets StackTrace field to given value.

### HasStackTrace

`func (o *ShardSearchFailureCause) HasStackTrace() bool`

HasStackTrace returns a boolean if a field has been set.

### GetSuppressed

`func (o *ShardSearchFailureCause) GetSuppressed() []ShardSearchFailureCauseSuppressedInner`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *ShardSearchFailureCause) GetSuppressedOk() (*[]ShardSearchFailureCauseSuppressedInner, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *ShardSearchFailureCause) SetSuppressed(v []ShardSearchFailureCauseSuppressedInner)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *ShardSearchFailureCause) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


