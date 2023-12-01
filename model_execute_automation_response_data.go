/*
Sumo Logic CSE API

 https://help.sumologic.com/APIs 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologiccse

import (
	"encoding/json"
)

// checks if the ExecuteAutomationResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecuteAutomationResponseData{}

// ExecuteAutomationResponseData struct for ExecuteAutomationResponseData
type ExecuteAutomationResponseData struct {
	ExecutionIds []string `json:"executionIds,omitempty"`
	PlaybooksWithErrors []string `json:"playbooksWithErrors,omitempty"`
}

// NewExecuteAutomationResponseData instantiates a new ExecuteAutomationResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecuteAutomationResponseData() *ExecuteAutomationResponseData {
	this := ExecuteAutomationResponseData{}
	return &this
}

// NewExecuteAutomationResponseDataWithDefaults instantiates a new ExecuteAutomationResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecuteAutomationResponseDataWithDefaults() *ExecuteAutomationResponseData {
	this := ExecuteAutomationResponseData{}
	return &this
}

// GetExecutionIds returns the ExecutionIds field value if set, zero value otherwise.
func (o *ExecuteAutomationResponseData) GetExecutionIds() []string {
	if o == nil || IsNil(o.ExecutionIds) {
		var ret []string
		return ret
	}
	return o.ExecutionIds
}

// GetExecutionIdsOk returns a tuple with the ExecutionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteAutomationResponseData) GetExecutionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExecutionIds) {
		return nil, false
	}
	return o.ExecutionIds, true
}

// HasExecutionIds returns a boolean if a field has been set.
func (o *ExecuteAutomationResponseData) HasExecutionIds() bool {
	if o != nil && !IsNil(o.ExecutionIds) {
		return true
	}

	return false
}

// SetExecutionIds gets a reference to the given []string and assigns it to the ExecutionIds field.
func (o *ExecuteAutomationResponseData) SetExecutionIds(v []string) {
	o.ExecutionIds = v
}

// GetPlaybooksWithErrors returns the PlaybooksWithErrors field value if set, zero value otherwise.
func (o *ExecuteAutomationResponseData) GetPlaybooksWithErrors() []string {
	if o == nil || IsNil(o.PlaybooksWithErrors) {
		var ret []string
		return ret
	}
	return o.PlaybooksWithErrors
}

// GetPlaybooksWithErrorsOk returns a tuple with the PlaybooksWithErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteAutomationResponseData) GetPlaybooksWithErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.PlaybooksWithErrors) {
		return nil, false
	}
	return o.PlaybooksWithErrors, true
}

// HasPlaybooksWithErrors returns a boolean if a field has been set.
func (o *ExecuteAutomationResponseData) HasPlaybooksWithErrors() bool {
	if o != nil && !IsNil(o.PlaybooksWithErrors) {
		return true
	}

	return false
}

// SetPlaybooksWithErrors gets a reference to the given []string and assigns it to the PlaybooksWithErrors field.
func (o *ExecuteAutomationResponseData) SetPlaybooksWithErrors(v []string) {
	o.PlaybooksWithErrors = v
}

func (o ExecuteAutomationResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecuteAutomationResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExecutionIds) {
		toSerialize["executionIds"] = o.ExecutionIds
	}
	if !IsNil(o.PlaybooksWithErrors) {
		toSerialize["playbooksWithErrors"] = o.PlaybooksWithErrors
	}
	return toSerialize, nil
}

type NullableExecuteAutomationResponseData struct {
	value *ExecuteAutomationResponseData
	isSet bool
}

func (v NullableExecuteAutomationResponseData) Get() *ExecuteAutomationResponseData {
	return v.value
}

func (v *NullableExecuteAutomationResponseData) Set(val *ExecuteAutomationResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteAutomationResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteAutomationResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteAutomationResponseData(val *ExecuteAutomationResponseData) *NullableExecuteAutomationResponseData {
	return &NullableExecuteAutomationResponseData{value: val, isSet: true}
}

func (v NullableExecuteAutomationResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteAutomationResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


