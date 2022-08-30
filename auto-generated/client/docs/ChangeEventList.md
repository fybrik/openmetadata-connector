# ChangeEventList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ChangeEvent**](ChangeEvent.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewChangeEventList

`func NewChangeEventList(data []ChangeEvent, ) *ChangeEventList`

NewChangeEventList instantiates a new ChangeEventList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeEventListWithDefaults

`func NewChangeEventListWithDefaults() *ChangeEventList`

NewChangeEventListWithDefaults instantiates a new ChangeEventList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ChangeEventList) GetData() []ChangeEvent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ChangeEventList) GetDataOk() (*[]ChangeEvent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ChangeEventList) SetData(v []ChangeEvent)`

SetData sets Data field to given value.


### GetPaging

`func (o *ChangeEventList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *ChangeEventList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *ChangeEventList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *ChangeEventList) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


