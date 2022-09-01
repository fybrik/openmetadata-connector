# TimeValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **int64** |  | [optional] 
**DaysFrac** | Pointer to **float64** |  | [optional] 
**Hours** | Pointer to **int64** |  | [optional] 
**HoursFrac** | Pointer to **float64** |  | [optional] 
**Micros** | Pointer to **int64** |  | [optional] 
**MicrosFrac** | Pointer to **float64** |  | [optional] 
**Millis** | Pointer to **int64** |  | [optional] 
**MillisFrac** | Pointer to **float64** |  | [optional] 
**Minutes** | Pointer to **int64** |  | [optional] 
**MinutesFrac** | Pointer to **float64** |  | [optional] 
**Nanos** | Pointer to **int64** |  | [optional] 
**Seconds** | Pointer to **int64** |  | [optional] 
**SecondsFrac** | Pointer to **float64** |  | [optional] 
**StringRep** | Pointer to **string** |  | [optional] 

## Methods

### NewTimeValue

`func NewTimeValue() *TimeValue`

NewTimeValue instantiates a new TimeValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeValueWithDefaults

`func NewTimeValueWithDefaults() *TimeValue`

NewTimeValueWithDefaults instantiates a new TimeValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *TimeValue) GetDays() int64`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *TimeValue) GetDaysOk() (*int64, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *TimeValue) SetDays(v int64)`

SetDays sets Days field to given value.

### HasDays

`func (o *TimeValue) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetDaysFrac

`func (o *TimeValue) GetDaysFrac() float64`

GetDaysFrac returns the DaysFrac field if non-nil, zero value otherwise.

### GetDaysFracOk

`func (o *TimeValue) GetDaysFracOk() (*float64, bool)`

GetDaysFracOk returns a tuple with the DaysFrac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysFrac

`func (o *TimeValue) SetDaysFrac(v float64)`

SetDaysFrac sets DaysFrac field to given value.

### HasDaysFrac

`func (o *TimeValue) HasDaysFrac() bool`

HasDaysFrac returns a boolean if a field has been set.

### GetHours

`func (o *TimeValue) GetHours() int64`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *TimeValue) GetHoursOk() (*int64, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *TimeValue) SetHours(v int64)`

SetHours sets Hours field to given value.

### HasHours

