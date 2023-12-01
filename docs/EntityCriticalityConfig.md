# EntityCriticalityConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | Human friendly and unique name. Examples: \&quot;Executive Laptop\&quot;, \&quot;Bastion Host\&quot; | 
**SeverityExpression** | **string** | Algebraic expression representing this entity&#39;s criticality. Examples: \&quot;severity * 2\&quot;, \&quot;severity - 5\&quot;, \&quot;severity / 3\&quot; | 
**EntityCount** | **int32** | Number of entities related to this criticality. | 

## Methods

### NewEntityCriticalityConfig

`func NewEntityCriticalityConfig(id string, name string, severityExpression string, entityCount int32, ) *EntityCriticalityConfig`

NewEntityCriticalityConfig instantiates a new EntityCriticalityConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityCriticalityConfigWithDefaults

`func NewEntityCriticalityConfigWithDefaults() *EntityCriticalityConfig`

NewEntityCriticalityConfigWithDefaults instantiates a new EntityCriticalityConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntityCriticalityConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityCriticalityConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityCriticalityConfig) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EntityCriticalityConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityCriticalityConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityCriticalityConfig) SetName(v string)`

SetName sets Name field to given value.


### GetSeverityExpression

`func (o *EntityCriticalityConfig) GetSeverityExpression() string`

GetSeverityExpression returns the SeverityExpression field if non-nil, zero value otherwise.

### GetSeverityExpressionOk

`func (o *EntityCriticalityConfig) GetSeverityExpressionOk() (*string, bool)`

GetSeverityExpressionOk returns a tuple with the SeverityExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityExpression

`func (o *EntityCriticalityConfig) SetSeverityExpression(v string)`

SetSeverityExpression sets SeverityExpression field to given value.


### GetEntityCount

`func (o *EntityCriticalityConfig) GetEntityCount() int32`

GetEntityCount returns the EntityCount field if non-nil, zero value otherwise.

### GetEntityCountOk

`func (o *EntityCriticalityConfig) GetEntityCountOk() (*int32, bool)`

GetEntityCountOk returns a tuple with the EntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityCount

`func (o *EntityCriticalityConfig) SetEntityCount(v int32)`

SetEntityCount sets EntityCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


