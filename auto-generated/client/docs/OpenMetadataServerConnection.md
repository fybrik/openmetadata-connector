# OpenMetadataServerConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**AuthProvider** | Pointer to **string** |  | [optional] 
**EnableVersionValidation** | Pointer to **bool** |  | [optional] 
**HostPort** | **string** |  | 
**IncludeDashboards** | Pointer to **bool** |  | [optional] 
**IncludeDatabaseServices** | Pointer to **bool** |  | [optional] 
**IncludeGlossaryTerms** | Pointer to **bool** |  | [optional] 
**IncludeMessagingServices** | Pointer to **bool** |  | [optional] 
**IncludeMlModels** | Pointer to **bool** |  | [optional] 
**IncludePipelineServices** | Pointer to **bool** |  | [optional] 
**IncludePipelines** | Pointer to **bool** |  | [optional] 
**IncludePolicy** | Pointer to **bool** |  | [optional] 
**IncludeTables** | Pointer to **bool** |  | [optional] 
**IncludeTags** | Pointer to **bool** |  | [optional] 
**IncludeTeams** | Pointer to **bool** |  | [optional] 
**IncludeTopics** | Pointer to **bool** |  | [optional] 
**IncludeUsers** | Pointer to **bool** |  | [optional] 
**LimitRecords** | Pointer to **int32** |  | [optional] 
**SecurityConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**SupportsMetadataExtraction** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewOpenMetadataServerConnection

`func NewOpenMetadataServerConnection(hostPort string, ) *OpenMetadataServerConnection`

NewOpenMetadataServerConnection instantiates a new OpenMetadataServerConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenMetadataServerConnectionWithDefaults

`func NewOpenMetadataServerConnectionWithDefaults() *OpenMetadataServerConnection`

NewOpenMetadataServerConnectionWithDefaults instantiates a new OpenMetadataServerConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *OpenMetadataServerConnection) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *OpenMetadataServerConnection) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *OpenMetadataServerConnection) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *OpenMetadataServerConnection) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetAuthProvider

`func (o *OpenMetadataServerConnection) GetAuthProvider() string`

GetAuthProvider returns the AuthProvider field if non-nil, zero value otherwise.

### GetAuthProviderOk

`func (o *OpenMetadataServerConnection) GetAuthProviderOk() (*string, bool)`

GetAuthProviderOk returns a tuple with the AuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProvider

`func (o *OpenMetadataServerConnection) SetAuthProvider(v string)`

SetAuthProvider sets AuthProvider field to given value.

### HasAuthProvider

`func (o *OpenMetadataServerConnection) HasAuthProvider() bool`

HasAuthProvider returns a boolean if a field has been set.

### GetEnableVersionValidation

`func (o *OpenMetadataServerConnection) GetEnableVersionValidation() bool`

GetEnableVersionValidation returns the EnableVersionValidation field if non-nil, zero value otherwise.

### GetEnableVersionValidationOk

`func (o *OpenMetadataServerConnection) GetEnableVersionValidationOk() (*bool, bool)`

GetEnableVersionValidationOk returns a tuple with the EnableVersionValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVersionValidation

`func (o *OpenMetadataServerConnection) SetEnableVersionValidation(v bool)`

SetEnableVersionValidation sets EnableVersionValidation field to given value.

### HasEnableVersionValidation

`func (o *OpenMetadataServerConnection) HasEnableVersionValidation() bool`

HasEnableVersionValidation returns a boolean if a field has been set.

### GetHostPort

`func (o *OpenMetadataServerConnection) GetHostPort() string`

GetHostPort returns the HostPort field if non-nil, zero value otherwise.

### GetHostPortOk

`func (o *OpenMetadataServerConnection) GetHostPortOk() (*string, bool)`

GetHostPortOk returns a tuple with the HostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPort

`func (o *OpenMetadataServerConnection) SetHostPort(v string)`

SetHostPort sets HostPort field to given value.


### GetIncludeDashboards

