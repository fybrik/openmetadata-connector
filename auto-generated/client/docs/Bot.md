# Bot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BotUser** | [**EntityReference**](EntityReference.md) |  | 
**ChangeDescription** | Pointer to [**ChangeDescription**](ChangeDescription.md) |  | [optional] 
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
**Tags** | Pointer to [**[]TagLabel**](TagLabel.md) |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **float64** |  | [optional] 

## Methods

### NewBot

`func NewBot(botUser EntityReference, id string, name string, ) *Bot`

NewBot instantiates a new Bot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotWithDefaults

`func NewBotWithDefaults() *Bot`

NewBotWithDefaults instantiates a new Bot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBotUser

`func (o *Bot) GetBotUser() EntityReference`

GetBotUser returns the BotUser field if non-nil, zero value otherwise.

### GetBotUserOk

`func (o *Bot) GetBotUserOk() (*EntityReference, bool)`

GetBotUserOk returns a tuple with the BotUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotUser

`func (o *Bot) SetBotUser(v EntityReference)`

SetBotUser sets BotUser field to given value.


### GetChangeDescription

`func (o *Bot) GetChangeDescription() ChangeDescription`

GetChangeDescription returns the ChangeDescription field if non-nil, zero value otherwise.

### GetChangeDescriptionOk

`func (o *Bot) GetChangeDescriptionOk() (*ChangeDescription, bool)`

GetChangeDescriptionOk returns a tuple with the ChangeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDescription

`func (o *Bot) SetChangeDescription(v ChangeDescription)`

SetChangeDescription sets ChangeDescription field to given value.

### HasChangeDescription

`func (o *Bot) HasChangeDescription() bool`

HasChangeDescription returns a boolean if a field has been set.

### GetDeleted

`func (o *Bot) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Bot) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Bot) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Bot) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDescription

`func (o *Bot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Bot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Bot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Bot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *Bot) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Bot) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Bot) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Bot) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtension

`func (o *Bot) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *Bot) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *Bot) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *Bot) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetFollowers

`func (o *Bot) GetFollowers() []EntityReference`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *Bot) GetFollowersOk() (*[]EntityReference, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *Bot) SetFollowers(v []EntityReference)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *Bot) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetFullyQualifiedName

`func (o *Bot) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *Bot) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *Bot) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.

### HasFullyQualifiedName

`func (o *Bot) HasFullyQualifiedName() bool`

HasFullyQualifiedName returns a boolean if a field has been set.

### GetHref

`func (o *Bot) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Bot) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Bot) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Bot) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *Bot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bot) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Bot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bot) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *Bot) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Bot) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Bot) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Bot) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTags

`func (o *Bot) GetTags() []TagLabel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Bot) GetTagsOk() (*[]TagLabel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Bot) SetTags(v []TagLabel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Bot) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Bot) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Bot) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Bot) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Bot) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Bot) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Bot) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Bot) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Bot) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetVersion

`func (o *Bot) GetVersion() float64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Bot) GetVersionOk() (*float64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Bot) SetVersion(v float64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Bot) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


