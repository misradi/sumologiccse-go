# GetInsightsResponseDataObjectsInner

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
**RecordSummaryFields** | [**[]GetInsightsResponseDataObjectsInnerRecordSummaryFieldsInner**](GetInsightsResponseDataObjectsInnerRecordSummaryFieldsInner.md) | The aggregated fields and values from the records of the Insight based on the recordSummaryFields query parameter | 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastUpdatedBy** | Pointer to **string** |  | [optional] 
**OrgId** | **string** |  | 
**TimeToDetection** | Pointer to **int32** |  | [optional] 
**TimeToResponse** | Pointer to **int32** |  | [optional] 
**TimeToRemediation** | Pointer to **int32** |  | [optional] 
**Tags** | **[]string** |  | 

## Methods

### NewGetInsightsResponseDataObjectsInner

`func NewGetInsightsResponseDataObjectsInner(id string, readableId string, name string, description string, timestamp time.Time, source string, created time.Time, severity string, status GetInsightsResponseDataObjectsInnerStatus, entity GetInsightsResponseDataObjectsInnerEntity, signals []GetInsightsResponseDataObjectsInnerSignalsInner, involvedEntities []GetInsightsResponseDataObjectsInnerSignalsInnerEntity, artifacts []GetInsightsResponseDataObjectsInnerArtifactsInner, recordSummaryFields []GetInsightsResponseDataObjectsInnerRecordSummaryFieldsInner, orgId string, tags []string, ) *GetInsightsResponseDataObjectsInner`

NewGetInsightsResponseDataObjectsInner instantiates a new GetInsightsResponseDataObjectsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInsightsResponseDataObjectsInnerWithDefaults

`func NewGetInsightsResponseDataObjectsInnerWithDefaults() *GetInsightsResponseDataObjectsInner`

NewGetInsightsResponseDataObjectsInnerWithDefaults instantiates a new GetInsightsResponseDataObjectsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInsightsResponseDataObjectsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInsightsResponseDataObjectsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInsightsResponseDataObjectsInner) SetId(v string)`

SetId sets Id field to given value.


### GetReadableId

`func (o *GetInsightsResponseDataObjectsInner) GetReadableId() string`

GetReadableId returns the ReadableId field if non-nil, zero value otherwise.

### GetReadableIdOk

`func (o *GetInsightsResponseDataObjectsInner) GetReadableIdOk() (*string, bool)`

GetReadableIdOk returns a tuple with the ReadableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadableId

`func (o *GetInsightsResponseDataObjectsInner) SetReadableId(v string)`

SetReadableId sets ReadableId field to given value.


### GetName

`func (o *GetInsightsResponseDataObjectsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInsightsResponseDataObjectsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInsightsResponseDataObjectsInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GetInsightsResponseDataObjectsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetInsightsResponseDataObjectsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetInsightsResponseDataObjectsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTimestamp

`func (o *GetInsightsResponseDataObjectsInner) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetInsightsResponseDataObjectsInner) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetInsightsResponseDataObjectsInner) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetSource

`func (o *GetInsightsResponseDataObjectsInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetInsightsResponseDataObjectsInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetInsightsResponseDataObjectsInner) SetSource(v string)`

SetSource sets Source field to given value.


### GetAssignedTo

`func (o *GetInsightsResponseDataObjectsInner) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *GetInsightsResponseDataObjectsInner) GetAssignedToOk() (*string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *GetInsightsResponseDataObjectsInner) SetAssignedTo(v string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *GetInsightsResponseDataObjectsInner) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetTeamAssignedTo

`func (o *GetInsightsResponseDataObjectsInner) GetTeamAssignedTo() string`

GetTeamAssignedTo returns the TeamAssignedTo field if non-nil, zero value otherwise.

### GetTeamAssignedToOk

`func (o *GetInsightsResponseDataObjectsInner) GetTeamAssignedToOk() (*string, bool)`

GetTeamAssignedToOk returns a tuple with the TeamAssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamAssignedTo

`func (o *GetInsightsResponseDataObjectsInner) SetTeamAssignedTo(v string)`

SetTeamAssignedTo sets TeamAssignedTo field to given value.

### HasTeamAssignedTo

`func (o *GetInsightsResponseDataObjectsInner) HasTeamAssignedTo() bool`

HasTeamAssignedTo returns a boolean if a field has been set.

### GetCreated

`func (o *GetInsightsResponseDataObjectsInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetInsightsResponseDataObjectsInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetInsightsResponseDataObjectsInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetClosed

`func (o *GetInsightsResponseDataObjectsInner) GetClosed() time.Time`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *GetInsightsResponseDataObjectsInner) GetClosedOk() (*time.Time, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *GetInsightsResponseDataObjectsInner) SetClosed(v time.Time)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *GetInsightsResponseDataObjectsInner) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetClosedBy

`func (o *GetInsightsResponseDataObjectsInner) GetClosedBy() string`

GetClosedBy returns the ClosedBy field if non-nil, zero value otherwise.

### GetClosedByOk

`func (o *GetInsightsResponseDataObjectsInner) GetClosedByOk() (*string, bool)`

GetClosedByOk returns a tuple with the ClosedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedBy

`func (o *GetInsightsResponseDataObjectsInner) SetClosedBy(v string)`

SetClosedBy sets ClosedBy field to given value.

### HasClosedBy

`func (o *GetInsightsResponseDataObjectsInner) HasClosedBy() bool`

HasClosedBy returns a boolean if a field has been set.

### GetSeverity

`func (o *GetInsightsResponseDataObjectsInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetInsightsResponseDataObjectsInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetInsightsResponseDataObjectsInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetConfidence

`func (o *GetInsightsResponseDataObjectsInner) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *GetInsightsResponseDataObjectsInner) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *GetInsightsResponseDataObjectsInner) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *GetInsightsResponseDataObjectsInner) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetAssignee

`func (o *GetInsightsResponseDataObjectsInner) GetAssignee() GetInsightsResponseDataObjectsInnerAssignee`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *GetInsightsResponseDataObjectsInner) GetAssigneeOk() (*GetInsightsResponseDataObjectsInnerAssignee, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *GetInsightsResponseDataObjectsInner) SetAssignee(v GetInsightsResponseDataObjectsInnerAssignee)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *GetInsightsResponseDataObjectsInner) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetStatus

`func (o *GetInsightsResponseDataObjectsInner) GetStatus() GetInsightsResponseDataObjectsInnerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInsightsResponseDataObjectsInner) GetStatusOk() (*GetInsightsResponseDataObjectsInnerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInsightsResponseDataObjectsInner) SetStatus(v GetInsightsResponseDataObjectsInnerStatus)`

SetStatus sets Status field to given value.


### GetResolution

`func (o *GetInsightsResponseDataObjectsInner) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *GetInsightsResponseDataObjectsInner) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *GetInsightsResponseDataObjectsInner) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *GetInsightsResponseDataObjectsInner) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetSubResolution

`func (o *GetInsightsResponseDataObjectsInner) GetSubResolution() string`

GetSubResolution returns the SubResolution field if non-nil, zero value otherwise.

### GetSubResolutionOk

`func (o *GetInsightsResponseDataObjectsInner) GetSubResolutionOk() (*string, bool)`

GetSubResolutionOk returns a tuple with the SubResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubResolution

`func (o *GetInsightsResponseDataObjectsInner) SetSubResolution(v string)`

SetSubResolution sets SubResolution field to given value.

### HasSubResolution

`func (o *GetInsightsResponseDataObjectsInner) HasSubResolution() bool`

HasSubResolution returns a boolean if a field has been set.

### GetEntity

`func (o *GetInsightsResponseDataObjectsInner) GetEntity() GetInsightsResponseDataObjectsInnerEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GetInsightsResponseDataObjectsInner) GetEntityOk() (*GetInsightsResponseDataObjectsInnerEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GetInsightsResponseDataObjectsInner) SetEntity(v GetInsightsResponseDataObjectsInnerEntity)`

SetEntity sets Entity field to given value.


### GetSignals

`func (o *GetInsightsResponseDataObjectsInner) GetSignals() []GetInsightsResponseDataObjectsInnerSignalsInner`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *GetInsightsResponseDataObjectsInner) GetSignalsOk() (*[]GetInsightsResponseDataObjectsInnerSignalsInner, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *GetInsightsResponseDataObjectsInner) SetSignals(v []GetInsightsResponseDataObjectsInnerSignalsInner)`

SetSignals sets Signals field to given value.


### GetInvolvedEntities

`func (o *GetInsightsResponseDataObjectsInner) GetInvolvedEntities() []GetInsightsResponseDataObjectsInnerSignalsInnerEntity`

GetInvolvedEntities returns the InvolvedEntities field if non-nil, zero value otherwise.

### GetInvolvedEntitiesOk

`func (o *GetInsightsResponseDataObjectsInner) GetInvolvedEntitiesOk() (*[]GetInsightsResponseDataObjectsInnerSignalsInnerEntity, bool)`

GetInvolvedEntitiesOk returns a tuple with the InvolvedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvolvedEntities

`func (o *GetInsightsResponseDataObjectsInner) SetInvolvedEntities(v []GetInsightsResponseDataObjectsInnerSignalsInnerEntity)`

SetInvolvedEntities sets InvolvedEntities field to given value.


### GetArtifacts

`func (o *GetInsightsResponseDataObjectsInner) GetArtifacts() []GetInsightsResponseDataObjectsInnerArtifactsInner`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *GetInsightsResponseDataObjectsInner) GetArtifactsOk() (*[]GetInsightsResponseDataObjectsInnerArtifactsInner, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *GetInsightsResponseDataObjectsInner) SetArtifacts(v []GetInsightsResponseDataObjectsInnerArtifactsInner)`

