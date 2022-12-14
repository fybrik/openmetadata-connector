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

// ShardSearchFailure struct for ShardSearchFailure
type ShardSearchFailure struct {
	Cause    *ShardSearchFailureCause `json:"cause,omitempty"`
	Fragment *bool                    `json:"fragment,omitempty"`
}

// NewShardSearchFailure instantiates a new ShardSearchFailure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardSearchFailure() *ShardSearchFailure {
	this := ShardSearchFailure{}
	return &this
}

// NewShardSearchFailureWithDefaults instantiates a new ShardSearchFailure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardSearchFailureWithDefaults() *ShardSearchFailure {
	this := ShardSearchFailure{}
	return &this
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *ShardSearchFailure) GetCause() ShardSearchFailureCause {
	if o == nil || o.Cause == nil {
		var ret ShardSearchFailureCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardSearchFailure) GetCauseOk() (*ShardSearchFailureCause, bool) {
	if o == nil || o.Cause == nil {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *ShardSearchFailure) HasCause() bool {
	if o != nil && o.Cause != nil {
		return true
	}

	return false
}

// SetCause gets a reference to the given ShardSearchFailureCause and assigns it to the Cause field.
func (o *ShardSearchFailure) SetCause(v ShardSearchFailureCause) {
	o.Cause = &v
}

// GetFragment returns the Fragment field value if set, zero value otherwise.
func (o *ShardSearchFailure) GetFragment() bool {
	if o == nil || o.Fragment == nil {
		var ret bool
		return ret
	}
	return *o.Fragment
}

// GetFragmentOk returns a tuple with the Fragment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardSearchFailure) GetFragmentOk() (*bool, bool) {
	if o == nil || o.Fragment == nil {
		return nil, false
	}
	return o.Fragment, true
}

// HasFragment returns a boolean if a field has been set.
func (o *ShardSearchFailure) HasFragment() bool {
	if o != nil && o.Fragment != nil {
		return true
	}

	return false
}

// SetFragment gets a reference to the given bool and assigns it to the Fragment field.
func (o *ShardSearchFailure) SetFragment(v bool) {
	o.Fragment = &v
}

func (o ShardSearchFailure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cause != nil {
		toSerialize["cause"] = o.Cause
	}
	if o.Fragment != nil {
		toSerialize["fragment"] = o.Fragment
	}
	return json.Marshal(toSerialize)
}

type NullableShardSearchFailure struct {
	value *ShardSearchFailure
	isSet bool
}

func (v NullableShardSearchFailure) Get() *ShardSearchFailure {
	return v.value
}

func (v *NullableShardSearchFailure) Set(val *ShardSearchFailure) {
	v.value = val
	v.isSet = true
}

func (v NullableShardSearchFailure) IsSet() bool {
	return v.isSet
}

func (v *NullableShardSearchFailure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShardSearchFailure(val *ShardSearchFailure) *NullableShardSearchFailure {
	return &NullableShardSearchFailure{value: val, isSet: true}
}

func (v NullableShardSearchFailure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShardSearchFailure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
