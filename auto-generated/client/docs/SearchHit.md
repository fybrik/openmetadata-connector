# SearchHit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterAlias** | Pointer to **string** |  | [optional] 
**Explanation** | Pointer to [**Explanation**](Explanation.md) |  | [optional] 
**Fields** | Pointer to [**map[string]DocumentField**](DocumentField.md) |  | [optional] 
**Fragment** | Pointer to **bool** |  | [optional] 
**HighlightFields** | Pointer to [**map[string]HighlightField**](HighlightField.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Index** | Pointer to **string** |  | [optional] 
**InnerHits** | Pointer to [**map[string]SearchHits**](SearchHits.md) |  | [optional] 
**MatchedQueries** | Pointer to **[]string** |  | [optional] 
**NestedIdentity** | Pointer to [**NestedIdentity**](NestedIdentity.md) |  | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**RawSortValues** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Score** | Pointer to **float32** |  | [optional] 
**SeqNo** | Pointer to **int64** |  | [optional] 
**Shard** | Pointer to [**SearchShardTarget**](SearchShardTarget.md) |  | [optional] 
**SortValues** | Pointer to **[]map[string]interface{}** |  | [optional] 
**SourceAsMap** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**SourceAsString** | Pointer to **string** |  | [optional] 
**SourceRef** | Pointer to [**BytesReference**](BytesReference.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewSearchHit

`func NewSearchHit() *SearchHit`

NewSearchHit instantiates a new SearchHit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchHitWithDefaults

`func NewSearchHitWithDefaults() *SearchHit`

NewSearchHitWithDefaults instantiates a new SearchHit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterAlias

`func (o *SearchHit) GetClusterAlias() string`

GetClusterAlias returns the ClusterAlias field if non-nil, zero value otherwise.

### GetClusterAliasOk

`func (o *SearchHit) GetClusterAliasOk() (*string, bool)`

GetClusterAliasOk returns a tuple with the ClusterAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterAlias

`func (o *SearchHit) SetClusterAlias(v string)`

SetClusterAlias sets ClusterAlias field to given value.

### HasClusterAlias

`func (o *SearchHit) HasClusterAlias() bool`

HasClusterAlias returns a boolean if a field has been set.

### GetExplanation

`func (o *SearchHit) GetExplanation() Explanation`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *SearchHit) GetExplanationOk() (*Explanation, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *SearchHit) SetExplanation(v Explanation)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *SearchHit) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### GetFields

`func (o *SearchHit) GetFields() map[string]DocumentField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SearchHit) GetFieldsOk() (*map[string]DocumentField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SearchHit) SetFields(v map[string]DocumentField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *SearchHit) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFragment

`func (o *SearchHit) GetFragment() bool`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *SearchHit) GetFragmentOk() (*bool, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *SearchHit) SetFragment(v bool)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *SearchHit) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetHighlightFields

`func (o *SearchHit) GetHighlightFields() map[string]HighlightField`

GetHighlightFields returns the HighlightFields field if non-nil, zero value otherwise.

### GetHighlightFieldsOk

`func (o *SearchHit) GetHighlightFieldsOk() (*map[string]HighlightField, bool)`

GetHighlightFieldsOk returns a tuple with the HighlightFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightFields

`func (o *SearchHit) SetHighlightFields(v map[string]HighlightField)`

SetHighlightFields sets HighlightFields field to given value.

### HasHighlightFields

`func (o *SearchHit) HasHighlightFields() bool`

HasHighlightFields returns a boolean if a field has been set.

### GetId

`func (o *SearchHit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchHit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchHit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchHit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndex

`func (o *SearchHit) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SearchHit) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SearchHit) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *SearchHit) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetInnerHits

`func (o *SearchHit) GetInnerHits() map[string]SearchHits`

GetInnerHits returns the InnerHits field if non-nil, zero value otherwise.

### GetInnerHitsOk

`func (o *SearchHit) GetInnerHitsOk() (*map[string]SearchHits, bool)`

GetInnerHitsOk returns a tuple with the InnerHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerHits

`func (o *SearchHit) SetInnerHits(v map[string]SearchHits)`

SetInnerHits sets InnerHits field to given value.

### HasInnerHits

`func (o *SearchHit) HasInnerHits() bool`

HasInnerHits returns a boolean if a field has been set.

### GetMatchedQueries

`func (o *SearchHit) GetMatchedQueries() []string`

GetMatchedQueries returns the MatchedQueries field if non-nil, zero value otherwise.

### GetMatchedQueriesOk

`func (o *SearchHit) GetMatchedQueriesOk() (*[]string, bool)`

GetMatchedQueriesOk returns a tuple with the MatchedQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedQueries

`func (o *SearchHit) SetMatchedQueries(v []string)`

SetMatchedQueries sets MatchedQueries field to given value.

### HasMatchedQueries

`func (o *SearchHit) HasMatchedQueries() bool`

HasMatchedQueries returns a boolean if a field has been set.

### GetNestedIdentity

`func (o *SearchHit) GetNestedIdentity() NestedIdentity`

GetNestedIdentity returns the NestedIdentity field if non-nil, zero value otherwise.

### GetNestedIdentityOk

`func (o *SearchHit) GetNestedIdentityOk() (*NestedIdentity, bool)`

GetNestedIdentityOk returns a tuple with the NestedIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedIdentity

`func (o *SearchHit) SetNestedIdentity(v NestedIdentity)`

SetNestedIdentity sets NestedIdentity field to given value.

### HasNestedIdentity

`func (o *SearchHit) HasNestedIdentity() bool`

HasNestedIdentity returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *SearchHit) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *SearchHit) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *SearchHit) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *SearchHit) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetRawSortValues

`func (o *SearchHit) GetRawSortValues() []map[string]interface{}`

GetRawSortValues returns the RawSortValues field if non-nil, zero value otherwise.

### GetRawSortValuesOk

`func (o *SearchHit) GetRawSortValuesOk() (*[]map[string]interface{}, bool)`

GetRawSortValuesOk returns a tuple with the RawSortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSortValues

`func (o *SearchHit) SetRawSortValues(v []map[string]interface{})`

SetRawSortValues sets RawSortValues field to given value.

### HasRawSortValues

`func (o *SearchHit) HasRawSortValues() bool`

HasRawSortValues returns a boolean if a field has been set.

### GetScore

`func (o *SearchHit) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SearchHit) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SearchHit) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *SearchHit) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSeqNo

