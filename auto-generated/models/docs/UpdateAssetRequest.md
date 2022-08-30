# UpdateAssetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetID** | **string** |  | 
**Columns** | Pointer to [**[]ResourceColumn**](ResourceColumn.md) | New columns associated with the asset | [optional] 
**Name** | Pointer to **string** | New name of the resource | [optional] 
**Owner** | Pointer to **string** | New owner of the resource | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateAssetRequest

`func NewUpdateAssetRequest(assetID string, ) *UpdateAssetRequest`

NewUpdateAssetRequest instantiates a new UpdateAssetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAssetRequestWithDefaults

`func NewUpdateAssetRequestWithDefaults() *UpdateAssetRequest`

NewUpdateAssetRequestWithDefaults instantiates a new UpdateAssetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetID

`func (o *UpdateAssetRequest) GetAssetID() string`

GetAssetID returns the AssetID field if non-nil, zero value otherwise.

### GetAssetIDOk

`func (o *UpdateAssetRequest) GetAssetIDOk() (*string, bool)`

GetAssetIDOk returns a tuple with the AssetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetID

`func (o *UpdateAssetRequest) SetAssetID(v string)`

SetAssetID sets AssetID field to given value.


### GetColumns

`func (o *UpdateAssetRequest) GetColumns() []ResourceColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *UpdateAssetRequest) GetColumnsOk() (*[]ResourceColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *UpdateAssetRequest) SetColumns(v []ResourceColumn)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *UpdateAssetRequest) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetName

`func (o *UpdateAssetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAssetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAssetRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAssetRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateAssetRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateAssetRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateAssetRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdateAssetRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTags

`func (o *UpdateAssetRequest) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateAssetRequest) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateAssetRequest) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateAssetRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


