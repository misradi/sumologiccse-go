# MatchListItem

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

### NewMatchListItem

`func NewMatchListItem(id string, value string, active bool, listName string, ) *MatchListItem`

NewMatchListItem instantiates a new MatchListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchListItemWithDefaults

`func NewMatchListItemWithDefaults() *MatchListItem`

NewMatchListItemWithDefaults instantiates a new MatchListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MatchListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MatchListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MatchListItem) SetId(v string)`

SetId sets Id field to given value.


### GetValue

`func (o *MatchListItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MatchListItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MatchListItem) SetValue(v string)`

SetValue sets Value field to given value.


### GetActive

`func (o *MatchListItem) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MatchListItem) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MatchListItem) SetActive(v bool)`

SetActive sets Active field to given value.


### GetExpiration

`func (o *MatchListItem) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *MatchListItem) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *MatchListItem) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *MatchListItem) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetMeta

`func (o *MatchListItem) GetMeta() GetThreatIntelIndicatorResponseDataMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MatchListItem) GetMetaOk() (*GetThreatIntelIndicatorResponseDataMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MatchListItem) SetMeta(v GetThreatIntelIndicatorResponseDataMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MatchListItem) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetListName

`func (o *MatchListItem) GetListName() string`

GetListName returns the ListName field if non-nil, zero value otherwise.

### GetListNameOk

`func (o *MatchListItem) GetListNameOk() (*string, bool)`

GetListNameOk returns a tuple with the ListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListName

`func (o *MatchListItem) SetListName(v string)`

SetListName sets ListName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


