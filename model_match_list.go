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

// checks if the MatchList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MatchList{}

// MatchList struct for MatchList
type MatchList struct {
	Id string `json:"id"`
	// The name of the List.
	Name string `json:"name"`
	// A description of the List.
	Description *string `json:"description,omitempty"`
	// The default time-to-live (in seconds) for new Items added to this List. This default is only used to default the field in the UI, and is not used at all when new Items are added via the API.
	DefaultTtl *int32 `json:"defaultTtl,omitempty"`
	// The column that Items in this List are matched against.
	TargetColumn string `json:"targetColumn"`
	Created *time.Time `json:"created,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
}

type _MatchList MatchList

// NewMatchList instantiates a new MatchList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatchList(id string, name string, targetColumn string) *MatchList {
	this := MatchList{}
	this.Id = id
	this.Name = name
	this.TargetColumn = targetColumn
	return &this
}

// NewMatchListWithDefaults instantiates a new MatchList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchListWithDefaults() *MatchList {
	this := MatchList{}
	return &this
}

// GetId returns the Id field value
func (o *MatchList) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MatchList) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MatchList) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *MatchList) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MatchList) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MatchList) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MatchList) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchList) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MatchList) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MatchList) SetDescription(v string) {
	o.Description = &v
}

// GetDefaultTtl returns the DefaultTtl field value if set, zero value otherwise.
func (o *MatchList) GetDefaultTtl() int32 {
	if o == nil || IsNil(o.DefaultTtl) {
		var ret int32
		return ret
	}
	return *o.DefaultTtl
}

// GetDefaultTtlOk returns a tuple with the DefaultTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchList) GetDefaultTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultTtl) {
		return nil, false
	}
	return o.DefaultTtl, true
}

// HasDefaultTtl returns a boolean if a field has been set.
func (o *MatchList) HasDefaultTtl() bool {
	if o != nil && !IsNil(o.DefaultTtl) {
		return true
	}

	return false
}

// SetDefaultTtl gets a reference to the given int32 and assigns it to the DefaultTtl field.
func (o *MatchList) SetDefaultTtl(v int32) {
	o.DefaultTtl = &v
}

// GetTargetColumn returns the TargetColumn field value
func (o *MatchList) GetTargetColumn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetColumn
}

// GetTargetColumnOk returns a tuple with the TargetColumn field value
// and a boolean to check if the value has been set.
func (o *MatchList) GetTargetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetColumn, true
}

// SetTargetColumn sets field value
func (o *MatchList) SetTargetColumn(v string) {
	o.TargetColumn = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *MatchList) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchList) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *MatchList) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *MatchList) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *MatchList) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchList) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *MatchList) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *MatchList) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *MatchList) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchList) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *MatchList) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *MatchList) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *MatchList) GetLastUpdatedBy() string {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchList) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *MatchList) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *MatchList) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

func (o MatchList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatchList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DefaultTtl) {
		toSerialize["defaultTtl"] = o.DefaultTtl
	}
	toSerialize["targetColumn"] = o.TargetColumn
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.LastUpdatedBy) {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	return toSerialize, nil
}

func (o *MatchList) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"targetColumn",
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

	varMatchList := _MatchList{}

	err = json.Unmarshal(bytes, &varMatchList)

	if err != nil {
		return err
	}

	*o = MatchList(varMatchList)

	return err
}

type NullableMatchList struct {
	value *MatchList
	isSet bool
}

func (v NullableMatchList) Get() *MatchList {
	return v.value
}

func (v *NullableMatchList) Set(val *MatchList) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchList) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchList(val *MatchList) *NullableMatchList {
	return &NullableMatchList{value: val, isSet: true}
}

func (v NullableMatchList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


