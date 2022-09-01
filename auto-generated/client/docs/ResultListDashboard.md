# ResultListDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Dashboard**](Dashboard.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewResultListDashboard

`func NewResultListDashboard(data []Dashboard, ) *ResultListDashboard`

NewResultListDashboard instantiates a new ResultListDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultListDashboardWithDefaults

`func NewResultListDashboardWithDefaults() *ResultListDashboard`

NewResultListDashboardWithDefaults instantiates a new ResultListDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResultListDashboard) GetData() []Dashboard`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultListDashboard) GetDataOk() (*[]Dashboard, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultListDashboard) SetData(v []Dashboard)`

SetData sets Data field to given value.


### GetPaging

`func (o *ResultListDashboard) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *ResultListDashboard) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *ResultListDashboard) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *ResultListDashboard) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


