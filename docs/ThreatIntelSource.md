# ThreatIntelSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | The name of the Source. | 
**Description** | Pointer to **string** | A description of the Source. | [optional] 
**SourceType** | **string** |  | 
**Created** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastUpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewThreatIntelSource

`func NewThreatIntelSource(id string, name string, sourceType string, ) *ThreatIntelSource`

NewThreatIntelSource instantiates a new ThreatIntelSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatIntelSourceWithDefaults

`func NewThreatIntelSourceWithDefaults() *ThreatIntelSource`

NewThreatIntelSourceWithDefaults instantiates a new ThreatIntelSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThreatIntelSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThreatIntelSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThreatIntelSource) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ThreatIntelSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThreatIntelSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThreatIntelSource) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ThreatIntelSource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThreatIntelSource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThreatIntelSource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThreatIntelSource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSourceType

`func (o *ThreatIntelSource) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ThreatIntelSource) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ThreatIntelSource) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetCreated

`func (o *ThreatIntelSource) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ThreatIntelSource) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ThreatIntelSource) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ThreatIntelSource) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ThreatIntelSource) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ThreatIntelSource) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ThreatIntelSource) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ThreatIntelSource) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ThreatIntelSource) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ThreatIntelSource) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ThreatIntelSource) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ThreatIntelSource) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *ThreatIntelSource) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *ThreatIntelSource) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *ThreatIntelSource) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *ThreatIntelSource) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


