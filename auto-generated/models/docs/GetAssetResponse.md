# GetAssetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | **string** | Vault plugin path where the data credentials will be stored as kubernetes secrets This value is assumed to be known to the catalog connector. | 
**Details** | [**ResourceDetails**](ResourceDetails.md) |  | 
**ResourceMetadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 

## Methods

### NewGetAssetResponse

`func NewGetAssetResponse(credentials string, details ResourceDetails, resourceMetadata ResourceMetadata, ) *GetAssetResponse`

NewGetAssetResponse instantiates a new GetAssetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetResponseWithDefaults

`func NewGetAssetResponseWithDefaults() *GetAssetResponse`

NewGetAssetResponseWithDefaults instantiates a new GetAssetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *GetAssetResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GetAssetResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GetAssetResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetDetails

`func (o *GetAssetResponse) GetDetails() ResourceDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetAssetResponse) GetDetailsOk() (*ResourceDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetAssetResponse) SetDetails(v ResourceDetails)`

SetDetails sets Details field to given value.


### GetResourceMetadata

`func (o *GetAssetResponse) GetResourceMetadata() ResourceMetadata`

GetResourceMetadata returns the ResourceMetadata field if non-nil, zero value otherwise.

### GetResourceMetadataOk

`func (o *GetAssetResponse) GetResourceMetadataOk() (*ResourceMetadata, bool)`

GetResourceMetadataOk returns a tuple with the ResourceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceMetadata

`func (o *GetAssetResponse) SetResourceMetadata(v ResourceMetadata)`

SetResourceMetadata sets ResourceMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


