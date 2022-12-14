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

// EntitiesEdge struct for EntitiesEdge
type EntitiesEdge struct {
	Description    *string         `json:"description,omitempty"`
	FromEntity     EntityReference `json:"fromEntity"`
	LineageDetails *LineageDetails `json:"lineageDetails,omitempty"`
	ToEntity       EntityReference `json:"toEntity"`
}

// NewEntitiesEdge instantiates a new EntitiesEdge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitiesEdge(fromEntity EntityReference, toEntity EntityReference) *EntitiesEdge {
	this := EntitiesEdge{}
	this.FromEntity = fromEntity
	this.ToEntity = toEntity
	return &this
}

// NewEntitiesEdgeWithDefaults instantiates a new EntitiesEdge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitiesEdgeWithDefaults() *EntitiesEdge {
	this := EntitiesEdge{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitiesEdge) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitiesEdge) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitiesEdge) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitiesEdge) SetDescription(v string) {
	o.Description = &v
}

// GetFromEntity returns the FromEntity field value
func (o *EntitiesEdge) GetFromEntity() EntityReference {
	if o == nil {
		var ret EntityReference
		return ret
	}

	return o.FromEntity
}

// GetFromEntityOk returns a tuple with the FromEntity field value
// and a boolean to check if the value has been set.
func (o *EntitiesEdge) GetFromEntityOk() (*EntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromEntity, true
}

// SetFromEntity sets field value
func (o *EntitiesEdge) SetFromEntity(v EntityReference) {
	o.FromEntity = v
}

// GetLineageDetails returns the LineageDetails field value if set, zero value otherwise.
func (o *EntitiesEdge) GetLineageDetails() LineageDetails {
	if o == nil || o.LineageDetails == nil {
		var ret LineageDetails
		return ret
	}
	return *o.LineageDetails
}

// GetLineageDetailsOk returns a tuple with the LineageDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitiesEdge) GetLineageDetailsOk() (*LineageDetails, bool) {
	if o == nil || o.LineageDetails == nil {
		return nil, false
	}
	return o.LineageDetails, true
}

// HasLineageDetails returns a boolean if a field has been set.
func (o *EntitiesEdge) HasLineageDetails() bool {
	if o != nil && o.LineageDetails != nil {
		return true
	}

	return false
}

// SetLineageDetails gets a reference to the given LineageDetails and assigns it to the LineageDetails field.
func (o *EntitiesEdge) SetLineageDetails(v LineageDetails) {
	o.LineageDetails = &v
}

// GetToEntity returns the ToEntity field value
func (o *EntitiesEdge) GetToEntity() EntityReference {
	if o == nil {
		var ret EntityReference
		return ret
	}

	return o.ToEntity
}

// GetToEntityOk returns a tuple with the ToEntity field value
// and a boolean to check if the value has been set.
func (o *EntitiesEdge) GetToEntityOk() (*EntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToEntity, true
}

// SetToEntity sets field value
func (o *EntitiesEdge) SetToEntity(v EntityReference) {
	o.ToEntity = v
}

func (o EntitiesEdge) MarshalJSON() ([]byte, error) {
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

type NullableEntitiesEdge struct {
	value *EntitiesEdge
	isSet bool
}

func (v NullableEntitiesEdge) Get() *EntitiesEdge {
	return v.value
}

func (v *NullableEntitiesEdge) Set(val *EntitiesEdge) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitiesEdge) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitiesEdge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitiesEdge(val *EntitiesEdge) *NullableEntitiesEdge {
	return &NullableEntitiesEdge{value: val, isSet: true}
}

func (v NullableEntitiesEdge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitiesEdge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
