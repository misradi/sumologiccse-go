# Insight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ReadableId** | **string** | A human-readable ID in the format \&quot;INSIGHT-542\&quot;. This is technically nullable, but in reality it will always be populated in every query other than the cross-type search query. | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**Source** | **string** |  | 
**AssignedTo** | Pointer to **string** | The user that this Insight is assigned to | [optional] 
**TeamAssignedTo** | Pointer to **string** | The team that this Insight is assigned to | [optional] 
**Created** | **time.Time** |  | 
**Closed** | Pointer to **time.Time** |  | [optional] 
**ClosedBy** | Pointer to **string** |  | [optional] 
**Severity** | **string** |  | 
**Confidence** | Pointer to **float64** | A 0-100 value of the ML-based confidence score for the Insight | [optional] 
**Assignee** | Pointer to [**GetInsightsResponseDataObjectsInnerAssignee**](GetInsightsResponseDataObjectsInnerAssignee.md) |  | [optional] 
**Status** | [**GetInsightsResponseDataObjectsInnerStatus**](GetInsightsResponseDataObjectsInnerStatus.md) |  | 
**Resolution** | Pointer to **string** |  | [optional] 
**SubResolution** | Pointer to **string** |  | [optional] 
**Entity** | [**GetInsightsResponseDataObjectsInnerEntity**](GetInsightsResponseDataObjectsInnerEntity.md) |  | 
**Signals** | [**[]GetInsightsResponseDataObjectsInnerSignalsInner**](GetInsightsResponseDataObjectsInnerSignalsInner.md) |  | 
**InvolvedEntities** | [**[]GetInsightsResponseDataObjectsInnerSignalsInnerEntity**](GetInsightsResponseDataObjectsInnerSignalsInnerEntity.md) |  | 
**Artifacts** | [**[]GetInsightsResponseDataObjectsInnerArtifactsInner**](GetInsightsResponseDataObjectsInnerArtifactsInner.md) |  | 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastUpdatedBy** | Pointer to **string** |  | [optional] 
**OrgId** | **string** |  | 
**TimeToDetection** | Pointer to **int32** |  | [optional] 
**TimeToResponse** | Pointer to **int32** |  | [optional] 
**TimeToRemediation** | Pointer to **int32** |  | [optional] 
**Tags** | **[]string** |  | 

## Methods

### NewInsight

`func NewInsight(id string, readableId string, name string, description string, timestamp time.Time, source string, created time.Time, severity string, status GetInsightsResponseDataObjectsInnerStatus, entity GetInsightsResponseDataObjectsInnerEntity, signals []GetInsightsResponseDataObjectsInnerSignalsInner, involvedEntities []GetInsightsResponseDataObjectsInnerSignalsInnerEntity, artifacts []GetInsightsResponseDataObjectsInnerArtifactsInner, orgId string, tags []string, ) *Insight`

NewInsight instantiates a new Insight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightWithDefaults

`func NewInsightWithDefaults() *Insight`

NewInsightWithDefaults instantiates a new Insight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Insight) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Insight) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Insight) SetId(v string)`

SetId sets Id field to given value.


### GetReadableId

`func (o *Insight) GetReadableId() string`

GetReadableId returns the ReadableId field if non-nil, zero value otherwise.

### GetReadableIdOk

`func (o *Insight) GetReadableIdOk() (*string, bool)`

GetReadableIdOk returns a tuple with the ReadableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadableId

`func (o *Insight) SetReadableId(v string)`

SetReadableId sets ReadableId field to given value.


### GetName

`func (o *Insight) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Insight) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Insight) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Insight) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Insight) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Insight) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTimestamp

`func (o *Insight) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Insight) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Insight) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetSource

`func (o *Insight) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Insight) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Insight) SetSource(v string)`

SetSource sets Source field to given value.


### GetAssignedTo

`func (o *Insight) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *Insight) GetAssignedToOk() (*string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *Insight) SetAssignedTo(v string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *Insight) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetTeamAssignedTo

`func (o *Insight) GetTeamAssignedTo() string`

GetTeamAssignedTo returns the TeamAssignedTo field if non-nil, zero value otherwise.

### GetTeamAssignedToOk

`func (o *Insight) GetTeamAssignedToOk() (*string, bool)`

GetTeamAssignedToOk returns a tuple with the TeamAssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamAssignedTo

`func (o *Insight) SetTeamAssignedTo(v string)`

SetTeamAssignedTo sets TeamAssignedTo field to given value.

### HasTeamAssignedTo

`func (o *Insight) HasTeamAssignedTo() bool`

HasTeamAssignedTo returns a boolean if a field has been set.

### GetCreated

`func (o *Insight) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Insight) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Insight) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetClosed

`func (o *Insight) GetClosed() time.Time`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *Insight) GetClosedOk() (*time.Time, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *Insight) SetClosed(v time.Time)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *Insight) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetClosedBy

`func (o *Insight) GetClosedBy() string`

GetClosedBy returns the ClosedBy field if non-nil, zero value otherwise.

### GetClosedByOk

`func (o *Insight) GetClosedByOk() (*string, bool)`

GetClosedByOk returns a tuple with the ClosedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedBy

`func (o *Insight) SetClosedBy(v string)`

SetClosedBy sets ClosedBy field to given value.

### HasClosedBy

`func (o *Insight) HasClosedBy() bool`

HasClosedBy returns a boolean if a field has been set.

### GetSeverity

`func (o *Insight) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Insight) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Insight) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetConfidence

