# UpdateEntityCriticalityConfigRequestBodyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeverityExpression** | **string** | Algebraic expression representing this entity&#39;s criticality. Examples: \&quot;severity * 2\&quot;, \&quot;severity - 5\&quot;, \&quot;severity / 3\&quot; | 

## Methods

### NewUpdateEntityCriticalityConfigRequestBodyFields

`func NewUpdateEntityCriticalityConfigRequestBodyFields(severityExpression string, ) *UpdateEntityCriticalityConfigRequestBodyFields`

NewUpdateEntityCriticalityConfigRequestBodyFields instantiates a new UpdateEntityCriticalityConfigRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEntityCriticalityConfigRequestBodyFieldsWithDefaults

`func NewUpdateEntityCriticalityConfigRequestBodyFieldsWithDefaults() *UpdateEntityCriticalityConfigRequestBodyFields`

NewUpdateEntityCriticalityConfigRequestBodyFieldsWithDefaults instantiates a new UpdateEntityCriticalityConfigRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeverityExpression

`func (o *UpdateEntityCriticalityConfigRequestBodyFields) GetSeverityExpression() string`

GetSeverityExpression returns the SeverityExpression field if non-nil, zero value otherwise.

### GetSeverityExpressionOk

`func (o *UpdateEntityCriticalityConfigRequestBodyFields) GetSeverityExpressionOk() (*string, bool)`

GetSeverityExpressionOk returns a tuple with the SeverityExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityExpression

`func (o *UpdateEntityCriticalityConfigRequestBodyFields) SetSeverityExpression(v string)`

SetSeverityExpression sets SeverityExpression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


