/*
Sumo Logic CSE API

 https://help.sumologic.com/APIs 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologiccse

import (
	"encoding/json"
)

// checks if the GetRulesResponseDataObjectsInnerOneOfNameOverride type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRulesResponseDataObjectsInnerOneOfNameOverride{}

// GetRulesResponseDataObjectsInnerOneOfNameOverride struct for GetRulesResponseDataObjectsInnerOneOfNameOverride
type GetRulesResponseDataObjectsInnerOneOfNameOverride struct {
	Original *string `json:"original,omitempty"`
	Override *string `json:"override,omitempty"`
}

// NewGetRulesResponseDataObjectsInnerOneOfNameOverride instantiates a new GetRulesResponseDataObjectsInnerOneOfNameOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRulesResponseDataObjectsInnerOneOfNameOverride() *GetRulesResponseDataObjectsInnerOneOfNameOverride {
	this := GetRulesResponseDataObjectsInnerOneOfNameOverride{}
	return &this
}

// NewGetRulesResponseDataObjectsInnerOneOfNameOverrideWithDefaults instantiates a new GetRulesResponseDataObjectsInnerOneOfNameOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRulesResponseDataObjectsInnerOneOfNameOverrideWithDefaults() *GetRulesResponseDataObjectsInnerOneOfNameOverride {
	this := GetRulesResponseDataObjectsInnerOneOfNameOverride{}
	return &this
}

// GetOriginal returns the Original field value if set, zero value otherwise.
func (o *GetRulesResponseDataObjectsInnerOneOfNameOverride) GetOriginal() string {
	if o == nil || IsNil(o.Original) {
		var ret string
		return ret
	}
	return *o.Original
}

// GetOriginalOk returns a tuple with the Original field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRulesResponseDataObjectsInnerOneOfNameOverride) GetOriginalOk() (*string, bool) {
	if o == nil || IsNil(o.Original) {
		return nil, false
	}
	return o.Original, true
}

// HasOriginal returns a boolean if a field has been set.
func (o *GetRulesResponseDataObjectsInnerOneOfNameOverride) HasOriginal() bool {
	if o != nil && !IsNil(o.Original) {
		return true
	}

	return false
}

// SetOriginal gets a reference to the given string and assigns it to the Original field.
func (o *GetRulesResponseDataObjectsInnerOneOfNameOverride) SetOriginal(v string) {
	o.Original = &v
}

// GetOverride returns the Override field value if set, zero value otherwise.
func (o *GetRulesResponseDataObjectsInnerOneOfNameOverride) GetOverride() string {
	if o == nil || IsNil(o.Override) {
		var ret string
		return ret
	}
	return *o.Override
}

// GetOverrideOk returns a tuple with the Override field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRulesResponseDataObjectsInnerOneOfNameOverride) GetOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.Override) {
		return nil, false
	}
	return o.Override, true
}

// HasOverride returns a boolean if a field has been set.
func (o *GetRulesResponseDataObjectsInnerOneOfNameOverride) HasOverride() bool {
	if o != nil && !IsNil(o.Override) {
		return true
	}

	return false
}

// SetOverride gets a reference to the given string and assigns it to the Override field.
func (o *GetRulesResponseDataObjectsInnerOneOfNameOverride) SetOverride(v string) {
	o.Override = &v
}

func (o GetRulesResponseDataObjectsInnerOneOfNameOverride) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRulesResponseDataObjectsInnerOneOfNameOverride) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Original) {
		toSerialize["original"] = o.Original
	}
	if !IsNil(o.Override) {
		toSerialize["override"] = o.Override
	}
	return toSerialize, nil
}

type NullableGetRulesResponseDataObjectsInnerOneOfNameOverride struct {
	value *GetRulesResponseDataObjectsInnerOneOfNameOverride
	isSet bool
}

func (v NullableGetRulesResponseDataObjectsInnerOneOfNameOverride) Get() *GetRulesResponseDataObjectsInnerOneOfNameOverride {
	return v.value
}

func (v *NullableGetRulesResponseDataObjectsInnerOneOfNameOverride) Set(val *GetRulesResponseDataObjectsInnerOneOfNameOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRulesResponseDataObjectsInnerOneOfNameOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRulesResponseDataObjectsInnerOneOfNameOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRulesResponseDataObjectsInnerOneOfNameOverride(val *GetRulesResponseDataObjectsInnerOneOfNameOverride) *NullableGetRulesResponseDataObjectsInnerOneOfNameOverride {
	return &NullableGetRulesResponseDataObjectsInnerOneOfNameOverride{value: val, isSet: true}
}

func (v NullableGetRulesResponseDataObjectsInnerOneOfNameOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRulesResponseDataObjectsInnerOneOfNameOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


