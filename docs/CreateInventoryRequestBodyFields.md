# CreateInventoryRequestBodyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ConfigurationType** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**InventoryType** | Pointer to **string** |  | [optional] 
**InventorySource** | **string** |  | 
**Group** | Pointer to **string** |  | [optional] 
**InventoryKey** | Pointer to **string** |  | [optional] 
**InventoryValue** | Pointer to **string** |  | [optional] 
**TagSchema** | Pointer to **string** |  | [optional] 
**DynamicTags** | Pointer to **bool** |  | [optional] 
**Criticality** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Suppressed** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateInventoryRequestBodyFields

`func NewCreateInventoryRequestBodyFields(name string, configurationType string, inventorySource string, ) *CreateInventoryRequestBodyFields`

NewCreateInventoryRequestBodyFields instantiates a new CreateInventoryRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInventoryRequestBodyFieldsWithDefaults

`func NewCreateInventoryRequestBodyFieldsWithDefaults() *CreateInventoryRequestBodyFields`

NewCreateInventoryRequestBodyFieldsWithDefaults instantiates a new CreateInventoryRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateInventoryRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInventoryRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInventoryRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.


### GetConfigurationType

`func (o *CreateInventoryRequestBodyFields) GetConfigurationType() string`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *CreateInventoryRequestBodyFields) GetConfigurationTypeOk() (*string, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *CreateInventoryRequestBodyFields) SetConfigurationType(v string)`

SetConfigurationType sets ConfigurationType field to given value.


### GetDescription

`func (o *CreateInventoryRequestBodyFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateInventoryRequestBodyFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateInventoryRequestBodyFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateInventoryRequestBodyFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInventoryType

`func (o *CreateInventoryRequestBodyFields) GetInventoryType() string`

GetInventoryType returns the InventoryType field if non-nil, zero value otherwise.

### GetInventoryTypeOk

`func (o *CreateInventoryRequestBodyFields) GetInventoryTypeOk() (*string, bool)`

GetInventoryTypeOk returns a tuple with the InventoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryType

`func (o *CreateInventoryRequestBodyFields) SetInventoryType(v string)`

SetInventoryType sets InventoryType field to given value.

### HasInventoryType

`func (o *CreateInventoryRequestBodyFields) HasInventoryType() bool`

HasInventoryType returns a boolean if a field has been set.

### GetInventorySource

`func (o *CreateInventoryRequestBodyFields) GetInventorySource() string`

GetInventorySource returns the InventorySource field if non-nil, zero value otherwise.

### GetInventorySourceOk

`func (o *CreateInventoryRequestBodyFields) GetInventorySourceOk() (*string, bool)`

GetInventorySourceOk returns a tuple with the InventorySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySource

`func (o *CreateInventoryRequestBodyFields) SetInventorySource(v string)`

SetInventorySource sets InventorySource field to given value.


### GetGroup

`func (o *CreateInventoryRequestBodyFields) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CreateInventoryRequestBodyFields) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CreateInventoryRequestBodyFields) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CreateInventoryRequestBodyFields) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetInventoryKey

`func (o *CreateInventoryRequestBodyFields) GetInventoryKey() string`

GetInventoryKey returns the InventoryKey field if non-nil, zero value otherwise.

### GetInventoryKeyOk

`func (o *CreateInventoryRequestBodyFields) GetInventoryKeyOk() (*string, bool)`

GetInventoryKeyOk returns a tuple with the InventoryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryKey

`func (o *CreateInventoryRequestBodyFields) SetInventoryKey(v string)`

SetInventoryKey sets InventoryKey field to given value.

### HasInventoryKey

`func (o *CreateInventoryRequestBodyFields) HasInventoryKey() bool`

HasInventoryKey returns a boolean if a field has been set.

### GetInventoryValue

`func (o *CreateInventoryRequestBodyFields) GetInventoryValue() string`

GetInventoryValue returns the InventoryValue field if non-nil, zero value otherwise.

### GetInventoryValueOk

`func (o *CreateInventoryRequestBodyFields) GetInventoryValueOk() (*string, bool)`

GetInventoryValueOk returns a tuple with the InventoryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryValue

`func (o *CreateInventoryRequestBodyFields) SetInventoryValue(v string)`

SetInventoryValue sets InventoryValue field to given value.

### HasInventoryValue

`func (o *CreateInventoryRequestBodyFields) HasInventoryValue() bool`

HasInventoryValue returns a boolean if a field has been set.

### GetTagSchema

`func (o *CreateInventoryRequestBodyFields) GetTagSchema() string`

GetTagSchema returns the TagSchema field if non-nil, zero value otherwise.

### GetTagSchemaOk

`func (o *CreateInventoryRequestBodyFields) GetTagSchemaOk() (*string, bool)`

GetTagSchemaOk returns a tuple with the TagSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagSchema

`func (o *CreateInventoryRequestBodyFields) SetTagSchema(v string)`

SetTagSchema sets TagSchema field to given value.

### HasTagSchema

`func (o *CreateInventoryRequestBodyFields) HasTagSchema() bool`

HasTagSchema returns a boolean if a field has been set.

### GetDynamicTags

`func (o *CreateInventoryRequestBodyFields) GetDynamicTags() bool`

GetDynamicTags returns the DynamicTags field if non-nil, zero value otherwise.

### GetDynamicTagsOk

`func (o *CreateInventoryRequestBodyFields) GetDynamicTagsOk() (*bool, bool)`

GetDynamicTagsOk returns a tuple with the DynamicTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicTags

`func (o *CreateInventoryRequestBodyFields) SetDynamicTags(v bool)`

SetDynamicTags sets DynamicTags field to given value.

### HasDynamicTags

`func (o *CreateInventoryRequestBodyFields) HasDynamicTags() bool`

HasDynamicTags returns a boolean if a field has been set.

### GetCriticality

`func (o *CreateInventoryRequestBodyFields) GetCriticality() string`

GetCriticality returns the Criticality field if non-nil, zero value otherwise.

### GetCriticalityOk

`func (o *CreateInventoryRequestBodyFields) GetCriticalityOk() (*string, bool)`

GetCriticalityOk returns a tuple with the Criticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticality

`func (o *CreateInventoryRequestBodyFields) SetCriticality(v string)`

SetCriticality sets Criticality field to given value.

### HasCriticality

`func (o *CreateInventoryRequestBodyFields) HasCriticality() bool`

HasCriticality returns a boolean if a field has been set.

### GetTags

`func (o *CreateInventoryRequestBodyFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateInventoryRequestBodyFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateInventoryRequestBodyFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateInventoryRequestBodyFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSuppressed

`func (o *CreateInventoryRequestBodyFields) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *CreateInventoryRequestBodyFields) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *CreateInventoryRequestBodyFields) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *CreateInventoryRequestBodyFields) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


