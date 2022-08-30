/*
OpenMetadata Apis

# Overview  OpenMetadata supports REST APIs for getting metadata in and out of metadata store. The API resources are grouped under following categories: - **Data assets** - includes resources for data entities, such as `databases`, `tables`, and `topics`. Resources for data assets created from data, such as `dashboards`, `reports`, `metrics`, and `ML Features` are part of this collection. `pipelines`, `notebooks`, etc. that are used for creating data assets are also available as resources as of this collection. - **Teams and Users** - includes `users`, `teams`, a special type of user called `bots` that performs many automated tasks such as ingestion. - **Services** - are services that OpenMetadata integrates with. Currently `databaseService` is the only service under this collection that represents data sources. In the future, services related to Dashboards, Reports, ETL pipelines will be added under this collection. - **Glossary** - OpenMetadata supports hierarchical tags that can be used to build business vocabulary to describe and classify data available under `tags` resource.

API version: 0.11.0
Contact: openmetadata-dev@googlegroups.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SearchHits struct for SearchHits
type SearchHits struct {
	CollapseField  *string                  `json:"collapseField,omitempty"`
	CollapseValues []map[string]interface{} `json:"collapseValues,omitempty"`
	Fragment       *bool                    `json:"fragment,omitempty"`
	Hits           []SearchHit              `json:"hits,omitempty"`
	MaxScore       *float32                 `json:"maxScore,omitempty"`
	SortFields     []SortField              `json:"sortFields,omitempty"`
	TotalHits      *TotalHits               `json:"totalHits,omitempty"`
}

// NewSearchHits instantiates a new SearchHits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchHits() *SearchHits {
	this := SearchHits{}
	return &this
}

// NewSearchHitsWithDefaults instantiates a new SearchHits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchHitsWithDefaults() *SearchHits {
	this := SearchHits{}
	return &this
}

// GetCollapseField returns the CollapseField field value if set, zero value otherwise.
func (o *SearchHits) GetCollapseField() string {
	if o == nil || o.CollapseField == nil {
		var ret string
		return ret
	}
	return *o.CollapseField
}

// GetCollapseFieldOk returns a tuple with the CollapseField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHits) GetCollapseFieldOk() (*string, bool) {
	if o == nil || o.CollapseField == nil {
		return nil, false
	}
	return o.CollapseField, true
}

// HasCollapseField returns a boolean if a field has been set.
func (o *SearchHits) HasCollapseField() bool {
	if o != nil && o.CollapseField != nil {
		return true
	}

	return false
}

// SetCollapseField gets a reference to the given string and assigns it to the CollapseField field.
func (o *SearchHits) SetCollapseField(v string) {
	o.CollapseField = &v
}

// GetCollapseValues returns the CollapseValues field value if set, zero value otherwise.
func (o *SearchHits) GetCollapseValues() []map[string]interface{} {
	if o == nil || o.CollapseValues == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.CollapseValues
}

// GetCollapseValuesOk returns a tuple with the CollapseValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHits) GetCollapseValuesOk() ([]map[string]interface{}, bool) {
	if o == nil || o.CollapseValues == nil {
		return nil, false
	}
	return o.CollapseValues, true
}

// HasCollapseValues returns a boolean if a field has been set.
func (o *SearchHits) HasCollapseValues() bool {
	if o != nil && o.CollapseValues != nil {
		return true
	}

	return false
}

// SetCollapseValues gets a reference to the given []map[string]interface{} and assigns it to the CollapseValues field.
func (o *SearchHits) SetCollapseValues(v []map[string]interface{}) {
	o.CollapseValues = v
}

// GetFragment returns the Fragment field value if set, zero value otherwise.
func (o *SearchHits) GetFragment() bool {
	if o == nil || o.Fragment == nil {
		var ret bool
		return ret
	}
	return *o.Fragment
}

// GetFragmentOk returns a tuple with the Fragment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHits) GetFragmentOk() (*bool, bool) {
	if o == nil || o.Fragment == nil {
		return nil, false
	}
	return o.Fragment, true
}

// HasFragment returns a boolean if a field has been set.
func (o *SearchHits) HasFragment() bool {
	if o != nil && o.Fragment != nil {
		return true
	}

	return false
}

// SetFragment gets a reference to the given bool and assigns it to the Fragment field.
func (o *SearchHits) SetFragment(v bool) {
	o.Fragment = &v
}

// GetHits returns the Hits field value if set, zero value otherwise.
func (o *SearchHits) GetHits() []SearchHit {
	if o == nil || o.Hits == nil {
		var ret []SearchHit
		return ret
	}
	return o.Hits
}

// GetHitsOk returns a tuple with the Hits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHits) GetHitsOk() ([]SearchHit, bool) {
	if o == nil || o.Hits == nil {
		return nil, false
	}
	return o.Hits, true
}

// HasHits returns a boolean if a field has been set.
func (o *SearchHits) HasHits() bool {
	if o != nil && o.Hits != nil {
		return true
	}

	return false
}

// SetHits gets a reference to the given []SearchHit and assigns it to the Hits field.
func (o *SearchHits) SetHits(v []SearchHit) {
	o.Hits = v
}

// GetMaxScore returns the MaxScore field value if set, zero value otherwise.
func (o *SearchHits) GetMaxScore() float32 {
	if o == nil || o.MaxScore == nil {
		var ret float32
		return ret
	}
	return *o.MaxScore
}

// GetMaxScoreOk returns a tuple with the MaxScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHits) GetMaxScoreOk() (*float32, bool) {
	if o == nil || o.MaxScore == nil {
		return nil, false
	}
	return o.MaxScore, true
}

// HasMaxScore returns a boolean if a field has been set.
func (o *SearchHits) HasMaxScore() bool {
	if o != nil && o.MaxScore != nil {
		return true
	}

	return false
}

// SetMaxScore gets a reference to the given float32 and assigns it to the MaxScore field.
func (o *SearchHits) SetMaxScore(v float32) {
	o.MaxScore = &v
}

// GetSortFields returns the SortFields field value if set, zero value otherwise.
func (o *SearchHits) GetSortFields() []SortField {
	if o == nil || o.SortFields == nil {
		var ret []SortField
		return ret
	}
	return o.SortFields
}

// GetSortFieldsOk returns a tuple with the SortFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHits) GetSortFieldsOk() ([]SortField, bool) {
	if o == nil || o.SortFields == nil {
		return nil, false
	}
	return o.SortFields, true
}

// HasSortFields returns a boolean if a field has been set.
func (o *SearchHits) HasSortFields() bool {
	if o != nil && o.SortFields != nil {
		return true
	}

	return false
}

// SetSortFields gets a reference to the given []SortField and assigns it to the SortFields field.
func (o *SearchHits) SetSortFields(v []SortField) {
	o.SortFields = v
}

// GetTotalHits returns the TotalHits field value if set, zero value otherwise.
func (o *SearchHits) GetTotalHits() TotalHits {
	if o == nil || o.TotalHits == nil {
		var ret TotalHits
		return ret
	}
	return *o.TotalHits
}

// GetTotalHitsOk returns a tuple with the TotalHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHits) GetTotalHitsOk() (*TotalHits, bool) {
	if o == nil || o.TotalHits == nil {
		return nil, false
	}
	return o.TotalHits, true
}

// HasTotalHits returns a boolean if a field has been set.
func (o *SearchHits) HasTotalHits() bool {
	if o != nil && o.TotalHits != nil {
		return true
	}

	return false
}

// SetTotalHits gets a reference to the given TotalHits and assigns it to the TotalHits field.
func (o *SearchHits) SetTotalHits(v TotalHits) {
	o.TotalHits = &v
}

func (o SearchHits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CollapseField != nil {
		toSerialize["collapseField"] = o.CollapseField
	}
	if o.CollapseValues != nil {
		toSerialize["collapseValues"] = o.CollapseValues
	}
	if o.Fragment != nil {
		toSerialize["fragment"] = o.Fragment
	}
	if o.Hits != nil {
		toSerialize["hits"] = o.Hits
	}
	if o.MaxScore != nil {
		toSerialize["maxScore"] = o.MaxScore
	}
	if o.SortFields != nil {
		toSerialize["sortFields"] = o.SortFields
	}
	if o.TotalHits != nil {
		toSerialize["totalHits"] = o.TotalHits
	}
	return json.Marshal(toSerialize)
}

type NullableSearchHits struct {
	value *SearchHits
	isSet bool
}

func (v NullableSearchHits) Get() *SearchHits {
	return v.value
}

func (v *NullableSearchHits) Set(val *SearchHits) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchHits) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchHits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchHits(val *SearchHits) *NullableSearchHits {
	return &NullableSearchHits{value: val, isSet: true}
}

func (v NullableSearchHits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchHits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
