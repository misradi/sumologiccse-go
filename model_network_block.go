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

// checks if the NetworkBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkBlock{}

// NetworkBlock struct for NetworkBlock
type NetworkBlock struct {
	Id string `json:"id"`
	AddressBlock string `json:"addressBlock"`
	// The name of the List.
	Label string `json:"label"`
	Internal bool `json:"internal"`
	SuppressesSignals bool `json:"suppressesSignals"`
}

type _NetworkBlock NetworkBlock

// NewNetworkBlock instantiates a new NetworkBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkBlock(id string, addressBlock string, label string, internal bool, suppressesSignals bool) *NetworkBlock {
	this := NetworkBlock{}
	this.Id = id
	this.AddressBlock = addressBlock
	this.Label = label
	this.Internal = internal
	this.SuppressesSignals = suppressesSignals
	return &this
}

// NewNetworkBlockWithDefaults instantiates a new NetworkBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkBlockWithDefaults() *NetworkBlock {
	this := NetworkBlock{}
	return &this
}

// GetId returns the Id field value
func (o *NetworkBlock) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetworkBlock) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetworkBlock) SetId(v string) {
	o.Id = v
}

// GetAddressBlock returns the AddressBlock field value
func (o *NetworkBlock) GetAddressBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressBlock
}

// GetAddressBlockOk returns a tuple with the AddressBlock field value
// and a boolean to check if the value has been set.
func (o *NetworkBlock) GetAddressBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressBlock, true
}

// SetAddressBlock sets field value
func (o *NetworkBlock) SetAddressBlock(v string) {
	o.AddressBlock = v
}

// GetLabel returns the Label field value
func (o *NetworkBlock) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *NetworkBlock) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *NetworkBlock) SetLabel(v string) {
	o.Label = v
}

// GetInternal returns the Internal field value
func (o *NetworkBlock) GetInternal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Internal
}

// GetInternalOk returns a tuple with the Internal field value
// and a boolean to check if the value has been set.
func (o *NetworkBlock) GetInternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Internal, true
}

// SetInternal sets field value
func (o *NetworkBlock) SetInternal(v bool) {
	o.Internal = v
}

// GetSuppressesSignals returns the SuppressesSignals field value
func (o *NetworkBlock) GetSuppressesSignals() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SuppressesSignals
}

// GetSuppressesSignalsOk returns a tuple with the SuppressesSignals field value
// and a boolean to check if the value has been set.
func (o *NetworkBlock) GetSuppressesSignalsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuppressesSignals, true
}

// SetSuppressesSignals sets field value
func (o *NetworkBlock) SetSuppressesSignals(v bool) {
	o.SuppressesSignals = v
}

func (o NetworkBlock) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["addressBlock"] = o.AddressBlock
	toSerialize["label"] = o.Label
	toSerialize["internal"] = o.Internal
	toSerialize["suppressesSignals"] = o.SuppressesSignals
	return toSerialize, nil
}

func (o *NetworkBlock) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"addressBlock",
		"label",
		"internal",
		"suppressesSignals",
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

	varNetworkBlock := _NetworkBlock{}

	err = json.Unmarshal(bytes, &varNetworkBlock)

	if err != nil {
		return err
	}

	*o = NetworkBlock(varNetworkBlock)

	return err
}

type NullableNetworkBlock struct {
	value *NetworkBlock
	isSet bool
}

func (v NullableNetworkBlock) Get() *NetworkBlock {
	return v.value
}

func (v *NullableNetworkBlock) Set(val *NetworkBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkBlock(val *NetworkBlock) *NullableNetworkBlock {
	return &NullableNetworkBlock{value: val, isSet: true}
}

func (v NullableNetworkBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


