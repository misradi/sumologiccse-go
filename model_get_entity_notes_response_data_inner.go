/*
Sumo Logic CSE API

 https://help.sumologic.com/APIs 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologiccse

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the GetEntityNotesResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEntityNotesResponseDataInner{}

// GetEntityNotesResponseDataInner struct for GetEntityNotesResponseDataInner
type GetEntityNotesResponseDataInner struct {
	Id string `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	LastUpdated time.Time `json:"lastUpdated"`
	Note string `json:"note"`
	Author *GetEntityNotesResponseDataInnerAuthor `json:"author,omitempty"`
}

type _GetEntityNotesResponseDataInner GetEntityNotesResponseDataInner

// NewGetEntityNotesResponseDataInner instantiates a new GetEntityNotesResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEntityNotesResponseDataInner(id string, timestamp time.Time, lastUpdated time.Time, note string) *GetEntityNotesResponseDataInner {
	this := GetEntityNotesResponseDataInner{}
	this.Id = id
	this.Timestamp = timestamp
	this.LastUpdated = lastUpdated
	this.Note = note
	return &this
}

// NewGetEntityNotesResponseDataInnerWithDefaults instantiates a new GetEntityNotesResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEntityNotesResponseDataInnerWithDefaults() *GetEntityNotesResponseDataInner {
	this := GetEntityNotesResponseDataInner{}
	return &this
}

// GetId returns the Id field value
func (o *GetEntityNotesResponseDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetEntityNotesResponseDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetEntityNotesResponseDataInner) SetId(v string) {
	o.Id = v
}

// GetTimestamp returns the Timestamp field value
func (o *GetEntityNotesResponseDataInner) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *GetEntityNotesResponseDataInner) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *GetEntityNotesResponseDataInner) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *GetEntityNotesResponseDataInner) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *GetEntityNotesResponseDataInner) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *GetEntityNotesResponseDataInner) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetNote returns the Note field value
func (o *GetEntityNotesResponseDataInner) GetNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Note
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
func (o *GetEntityNotesResponseDataInner) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Note, true
}

// SetNote sets field value
func (o *GetEntityNotesResponseDataInner) SetNote(v string) {
	o.Note = v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *GetEntityNotesResponseDataInner) GetAuthor() GetEntityNotesResponseDataInnerAuthor {
	if o == nil || IsNil(o.Author) {
		var ret GetEntityNotesResponseDataInnerAuthor
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEntityNotesResponseDataInner) GetAuthorOk() (*GetEntityNotesResponseDataInnerAuthor, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *GetEntityNotesResponseDataInner) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given GetEntityNotesResponseDataInnerAuthor and assigns it to the Author field.
func (o *GetEntityNotesResponseDataInner) SetAuthor(v GetEntityNotesResponseDataInnerAuthor) {
	o.Author = &v
}

func (o GetEntityNotesResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEntityNotesResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["note"] = o.Note
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	return toSerialize, nil
}

func (o *GetEntityNotesResponseDataInner) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"timestamp",
		"lastUpdated",
		"note",
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

	varGetEntityNotesResponseDataInner := _GetEntityNotesResponseDataInner{}

	err = json.Unmarshal(bytes, &varGetEntityNotesResponseDataInner)

	if err != nil {
		return err
	}

	*o = GetEntityNotesResponseDataInner(varGetEntityNotesResponseDataInner)

	return err
}

type NullableGetEntityNotesResponseDataInner struct {
	value *GetEntityNotesResponseDataInner
	isSet bool
}

func (v NullableGetEntityNotesResponseDataInner) Get() *GetEntityNotesResponseDataInner {
	return v.value
}

func (v *NullableGetEntityNotesResponseDataInner) Set(val *GetEntityNotesResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEntityNotesResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEntityNotesResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEntityNotesResponseDataInner(val *GetEntityNotesResponseDataInner) *NullableGetEntityNotesResponseDataInner {
	return &NullableGetEntityNotesResponseDataInner{value: val, isSet: true}
}

func (v NullableGetEntityNotesResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEntityNotesResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


