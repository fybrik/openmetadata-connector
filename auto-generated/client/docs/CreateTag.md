# CreateTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedTags** | Pointer to **[]string** |  | [optional] 
**Description** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 

## Methods

### NewCreateTag

`func NewCreateTag(description string, name string, ) *CreateTag`

NewCreateTag instantiates a new CreateTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTagWithDefaults

`func NewCreateTagWithDefaults() *CreateTag`

NewCreateTagWithDefaults instantiates a new CreateTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedTags

`func (o *CreateTag) GetAssociatedTags() []string`

GetAssociatedTags returns the AssociatedTags field if non-nil, zero value otherwise.

### GetAssociatedTagsOk

`func (o *CreateTag) GetAssociatedTagsOk() (*[]string, bool)`

GetAssociatedTagsOk returns a tuple with the AssociatedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedTags

`func (o *CreateTag) SetAssociatedTags(v []string)`

SetAssociatedTags sets AssociatedTags field to given value.

### HasAssociatedTags

`func (o *CreateTag) HasAssociatedTags() bool`

HasAssociatedTags returns a boolean if a field has been set.

### GetDescription

`func (o *CreateTag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTag) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayName

`func (o *CreateTag) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateTag) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateTag) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateTag) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtension

`func (o *CreateTag) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *CreateTag) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *CreateTag) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *CreateTag) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetName

`func (o *CreateTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTag) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *CreateTag) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateTag) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateTag) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateTag) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


