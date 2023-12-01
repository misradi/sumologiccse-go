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

// checks if the CreateInsightResolutionRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInsightResolutionRequestBodyFields{}

// CreateInsightResolutionRequestBodyFields struct for CreateInsightResolutionRequestBodyFields
type CreateInsightResolutionRequestBodyFields struct {
	Name string `json:"name"`
	ParentId *int32 `json:"parentId,omitempty"`
	Description *string `json:"description,omitempty"`
}

type _CreateInsightResolutionRequestBodyFields CreateInsightResolutionRequestBodyFields

// NewCreateInsightResolutionRequestBodyFields instantiates a new CreateInsightResolutionRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInsightResolutionRequestBodyFields(name string) *CreateInsightResolutionRequestBodyFields {
	this := CreateInsightResolutionRequestBodyFields{}
	this.Name = name
	return &this
}

// NewCreateInsightResolutionRequestBodyFieldsWithDefaults instantiates a new CreateInsightResolutionRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInsightResolutionRequestBodyFieldsWithDefaults() *CreateInsightResolutionRequestBodyFields {
	this := CreateInsightResolutionRequestBodyFields{}
	return &this
}

// GetName returns the Name field value
func (o *CreateInsightResolutionRequestBodyFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateInsightResolutionRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateInsightResolutionRequestBodyFields) SetName(v string) {
	o.Name = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *CreateInsightResolutionRequestBodyFields) GetParentId() int32 {
	if o == nil || IsNil(o.ParentId) {
		var ret int32
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInsightResolutionRequestBodyFields) GetParentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *CreateInsightResolutionRequestBodyFields) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given int32 and assigns it to the ParentId field.
func (o *CreateInsightResolutionRequestBodyFields) SetParentId(v int32) {
	o.ParentId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateInsightResolutionRequestBodyFields) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInsightResolutionRequestBodyFields) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateInsightResolutionRequestBodyFields) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateInsightResolutionRequestBodyFields) SetDescription(v string) {
	o.Description = &v
}

func (o CreateInsightResolutionRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInsightResolutionRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

func (o *CreateInsightResolutionRequestBodyFields) UnmarshalJSON(bytes []byte) (err error) {
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

	varCreateInsightResolutionRequestBodyFields := _CreateInsightResolutionRequestBodyFields{}

	err = json.Unmarshal(bytes, &varCreateInsightResolutionRequestBodyFields)

	if err != nil {
		return err
	}

	*o = CreateInsightResolutionRequestBodyFields(varCreateInsightResolutionRequestBodyFields)

	return err
}

type NullableCreateInsightResolutionRequestBodyFields struct {
	value *CreateInsightResolutionRequestBodyFields
	isSet bool
}

func (v NullableCreateInsightResolutionRequestBodyFields) Get() *CreateInsightResolutionRequestBodyFields {
	return v.value
}

func (v *NullableCreateInsightResolutionRequestBodyFields) Set(val *CreateInsightResolutionRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInsightResolutionRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInsightResolutionRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInsightResolutionRequestBodyFields(val *CreateInsightResolutionRequestBodyFields) *NullableCreateInsightResolutionRequestBodyFields {
	return &NullableCreateInsightResolutionRequestBodyFields{value: val, isSet: true}
}

func (v NullableCreateInsightResolutionRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInsightResolutionRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


