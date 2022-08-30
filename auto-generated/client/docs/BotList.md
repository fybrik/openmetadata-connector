# BotList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Bot**](Bot.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewBotList

`func NewBotList(data []Bot, ) *BotList`

NewBotList instantiates a new BotList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotListWithDefaults

`func NewBotListWithDefaults() *BotList`

NewBotListWithDefaults instantiates a new BotList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BotList) GetData() []Bot`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BotList) GetDataOk() (*[]Bot, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BotList) SetData(v []Bot)`

SetData sets Data field to given value.


### GetPaging

`func (o *BotList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *BotList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *BotList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *BotList) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


