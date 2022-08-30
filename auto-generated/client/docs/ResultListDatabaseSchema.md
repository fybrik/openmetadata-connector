# ResultListDatabaseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DatabaseSchema**](DatabaseSchema.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewResultListDatabaseSchema

`func NewResultListDatabaseSchema(data []DatabaseSchema, ) *ResultListDatabaseSchema`

NewResultListDatabaseSchema instantiates a new ResultListDatabaseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultListDatabaseSchemaWithDefaults

`func NewResultListDatabaseSchemaWithDefaults() *ResultListDatabaseSchema`

NewResultListDatabaseSchemaWithDefaults instantiates a new ResultListDatabaseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResultListDatabaseSchema) GetData() []DatabaseSchema`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultListDatabaseSchema) GetDataOk() (*[]DatabaseSchema, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultListDatabaseSchema) SetData(v []DatabaseSchema)`

SetData sets Data field to given value.


### GetPaging

`func (o *ResultListDatabaseSchema) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *ResultListDatabaseSchema) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *ResultListDatabaseSchema) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *ResultListDatabaseSchema) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


