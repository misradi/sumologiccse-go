# UpdateEntityCriticalityResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ActivityScore** | **int32** |  | 
**LastSeen** | Pointer to **time.Time** |  | [optional] 
**FirstSeen** | Pointer to **time.Time** |  | [optional] 
**IsSuppressed** | **bool** |  | 
**Name** | **string** |  | 
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

### NewUpdateEntityCriticalityResponseData

`func NewUpdateEntityCriticalityResponseData(id string, activityScore int32, isSuppressed bool, name string, tags []string, recentSignalSeverity int32, entityType string, value string, ) *UpdateEntityCriticalityResponseData`

NewUpdateEntityCriticalityResponseData instantiates a new UpdateEntityCriticalityResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEntityCriticalityResponseDataWithDefaults

`func NewUpdateEntityCriticalityResponseDataWithDefaults() *UpdateEntityCriticalityResponseData`

NewUpdateEntityCriticalityResponseDataWithDefaults instantiates a new UpdateEntityCriticalityResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateEntityCriticalityResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateEntityCriticalityResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateEntityCriticalityResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetActivityScore

`func (o *UpdateEntityCriticalityResponseData) GetActivityScore() int32`

GetActivityScore returns the ActivityScore field if non-nil, zero value otherwise.

### GetActivityScoreOk

`func (o *UpdateEntityCriticalityResponseData) GetActivityScoreOk() (*int32, bool)`

GetActivityScoreOk returns a tuple with the ActivityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityScore

`func (o *UpdateEntityCriticalityResponseData) SetActivityScore(v int32)`

SetActivityScore sets ActivityScore field to given value.


### GetLastSeen

`func (o *UpdateEntityCriticalityResponseData) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *UpdateEntityCriticalityResponseData) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *UpdateEntityCriticalityResponseData) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *UpdateEntityCriticalityResponseData) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetFirstSeen

`func (o *UpdateEntityCriticalityResponseData) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *UpdateEntityCriticalityResponseData) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *UpdateEntityCriticalityResponseData) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *UpdateEntityCriticalityResponseData) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetIsSuppressed

`func (o *UpdateEntityCriticalityResponseData) GetIsSuppressed() bool`

GetIsSuppressed returns the IsSuppressed field if non-nil, zero value otherwise.

### GetIsSuppressedOk

`func (o *UpdateEntityCriticalityResponseData) GetIsSuppressedOk() (*bool, bool)`

GetIsSuppressedOk returns a tuple with the IsSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuppressed

`func (o *UpdateEntityCriticalityResponseData) SetIsSuppressed(v bool)`

SetIsSuppressed sets IsSuppressed field to given value.


### GetName

`func (o *UpdateEntityCriticalityResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateEntityCriticalityResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateEntityCriticalityResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *UpdateEntityCriticalityResponseData) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateEntityCriticalityResponseData) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateEntityCriticalityResponseData) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetCriticality

`func (o *UpdateEntityCriticalityResponseData) GetCriticality() string`

GetCriticality returns the Criticality field if non-nil, zero value otherwise.

### GetCriticalityOk

`func (o *UpdateEntityCriticalityResponseData) GetCriticalityOk() (*string, bool)`

GetCriticalityOk returns a tuple with the Criticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticality

`func (o *UpdateEntityCriticalityResponseData) SetCriticality(v string)`

SetCriticality sets Criticality field to given value.

### HasCriticality

`func (o *UpdateEntityCriticalityResponseData) HasCriticality() bool`

HasCriticality returns a boolean if a field has been set.

### GetReputation

`func (o *UpdateEntityCriticalityResponseData) GetReputation() string`

GetReputation returns the Reputation field if non-nil, zero value otherwise.

### GetReputationOk

`func (o *UpdateEntityCriticalityResponseData) GetReputationOk() (*string, bool)`

GetReputationOk returns a tuple with the Reputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputation

`func (o *UpdateEntityCriticalityResponseData) SetReputation(v string)`

SetReputation sets Reputation field to given value.

### HasReputation

`func (o *UpdateEntityCriticalityResponseData) HasReputation() bool`

HasReputation returns a boolean if a field has been set.

### GetRecentSignalSeverity

`func (o *UpdateEntityCriticalityResponseData) GetRecentSignalSeverity() int32`

GetRecentSignalSeverity returns the RecentSignalSeverity field if non-nil, zero value otherwise.

### GetRecentSignalSeverityOk

`func (o *UpdateEntityCriticalityResponseData) GetRecentSignalSeverityOk() (*int32, bool)`

GetRecentSignalSeverityOk returns a tuple with the RecentSignalSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentSignalSeverity

`func (o *UpdateEntityCriticalityResponseData) SetRecentSignalSeverity(v int32)`

SetRecentSignalSeverity sets RecentSignalSeverity field to given value.


### GetEntityType

`func (o *UpdateEntityCriticalityResponseData) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *UpdateEntityCriticalityResponseData) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *UpdateEntityCriticalityResponseData) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetSensorZone

`func (o *UpdateEntityCriticalityResponseData) GetSensorZone() string`

GetSensorZone returns the SensorZone field if non-nil, zero value otherwise.

### GetSensorZoneOk

`func (o *UpdateEntityCriticalityResponseData) GetSensorZoneOk() (*string, bool)`

GetSensorZoneOk returns a tuple with the SensorZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorZone

`func (o *UpdateEntityCriticalityResponseData) SetSensorZone(v string)`

SetSensorZone sets SensorZone field to given value.

### HasSensorZone

`func (o *UpdateEntityCriticalityResponseData) HasSensorZone() bool`

HasSensorZone returns a boolean if a field has been set.

### GetValue

`func (o *UpdateEntityCriticalityResponseData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateEntityCriticalityResponseData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateEntityCriticalityResponseData) SetValue(v string)`

SetValue sets Value field to given value.


### GetHostname

`func (o *UpdateEntityCriticalityResponseData) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *UpdateEntityCriticalityResponseData) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *UpdateEntityCriticalityResponseData) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *UpdateEntityCriticalityResponseData) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetMacAddress

`func (o *UpdateEntityCriticalityResponseData) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *UpdateEntityCriticalityResponseData) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *UpdateEntityCriticalityResponseData) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *UpdateEntityCriticalityResponseData) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


