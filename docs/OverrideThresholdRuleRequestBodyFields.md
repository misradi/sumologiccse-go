# OverrideThresholdRuleRequestBodyFields

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
**Limit** | Pointer to **int32** |  | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**WindowSize** | Pointer to **string** |  | [optional] 
**WindowSizeMilliseconds** | Pointer to **string** | Can be used instead of windowSize. | [optional] 

## Methods

### NewOverrideThresholdRuleRequestBodyFields

`func NewOverrideThresholdRuleRequestBodyFields(tuningExpressionIds []string, ) *OverrideThresholdRuleRequestBodyFields`

NewOverrideThresholdRuleRequestBodyFields instantiates a new OverrideThresholdRuleRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverrideThresholdRuleRequestBodyFieldsWithDefaults

`func NewOverrideThresholdRuleRequestBodyFieldsWithDefaults() *OverrideThresholdRuleRequestBodyFields`

NewOverrideThresholdRuleRequestBodyFieldsWithDefaults instantiates a new OverrideThresholdRuleRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTuningExpressionIds

`func (o *OverrideThresholdRuleRequestBodyFields) GetTuningExpressionIds() []string`

GetTuningExpressionIds returns the TuningExpressionIds field if non-nil, zero value otherwise.

### GetTuningExpressionIdsOk

`func (o *OverrideThresholdRuleRequestBodyFields) GetTuningExpressionIdsOk() (*[]string, bool)`

GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuningExpressionIds

`func (o *OverrideThresholdRuleRequestBodyFields) SetTuningExpressionIds(v []string)`

SetTuningExpressionIds sets TuningExpressionIds field to given value.


### GetIsPrototype

`func (o *OverrideThresholdRuleRequestBodyFields) GetIsPrototype() bool`

GetIsPrototype returns the IsPrototype field if non-nil, zero value otherwise.

### GetIsPrototypeOk

`func (o *OverrideThresholdRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool)`

GetIsPrototypeOk returns a tuple with the IsPrototype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototype

`func (o *OverrideThresholdRuleRequestBodyFields) SetIsPrototype(v bool)`

SetIsPrototype sets IsPrototype field to given value.

### HasIsPrototype

`func (o *OverrideThresholdRuleRequestBodyFields) HasIsPrototype() bool`

HasIsPrototype returns a boolean if a field has been set.

### GetEntitySelectors

`func (o *OverrideThresholdRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner`

GetEntitySelectors returns the EntitySelectors field if non-nil, zero value otherwise.

### GetEntitySelectorsOk

`func (o *OverrideThresholdRuleRequestBodyFields) GetEntitySelectorsOk() (*[]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool)`

GetEntitySelectorsOk returns a tuple with the EntitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectors

`func (o *OverrideThresholdRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner)`

SetEntitySelectors sets EntitySelectors field to given value.

### HasEntitySelectors

`func (o *OverrideThresholdRuleRequestBodyFields) HasEntitySelectors() bool`

HasEntitySelectors returns a boolean if a field has been set.

### GetName

`func (o *OverrideThresholdRuleRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OverrideThresholdRuleRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OverrideThresholdRuleRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OverrideThresholdRuleRequestBodyFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleDescription

`func (o *OverrideThresholdRuleRequestBodyFields) GetRuleDescription() string`

GetRuleDescription returns the RuleDescription field if non-nil, zero value otherwise.

### GetRuleDescriptionOk

`func (o *OverrideThresholdRuleRequestBodyFields) GetRuleDescriptionOk() (*string, bool)`

GetRuleDescriptionOk returns a tuple with the RuleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDescription

`func (o *OverrideThresholdRuleRequestBodyFields) SetRuleDescription(v string)`

SetRuleDescription sets RuleDescription field to given value.

### HasRuleDescription

`func (o *OverrideThresholdRuleRequestBodyFields) HasRuleDescription() bool`

HasRuleDescription returns a boolean if a field has been set.

### GetSummaryExpression

`func (o *OverrideThresholdRuleRequestBodyFields) GetSummaryExpression() string`

GetSummaryExpression returns the SummaryExpression field if non-nil, zero value otherwise.

### GetSummaryExpressionOk

`func (o *OverrideThresholdRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool)`

GetSummaryExpressionOk returns a tuple with the SummaryExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpression

`func (o *OverrideThresholdRuleRequestBodyFields) SetSummaryExpression(v string)`

SetSummaryExpression sets SummaryExpression field to given value.

### HasSummaryExpression

`func (o *OverrideThresholdRuleRequestBodyFields) HasSummaryExpression() bool`

HasSummaryExpression returns a boolean if a field has been set.

### GetTags

`func (o *OverrideThresholdRuleRequestBodyFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OverrideThresholdRuleRequestBodyFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OverrideThresholdRuleRequestBodyFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OverrideThresholdRuleRequestBodyFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDescription

`func (o *OverrideThresholdRuleRequestBodyFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OverrideThresholdRuleRequestBodyFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OverrideThresholdRuleRequestBodyFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OverrideThresholdRuleRequestBodyFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroupByFields

`func (o *OverrideThresholdRuleRequestBodyFields) GetGroupByFields() []string`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *OverrideThresholdRuleRequestBodyFields) GetGroupByFieldsOk() (*[]string, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *OverrideThresholdRuleRequestBodyFields) SetGroupByFields(v []string)`

SetGroupByFields sets GroupByFields field to given value.

### HasGroupByFields

`func (o *OverrideThresholdRuleRequestBodyFields) HasGroupByFields() bool`

HasGroupByFields returns a boolean if a field has been set.

### GetLimit

`func (o *OverrideThresholdRuleRequestBodyFields) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *OverrideThresholdRuleRequestBodyFields) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *OverrideThresholdRuleRequestBodyFields) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *OverrideThresholdRuleRequestBodyFields) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetScore

`func (o *OverrideThresholdRuleRequestBodyFields) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *OverrideThresholdRuleRequestBodyFields) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *OverrideThresholdRuleRequestBodyFields) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *OverrideThresholdRuleRequestBodyFields) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetWindowSize

`func (o *OverrideThresholdRuleRequestBodyFields) GetWindowSize() string`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *OverrideThresholdRuleRequestBodyFields) GetWindowSizeOk() (*string, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *OverrideThresholdRuleRequestBodyFields) SetWindowSize(v string)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *OverrideThresholdRuleRequestBodyFields) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.

### GetWindowSizeMilliseconds

`func (o *OverrideThresholdRuleRequestBodyFields) GetWindowSizeMilliseconds() string`

GetWindowSizeMilliseconds returns the WindowSizeMilliseconds field if non-nil, zero value otherwise.

### GetWindowSizeMillisecondsOk

`func (o *OverrideThresholdRuleRequestBodyFields) GetWindowSizeMillisecondsOk() (*string, bool)`

GetWindowSizeMillisecondsOk returns a tuple with the WindowSizeMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSizeMilliseconds

`func (o *OverrideThresholdRuleRequestBodyFields) SetWindowSizeMilliseconds(v string)`

SetWindowSizeMilliseconds sets WindowSizeMilliseconds field to given value.

### HasWindowSizeMilliseconds

`func (o *OverrideThresholdRuleRequestBodyFields) HasWindowSizeMilliseconds() bool`

HasWindowSizeMilliseconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


