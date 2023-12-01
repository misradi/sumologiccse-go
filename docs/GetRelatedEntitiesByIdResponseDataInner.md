# GetRelatedEntitiesByIdResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**EntityType** | **string** |  | 
**Value** | **string** |  | 
**ActivityScore** | **int32** |  | 
**Tags** | **[]string** |  | 
**Criticality** | Pointer to **string** |  | [optional] 
**IsSuppressed** | **bool** |  | 

## Methods

### NewGetRelatedEntitiesByIdResponseDataInner

`func NewGetRelatedEntitiesByIdResponseDataInner(id string, entityType string, value string, activityScore int32, tags []string, isSuppressed bool, ) *GetRelatedEntitiesByIdResponseDataInner`

NewGetRelatedEntitiesByIdResponseDataInner instantiates a new GetRelatedEntitiesByIdResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRelatedEntitiesByIdResponseDataInnerWithDefaults

`func NewGetRelatedEntitiesByIdResponseDataInnerWithDefaults() *GetRelatedEntitiesByIdResponseDataInner`

NewGetRelatedEntitiesByIdResponseDataInnerWithDefaults instantiates a new GetRelatedEntitiesByIdResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetRelatedEntitiesByIdResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRelatedEntitiesByIdResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRelatedEntitiesByIdResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetEntityType

`func (o *GetRelatedEntitiesByIdResponseDataInner) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *GetRelatedEntitiesByIdResponseDataInner) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *GetRelatedEntitiesByIdResponseDataInner) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetValue

`func (o *GetRelatedEntitiesByIdResponseDataInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetRelatedEntitiesByIdResponseDataInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetRelatedEntitiesByIdResponseDataInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetActivityScore

`func (o *GetRelatedEntitiesByIdResponseDataInner) GetActivityScore() int32`

GetActivityScore returns the ActivityScore field if non-nil, zero value otherwise.

### GetActivityScoreOk

`func (o *GetRelatedEntitiesByIdResponseDataInner) GetActivityScoreOk() (*int32, bool)`

GetActivityScoreOk returns a tuple with the ActivityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityScore

`func (o *GetRelatedEntitiesByIdResponseDataInner) SetActivityScore(v int32)`

SetActivityScore sets ActivityScore field to given value.


### GetTags

`func (o *GetRelatedEntitiesByIdResponseDataInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetRelatedEntitiesByIdResponseDataInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetRelatedEntitiesByIdResponseDataInner) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetCriticality

`func (o *GetRelatedEntitiesByIdResponseDataInner) GetCriticality() string`

GetCriticality returns the Criticality field if non-nil, zero value otherwise.

### GetCriticalityOk

`func (o *GetRelatedEntitiesByIdResponseDataInner) GetCriticalityOk() (*string, bool)`

GetCriticalityOk returns a tuple with the Criticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticality

`func (o *GetRelatedEntitiesByIdResponseDataInner) SetCriticality(v string)`

SetCriticality sets Criticality field to given value.

### HasCriticality

`func (o *GetRelatedEntitiesByIdResponseDataInner) HasCriticality() bool`

HasCriticality returns a boolean if a field has been set.

### GetIsSuppressed

`func (o *GetRelatedEntitiesByIdResponseDataInner) GetIsSuppressed() bool`

GetIsSuppressed returns the IsSuppressed field if non-nil, zero value otherwise.

### GetIsSuppressedOk

`func (o *GetRelatedEntitiesByIdResponseDataInner) GetIsSuppressedOk() (*bool, bool)`

GetIsSuppressedOk returns a tuple with the IsSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuppressed

`func (o *GetRelatedEntitiesByIdResponseDataInner) SetIsSuppressed(v bool)`

SetIsSuppressed sets IsSuppressed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


