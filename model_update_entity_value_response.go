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

// checks if the UpdateEntityValueResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEntityValueResponse{}

// UpdateEntityValueResponse struct for UpdateEntityValueResponse
type UpdateEntityValueResponse struct {
	Data EntityValue `json:"data"`
}

type _UpdateEntityValueResponse UpdateEntityValueResponse

// NewUpdateEntityValueResponse instantiates a new UpdateEntityValueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEntityValueResponse(data EntityValue) *UpdateEntityValueResponse {
	this := UpdateEntityValueResponse{}
	this.Data = data
	return &this
}

// NewUpdateEntityValueResponseWithDefaults instantiates a new UpdateEntityValueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEntityValueResponseWithDefaults() *UpdateEntityValueResponse {
	this := UpdateEntityValueResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateEntityValueResponse) GetData() EntityValue {
	if o == nil {
		var ret EntityValue
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateEntityValueResponse) GetDataOk() (*EntityValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateEntityValueResponse) SetData(v EntityValue) {
	o.Data = v
}

func (o UpdateEntityValueResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEntityValueResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UpdateEntityValueResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varUpdateEntityValueResponse := _UpdateEntityValueResponse{}

	err = json.Unmarshal(bytes, &varUpdateEntityValueResponse)

	if err != nil {
		return err
	}

	*o = UpdateEntityValueResponse(varUpdateEntityValueResponse)

	return err
}

type NullableUpdateEntityValueResponse struct {
	value *UpdateEntityValueResponse
	isSet bool
}

func (v NullableUpdateEntityValueResponse) Get() *UpdateEntityValueResponse {
	return v.value
}

func (v *NullableUpdateEntityValueResponse) Set(val *UpdateEntityValueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEntityValueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEntityValueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEntityValueResponse(val *UpdateEntityValueResponse) *NullableUpdateEntityValueResponse {
	return &NullableUpdateEntityValueResponse{value: val, isSet: true}
}

func (v NullableUpdateEntityValueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEntityValueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