`func (o *TimeValue) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetHoursFrac

`func (o *TimeValue) GetHoursFrac() float64`

GetHoursFrac returns the HoursFrac field if non-nil, zero value otherwise.

### GetHoursFracOk

`func (o *TimeValue) GetHoursFracOk() (*float64, bool)`

GetHoursFracOk returns a tuple with the HoursFrac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursFrac

`func (o *TimeValue) SetHoursFrac(v float64)`

SetHoursFrac sets HoursFrac field to given value.

### HasHoursFrac

`func (o *TimeValue) HasHoursFrac() bool`

HasHoursFrac returns a boolean if a field has been set.

### GetMicros

`func (o *TimeValue) GetMicros() int64`

GetMicros returns the Micros field if non-nil, zero value otherwise.

### GetMicrosOk

`func (o *TimeValue) GetMicrosOk() (*int64, bool)`

GetMicrosOk returns a tuple with the Micros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicros

`func (o *TimeValue) SetMicros(v int64)`

SetMicros sets Micros field to given value.

### HasMicros

`func (o *TimeValue) HasMicros() bool`

HasMicros returns a boolean if a field has been set.

### GetMicrosFrac

`func (o *TimeValue) GetMicrosFrac() float64`

GetMicrosFrac returns the MicrosFrac field if non-nil, zero value otherwise.

### GetMicrosFracOk

`func (o *TimeValue) GetMicrosFracOk() (*float64, bool)`

GetMicrosFracOk returns a tuple with the MicrosFrac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosFrac

`func (o *TimeValue) SetMicrosFrac(v float64)`

SetMicrosFrac sets MicrosFrac field to given value.

### HasMicrosFrac

`func (o *TimeValue) HasMicrosFrac() bool`

HasMicrosFrac returns a boolean if a field has been set.

### GetMillis

`func (o *TimeValue) GetMillis() int64`

GetMillis returns the Millis field if non-nil, zero value otherwise.

### GetMillisOk

`func (o *TimeValue) GetMillisOk() (*int64, bool)`

GetMillisOk returns a tuple with the Millis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillis

`func (o *TimeValue) SetMillis(v int64)`

SetMillis sets Millis field to given value.

### HasMillis

`func (o *TimeValue) HasMillis() bool`

HasMillis returns a boolean if a field has been set.

### GetMillisFrac

`func (o *TimeValue) GetMillisFrac() float64`

GetMillisFrac returns the MillisFrac field if non-nil, zero value otherwise.

### GetMillisFracOk

`func (o *TimeValue) GetMillisFracOk() (*float64, bool)`

GetMillisFracOk returns a tuple with the MillisFrac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillisFrac

`func (o *TimeValue) SetMillisFrac(v float64)`

SetMillisFrac sets MillisFrac field to given value.

### HasMillisFrac

`func (o *TimeValue) HasMillisFrac() bool`

HasMillisFrac returns a boolean if a field has been set.

### GetMinutes

`func (o *TimeValue) GetMinutes() int64`

GetMinutes returns the Minutes field if non-nil, zero value otherwise.

### GetMinutesOk

`func (o *TimeValue) GetMinutesOk() (*int64, bool)`

GetMinutesOk returns a tuple with the Minutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutes

`func (o *TimeValue) SetMinutes(v int64)`

SetMinutes sets Minutes field to given value.

### HasMinutes

`func (o *TimeValue) HasMinutes() bool`

HasMinutes returns a boolean if a field has been set.

### GetMinutesFrac

`func (o *TimeValue) GetMinutesFrac() float64`

GetMinutesFrac returns the MinutesFrac field if non-nil, zero value otherwise.

### GetMinutesFracOk

`func (o *TimeValue) GetMinutesFracOk() (*float64, bool)`

GetMinutesFracOk returns a tuple with the MinutesFrac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesFrac

`func (o *TimeValue) SetMinutesFrac(v float64)`

SetMinutesFrac sets MinutesFrac field to given value.

### HasMinutesFrac

`func (o *TimeValue) HasMinutesFrac() bool`

HasMinutesFrac returns a boolean if a field has been set.

### GetNanos

`func (o *TimeValue) GetNanos() int64`

GetNanos returns the Nanos field if non-nil, zero value otherwise.

### GetNanosOk

`func (o *TimeValue) GetNanosOk() (*int64, bool)`

GetNanosOk returns a tuple with the Nanos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNanos

`func (o *TimeValue) SetNanos(v int64)`

SetNanos sets Nanos field to given value.

### HasNanos

`func (o *TimeValue) HasNanos() bool`

HasNanos returns a boolean if a field has been set.

### GetSeconds

`func (o *TimeValue) GetSeconds() int64`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *TimeValue) GetSecondsOk() (*int64, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *TimeValue) SetSeconds(v int64)`

SetSeconds sets Seconds field to given value.

### HasSeconds

`func (o *TimeValue) HasSeconds() bool`

HasSeconds returns a boolean if a field has been set.

### GetSecondsFrac

`func (o *TimeValue) GetSecondsFrac() float64`

GetSecondsFrac returns the SecondsFrac field if non-nil, zero value otherwise.

### GetSecondsFracOk

`func (o *TimeValue) GetSecondsFracOk() (*float64, bool)`

GetSecondsFracOk returns a tuple with the SecondsFrac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsFrac

`func (o *TimeValue) SetSecondsFrac(v float64)`

SetSecondsFrac sets SecondsFrac field to given value.

### HasSecondsFrac

`func (o *TimeValue) HasSecondsFrac() bool`

HasSecondsFrac returns a boolean if a field has been set.

### GetStringRep

`func (o *TimeValue) GetStringRep() string`

GetStringRep returns the StringRep field if non-nil, zero value otherwise.

### GetStringRepOk

`func (o *TimeValue) GetStringRepOk() (*string, bool)`

GetStringRepOk returns a tuple with the StringRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringRep

`func (o *TimeValue) SetStringRep(v string)`

SetStringRep sets StringRep field to given value.

### HasStringRep

`func (o *TimeValue) HasStringRep() bool`

HasStringRep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


