# ThreadList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Thread**](Thread.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewThreadList

`func NewThreadList(data []Thread, ) *ThreadList`

NewThreadList instantiates a new ThreadList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadListWithDefaults

`func NewThreadListWithDefaults() *ThreadList`

NewThreadListWithDefaults instantiates a new ThreadList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ThreadList) GetData() []Thread`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ThreadList) GetDataOk() (*[]Thread, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ThreadList) SetData(v []Thread)`

SetData sets Data field to given value.


### GetPaging

`func (o *ThreadList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *ThreadList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *ThreadList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *ThreadList) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


