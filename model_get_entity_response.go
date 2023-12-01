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

// checks if the GetEntityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEntityResponse{}

// GetEntityResponse struct for GetEntityResponse
type GetEntityResponse struct {
	Data GetEntityResponseData `json:"data"`
}

type _GetEntityResponse GetEntityResponse

// NewGetEntityResponse instantiates a new GetEntityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEntityResponse(data GetEntityResponseData) *GetEntityResponse {
	this := GetEntityResponse{}
	this.Data = data
	return &this
}

// NewGetEntityResponseWithDefaults instantiates a new GetEntityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEntityResponseWithDefaults() *GetEntityResponse {
	this := GetEntityResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetEntityResponse) GetData() GetEntityResponseData {
	if o == nil {
		var ret GetEntityResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetEntityResponse) GetDataOk() (*GetEntityResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetEntityResponse) SetData(v GetEntityResponseData) {
	o.Data = v
}

func (o GetEntityResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEntityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetEntityResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varGetEntityResponse := _GetEntityResponse{}

	err = json.Unmarshal(bytes, &varGetEntityResponse)

	if err != nil {
		return err
	}

	*o = GetEntityResponse(varGetEntityResponse)

	return err
}

type NullableGetEntityResponse struct {
	value *GetEntityResponse
	isSet bool
}

func (v NullableGetEntityResponse) Get() *GetEntityResponse {
	return v.value
}

func (v *NullableGetEntityResponse) Set(val *GetEntityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEntityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEntityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEntityResponse(val *GetEntityResponse) *NullableGetEntityResponse {
	return &NullableGetEntityResponse{value: val, isSet: true}
}

func (v NullableGetEntityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEntityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

