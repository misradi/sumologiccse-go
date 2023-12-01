# EnrichmentNormalizedIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Indicator** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Imported** | Pointer to **time.Time** |  | [optional] 
**ValidFrom** | Pointer to **time.Time** |  | [optional] 
**ValidUntil** | Pointer to **time.Time** |  | [optional] 
**Confidence** | Pointer to **int32** |  | [optional] 
**ThreatType** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to **string** |  | [optional] 

## Methods

### NewEnrichmentNormalizedIndicator

`func NewEnrichmentNormalizedIndicator() *EnrichmentNormalizedIndicator`

NewEnrichmentNormalizedIndicator instantiates a new EnrichmentNormalizedIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrichmentNormalizedIndicatorWithDefaults

`func NewEnrichmentNormalizedIndicatorWithDefaults() *EnrichmentNormalizedIndicator`

NewEnrichmentNormalizedIndicatorWithDefaults instantiates a new EnrichmentNormalizedIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnrichmentNormalizedIndicator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnrichmentNormalizedIndicator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnrichmentNormalizedIndicator) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnrichmentNormalizedIndicator) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndicator

`func (o *EnrichmentNormalizedIndicator) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *EnrichmentNormalizedIndicator) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *EnrichmentNormalizedIndicator) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.

### HasIndicator

`func (o *EnrichmentNormalizedIndicator) HasIndicator() bool`

HasIndicator returns a boolean if a field has been set.

### GetType

`func (o *EnrichmentNormalizedIndicator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnrichmentNormalizedIndicator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnrichmentNormalizedIndicator) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnrichmentNormalizedIndicator) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSource

`func (o *EnrichmentNormalizedIndicator) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EnrichmentNormalizedIndicator) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EnrichmentNormalizedIndicator) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *EnrichmentNormalizedIndicator) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetImported

`func (o *EnrichmentNormalizedIndicator) GetImported() time.Time`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *EnrichmentNormalizedIndicator) GetImportedOk() (*time.Time, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *EnrichmentNormalizedIndicator) SetImported(v time.Time)`

SetImported sets Imported field to given value.

### HasImported

`func (o *EnrichmentNormalizedIndicator) HasImported() bool`

HasImported returns a boolean if a field has been set.

### GetValidFrom

`func (o *EnrichmentNormalizedIndicator) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *EnrichmentNormalizedIndicator) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *EnrichmentNormalizedIndicator) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *EnrichmentNormalizedIndicator) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidUntil

`func (o *EnrichmentNormalizedIndicator) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *EnrichmentNormalizedIndicator) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *EnrichmentNormalizedIndicator) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *EnrichmentNormalizedIndicator) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### GetConfidence

`func (o *EnrichmentNormalizedIndicator) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *EnrichmentNormalizedIndicator) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *EnrichmentNormalizedIndicator) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *EnrichmentNormalizedIndicator) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetThreatType

`func (o *EnrichmentNormalizedIndicator) GetThreatType() string`

GetThreatType returns the ThreatType field if non-nil, zero value otherwise.

### GetThreatTypeOk

`func (o *EnrichmentNormalizedIndicator) GetThreatTypeOk() (*string, bool)`

GetThreatTypeOk returns a tuple with the ThreatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatType

`func (o *EnrichmentNormalizedIndicator) SetThreatType(v string)`

SetThreatType sets ThreatType field to given value.

### HasThreatType

`func (o *EnrichmentNormalizedIndicator) HasThreatType() bool`

HasThreatType returns a boolean if a field has been set.

### GetFields

`func (o *EnrichmentNormalizedIndicator) GetFields() string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *EnrichmentNormalizedIndicator) GetFieldsOk() (*string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *EnrichmentNormalizedIndicator) SetFields(v string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *EnrichmentNormalizedIndicator) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


