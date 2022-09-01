# DatabaseServiceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DatabaseService**](DatabaseService.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewDatabaseServiceList

`func NewDatabaseServiceList(data []DatabaseService, ) *DatabaseServiceList`

NewDatabaseServiceList instantiates a new DatabaseServiceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseServiceListWithDefaults

`func NewDatabaseServiceListWithDefaults() *DatabaseServiceList`

NewDatabaseServiceListWithDefaults instantiates a new DatabaseServiceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DatabaseServiceList) GetData() []DatabaseService`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DatabaseServiceList) GetDataOk() (*[]DatabaseService, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DatabaseServiceList) SetData(v []DatabaseService)`

SetData sets Data field to given value.


### GetPaging

`func (o *DatabaseServiceList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *DatabaseServiceList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *DatabaseServiceList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *DatabaseServiceList) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


