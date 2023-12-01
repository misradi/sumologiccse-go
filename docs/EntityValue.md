# EntityValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ConfigurationType** | **string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Criticality** | Pointer to **string** |  | [optional] 
**Suppressed** | Pointer to **bool** |  | [optional] 
**Created** | **time.Time** |  | 
**CreatedBy** | **string** |  | 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastUpdatedBy** | Pointer to **string** |  | [optional] 
**EntityNamespace** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**Suffix** | Pointer to **string** |  | [optional] 
**NetworkBlock** | Pointer to **string** |  | [optional] 

## Methods

### NewEntityValue

`func NewEntityValue(id string, name string, configurationType string, created time.Time, createdBy string, ) *EntityValue`

NewEntityValue instantiates a new EntityValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityValueWithDefaults

`func NewEntityValueWithDefaults() *EntityValue`

NewEntityValueWithDefaults instantiates a new EntityValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntityValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityValue) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EntityValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityValue) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EntityValue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityValue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityValue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntityValue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfigurationType

`func (o *EntityValue) GetConfigurationType() string`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *EntityValue) GetConfigurationTypeOk() (*string, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *EntityValue) SetConfigurationType(v string)`

SetConfigurationType sets ConfigurationType field to given value.


### GetTags

`func (o *EntityValue) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EntityValue) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EntityValue) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EntityValue) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCriticality

`func (o *EntityValue) GetCriticality() string`

GetCriticality returns the Criticality field if non-nil, zero value otherwise.

### GetCriticalityOk

`func (o *EntityValue) GetCriticalityOk() (*string, bool)`

GetCriticalityOk returns a tuple with the Criticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticality

`func (o *EntityValue) SetCriticality(v string)`

SetCriticality sets Criticality field to given value.

### HasCriticality

`func (o *EntityValue) HasCriticality() bool`

HasCriticality returns a boolean if a field has been set.

### GetSuppressed

`func (o *EntityValue) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *EntityValue) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *EntityValue) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *EntityValue) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetCreated

`func (o *EntityValue) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EntityValue) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EntityValue) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCreatedBy

`func (o *EntityValue) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EntityValue) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EntityValue) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastUpdated

`func (o *EntityValue) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *EntityValue) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *EntityValue) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *EntityValue) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *EntityValue) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *EntityValue) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *EntityValue) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *EntityValue) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetEntityNamespace

`func (o *EntityValue) GetEntityNamespace() string`

GetEntityNamespace returns the EntityNamespace field if non-nil, zero value otherwise.

### GetEntityNamespaceOk

`func (o *EntityValue) GetEntityNamespaceOk() (*string, bool)`

GetEntityNamespaceOk returns a tuple with the EntityNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityNamespace

`func (o *EntityValue) SetEntityNamespace(v string)`

SetEntityNamespace sets EntityNamespace field to given value.

### HasEntityNamespace

`func (o *EntityValue) HasEntityNamespace() bool`

HasEntityNamespace returns a boolean if a field has been set.

### GetEntityType

`func (o *EntityValue) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *EntityValue) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *EntityValue) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *EntityValue) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetPrefix

`func (o *EntityValue) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *EntityValue) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *EntityValue) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *EntityValue) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetSuffix

`func (o *EntityValue) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *EntityValue) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *EntityValue) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *EntityValue) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### GetNetworkBlock

`func (o *EntityValue) GetNetworkBlock() string`

GetNetworkBlock returns the NetworkBlock field if non-nil, zero value otherwise.

### GetNetworkBlockOk

`func (o *EntityValue) GetNetworkBlockOk() (*string, bool)`

GetNetworkBlockOk returns a tuple with the NetworkBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBlock

`func (o *EntityValue) SetNetworkBlock(v string)`

SetNetworkBlock sets NetworkBlock field to given value.

### HasNetworkBlock

`func (o *EntityValue) HasNetworkBlock() bool`

HasNetworkBlock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


