# MlFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FeatureAlgorithm** | Pointer to **string** |  | [optional] 
**FeatureSources** | Pointer to [**[]MlFeatureSource**](MlFeatureSource.md) |  | [optional] 
**FullyQualifiedName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagLabel**](TagLabel.md) |  | [optional] 

## Methods

### NewMlFeature

`func NewMlFeature() *MlFeature`

NewMlFeature instantiates a new MlFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlFeatureWithDefaults

`func NewMlFeatureWithDefaults() *MlFeature`

NewMlFeatureWithDefaults instantiates a new MlFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *MlFeature) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *MlFeature) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *MlFeature) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *MlFeature) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDescription

`func (o *MlFeature) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MlFeature) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MlFeature) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MlFeature) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatureAlgorithm

`func (o *MlFeature) GetFeatureAlgorithm() string`

GetFeatureAlgorithm returns the FeatureAlgorithm field if non-nil, zero value otherwise.

### GetFeatureAlgorithmOk

`func (o *MlFeature) GetFeatureAlgorithmOk() (*string, bool)`

GetFeatureAlgorithmOk returns a tuple with the FeatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureAlgorithm

`func (o *MlFeature) SetFeatureAlgorithm(v string)`

SetFeatureAlgorithm sets FeatureAlgorithm field to given value.

### HasFeatureAlgorithm

`func (o *MlFeature) HasFeatureAlgorithm() bool`

HasFeatureAlgorithm returns a boolean if a field has been set.

### GetFeatureSources

`func (o *MlFeature) GetFeatureSources() []MlFeatureSource`

GetFeatureSources returns the FeatureSources field if non-nil, zero value otherwise.

### GetFeatureSourcesOk

`func (o *MlFeature) GetFeatureSourcesOk() (*[]MlFeatureSource, bool)`

GetFeatureSourcesOk returns a tuple with the FeatureSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSources

`func (o *MlFeature) SetFeatureSources(v []MlFeatureSource)`

SetFeatureSources sets FeatureSources field to given value.

### HasFeatureSources

`func (o *MlFeature) HasFeatureSources() bool`

HasFeatureSources returns a boolean if a field has been set.

### GetFullyQualifiedName

`func (o *MlFeature) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *MlFeature) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *MlFeature) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.

### HasFullyQualifiedName

`func (o *MlFeature) HasFullyQualifiedName() bool`

HasFullyQualifiedName returns a boolean if a field has been set.

### GetName

`func (o *MlFeature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MlFeature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MlFeature) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MlFeature) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *MlFeature) GetTags() []TagLabel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MlFeature) GetTagsOk() (*[]TagLabel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MlFeature) SetTags(v []TagLabel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MlFeature) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


