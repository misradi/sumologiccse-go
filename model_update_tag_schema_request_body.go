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

// checks if the UpdateTagSchemaRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTagSchemaRequestBody{}

// UpdateTagSchemaRequestBody struct for UpdateTagSchemaRequestBody
type UpdateTagSchemaRequestBody struct {
	Fields UpdateTagSchemaRequestBodyFields `json:"fields"`
}

type _UpdateTagSchemaRequestBody UpdateTagSchemaRequestBody

// NewUpdateTagSchemaRequestBody instantiates a new UpdateTagSchemaRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTagSchemaRequestBody(fields UpdateTagSchemaRequestBodyFields) *UpdateTagSchemaRequestBody {
	this := UpdateTagSchemaRequestBody{}
	this.Fields = fields
	return &this
}

// NewUpdateTagSchemaRequestBodyWithDefaults instantiates a new UpdateTagSchemaRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTagSchemaRequestBodyWithDefaults() *UpdateTagSchemaRequestBody {
	this := UpdateTagSchemaRequestBody{}
	return &this
}

// GetFields returns the Fields field value
func (o *UpdateTagSchemaRequestBody) GetFields() UpdateTagSchemaRequestBodyFields {
	if o == nil {
		var ret UpdateTagSchemaRequestBodyFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *UpdateTagSchemaRequestBody) GetFieldsOk() (*UpdateTagSchemaRequestBodyFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *UpdateTagSchemaRequestBody) SetFields(v UpdateTagSchemaRequestBodyFields) {
	o.Fields = v
}

func (o UpdateTagSchemaRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTagSchemaRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *UpdateTagSchemaRequestBody) UnmarshalJSON(bytes []byte) (err error) {
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

	varUpdateTagSchemaRequestBody := _UpdateTagSchemaRequestBody{}

	err = json.Unmarshal(bytes, &varUpdateTagSchemaRequestBody)

	if err != nil {
		return err
	}

	*o = UpdateTagSchemaRequestBody(varUpdateTagSchemaRequestBody)

	return err
}

type NullableUpdateTagSchemaRequestBody struct {
	value *UpdateTagSchemaRequestBody
	isSet bool
}

func (v NullableUpdateTagSchemaRequestBody) Get() *UpdateTagSchemaRequestBody {
	return v.value
}

func (v *NullableUpdateTagSchemaRequestBody) Set(val *UpdateTagSchemaRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTagSchemaRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTagSchemaRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTagSchemaRequestBody(val *UpdateTagSchemaRequestBody) *NullableUpdateTagSchemaRequestBody {
	return &NullableUpdateTagSchemaRequestBody{value: val, isSet: true}
}

func (v NullableUpdateTagSchemaRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTagSchemaRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


