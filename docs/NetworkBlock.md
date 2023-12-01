# NetworkBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AddressBlock** | **string** |  | 
**Label** | **string** | The name of the List. | 
**Internal** | **bool** |  | 
**SuppressesSignals** | **bool** |  | 

## Methods

### NewNetworkBlock

`func NewNetworkBlock(id string, addressBlock string, label string, internal bool, suppressesSignals bool, ) *NetworkBlock`

NewNetworkBlock instantiates a new NetworkBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkBlockWithDefaults

`func NewNetworkBlockWithDefaults() *NetworkBlock`

NewNetworkBlockWithDefaults instantiates a new NetworkBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkBlock) SetId(v string)`

SetId sets Id field to given value.


### GetAddressBlock

`func (o *NetworkBlock) GetAddressBlock() string`

GetAddressBlock returns the AddressBlock field if non-nil, zero value otherwise.

### GetAddressBlockOk

`func (o *NetworkBlock) GetAddressBlockOk() (*string, bool)`

GetAddressBlockOk returns a tuple with the AddressBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBlock

`func (o *NetworkBlock) SetAddressBlock(v string)`

SetAddressBlock sets AddressBlock field to given value.


### GetLabel

`func (o *NetworkBlock) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *NetworkBlock) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *NetworkBlock) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetInternal

`func (o *NetworkBlock) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *NetworkBlock) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *NetworkBlock) SetInternal(v bool)`

SetInternal sets Internal field to given value.


### GetSuppressesSignals

`func (o *NetworkBlock) GetSuppressesSignals() bool`

GetSuppressesSignals returns the SuppressesSignals field if non-nil, zero value otherwise.

### GetSuppressesSignalsOk

`func (o *NetworkBlock) GetSuppressesSignalsOk() (*bool, bool)`

GetSuppressesSignalsOk returns a tuple with the SuppressesSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressesSignals

`func (o *NetworkBlock) SetSuppressesSignals(v bool)`

SetSuppressesSignals sets SuppressesSignals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


