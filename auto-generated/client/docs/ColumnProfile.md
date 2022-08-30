# ColumnProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomMetricsProfile** | Pointer to [**[]CustomMetricProfile**](CustomMetricProfile.md) |  | [optional] 
**DistinctCount** | Pointer to **float64** |  | [optional] 
**DistinctProportion** | Pointer to **float64** |  | [optional] 
**DuplicateCount** | Pointer to **float64** |  | [optional] 
**Histogram** | Pointer to [**Histogram**](Histogram.md) |  | [optional] 
**Max** | Pointer to **float64** |  | [optional] 
**MaxLength** | Pointer to **float64** |  | [optional] 
**Mean** | Pointer to **float64** |  | [optional] 
**Min** | Pointer to **float64** |  | [optional] 
**MinLength** | Pointer to **float64** |  | [optional] 
**MissingCount** | Pointer to **float64** |  | [optional] 
**MissingPercentage** | Pointer to **float64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NullCount** | Pointer to **float64** |  | [optional] 
**NullProportion** | Pointer to **float64** |  | [optional] 
**Stddev** | Pointer to **float64** |  | [optional] 
**Sum** | Pointer to **float64** |  | [optional] 
**UniqueCount** | Pointer to **float64** |  | [optional] 
**UniqueProportion** | Pointer to **float64** |  | [optional] 
**ValidCount** | Pointer to **float64** |  | [optional] 
**ValuesCount** | Pointer to **float64** |  | [optional] 
**ValuesPercentage** | Pointer to **float64** |  | [optional] 
**Variance** | Pointer to **float64** |  | [optional] 

## Methods

### NewColumnProfile

`func NewColumnProfile() *ColumnProfile`

NewColumnProfile instantiates a new ColumnProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColumnProfileWithDefaults

`func NewColumnProfileWithDefaults() *ColumnProfile`

NewColumnProfileWithDefaults instantiates a new ColumnProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomMetricsProfile

`func (o *ColumnProfile) GetCustomMetricsProfile() []CustomMetricProfile`

GetCustomMetricsProfile returns the CustomMetricsProfile field if non-nil, zero value otherwise.

### GetCustomMetricsProfileOk

`func (o *ColumnProfile) GetCustomMetricsProfileOk() (*[]CustomMetricProfile, bool)`

GetCustomMetricsProfileOk returns a tuple with the CustomMetricsProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetricsProfile

`func (o *ColumnProfile) SetCustomMetricsProfile(v []CustomMetricProfile)`

SetCustomMetricsProfile sets CustomMetricsProfile field to given value.

### HasCustomMetricsProfile

`func (o *ColumnProfile) HasCustomMetricsProfile() bool`

HasCustomMetricsProfile returns a boolean if a field has been set.

### GetDistinctCount

`func (o *ColumnProfile) GetDistinctCount() float64`

GetDistinctCount returns the DistinctCount field if non-nil, zero value otherwise.

### GetDistinctCountOk

`func (o *ColumnProfile) GetDistinctCountOk() (*float64, bool)`

GetDistinctCountOk returns a tuple with the DistinctCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctCount

`func (o *ColumnProfile) SetDistinctCount(v float64)`

SetDistinctCount sets DistinctCount field to given value.

### HasDistinctCount

`func (o *ColumnProfile) HasDistinctCount() bool`

HasDistinctCount returns a boolean if a field has been set.

### GetDistinctProportion

`func (o *ColumnProfile) GetDistinctProportion() float64`

GetDistinctProportion returns the DistinctProportion field if non-nil, zero value otherwise.

### GetDistinctProportionOk

`func (o *ColumnProfile) GetDistinctProportionOk() (*float64, bool)`

GetDistinctProportionOk returns a tuple with the DistinctProportion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctProportion

`func (o *ColumnProfile) SetDistinctProportion(v float64)`

SetDistinctProportion sets DistinctProportion field to given value.

### HasDistinctProportion

`func (o *ColumnProfile) HasDistinctProportion() bool`

HasDistinctProportion returns a boolean if a field has been set.

### GetDuplicateCount

`func (o *ColumnProfile) GetDuplicateCount() float64`

GetDuplicateCount returns the DuplicateCount field if non-nil, zero value otherwise.

### GetDuplicateCountOk

`func (o *ColumnProfile) GetDuplicateCountOk() (*float64, bool)`

GetDuplicateCountOk returns a tuple with the DuplicateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateCount

`func (o *ColumnProfile) SetDuplicateCount(v float64)`

