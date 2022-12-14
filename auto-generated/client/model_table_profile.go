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

// TableProfile struct for TableProfile
type TableProfile struct {
	ColumnCount   *float64                 `json:"columnCount,omitempty"`
	ColumnNames   []map[string]interface{} `json:"columnNames,omitempty"`
	ColumnProfile []ColumnProfile          `json:"columnProfile,omitempty"`
	ProfileDate   *string                  `json:"profileDate,omitempty"`
	ProfileQuery  *string                  `json:"profileQuery,omitempty"`
	ProfileSample *float64                 `json:"profileSample,omitempty"`
	RowCount      *float64                 `json:"rowCount,omitempty"`
}

// NewTableProfile instantiates a new TableProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableProfile() *TableProfile {
	this := TableProfile{}
	return &this
}

// NewTableProfileWithDefaults instantiates a new TableProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableProfileWithDefaults() *TableProfile {
	this := TableProfile{}
	return &this
}

// GetColumnCount returns the ColumnCount field value if set, zero value otherwise.
func (o *TableProfile) GetColumnCount() float64 {
	if o == nil || o.ColumnCount == nil {
		var ret float64
		return ret
	}
	return *o.ColumnCount
}

// GetColumnCountOk returns a tuple with the ColumnCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableProfile) GetColumnCountOk() (*float64, bool) {
	if o == nil || o.ColumnCount == nil {
		return nil, false
	}
	return o.ColumnCount, true
}

// HasColumnCount returns a boolean if a field has been set.
func (o *TableProfile) HasColumnCount() bool {
	if o != nil && o.ColumnCount != nil {
		return true
	}

	return false
}

// SetColumnCount gets a reference to the given float64 and assigns it to the ColumnCount field.
func (o *TableProfile) SetColumnCount(v float64) {
	o.ColumnCount = &v
}

// GetColumnNames returns the ColumnNames field value if set, zero value otherwise.
func (o *TableProfile) GetColumnNames() []map[string]interface{} {
	if o == nil || o.ColumnNames == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.ColumnNames
}

// GetColumnNamesOk returns a tuple with the ColumnNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableProfile) GetColumnNamesOk() ([]map[string]interface{}, bool) {
	if o == nil || o.ColumnNames == nil {
		return nil, false
	}
	return o.ColumnNames, true
}

// HasColumnNames returns a boolean if a field has been set.
func (o *TableProfile) HasColumnNames() bool {
	if o != nil && o.ColumnNames != nil {
		return true
	}

	return false
}

// SetColumnNames gets a reference to the given []map[string]interface{} and assigns it to the ColumnNames field.
func (o *TableProfile) SetColumnNames(v []map[string]interface{}) {
	o.ColumnNames = v
}

// GetColumnProfile returns the ColumnProfile field value if set, zero value otherwise.
func (o *TableProfile) GetColumnProfile() []ColumnProfile {
	if o == nil || o.ColumnProfile == nil {
		var ret []ColumnProfile
		return ret
	}
	return o.ColumnProfile
}

// GetColumnProfileOk returns a tuple with the ColumnProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableProfile) GetColumnProfileOk() ([]ColumnProfile, bool) {
	if o == nil || o.ColumnProfile == nil {
		return nil, false
	}
	return o.ColumnProfile, true
}

// HasColumnProfile returns a boolean if a field has been set.
func (o *TableProfile) HasColumnProfile() bool {
	if o != nil && o.ColumnProfile != nil {
		return true
	}

	return false
}

// SetColumnProfile gets a reference to the given []ColumnProfile and assigns it to the ColumnProfile field.
func (o *TableProfile) SetColumnProfile(v []ColumnProfile) {
	o.ColumnProfile = v
}

// GetProfileDate returns the ProfileDate field value if set, zero value otherwise.
func (o *TableProfile) GetProfileDate() string {
	if o == nil || o.ProfileDate == nil {
		var ret string
		return ret
	}
	return *o.ProfileDate
}

// GetProfileDateOk returns a tuple with the ProfileDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableProfile) GetProfileDateOk() (*string, bool) {
	if o == nil || o.ProfileDate == nil {
		return nil, false
	}
	return o.ProfileDate, true
}

