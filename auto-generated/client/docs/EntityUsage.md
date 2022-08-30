# EntityUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | [**EntityReference**](EntityReference.md) |  | 
**Usage** | [**[]UsageDetails**](UsageDetails.md) |  | 

## Methods

### NewEntityUsage

`func NewEntityUsage(entity EntityReference, usage []UsageDetails, ) *EntityUsage`

NewEntityUsage instantiates a new EntityUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityUsageWithDefaults

`func NewEntityUsageWithDefaults() *EntityUsage`

NewEntityUsageWithDefaults instantiates a new EntityUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *EntityUsage) GetEntity() EntityReference`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *EntityUsage) GetEntityOk() (*EntityReference, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *EntityUsage) SetEntity(v EntityReference)`

SetEntity sets Entity field to given value.


### GetUsage

`func (o *EntityUsage) GetUsage() []UsageDetails`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *EntityUsage) GetUsageOk() (*[]UsageDetails, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *EntityUsage) SetUsage(v []UsageDetails)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


