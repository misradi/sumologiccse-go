# GetEntityDomainConfigurationResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NormalizeUsernames** | **bool** |  | 
**NormalizeHostnames** | **bool** |  | 
**DefaultNormalizedDomain** | Pointer to **string** |  | [optional] 
**DomainMappings** | [**[]GetEntityDomainConfigurationResponseDataDomainMappingsInner**](GetEntityDomainConfigurationResponseDataDomainMappingsInner.md) |  | 
**AdDomainNormalizationEnabled** | **bool** |  | 
**FqdnNormalizationEnabled** | **bool** |  | 
**AwsNormalizationEnabled** | **bool** |  | 

## Methods

### NewGetEntityDomainConfigurationResponseData

`func NewGetEntityDomainConfigurationResponseData(normalizeUsernames bool, normalizeHostnames bool, domainMappings []GetEntityDomainConfigurationResponseDataDomainMappingsInner, adDomainNormalizationEnabled bool, fqdnNormalizationEnabled bool, awsNormalizationEnabled bool, ) *GetEntityDomainConfigurationResponseData`

NewGetEntityDomainConfigurationResponseData instantiates a new GetEntityDomainConfigurationResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntityDomainConfigurationResponseDataWithDefaults

`func NewGetEntityDomainConfigurationResponseDataWithDefaults() *GetEntityDomainConfigurationResponseData`

NewGetEntityDomainConfigurationResponseDataWithDefaults instantiates a new GetEntityDomainConfigurationResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNormalizeUsernames

`func (o *GetEntityDomainConfigurationResponseData) GetNormalizeUsernames() bool`

GetNormalizeUsernames returns the NormalizeUsernames field if non-nil, zero value otherwise.

### GetNormalizeUsernamesOk

`func (o *GetEntityDomainConfigurationResponseData) GetNormalizeUsernamesOk() (*bool, bool)`

GetNormalizeUsernamesOk returns a tuple with the NormalizeUsernames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizeUsernames

`func (o *GetEntityDomainConfigurationResponseData) SetNormalizeUsernames(v bool)`

SetNormalizeUsernames sets NormalizeUsernames field to given value.


### GetNormalizeHostnames

`func (o *GetEntityDomainConfigurationResponseData) GetNormalizeHostnames() bool`

GetNormalizeHostnames returns the NormalizeHostnames field if non-nil, zero value otherwise.

### GetNormalizeHostnamesOk

`func (o *GetEntityDomainConfigurationResponseData) GetNormalizeHostnamesOk() (*bool, bool)`

GetNormalizeHostnamesOk returns a tuple with the NormalizeHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizeHostnames

`func (o *GetEntityDomainConfigurationResponseData) SetNormalizeHostnames(v bool)`

SetNormalizeHostnames sets NormalizeHostnames field to given value.


### GetDefaultNormalizedDomain

`func (o *GetEntityDomainConfigurationResponseData) GetDefaultNormalizedDomain() string`

GetDefaultNormalizedDomain returns the DefaultNormalizedDomain field if non-nil, zero value otherwise.

### GetDefaultNormalizedDomainOk

`func (o *GetEntityDomainConfigurationResponseData) GetDefaultNormalizedDomainOk() (*string, bool)`

GetDefaultNormalizedDomainOk returns a tuple with the DefaultNormalizedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNormalizedDomain

`func (o *GetEntityDomainConfigurationResponseData) SetDefaultNormalizedDomain(v string)`

SetDefaultNormalizedDomain sets DefaultNormalizedDomain field to given value.

### HasDefaultNormalizedDomain

`func (o *GetEntityDomainConfigurationResponseData) HasDefaultNormalizedDomain() bool`

HasDefaultNormalizedDomain returns a boolean if a field has been set.

### GetDomainMappings

`func (o *GetEntityDomainConfigurationResponseData) GetDomainMappings() []GetEntityDomainConfigurationResponseDataDomainMappingsInner`

GetDomainMappings returns the DomainMappings field if non-nil, zero value otherwise.

### GetDomainMappingsOk

`func (o *GetEntityDomainConfigurationResponseData) GetDomainMappingsOk() (*[]GetEntityDomainConfigurationResponseDataDomainMappingsInner, bool)`

GetDomainMappingsOk returns a tuple with the DomainMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainMappings

`func (o *GetEntityDomainConfigurationResponseData) SetDomainMappings(v []GetEntityDomainConfigurationResponseDataDomainMappingsInner)`

SetDomainMappings sets DomainMappings field to given value.


### GetAdDomainNormalizationEnabled

`func (o *GetEntityDomainConfigurationResponseData) GetAdDomainNormalizationEnabled() bool`

GetAdDomainNormalizationEnabled returns the AdDomainNormalizationEnabled field if non-nil, zero value otherwise.

### GetAdDomainNormalizationEnabledOk

`func (o *GetEntityDomainConfigurationResponseData) GetAdDomainNormalizationEnabledOk() (*bool, bool)`

GetAdDomainNormalizationEnabledOk returns a tuple with the AdDomainNormalizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainNormalizationEnabled

`func (o *GetEntityDomainConfigurationResponseData) SetAdDomainNormalizationEnabled(v bool)`

SetAdDomainNormalizationEnabled sets AdDomainNormalizationEnabled field to given value.


### GetFqdnNormalizationEnabled

`func (o *GetEntityDomainConfigurationResponseData) GetFqdnNormalizationEnabled() bool`

GetFqdnNormalizationEnabled returns the FqdnNormalizationEnabled field if non-nil, zero value otherwise.

### GetFqdnNormalizationEnabledOk

`func (o *GetEntityDomainConfigurationResponseData) GetFqdnNormalizationEnabledOk() (*bool, bool)`

GetFqdnNormalizationEnabledOk returns a tuple with the FqdnNormalizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdnNormalizationEnabled

`func (o *GetEntityDomainConfigurationResponseData) SetFqdnNormalizationEnabled(v bool)`

SetFqdnNormalizationEnabled sets FqdnNormalizationEnabled field to given value.


### GetAwsNormalizationEnabled

`func (o *GetEntityDomainConfigurationResponseData) GetAwsNormalizationEnabled() bool`

GetAwsNormalizationEnabled returns the AwsNormalizationEnabled field if non-nil, zero value otherwise.

### GetAwsNormalizationEnabledOk

`func (o *GetEntityDomainConfigurationResponseData) GetAwsNormalizationEnabledOk() (*bool, bool)`

GetAwsNormalizationEnabledOk returns a tuple with the AwsNormalizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsNormalizationEnabled

`func (o *GetEntityDomainConfigurationResponseData) SetAwsNormalizationEnabled(v bool)`

SetAwsNormalizationEnabled sets AwsNormalizationEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


