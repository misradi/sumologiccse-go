# UpdateThreatIntelIndicatorResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Value** | **string** |  | 
**Active** | **bool** |  | 
**Expiration** | Pointer to **time.Time** |  | [optional] 
**Meta** | Pointer to [**GetThreatIntelIndicatorResponseDataMeta**](GetThreatIntelIndicatorResponseDataMeta.md) |  | [optional] 

## Methods

### NewUpdateThreatIntelIndicatorResponseData

`func NewUpdateThreatIntelIndicatorResponseData(id string, value string, active bool, ) *UpdateThreatIntelIndicatorResponseData`

NewUpdateThreatIntelIndicatorResponseData instantiates a new UpdateThreatIntelIndicatorResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateThreatIntelIndicatorResponseDataWithDefaults

`func NewUpdateThreatIntelIndicatorResponseDataWithDefaults() *UpdateThreatIntelIndicatorResponseData`

NewUpdateThreatIntelIndicatorResponseDataWithDefaults instantiates a new UpdateThreatIntelIndicatorResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateThreatIntelIndicatorResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateThreatIntelIndicatorResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateThreatIntelIndicatorResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetValue

`func (o *UpdateThreatIntelIndicatorResponseData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateThreatIntelIndicatorResponseData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateThreatIntelIndicatorResponseData) SetValue(v string)`

SetValue sets Value field to given value.


### GetActive

`func (o *UpdateThreatIntelIndicatorResponseData) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateThreatIntelIndicatorResponseData) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateThreatIntelIndicatorResponseData) SetActive(v bool)`

SetActive sets Active field to given value.


### GetExpiration

`func (o *UpdateThreatIntelIndicatorResponseData) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *UpdateThreatIntelIndicatorResponseData) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *UpdateThreatIntelIndicatorResponseData) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *UpdateThreatIntelIndicatorResponseData) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetMeta

`func (o *UpdateThreatIntelIndicatorResponseData) GetMeta() GetThreatIntelIndicatorResponseDataMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UpdateThreatIntelIndicatorResponseData) GetMetaOk() (*GetThreatIntelIndicatorResponseDataMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UpdateThreatIntelIndicatorResponseData) SetMeta(v GetThreatIntelIndicatorResponseDataMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UpdateThreatIntelIndicatorResponseData) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


