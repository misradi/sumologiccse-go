# CreateEntityCriticalityConfigRequestBodyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human friendly and unique name. Examples: \&quot;Executive Laptop\&quot;, \&quot;Bastion Host\&quot; | 
**SeverityExpression** | **string** | Algebraic expression representing this entity&#39;s criticality. Examples: \&quot;severity * 2\&quot;, \&quot;severity - 5\&quot;, \&quot;severity / 3\&quot; | 

## Methods

### NewCreateEntityCriticalityConfigRequestBodyFields

`func NewCreateEntityCriticalityConfigRequestBodyFields(name string, severityExpression string, ) *CreateEntityCriticalityConfigRequestBodyFields`

NewCreateEntityCriticalityConfigRequestBodyFields instantiates a new CreateEntityCriticalityConfigRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEntityCriticalityConfigRequestBodyFieldsWithDefaults

`func NewCreateEntityCriticalityConfigRequestBodyFieldsWithDefaults() *CreateEntityCriticalityConfigRequestBodyFields`

NewCreateEntityCriticalityConfigRequestBodyFieldsWithDefaults instantiates a new CreateEntityCriticalityConfigRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateEntityCriticalityConfigRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEntityCriticalityConfigRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEntityCriticalityConfigRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.


### GetSeverityExpression

`func (o *CreateEntityCriticalityConfigRequestBodyFields) GetSeverityExpression() string`

GetSeverityExpression returns the SeverityExpression field if non-nil, zero value otherwise.

### GetSeverityExpressionOk

`func (o *CreateEntityCriticalityConfigRequestBodyFields) GetSeverityExpressionOk() (*string, bool)`

GetSeverityExpressionOk returns a tuple with the SeverityExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityExpression

`func (o *CreateEntityCriticalityConfigRequestBodyFields) SetSeverityExpression(v string)`

SetSeverityExpression sets SeverityExpression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


