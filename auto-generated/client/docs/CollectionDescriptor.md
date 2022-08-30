# CollectionDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to [**CollectionInfo**](CollectionInfo.md) |  | [optional] 

## Methods

### NewCollectionDescriptor

`func NewCollectionDescriptor() *CollectionDescriptor`

NewCollectionDescriptor instantiates a new CollectionDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionDescriptorWithDefaults

`func NewCollectionDescriptorWithDefaults() *CollectionDescriptor`

NewCollectionDescriptorWithDefaults instantiates a new CollectionDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *CollectionDescriptor) GetCollection() CollectionInfo`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *CollectionDescriptor) GetCollectionOk() (*CollectionInfo, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *CollectionDescriptor) SetCollection(v CollectionInfo)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *CollectionDescriptor) HasCollection() bool`

HasCollection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


