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

// CreateMessagingService struct for CreateMessagingService
type CreateMessagingService struct {
	Connection  MessagingConnection    `json:"connection"`
	Description *string                `json:"description,omitempty"`
	DisplayName *string                `json:"displayName,omitempty"`
	Extension   map[string]interface{} `json:"extension,omitempty"`
	Name        string                 `json:"name"`
	Owner       *EntityReference       `json:"owner,omitempty"`
	ServiceType string                 `json:"serviceType"`
}

// NewCreateMessagingService instantiates a new CreateMessagingService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMessagingService(connection MessagingConnection, name string, serviceType string) *CreateMessagingService {
	this := CreateMessagingService{}
	this.Connection = connection
	this.Name = name
	this.ServiceType = serviceType
	return &this
}

// NewCreateMessagingServiceWithDefaults instantiates a new CreateMessagingService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMessagingServiceWithDefaults() *CreateMessagingService {
	this := CreateMessagingService{}
	return &this
}

// GetConnection returns the Connection field value
func (o *CreateMessagingService) GetConnection() MessagingConnection {
	if o == nil {
		var ret MessagingConnection
		return ret
	}

	return o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value
// and a boolean to check if the value has been set.
func (o *CreateMessagingService) GetConnectionOk() (*MessagingConnection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connection, true
}

// SetConnection sets field value
func (o *CreateMessagingService) SetConnection(v MessagingConnection) {
	o.Connection = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateMessagingService) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessagingService) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateMessagingService) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateMessagingService) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CreateMessagingService) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessagingService) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateMessagingService) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CreateMessagingService) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *CreateMessagingService) GetExtension() map[string]interface{} {
	if o == nil || o.Extension == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessagingService) GetExtensionOk() (map[string]interface{}, bool) {
	if o == nil || o.Extension == nil {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *CreateMessagingService) HasExtension() bool {
	if o != nil && o.Extension != nil {
		return true
	}

	return false
}

// SetExtension gets a reference to the given map[string]interface{} and assigns it to the Extension field.
func (o *CreateMessagingService) SetExtension(v map[string]interface{}) {
	o.Extension = v
}

// GetName returns the Name field value
func (o *CreateMessagingService) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateMessagingService) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateMessagingService) SetName(v string) {
	o.Name = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *CreateMessagingService) GetOwner() EntityReference {
	if o == nil || o.Owner == nil {
		var ret EntityReference
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessagingService) GetOwnerOk() (*EntityReference, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *CreateMessagingService) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given EntityReference and assigns it to the Owner field.
func (o *CreateMessagingService) SetOwner(v EntityReference) {
	o.Owner = &v
}

// GetServiceType returns the ServiceType field value
func (o *CreateMessagingService) GetServiceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value
// and a boolean to check if the value has been set.
func (o *CreateMessagingService) GetServiceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceType, true
}

// SetServiceType sets field value
func (o *CreateMessagingService) SetServiceType(v string) {
	o.ServiceType = v
}

func (o CreateMessagingService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["connection"] = o.Connection
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Extension != nil {
		toSerialize["extension"] = o.Extension
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if true {
		toSerialize["serviceType"] = o.ServiceType
	}
	return json.Marshal(toSerialize)
}

type NullableCreateMessagingService struct {
	value *CreateMessagingService
	isSet bool
}

func (v NullableCreateMessagingService) Get() *CreateMessagingService {
	return v.value
}

func (v *NullableCreateMessagingService) Set(val *CreateMessagingService) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMessagingService) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMessagingService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMessagingService(val *CreateMessagingService) *NullableCreateMessagingService {
	return &NullableCreateMessagingService{value: val, isSet: true}
}

func (v NullableCreateMessagingService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMessagingService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
