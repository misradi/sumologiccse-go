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

// checks if the ContextAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextAction{}

// ContextAction struct for ContextAction
type ContextAction struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Template string `json:"template"`
	IocTypes []string `json:"iocTypes"`
	EntityTypes []string `json:"entityTypes"`
	RecordFields []string `json:"recordFields"`
	AllRecordFields bool `json:"allRecordFields"`
	Enabled bool `json:"enabled"`
}

type _ContextAction ContextAction

// NewContextAction instantiates a new ContextAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextAction(id string, name string, type_ string, template string, iocTypes []string, entityTypes []string, recordFields []string, allRecordFields bool, enabled bool) *ContextAction {
	this := ContextAction{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Template = template
	this.IocTypes = iocTypes
	this.EntityTypes = entityTypes
	this.RecordFields = recordFields
	this.AllRecordFields = allRecordFields
	this.Enabled = enabled
	return &this
}

// NewContextActionWithDefaults instantiates a new ContextAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextActionWithDefaults() *ContextAction {
	this := ContextAction{}
	return &this
}

// GetId returns the Id field value
func (o *ContextAction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContextAction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContextAction) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ContextAction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContextAction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContextAction) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ContextAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContextAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContextAction) SetType(v string) {
	o.Type = v
}

// GetTemplate returns the Template field value
func (o *ContextAction) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *ContextAction) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *ContextAction) SetTemplate(v string) {
	o.Template = v
}

// GetIocTypes returns the IocTypes field value
func (o *ContextAction) GetIocTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IocTypes
}

// GetIocTypesOk returns a tuple with the IocTypes field value
// and a boolean to check if the value has been set.
func (o *ContextAction) GetIocTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IocTypes, true
}

// SetIocTypes sets field value
func (o *ContextAction) SetIocTypes(v []string) {
	o.IocTypes = v
}

// GetEntityTypes returns the EntityTypes field value
func (o *ContextAction) GetEntityTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EntityTypes
}

// GetEntityTypesOk returns a tuple with the EntityTypes field value
// and a boolean to check if the value has been set.
func (o *ContextAction) GetEntityTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntityTypes, true
}

// SetEntityTypes sets field value
func (o *ContextAction) SetEntityTypes(v []string) {
	o.EntityTypes = v
}

// GetRecordFields returns the RecordFields field value
func (o *ContextAction) GetRecordFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RecordFields
}

// GetRecordFieldsOk returns a tuple with the RecordFields field value
// and a boolean to check if the value has been set.
func (o *ContextAction) GetRecordFieldsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecordFields, true
}

// SetRecordFields sets field value
func (o *ContextAction) SetRecordFields(v []string) {
	o.RecordFields = v
}

// GetAllRecordFields returns the AllRecordFields field value
func (o *ContextAction) GetAllRecordFields() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllRecordFields
}

// GetAllRecordFieldsOk returns a tuple with the AllRecordFields field value
// and a boolean to check if the value has been set.
func (o *ContextAction) GetAllRecordFieldsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllRecordFields, true
}

// SetAllRecordFields sets field value
func (o *ContextAction) SetAllRecordFields(v bool) {
	o.AllRecordFields = v
}

// GetEnabled returns the Enabled field value
func (o *ContextAction) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ContextAction) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ContextAction) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ContextAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["template"] = o.Template
	toSerialize["iocTypes"] = o.IocTypes
	toSerialize["entityTypes"] = o.EntityTypes
	toSerialize["recordFields"] = o.RecordFields
	toSerialize["allRecordFields"] = o.AllRecordFields
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

func (o *ContextAction) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"type",
		"template",
		"iocTypes",
		"entityTypes",
		"recordFields",
		"allRecordFields",
		"enabled",
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

	varContextAction := _ContextAction{}

	err = json.Unmarshal(bytes, &varContextAction)

	if err != nil {
		return err
	}

	*o = ContextAction(varContextAction)

	return err
}

type NullableContextAction struct {
	value *ContextAction
	isSet bool
}

func (v NullableContextAction) Get() *ContextAction {
	return v.value
}

func (v *NullableContextAction) Set(val *ContextAction) {
	v.value = val
	v.isSet = true
}

func (v NullableContextAction) IsSet() bool {
	return v.isSet
}

func (v *NullableContextAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextAction(val *ContextAction) *NullableContextAction {
	return &NullableContextAction{value: val, isSet: true}
}

func (v NullableContextAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


