# GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueId** | **string** |  | 
**Source** | **string** |  | 
**Timestamp** | **string** |  | 
**ParsedTime** | **string** |  | 
**Groups** | **[]string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**ComputerName** | Pointer to **string** |  | [optional] 
**NormalizedComputerName** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**NormalizedHostname** | Pointer to **string** |  | [optional] 
**Ip** | **[]string** |  | 
**NatIp** | **[]string** |  | 
**Mac** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 

## Methods

### NewGetEntitiesResponseDataObjectsInnerInventoryInnerOneOf

`func NewGetEntitiesResponseDataObjectsInnerInventoryInnerOneOf(uniqueId string, source string, timestamp string, parsedTime string, groups []string, metadata map[string]interface{}, ip []string, natIp []string, ) *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf`

NewGetEntitiesResponseDataObjectsInnerInventoryInnerOneOf instantiates a new GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntitiesResponseDataObjectsInnerInventoryInnerOneOfWithDefaults

`func NewGetEntitiesResponseDataObjectsInnerInventoryInnerOneOfWithDefaults() *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf`

NewGetEntitiesResponseDataObjectsInnerInventoryInnerOneOfWithDefaults instantiates a new GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueId

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.


### GetSource

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) SetSource(v string)`

SetSource sets Source field to given value.


### GetTimestamp

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetParsedTime

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetParsedTime() string`

GetParsedTime returns the ParsedTime field if non-nil, zero value otherwise.

### GetParsedTimeOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetParsedTimeOk() (*string, bool)`

GetParsedTimeOk returns a tuple with the ParsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedTime

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) SetParsedTime(v string)`

SetParsedTime sets ParsedTime field to given value.


### GetGroups

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) SetGroups(v []string)`

SetGroups sets Groups field to given value.


### GetMetadata

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetComputerName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetComputerName() string`

GetComputerName returns the ComputerName field if non-nil, zero value otherwise.

### GetComputerNameOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetComputerNameOk() (*string, bool)`

GetComputerNameOk returns a tuple with the ComputerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) SetComputerName(v string)`

SetComputerName sets ComputerName field to given value.

### HasComputerName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) HasComputerName() bool`

HasComputerName returns a boolean if a field has been set.

### GetNormalizedComputerName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetNormalizedComputerName() string`

GetNormalizedComputerName returns the NormalizedComputerName field if non-nil, zero value otherwise.

### GetNormalizedComputerNameOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetNormalizedComputerNameOk() (*string, bool)`

GetNormalizedComputerNameOk returns a tuple with the NormalizedComputerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedComputerName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) SetNormalizedComputerName(v string)`

SetNormalizedComputerName sets NormalizedComputerName field to given value.

### HasNormalizedComputerName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) HasNormalizedComputerName() bool`

HasNormalizedComputerName returns a boolean if a field has been set.

### GetHostname

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetNormalizedHostname

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetNormalizedHostname() string`

GetNormalizedHostname returns the NormalizedHostname field if non-nil, zero value otherwise.

### GetNormalizedHostnameOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetNormalizedHostnameOk() (*string, bool)`

GetNormalizedHostnameOk returns a tuple with the NormalizedHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedHostname

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) SetNormalizedHostname(v string)`

SetNormalizedHostname sets NormalizedHostname field to given value.

### HasNormalizedHostname

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) HasNormalizedHostname() bool`

HasNormalizedHostname returns a boolean if a field has been set.

### GetIp

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetIp() []string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetIpOk() (*[]string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) SetIp(v []string)`

SetIp sets Ip field to given value.


### GetNatIp

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetNatIp() []string`

GetNatIp returns the NatIp field if non-nil, zero value otherwise.

### GetNatIpOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetNatIpOk() (*[]string, bool)`

GetNatIpOk returns a tuple with the NatIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatIp

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) SetNatIp(v []string)`

SetNatIp sets NatIp field to given value.


### GetMac

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetOs

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOsVersion

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetLocation

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


