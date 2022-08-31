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

// TableJoins struct for TableJoins
type TableJoins struct {
	ColumnJoins      []ColumnJoin `json:"columnJoins,omitempty"`
	DayCount         *int32       `json:"dayCount,omitempty"`
	DirectTableJoins []JoinedWith `json:"directTableJoins,omitempty"`
	StartDate        *string      `json:"startDate,omitempty"`
}

// NewTableJoins instantiates a new TableJoins object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableJoins() *TableJoins {
	this := TableJoins{}
	return &this
}

// NewTableJoinsWithDefaults instantiates a new TableJoins object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableJoinsWithDefaults() *TableJoins {
	this := TableJoins{}
	return &this
}

// GetColumnJoins returns the ColumnJoins field value if set, zero value otherwise.
func (o *TableJoins) GetColumnJoins() []ColumnJoin {
	if o == nil || o.ColumnJoins == nil {
		var ret []ColumnJoin
		return ret
	}
	return o.ColumnJoins
}

// GetColumnJoinsOk returns a tuple with the ColumnJoins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableJoins) GetColumnJoinsOk() ([]ColumnJoin, bool) {
	if o == nil || o.ColumnJoins == nil {
		return nil, false
	}
	return o.ColumnJoins, true
}

// HasColumnJoins returns a boolean if a field has been set.
func (o *TableJoins) HasColumnJoins() bool {
	if o != nil && o.ColumnJoins != nil {
		return true
	}

	return false
}

// SetColumnJoins gets a reference to the given []ColumnJoin and assigns it to the ColumnJoins field.
func (o *TableJoins) SetColumnJoins(v []ColumnJoin) {
	o.ColumnJoins = v
}

// GetDayCount returns the DayCount field value if set, zero value otherwise.
func (o *TableJoins) GetDayCount() int32 {
	if o == nil || o.DayCount == nil {
		var ret int32
		return ret
	}
	return *o.DayCount
}

// GetDayCountOk returns a tuple with the DayCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableJoins) GetDayCountOk() (*int32, bool) {
	if o == nil || o.DayCount == nil {
		return nil, false
	}
	return o.DayCount, true
}

// HasDayCount returns a boolean if a field has been set.
func (o *TableJoins) HasDayCount() bool {
	if o != nil && o.DayCount != nil {
		return true
	}

	return false
}

// SetDayCount gets a reference to the given int32 and assigns it to the DayCount field.
func (o *TableJoins) SetDayCount(v int32) {
	o.DayCount = &v
}

// GetDirectTableJoins returns the DirectTableJoins field value if set, zero value otherwise.
func (o *TableJoins) GetDirectTableJoins() []JoinedWith {
	if o == nil || o.DirectTableJoins == nil {
		var ret []JoinedWith
		return ret
	}
	return o.DirectTableJoins
}

// GetDirectTableJoinsOk returns a tuple with the DirectTableJoins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableJoins) GetDirectTableJoinsOk() ([]JoinedWith, bool) {
	if o == nil || o.DirectTableJoins == nil {
		return nil, false
	}
	return o.DirectTableJoins, true
}

// HasDirectTableJoins returns a boolean if a field has been set.
func (o *TableJoins) HasDirectTableJoins() bool {
	if o != nil && o.DirectTableJoins != nil {
		return true
	}

	return false
}

// SetDirectTableJoins gets a reference to the given []JoinedWith and assigns it to the DirectTableJoins field.
func (o *TableJoins) SetDirectTableJoins(v []JoinedWith) {
	o.DirectTableJoins = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *TableJoins) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableJoins) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *TableJoins) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *TableJoins) SetStartDate(v string) {
	o.StartDate = &v
}

func (o TableJoins) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ColumnJoins != nil {
		toSerialize["columnJoins"] = o.ColumnJoins
	}
	if o.DayCount != nil {
		toSerialize["dayCount"] = o.DayCount
	}
	if o.DirectTableJoins != nil {
		toSerialize["directTableJoins"] = o.DirectTableJoins
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	return json.Marshal(toSerialize)
}

type NullableTableJoins struct {
	value *TableJoins
	isSet bool
}

func (v NullableTableJoins) Get() *TableJoins {
	return v.value
}

func (v *NullableTableJoins) Set(val *TableJoins) {
	v.value = val
	v.isSet = true
}

func (v NullableTableJoins) IsSet() bool {
	return v.isSet
}

func (v *NullableTableJoins) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableJoins(val *TableJoins) *NullableTableJoins {
	return &NullableTableJoins{value: val, isSet: true}
}

func (v NullableTableJoins) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableJoins) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}