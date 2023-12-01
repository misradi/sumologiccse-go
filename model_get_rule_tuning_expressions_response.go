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

// checks if the GetRuleTuningExpressionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRuleTuningExpressionsResponse{}

// GetRuleTuningExpressionsResponse struct for GetRuleTuningExpressionsResponse
type GetRuleTuningExpressionsResponse struct {
	Data GetRuleTuningExpressionsResponseData `json:"data"`
}

type _GetRuleTuningExpressionsResponse GetRuleTuningExpressionsResponse

// NewGetRuleTuningExpressionsResponse instantiates a new GetRuleTuningExpressionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRuleTuningExpressionsResponse(data GetRuleTuningExpressionsResponseData) *GetRuleTuningExpressionsResponse {
	this := GetRuleTuningExpressionsResponse{}
	this.Data = data
	return &this
}

// NewGetRuleTuningExpressionsResponseWithDefaults instantiates a new GetRuleTuningExpressionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRuleTuningExpressionsResponseWithDefaults() *GetRuleTuningExpressionsResponse {
	this := GetRuleTuningExpressionsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetRuleTuningExpressionsResponse) GetData() GetRuleTuningExpressionsResponseData {
	if o == nil {
		var ret GetRuleTuningExpressionsResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetRuleTuningExpressionsResponse) GetDataOk() (*GetRuleTuningExpressionsResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetRuleTuningExpressionsResponse) SetData(v GetRuleTuningExpressionsResponseData) {
	o.Data = v
}

func (o GetRuleTuningExpressionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRuleTuningExpressionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetRuleTuningExpressionsResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varGetRuleTuningExpressionsResponse := _GetRuleTuningExpressionsResponse{}

	err = json.Unmarshal(bytes, &varGetRuleTuningExpressionsResponse)

	if err != nil {
		return err
	}

	*o = GetRuleTuningExpressionsResponse(varGetRuleTuningExpressionsResponse)

	return err
}

type NullableGetRuleTuningExpressionsResponse struct {
	value *GetRuleTuningExpressionsResponse
	isSet bool
}

func (v NullableGetRuleTuningExpressionsResponse) Get() *GetRuleTuningExpressionsResponse {
	return v.value
}

func (v *NullableGetRuleTuningExpressionsResponse) Set(val *GetRuleTuningExpressionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRuleTuningExpressionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRuleTuningExpressionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRuleTuningExpressionsResponse(val *GetRuleTuningExpressionsResponse) *NullableGetRuleTuningExpressionsResponse {
	return &NullableGetRuleTuningExpressionsResponse{value: val, isSet: true}
}

func (v NullableGetRuleTuningExpressionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRuleTuningExpressionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


