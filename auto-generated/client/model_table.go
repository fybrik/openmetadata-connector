/*
OpenMetadata Apis

# Overview  OpenMetadata supports REST APIs for getting metadata in and out of metadata store. The API resources are grouped under following categories: - **Data assets** - includes resources for data entities, such as `databases`, `tables`, and `topics`. Resources for data assets created from data, such as `dashboards`, `reports`, `metrics`, and `ML Features` are part of this collection. `pipelines`, `notebooks`, etc. that are used for creating data assets are also available as resources as of this collection. - **Teams and Users** - includes `users`, `teams`, a special type of user called `bots` that performs many automated tasks such as ingestion. - **Services** - are services that OpenMetadata integrates with. Currently `databaseService` is the only service under this collection that represents data sources. In the future, services related to Dashboards, Reports, ETL pipelines will be added under this collection. - **Glossary** - OpenMetadata supports hierarchical tags that can be used to build business vocabulary to describe and classify data available under `tags` resource.

API version: 0.11.0
Contact: openmetadata-dev@googlegroups.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Table struct for Table
type Table struct {
	ChangeDescription  *ChangeDescription     `json:"changeDescription,omitempty"`
	Columns            []Column               `json:"columns"`
	DataModel          *DataModel             `json:"dataModel,omitempty"`
	Database           *EntityReference       `json:"database,omitempty"`
	DatabaseSchema     *EntityReference       `json:"databaseSchema,omitempty"`
	Deleted            *bool                  `json:"deleted,omitempty"`
	Description        *string                `json:"description,omitempty"`
	DisplayName        *string                `json:"displayName,omitempty"`
	Extension          map[string]interface{} `json:"extension,omitempty"`
	Followers          []EntityReference      `json:"followers,omitempty"`
	FullyQualifiedName *string                `json:"fullyQualifiedName,omitempty"`
	Href               *string                `json:"href,omitempty"`
	Id                 string                 `json:"id"`
	Joins              *TableJoins            `json:"joins,omitempty"`
	Location           *EntityReference       `json:"location,omitempty"`
	Name               string                 `json:"name"`
	Owner              *EntityReference       `json:"owner,omitempty"`
	ProfileQuery       *string                `json:"profileQuery,omitempty"`
	ProfileSample      *float64               `json:"profileSample,omitempty"`
	SampleData         *TableData             `json:"sampleData,omitempty"`
	Service            *EntityReference       `json:"service,omitempty"`
	ServiceType        *string                `json:"serviceType,omitempty"`
	TableConstraints   []TableConstraint      `json:"tableConstraints,omitempty"`
	TablePartition     *TablePartition        `json:"tablePartition,omitempty"`
	TableProfile       []TableProfile         `json:"tableProfile,omitempty"`
	TableQueries       []SQLQuery             `json:"tableQueries,omitempty"`
	TableTests         []TableTest            `json:"tableTests,omitempty"`
	TableType          *string                `json:"tableType,omitempty"`
	Tags               []TagLabel             `json:"tags,omitempty"`
	UpdatedAt          *int64                 `json:"updatedAt,omitempty"`
	UpdatedBy          *string                `json:"updatedBy,omitempty"`
	UsageSummary       *UsageDetails          `json:"usageSummary,omitempty"`
	Version            *float64               `json:"version,omitempty"`
	ViewDefinition     *string                `json:"viewDefinition,omitempty"`
}

// NewTable instantiates a new Table object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTable(columns []Column, id string, name string) *Table {
	this := Table{}
	this.Columns = columns
	this.Id = id
	this.Name = name
	return &this
}

// NewTableWithDefaults instantiates a new Table object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableWithDefaults() *Table {
	this := Table{}
	return &this
}

// GetChangeDescription returns the ChangeDescription field value if set, zero value otherwise.
func (o *Table) GetChangeDescription() ChangeDescription {
	if o == nil || o.ChangeDescription == nil {
		var ret ChangeDescription
		return ret
	}
	return *o.ChangeDescription
}

// GetChangeDescriptionOk returns a tuple with the ChangeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetChangeDescriptionOk() (*ChangeDescription, bool) {
	if o == nil || o.ChangeDescription == nil {
		return nil, false
	}
	return o.ChangeDescription, true
}

// HasChangeDescription returns a boolean if a field has been set.
func (o *Table) HasChangeDescription() bool {
	if o != nil && o.ChangeDescription != nil {
		return true
	}

	return false
}

// SetChangeDescription gets a reference to the given ChangeDescription and assigns it to the ChangeDescription field.
func (o *Table) SetChangeDescription(v ChangeDescription) {
	o.ChangeDescription = &v
}

// GetColumns returns the Columns field value
func (o *Table) GetColumns() []Column {
	if o == nil {
		var ret []Column
		return ret
	}

	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *Table) GetColumnsOk() ([]Column, bool) {
	if o == nil {
		return nil, false
	}
	return o.Columns, true
}

// SetColumns sets field value
func (o *Table) SetColumns(v []Column) {
	o.Columns = v
}

// GetDataModel returns the DataModel field value if set, zero value otherwise.
func (o *Table) GetDataModel() DataModel {
	if o == nil || o.DataModel == nil {
		var ret DataModel
		return ret
	}
	return *o.DataModel
}

// GetDataModelOk returns a tuple with the DataModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetDataModelOk() (*DataModel, bool) {
	if o == nil || o.DataModel == nil {
		return nil, false
	}
	return o.DataModel, true
}

// HasDataModel returns a boolean if a field has been set.
func (o *Table) HasDataModel() bool {
	if o != nil && o.DataModel != nil {
		return true
	}

	return false
}

// SetDataModel gets a reference to the given DataModel and assigns it to the DataModel field.
func (o *Table) SetDataModel(v DataModel) {
	o.DataModel = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *Table) GetDatabase() EntityReference {
	if o == nil || o.Database == nil {
		var ret EntityReference
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetDatabaseOk() (*EntityReference, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *Table) HasDatabase() bool {
	if o != nil && o.Database != nil {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given EntityReference and assigns it to the Database field.
func (o *Table) SetDatabase(v EntityReference) {
	o.Database = &v
}

// GetDatabaseSchema returns the DatabaseSchema field value if set, zero value otherwise.
func (o *Table) GetDatabaseSchema() EntityReference {
	if o == nil || o.DatabaseSchema == nil {
		var ret EntityReference
		return ret
	}
	return *o.DatabaseSchema
}

// GetDatabaseSchemaOk returns a tuple with the DatabaseSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetDatabaseSchemaOk() (*EntityReference, bool) {
	if o == nil || o.DatabaseSchema == nil {
		return nil, false
	}
	return o.DatabaseSchema, true
}

// HasDatabaseSchema returns a boolean if a field has been set.
func (o *Table) HasDatabaseSchema() bool {
	if o != nil && o.DatabaseSchema != nil {
		return true
	}

	return false
}

// SetDatabaseSchema gets a reference to the given EntityReference and assigns it to the DatabaseSchema field.
func (o *Table) SetDatabaseSchema(v EntityReference) {
	o.DatabaseSchema = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Table) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Table) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *Table) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Table) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Table) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Table) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Table) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Table) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Table) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *Table) GetExtension() map[string]interface{} {
	if o == nil || o.Extension == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetExtensionOk() (map[string]interface{}, bool) {
	if o == nil || o.Extension == nil {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *Table) HasExtension() bool {
	if o != nil && o.Extension != nil {
		return true
	}

	return false
}

// SetExtension gets a reference to the given map[string]interface{} and assigns it to the Extension field.
func (o *Table) SetExtension(v map[string]interface{}) {
	o.Extension = v
}

// GetFollowers returns the Followers field value if set, zero value otherwise.
func (o *Table) GetFollowers() []EntityReference {
	if o == nil || o.Followers == nil {
		var ret []EntityReference
		return ret
	}
	return o.Followers
}

// GetFollowersOk returns a tuple with the Followers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetFollowersOk() ([]EntityReference, bool) {
	if o == nil || o.Followers == nil {
		return nil, false
	}
	return o.Followers, true
}

// HasFollowers returns a boolean if a field has been set.
func (o *Table) HasFollowers() bool {
	if o != nil && o.Followers != nil {
		return true
	}

	return false
}

// SetFollowers gets a reference to the given []EntityReference and assigns it to the Followers field.
func (o *Table) SetFollowers(v []EntityReference) {
	o.Followers = v
}

// GetFullyQualifiedName returns the FullyQualifiedName field value if set, zero value otherwise.
func (o *Table) GetFullyQualifiedName() string {
	if o == nil || o.FullyQualifiedName == nil {
		var ret string
		return ret
	}
	return *o.FullyQualifiedName
}

// GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetFullyQualifiedNameOk() (*string, bool) {
	if o == nil || o.FullyQualifiedName == nil {
		return nil, false
	}
	return o.FullyQualifiedName, true
}

// HasFullyQualifiedName returns a boolean if a field has been set.
func (o *Table) HasFullyQualifiedName() bool {
	if o != nil && o.FullyQualifiedName != nil {
		return true
	}

	return false
}

// SetFullyQualifiedName gets a reference to the given string and assigns it to the FullyQualifiedName field.
func (o *Table) SetFullyQualifiedName(v string) {
	o.FullyQualifiedName = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Table) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Table) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Table) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value
func (o *Table) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Table) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Table) SetId(v string) {
	o.Id = v
}

// GetJoins returns the Joins field value if set, zero value otherwise.
func (o *Table) GetJoins() TableJoins {
	if o == nil || o.Joins == nil {
		var ret TableJoins
		return ret
	}
	return *o.Joins
}

// GetJoinsOk returns a tuple with the Joins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetJoinsOk() (*TableJoins, bool) {
	if o == nil || o.Joins == nil {
		return nil, false
	}
	return o.Joins, true
}

// HasJoins returns a boolean if a field has been set.
func (o *Table) HasJoins() bool {
	if o != nil && o.Joins != nil {
		return true
	}

	return false
}

// SetJoins gets a reference to the given TableJoins and assigns it to the Joins field.
func (o *Table) SetJoins(v TableJoins) {
	o.Joins = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Table) GetLocation() EntityReference {
	if o == nil || o.Location == nil {
		var ret EntityReference
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetLocationOk() (*EntityReference, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Table) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given EntityReference and assigns it to the Location field.
func (o *Table) SetLocation(v EntityReference) {
	o.Location = &v
}

// GetName returns the Name field value
func (o *Table) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Table) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Table) SetName(v string) {
	o.Name = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Table) GetOwner() EntityReference {
	if o == nil || o.Owner == nil {
		var ret EntityReference
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetOwnerOk() (*EntityReference, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Table) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given EntityReference and assigns it to the Owner field.
func (o *Table) SetOwner(v EntityReference) {
	o.Owner = &v
}

// GetProfileQuery returns the ProfileQuery field value if set, zero value otherwise.
func (o *Table) GetProfileQuery() string {
	if o == nil || o.ProfileQuery == nil {
		var ret string
		return ret
	}
	return *o.ProfileQuery
}

// GetProfileQueryOk returns a tuple with the ProfileQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetProfileQueryOk() (*string, bool) {
	if o == nil || o.ProfileQuery == nil {
		return nil, false
	}
	return o.ProfileQuery, true
}

// HasProfileQuery returns a boolean if a field has been set.
func (o *Table) HasProfileQuery() bool {
	if o != nil && o.ProfileQuery != nil {
		return true
	}

	return false
}

// SetProfileQuery gets a reference to the given string and assigns it to the ProfileQuery field.
func (o *Table) SetProfileQuery(v string) {
	o.ProfileQuery = &v
}

// GetProfileSample returns the ProfileSample field value if set, zero value otherwise.
func (o *Table) GetProfileSample() float64 {
	if o == nil || o.ProfileSample == nil {
		var ret float64
		return ret
	}
	return *o.ProfileSample
}

// GetProfileSampleOk returns a tuple with the ProfileSample field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetProfileSampleOk() (*float64, bool) {
	if o == nil || o.ProfileSample == nil {
		return nil, false
	}
	return o.ProfileSample, true
}

// HasProfileSample returns a boolean if a field has been set.
func (o *Table) HasProfileSample() bool {
	if o != nil && o.ProfileSample != nil {
		return true
	}

	return false
}

// SetProfileSample gets a reference to the given float64 and assigns it to the ProfileSample field.
func (o *Table) SetProfileSample(v float64) {
	o.ProfileSample = &v
}

// GetSampleData returns the SampleData field value if set, zero value otherwise.
func (o *Table) GetSampleData() TableData {
	if o == nil || o.SampleData == nil {
		var ret TableData
		return ret
	}
	return *o.SampleData
}

// GetSampleDataOk returns a tuple with the SampleData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetSampleDataOk() (*TableData, bool) {
	if o == nil || o.SampleData == nil {
		return nil, false
	}
	return o.SampleData, true
}

// HasSampleData returns a boolean if a field has been set.
func (o *Table) HasSampleData() bool {
	if o != nil && o.SampleData != nil {
		return true
	}

	return false
}

// SetSampleData gets a reference to the given TableData and assigns it to the SampleData field.
func (o *Table) SetSampleData(v TableData) {
	o.SampleData = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *Table) GetService() EntityReference {
	if o == nil || o.Service == nil {
		var ret EntityReference
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetServiceOk() (*EntityReference, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *Table) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given EntityReference and assigns it to the Service field.
func (o *Table) SetService(v EntityReference) {
	o.Service = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *Table) GetServiceType() string {
	if o == nil || o.ServiceType == nil {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetServiceTypeOk() (*string, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *Table) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *Table) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetTableConstraints returns the TableConstraints field value if set, zero value otherwise.
func (o *Table) GetTableConstraints() []TableConstraint {
	if o == nil || o.TableConstraints == nil {
		var ret []TableConstraint
		return ret
	}
	return o.TableConstraints
}

// GetTableConstraintsOk returns a tuple with the TableConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetTableConstraintsOk() ([]TableConstraint, bool) {
	if o == nil || o.TableConstraints == nil {
		return nil, false
	}
	return o.TableConstraints, true
}

// HasTableConstraints returns a boolean if a field has been set.
func (o *Table) HasTableConstraints() bool {
	if o != nil && o.TableConstraints != nil {
		return true
	}

	return false
}

// SetTableConstraints gets a reference to the given []TableConstraint and assigns it to the TableConstraints field.
func (o *Table) SetTableConstraints(v []TableConstraint) {
	o.TableConstraints = v
}

// GetTablePartition returns the TablePartition field value if set, zero value otherwise.
func (o *Table) GetTablePartition() TablePartition {
	if o == nil || o.TablePartition == nil {
		var ret TablePartition
		return ret
	}
	return *o.TablePartition
}

// GetTablePartitionOk returns a tuple with the TablePartition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetTablePartitionOk() (*TablePartition, bool) {
	if o == nil || o.TablePartition == nil {
		return nil, false
	}
	return o.TablePartition, true
}

// HasTablePartition returns a boolean if a field has been set.
func (o *Table) HasTablePartition() bool {
	if o != nil && o.TablePartition != nil {
		return true
	}

	return false
}

// SetTablePartition gets a reference to the given TablePartition and assigns it to the TablePartition field.
func (o *Table) SetTablePartition(v TablePartition) {
	o.TablePartition = &v
}

// GetTableProfile returns the TableProfile field value if set, zero value otherwise.
func (o *Table) GetTableProfile() []TableProfile {
	if o == nil || o.TableProfile == nil {
		var ret []TableProfile
		return ret
	}
	return o.TableProfile
}

// GetTableProfileOk returns a tuple with the TableProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetTableProfileOk() ([]TableProfile, bool) {
	if o == nil || o.TableProfile == nil {
		return nil, false
	}
	return o.TableProfile, true
}

// HasTableProfile returns a boolean if a field has been set.
func (o *Table) HasTableProfile() bool {
	if o != nil && o.TableProfile != nil {
		return true
	}

	return false
}

// SetTableProfile gets a reference to the given []TableProfile and assigns it to the TableProfile field.
func (o *Table) SetTableProfile(v []TableProfile) {
	o.TableProfile = v
}

// GetTableQueries returns the TableQueries field value if set, zero value otherwise.
func (o *Table) GetTableQueries() []SQLQuery {
	if o == nil || o.TableQueries == nil {
		var ret []SQLQuery
		return ret
	}
	return o.TableQueries
}

// GetTableQueriesOk returns a tuple with the TableQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetTableQueriesOk() ([]SQLQuery, bool) {
	if o == nil || o.TableQueries == nil {
		return nil, false
	}
	return o.TableQueries, true
}

// HasTableQueries returns a boolean if a field has been set.
func (o *Table) HasTableQueries() bool {
	if o != nil && o.TableQueries != nil {
		return true
	}

	return false
}

// SetTableQueries gets a reference to the given []SQLQuery and assigns it to the TableQueries field.
func (o *Table) SetTableQueries(v []SQLQuery) {
	o.TableQueries = v
}

// GetTableTests returns the TableTests field value if set, zero value otherwise.
func (o *Table) GetTableTests() []TableTest {
	if o == nil || o.TableTests == nil {
		var ret []TableTest
		return ret
	}
	return o.TableTests
}

// GetTableTestsOk returns a tuple with the TableTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetTableTestsOk() ([]TableTest, bool) {
	if o == nil || o.TableTests == nil {
		return nil, false
	}
	return o.TableTests, true
}

// HasTableTests returns a boolean if a field has been set.
func (o *Table) HasTableTests() bool {
	if o != nil && o.TableTests != nil {
		return true
	}

	return false
}

// SetTableTests gets a reference to the given []TableTest and assigns it to the TableTests field.
func (o *Table) SetTableTests(v []TableTest) {
	o.TableTests = v
}

// GetTableType returns the TableType field value if set, zero value otherwise.
func (o *Table) GetTableType() string {
	if o == nil || o.TableType == nil {
		var ret string
		return ret
	}
	return *o.TableType
}

// GetTableTypeOk returns a tuple with the TableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetTableTypeOk() (*string, bool) {
	if o == nil || o.TableType == nil {
		return nil, false
	}
	return o.TableType, true
}

// HasTableType returns a boolean if a field has been set.
func (o *Table) HasTableType() bool {
	if o != nil && o.TableType != nil {
		return true
	}

	return false
}

// SetTableType gets a reference to the given string and assigns it to the TableType field.
func (o *Table) SetTableType(v string) {
	o.TableType = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Table) GetTags() []TagLabel {
	if o == nil || o.Tags == nil {
		var ret []TagLabel
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetTagsOk() ([]TagLabel, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Table) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagLabel and assigns it to the Tags field.
func (o *Table) SetTags(v []TagLabel) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Table) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Table) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *Table) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *Table) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *Table) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *Table) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetUsageSummary returns the UsageSummary field value if set, zero value otherwise.
func (o *Table) GetUsageSummary() UsageDetails {
	if o == nil || o.UsageSummary == nil {
		var ret UsageDetails
		return ret
	}
	return *o.UsageSummary
}

// GetUsageSummaryOk returns a tuple with the UsageSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetUsageSummaryOk() (*UsageDetails, bool) {
	if o == nil || o.UsageSummary == nil {
		return nil, false
	}
	return o.UsageSummary, true
}

// HasUsageSummary returns a boolean if a field has been set.
func (o *Table) HasUsageSummary() bool {
	if o != nil && o.UsageSummary != nil {
		return true
	}

	return false
}

// SetUsageSummary gets a reference to the given UsageDetails and assigns it to the UsageSummary field.
func (o *Table) SetUsageSummary(v UsageDetails) {
	o.UsageSummary = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Table) GetVersion() float64 {
	if o == nil || o.Version == nil {
		var ret float64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetVersionOk() (*float64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Table) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given float64 and assigns it to the Version field.
func (o *Table) SetVersion(v float64) {
	o.Version = &v
}

// GetViewDefinition returns the ViewDefinition field value if set, zero value otherwise.
func (o *Table) GetViewDefinition() string {
	if o == nil || o.ViewDefinition == nil {
		var ret string
		return ret
	}
	return *o.ViewDefinition
}

// GetViewDefinitionOk returns a tuple with the ViewDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table) GetViewDefinitionOk() (*string, bool) {
	if o == nil || o.ViewDefinition == nil {
		return nil, false
	}
	return o.ViewDefinition, true
}

// HasViewDefinition returns a boolean if a field has been set.
func (o *Table) HasViewDefinition() bool {
	if o != nil && o.ViewDefinition != nil {
		return true
	}

	return false
}

// SetViewDefinition gets a reference to the given string and assigns it to the ViewDefinition field.
func (o *Table) SetViewDefinition(v string) {
	o.ViewDefinition = &v
}

func (o Table) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangeDescription != nil {
		toSerialize["changeDescription"] = o.ChangeDescription
	}
	if true {
		toSerialize["columns"] = o.Columns
	}
	if o.DataModel != nil {
		toSerialize["dataModel"] = o.DataModel
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.DatabaseSchema != nil {
		toSerialize["databaseSchema"] = o.DatabaseSchema
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Extension != nil {
		toSerialize["extension"] = o.Extension
	}
	if o.Followers != nil {
		toSerialize["followers"] = o.Followers
	}
	if o.FullyQualifiedName != nil {
		toSerialize["fullyQualifiedName"] = o.FullyQualifiedName
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Joins != nil {
		toSerialize["joins"] = o.Joins
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.ProfileQuery != nil {
		toSerialize["profileQuery"] = o.ProfileQuery
	}
	if o.ProfileSample != nil {
		toSerialize["profileSample"] = o.ProfileSample
	}
	if o.SampleData != nil {
		toSerialize["sampleData"] = o.SampleData
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.ServiceType != nil {
		toSerialize["serviceType"] = o.ServiceType
	}
	if o.TableConstraints != nil {
		toSerialize["tableConstraints"] = o.TableConstraints
	}
	if o.TablePartition != nil {
		toSerialize["tablePartition"] = o.TablePartition
	}
	if o.TableProfile != nil {
		toSerialize["tableProfile"] = o.TableProfile
	}
	if o.TableQueries != nil {
		toSerialize["tableQueries"] = o.TableQueries
	}
	if o.TableTests != nil {
		toSerialize["tableTests"] = o.TableTests
	}
	if o.TableType != nil {
		toSerialize["tableType"] = o.TableType
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if o.UsageSummary != nil {
		toSerialize["usageSummary"] = o.UsageSummary
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.ViewDefinition != nil {
		toSerialize["viewDefinition"] = o.ViewDefinition
	}
	return json.Marshal(toSerialize)
}

type NullableTable struct {
	value *Table
	isSet bool
}

func (v NullableTable) Get() *Table {
	return v.value
}

func (v *NullableTable) Set(val *Table) {
	v.value = val
	v.isSet = true
}

func (v NullableTable) IsSet() bool {
	return v.isSet
}

func (v *NullableTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTable(val *Table) *NullableTable {
	return &NullableTable{value: val, isSet: true}
}

func (v NullableTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
