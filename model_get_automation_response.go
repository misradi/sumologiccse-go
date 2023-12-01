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

// checks if the GetAutomationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAutomationResponse{}

// GetAutomationResponse struct for GetAutomationResponse
type GetAutomationResponse struct {
	Data Automation `json:"data"`
}

type _GetAutomationResponse GetAutomationResponse

// NewGetAutomationResponse instantiates a new GetAutomationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAutomationResponse(data Automation) *GetAutomationResponse {
	this := GetAutomationResponse{}
	this.Data = data
	return &this
}

// NewGetAutomationResponseWithDefaults instantiates a new GetAutomationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAutomationResponseWithDefaults() *GetAutomationResponse {
	this := GetAutomationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetAutomationResponse) GetData() Automation {
	if o == nil {
		var ret Automation
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetAutomationResponse) GetDataOk() (*Automation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetAutomationResponse) SetData(v Automation) {
	o.Data = v
}

func (o GetAutomationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAutomationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetAutomationResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varGetAutomationResponse := _GetAutomationResponse{}

	err = json.Unmarshal(bytes, &varGetAutomationResponse)

	if err != nil {
		return err
	}

	*o = GetAutomationResponse(varGetAutomationResponse)

	return err
}

type NullableGetAutomationResponse struct {
	value *GetAutomationResponse
	isSet bool
}

func (v NullableGetAutomationResponse) Get() *GetAutomationResponse {
	return v.value
}

func (v *NullableGetAutomationResponse) Set(val *GetAutomationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAutomationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAutomationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAutomationResponse(val *GetAutomationResponse) *NullableGetAutomationResponse {
	return &NullableGetAutomationResponse{value: val, isSet: true}
}

func (v NullableGetAutomationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAutomationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


