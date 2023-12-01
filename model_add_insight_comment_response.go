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

// checks if the AddInsightCommentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddInsightCommentResponse{}

// AddInsightCommentResponse struct for AddInsightCommentResponse
type AddInsightCommentResponse struct {
	Data GetInsightCommentsResponseDataCommentsInner `json:"data"`
}

type _AddInsightCommentResponse AddInsightCommentResponse

// NewAddInsightCommentResponse instantiates a new AddInsightCommentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddInsightCommentResponse(data GetInsightCommentsResponseDataCommentsInner) *AddInsightCommentResponse {
	this := AddInsightCommentResponse{}
	this.Data = data
	return &this
}

// NewAddInsightCommentResponseWithDefaults instantiates a new AddInsightCommentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddInsightCommentResponseWithDefaults() *AddInsightCommentResponse {
	this := AddInsightCommentResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AddInsightCommentResponse) GetData() GetInsightCommentsResponseDataCommentsInner {
	if o == nil {
		var ret GetInsightCommentsResponseDataCommentsInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddInsightCommentResponse) GetDataOk() (*GetInsightCommentsResponseDataCommentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AddInsightCommentResponse) SetData(v GetInsightCommentsResponseDataCommentsInner) {
	o.Data = v
}

func (o AddInsightCommentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddInsightCommentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AddInsightCommentResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varAddInsightCommentResponse := _AddInsightCommentResponse{}

	err = json.Unmarshal(bytes, &varAddInsightCommentResponse)

	if err != nil {
		return err
	}

	*o = AddInsightCommentResponse(varAddInsightCommentResponse)

	return err
}

type NullableAddInsightCommentResponse struct {
	value *AddInsightCommentResponse
	isSet bool
}

func (v NullableAddInsightCommentResponse) Get() *AddInsightCommentResponse {
	return v.value
}

func (v *NullableAddInsightCommentResponse) Set(val *AddInsightCommentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddInsightCommentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddInsightCommentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddInsightCommentResponse(val *AddInsightCommentResponse) *NullableAddInsightCommentResponse {
	return &NullableAddInsightCommentResponse{value: val, isSet: true}
}

func (v NullableAddInsightCommentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddInsightCommentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


