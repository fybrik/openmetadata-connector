# MlModelServiceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]MlModelService**](MlModelService.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewMlModelServiceList

`func NewMlModelServiceList(data []MlModelService, ) *MlModelServiceList`

NewMlModelServiceList instantiates a new MlModelServiceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlModelServiceListWithDefaults

`func NewMlModelServiceListWithDefaults() *MlModelServiceList`

NewMlModelServiceListWithDefaults instantiates a new MlModelServiceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MlModelServiceList) GetData() []MlModelService`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MlModelServiceList) GetDataOk() (*[]MlModelService, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MlModelServiceList) SetData(v []MlModelService)`

SetData sets Data field to given value.


### GetPaging

`func (o *MlModelServiceList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *MlModelServiceList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *MlModelServiceList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *MlModelServiceList) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


