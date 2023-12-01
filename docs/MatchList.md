# MatchList

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

### NewMatchList

`func NewMatchList(id string, name string, targetColumn string, ) *MatchList`

NewMatchList instantiates a new MatchList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchListWithDefaults

`func NewMatchListWithDefaults() *MatchList`

NewMatchListWithDefaults instantiates a new MatchList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MatchList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MatchList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MatchList) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MatchList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MatchList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MatchList) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MatchList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MatchList) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MatchList) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MatchList) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultTtl

`func (o *MatchList) GetDefaultTtl() int32`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *MatchList) GetDefaultTtlOk() (*int32, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *MatchList) SetDefaultTtl(v int32)`

SetDefaultTtl sets DefaultTtl field to given value.

### HasDefaultTtl

`func (o *MatchList) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.

### GetTargetColumn

`func (o *MatchList) GetTargetColumn() string`

GetTargetColumn returns the TargetColumn field if non-nil, zero value otherwise.

### GetTargetColumnOk

`func (o *MatchList) GetTargetColumnOk() (*string, bool)`

GetTargetColumnOk returns a tuple with the TargetColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetColumn

`func (o *MatchList) SetTargetColumn(v string)`

SetTargetColumn sets TargetColumn field to given value.


### GetCreated

`func (o *MatchList) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MatchList) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MatchList) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *MatchList) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *MatchList) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MatchList) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MatchList) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MatchList) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastUpdated

`func (o *MatchList) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *MatchList) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *MatchList) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *MatchList) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *MatchList) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *MatchList) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *MatchList) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *MatchList) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


