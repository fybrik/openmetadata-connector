# SortField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BytesComparator** | Pointer to **map[string]interface{}** |  | [optional] 
**CanUsePoints** | Pointer to **bool** |  | [optional] 
**ComparatorSource** | Pointer to **map[string]interface{}** |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**IndexSorter** | Pointer to [**IndexSorter**](IndexSorter.md) |  | [optional] 
**MissingValue** | Pointer to **map[string]interface{}** |  | [optional] 
**Reverse** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewSortField

`func NewSortField() *SortField`

NewSortField instantiates a new SortField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortFieldWithDefaults

`func NewSortFieldWithDefaults() *SortField`

NewSortFieldWithDefaults instantiates a new SortField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytesComparator

`func (o *SortField) GetBytesComparator() map[string]interface{}`

GetBytesComparator returns the BytesComparator field if non-nil, zero value otherwise.

### GetBytesComparatorOk

`func (o *SortField) GetBytesComparatorOk() (*map[string]interface{}, bool)`

GetBytesComparatorOk returns a tuple with the BytesComparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesComparator

`func (o *SortField) SetBytesComparator(v map[string]interface{})`

SetBytesComparator sets BytesComparator field to given value.

### HasBytesComparator

`func (o *SortField) HasBytesComparator() bool`

HasBytesComparator returns a boolean if a field has been set.

### GetCanUsePoints

`func (o *SortField) GetCanUsePoints() bool`

GetCanUsePoints returns the CanUsePoints field if non-nil, zero value otherwise.

### GetCanUsePointsOk

`func (o *SortField) GetCanUsePointsOk() (*bool, bool)`

GetCanUsePointsOk returns a tuple with the CanUsePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUsePoints

`func (o *SortField) SetCanUsePoints(v bool)`

SetCanUsePoints sets CanUsePoints field to given value.

### HasCanUsePoints

`func (o *SortField) HasCanUsePoints() bool`

HasCanUsePoints returns a boolean if a field has been set.

### GetComparatorSource

`func (o *SortField) GetComparatorSource() map[string]interface{}`

GetComparatorSource returns the ComparatorSource field if non-nil, zero value otherwise.

### GetComparatorSourceOk

`func (o *SortField) GetComparatorSourceOk() (*map[string]interface{}, bool)`

GetComparatorSourceOk returns a tuple with the ComparatorSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparatorSource

`func (o *SortField) SetComparatorSource(v map[string]interface{})`

SetComparatorSource sets ComparatorSource field to given value.

### HasComparatorSource

`func (o *SortField) HasComparatorSource() bool`

HasComparatorSource returns a boolean if a field has been set.

### GetField

`func (o *SortField) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *SortField) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *SortField) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *SortField) HasField() bool`

HasField returns a boolean if a field has been set.

### GetIndexSorter

`func (o *SortField) GetIndexSorter() IndexSorter`

GetIndexSorter returns the IndexSorter field if non-nil, zero value otherwise.

### GetIndexSorterOk

`func (o *SortField) GetIndexSorterOk() (*IndexSorter, bool)`

GetIndexSorterOk returns a tuple with the IndexSorter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexSorter

`func (o *SortField) SetIndexSorter(v IndexSorter)`

SetIndexSorter sets IndexSorter field to given value.

### HasIndexSorter

`func (o *SortField) HasIndexSorter() bool`

HasIndexSorter returns a boolean if a field has been set.

### GetMissingValue

`func (o *SortField) GetMissingValue() map[string]interface{}`

GetMissingValue returns the MissingValue field if non-nil, zero value otherwise.

### GetMissingValueOk

`func (o *SortField) GetMissingValueOk() (*map[string]interface{}, bool)`

GetMissingValueOk returns a tuple with the MissingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingValue

`func (o *SortField) SetMissingValue(v map[string]interface{})`

SetMissingValue sets MissingValue field to given value.

### HasMissingValue

`func (o *SortField) HasMissingValue() bool`

HasMissingValue returns a boolean if a field has been set.

### GetReverse

`func (o *SortField) GetReverse() bool`

GetReverse returns the Reverse field if non-nil, zero value otherwise.

### GetReverseOk

`func (o *SortField) GetReverseOk() (*bool, bool)`

GetReverseOk returns a tuple with the Reverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverse

`func (o *SortField) SetReverse(v bool)`

SetReverse sets Reverse field to given value.

### HasReverse

`func (o *SortField) HasReverse() bool`

HasReverse returns a boolean if a field has been set.

### GetType

`func (o *SortField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SortField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SortField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SortField) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


