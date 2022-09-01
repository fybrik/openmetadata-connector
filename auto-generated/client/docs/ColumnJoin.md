# ColumnJoin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnName** | Pointer to **string** |  | [optional] 
**JoinedWith** | Pointer to [**[]JoinedWith**](JoinedWith.md) |  | [optional] 

## Methods

### NewColumnJoin

`func NewColumnJoin() *ColumnJoin`

NewColumnJoin instantiates a new ColumnJoin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColumnJoinWithDefaults

`func NewColumnJoinWithDefaults() *ColumnJoin`

NewColumnJoinWithDefaults instantiates a new ColumnJoin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnName

`func (o *ColumnJoin) GetColumnName() string`

GetColumnName returns the ColumnName field if non-nil, zero value otherwise.

### GetColumnNameOk

`func (o *ColumnJoin) GetColumnNameOk() (*string, bool)`

GetColumnNameOk returns a tuple with the ColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnName

`func (o *ColumnJoin) SetColumnName(v string)`

SetColumnName sets ColumnName field to given value.

### HasColumnName

`func (o *ColumnJoin) HasColumnName() bool`

HasColumnName returns a boolean if a field has been set.

### GetJoinedWith

`func (o *ColumnJoin) GetJoinedWith() []JoinedWith`

GetJoinedWith returns the JoinedWith field if non-nil, zero value otherwise.

### GetJoinedWithOk

`func (o *ColumnJoin) GetJoinedWithOk() (*[]JoinedWith, bool)`

GetJoinedWithOk returns a tuple with the JoinedWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedWith

`func (o *ColumnJoin) SetJoinedWith(v []JoinedWith)`

SetJoinedWith sets JoinedWith field to given value.

### HasJoinedWith

`func (o *ColumnJoin) HasJoinedWith() bool`

HasJoinedWith returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


