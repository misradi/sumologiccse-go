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

// checks if the CreateCustomerSourcedEntityLookupTableRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCustomerSourcedEntityLookupTableRequestBodyFields{}

// CreateCustomerSourcedEntityLookupTableRequestBodyFields struct for CreateCustomerSourcedEntityLookupTableRequestBodyFields
type CreateCustomerSourcedEntityLookupTableRequestBodyFields struct {
	Type string `json:"type"`
	KeyColumnName string `json:"keyColumnName"`
	ValueColumnName string `json:"valueColumnName"`
	TablePath string `json:"tablePath"`
	SourceCategory *string `json:"sourceCategory,omitempty"`
}

type _CreateCustomerSourcedEntityLookupTableRequestBodyFields CreateCustomerSourcedEntityLookupTableRequestBodyFields

// NewCreateCustomerSourcedEntityLookupTableRequestBodyFields instantiates a new CreateCustomerSourcedEntityLookupTableRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomerSourcedEntityLookupTableRequestBodyFields(type_ string, keyColumnName string, valueColumnName string, tablePath string) *CreateCustomerSourcedEntityLookupTableRequestBodyFields {
	this := CreateCustomerSourcedEntityLookupTableRequestBodyFields{}
	this.Type = type_
	this.KeyColumnName = keyColumnName
	this.ValueColumnName = valueColumnName
	this.TablePath = tablePath
	return &this
}

// NewCreateCustomerSourcedEntityLookupTableRequestBodyFieldsWithDefaults instantiates a new CreateCustomerSourcedEntityLookupTableRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomerSourcedEntityLookupTableRequestBodyFieldsWithDefaults() *CreateCustomerSourcedEntityLookupTableRequestBodyFields {
	this := CreateCustomerSourcedEntityLookupTableRequestBodyFields{}
	return &this
}

// GetType returns the Type field value
func (o *CreateCustomerSourcedEntityLookupTableRequestBodyFields) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerSourcedEntityLookupTableRequestBodyFields) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateCustomerSourcedEntityLookupTableRequestBodyFields) SetType(v string) {
	o.Type = v
}

// GetKeyColumnName returns the KeyColumnName field value
func (o *CreateCustomerSourcedEntityLookupTableRequestBodyFields) GetKeyColumnName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyColumnName
}

// GetKeyColumnNameOk returns a tuple with the KeyColumnName field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerSourcedEntityLookupTableRequestBodyFields) GetKeyColumnNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyColumnName, true
}

// SetKeyColumnName sets field value
func (o *CreateCustomerSourcedEntityLookupTableRequestBodyFields) SetKeyColumnName(v string) {
	o.KeyColumnName = v
}

// GetValueColumnName returns the ValueColumnName field value
func (o *CreateCustomerSourcedEntityLookupTableRequestBodyFields) GetValueColumnName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValueColumnName
}

// GetValueColumnNameOk returns a tuple with the ValueColumnName field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerSourcedEntityLookupTableRequestBodyFields) GetValueColumnNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueColumnName, true
}

// SetValueColumnName sets field value
func (o *CreateCustomerSourcedEntityLookupTableRequestBodyFields) SetValueColumnName(v string) {
	o.ValueColumnName = v
}

// GetTablePath returns the TablePath field value
func (o *CreateCustomerSourcedEntityLookupTableRequestBodyFields) GetTablePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TablePath
}

// GetTablePathOk returns a tuple with the TablePath field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerSourcedEntityLookupTableRequestBodyFields) GetTablePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TablePath, true
}

// SetTablePath sets field value
func (o *CreateCustomerSourcedEntityLookupTableRequestBodyFields) SetTablePath(v string) {
	o.TablePath = v
}

// GetSourceCategory returns the SourceCategory field value if set, zero value otherwise.
func (o *CreateCustomerSourcedEntityLookupTableRequestBodyFields) GetSourceCategory() string {
	if o == nil || IsNil(o.SourceCategory) {
		var ret string
		return ret
	}
	return *o.SourceCategory
}

// GetSourceCategoryOk returns a tuple with the SourceCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerSourcedEntityLookupTableRequestBodyFields) GetSourceCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.SourceCategory) {
		return nil, false
	}
	return o.SourceCategory, true
}

// HasSourceCategory returns a boolean if a field has been set.
func (o *CreateCustomerSourcedEntityLookupTableRequestBodyFields) HasSourceCategory() bool {
	if o != nil && !IsNil(o.SourceCategory) {
		return true
	}

	return false
}

// SetSourceCategory gets a reference to the given string and assigns it to the SourceCategory field.
func (o *CreateCustomerSourcedEntityLookupTableRequestBodyFields) SetSourceCategory(v string) {
	o.SourceCategory = &v
}

func (o CreateCustomerSourcedEntityLookupTableRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustomerSourcedEntityLookupTableRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["keyColumnName"] = o.KeyColumnName
	toSerialize["valueColumnName"] = o.ValueColumnName
	toSerialize["tablePath"] = o.TablePath
	if !IsNil(o.SourceCategory) {
		toSerialize["sourceCategory"] = o.SourceCategory
	}
	return toSerialize, nil
}

func (o *CreateCustomerSourcedEntityLookupTableRequestBodyFields) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"keyColumnName",
		"valueColumnName",
		"tablePath",
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

	varCreateCustomerSourcedEntityLookupTableRequestBodyFields := _CreateCustomerSourcedEntityLookupTableRequestBodyFields{}

	err = json.Unmarshal(bytes, &varCreateCustomerSourcedEntityLookupTableRequestBodyFields)

	if err != nil {
		return err
	}

	*o = CreateCustomerSourcedEntityLookupTableRequestBodyFields(varCreateCustomerSourcedEntityLookupTableRequestBodyFields)

	return err
}

type NullableCreateCustomerSourcedEntityLookupTableRequestBodyFields struct {
	value *CreateCustomerSourcedEntityLookupTableRequestBodyFields
	isSet bool
}

func (v NullableCreateCustomerSourcedEntityLookupTableRequestBodyFields) Get() *CreateCustomerSourcedEntityLookupTableRequestBodyFields {
	return v.value
}

func (v *NullableCreateCustomerSourcedEntityLookupTableRequestBodyFields) Set(val *CreateCustomerSourcedEntityLookupTableRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomerSourcedEntityLookupTableRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomerSourcedEntityLookupTableRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomerSourcedEntityLookupTableRequestBodyFields(val *CreateCustomerSourcedEntityLookupTableRequestBodyFields) *NullableCreateCustomerSourcedEntityLookupTableRequestBodyFields {
	return &NullableCreateCustomerSourcedEntityLookupTableRequestBodyFields{value: val, isSet: true}
}

func (v NullableCreateCustomerSourcedEntityLookupTableRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomerSourcedEntityLookupTableRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

