# ColumnTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnName** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ExecutionFrequency** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Results** | Pointer to [**[]TestCaseResult**](TestCaseResult.md) |  | [optional] 
**TestCase** | [**ColumnTestCase**](ColumnTestCase.md) |  | 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewColumnTest

`func NewColumnTest(columnName string, name string, testCase ColumnTestCase, ) *ColumnTest`

NewColumnTest instantiates a new ColumnTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColumnTestWithDefaults

`func NewColumnTestWithDefaults() *ColumnTest`

NewColumnTestWithDefaults instantiates a new ColumnTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnName

`func (o *ColumnTest) GetColumnName() string`

GetColumnName returns the ColumnName field if non-nil, zero value otherwise.

### GetColumnNameOk

`func (o *ColumnTest) GetColumnNameOk() (*string, bool)`

GetColumnNameOk returns a tuple with the ColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnName

`func (o *ColumnTest) SetColumnName(v string)`

SetColumnName sets ColumnName field to given value.


### GetDescription

`func (o *ColumnTest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ColumnTest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ColumnTest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ColumnTest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecutionFrequency

`func (o *ColumnTest) GetExecutionFrequency() string`

GetExecutionFrequency returns the ExecutionFrequency field if non-nil, zero value otherwise.

### GetExecutionFrequencyOk

`func (o *ColumnTest) GetExecutionFrequencyOk() (*string, bool)`

GetExecutionFrequencyOk returns a tuple with the ExecutionFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionFrequency

`func (o *ColumnTest) SetExecutionFrequency(v string)`

SetExecutionFrequency sets ExecutionFrequency field to given value.

### HasExecutionFrequency

`func (o *ColumnTest) HasExecutionFrequency() bool`

HasExecutionFrequency returns a boolean if a field has been set.

### GetId

`func (o *ColumnTest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ColumnTest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ColumnTest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ColumnTest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ColumnTest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ColumnTest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ColumnTest) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *ColumnTest) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ColumnTest) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ColumnTest) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ColumnTest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetResults

`func (o *ColumnTest) GetResults() []TestCaseResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ColumnTest) GetResultsOk() (*[]TestCaseResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ColumnTest) SetResults(v []TestCaseResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *ColumnTest) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTestCase

`func (o *ColumnTest) GetTestCase() ColumnTestCase`

GetTestCase returns the TestCase field if non-nil, zero value otherwise.

### GetTestCaseOk

`func (o *ColumnTest) GetTestCaseOk() (*ColumnTestCase, bool)`

GetTestCaseOk returns a tuple with the TestCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCase

`func (o *ColumnTest) SetTestCase(v ColumnTestCase)`

SetTestCase sets TestCase field to given value.


### GetUpdatedAt

`func (o *ColumnTest) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ColumnTest) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ColumnTest) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ColumnTest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ColumnTest) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ColumnTest) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ColumnTest) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ColumnTest) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


