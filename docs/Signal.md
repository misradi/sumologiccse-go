# Signal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Summary** | **string** |  | 
**Stage** | **string** |  | 
**Created** | **time.Time** | When this signal was generated from a rule. | 
**Timestamp** | **time.Time** | Timestamp of first log record for this signal. | 
**Severity** | **int32** |  | 
**RecordCount** | **int32** | The total number of Records (including the sum of primaryRecordsJson and extraRecordsJson | 
**RecordTypes** | **[]string** |  | 
**ContentType** | Pointer to **string** |  | [optional] 
**AllRecords** | **[]map[string]interface{}** | A JSON-stringified array of all Records associated with this Signal. | 
**Suppressed** | Pointer to **bool** | Whether the Signal is suppressed | [optional] 
**RuleId** | Pointer to **string** |  | [optional] 
**Tags** | **[]string** |  | 
**Entity** | [**SignalEntity**](SignalEntity.md) |  | 
**InvolvedEntities** | [**[]SignalEntity**](SignalEntity.md) |  | 
**Artifacts** | Pointer to [**[]GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner**](GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner.md) |  | [optional] 

## Methods

### NewSignal

`func NewSignal(id string, name string, summary string, stage string, created time.Time, timestamp time.Time, severity int32, recordCount int32, recordTypes []string, allRecords []map[string]interface{}, tags []string, entity SignalEntity, involvedEntities []SignalEntity, ) *Signal`

NewSignal instantiates a new Signal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalWithDefaults

`func NewSignalWithDefaults() *Signal`

NewSignalWithDefaults instantiates a new Signal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Signal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Signal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Signal) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Signal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Signal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Signal) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Signal) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Signal) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Signal) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Signal) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSummary

`func (o *Signal) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Signal) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Signal) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetStage

`func (o *Signal) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *Signal) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *Signal) SetStage(v string)`

SetStage sets Stage field to given value.


### GetCreated

`func (o *Signal) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Signal) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Signal) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetTimestamp

`func (o *Signal) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Signal) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Signal) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetSeverity

`func (o *Signal) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Signal) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Signal) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.


### GetRecordCount

`func (o *Signal) GetRecordCount() int32`

GetRecordCount returns the RecordCount field if non-nil, zero value otherwise.

### GetRecordCountOk

`func (o *Signal) GetRecordCountOk() (*int32, bool)`

GetRecordCountOk returns a tuple with the RecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordCount

`func (o *Signal) SetRecordCount(v int32)`

SetRecordCount sets RecordCount field to given value.


### GetRecordTypes

`func (o *Signal) GetRecordTypes() []string`

GetRecordTypes returns the RecordTypes field if non-nil, zero value otherwise.

### GetRecordTypesOk

`func (o *Signal) GetRecordTypesOk() (*[]string, bool)`

GetRecordTypesOk returns a tuple with the RecordTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTypes

`func (o *Signal) SetRecordTypes(v []string)`

SetRecordTypes sets RecordTypes field to given value.


### GetContentType

`func (o *Signal) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *Signal) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *Signal) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *Signal) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetAllRecords

`func (o *Signal) GetAllRecords() []map[string]interface{}`

GetAllRecords returns the AllRecords field if non-nil, zero value otherwise.

### GetAllRecordsOk

`func (o *Signal) GetAllRecordsOk() (*[]map[string]interface{}, bool)`

GetAllRecordsOk returns a tuple with the AllRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRecords

`func (o *Signal) SetAllRecords(v []map[string]interface{})`

SetAllRecords sets AllRecords field to given value.


### GetSuppressed

`func (o *Signal) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *Signal) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *Signal) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *Signal) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetRuleId

`func (o *Signal) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *Signal) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *Signal) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *Signal) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetTags

`func (o *Signal) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Signal) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Signal) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetEntity

`func (o *Signal) GetEntity() SignalEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *Signal) GetEntityOk() (*SignalEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *Signal) SetEntity(v SignalEntity)`

SetEntity sets Entity field to given value.


### GetInvolvedEntities

`func (o *Signal) GetInvolvedEntities() []SignalEntity`

GetInvolvedEntities returns the InvolvedEntities field if non-nil, zero value otherwise.

### GetInvolvedEntitiesOk

`func (o *Signal) GetInvolvedEntitiesOk() (*[]SignalEntity, bool)`

GetInvolvedEntitiesOk returns a tuple with the InvolvedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvolvedEntities

`func (o *Signal) SetInvolvedEntities(v []SignalEntity)`

SetInvolvedEntities sets InvolvedEntities field to given value.


### GetArtifacts

`func (o *Signal) GetArtifacts() []GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *Signal) GetArtifactsOk() (*[]GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *Signal) SetArtifacts(v []GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *Signal) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