SetArtifacts sets Artifacts field to given value.


### GetRecordSummaryFields

`func (o *GetInsightsResponseDataObjectsInner) GetRecordSummaryFields() []GetInsightsResponseDataObjectsInnerRecordSummaryFieldsInner`

GetRecordSummaryFields returns the RecordSummaryFields field if non-nil, zero value otherwise.

### GetRecordSummaryFieldsOk

`func (o *GetInsightsResponseDataObjectsInner) GetRecordSummaryFieldsOk() (*[]GetInsightsResponseDataObjectsInnerRecordSummaryFieldsInner, bool)`

GetRecordSummaryFieldsOk returns a tuple with the RecordSummaryFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordSummaryFields

`func (o *GetInsightsResponseDataObjectsInner) SetRecordSummaryFields(v []GetInsightsResponseDataObjectsInnerRecordSummaryFieldsInner)`

SetRecordSummaryFields sets RecordSummaryFields field to given value.


### GetLastUpdated

`func (o *GetInsightsResponseDataObjectsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetInsightsResponseDataObjectsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetInsightsResponseDataObjectsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetInsightsResponseDataObjectsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *GetInsightsResponseDataObjectsInner) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *GetInsightsResponseDataObjectsInner) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *GetInsightsResponseDataObjectsInner) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *GetInsightsResponseDataObjectsInner) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetOrgId

`func (o *GetInsightsResponseDataObjectsInner) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *GetInsightsResponseDataObjectsInner) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *GetInsightsResponseDataObjectsInner) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetTimeToDetection

`func (o *GetInsightsResponseDataObjectsInner) GetTimeToDetection() int32`

GetTimeToDetection returns the TimeToDetection field if non-nil, zero value otherwise.

### GetTimeToDetectionOk

`func (o *GetInsightsResponseDataObjectsInner) GetTimeToDetectionOk() (*int32, bool)`

GetTimeToDetectionOk returns a tuple with the TimeToDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToDetection

`func (o *GetInsightsResponseDataObjectsInner) SetTimeToDetection(v int32)`

SetTimeToDetection sets TimeToDetection field to given value.

### HasTimeToDetection

`func (o *GetInsightsResponseDataObjectsInner) HasTimeToDetection() bool`

HasTimeToDetection returns a boolean if a field has been set.

### GetTimeToResponse

`func (o *GetInsightsResponseDataObjectsInner) GetTimeToResponse() int32`

GetTimeToResponse returns the TimeToResponse field if non-nil, zero value otherwise.

### GetTimeToResponseOk

`func (o *GetInsightsResponseDataObjectsInner) GetTimeToResponseOk() (*int32, bool)`

GetTimeToResponseOk returns a tuple with the TimeToResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToResponse

`func (o *GetInsightsResponseDataObjectsInner) SetTimeToResponse(v int32)`

SetTimeToResponse sets TimeToResponse field to given value.

### HasTimeToResponse

`func (o *GetInsightsResponseDataObjectsInner) HasTimeToResponse() bool`

HasTimeToResponse returns a boolean if a field has been set.

### GetTimeToRemediation

`func (o *GetInsightsResponseDataObjectsInner) GetTimeToRemediation() int32`

GetTimeToRemediation returns the TimeToRemediation field if non-nil, zero value otherwise.

### GetTimeToRemediationOk

`func (o *GetInsightsResponseDataObjectsInner) GetTimeToRemediationOk() (*int32, bool)`

GetTimeToRemediationOk returns a tuple with the TimeToRemediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToRemediation

`func (o *GetInsightsResponseDataObjectsInner) SetTimeToRemediation(v int32)`

SetTimeToRemediation sets TimeToRemediation field to given value.

### HasTimeToRemediation

`func (o *GetInsightsResponseDataObjectsInner) HasTimeToRemediation() bool`

HasTimeToRemediation returns a boolean if a field has been set.

### GetTags

`func (o *GetInsightsResponseDataObjectsInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetInsightsResponseDataObjectsInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetInsightsResponseDataObjectsInner) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


