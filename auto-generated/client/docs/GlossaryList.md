# GlossaryList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Glossary**](Glossary.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewGlossaryList

`func NewGlossaryList(data []Glossary, ) *GlossaryList`

NewGlossaryList instantiates a new GlossaryList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlossaryListWithDefaults

`func NewGlossaryListWithDefaults() *GlossaryList`

NewGlossaryListWithDefaults instantiates a new GlossaryList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GlossaryList) GetData() []Glossary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GlossaryList) GetDataOk() (*[]Glossary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GlossaryList) SetData(v []Glossary)`

SetData sets Data field to given value.


### GetPaging

`func (o *GlossaryList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *GlossaryList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *GlossaryList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *GlossaryList) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


