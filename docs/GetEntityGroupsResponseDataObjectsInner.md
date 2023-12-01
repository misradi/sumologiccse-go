# GetEntityGroupsResponseDataObjectsInner

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
**EntityNamespace** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**Suffix** | Pointer to **string** |  | [optional] 
**NetworkBlock** | Pointer to **string** |  | [optional] 

## Methods

### NewGetEntityGroupsResponseDataObjectsInner

`func NewGetEntityGroupsResponseDataObjectsInner(id string, name string, configurationType string, created time.Time, createdBy string, ) *GetEntityGroupsResponseDataObjectsInner`

NewGetEntityGroupsResponseDataObjectsInner instantiates a new GetEntityGroupsResponseDataObjectsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntityGroupsResponseDataObjectsInnerWithDefaults

`func NewGetEntityGroupsResponseDataObjectsInnerWithDefaults() *GetEntityGroupsResponseDataObjectsInner`

NewGetEntityGroupsResponseDataObjectsInnerWithDefaults instantiates a new GetEntityGroupsResponseDataObjectsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetEntityGroupsResponseDataObjectsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEntityGroupsResponseDataObjectsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GetEntityGroupsResponseDataObjectsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetEntityGroupsResponseDataObjectsInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GetEntityGroupsResponseDataObjectsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetEntityGroupsResponseDataObjectsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetEntityGroupsResponseDataObjectsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfigurationType

`func (o *GetEntityGroupsResponseDataObjectsInner) GetConfigurationType() string`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetConfigurationTypeOk() (*string, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *GetEntityGroupsResponseDataObjectsInner) SetConfigurationType(v string)`

SetConfigurationType sets ConfigurationType field to given value.


### GetTags

`func (o *GetEntityGroupsResponseDataObjectsInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetEntityGroupsResponseDataObjectsInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetEntityGroupsResponseDataObjectsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCriticality

`func (o *GetEntityGroupsResponseDataObjectsInner) GetCriticality() string`

GetCriticality returns the Criticality field if non-nil, zero value otherwise.

### GetCriticalityOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetCriticalityOk() (*string, bool)`

GetCriticalityOk returns a tuple with the Criticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticality

`func (o *GetEntityGroupsResponseDataObjectsInner) SetCriticality(v string)`

SetCriticality sets Criticality field to given value.

### HasCriticality

`func (o *GetEntityGroupsResponseDataObjectsInner) HasCriticality() bool`

HasCriticality returns a boolean if a field has been set.

### GetSuppressed

`func (o *GetEntityGroupsResponseDataObjectsInner) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *GetEntityGroupsResponseDataObjectsInner) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *GetEntityGroupsResponseDataObjectsInner) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetCreated

`func (o *GetEntityGroupsResponseDataObjectsInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetEntityGroupsResponseDataObjectsInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCreatedBy

`func (o *GetEntityGroupsResponseDataObjectsInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetEntityGroupsResponseDataObjectsInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastUpdated

`func (o *GetEntityGroupsResponseDataObjectsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetEntityGroupsResponseDataObjectsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetEntityGroupsResponseDataObjectsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *GetEntityGroupsResponseDataObjectsInner) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *GetEntityGroupsResponseDataObjectsInner) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *GetEntityGroupsResponseDataObjectsInner) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetInventoryType

`func (o *GetEntityGroupsResponseDataObjectsInner) GetInventoryType() string`

GetInventoryType returns the InventoryType field if non-nil, zero value otherwise.

### GetInventoryTypeOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetInventoryTypeOk() (*string, bool)`

GetInventoryTypeOk returns a tuple with the InventoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryType

`func (o *GetEntityGroupsResponseDataObjectsInner) SetInventoryType(v string)`

SetInventoryType sets InventoryType field to given value.

### HasInventoryType

`func (o *GetEntityGroupsResponseDataObjectsInner) HasInventoryType() bool`

HasInventoryType returns a boolean if a field has been set.

### GetInventorySource

`func (o *GetEntityGroupsResponseDataObjectsInner) GetInventorySource() string`

GetInventorySource returns the InventorySource field if non-nil, zero value otherwise.

### GetInventorySourceOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetInventorySourceOk() (*string, bool)`

GetInventorySourceOk returns a tuple with the InventorySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySource

`func (o *GetEntityGroupsResponseDataObjectsInner) SetInventorySource(v string)`

SetInventorySource sets InventorySource field to given value.

### HasInventorySource

`func (o *GetEntityGroupsResponseDataObjectsInner) HasInventorySource() bool`

HasInventorySource returns a boolean if a field has been set.

### GetGroup

