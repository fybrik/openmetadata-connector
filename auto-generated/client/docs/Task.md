# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DownstreamTasks** | Pointer to **[]string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**FullyQualifiedName** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**StartDate** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagLabel**](TagLabel.md) |  | [optional] 
**TaskSQL** | Pointer to **string** |  | [optional] 
**TaskType** | Pointer to **string** |  | [optional] 
**TaskUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewTask

`func NewTask(name string, ) *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Task) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Task) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Task) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Task) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *Task) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Task) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Task) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Task) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDownstreamTasks

`func (o *Task) GetDownstreamTasks() []string`

GetDownstreamTasks returns the DownstreamTasks field if non-nil, zero value otherwise.

### GetDownstreamTasksOk

`func (o *Task) GetDownstreamTasksOk() (*[]string, bool)`

GetDownstreamTasksOk returns a tuple with the DownstreamTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreamTasks

`func (o *Task) SetDownstreamTasks(v []string)`

SetDownstreamTasks sets DownstreamTasks field to given value.

### HasDownstreamTasks

`func (o *Task) HasDownstreamTasks() bool`

HasDownstreamTasks returns a boolean if a field has been set.

### GetEndDate

`func (o *Task) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Task) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Task) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Task) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetFullyQualifiedName

`func (o *Task) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *Task) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *Task) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.

### HasFullyQualifiedName

`func (o *Task) HasFullyQualifiedName() bool`

HasFullyQualifiedName returns a boolean if a field has been set.

### GetName

`func (o *Task) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Task) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Task) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *Task) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Task) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Task) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Task) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTags

`func (o *Task) GetTags() []TagLabel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Task) GetTagsOk() (*[]TagLabel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Task) SetTags(v []TagLabel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Task) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTaskSQL

`func (o *Task) GetTaskSQL() string`

GetTaskSQL returns the TaskSQL field if non-nil, zero value otherwise.

### GetTaskSQLOk

`func (o *Task) GetTaskSQLOk() (*string, bool)`

GetTaskSQLOk returns a tuple with the TaskSQL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSQL

`func (o *Task) SetTaskSQL(v string)`

SetTaskSQL sets TaskSQL field to given value.

### HasTaskSQL

`func (o *Task) HasTaskSQL() bool`

HasTaskSQL returns a boolean if a field has been set.

### GetTaskType

`func (o *Task) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *Task) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *Task) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *Task) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetTaskUrl

`func (o *Task) GetTaskUrl() string`

GetTaskUrl returns the TaskUrl field if non-nil, zero value otherwise.

### GetTaskUrlOk

`func (o *Task) GetTaskUrlOk() (*string, bool)`

GetTaskUrlOk returns a tuple with the TaskUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUrl

`func (o *Task) SetTaskUrl(v string)`

SetTaskUrl sets TaskUrl field to given value.

### HasTaskUrl

`func (o *Task) HasTaskUrl() bool`

HasTaskUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


