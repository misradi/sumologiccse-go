# OverrideChainRuleRequestBodyFields

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
**Description** | Pointer to **string** |  | [optional] 
**GroupByFields** | Pointer to **[]string** |  | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**WindowSize** | Pointer to **string** |  | [optional] 
**WindowSizeMilliseconds** | Pointer to **string** | Can be used instead of windowSize. | [optional] 
**Limits** | Pointer to [**[]OverrideChainRuleRequestBodyFieldsLimitsInner**](OverrideChainRuleRequestBodyFieldsLimitsInner.md) |  | [optional] 

## Methods

### NewOverrideChainRuleRequestBodyFields

`func NewOverrideChainRuleRequestBodyFields(tuningExpressionIds []string, ) *OverrideChainRuleRequestBodyFields`

NewOverrideChainRuleRequestBodyFields instantiates a new OverrideChainRuleRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverrideChainRuleRequestBodyFieldsWithDefaults

`func NewOverrideChainRuleRequestBodyFieldsWithDefaults() *OverrideChainRuleRequestBodyFields`

NewOverrideChainRuleRequestBodyFieldsWithDefaults instantiates a new OverrideChainRuleRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTuningExpressionIds

`func (o *OverrideChainRuleRequestBodyFields) GetTuningExpressionIds() []string`

GetTuningExpressionIds returns the TuningExpressionIds field if non-nil, zero value otherwise.

### GetTuningExpressionIdsOk

`func (o *OverrideChainRuleRequestBodyFields) GetTuningExpressionIdsOk() (*[]string, bool)`

GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuningExpressionIds

`func (o *OverrideChainRuleRequestBodyFields) SetTuningExpressionIds(v []string)`

SetTuningExpressionIds sets TuningExpressionIds field to given value.


### GetIsPrototype

`func (o *OverrideChainRuleRequestBodyFields) GetIsPrototype() bool`

GetIsPrototype returns the IsPrototype field if non-nil, zero value otherwise.

### GetIsPrototypeOk

`func (o *OverrideChainRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool)`

GetIsPrototypeOk returns a tuple with the IsPrototype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototype

`func (o *OverrideChainRuleRequestBodyFields) SetIsPrototype(v bool)`

SetIsPrototype sets IsPrototype field to given value.

### HasIsPrototype

`func (o *OverrideChainRuleRequestBodyFields) HasIsPrototype() bool`

HasIsPrototype returns a boolean if a field has been set.

### GetEntitySelectors

`func (o *OverrideChainRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner`

GetEntitySelectors returns the EntitySelectors field if non-nil, zero value otherwise.

### GetEntitySelectorsOk

`func (o *OverrideChainRuleRequestBodyFields) GetEntitySelectorsOk() (*[]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool)`

GetEntitySelectorsOk returns a tuple with the EntitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectors

`func (o *OverrideChainRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner)`

SetEntitySelectors sets EntitySelectors field to given value.

### HasEntitySelectors

`func (o *OverrideChainRuleRequestBodyFields) HasEntitySelectors() bool`

HasEntitySelectors returns a boolean if a field has been set.

### GetName

`func (o *OverrideChainRuleRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OverrideChainRuleRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OverrideChainRuleRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OverrideChainRuleRequestBodyFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleDescription

`func (o *OverrideChainRuleRequestBodyFields) GetRuleDescription() string`

GetRuleDescription returns the RuleDescription field if non-nil, zero value otherwise.

### GetRuleDescriptionOk

`func (o *OverrideChainRuleRequestBodyFields) GetRuleDescriptionOk() (*string, bool)`

GetRuleDescriptionOk returns a tuple with the RuleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDescription

`func (o *OverrideChainRuleRequestBodyFields) SetRuleDescription(v string)`

SetRuleDescription sets RuleDescription field to given value.

### HasRuleDescription

`func (o *OverrideChainRuleRequestBodyFields) HasRuleDescription() bool`

HasRuleDescription returns a boolean if a field has been set.

### GetSummaryExpression

`func (o *OverrideChainRuleRequestBodyFields) GetSummaryExpression() string`

GetSummaryExpression returns the SummaryExpression field if non-nil, zero value otherwise.

### GetSummaryExpressionOk

`func (o *OverrideChainRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool)`

GetSummaryExpressionOk returns a tuple with the SummaryExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpression

`func (o *OverrideChainRuleRequestBodyFields) SetSummaryExpression(v string)`

SetSummaryExpression sets SummaryExpression field to given value.

### HasSummaryExpression

`func (o *OverrideChainRuleRequestBodyFields) HasSummaryExpression() bool`

HasSummaryExpression returns a boolean if a field has been set.

### GetTags

`func (o *OverrideChainRuleRequestBodyFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OverrideChainRuleRequestBodyFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OverrideChainRuleRequestBodyFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OverrideChainRuleRequestBodyFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDescription

`func (o *OverrideChainRuleRequestBodyFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OverrideChainRuleRequestBodyFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OverrideChainRuleRequestBodyFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OverrideChainRuleRequestBodyFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroupByFields

`func (o *OverrideChainRuleRequestBodyFields) GetGroupByFields() []string`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *OverrideChainRuleRequestBodyFields) GetGroupByFieldsOk() (*[]string, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *OverrideChainRuleRequestBodyFields) SetGroupByFields(v []string)`

SetGroupByFields sets GroupByFields field to given value.

### HasGroupByFields

`func (o *OverrideChainRuleRequestBodyFields) HasGroupByFields() bool`

HasGroupByFields returns a boolean if a field has been set.

### GetScore

`func (o *OverrideChainRuleRequestBodyFields) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *OverrideChainRuleRequestBodyFields) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *OverrideChainRuleRequestBodyFields) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *OverrideChainRuleRequestBodyFields) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetWindowSize

`func (o *OverrideChainRuleRequestBodyFields) GetWindowSize() string`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *OverrideChainRuleRequestBodyFields) GetWindowSizeOk() (*string, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *OverrideChainRuleRequestBodyFields) SetWindowSize(v string)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *OverrideChainRuleRequestBodyFields) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.

### GetWindowSizeMilliseconds

`func (o *OverrideChainRuleRequestBodyFields) GetWindowSizeMilliseconds() string`

GetWindowSizeMilliseconds returns the WindowSizeMilliseconds field if non-nil, zero value otherwise.

### GetWindowSizeMillisecondsOk

`func (o *OverrideChainRuleRequestBodyFields) GetWindowSizeMillisecondsOk() (*string, bool)`

GetWindowSizeMillisecondsOk returns a tuple with the WindowSizeMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSizeMilliseconds

`func (o *OverrideChainRuleRequestBodyFields) SetWindowSizeMilliseconds(v string)`

SetWindowSizeMilliseconds sets WindowSizeMilliseconds field to given value.

### HasWindowSizeMilliseconds

`func (o *OverrideChainRuleRequestBodyFields) HasWindowSizeMilliseconds() bool`

HasWindowSizeMilliseconds returns a boolean if a field has been set.

### GetLimits

`func (o *OverrideChainRuleRequestBodyFields) GetLimits() []OverrideChainRuleRequestBodyFieldsLimitsInner`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *OverrideChainRuleRequestBodyFields) GetLimitsOk() (*[]OverrideChainRuleRequestBodyFieldsLimitsInner, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *OverrideChainRuleRequestBodyFields) SetLimits(v []OverrideChainRuleRequestBodyFieldsLimitsInner)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *OverrideChainRuleRequestBodyFields) HasLimits() bool`

HasLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


