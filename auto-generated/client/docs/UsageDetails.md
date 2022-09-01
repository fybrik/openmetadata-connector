# UsageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DailyStats** | [**UsageStats**](UsageStats.md) |  | 
**Date** | **string** |  | 
**MonthlyStats** | Pointer to [**UsageStats**](UsageStats.md) |  | [optional] 
**WeeklyStats** | Pointer to [**UsageStats**](UsageStats.md) |  | [optional] 

## Methods

### NewUsageDetails

`func NewUsageDetails(dailyStats UsageStats, date string, ) *UsageDetails`

NewUsageDetails instantiates a new UsageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageDetailsWithDefaults

`func NewUsageDetailsWithDefaults() *UsageDetails`

NewUsageDetailsWithDefaults instantiates a new UsageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDailyStats

`func (o *UsageDetails) GetDailyStats() UsageStats`

GetDailyStats returns the DailyStats field if non-nil, zero value otherwise.

### GetDailyStatsOk

`func (o *UsageDetails) GetDailyStatsOk() (*UsageStats, bool)`

GetDailyStatsOk returns a tuple with the DailyStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyStats

`func (o *UsageDetails) SetDailyStats(v UsageStats)`

SetDailyStats sets DailyStats field to given value.


### GetDate

`func (o *UsageDetails) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *UsageDetails) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *UsageDetails) SetDate(v string)`

SetDate sets Date field to given value.


### GetMonthlyStats

`func (o *UsageDetails) GetMonthlyStats() UsageStats`

GetMonthlyStats returns the MonthlyStats field if non-nil, zero value otherwise.

### GetMonthlyStatsOk

`func (o *UsageDetails) GetMonthlyStatsOk() (*UsageStats, bool)`

GetMonthlyStatsOk returns a tuple with the MonthlyStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyStats

`func (o *UsageDetails) SetMonthlyStats(v UsageStats)`

SetMonthlyStats sets MonthlyStats field to given value.

### HasMonthlyStats

`func (o *UsageDetails) HasMonthlyStats() bool`

HasMonthlyStats returns a boolean if a field has been set.

### GetWeeklyStats

`func (o *UsageDetails) GetWeeklyStats() UsageStats`

GetWeeklyStats returns the WeeklyStats field if non-nil, zero value otherwise.

### GetWeeklyStatsOk

`func (o *UsageDetails) GetWeeklyStatsOk() (*UsageStats, bool)`

GetWeeklyStatsOk returns a tuple with the WeeklyStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklyStats

`func (o *UsageDetails) SetWeeklyStats(v UsageStats)`

SetWeeklyStats sets WeeklyStats field to given value.

### HasWeeklyStats

`func (o *UsageDetails) HasWeeklyStats() bool`

HasWeeklyStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


