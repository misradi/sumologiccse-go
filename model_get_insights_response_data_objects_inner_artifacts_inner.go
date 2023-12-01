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

// checks if the GetInsightsResponseDataObjectsInnerArtifactsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInsightsResponseDataObjectsInnerArtifactsInner{}

// GetInsightsResponseDataObjectsInnerArtifactsInner struct for GetInsightsResponseDataObjectsInnerArtifactsInner
type GetInsightsResponseDataObjectsInnerArtifactsInner struct {
	Name string `json:"name"`
	Value string `json:"value"`
	RecordUid string `json:"recordUid"`
	RecordStream string `json:"recordStream"`
	Signal GetInsightsResponseDataObjectsInnerArtifactsInnerSignal `json:"signal"`
}

type _GetInsightsResponseDataObjectsInnerArtifactsInner GetInsightsResponseDataObjectsInnerArtifactsInner

// NewGetInsightsResponseDataObjectsInnerArtifactsInner instantiates a new GetInsightsResponseDataObjectsInnerArtifactsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInsightsResponseDataObjectsInnerArtifactsInner(name string, value string, recordUid string, recordStream string, signal GetInsightsResponseDataObjectsInnerArtifactsInnerSignal) *GetInsightsResponseDataObjectsInnerArtifactsInner {
	this := GetInsightsResponseDataObjectsInnerArtifactsInner{}
	this.Name = name
	this.Value = value
	this.RecordUid = recordUid
	this.RecordStream = recordStream
	this.Signal = signal
	return &this
}

// NewGetInsightsResponseDataObjectsInnerArtifactsInnerWithDefaults instantiates a new GetInsightsResponseDataObjectsInnerArtifactsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInsightsResponseDataObjectsInnerArtifactsInnerWithDefaults() *GetInsightsResponseDataObjectsInnerArtifactsInner {
	this := GetInsightsResponseDataObjectsInnerArtifactsInner{}
	return &this
}

// GetName returns the Name field value
func (o *GetInsightsResponseDataObjectsInnerArtifactsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerArtifactsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetInsightsResponseDataObjectsInnerArtifactsInner) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *GetInsightsResponseDataObjectsInnerArtifactsInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerArtifactsInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetInsightsResponseDataObjectsInnerArtifactsInner) SetValue(v string) {
	o.Value = v
}

// GetRecordUid returns the RecordUid field value
func (o *GetInsightsResponseDataObjectsInnerArtifactsInner) GetRecordUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecordUid
}

// GetRecordUidOk returns a tuple with the RecordUid field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerArtifactsInner) GetRecordUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordUid, true
}

// SetRecordUid sets field value
func (o *GetInsightsResponseDataObjectsInnerArtifactsInner) SetRecordUid(v string) {
	o.RecordUid = v
}

// GetRecordStream returns the RecordStream field value
func (o *GetInsightsResponseDataObjectsInnerArtifactsInner) GetRecordStream() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecordStream
}

// GetRecordStreamOk returns a tuple with the RecordStream field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerArtifactsInner) GetRecordStreamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordStream, true
}

// SetRecordStream sets field value
func (o *GetInsightsResponseDataObjectsInnerArtifactsInner) SetRecordStream(v string) {
	o.RecordStream = v
}

// GetSignal returns the Signal field value
func (o *GetInsightsResponseDataObjectsInnerArtifactsInner) GetSignal() GetInsightsResponseDataObjectsInnerArtifactsInnerSignal {
	if o == nil {
		var ret GetInsightsResponseDataObjectsInnerArtifactsInnerSignal
		return ret
	}

	return o.Signal
}

// GetSignalOk returns a tuple with the Signal field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerArtifactsInner) GetSignalOk() (*GetInsightsResponseDataObjectsInnerArtifactsInnerSignal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signal, true
}

// SetSignal sets field value
func (o *GetInsightsResponseDataObjectsInnerArtifactsInner) SetSignal(v GetInsightsResponseDataObjectsInnerArtifactsInnerSignal) {
	o.Signal = v
}

func (o GetInsightsResponseDataObjectsInnerArtifactsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInsightsResponseDataObjectsInnerArtifactsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	toSerialize["recordUid"] = o.RecordUid
	toSerialize["recordStream"] = o.RecordStream
	toSerialize["signal"] = o.Signal
	return toSerialize, nil
}

func (o *GetInsightsResponseDataObjectsInnerArtifactsInner) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
		"recordUid",
		"recordStream",
		"signal",
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

	varGetInsightsResponseDataObjectsInnerArtifactsInner := _GetInsightsResponseDataObjectsInnerArtifactsInner{}

	err = json.Unmarshal(bytes, &varGetInsightsResponseDataObjectsInnerArtifactsInner)

	if err != nil {
		return err
	}

	*o = GetInsightsResponseDataObjectsInnerArtifactsInner(varGetInsightsResponseDataObjectsInnerArtifactsInner)

	return err
}

type NullableGetInsightsResponseDataObjectsInnerArtifactsInner struct {
	value *GetInsightsResponseDataObjectsInnerArtifactsInner
	isSet bool
}

func (v NullableGetInsightsResponseDataObjectsInnerArtifactsInner) Get() *GetInsightsResponseDataObjectsInnerArtifactsInner {
	return v.value
}

func (v *NullableGetInsightsResponseDataObjectsInnerArtifactsInner) Set(val *GetInsightsResponseDataObjectsInnerArtifactsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInsightsResponseDataObjectsInnerArtifactsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInsightsResponseDataObjectsInnerArtifactsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInsightsResponseDataObjectsInnerArtifactsInner(val *GetInsightsResponseDataObjectsInnerArtifactsInner) *NullableGetInsightsResponseDataObjectsInnerArtifactsInner {
	return &NullableGetInsightsResponseDataObjectsInnerArtifactsInner{value: val, isSet: true}
}

func (v NullableGetInsightsResponseDataObjectsInnerArtifactsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInsightsResponseDataObjectsInnerArtifactsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


