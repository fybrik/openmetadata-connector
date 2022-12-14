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

// Thread struct for Thread
type Thread struct {
	About       string       `json:"about"`
	AddressedTo *string      `json:"addressedTo,omitempty"`
	CreatedBy   *string      `json:"createdBy,omitempty"`
	EntityId    *string      `json:"entityId,omitempty"`
	Href        *string      `json:"href,omitempty"`
	Id          string       `json:"id"`
	Message     string       `json:"message"`
	Posts       []Post       `json:"posts,omitempty"`
	PostsCount  *int32       `json:"postsCount,omitempty"`
	Reactions   []Reaction   `json:"reactions,omitempty"`
	Resolved    *bool        `json:"resolved,omitempty"`
	Task        *TaskDetails `json:"task,omitempty"`
	ThreadTs    *int64       `json:"threadTs,omitempty"`
	Type        *string      `json:"type,omitempty"`
	UpdatedAt   *int64       `json:"updatedAt,omitempty"`
	UpdatedBy   *string      `json:"updatedBy,omitempty"`
}

// NewThread instantiates a new Thread object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThread(about string, id string, message string) *Thread {
	this := Thread{}
	this.About = about
	this.Id = id
	this.Message = message
	return &this
}

// NewThreadWithDefaults instantiates a new Thread object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreadWithDefaults() *Thread {
	this := Thread{}
	return &this
}

// GetAbout returns the About field value
func (o *Thread) GetAbout() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.About
}

// GetAboutOk returns a tuple with the About field value
// and a boolean to check if the value has been set.
func (o *Thread) GetAboutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.About, true
}

// SetAbout sets field value
func (o *Thread) SetAbout(v string) {
	o.About = v
}

// GetAddressedTo returns the AddressedTo field value if set, zero value otherwise.
func (o *Thread) GetAddressedTo() string {
	if o == nil || o.AddressedTo == nil {
		var ret string
		return ret
	}
	return *o.AddressedTo
}

// GetAddressedToOk returns a tuple with the AddressedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Thread) GetAddressedToOk() (*string, bool) {
	if o == nil || o.AddressedTo == nil {
		return nil, false
	}
	return o.AddressedTo, true
}

// HasAddressedTo returns a boolean if a field has been set.
func (o *Thread) HasAddressedTo() bool {
	if o != nil && o.AddressedTo != nil {
		return true
	}

	return false
}

// SetAddressedTo gets a reference to the given string and assigns it to the AddressedTo field.
func (o *Thread) SetAddressedTo(v string) {
	o.AddressedTo = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Thread) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Thread) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Thread) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Thread) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *Thread) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Thread) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *Thread) HasEntityId() bool {
	if o != nil && o.EntityId != nil {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *Thread) SetEntityId(v string) {
	o.EntityId = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Thread) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Thread) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Thread) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Thread) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value
func (o *Thread) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Thread) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Thread) SetId(v string) {
	o.Id = v
}

// GetMessage returns the Message field value
func (o *Thread) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Thread) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *Thread) SetMessage(v string) {
	o.Message = v
}

// GetPosts returns the Posts field value if set, zero value otherwise.
func (o *Thread) GetPosts() []Post {
	if o == nil || o.Posts == nil {
		var ret []Post
		return ret
	}
	return o.Posts
}

// GetPostsOk returns a tuple with the Posts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Thread) GetPostsOk() ([]Post, bool) {
	if o == nil || o.Posts == nil {
		return nil, false
	}
	return o.Posts, true
}

// HasPosts returns a boolean if a field has been set.
func (o *Thread) HasPosts() bool {
	if o != nil && o.Posts != nil {
		return true
	}

	return false
}

// SetPosts gets a reference to the given []Post and assigns it to the Posts field.
func (o *Thread) SetPosts(v []Post) {
	o.Posts = v
}

// GetPostsCount returns the PostsCount field value if set, zero value otherwise.
func (o *Thread) GetPostsCount() int32 {
	if o == nil || o.PostsCount == nil {
		var ret int32
		return ret
	}
	return *o.PostsCount
}

// GetPostsCountOk returns a tuple with the PostsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Thread) GetPostsCountOk() (*int32, bool) {
	if o == nil || o.PostsCount == nil {
		return nil, false
	}
	return o.PostsCount, true
}

// HasPostsCount returns a boolean if a field has been set.
func (o *Thread) HasPostsCount() bool {
	if o != nil && o.PostsCount != nil {
		return true
	}

	return false
}

