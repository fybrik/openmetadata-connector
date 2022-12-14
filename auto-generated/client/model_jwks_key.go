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

// JWKSKey struct for JWKSKey
type JWKSKey struct {
	Alg string `json:"alg"`
	E   string `json:"e"`
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	N   string `json:"n"`
	Use string `json:"use"`
}

// NewJWKSKey instantiates a new JWKSKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJWKSKey(alg string, e string, kid string, kty string, n string, use string) *JWKSKey {
	this := JWKSKey{}
	this.Alg = alg
	this.E = e
	this.Kid = kid
	this.Kty = kty
	this.N = n
	this.Use = use
	return &this
}

// NewJWKSKeyWithDefaults instantiates a new JWKSKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJWKSKeyWithDefaults() *JWKSKey {
	this := JWKSKey{}
	return &this
}

// GetAlg returns the Alg field value
func (o *JWKSKey) GetAlg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alg
}

// GetAlgOk returns a tuple with the Alg field value
// and a boolean to check if the value has been set.
func (o *JWKSKey) GetAlgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alg, true
}

// SetAlg sets field value
func (o *JWKSKey) SetAlg(v string) {
	o.Alg = v
}

// GetE returns the E field value
func (o *JWKSKey) GetE() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.E
}

// GetEOk returns a tuple with the E field value
// and a boolean to check if the value has been set.
func (o *JWKSKey) GetEOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.E, true
}

// SetE sets field value
func (o *JWKSKey) SetE(v string) {
	o.E = v
}

// GetKid returns the Kid field value
func (o *JWKSKey) GetKid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kid
}

// GetKidOk returns a tuple with the Kid field value
// and a boolean to check if the value has been set.
func (o *JWKSKey) GetKidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kid, true
}

// SetKid sets field value
func (o *JWKSKey) SetKid(v string) {
	o.Kid = v
}

// GetKty returns the Kty field value
func (o *JWKSKey) GetKty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kty
}

// GetKtyOk returns a tuple with the Kty field value
// and a boolean to check if the value has been set.
func (o *JWKSKey) GetKtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kty, true
}

// SetKty sets field value
func (o *JWKSKey) SetKty(v string) {
	o.Kty = v
}

// GetN returns the N field value
func (o *JWKSKey) GetN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.N
}

// GetNOk returns a tuple with the N field value
// and a boolean to check if the value has been set.
func (o *JWKSKey) GetNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.N, true
}

// SetN sets field value
func (o *JWKSKey) SetN(v string) {
	o.N = v
}

// GetUse returns the Use field value
func (o *JWKSKey) GetUse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Use
}

// GetUseOk returns a tuple with the Use field value
// and a boolean to check if the value has been set.
func (o *JWKSKey) GetUseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Use, true
}

// SetUse sets field value
func (o *JWKSKey) SetUse(v string) {
	o.Use = v
}

func (o JWKSKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alg"] = o.Alg
	}
	if true {
		toSerialize["e"] = o.E
	}
	if true {
		toSerialize["kid"] = o.Kid
	}
	if true {
		toSerialize["kty"] = o.Kty
	}
	if true {
		toSerialize["n"] = o.N
	}
	if true {
		toSerialize["use"] = o.Use
	}
	return json.Marshal(toSerialize)
}

type NullableJWKSKey struct {
	value *JWKSKey
	isSet bool
}

func (v NullableJWKSKey) Get() *JWKSKey {
	return v.value
}

func (v *NullableJWKSKey) Set(val *JWKSKey) {
	v.value = val
	v.isSet = true
}

func (v NullableJWKSKey) IsSet() bool {
	return v.isSet
}

func (v *NullableJWKSKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJWKSKey(val *JWKSKey) *NullableJWKSKey {
	return &NullableJWKSKey{value: val, isSet: true}
}

func (v NullableJWKSKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJWKSKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