// HasProfileDate returns a boolean if a field has been set.
func (o *TableProfile) HasProfileDate() bool {
	if o != nil && o.ProfileDate != nil {
		return true
	}

	return false
}

// SetProfileDate gets a reference to the given string and assigns it to the ProfileDate field.
func (o *TableProfile) SetProfileDate(v string) {
	o.ProfileDate = &v
}

// GetProfileQuery returns the ProfileQuery field value if set, zero value otherwise.
func (o *TableProfile) GetProfileQuery() string {
	if o == nil || o.ProfileQuery == nil {
		var ret string
		return ret
	}
	return *o.ProfileQuery
}

// GetProfileQueryOk returns a tuple with the ProfileQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableProfile) GetProfileQueryOk() (*string, bool) {
	if o == nil || o.ProfileQuery == nil {
		return nil, false
	}
	return o.ProfileQuery, true
}

// HasProfileQuery returns a boolean if a field has been set.
func (o *TableProfile) HasProfileQuery() bool {
	if o != nil && o.ProfileQuery != nil {
		return true
	}

	return false
}

// SetProfileQuery gets a reference to the given string and assigns it to the ProfileQuery field.
func (o *TableProfile) SetProfileQuery(v string) {
	o.ProfileQuery = &v
}

// GetProfileSample returns the ProfileSample field value if set, zero value otherwise.
func (o *TableProfile) GetProfileSample() float64 {
	if o == nil || o.ProfileSample == nil {
		var ret float64
		return ret
	}
	return *o.ProfileSample
}

// GetProfileSampleOk returns a tuple with the ProfileSample field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableProfile) GetProfileSampleOk() (*float64, bool) {
	if o == nil || o.ProfileSample == nil {
		return nil, false
	}
	return o.ProfileSample, true
}

// HasProfileSample returns a boolean if a field has been set.
func (o *TableProfile) HasProfileSample() bool {
	if o != nil && o.ProfileSample != nil {
		return true
	}

	return false
}

// SetProfileSample gets a reference to the given float64 and assigns it to the ProfileSample field.
func (o *TableProfile) SetProfileSample(v float64) {
	o.ProfileSample = &v
}

// GetRowCount returns the RowCount field value if set, zero value otherwise.
func (o *TableProfile) GetRowCount() float64 {
	if o == nil || o.RowCount == nil {
		var ret float64
		return ret
	}
	return *o.RowCount
}

// GetRowCountOk returns a tuple with the RowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableProfile) GetRowCountOk() (*float64, bool) {
	if o == nil || o.RowCount == nil {
		return nil, false
	}
	return o.RowCount, true
}

// HasRowCount returns a boolean if a field has been set.
func (o *TableProfile) HasRowCount() bool {
	if o != nil && o.RowCount != nil {
		return true
	}

	return false
}

// SetRowCount gets a reference to the given float64 and assigns it to the RowCount field.
func (o *TableProfile) SetRowCount(v float64) {
	o.RowCount = &v
}

func (o TableProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ColumnCount != nil {
		toSerialize["columnCount"] = o.ColumnCount
	}
	if o.ColumnNames != nil {
		toSerialize["columnNames"] = o.ColumnNames
	}
	if o.ColumnProfile != nil {
		toSerialize["columnProfile"] = o.ColumnProfile
	}
	if o.ProfileDate != nil {
		toSerialize["profileDate"] = o.ProfileDate
	}
	if o.ProfileQuery != nil {
		toSerialize["profileQuery"] = o.ProfileQuery
	}
	if o.ProfileSample != nil {
		toSerialize["profileSample"] = o.ProfileSample
	}
	if o.RowCount != nil {
		toSerialize["rowCount"] = o.RowCount
	}
	return json.Marshal(toSerialize)
}

type NullableTableProfile struct {
	value *TableProfile
	isSet bool
}

func (v NullableTableProfile) Get() *TableProfile {
	return v.value
}

func (v *NullableTableProfile) Set(val *TableProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableTableProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableTableProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableProfile(val *TableProfile) *NullableTableProfile {
	return &NullableTableProfile{value: val, isSet: true}
}

func (v NullableTableProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
