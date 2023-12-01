# CreateFirstSeenRuleRequestBodyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetField** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** |  | 
**EntitySelectors** | Pointer to [**[]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner**](CreateMatchRuleRequestBodyFieldsEntitySelectorsInner.md) |  | [optional] 
**IsPrototype** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**ParentJaskId** | Pointer to **string** |  | [optional] 
**SummaryExpression** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TuningExpressionIds** | Pointer to **[]string** |  | [optional] 
**DescriptionExpression** | **string** | The description to be used on the generated Signal | 
**NameExpression** | **string** |  | 
**FilterExpression** | **string** |  | 
**ValueFields** | Pointer to **[]string** | The value(s) to build an expression from - can be used instead of valueExpression | [optional] 
**ValueExpression** | Pointer to **string** |  | [optional] 
**GroupByFields** | **[]string** |  | 
**Score** | **int32** |  | 
**Version** | **int32** |  | 
**BaselineWindowSize** | **string** | milliseconds | 
**RetentionWindowSize** | **string** | milliseconds | 
**BaselineType** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateFirstSeenRuleRequestBodyFields

`func NewCreateFirstSeenRuleRequestBodyFields(enabled bool, name string, descriptionExpression string, nameExpression string, filterExpression string, groupByFields []string, score int32, version int32, baselineWindowSize string, retentionWindowSize string, ) *CreateFirstSeenRuleRequestBodyFields`

NewCreateFirstSeenRuleRequestBodyFields instantiates a new CreateFirstSeenRuleRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirstSeenRuleRequestBodyFieldsWithDefaults

`func NewCreateFirstSeenRuleRequestBodyFieldsWithDefaults() *CreateFirstSeenRuleRequestBodyFields`

NewCreateFirstSeenRuleRequestBodyFieldsWithDefaults instantiates a new CreateFirstSeenRuleRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetField

`func (o *CreateFirstSeenRuleRequestBodyFields) GetAssetField() string`

GetAssetField returns the AssetField field if non-nil, zero value otherwise.

### GetAssetFieldOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetAssetFieldOk() (*string, bool)`

GetAssetFieldOk returns a tuple with the AssetField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetField

`func (o *CreateFirstSeenRuleRequestBodyFields) SetAssetField(v string)`

SetAssetField sets AssetField field to given value.

### HasAssetField

`func (o *CreateFirstSeenRuleRequestBodyFields) HasAssetField() bool`

HasAssetField returns a boolean if a field has been set.

### GetCategory

