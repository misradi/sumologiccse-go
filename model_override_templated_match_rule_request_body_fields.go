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

// checks if the OverrideTemplatedMatchRuleRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OverrideTemplatedMatchRuleRequestBodyFields{}

// OverrideTemplatedMatchRuleRequestBodyFields struct for OverrideTemplatedMatchRuleRequestBodyFields
type OverrideTemplatedMatchRuleRequestBodyFields struct {
	TuningExpressionIds []string `json:"tuningExpressionIds"`
	IsPrototype *bool `json:"isPrototype,omitempty"`
	EntitySelectors []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner `json:"entitySelectors,omitempty"`
	Name *string `json:"name,omitempty"`
	RuleDescription *string `json:"ruleDescription,omitempty"`
	SummaryExpression *string `json:"summaryExpression,omitempty"`
	Tags []string `json:"tags,omitempty"`
	DescriptionExpression *string `json:"descriptionExpression,omitempty"`
	NameExpression *string `json:"nameExpression,omitempty"`
	ScoreMapping *GetRulesResponseDataObjectsInnerOneOf2ScoreMapping `json:"scoreMapping,omitempty"`
}

type _OverrideTemplatedMatchRuleRequestBodyFields OverrideTemplatedMatchRuleRequestBodyFields

// NewOverrideTemplatedMatchRuleRequestBodyFields instantiates a new OverrideTemplatedMatchRuleRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverrideTemplatedMatchRuleRequestBodyFields(tuningExpressionIds []string) *OverrideTemplatedMatchRuleRequestBodyFields {
	this := OverrideTemplatedMatchRuleRequestBodyFields{}
	this.TuningExpressionIds = tuningExpressionIds
	return &this
}

// NewOverrideTemplatedMatchRuleRequestBodyFieldsWithDefaults instantiates a new OverrideTemplatedMatchRuleRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverrideTemplatedMatchRuleRequestBodyFieldsWithDefaults() *OverrideTemplatedMatchRuleRequestBodyFields {
	this := OverrideTemplatedMatchRuleRequestBodyFields{}
	return &this
}

// GetTuningExpressionIds returns the TuningExpressionIds field value
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetTuningExpressionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TuningExpressionIds
}

// GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field value
// and a boolean to check if the value has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetTuningExpressionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TuningExpressionIds, true
}

// SetTuningExpressionIds sets field value
func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetTuningExpressionIds(v []string) {
	o.TuningExpressionIds = v
}

// GetIsPrototype returns the IsPrototype field value if set, zero value otherwise.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetIsPrototype() bool {
	if o == nil || IsNil(o.IsPrototype) {
		var ret bool
		return ret
	}
	return *o.IsPrototype
}

// GetIsPrototypeOk returns a tuple with the IsPrototype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrototype) {
		return nil, false
	}
	return o.IsPrototype, true
}

// HasIsPrototype returns a boolean if a field has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) HasIsPrototype() bool {
	if o != nil && !IsNil(o.IsPrototype) {
		return true
	}

	return false
}

// SetIsPrototype gets a reference to the given bool and assigns it to the IsPrototype field.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetIsPrototype(v bool) {
	o.IsPrototype = &v
}

// GetEntitySelectors returns the EntitySelectors field value if set, zero value otherwise.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner {
	if o == nil || IsNil(o.EntitySelectors) {
		var ret []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner
		return ret
	}
	return o.EntitySelectors
}

// GetEntitySelectorsOk returns a tuple with the EntitySelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetEntitySelectorsOk() ([]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool) {
	if o == nil || IsNil(o.EntitySelectors) {
		return nil, false
	}
	return o.EntitySelectors, true
}

// HasEntitySelectors returns a boolean if a field has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) HasEntitySelectors() bool {
	if o != nil && !IsNil(o.EntitySelectors) {
		return true
	}

	return false
}

// SetEntitySelectors gets a reference to the given []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner and assigns it to the EntitySelectors field.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner) {
	o.EntitySelectors = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetName(v string) {
	o.Name = &v
}

// GetRuleDescription returns the RuleDescription field value if set, zero value otherwise.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetRuleDescription() string {
	if o == nil || IsNil(o.RuleDescription) {
		var ret string
		return ret
	}
	return *o.RuleDescription
}

// GetRuleDescriptionOk returns a tuple with the RuleDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetRuleDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.RuleDescription) {
		return nil, false
	}
	return o.RuleDescription, true
}

// HasRuleDescription returns a boolean if a field has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) HasRuleDescription() bool {
	if o != nil && !IsNil(o.RuleDescription) {
		return true
	}

	return false
}

// SetRuleDescription gets a reference to the given string and assigns it to the RuleDescription field.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetRuleDescription(v string) {
	o.RuleDescription = &v
}

// GetSummaryExpression returns the SummaryExpression field value if set, zero value otherwise.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetSummaryExpression() string {
	if o == nil || IsNil(o.SummaryExpression) {
		var ret string
		return ret
	}
	return *o.SummaryExpression
}

// GetSummaryExpressionOk returns a tuple with the SummaryExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.SummaryExpression) {
		return nil, false
	}
	return o.SummaryExpression, true
}

// HasSummaryExpression returns a boolean if a field has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) HasSummaryExpression() bool {
	if o != nil && !IsNil(o.SummaryExpression) {
		return true
	}

	return false
}

