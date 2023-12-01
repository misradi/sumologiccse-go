# CreateAutomationRequestBodyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaybookId** | **string** |  | 
**CseResourceType** | **string** |  | 
**CseResourceSubTypes** | Pointer to **[]string** |  | [optional] 
**ExecutionTypes** | **[]string** |  | 
**Enabled** | **bool** |  | 

## Methods

### NewCreateAutomationRequestBodyFields

`func NewCreateAutomationRequestBodyFields(playbookId string, cseResourceType string, executionTypes []string, enabled bool, ) *CreateAutomationRequestBodyFields`

NewCreateAutomationRequestBodyFields instantiates a new CreateAutomationRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutomationRequestBodyFieldsWithDefaults

`func NewCreateAutomationRequestBodyFieldsWithDefaults() *CreateAutomationRequestBodyFields`

NewCreateAutomationRequestBodyFieldsWithDefaults instantiates a new CreateAutomationRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaybookId

`func (o *CreateAutomationRequestBodyFields) GetPlaybookId() string`

GetPlaybookId returns the PlaybookId field if non-nil, zero value otherwise.

### GetPlaybookIdOk

`func (o *CreateAutomationRequestBodyFields) GetPlaybookIdOk() (*string, bool)`

GetPlaybookIdOk returns a tuple with the PlaybookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookId

`func (o *CreateAutomationRequestBodyFields) SetPlaybookId(v string)`

SetPlaybookId sets PlaybookId field to given value.


### GetCseResourceType

`func (o *CreateAutomationRequestBodyFields) GetCseResourceType() string`

GetCseResourceType returns the CseResourceType field if non-nil, zero value otherwise.

### GetCseResourceTypeOk

`func (o *CreateAutomationRequestBodyFields) GetCseResourceTypeOk() (*string, bool)`

GetCseResourceTypeOk returns a tuple with the CseResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCseResourceType

`func (o *CreateAutomationRequestBodyFields) SetCseResourceType(v string)`

SetCseResourceType sets CseResourceType field to given value.


### GetCseResourceSubTypes

`func (o *CreateAutomationRequestBodyFields) GetCseResourceSubTypes() []string`

GetCseResourceSubTypes returns the CseResourceSubTypes field if non-nil, zero value otherwise.

### GetCseResourceSubTypesOk

`func (o *CreateAutomationRequestBodyFields) GetCseResourceSubTypesOk() (*[]string, bool)`

GetCseResourceSubTypesOk returns a tuple with the CseResourceSubTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCseResourceSubTypes

`func (o *CreateAutomationRequestBodyFields) SetCseResourceSubTypes(v []string)`

SetCseResourceSubTypes sets CseResourceSubTypes field to given value.

### HasCseResourceSubTypes

`func (o *CreateAutomationRequestBodyFields) HasCseResourceSubTypes() bool`

HasCseResourceSubTypes returns a boolean if a field has been set.

### GetExecutionTypes

`func (o *CreateAutomationRequestBodyFields) GetExecutionTypes() []string`

GetExecutionTypes returns the ExecutionTypes field if non-nil, zero value otherwise.

### GetExecutionTypesOk

`func (o *CreateAutomationRequestBodyFields) GetExecutionTypesOk() (*[]string, bool)`

GetExecutionTypesOk returns a tuple with the ExecutionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTypes

`func (o *CreateAutomationRequestBodyFields) SetExecutionTypes(v []string)`

SetExecutionTypes sets ExecutionTypes field to given value.


### GetEnabled

`func (o *CreateAutomationRequestBodyFields) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateAutomationRequestBodyFields) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateAutomationRequestBodyFields) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


