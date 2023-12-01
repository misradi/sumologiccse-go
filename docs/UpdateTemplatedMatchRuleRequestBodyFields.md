# UpdateTemplatedMatchRuleRequestBodyFields

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
**Expression** | **string** |  | 
**NameExpression** | **string** |  | 
**ScoreMapping** | [**GetRulesResponseDataObjectsInnerOneOf2ScoreMapping**](GetRulesResponseDataObjectsInnerOneOf2ScoreMapping.md) |  | 
**Stream** | **string** |  | 

## Methods

### NewUpdateTemplatedMatchRuleRequestBodyFields

`func NewUpdateTemplatedMatchRuleRequestBodyFields(name string, descriptionExpression string, expression string, nameExpression string, scoreMapping GetRulesResponseDataObjectsInnerOneOf2ScoreMapping, stream string, ) *UpdateTemplatedMatchRuleRequestBodyFields`

NewUpdateTemplatedMatchRuleRequestBodyFields instantiates a new UpdateTemplatedMatchRuleRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTemplatedMatchRuleRequestBodyFieldsWithDefaults

`func NewUpdateTemplatedMatchRuleRequestBodyFieldsWithDefaults() *UpdateTemplatedMatchRuleRequestBodyFields`

NewUpdateTemplatedMatchRuleRequestBodyFieldsWithDefaults instantiates a new UpdateTemplatedMatchRuleRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetField

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetAssetField() string`

GetAssetField returns the AssetField field if non-nil, zero value otherwise.

### GetAssetFieldOk

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetAssetFieldOk() (*string, bool)`

GetAssetFieldOk returns a tuple with the AssetField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetField

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) SetAssetField(v string)`

SetAssetField sets AssetField field to given value.

### HasAssetField

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) HasAssetField() bool`

HasAssetField returns a boolean if a field has been set.

### GetCategory

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEntitySelectors

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner`

GetEntitySelectors returns the EntitySelectors field if non-nil, zero value otherwise.

### GetEntitySelectorsOk

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetEntitySelectorsOk() (*[]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool)`

GetEntitySelectorsOk returns a tuple with the EntitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectors

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner)`

SetEntitySelectors sets EntitySelectors field to given value.

### HasEntitySelectors

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) HasEntitySelectors() bool`

HasEntitySelectors returns a boolean if a field has been set.

### GetIsPrototype

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetIsPrototype() bool`

GetIsPrototype returns the IsPrototype field if non-nil, zero value otherwise.

### GetIsPrototypeOk

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool)`

GetIsPrototypeOk returns a tuple with the IsPrototype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototype

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) SetIsPrototype(v bool)`

SetIsPrototype sets IsPrototype field to given value.

### HasIsPrototype

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) HasIsPrototype() bool`

HasIsPrototype returns a boolean if a field has been set.

### GetName

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.


### GetParentJaskId

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetParentJaskId() string`

GetParentJaskId returns the ParentJaskId field if non-nil, zero value otherwise.

### GetParentJaskIdOk

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetParentJaskIdOk() (*string, bool)`

GetParentJaskIdOk returns a tuple with the ParentJaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJaskId

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) SetParentJaskId(v string)`

SetParentJaskId sets ParentJaskId field to given value.

### HasParentJaskId

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) HasParentJaskId() bool`

HasParentJaskId returns a boolean if a field has been set.

### GetSummaryExpression

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetSummaryExpression() string`

GetSummaryExpression returns the SummaryExpression field if non-nil, zero value otherwise.

### GetSummaryExpressionOk

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool)`

GetSummaryExpressionOk returns a tuple with the SummaryExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpression

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) SetSummaryExpression(v string)`

SetSummaryExpression sets SummaryExpression field to given value.

### HasSummaryExpression

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) HasSummaryExpression() bool`

HasSummaryExpression returns a boolean if a field has been set.

### GetTags

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTuningExpressionIds

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetTuningExpressionIds() []string`

GetTuningExpressionIds returns the TuningExpressionIds field if non-nil, zero value otherwise.

### GetTuningExpressionIdsOk

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetTuningExpressionIdsOk() (*[]string, bool)`

GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuningExpressionIds

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) SetTuningExpressionIds(v []string)`

SetTuningExpressionIds sets TuningExpressionIds field to given value.

### HasTuningExpressionIds

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) HasTuningExpressionIds() bool`

HasTuningExpressionIds returns a boolean if a field has been set.

### GetDescriptionExpression

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetDescriptionExpression() string`

GetDescriptionExpression returns the DescriptionExpression field if non-nil, zero value otherwise.

### GetDescriptionExpressionOk

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetDescriptionExpressionOk() (*string, bool)`

GetDescriptionExpressionOk returns a tuple with the DescriptionExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionExpression

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) SetDescriptionExpression(v string)`

SetDescriptionExpression sets DescriptionExpression field to given value.


### GetExpression

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetNameExpression

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetNameExpression() string`

GetNameExpression returns the NameExpression field if non-nil, zero value otherwise.

### GetNameExpressionOk

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetNameExpressionOk() (*string, bool)`

GetNameExpressionOk returns a tuple with the NameExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameExpression

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) SetNameExpression(v string)`

SetNameExpression sets NameExpression field to given value.


### GetScoreMapping

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetScoreMapping() GetRulesResponseDataObjectsInnerOneOf2ScoreMapping`

GetScoreMapping returns the ScoreMapping field if non-nil, zero value otherwise.

### GetScoreMappingOk

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetScoreMappingOk() (*GetRulesResponseDataObjectsInnerOneOf2ScoreMapping, bool)`

GetScoreMappingOk returns a tuple with the ScoreMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreMapping

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) SetScoreMapping(v GetRulesResponseDataObjectsInnerOneOf2ScoreMapping)`

SetScoreMapping sets ScoreMapping field to given value.


### GetStream

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *UpdateTemplatedMatchRuleRequestBodyFields) SetStream(v string)`

SetStream sets Stream field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


