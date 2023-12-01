# CreateThresholdRuleRequestBodyFields

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
**CountDistinct** | **bool** |  | 
**CountField** | Pointer to **string** |  | [optional] 
**Expression** | **string** |  | 
**Limit** | **int32** |  | 
**Score** | **int32** |  | 
**Stream** | **string** |  | 
**Version** | **int32** |  | 
**WindowSize** | **string** |  | 
**WindowSizeMilliseconds** | Pointer to **string** | Can be used instead of windowSize. | [optional] 
**GroupByFields** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateThresholdRuleRequestBodyFields

`func NewCreateThresholdRuleRequestBodyFields(enabled bool, name string, description string, countDistinct bool, expression string, limit int32, score int32, stream string, version int32, windowSize string, ) *CreateThresholdRuleRequestBodyFields`

NewCreateThresholdRuleRequestBodyFields instantiates a new CreateThresholdRuleRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateThresholdRuleRequestBodyFieldsWithDefaults

`func NewCreateThresholdRuleRequestBodyFieldsWithDefaults() *CreateThresholdRuleRequestBodyFields`

NewCreateThresholdRuleRequestBodyFieldsWithDefaults instantiates a new CreateThresholdRuleRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetField

`func (o *CreateThresholdRuleRequestBodyFields) GetAssetField() string`

GetAssetField returns the AssetField field if non-nil, zero value otherwise.

### GetAssetFieldOk

`func (o *CreateThresholdRuleRequestBodyFields) GetAssetFieldOk() (*string, bool)`

GetAssetFieldOk returns a tuple with the AssetField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetField

`func (o *CreateThresholdRuleRequestBodyFields) SetAssetField(v string)`

SetAssetField sets AssetField field to given value.

### HasAssetField

`func (o *CreateThresholdRuleRequestBodyFields) HasAssetField() bool`

HasAssetField returns a boolean if a field has been set.

### GetCategory

`func (o *CreateThresholdRuleRequestBodyFields) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateThresholdRuleRequestBodyFields) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateThresholdRuleRequestBodyFields) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CreateThresholdRuleRequestBodyFields) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateThresholdRuleRequestBodyFields) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateThresholdRuleRequestBodyFields) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateThresholdRuleRequestBodyFields) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEntitySelectors

`func (o *CreateThresholdRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner`

GetEntitySelectors returns the EntitySelectors field if non-nil, zero value otherwise.

### GetEntitySelectorsOk

`func (o *CreateThresholdRuleRequestBodyFields) GetEntitySelectorsOk() (*[]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool)`

GetEntitySelectorsOk returns a tuple with the EntitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectors

`func (o *CreateThresholdRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner)`

SetEntitySelectors sets EntitySelectors field to given value.

### HasEntitySelectors

`func (o *CreateThresholdRuleRequestBodyFields) HasEntitySelectors() bool`

HasEntitySelectors returns a boolean if a field has been set.

### GetIsPrototype

`func (o *CreateThresholdRuleRequestBodyFields) GetIsPrototype() bool`

GetIsPrototype returns the IsPrototype field if non-nil, zero value otherwise.

### GetIsPrototypeOk

`func (o *CreateThresholdRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool)`

GetIsPrototypeOk returns a tuple with the IsPrototype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototype

`func (o *CreateThresholdRuleRequestBodyFields) SetIsPrototype(v bool)`

SetIsPrototype sets IsPrototype field to given value.

### HasIsPrototype

`func (o *CreateThresholdRuleRequestBodyFields) HasIsPrototype() bool`

HasIsPrototype returns a boolean if a field has been set.

### GetName

`func (o *CreateThresholdRuleRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateThresholdRuleRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateThresholdRuleRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.


### GetParentJaskId

`func (o *CreateThresholdRuleRequestBodyFields) GetParentJaskId() string`

GetParentJaskId returns the ParentJaskId field if non-nil, zero value otherwise.

### GetParentJaskIdOk

`func (o *CreateThresholdRuleRequestBodyFields) GetParentJaskIdOk() (*string, bool)`

GetParentJaskIdOk returns a tuple with the ParentJaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJaskId

`func (o *CreateThresholdRuleRequestBodyFields) SetParentJaskId(v string)`

SetParentJaskId sets ParentJaskId field to given value.

### HasParentJaskId

`func (o *CreateThresholdRuleRequestBodyFields) HasParentJaskId() bool`

HasParentJaskId returns a boolean if a field has been set.

### GetSummaryExpression

`func (o *CreateThresholdRuleRequestBodyFields) GetSummaryExpression() string`

GetSummaryExpression returns the SummaryExpression field if non-nil, zero value otherwise.

### GetSummaryExpressionOk

`func (o *CreateThresholdRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool)`

GetSummaryExpressionOk returns a tuple with the SummaryExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpression

`func (o *CreateThresholdRuleRequestBodyFields) SetSummaryExpression(v string)`

SetSummaryExpression sets SummaryExpression field to given value.

### HasSummaryExpression

`func (o *CreateThresholdRuleRequestBodyFields) HasSummaryExpression() bool`

HasSummaryExpression returns a boolean if a field has been set.

### GetTags

`func (o *CreateThresholdRuleRequestBodyFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateThresholdRuleRequestBodyFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateThresholdRuleRequestBodyFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateThresholdRuleRequestBodyFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTuningExpressionIds

`func (o *CreateThresholdRuleRequestBodyFields) GetTuningExpressionIds() []string`

GetTuningExpressionIds returns the TuningExpressionIds field if non-nil, zero value otherwise.

### GetTuningExpressionIdsOk

`func (o *CreateThresholdRuleRequestBodyFields) GetTuningExpressionIdsOk() (*[]string, bool)`

GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuningExpressionIds

`func (o *CreateThresholdRuleRequestBodyFields) SetTuningExpressionIds(v []string)`

SetTuningExpressionIds sets TuningExpressionIds field to given value.

### HasTuningExpressionIds

`func (o *CreateThresholdRuleRequestBodyFields) HasTuningExpressionIds() bool`

HasTuningExpressionIds returns a boolean if a field has been set.

### GetDescription

`func (o *CreateThresholdRuleRequestBodyFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateThresholdRuleRequestBodyFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateThresholdRuleRequestBodyFields) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCountDistinct

`func (o *CreateThresholdRuleRequestBodyFields) GetCountDistinct() bool`

GetCountDistinct returns the CountDistinct field if non-nil, zero value otherwise.

### GetCountDistinctOk

`func (o *CreateThresholdRuleRequestBodyFields) GetCountDistinctOk() (*bool, bool)`

GetCountDistinctOk returns a tuple with the CountDistinct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountDistinct

`func (o *CreateThresholdRuleRequestBodyFields) SetCountDistinct(v bool)`

SetCountDistinct sets CountDistinct field to given value.


### GetCountField

`func (o *CreateThresholdRuleRequestBodyFields) GetCountField() string`

GetCountField returns the CountField field if non-nil, zero value otherwise.

### GetCountFieldOk

`func (o *CreateThresholdRuleRequestBodyFields) GetCountFieldOk() (*string, bool)`

GetCountFieldOk returns a tuple with the CountField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountField

`func (o *CreateThresholdRuleRequestBodyFields) SetCountField(v string)`

SetCountField sets CountField field to given value.

### HasCountField

`func (o *CreateThresholdRuleRequestBodyFields) HasCountField() bool`

HasCountField returns a boolean if a field has been set.

### GetExpression

`func (o *CreateThresholdRuleRequestBodyFields) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *CreateThresholdRuleRequestBodyFields) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *CreateThresholdRuleRequestBodyFields) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetLimit

`func (o *CreateThresholdRuleRequestBodyFields) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CreateThresholdRuleRequestBodyFields) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CreateThresholdRuleRequestBodyFields) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetScore

