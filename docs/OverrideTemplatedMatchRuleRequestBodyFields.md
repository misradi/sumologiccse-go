# OverrideTemplatedMatchRuleRequestBodyFields

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
**DescriptionExpression** | Pointer to **string** |  | [optional] 
**NameExpression** | Pointer to **string** |  | [optional] 
**ScoreMapping** | Pointer to [**GetRulesResponseDataObjectsInnerOneOf2ScoreMapping**](GetRulesResponseDataObjectsInnerOneOf2ScoreMapping.md) |  | [optional] 

## Methods

### NewOverrideTemplatedMatchRuleRequestBodyFields

`func NewOverrideTemplatedMatchRuleRequestBodyFields(tuningExpressionIds []string, ) *OverrideTemplatedMatchRuleRequestBodyFields`

NewOverrideTemplatedMatchRuleRequestBodyFields instantiates a new OverrideTemplatedMatchRuleRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverrideTemplatedMatchRuleRequestBodyFieldsWithDefaults

`func NewOverrideTemplatedMatchRuleRequestBodyFieldsWithDefaults() *OverrideTemplatedMatchRuleRequestBodyFields`

NewOverrideTemplatedMatchRuleRequestBodyFieldsWithDefaults instantiates a new OverrideTemplatedMatchRuleRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTuningExpressionIds

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetTuningExpressionIds() []string`

GetTuningExpressionIds returns the TuningExpressionIds field if non-nil, zero value otherwise.

### GetTuningExpressionIdsOk

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetTuningExpressionIdsOk() (*[]string, bool)`

GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuningExpressionIds

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetTuningExpressionIds(v []string)`

SetTuningExpressionIds sets TuningExpressionIds field to given value.


### GetIsPrototype

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetIsPrototype() bool`

GetIsPrototype returns the IsPrototype field if non-nil, zero value otherwise.

### GetIsPrototypeOk

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool)`

GetIsPrototypeOk returns a tuple with the IsPrototype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototype

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetIsPrototype(v bool)`

SetIsPrototype sets IsPrototype field to given value.

### HasIsPrototype

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) HasIsPrototype() bool`

HasIsPrototype returns a boolean if a field has been set.

### GetEntitySelectors

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner`

GetEntitySelectors returns the EntitySelectors field if non-nil, zero value otherwise.

### GetEntitySelectorsOk

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetEntitySelectorsOk() (*[]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool)`

GetEntitySelectorsOk returns a tuple with the EntitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectors

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner)`

SetEntitySelectors sets EntitySelectors field to given value.

### HasEntitySelectors

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) HasEntitySelectors() bool`

HasEntitySelectors returns a boolean if a field has been set.

### GetName

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleDescription

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetRuleDescription() string`

GetRuleDescription returns the RuleDescription field if non-nil, zero value otherwise.

### GetRuleDescriptionOk

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetRuleDescriptionOk() (*string, bool)`

GetRuleDescriptionOk returns a tuple with the RuleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDescription

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetRuleDescription(v string)`

SetRuleDescription sets RuleDescription field to given value.

### HasRuleDescription

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) HasRuleDescription() bool`

HasRuleDescription returns a boolean if a field has been set.

### GetSummaryExpression

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetSummaryExpression() string`

GetSummaryExpression returns the SummaryExpression field if non-nil, zero value otherwise.

### GetSummaryExpressionOk

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool)`

GetSummaryExpressionOk returns a tuple with the SummaryExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpression

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetSummaryExpression(v string)`

SetSummaryExpression sets SummaryExpression field to given value.

### HasSummaryExpression

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) HasSummaryExpression() bool`

HasSummaryExpression returns a boolean if a field has been set.

### GetTags

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDescriptionExpression

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetDescriptionExpression() string`

GetDescriptionExpression returns the DescriptionExpression field if non-nil, zero value otherwise.

### GetDescriptionExpressionOk

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetDescriptionExpressionOk() (*string, bool)`

GetDescriptionExpressionOk returns a tuple with the DescriptionExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionExpression

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetDescriptionExpression(v string)`

SetDescriptionExpression sets DescriptionExpression field to given value.

### HasDescriptionExpression

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) HasDescriptionExpression() bool`

HasDescriptionExpression returns a boolean if a field has been set.

### GetNameExpression

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetNameExpression() string`

GetNameExpression returns the NameExpression field if non-nil, zero value otherwise.

### GetNameExpressionOk

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetNameExpressionOk() (*string, bool)`

GetNameExpressionOk returns a tuple with the NameExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameExpression

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetNameExpression(v string)`

SetNameExpression sets NameExpression field to given value.

### HasNameExpression

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) HasNameExpression() bool`

HasNameExpression returns a boolean if a field has been set.

### GetScoreMapping

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetScoreMapping() GetRulesResponseDataObjectsInnerOneOf2ScoreMapping`

GetScoreMapping returns the ScoreMapping field if non-nil, zero value otherwise.

### GetScoreMappingOk

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetScoreMappingOk() (*GetRulesResponseDataObjectsInnerOneOf2ScoreMapping, bool)`

GetScoreMappingOk returns a tuple with the ScoreMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreMapping

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetScoreMapping(v GetRulesResponseDataObjectsInnerOneOf2ScoreMapping)`

SetScoreMapping sets ScoreMapping field to given value.

### HasScoreMapping

`func (o *OverrideTemplatedMatchRuleRequestBodyFields) HasScoreMapping() bool`

HasScoreMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


