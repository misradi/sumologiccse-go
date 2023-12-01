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

// checks if the AddEntityNoteRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddEntityNoteRequestBody{}

// AddEntityNoteRequestBody struct for AddEntityNoteRequestBody
type AddEntityNoteRequestBody struct {
	Note string `json:"note"`
}

type _AddEntityNoteRequestBody AddEntityNoteRequestBody

// NewAddEntityNoteRequestBody instantiates a new AddEntityNoteRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddEntityNoteRequestBody(note string) *AddEntityNoteRequestBody {
	this := AddEntityNoteRequestBody{}
	this.Note = note
	return &this
}

// NewAddEntityNoteRequestBodyWithDefaults instantiates a new AddEntityNoteRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddEntityNoteRequestBodyWithDefaults() *AddEntityNoteRequestBody {
	this := AddEntityNoteRequestBody{}
	return &this
}

// GetNote returns the Note field value
func (o *AddEntityNoteRequestBody) GetNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Note
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
func (o *AddEntityNoteRequestBody) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Note, true
}

// SetNote sets field value
func (o *AddEntityNoteRequestBody) SetNote(v string) {
	o.Note = v
}

func (o AddEntityNoteRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddEntityNoteRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["note"] = o.Note
	return toSerialize, nil
}

func (o *AddEntityNoteRequestBody) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"note",
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

	varAddEntityNoteRequestBody := _AddEntityNoteRequestBody{}

	err = json.Unmarshal(bytes, &varAddEntityNoteRequestBody)

	if err != nil {
		return err
	}

	*o = AddEntityNoteRequestBody(varAddEntityNoteRequestBody)

	return err
}

type NullableAddEntityNoteRequestBody struct {
	value *AddEntityNoteRequestBody
	isSet bool
}

func (v NullableAddEntityNoteRequestBody) Get() *AddEntityNoteRequestBody {
	return v.value
}

func (v *NullableAddEntityNoteRequestBody) Set(val *AddEntityNoteRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAddEntityNoteRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAddEntityNoteRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddEntityNoteRequestBody(val *AddEntityNoteRequestBody) *NullableAddEntityNoteRequestBody {
	return &NullableAddEntityNoteRequestBody{value: val, isSet: true}
}

func (v NullableAddEntityNoteRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddEntityNoteRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


