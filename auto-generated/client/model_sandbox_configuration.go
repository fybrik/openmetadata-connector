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

// SandboxConfiguration struct for SandboxConfiguration
type SandboxConfiguration struct {
	SandboxModeEnabled *bool `json:"sandboxModeEnabled,omitempty"`
}

// NewSandboxConfiguration instantiates a new SandboxConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxConfiguration() *SandboxConfiguration {
	this := SandboxConfiguration{}
	return &this
}

// NewSandboxConfigurationWithDefaults instantiates a new SandboxConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxConfigurationWithDefaults() *SandboxConfiguration {
	this := SandboxConfiguration{}
	return &this
}

// GetSandboxModeEnabled returns the SandboxModeEnabled field value if set, zero value otherwise.
func (o *SandboxConfiguration) GetSandboxModeEnabled() bool {
	if o == nil || o.SandboxModeEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SandboxModeEnabled
}

// GetSandboxModeEnabledOk returns a tuple with the SandboxModeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxConfiguration) GetSandboxModeEnabledOk() (*bool, bool) {
	if o == nil || o.SandboxModeEnabled == nil {
		return nil, false
	}
	return o.SandboxModeEnabled, true
}

// HasSandboxModeEnabled returns a boolean if a field has been set.
func (o *SandboxConfiguration) HasSandboxModeEnabled() bool {
	if o != nil && o.SandboxModeEnabled != nil {
		return true
	}

	return false
}

// SetSandboxModeEnabled gets a reference to the given bool and assigns it to the SandboxModeEnabled field.
func (o *SandboxConfiguration) SetSandboxModeEnabled(v bool) {
	o.SandboxModeEnabled = &v
}

func (o SandboxConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SandboxModeEnabled != nil {
		toSerialize["sandboxModeEnabled"] = o.SandboxModeEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxConfiguration struct {
	value *SandboxConfiguration
	isSet bool
}

func (v NullableSandboxConfiguration) Get() *SandboxConfiguration {
	return v.value
}

func (v *NullableSandboxConfiguration) Set(val *SandboxConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxConfiguration(val *SandboxConfiguration) *NullableSandboxConfiguration {
	return &NullableSandboxConfiguration{value: val, isSet: true}
}

func (v NullableSandboxConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