`func (o *CreateFirstSeenRuleRequestBodyFields) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateFirstSeenRuleRequestBodyFields) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CreateFirstSeenRuleRequestBodyFields) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateFirstSeenRuleRequestBodyFields) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateFirstSeenRuleRequestBodyFields) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEntitySelectors

`func (o *CreateFirstSeenRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner`

GetEntitySelectors returns the EntitySelectors field if non-nil, zero value otherwise.

### GetEntitySelectorsOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetEntitySelectorsOk() (*[]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool)`

GetEntitySelectorsOk returns a tuple with the EntitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectors

`func (o *CreateFirstSeenRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner)`

SetEntitySelectors sets EntitySelectors field to given value.

### HasEntitySelectors

`func (o *CreateFirstSeenRuleRequestBodyFields) HasEntitySelectors() bool`

HasEntitySelectors returns a boolean if a field has been set.

### GetIsPrototype

`func (o *CreateFirstSeenRuleRequestBodyFields) GetIsPrototype() bool`

GetIsPrototype returns the IsPrototype field if non-nil, zero value otherwise.

### GetIsPrototypeOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool)`

GetIsPrototypeOk returns a tuple with the IsPrototype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototype

`func (o *CreateFirstSeenRuleRequestBodyFields) SetIsPrototype(v bool)`

SetIsPrototype sets IsPrototype field to given value.

### HasIsPrototype

`func (o *CreateFirstSeenRuleRequestBodyFields) HasIsPrototype() bool`

HasIsPrototype returns a boolean if a field has been set.

### GetName

`func (o *CreateFirstSeenRuleRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFirstSeenRuleRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.


### GetParentJaskId

`func (o *CreateFirstSeenRuleRequestBodyFields) GetParentJaskId() string`

GetParentJaskId returns the ParentJaskId field if non-nil, zero value otherwise.

### GetParentJaskIdOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetParentJaskIdOk() (*string, bool)`

GetParentJaskIdOk returns a tuple with the ParentJaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJaskId

`func (o *CreateFirstSeenRuleRequestBodyFields) SetParentJaskId(v string)`

SetParentJaskId sets ParentJaskId field to given value.

### HasParentJaskId

`func (o *CreateFirstSeenRuleRequestBodyFields) HasParentJaskId() bool`

HasParentJaskId returns a boolean if a field has been set.

### GetSummaryExpression

`func (o *CreateFirstSeenRuleRequestBodyFields) GetSummaryExpression() string`

GetSummaryExpression returns the SummaryExpression field if non-nil, zero value otherwise.

### GetSummaryExpressionOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool)`

GetSummaryExpressionOk returns a tuple with the SummaryExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpression

`func (o *CreateFirstSeenRuleRequestBodyFields) SetSummaryExpression(v string)`

SetSummaryExpression sets SummaryExpression field to given value.

### HasSummaryExpression

`func (o *CreateFirstSeenRuleRequestBodyFields) HasSummaryExpression() bool`

HasSummaryExpression returns a boolean if a field has been set.

### GetTags

`func (o *CreateFirstSeenRuleRequestBodyFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateFirstSeenRuleRequestBodyFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateFirstSeenRuleRequestBodyFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTuningExpressionIds

`func (o *CreateFirstSeenRuleRequestBodyFields) GetTuningExpressionIds() []string`

GetTuningExpressionIds returns the TuningExpressionIds field if non-nil, zero value otherwise.

### GetTuningExpressionIdsOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetTuningExpressionIdsOk() (*[]string, bool)`

GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuningExpressionIds

`func (o *CreateFirstSeenRuleRequestBodyFields) SetTuningExpressionIds(v []string)`

SetTuningExpressionIds sets TuningExpressionIds field to given value.

### HasTuningExpressionIds

`func (o *CreateFirstSeenRuleRequestBodyFields) HasTuningExpressionIds() bool`

HasTuningExpressionIds returns a boolean if a field has been set.

### GetDescriptionExpression

`func (o *CreateFirstSeenRuleRequestBodyFields) GetDescriptionExpression() string`

GetDescriptionExpression returns the DescriptionExpression field if non-nil, zero value otherwise.

### GetDescriptionExpressionOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetDescriptionExpressionOk() (*string, bool)`

GetDescriptionExpressionOk returns a tuple with the DescriptionExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionExpression

`func (o *CreateFirstSeenRuleRequestBodyFields) SetDescriptionExpression(v string)`

SetDescriptionExpression sets DescriptionExpression field to given value.


### GetNameExpression

`func (o *CreateFirstSeenRuleRequestBodyFields) GetNameExpression() string`

GetNameExpression returns the NameExpression field if non-nil, zero value otherwise.

### GetNameExpressionOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetNameExpressionOk() (*string, bool)`

GetNameExpressionOk returns a tuple with the NameExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameExpression

`func (o *CreateFirstSeenRuleRequestBodyFields) SetNameExpression(v string)`

SetNameExpression sets NameExpression field to given value.


### GetFilterExpression

`func (o *CreateFirstSeenRuleRequestBodyFields) GetFilterExpression() string`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetFilterExpressionOk() (*string, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *CreateFirstSeenRuleRequestBodyFields) SetFilterExpression(v string)`

SetFilterExpression sets FilterExpression field to given value.


### GetValueFields

`func (o *CreateFirstSeenRuleRequestBodyFields) GetValueFields() []string`

GetValueFields returns the ValueFields field if non-nil, zero value otherwise.

### GetValueFieldsOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetValueFieldsOk() (*[]string, bool)`

GetValueFieldsOk returns a tuple with the ValueFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFields

`func (o *CreateFirstSeenRuleRequestBodyFields) SetValueFields(v []string)`

SetValueFields sets ValueFields field to given value.

### HasValueFields

`func (o *CreateFirstSeenRuleRequestBodyFields) HasValueFields() bool`

HasValueFields returns a boolean if a field has been set.

### GetValueExpression

`func (o *CreateFirstSeenRuleRequestBodyFields) GetValueExpression() string`

GetValueExpression returns the ValueExpression field if non-nil, zero value otherwise.

### GetValueExpressionOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetValueExpressionOk() (*string, bool)`

GetValueExpressionOk returns a tuple with the ValueExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueExpression

`func (o *CreateFirstSeenRuleRequestBodyFields) SetValueExpression(v string)`

SetValueExpression sets ValueExpression field to given value.

### HasValueExpression

`func (o *CreateFirstSeenRuleRequestBodyFields) HasValueExpression() bool`

HasValueExpression returns a boolean if a field has been set.

### GetGroupByFields

`func (o *CreateFirstSeenRuleRequestBodyFields) GetGroupByFields() []string`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetGroupByFieldsOk() (*[]string, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *CreateFirstSeenRuleRequestBodyFields) SetGroupByFields(v []string)`

SetGroupByFields sets GroupByFields field to given value.


### GetScore

`func (o *CreateFirstSeenRuleRequestBodyFields) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CreateFirstSeenRuleRequestBodyFields) SetScore(v int32)`

SetScore sets Score field to given value.


### GetVersion

`func (o *CreateFirstSeenRuleRequestBodyFields) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateFirstSeenRuleRequestBodyFields) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetBaselineWindowSize

