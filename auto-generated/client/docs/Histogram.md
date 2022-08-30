# Histogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boundaries** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Frequencies** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewHistogram

`func NewHistogram() *Histogram`

NewHistogram instantiates a new Histogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistogramWithDefaults

`func NewHistogramWithDefaults() *Histogram`

NewHistogramWithDefaults instantiates a new Histogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoundaries

`func (o *Histogram) GetBoundaries() []map[string]interface{}`

GetBoundaries returns the Boundaries field if non-nil, zero value otherwise.

### GetBoundariesOk

`func (o *Histogram) GetBoundariesOk() (*[]map[string]interface{}, bool)`

GetBoundariesOk returns a tuple with the Boundaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundaries

`func (o *Histogram) SetBoundaries(v []map[string]interface{})`

SetBoundaries sets Boundaries field to given value.

### HasBoundaries

`func (o *Histogram) HasBoundaries() bool`

HasBoundaries returns a boolean if a field has been set.

### GetFrequencies

`func (o *Histogram) GetFrequencies() []map[string]interface{}`

GetFrequencies returns the Frequencies field if non-nil, zero value otherwise.

### GetFrequenciesOk

`func (o *Histogram) GetFrequenciesOk() (*[]map[string]interface{}, bool)`

GetFrequenciesOk returns a tuple with the Frequencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencies

`func (o *Histogram) SetFrequencies(v []map[string]interface{})`

SetFrequencies sets Frequencies field to given value.

### HasFrequencies

`func (o *Histogram) HasFrequencies() bool`

HasFrequencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


