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

// ResultListUser struct for ResultListUser
type ResultListUser struct {
	Data   []User  `json:"data"`
	Paging *Paging `json:"paging,omitempty"`
}

// NewResultListUser instantiates a new ResultListUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultListUser(data []User) *ResultListUser {
	this := ResultListUser{}
	this.Data = data
	return &this
}

// NewResultListUserWithDefaults instantiates a new ResultListUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultListUserWithDefaults() *ResultListUser {
	this := ResultListUser{}
	return &this
}

// GetData returns the Data field value
func (o *ResultListUser) GetData() []User {
	if o == nil {
		var ret []User
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResultListUser) GetDataOk() ([]User, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ResultListUser) SetData(v []User) {
	o.Data = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *ResultListUser) GetPaging() Paging {
	if o == nil || o.Paging == nil {
		var ret Paging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultListUser) GetPagingOk() (*Paging, bool) {
	if o == nil || o.Paging == nil {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *ResultListUser) HasPaging() bool {
	if o != nil && o.Paging != nil {
		return true
	}

	return false
}

// SetPaging gets a reference to the given Paging and assigns it to the Paging field.
func (o *ResultListUser) SetPaging(v Paging) {
	o.Paging = &v
}

func (o ResultListUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Paging != nil {
		toSerialize["paging"] = o.Paging
	}
	return json.Marshal(toSerialize)
}

type NullableResultListUser struct {
	value *ResultListUser
	isSet bool
}

func (v NullableResultListUser) Get() *ResultListUser {
	return v.value
}

func (v *NullableResultListUser) Set(val *ResultListUser) {
	v.value = val
	v.isSet = true
}

func (v NullableResultListUser) IsSet() bool {
	return v.isSet
}

func (v *NullableResultListUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultListUser(val *ResultListUser) *NullableResultListUser {
	return &NullableResultListUser{value: val, isSet: true}
}

func (v NullableResultListUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultListUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