`func (o *GetEntityGroupsResponseDataObjectsInner) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetEntityGroupsResponseDataObjectsInner) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GetEntityGroupsResponseDataObjectsInner) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetInventoryKey

`func (o *GetEntityGroupsResponseDataObjectsInner) GetInventoryKey() string`

GetInventoryKey returns the InventoryKey field if non-nil, zero value otherwise.

### GetInventoryKeyOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetInventoryKeyOk() (*string, bool)`

GetInventoryKeyOk returns a tuple with the InventoryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryKey

`func (o *GetEntityGroupsResponseDataObjectsInner) SetInventoryKey(v string)`

SetInventoryKey sets InventoryKey field to given value.

### HasInventoryKey

`func (o *GetEntityGroupsResponseDataObjectsInner) HasInventoryKey() bool`

HasInventoryKey returns a boolean if a field has been set.

### GetInventoryValue

`func (o *GetEntityGroupsResponseDataObjectsInner) GetInventoryValue() string`

GetInventoryValue returns the InventoryValue field if non-nil, zero value otherwise.

### GetInventoryValueOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetInventoryValueOk() (*string, bool)`

GetInventoryValueOk returns a tuple with the InventoryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryValue

`func (o *GetEntityGroupsResponseDataObjectsInner) SetInventoryValue(v string)`

SetInventoryValue sets InventoryValue field to given value.

### HasInventoryValue

`func (o *GetEntityGroupsResponseDataObjectsInner) HasInventoryValue() bool`

HasInventoryValue returns a boolean if a field has been set.

### GetTagSchema

`func (o *GetEntityGroupsResponseDataObjectsInner) GetTagSchema() string`

GetTagSchema returns the TagSchema field if non-nil, zero value otherwise.

### GetTagSchemaOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetTagSchemaOk() (*string, bool)`

GetTagSchemaOk returns a tuple with the TagSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagSchema

`func (o *GetEntityGroupsResponseDataObjectsInner) SetTagSchema(v string)`

SetTagSchema sets TagSchema field to given value.

### HasTagSchema

`func (o *GetEntityGroupsResponseDataObjectsInner) HasTagSchema() bool`

HasTagSchema returns a boolean if a field has been set.

### GetDynamicTags

`func (o *GetEntityGroupsResponseDataObjectsInner) GetDynamicTags() bool`

GetDynamicTags returns the DynamicTags field if non-nil, zero value otherwise.

### GetDynamicTagsOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetDynamicTagsOk() (*bool, bool)`

GetDynamicTagsOk returns a tuple with the DynamicTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicTags

`func (o *GetEntityGroupsResponseDataObjectsInner) SetDynamicTags(v bool)`

SetDynamicTags sets DynamicTags field to given value.

### HasDynamicTags

`func (o *GetEntityGroupsResponseDataObjectsInner) HasDynamicTags() bool`

HasDynamicTags returns a boolean if a field has been set.

### GetEntityNamespace

`func (o *GetEntityGroupsResponseDataObjectsInner) GetEntityNamespace() string`

GetEntityNamespace returns the EntityNamespace field if non-nil, zero value otherwise.

### GetEntityNamespaceOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetEntityNamespaceOk() (*string, bool)`

GetEntityNamespaceOk returns a tuple with the EntityNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityNamespace

`func (o *GetEntityGroupsResponseDataObjectsInner) SetEntityNamespace(v string)`

SetEntityNamespace sets EntityNamespace field to given value.

### HasEntityNamespace

`func (o *GetEntityGroupsResponseDataObjectsInner) HasEntityNamespace() bool`

HasEntityNamespace returns a boolean if a field has been set.

### GetEntityType

`func (o *GetEntityGroupsResponseDataObjectsInner) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *GetEntityGroupsResponseDataObjectsInner) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *GetEntityGroupsResponseDataObjectsInner) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetPrefix

`func (o *GetEntityGroupsResponseDataObjectsInner) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetEntityGroupsResponseDataObjectsInner) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GetEntityGroupsResponseDataObjectsInner) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetSuffix

`func (o *GetEntityGroupsResponseDataObjectsInner) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *GetEntityGroupsResponseDataObjectsInner) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *GetEntityGroupsResponseDataObjectsInner) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### GetNetworkBlock

`func (o *GetEntityGroupsResponseDataObjectsInner) GetNetworkBlock() string`

GetNetworkBlock returns the NetworkBlock field if non-nil, zero value otherwise.

### GetNetworkBlockOk

`func (o *GetEntityGroupsResponseDataObjectsInner) GetNetworkBlockOk() (*string, bool)`

GetNetworkBlockOk returns a tuple with the NetworkBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBlock

`func (o *GetEntityGroupsResponseDataObjectsInner) SetNetworkBlock(v string)`

SetNetworkBlock sets NetworkBlock field to given value.

### HasNetworkBlock

`func (o *GetEntityGroupsResponseDataObjectsInner) HasNetworkBlock() bool`

HasNetworkBlock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


