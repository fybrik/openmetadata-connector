# CreateGlossaryTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **map[string]interface{}** |  | [optional] 
**Glossary** | [**EntityReference**](EntityReference.md) |  | 
**Name** | **string** |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Parent** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**References** | Pointer to [**[]TermReference**](TermReference.md) |  | [optional] 
**RelatedTerms** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**Reviewers** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**Synonyms** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to [**[]TagLabel**](TagLabel.md) |  | [optional] 

## Methods

### NewCreateGlossaryTerm

`func NewCreateGlossaryTerm(description string, glossary EntityReference, name string, ) *CreateGlossaryTerm`

NewCreateGlossaryTerm instantiates a new CreateGlossaryTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGlossaryTermWithDefaults

`func NewCreateGlossaryTermWithDefaults() *CreateGlossaryTerm`

NewCreateGlossaryTermWithDefaults instantiates a new CreateGlossaryTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateGlossaryTerm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGlossaryTerm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGlossaryTerm) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayName

`func (o *CreateGlossaryTerm) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateGlossaryTerm) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateGlossaryTerm) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateGlossaryTerm) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtension

`func (o *CreateGlossaryTerm) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *CreateGlossaryTerm) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *CreateGlossaryTerm) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *CreateGlossaryTerm) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetGlossary

`func (o *CreateGlossaryTerm) GetGlossary() EntityReference`

GetGlossary returns the Glossary field if non-nil, zero value otherwise.

### GetGlossaryOk

`func (o *CreateGlossaryTerm) GetGlossaryOk() (*EntityReference, bool)`

GetGlossaryOk returns a tuple with the Glossary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlossary

`func (o *CreateGlossaryTerm) SetGlossary(v EntityReference)`

SetGlossary sets Glossary field to given value.


### GetName

`func (o *CreateGlossaryTerm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGlossaryTerm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGlossaryTerm) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *CreateGlossaryTerm) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateGlossaryTerm) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateGlossaryTerm) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateGlossaryTerm) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParent

`func (o *CreateGlossaryTerm) GetParent() EntityReference`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *CreateGlossaryTerm) GetParentOk() (*EntityReference, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *CreateGlossaryTerm) SetParent(v EntityReference)`

SetParent sets Parent field to given value.

### HasParent

`func (o *CreateGlossaryTerm) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetReferences

`func (o *CreateGlossaryTerm) GetReferences() []TermReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *CreateGlossaryTerm) GetReferencesOk() (*[]TermReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *CreateGlossaryTerm) SetReferences(v []TermReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *CreateGlossaryTerm) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelatedTerms

`func (o *CreateGlossaryTerm) GetRelatedTerms() []EntityReference`

GetRelatedTerms returns the RelatedTerms field if non-nil, zero value otherwise.

### GetRelatedTermsOk

`func (o *CreateGlossaryTerm) GetRelatedTermsOk() (*[]EntityReference, bool)`

GetRelatedTermsOk returns a tuple with the RelatedTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedTerms

`func (o *CreateGlossaryTerm) SetRelatedTerms(v []EntityReference)`

SetRelatedTerms sets RelatedTerms field to given value.

### HasRelatedTerms

`func (o *CreateGlossaryTerm) HasRelatedTerms() bool`

HasRelatedTerms returns a boolean if a field has been set.

### GetReviewers

`func (o *CreateGlossaryTerm) GetReviewers() []EntityReference`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *CreateGlossaryTerm) GetReviewersOk() (*[]EntityReference, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *CreateGlossaryTerm) SetReviewers(v []EntityReference)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *CreateGlossaryTerm) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetSynonyms

`func (o *CreateGlossaryTerm) GetSynonyms() []string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *CreateGlossaryTerm) GetSynonymsOk() (*[]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *CreateGlossaryTerm) SetSynonyms(v []string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *CreateGlossaryTerm) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### GetTags

`func (o *CreateGlossaryTerm) GetTags() []TagLabel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateGlossaryTerm) GetTagsOk() (*[]TagLabel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateGlossaryTerm) SetTags(v []TagLabel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateGlossaryTerm) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


