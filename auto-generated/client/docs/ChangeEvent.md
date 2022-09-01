# ChangeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeDescription** | Pointer to [**ChangeDescription**](ChangeDescription.md) |  | [optional] 
**CurrentVersion** | Pointer to **float64** |  | [optional] 
**Entity** | Pointer to **map[string]interface{}** |  | [optional] 
**EntityFullyQualifiedName** | Pointer to **string** |  | [optional] 
**EntityId** | **string** |  | 
**EntityType** | **string** |  | 
**EventType** | **string** |  | 
**PreviousVersion** | Pointer to **float64** |  | [optional] 
**Timestamp** | **int64** |  | 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewChangeEvent

`func NewChangeEvent(entityId string, entityType string, eventType string, timestamp int64, ) *ChangeEvent`

NewChangeEvent instantiates a new ChangeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeEventWithDefaults

`func NewChangeEventWithDefaults() *ChangeEvent`

NewChangeEventWithDefaults instantiates a new ChangeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeDescription

`func (o *ChangeEvent) GetChangeDescription() ChangeDescription`

GetChangeDescription returns the ChangeDescription field if non-nil, zero value otherwise.

### GetChangeDescriptionOk

`func (o *ChangeEvent) GetChangeDescriptionOk() (*ChangeDescription, bool)`

GetChangeDescriptionOk returns a tuple with the ChangeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDescription

`func (o *ChangeEvent) SetChangeDescription(v ChangeDescription)`

SetChangeDescription sets ChangeDescription field to given value.

### HasChangeDescription

`func (o *ChangeEvent) HasChangeDescription() bool`

HasChangeDescription returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *ChangeEvent) GetCurrentVersion() float64`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *ChangeEvent) GetCurrentVersionOk() (*float64, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *ChangeEvent) SetCurrentVersion(v float64)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *ChangeEvent) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetEntity

`func (o *ChangeEvent) GetEntity() map[string]interface{}`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ChangeEvent) GetEntityOk() (*map[string]interface{}, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *ChangeEvent) SetEntity(v map[string]interface{})`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *ChangeEvent) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetEntityFullyQualifiedName

`func (o *ChangeEvent) GetEntityFullyQualifiedName() string`

GetEntityFullyQualifiedName returns the EntityFullyQualifiedName field if non-nil, zero value otherwise.

### GetEntityFullyQualifiedNameOk

`func (o *ChangeEvent) GetEntityFullyQualifiedNameOk() (*string, bool)`

GetEntityFullyQualifiedNameOk returns a tuple with the EntityFullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityFullyQualifiedName

`func (o *ChangeEvent) SetEntityFullyQualifiedName(v string)`

SetEntityFullyQualifiedName sets EntityFullyQualifiedName field to given value.

### HasEntityFullyQualifiedName

`func (o *ChangeEvent) HasEntityFullyQualifiedName() bool`

HasEntityFullyQualifiedName returns a boolean if a field has been set.

### GetEntityId

`func (o *ChangeEvent) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ChangeEvent) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ChangeEvent) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *ChangeEvent) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ChangeEvent) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ChangeEvent) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetEventType

`func (o *ChangeEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ChangeEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ChangeEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetPreviousVersion

`func (o *ChangeEvent) GetPreviousVersion() float64`

GetPreviousVersion returns the PreviousVersion field if non-nil, zero value otherwise.

### GetPreviousVersionOk

`func (o *ChangeEvent) GetPreviousVersionOk() (*float64, bool)`

GetPreviousVersionOk returns a tuple with the PreviousVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersion

`func (o *ChangeEvent) SetPreviousVersion(v float64)`

SetPreviousVersion sets PreviousVersion field to given value.

### HasPreviousVersion

`func (o *ChangeEvent) HasPreviousVersion() bool`

HasPreviousVersion returns a boolean if a field has been set.

### GetTimestamp

`func (o *ChangeEvent) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ChangeEvent) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ChangeEvent) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetUserName

`func (o *ChangeEvent) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ChangeEvent) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ChangeEvent) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ChangeEvent) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