// SetPostsCount gets a reference to the given int32 and assigns it to the PostsCount field.
func (o *Thread) SetPostsCount(v int32) {
	o.PostsCount = &v
}

// GetReactions returns the Reactions field value if set, zero value otherwise.
func (o *Thread) GetReactions() []Reaction {
	if o == nil || o.Reactions == nil {
		var ret []Reaction
		return ret
	}
	return o.Reactions
}

// GetReactionsOk returns a tuple with the Reactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Thread) GetReactionsOk() ([]Reaction, bool) {
	if o == nil || o.Reactions == nil {
		return nil, false
	}
	return o.Reactions, true
}

// HasReactions returns a boolean if a field has been set.
func (o *Thread) HasReactions() bool {
	if o != nil && o.Reactions != nil {
		return true
	}

	return false
}

// SetReactions gets a reference to the given []Reaction and assigns it to the Reactions field.
func (o *Thread) SetReactions(v []Reaction) {
	o.Reactions = v
}

// GetResolved returns the Resolved field value if set, zero value otherwise.
func (o *Thread) GetResolved() bool {
	if o == nil || o.Resolved == nil {
		var ret bool
		return ret
	}
	return *o.Resolved
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Thread) GetResolvedOk() (*bool, bool) {
	if o == nil || o.Resolved == nil {
		return nil, false
	}
	return o.Resolved, true
}

// HasResolved returns a boolean if a field has been set.
func (o *Thread) HasResolved() bool {
	if o != nil && o.Resolved != nil {
		return true
	}

	return false
}

// SetResolved gets a reference to the given bool and assigns it to the Resolved field.
func (o *Thread) SetResolved(v bool) {
	o.Resolved = &v
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *Thread) GetTask() TaskDetails {
	if o == nil || o.Task == nil {
		var ret TaskDetails
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Thread) GetTaskOk() (*TaskDetails, bool) {
	if o == nil || o.Task == nil {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *Thread) HasTask() bool {
	if o != nil && o.Task != nil {
		return true
	}

	return false
}

// SetTask gets a reference to the given TaskDetails and assigns it to the Task field.
func (o *Thread) SetTask(v TaskDetails) {
	o.Task = &v
}

// GetThreadTs returns the ThreadTs field value if set, zero value otherwise.
func (o *Thread) GetThreadTs() int64 {
	if o == nil || o.ThreadTs == nil {
		var ret int64
		return ret
	}
	return *o.ThreadTs
}

// GetThreadTsOk returns a tuple with the ThreadTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Thread) GetThreadTsOk() (*int64, bool) {
	if o == nil || o.ThreadTs == nil {
		return nil, false
	}
	return o.ThreadTs, true
}

// HasThreadTs returns a boolean if a field has been set.
func (o *Thread) HasThreadTs() bool {
	if o != nil && o.ThreadTs != nil {
		return true
	}

	return false
}

// SetThreadTs gets a reference to the given int64 and assigns it to the ThreadTs field.
func (o *Thread) SetThreadTs(v int64) {
	o.ThreadTs = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Thread) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Thread) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Thread) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Thread) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Thread) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Thread) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Thread) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *Thread) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *Thread) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Thread) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *Thread) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *Thread) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

func (o Thread) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["about"] = o.About
	}
	if o.AddressedTo != nil {
		toSerialize["addressedTo"] = o.AddressedTo
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.EntityId != nil {
		toSerialize["entityId"] = o.EntityId
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if o.Posts != nil {
		toSerialize["posts"] = o.Posts
	}
	if o.PostsCount != nil {
		toSerialize["postsCount"] = o.PostsCount
	}
	if o.Reactions != nil {
		toSerialize["reactions"] = o.Reactions
	}
	if o.Resolved != nil {
		toSerialize["resolved"] = o.Resolved
	}
	if o.Task != nil {
		toSerialize["task"] = o.Task
	}
	if o.ThreadTs != nil {
		toSerialize["threadTs"] = o.ThreadTs
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	return json.Marshal(toSerialize)
}

type NullableThread struct {
	value *Thread
	isSet bool
}

func (v NullableThread) Get() *Thread {
	return v.value
}

func (v *NullableThread) Set(val *Thread) {
	v.value = val
	v.isSet = true
}

func (v NullableThread) IsSet() bool {
	return v.isSet
}

func (v *NullableThread) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThread(val *Thread) *NullableThread {
	return &NullableThread{value: val, isSet: true}
}

func (v NullableThread) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThread) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
