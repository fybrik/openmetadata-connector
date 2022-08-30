# CreateDashboardService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connection** | [**DashboardConnection**](DashboardConnection.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**ServiceType** | **string** |  | 

## Methods

### NewCreateDashboardService

`func NewCreateDashboardService(connection DashboardConnection, name string, serviceType string, ) *CreateDashboardService`

NewCreateDashboardService instantiates a new CreateDashboardService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDashboardServiceWithDefaults

`func NewCreateDashboardServiceWithDefaults() *CreateDashboardService`

NewCreateDashboardServiceWithDefaults instantiates a new CreateDashboardService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnection

`func (o *CreateDashboardService) GetConnection() DashboardConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *CreateDashboardService) GetConnectionOk() (*DashboardConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *CreateDashboardService) SetConnection(v DashboardConnection)`

SetConnection sets Connection field to given value.


### GetDescription

`func (o *CreateDashboardService) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDashboardService) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDashboardService) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDashboardService) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateDashboardService) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateDashboardService) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateDashboardService) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateDashboardService) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtension

`func (o *CreateDashboardService) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *CreateDashboardService) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *CreateDashboardService) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *CreateDashboardService) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetName

`func (o *CreateDashboardService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDashboardService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDashboardService) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *CreateDashboardService) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateDashboardService) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateDashboardService) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateDashboardService) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetServiceType

`func (o *CreateDashboardService) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CreateDashboardService) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CreateDashboardService) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


