# UsageStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**PercentileRank** | Pointer to **float64** |  | [optional] 

## Methods

### NewUsageStats

`func NewUsageStats(count int32, ) *UsageStats`

NewUsageStats instantiates a new UsageStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageStatsWithDefaults

`func NewUsageStatsWithDefaults() *UsageStats`

NewUsageStatsWithDefaults instantiates a new UsageStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UsageStats) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UsageStats) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UsageStats) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPercentileRank

`func (o *UsageStats) GetPercentileRank() float64`

GetPercentileRank returns the PercentileRank field if non-nil, zero value otherwise.

### GetPercentileRankOk

`func (o *UsageStats) GetPercentileRankOk() (*float64, bool)`

GetPercentileRankOk returns a tuple with the PercentileRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileRank

`func (o *UsageStats) SetPercentileRank(v float64)`

SetPercentileRank sets PercentileRank field to given value.

### HasPercentileRank

`func (o *UsageStats) HasPercentileRank() bool`

HasPercentileRank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


