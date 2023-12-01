# CreateCustomInsightRequestBodyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Severity** | **string** |  | 
**Ordered** | **bool** |  | 
**Enabled** | **bool** |  | 
**Tags** | **[]string** |  | 
**DynamicSeverity** | Pointer to [**[]CreateCustomInsightRequestBodyFieldsDynamicSeverityInner**](CreateCustomInsightRequestBodyFieldsDynamicSeverityInner.md) |  | [optional] 
**RuleIds** | Pointer to **[]string** |  | [optional] 
**SignalNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateCustomInsightRequestBodyFields

`func NewCreateCustomInsightRequestBodyFields(name string, description string, severity string, ordered bool, enabled bool, tags []string, ) *CreateCustomInsightRequestBodyFields`

NewCreateCustomInsightRequestBodyFields instantiates a new CreateCustomInsightRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomInsightRequestBodyFieldsWithDefaults

`func NewCreateCustomInsightRequestBodyFieldsWithDefaults() *CreateCustomInsightRequestBodyFields`

NewCreateCustomInsightRequestBodyFieldsWithDefaults instantiates a new CreateCustomInsightRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCustomInsightRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCustomInsightRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCustomInsightRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateCustomInsightRequestBodyFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCustomInsightRequestBodyFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCustomInsightRequestBodyFields) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSeverity

`func (o *CreateCustomInsightRequestBodyFields) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CreateCustomInsightRequestBodyFields) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CreateCustomInsightRequestBodyFields) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetOrdered

`func (o *CreateCustomInsightRequestBodyFields) GetOrdered() bool`

GetOrdered returns the Ordered field if non-nil, zero value otherwise.

### GetOrderedOk

`func (o *CreateCustomInsightRequestBodyFields) GetOrderedOk() (*bool, bool)`

GetOrderedOk returns a tuple with the Ordered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdered

`func (o *CreateCustomInsightRequestBodyFields) SetOrdered(v bool)`

SetOrdered sets Ordered field to given value.


### GetEnabled

`func (o *CreateCustomInsightRequestBodyFields) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateCustomInsightRequestBodyFields) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateCustomInsightRequestBodyFields) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetTags

`func (o *CreateCustomInsightRequestBodyFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateCustomInsightRequestBodyFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateCustomInsightRequestBodyFields) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetDynamicSeverity

`func (o *CreateCustomInsightRequestBodyFields) GetDynamicSeverity() []CreateCustomInsightRequestBodyFieldsDynamicSeverityInner`

GetDynamicSeverity returns the DynamicSeverity field if non-nil, zero value otherwise.

### GetDynamicSeverityOk

`func (o *CreateCustomInsightRequestBodyFields) GetDynamicSeverityOk() (*[]CreateCustomInsightRequestBodyFieldsDynamicSeverityInner, bool)`

GetDynamicSeverityOk returns a tuple with the DynamicSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicSeverity

`func (o *CreateCustomInsightRequestBodyFields) SetDynamicSeverity(v []CreateCustomInsightRequestBodyFieldsDynamicSeverityInner)`

SetDynamicSeverity sets DynamicSeverity field to given value.

### HasDynamicSeverity

`func (o *CreateCustomInsightRequestBodyFields) HasDynamicSeverity() bool`

HasDynamicSeverity returns a boolean if a field has been set.

### GetRuleIds

`func (o *CreateCustomInsightRequestBodyFields) GetRuleIds() []string`

GetRuleIds returns the RuleIds field if non-nil, zero value otherwise.

### GetRuleIdsOk

`func (o *CreateCustomInsightRequestBodyFields) GetRuleIdsOk() (*[]string, bool)`

GetRuleIdsOk returns a tuple with the RuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIds

`func (o *CreateCustomInsightRequestBodyFields) SetRuleIds(v []string)`

SetRuleIds sets RuleIds field to given value.

### HasRuleIds

`func (o *CreateCustomInsightRequestBodyFields) HasRuleIds() bool`

HasRuleIds returns a boolean if a field has been set.

### GetSignalNames

`func (o *CreateCustomInsightRequestBodyFields) GetSignalNames() []string`

GetSignalNames returns the SignalNames field if non-nil, zero value otherwise.

### GetSignalNamesOk

`func (o *CreateCustomInsightRequestBodyFields) GetSignalNamesOk() (*[]string, bool)`

GetSignalNamesOk returns a tuple with the SignalNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalNames

`func (o *CreateCustomInsightRequestBodyFields) SetSignalNames(v []string)`

SetSignalNames sets SignalNames field to given value.

### HasSignalNames

`func (o *CreateCustomInsightRequestBodyFields) HasSignalNames() bool`

HasSignalNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


