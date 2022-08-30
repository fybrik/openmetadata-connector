# MlStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageRepository** | Pointer to **string** |  | [optional] 
**Storage** | Pointer to **string** |  | [optional] 

## Methods

### NewMlStore

`func NewMlStore() *MlStore`

NewMlStore instantiates a new MlStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlStoreWithDefaults

`func NewMlStoreWithDefaults() *MlStore`

NewMlStoreWithDefaults instantiates a new MlStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageRepository

`func (o *MlStore) GetImageRepository() string`

GetImageRepository returns the ImageRepository field if non-nil, zero value otherwise.

### GetImageRepositoryOk

`func (o *MlStore) GetImageRepositoryOk() (*string, bool)`

GetImageRepositoryOk returns a tuple with the ImageRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRepository

`func (o *MlStore) SetImageRepository(v string)`

SetImageRepository sets ImageRepository field to given value.

### HasImageRepository

`func (o *MlStore) HasImageRepository() bool`

HasImageRepository returns a boolean if a field has been set.

### GetStorage

`func (o *MlStore) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *MlStore) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *MlStore) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *MlStore) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


