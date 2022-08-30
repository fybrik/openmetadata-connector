# Reaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReactionType** | **string** |  | 
**User** | [**EntityReference**](EntityReference.md) |  | 

## Methods

### NewReaction

`func NewReaction(reactionType string, user EntityReference, ) *Reaction`

NewReaction instantiates a new Reaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionWithDefaults

`func NewReactionWithDefaults() *Reaction`

NewReactionWithDefaults instantiates a new Reaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReactionType

`func (o *Reaction) GetReactionType() string`

GetReactionType returns the ReactionType field if non-nil, zero value otherwise.

### GetReactionTypeOk

`func (o *Reaction) GetReactionTypeOk() (*string, bool)`

GetReactionTypeOk returns a tuple with the ReactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionType

`func (o *Reaction) SetReactionType(v string)`

SetReactionType sets ReactionType field to given value.


### GetUser

`func (o *Reaction) GetUser() EntityReference`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Reaction) GetUserOk() (*EntityReference, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Reaction) SetUser(v EntityReference)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


