# LineageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnsLineage** | [**[]ColumnLineage**](ColumnLineage.md) |  | 
**Pipeline** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**SqlQuery** | **string** |  | 

## Methods

### NewLineageDetails

`func NewLineageDetails(columnsLineage []ColumnLineage, sqlQuery string, ) *LineageDetails`

NewLineageDetails instantiates a new LineageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineageDetailsWithDefaults

`func NewLineageDetailsWithDefaults() *LineageDetails`

NewLineageDetailsWithDefaults instantiates a new LineageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnsLineage

`func (o *LineageDetails) GetColumnsLineage() []ColumnLineage`

GetColumnsLineage returns the ColumnsLineage field if non-nil, zero value otherwise.

### GetColumnsLineageOk

`func (o *LineageDetails) GetColumnsLineageOk() (*[]ColumnLineage, bool)`

GetColumnsLineageOk returns a tuple with the ColumnsLineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnsLineage

`func (o *LineageDetails) SetColumnsLineage(v []ColumnLineage)`

SetColumnsLineage sets ColumnsLineage field to given value.


### GetPipeline

`func (o *LineageDetails) GetPipeline() EntityReference`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *LineageDetails) GetPipelineOk() (*EntityReference, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *LineageDetails) SetPipeline(v EntityReference)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *LineageDetails) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetSqlQuery

`func (o *LineageDetails) GetSqlQuery() string`

GetSqlQuery returns the SqlQuery field if non-nil, zero value otherwise.

### GetSqlQueryOk

`func (o *LineageDetails) GetSqlQueryOk() (*string, bool)`

GetSqlQueryOk returns a tuple with the SqlQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlQuery

`func (o *LineageDetails) SetSqlQuery(v string)`

SetSqlQuery sets SqlQuery field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


