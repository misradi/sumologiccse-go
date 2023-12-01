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

// checks if the Inventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Inventory{}

// Inventory struct for Inventory
type Inventory struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	ConfigurationType string `json:"configurationType"`
	Tags []string `json:"tags,omitempty"`
	Criticality *string `json:"criticality,omitempty"`
	Suppressed *bool `json:"suppressed,omitempty"`
	Created time.Time `json:"created"`
	CreatedBy string `json:"createdBy"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	InventoryType *string `json:"inventoryType,omitempty"`
	InventorySource *string `json:"inventorySource,omitempty"`
	Group *string `json:"group,omitempty"`
	InventoryKey *string `json:"inventoryKey,omitempty"`
	InventoryValue *string `json:"inventoryValue,omitempty"`
	TagSchema *string `json:"tagSchema,omitempty"`
	DynamicTags *bool `json:"dynamicTags,omitempty"`
}

type _Inventory Inventory

// NewInventory instantiates a new Inventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventory(id string, name string, configurationType string, created time.Time, createdBy string) *Inventory {
	this := Inventory{}
	this.Id = id
	this.Name = name
	this.ConfigurationType = configurationType
	this.Created = created
	this.CreatedBy = createdBy
	return &this
}

// NewInventoryWithDefaults instantiates a new Inventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryWithDefaults() *Inventory {
	this := Inventory{}
	return &this
}

// GetId returns the Id field value
func (o *Inventory) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Inventory) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Inventory) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Inventory) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Inventory) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Inventory) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Inventory) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Inventory) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Inventory) SetDescription(v string) {
	o.Description = &v
}

// GetConfigurationType returns the ConfigurationType field value
func (o *Inventory) GetConfigurationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigurationType
}

// GetConfigurationTypeOk returns a tuple with the ConfigurationType field value
// and a boolean to check if the value has been set.
func (o *Inventory) GetConfigurationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationType, true
}

// SetConfigurationType sets field value
func (o *Inventory) SetConfigurationType(v string) {
	o.ConfigurationType = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Inventory) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Inventory) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Inventory) SetTags(v []string) {
	o.Tags = v
}

// GetCriticality returns the Criticality field value if set, zero value otherwise.
func (o *Inventory) GetCriticality() string {
	if o == nil || IsNil(o.Criticality) {
		var ret string
		return ret
	}
	return *o.Criticality
}

// GetCriticalityOk returns a tuple with the Criticality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetCriticalityOk() (*string, bool) {
	if o == nil || IsNil(o.Criticality) {
		return nil, false
	}
	return o.Criticality, true
}

// HasCriticality returns a boolean if a field has been set.
func (o *Inventory) HasCriticality() bool {
	if o != nil && !IsNil(o.Criticality) {
		return true
	}

	return false
}

// SetCriticality gets a reference to the given string and assigns it to the Criticality field.
func (o *Inventory) SetCriticality(v string) {
	o.Criticality = &v
}

// GetSuppressed returns the Suppressed field value if set, zero value otherwise.
func (o *Inventory) GetSuppressed() bool {
	if o == nil || IsNil(o.Suppressed) {
		var ret bool
		return ret
	}
	return *o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetSuppressedOk() (*bool, bool) {
	if o == nil || IsNil(o.Suppressed) {
		return nil, false
	}
	return o.Suppressed, true
}

// HasSuppressed returns a boolean if a field has been set.
func (o *Inventory) HasSuppressed() bool {
	if o != nil && !IsNil(o.Suppressed) {
		return true
	}

	return false
}

// SetSuppressed gets a reference to the given bool and assigns it to the Suppressed field.
func (o *Inventory) SetSuppressed(v bool) {
	o.Suppressed = &v
}

// GetCreated returns the Created field value
func (o *Inventory) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Inventory) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Inventory) SetCreated(v time.Time) {
	o.Created = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *Inventory) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *Inventory) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *Inventory) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Inventory) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Inventory) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Inventory) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *Inventory) GetLastUpdatedBy() string {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *Inventory) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *Inventory) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetInventoryType returns the InventoryType field value if set, zero value otherwise.
func (o *Inventory) GetInventoryType() string {
	if o == nil || IsNil(o.InventoryType) {
		var ret string
		return ret
	}
	return *o.InventoryType
}

// GetInventoryTypeOk returns a tuple with the InventoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetInventoryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InventoryType) {
		return nil, false
	}
	return o.InventoryType, true
}

// HasInventoryType returns a boolean if a field has been set.
func (o *Inventory) HasInventoryType() bool {
	if o != nil && !IsNil(o.InventoryType) {
		return true
	}

	return false
}

// SetInventoryType gets a reference to the given string and assigns it to the InventoryType field.
func (o *Inventory) SetInventoryType(v string) {
	o.InventoryType = &v
}

// GetInventorySource returns the InventorySource field value if set, zero value otherwise.
func (o *Inventory) GetInventorySource() string {
	if o == nil || IsNil(o.InventorySource) {
		var ret string
		return ret
	}
	return *o.InventorySource
}

// GetInventorySourceOk returns a tuple with the InventorySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetInventorySourceOk() (*string, bool) {
	if o == nil || IsNil(o.InventorySource) {
		return nil, false
	}
	return o.InventorySource, true
}

// HasInventorySource returns a boolean if a field has been set.
func (o *Inventory) HasInventorySource() bool {
	if o != nil && !IsNil(o.InventorySource) {
		return true
	}

	return false
}

// SetInventorySource gets a reference to the given string and assigns it to the InventorySource field.
func (o *Inventory) SetInventorySource(v string) {
	o.InventorySource = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *Inventory) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *Inventory) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *Inventory) SetGroup(v string) {
	o.Group = &v
}

// GetInventoryKey returns the InventoryKey field value if set, zero value otherwise.
func (o *Inventory) GetInventoryKey() string {
	if o == nil || IsNil(o.InventoryKey) {
		var ret string
		return ret
	}
	return *o.InventoryKey
}

// GetInventoryKeyOk returns a tuple with the InventoryKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetInventoryKeyOk() (*string, bool) {
	if o == nil || IsNil(o.InventoryKey) {
		return nil, false
	}
	return o.InventoryKey, true
}

// HasInventoryKey returns a boolean if a field has been set.
func (o *Inventory) HasInventoryKey() bool {
	if o != nil && !IsNil(o.InventoryKey) {
		return true
	}

	return false
}

// SetInventoryKey gets a reference to the given string and assigns it to the InventoryKey field.
func (o *Inventory) SetInventoryKey(v string) {
	o.InventoryKey = &v
}

// GetInventoryValue returns the InventoryValue field value if set, zero value otherwise.
func (o *Inventory) GetInventoryValue() string {
	if o == nil || IsNil(o.InventoryValue) {
		var ret string
		return ret
	}
	return *o.InventoryValue
}

// GetInventoryValueOk returns a tuple with the InventoryValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetInventoryValueOk() (*string, bool) {
	if o == nil || IsNil(o.InventoryValue) {
		return nil, false
	}
	return o.InventoryValue, true
}

// HasInventoryValue returns a boolean if a field has been set.
func (o *Inventory) HasInventoryValue() bool {
	if o != nil && !IsNil(o.InventoryValue) {
		return true
	}

	return false
}

// SetInventoryValue gets a reference to the given string and assigns it to the InventoryValue field.
func (o *Inventory) SetInventoryValue(v string) {
	o.InventoryValue = &v
}

// GetTagSchema returns the TagSchema field value if set, zero value otherwise.
func (o *Inventory) GetTagSchema() string {
	if o == nil || IsNil(o.TagSchema) {
		var ret string
		return ret
	}
	return *o.TagSchema
}

// GetTagSchemaOk returns a tuple with the TagSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetTagSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.TagSchema) {
		return nil, false
	}
	return o.TagSchema, true
}

// HasTagSchema returns a boolean if a field has been set.
func (o *Inventory) HasTagSchema() bool {
	if o != nil && !IsNil(o.TagSchema) {
		return true
	}

	return false
}

// SetTagSchema gets a reference to the given string and assigns it to the TagSchema field.
func (o *Inventory) SetTagSchema(v string) {
	o.TagSchema = &v
}

// GetDynamicTags returns the DynamicTags field value if set, zero value otherwise.
func (o *Inventory) GetDynamicTags() bool {
	if o == nil || IsNil(o.DynamicTags) {
		var ret bool
		return ret
	}
	return *o.DynamicTags
}

// GetDynamicTagsOk returns a tuple with the DynamicTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetDynamicTagsOk() (*bool, bool) {
	if o == nil || IsNil(o.DynamicTags) {
		return nil, false
	}
	return o.DynamicTags, true
}

// HasDynamicTags returns a boolean if a field has been set.
func (o *Inventory) HasDynamicTags() bool {
	if o != nil && !IsNil(o.DynamicTags) {
		return true
	}

	return false
}

// SetDynamicTags gets a reference to the given bool and assigns it to the DynamicTags field.
func (o *Inventory) SetDynamicTags(v bool) {
	o.DynamicTags = &v
}

func (o Inventory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Inventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["configurationType"] = o.ConfigurationType
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Criticality) {
		toSerialize["criticality"] = o.Criticality
	}
	if !IsNil(o.Suppressed) {
		toSerialize["suppressed"] = o.Suppressed
	}
	toSerialize["created"] = o.Created
	toSerialize["createdBy"] = o.CreatedBy
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.LastUpdatedBy) {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	if !IsNil(o.InventoryType) {
		toSerialize["inventoryType"] = o.InventoryType
	}
	if !IsNil(o.InventorySource) {
		toSerialize["inventorySource"] = o.InventorySource
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.InventoryKey) {
		toSerialize["inventoryKey"] = o.InventoryKey
	}
	if !IsNil(o.InventoryValue) {
		toSerialize["inventoryValue"] = o.InventoryValue
	}
	if !IsNil(o.TagSchema) {
		toSerialize["tagSchema"] = o.TagSchema
	}
	if !IsNil(o.DynamicTags) {
		toSerialize["dynamicTags"] = o.DynamicTags
	}
	return toSerialize, nil
}

func (o *Inventory) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"configurationType",
		"created",
		"createdBy",
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

	varInventory := _Inventory{}

	err = json.Unmarshal(bytes, &varInventory)

	if err != nil {
		return err
	}

	*o = Inventory(varInventory)

	return err
}

type NullableInventory struct {
	value *Inventory
	isSet bool
}

func (v NullableInventory) Get() *Inventory {
	return v.value
}

func (v *NullableInventory) Set(val *Inventory) {
	v.value = val
	v.isSet = true
}

func (v NullableInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventory(val *Inventory) *NullableInventory {
	return &NullableInventory{value: val, isSet: true}
}

func (v NullableInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


