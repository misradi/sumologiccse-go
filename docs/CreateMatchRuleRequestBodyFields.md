# CreateMatchRuleRequestBodyFields

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
**Description** | **string** | The description to be used on the generated Signal | 
**Expression** | **string** |  | 
**Score** | **int32** |  | 
**Stream** | **string** |  | 

## Methods

### NewCreateMatchRuleRequestBodyFields

`func NewCreateMatchRuleRequestBodyFields(enabled bool, name string, description string, expression string, score int32, stream string, ) *CreateMatchRuleRequestBodyFields`

NewCreateMatchRuleRequestBodyFields instantiates a new CreateMatchRuleRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMatchRuleRequestBodyFieldsWithDefaults

`func NewCreateMatchRuleRequestBodyFieldsWithDefaults() *CreateMatchRuleRequestBodyFields`

NewCreateMatchRuleRequestBodyFieldsWithDefaults instantiates a new CreateMatchRuleRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetField

`func (o *CreateMatchRuleRequestBodyFields) GetAssetField() string`

GetAssetField returns the AssetField field if non-nil, zero value otherwise.

### GetAssetFieldOk

`func (o *CreateMatchRuleRequestBodyFields) GetAssetFieldOk() (*string, bool)`

GetAssetFieldOk returns a tuple with the AssetField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetField

`func (o *CreateMatchRuleRequestBodyFields) SetAssetField(v string)`

SetAssetField sets AssetField field to given value.

### HasAssetField

`func (o *CreateMatchRuleRequestBodyFields) HasAssetField() bool`

HasAssetField returns a boolean if a field has been set.

### GetCategory

`func (o *CreateMatchRuleRequestBodyFields) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateMatchRuleRequestBodyFields) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateMatchRuleRequestBodyFields) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CreateMatchRuleRequestBodyFields) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateMatchRuleRequestBodyFields) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateMatchRuleRequestBodyFields) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateMatchRuleRequestBodyFields) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEntitySelectors

`func (o *CreateMatchRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner`

GetEntitySelectors returns the EntitySelectors field if non-nil, zero value otherwise.

### GetEntitySelectorsOk

`func (o *CreateMatchRuleRequestBodyFields) GetEntitySelectorsOk() (*[]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool)`

GetEntitySelectorsOk returns a tuple with the EntitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectors

`func (o *CreateMatchRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner)`

SetEntitySelectors sets EntitySelectors field to given value.

### HasEntitySelectors

`func (o *CreateMatchRuleRequestBodyFields) HasEntitySelectors() bool`

HasEntitySelectors returns a boolean if a field has been set.

### GetIsPrototype

`func (o *CreateMatchRuleRequestBodyFields) GetIsPrototype() bool`

GetIsPrototype returns the IsPrototype field if non-nil, zero value otherwise.

### GetIsPrototypeOk

`func (o *CreateMatchRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool)`

GetIsPrototypeOk returns a tuple with the IsPrototype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototype

`func (o *CreateMatchRuleRequestBodyFields) SetIsPrototype(v bool)`

SetIsPrototype sets IsPrototype field to given value.

### HasIsPrototype

`func (o *CreateMatchRuleRequestBodyFields) HasIsPrototype() bool`

HasIsPrototype returns a boolean if a field has been set.

### GetName

`func (o *CreateMatchRuleRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMatchRuleRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMatchRuleRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.


### GetParentJaskId

`func (o *CreateMatchRuleRequestBodyFields) GetParentJaskId() string`

GetParentJaskId returns the ParentJaskId field if non-nil, zero value otherwise.

### GetParentJaskIdOk

`func (o *CreateMatchRuleRequestBodyFields) GetParentJaskIdOk() (*string, bool)`

GetParentJaskIdOk returns a tuple with the ParentJaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJaskId

`func (o *CreateMatchRuleRequestBodyFields) SetParentJaskId(v string)`

SetParentJaskId sets ParentJaskId field to given value.

### HasParentJaskId

`func (o *CreateMatchRuleRequestBodyFields) HasParentJaskId() bool`

HasParentJaskId returns a boolean if a field has been set.

### GetSummaryExpression

`func (o *CreateMatchRuleRequestBodyFields) GetSummaryExpression() string`

GetSummaryExpression returns the SummaryExpression field if non-nil, zero value otherwise.

### GetSummaryExpressionOk

`func (o *CreateMatchRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool)`

GetSummaryExpressionOk returns a tuple with the SummaryExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpression

`func (o *CreateMatchRuleRequestBodyFields) SetSummaryExpression(v string)`

SetSummaryExpression sets SummaryExpression field to given value.

### HasSummaryExpression

`func (o *CreateMatchRuleRequestBodyFields) HasSummaryExpression() bool`

HasSummaryExpression returns a boolean if a field has been set.

### GetTags

`func (o *CreateMatchRuleRequestBodyFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateMatchRuleRequestBodyFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateMatchRuleRequestBodyFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateMatchRuleRequestBodyFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTuningExpressionIds

`func (o *CreateMatchRuleRequestBodyFields) GetTuningExpressionIds() []string`

GetTuningExpressionIds returns the TuningExpressionIds field if non-nil, zero value otherwise.

### GetTuningExpressionIdsOk

`func (o *CreateMatchRuleRequestBodyFields) GetTuningExpressionIdsOk() (*[]string, bool)`

GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuningExpressionIds

`func (o *CreateMatchRuleRequestBodyFields) SetTuningExpressionIds(v []string)`

SetTuningExpressionIds sets TuningExpressionIds field to given value.

### HasTuningExpressionIds

`func (o *CreateMatchRuleRequestBodyFields) HasTuningExpressionIds() bool`

HasTuningExpressionIds returns a boolean if a field has been set.

### GetDescription

`func (o *CreateMatchRuleRequestBodyFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMatchRuleRequestBodyFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMatchRuleRequestBodyFields) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExpression

`func (o *CreateMatchRuleRequestBodyFields) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *CreateMatchRuleRequestBodyFields) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *CreateMatchRuleRequestBodyFields) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetScore

`func (o *CreateMatchRuleRequestBodyFields) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CreateMatchRuleRequestBodyFields) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CreateMatchRuleRequestBodyFields) SetScore(v int32)`

SetScore sets Score field to given value.


### GetStream

`func (o *CreateMatchRuleRequestBodyFields) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *CreateMatchRuleRequestBodyFields) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *CreateMatchRuleRequestBodyFields) SetStream(v string)`

SetStream sets Stream field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


