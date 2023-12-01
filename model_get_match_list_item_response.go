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

// checks if the GetMatchListItemResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMatchListItemResponse{}

// GetMatchListItemResponse struct for GetMatchListItemResponse
type GetMatchListItemResponse struct {
	Data MatchListItem `json:"data"`
}

type _GetMatchListItemResponse GetMatchListItemResponse

// NewGetMatchListItemResponse instantiates a new GetMatchListItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMatchListItemResponse(data MatchListItem) *GetMatchListItemResponse {
	this := GetMatchListItemResponse{}
	this.Data = data
	return &this
}

// NewGetMatchListItemResponseWithDefaults instantiates a new GetMatchListItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMatchListItemResponseWithDefaults() *GetMatchListItemResponse {
	this := GetMatchListItemResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetMatchListItemResponse) GetData() MatchListItem {
	if o == nil {
		var ret MatchListItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetMatchListItemResponse) GetDataOk() (*MatchListItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetMatchListItemResponse) SetData(v MatchListItem) {
	o.Data = v
}

func (o GetMatchListItemResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMatchListItemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetMatchListItemResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varGetMatchListItemResponse := _GetMatchListItemResponse{}

	err = json.Unmarshal(bytes, &varGetMatchListItemResponse)

	if err != nil {
		return err
	}

	*o = GetMatchListItemResponse(varGetMatchListItemResponse)

	return err
}

type NullableGetMatchListItemResponse struct {
	value *GetMatchListItemResponse
	isSet bool
}

func (v NullableGetMatchListItemResponse) Get() *GetMatchListItemResponse {
	return v.value
}

func (v *NullableGetMatchListItemResponse) Set(val *GetMatchListItemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMatchListItemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMatchListItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMatchListItemResponse(val *GetMatchListItemResponse) *NullableGetMatchListItemResponse {
	return &NullableGetMatchListItemResponse{value: val, isSet: true}
}

func (v NullableGetMatchListItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMatchListItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


