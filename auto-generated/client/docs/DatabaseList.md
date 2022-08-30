# DatabaseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Database**](Database.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewDatabaseList

`func NewDatabaseList(data []Database, ) *DatabaseList`

NewDatabaseList instantiates a new DatabaseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseListWithDefaults

`func NewDatabaseListWithDefaults() *DatabaseList`

NewDatabaseListWithDefaults instantiates a new DatabaseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DatabaseList) GetData() []Database`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DatabaseList) GetDataOk() (*[]Database, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DatabaseList) SetData(v []Database)`

SetData sets Data field to given value.


### GetPaging

`func (o *DatabaseList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *DatabaseList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *DatabaseList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *DatabaseList) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


