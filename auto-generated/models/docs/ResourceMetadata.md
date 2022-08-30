# ResourceMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to [**[]ResourceColumn**](ResourceColumn.md) | Columns associated with the asset | [optional] 
**Geography** | Pointer to **string** | Geography of the resource | [optional] 
**Name** | Pointer to **string** | Name of the resource | [optional] 
**Owner** | Pointer to **string** | Owner of the resource | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResourceMetadata

`func NewResourceMetadata() *ResourceMetadata`

NewResourceMetadata instantiates a new ResourceMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceMetadataWithDefaults

`func NewResourceMetadataWithDefaults() *ResourceMetadata`

NewResourceMetadataWithDefaults instantiates a new ResourceMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *ResourceMetadata) GetColumns() []ResourceColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *ResourceMetadata) GetColumnsOk() (*[]ResourceColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *ResourceMetadata) SetColumns(v []ResourceColumn)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *ResourceMetadata) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetGeography

`func (o *ResourceMetadata) GetGeography() string`

GetGeography returns the Geography field if non-nil, zero value otherwise.

### GetGeographyOk

`func (o *ResourceMetadata) GetGeographyOk() (*string, bool)`

GetGeographyOk returns a tuple with the Geography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeography

`func (o *ResourceMetadata) SetGeography(v string)`

SetGeography sets Geography field to given value.

### HasGeography

`func (o *ResourceMetadata) HasGeography() bool`

HasGeography returns a boolean if a field has been set.

### GetName

`func (o *ResourceMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *ResourceMetadata) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ResourceMetadata) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ResourceMetadata) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ResourceMetadata) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTags

`func (o *ResourceMetadata) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResourceMetadata) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResourceMetadata) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ResourceMetadata) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


