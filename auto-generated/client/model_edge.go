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

// Edge struct for Edge
type Edge struct {
	Description    *string         `json:"description,omitempty"`
	FromEntity     string          `json:"fromEntity"`
	LineageDetails *LineageDetails `json:"lineageDetails,omitempty"`
	ToEntity       string          `json:"toEntity"`
}

// NewEdge instantiates a new Edge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdge(fromEntity string, toEntity string) *Edge {
	this := Edge{}
	this.FromEntity = fromEntity
	this.ToEntity = toEntity
	return &this
}

// NewEdgeWithDefaults instantiates a new Edge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeWithDefaults() *Edge {
	this := Edge{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Edge) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Edge) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Edge) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Edge) SetDescription(v string) {
	o.Description = &v
}

// GetFromEntity returns the FromEntity field value
func (o *Edge) GetFromEntity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromEntity
}

// GetFromEntityOk returns a tuple with the FromEntity field value
// and a boolean to check if the value has been set.
func (o *Edge) GetFromEntityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromEntity, true
}

// SetFromEntity sets field value
func (o *Edge) SetFromEntity(v string) {
	o.FromEntity = v
}

// GetLineageDetails returns the LineageDetails field value if set, zero value otherwise.
func (o *Edge) GetLineageDetails() LineageDetails {
	if o == nil || o.LineageDetails == nil {
		var ret LineageDetails
		return ret
	}
	return *o.LineageDetails
}

// GetLineageDetailsOk returns a tuple with the LineageDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Edge) GetLineageDetailsOk() (*LineageDetails, bool) {
	if o == nil || o.LineageDetails == nil {
		return nil, false
	}
	return o.LineageDetails, true
}

// HasLineageDetails returns a boolean if a field has been set.
func (o *Edge) HasLineageDetails() bool {
	if o != nil && o.LineageDetails != nil {
		return true
	}

	return false
}

// SetLineageDetails gets a reference to the given LineageDetails and assigns it to the LineageDetails field.
func (o *Edge) SetLineageDetails(v LineageDetails) {
	o.LineageDetails = &v
}

// GetToEntity returns the ToEntity field value
func (o *Edge) GetToEntity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToEntity
}

// GetToEntityOk returns a tuple with the ToEntity field value
// and a boolean to check if the value has been set.
func (o *Edge) GetToEntityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToEntity, true
}

// SetToEntity sets field value
func (o *Edge) SetToEntity(v string) {
	o.ToEntity = v
}

func (o Edge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["fromEntity"] = o.FromEntity
	}
	if o.LineageDetails != nil {
		toSerialize["lineageDetails"] = o.LineageDetails
	}
	if true {
		toSerialize["toEntity"] = o.ToEntity
	}
	return json.Marshal(toSerialize)
}

type NullableEdge struct {
	value *Edge
	isSet bool
}

func (v NullableEdge) Get() *Edge {
	return v.value
}

func (v *NullableEdge) Set(val *Edge) {
	v.value = val
	v.isSet = true
}

func (v NullableEdge) IsSet() bool {
	return v.isSet
}

func (v *NullableEdge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdge(val *Edge) *NullableEdge {
	return &NullableEdge{value: val, isSet: true}
}

func (v NullableEdge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}