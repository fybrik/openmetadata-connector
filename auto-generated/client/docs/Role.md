# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeDescription** | Pointer to [**ChangeDescription**](ChangeDescription.md) |  | [optional] 
**DefaultRole** | Pointer to **bool** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **map[string]interface{}** |  | [optional] 
**Followers** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**FullyQualifiedName** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Policies** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**Tags** | Pointer to [**[]TagLabel**](TagLabel.md) |  | [optional] 
**Teams** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Users** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**Version** | Pointer to **float64** |  | [optional] 

## Methods

### NewRole

`func NewRole(id string, name string, ) *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeDescription

`func (o *Role) GetChangeDescription() ChangeDescription`

GetChangeDescription returns the ChangeDescription field if non-nil, zero value otherwise.

### GetChangeDescriptionOk

`func (o *Role) GetChangeDescriptionOk() (*ChangeDescription, bool)`

GetChangeDescriptionOk returns a tuple with the ChangeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDescription

`func (o *Role) SetChangeDescription(v ChangeDescription)`

SetChangeDescription sets ChangeDescription field to given value.

### HasChangeDescription

`func (o *Role) HasChangeDescription() bool`

HasChangeDescription returns a boolean if a field has been set.

### GetDefaultRole

`func (o *Role) GetDefaultRole() bool`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *Role) GetDefaultRoleOk() (*bool, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *Role) SetDefaultRole(v bool)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *Role) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetDeleted

`func (o *Role) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Role) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Role) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Role) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDescription

`func (o *Role) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Role) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Role) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Role) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *Role) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Role) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Role) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Role) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtension

`func (o *Role) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *Role) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *Role) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *Role) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetFollowers

`func (o *Role) GetFollowers() []EntityReference`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *Role) GetFollowersOk() (*[]EntityReference, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *Role) SetFollowers(v []EntityReference)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *Role) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetFullyQualifiedName

`func (o *Role) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *Role) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *Role) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.

### HasFullyQualifiedName

`func (o *Role) HasFullyQualifiedName() bool`

HasFullyQualifiedName returns a boolean if a field has been set.

### GetHref

`func (o *Role) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Role) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Role) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Role) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *Role) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Role) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Role) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Role) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Role) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Role) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *Role) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Role) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Role) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Role) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPolicies

`func (o *Role) GetPolicies() []EntityReference`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *Role) GetPoliciesOk() (*[]EntityReference, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *Role) SetPolicies(v []EntityReference)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *Role) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetTags

`func (o *Role) GetTags() []TagLabel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Role) GetTagsOk() (*[]TagLabel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Role) SetTags(v []TagLabel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Role) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTeams

`func (o *Role) GetTeams() []EntityReference`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *Role) GetTeamsOk() (*[]EntityReference, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *Role) SetTeams(v []EntityReference)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *Role) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Role) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Role) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Role) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Role) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Role) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Role) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Role) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Role) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUsers

`func (o *Role) GetUsers() []EntityReference`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Role) GetUsersOk() (*[]EntityReference, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Role) SetUsers(v []EntityReference)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Role) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetVersion

`func (o *Role) GetVersion() float64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Role) GetVersionOk() (*float64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Role) SetVersion(v float64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Role) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


