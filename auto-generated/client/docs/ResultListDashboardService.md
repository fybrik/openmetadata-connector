# ResultListDashboardService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DashboardService**](DashboardService.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewResultListDashboardService

`func NewResultListDashboardService(data []DashboardService, ) *ResultListDashboardService`

NewResultListDashboardService instantiates a new ResultListDashboardService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultListDashboardServiceWithDefaults

`func NewResultListDashboardServiceWithDefaults() *ResultListDashboardService`

NewResultListDashboardServiceWithDefaults instantiates a new ResultListDashboardService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResultListDashboardService) GetData() []DashboardService`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultListDashboardService) GetDataOk() (*[]DashboardService, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultListDashboardService) SetData(v []DashboardService)`

SetData sets Data field to given value.


### GetPaging

`func (o *ResultListDashboardService) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *ResultListDashboardService) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *ResultListDashboardService) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *ResultListDashboardService) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


