# CustomProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Name** | **string** |  | 
**PropertyType** | [**EntityReference**](EntityReference.md) |  | 

## Methods

### NewCustomProperty

`func NewCustomProperty(description string, name string, propertyType EntityReference, ) *CustomProperty`

NewCustomProperty instantiates a new CustomProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomPropertyWithDefaults

`func NewCustomPropertyWithDefaults() *CustomProperty`

NewCustomPropertyWithDefaults instantiates a new CustomProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CustomProperty) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomProperty) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomProperty) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *CustomProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomProperty) SetName(v string)`

SetName sets Name field to given value.


### GetPropertyType

`func (o *CustomProperty) GetPropertyType() EntityReference`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *CustomProperty) GetPropertyTypeOk() (*EntityReference, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *CustomProperty) SetPropertyType(v EntityReference)`

SetPropertyType sets PropertyType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


