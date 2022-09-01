# JWTAuthMechanism

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JWTToken** | **string** |  | 
**JWTTokenExpiresAt** | Pointer to **int64** |  | [optional] 
**JWTTokenExpiry** | **string** |  | 

## Methods

### NewJWTAuthMechanism

`func NewJWTAuthMechanism(jWTToken string, jWTTokenExpiry string, ) *JWTAuthMechanism`

NewJWTAuthMechanism instantiates a new JWTAuthMechanism object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJWTAuthMechanismWithDefaults

`func NewJWTAuthMechanismWithDefaults() *JWTAuthMechanism`

NewJWTAuthMechanismWithDefaults instantiates a new JWTAuthMechanism object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJWTToken

`func (o *JWTAuthMechanism) GetJWTToken() string`

GetJWTToken returns the JWTToken field if non-nil, zero value otherwise.

### GetJWTTokenOk

`func (o *JWTAuthMechanism) GetJWTTokenOk() (*string, bool)`

GetJWTTokenOk returns a tuple with the JWTToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJWTToken

`func (o *JWTAuthMechanism) SetJWTToken(v string)`

SetJWTToken sets JWTToken field to given value.


### GetJWTTokenExpiresAt

`func (o *JWTAuthMechanism) GetJWTTokenExpiresAt() int64`

GetJWTTokenExpiresAt returns the JWTTokenExpiresAt field if non-nil, zero value otherwise.

### GetJWTTokenExpiresAtOk

`func (o *JWTAuthMechanism) GetJWTTokenExpiresAtOk() (*int64, bool)`

GetJWTTokenExpiresAtOk returns a tuple with the JWTTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJWTTokenExpiresAt

`func (o *JWTAuthMechanism) SetJWTTokenExpiresAt(v int64)`

SetJWTTokenExpiresAt sets JWTTokenExpiresAt field to given value.

### HasJWTTokenExpiresAt

`func (o *JWTAuthMechanism) HasJWTTokenExpiresAt() bool`

HasJWTTokenExpiresAt returns a boolean if a field has been set.

### GetJWTTokenExpiry

`func (o *JWTAuthMechanism) GetJWTTokenExpiry() string`

GetJWTTokenExpiry returns the JWTTokenExpiry field if non-nil, zero value otherwise.

### GetJWTTokenExpiryOk

`func (o *JWTAuthMechanism) GetJWTTokenExpiryOk() (*string, bool)`

GetJWTTokenExpiryOk returns a tuple with the JWTTokenExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJWTTokenExpiry

`func (o *JWTAuthMechanism) SetJWTTokenExpiry(v string)`

SetJWTTokenExpiry sets JWTTokenExpiry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


