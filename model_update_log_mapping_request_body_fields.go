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

// checks if the UpdateLogMappingRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateLogMappingRequestBodyFields{}

// UpdateLogMappingRequestBodyFields struct for UpdateLogMappingRequestBodyFields
type UpdateLogMappingRequestBodyFields struct {
	Name string `json:"name"`
	ParentId *string `json:"parentId,omitempty"`
	Fields []CreateLogMappingRequestBodyFieldsFieldsInner `json:"fields"`
	SkippedValues []string `json:"skippedValues,omitempty"`
	UnstructuredFields *CreateLogMappingRequestBodyFieldsUnstructuredFields `json:"unstructuredFields,omitempty"`
	StructuredFields *CreateLogMappingRequestBodyFieldsStructuredFields `json:"structuredFields,omitempty"`
	StructuredInputs []CreateLogMappingRequestBodyFieldsStructuredFields `json:"structuredInputs,omitempty"`
	RecordType string `json:"recordType"`
	ProductGuid string `json:"productGuid"`
	RelatesEntities *bool `json:"relatesEntities,omitempty"`
}

type _UpdateLogMappingRequestBodyFields UpdateLogMappingRequestBodyFields

// NewUpdateLogMappingRequestBodyFields instantiates a new UpdateLogMappingRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLogMappingRequestBodyFields(name string, fields []CreateLogMappingRequestBodyFieldsFieldsInner, recordType string, productGuid string) *UpdateLogMappingRequestBodyFields {
	this := UpdateLogMappingRequestBodyFields{}
	this.Name = name
	this.Fields = fields
	this.RecordType = recordType
	this.ProductGuid = productGuid
	return &this
}

// NewUpdateLogMappingRequestBodyFieldsWithDefaults instantiates a new UpdateLogMappingRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLogMappingRequestBodyFieldsWithDefaults() *UpdateLogMappingRequestBodyFields {
	this := UpdateLogMappingRequestBodyFields{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateLogMappingRequestBodyFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateLogMappingRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateLogMappingRequestBodyFields) SetName(v string) {
	o.Name = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *UpdateLogMappingRequestBodyFields) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLogMappingRequestBodyFields) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *UpdateLogMappingRequestBodyFields) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *UpdateLogMappingRequestBodyFields) SetParentId(v string) {
	o.ParentId = &v
}

// GetFields returns the Fields field value
func (o *UpdateLogMappingRequestBodyFields) GetFields() []CreateLogMappingRequestBodyFieldsFieldsInner {
	if o == nil {
		var ret []CreateLogMappingRequestBodyFieldsFieldsInner
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *UpdateLogMappingRequestBodyFields) GetFieldsOk() ([]CreateLogMappingRequestBodyFieldsFieldsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *UpdateLogMappingRequestBodyFields) SetFields(v []CreateLogMappingRequestBodyFieldsFieldsInner) {
	o.Fields = v
}

// GetSkippedValues returns the SkippedValues field value if set, zero value otherwise.
func (o *UpdateLogMappingRequestBodyFields) GetSkippedValues() []string {
	if o == nil || IsNil(o.SkippedValues) {
		var ret []string
		return ret
	}
	return o.SkippedValues
}

// GetSkippedValuesOk returns a tuple with the SkippedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLogMappingRequestBodyFields) GetSkippedValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.SkippedValues) {
		return nil, false
	}
	return o.SkippedValues, true
}

// HasSkippedValues returns a boolean if a field has been set.
func (o *UpdateLogMappingRequestBodyFields) HasSkippedValues() bool {
	if o != nil && !IsNil(o.SkippedValues) {
		return true
	}

	return false
}

// SetSkippedValues gets a reference to the given []string and assigns it to the SkippedValues field.
func (o *UpdateLogMappingRequestBodyFields) SetSkippedValues(v []string) {
	o.SkippedValues = v
}

