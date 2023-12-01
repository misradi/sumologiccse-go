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

// checks if the GetInsightCommentsResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInsightCommentsResponseData{}

// GetInsightCommentsResponseData struct for GetInsightCommentsResponseData
type GetInsightCommentsResponseData struct {
	Comments []GetInsightCommentsResponseDataCommentsInner `json:"comments"`
}

type _GetInsightCommentsResponseData GetInsightCommentsResponseData

// NewGetInsightCommentsResponseData instantiates a new GetInsightCommentsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInsightCommentsResponseData(comments []GetInsightCommentsResponseDataCommentsInner) *GetInsightCommentsResponseData {
	this := GetInsightCommentsResponseData{}
	this.Comments = comments
	return &this
}

// NewGetInsightCommentsResponseDataWithDefaults instantiates a new GetInsightCommentsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInsightCommentsResponseDataWithDefaults() *GetInsightCommentsResponseData {
	this := GetInsightCommentsResponseData{}
	return &this
}

// GetComments returns the Comments field value
func (o *GetInsightCommentsResponseData) GetComments() []GetInsightCommentsResponseDataCommentsInner {
	if o == nil {
		var ret []GetInsightCommentsResponseDataCommentsInner
		return ret
	}

	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value
// and a boolean to check if the value has been set.
func (o *GetInsightCommentsResponseData) GetCommentsOk() ([]GetInsightCommentsResponseDataCommentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comments, true
}

// SetComments sets field value
func (o *GetInsightCommentsResponseData) SetComments(v []GetInsightCommentsResponseDataCommentsInner) {
	o.Comments = v
}

func (o GetInsightCommentsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInsightCommentsResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["comments"] = o.Comments
	return toSerialize, nil
}

func (o *GetInsightCommentsResponseData) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"comments",
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

	varGetInsightCommentsResponseData := _GetInsightCommentsResponseData{}

	err = json.Unmarshal(bytes, &varGetInsightCommentsResponseData)

	if err != nil {
		return err
	}

	*o = GetInsightCommentsResponseData(varGetInsightCommentsResponseData)

	return err
}

type NullableGetInsightCommentsResponseData struct {
	value *GetInsightCommentsResponseData
	isSet bool
}

func (v NullableGetInsightCommentsResponseData) Get() *GetInsightCommentsResponseData {
	return v.value
}

func (v *NullableGetInsightCommentsResponseData) Set(val *GetInsightCommentsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInsightCommentsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInsightCommentsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInsightCommentsResponseData(val *GetInsightCommentsResponseData) *NullableGetInsightCommentsResponseData {
	return &NullableGetInsightCommentsResponseData{value: val, isSet: true}
}

func (v NullableGetInsightCommentsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInsightCommentsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


