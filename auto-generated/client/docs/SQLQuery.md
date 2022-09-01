# SQLQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checksum** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **float64** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**QueryDate** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Vote** | Pointer to **float64** |  | [optional] 

## Methods

### NewSQLQuery

`func NewSQLQuery() *SQLQuery`

NewSQLQuery instantiates a new SQLQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSQLQueryWithDefaults

`func NewSQLQueryWithDefaults() *SQLQuery`

NewSQLQueryWithDefaults instantiates a new SQLQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecksum

`func (o *SQLQuery) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *SQLQuery) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *SQLQuery) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *SQLQuery) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetDuration

`func (o *SQLQuery) GetDuration() float64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SQLQuery) GetDurationOk() (*float64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SQLQuery) SetDuration(v float64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SQLQuery) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetQuery

`func (o *SQLQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SQLQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SQLQuery) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SQLQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetQueryDate

`func (o *SQLQuery) GetQueryDate() string`

GetQueryDate returns the QueryDate field if non-nil, zero value otherwise.

### GetQueryDateOk

`func (o *SQLQuery) GetQueryDateOk() (*string, bool)`

GetQueryDateOk returns a tuple with the QueryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryDate

`func (o *SQLQuery) SetQueryDate(v string)`

SetQueryDate sets QueryDate field to given value.

### HasQueryDate

`func (o *SQLQuery) HasQueryDate() bool`

HasQueryDate returns a boolean if a field has been set.

### GetUser

`func (o *SQLQuery) GetUser() EntityReference`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SQLQuery) GetUserOk() (*EntityReference, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SQLQuery) SetUser(v EntityReference)`

SetUser sets User field to given value.

### HasUser

`func (o *SQLQuery) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVote

`func (o *SQLQuery) GetVote() float64`

GetVote returns the Vote field if non-nil, zero value otherwise.

### GetVoteOk

`func (o *SQLQuery) GetVoteOk() (*float64, bool)`

GetVoteOk returns a tuple with the Vote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVote

`func (o *SQLQuery) SetVote(v float64)`

SetVote sets Vote field to given value.

### HasVote

`func (o *SQLQuery) HasVote() bool`

HasVote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


