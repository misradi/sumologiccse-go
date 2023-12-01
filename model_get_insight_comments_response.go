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

// checks if the GetInsightCommentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInsightCommentsResponse{}

// GetInsightCommentsResponse struct for GetInsightCommentsResponse
type GetInsightCommentsResponse struct {
	Data GetInsightCommentsResponseData `json:"data"`
}

type _GetInsightCommentsResponse GetInsightCommentsResponse

// NewGetInsightCommentsResponse instantiates a new GetInsightCommentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInsightCommentsResponse(data GetInsightCommentsResponseData) *GetInsightCommentsResponse {
	this := GetInsightCommentsResponse{}
	this.Data = data
	return &this
}

// NewGetInsightCommentsResponseWithDefaults instantiates a new GetInsightCommentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInsightCommentsResponseWithDefaults() *GetInsightCommentsResponse {
	this := GetInsightCommentsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetInsightCommentsResponse) GetData() GetInsightCommentsResponseData {
	if o == nil {
		var ret GetInsightCommentsResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetInsightCommentsResponse) GetDataOk() (*GetInsightCommentsResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetInsightCommentsResponse) SetData(v GetInsightCommentsResponseData) {
	o.Data = v
}

func (o GetInsightCommentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInsightCommentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetInsightCommentsResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varGetInsightCommentsResponse := _GetInsightCommentsResponse{}

	err = json.Unmarshal(bytes, &varGetInsightCommentsResponse)

	if err != nil {
		return err
	}

	*o = GetInsightCommentsResponse(varGetInsightCommentsResponse)

	return err
}

type NullableGetInsightCommentsResponse struct {
	value *GetInsightCommentsResponse
	isSet bool
}

func (v NullableGetInsightCommentsResponse) Get() *GetInsightCommentsResponse {
	return v.value
}

func (v *NullableGetInsightCommentsResponse) Set(val *GetInsightCommentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInsightCommentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInsightCommentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInsightCommentsResponse(val *GetInsightCommentsResponse) *NullableGetInsightCommentsResponse {
	return &NullableGetInsightCommentsResponse{value: val, isSet: true}
}

func (v NullableGetInsightCommentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInsightCommentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

