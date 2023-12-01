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

// checks if the UpdateAggregationRuleRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAggregationRuleRequestBody{}

// UpdateAggregationRuleRequestBody struct for UpdateAggregationRuleRequestBody
type UpdateAggregationRuleRequestBody struct {
	Fields UpdateAggregationRuleRequestBodyFields `json:"fields"`
}

type _UpdateAggregationRuleRequestBody UpdateAggregationRuleRequestBody

// NewUpdateAggregationRuleRequestBody instantiates a new UpdateAggregationRuleRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAggregationRuleRequestBody(fields UpdateAggregationRuleRequestBodyFields) *UpdateAggregationRuleRequestBody {
	this := UpdateAggregationRuleRequestBody{}
	this.Fields = fields
	return &this
}

// NewUpdateAggregationRuleRequestBodyWithDefaults instantiates a new UpdateAggregationRuleRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAggregationRuleRequestBodyWithDefaults() *UpdateAggregationRuleRequestBody {
	this := UpdateAggregationRuleRequestBody{}
	return &this
}

// GetFields returns the Fields field value
func (o *UpdateAggregationRuleRequestBody) GetFields() UpdateAggregationRuleRequestBodyFields {
	if o == nil {
		var ret UpdateAggregationRuleRequestBodyFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *UpdateAggregationRuleRequestBody) GetFieldsOk() (*UpdateAggregationRuleRequestBodyFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *UpdateAggregationRuleRequestBody) SetFields(v UpdateAggregationRuleRequestBodyFields) {
	o.Fields = v
}

func (o UpdateAggregationRuleRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAggregationRuleRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *UpdateAggregationRuleRequestBody) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fields",
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

	varUpdateAggregationRuleRequestBody := _UpdateAggregationRuleRequestBody{}

	err = json.Unmarshal(bytes, &varUpdateAggregationRuleRequestBody)

	if err != nil {
		return err
	}

	*o = UpdateAggregationRuleRequestBody(varUpdateAggregationRuleRequestBody)

	return err
}

type NullableUpdateAggregationRuleRequestBody struct {
	value *UpdateAggregationRuleRequestBody
	isSet bool
}

func (v NullableUpdateAggregationRuleRequestBody) Get() *UpdateAggregationRuleRequestBody {
	return v.value
}

func (v *NullableUpdateAggregationRuleRequestBody) Set(val *UpdateAggregationRuleRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAggregationRuleRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAggregationRuleRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAggregationRuleRequestBody(val *UpdateAggregationRuleRequestBody) *NullableUpdateAggregationRuleRequestBody {
	return &NullableUpdateAggregationRuleRequestBody{value: val, isSet: true}
}

func (v NullableUpdateAggregationRuleRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAggregationRuleRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


