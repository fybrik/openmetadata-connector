# Explanation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**[]Explanation**](Explanation.md) |  | [optional] 
**Match** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 

## Methods

### NewExplanation

`func NewExplanation() *Explanation`

NewExplanation instantiates a new Explanation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExplanationWithDefaults

`func NewExplanationWithDefaults() *Explanation`

NewExplanationWithDefaults instantiates a new Explanation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Explanation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Explanation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Explanation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Explanation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetails

`func (o *Explanation) GetDetails() []Explanation`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Explanation) GetDetailsOk() (*[]Explanation, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Explanation) SetDetails(v []Explanation)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Explanation) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMatch

`func (o *Explanation) GetMatch() bool`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *Explanation) GetMatchOk() (*bool, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *Explanation) SetMatch(v bool)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *Explanation) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetValue

`func (o *Explanation) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Explanation) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Explanation) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *Explanation) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


