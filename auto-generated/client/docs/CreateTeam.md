# CreateTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultRoles** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **map[string]interface{}** |  | [optional] 
**IsJoinable** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Profile** | Pointer to [**Profile**](Profile.md) |  | [optional] 
**Users** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateTeam

`func NewCreateTeam(name string, ) *CreateTeam`

NewCreateTeam instantiates a new CreateTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTeamWithDefaults

`func NewCreateTeamWithDefaults() *CreateTeam`

NewCreateTeamWithDefaults instantiates a new CreateTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultRoles

`func (o *CreateTeam) GetDefaultRoles() []string`

GetDefaultRoles returns the DefaultRoles field if non-nil, zero value otherwise.

### GetDefaultRolesOk

`func (o *CreateTeam) GetDefaultRolesOk() (*[]string, bool)`

GetDefaultRolesOk returns a tuple with the DefaultRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoles

`func (o *CreateTeam) SetDefaultRoles(v []string)`

SetDefaultRoles sets DefaultRoles field to given value.

### HasDefaultRoles

`func (o *CreateTeam) HasDefaultRoles() bool`

HasDefaultRoles returns a boolean if a field has been set.

### GetDescription

`func (o *CreateTeam) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTeam) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTeam) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTeam) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateTeam) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateTeam) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateTeam) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateTeam) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtension

`func (o *CreateTeam) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *CreateTeam) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *CreateTeam) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *CreateTeam) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetIsJoinable

`func (o *CreateTeam) GetIsJoinable() bool`

GetIsJoinable returns the IsJoinable field if non-nil, zero value otherwise.

### GetIsJoinableOk

`func (o *CreateTeam) GetIsJoinableOk() (*bool, bool)`

GetIsJoinableOk returns a tuple with the IsJoinable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsJoinable

`func (o *CreateTeam) SetIsJoinable(v bool)`

SetIsJoinable sets IsJoinable field to given value.

### HasIsJoinable

`func (o *CreateTeam) HasIsJoinable() bool`

HasIsJoinable returns a boolean if a field has been set.

### GetName

`func (o *CreateTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTeam) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *CreateTeam) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateTeam) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateTeam) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateTeam) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProfile

`func (o *CreateTeam) GetProfile() Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CreateTeam) GetProfileOk() (*Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CreateTeam) SetProfile(v Profile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *CreateTeam) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetUsers

`func (o *CreateTeam) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *CreateTeam) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *CreateTeam) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *CreateTeam) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


