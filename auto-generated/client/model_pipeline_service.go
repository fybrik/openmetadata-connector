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

// PipelineService struct for PipelineService
type PipelineService struct {
	ChangeDescription  *ChangeDescription     `json:"changeDescription,omitempty"`
	Connection         PipelineConnection     `json:"connection"`
	Deleted            *bool                  `json:"deleted,omitempty"`
	Description        *string                `json:"description,omitempty"`
	DisplayName        *string                `json:"displayName,omitempty"`
	Extension          map[string]interface{} `json:"extension,omitempty"`
	Followers          []EntityReference      `json:"followers,omitempty"`
	FullyQualifiedName *string                `json:"fullyQualifiedName,omitempty"`
	Href               *string                `json:"href,omitempty"`
	Id                 string                 `json:"id"`
	Name               string                 `json:"name"`
	Owner              *EntityReference       `json:"owner,omitempty"`
	Pipelines          []EntityReference      `json:"pipelines,omitempty"`
	ServiceType        string                 `json:"serviceType"`
	Tags               []TagLabel             `json:"tags,omitempty"`
	UpdatedAt          *int64                 `json:"updatedAt,omitempty"`
	UpdatedBy          *string                `json:"updatedBy,omitempty"`
	Version            *float64               `json:"version,omitempty"`
}

// NewPipelineService instantiates a new PipelineService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineService(connection PipelineConnection, id string, name string, serviceType string) *PipelineService {
	this := PipelineService{}
	this.Connection = connection
	this.Id = id
	this.Name = name
	this.ServiceType = serviceType
	return &this
}

// NewPipelineServiceWithDefaults instantiates a new PipelineService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineServiceWithDefaults() *PipelineService {
	this := PipelineService{}
	return &this
}

// GetChangeDescription returns the ChangeDescription field value if set, zero value otherwise.
func (o *PipelineService) GetChangeDescription() ChangeDescription {
	if o == nil || o.ChangeDescription == nil {
		var ret ChangeDescription
		return ret
	}
	return *o.ChangeDescription
}

// GetChangeDescriptionOk returns a tuple with the ChangeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineService) GetChangeDescriptionOk() (*ChangeDescription, bool) {
	if o == nil || o.ChangeDescription == nil {
		return nil, false
	}
	return o.ChangeDescription, true
}

// HasChangeDescription returns a boolean if a field has been set.
func (o *PipelineService) HasChangeDescription() bool {
	if o != nil && o.ChangeDescription != nil {
		return true
	}

	return false
}

// SetChangeDescription gets a reference to the given ChangeDescription and assigns it to the ChangeDescription field.
func (o *PipelineService) SetChangeDescription(v ChangeDescription) {
	o.ChangeDescription = &v
}

// GetConnection returns the Connection field value
func (o *PipelineService) GetConnection() PipelineConnection {
	if o == nil {
		var ret PipelineConnection
		return ret
	}

	return o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value
// and a boolean to check if the value has been set.
func (o *PipelineService) GetConnectionOk() (*PipelineConnection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connection, true
}

// SetConnection sets field value
func (o *PipelineService) SetConnection(v PipelineConnection) {
	o.Connection = v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *PipelineService) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineService) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *PipelineService) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *PipelineService) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PipelineService) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineService) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PipelineService) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PipelineService) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PipelineService) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineService) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PipelineService) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PipelineService) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *PipelineService) GetExtension() map[string]interface{} {
	if o == nil || o.Extension == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineService) GetExtensionOk() (map[string]interface{}, bool) {
	if o == nil || o.Extension == nil {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *PipelineService) HasExtension() bool {
	if o != nil && o.Extension != nil {
		return true
	}

	return false
}

// SetExtension gets a reference to the given map[string]interface{} and assigns it to the Extension field.
func (o *PipelineService) SetExtension(v map[string]interface{}) {
	o.Extension = v
}

// GetFollowers returns the Followers field value if set, zero value otherwise.
func (o *PipelineService) GetFollowers() []EntityReference {
	if o == nil || o.Followers == nil {
		var ret []EntityReference
		return ret
	}
	return o.Followers
}

// GetFollowersOk returns a tuple with the Followers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineService) GetFollowersOk() ([]EntityReference, bool) {
	if o == nil || o.Followers == nil {
		return nil, false
	}
	return o.Followers, true
}

// HasFollowers returns a boolean if a field has been set.
func (o *PipelineService) HasFollowers() bool {
	if o != nil && o.Followers != nil {
		return true
	}

	return false
}