`func (o *CreateFirstSeenRuleRequestBodyFields) GetBaselineWindowSize() string`

GetBaselineWindowSize returns the BaselineWindowSize field if non-nil, zero value otherwise.

### GetBaselineWindowSizeOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetBaselineWindowSizeOk() (*string, bool)`

GetBaselineWindowSizeOk returns a tuple with the BaselineWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineWindowSize

`func (o *CreateFirstSeenRuleRequestBodyFields) SetBaselineWindowSize(v string)`

SetBaselineWindowSize sets BaselineWindowSize field to given value.


### GetRetentionWindowSize

`func (o *CreateFirstSeenRuleRequestBodyFields) GetRetentionWindowSize() string`

GetRetentionWindowSize returns the RetentionWindowSize field if non-nil, zero value otherwise.

### GetRetentionWindowSizeOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetRetentionWindowSizeOk() (*string, bool)`

GetRetentionWindowSizeOk returns a tuple with the RetentionWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionWindowSize

`func (o *CreateFirstSeenRuleRequestBodyFields) SetRetentionWindowSize(v string)`

SetRetentionWindowSize sets RetentionWindowSize field to given value.


### GetBaselineType

`func (o *CreateFirstSeenRuleRequestBodyFields) GetBaselineType() string`

GetBaselineType returns the BaselineType field if non-nil, zero value otherwise.

### GetBaselineTypeOk

`func (o *CreateFirstSeenRuleRequestBodyFields) GetBaselineTypeOk() (*string, bool)`

GetBaselineTypeOk returns a tuple with the BaselineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineType

`func (o *CreateFirstSeenRuleRequestBodyFields) SetBaselineType(v string)`

SetBaselineType sets BaselineType field to given value.

### HasBaselineType

`func (o *CreateFirstSeenRuleRequestBodyFields) HasBaselineType() bool`

HasBaselineType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


