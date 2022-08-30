# TypeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Type**](Type.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewTypeList

`func NewTypeList(data []Type, ) *TypeList`

NewTypeList instantiates a new TypeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypeListWithDefaults

`func NewTypeListWithDefaults() *TypeList`

NewTypeListWithDefaults instantiates a new TypeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TypeList) GetData() []Type`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TypeList) GetDataOk() (*[]Type, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TypeList) SetData(v []Type)`

SetData sets Data field to given value.


### GetPaging

`func (o *TypeList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *TypeList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *TypeList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *TypeList) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


