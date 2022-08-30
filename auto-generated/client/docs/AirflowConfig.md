# AirflowConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Concurrency** | Pointer to **int32** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**MaxActiveRuns** | Pointer to **int32** |  | [optional] 
**PausePipeline** | Pointer to **bool** |  | [optional] 
**PipelineCatchup** | Pointer to **bool** |  | [optional] 
**PipelineTimezone** | Pointer to **string** |  | [optional] 
**Retries** | Pointer to **int32** |  | [optional] 
**RetryDelay** | Pointer to **int32** |  | [optional] 
**ScheduleInterval** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**WorkflowDefaultView** | Pointer to **string** |  | [optional] 
**WorkflowDefaultViewOrientation** | Pointer to **string** |  | [optional] 
**WorkflowTimeout** | Pointer to **int32** |  | [optional] 

## Methods

### NewAirflowConfig

`func NewAirflowConfig() *AirflowConfig`

NewAirflowConfig instantiates a new AirflowConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirflowConfigWithDefaults

`func NewAirflowConfigWithDefaults() *AirflowConfig`

NewAirflowConfigWithDefaults instantiates a new AirflowConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConcurrency

`func (o *AirflowConfig) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *AirflowConfig) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *AirflowConfig) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *AirflowConfig) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### GetEmail

`func (o *AirflowConfig) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AirflowConfig) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AirflowConfig) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AirflowConfig) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEndDate

`func (o *AirflowConfig) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AirflowConfig) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AirflowConfig) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AirflowConfig) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetMaxActiveRuns

`func (o *AirflowConfig) GetMaxActiveRuns() int32`

GetMaxActiveRuns returns the MaxActiveRuns field if non-nil, zero value otherwise.

### GetMaxActiveRunsOk

`func (o *AirflowConfig) GetMaxActiveRunsOk() (*int32, bool)`

GetMaxActiveRunsOk returns a tuple with the MaxActiveRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActiveRuns

`func (o *AirflowConfig) SetMaxActiveRuns(v int32)`

SetMaxActiveRuns sets MaxActiveRuns field to given value.

### HasMaxActiveRuns

`func (o *AirflowConfig) HasMaxActiveRuns() bool`

HasMaxActiveRuns returns a boolean if a field has been set.

### GetPausePipeline

`func (o *AirflowConfig) GetPausePipeline() bool`

GetPausePipeline returns the PausePipeline field if non-nil, zero value otherwise.

### GetPausePipelineOk

`func (o *AirflowConfig) GetPausePipelineOk() (*bool, bool)`

GetPausePipelineOk returns a tuple with the PausePipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausePipeline

`func (o *AirflowConfig) SetPausePipeline(v bool)`

SetPausePipeline sets PausePipeline field to given value.

### HasPausePipeline

`func (o *AirflowConfig) HasPausePipeline() bool`

HasPausePipeline returns a boolean if a field has been set.

### GetPipelineCatchup

`func (o *AirflowConfig) GetPipelineCatchup() bool`

GetPipelineCatchup returns the PipelineCatchup field if non-nil, zero value otherwise.

### GetPipelineCatchupOk

`func (o *AirflowConfig) GetPipelineCatchupOk() (*bool, bool)`

GetPipelineCatchupOk returns a tuple with the PipelineCatchup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineCatchup

`func (o *AirflowConfig) SetPipelineCatchup(v bool)`

SetPipelineCatchup sets PipelineCatchup field to given value.

### HasPipelineCatchup

`func (o *AirflowConfig) HasPipelineCatchup() bool`

HasPipelineCatchup returns a boolean if a field has been set.

### GetPipelineTimezone

`func (o *AirflowConfig) GetPipelineTimezone() string`

GetPipelineTimezone returns the PipelineTimezone field if non-nil, zero value otherwise.

### GetPipelineTimezoneOk

`func (o *AirflowConfig) GetPipelineTimezoneOk() (*string, bool)`

GetPipelineTimezoneOk returns a tuple with the PipelineTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineTimezone

`func (o *AirflowConfig) SetPipelineTimezone(v string)`

SetPipelineTimezone sets PipelineTimezone field to given value.

### HasPipelineTimezone

`func (o *AirflowConfig) HasPipelineTimezone() bool`

HasPipelineTimezone returns a boolean if a field has been set.

### GetRetries

`func (o *AirflowConfig) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *AirflowConfig) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *AirflowConfig) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *AirflowConfig) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetRetryDelay

