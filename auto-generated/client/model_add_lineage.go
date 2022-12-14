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

// AddLineage struct for AddLineage
type AddLineage struct {
	Description *string      `json:"description,omitempty"`
	Edge        EntitiesEdge `json:"edge"`
}

// NewAddLineage instantiates a new AddLineage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddLineage(edge EntitiesEdge) *AddLineage {
	this := AddLineage{}
	this.Edge = edge
	return &this
}

// NewAddLineageWithDefaults instantiates a new AddLineage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddLineageWithDefaults() *AddLineage {
	this := AddLineage{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddLineage) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLineage) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddLineage) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddLineage) SetDescription(v string) {
	o.Description = &v
}

// GetEdge returns the Edge field value
func (o *AddLineage) GetEdge() EntitiesEdge {
	if o == nil {
		var ret EntitiesEdge
		return ret
	}

	return o.Edge
}

// GetEdgeOk returns a tuple with the Edge field value
// and a boolean to check if the value has been set.
func (o *AddLineage) GetEdgeOk() (*EntitiesEdge, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Edge, true
}

// SetEdge sets field value
func (o *AddLineage) SetEdge(v EntitiesEdge) {
	o.Edge = v
}

func (o AddLineage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["edge"] = o.Edge
	}
	return json.Marshal(toSerialize)
}

type NullableAddLineage struct {
	value *AddLineage
	isSet bool
}

func (v NullableAddLineage) Get() *AddLineage {
	return v.value
}

func (v *NullableAddLineage) Set(val *AddLineage) {
	v.value = val
	v.isSet = true
}

func (v NullableAddLineage) IsSet() bool {
	return v.isSet
}

func (v *NullableAddLineage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddLineage(val *AddLineage) *NullableAddLineage {
	return &NullableAddLineage{value: val, isSet: true}
}

func (v NullableAddLineage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLineage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
