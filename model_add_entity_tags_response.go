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

// checks if the AddEntityTagsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddEntityTagsResponse{}

// AddEntityTagsResponse struct for AddEntityTagsResponse
type AddEntityTagsResponse struct {
	Data UpdateEntityCriticalityResponseData `json:"data"`
}

type _AddEntityTagsResponse AddEntityTagsResponse

// NewAddEntityTagsResponse instantiates a new AddEntityTagsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddEntityTagsResponse(data UpdateEntityCriticalityResponseData) *AddEntityTagsResponse {
	this := AddEntityTagsResponse{}
	this.Data = data
	return &this
}

// NewAddEntityTagsResponseWithDefaults instantiates a new AddEntityTagsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddEntityTagsResponseWithDefaults() *AddEntityTagsResponse {
	this := AddEntityTagsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AddEntityTagsResponse) GetData() UpdateEntityCriticalityResponseData {
	if o == nil {
		var ret UpdateEntityCriticalityResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddEntityTagsResponse) GetDataOk() (*UpdateEntityCriticalityResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AddEntityTagsResponse) SetData(v UpdateEntityCriticalityResponseData) {
	o.Data = v
}

func (o AddEntityTagsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddEntityTagsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AddEntityTagsResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varAddEntityTagsResponse := _AddEntityTagsResponse{}

	err = json.Unmarshal(bytes, &varAddEntityTagsResponse)

	if err != nil {
		return err
	}

	*o = AddEntityTagsResponse(varAddEntityTagsResponse)

	return err
}

type NullableAddEntityTagsResponse struct {
	value *AddEntityTagsResponse
	isSet bool
}

func (v NullableAddEntityTagsResponse) Get() *AddEntityTagsResponse {
	return v.value
}

func (v *NullableAddEntityTagsResponse) Set(val *AddEntityTagsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddEntityTagsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddEntityTagsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddEntityTagsResponse(val *AddEntityTagsResponse) *NullableAddEntityTagsResponse {
	return &NullableAddEntityTagsResponse{value: val, isSet: true}
}

func (v NullableAddEntityTagsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddEntityTagsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


