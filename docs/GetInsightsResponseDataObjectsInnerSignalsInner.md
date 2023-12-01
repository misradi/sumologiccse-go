# GetInsightsResponseDataObjectsInnerSignalsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Stage** | **string** |  | 
**Created** | **time.Time** | When this signal was generated from a rule. | 
**Timestamp** | **time.Time** | Timestamp of first log record for this signal. | 
**Severity** | **int32** |  | 
**RecordCount** | **int32** | The total number of Records (including the sum of primaryRecordsJson and extraRecordsJson | 
**RecordTypes** | **[]string** |  | 
**ContentType** | Pointer to **string** |  | [optional] 
**AllRecords** | **[]map[string]interface{}** | A JSON-stringified array of all Records associated with this Signal. | 
**RuleId** | Pointer to **string** |  | [optional] 
**Tags** | **[]string** |  | 
**Entity** | [**GetInsightsResponseDataObjectsInnerSignalsInnerEntity**](GetInsightsResponseDataObjectsInnerSignalsInnerEntity.md) |  | 
**Artifacts** | Pointer to [**[]GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner**](GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner.md) |  | [optional] 

## Methods

### NewGetInsightsResponseDataObjectsInnerSignalsInner

`func NewGetInsightsResponseDataObjectsInnerSignalsInner(id string, name string, stage string, created time.Time, timestamp time.Time, severity int32, recordCount int32, recordTypes []string, allRecords []map[string]interface{}, tags []string, entity GetInsightsResponseDataObjectsInnerSignalsInnerEntity, ) *GetInsightsResponseDataObjectsInnerSignalsInner`

NewGetInsightsResponseDataObjectsInnerSignalsInner instantiates a new GetInsightsResponseDataObjectsInnerSignalsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInsightsResponseDataObjectsInnerSignalsInnerWithDefaults

`func NewGetInsightsResponseDataObjectsInnerSignalsInnerWithDefaults() *GetInsightsResponseDataObjectsInnerSignalsInner`

NewGetInsightsResponseDataObjectsInnerSignalsInnerWithDefaults instantiates a new GetInsightsResponseDataObjectsInnerSignalsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStage

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) SetStage(v string)`

SetStage sets Stage field to given value.


### GetCreated

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetTimestamp

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetSeverity

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.


### GetRecordCount

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetRecordCount() int32`

GetRecordCount returns the RecordCount field if non-nil, zero value otherwise.

### GetRecordCountOk

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetRecordCountOk() (*int32, bool)`

GetRecordCountOk returns a tuple with the RecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordCount

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) SetRecordCount(v int32)`

SetRecordCount sets RecordCount field to given value.


### GetRecordTypes

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetRecordTypes() []string`

GetRecordTypes returns the RecordTypes field if non-nil, zero value otherwise.

### GetRecordTypesOk

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetRecordTypesOk() (*[]string, bool)`

GetRecordTypesOk returns a tuple with the RecordTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTypes

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) SetRecordTypes(v []string)`

SetRecordTypes sets RecordTypes field to given value.


### GetContentType

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetAllRecords

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetAllRecords() []map[string]interface{}`

GetAllRecords returns the AllRecords field if non-nil, zero value otherwise.

### GetAllRecordsOk

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetAllRecordsOk() (*[]map[string]interface{}, bool)`

GetAllRecordsOk returns a tuple with the AllRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRecords

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) SetAllRecords(v []map[string]interface{})`

SetAllRecords sets AllRecords field to given value.


### GetRuleId

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetTags

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetEntity

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetEntity() GetInsightsResponseDataObjectsInnerSignalsInnerEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetEntityOk() (*GetInsightsResponseDataObjectsInnerSignalsInnerEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) SetEntity(v GetInsightsResponseDataObjectsInnerSignalsInnerEntity)`

SetEntity sets Entity field to given value.


### GetArtifacts

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetArtifacts() []GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) GetArtifactsOk() (*[]GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) SetArtifacts(v []GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *GetInsightsResponseDataObjectsInnerSignalsInner) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


