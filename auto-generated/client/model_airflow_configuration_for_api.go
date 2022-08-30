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

// AirflowConfigurationForAPI struct for AirflowConfigurationForAPI
type AirflowConfigurationForAPI struct {
	ApiEndpoint string `json:"apiEndpoint"`
}

// NewAirflowConfigurationForAPI instantiates a new AirflowConfigurationForAPI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAirflowConfigurationForAPI(apiEndpoint string) *AirflowConfigurationForAPI {
	this := AirflowConfigurationForAPI{}
	this.ApiEndpoint = apiEndpoint
	return &this
}

// NewAirflowConfigurationForAPIWithDefaults instantiates a new AirflowConfigurationForAPI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAirflowConfigurationForAPIWithDefaults() *AirflowConfigurationForAPI {
	this := AirflowConfigurationForAPI{}
	return &this
}

// GetApiEndpoint returns the ApiEndpoint field value
func (o *AirflowConfigurationForAPI) GetApiEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiEndpoint
}

// GetApiEndpointOk returns a tuple with the ApiEndpoint field value
// and a boolean to check if the value has been set.
func (o *AirflowConfigurationForAPI) GetApiEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiEndpoint, true
}

// SetApiEndpoint sets field value
func (o *AirflowConfigurationForAPI) SetApiEndpoint(v string) {
	o.ApiEndpoint = v
}

func (o AirflowConfigurationForAPI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apiEndpoint"] = o.ApiEndpoint
	}
	return json.Marshal(toSerialize)
}

type NullableAirflowConfigurationForAPI struct {
	value *AirflowConfigurationForAPI
	isSet bool
}

func (v NullableAirflowConfigurationForAPI) Get() *AirflowConfigurationForAPI {
	return v.value
}

func (v *NullableAirflowConfigurationForAPI) Set(val *AirflowConfigurationForAPI) {
	v.value = val
	v.isSet = true
}

func (v NullableAirflowConfigurationForAPI) IsSet() bool {
	return v.isSet
}

func (v *NullableAirflowConfigurationForAPI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAirflowConfigurationForAPI(val *AirflowConfigurationForAPI) *NullableAirflowConfigurationForAPI {
	return &NullableAirflowConfigurationForAPI{value: val, isSet: true}
}

func (v NullableAirflowConfigurationForAPI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAirflowConfigurationForAPI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
