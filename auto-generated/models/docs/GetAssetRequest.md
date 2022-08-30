# GetAssetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetID** | **string** |  | 
**OperationType** | [**OperationType**](OperationType.md) |  | 

## Methods

### NewGetAssetRequest

`func NewGetAssetRequest(assetID string, operationType OperationType, ) *GetAssetRequest`

NewGetAssetRequest instantiates a new GetAssetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetRequestWithDefaults

`func NewGetAssetRequestWithDefaults() *GetAssetRequest`

NewGetAssetRequestWithDefaults instantiates a new GetAssetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetID

`func (o *GetAssetRequest) GetAssetID() string`

GetAssetID returns the AssetID field if non-nil, zero value otherwise.

### GetAssetIDOk

`func (o *GetAssetRequest) GetAssetIDOk() (*string, bool)`

GetAssetIDOk returns a tuple with the AssetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetID

`func (o *GetAssetRequest) SetAssetID(v string)`

SetAssetID sets AssetID field to given value.


### GetOperationType

`func (o *GetAssetRequest) GetOperationType() OperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *GetAssetRequest) GetOperationTypeOk() (*OperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *GetAssetRequest) SetOperationType(v OperationType)`

SetOperationType sets OperationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