`func (o *OpenMetadataServerConnection) GetIncludeDashboards() bool`

GetIncludeDashboards returns the IncludeDashboards field if non-nil, zero value otherwise.

### GetIncludeDashboardsOk

`func (o *OpenMetadataServerConnection) GetIncludeDashboardsOk() (*bool, bool)`

GetIncludeDashboardsOk returns a tuple with the IncludeDashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDashboards

`func (o *OpenMetadataServerConnection) SetIncludeDashboards(v bool)`

SetIncludeDashboards sets IncludeDashboards field to given value.

### HasIncludeDashboards

`func (o *OpenMetadataServerConnection) HasIncludeDashboards() bool`

HasIncludeDashboards returns a boolean if a field has been set.

### GetIncludeDatabaseServices

`func (o *OpenMetadataServerConnection) GetIncludeDatabaseServices() bool`

GetIncludeDatabaseServices returns the IncludeDatabaseServices field if non-nil, zero value otherwise.

### GetIncludeDatabaseServicesOk

`func (o *OpenMetadataServerConnection) GetIncludeDatabaseServicesOk() (*bool, bool)`

GetIncludeDatabaseServicesOk returns a tuple with the IncludeDatabaseServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDatabaseServices

`func (o *OpenMetadataServerConnection) SetIncludeDatabaseServices(v bool)`

SetIncludeDatabaseServices sets IncludeDatabaseServices field to given value.

### HasIncludeDatabaseServices

`func (o *OpenMetadataServerConnection) HasIncludeDatabaseServices() bool`

HasIncludeDatabaseServices returns a boolean if a field has been set.

### GetIncludeGlossaryTerms

`func (o *OpenMetadataServerConnection) GetIncludeGlossaryTerms() bool`

GetIncludeGlossaryTerms returns the IncludeGlossaryTerms field if non-nil, zero value otherwise.

### GetIncludeGlossaryTermsOk

`func (o *OpenMetadataServerConnection) GetIncludeGlossaryTermsOk() (*bool, bool)`

GetIncludeGlossaryTermsOk returns a tuple with the IncludeGlossaryTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeGlossaryTerms

`func (o *OpenMetadataServerConnection) SetIncludeGlossaryTerms(v bool)`

SetIncludeGlossaryTerms sets IncludeGlossaryTerms field to given value.

### HasIncludeGlossaryTerms

`func (o *OpenMetadataServerConnection) HasIncludeGlossaryTerms() bool`

HasIncludeGlossaryTerms returns a boolean if a field has been set.

### GetIncludeMessagingServices

`func (o *OpenMetadataServerConnection) GetIncludeMessagingServices() bool`

GetIncludeMessagingServices returns the IncludeMessagingServices field if non-nil, zero value otherwise.

### GetIncludeMessagingServicesOk

`func (o *OpenMetadataServerConnection) GetIncludeMessagingServicesOk() (*bool, bool)`

GetIncludeMessagingServicesOk returns a tuple with the IncludeMessagingServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMessagingServices

`func (o *OpenMetadataServerConnection) SetIncludeMessagingServices(v bool)`

SetIncludeMessagingServices sets IncludeMessagingServices field to given value.

### HasIncludeMessagingServices

`func (o *OpenMetadataServerConnection) HasIncludeMessagingServices() bool`

HasIncludeMessagingServices returns a boolean if a field has been set.

### GetIncludeMlModels

`func (o *OpenMetadataServerConnection) GetIncludeMlModels() bool`

GetIncludeMlModels returns the IncludeMlModels field if non-nil, zero value otherwise.

### GetIncludeMlModelsOk

`func (o *OpenMetadataServerConnection) GetIncludeMlModelsOk() (*bool, bool)`

GetIncludeMlModelsOk returns a tuple with the IncludeMlModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMlModels

`func (o *OpenMetadataServerConnection) SetIncludeMlModels(v bool)`

SetIncludeMlModels sets IncludeMlModels field to given value.

### HasIncludeMlModels

`func (o *OpenMetadataServerConnection) HasIncludeMlModels() bool`

