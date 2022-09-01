# TablePartition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | **[]string** |  | 
**Interval** | **string** |  | 
**IntervalType** | **string** |  | 

## Methods

### NewTablePartition

`func NewTablePartition(columns []string, interval string, intervalType string, ) *TablePartition`

NewTablePartition instantiates a new TablePartition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTablePartitionWithDefaults

`func NewTablePartitionWithDefaults() *TablePartition`

NewTablePartitionWithDefaults instantiates a new TablePartition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *TablePartition) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TablePartition) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TablePartition) SetColumns(v []string)`

SetColumns sets Columns field to given value.


### GetInterval

`func (o *TablePartition) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *TablePartition) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *TablePartition) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetIntervalType

`func (o *TablePartition) GetIntervalType() string`

GetIntervalType returns the IntervalType field if non-nil, zero value otherwise.

### GetIntervalTypeOk

`func (o *TablePartition) GetIntervalTypeOk() (*string, bool)`

GetIntervalTypeOk returns a tuple with the IntervalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalType

`func (o *TablePartition) SetIntervalType(v string)`

SetIntervalType sets IntervalType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


