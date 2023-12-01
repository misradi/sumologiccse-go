# UpdateFirstSeenRuleRequestBodyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetField** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
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

### NewUpdateFirstSeenRuleRequestBodyFields

`func NewUpdateFirstSeenRuleRequestBodyFields(name string, descriptionExpression string, nameExpression string, filterExpression string, groupByFields []string, score int32, version int32, baselineWindowSize string, retentionWindowSize string, ) *UpdateFirstSeenRuleRequestBodyFields`

NewUpdateFirstSeenRuleRequestBodyFields instantiates a new UpdateFirstSeenRuleRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFirstSeenRuleRequestBodyFieldsWithDefaults

`func NewUpdateFirstSeenRuleRequestBodyFieldsWithDefaults() *UpdateFirstSeenRuleRequestBodyFields`

NewUpdateFirstSeenRuleRequestBodyFieldsWithDefaults instantiates a new UpdateFirstSeenRuleRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetField

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetAssetField() string`

GetAssetField returns the AssetField field if non-nil, zero value otherwise.

### GetAssetFieldOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetAssetFieldOk() (*string, bool)`

GetAssetFieldOk returns a tuple with the AssetField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetField

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetAssetField(v string)`

SetAssetField sets AssetField field to given value.

### HasAssetField

`func (o *UpdateFirstSeenRuleRequestBodyFields) HasAssetField() bool`

HasAssetField returns a boolean if a field has been set.

### GetCategory

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateFirstSeenRuleRequestBodyFields) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateFirstSeenRuleRequestBodyFields) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEntitySelectors

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner`

GetEntitySelectors returns the EntitySelectors field if non-nil, zero value otherwise.

### GetEntitySelectorsOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetEntitySelectorsOk() (*[]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool)`

GetEntitySelectorsOk returns a tuple with the EntitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectors

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner)`

SetEntitySelectors sets EntitySelectors field to given value.

### HasEntitySelectors

`func (o *UpdateFirstSeenRuleRequestBodyFields) HasEntitySelectors() bool`

HasEntitySelectors returns a boolean if a field has been set.

### GetIsPrototype

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetIsPrototype() bool`

GetIsPrototype returns the IsPrototype field if non-nil, zero value otherwise.

### GetIsPrototypeOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool)`

GetIsPrototypeOk returns a tuple with the IsPrototype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototype

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetIsPrototype(v bool)`

SetIsPrototype sets IsPrototype field to given value.

### HasIsPrototype

`func (o *UpdateFirstSeenRuleRequestBodyFields) HasIsPrototype() bool`

HasIsPrototype returns a boolean if a field has been set.

### GetName

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.


### GetParentJaskId

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetParentJaskId() string`

GetParentJaskId returns the ParentJaskId field if non-nil, zero value otherwise.

### GetParentJaskIdOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetParentJaskIdOk() (*string, bool)`

GetParentJaskIdOk returns a tuple with the ParentJaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJaskId

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetParentJaskId(v string)`

SetParentJaskId sets ParentJaskId field to given value.

### HasParentJaskId

`func (o *UpdateFirstSeenRuleRequestBodyFields) HasParentJaskId() bool`

HasParentJaskId returns a boolean if a field has been set.

### GetSummaryExpression

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetSummaryExpression() string`

GetSummaryExpression returns the SummaryExpression field if non-nil, zero value otherwise.

### GetSummaryExpressionOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool)`

GetSummaryExpressionOk returns a tuple with the SummaryExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpression

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetSummaryExpression(v string)`

SetSummaryExpression sets SummaryExpression field to given value.

### HasSummaryExpression

`func (o *UpdateFirstSeenRuleRequestBodyFields) HasSummaryExpression() bool`

HasSummaryExpression returns a boolean if a field has been set.

### GetTags

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateFirstSeenRuleRequestBodyFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTuningExpressionIds

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetTuningExpressionIds() []string`

GetTuningExpressionIds returns the TuningExpressionIds field if non-nil, zero value otherwise.

### GetTuningExpressionIdsOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetTuningExpressionIdsOk() (*[]string, bool)`

GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuningExpressionIds

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetTuningExpressionIds(v []string)`

SetTuningExpressionIds sets TuningExpressionIds field to given value.

### HasTuningExpressionIds

`func (o *UpdateFirstSeenRuleRequestBodyFields) HasTuningExpressionIds() bool`

HasTuningExpressionIds returns a boolean if a field has been set.

### GetDescriptionExpression

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetDescriptionExpression() string`

GetDescriptionExpression returns the DescriptionExpression field if non-nil, zero value otherwise.

### GetDescriptionExpressionOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetDescriptionExpressionOk() (*string, bool)`

GetDescriptionExpressionOk returns a tuple with the DescriptionExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionExpression

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetDescriptionExpression(v string)`

SetDescriptionExpression sets DescriptionExpression field to given value.


### GetNameExpression

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetNameExpression() string`

GetNameExpression returns the NameExpression field if non-nil, zero value otherwise.

### GetNameExpressionOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetNameExpressionOk() (*string, bool)`

GetNameExpressionOk returns a tuple with the NameExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameExpression

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetNameExpression(v string)`

SetNameExpression sets NameExpression field to given value.


### GetFilterExpression

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetFilterExpression() string`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetFilterExpressionOk() (*string, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetFilterExpression(v string)`

SetFilterExpression sets FilterExpression field to given value.


### GetValueFields

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetValueFields() []string`

GetValueFields returns the ValueFields field if non-nil, zero value otherwise.

### GetValueFieldsOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetValueFieldsOk() (*[]string, bool)`

GetValueFieldsOk returns a tuple with the ValueFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFields

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetValueFields(v []string)`

SetValueFields sets ValueFields field to given value.

### HasValueFields

`func (o *UpdateFirstSeenRuleRequestBodyFields) HasValueFields() bool`

HasValueFields returns a boolean if a field has been set.

### GetValueExpression

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetValueExpression() string`

GetValueExpression returns the ValueExpression field if non-nil, zero value otherwise.

### GetValueExpressionOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetValueExpressionOk() (*string, bool)`

GetValueExpressionOk returns a tuple with the ValueExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueExpression

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetValueExpression(v string)`

SetValueExpression sets ValueExpression field to given value.

### HasValueExpression

`func (o *UpdateFirstSeenRuleRequestBodyFields) HasValueExpression() bool`

HasValueExpression returns a boolean if a field has been set.

### GetGroupByFields

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetGroupByFields() []string`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetGroupByFieldsOk() (*[]string, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetGroupByFields(v []string)`

SetGroupByFields sets GroupByFields field to given value.


### GetScore

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetScore(v int32)`

SetScore sets Score field to given value.


### GetVersion

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetBaselineWindowSize

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetBaselineWindowSize() string`

GetBaselineWindowSize returns the BaselineWindowSize field if non-nil, zero value otherwise.

### GetBaselineWindowSizeOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetBaselineWindowSizeOk() (*string, bool)`

GetBaselineWindowSizeOk returns a tuple with the BaselineWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineWindowSize

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetBaselineWindowSize(v string)`

SetBaselineWindowSize sets BaselineWindowSize field to given value.


### GetRetentionWindowSize

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetRetentionWindowSize() string`

GetRetentionWindowSize returns the RetentionWindowSize field if non-nil, zero value otherwise.

### GetRetentionWindowSizeOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetRetentionWindowSizeOk() (*string, bool)`

GetRetentionWindowSizeOk returns a tuple with the RetentionWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionWindowSize

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetRetentionWindowSize(v string)`

SetRetentionWindowSize sets RetentionWindowSize field to given value.


### GetBaselineType

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetBaselineType() string`

GetBaselineType returns the BaselineType field if non-nil, zero value otherwise.

### GetBaselineTypeOk

`func (o *UpdateFirstSeenRuleRequestBodyFields) GetBaselineTypeOk() (*string, bool)`

GetBaselineTypeOk returns a tuple with the BaselineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineType

`func (o *UpdateFirstSeenRuleRequestBodyFields) SetBaselineType(v string)`

SetBaselineType sets BaselineType field to given value.

### HasBaselineType

`func (o *UpdateFirstSeenRuleRequestBodyFields) HasBaselineType() bool`

HasBaselineType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


