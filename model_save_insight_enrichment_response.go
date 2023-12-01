/*
Sumo Logic CSE API

 https://help.sumologic.com/APIs 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologiccse

import (
	"encoding/json"
	"fmt"
)

// checks if the SaveInsightEnrichmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveInsightEnrichmentResponse{}

// SaveInsightEnrichmentResponse struct for SaveInsightEnrichmentResponse
type SaveInsightEnrichmentResponse struct {
	Data Enrichment `json:"data"`
}

type _SaveInsightEnrichmentResponse SaveInsightEnrichmentResponse

// NewSaveInsightEnrichmentResponse instantiates a new SaveInsightEnrichmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveInsightEnrichmentResponse(data Enrichment) *SaveInsightEnrichmentResponse {
	this := SaveInsightEnrichmentResponse{}
	this.Data = data
	return &this
}

// NewSaveInsightEnrichmentResponseWithDefaults instantiates a new SaveInsightEnrichmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveInsightEnrichmentResponseWithDefaults() *SaveInsightEnrichmentResponse {
	this := SaveInsightEnrichmentResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SaveInsightEnrichmentResponse) GetData() Enrichment {
	if o == nil {
		var ret Enrichment
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SaveInsightEnrichmentResponse) GetDataOk() (*Enrichment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SaveInsightEnrichmentResponse) SetData(v Enrichment) {
	o.Data = v
}

func (o SaveInsightEnrichmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveInsightEnrichmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SaveInsightEnrichmentResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varSaveInsightEnrichmentResponse := _SaveInsightEnrichmentResponse{}

	err = json.Unmarshal(bytes, &varSaveInsightEnrichmentResponse)

	if err != nil {
		return err
	}

	*o = SaveInsightEnrichmentResponse(varSaveInsightEnrichmentResponse)

	return err
}

type NullableSaveInsightEnrichmentResponse struct {
	value *SaveInsightEnrichmentResponse
	isSet bool
}

func (v NullableSaveInsightEnrichmentResponse) Get() *SaveInsightEnrichmentResponse {
	return v.value
}

func (v *NullableSaveInsightEnrichmentResponse) Set(val *SaveInsightEnrichmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveInsightEnrichmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveInsightEnrichmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveInsightEnrichmentResponse(val *SaveInsightEnrichmentResponse) *NullableSaveInsightEnrichmentResponse {
	return &NullableSaveInsightEnrichmentResponse{value: val, isSet: true}
}

func (v NullableSaveInsightEnrichmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveInsightEnrichmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