// SetSummaryExpression gets a reference to the given string and assigns it to the SummaryExpression field.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetSummaryExpression(v string) {
	o.SummaryExpression = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetTags(v []string) {
	o.Tags = v
}

// GetDescriptionExpression returns the DescriptionExpression field value if set, zero value otherwise.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetDescriptionExpression() string {
	if o == nil || IsNil(o.DescriptionExpression) {
		var ret string
		return ret
	}
	return *o.DescriptionExpression
}

// GetDescriptionExpressionOk returns a tuple with the DescriptionExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetDescriptionExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.DescriptionExpression) {
		return nil, false
	}
	return o.DescriptionExpression, true
}

// HasDescriptionExpression returns a boolean if a field has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) HasDescriptionExpression() bool {
	if o != nil && !IsNil(o.DescriptionExpression) {
		return true
	}

	return false
}

// SetDescriptionExpression gets a reference to the given string and assigns it to the DescriptionExpression field.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetDescriptionExpression(v string) {
	o.DescriptionExpression = &v
}

// GetNameExpression returns the NameExpression field value if set, zero value otherwise.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetNameExpression() string {
	if o == nil || IsNil(o.NameExpression) {
		var ret string
		return ret
	}
	return *o.NameExpression
}

// GetNameExpressionOk returns a tuple with the NameExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetNameExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.NameExpression) {
		return nil, false
	}
	return o.NameExpression, true
}

// HasNameExpression returns a boolean if a field has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) HasNameExpression() bool {
	if o != nil && !IsNil(o.NameExpression) {
		return true
	}

	return false
}

// SetNameExpression gets a reference to the given string and assigns it to the NameExpression field.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetNameExpression(v string) {
	o.NameExpression = &v
}

// GetScoreMapping returns the ScoreMapping field value if set, zero value otherwise.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetScoreMapping() GetRulesResponseDataObjectsInnerOneOf2ScoreMapping {
	if o == nil || IsNil(o.ScoreMapping) {
		var ret GetRulesResponseDataObjectsInnerOneOf2ScoreMapping
		return ret
	}
	return *o.ScoreMapping
}

// GetScoreMappingOk returns a tuple with the ScoreMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) GetScoreMappingOk() (*GetRulesResponseDataObjectsInnerOneOf2ScoreMapping, bool) {
	if o == nil || IsNil(o.ScoreMapping) {
		return nil, false
	}
	return o.ScoreMapping, true
}

// HasScoreMapping returns a boolean if a field has been set.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) HasScoreMapping() bool {
	if o != nil && !IsNil(o.ScoreMapping) {
		return true
	}

	return false
}

// SetScoreMapping gets a reference to the given GetRulesResponseDataObjectsInnerOneOf2ScoreMapping and assigns it to the ScoreMapping field.
func (o *OverrideTemplatedMatchRuleRequestBodyFields) SetScoreMapping(v GetRulesResponseDataObjectsInnerOneOf2ScoreMapping) {
	o.ScoreMapping = &v
}

func (o OverrideTemplatedMatchRuleRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverrideTemplatedMatchRuleRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tuningExpressionIds"] = o.TuningExpressionIds
	if !IsNil(o.IsPrototype) {
		toSerialize["isPrototype"] = o.IsPrototype
	}
	if !IsNil(o.EntitySelectors) {
		toSerialize["entitySelectors"] = o.EntitySelectors
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RuleDescription) {
		toSerialize["ruleDescription"] = o.RuleDescription
	}
	if !IsNil(o.SummaryExpression) {
		toSerialize["summaryExpression"] = o.SummaryExpression
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.DescriptionExpression) {
		toSerialize["descriptionExpression"] = o.DescriptionExpression
	}
	if !IsNil(o.NameExpression) {
		toSerialize["nameExpression"] = o.NameExpression
	}
	if !IsNil(o.ScoreMapping) {
		toSerialize["scoreMapping"] = o.ScoreMapping
	}
	return toSerialize, nil
}

func (o *OverrideTemplatedMatchRuleRequestBodyFields) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tuningExpressionIds",
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

	varOverrideTemplatedMatchRuleRequestBodyFields := _OverrideTemplatedMatchRuleRequestBodyFields{}

	err = json.Unmarshal(bytes, &varOverrideTemplatedMatchRuleRequestBodyFields)

	if err != nil {
		return err
	}

	*o = OverrideTemplatedMatchRuleRequestBodyFields(varOverrideTemplatedMatchRuleRequestBodyFields)

	return err
}

type NullableOverrideTemplatedMatchRuleRequestBodyFields struct {
	value *OverrideTemplatedMatchRuleRequestBodyFields
	isSet bool
}

func (v NullableOverrideTemplatedMatchRuleRequestBodyFields) Get() *OverrideTemplatedMatchRuleRequestBodyFields {
	return v.value
}

func (v *NullableOverrideTemplatedMatchRuleRequestBodyFields) Set(val *OverrideTemplatedMatchRuleRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideTemplatedMatchRuleRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideTemplatedMatchRuleRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideTemplatedMatchRuleRequestBodyFields(val *OverrideTemplatedMatchRuleRequestBodyFields) *NullableOverrideTemplatedMatchRuleRequestBodyFields {
	return &NullableOverrideTemplatedMatchRuleRequestBodyFields{value: val, isSet: true}
}

func (v NullableOverrideTemplatedMatchRuleRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideTemplatedMatchRuleRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