SetDuplicateCount sets DuplicateCount field to given value.

### HasDuplicateCount

`func (o *ColumnProfile) HasDuplicateCount() bool`

HasDuplicateCount returns a boolean if a field has been set.

### GetHistogram

`func (o *ColumnProfile) GetHistogram() Histogram`

GetHistogram returns the Histogram field if non-nil, zero value otherwise.

### GetHistogramOk

`func (o *ColumnProfile) GetHistogramOk() (*Histogram, bool)`

GetHistogramOk returns a tuple with the Histogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogram

`func (o *ColumnProfile) SetHistogram(v Histogram)`

SetHistogram sets Histogram field to given value.

### HasHistogram

`func (o *ColumnProfile) HasHistogram() bool`

HasHistogram returns a boolean if a field has been set.

### GetMax

`func (o *ColumnProfile) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ColumnProfile) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ColumnProfile) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *ColumnProfile) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMaxLength

`func (o *ColumnProfile) GetMaxLength() float64`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *ColumnProfile) GetMaxLengthOk() (*float64, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *ColumnProfile) SetMaxLength(v float64)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *ColumnProfile) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMean

`func (o *ColumnProfile) GetMean() float64`

GetMean returns the Mean field if non-nil, zero value otherwise.

### GetMeanOk

`func (o *ColumnProfile) GetMeanOk() (*float64, bool)`

GetMeanOk returns a tuple with the Mean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMean

`func (o *ColumnProfile) SetMean(v float64)`

SetMean sets Mean field to given value.

### HasMean

`func (o *ColumnProfile) HasMean() bool`

HasMean returns a boolean if a field has been set.

### GetMin

`func (o *ColumnProfile) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *ColumnProfile) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *ColumnProfile) SetMin(v float64)`

SetMin sets Min field to given value.

### HasMin

`func (o *ColumnProfile) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMinLength

`func (o *ColumnProfile) GetMinLength() float64`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *ColumnProfile) GetMinLengthOk() (*float64, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *ColumnProfile) SetMinLength(v float64)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *ColumnProfile) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMissingCount

`func (o *ColumnProfile) GetMissingCount() float64`

GetMissingCount returns the MissingCount field if non-nil, zero value otherwise.

### GetMissingCountOk

`func (o *ColumnProfile) GetMissingCountOk() (*float64, bool)`

GetMissingCountOk returns a tuple with the MissingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingCount

`func (o *ColumnProfile) SetMissingCount(v float64)`

SetMissingCount sets MissingCount field to given value.

### HasMissingCount

`func (o *ColumnProfile) HasMissingCount() bool`

HasMissingCount returns a boolean if a field has been set.

### GetMissingPercentage

`func (o *ColumnProfile) GetMissingPercentage() float64`

GetMissingPercentage returns the MissingPercentage field if non-nil, zero value otherwise.

### GetMissingPercentageOk

`func (o *ColumnProfile) GetMissingPercentageOk() (*float64, bool)`

GetMissingPercentageOk returns a tuple with the MissingPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingPercentage

`func (o *ColumnProfile) SetMissingPercentage(v float64)`

SetMissingPercentage sets MissingPercentage field to given value.

### HasMissingPercentage

`func (o *ColumnProfile) HasMissingPercentage() bool`

HasMissingPercentage returns a boolean if a field has been set.

### GetName

`func (o *ColumnProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ColumnProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ColumnProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ColumnProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNullCount

`func (o *ColumnProfile) GetNullCount() float64`

GetNullCount returns the NullCount field if non-nil, zero value otherwise.

### GetNullCountOk

`func (o *ColumnProfile) GetNullCountOk() (*float64, bool)`

GetNullCountOk returns a tuple with the NullCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullCount

`func (o *ColumnProfile) SetNullCount(v float64)`

SetNullCount sets NullCount field to given value.

### HasNullCount

`func (o *ColumnProfile) HasNullCount() bool`

HasNullCount returns a boolean if a field has been set.

### GetNullProportion

`func (o *ColumnProfile) GetNullProportion() float64`

GetNullProportion returns the NullProportion field if non-nil, zero value otherwise.

### GetNullProportionOk

`func (o *ColumnProfile) GetNullProportionOk() (*float64, bool)`

GetNullProportionOk returns a tuple with the NullProportion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullProportion

`func (o *ColumnProfile) SetNullProportion(v float64)`

SetNullProportion sets NullProportion field to given value.

### HasNullProportion

`func (o *ColumnProfile) HasNullProportion() bool`

HasNullProportion returns a boolean if a field has been set.

### GetStddev

`func (o *ColumnProfile) GetStddev() float64`

GetStddev returns the Stddev field if non-nil, zero value otherwise.

### GetStddevOk

`func (o *ColumnProfile) GetStddevOk() (*float64, bool)`

GetStddevOk returns a tuple with the Stddev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStddev

`func (o *ColumnProfile) SetStddev(v float64)`

SetStddev sets Stddev field to given value.

### HasStddev

`func (o *ColumnProfile) HasStddev() bool`

HasStddev returns a boolean if a field has been set.

### GetSum

`func (o *ColumnProfile) GetSum() float64`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *ColumnProfile) GetSumOk() (*float64, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *ColumnProfile) SetSum(v float64)`

