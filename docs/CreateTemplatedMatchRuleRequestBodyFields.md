# CreateTemplatedMatchRuleRequestBodyFields

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
**Expression** | **string** |  | 
**NameExpression** | **string** |  | 
**ScoreMapping** | [**GetRulesResponseDataObjectsInnerOneOf2ScoreMapping**](GetRulesResponseDataObjectsInnerOneOf2ScoreMapping.md) |  | 
**Stream** | **string** |  | 

## Methods

### NewCreateTemplatedMatchRuleRequestBodyFields

`func NewCreateTemplatedMatchRuleRequestBodyFields(enabled bool, name string, descriptionExpression string, expression string, nameExpression string, scoreMapping GetRulesResponseDataObjectsInnerOneOf2ScoreMapping, stream string, ) *CreateTemplatedMatchRuleRequestBodyFields`

NewCreateTemplatedMatchRuleRequestBodyFields instantiates a new CreateTemplatedMatchRuleRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTemplatedMatchRuleRequestBodyFieldsWithDefaults

`func NewCreateTemplatedMatchRuleRequestBodyFieldsWithDefaults() *CreateTemplatedMatchRuleRequestBodyFields`

NewCreateTemplatedMatchRuleRequestBodyFieldsWithDefaults instantiates a new CreateTemplatedMatchRuleRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetField

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetAssetField() string`

GetAssetField returns the AssetField field if non-nil, zero value otherwise.

### GetAssetFieldOk

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetAssetFieldOk() (*string, bool)`

GetAssetFieldOk returns a tuple with the AssetField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetField

`func (o *CreateTemplatedMatchRuleRequestBodyFields) SetAssetField(v string)`

SetAssetField sets AssetField field to given value.

### HasAssetField

`func (o *CreateTemplatedMatchRuleRequestBodyFields) HasAssetField() bool`

HasAssetField returns a boolean if a field has been set.

### GetCategory

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateTemplatedMatchRuleRequestBodyFields) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CreateTemplatedMatchRuleRequestBodyFields) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateTemplatedMatchRuleRequestBodyFields) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEntitySelectors

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner`

GetEntitySelectors returns the EntitySelectors field if non-nil, zero value otherwise.

### GetEntitySelectorsOk

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetEntitySelectorsOk() (*[]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool)`

GetEntitySelectorsOk returns a tuple with the EntitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectors

`func (o *CreateTemplatedMatchRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner)`

SetEntitySelectors sets EntitySelectors field to given value.

### HasEntitySelectors

`func (o *CreateTemplatedMatchRuleRequestBodyFields) HasEntitySelectors() bool`

HasEntitySelectors returns a boolean if a field has been set.

### GetIsPrototype

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetIsPrototype() bool`

GetIsPrototype returns the IsPrototype field if non-nil, zero value otherwise.

### GetIsPrototypeOk

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool)`

GetIsPrototypeOk returns a tuple with the IsPrototype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototype

`func (o *CreateTemplatedMatchRuleRequestBodyFields) SetIsPrototype(v bool)`

SetIsPrototype sets IsPrototype field to given value.

### HasIsPrototype

`func (o *CreateTemplatedMatchRuleRequestBodyFields) HasIsPrototype() bool`

HasIsPrototype returns a boolean if a field has been set.

### GetName

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTemplatedMatchRuleRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.


### GetParentJaskId

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetParentJaskId() string`

GetParentJaskId returns the ParentJaskId field if non-nil, zero value otherwise.

### GetParentJaskIdOk

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetParentJaskIdOk() (*string, bool)`

GetParentJaskIdOk returns a tuple with the ParentJaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJaskId

`func (o *CreateTemplatedMatchRuleRequestBodyFields) SetParentJaskId(v string)`

SetParentJaskId sets ParentJaskId field to given value.

### HasParentJaskId

`func (o *CreateTemplatedMatchRuleRequestBodyFields) HasParentJaskId() bool`

HasParentJaskId returns a boolean if a field has been set.

### GetSummaryExpression

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetSummaryExpression() string`

GetSummaryExpression returns the SummaryExpression field if non-nil, zero value otherwise.

### GetSummaryExpressionOk

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool)`

GetSummaryExpressionOk returns a tuple with the SummaryExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpression

`func (o *CreateTemplatedMatchRuleRequestBodyFields) SetSummaryExpression(v string)`

SetSummaryExpression sets SummaryExpression field to given value.

### HasSummaryExpression

`func (o *CreateTemplatedMatchRuleRequestBodyFields) HasSummaryExpression() bool`

HasSummaryExpression returns a boolean if a field has been set.

### GetTags

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateTemplatedMatchRuleRequestBodyFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateTemplatedMatchRuleRequestBodyFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTuningExpressionIds

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetTuningExpressionIds() []string`

GetTuningExpressionIds returns the TuningExpressionIds field if non-nil, zero value otherwise.

### GetTuningExpressionIdsOk

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetTuningExpressionIdsOk() (*[]string, bool)`

GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuningExpressionIds

`func (o *CreateTemplatedMatchRuleRequestBodyFields) SetTuningExpressionIds(v []string)`

SetTuningExpressionIds sets TuningExpressionIds field to given value.

### HasTuningExpressionIds

`func (o *CreateTemplatedMatchRuleRequestBodyFields) HasTuningExpressionIds() bool`

HasTuningExpressionIds returns a boolean if a field has been set.

### GetDescriptionExpression

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetDescriptionExpression() string`

GetDescriptionExpression returns the DescriptionExpression field if non-nil, zero value otherwise.

### GetDescriptionExpressionOk

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetDescriptionExpressionOk() (*string, bool)`

GetDescriptionExpressionOk returns a tuple with the DescriptionExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionExpression

`func (o *CreateTemplatedMatchRuleRequestBodyFields) SetDescriptionExpression(v string)`

SetDescriptionExpression sets DescriptionExpression field to given value.


### GetExpression

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *CreateTemplatedMatchRuleRequestBodyFields) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetNameExpression

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetNameExpression() string`

GetNameExpression returns the NameExpression field if non-nil, zero value otherwise.

### GetNameExpressionOk

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetNameExpressionOk() (*string, bool)`

GetNameExpressionOk returns a tuple with the NameExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameExpression

`func (o *CreateTemplatedMatchRuleRequestBodyFields) SetNameExpression(v string)`

SetNameExpression sets NameExpression field to given value.


### GetScoreMapping

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetScoreMapping() GetRulesResponseDataObjectsInnerOneOf2ScoreMapping`

GetScoreMapping returns the ScoreMapping field if non-nil, zero value otherwise.

### GetScoreMappingOk

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetScoreMappingOk() (*GetRulesResponseDataObjectsInnerOneOf2ScoreMapping, bool)`

GetScoreMappingOk returns a tuple with the ScoreMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreMapping

`func (o *CreateTemplatedMatchRuleRequestBodyFields) SetScoreMapping(v GetRulesResponseDataObjectsInnerOneOf2ScoreMapping)`

SetScoreMapping sets ScoreMapping field to given value.


### GetStream

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *CreateTemplatedMatchRuleRequestBodyFields) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *CreateTemplatedMatchRuleRequestBodyFields) SetStream(v string)`

SetStream sets Stream field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


