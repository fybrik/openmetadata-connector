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

// CollectionDescriptor struct for CollectionDescriptor
type CollectionDescriptor struct {
	Collection *CollectionInfo `json:"collection,omitempty"`
}

// NewCollectionDescriptor instantiates a new CollectionDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionDescriptor() *CollectionDescriptor {
	this := CollectionDescriptor{}
	return &this
}

// NewCollectionDescriptorWithDefaults instantiates a new CollectionDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionDescriptorWithDefaults() *CollectionDescriptor {
	this := CollectionDescriptor{}
	return &this
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *CollectionDescriptor) GetCollection() CollectionInfo {
	if o == nil || o.Collection == nil {
		var ret CollectionInfo
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionDescriptor) GetCollectionOk() (*CollectionInfo, bool) {
	if o == nil || o.Collection == nil {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *CollectionDescriptor) HasCollection() bool {
	if o != nil && o.Collection != nil {
		return true
	}

	return false
}

// SetCollection gets a reference to the given CollectionInfo and assigns it to the Collection field.
func (o *CollectionDescriptor) SetCollection(v CollectionInfo) {
	o.Collection = &v
}

func (o CollectionDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Collection != nil {
		toSerialize["collection"] = o.Collection
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionDescriptor struct {
	value *CollectionDescriptor
	isSet bool
}

func (v NullableCollectionDescriptor) Get() *CollectionDescriptor {
	return v.value
}

func (v *NullableCollectionDescriptor) Set(val *CollectionDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionDescriptor(val *CollectionDescriptor) *NullableCollectionDescriptor {
	return &NullableCollectionDescriptor{value: val, isSet: true}
}

func (v NullableCollectionDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
