# GetThreatIntelIndicatorResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Value** | **string** |  | 
**Active** | **bool** |  | 
**Expiration** | Pointer to **time.Time** |  | [optional] 
**Meta** | Pointer to [**GetThreatIntelIndicatorResponseDataMeta**](GetThreatIntelIndicatorResponseDataMeta.md) |  | [optional] 
**Source** | [**GetThreatIntelIndicatorResponseDataSource**](GetThreatIntelIndicatorResponseDataSource.md) |  | 

## Methods

### NewGetThreatIntelIndicatorResponseData

`func NewGetThreatIntelIndicatorResponseData(id string, value string, active bool, source GetThreatIntelIndicatorResponseDataSource, ) *GetThreatIntelIndicatorResponseData`

NewGetThreatIntelIndicatorResponseData instantiates a new GetThreatIntelIndicatorResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetThreatIntelIndicatorResponseDataWithDefaults

`func NewGetThreatIntelIndicatorResponseDataWithDefaults() *GetThreatIntelIndicatorResponseData`

NewGetThreatIntelIndicatorResponseDataWithDefaults instantiates a new GetThreatIntelIndicatorResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetThreatIntelIndicatorResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetThreatIntelIndicatorResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetThreatIntelIndicatorResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetValue

`func (o *GetThreatIntelIndicatorResponseData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetThreatIntelIndicatorResponseData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetThreatIntelIndicatorResponseData) SetValue(v string)`

SetValue sets Value field to given value.


### GetActive

`func (o *GetThreatIntelIndicatorResponseData) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetThreatIntelIndicatorResponseData) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetThreatIntelIndicatorResponseData) SetActive(v bool)`

SetActive sets Active field to given value.


### GetExpiration

`func (o *GetThreatIntelIndicatorResponseData) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *GetThreatIntelIndicatorResponseData) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *GetThreatIntelIndicatorResponseData) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *GetThreatIntelIndicatorResponseData) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetMeta

`func (o *GetThreatIntelIndicatorResponseData) GetMeta() GetThreatIntelIndicatorResponseDataMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetThreatIntelIndicatorResponseData) GetMetaOk() (*GetThreatIntelIndicatorResponseDataMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetThreatIntelIndicatorResponseData) SetMeta(v GetThreatIntelIndicatorResponseDataMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetThreatIntelIndicatorResponseData) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSource

`func (o *GetThreatIntelIndicatorResponseData) GetSource() GetThreatIntelIndicatorResponseDataSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetThreatIntelIndicatorResponseData) GetSourceOk() (*GetThreatIntelIndicatorResponseDataSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetThreatIntelIndicatorResponseData) SetSource(v GetThreatIntelIndicatorResponseDataSource)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


