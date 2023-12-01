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

// checks if the GetCustomInsightResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCustomInsightResponse{}

// GetCustomInsightResponse struct for GetCustomInsightResponse
type GetCustomInsightResponse struct {
	Data CustomInsight `json:"data"`
}

type _GetCustomInsightResponse GetCustomInsightResponse

// NewGetCustomInsightResponse instantiates a new GetCustomInsightResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCustomInsightResponse(data CustomInsight) *GetCustomInsightResponse {
	this := GetCustomInsightResponse{}
	this.Data = data
	return &this
}

// NewGetCustomInsightResponseWithDefaults instantiates a new GetCustomInsightResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCustomInsightResponseWithDefaults() *GetCustomInsightResponse {
	this := GetCustomInsightResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetCustomInsightResponse) GetData() CustomInsight {
	if o == nil {
		var ret CustomInsight
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCustomInsightResponse) GetDataOk() (*CustomInsight, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetCustomInsightResponse) SetData(v CustomInsight) {
	o.Data = v
}

func (o GetCustomInsightResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCustomInsightResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetCustomInsightResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varGetCustomInsightResponse := _GetCustomInsightResponse{}

	err = json.Unmarshal(bytes, &varGetCustomInsightResponse)

	if err != nil {
		return err
	}

	*o = GetCustomInsightResponse(varGetCustomInsightResponse)

	return err
}

type NullableGetCustomInsightResponse struct {
	value *GetCustomInsightResponse
	isSet bool
}

func (v NullableGetCustomInsightResponse) Get() *GetCustomInsightResponse {
	return v.value
}

func (v *NullableGetCustomInsightResponse) Set(val *GetCustomInsightResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCustomInsightResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCustomInsightResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCustomInsightResponse(val *GetCustomInsightResponse) *NullableGetCustomInsightResponse {
	return &NullableGetCustomInsightResponse{value: val, isSet: true}
}

func (v NullableGetCustomInsightResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCustomInsightResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


