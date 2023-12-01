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

// checks if the UpdateMatchRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMatchRuleResponse{}

// UpdateMatchRuleResponse struct for UpdateMatchRuleResponse
type UpdateMatchRuleResponse struct {
	Data MatchRule `json:"data"`
}

type _UpdateMatchRuleResponse UpdateMatchRuleResponse

// NewUpdateMatchRuleResponse instantiates a new UpdateMatchRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMatchRuleResponse(data MatchRule) *UpdateMatchRuleResponse {
	this := UpdateMatchRuleResponse{}
	this.Data = data
	return &this
}

// NewUpdateMatchRuleResponseWithDefaults instantiates a new UpdateMatchRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMatchRuleResponseWithDefaults() *UpdateMatchRuleResponse {
	this := UpdateMatchRuleResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateMatchRuleResponse) GetData() MatchRule {
	if o == nil {
		var ret MatchRule
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateMatchRuleResponse) GetDataOk() (*MatchRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateMatchRuleResponse) SetData(v MatchRule) {
	o.Data = v
}

func (o UpdateMatchRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMatchRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UpdateMatchRuleResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varUpdateMatchRuleResponse := _UpdateMatchRuleResponse{}

	err = json.Unmarshal(bytes, &varUpdateMatchRuleResponse)

	if err != nil {
		return err
	}

	*o = UpdateMatchRuleResponse(varUpdateMatchRuleResponse)

	return err
}

type NullableUpdateMatchRuleResponse struct {
	value *UpdateMatchRuleResponse
	isSet bool
}

func (v NullableUpdateMatchRuleResponse) Get() *UpdateMatchRuleResponse {
	return v.value
}

func (v *NullableUpdateMatchRuleResponse) Set(val *UpdateMatchRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMatchRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMatchRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMatchRuleResponse(val *UpdateMatchRuleResponse) *NullableUpdateMatchRuleResponse {
	return &NullableUpdateMatchRuleResponse{value: val, isSet: true}
}

func (v NullableUpdateMatchRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMatchRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


