# Thread

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**About** | **string** |  | 
**AddressedTo** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Message** | **string** |  | 
**Posts** | Pointer to [**[]Post**](Post.md) |  | [optional] 
**PostsCount** | Pointer to **int32** |  | [optional] 
**Reactions** | Pointer to [**[]Reaction**](Reaction.md) |  | [optional] 
**Resolved** | Pointer to **bool** |  | [optional] 
**Task** | Pointer to [**TaskDetails**](TaskDetails.md) |  | [optional] 
**ThreadTs** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewThread

`func NewThread(about string, id string, message string, ) *Thread`

NewThread instantiates a new Thread object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadWithDefaults

`func NewThreadWithDefaults() *Thread`

NewThreadWithDefaults instantiates a new Thread object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbout

`func (o *Thread) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *Thread) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *Thread) SetAbout(v string)`

SetAbout sets About field to given value.


### GetAddressedTo

`func (o *Thread) GetAddressedTo() string`

GetAddressedTo returns the AddressedTo field if non-nil, zero value otherwise.

### GetAddressedToOk

`func (o *Thread) GetAddressedToOk() (*string, bool)`

GetAddressedToOk returns a tuple with the AddressedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressedTo

`func (o *Thread) SetAddressedTo(v string)`

SetAddressedTo sets AddressedTo field to given value.

### HasAddressedTo

`func (o *Thread) HasAddressedTo() bool`

HasAddressedTo returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Thread) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Thread) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Thread) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Thread) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEntityId

`func (o *Thread) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *Thread) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *Thread) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *Thread) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetHref

`func (o *Thread) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Thread) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Thread) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Thread) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *Thread) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Thread) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Thread) SetId(v string)`

SetId sets Id field to given value.


### GetMessage

`func (o *Thread) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Thread) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Thread) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPosts

`func (o *Thread) GetPosts() []Post`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *Thread) GetPostsOk() (*[]Post, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosts

`func (o *Thread) SetPosts(v []Post)`

SetPosts sets Posts field to given value.

### HasPosts

`func (o *Thread) HasPosts() bool`

HasPosts returns a boolean if a field has been set.

### GetPostsCount

`func (o *Thread) GetPostsCount() int32`

GetPostsCount returns the PostsCount field if non-nil, zero value otherwise.

### GetPostsCountOk

`func (o *Thread) GetPostsCountOk() (*int32, bool)`

GetPostsCountOk returns a tuple with the PostsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostsCount

`func (o *Thread) SetPostsCount(v int32)`

SetPostsCount sets PostsCount field to given value.

### HasPostsCount

`func (o *Thread) HasPostsCount() bool`

HasPostsCount returns a boolean if a field has been set.

### GetReactions

`func (o *Thread) GetReactions() []Reaction`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *Thread) GetReactionsOk() (*[]Reaction, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *Thread) SetReactions(v []Reaction)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *Thread) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetResolved

`func (o *Thread) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *Thread) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *Thread) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *Thread) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetTask

`func (o *Thread) GetTask() TaskDetails`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *Thread) GetTaskOk() (*TaskDetails, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *Thread) SetTask(v TaskDetails)`

SetTask sets Task field to given value.

### HasTask

`func (o *Thread) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetThreadTs

`func (o *Thread) GetThreadTs() int64`

GetThreadTs returns the ThreadTs field if non-nil, zero value otherwise.

### GetThreadTsOk

`func (o *Thread) GetThreadTsOk() (*int64, bool)`

GetThreadTsOk returns a tuple with the ThreadTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadTs

`func (o *Thread) SetThreadTs(v int64)`

SetThreadTs sets ThreadTs field to given value.

### HasThreadTs

`func (o *Thread) HasThreadTs() bool`

HasThreadTs returns a boolean if a field has been set.

### GetType

`func (o *Thread) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Thread) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Thread) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Thread) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Thread) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Thread) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Thread) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Thread) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Thread) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Thread) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Thread) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Thread) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


