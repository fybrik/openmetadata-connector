# SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregations** | Pointer to [**Aggregations**](Aggregations.md) |  | [optional] 
**Clusters** | Pointer to [**Clusters**](Clusters.md) |  | [optional] 
**FailedShards** | Pointer to **int32** |  | [optional] 
**Fragment** | Pointer to **bool** |  | [optional] 
**Hits** | Pointer to [**SearchHits**](SearchHits.md) |  | [optional] 
**InternalResponse** | Pointer to [**SearchResponseSections**](SearchResponseSections.md) |  | [optional] 
**NumReducePhases** | Pointer to **int32** |  | [optional] 
**ProfileResults** | Pointer to [**map[string]ProfileShardResult**](ProfileShardResult.md) |  | [optional] 
**ScrollId** | Pointer to **string** |  | [optional] 
**ShardFailures** | Pointer to [**[]ShardSearchFailure**](ShardSearchFailure.md) |  | [optional] 
**SkippedShards** | Pointer to **int32** |  | [optional] 
**SuccessfulShards** | Pointer to **int32** |  | [optional] 
**Suggest** | Pointer to [**Suggest**](Suggest.md) |  | [optional] 
**TerminatedEarly** | Pointer to **bool** |  | [optional] 
**TimedOut** | Pointer to **bool** |  | [optional] 
**Took** | Pointer to [**TimeValue**](TimeValue.md) |  | [optional] 
**TotalShards** | Pointer to **int32** |  | [optional] 

## Methods

### NewSearchResponse

`func NewSearchResponse() *SearchResponse`

NewSearchResponse instantiates a new SearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseWithDefaults

`func NewSearchResponseWithDefaults() *SearchResponse`

NewSearchResponseWithDefaults instantiates a new SearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregations

`func (o *SearchResponse) GetAggregations() Aggregations`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *SearchResponse) GetAggregationsOk() (*Aggregations, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *SearchResponse) SetAggregations(v Aggregations)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *SearchResponse) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetClusters

`func (o *SearchResponse) GetClusters() Clusters`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *SearchResponse) GetClustersOk() (*Clusters, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *SearchResponse) SetClusters(v Clusters)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *SearchResponse) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetFailedShards

`func (o *SearchResponse) GetFailedShards() int32`

GetFailedShards returns the FailedShards field if non-nil, zero value otherwise.

### GetFailedShardsOk

`func (o *SearchResponse) GetFailedShardsOk() (*int32, bool)`

GetFailedShardsOk returns a tuple with the FailedShards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedShards

`func (o *SearchResponse) SetFailedShards(v int32)`

SetFailedShards sets FailedShards field to given value.

### HasFailedShards

`func (o *SearchResponse) HasFailedShards() bool`

HasFailedShards returns a boolean if a field has been set.

### GetFragment

`func (o *SearchResponse) GetFragment() bool`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *SearchResponse) GetFragmentOk() (*bool, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *SearchResponse) SetFragment(v bool)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *SearchResponse) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetHits

`func (o *SearchResponse) GetHits() SearchHits`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *SearchResponse) GetHitsOk() (*SearchHits, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *SearchResponse) SetHits(v SearchHits)`

SetHits sets Hits field to given value.

### HasHits

`func (o *SearchResponse) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetInternalResponse

`func (o *SearchResponse) GetInternalResponse() SearchResponseSections`

GetInternalResponse returns the InternalResponse field if non-nil, zero value otherwise.

### GetInternalResponseOk

`func (o *SearchResponse) GetInternalResponseOk() (*SearchResponseSections, bool)`

GetInternalResponseOk returns a tuple with the InternalResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalResponse

`func (o *SearchResponse) SetInternalResponse(v SearchResponseSections)`

SetInternalResponse sets InternalResponse field to given value.

### HasInternalResponse

`func (o *SearchResponse) HasInternalResponse() bool`

HasInternalResponse returns a boolean if a field has been set.

### GetNumReducePhases

`func (o *SearchResponse) GetNumReducePhases() int32`

GetNumReducePhases returns the NumReducePhases field if non-nil, zero value otherwise.

### GetNumReducePhasesOk

`func (o *SearchResponse) GetNumReducePhasesOk() (*int32, bool)`

GetNumReducePhasesOk returns a tuple with the NumReducePhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReducePhases

`func (o *SearchResponse) SetNumReducePhases(v int32)`

SetNumReducePhases sets NumReducePhases field to given value.

### HasNumReducePhases

`func (o *SearchResponse) HasNumReducePhases() bool`

HasNumReducePhases returns a boolean if a field has been set.

### GetProfileResults

`func (o *SearchResponse) GetProfileResults() map[string]ProfileShardResult`

GetProfileResults returns the ProfileResults field if non-nil, zero value otherwise.

### GetProfileResultsOk

`func (o *SearchResponse) GetProfileResultsOk() (*map[string]ProfileShardResult, bool)`

GetProfileResultsOk returns a tuple with the ProfileResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileResults

`func (o *SearchResponse) SetProfileResults(v map[string]ProfileShardResult)`

SetProfileResults sets ProfileResults field to given value.

### HasProfileResults

`func (o *SearchResponse) HasProfileResults() bool`

HasProfileResults returns a boolean if a field has been set.

### GetScrollId

