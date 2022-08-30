# EntityHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | **string** |  | 
**Versions** | **[]map[string]interface{}** |  | 

## Methods

### NewEntityHistory

`func NewEntityHistory(entityType string, versions []map[string]interface{}, ) *EntityHistory`

NewEntityHistory instantiates a new EntityHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityHistoryWithDefaults

`func NewEntityHistoryWithDefaults() *EntityHistory`

NewEntityHistoryWithDefaults instantiates a new EntityHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *EntityHistory) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *EntityHistory) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *EntityHistory) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetVersions

`func (o *EntityHistory) GetVersions() []map[string]interface{}`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *EntityHistory) GetVersionsOk() (*[]map[string]interface{}, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *EntityHistory) SetVersions(v []map[string]interface{})`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


