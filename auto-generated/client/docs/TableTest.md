# TableTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**ExecutionFrequency** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Results** | Pointer to [**[]TestCaseResult**](TestCaseResult.md) |  | [optional] 
**TestCase** | [**TableTestCase**](TableTestCase.md) |  | 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewTableTest

`func NewTableTest(name string, testCase TableTestCase, ) *TableTest`

NewTableTest instantiates a new TableTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableTestWithDefaults

`func NewTableTestWithDefaults() *TableTest`

NewTableTestWithDefaults instantiates a new TableTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TableTest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TableTest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TableTest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TableTest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecutionFrequency

`func (o *TableTest) GetExecutionFrequency() string`

GetExecutionFrequency returns the ExecutionFrequency field if non-nil, zero value otherwise.

### GetExecutionFrequencyOk

`func (o *TableTest) GetExecutionFrequencyOk() (*string, bool)`

GetExecutionFrequencyOk returns a tuple with the ExecutionFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionFrequency

`func (o *TableTest) SetExecutionFrequency(v string)`

SetExecutionFrequency sets ExecutionFrequency field to given value.

### HasExecutionFrequency

`func (o *TableTest) HasExecutionFrequency() bool`

HasExecutionFrequency returns a boolean if a field has been set.

### GetId

`func (o *TableTest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TableTest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TableTest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TableTest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TableTest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TableTest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TableTest) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *TableTest) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *TableTest) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *TableTest) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *TableTest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetResults

`func (o *TableTest) GetResults() []TestCaseResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TableTest) GetResultsOk() (*[]TestCaseResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TableTest) SetResults(v []TestCaseResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *TableTest) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTestCase

`func (o *TableTest) GetTestCase() TableTestCase`

GetTestCase returns the TestCase field if non-nil, zero value otherwise.

### GetTestCaseOk

`func (o *TableTest) GetTestCaseOk() (*TableTestCase, bool)`

GetTestCaseOk returns a tuple with the TestCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCase

`func (o *TableTest) SetTestCase(v TableTestCase)`

SetTestCase sets TestCase field to given value.


### GetUpdatedAt

`func (o *TableTest) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TableTest) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TableTest) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TableTest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *TableTest) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *TableTest) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *TableTest) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *TableTest) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


