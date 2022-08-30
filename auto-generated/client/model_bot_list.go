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

// BotList struct for BotList
type BotList struct {
	Data   []Bot   `json:"data"`
	Paging *Paging `json:"paging,omitempty"`
}

// NewBotList instantiates a new BotList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotList(data []Bot) *BotList {
	this := BotList{}
	this.Data = data
	return &this
}

// NewBotListWithDefaults instantiates a new BotList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotListWithDefaults() *BotList {
	this := BotList{}
	return &this
}

// GetData returns the Data field value
func (o *BotList) GetData() []Bot {
	if o == nil {
		var ret []Bot
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BotList) GetDataOk() ([]Bot, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BotList) SetData(v []Bot) {
	o.Data = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *BotList) GetPaging() Paging {
	if o == nil || o.Paging == nil {
		var ret Paging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotList) GetPagingOk() (*Paging, bool) {
	if o == nil || o.Paging == nil {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *BotList) HasPaging() bool {
	if o != nil && o.Paging != nil {
		return true
	}

	return false
}

// SetPaging gets a reference to the given Paging and assigns it to the Paging field.
func (o *BotList) SetPaging(v Paging) {
	o.Paging = &v
}

func (o BotList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Paging != nil {
		toSerialize["paging"] = o.Paging
	}
	return json.Marshal(toSerialize)
}

type NullableBotList struct {
	value *BotList
	isSet bool
}

func (v NullableBotList) Get() *BotList {
	return v.value
}

func (v *NullableBotList) Set(val *BotList) {
	v.value = val
	v.isSet = true
}

func (v NullableBotList) IsSet() bool {
	return v.isSet
}

func (v *NullableBotList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBotList(val *BotList) *NullableBotList {
	return &NullableBotList{value: val, isSet: true}
}

func (v NullableBotList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBotList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
