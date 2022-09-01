# CreateThread

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**About** | **string** |  | 
**AddressedTo** | Pointer to **string** |  | [optional] 
**From** | **string** |  | 
**Message** | **string** |  | 
**TaskDetails** | Pointer to [**CreateTaskDetails**](CreateTaskDetails.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateThread

`func NewCreateThread(about string, from string, message string, ) *CreateThread`

NewCreateThread instantiates a new CreateThread object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateThreadWithDefaults

`func NewCreateThreadWithDefaults() *CreateThread`

NewCreateThreadWithDefaults instantiates a new CreateThread object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbout

`func (o *CreateThread) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *CreateThread) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *CreateThread) SetAbout(v string)`

SetAbout sets About field to given value.


### GetAddressedTo

`func (o *CreateThread) GetAddressedTo() string`

GetAddressedTo returns the AddressedTo field if non-nil, zero value otherwise.

### GetAddressedToOk

`func (o *CreateThread) GetAddressedToOk() (*string, bool)`

GetAddressedToOk returns a tuple with the AddressedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressedTo

`func (o *CreateThread) SetAddressedTo(v string)`

SetAddressedTo sets AddressedTo field to given value.

### HasAddressedTo

`func (o *CreateThread) HasAddressedTo() bool`

HasAddressedTo returns a boolean if a field has been set.

### GetFrom

`func (o *CreateThread) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CreateThread) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CreateThread) SetFrom(v string)`

SetFrom sets From field to given value.


### GetMessage

`func (o *CreateThread) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateThread) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateThread) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTaskDetails

`func (o *CreateThread) GetTaskDetails() CreateTaskDetails`

GetTaskDetails returns the TaskDetails field if non-nil, zero value otherwise.

### GetTaskDetailsOk

`func (o *CreateThread) GetTaskDetailsOk() (*CreateTaskDetails, bool)`

GetTaskDetailsOk returns a tuple with the TaskDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDetails

`func (o *CreateThread) SetTaskDetails(v CreateTaskDetails)`

SetTaskDetails sets TaskDetails field to given value.

### HasTaskDetails

`func (o *CreateThread) HasTaskDetails() bool`

HasTaskDetails returns a boolean if a field has been set.

### GetType

`func (o *CreateThread) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateThread) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateThread) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateThread) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


