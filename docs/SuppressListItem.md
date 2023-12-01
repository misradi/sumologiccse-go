# SuppressListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Value** | **string** |  | 
**Active** | **bool** |  | 
**Expiration** | Pointer to **time.Time** |  | [optional] 
**Meta** | Pointer to [**GetThreatIntelIndicatorResponseDataMeta**](GetThreatIntelIndicatorResponseDataMeta.md) |  | [optional] 
**ListName** | **string** |  | 

## Methods

### NewSuppressListItem

`func NewSuppressListItem(id string, value string, active bool, listName string, ) *SuppressListItem`

NewSuppressListItem instantiates a new SuppressListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuppressListItemWithDefaults

`func NewSuppressListItemWithDefaults() *SuppressListItem`

NewSuppressListItemWithDefaults instantiates a new SuppressListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SuppressListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuppressListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuppressListItem) SetId(v string)`

SetId sets Id field to given value.


### GetValue

`func (o *SuppressListItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SuppressListItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SuppressListItem) SetValue(v string)`

SetValue sets Value field to given value.


### GetActive

`func (o *SuppressListItem) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SuppressListItem) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SuppressListItem) SetActive(v bool)`

SetActive sets Active field to given value.


### GetExpiration

`func (o *SuppressListItem) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *SuppressListItem) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *SuppressListItem) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *SuppressListItem) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetMeta

`func (o *SuppressListItem) GetMeta() GetThreatIntelIndicatorResponseDataMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SuppressListItem) GetMetaOk() (*GetThreatIntelIndicatorResponseDataMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SuppressListItem) SetMeta(v GetThreatIntelIndicatorResponseDataMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SuppressListItem) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetListName

`func (o *SuppressListItem) GetListName() string`

GetListName returns the ListName field if non-nil, zero value otherwise.

### GetListNameOk

`func (o *SuppressListItem) GetListNameOk() (*string, bool)`

GetListNameOk returns a tuple with the ListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListName

`func (o *SuppressListItem) SetListName(v string)`

SetListName sets ListName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


