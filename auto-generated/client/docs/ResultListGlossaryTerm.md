# ResultListGlossaryTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GlossaryTerm**](GlossaryTerm.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewResultListGlossaryTerm

`func NewResultListGlossaryTerm(data []GlossaryTerm, ) *ResultListGlossaryTerm`

NewResultListGlossaryTerm instantiates a new ResultListGlossaryTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultListGlossaryTermWithDefaults

`func NewResultListGlossaryTermWithDefaults() *ResultListGlossaryTerm`

NewResultListGlossaryTermWithDefaults instantiates a new ResultListGlossaryTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResultListGlossaryTerm) GetData() []GlossaryTerm`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultListGlossaryTerm) GetDataOk() (*[]GlossaryTerm, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultListGlossaryTerm) SetData(v []GlossaryTerm)`

SetData sets Data field to given value.


### GetPaging

`func (o *ResultListGlossaryTerm) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *ResultListGlossaryTerm) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *ResultListGlossaryTerm) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *ResultListGlossaryTerm) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