`func (o *Insight) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *Insight) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *Insight) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *Insight) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetAssignee

`func (o *Insight) GetAssignee() GetInsightsResponseDataObjectsInnerAssignee`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *Insight) GetAssigneeOk() (*GetInsightsResponseDataObjectsInnerAssignee, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *Insight) SetAssignee(v GetInsightsResponseDataObjectsInnerAssignee)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *Insight) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetStatus

`func (o *Insight) GetStatus() GetInsightsResponseDataObjectsInnerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Insight) GetStatusOk() (*GetInsightsResponseDataObjectsInnerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Insight) SetStatus(v GetInsightsResponseDataObjectsInnerStatus)`

SetStatus sets Status field to given value.


### GetResolution

`func (o *Insight) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *Insight) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *Insight) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *Insight) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetSubResolution

`func (o *Insight) GetSubResolution() string`

GetSubResolution returns the SubResolution field if non-nil, zero value otherwise.

### GetSubResolutionOk

`func (o *Insight) GetSubResolutionOk() (*string, bool)`

GetSubResolutionOk returns a tuple with the SubResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubResolution

`func (o *Insight) SetSubResolution(v string)`

SetSubResolution sets SubResolution field to given value.

### HasSubResolution

`func (o *Insight) HasSubResolution() bool`

HasSubResolution returns a boolean if a field has been set.

### GetEntity

`func (o *Insight) GetEntity() GetInsightsResponseDataObjectsInnerEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *Insight) GetEntityOk() (*GetInsightsResponseDataObjectsInnerEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *Insight) SetEntity(v GetInsightsResponseDataObjectsInnerEntity)`

SetEntity sets Entity field to given value.


### GetSignals

`func (o *Insight) GetSignals() []GetInsightsResponseDataObjectsInnerSignalsInner`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *Insight) GetSignalsOk() (*[]GetInsightsResponseDataObjectsInnerSignalsInner, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *Insight) SetSignals(v []GetInsightsResponseDataObjectsInnerSignalsInner)`

SetSignals sets Signals field to given value.


### GetInvolvedEntities

`func (o *Insight) GetInvolvedEntities() []GetInsightsResponseDataObjectsInnerSignalsInnerEntity`

GetInvolvedEntities returns the InvolvedEntities field if non-nil, zero value otherwise.

### GetInvolvedEntitiesOk

`func (o *Insight) GetInvolvedEntitiesOk() (*[]GetInsightsResponseDataObjectsInnerSignalsInnerEntity, bool)`

GetInvolvedEntitiesOk returns a tuple with the InvolvedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvolvedEntities

`func (o *Insight) SetInvolvedEntities(v []GetInsightsResponseDataObjectsInnerSignalsInnerEntity)`

SetInvolvedEntities sets InvolvedEntities field to given value.


### GetArtifacts

`func (o *Insight) GetArtifacts() []GetInsightsResponseDataObjectsInnerArtifactsInner`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *Insight) GetArtifactsOk() (*[]GetInsightsResponseDataObjectsInnerArtifactsInner, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *Insight) SetArtifacts(v []GetInsightsResponseDataObjectsInnerArtifactsInner)`

SetArtifacts sets Artifacts field to given value.


### GetLastUpdated

`func (o *Insight) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Insight) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Insight) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Insight) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *Insight) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *Insight) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *Insight) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *Insight) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetOrgId

`func (o *Insight) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Insight) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Insight) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetTimeToDetection

`func (o *Insight) GetTimeToDetection() int32`

GetTimeToDetection returns the TimeToDetection field if non-nil, zero value otherwise.

### GetTimeToDetectionOk

`func (o *Insight) GetTimeToDetectionOk() (*int32, bool)`

GetTimeToDetectionOk returns a tuple with the TimeToDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToDetection

`func (o *Insight) SetTimeToDetection(v int32)`

SetTimeToDetection sets TimeToDetection field to given value.

### HasTimeToDetection

`func (o *Insight) HasTimeToDetection() bool`

HasTimeToDetection returns a boolean if a field has been set.

### GetTimeToResponse

`func (o *Insight) GetTimeToResponse() int32`

GetTimeToResponse returns the TimeToResponse field if non-nil, zero value otherwise.

### GetTimeToResponseOk

`func (o *Insight) GetTimeToResponseOk() (*int32, bool)`

GetTimeToResponseOk returns a tuple with the TimeToResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToResponse

`func (o *Insight) SetTimeToResponse(v int32)`

SetTimeToResponse sets TimeToResponse field to given value.

### HasTimeToResponse

`func (o *Insight) HasTimeToResponse() bool`

HasTimeToResponse returns a boolean if a field has been set.

### GetTimeToRemediation

`func (o *Insight) GetTimeToRemediation() int32`

GetTimeToRemediation returns the TimeToRemediation field if non-nil, zero value otherwise.

### GetTimeToRemediationOk

`func (o *Insight) GetTimeToRemediationOk() (*int32, bool)`

GetTimeToRemediationOk returns a tuple with the TimeToRemediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToRemediation

`func (o *Insight) SetTimeToRemediation(v int32)`

SetTimeToRemediation sets TimeToRemediation field to given value.

### HasTimeToRemediation

`func (o *Insight) HasTimeToRemediation() bool`

HasTimeToRemediation returns a boolean if a field has been set.

### GetTags

`func (o *Insight) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Insight) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Insight) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


