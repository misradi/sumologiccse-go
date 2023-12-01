# UpdateSuppressListItemRequestBodyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** |  | 
**Expiration** | Pointer to **time.Time** |  | [optional] 
**Description** | **string** |  | 

## Methods

### NewUpdateSuppressListItemRequestBodyFields

`func NewUpdateSuppressListItemRequestBodyFields(active bool, description string, ) *UpdateSuppressListItemRequestBodyFields`

NewUpdateSuppressListItemRequestBodyFields instantiates a new UpdateSuppressListItemRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSuppressListItemRequestBodyFieldsWithDefaults

`func NewUpdateSuppressListItemRequestBodyFieldsWithDefaults() *UpdateSuppressListItemRequestBodyFields`

NewUpdateSuppressListItemRequestBodyFieldsWithDefaults instantiates a new UpdateSuppressListItemRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdateSuppressListItemRequestBodyFields) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateSuppressListItemRequestBodyFields) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateSuppressListItemRequestBodyFields) SetActive(v bool)`

SetActive sets Active field to given value.


### GetExpiration

`func (o *UpdateSuppressListItemRequestBodyFields) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *UpdateSuppressListItemRequestBodyFields) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *UpdateSuppressListItemRequestBodyFields) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *UpdateSuppressListItemRequestBodyFields) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSuppressListItemRequestBodyFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSuppressListItemRequestBodyFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSuppressListItemRequestBodyFields) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


