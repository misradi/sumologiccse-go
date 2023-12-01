# OverrideOutlierRuleRequestBodyFields

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
**NameExpression** | Pointer to **string** |  | [optional] 
**FloorValue** | Pointer to **string** | milliseconds | [optional] 
**GroupByFields** | Pointer to **[]string** |  | [optional] 
**DescriptionExpression** | Pointer to **string** |  | [optional] 
**BaselineWindowSize** | Pointer to **string** | milliseconds | [optional] 
**DeviationThreshold** | Pointer to **string** | milliseconds | [optional] 
**RetentionWindowSize** | Pointer to **string** | milliseconds | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**WindowSize** | Pointer to **string** |  | [optional] 
**WindowSizeMilliseconds** | Pointer to **string** | Can be used instead of windowSize. | [optional] 

## Methods

### NewOverrideOutlierRuleRequestBodyFields

`func NewOverrideOutlierRuleRequestBodyFields(tuningExpressionIds []string, ) *OverrideOutlierRuleRequestBodyFields`

NewOverrideOutlierRuleRequestBodyFields instantiates a new OverrideOutlierRuleRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverrideOutlierRuleRequestBodyFieldsWithDefaults

`func NewOverrideOutlierRuleRequestBodyFieldsWithDefaults() *OverrideOutlierRuleRequestBodyFields`

NewOverrideOutlierRuleRequestBodyFieldsWithDefaults instantiates a new OverrideOutlierRuleRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTuningExpressionIds

`func (o *OverrideOutlierRuleRequestBodyFields) GetTuningExpressionIds() []string`

GetTuningExpressionIds returns the TuningExpressionIds field if non-nil, zero value otherwise.

### GetTuningExpressionIdsOk

`func (o *OverrideOutlierRuleRequestBodyFields) GetTuningExpressionIdsOk() (*[]string, bool)`

GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuningExpressionIds

`func (o *OverrideOutlierRuleRequestBodyFields) SetTuningExpressionIds(v []string)`

SetTuningExpressionIds sets TuningExpressionIds field to given value.


### GetIsPrototype

`func (o *OverrideOutlierRuleRequestBodyFields) GetIsPrototype() bool`

GetIsPrototype returns the IsPrototype field if non-nil, zero value otherwise.

### GetIsPrototypeOk

`func (o *OverrideOutlierRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool)`

GetIsPrototypeOk returns a tuple with the IsPrototype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototype

`func (o *OverrideOutlierRuleRequestBodyFields) SetIsPrototype(v bool)`

SetIsPrototype sets IsPrototype field to given value.

### HasIsPrototype

`func (o *OverrideOutlierRuleRequestBodyFields) HasIsPrototype() bool`

HasIsPrototype returns a boolean if a field has been set.

### GetEntitySelectors

`func (o *OverrideOutlierRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner`

GetEntitySelectors returns the EntitySelectors field if non-nil, zero value otherwise.

### GetEntitySelectorsOk

`func (o *OverrideOutlierRuleRequestBodyFields) GetEntitySelectorsOk() (*[]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool)`

GetEntitySelectorsOk returns a tuple with the EntitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectors

`func (o *OverrideOutlierRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner)`

SetEntitySelectors sets EntitySelectors field to given value.

### HasEntitySelectors

`func (o *OverrideOutlierRuleRequestBodyFields) HasEntitySelectors() bool`

HasEntitySelectors returns a boolean if a field has been set.

### GetName

