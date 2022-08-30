# CreateChart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartType** | Pointer to **string** |  | [optional] 
**ChartUrl** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**Owner** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Service** | [**EntityReference**](EntityReference.md) |  | 
**Tables** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**Tags** | Pointer to [**[]TagLabel**](TagLabel.md) |  | [optional] 

## Methods

### NewCreateChart

`func NewCreateChart(name string, service EntityReference, ) *CreateChart`

NewCreateChart instantiates a new CreateChart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateChartWithDefaults

`func NewCreateChartWithDefaults() *CreateChart`

NewCreateChartWithDefaults instantiates a new CreateChart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartType

`func (o *CreateChart) GetChartType() string`

GetChartType returns the ChartType field if non-nil, zero value otherwise.

### GetChartTypeOk

`func (o *CreateChart) GetChartTypeOk() (*string, bool)`

GetChartTypeOk returns a tuple with the ChartType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartType

`func (o *CreateChart) SetChartType(v string)`

SetChartType sets ChartType field to given value.

### HasChartType

`func (o *CreateChart) HasChartType() bool`

HasChartType returns a boolean if a field has been set.

### GetChartUrl

`func (o *CreateChart) GetChartUrl() string`

GetChartUrl returns the ChartUrl field if non-nil, zero value otherwise.

### GetChartUrlOk

`func (o *CreateChart) GetChartUrlOk() (*string, bool)`

GetChartUrlOk returns a tuple with the ChartUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartUrl

`func (o *CreateChart) SetChartUrl(v string)`

SetChartUrl sets ChartUrl field to given value.

### HasChartUrl

`func (o *CreateChart) HasChartUrl() bool`

HasChartUrl returns a boolean if a field has been set.

### GetDescription

`func (o *CreateChart) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateChart) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateChart) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateChart) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateChart) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateChart) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateChart) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateChart) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtension

`func (o *CreateChart) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *CreateChart) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *CreateChart) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *CreateChart) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetName

`func (o *CreateChart) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateChart) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateChart) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *CreateChart) GetOwner() EntityReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateChart) GetOwnerOk() (*EntityReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateChart) SetOwner(v EntityReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateChart) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetService

`func (o *CreateChart) GetService() EntityReference`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CreateChart) GetServiceOk() (*EntityReference, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CreateChart) SetService(v EntityReference)`

SetService sets Service field to given value.


### GetTables

`func (o *CreateChart) GetTables() []EntityReference`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *CreateChart) GetTablesOk() (*[]EntityReference, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *CreateChart) SetTables(v []EntityReference)`

SetTables sets Tables field to given value.

### HasTables

`func (o *CreateChart) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetTags

`func (o *CreateChart) GetTags() []TagLabel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateChart) GetTagsOk() (*[]TagLabel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateChart) SetTags(v []TagLabel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateChart) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