// SetFollowers gets a reference to the given []EntityReference and assigns it to the Followers field.
func (o *PipelineService) SetFollowers(v []EntityReference) {
	o.Followers = v
}

// GetFullyQualifiedName returns the FullyQualifiedName field value if set, zero value otherwise.
func (o *PipelineService) GetFullyQualifiedName() string {
	if o == nil || o.FullyQualifiedName == nil {
		var ret string
		return ret
	}
	return *o.FullyQualifiedName
}

// GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineService) GetFullyQualifiedNameOk() (*string, bool) {
	if o == nil || o.FullyQualifiedName == nil {
		return nil, false
	}
	return o.FullyQualifiedName, true
}

// HasFullyQualifiedName returns a boolean if a field has been set.
func (o *PipelineService) HasFullyQualifiedName() bool {
	if o != nil && o.FullyQualifiedName != nil {
		return true
	}

	return false
}

// SetFullyQualifiedName gets a reference to the given string and assigns it to the FullyQualifiedName field.
func (o *PipelineService) SetFullyQualifiedName(v string) {
	o.FullyQualifiedName = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *PipelineService) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineService) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *PipelineService) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *PipelineService) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value
func (o *PipelineService) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PipelineService) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PipelineService) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PipelineService) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PipelineService) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PipelineService) SetName(v string) {
	o.Name = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *PipelineService) GetOwner() EntityReference {
	if o == nil || o.Owner == nil {
		var ret EntityReference
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineService) GetOwnerOk() (*EntityReference, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *PipelineService) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given EntityReference and assigns it to the Owner field.
func (o *PipelineService) SetOwner(v EntityReference) {
	o.Owner = &v
}

// GetPipelines returns the Pipelines field value if set, zero value otherwise.
func (o *PipelineService) GetPipelines() []EntityReference {
	if o == nil || o.Pipelines == nil {
		var ret []EntityReference
		return ret
	}
	return o.Pipelines
}

// GetPipelinesOk returns a tuple with the Pipelines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineService) GetPipelinesOk() ([]EntityReference, bool) {
	if o == nil || o.Pipelines == nil {
		return nil, false
	}
	return o.Pipelines, true
}

// HasPipelines returns a boolean if a field has been set.
func (o *PipelineService) HasPipelines() bool {
	if o != nil && o.Pipelines != nil {
		return true
	}

	return false
}

// SetPipelines gets a reference to the given []EntityReference and assigns it to the Pipelines field.
func (o *PipelineService) SetPipelines(v []EntityReference) {
	o.Pipelines = v
}

// GetServiceType returns the ServiceType field value
func (o *PipelineService) GetServiceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value
// and a boolean to check if the value has been set.
func (o *PipelineService) GetServiceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceType, true
}

// SetServiceType sets field value
func (o *PipelineService) SetServiceType(v string) {
	o.ServiceType = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PipelineService) GetTags() []TagLabel {
	if o == nil || o.Tags == nil {
		var ret []TagLabel
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineService) GetTagsOk() ([]TagLabel, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PipelineService) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagLabel and assigns it to the Tags field.
func (o *PipelineService) SetTags(v []TagLabel) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PipelineService) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineService) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PipelineService) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *PipelineService) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *PipelineService) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineService) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *PipelineService) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *PipelineService) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PipelineService) GetVersion() float64 {
	if o == nil || o.Version == nil {
		var ret float64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineService) GetVersionOk() (*float64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PipelineService) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given float64 and assigns it to the Version field.
func (o *PipelineService) SetVersion(v float64) {
	o.Version = &v
}

func (o PipelineService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangeDescription != nil {
		toSerialize["changeDescription"] = o.ChangeDescription
	}
	if true {
		toSerialize["connection"] = o.Connection
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
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
	if o.Followers != nil {
		toSerialize["followers"] = o.Followers
	}
	if o.FullyQualifiedName != nil {
		toSerialize["fullyQualifiedName"] = o.FullyQualifiedName
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Pipelines != nil {
		toSerialize["pipelines"] = o.Pipelines
	}
	if true {
		toSerialize["serviceType"] = o.ServiceType
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineService struct {
	value *PipelineService
	isSet bool
}

func (v NullablePipelineService) Get() *PipelineService {
	return v.value
}

func (v *NullablePipelineService) Set(val *PipelineService) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineService) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineService(val *PipelineService) *NullablePipelineService {
	return &NullablePipelineService{value: val, isSet: true}
}

func (v NullablePipelineService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
