# OverrideFirstSeenRuleRequestBodyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TuningExpressionIds** | **[]string** |  | 
**IsPrototype** | Pointer to **bool** |  | [optional] 
**EntitySelectors** | Pointer to [**[]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner**](CreateMatchRuleRequestBodyFieldsEntitySelectorsInner.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RuleDescription** | Pointer to **string** |  | [optional] 
**SummaryExpression** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**BaselineWindowSize** | Pointer to **string** | milliseconds | [optional] 
**DescriptionExpression** | Pointer to **string** |  | [optional] 
**GroupByFields** | Pointer to **[]string** |  | [optional] 
**NameExpression** | Pointer to **string** |  | [optional] 
**RetentionWindowSize** | Pointer to **string** | milliseconds | [optional] 
**Score** | Pointer to **int32** |  | [optional] 

## Methods

### NewOverrideFirstSeenRuleRequestBodyFields

`func NewOverrideFirstSeenRuleRequestBodyFields(tuningExpressionIds []string, ) *OverrideFirstSeenRuleRequestBodyFields`

NewOverrideFirstSeenRuleRequestBodyFields instantiates a new OverrideFirstSeenRuleRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverrideFirstSeenRuleRequestBodyFieldsWithDefaults

`func NewOverrideFirstSeenRuleRequestBodyFieldsWithDefaults() *OverrideFirstSeenRuleRequestBodyFields`

NewOverrideFirstSeenRuleRequestBodyFieldsWithDefaults instantiates a new OverrideFirstSeenRuleRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTuningExpressionIds

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetTuningExpressionIds() []string`

GetTuningExpressionIds returns the TuningExpressionIds field if non-nil, zero value otherwise.

### GetTuningExpressionIdsOk

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetTuningExpressionIdsOk() (*[]string, bool)`

GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuningExpressionIds

`func (o *OverrideFirstSeenRuleRequestBodyFields) SetTuningExpressionIds(v []string)`

SetTuningExpressionIds sets TuningExpressionIds field to given value.


### GetIsPrototype

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetIsPrototype() bool`

GetIsPrototype returns the IsPrototype field if non-nil, zero value otherwise.

### GetIsPrototypeOk

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool)`

GetIsPrototypeOk returns a tuple with the IsPrototype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototype

`func (o *OverrideFirstSeenRuleRequestBodyFields) SetIsPrototype(v bool)`

SetIsPrototype sets IsPrototype field to given value.

### HasIsPrototype

`func (o *OverrideFirstSeenRuleRequestBodyFields) HasIsPrototype() bool`

HasIsPrototype returns a boolean if a field has been set.

### GetEntitySelectors

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner`

GetEntitySelectors returns the EntitySelectors field if non-nil, zero value otherwise.

### GetEntitySelectorsOk

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetEntitySelectorsOk() (*[]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool)`

GetEntitySelectorsOk returns a tuple with the EntitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectors

`func (o *OverrideFirstSeenRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner)`

SetEntitySelectors sets EntitySelectors field to given value.

### HasEntitySelectors

`func (o *OverrideFirstSeenRuleRequestBodyFields) HasEntitySelectors() bool`

HasEntitySelectors returns a boolean if a field has been set.

### GetName

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OverrideFirstSeenRuleRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OverrideFirstSeenRuleRequestBodyFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleDescription

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetRuleDescription() string`

GetRuleDescription returns the RuleDescription field if non-nil, zero value otherwise.

### GetRuleDescriptionOk

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetRuleDescriptionOk() (*string, bool)`

GetRuleDescriptionOk returns a tuple with the RuleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDescription

`func (o *OverrideFirstSeenRuleRequestBodyFields) SetRuleDescription(v string)`

SetRuleDescription sets RuleDescription field to given value.

### HasRuleDescription

`func (o *OverrideFirstSeenRuleRequestBodyFields) HasRuleDescription() bool`

HasRuleDescription returns a boolean if a field has been set.

### GetSummaryExpression

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetSummaryExpression() string`

GetSummaryExpression returns the SummaryExpression field if non-nil, zero value otherwise.

### GetSummaryExpressionOk

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool)`

GetSummaryExpressionOk returns a tuple with the SummaryExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpression

`func (o *OverrideFirstSeenRuleRequestBodyFields) SetSummaryExpression(v string)`

SetSummaryExpression sets SummaryExpression field to given value.

### HasSummaryExpression

`func (o *OverrideFirstSeenRuleRequestBodyFields) HasSummaryExpression() bool`

HasSummaryExpression returns a boolean if a field has been set.

### GetTags

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OverrideFirstSeenRuleRequestBodyFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OverrideFirstSeenRuleRequestBodyFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetBaselineWindowSize

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetBaselineWindowSize() string`

GetBaselineWindowSize returns the BaselineWindowSize field if non-nil, zero value otherwise.

### GetBaselineWindowSizeOk

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetBaselineWindowSizeOk() (*string, bool)`

GetBaselineWindowSizeOk returns a tuple with the BaselineWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineWindowSize

`func (o *OverrideFirstSeenRuleRequestBodyFields) SetBaselineWindowSize(v string)`

SetBaselineWindowSize sets BaselineWindowSize field to given value.

### HasBaselineWindowSize

`func (o *OverrideFirstSeenRuleRequestBodyFields) HasBaselineWindowSize() bool`

HasBaselineWindowSize returns a boolean if a field has been set.

### GetDescriptionExpression

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetDescriptionExpression() string`

GetDescriptionExpression returns the DescriptionExpression field if non-nil, zero value otherwise.

### GetDescriptionExpressionOk

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetDescriptionExpressionOk() (*string, bool)`

GetDescriptionExpressionOk returns a tuple with the DescriptionExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionExpression

`func (o *OverrideFirstSeenRuleRequestBodyFields) SetDescriptionExpression(v string)`

SetDescriptionExpression sets DescriptionExpression field to given value.

### HasDescriptionExpression

`func (o *OverrideFirstSeenRuleRequestBodyFields) HasDescriptionExpression() bool`

HasDescriptionExpression returns a boolean if a field has been set.

### GetGroupByFields

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetGroupByFields() []string`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetGroupByFieldsOk() (*[]string, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *OverrideFirstSeenRuleRequestBodyFields) SetGroupByFields(v []string)`

SetGroupByFields sets GroupByFields field to given value.

### HasGroupByFields

`func (o *OverrideFirstSeenRuleRequestBodyFields) HasGroupByFields() bool`

HasGroupByFields returns a boolean if a field has been set.

### GetNameExpression

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetNameExpression() string`

GetNameExpression returns the NameExpression field if non-nil, zero value otherwise.

### GetNameExpressionOk

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetNameExpressionOk() (*string, bool)`

GetNameExpressionOk returns a tuple with the NameExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameExpression

`func (o *OverrideFirstSeenRuleRequestBodyFields) SetNameExpression(v string)`

SetNameExpression sets NameExpression field to given value.

### HasNameExpression

`func (o *OverrideFirstSeenRuleRequestBodyFields) HasNameExpression() bool`

HasNameExpression returns a boolean if a field has been set.

### GetRetentionWindowSize

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetRetentionWindowSize() string`

GetRetentionWindowSize returns the RetentionWindowSize field if non-nil, zero value otherwise.

### GetRetentionWindowSizeOk

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetRetentionWindowSizeOk() (*string, bool)`

GetRetentionWindowSizeOk returns a tuple with the RetentionWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionWindowSize

`func (o *OverrideFirstSeenRuleRequestBodyFields) SetRetentionWindowSize(v string)`

SetRetentionWindowSize sets RetentionWindowSize field to given value.

### HasRetentionWindowSize

`func (o *OverrideFirstSeenRuleRequestBodyFields) HasRetentionWindowSize() bool`

HasRetentionWindowSize returns a boolean if a field has been set.

### GetScore

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *OverrideFirstSeenRuleRequestBodyFields) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *OverrideFirstSeenRuleRequestBodyFields) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *OverrideFirstSeenRuleRequestBodyFields) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


