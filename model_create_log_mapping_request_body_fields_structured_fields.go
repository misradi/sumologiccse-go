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

// checks if the CreateLogMappingRequestBodyFieldsStructuredFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLogMappingRequestBodyFieldsStructuredFields{}

// CreateLogMappingRequestBodyFieldsStructuredFields struct for CreateLogMappingRequestBodyFieldsStructuredFields
type CreateLogMappingRequestBodyFieldsStructuredFields struct {
	LogFormat string `json:"logFormat"`
	Vendor string `json:"vendor"`
	Product string `json:"product"`
	EventIdPattern string `json:"eventIdPattern"`
}

type _CreateLogMappingRequestBodyFieldsStructuredFields CreateLogMappingRequestBodyFieldsStructuredFields

// NewCreateLogMappingRequestBodyFieldsStructuredFields instantiates a new CreateLogMappingRequestBodyFieldsStructuredFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLogMappingRequestBodyFieldsStructuredFields(logFormat string, vendor string, product string, eventIdPattern string) *CreateLogMappingRequestBodyFieldsStructuredFields {
	this := CreateLogMappingRequestBodyFieldsStructuredFields{}
	this.LogFormat = logFormat
	this.Vendor = vendor
	this.Product = product
	this.EventIdPattern = eventIdPattern
	return &this
}

// NewCreateLogMappingRequestBodyFieldsStructuredFieldsWithDefaults instantiates a new CreateLogMappingRequestBodyFieldsStructuredFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLogMappingRequestBodyFieldsStructuredFieldsWithDefaults() *CreateLogMappingRequestBodyFieldsStructuredFields {
	this := CreateLogMappingRequestBodyFieldsStructuredFields{}
	return &this
}

// GetLogFormat returns the LogFormat field value
func (o *CreateLogMappingRequestBodyFieldsStructuredFields) GetLogFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFormat
}

// GetLogFormatOk returns a tuple with the LogFormat field value
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsStructuredFields) GetLogFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogFormat, true
}

// SetLogFormat sets field value
func (o *CreateLogMappingRequestBodyFieldsStructuredFields) SetLogFormat(v string) {
	o.LogFormat = v
}

// GetVendor returns the Vendor field value
func (o *CreateLogMappingRequestBodyFieldsStructuredFields) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsStructuredFields) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *CreateLogMappingRequestBodyFieldsStructuredFields) SetVendor(v string) {
	o.Vendor = v
}

// GetProduct returns the Product field value
func (o *CreateLogMappingRequestBodyFieldsStructuredFields) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsStructuredFields) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *CreateLogMappingRequestBodyFieldsStructuredFields) SetProduct(v string) {
	o.Product = v
}

// GetEventIdPattern returns the EventIdPattern field value
func (o *CreateLogMappingRequestBodyFieldsStructuredFields) GetEventIdPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventIdPattern
}

// GetEventIdPatternOk returns a tuple with the EventIdPattern field value
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsStructuredFields) GetEventIdPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIdPattern, true
}

// SetEventIdPattern sets field value
func (o *CreateLogMappingRequestBodyFieldsStructuredFields) SetEventIdPattern(v string) {
	o.EventIdPattern = v
}

func (o CreateLogMappingRequestBodyFieldsStructuredFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLogMappingRequestBodyFieldsStructuredFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["logFormat"] = o.LogFormat
	toSerialize["vendor"] = o.Vendor
	toSerialize["product"] = o.Product
	toSerialize["eventIdPattern"] = o.EventIdPattern
	return toSerialize, nil
}

func (o *CreateLogMappingRequestBodyFieldsStructuredFields) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"logFormat",
		"vendor",
		"product",
		"eventIdPattern",
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

	varCreateLogMappingRequestBodyFieldsStructuredFields := _CreateLogMappingRequestBodyFieldsStructuredFields{}

	err = json.Unmarshal(bytes, &varCreateLogMappingRequestBodyFieldsStructuredFields)

	if err != nil {
		return err
	}

	*o = CreateLogMappingRequestBodyFieldsStructuredFields(varCreateLogMappingRequestBodyFieldsStructuredFields)

	return err
}

type NullableCreateLogMappingRequestBodyFieldsStructuredFields struct {
	value *CreateLogMappingRequestBodyFieldsStructuredFields
	isSet bool
}

func (v NullableCreateLogMappingRequestBodyFieldsStructuredFields) Get() *CreateLogMappingRequestBodyFieldsStructuredFields {
	return v.value
}

func (v *NullableCreateLogMappingRequestBodyFieldsStructuredFields) Set(val *CreateLogMappingRequestBodyFieldsStructuredFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLogMappingRequestBodyFieldsStructuredFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLogMappingRequestBodyFieldsStructuredFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLogMappingRequestBodyFieldsStructuredFields(val *CreateLogMappingRequestBodyFieldsStructuredFields) *NullableCreateLogMappingRequestBodyFieldsStructuredFields {
	return &NullableCreateLogMappingRequestBodyFieldsStructuredFields{value: val, isSet: true}
}

func (v NullableCreateLogMappingRequestBodyFieldsStructuredFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLogMappingRequestBodyFieldsStructuredFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


