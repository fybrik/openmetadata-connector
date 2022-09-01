# CreateAssetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to **string** | The vault plugin path where the destination data credentials will be stored as kubernetes secrets | [optional] 
**DestinationAssetID** | Pointer to **string** | Asset ID to be used for the created asset | [optional] 
**DestinationCatalogID** | **string** | The destination catalog id in which the new asset will be created based on the information provided in ResourceMetadata and ResourceDetails field | 
**Details** | [**ResourceDetails**](ResourceDetails.md) |  | 
**ResourceMetadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 

## Methods

### NewCreateAssetRequest

`func NewCreateAssetRequest(destinationCatalogID string, details ResourceDetails, resourceMetadata ResourceMetadata, ) *CreateAssetRequest`

NewCreateAssetRequest instantiates a new CreateAssetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAssetRequestWithDefaults

`func NewCreateAssetRequestWithDefaults() *CreateAssetRequest`

NewCreateAssetRequestWithDefaults instantiates a new CreateAssetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *CreateAssetRequest) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CreateAssetRequest) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CreateAssetRequest) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CreateAssetRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetDestinationAssetID

`func (o *CreateAssetRequest) GetDestinationAssetID() string`

GetDestinationAssetID returns the DestinationAssetID field if non-nil, zero value otherwise.

### GetDestinationAssetIDOk

`func (o *CreateAssetRequest) GetDestinationAssetIDOk() (*string, bool)`

GetDestinationAssetIDOk returns a tuple with the DestinationAssetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAssetID

`func (o *CreateAssetRequest) SetDestinationAssetID(v string)`

SetDestinationAssetID sets DestinationAssetID field to given value.

### HasDestinationAssetID

`func (o *CreateAssetRequest) HasDestinationAssetID() bool`

HasDestinationAssetID returns a boolean if a field has been set.

### GetDestinationCatalogID

`func (o *CreateAssetRequest) GetDestinationCatalogID() string`

GetDestinationCatalogID returns the DestinationCatalogID field if non-nil, zero value otherwise.

### GetDestinationCatalogIDOk

`func (o *CreateAssetRequest) GetDestinationCatalogIDOk() (*string, bool)`

GetDestinationCatalogIDOk returns a tuple with the DestinationCatalogID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCatalogID

`func (o *CreateAssetRequest) SetDestinationCatalogID(v string)`

SetDestinationCatalogID sets DestinationCatalogID field to given value.


### GetDetails

`func (o *CreateAssetRequest) GetDetails() ResourceDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CreateAssetRequest) GetDetailsOk() (*ResourceDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CreateAssetRequest) SetDetails(v ResourceDetails)`

SetDetails sets Details field to given value.


### GetResourceMetadata

`func (o *CreateAssetRequest) GetResourceMetadata() ResourceMetadata`

GetResourceMetadata returns the ResourceMetadata field if non-nil, zero value otherwise.

### GetResourceMetadataOk

`func (o *CreateAssetRequest) GetResourceMetadataOk() (*ResourceMetadata, bool)`

GetResourceMetadataOk returns a tuple with the ResourceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceMetadata

`func (o *CreateAssetRequest) SetResourceMetadata(v ResourceMetadata)`

SetResourceMetadata sets ResourceMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


