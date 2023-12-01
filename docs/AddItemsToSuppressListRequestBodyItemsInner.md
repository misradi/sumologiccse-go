# AddItemsToSuppressListRequestBodyItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Active** | **bool** |  | 
**Expiration** | Pointer to **time.Time** |  | [optional] 
**Description** | **string** |  | 

## Methods

### NewAddItemsToSuppressListRequestBodyItemsInner

`func NewAddItemsToSuppressListRequestBodyItemsInner(value string, active bool, description string, ) *AddItemsToSuppressListRequestBodyItemsInner`

NewAddItemsToSuppressListRequestBodyItemsInner instantiates a new AddItemsToSuppressListRequestBodyItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddItemsToSuppressListRequestBodyItemsInnerWithDefaults

`func NewAddItemsToSuppressListRequestBodyItemsInnerWithDefaults() *AddItemsToSuppressListRequestBodyItemsInner`

NewAddItemsToSuppressListRequestBodyItemsInnerWithDefaults instantiates a new AddItemsToSuppressListRequestBodyItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *AddItemsToSuppressListRequestBodyItemsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AddItemsToSuppressListRequestBodyItemsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AddItemsToSuppressListRequestBodyItemsInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetActive

`func (o *AddItemsToSuppressListRequestBodyItemsInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddItemsToSuppressListRequestBodyItemsInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddItemsToSuppressListRequestBodyItemsInner) SetActive(v bool)`

SetActive sets Active field to given value.


### GetExpiration

`func (o *AddItemsToSuppressListRequestBodyItemsInner) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *AddItemsToSuppressListRequestBodyItemsInner) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *AddItemsToSuppressListRequestBodyItemsInner) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *AddItemsToSuppressListRequestBodyItemsInner) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetDescription

`func (o *AddItemsToSuppressListRequestBodyItemsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddItemsToSuppressListRequestBodyItemsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddItemsToSuppressListRequestBodyItemsInner) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


