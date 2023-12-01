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

// checks if the GetNetworkBlockResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkBlockResponse{}

// GetNetworkBlockResponse struct for GetNetworkBlockResponse
type GetNetworkBlockResponse struct {
	Data NetworkBlock `json:"data"`
}

type _GetNetworkBlockResponse GetNetworkBlockResponse

// NewGetNetworkBlockResponse instantiates a new GetNetworkBlockResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkBlockResponse(data NetworkBlock) *GetNetworkBlockResponse {
	this := GetNetworkBlockResponse{}
	this.Data = data
	return &this
}

// NewGetNetworkBlockResponseWithDefaults instantiates a new GetNetworkBlockResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkBlockResponseWithDefaults() *GetNetworkBlockResponse {
	this := GetNetworkBlockResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetNetworkBlockResponse) GetData() NetworkBlock {
	if o == nil {
		var ret NetworkBlock
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetNetworkBlockResponse) GetDataOk() (*NetworkBlock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetNetworkBlockResponse) SetData(v NetworkBlock) {
	o.Data = v
}

func (o GetNetworkBlockResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkBlockResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetNetworkBlockResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varGetNetworkBlockResponse := _GetNetworkBlockResponse{}

	err = json.Unmarshal(bytes, &varGetNetworkBlockResponse)

	if err != nil {
		return err
	}

	*o = GetNetworkBlockResponse(varGetNetworkBlockResponse)

	return err
}

type NullableGetNetworkBlockResponse struct {
	value *GetNetworkBlockResponse
	isSet bool
}

func (v NullableGetNetworkBlockResponse) Get() *GetNetworkBlockResponse {
	return v.value
}

func (v *NullableGetNetworkBlockResponse) Set(val *GetNetworkBlockResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkBlockResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkBlockResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkBlockResponse(val *GetNetworkBlockResponse) *NullableGetNetworkBlockResponse {
	return &NullableGetNetworkBlockResponse{value: val, isSet: true}
}

func (v NullableGetNetworkBlockResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkBlockResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


