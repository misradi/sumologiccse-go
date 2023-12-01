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

// checks if the GetCustomEntityTypesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCustomEntityTypesResponse{}

// GetCustomEntityTypesResponse struct for GetCustomEntityTypesResponse
type GetCustomEntityTypesResponse struct {
	Data GetCustomEntityTypesResponseData `json:"data"`
}

type _GetCustomEntityTypesResponse GetCustomEntityTypesResponse

// NewGetCustomEntityTypesResponse instantiates a new GetCustomEntityTypesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCustomEntityTypesResponse(data GetCustomEntityTypesResponseData) *GetCustomEntityTypesResponse {
	this := GetCustomEntityTypesResponse{}
	this.Data = data
	return &this
}

// NewGetCustomEntityTypesResponseWithDefaults instantiates a new GetCustomEntityTypesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCustomEntityTypesResponseWithDefaults() *GetCustomEntityTypesResponse {
	this := GetCustomEntityTypesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetCustomEntityTypesResponse) GetData() GetCustomEntityTypesResponseData {
	if o == nil {
		var ret GetCustomEntityTypesResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCustomEntityTypesResponse) GetDataOk() (*GetCustomEntityTypesResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetCustomEntityTypesResponse) SetData(v GetCustomEntityTypesResponseData) {
	o.Data = v
}

func (o GetCustomEntityTypesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCustomEntityTypesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetCustomEntityTypesResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varGetCustomEntityTypesResponse := _GetCustomEntityTypesResponse{}

	err = json.Unmarshal(bytes, &varGetCustomEntityTypesResponse)

	if err != nil {
		return err
	}

	*o = GetCustomEntityTypesResponse(varGetCustomEntityTypesResponse)

	return err
}

type NullableGetCustomEntityTypesResponse struct {
	value *GetCustomEntityTypesResponse
	isSet bool
}

func (v NullableGetCustomEntityTypesResponse) Get() *GetCustomEntityTypesResponse {
	return v.value
}

func (v *NullableGetCustomEntityTypesResponse) Set(val *GetCustomEntityTypesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCustomEntityTypesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCustomEntityTypesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCustomEntityTypesResponse(val *GetCustomEntityTypesResponse) *NullableGetCustomEntityTypesResponse {
	return &NullableGetCustomEntityTypesResponse{value: val, isSet: true}
}

func (v NullableGetCustomEntityTypesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCustomEntityTypesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


