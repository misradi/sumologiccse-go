# CustomInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Severity** | **string** |  | 
**DynamicSeverity** | Pointer to [**[]CreateCustomInsightRequestBodyFieldsDynamicSeverityInner**](CreateCustomInsightRequestBodyFieldsDynamicSeverityInner.md) |  | [optional] 
**Ordered** | **bool** |  | 
**Enabled** | **bool** |  | 
**Created** | **time.Time** |  | 
**LastUpdated** | **time.Time** |  | 
**RuleIds** | **[]string** |  | 
**SignalNames** | **[]string** |  | 
**Tags** | **[]string** |  | 

## Methods

### NewCustomInsight

`func NewCustomInsight(id string, name string, description string, severity string, ordered bool, enabled bool, created time.Time, lastUpdated time.Time, ruleIds []string, signalNames []string, tags []string, ) *CustomInsight`

NewCustomInsight instantiates a new CustomInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomInsightWithDefaults

`func NewCustomInsightWithDefaults() *CustomInsight`

NewCustomInsightWithDefaults instantiates a new CustomInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomInsight) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomInsight) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomInsight) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CustomInsight) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomInsight) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomInsight) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CustomInsight) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomInsight) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomInsight) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSeverity

`func (o *CustomInsight) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CustomInsight) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CustomInsight) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetDynamicSeverity

`func (o *CustomInsight) GetDynamicSeverity() []CreateCustomInsightRequestBodyFieldsDynamicSeverityInner`

GetDynamicSeverity returns the DynamicSeverity field if non-nil, zero value otherwise.

### GetDynamicSeverityOk

`func (o *CustomInsight) GetDynamicSeverityOk() (*[]CreateCustomInsightRequestBodyFieldsDynamicSeverityInner, bool)`

GetDynamicSeverityOk returns a tuple with the DynamicSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicSeverity

`func (o *CustomInsight) SetDynamicSeverity(v []CreateCustomInsightRequestBodyFieldsDynamicSeverityInner)`

SetDynamicSeverity sets DynamicSeverity field to given value.

### HasDynamicSeverity

`func (o *CustomInsight) HasDynamicSeverity() bool`

HasDynamicSeverity returns a boolean if a field has been set.

### GetOrdered

`func (o *CustomInsight) GetOrdered() bool`

GetOrdered returns the Ordered field if non-nil, zero value otherwise.

### GetOrderedOk

`func (o *CustomInsight) GetOrderedOk() (*bool, bool)`

GetOrderedOk returns a tuple with the Ordered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdered

`func (o *CustomInsight) SetOrdered(v bool)`

SetOrdered sets Ordered field to given value.


### GetEnabled

`func (o *CustomInsight) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustomInsight) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustomInsight) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCreated

`func (o *CustomInsight) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomInsight) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomInsight) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *CustomInsight) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CustomInsight) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CustomInsight) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetRuleIds

`func (o *CustomInsight) GetRuleIds() []string`

GetRuleIds returns the RuleIds field if non-nil, zero value otherwise.

### GetRuleIdsOk

`func (o *CustomInsight) GetRuleIdsOk() (*[]string, bool)`

GetRuleIdsOk returns a tuple with the RuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIds

`func (o *CustomInsight) SetRuleIds(v []string)`

SetRuleIds sets RuleIds field to given value.


### GetSignalNames

`func (o *CustomInsight) GetSignalNames() []string`

GetSignalNames returns the SignalNames field if non-nil, zero value otherwise.

### GetSignalNamesOk

`func (o *CustomInsight) GetSignalNamesOk() (*[]string, bool)`

GetSignalNamesOk returns a tuple with the SignalNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalNames

`func (o *CustomInsight) SetSignalNames(v []string)`

SetSignalNames sets SignalNames field to given value.


### GetTags

`func (o *CustomInsight) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CustomInsight) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CustomInsight) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


