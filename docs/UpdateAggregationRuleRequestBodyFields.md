# UpdateAggregationRuleRequestBodyFields

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
**AggregationFunctions** | [**[]GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner**](GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner.md) |  | 
**DescriptionExpression** | **string** |  | 
**GroupByAsset** | **bool** |  | 
**GroupByFields** | **[]string** |  | 
**MatchExpression** | **string** |  | 
**NameExpression** | **string** |  | 
**ScoreMapping** | [**GetRulesResponseDataObjectsInnerOneOf2ScoreMapping**](GetRulesResponseDataObjectsInnerOneOf2ScoreMapping.md) |  | 
**Stream** | **string** |  | 
**TriggerExpression** | **string** |  | 
**WindowSize** | **string** |  | 
**WindowSizeMilliseconds** | Pointer to **string** | Can be used instead of windowSize. | [optional] 

## Methods

### NewUpdateAggregationRuleRequestBodyFields

`func NewUpdateAggregationRuleRequestBodyFields(name string, aggregationFunctions []GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner, descriptionExpression string, groupByAsset bool, groupByFields []string, matchExpression string, nameExpression string, scoreMapping GetRulesResponseDataObjectsInnerOneOf2ScoreMapping, stream string, triggerExpression string, windowSize string, ) *UpdateAggregationRuleRequestBodyFields`

NewUpdateAggregationRuleRequestBodyFields instantiates a new UpdateAggregationRuleRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAggregationRuleRequestBodyFieldsWithDefaults

`func NewUpdateAggregationRuleRequestBodyFieldsWithDefaults() *UpdateAggregationRuleRequestBodyFields`

NewUpdateAggregationRuleRequestBodyFieldsWithDefaults instantiates a new UpdateAggregationRuleRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetField

`func (o *UpdateAggregationRuleRequestBodyFields) GetAssetField() string`

GetAssetField returns the AssetField field if non-nil, zero value otherwise.

### GetAssetFieldOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetAssetFieldOk() (*string, bool)`

GetAssetFieldOk returns a tuple with the AssetField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetField

`func (o *UpdateAggregationRuleRequestBodyFields) SetAssetField(v string)`

SetAssetField sets AssetField field to given value.

### HasAssetField

`func (o *UpdateAggregationRuleRequestBodyFields) HasAssetField() bool`

HasAssetField returns a boolean if a field has been set.

### GetCategory

