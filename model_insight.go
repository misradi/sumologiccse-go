/*
Sumo Logic CSE API

 https://help.sumologic.com/APIs 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologiccse

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the Insight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Insight{}

// Insight struct for Insight
type Insight struct {
	Id string `json:"id"`
	// A human-readable ID in the format \"INSIGHT-542\". This is technically nullable, but in reality it will always be populated in every query other than the cross-type search query.
	ReadableId string `json:"readableId"`
	Name string `json:"name"`
	Description string `json:"description"`
	Timestamp time.Time `json:"timestamp"`
	Source string `json:"source"`
	// The user that this Insight is assigned to
	AssignedTo *string `json:"assignedTo,omitempty"`
	// The team that this Insight is assigned to
	TeamAssignedTo *string `json:"teamAssignedTo,omitempty"`
	Created time.Time `json:"created"`
	Closed *time.Time `json:"closed,omitempty"`
	ClosedBy *string `json:"closedBy,omitempty"`
	Severity string `json:"severity"`
	// A 0-100 value of the ML-based confidence score for the Insight
	Confidence *float64 `json:"confidence,omitempty"`
	Assignee *GetInsightsResponseDataObjectsInnerAssignee `json:"assignee,omitempty"`
	Status GetInsightsResponseDataObjectsInnerStatus `json:"status"`
	Resolution *string `json:"resolution,omitempty"`
	SubResolution *string `json:"subResolution,omitempty"`
	Entity GetInsightsResponseDataObjectsInnerEntity `json:"entity"`
	Signals []GetInsightsResponseDataObjectsInnerSignalsInner `json:"signals"`
	InvolvedEntities []GetInsightsResponseDataObjectsInnerSignalsInnerEntity `json:"involvedEntities"`
	Artifacts []GetInsightsResponseDataObjectsInnerArtifactsInner `json:"artifacts"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	OrgId string `json:"orgId"`
	TimeToDetection *int32 `json:"timeToDetection,omitempty"`
	TimeToResponse *int32 `json:"timeToResponse,omitempty"`
	TimeToRemediation *int32 `json:"timeToRemediation,omitempty"`
	Tags []string `json:"tags"`
}

type _Insight Insight

// NewInsight instantiates a new Insight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsight(id string, readableId string, name string, description string, timestamp time.Time, source string, created time.Time, severity string, status GetInsightsResponseDataObjectsInnerStatus, entity GetInsightsResponseDataObjectsInnerEntity, signals []GetInsightsResponseDataObjectsInnerSignalsInner, involvedEntities []GetInsightsResponseDataObjectsInnerSignalsInnerEntity, artifacts []GetInsightsResponseDataObjectsInnerArtifactsInner, orgId string, tags []string) *Insight {
	this := Insight{}
	this.Id = id
	this.ReadableId = readableId
	this.Name = name
	this.Description = description
	this.Timestamp = timestamp
	this.Source = source
	this.Created = created
	this.Severity = severity
	this.Status = status
	this.Entity = entity
	this.Signals = signals
	this.InvolvedEntities = involvedEntities
	this.Artifacts = artifacts
	this.OrgId = orgId
	this.Tags = tags
	return &this
}

// NewInsightWithDefaults instantiates a new Insight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsightWithDefaults() *Insight {
	this := Insight{}
	return &this
}

// GetId returns the Id field value
func (o *Insight) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Insight) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Insight) SetId(v string) {
	o.Id = v
}

// GetReadableId returns the ReadableId field value
func (o *Insight) GetReadableId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReadableId
}

// GetReadableIdOk returns a tuple with the ReadableId field value
// and a boolean to check if the value has been set.
func (o *Insight) GetReadableIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadableId, true
}

// SetReadableId sets field value
func (o *Insight) SetReadableId(v string) {
	o.ReadableId = v
}

// GetName returns the Name field value
func (o *Insight) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Insight) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Insight) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Insight) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Insight) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Insight) SetDescription(v string) {
	o.Description = v
}

// GetTimestamp returns the Timestamp field value
func (o *Insight) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *Insight) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *Insight) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetSource returns the Source field value
func (o *Insight) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *Insight) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *Insight) SetSource(v string) {
	o.Source = v
}

// GetAssignedTo returns the AssignedTo field value if set, zero value otherwise.
func (o *Insight) GetAssignedTo() string {
	if o == nil || IsNil(o.AssignedTo) {
		var ret string
		return ret
	}
	return *o.AssignedTo
}

// GetAssignedToOk returns a tuple with the AssignedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetAssignedToOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedTo) {
		return nil, false
	}
	return o.AssignedTo, true
}

// HasAssignedTo returns a boolean if a field has been set.
func (o *Insight) HasAssignedTo() bool {
	if o != nil && !IsNil(o.AssignedTo) {
		return true
	}

	return false
}

// SetAssignedTo gets a reference to the given string and assigns it to the AssignedTo field.
func (o *Insight) SetAssignedTo(v string) {
	o.AssignedTo = &v
}

// GetTeamAssignedTo returns the TeamAssignedTo field value if set, zero value otherwise.
func (o *Insight) GetTeamAssignedTo() string {
	if o == nil || IsNil(o.TeamAssignedTo) {
		var ret string
		return ret
	}
	return *o.TeamAssignedTo
}

// GetTeamAssignedToOk returns a tuple with the TeamAssignedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetTeamAssignedToOk() (*string, bool) {
	if o == nil || IsNil(o.TeamAssignedTo) {
		return nil, false
	}
	return o.TeamAssignedTo, true
}

// HasTeamAssignedTo returns a boolean if a field has been set.
func (o *Insight) HasTeamAssignedTo() bool {
	if o != nil && !IsNil(o.TeamAssignedTo) {
		return true
	}

	return false
}

// SetTeamAssignedTo gets a reference to the given string and assigns it to the TeamAssignedTo field.
func (o *Insight) SetTeamAssignedTo(v string) {
	o.TeamAssignedTo = &v
}

// GetCreated returns the Created field value
func (o *Insight) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Insight) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Insight) SetCreated(v time.Time) {
	o.Created = v
}

// GetClosed returns the Closed field value if set, zero value otherwise.
func (o *Insight) GetClosed() time.Time {
	if o == nil || IsNil(o.Closed) {
		var ret time.Time
		return ret
	}
	return *o.Closed
}

// GetClosedOk returns a tuple with the Closed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetClosedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Closed) {
		return nil, false
	}
	return o.Closed, true
}

// HasClosed returns a boolean if a field has been set.
func (o *Insight) HasClosed() bool {
	if o != nil && !IsNil(o.Closed) {
		return true
	}

	return false
}

// SetClosed gets a reference to the given time.Time and assigns it to the Closed field.
func (o *Insight) SetClosed(v time.Time) {
	o.Closed = &v
}

// GetClosedBy returns the ClosedBy field value if set, zero value otherwise.
func (o *Insight) GetClosedBy() string {
	if o == nil || IsNil(o.ClosedBy) {
		var ret string
		return ret
	}
	return *o.ClosedBy
}

// GetClosedByOk returns a tuple with the ClosedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetClosedByOk() (*string, bool) {
	if o == nil || IsNil(o.ClosedBy) {
		return nil, false
	}
	return o.ClosedBy, true
}

// HasClosedBy returns a boolean if a field has been set.
func (o *Insight) HasClosedBy() bool {
	if o != nil && !IsNil(o.ClosedBy) {
		return true
	}

	return false
}

// SetClosedBy gets a reference to the given string and assigns it to the ClosedBy field.
func (o *Insight) SetClosedBy(v string) {
	o.ClosedBy = &v
}

// GetSeverity returns the Severity field value
func (o *Insight) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *Insight) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *Insight) SetSeverity(v string) {
	o.Severity = v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *Insight) GetConfidence() float64 {
	if o == nil || IsNil(o.Confidence) {
		var ret float64
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetConfidenceOk() (*float64, bool) {
	if o == nil || IsNil(o.Confidence) {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *Insight) HasConfidence() bool {
	if o != nil && !IsNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given float64 and assigns it to the Confidence field.
func (o *Insight) SetConfidence(v float64) {
	o.Confidence = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *Insight) GetAssignee() GetInsightsResponseDataObjectsInnerAssignee {
	if o == nil || IsNil(o.Assignee) {
		var ret GetInsightsResponseDataObjectsInnerAssignee
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetAssigneeOk() (*GetInsightsResponseDataObjectsInnerAssignee, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *Insight) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given GetInsightsResponseDataObjectsInnerAssignee and assigns it to the Assignee field.
func (o *Insight) SetAssignee(v GetInsightsResponseDataObjectsInnerAssignee) {
	o.Assignee = &v
}

// GetStatus returns the Status field value
func (o *Insight) GetStatus() GetInsightsResponseDataObjectsInnerStatus {
	if o == nil {
		var ret GetInsightsResponseDataObjectsInnerStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Insight) GetStatusOk() (*GetInsightsResponseDataObjectsInnerStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Insight) SetStatus(v GetInsightsResponseDataObjectsInnerStatus) {
	o.Status = v
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *Insight) GetResolution() string {
	if o == nil || IsNil(o.Resolution) {
		var ret string
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetResolutionOk() (*string, bool) {
	if o == nil || IsNil(o.Resolution) {
		return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *Insight) HasResolution() bool {
	if o != nil && !IsNil(o.Resolution) {
		return true
	}

	return false
}

// SetResolution gets a reference to the given string and assigns it to the Resolution field.
func (o *Insight) SetResolution(v string) {
	o.Resolution = &v
}

// GetSubResolution returns the SubResolution field value if set, zero value otherwise.
func (o *Insight) GetSubResolution() string {
	if o == nil || IsNil(o.SubResolution) {
		var ret string
		return ret
	}
	return *o.SubResolution
}

// GetSubResolutionOk returns a tuple with the SubResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetSubResolutionOk() (*string, bool) {
	if o == nil || IsNil(o.SubResolution) {
		return nil, false
	}
	return o.SubResolution, true
}

// HasSubResolution returns a boolean if a field has been set.
func (o *Insight) HasSubResolution() bool {
	if o != nil && !IsNil(o.SubResolution) {
		return true
	}

	return false
}

// SetSubResolution gets a reference to the given string and assigns it to the SubResolution field.
func (o *Insight) SetSubResolution(v string) {
	o.SubResolution = &v
}

// GetEntity returns the Entity field value
func (o *Insight) GetEntity() GetInsightsResponseDataObjectsInnerEntity {
	if o == nil {
		var ret GetInsightsResponseDataObjectsInnerEntity
		return ret
	}

	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value
// and a boolean to check if the value has been set.
func (o *Insight) GetEntityOk() (*GetInsightsResponseDataObjectsInnerEntity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entity, true
}

// SetEntity sets field value
func (o *Insight) SetEntity(v GetInsightsResponseDataObjectsInnerEntity) {
	o.Entity = v
}

// GetSignals returns the Signals field value
func (o *Insight) GetSignals() []GetInsightsResponseDataObjectsInnerSignalsInner {
	if o == nil {
		var ret []GetInsightsResponseDataObjectsInnerSignalsInner
		return ret
	}

	return o.Signals
}

// GetSignalsOk returns a tuple with the Signals field value
// and a boolean to check if the value has been set.
func (o *Insight) GetSignalsOk() ([]GetInsightsResponseDataObjectsInnerSignalsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signals, true
}

// SetSignals sets field value
func (o *Insight) SetSignals(v []GetInsightsResponseDataObjectsInnerSignalsInner) {
	o.Signals = v
}

// GetInvolvedEntities returns the InvolvedEntities field value
func (o *Insight) GetInvolvedEntities() []GetInsightsResponseDataObjectsInnerSignalsInnerEntity {
	if o == nil {
		var ret []GetInsightsResponseDataObjectsInnerSignalsInnerEntity
		return ret
	}

	return o.InvolvedEntities
}

// GetInvolvedEntitiesOk returns a tuple with the InvolvedEntities field value
// and a boolean to check if the value has been set.
func (o *Insight) GetInvolvedEntitiesOk() ([]GetInsightsResponseDataObjectsInnerSignalsInnerEntity, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvolvedEntities, true
}

// SetInvolvedEntities sets field value
func (o *Insight) SetInvolvedEntities(v []GetInsightsResponseDataObjectsInnerSignalsInnerEntity) {
	o.InvolvedEntities = v
}

// GetArtifacts returns the Artifacts field value
func (o *Insight) GetArtifacts() []GetInsightsResponseDataObjectsInnerArtifactsInner {
	if o == nil {
		var ret []GetInsightsResponseDataObjectsInnerArtifactsInner
		return ret
	}

	return o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value
// and a boolean to check if the value has been set.
func (o *Insight) GetArtifactsOk() ([]GetInsightsResponseDataObjectsInnerArtifactsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Artifacts, true
}

// SetArtifacts sets field value
func (o *Insight) SetArtifacts(v []GetInsightsResponseDataObjectsInnerArtifactsInner) {
	o.Artifacts = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Insight) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Insight) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Insight) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *Insight) GetLastUpdatedBy() string {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *Insight) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *Insight) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetOrgId returns the OrgId field value
func (o *Insight) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *Insight) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *Insight) SetOrgId(v string) {
	o.OrgId = v
}

// GetTimeToDetection returns the TimeToDetection field value if set, zero value otherwise.
func (o *Insight) GetTimeToDetection() int32 {
	if o == nil || IsNil(o.TimeToDetection) {
		var ret int32
		return ret
	}
	return *o.TimeToDetection
}

// GetTimeToDetectionOk returns a tuple with the TimeToDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetTimeToDetectionOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeToDetection) {
		return nil, false
	}
	return o.TimeToDetection, true
}

// HasTimeToDetection returns a boolean if a field has been set.
func (o *Insight) HasTimeToDetection() bool {
	if o != nil && !IsNil(o.TimeToDetection) {
		return true
	}

	return false
}

// SetTimeToDetection gets a reference to the given int32 and assigns it to the TimeToDetection field.
func (o *Insight) SetTimeToDetection(v int32) {
	o.TimeToDetection = &v
}

// GetTimeToResponse returns the TimeToResponse field value if set, zero value otherwise.
func (o *Insight) GetTimeToResponse() int32 {
	if o == nil || IsNil(o.TimeToResponse) {
		var ret int32
		return ret
	}
	return *o.TimeToResponse
}

// GetTimeToResponseOk returns a tuple with the TimeToResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetTimeToResponseOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeToResponse) {
		return nil, false
	}
	return o.TimeToResponse, true
}

// HasTimeToResponse returns a boolean if a field has been set.
func (o *Insight) HasTimeToResponse() bool {
	if o != nil && !IsNil(o.TimeToResponse) {
		return true
	}

	return false
}

// SetTimeToResponse gets a reference to the given int32 and assigns it to the TimeToResponse field.
func (o *Insight) SetTimeToResponse(v int32) {
	o.TimeToResponse = &v
}

// GetTimeToRemediation returns the TimeToRemediation field value if set, zero value otherwise.
func (o *Insight) GetTimeToRemediation() int32 {
	if o == nil || IsNil(o.TimeToRemediation) {
		var ret int32
		return ret
	}
	return *o.TimeToRemediation
}

// GetTimeToRemediationOk returns a tuple with the TimeToRemediation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Insight) GetTimeToRemediationOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeToRemediation) {
		return nil, false
	}
	return o.TimeToRemediation, true
}

// HasTimeToRemediation returns a boolean if a field has been set.
func (o *Insight) HasTimeToRemediation() bool {
	if o != nil && !IsNil(o.TimeToRemediation) {
		return true
	}

	return false
}

// SetTimeToRemediation gets a reference to the given int32 and assigns it to the TimeToRemediation field.
func (o *Insight) SetTimeToRemediation(v int32) {
	o.TimeToRemediation = &v
}

// GetTags returns the Tags field value
func (o *Insight) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *Insight) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *Insight) SetTags(v []string) {
	o.Tags = v
}

func (o Insight) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Insight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["readableId"] = o.ReadableId
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["source"] = o.Source
	if !IsNil(o.AssignedTo) {
		toSerialize["assignedTo"] = o.AssignedTo
	}
	if !IsNil(o.TeamAssignedTo) {
		toSerialize["teamAssignedTo"] = o.TeamAssignedTo
	}
	toSerialize["created"] = o.Created
	if !IsNil(o.Closed) {
		toSerialize["closed"] = o.Closed
	}
	if !IsNil(o.ClosedBy) {
		toSerialize["closedBy"] = o.ClosedBy
	}
	toSerialize["severity"] = o.Severity
	if !IsNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.Resolution) {
		toSerialize["resolution"] = o.Resolution
	}
	if !IsNil(o.SubResolution) {
		toSerialize["subResolution"] = o.SubResolution
	}
	toSerialize["entity"] = o.Entity
	toSerialize["signals"] = o.Signals
	toSerialize["involvedEntities"] = o.InvolvedEntities
	toSerialize["artifacts"] = o.Artifacts
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.LastUpdatedBy) {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	toSerialize["orgId"] = o.OrgId
	if !IsNil(o.TimeToDetection) {
		toSerialize["timeToDetection"] = o.TimeToDetection
	}
	if !IsNil(o.TimeToResponse) {
		toSerialize["timeToResponse"] = o.TimeToResponse
	}
	if !IsNil(o.TimeToRemediation) {
		toSerialize["timeToRemediation"] = o.TimeToRemediation
	}
	toSerialize["tags"] = o.Tags
	return toSerialize, nil
}

func (o *Insight) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"readableId",
		"name",
		"description",
		"timestamp",
		"source",
		"created",
		"severity",
		"status",
		"entity",
		"signals",
		"involvedEntities",
		"artifacts",
		"orgId",
		"tags",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInsight := _Insight{}

	err = json.Unmarshal(bytes, &varInsight)

	if err != nil {
		return err
	}

	*o = Insight(varInsight)

	return err
}

type NullableInsight struct {
	value *Insight
	isSet bool
}

func (v NullableInsight) Get() *Insight {
	return v.value
}

func (v *NullableInsight) Set(val *Insight) {
	v.value = val
	v.isSet = true
}

func (v NullableInsight) IsSet() bool {
	return v.isSet
}

func (v *NullableInsight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsight(val *Insight) *NullableInsight {
	return &NullableInsight{value: val, isSet: true}
}

func (v NullableInsight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

