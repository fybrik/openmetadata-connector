# GlossaryTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeDescription** | Pointer to [**ChangeDescription**](ChangeDescription.md) |  | [optional] 
**Children** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Description** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **map[string]interface{}** |  | [optional] 
**Followers** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**FullyQualifiedName** | Pointer to **string** |  | [optional] 
**Glossary** | [**EntityReference**](EntityReference.md) |  | 
**Href** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Parent** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**References** | Pointer to [**[]TermReference**](TermReference.md) |  | [optional] 
**RelatedTerms** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**Reviewers** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Synonyms** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to [**[]TagLabel**](TagLabel.md) |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**UsageCount** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **float64** |  | [optional] 

## Methods

### NewGlossaryTerm

`func NewGlossaryTerm(description string, glossary EntityReference, id string, name string, ) *GlossaryTerm`

NewGlossaryTerm instantiates a new GlossaryTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlossaryTermWithDefaults

`func NewGlossaryTermWithDefaults() *GlossaryTerm`

NewGlossaryTermWithDefaults instantiates a new GlossaryTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeDescription

`func (o *GlossaryTerm) GetChangeDescription() ChangeDescription`

GetChangeDescription returns the ChangeDescription field if non-nil, zero value otherwise.

### GetChangeDescriptionOk

`func (o *GlossaryTerm) GetChangeDescriptionOk() (*ChangeDescription, bool)`

GetChangeDescriptionOk returns a tuple with the ChangeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDescription

`func (o *GlossaryTerm) SetChangeDescription(v ChangeDescription)`

SetChangeDescription sets ChangeDescription field to given value.

### HasChangeDescription

`func (o *GlossaryTerm) HasChangeDescription() bool`

HasChangeDescription returns a boolean if a field has been set.

### GetChildren

`func (o *GlossaryTerm) GetChildren() []EntityReference`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *GlossaryTerm) GetChildrenOk() (*[]EntityReference, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *GlossaryTerm) SetChildren(v []EntityReference)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *GlossaryTerm) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetDeleted

`func (o *GlossaryTerm) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *GlossaryTerm) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *GlossaryTerm) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *GlossaryTerm) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDescription

`func (o *GlossaryTerm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GlossaryTerm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GlossaryTerm) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayName

`func (o *GlossaryTerm) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GlossaryTerm) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GlossaryTerm) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GlossaryTerm) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtension

`func (o *GlossaryTerm) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *GlossaryTerm) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *GlossaryTerm) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *GlossaryTerm) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetFollowers

`func (o *GlossaryTerm) GetFollowers() []EntityReference`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *GlossaryTerm) GetFollowersOk() (*[]EntityReference, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *GlossaryTerm) SetFollowers(v []EntityReference)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *GlossaryTerm) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetFullyQualifiedName

`func (o *GlossaryTerm) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *GlossaryTerm) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *GlossaryTerm) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.

### HasFullyQualifiedName

`func (o *GlossaryTerm) HasFullyQualifiedName() bool`

HasFullyQualifiedName returns a boolean if a field has been set.

### GetGlossary

`func (o *GlossaryTerm) GetGlossary() EntityReference`

GetGlossary returns the Glossary field if non-nil, zero value otherwise.

### GetGlossaryOk

`func (o *GlossaryTerm) GetGlossaryOk() (*EntityReference, bool)`

GetGlossaryOk returns a tuple with the Glossary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlossary

`func (o *GlossaryTerm) SetGlossary(v EntityReference)`

SetGlossary sets Glossary field to given value.


### GetHref

`func (o *GlossaryTerm) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *GlossaryTerm) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *GlossaryTerm) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *GlossaryTerm) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *GlossaryTerm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlossaryTerm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlossaryTerm) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GlossaryTerm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlossaryTerm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlossaryTerm) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *GlossaryTerm) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GlossaryTerm) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GlossaryTerm) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GlossaryTerm) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParent

`func (o *GlossaryTerm) GetParent() EntityReference`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GlossaryTerm) GetParentOk() (*EntityReference, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GlossaryTerm) SetParent(v EntityReference)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GlossaryTerm) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetReferences

`func (o *GlossaryTerm) GetReferences() []TermReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *GlossaryTerm) GetReferencesOk() (*[]TermReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *GlossaryTerm) SetReferences(v []TermReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *GlossaryTerm) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelatedTerms

`func (o *GlossaryTerm) GetRelatedTerms() []EntityReference`

GetRelatedTerms returns the RelatedTerms field if non-nil, zero value otherwise.

### GetRelatedTermsOk

`func (o *GlossaryTerm) GetRelatedTermsOk() (*[]EntityReference, bool)`

GetRelatedTermsOk returns a tuple with the RelatedTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedTerms

`func (o *GlossaryTerm) SetRelatedTerms(v []EntityReference)`

SetRelatedTerms sets RelatedTerms field to given value.

### HasRelatedTerms

`func (o *GlossaryTerm) HasRelatedTerms() bool`

HasRelatedTerms returns a boolean if a field has been set.

### GetReviewers

`func (o *GlossaryTerm) GetReviewers() []EntityReference`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *GlossaryTerm) GetReviewersOk() (*[]EntityReference, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *GlossaryTerm) SetReviewers(v []EntityReference)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *GlossaryTerm) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetStatus

`func (o *GlossaryTerm) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GlossaryTerm) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GlossaryTerm) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GlossaryTerm) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSynonyms

`func (o *GlossaryTerm) GetSynonyms() []string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *GlossaryTerm) GetSynonymsOk() (*[]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *GlossaryTerm) SetSynonyms(v []string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *GlossaryTerm) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### GetTags

`func (o *GlossaryTerm) GetTags() []TagLabel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GlossaryTerm) GetTagsOk() (*[]TagLabel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GlossaryTerm) SetTags(v []TagLabel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GlossaryTerm) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GlossaryTerm) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GlossaryTerm) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GlossaryTerm) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GlossaryTerm) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *GlossaryTerm) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *GlossaryTerm) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *GlossaryTerm) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *GlossaryTerm) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUsageCount

`func (o *GlossaryTerm) GetUsageCount() int32`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *GlossaryTerm) GetUsageCountOk() (*int32, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *GlossaryTerm) SetUsageCount(v int32)`

SetUsageCount sets UsageCount field to given value.

### HasUsageCount

`func (o *GlossaryTerm) HasUsageCount() bool`

HasUsageCount returns a boolean if a field has been set.

### GetVersion

`func (o *GlossaryTerm) GetVersion() float64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GlossaryTerm) GetVersionOk() (*float64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GlossaryTerm) SetVersion(v float64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GlossaryTerm) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


