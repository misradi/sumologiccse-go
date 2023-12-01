# Enrichment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Detail** | **map[string]interface{}** | A map of the enrichment details | 
**ExternalUrl** | Pointer to **string** |  | [optional] 
**Reputation** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**NormalizedIndicator** | Pointer to [**EnrichmentNormalizedIndicator**](EnrichmentNormalizedIndicator.md) |  | [optional] 

## Methods

### NewEnrichment

`func NewEnrichment(type_ string, detail map[string]interface{}, ) *Enrichment`

NewEnrichment instantiates a new Enrichment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrichmentWithDefaults

`func NewEnrichmentWithDefaults() *Enrichment`

NewEnrichmentWithDefaults instantiates a new Enrichment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Enrichment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Enrichment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Enrichment) SetType(v string)`

SetType sets Type field to given value.


### GetDetail

`func (o *Enrichment) GetDetail() map[string]interface{}`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Enrichment) GetDetailOk() (*map[string]interface{}, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Enrichment) SetDetail(v map[string]interface{})`

SetDetail sets Detail field to given value.


### GetExternalUrl

`func (o *Enrichment) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *Enrichment) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *Enrichment) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.

### HasExternalUrl

`func (o *Enrichment) HasExternalUrl() bool`

HasExternalUrl returns a boolean if a field has been set.

### GetReputation

`func (o *Enrichment) GetReputation() string`

GetReputation returns the Reputation field if non-nil, zero value otherwise.

### GetReputationOk

`func (o *Enrichment) GetReputationOk() (*string, bool)`

GetReputationOk returns a tuple with the Reputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputation

`func (o *Enrichment) SetReputation(v string)`

SetReputation sets Reputation field to given value.

### HasReputation

`func (o *Enrichment) HasReputation() bool`

HasReputation returns a boolean if a field has been set.

### GetExpiresAt

`func (o *Enrichment) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Enrichment) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Enrichment) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Enrichment) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetNormalizedIndicator

`func (o *Enrichment) GetNormalizedIndicator() EnrichmentNormalizedIndicator`

GetNormalizedIndicator returns the NormalizedIndicator field if non-nil, zero value otherwise.

### GetNormalizedIndicatorOk

`func (o *Enrichment) GetNormalizedIndicatorOk() (*EnrichmentNormalizedIndicator, bool)`

GetNormalizedIndicatorOk returns a tuple with the NormalizedIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedIndicator

`func (o *Enrichment) SetNormalizedIndicator(v EnrichmentNormalizedIndicator)`

SetNormalizedIndicator sets NormalizedIndicator field to given value.

### HasNormalizedIndicator

`func (o *Enrichment) HasNormalizedIndicator() bool`

HasNormalizedIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


