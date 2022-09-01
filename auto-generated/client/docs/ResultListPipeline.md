# ResultListPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Pipeline**](Pipeline.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewResultListPipeline

`func NewResultListPipeline(data []Pipeline, ) *ResultListPipeline`

NewResultListPipeline instantiates a new ResultListPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultListPipelineWithDefaults

`func NewResultListPipelineWithDefaults() *ResultListPipeline`

NewResultListPipelineWithDefaults instantiates a new ResultListPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResultListPipeline) GetData() []Pipeline`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultListPipeline) GetDataOk() (*[]Pipeline, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultListPipeline) SetData(v []Pipeline)`

SetData sets Data field to given value.


### GetPaging

`func (o *ResultListPipeline) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *ResultListPipeline) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *ResultListPipeline) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *ResultListPipeline) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


