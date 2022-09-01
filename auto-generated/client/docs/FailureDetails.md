# FailureDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastFailedAt** | Pointer to **int64** |  | [optional] 
**LastFailedReason** | Pointer to **string** |  | [optional] 
**LastFailedStatusCode** | Pointer to **int32** |  | [optional] 
**LastSuccessfulAt** | Pointer to **int64** |  | [optional] 
**NextAttempt** | Pointer to **int64** |  | [optional] 

## Methods

### NewFailureDetails

`func NewFailureDetails() *FailureDetails`

NewFailureDetails instantiates a new FailureDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureDetailsWithDefaults

`func NewFailureDetailsWithDefaults() *FailureDetails`

NewFailureDetailsWithDefaults instantiates a new FailureDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastFailedAt

`func (o *FailureDetails) GetLastFailedAt() int64`

GetLastFailedAt returns the LastFailedAt field if non-nil, zero value otherwise.

### GetLastFailedAtOk

`func (o *FailureDetails) GetLastFailedAtOk() (*int64, bool)`

GetLastFailedAtOk returns a tuple with the LastFailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailedAt

`func (o *FailureDetails) SetLastFailedAt(v int64)`

SetLastFailedAt sets LastFailedAt field to given value.

### HasLastFailedAt

`func (o *FailureDetails) HasLastFailedAt() bool`

HasLastFailedAt returns a boolean if a field has been set.

### GetLastFailedReason

`func (o *FailureDetails) GetLastFailedReason() string`

GetLastFailedReason returns the LastFailedReason field if non-nil, zero value otherwise.

### GetLastFailedReasonOk

`func (o *FailureDetails) GetLastFailedReasonOk() (*string, bool)`

GetLastFailedReasonOk returns a tuple with the LastFailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailedReason

`func (o *FailureDetails) SetLastFailedReason(v string)`

SetLastFailedReason sets LastFailedReason field to given value.

### HasLastFailedReason

`func (o *FailureDetails) HasLastFailedReason() bool`

HasLastFailedReason returns a boolean if a field has been set.

### GetLastFailedStatusCode

`func (o *FailureDetails) GetLastFailedStatusCode() int32`

GetLastFailedStatusCode returns the LastFailedStatusCode field if non-nil, zero value otherwise.

### GetLastFailedStatusCodeOk

`func (o *FailureDetails) GetLastFailedStatusCodeOk() (*int32, bool)`

GetLastFailedStatusCodeOk returns a tuple with the LastFailedStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailedStatusCode

`func (o *FailureDetails) SetLastFailedStatusCode(v int32)`

SetLastFailedStatusCode sets LastFailedStatusCode field to given value.

### HasLastFailedStatusCode

`func (o *FailureDetails) HasLastFailedStatusCode() bool`

HasLastFailedStatusCode returns a boolean if a field has been set.

### GetLastSuccessfulAt

`func (o *FailureDetails) GetLastSuccessfulAt() int64`

GetLastSuccessfulAt returns the LastSuccessfulAt field if non-nil, zero value otherwise.

### GetLastSuccessfulAtOk

`func (o *FailureDetails) GetLastSuccessfulAtOk() (*int64, bool)`

GetLastSuccessfulAtOk returns a tuple with the LastSuccessfulAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulAt

`func (o *FailureDetails) SetLastSuccessfulAt(v int64)`

SetLastSuccessfulAt sets LastSuccessfulAt field to given value.

### HasLastSuccessfulAt

`func (o *FailureDetails) HasLastSuccessfulAt() bool`

HasLastSuccessfulAt returns a boolean if a field has been set.

### GetNextAttempt

`func (o *FailureDetails) GetNextAttempt() int64`

GetNextAttempt returns the NextAttempt field if non-nil, zero value otherwise.

### GetNextAttemptOk

`func (o *FailureDetails) GetNextAttemptOk() (*int64, bool)`

GetNextAttemptOk returns a tuple with the NextAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAttempt

`func (o *FailureDetails) SetNextAttempt(v int64)`

SetNextAttempt sets NextAttempt field to given value.

### HasNextAttempt

`func (o *FailureDetails) HasNextAttempt() bool`

HasNextAttempt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


