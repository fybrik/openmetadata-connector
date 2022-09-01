# ThreadCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counts** | [**[]EntityLinkThreadCount**](EntityLinkThreadCount.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewThreadCount

`func NewThreadCount(counts []EntityLinkThreadCount, totalCount int32, ) *ThreadCount`

NewThreadCount instantiates a new ThreadCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadCountWithDefaults

`func NewThreadCountWithDefaults() *ThreadCount`

NewThreadCountWithDefaults instantiates a new ThreadCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounts

`func (o *ThreadCount) GetCounts() []EntityLinkThreadCount`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *ThreadCount) GetCountsOk() (*[]EntityLinkThreadCount, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *ThreadCount) SetCounts(v []EntityLinkThreadCount)`

SetCounts sets Counts field to given value.


### GetTotalCount

`func (o *ThreadCount) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ThreadCount) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ThreadCount) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