// GetUnstructuredFields returns the UnstructuredFields field value if set, zero value otherwise.
func (o *UpdateLogMappingRequestBodyFields) GetUnstructuredFields() CreateLogMappingRequestBodyFieldsUnstructuredFields {
	if o == nil || IsNil(o.UnstructuredFields) {
		var ret CreateLogMappingRequestBodyFieldsUnstructuredFields
		return ret
	}
	return *o.UnstructuredFields
}

// GetUnstructuredFieldsOk returns a tuple with the UnstructuredFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLogMappingRequestBodyFields) GetUnstructuredFieldsOk() (*CreateLogMappingRequestBodyFieldsUnstructuredFields, bool) {
	if o == nil || IsNil(o.UnstructuredFields) {
		return nil, false
	}
	return o.UnstructuredFields, true
}

// HasUnstructuredFields returns a boolean if a field has been set.
func (o *UpdateLogMappingRequestBodyFields) HasUnstructuredFields() bool {
	if o != nil && !IsNil(o.UnstructuredFields) {
		return true
	}

	return false
}

// SetUnstructuredFields gets a reference to the given CreateLogMappingRequestBodyFieldsUnstructuredFields and assigns it to the UnstructuredFields field.
func (o *UpdateLogMappingRequestBodyFields) SetUnstructuredFields(v CreateLogMappingRequestBodyFieldsUnstructuredFields) {
	o.UnstructuredFields = &v
}

// GetStructuredFields returns the StructuredFields field value if set, zero value otherwise.
func (o *UpdateLogMappingRequestBodyFields) GetStructuredFields() CreateLogMappingRequestBodyFieldsStructuredFields {
	if o == nil || IsNil(o.StructuredFields) {
		var ret CreateLogMappingRequestBodyFieldsStructuredFields
		return ret
	}
	return *o.StructuredFields
}

// GetStructuredFieldsOk returns a tuple with the StructuredFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLogMappingRequestBodyFields) GetStructuredFieldsOk() (*CreateLogMappingRequestBodyFieldsStructuredFields, bool) {
	if o == nil || IsNil(o.StructuredFields) {
		return nil, false
	}
	return o.StructuredFields, true
}

// HasStructuredFields returns a boolean if a field has been set.
func (o *UpdateLogMappingRequestBodyFields) HasStructuredFields() bool {
	if o != nil && !IsNil(o.StructuredFields) {
		return true
	}

	return false
}

// SetStructuredFields gets a reference to the given CreateLogMappingRequestBodyFieldsStructuredFields and assigns it to the StructuredFields field.
func (o *UpdateLogMappingRequestBodyFields) SetStructuredFields(v CreateLogMappingRequestBodyFieldsStructuredFields) {
	o.StructuredFields = &v
}

// GetStructuredInputs returns the StructuredInputs field value if set, zero value otherwise.
func (o *UpdateLogMappingRequestBodyFields) GetStructuredInputs() []CreateLogMappingRequestBodyFieldsStructuredFields {
	if o == nil || IsNil(o.StructuredInputs) {
		var ret []CreateLogMappingRequestBodyFieldsStructuredFields
		return ret
	}
	return o.StructuredInputs
}

// GetStructuredInputsOk returns a tuple with the StructuredInputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLogMappingRequestBodyFields) GetStructuredInputsOk() ([]CreateLogMappingRequestBodyFieldsStructuredFields, bool) {
	if o == nil || IsNil(o.StructuredInputs) {
		return nil, false
	}
	return o.StructuredInputs, true
}

// HasStructuredInputs returns a boolean if a field has been set.
func (o *UpdateLogMappingRequestBodyFields) HasStructuredInputs() bool {
	if o != nil && !IsNil(o.StructuredInputs) {
		return true
	}

	return false
}

// SetStructuredInputs gets a reference to the given []CreateLogMappingRequestBodyFieldsStructuredFields and assigns it to the StructuredInputs field.
func (o *UpdateLogMappingRequestBodyFields) SetStructuredInputs(v []CreateLogMappingRequestBodyFieldsStructuredFields) {
	o.StructuredInputs = v
}