HasIncludeMlModels returns a boolean if a field has been set.

### GetIncludePipelineServices

`func (o *OpenMetadataServerConnection) GetIncludePipelineServices() bool`

GetIncludePipelineServices returns the IncludePipelineServices field if non-nil, zero value otherwise.

### GetIncludePipelineServicesOk

`func (o *OpenMetadataServerConnection) GetIncludePipelineServicesOk() (*bool, bool)`

GetIncludePipelineServicesOk returns a tuple with the IncludePipelineServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePipelineServices

`func (o *OpenMetadataServerConnection) SetIncludePipelineServices(v bool)`

SetIncludePipelineServices sets IncludePipelineServices field to given value.

### HasIncludePipelineServices

`func (o *OpenMetadataServerConnection) HasIncludePipelineServices() bool`

HasIncludePipelineServices returns a boolean if a field has been set.

### GetIncludePipelines

`func (o *OpenMetadataServerConnection) GetIncludePipelines() bool`

GetIncludePipelines returns the IncludePipelines field if non-nil, zero value otherwise.

### GetIncludePipelinesOk

`func (o *OpenMetadataServerConnection) GetIncludePipelinesOk() (*bool, bool)`

GetIncludePipelinesOk returns a tuple with the IncludePipelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePipelines

`func (o *OpenMetadataServerConnection) SetIncludePipelines(v bool)`

SetIncludePipelines sets IncludePipelines field to given value.

### HasIncludePipelines

`func (o *OpenMetadataServerConnection) HasIncludePipelines() bool`

HasIncludePipelines returns a boolean if a field has been set.

### GetIncludePolicy

`func (o *OpenMetadataServerConnection) GetIncludePolicy() bool`

GetIncludePolicy returns the IncludePolicy field if non-nil, zero value otherwise.

### GetIncludePolicyOk

`func (o *OpenMetadataServerConnection) GetIncludePolicyOk() (*bool, bool)`

GetIncludePolicyOk returns a tuple with the IncludePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePolicy

`func (o *OpenMetadataServerConnection) SetIncludePolicy(v bool)`

SetIncludePolicy sets IncludePolicy field to given value.

### HasIncludePolicy

`func (o *OpenMetadataServerConnection) HasIncludePolicy() bool`

HasIncludePolicy returns a boolean if a field has been set.

### GetIncludeTables

`func (o *OpenMetadataServerConnection) GetIncludeTables() bool`

GetIncludeTables returns the IncludeTables field if non-nil, zero value otherwise.

### GetIncludeTablesOk

`func (o *OpenMetadataServerConnection) GetIncludeTablesOk() (*bool, bool)`

GetIncludeTablesOk returns a tuple with the IncludeTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTables

`func (o *OpenMetadataServerConnection) SetIncludeTables(v bool)`

SetIncludeTables sets IncludeTables field to given value.

### HasIncludeTables

`func (o *OpenMetadataServerConnection) HasIncludeTables() bool`

HasIncludeTables returns a boolean if a field has been set.

### GetIncludeTags

`func (o *OpenMetadataServerConnection) GetIncludeTags() bool`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *OpenMetadataServerConnection) GetIncludeTagsOk() (*bool, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *OpenMetadataServerConnection) SetIncludeTags(v bool)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *OpenMetadataServerConnection) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### GetIncludeTeams

`func (o *OpenMetadataServerConnection) GetIncludeTeams() bool`

GetIncludeTeams returns the IncludeTeams field if non-nil, zero value otherwise.

### GetIncludeTeamsOk

`func (o *OpenMetadataServerConnection) GetIncludeTeamsOk() (*bool, bool)`

GetIncludeTeamsOk returns a tuple with the IncludeTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTeams

`func (o *OpenMetadataServerConnection) SetIncludeTeams(v bool)`

SetIncludeTeams sets IncludeTeams field to given value.

### HasIncludeTeams

`func (o *OpenMetadataServerConnection) HasIncludeTeams() bool`