`func (o *SearchHit) GetSeqNo() int64`

GetSeqNo returns the SeqNo field if non-nil, zero value otherwise.

### GetSeqNoOk

`func (o *SearchHit) GetSeqNoOk() (*int64, bool)`

GetSeqNoOk returns a tuple with the SeqNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNo

`func (o *SearchHit) SetSeqNo(v int64)`

SetSeqNo sets SeqNo field to given value.

### HasSeqNo

`func (o *SearchHit) HasSeqNo() bool`

HasSeqNo returns a boolean if a field has been set.

### GetShard

`func (o *SearchHit) GetShard() SearchShardTarget`

GetShard returns the Shard field if non-nil, zero value otherwise.

### GetShardOk

`func (o *SearchHit) GetShardOk() (*SearchShardTarget, bool)`

GetShardOk returns a tuple with the Shard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShard

`func (o *SearchHit) SetShard(v SearchShardTarget)`

SetShard sets Shard field to given value.

### HasShard

`func (o *SearchHit) HasShard() bool`

HasShard returns a boolean if a field has been set.

### GetSortValues

`func (o *SearchHit) GetSortValues() []map[string]interface{}`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *SearchHit) GetSortValuesOk() (*[]map[string]interface{}, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *SearchHit) SetSortValues(v []map[string]interface{})`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *SearchHit) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetSourceAsMap

`func (o *SearchHit) GetSourceAsMap() map[string]map[string]interface{}`

GetSourceAsMap returns the SourceAsMap field if non-nil, zero value otherwise.

### GetSourceAsMapOk

`func (o *SearchHit) GetSourceAsMapOk() (*map[string]map[string]interface{}, bool)`

GetSourceAsMapOk returns a tuple with the SourceAsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAsMap

`func (o *SearchHit) SetSourceAsMap(v map[string]map[string]interface{})`

SetSourceAsMap sets SourceAsMap field to given value.

### HasSourceAsMap

`func (o *SearchHit) HasSourceAsMap() bool`

HasSourceAsMap returns a boolean if a field has been set.

### GetSourceAsString

`func (o *SearchHit) GetSourceAsString() string`

GetSourceAsString returns the SourceAsString field if non-nil, zero value otherwise.

### GetSourceAsStringOk

`func (o *SearchHit) GetSourceAsStringOk() (*string, bool)`

GetSourceAsStringOk returns a tuple with the SourceAsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAsString

`func (o *SearchHit) SetSourceAsString(v string)`

SetSourceAsString sets SourceAsString field to given value.

### HasSourceAsString

`func (o *SearchHit) HasSourceAsString() bool`

HasSourceAsString returns a boolean if a field has been set.

### GetSourceRef

`func (o *SearchHit) GetSourceRef() BytesReference`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *SearchHit) GetSourceRefOk() (*BytesReference, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *SearchHit) SetSourceRef(v BytesReference)`

SetSourceRef sets SourceRef field to given value.

### HasSourceRef

`func (o *SearchHit) HasSourceRef() bool`

HasSourceRef returns a boolean if a field has been set.

### GetType

`func (o *SearchHit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchHit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchHit) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SearchHit) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *SearchHit) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SearchHit) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SearchHit) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SearchHit) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


