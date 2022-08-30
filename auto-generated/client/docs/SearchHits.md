# SearchHits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollapseField** | Pointer to **string** |  | [optional] 
**CollapseValues** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Fragment** | Pointer to **bool** |  | [optional] 
**Hits** | Pointer to [**[]SearchHit**](SearchHit.md) |  | [optional] 
**MaxScore** | Pointer to **float32** |  | [optional] 
**SortFields** | Pointer to [**[]SortField**](SortField.md) |  | [optional] 
**TotalHits** | Pointer to [**TotalHits**](TotalHits.md) |  | [optional] 

## Methods

### NewSearchHits

`func NewSearchHits() *SearchHits`

NewSearchHits instantiates a new SearchHits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchHitsWithDefaults

`func NewSearchHitsWithDefaults() *SearchHits`

NewSearchHitsWithDefaults instantiates a new SearchHits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollapseField

`func (o *SearchHits) GetCollapseField() string`

GetCollapseField returns the CollapseField field if non-nil, zero value otherwise.

### GetCollapseFieldOk

`func (o *SearchHits) GetCollapseFieldOk() (*string, bool)`

GetCollapseFieldOk returns a tuple with the CollapseField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapseField

`func (o *SearchHits) SetCollapseField(v string)`

SetCollapseField sets CollapseField field to given value.

### HasCollapseField

`func (o *SearchHits) HasCollapseField() bool`

HasCollapseField returns a boolean if a field has been set.

### GetCollapseValues

`func (o *SearchHits) GetCollapseValues() []map[string]interface{}`

GetCollapseValues returns the CollapseValues field if non-nil, zero value otherwise.

### GetCollapseValuesOk

`func (o *SearchHits) GetCollapseValuesOk() (*[]map[string]interface{}, bool)`

GetCollapseValuesOk returns a tuple with the CollapseValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapseValues

`func (o *SearchHits) SetCollapseValues(v []map[string]interface{})`

SetCollapseValues sets CollapseValues field to given value.

### HasCollapseValues

`func (o *SearchHits) HasCollapseValues() bool`

HasCollapseValues returns a boolean if a field has been set.

### GetFragment

`func (o *SearchHits) GetFragment() bool`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *SearchHits) GetFragmentOk() (*bool, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *SearchHits) SetFragment(v bool)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *SearchHits) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetHits

`func (o *SearchHits) GetHits() []SearchHit`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *SearchHits) GetHitsOk() (*[]SearchHit, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *SearchHits) SetHits(v []SearchHit)`

SetHits sets Hits field to given value.

### HasHits

`func (o *SearchHits) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetMaxScore

`func (o *SearchHits) GetMaxScore() float32`

GetMaxScore returns the MaxScore field if non-nil, zero value otherwise.

### GetMaxScoreOk

`func (o *SearchHits) GetMaxScoreOk() (*float32, bool)`

GetMaxScoreOk returns a tuple with the MaxScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScore

`func (o *SearchHits) SetMaxScore(v float32)`

SetMaxScore sets MaxScore field to given value.

### HasMaxScore

`func (o *SearchHits) HasMaxScore() bool`

HasMaxScore returns a boolean if a field has been set.

### GetSortFields

`func (o *SearchHits) GetSortFields() []SortField`

GetSortFields returns the SortFields field if non-nil, zero value otherwise.

### GetSortFieldsOk

`func (o *SearchHits) GetSortFieldsOk() (*[]SortField, bool)`

GetSortFieldsOk returns a tuple with the SortFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortFields

`func (o *SearchHits) SetSortFields(v []SortField)`

SetSortFields sets SortFields field to given value.

### HasSortFields

`func (o *SearchHits) HasSortFields() bool`

HasSortFields returns a boolean if a field has been set.

### GetTotalHits

`func (o *SearchHits) GetTotalHits() TotalHits`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *SearchHits) GetTotalHitsOk() (*TotalHits, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *SearchHits) SetTotalHits(v TotalHits)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *SearchHits) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


