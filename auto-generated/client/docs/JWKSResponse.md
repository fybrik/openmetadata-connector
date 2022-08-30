# JWKSResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | [**[]JWKSKey**](JWKSKey.md) |  | 

## Methods

### NewJWKSResponse

`func NewJWKSResponse(keys []JWKSKey, ) *JWKSResponse`

NewJWKSResponse instantiates a new JWKSResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJWKSResponseWithDefaults

`func NewJWKSResponseWithDefaults() *JWKSResponse`

NewJWKSResponseWithDefaults instantiates a new JWKSResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *JWKSResponse) GetKeys() []JWKSKey`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *JWKSResponse) GetKeysOk() (*[]JWKSKey, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *JWKSResponse) SetKeys(v []JWKSKey)`

SetKeys sets Keys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


