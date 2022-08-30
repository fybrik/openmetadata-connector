# TaskDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignees** | [**[]EntityReference**](EntityReference.md) |  | 
**ClosedAt** | Pointer to **int64** |  | [optional] 
**ClosedBy** | Pointer to **string** |  | [optional] 
**Id** | **int32** |  | 
**NewValue** | Pointer to **string** |  | [optional] 
**OldValue** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Suggestion** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewTaskDetails

`func NewTaskDetails(assignees []EntityReference, id int32, type_ string, ) *TaskDetails`

NewTaskDetails instantiates a new TaskDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskDetailsWithDefaults

`func NewTaskDetailsWithDefaults() *TaskDetails`

NewTaskDetailsWithDefaults instantiates a new TaskDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignees

`func (o *TaskDetails) GetAssignees() []EntityReference`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *TaskDetails) GetAssigneesOk() (*[]EntityReference, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *TaskDetails) SetAssignees(v []EntityReference)`

SetAssignees sets Assignees field to given value.


### GetClosedAt

`func (o *TaskDetails) GetClosedAt() int64`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *TaskDetails) GetClosedAtOk() (*int64, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *TaskDetails) SetClosedAt(v int64)`

SetClosedAt sets ClosedAt field to given value.

### HasClosedAt

`func (o *TaskDetails) HasClosedAt() bool`

HasClosedAt returns a boolean if a field has been set.

### GetClosedBy

`func (o *TaskDetails) GetClosedBy() string`

GetClosedBy returns the ClosedBy field if non-nil, zero value otherwise.

### GetClosedByOk

`func (o *TaskDetails) GetClosedByOk() (*string, bool)`

GetClosedByOk returns a tuple with the ClosedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedBy

`func (o *TaskDetails) SetClosedBy(v string)`

SetClosedBy sets ClosedBy field to given value.

### HasClosedBy

`func (o *TaskDetails) HasClosedBy() bool`

HasClosedBy returns a boolean if a field has been set.

### GetId

`func (o *TaskDetails) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskDetails) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskDetails) SetId(v int32)`

SetId sets Id field to given value.


### GetNewValue

`func (o *TaskDetails) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *TaskDetails) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *TaskDetails) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *TaskDetails) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### GetOldValue

`func (o *TaskDetails) GetOldValue() string`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *TaskDetails) GetOldValueOk() (*string, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *TaskDetails) SetOldValue(v string)`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *TaskDetails) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### GetStatus

`func (o *TaskDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuggestion

`func (o *TaskDetails) GetSuggestion() string`

GetSuggestion returns the Suggestion field if non-nil, zero value otherwise.

### GetSuggestionOk

`func (o *TaskDetails) GetSuggestionOk() (*string, bool)`

GetSuggestionOk returns a tuple with the Suggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestion

`func (o *TaskDetails) SetSuggestion(v string)`

SetSuggestion sets Suggestion field to given value.

### HasSuggestion

`func (o *TaskDetails) HasSuggestion() bool`

HasSuggestion returns a boolean if a field has been set.

### GetType

`func (o *TaskDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskDetails) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


