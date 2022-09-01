# CreateColumnTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnName** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ExecutionFrequency** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Result** | Pointer to [**TestCaseResult**](TestCaseResult.md) |  | [optional] 
**TestCase** | [**ColumnTestCase**](ColumnTestCase.md) |  | 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateColumnTest

`func NewCreateColumnTest(columnName string, testCase ColumnTestCase, ) *CreateColumnTest`

NewCreateColumnTest instantiates a new CreateColumnTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateColumnTestWithDefaults

`func NewCreateColumnTestWithDefaults() *CreateColumnTest`

NewCreateColumnTestWithDefaults instantiates a new CreateColumnTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnName

`func (o *CreateColumnTest) GetColumnName() string`

GetColumnName returns the ColumnName field if non-nil, zero value otherwise.

### GetColumnNameOk

`func (o *CreateColumnTest) GetColumnNameOk() (*string, bool)`

GetColumnNameOk returns a tuple with the ColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnName

`func (o *CreateColumnTest) SetColumnName(v string)`

SetColumnName sets ColumnName field to given value.


### GetDescription

`func (o *CreateColumnTest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateColumnTest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateColumnTest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateColumnTest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecutionFrequency

`func (o *CreateColumnTest) GetExecutionFrequency() string`

GetExecutionFrequency returns the ExecutionFrequency field if non-nil, zero value otherwise.

### GetExecutionFrequencyOk

`func (o *CreateColumnTest) GetExecutionFrequencyOk() (*string, bool)`

GetExecutionFrequencyOk returns a tuple with the ExecutionFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionFrequency

`func (o *CreateColumnTest) SetExecutionFrequency(v string)`

SetExecutionFrequency sets ExecutionFrequency field to given value.

### HasExecutionFrequency

`func (o *CreateColumnTest) HasExecutionFrequency() bool`

HasExecutionFrequency returns a boolean if a field has been set.

### GetOwner

`func (o *CreateColumnTest) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateColumnTest) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateColumnTest) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateColumnTest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetResult

`func (o *CreateColumnTest) GetResult() TestCaseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CreateColumnTest) GetResultOk() (*TestCaseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CreateColumnTest) SetResult(v TestCaseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *CreateColumnTest) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetTestCase

`func (o *CreateColumnTest) GetTestCase() ColumnTestCase`

GetTestCase returns the TestCase field if non-nil, zero value otherwise.

### GetTestCaseOk

`func (o *CreateColumnTest) GetTestCaseOk() (*ColumnTestCase, bool)`

GetTestCaseOk returns a tuple with the TestCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCase

`func (o *CreateColumnTest) SetTestCase(v ColumnTestCase)`

SetTestCase sets TestCase field to given value.


### GetUpdatedAt

`func (o *CreateColumnTest) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateColumnTest) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateColumnTest) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateColumnTest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *CreateColumnTest) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CreateColumnTest) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *CreateColumnTest) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *CreateColumnTest) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