`func (o *UpdateAggregationRuleRequestBodyFields) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateAggregationRuleRequestBodyFields) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateAggregationRuleRequestBodyFields) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateAggregationRuleRequestBodyFields) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateAggregationRuleRequestBodyFields) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateAggregationRuleRequestBodyFields) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEntitySelectors

`func (o *UpdateAggregationRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner`

GetEntitySelectors returns the EntitySelectors field if non-nil, zero value otherwise.

### GetEntitySelectorsOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetEntitySelectorsOk() (*[]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool)`

GetEntitySelectorsOk returns a tuple with the EntitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectors

`func (o *UpdateAggregationRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner)`

SetEntitySelectors sets EntitySelectors field to given value.

### HasEntitySelectors

`func (o *UpdateAggregationRuleRequestBodyFields) HasEntitySelectors() bool`

HasEntitySelectors returns a boolean if a field has been set.

### GetIsPrototype

`func (o *UpdateAggregationRuleRequestBodyFields) GetIsPrototype() bool`

GetIsPrototype returns the IsPrototype field if non-nil, zero value otherwise.

### GetIsPrototypeOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool)`

GetIsPrototypeOk returns a tuple with the IsPrototype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototype

`func (o *UpdateAggregationRuleRequestBodyFields) SetIsPrototype(v bool)`

SetIsPrototype sets IsPrototype field to given value.

### HasIsPrototype

`func (o *UpdateAggregationRuleRequestBodyFields) HasIsPrototype() bool`

HasIsPrototype returns a boolean if a field has been set.

### GetName

`func (o *UpdateAggregationRuleRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAggregationRuleRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.


### GetParentJaskId

`func (o *UpdateAggregationRuleRequestBodyFields) GetParentJaskId() string`

GetParentJaskId returns the ParentJaskId field if non-nil, zero value otherwise.

### GetParentJaskIdOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetParentJaskIdOk() (*string, bool)`

GetParentJaskIdOk returns a tuple with the ParentJaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJaskId

`func (o *UpdateAggregationRuleRequestBodyFields) SetParentJaskId(v string)`

SetParentJaskId sets ParentJaskId field to given value.

### HasParentJaskId

`func (o *UpdateAggregationRuleRequestBodyFields) HasParentJaskId() bool`

HasParentJaskId returns a boolean if a field has been set.

### GetSummaryExpression

`func (o *UpdateAggregationRuleRequestBodyFields) GetSummaryExpression() string`

GetSummaryExpression returns the SummaryExpression field if non-nil, zero value otherwise.

### GetSummaryExpressionOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool)`

GetSummaryExpressionOk returns a tuple with the SummaryExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpression

`func (o *UpdateAggregationRuleRequestBodyFields) SetSummaryExpression(v string)`

SetSummaryExpression sets SummaryExpression field to given value.

### HasSummaryExpression

`func (o *UpdateAggregationRuleRequestBodyFields) HasSummaryExpression() bool`

HasSummaryExpression returns a boolean if a field has been set.

### GetTags

`func (o *UpdateAggregationRuleRequestBodyFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateAggregationRuleRequestBodyFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateAggregationRuleRequestBodyFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTuningExpressionIds

`func (o *UpdateAggregationRuleRequestBodyFields) GetTuningExpressionIds() []string`

GetTuningExpressionIds returns the TuningExpressionIds field if non-nil, zero value otherwise.

### GetTuningExpressionIdsOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetTuningExpressionIdsOk() (*[]string, bool)`

GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuningExpressionIds

`func (o *UpdateAggregationRuleRequestBodyFields) SetTuningExpressionIds(v []string)`

SetTuningExpressionIds sets TuningExpressionIds field to given value.

### HasTuningExpressionIds

`func (o *UpdateAggregationRuleRequestBodyFields) HasTuningExpressionIds() bool`

HasTuningExpressionIds returns a boolean if a field has been set.

### GetAggregationFunctions

`func (o *UpdateAggregationRuleRequestBodyFields) GetAggregationFunctions() []GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner`

GetAggregationFunctions returns the AggregationFunctions field if non-nil, zero value otherwise.

### GetAggregationFunctionsOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetAggregationFunctionsOk() (*[]GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner, bool)`

GetAggregationFunctionsOk returns a tuple with the AggregationFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationFunctions

`func (o *UpdateAggregationRuleRequestBodyFields) SetAggregationFunctions(v []GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner)`

SetAggregationFunctions sets AggregationFunctions field to given value.


### GetDescriptionExpression

`func (o *UpdateAggregationRuleRequestBodyFields) GetDescriptionExpression() string`

GetDescriptionExpression returns the DescriptionExpression field if non-nil, zero value otherwise.

### GetDescriptionExpressionOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetDescriptionExpressionOk() (*string, bool)`

GetDescriptionExpressionOk returns a tuple with the DescriptionExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionExpression

`func (o *UpdateAggregationRuleRequestBodyFields) SetDescriptionExpression(v string)`

SetDescriptionExpression sets DescriptionExpression field to given value.


### GetGroupByAsset

`func (o *UpdateAggregationRuleRequestBodyFields) GetGroupByAsset() bool`

GetGroupByAsset returns the GroupByAsset field if non-nil, zero value otherwise.

### GetGroupByAssetOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetGroupByAssetOk() (*bool, bool)`

GetGroupByAssetOk returns a tuple with the GroupByAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByAsset

`func (o *UpdateAggregationRuleRequestBodyFields) SetGroupByAsset(v bool)`

SetGroupByAsset sets GroupByAsset field to given value.


### GetGroupByFields

`func (o *UpdateAggregationRuleRequestBodyFields) GetGroupByFields() []string`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetGroupByFieldsOk() (*[]string, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *UpdateAggregationRuleRequestBodyFields) SetGroupByFields(v []string)`

SetGroupByFields sets GroupByFields field to given value.


### GetMatchExpression

`func (o *UpdateAggregationRuleRequestBodyFields) GetMatchExpression() string`

GetMatchExpression returns the MatchExpression field if non-nil, zero value otherwise.

### GetMatchExpressionOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetMatchExpressionOk() (*string, bool)`

GetMatchExpressionOk returns a tuple with the MatchExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchExpression

`func (o *UpdateAggregationRuleRequestBodyFields) SetMatchExpression(v string)`

SetMatchExpression sets MatchExpression field to given value.


### GetNameExpression

`func (o *UpdateAggregationRuleRequestBodyFields) GetNameExpression() string`

GetNameExpression returns the NameExpression field if non-nil, zero value otherwise.

### GetNameExpressionOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetNameExpressionOk() (*string, bool)`

GetNameExpressionOk returns a tuple with the NameExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameExpression

`func (o *UpdateAggregationRuleRequestBodyFields) SetNameExpression(v string)`

SetNameExpression sets NameExpression field to given value.


### GetScoreMapping

`func (o *UpdateAggregationRuleRequestBodyFields) GetScoreMapping() GetRulesResponseDataObjectsInnerOneOf2ScoreMapping`

GetScoreMapping returns the ScoreMapping field if non-nil, zero value otherwise.

### GetScoreMappingOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetScoreMappingOk() (*GetRulesResponseDataObjectsInnerOneOf2ScoreMapping, bool)`

GetScoreMappingOk returns a tuple with the ScoreMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreMapping

`func (o *UpdateAggregationRuleRequestBodyFields) SetScoreMapping(v GetRulesResponseDataObjectsInnerOneOf2ScoreMapping)`

SetScoreMapping sets ScoreMapping field to given value.


### GetStream

`func (o *UpdateAggregationRuleRequestBodyFields) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *UpdateAggregationRuleRequestBodyFields) SetStream(v string)`

SetStream sets Stream field to given value.


### GetTriggerExpression

`func (o *UpdateAggregationRuleRequestBodyFields) GetTriggerExpression() string`

GetTriggerExpression returns the TriggerExpression field if non-nil, zero value otherwise.

### GetTriggerExpressionOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetTriggerExpressionOk() (*string, bool)`

GetTriggerExpressionOk returns a tuple with the TriggerExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerExpression

`func (o *UpdateAggregationRuleRequestBodyFields) SetTriggerExpression(v string)`

SetTriggerExpression sets TriggerExpression field to given value.


### GetWindowSize

`func (o *UpdateAggregationRuleRequestBodyFields) GetWindowSize() string`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetWindowSizeOk() (*string, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *UpdateAggregationRuleRequestBodyFields) SetWindowSize(v string)`

SetWindowSize sets WindowSize field to given value.


### GetWindowSizeMilliseconds

`func (o *UpdateAggregationRuleRequestBodyFields) GetWindowSizeMilliseconds() string`

GetWindowSizeMilliseconds returns the WindowSizeMilliseconds field if non-nil, zero value otherwise.

### GetWindowSizeMillisecondsOk

`func (o *UpdateAggregationRuleRequestBodyFields) GetWindowSizeMillisecondsOk() (*string, bool)`

GetWindowSizeMillisecondsOk returns a tuple with the WindowSizeMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSizeMilliseconds

`func (o *UpdateAggregationRuleRequestBodyFields) SetWindowSizeMilliseconds(v string)`

SetWindowSizeMilliseconds sets WindowSizeMilliseconds field to given value.

### HasWindowSizeMilliseconds

`func (o *UpdateAggregationRuleRequestBodyFields) HasWindowSizeMilliseconds() bool`

HasWindowSizeMilliseconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


