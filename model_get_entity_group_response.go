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

// checks if the GetEntityGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEntityGroupResponse{}

// GetEntityGroupResponse struct for GetEntityGroupResponse
type GetEntityGroupResponse struct {
	Data GetEntityGroupsResponseDataObjectsInner `json:"data"`
}

type _GetEntityGroupResponse GetEntityGroupResponse

// NewGetEntityGroupResponse instantiates a new GetEntityGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEntityGroupResponse(data GetEntityGroupsResponseDataObjectsInner) *GetEntityGroupResponse {
	this := GetEntityGroupResponse{}
	this.Data = data
	return &this
}

// NewGetEntityGroupResponseWithDefaults instantiates a new GetEntityGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEntityGroupResponseWithDefaults() *GetEntityGroupResponse {
	this := GetEntityGroupResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetEntityGroupResponse) GetData() GetEntityGroupsResponseDataObjectsInner {
	if o == nil {
		var ret GetEntityGroupsResponseDataObjectsInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetEntityGroupResponse) GetDataOk() (*GetEntityGroupsResponseDataObjectsInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetEntityGroupResponse) SetData(v GetEntityGroupsResponseDataObjectsInner) {
	o.Data = v
}

func (o GetEntityGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEntityGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetEntityGroupResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varGetEntityGroupResponse := _GetEntityGroupResponse{}

	err = json.Unmarshal(bytes, &varGetEntityGroupResponse)

	if err != nil {
		return err
	}

	*o = GetEntityGroupResponse(varGetEntityGroupResponse)

	return err
}

type NullableGetEntityGroupResponse struct {
	value *GetEntityGroupResponse
	isSet bool
}

func (v NullableGetEntityGroupResponse) Get() *GetEntityGroupResponse {
	return v.value
}

func (v *NullableGetEntityGroupResponse) Set(val *GetEntityGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEntityGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEntityGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEntityGroupResponse(val *GetEntityGroupResponse) *NullableGetEntityGroupResponse {
	return &NullableGetEntityGroupResponse{value: val, isSet: true}
}

func (v NullableGetEntityGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEntityGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


