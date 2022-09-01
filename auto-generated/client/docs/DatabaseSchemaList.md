# DatabaseSchemaList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DatabaseSchema**](DatabaseSchema.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewDatabaseSchemaList

`func NewDatabaseSchemaList(data []DatabaseSchema, ) *DatabaseSchemaList`

NewDatabaseSchemaList instantiates a new DatabaseSchemaList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseSchemaListWithDefaults

`func NewDatabaseSchemaListWithDefaults() *DatabaseSchemaList`

NewDatabaseSchemaListWithDefaults instantiates a new DatabaseSchemaList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DatabaseSchemaList) GetData() []DatabaseSchema`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DatabaseSchemaList) GetDataOk() (*[]DatabaseSchema, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DatabaseSchemaList) SetData(v []DatabaseSchema)`

SetData sets Data field to given value.


### GetPaging

`func (o *DatabaseSchemaList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *DatabaseSchemaList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *DatabaseSchemaList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *DatabaseSchemaList) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