`func (o *AirflowConfig) GetRetryDelay() int32`

GetRetryDelay returns the RetryDelay field if non-nil, zero value otherwise.

### GetRetryDelayOk

`func (o *AirflowConfig) GetRetryDelayOk() (*int32, bool)`

GetRetryDelayOk returns a tuple with the RetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelay

`func (o *AirflowConfig) SetRetryDelay(v int32)`

SetRetryDelay sets RetryDelay field to given value.

### HasRetryDelay

`func (o *AirflowConfig) HasRetryDelay() bool`

HasRetryDelay returns a boolean if a field has been set.

### GetScheduleInterval

`func (o *AirflowConfig) GetScheduleInterval() string`

GetScheduleInterval returns the ScheduleInterval field if non-nil, zero value otherwise.

### GetScheduleIntervalOk

`func (o *AirflowConfig) GetScheduleIntervalOk() (*string, bool)`

GetScheduleIntervalOk returns a tuple with the ScheduleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleInterval

`func (o *AirflowConfig) SetScheduleInterval(v string)`

SetScheduleInterval sets ScheduleInterval field to given value.

### HasScheduleInterval

`func (o *AirflowConfig) HasScheduleInterval() bool`

HasScheduleInterval returns a boolean if a field has been set.

### GetStartDate

`func (o *AirflowConfig) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AirflowConfig) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AirflowConfig) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AirflowConfig) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetWorkflowDefaultView

`func (o *AirflowConfig) GetWorkflowDefaultView() string`

GetWorkflowDefaultView returns the WorkflowDefaultView field if non-nil, zero value otherwise.

### GetWorkflowDefaultViewOk

`func (o *AirflowConfig) GetWorkflowDefaultViewOk() (*string, bool)`

GetWorkflowDefaultViewOk returns a tuple with the WorkflowDefaultView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefaultView

`func (o *AirflowConfig) SetWorkflowDefaultView(v string)`

SetWorkflowDefaultView sets WorkflowDefaultView field to given value.

### HasWorkflowDefaultView

`func (o *AirflowConfig) HasWorkflowDefaultView() bool`

HasWorkflowDefaultView returns a boolean if a field has been set.

### GetWorkflowDefaultViewOrientation

`func (o *AirflowConfig) GetWorkflowDefaultViewOrientation() string`

GetWorkflowDefaultViewOrientation returns the WorkflowDefaultViewOrientation field if non-nil, zero value otherwise.

### GetWorkflowDefaultViewOrientationOk

`func (o *AirflowConfig) GetWorkflowDefaultViewOrientationOk() (*string, bool)`

GetWorkflowDefaultViewOrientationOk returns a tuple with the WorkflowDefaultViewOrientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefaultViewOrientation

`func (o *AirflowConfig) SetWorkflowDefaultViewOrientation(v string)`

SetWorkflowDefaultViewOrientation sets WorkflowDefaultViewOrientation field to given value.

### HasWorkflowDefaultViewOrientation

`func (o *AirflowConfig) HasWorkflowDefaultViewOrientation() bool`

HasWorkflowDefaultViewOrientation returns a boolean if a field has been set.

### GetWorkflowTimeout

`func (o *AirflowConfig) GetWorkflowTimeout() int32`

GetWorkflowTimeout returns the WorkflowTimeout field if non-nil, zero value otherwise.

### GetWorkflowTimeoutOk

`func (o *AirflowConfig) GetWorkflowTimeoutOk() (*int32, bool)`

GetWorkflowTimeoutOk returns a tuple with the WorkflowTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTimeout

`func (o *AirflowConfig) SetWorkflowTimeout(v int32)`

SetWorkflowTimeout sets WorkflowTimeout field to given value.

### HasWorkflowTimeout

`func (o *AirflowConfig) HasWorkflowTimeout() bool`

HasWorkflowTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


