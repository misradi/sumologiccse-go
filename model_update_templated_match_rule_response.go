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

// checks if the UpdateTemplatedMatchRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTemplatedMatchRuleResponse{}

// UpdateTemplatedMatchRuleResponse struct for UpdateTemplatedMatchRuleResponse
type UpdateTemplatedMatchRuleResponse struct {
	Data TemplatedMatchRule `json:"data"`
}

type _UpdateTemplatedMatchRuleResponse UpdateTemplatedMatchRuleResponse

// NewUpdateTemplatedMatchRuleResponse instantiates a new UpdateTemplatedMatchRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTemplatedMatchRuleResponse(data TemplatedMatchRule) *UpdateTemplatedMatchRuleResponse {
	this := UpdateTemplatedMatchRuleResponse{}
	this.Data = data
	return &this
}

// NewUpdateTemplatedMatchRuleResponseWithDefaults instantiates a new UpdateTemplatedMatchRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTemplatedMatchRuleResponseWithDefaults() *UpdateTemplatedMatchRuleResponse {
	this := UpdateTemplatedMatchRuleResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateTemplatedMatchRuleResponse) GetData() TemplatedMatchRule {
	if o == nil {
		var ret TemplatedMatchRule
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateTemplatedMatchRuleResponse) GetDataOk() (*TemplatedMatchRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateTemplatedMatchRuleResponse) SetData(v TemplatedMatchRule) {
	o.Data = v
}

func (o UpdateTemplatedMatchRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTemplatedMatchRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UpdateTemplatedMatchRuleResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varUpdateTemplatedMatchRuleResponse := _UpdateTemplatedMatchRuleResponse{}

	err = json.Unmarshal(bytes, &varUpdateTemplatedMatchRuleResponse)

	if err != nil {
		return err
	}

	*o = UpdateTemplatedMatchRuleResponse(varUpdateTemplatedMatchRuleResponse)

	return err
}

type NullableUpdateTemplatedMatchRuleResponse struct {
	value *UpdateTemplatedMatchRuleResponse
	isSet bool
}

func (v NullableUpdateTemplatedMatchRuleResponse) Get() *UpdateTemplatedMatchRuleResponse {
	return v.value
}

func (v *NullableUpdateTemplatedMatchRuleResponse) Set(val *UpdateTemplatedMatchRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTemplatedMatchRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTemplatedMatchRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTemplatedMatchRuleResponse(val *UpdateTemplatedMatchRuleResponse) *NullableUpdateTemplatedMatchRuleResponse {
	return &NullableUpdateTemplatedMatchRuleResponse{value: val, isSet: true}
}

func (v NullableUpdateTemplatedMatchRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTemplatedMatchRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

