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

// EntityLineage struct for EntityLineage
type EntityLineage struct {
	DownstreamEdges []Edge            `json:"downstreamEdges,omitempty"`
	Entity          EntityReference   `json:"entity"`
	Nodes           []EntityReference `json:"nodes,omitempty"`
	UpstreamEdges   []Edge            `json:"upstreamEdges,omitempty"`
}

// NewEntityLineage instantiates a new EntityLineage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityLineage(entity EntityReference) *EntityLineage {
	this := EntityLineage{}
	this.Entity = entity
	return &this
}

// NewEntityLineageWithDefaults instantiates a new EntityLineage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityLineageWithDefaults() *EntityLineage {
	this := EntityLineage{}
	return &this
}

// GetDownstreamEdges returns the DownstreamEdges field value if set, zero value otherwise.
func (o *EntityLineage) GetDownstreamEdges() []Edge {
	if o == nil || o.DownstreamEdges == nil {
		var ret []Edge
		return ret
	}
	return o.DownstreamEdges
}

// GetDownstreamEdgesOk returns a tuple with the DownstreamEdges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityLineage) GetDownstreamEdgesOk() ([]Edge, bool) {
	if o == nil || o.DownstreamEdges == nil {
		return nil, false
	}
	return o.DownstreamEdges, true
}

// HasDownstreamEdges returns a boolean if a field has been set.
func (o *EntityLineage) HasDownstreamEdges() bool {
	if o != nil && o.DownstreamEdges != nil {
		return true
	}

	return false
}

// SetDownstreamEdges gets a reference to the given []Edge and assigns it to the DownstreamEdges field.
func (o *EntityLineage) SetDownstreamEdges(v []Edge) {
	o.DownstreamEdges = v
}

// GetEntity returns the Entity field value
func (o *EntityLineage) GetEntity() EntityReference {
	if o == nil {
		var ret EntityReference
		return ret
	}

	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value
// and a boolean to check if the value has been set.
func (o *EntityLineage) GetEntityOk() (*EntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entity, true
}

// SetEntity sets field value
func (o *EntityLineage) SetEntity(v EntityReference) {
	o.Entity = v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *EntityLineage) GetNodes() []EntityReference {
	if o == nil || o.Nodes == nil {
		var ret []EntityReference
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityLineage) GetNodesOk() ([]EntityReference, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *EntityLineage) HasNodes() bool {
	if o != nil && o.Nodes != nil {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []EntityReference and assigns it to the Nodes field.
func (o *EntityLineage) SetNodes(v []EntityReference) {
	o.Nodes = v
}

// GetUpstreamEdges returns the UpstreamEdges field value if set, zero value otherwise.
func (o *EntityLineage) GetUpstreamEdges() []Edge {
	if o == nil || o.UpstreamEdges == nil {
		var ret []Edge
		return ret
	}
	return o.UpstreamEdges
}

// GetUpstreamEdgesOk returns a tuple with the UpstreamEdges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityLineage) GetUpstreamEdgesOk() ([]Edge, bool) {
	if o == nil || o.UpstreamEdges == nil {
		return nil, false
	}
	return o.UpstreamEdges, true
}

// HasUpstreamEdges returns a boolean if a field has been set.
func (o *EntityLineage) HasUpstreamEdges() bool {
	if o != nil && o.UpstreamEdges != nil {
		return true
	}

	return false
}

// SetUpstreamEdges gets a reference to the given []Edge and assigns it to the UpstreamEdges field.
func (o *EntityLineage) SetUpstreamEdges(v []Edge) {
	o.UpstreamEdges = v
}

func (o EntityLineage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DownstreamEdges != nil {
		toSerialize["downstreamEdges"] = o.DownstreamEdges
	}
	if true {
		toSerialize["entity"] = o.Entity
	}
	if o.Nodes != nil {
		toSerialize["nodes"] = o.Nodes
	}
	if o.UpstreamEdges != nil {
		toSerialize["upstreamEdges"] = o.UpstreamEdges
	}
	return json.Marshal(toSerialize)
}

type NullableEntityLineage struct {
	value *EntityLineage
	isSet bool
}

func (v NullableEntityLineage) Get() *EntityLineage {
	return v.value
}

func (v *NullableEntityLineage) Set(val *EntityLineage) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityLineage) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityLineage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityLineage(val *EntityLineage) *NullableEntityLineage {
	return &NullableEntityLineage{value: val, isSet: true}
}

func (v NullableEntityLineage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityLineage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}