SetSum sets Sum field to given value.

### HasSum

`func (o *ColumnProfile) HasSum() bool`

HasSum returns a boolean if a field has been set.

### GetUniqueCount

`func (o *ColumnProfile) GetUniqueCount() float64`

GetUniqueCount returns the UniqueCount field if non-nil, zero value otherwise.

### GetUniqueCountOk

`func (o *ColumnProfile) GetUniqueCountOk() (*float64, bool)`

GetUniqueCountOk returns a tuple with the UniqueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueCount

`func (o *ColumnProfile) SetUniqueCount(v float64)`

SetUniqueCount sets UniqueCount field to given value.

### HasUniqueCount

`func (o *ColumnProfile) HasUniqueCount() bool`

HasUniqueCount returns a boolean if a field has been set.

### GetUniqueProportion

`func (o *ColumnProfile) GetUniqueProportion() float64`

GetUniqueProportion returns the UniqueProportion field if non-nil, zero value otherwise.

### GetUniqueProportionOk

`func (o *ColumnProfile) GetUniqueProportionOk() (*float64, bool)`

GetUniqueProportionOk returns a tuple with the UniqueProportion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueProportion

`func (o *ColumnProfile) SetUniqueProportion(v float64)`

SetUniqueProportion sets UniqueProportion field to given value.

### HasUniqueProportion

`func (o *ColumnProfile) HasUniqueProportion() bool`

HasUniqueProportion returns a boolean if a field has been set.

### GetValidCount

`func (o *ColumnProfile) GetValidCount() float64`

GetValidCount returns the ValidCount field if non-nil, zero value otherwise.

### GetValidCountOk

`func (o *ColumnProfile) GetValidCountOk() (*float64, bool)`

GetValidCountOk returns a tuple with the ValidCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidCount

`func (o *ColumnProfile) SetValidCount(v float64)`

SetValidCount sets ValidCount field to given value.

### HasValidCount

`func (o *ColumnProfile) HasValidCount() bool`

HasValidCount returns a boolean if a field has been set.

### GetValuesCount

`func (o *ColumnProfile) GetValuesCount() float64`

GetValuesCount returns the ValuesCount field if non-nil, zero value otherwise.

### GetValuesCountOk

`func (o *ColumnProfile) GetValuesCountOk() (*float64, bool)`

GetValuesCountOk returns a tuple with the ValuesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesCount

`func (o *ColumnProfile) SetValuesCount(v float64)`

SetValuesCount sets ValuesCount field to given value.

### HasValuesCount

`func (o *ColumnProfile) HasValuesCount() bool`

HasValuesCount returns a boolean if a field has been set.

### GetValuesPercentage

`func (o *ColumnProfile) GetValuesPercentage() float64`

GetValuesPercentage returns the ValuesPercentage field if non-nil, zero value otherwise.

### GetValuesPercentageOk

`func (o *ColumnProfile) GetValuesPercentageOk() (*float64, bool)`

GetValuesPercentageOk returns a tuple with the ValuesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesPercentage

`func (o *ColumnProfile) SetValuesPercentage(v float64)`

SetValuesPercentage sets ValuesPercentage field to given value.

### HasValuesPercentage

`func (o *ColumnProfile) HasValuesPercentage() bool`

HasValuesPercentage returns a boolean if a field has been set.

### GetVariance

`func (o *ColumnProfile) GetVariance() float64`

GetVariance returns the Variance field if non-nil, zero value otherwise.

### GetVarianceOk

`func (o *ColumnProfile) GetVarianceOk() (*float64, bool)`

GetVarianceOk returns a tuple with the Variance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariance

`func (o *ColumnProfile) SetVariance(v float64)`

SetVariance sets Variance field to given value.

### HasVariance

`func (o *ColumnProfile) HasVariance() bool`

HasVariance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


