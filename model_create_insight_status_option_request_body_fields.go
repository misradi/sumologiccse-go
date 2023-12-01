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

// checks if the CreateInsightStatusOptionRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInsightStatusOptionRequestBodyFields{}

// CreateInsightStatusOptionRequestBodyFields struct for CreateInsightStatusOptionRequestBodyFields
type CreateInsightStatusOptionRequestBodyFields struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Color *string `json:"color,omitempty"`
}

type _CreateInsightStatusOptionRequestBodyFields CreateInsightStatusOptionRequestBodyFields

// NewCreateInsightStatusOptionRequestBodyFields instantiates a new CreateInsightStatusOptionRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInsightStatusOptionRequestBodyFields(name string) *CreateInsightStatusOptionRequestBodyFields {
	this := CreateInsightStatusOptionRequestBodyFields{}
	this.Name = name
	return &this
}

// NewCreateInsightStatusOptionRequestBodyFieldsWithDefaults instantiates a new CreateInsightStatusOptionRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInsightStatusOptionRequestBodyFieldsWithDefaults() *CreateInsightStatusOptionRequestBodyFields {
	this := CreateInsightStatusOptionRequestBodyFields{}
	return &this
}

// GetName returns the Name field value
func (o *CreateInsightStatusOptionRequestBodyFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateInsightStatusOptionRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateInsightStatusOptionRequestBodyFields) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateInsightStatusOptionRequestBodyFields) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInsightStatusOptionRequestBodyFields) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateInsightStatusOptionRequestBodyFields) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateInsightStatusOptionRequestBodyFields) SetDescription(v string) {
	o.Description = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *CreateInsightStatusOptionRequestBodyFields) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInsightStatusOptionRequestBodyFields) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *CreateInsightStatusOptionRequestBodyFields) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *CreateInsightStatusOptionRequestBodyFields) SetColor(v string) {
	o.Color = &v
}

func (o CreateInsightStatusOptionRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInsightStatusOptionRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	return toSerialize, nil
}

func (o *CreateInsightStatusOptionRequestBodyFields) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateInsightStatusOptionRequestBodyFields := _CreateInsightStatusOptionRequestBodyFields{}

	err = json.Unmarshal(bytes, &varCreateInsightStatusOptionRequestBodyFields)

	if err != nil {
		return err
	}

	*o = CreateInsightStatusOptionRequestBodyFields(varCreateInsightStatusOptionRequestBodyFields)

	return err
}

type NullableCreateInsightStatusOptionRequestBodyFields struct {
	value *CreateInsightStatusOptionRequestBodyFields
	isSet bool
}

func (v NullableCreateInsightStatusOptionRequestBodyFields) Get() *CreateInsightStatusOptionRequestBodyFields {
	return v.value
}

func (v *NullableCreateInsightStatusOptionRequestBodyFields) Set(val *CreateInsightStatusOptionRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInsightStatusOptionRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInsightStatusOptionRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInsightStatusOptionRequestBodyFields(val *CreateInsightStatusOptionRequestBodyFields) *NullableCreateInsightStatusOptionRequestBodyFields {
	return &NullableCreateInsightStatusOptionRequestBodyFields{value: val, isSet: true}
}

func (v NullableCreateInsightStatusOptionRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInsightStatusOptionRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


