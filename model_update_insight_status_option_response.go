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

// checks if the UpdateInsightStatusOptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateInsightStatusOptionResponse{}

// UpdateInsightStatusOptionResponse struct for UpdateInsightStatusOptionResponse
type UpdateInsightStatusOptionResponse struct {
	Data InsightStatus `json:"data"`
}

type _UpdateInsightStatusOptionResponse UpdateInsightStatusOptionResponse

// NewUpdateInsightStatusOptionResponse instantiates a new UpdateInsightStatusOptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInsightStatusOptionResponse(data InsightStatus) *UpdateInsightStatusOptionResponse {
	this := UpdateInsightStatusOptionResponse{}
	this.Data = data
	return &this
}

// NewUpdateInsightStatusOptionResponseWithDefaults instantiates a new UpdateInsightStatusOptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInsightStatusOptionResponseWithDefaults() *UpdateInsightStatusOptionResponse {
	this := UpdateInsightStatusOptionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateInsightStatusOptionResponse) GetData() InsightStatus {
	if o == nil {
		var ret InsightStatus
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateInsightStatusOptionResponse) GetDataOk() (*InsightStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateInsightStatusOptionResponse) SetData(v InsightStatus) {
	o.Data = v
}

func (o UpdateInsightStatusOptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateInsightStatusOptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UpdateInsightStatusOptionResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varUpdateInsightStatusOptionResponse := _UpdateInsightStatusOptionResponse{}

	err = json.Unmarshal(bytes, &varUpdateInsightStatusOptionResponse)

	if err != nil {
		return err
	}

	*o = UpdateInsightStatusOptionResponse(varUpdateInsightStatusOptionResponse)

	return err
}

type NullableUpdateInsightStatusOptionResponse struct {
	value *UpdateInsightStatusOptionResponse
	isSet bool
}

func (v NullableUpdateInsightStatusOptionResponse) Get() *UpdateInsightStatusOptionResponse {
	return v.value
}

func (v *NullableUpdateInsightStatusOptionResponse) Set(val *UpdateInsightStatusOptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateInsightStatusOptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateInsightStatusOptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateInsightStatusOptionResponse(val *UpdateInsightStatusOptionResponse) *NullableUpdateInsightStatusOptionResponse {
	return &NullableUpdateInsightStatusOptionResponse{value: val, isSet: true}
}

func (v NullableUpdateInsightStatusOptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateInsightStatusOptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


