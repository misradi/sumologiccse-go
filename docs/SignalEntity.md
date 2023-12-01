# SignalEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**EntityType** | **string** |  | 
**Value** | **string** |  | 
**Hostname** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**SensorZone** | Pointer to **string** |  | [optional] 

## Methods

### NewSignalEntity

`func NewSignalEntity(id string, name string, entityType string, value string, ) *SignalEntity`

NewSignalEntity instantiates a new SignalEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalEntityWithDefaults

`func NewSignalEntityWithDefaults() *SignalEntity`

NewSignalEntityWithDefaults instantiates a new SignalEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SignalEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignalEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignalEntity) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SignalEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignalEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignalEntity) SetName(v string)`

SetName sets Name field to given value.


### GetEntityType

`func (o *SignalEntity) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *SignalEntity) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *SignalEntity) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetValue

`func (o *SignalEntity) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SignalEntity) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SignalEntity) SetValue(v string)`

SetValue sets Value field to given value.


### GetHostname

`func (o *SignalEntity) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SignalEntity) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SignalEntity) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SignalEntity) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetMacAddress

`func (o *SignalEntity) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *SignalEntity) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *SignalEntity) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *SignalEntity) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetSensorZone

`func (o *SignalEntity) GetSensorZone() string`

GetSensorZone returns the SensorZone field if non-nil, zero value otherwise.

### GetSensorZoneOk

`func (o *SignalEntity) GetSensorZoneOk() (*string, bool)`

GetSensorZoneOk returns a tuple with the SensorZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorZone

`func (o *SignalEntity) SetSensorZone(v string)`

SetSensorZone sets SensorZone field to given value.

### HasSensorZone

`func (o *SignalEntity) HasSensorZone() bool`

HasSensorZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


