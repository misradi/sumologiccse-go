# CreateEntityValueRequestBodyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ConfigurationType** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**EntityNamespace** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**Suffix** | Pointer to **string** |  | [optional] 
**NetworkBlock** | Pointer to **string** |  | [optional] 
**Criticality** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Suppressed** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateEntityValueRequestBodyFields

`func NewCreateEntityValueRequestBodyFields(name string, configurationType string, ) *CreateEntityValueRequestBodyFields`

NewCreateEntityValueRequestBodyFields instantiates a new CreateEntityValueRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEntityValueRequestBodyFieldsWithDefaults

`func NewCreateEntityValueRequestBodyFieldsWithDefaults() *CreateEntityValueRequestBodyFields`

NewCreateEntityValueRequestBodyFieldsWithDefaults instantiates a new CreateEntityValueRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateEntityValueRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEntityValueRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEntityValueRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.


### GetConfigurationType

`func (o *CreateEntityValueRequestBodyFields) GetConfigurationType() string`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *CreateEntityValueRequestBodyFields) GetConfigurationTypeOk() (*string, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *CreateEntityValueRequestBodyFields) SetConfigurationType(v string)`

SetConfigurationType sets ConfigurationType field to given value.


### GetDescription

`func (o *CreateEntityValueRequestBodyFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateEntityValueRequestBodyFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateEntityValueRequestBodyFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateEntityValueRequestBodyFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntityNamespace

`func (o *CreateEntityValueRequestBodyFields) GetEntityNamespace() string`

GetEntityNamespace returns the EntityNamespace field if non-nil, zero value otherwise.

### GetEntityNamespaceOk

`func (o *CreateEntityValueRequestBodyFields) GetEntityNamespaceOk() (*string, bool)`

GetEntityNamespaceOk returns a tuple with the EntityNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityNamespace

`func (o *CreateEntityValueRequestBodyFields) SetEntityNamespace(v string)`

SetEntityNamespace sets EntityNamespace field to given value.

### HasEntityNamespace

`func (o *CreateEntityValueRequestBodyFields) HasEntityNamespace() bool`

HasEntityNamespace returns a boolean if a field has been set.

### GetEntityType

`func (o *CreateEntityValueRequestBodyFields) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CreateEntityValueRequestBodyFields) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CreateEntityValueRequestBodyFields) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *CreateEntityValueRequestBodyFields) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetPrefix

`func (o *CreateEntityValueRequestBodyFields) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CreateEntityValueRequestBodyFields) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CreateEntityValueRequestBodyFields) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *CreateEntityValueRequestBodyFields) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetSuffix

`func (o *CreateEntityValueRequestBodyFields) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *CreateEntityValueRequestBodyFields) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *CreateEntityValueRequestBodyFields) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *CreateEntityValueRequestBodyFields) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### GetNetworkBlock

`func (o *CreateEntityValueRequestBodyFields) GetNetworkBlock() string`

GetNetworkBlock returns the NetworkBlock field if non-nil, zero value otherwise.

### GetNetworkBlockOk

`func (o *CreateEntityValueRequestBodyFields) GetNetworkBlockOk() (*string, bool)`

GetNetworkBlockOk returns a tuple with the NetworkBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBlock

`func (o *CreateEntityValueRequestBodyFields) SetNetworkBlock(v string)`

SetNetworkBlock sets NetworkBlock field to given value.

### HasNetworkBlock

`func (o *CreateEntityValueRequestBodyFields) HasNetworkBlock() bool`

HasNetworkBlock returns a boolean if a field has been set.

### GetCriticality

`func (o *CreateEntityValueRequestBodyFields) GetCriticality() string`

GetCriticality returns the Criticality field if non-nil, zero value otherwise.

### GetCriticalityOk

`func (o *CreateEntityValueRequestBodyFields) GetCriticalityOk() (*string, bool)`

GetCriticalityOk returns a tuple with the Criticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticality

`func (o *CreateEntityValueRequestBodyFields) SetCriticality(v string)`

SetCriticality sets Criticality field to given value.

### HasCriticality

`func (o *CreateEntityValueRequestBodyFields) HasCriticality() bool`

HasCriticality returns a boolean if a field has been set.

### GetTags

`func (o *CreateEntityValueRequestBodyFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateEntityValueRequestBodyFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateEntityValueRequestBodyFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateEntityValueRequestBodyFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSuppressed

`func (o *CreateEntityValueRequestBodyFields) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *CreateEntityValueRequestBodyFields) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *CreateEntityValueRequestBodyFields) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *CreateEntityValueRequestBodyFields) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


