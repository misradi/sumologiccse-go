# CreateNetworkBlockRequestBodyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressBlock** | **string** |  | 
**Label** | Pointer to **string** |  | [optional] 
**Internal** | Pointer to **bool** |  | [optional] 
**SuppressesSignals** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateNetworkBlockRequestBodyFields

`func NewCreateNetworkBlockRequestBodyFields(addressBlock string, ) *CreateNetworkBlockRequestBodyFields`

NewCreateNetworkBlockRequestBodyFields instantiates a new CreateNetworkBlockRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkBlockRequestBodyFieldsWithDefaults

`func NewCreateNetworkBlockRequestBodyFieldsWithDefaults() *CreateNetworkBlockRequestBodyFields`

NewCreateNetworkBlockRequestBodyFieldsWithDefaults instantiates a new CreateNetworkBlockRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressBlock

`func (o *CreateNetworkBlockRequestBodyFields) GetAddressBlock() string`

GetAddressBlock returns the AddressBlock field if non-nil, zero value otherwise.

### GetAddressBlockOk

`func (o *CreateNetworkBlockRequestBodyFields) GetAddressBlockOk() (*string, bool)`

GetAddressBlockOk returns a tuple with the AddressBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBlock

`func (o *CreateNetworkBlockRequestBodyFields) SetAddressBlock(v string)`

SetAddressBlock sets AddressBlock field to given value.


### GetLabel

`func (o *CreateNetworkBlockRequestBodyFields) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateNetworkBlockRequestBodyFields) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateNetworkBlockRequestBodyFields) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateNetworkBlockRequestBodyFields) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetInternal

`func (o *CreateNetworkBlockRequestBodyFields) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *CreateNetworkBlockRequestBodyFields) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *CreateNetworkBlockRequestBodyFields) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *CreateNetworkBlockRequestBodyFields) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetSuppressesSignals

`func (o *CreateNetworkBlockRequestBodyFields) GetSuppressesSignals() bool`

GetSuppressesSignals returns the SuppressesSignals field if non-nil, zero value otherwise.

### GetSuppressesSignalsOk

`func (o *CreateNetworkBlockRequestBodyFields) GetSuppressesSignalsOk() (*bool, bool)`

GetSuppressesSignalsOk returns a tuple with the SuppressesSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressesSignals

`func (o *CreateNetworkBlockRequestBodyFields) SetSuppressesSignals(v bool)`

SetSuppressesSignals sets SuppressesSignals field to given value.

### HasSuppressesSignals

`func (o *CreateNetworkBlockRequestBodyFields) HasSuppressesSignals() bool`

HasSuppressesSignals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


