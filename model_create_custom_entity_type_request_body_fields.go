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

// checks if the CreateCustomEntityTypeRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCustomEntityTypeRequestBodyFields{}

// CreateCustomEntityTypeRequestBodyFields struct for CreateCustomEntityTypeRequestBodyFields
type CreateCustomEntityTypeRequestBodyFields struct {
	// Human friend and unique name. Examples: \"Ip Address\", \"Username\", \"Mac Address\".
	Name string `json:"name"`
	// Record schema fields. Examples: \"file_hash_md5\", \"file_hash_sha1\".
	Fields []string `json:"fields"`
	// Machine friendly and unique identifier. Examples: \"ip\", \"username\", \"mac\".
	Identifier string `json:"identifier"`
}

type _CreateCustomEntityTypeRequestBodyFields CreateCustomEntityTypeRequestBodyFields

// NewCreateCustomEntityTypeRequestBodyFields instantiates a new CreateCustomEntityTypeRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomEntityTypeRequestBodyFields(name string, fields []string, identifier string) *CreateCustomEntityTypeRequestBodyFields {
	this := CreateCustomEntityTypeRequestBodyFields{}
	this.Name = name
	this.Fields = fields
	this.Identifier = identifier
	return &this
}

// NewCreateCustomEntityTypeRequestBodyFieldsWithDefaults instantiates a new CreateCustomEntityTypeRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomEntityTypeRequestBodyFieldsWithDefaults() *CreateCustomEntityTypeRequestBodyFields {
	this := CreateCustomEntityTypeRequestBodyFields{}
	return &this
}

// GetName returns the Name field value
func (o *CreateCustomEntityTypeRequestBodyFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateCustomEntityTypeRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateCustomEntityTypeRequestBodyFields) SetName(v string) {
	o.Name = v
}

// GetFields returns the Fields field value
func (o *CreateCustomEntityTypeRequestBodyFields) GetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *CreateCustomEntityTypeRequestBodyFields) GetFieldsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *CreateCustomEntityTypeRequestBodyFields) SetFields(v []string) {
	o.Fields = v
}

// GetIdentifier returns the Identifier field value
func (o *CreateCustomEntityTypeRequestBodyFields) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *CreateCustomEntityTypeRequestBodyFields) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *CreateCustomEntityTypeRequestBodyFields) SetIdentifier(v string) {
	o.Identifier = v
}

func (o CreateCustomEntityTypeRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustomEntityTypeRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["fields"] = o.Fields
	toSerialize["identifier"] = o.Identifier
	return toSerialize, nil
}

func (o *CreateCustomEntityTypeRequestBodyFields) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"fields",
		"identifier",
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

	varCreateCustomEntityTypeRequestBodyFields := _CreateCustomEntityTypeRequestBodyFields{}

	err = json.Unmarshal(bytes, &varCreateCustomEntityTypeRequestBodyFields)

	if err != nil {
		return err
	}

	*o = CreateCustomEntityTypeRequestBodyFields(varCreateCustomEntityTypeRequestBodyFields)

	return err
}

type NullableCreateCustomEntityTypeRequestBodyFields struct {
	value *CreateCustomEntityTypeRequestBodyFields
	isSet bool
}

func (v NullableCreateCustomEntityTypeRequestBodyFields) Get() *CreateCustomEntityTypeRequestBodyFields {
	return v.value
}

func (v *NullableCreateCustomEntityTypeRequestBodyFields) Set(val *CreateCustomEntityTypeRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomEntityTypeRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomEntityTypeRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomEntityTypeRequestBodyFields(val *CreateCustomEntityTypeRequestBodyFields) *NullableCreateCustomEntityTypeRequestBodyFields {
	return &NullableCreateCustomEntityTypeRequestBodyFields{value: val, isSet: true}
}

func (v NullableCreateCustomEntityTypeRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomEntityTypeRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

