# CreateTaskDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignees** | [**[]EntityReference**](EntityReference.md) |  | 
**OldValue** | Pointer to **string** |  | [optional] 
**Suggestion** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewCreateTaskDetails

`func NewCreateTaskDetails(assignees []EntityReference, type_ string, ) *CreateTaskDetails`

NewCreateTaskDetails instantiates a new CreateTaskDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTaskDetailsWithDefaults

`func NewCreateTaskDetailsWithDefaults() *CreateTaskDetails`

NewCreateTaskDetailsWithDefaults instantiates a new CreateTaskDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignees

`func (o *CreateTaskDetails) GetAssignees() []EntityReference`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *CreateTaskDetails) GetAssigneesOk() (*[]EntityReference, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *CreateTaskDetails) SetAssignees(v []EntityReference)`

SetAssignees sets Assignees field to given value.


### GetOldValue

`func (o *CreateTaskDetails) GetOldValue() string`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *CreateTaskDetails) GetOldValueOk() (*string, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *CreateTaskDetails) SetOldValue(v string)`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *CreateTaskDetails) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### GetSuggestion

`func (o *CreateTaskDetails) GetSuggestion() string`

GetSuggestion returns the Suggestion field if non-nil, zero value otherwise.

### GetSuggestionOk

`func (o *CreateTaskDetails) GetSuggestionOk() (*string, bool)`

GetSuggestionOk returns a tuple with the Suggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestion

`func (o *CreateTaskDetails) SetSuggestion(v string)`

SetSuggestion sets Suggestion field to given value.

### HasSuggestion

`func (o *CreateTaskDetails) HasSuggestion() bool`

HasSuggestion returns a boolean if a field has been set.

### GetType

`func (o *CreateTaskDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTaskDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTaskDetails) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