`func (o *OverrideOutlierRuleRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OverrideOutlierRuleRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OverrideOutlierRuleRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OverrideOutlierRuleRequestBodyFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleDescription

`func (o *OverrideOutlierRuleRequestBodyFields) GetRuleDescription() string`

GetRuleDescription returns the RuleDescription field if non-nil, zero value otherwise.

### GetRuleDescriptionOk

`func (o *OverrideOutlierRuleRequestBodyFields) GetRuleDescriptionOk() (*string, bool)`

GetRuleDescriptionOk returns a tuple with the RuleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDescription

`func (o *OverrideOutlierRuleRequestBodyFields) SetRuleDescription(v string)`

SetRuleDescription sets RuleDescription field to given value.

### HasRuleDescription

`func (o *OverrideOutlierRuleRequestBodyFields) HasRuleDescription() bool`

HasRuleDescription returns a boolean if a field has been set.

### GetSummaryExpression

`func (o *OverrideOutlierRuleRequestBodyFields) GetSummaryExpression() string`

GetSummaryExpression returns the SummaryExpression field if non-nil, zero value otherwise.

### GetSummaryExpressionOk

`func (o *OverrideOutlierRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool)`

GetSummaryExpressionOk returns a tuple with the SummaryExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpression

`func (o *OverrideOutlierRuleRequestBodyFields) SetSummaryExpression(v string)`

SetSummaryExpression sets SummaryExpression field to given value.

### HasSummaryExpression

`func (o *OverrideOutlierRuleRequestBodyFields) HasSummaryExpression() bool`

HasSummaryExpression returns a boolean if a field has been set.

### GetTags

`func (o *OverrideOutlierRuleRequestBodyFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OverrideOutlierRuleRequestBodyFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OverrideOutlierRuleRequestBodyFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OverrideOutlierRuleRequestBodyFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNameExpression

`func (o *OverrideOutlierRuleRequestBodyFields) GetNameExpression() string`

GetNameExpression returns the NameExpression field if non-nil, zero value otherwise.

### GetNameExpressionOk

`func (o *OverrideOutlierRuleRequestBodyFields) GetNameExpressionOk() (*string, bool)`

GetNameExpressionOk returns a tuple with the NameExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameExpression

`func (o *OverrideOutlierRuleRequestBodyFields) SetNameExpression(v string)`

SetNameExpression sets NameExpression field to given value.

### HasNameExpression

`func (o *OverrideOutlierRuleRequestBodyFields) HasNameExpression() bool`

HasNameExpression returns a boolean if a field has been set.

### GetFloorValue

`func (o *OverrideOutlierRuleRequestBodyFields) GetFloorValue() string`

GetFloorValue returns the FloorValue field if non-nil, zero value otherwise.

### GetFloorValueOk

`func (o *OverrideOutlierRuleRequestBodyFields) GetFloorValueOk() (*string, bool)`

GetFloorValueOk returns a tuple with the FloorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorValue

`func (o *OverrideOutlierRuleRequestBodyFields) SetFloorValue(v string)`

SetFloorValue sets FloorValue field to given value.

### HasFloorValue

`func (o *OverrideOutlierRuleRequestBodyFields) HasFloorValue() bool`

HasFloorValue returns a boolean if a field has been set.

### GetGroupByFields

`func (o *OverrideOutlierRuleRequestBodyFields) GetGroupByFields() []string`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *OverrideOutlierRuleRequestBodyFields) GetGroupByFieldsOk() (*[]string, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *OverrideOutlierRuleRequestBodyFields) SetGroupByFields(v []string)`

SetGroupByFields sets GroupByFields field to given value.

### HasGroupByFields

`func (o *OverrideOutlierRuleRequestBodyFields) HasGroupByFields() bool`

HasGroupByFields returns a boolean if a field has been set.

### GetDescriptionExpression

`func (o *OverrideOutlierRuleRequestBodyFields) GetDescriptionExpression() string`

GetDescriptionExpression returns the DescriptionExpression field if non-nil, zero value otherwise.

### GetDescriptionExpressionOk

`func (o *OverrideOutlierRuleRequestBodyFields) GetDescriptionExpressionOk() (*string, bool)`

GetDescriptionExpressionOk returns a tuple with the DescriptionExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionExpression

`func (o *OverrideOutlierRuleRequestBodyFields) SetDescriptionExpression(v string)`

SetDescriptionExpression sets DescriptionExpression field to given value.

### HasDescriptionExpression

`func (o *OverrideOutlierRuleRequestBodyFields) HasDescriptionExpression() bool`

HasDescriptionExpression returns a boolean if a field has been set.

### GetBaselineWindowSize

`func (o *OverrideOutlierRuleRequestBodyFields) GetBaselineWindowSize() string`

GetBaselineWindowSize returns the BaselineWindowSize field if non-nil, zero value otherwise.

### GetBaselineWindowSizeOk

`func (o *OverrideOutlierRuleRequestBodyFields) GetBaselineWindowSizeOk() (*string, bool)`

GetBaselineWindowSizeOk returns a tuple with the BaselineWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineWindowSize

`func (o *OverrideOutlierRuleRequestBodyFields) SetBaselineWindowSize(v string)`

SetBaselineWindowSize sets BaselineWindowSize field to given value.

### HasBaselineWindowSize

`func (o *OverrideOutlierRuleRequestBodyFields) HasBaselineWindowSize() bool`

HasBaselineWindowSize returns a boolean if a field has been set.

### GetDeviationThreshold

`func (o *OverrideOutlierRuleRequestBodyFields) GetDeviationThreshold() string`

GetDeviationThreshold returns the DeviationThreshold field if non-nil, zero value otherwise.

### GetDeviationThresholdOk

`func (o *OverrideOutlierRuleRequestBodyFields) GetDeviationThresholdOk() (*string, bool)`

GetDeviationThresholdOk returns a tuple with the DeviationThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviationThreshold

`func (o *OverrideOutlierRuleRequestBodyFields) SetDeviationThreshold(v string)`

SetDeviationThreshold sets DeviationThreshold field to given value.

### HasDeviationThreshold

`func (o *OverrideOutlierRuleRequestBodyFields) HasDeviationThreshold() bool`

HasDeviationThreshold returns a boolean if a field has been set.

### GetRetentionWindowSize

`func (o *OverrideOutlierRuleRequestBodyFields) GetRetentionWindowSize() string`

GetRetentionWindowSize returns the RetentionWindowSize field if non-nil, zero value otherwise.

### GetRetentionWindowSizeOk

`func (o *OverrideOutlierRuleRequestBodyFields) GetRetentionWindowSizeOk() (*string, bool)`

GetRetentionWindowSizeOk returns a tuple with the RetentionWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionWindowSize

`func (o *OverrideOutlierRuleRequestBodyFields) SetRetentionWindowSize(v string)`

SetRetentionWindowSize sets RetentionWindowSize field to given value.

### HasRetentionWindowSize

`func (o *OverrideOutlierRuleRequestBodyFields) HasRetentionWindowSize() bool`

HasRetentionWindowSize returns a boolean if a field has been set.

### GetScore

`func (o *OverrideOutlierRuleRequestBodyFields) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *OverrideOutlierRuleRequestBodyFields) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *OverrideOutlierRuleRequestBodyFields) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *OverrideOutlierRuleRequestBodyFields) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetWindowSize

`func (o *OverrideOutlierRuleRequestBodyFields) GetWindowSize() string`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *OverrideOutlierRuleRequestBodyFields) GetWindowSizeOk() (*string, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *OverrideOutlierRuleRequestBodyFields) SetWindowSize(v string)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *OverrideOutlierRuleRequestBodyFields) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.

### GetWindowSizeMilliseconds

`func (o *OverrideOutlierRuleRequestBodyFields) GetWindowSizeMilliseconds() string`

GetWindowSizeMilliseconds returns the WindowSizeMilliseconds field if non-nil, zero value otherwise.

### GetWindowSizeMillisecondsOk

`func (o *OverrideOutlierRuleRequestBodyFields) GetWindowSizeMillisecondsOk() (*string, bool)`

GetWindowSizeMillisecondsOk returns a tuple with the WindowSizeMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSizeMilliseconds

`func (o *OverrideOutlierRuleRequestBodyFields) SetWindowSizeMilliseconds(v string)`

SetWindowSizeMilliseconds sets WindowSizeMilliseconds field to given value.

### HasWindowSizeMilliseconds

`func (o *OverrideOutlierRuleRequestBodyFields) HasWindowSizeMilliseconds() bool`

HasWindowSizeMilliseconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