HasIncludeTeams returns a boolean if a field has been set.

### GetIncludeTopics

`func (o *OpenMetadataServerConnection) GetIncludeTopics() bool`

GetIncludeTopics returns the IncludeTopics field if non-nil, zero value otherwise.

### GetIncludeTopicsOk

`func (o *OpenMetadataServerConnection) GetIncludeTopicsOk() (*bool, bool)`

GetIncludeTopicsOk returns a tuple with the IncludeTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTopics

`func (o *OpenMetadataServerConnection) SetIncludeTopics(v bool)`

SetIncludeTopics sets IncludeTopics field to given value.

### HasIncludeTopics

`func (o *OpenMetadataServerConnection) HasIncludeTopics() bool`

HasIncludeTopics returns a boolean if a field has been set.

### GetIncludeUsers

`func (o *OpenMetadataServerConnection) GetIncludeUsers() bool`

GetIncludeUsers returns the IncludeUsers field if non-nil, zero value otherwise.

### GetIncludeUsersOk

`func (o *OpenMetadataServerConnection) GetIncludeUsersOk() (*bool, bool)`

GetIncludeUsersOk returns a tuple with the IncludeUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUsers

`func (o *OpenMetadataServerConnection) SetIncludeUsers(v bool)`

SetIncludeUsers sets IncludeUsers field to given value.

### HasIncludeUsers

`func (o *OpenMetadataServerConnection) HasIncludeUsers() bool`

HasIncludeUsers returns a boolean if a field has been set.

### GetLimitRecords

`func (o *OpenMetadataServerConnection) GetLimitRecords() int32`

GetLimitRecords returns the LimitRecords field if non-nil, zero value otherwise.

### GetLimitRecordsOk

`func (o *OpenMetadataServerConnection) GetLimitRecordsOk() (*int32, bool)`

GetLimitRecordsOk returns a tuple with the LimitRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitRecords

`func (o *OpenMetadataServerConnection) SetLimitRecords(v int32)`

SetLimitRecords sets LimitRecords field to given value.

### HasLimitRecords

`func (o *OpenMetadataServerConnection) HasLimitRecords() bool`

HasLimitRecords returns a boolean if a field has been set.

### GetSecurityConfig

`func (o *OpenMetadataServerConnection) GetSecurityConfig() map[string]interface{}`

GetSecurityConfig returns the SecurityConfig field if non-nil, zero value otherwise.

### GetSecurityConfigOk

`func (o *OpenMetadataServerConnection) GetSecurityConfigOk() (*map[string]interface{}, bool)`

GetSecurityConfigOk returns a tuple with the SecurityConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityConfig

`func (o *OpenMetadataServerConnection) SetSecurityConfig(v map[string]interface{})`

SetSecurityConfig sets SecurityConfig field to given value.

### HasSecurityConfig

`func (o *OpenMetadataServerConnection) HasSecurityConfig() bool`

HasSecurityConfig returns a boolean if a field has been set.

### GetSupportsMetadataExtraction

`func (o *OpenMetadataServerConnection) GetSupportsMetadataExtraction() bool`

GetSupportsMetadataExtraction returns the SupportsMetadataExtraction field if non-nil, zero value otherwise.

### GetSupportsMetadataExtractionOk

`func (o *OpenMetadataServerConnection) GetSupportsMetadataExtractionOk() (*bool, bool)`

GetSupportsMetadataExtractionOk returns a tuple with the SupportsMetadataExtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMetadataExtraction

`func (o *OpenMetadataServerConnection) SetSupportsMetadataExtraction(v bool)`

SetSupportsMetadataExtraction sets SupportsMetadataExtraction field to given value.

### HasSupportsMetadataExtraction

`func (o *OpenMetadataServerConnection) HasSupportsMetadataExtraction() bool`

HasSupportsMetadataExtraction returns a boolean if a field has been set.

### GetType

`func (o *OpenMetadataServerConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpenMetadataServerConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpenMetadataServerConnection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OpenMetadataServerConnection) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


