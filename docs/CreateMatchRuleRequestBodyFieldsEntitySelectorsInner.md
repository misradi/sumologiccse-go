# CreateMatchRuleRequestBodyFieldsEntitySelectorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | **string** | The field from the record containing the Entity value to generate the Signal on | 
**EntityType** | **string** | The identifier of the type of the Entity to generate the Signal on. Either one of the built-in entity types (_ip, _hostname, _mac, _username) or a custom entity type defined in the UI. | 

## Methods

### NewCreateMatchRuleRequestBodyFieldsEntitySelectorsInner

`func NewCreateMatchRuleRequestBodyFieldsEntitySelectorsInner(expression string, entityType string, ) *CreateMatchRuleRequestBodyFieldsEntitySelectorsInner`

NewCreateMatchRuleRequestBodyFieldsEntitySelectorsInner instantiates a new CreateMatchRuleRequestBodyFieldsEntitySelectorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMatchRuleRequestBodyFieldsEntitySelectorsInnerWithDefaults

`func NewCreateMatchRuleRequestBodyFieldsEntitySelectorsInnerWithDefaults() *CreateMatchRuleRequestBodyFieldsEntitySelectorsInner`

NewCreateMatchRuleRequestBodyFieldsEntitySelectorsInnerWithDefaults instantiates a new CreateMatchRuleRequestBodyFieldsEntitySelectorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *CreateMatchRuleRequestBodyFieldsEntitySelectorsInner) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *CreateMatchRuleRequestBodyFieldsEntitySelectorsInner) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *CreateMatchRuleRequestBodyFieldsEntitySelectorsInner) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetEntityType

`func (o *CreateMatchRuleRequestBodyFieldsEntitySelectorsInner) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CreateMatchRuleRequestBodyFieldsEntitySelectorsInner) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CreateMatchRuleRequestBodyFieldsEntitySelectorsInner) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