`func (o *CreateThresholdRuleRequestBodyFields) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CreateThresholdRuleRequestBodyFields) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CreateThresholdRuleRequestBodyFields) SetScore(v int32)`

SetScore sets Score field to given value.


### GetStream

`func (o *CreateThresholdRuleRequestBodyFields) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *CreateThresholdRuleRequestBodyFields) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *CreateThresholdRuleRequestBodyFields) SetStream(v string)`

SetStream sets Stream field to given value.


### GetVersion

`func (o *CreateThresholdRuleRequestBodyFields) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateThresholdRuleRequestBodyFields) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateThresholdRuleRequestBodyFields) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetWindowSize

`func (o *CreateThresholdRuleRequestBodyFields) GetWindowSize() string`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *CreateThresholdRuleRequestBodyFields) GetWindowSizeOk() (*string, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *CreateThresholdRuleRequestBodyFields) SetWindowSize(v string)`

SetWindowSize sets WindowSize field to given value.


### GetWindowSizeMilliseconds

`func (o *CreateThresholdRuleRequestBodyFields) GetWindowSizeMilliseconds() string`

GetWindowSizeMilliseconds returns the WindowSizeMilliseconds field if non-nil, zero value otherwise.

### GetWindowSizeMillisecondsOk

`func (o *CreateThresholdRuleRequestBodyFields) GetWindowSizeMillisecondsOk() (*string, bool)`

GetWindowSizeMillisecondsOk returns a tuple with the WindowSizeMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSizeMilliseconds

`func (o *CreateThresholdRuleRequestBodyFields) SetWindowSizeMilliseconds(v string)`

SetWindowSizeMilliseconds sets WindowSizeMilliseconds field to given value.

### HasWindowSizeMilliseconds

`func (o *CreateThresholdRuleRequestBodyFields) HasWindowSizeMilliseconds() bool`

HasWindowSizeMilliseconds returns a boolean if a field has been set.

### GetGroupByFields

`func (o *CreateThresholdRuleRequestBodyFields) GetGroupByFields() []string`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *CreateThresholdRuleRequestBodyFields) GetGroupByFieldsOk() (*[]string, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *CreateThresholdRuleRequestBodyFields) SetGroupByFields(v []string)`

SetGroupByFields sets GroupByFields field to given value.

### HasGroupByFields

`func (o *CreateThresholdRuleRequestBodyFields) HasGroupByFields() bool`

HasGroupByFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


