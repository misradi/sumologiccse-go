# GetEntityResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ActivityScore** | **int32** |  | 
**LastSeen** | Pointer to **time.Time** |  | [optional] 
**FirstSeen** | Pointer to **time.Time** |  | [optional] 
**IsSuppressed** | **bool** |  | 
**Name** | **string** |  | 
**Inventory** | [**[]GetEntitiesResponseDataObjectsInnerInventoryInner**](GetEntitiesResponseDataObjectsInnerInventoryInner.md) |  | 
**Tags** | **[]string** |  | 
**Criticality** | Pointer to **string** |  | [optional] 
**Reputation** | Pointer to **string** |  | [optional] 
**RecentSignalSeverity** | **int32** |  | 
**EntityType** | **string** |  | 
**SensorZone** | Pointer to **string** |  | [optional] 
**Value** | **string** |  | 
**Hostname** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 

## Methods

### NewGetEntityResponseData

`func NewGetEntityResponseData(id string, activityScore int32, isSuppressed bool, name string, inventory []GetEntitiesResponseDataObjectsInnerInventoryInner, tags []string, recentSignalSeverity int32, entityType string, value string, ) *GetEntityResponseData`

NewGetEntityResponseData instantiates a new GetEntityResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntityResponseDataWithDefaults

`func NewGetEntityResponseDataWithDefaults() *GetEntityResponseData`

NewGetEntityResponseDataWithDefaults instantiates a new GetEntityResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetEntityResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEntityResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEntityResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetActivityScore

`func (o *GetEntityResponseData) GetActivityScore() int32`

GetActivityScore returns the ActivityScore field if non-nil, zero value otherwise.

### GetActivityScoreOk

`func (o *GetEntityResponseData) GetActivityScoreOk() (*int32, bool)`

GetActivityScoreOk returns a tuple with the ActivityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityScore

`func (o *GetEntityResponseData) SetActivityScore(v int32)`

SetActivityScore sets ActivityScore field to given value.


### GetLastSeen

`func (o *GetEntityResponseData) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *GetEntityResponseData) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *GetEntityResponseData) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *GetEntityResponseData) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetFirstSeen

`func (o *GetEntityResponseData) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *GetEntityResponseData) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *GetEntityResponseData) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *GetEntityResponseData) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetIsSuppressed

`func (o *GetEntityResponseData) GetIsSuppressed() bool`

GetIsSuppressed returns the IsSuppressed field if non-nil, zero value otherwise.

### GetIsSuppressedOk

`func (o *GetEntityResponseData) GetIsSuppressedOk() (*bool, bool)`

GetIsSuppressedOk returns a tuple with the IsSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuppressed

`func (o *GetEntityResponseData) SetIsSuppressed(v bool)`

SetIsSuppressed sets IsSuppressed field to given value.


### GetName

`func (o *GetEntityResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetEntityResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetEntityResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetInventory

`func (o *GetEntityResponseData) GetInventory() []GetEntitiesResponseDataObjectsInnerInventoryInner`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *GetEntityResponseData) GetInventoryOk() (*[]GetEntitiesResponseDataObjectsInnerInventoryInner, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *GetEntityResponseData) SetInventory(v []GetEntitiesResponseDataObjectsInnerInventoryInner)`

SetInventory sets Inventory field to given value.


### GetTags

`func (o *GetEntityResponseData) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetEntityResponseData) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetEntityResponseData) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetCriticality

`func (o *GetEntityResponseData) GetCriticality() string`

GetCriticality returns the Criticality field if non-nil, zero value otherwise.

### GetCriticalityOk

`func (o *GetEntityResponseData) GetCriticalityOk() (*string, bool)`

GetCriticalityOk returns a tuple with the Criticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticality

`func (o *GetEntityResponseData) SetCriticality(v string)`

SetCriticality sets Criticality field to given value.

### HasCriticality

`func (o *GetEntityResponseData) HasCriticality() bool`

HasCriticality returns a boolean if a field has been set.

### GetReputation

`func (o *GetEntityResponseData) GetReputation() string`

GetReputation returns the Reputation field if non-nil, zero value otherwise.

### GetReputationOk

`func (o *GetEntityResponseData) GetReputationOk() (*string, bool)`

GetReputationOk returns a tuple with the Reputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputation

`func (o *GetEntityResponseData) SetReputation(v string)`

SetReputation sets Reputation field to given value.

### HasReputation

`func (o *GetEntityResponseData) HasReputation() bool`

HasReputation returns a boolean if a field has been set.

### GetRecentSignalSeverity

`func (o *GetEntityResponseData) GetRecentSignalSeverity() int32`

GetRecentSignalSeverity returns the RecentSignalSeverity field if non-nil, zero value otherwise.

### GetRecentSignalSeverityOk

`func (o *GetEntityResponseData) GetRecentSignalSeverityOk() (*int32, bool)`

GetRecentSignalSeverityOk returns a tuple with the RecentSignalSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentSignalSeverity

`func (o *GetEntityResponseData) SetRecentSignalSeverity(v int32)`

SetRecentSignalSeverity sets RecentSignalSeverity field to given value.


### GetEntityType

`func (o *GetEntityResponseData) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *GetEntityResponseData) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *GetEntityResponseData) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetSensorZone

`func (o *GetEntityResponseData) GetSensorZone() string`

GetSensorZone returns the SensorZone field if non-nil, zero value otherwise.

### GetSensorZoneOk

`func (o *GetEntityResponseData) GetSensorZoneOk() (*string, bool)`

GetSensorZoneOk returns a tuple with the SensorZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorZone

`func (o *GetEntityResponseData) SetSensorZone(v string)`

SetSensorZone sets SensorZone field to given value.

### HasSensorZone

`func (o *GetEntityResponseData) HasSensorZone() bool`

HasSensorZone returns a boolean if a field has been set.

### GetValue

`func (o *GetEntityResponseData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetEntityResponseData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetEntityResponseData) SetValue(v string)`

SetValue sets Value field to given value.


### GetHostname

`func (o *GetEntityResponseData) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetEntityResponseData) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetEntityResponseData) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetEntityResponseData) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetMacAddress

`func (o *GetEntityResponseData) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *GetEntityResponseData) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *GetEntityResponseData) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *GetEntityResponseData) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


