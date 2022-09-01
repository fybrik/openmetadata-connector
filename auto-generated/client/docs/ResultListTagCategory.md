# ResultListTagCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TagCategory**](TagCategory.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewResultListTagCategory

`func NewResultListTagCategory(data []TagCategory, ) *ResultListTagCategory`

NewResultListTagCategory instantiates a new ResultListTagCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultListTagCategoryWithDefaults

`func NewResultListTagCategoryWithDefaults() *ResultListTagCategory`

NewResultListTagCategoryWithDefaults instantiates a new ResultListTagCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResultListTagCategory) GetData() []TagCategory`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultListTagCategory) GetDataOk() (*[]TagCategory, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultListTagCategory) SetData(v []TagCategory)`

SetData sets Data field to given value.


### GetPaging

`func (o *ResultListTagCategory) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *ResultListTagCategory) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *ResultListTagCategory) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *ResultListTagCategory) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


