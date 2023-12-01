# Automation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**PlaybookId** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**CseResourceType** | **string** |  | 
**CseResourceSubTypes** | Pointer to **[]string** |  | [optional] 
**ExecutionTypes** | **[]string** |  | 
**Enabled** | **bool** |  | 

## Methods

### NewAutomation

`func NewAutomation(id string, playbookId string, name string, cseResourceType string, executionTypes []string, enabled bool, ) *Automation`

NewAutomation instantiates a new Automation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationWithDefaults

`func NewAutomationWithDefaults() *Automation`

NewAutomationWithDefaults instantiates a new Automation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Automation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Automation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Automation) SetId(v string)`

SetId sets Id field to given value.


### GetPlaybookId

`func (o *Automation) GetPlaybookId() string`

GetPlaybookId returns the PlaybookId field if non-nil, zero value otherwise.

### GetPlaybookIdOk

`func (o *Automation) GetPlaybookIdOk() (*string, bool)`

GetPlaybookIdOk returns a tuple with the PlaybookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookId

`func (o *Automation) SetPlaybookId(v string)`

SetPlaybookId sets PlaybookId field to given value.


### GetName

`func (o *Automation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Automation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Automation) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Automation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Automation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Automation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Automation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCseResourceType

`func (o *Automation) GetCseResourceType() string`

GetCseResourceType returns the CseResourceType field if non-nil, zero value otherwise.

### GetCseResourceTypeOk

`func (o *Automation) GetCseResourceTypeOk() (*string, bool)`

GetCseResourceTypeOk returns a tuple with the CseResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCseResourceType

`func (o *Automation) SetCseResourceType(v string)`

SetCseResourceType sets CseResourceType field to given value.


### GetCseResourceSubTypes

`func (o *Automation) GetCseResourceSubTypes() []string`

GetCseResourceSubTypes returns the CseResourceSubTypes field if non-nil, zero value otherwise.

### GetCseResourceSubTypesOk

`func (o *Automation) GetCseResourceSubTypesOk() (*[]string, bool)`

GetCseResourceSubTypesOk returns a tuple with the CseResourceSubTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCseResourceSubTypes

`func (o *Automation) SetCseResourceSubTypes(v []string)`

SetCseResourceSubTypes sets CseResourceSubTypes field to given value.

### HasCseResourceSubTypes

`func (o *Automation) HasCseResourceSubTypes() bool`

HasCseResourceSubTypes returns a boolean if a field has been set.

### GetExecutionTypes

`func (o *Automation) GetExecutionTypes() []string`

GetExecutionTypes returns the ExecutionTypes field if non-nil, zero value otherwise.

### GetExecutionTypesOk

`func (o *Automation) GetExecutionTypesOk() (*[]string, bool)`

GetExecutionTypesOk returns a tuple with the ExecutionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTypes

`func (o *Automation) SetExecutionTypes(v []string)`

SetExecutionTypes sets ExecutionTypes field to given value.


### GetEnabled

`func (o *Automation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Automation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Automation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


