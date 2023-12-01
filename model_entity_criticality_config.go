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

// checks if the EntityCriticalityConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityCriticalityConfig{}

// EntityCriticalityConfig struct for EntityCriticalityConfig
type EntityCriticalityConfig struct {
	Id string `json:"id"`
	// Human friendly and unique name. Examples: \"Executive Laptop\", \"Bastion Host\"
	Name string `json:"name"`
	// Algebraic expression representing this entity's criticality. Examples: \"severity * 2\", \"severity - 5\", \"severity / 3\"
	SeverityExpression string `json:"severityExpression"`
	// Number of entities related to this criticality.
	EntityCount int32 `json:"entityCount"`
}

type _EntityCriticalityConfig EntityCriticalityConfig

// NewEntityCriticalityConfig instantiates a new EntityCriticalityConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityCriticalityConfig(id string, name string, severityExpression string, entityCount int32) *EntityCriticalityConfig {
	this := EntityCriticalityConfig{}
	this.Id = id
	this.Name = name
	this.SeverityExpression = severityExpression
	this.EntityCount = entityCount
	return &this
}

// NewEntityCriticalityConfigWithDefaults instantiates a new EntityCriticalityConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityCriticalityConfigWithDefaults() *EntityCriticalityConfig {
	this := EntityCriticalityConfig{}
	return &this
}

// GetId returns the Id field value
func (o *EntityCriticalityConfig) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntityCriticalityConfig) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntityCriticalityConfig) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EntityCriticalityConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntityCriticalityConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntityCriticalityConfig) SetName(v string) {
	o.Name = v
}

// GetSeverityExpression returns the SeverityExpression field value
func (o *EntityCriticalityConfig) GetSeverityExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeverityExpression
}

// GetSeverityExpressionOk returns a tuple with the SeverityExpression field value
// and a boolean to check if the value has been set.
func (o *EntityCriticalityConfig) GetSeverityExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeverityExpression, true
}

// SetSeverityExpression sets field value
func (o *EntityCriticalityConfig) SetSeverityExpression(v string) {
	o.SeverityExpression = v
}

// GetEntityCount returns the EntityCount field value
func (o *EntityCriticalityConfig) GetEntityCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EntityCount
}

// GetEntityCountOk returns a tuple with the EntityCount field value
// and a boolean to check if the value has been set.
func (o *EntityCriticalityConfig) GetEntityCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityCount, true
}

// SetEntityCount sets field value
func (o *EntityCriticalityConfig) SetEntityCount(v int32) {
	o.EntityCount = v
}

func (o EntityCriticalityConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityCriticalityConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["severityExpression"] = o.SeverityExpression
	toSerialize["entityCount"] = o.EntityCount
	return toSerialize, nil
}

func (o *EntityCriticalityConfig) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"severityExpression",
		"entityCount",
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

	varEntityCriticalityConfig := _EntityCriticalityConfig{}

	err = json.Unmarshal(bytes, &varEntityCriticalityConfig)

	if err != nil {
		return err
	}

	*o = EntityCriticalityConfig(varEntityCriticalityConfig)

	return err
}

type NullableEntityCriticalityConfig struct {
	value *EntityCriticalityConfig
	isSet bool
}

func (v NullableEntityCriticalityConfig) Get() *EntityCriticalityConfig {
	return v.value
}

func (v *NullableEntityCriticalityConfig) Set(val *EntityCriticalityConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityCriticalityConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityCriticalityConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityCriticalityConfig(val *EntityCriticalityConfig) *NullableEntityCriticalityConfig {
	return &NullableEntityCriticalityConfig{value: val, isSet: true}
}

func (v NullableEntityCriticalityConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityCriticalityConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


