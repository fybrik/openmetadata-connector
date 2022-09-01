# TableProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnCount** | Pointer to **float64** |  | [optional] 
**ColumnNames** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ColumnProfile** | Pointer to [**[]ColumnProfile**](ColumnProfile.md) |  | [optional] 
**ProfileDate** | Pointer to **string** |  | [optional] 
**ProfileQuery** | Pointer to **string** |  | [optional] 
**ProfileSample** | Pointer to **float64** |  | [optional] 
**RowCount** | Pointer to **float64** |  | [optional] 

## Methods

### NewTableProfile

`func NewTableProfile() *TableProfile`

NewTableProfile instantiates a new TableProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableProfileWithDefaults

`func NewTableProfileWithDefaults() *TableProfile`

NewTableProfileWithDefaults instantiates a new TableProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnCount

`func (o *TableProfile) GetColumnCount() float64`

GetColumnCount returns the ColumnCount field if non-nil, zero value otherwise.

### GetColumnCountOk

`func (o *TableProfile) GetColumnCountOk() (*float64, bool)`

GetColumnCountOk returns a tuple with the ColumnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnCount

`func (o *TableProfile) SetColumnCount(v float64)`

SetColumnCount sets ColumnCount field to given value.

### HasColumnCount

`func (o *TableProfile) HasColumnCount() bool`

HasColumnCount returns a boolean if a field has been set.

### GetColumnNames

`func (o *TableProfile) GetColumnNames() []map[string]interface{}`

GetColumnNames returns the ColumnNames field if non-nil, zero value otherwise.

### GetColumnNamesOk

`func (o *TableProfile) GetColumnNamesOk() (*[]map[string]interface{}, bool)`

GetColumnNamesOk returns a tuple with the ColumnNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnNames

`func (o *TableProfile) SetColumnNames(v []map[string]interface{})`

SetColumnNames sets ColumnNames field to given value.

### HasColumnNames

`func (o *TableProfile) HasColumnNames() bool`

HasColumnNames returns a boolean if a field has been set.

### GetColumnProfile

`func (o *TableProfile) GetColumnProfile() []ColumnProfile`

GetColumnProfile returns the ColumnProfile field if non-nil, zero value otherwise.

### GetColumnProfileOk

`func (o *TableProfile) GetColumnProfileOk() (*[]ColumnProfile, bool)`

GetColumnProfileOk returns a tuple with the ColumnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnProfile

`func (o *TableProfile) SetColumnProfile(v []ColumnProfile)`

SetColumnProfile sets ColumnProfile field to given value.

### HasColumnProfile

`func (o *TableProfile) HasColumnProfile() bool`

HasColumnProfile returns a boolean if a field has been set.

### GetProfileDate

`func (o *TableProfile) GetProfileDate() string`

GetProfileDate returns the ProfileDate field if non-nil, zero value otherwise.

### GetProfileDateOk

`func (o *TableProfile) GetProfileDateOk() (*string, bool)`

GetProfileDateOk returns a tuple with the ProfileDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileDate

`func (o *TableProfile) SetProfileDate(v string)`

SetProfileDate sets ProfileDate field to given value.

### HasProfileDate

`func (o *TableProfile) HasProfileDate() bool`

HasProfileDate returns a boolean if a field has been set.

### GetProfileQuery

`func (o *TableProfile) GetProfileQuery() string`

GetProfileQuery returns the ProfileQuery field if non-nil, zero value otherwise.

### GetProfileQueryOk

`func (o *TableProfile) GetProfileQueryOk() (*string, bool)`

GetProfileQueryOk returns a tuple with the ProfileQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileQuery

`func (o *TableProfile) SetProfileQuery(v string)`

SetProfileQuery sets ProfileQuery field to given value.

### HasProfileQuery

`func (o *TableProfile) HasProfileQuery() bool`

HasProfileQuery returns a boolean if a field has been set.

### GetProfileSample

`func (o *TableProfile) GetProfileSample() float64`

GetProfileSample returns the ProfileSample field if non-nil, zero value otherwise.

### GetProfileSampleOk

`func (o *TableProfile) GetProfileSampleOk() (*float64, bool)`

GetProfileSampleOk returns a tuple with the ProfileSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileSample

`func (o *TableProfile) SetProfileSample(v float64)`

SetProfileSample sets ProfileSample field to given value.

### HasProfileSample

`func (o *TableProfile) HasProfileSample() bool`

HasProfileSample returns a boolean if a field has been set.

### GetRowCount

`func (o *TableProfile) GetRowCount() float64`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *TableProfile) GetRowCountOk() (*float64, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *TableProfile) SetRowCount(v float64)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *TableProfile) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


