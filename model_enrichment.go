/*
Sumo Logic CSE API

 https://help.sumologic.com/APIs 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologiccse

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the Enrichment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Enrichment{}

// Enrichment struct for Enrichment
type Enrichment struct {
	Type string `json:"type"`
	// A map of the enrichment details
	Detail map[string]interface{} `json:"detail"`
	ExternalUrl *string `json:"externalUrl,omitempty"`
	Reputation *string `json:"reputation,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	NormalizedIndicator *EnrichmentNormalizedIndicator `json:"normalizedIndicator,omitempty"`
}

type _Enrichment Enrichment

// NewEnrichment instantiates a new Enrichment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrichment(type_ string, detail map[string]interface{}) *Enrichment {
	this := Enrichment{}
	this.Type = type_
	this.Detail = detail
	return &this
}

// NewEnrichmentWithDefaults instantiates a new Enrichment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrichmentWithDefaults() *Enrichment {
	this := Enrichment{}
	return &this
}

// GetType returns the Type field value
func (o *Enrichment) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Enrichment) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Enrichment) SetType(v string) {
	o.Type = v
}

// GetDetail returns the Detail field value
func (o *Enrichment) GetDetail() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *Enrichment) GetDetailOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Detail, true
}

// SetDetail sets field value
func (o *Enrichment) SetDetail(v map[string]interface{}) {
	o.Detail = v
}

// GetExternalUrl returns the ExternalUrl field value if set, zero value otherwise.
func (o *Enrichment) GetExternalUrl() string {
	if o == nil || IsNil(o.ExternalUrl) {
		var ret string
		return ret
	}
	return *o.ExternalUrl
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Enrichment) GetExternalUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalUrl) {
		return nil, false
	}
	return o.ExternalUrl, true
}

// HasExternalUrl returns a boolean if a field has been set.
func (o *Enrichment) HasExternalUrl() bool {
	if o != nil && !IsNil(o.ExternalUrl) {
		return true
	}

	return false
}

// SetExternalUrl gets a reference to the given string and assigns it to the ExternalUrl field.
func (o *Enrichment) SetExternalUrl(v string) {
	o.ExternalUrl = &v
}

// GetReputation returns the Reputation field value if set, zero value otherwise.
func (o *Enrichment) GetReputation() string {
	if o == nil || IsNil(o.Reputation) {
		var ret string
		return ret
	}
	return *o.Reputation
}

// GetReputationOk returns a tuple with the Reputation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Enrichment) GetReputationOk() (*string, bool) {
	if o == nil || IsNil(o.Reputation) {
		return nil, false
	}
	return o.Reputation, true
}

// HasReputation returns a boolean if a field has been set.
func (o *Enrichment) HasReputation() bool {
	if o != nil && !IsNil(o.Reputation) {
		return true
	}

	return false
}

// SetReputation gets a reference to the given string and assigns it to the Reputation field.
func (o *Enrichment) SetReputation(v string) {
	o.Reputation = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *Enrichment) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Enrichment) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *Enrichment) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *Enrichment) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetNormalizedIndicator returns the NormalizedIndicator field value if set, zero value otherwise.
func (o *Enrichment) GetNormalizedIndicator() EnrichmentNormalizedIndicator {
	if o == nil || IsNil(o.NormalizedIndicator) {
		var ret EnrichmentNormalizedIndicator
		return ret
	}
	return *o.NormalizedIndicator
}

// GetNormalizedIndicatorOk returns a tuple with the NormalizedIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Enrichment) GetNormalizedIndicatorOk() (*EnrichmentNormalizedIndicator, bool) {
	if o == nil || IsNil(o.NormalizedIndicator) {
		return nil, false
	}
	return o.NormalizedIndicator, true
}

// HasNormalizedIndicator returns a boolean if a field has been set.
func (o *Enrichment) HasNormalizedIndicator() bool {
	if o != nil && !IsNil(o.NormalizedIndicator) {
		return true
	}

	return false
}

// SetNormalizedIndicator gets a reference to the given EnrichmentNormalizedIndicator and assigns it to the NormalizedIndicator field.
func (o *Enrichment) SetNormalizedIndicator(v EnrichmentNormalizedIndicator) {
	o.NormalizedIndicator = &v
}

func (o Enrichment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Enrichment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["detail"] = o.Detail
	if !IsNil(o.ExternalUrl) {
		toSerialize["externalUrl"] = o.ExternalUrl
	}
	if !IsNil(o.Reputation) {
		toSerialize["reputation"] = o.Reputation
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !IsNil(o.NormalizedIndicator) {
		toSerialize["normalizedIndicator"] = o.NormalizedIndicator
	}
	return toSerialize, nil
}

func (o *Enrichment) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"detail",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEnrichment := _Enrichment{}

	err = json.Unmarshal(bytes, &varEnrichment)

	if err != nil {
		return err
	}

	*o = Enrichment(varEnrichment)

	return err
}

type NullableEnrichment struct {
	value *Enrichment
	isSet bool
}

func (v NullableEnrichment) Get() *Enrichment {
	return v.value
}

func (v *NullableEnrichment) Set(val *Enrichment) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrichment) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrichment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrichment(val *Enrichment) *NullableEnrichment {
	return &NullableEnrichment{value: val, isSet: true}
}

func (v NullableEnrichment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrichment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