`func (o *SearchResponse) GetScrollId() string`

GetScrollId returns the ScrollId field if non-nil, zero value otherwise.

### GetScrollIdOk

`func (o *SearchResponse) GetScrollIdOk() (*string, bool)`

GetScrollIdOk returns a tuple with the ScrollId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollId

`func (o *SearchResponse) SetScrollId(v string)`

SetScrollId sets ScrollId field to given value.

### HasScrollId

`func (o *SearchResponse) HasScrollId() bool`

HasScrollId returns a boolean if a field has been set.

### GetShardFailures

`func (o *SearchResponse) GetShardFailures() []ShardSearchFailure`

GetShardFailures returns the ShardFailures field if non-nil, zero value otherwise.

### GetShardFailuresOk

`func (o *SearchResponse) GetShardFailuresOk() (*[]ShardSearchFailure, bool)`

GetShardFailuresOk returns a tuple with the ShardFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardFailures

`func (o *SearchResponse) SetShardFailures(v []ShardSearchFailure)`

SetShardFailures sets ShardFailures field to given value.

### HasShardFailures

`func (o *SearchResponse) HasShardFailures() bool`

HasShardFailures returns a boolean if a field has been set.

### GetSkippedShards

`func (o *SearchResponse) GetSkippedShards() int32`

GetSkippedShards returns the SkippedShards field if non-nil, zero value otherwise.

### GetSkippedShardsOk

`func (o *SearchResponse) GetSkippedShardsOk() (*int32, bool)`

GetSkippedShardsOk returns a tuple with the SkippedShards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedShards

`func (o *SearchResponse) SetSkippedShards(v int32)`

SetSkippedShards sets SkippedShards field to given value.

### HasSkippedShards

`func (o *SearchResponse) HasSkippedShards() bool`

HasSkippedShards returns a boolean if a field has been set.

### GetSuccessfulShards

`func (o *SearchResponse) GetSuccessfulShards() int32`

GetSuccessfulShards returns the SuccessfulShards field if non-nil, zero value otherwise.

### GetSuccessfulShardsOk

`func (o *SearchResponse) GetSuccessfulShardsOk() (*int32, bool)`

GetSuccessfulShardsOk returns a tuple with the SuccessfulShards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulShards

`func (o *SearchResponse) SetSuccessfulShards(v int32)`

SetSuccessfulShards sets SuccessfulShards field to given value.

### HasSuccessfulShards

`func (o *SearchResponse) HasSuccessfulShards() bool`

HasSuccessfulShards returns a boolean if a field has been set.

### GetSuggest

`func (o *SearchResponse) GetSuggest() Suggest`

GetSuggest returns the Suggest field if non-nil, zero value otherwise.

### GetSuggestOk

`func (o *SearchResponse) GetSuggestOk() (*Suggest, bool)`

GetSuggestOk returns a tuple with the Suggest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggest

`func (o *SearchResponse) SetSuggest(v Suggest)`

SetSuggest sets Suggest field to given value.

### HasSuggest

`func (o *SearchResponse) HasSuggest() bool`

HasSuggest returns a boolean if a field has been set.

### GetTerminatedEarly

`func (o *SearchResponse) GetTerminatedEarly() bool`

GetTerminatedEarly returns the TerminatedEarly field if non-nil, zero value otherwise.

### GetTerminatedEarlyOk

`func (o *SearchResponse) GetTerminatedEarlyOk() (*bool, bool)`

GetTerminatedEarlyOk returns a tuple with the TerminatedEarly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedEarly

`func (o *SearchResponse) SetTerminatedEarly(v bool)`

SetTerminatedEarly sets TerminatedEarly field to given value.

### HasTerminatedEarly

`func (o *SearchResponse) HasTerminatedEarly() bool`

HasTerminatedEarly returns a boolean if a field has been set.

### GetTimedOut

`func (o *SearchResponse) GetTimedOut() bool`

GetTimedOut returns the TimedOut field if non-nil, zero value otherwise.

### GetTimedOutOk

`func (o *SearchResponse) GetTimedOutOk() (*bool, bool)`

GetTimedOutOk returns a tuple with the TimedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimedOut

`func (o *SearchResponse) SetTimedOut(v bool)`

SetTimedOut sets TimedOut field to given value.

### HasTimedOut

`func (o *SearchResponse) HasTimedOut() bool`

HasTimedOut returns a boolean if a field has been set.

### GetTook

`func (o *SearchResponse) GetTook() TimeValue`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *SearchResponse) GetTookOk() (*TimeValue, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *SearchResponse) SetTook(v TimeValue)`

SetTook sets Took field to given value.

### HasTook

`func (o *SearchResponse) HasTook() bool`

HasTook returns a boolean if a field has been set.

### GetTotalShards

`func (o *SearchResponse) GetTotalShards() int32`

GetTotalShards returns the TotalShards field if non-nil, zero value otherwise.

### GetTotalShardsOk

`func (o *SearchResponse) GetTotalShardsOk() (*int32, bool)`

GetTotalShardsOk returns a tuple with the TotalShards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShards

`func (o *SearchResponse) SetTotalShards(v int32)`

SetTotalShards sets TotalShards field to given value.

### HasTotalShards

`func (o *SearchResponse) HasTotalShards() bool`

HasTotalShards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


