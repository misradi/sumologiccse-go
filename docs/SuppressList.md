# SuppressList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | The name of the List. | 
**Description** | Pointer to **string** | A description of the List. | [optional] 
**DefaultTtl** | Pointer to **int32** | The default time-to-live (in seconds) for new Items added to this List. This default is only used to default the field in the UI, and is not used at all when new Items are added via the API. | [optional] 
**TargetColumn** | **string** | The column that Items in this List are matched against. | 
**Created** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastUpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewSuppressList

`func NewSuppressList(id string, name string, targetColumn string, ) *SuppressList`

NewSuppressList instantiates a new SuppressList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuppressListWithDefaults

`func NewSuppressListWithDefaults() *SuppressList`

NewSuppressListWithDefaults instantiates a new SuppressList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SuppressList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuppressList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuppressList) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SuppressList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SuppressList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SuppressList) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SuppressList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SuppressList) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SuppressList) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SuppressList) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultTtl

`func (o *SuppressList) GetDefaultTtl() int32`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *SuppressList) GetDefaultTtlOk() (*int32, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *SuppressList) SetDefaultTtl(v int32)`

SetDefaultTtl sets DefaultTtl field to given value.

### HasDefaultTtl

`func (o *SuppressList) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.

### GetTargetColumn

`func (o *SuppressList) GetTargetColumn() string`

GetTargetColumn returns the TargetColumn field if non-nil, zero value otherwise.

### GetTargetColumnOk

`func (o *SuppressList) GetTargetColumnOk() (*string, bool)`

GetTargetColumnOk returns a tuple with the TargetColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetColumn

`func (o *SuppressList) SetTargetColumn(v string)`

SetTargetColumn sets TargetColumn field to given value.


### GetCreated

`func (o *SuppressList) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SuppressList) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SuppressList) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SuppressList) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SuppressList) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SuppressList) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SuppressList) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SuppressList) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SuppressList) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SuppressList) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SuppressList) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SuppressList) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *SuppressList) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *SuppressList) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *SuppressList) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *SuppressList) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


