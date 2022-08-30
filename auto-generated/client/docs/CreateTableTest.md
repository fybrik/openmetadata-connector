# CreateTableTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**ExecutionFrequency** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Result** | Pointer to [**TestCaseResult**](TestCaseResult.md) |  | [optional] 
**TestCase** | [**TableTestCase**](TableTestCase.md) |  | 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateTableTest

`func NewCreateTableTest(testCase TableTestCase, ) *CreateTableTest`

NewCreateTableTest instantiates a new CreateTableTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTableTestWithDefaults

`func NewCreateTableTestWithDefaults() *CreateTableTest`

NewCreateTableTestWithDefaults instantiates a new CreateTableTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateTableTest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTableTest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTableTest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTableTest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecutionFrequency

`func (o *CreateTableTest) GetExecutionFrequency() string`

GetExecutionFrequency returns the ExecutionFrequency field if non-nil, zero value otherwise.

### GetExecutionFrequencyOk

`func (o *CreateTableTest) GetExecutionFrequencyOk() (*string, bool)`

GetExecutionFrequencyOk returns a tuple with the ExecutionFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionFrequency

`func (o *CreateTableTest) SetExecutionFrequency(v string)`

SetExecutionFrequency sets ExecutionFrequency field to given value.

### HasExecutionFrequency

`func (o *CreateTableTest) HasExecutionFrequency() bool`

HasExecutionFrequency returns a boolean if a field has been set.

### GetOwner

`func (o *CreateTableTest) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateTableTest) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateTableTest) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateTableTest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetResult

`func (o *CreateTableTest) GetResult() TestCaseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CreateTableTest) GetResultOk() (*TestCaseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CreateTableTest) SetResult(v TestCaseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *CreateTableTest) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetTestCase

`func (o *CreateTableTest) GetTestCase() TableTestCase`

GetTestCase returns the TestCase field if non-nil, zero value otherwise.

### GetTestCaseOk

`func (o *CreateTableTest) GetTestCaseOk() (*TableTestCase, bool)`

GetTestCaseOk returns a tuple with the TestCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCase

`func (o *CreateTableTest) SetTestCase(v TableTestCase)`

SetTestCase sets TestCase field to given value.


### GetUpdatedAt

`func (o *CreateTableTest) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateTableTest) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateTableTest) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateTableTest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *CreateTableTest) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CreateTableTest) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *CreateTableTest) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *CreateTableTest) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


