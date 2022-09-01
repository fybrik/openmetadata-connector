# MessagingServiceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]MessagingService**](MessagingService.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewMessagingServiceList

`func NewMessagingServiceList(data []MessagingService, ) *MessagingServiceList`

NewMessagingServiceList instantiates a new MessagingServiceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagingServiceListWithDefaults

`func NewMessagingServiceListWithDefaults() *MessagingServiceList`

NewMessagingServiceListWithDefaults instantiates a new MessagingServiceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MessagingServiceList) GetData() []MessagingService`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MessagingServiceList) GetDataOk() (*[]MessagingService, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MessagingServiceList) SetData(v []MessagingService)`

SetData sets Data field to given value.


### GetPaging

`func (o *MessagingServiceList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *MessagingServiceList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *MessagingServiceList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *MessagingServiceList) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


