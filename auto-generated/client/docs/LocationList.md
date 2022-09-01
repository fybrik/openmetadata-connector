# LocationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Location**](Location.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewLocationList

`func NewLocationList(data []Location, ) *LocationList`

NewLocationList instantiates a new LocationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationListWithDefaults

`func NewLocationListWithDefaults() *LocationList`

NewLocationListWithDefaults instantiates a new LocationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *LocationList) GetData() []Location`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LocationList) GetDataOk() (*[]Location, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LocationList) SetData(v []Location)`

SetData sets Data field to given value.


### GetPaging

`func (o *LocationList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *LocationList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *LocationList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *LocationList) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


