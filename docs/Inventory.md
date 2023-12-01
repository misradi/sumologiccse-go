# Inventory

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
**InventoryType** | Pointer to **string** |  | [optional] 
**InventorySource** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**InventoryKey** | Pointer to **string** |  | [optional] 
**InventoryValue** | Pointer to **string** |  | [optional] 
**TagSchema** | Pointer to **string** |  | [optional] 
**DynamicTags** | Pointer to **bool** |  | [optional] 

## Methods

### NewInventory

`func NewInventory(id string, name string, configurationType string, created time.Time, createdBy string, ) *Inventory`

NewInventory instantiates a new Inventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryWithDefaults

`func NewInventoryWithDefaults() *Inventory`

NewInventoryWithDefaults instantiates a new Inventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Inventory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Inventory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Inventory) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Inventory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Inventory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Inventory) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Inventory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Inventory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Inventory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Inventory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfigurationType

`func (o *Inventory) GetConfigurationType() string`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *Inventory) GetConfigurationTypeOk() (*string, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *Inventory) SetConfigurationType(v string)`

SetConfigurationType sets ConfigurationType field to given value.


### GetTags

`func (o *Inventory) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Inventory) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Inventory) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Inventory) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCriticality

`func (o *Inventory) GetCriticality() string`

GetCriticality returns the Criticality field if non-nil, zero value otherwise.

### GetCriticalityOk

`func (o *Inventory) GetCriticalityOk() (*string, bool)`

GetCriticalityOk returns a tuple with the Criticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticality

`func (o *Inventory) SetCriticality(v string)`

SetCriticality sets Criticality field to given value.

### HasCriticality

`func (o *Inventory) HasCriticality() bool`

HasCriticality returns a boolean if a field has been set.

### GetSuppressed

`func (o *Inventory) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *Inventory) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *Inventory) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *Inventory) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetCreated

`func (o *Inventory) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Inventory) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Inventory) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCreatedBy

`func (o *Inventory) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Inventory) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Inventory) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastUpdated

`func (o *Inventory) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Inventory) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Inventory) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Inventory) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *Inventory) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *Inventory) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *Inventory) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *Inventory) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetInventoryType

`func (o *Inventory) GetInventoryType() string`

GetInventoryType returns the InventoryType field if non-nil, zero value otherwise.

### GetInventoryTypeOk

`func (o *Inventory) GetInventoryTypeOk() (*string, bool)`

GetInventoryTypeOk returns a tuple with the InventoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryType

`func (o *Inventory) SetInventoryType(v string)`

SetInventoryType sets InventoryType field to given value.

### HasInventoryType

`func (o *Inventory) HasInventoryType() bool`

HasInventoryType returns a boolean if a field has been set.

### GetInventorySource

`func (o *Inventory) GetInventorySource() string`

GetInventorySource returns the InventorySource field if non-nil, zero value otherwise.

### GetInventorySourceOk

`func (o *Inventory) GetInventorySourceOk() (*string, bool)`

GetInventorySourceOk returns a tuple with the InventorySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySource

`func (o *Inventory) SetInventorySource(v string)`

SetInventorySource sets InventorySource field to given value.

### HasInventorySource

`func (o *Inventory) HasInventorySource() bool`

HasInventorySource returns a boolean if a field has been set.

### GetGroup

`func (o *Inventory) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Inventory) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Inventory) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Inventory) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetInventoryKey

`func (o *Inventory) GetInventoryKey() string`

GetInventoryKey returns the InventoryKey field if non-nil, zero value otherwise.

### GetInventoryKeyOk

`func (o *Inventory) GetInventoryKeyOk() (*string, bool)`

GetInventoryKeyOk returns a tuple with the InventoryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryKey

`func (o *Inventory) SetInventoryKey(v string)`

SetInventoryKey sets InventoryKey field to given value.

### HasInventoryKey

`func (o *Inventory) HasInventoryKey() bool`

HasInventoryKey returns a boolean if a field has been set.

### GetInventoryValue

`func (o *Inventory) GetInventoryValue() string`

GetInventoryValue returns the InventoryValue field if non-nil, zero value otherwise.

### GetInventoryValueOk

`func (o *Inventory) GetInventoryValueOk() (*string, bool)`

GetInventoryValueOk returns a tuple with the InventoryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryValue

`func (o *Inventory) SetInventoryValue(v string)`

SetInventoryValue sets InventoryValue field to given value.

### HasInventoryValue

`func (o *Inventory) HasInventoryValue() bool`

HasInventoryValue returns a boolean if a field has been set.

### GetTagSchema

`func (o *Inventory) GetTagSchema() string`

GetTagSchema returns the TagSchema field if non-nil, zero value otherwise.

### GetTagSchemaOk

`func (o *Inventory) GetTagSchemaOk() (*string, bool)`

GetTagSchemaOk returns a tuple with the TagSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagSchema

`func (o *Inventory) SetTagSchema(v string)`

SetTagSchema sets TagSchema field to given value.

### HasTagSchema

`func (o *Inventory) HasTagSchema() bool`

HasTagSchema returns a boolean if a field has been set.

### GetDynamicTags

`func (o *Inventory) GetDynamicTags() bool`

GetDynamicTags returns the DynamicTags field if non-nil, zero value otherwise.

### GetDynamicTagsOk

`func (o *Inventory) GetDynamicTagsOk() (*bool, bool)`

GetDynamicTagsOk returns a tuple with the DynamicTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicTags

`func (o *Inventory) SetDynamicTags(v bool)`

SetDynamicTags sets DynamicTags field to given value.

### HasDynamicTags

`func (o *Inventory) HasDynamicTags() bool`

HasDynamicTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


