# AuthorizerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminPrincipals** | **[]string** |  | 
**BotPrincipals** | **[]string** |  | 
**ClassName** | **string** |  | 
**ContainerRequestFilter** | **string** |  | 
**EnableSecureSocketConnection** | **bool** |  | 
**EnforcePrincipalDomain** | **bool** |  | 
**PrincipalDomain** | **string** |  | 

## Methods

### NewAuthorizerConfiguration

`func NewAuthorizerConfiguration(adminPrincipals []string, botPrincipals []string, className string, containerRequestFilter string, enableSecureSocketConnection bool, enforcePrincipalDomain bool, principalDomain string, ) *AuthorizerConfiguration`

NewAuthorizerConfiguration instantiates a new AuthorizerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizerConfigurationWithDefaults

`func NewAuthorizerConfigurationWithDefaults() *AuthorizerConfiguration`

NewAuthorizerConfigurationWithDefaults instantiates a new AuthorizerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminPrincipals

`func (o *AuthorizerConfiguration) GetAdminPrincipals() []string`

GetAdminPrincipals returns the AdminPrincipals field if non-nil, zero value otherwise.

### GetAdminPrincipalsOk

`func (o *AuthorizerConfiguration) GetAdminPrincipalsOk() (*[]string, bool)`

GetAdminPrincipalsOk returns a tuple with the AdminPrincipals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPrincipals

`func (o *AuthorizerConfiguration) SetAdminPrincipals(v []string)`

SetAdminPrincipals sets AdminPrincipals field to given value.


### GetBotPrincipals

`func (o *AuthorizerConfiguration) GetBotPrincipals() []string`

GetBotPrincipals returns the BotPrincipals field if non-nil, zero value otherwise.

### GetBotPrincipalsOk

`func (o *AuthorizerConfiguration) GetBotPrincipalsOk() (*[]string, bool)`

GetBotPrincipalsOk returns a tuple with the BotPrincipals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotPrincipals

`func (o *AuthorizerConfiguration) SetBotPrincipals(v []string)`

SetBotPrincipals sets BotPrincipals field to given value.


### GetClassName

`func (o *AuthorizerConfiguration) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *AuthorizerConfiguration) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *AuthorizerConfiguration) SetClassName(v string)`

SetClassName sets ClassName field to given value.


### GetContainerRequestFilter

`func (o *AuthorizerConfiguration) GetContainerRequestFilter() string`

GetContainerRequestFilter returns the ContainerRequestFilter field if non-nil, zero value otherwise.

### GetContainerRequestFilterOk

`func (o *AuthorizerConfiguration) GetContainerRequestFilterOk() (*string, bool)`

GetContainerRequestFilterOk returns a tuple with the ContainerRequestFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerRequestFilter

`func (o *AuthorizerConfiguration) SetContainerRequestFilter(v string)`

SetContainerRequestFilter sets ContainerRequestFilter field to given value.


### GetEnableSecureSocketConnection

`func (o *AuthorizerConfiguration) GetEnableSecureSocketConnection() bool`

GetEnableSecureSocketConnection returns the EnableSecureSocketConnection field if non-nil, zero value otherwise.

### GetEnableSecureSocketConnectionOk

`func (o *AuthorizerConfiguration) GetEnableSecureSocketConnectionOk() (*bool, bool)`

GetEnableSecureSocketConnectionOk returns a tuple with the EnableSecureSocketConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSecureSocketConnection

`func (o *AuthorizerConfiguration) SetEnableSecureSocketConnection(v bool)`

SetEnableSecureSocketConnection sets EnableSecureSocketConnection field to given value.


### GetEnforcePrincipalDomain

`func (o *AuthorizerConfiguration) GetEnforcePrincipalDomain() bool`

GetEnforcePrincipalDomain returns the EnforcePrincipalDomain field if non-nil, zero value otherwise.

### GetEnforcePrincipalDomainOk

`func (o *AuthorizerConfiguration) GetEnforcePrincipalDomainOk() (*bool, bool)`

GetEnforcePrincipalDomainOk returns a tuple with the EnforcePrincipalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcePrincipalDomain

`func (o *AuthorizerConfiguration) SetEnforcePrincipalDomain(v bool)`

SetEnforcePrincipalDomain sets EnforcePrincipalDomain field to given value.


### GetPrincipalDomain

`func (o *AuthorizerConfiguration) GetPrincipalDomain() string`

GetPrincipalDomain returns the PrincipalDomain field if non-nil, zero value otherwise.

### GetPrincipalDomainOk

`func (o *AuthorizerConfiguration) GetPrincipalDomainOk() (*string, bool)`

GetPrincipalDomainOk returns a tuple with the PrincipalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalDomain

`func (o *AuthorizerConfiguration) SetPrincipalDomain(v string)`

SetPrincipalDomain sets PrincipalDomain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


