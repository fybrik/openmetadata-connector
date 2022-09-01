# ResourceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connection** | [**Connection**](Connection.md) |  | 
**DataFormat** | Pointer to **string** |  | [optional] 

## Methods

### NewResourceDetails

`func NewResourceDetails(connection Connection, ) *ResourceDetails`

NewResourceDetails instantiates a new ResourceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceDetailsWithDefaults

`func NewResourceDetailsWithDefaults() *ResourceDetails`

NewResourceDetailsWithDefaults instantiates a new ResourceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnection

`func (o *ResourceDetails) GetConnection() Connection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ResourceDetails) GetConnectionOk() (*Connection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ResourceDetails) SetConnection(v Connection)`

SetConnection sets Connection field to given value.


### GetDataFormat

`func (o *ResourceDetails) GetDataFormat() string`

GetDataFormat returns the DataFormat field if non-nil, zero value otherwise.

### GetDataFormatOk

`func (o *ResourceDetails) GetDataFormatOk() (*string, bool)`

GetDataFormatOk returns a tuple with the DataFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFormat

`func (o *ResourceDetails) SetDataFormat(v string)`

SetDataFormat sets DataFormat field to given value.

### HasDataFormat

`func (o *ResourceDetails) HasDataFormat() bool`

HasDataFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


