# Post

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** |  | 
**Id** | **string** |  | 
**Message** | **string** |  | 
**PostTs** | Pointer to **int64** |  | [optional] 
**Reactions** | Pointer to [**[]Reaction**](Reaction.md) |  | [optional] 

## Methods

### NewPost

`func NewPost(from string, id string, message string, ) *Post`

NewPost instantiates a new Post object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostWithDefaults

`func NewPostWithDefaults() *Post`

NewPostWithDefaults instantiates a new Post object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *Post) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Post) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Post) SetFrom(v string)`

SetFrom sets From field to given value.


### GetId

`func (o *Post) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Post) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Post) SetId(v string)`

SetId sets Id field to given value.


### GetMessage

`func (o *Post) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Post) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Post) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPostTs

`func (o *Post) GetPostTs() int64`

GetPostTs returns the PostTs field if non-nil, zero value otherwise.

### GetPostTsOk

`func (o *Post) GetPostTsOk() (*int64, bool)`

GetPostTsOk returns a tuple with the PostTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostTs

`func (o *Post) SetPostTs(v int64)`

SetPostTs sets PostTs field to given value.

### HasPostTs

`func (o *Post) HasPostTs() bool`

HasPostTs returns a boolean if a field has been set.

### GetReactions

`func (o *Post) GetReactions() []Reaction`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *Post) GetReactionsOk() (*[]Reaction, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *Post) SetReactions(v []Reaction)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *Post) HasReactions() bool`

HasReactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


