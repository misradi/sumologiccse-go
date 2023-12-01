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

// checks if the GetInsightsResponseDataObjectsInnerSignalsInnerEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInsightsResponseDataObjectsInnerSignalsInnerEntity{}

// GetInsightsResponseDataObjectsInnerSignalsInnerEntity struct for GetInsightsResponseDataObjectsInnerSignalsInnerEntity
type GetInsightsResponseDataObjectsInnerSignalsInnerEntity struct {
	Id string `json:"id"`
	EntityType string `json:"entityType"`
	Name string `json:"name"`
	Value string `json:"value"`
	Hostname *string `json:"hostname,omitempty"`
	MacAddress *string `json:"macAddress,omitempty"`
	SensorZone *string `json:"sensorZone,omitempty"`
}

type _GetInsightsResponseDataObjectsInnerSignalsInnerEntity GetInsightsResponseDataObjectsInnerSignalsInnerEntity

// NewGetInsightsResponseDataObjectsInnerSignalsInnerEntity instantiates a new GetInsightsResponseDataObjectsInnerSignalsInnerEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInsightsResponseDataObjectsInnerSignalsInnerEntity(id string, entityType string, name string, value string) *GetInsightsResponseDataObjectsInnerSignalsInnerEntity {
	this := GetInsightsResponseDataObjectsInnerSignalsInnerEntity{}
	this.Id = id
	this.EntityType = entityType
	this.Name = name
	this.Value = value
	return &this
}

// NewGetInsightsResponseDataObjectsInnerSignalsInnerEntityWithDefaults instantiates a new GetInsightsResponseDataObjectsInnerSignalsInnerEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInsightsResponseDataObjectsInnerSignalsInnerEntityWithDefaults() *GetInsightsResponseDataObjectsInnerSignalsInnerEntity {
	this := GetInsightsResponseDataObjectsInnerSignalsInnerEntity{}
	return &this
}

// GetId returns the Id field value
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) SetId(v string) {
	o.Id = v
}

// GetEntityType returns the EntityType field value
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) SetEntityType(v string) {
	o.EntityType = v
}

// GetName returns the Name field value
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) SetValue(v string) {
	o.Value = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) SetHostname(v string) {
	o.Hostname = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetSensorZone returns the SensorZone field value if set, zero value otherwise.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) GetSensorZone() string {
	if o == nil || IsNil(o.SensorZone) {
		var ret string
		return ret
	}
	return *o.SensorZone
}

// GetSensorZoneOk returns a tuple with the SensorZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) GetSensorZoneOk() (*string, bool) {
	if o == nil || IsNil(o.SensorZone) {
		return nil, false
	}
	return o.SensorZone, true
}

// HasSensorZone returns a boolean if a field has been set.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) HasSensorZone() bool {
	if o != nil && !IsNil(o.SensorZone) {
		return true
	}

	return false
}

// SetSensorZone gets a reference to the given string and assigns it to the SensorZone field.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) SetSensorZone(v string) {
	o.SensorZone = &v
}

func (o GetInsightsResponseDataObjectsInnerSignalsInnerEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInsightsResponseDataObjectsInnerSignalsInnerEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["entityType"] = o.EntityType
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.MacAddress) {
		toSerialize["macAddress"] = o.MacAddress
	}
	if !IsNil(o.SensorZone) {
		toSerialize["sensorZone"] = o.SensorZone
	}
	return toSerialize, nil
}

func (o *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"entityType",
		"name",
		"value",
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

	varGetInsightsResponseDataObjectsInnerSignalsInnerEntity := _GetInsightsResponseDataObjectsInnerSignalsInnerEntity{}

	err = json.Unmarshal(bytes, &varGetInsightsResponseDataObjectsInnerSignalsInnerEntity)

	if err != nil {
		return err
	}

	*o = GetInsightsResponseDataObjectsInnerSignalsInnerEntity(varGetInsightsResponseDataObjectsInnerSignalsInnerEntity)

	return err
}

type NullableGetInsightsResponseDataObjectsInnerSignalsInnerEntity struct {
	value *GetInsightsResponseDataObjectsInnerSignalsInnerEntity
	isSet bool
}

func (v NullableGetInsightsResponseDataObjectsInnerSignalsInnerEntity) Get() *GetInsightsResponseDataObjectsInnerSignalsInnerEntity {
	return v.value
}

func (v *NullableGetInsightsResponseDataObjectsInnerSignalsInnerEntity) Set(val *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInsightsResponseDataObjectsInnerSignalsInnerEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInsightsResponseDataObjectsInnerSignalsInnerEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInsightsResponseDataObjectsInnerSignalsInnerEntity(val *GetInsightsResponseDataObjectsInnerSignalsInnerEntity) *NullableGetInsightsResponseDataObjectsInnerSignalsInnerEntity {
	return &NullableGetInsightsResponseDataObjectsInnerSignalsInnerEntity{value: val, isSet: true}
}

func (v NullableGetInsightsResponseDataObjectsInnerSignalsInnerEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInsightsResponseDataObjectsInnerSignalsInnerEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