// GetRecordType returns the RecordType field value
func (o *UpdateLogMappingRequestBodyFields) GetRecordType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value
// and a boolean to check if the value has been set.
func (o *UpdateLogMappingRequestBodyFields) GetRecordTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordType, true
}

// SetRecordType sets field value
func (o *UpdateLogMappingRequestBodyFields) SetRecordType(v string) {
	o.RecordType = v
}

// GetProductGuid returns the ProductGuid field value
func (o *UpdateLogMappingRequestBodyFields) GetProductGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductGuid
}

// GetProductGuidOk returns a tuple with the ProductGuid field value
// and a boolean to check if the value has been set.
func (o *UpdateLogMappingRequestBodyFields) GetProductGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductGuid, true
}

// SetProductGuid sets field value
func (o *UpdateLogMappingRequestBodyFields) SetProductGuid(v string) {
	o.ProductGuid = v
}

// GetRelatesEntities returns the RelatesEntities field value if set, zero value otherwise.
func (o *UpdateLogMappingRequestBodyFields) GetRelatesEntities() bool {
	if o == nil || IsNil(o.RelatesEntities) {
		var ret bool
		return ret
	}
	return *o.RelatesEntities
}

// GetRelatesEntitiesOk returns a tuple with the RelatesEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLogMappingRequestBodyFields) GetRelatesEntitiesOk() (*bool, bool) {
	if o == nil || IsNil(o.RelatesEntities) {
		return nil, false
	}
	return o.RelatesEntities, true
}

// HasRelatesEntities returns a boolean if a field has been set.
func (o *UpdateLogMappingRequestBodyFields) HasRelatesEntities() bool {
	if o != nil && !IsNil(o.RelatesEntities) {
		return true
	}

	return false
}

// SetRelatesEntities gets a reference to the given bool and assigns it to the RelatesEntities field.
func (o *UpdateLogMappingRequestBodyFields) SetRelatesEntities(v bool) {
	o.RelatesEntities = &v
}

func (o UpdateLogMappingRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateLogMappingRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}
	toSerialize["fields"] = o.Fields
	if !IsNil(o.SkippedValues) {
		toSerialize["skippedValues"] = o.SkippedValues
	}
	if !IsNil(o.UnstructuredFields) {
		toSerialize["unstructuredFields"] = o.UnstructuredFields
	}
	if !IsNil(o.StructuredFields) {
		toSerialize["structuredFields"] = o.StructuredFields
	}
	if !IsNil(o.StructuredInputs) {
		toSerialize["structuredInputs"] = o.StructuredInputs
	}
	toSerialize["recordType"] = o.RecordType
	toSerialize["productGuid"] = o.ProductGuid
	if !IsNil(o.RelatesEntities) {
		toSerialize["relatesEntities"] = o.RelatesEntities
	}
	return toSerialize, nil
}

func (o *UpdateLogMappingRequestBodyFields) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"fields",
		"recordType",
		"productGuid",
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

	varUpdateLogMappingRequestBodyFields := _UpdateLogMappingRequestBodyFields{}

	err = json.Unmarshal(bytes, &varUpdateLogMappingRequestBodyFields)

	if err != nil {
		return err
	}

	*o = UpdateLogMappingRequestBodyFields(varUpdateLogMappingRequestBodyFields)

	return err
}

type NullableUpdateLogMappingRequestBodyFields struct {
	value *UpdateLogMappingRequestBodyFields
	isSet bool
}

func (v NullableUpdateLogMappingRequestBodyFields) Get() *UpdateLogMappingRequestBodyFields {
	return v.value
}

func (v *NullableUpdateLogMappingRequestBodyFields) Set(val *UpdateLogMappingRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLogMappingRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLogMappingRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLogMappingRequestBodyFields(val *UpdateLogMappingRequestBodyFields) *NullableUpdateLogMappingRequestBodyFields {
	return &NullableUpdateLogMappingRequestBodyFields{value: val, isSet: true}
}

func (v NullableUpdateLogMappingRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLogMappingRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


