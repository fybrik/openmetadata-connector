# MlFeatureSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSource** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**DataType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FullyQualifiedName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagLabel**](TagLabel.md) |  | [optional] 

## Methods

### NewMlFeatureSource

`func NewMlFeatureSource() *MlFeatureSource`

NewMlFeatureSource instantiates a new MlFeatureSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlFeatureSourceWithDefaults

`func NewMlFeatureSourceWithDefaults() *MlFeatureSource`

NewMlFeatureSourceWithDefaults instantiates a new MlFeatureSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSource

`func (o *MlFeatureSource) GetDataSource() EntityReference`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *MlFeatureSource) GetDataSourceOk() (*EntityReference, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *MlFeatureSource) SetDataSource(v EntityReference)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *MlFeatureSource) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetDataType

`func (o *MlFeatureSource) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *MlFeatureSource) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *MlFeatureSource) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *MlFeatureSource) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDescription

`func (o *MlFeatureSource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MlFeatureSource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MlFeatureSource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MlFeatureSource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFullyQualifiedName

`func (o *MlFeatureSource) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *MlFeatureSource) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *MlFeatureSource) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.

### HasFullyQualifiedName

`func (o *MlFeatureSource) HasFullyQualifiedName() bool`

HasFullyQualifiedName returns a boolean if a field has been set.

### GetName

`func (o *MlFeatureSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MlFeatureSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MlFeatureSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MlFeatureSource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *MlFeatureSource) GetTags() []TagLabel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MlFeatureSource) GetTagsOk() (*[]TagLabel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MlFeatureSource) SetTags(v []TagLabel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MlFeatureSource) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


