# TableJoins

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnJoins** | Pointer to [**[]ColumnJoin**](ColumnJoin.md) |  | [optional] 
**DayCount** | Pointer to **int32** |  | [optional] 
**DirectTableJoins** | Pointer to [**[]JoinedWith**](JoinedWith.md) |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 

## Methods

### NewTableJoins

`func NewTableJoins() *TableJoins`

NewTableJoins instantiates a new TableJoins object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableJoinsWithDefaults

`func NewTableJoinsWithDefaults() *TableJoins`

NewTableJoinsWithDefaults instantiates a new TableJoins object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnJoins

`func (o *TableJoins) GetColumnJoins() []ColumnJoin`

GetColumnJoins returns the ColumnJoins field if non-nil, zero value otherwise.

### GetColumnJoinsOk

`func (o *TableJoins) GetColumnJoinsOk() (*[]ColumnJoin, bool)`

GetColumnJoinsOk returns a tuple with the ColumnJoins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnJoins

`func (o *TableJoins) SetColumnJoins(v []ColumnJoin)`

SetColumnJoins sets ColumnJoins field to given value.

### HasColumnJoins

`func (o *TableJoins) HasColumnJoins() bool`

HasColumnJoins returns a boolean if a field has been set.

### GetDayCount

`func (o *TableJoins) GetDayCount() int32`

GetDayCount returns the DayCount field if non-nil, zero value otherwise.

### GetDayCountOk

`func (o *TableJoins) GetDayCountOk() (*int32, bool)`

GetDayCountOk returns a tuple with the DayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayCount

`func (o *TableJoins) SetDayCount(v int32)`

SetDayCount sets DayCount field to given value.

### HasDayCount

`func (o *TableJoins) HasDayCount() bool`

HasDayCount returns a boolean if a field has been set.

### GetDirectTableJoins

`func (o *TableJoins) GetDirectTableJoins() []JoinedWith`

GetDirectTableJoins returns the DirectTableJoins field if non-nil, zero value otherwise.

### GetDirectTableJoinsOk

`func (o *TableJoins) GetDirectTableJoinsOk() (*[]JoinedWith, bool)`

GetDirectTableJoinsOk returns a tuple with the DirectTableJoins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectTableJoins

`func (o *TableJoins) SetDirectTableJoins(v []JoinedWith)`

SetDirectTableJoins sets DirectTableJoins field to given value.

### HasDirectTableJoins

`func (o *TableJoins) HasDirectTableJoins() bool`

HasDirectTableJoins returns a boolean if a field has been set.

### GetStartDate

`func (o *TableJoins) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TableJoins) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TableJoins) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TableJoins) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


