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

// checks if the UpdateContextActionRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateContextActionRequestBodyFields{}

// UpdateContextActionRequestBodyFields struct for UpdateContextActionRequestBodyFields
type UpdateContextActionRequestBodyFields struct {
	Name string `json:"name"`
	Type *string `json:"type,omitempty"`
	UrlTemplate *string `json:"urlTemplate,omitempty"`
	Template *string `json:"template,omitempty"`
	IocTypes []string `json:"iocTypes"`
	EntityTypes []string `json:"entityTypes,omitempty"`
	RecordFields []string `json:"recordFields,omitempty"`
	AllRecordFields *bool `json:"allRecordFields,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

type _UpdateContextActionRequestBodyFields UpdateContextActionRequestBodyFields

// NewUpdateContextActionRequestBodyFields instantiates a new UpdateContextActionRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateContextActionRequestBodyFields(name string, iocTypes []string) *UpdateContextActionRequestBodyFields {
	this := UpdateContextActionRequestBodyFields{}
	this.Name = name
	this.IocTypes = iocTypes
	return &this
}

// NewUpdateContextActionRequestBodyFieldsWithDefaults instantiates a new UpdateContextActionRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateContextActionRequestBodyFieldsWithDefaults() *UpdateContextActionRequestBodyFields {
	this := UpdateContextActionRequestBodyFields{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateContextActionRequestBodyFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateContextActionRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateContextActionRequestBodyFields) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateContextActionRequestBodyFields) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateContextActionRequestBodyFields) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateContextActionRequestBodyFields) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateContextActionRequestBodyFields) SetType(v string) {
	o.Type = &v
}

// GetUrlTemplate returns the UrlTemplate field value if set, zero value otherwise.
func (o *UpdateContextActionRequestBodyFields) GetUrlTemplate() string {
	if o == nil || IsNil(o.UrlTemplate) {
		var ret string
		return ret
	}
	return *o.UrlTemplate
}

// GetUrlTemplateOk returns a tuple with the UrlTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateContextActionRequestBodyFields) GetUrlTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.UrlTemplate) {
		return nil, false
	}
	return o.UrlTemplate, true
}

// HasUrlTemplate returns a boolean if a field has been set.
func (o *UpdateContextActionRequestBodyFields) HasUrlTemplate() bool {
	if o != nil && !IsNil(o.UrlTemplate) {
		return true
	}

	return false
}

// SetUrlTemplate gets a reference to the given string and assigns it to the UrlTemplate field.
func (o *UpdateContextActionRequestBodyFields) SetUrlTemplate(v string) {
	o.UrlTemplate = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *UpdateContextActionRequestBodyFields) GetTemplate() string {
	if o == nil || IsNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateContextActionRequestBodyFields) GetTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *UpdateContextActionRequestBodyFields) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *UpdateContextActionRequestBodyFields) SetTemplate(v string) {
	o.Template = &v
}

// GetIocTypes returns the IocTypes field value
func (o *UpdateContextActionRequestBodyFields) GetIocTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IocTypes
}

// GetIocTypesOk returns a tuple with the IocTypes field value
// and a boolean to check if the value has been set.
func (o *UpdateContextActionRequestBodyFields) GetIocTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IocTypes, true
}

// SetIocTypes sets field value
func (o *UpdateContextActionRequestBodyFields) SetIocTypes(v []string) {
	o.IocTypes = v
}

// GetEntityTypes returns the EntityTypes field value if set, zero value otherwise.
func (o *UpdateContextActionRequestBodyFields) GetEntityTypes() []string {
	if o == nil || IsNil(o.EntityTypes) {
		var ret []string
		return ret
	}
	return o.EntityTypes
}

// GetEntityTypesOk returns a tuple with the EntityTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateContextActionRequestBodyFields) GetEntityTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.EntityTypes) {
		return nil, false
	}
	return o.EntityTypes, true
}

// HasEntityTypes returns a boolean if a field has been set.
func (o *UpdateContextActionRequestBodyFields) HasEntityTypes() bool {
	if o != nil && !IsNil(o.EntityTypes) {
		return true
	}

	return false
}

// SetEntityTypes gets a reference to the given []string and assigns it to the EntityTypes field.
func (o *UpdateContextActionRequestBodyFields) SetEntityTypes(v []string) {
	o.EntityTypes = v
}

// GetRecordFields returns the RecordFields field value if set, zero value otherwise.
func (o *UpdateContextActionRequestBodyFields) GetRecordFields() []string {
	if o == nil || IsNil(o.RecordFields) {
		var ret []string
		return ret
	}
	return o.RecordFields
}

// GetRecordFieldsOk returns a tuple with the RecordFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateContextActionRequestBodyFields) GetRecordFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.RecordFields) {
		return nil, false
	}
	return o.RecordFields, true
}

// HasRecordFields returns a boolean if a field has been set.
func (o *UpdateContextActionRequestBodyFields) HasRecordFields() bool {
	if o != nil && !IsNil(o.RecordFields) {
		return true
	}

	return false
}

// SetRecordFields gets a reference to the given []string and assigns it to the RecordFields field.
func (o *UpdateContextActionRequestBodyFields) SetRecordFields(v []string) {
	o.RecordFields = v
}

// GetAllRecordFields returns the AllRecordFields field value if set, zero value otherwise.
func (o *UpdateContextActionRequestBodyFields) GetAllRecordFields() bool {
	if o == nil || IsNil(o.AllRecordFields) {
		var ret bool
		return ret
	}
	return *o.AllRecordFields
}

// GetAllRecordFieldsOk returns a tuple with the AllRecordFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateContextActionRequestBodyFields) GetAllRecordFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllRecordFields) {
		return nil, false
	}
	return o.AllRecordFields, true
}

// HasAllRecordFields returns a boolean if a field has been set.
func (o *UpdateContextActionRequestBodyFields) HasAllRecordFields() bool {
	if o != nil && !IsNil(o.AllRecordFields) {
		return true
	}

	return false
}

// SetAllRecordFields gets a reference to the given bool and assigns it to the AllRecordFields field.
func (o *UpdateContextActionRequestBodyFields) SetAllRecordFields(v bool) {
	o.AllRecordFields = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateContextActionRequestBodyFields) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateContextActionRequestBodyFields) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateContextActionRequestBodyFields) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateContextActionRequestBodyFields) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateContextActionRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateContextActionRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UrlTemplate) {
		toSerialize["urlTemplate"] = o.UrlTemplate
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	toSerialize["iocTypes"] = o.IocTypes
	if !IsNil(o.EntityTypes) {
		toSerialize["entityTypes"] = o.EntityTypes
	}
	if !IsNil(o.RecordFields) {
		toSerialize["recordFields"] = o.RecordFields
	}
	if !IsNil(o.AllRecordFields) {
		toSerialize["allRecordFields"] = o.AllRecordFields
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

func (o *UpdateContextActionRequestBodyFields) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"iocTypes",
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

	varUpdateContextActionRequestBodyFields := _UpdateContextActionRequestBodyFields{}

	err = json.Unmarshal(bytes, &varUpdateContextActionRequestBodyFields)

	if err != nil {
		return err
	}

	*o = UpdateContextActionRequestBodyFields(varUpdateContextActionRequestBodyFields)

	return err
}

type NullableUpdateContextActionRequestBodyFields struct {
	value *UpdateContextActionRequestBodyFields
	isSet bool
}

func (v NullableUpdateContextActionRequestBodyFields) Get() *UpdateContextActionRequestBodyFields {
	return v.value
}

func (v *NullableUpdateContextActionRequestBodyFields) Set(val *UpdateContextActionRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateContextActionRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateContextActionRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateContextActionRequestBodyFields(val *UpdateContextActionRequestBodyFields) *NullableUpdateContextActionRequestBodyFields {
	return &NullableUpdateContextActionRequestBodyFields{value: val, isSet: true}
}

func (v NullableUpdateContextActionRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateContextActionRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

