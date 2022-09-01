# ResourceColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the column | 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResourceColumn

`func NewResourceColumn(name string, ) *ResourceColumn`

NewResourceColumn instantiates a new ResourceColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceColumnWithDefaults

`func NewResourceColumnWithDefaults() *ResourceColumn`

NewResourceColumnWithDefaults instantiates a new ResourceColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResourceColumn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceColumn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceColumn) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *ResourceColumn) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResourceColumn) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResourceColumn) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ResourceColumn) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


