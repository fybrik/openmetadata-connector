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

// Paging struct for Paging
type Paging struct {
	After  *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Total  int32   `json:"total"`
}

// NewPaging instantiates a new Paging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaging(total int32) *Paging {
	this := Paging{}
	this.Total = total
	return &this
}

// NewPagingWithDefaults instantiates a new Paging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagingWithDefaults() *Paging {
	this := Paging{}
	return &this
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *Paging) GetAfter() string {
	if o == nil || o.After == nil {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Paging) GetAfterOk() (*string, bool) {
	if o == nil || o.After == nil {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *Paging) HasAfter() bool {
	if o != nil && o.After != nil {
		return true
	}

	return false
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *Paging) SetAfter(v string) {
	o.After = &v
}

// GetBefore returns the Before field value if set, zero value otherwise.
func (o *Paging) GetBefore() string {
	if o == nil || o.Before == nil {
		var ret string
		return ret
	}
	return *o.Before
}

// GetBeforeOk returns a tuple with the Before field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Paging) GetBeforeOk() (*string, bool) {
	if o == nil || o.Before == nil {
		return nil, false
	}
	return o.Before, true
}

// HasBefore returns a boolean if a field has been set.
func (o *Paging) HasBefore() bool {
	if o != nil && o.Before != nil {
		return true
	}

	return false
}

// SetBefore gets a reference to the given string and assigns it to the Before field.
func (o *Paging) SetBefore(v string) {
	o.Before = &v
}

// GetTotal returns the Total field value
func (o *Paging) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *Paging) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *Paging) SetTotal(v int32) {
	o.Total = v
}

func (o Paging) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.After != nil {
		toSerialize["after"] = o.After
	}
	if o.Before != nil {
		toSerialize["before"] = o.Before
	}
	if true {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullablePaging struct {
	value *Paging
	isSet bool
}

func (v NullablePaging) Get() *Paging {
	return v.value
}

func (v *NullablePaging) Set(val *Paging) {
	v.value = val
	v.isSet = true
}

func (v NullablePaging) IsSet() bool {
	return v.isSet
}

func (v *NullablePaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaging(val *Paging) *NullablePaging {
	return &NullablePaging{value: val, isSet: true}
}

func (v NullablePaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
