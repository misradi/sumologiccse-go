# GetEntitiesResponseDataObjectsInnerInventoryInner

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
**Username** | **string** |  | 
**NormalizedUsername** | Pointer to **string** |  | [optional] 
**GivenName** | Pointer to **string** |  | [optional] 
**MiddleName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Department** | Pointer to **string** |  | [optional] 
**Emails** | **[]string** |  | 

## Methods

### NewGetEntitiesResponseDataObjectsInnerInventoryInner

`func NewGetEntitiesResponseDataObjectsInnerInventoryInner(uniqueId string, source string, timestamp string, parsedTime string, groups []string, metadata map[string]interface{}, ip []string, natIp []string, username string, emails []string, ) *GetEntitiesResponseDataObjectsInnerInventoryInner`

NewGetEntitiesResponseDataObjectsInnerInventoryInner instantiates a new GetEntitiesResponseDataObjectsInnerInventoryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntitiesResponseDataObjectsInnerInventoryInnerWithDefaults

`func NewGetEntitiesResponseDataObjectsInnerInventoryInnerWithDefaults() *GetEntitiesResponseDataObjectsInnerInventoryInner`

NewGetEntitiesResponseDataObjectsInnerInventoryInnerWithDefaults instantiates a new GetEntitiesResponseDataObjectsInnerInventoryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueId

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.


### GetSource

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetSource(v string)`

SetSource sets Source field to given value.


### GetTimestamp

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetParsedTime

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetParsedTime() string`

GetParsedTime returns the ParsedTime field if non-nil, zero value otherwise.

### GetParsedTimeOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetParsedTimeOk() (*string, bool)`

GetParsedTimeOk returns a tuple with the ParsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedTime

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetParsedTime(v string)`

SetParsedTime sets ParsedTime field to given value.


### GetGroups

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetGroups(v []string)`

SetGroups sets Groups field to given value.


### GetMetadata

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetComputerName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetComputerName() string`

GetComputerName returns the ComputerName field if non-nil, zero value otherwise.

### GetComputerNameOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetComputerNameOk() (*string, bool)`

GetComputerNameOk returns a tuple with the ComputerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetComputerName(v string)`

SetComputerName sets ComputerName field to given value.

### HasComputerName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) HasComputerName() bool`

HasComputerName returns a boolean if a field has been set.

### GetNormalizedComputerName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetNormalizedComputerName() string`

GetNormalizedComputerName returns the NormalizedComputerName field if non-nil, zero value otherwise.

### GetNormalizedComputerNameOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetNormalizedComputerNameOk() (*string, bool)`

GetNormalizedComputerNameOk returns a tuple with the NormalizedComputerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedComputerName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetNormalizedComputerName(v string)`

SetNormalizedComputerName sets NormalizedComputerName field to given value.

### HasNormalizedComputerName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) HasNormalizedComputerName() bool`

HasNormalizedComputerName returns a boolean if a field has been set.

### GetHostname

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetNormalizedHostname

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetNormalizedHostname() string`

GetNormalizedHostname returns the NormalizedHostname field if non-nil, zero value otherwise.

### GetNormalizedHostnameOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetNormalizedHostnameOk() (*string, bool)`

GetNormalizedHostnameOk returns a tuple with the NormalizedHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedHostname

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetNormalizedHostname(v string)`

SetNormalizedHostname sets NormalizedHostname field to given value.

### HasNormalizedHostname

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) HasNormalizedHostname() bool`

HasNormalizedHostname returns a boolean if a field has been set.

### GetIp

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetIp() []string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetIpOk() (*[]string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetIp(v []string)`

SetIp sets Ip field to given value.


### GetNatIp

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetNatIp() []string`

GetNatIp returns the NatIp field if non-nil, zero value otherwise.

### GetNatIpOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetNatIpOk() (*[]string, bool)`

GetNatIpOk returns a tuple with the NatIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatIp

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetNatIp(v []string)`

SetNatIp sets NatIp field to given value.


### GetMac

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetOs

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOsVersion

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetLocation

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetUsername

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetNormalizedUsername

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetNormalizedUsername() string`

GetNormalizedUsername returns the NormalizedUsername field if non-nil, zero value otherwise.

### GetNormalizedUsernameOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetNormalizedUsernameOk() (*string, bool)`

GetNormalizedUsernameOk returns a tuple with the NormalizedUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedUsername

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetNormalizedUsername(v string)`

SetNormalizedUsername sets NormalizedUsername field to given value.

### HasNormalizedUsername

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) HasNormalizedUsername() bool`

HasNormalizedUsername returns a boolean if a field has been set.

### GetGivenName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetMiddleName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetLastName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetDepartment

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetEmails

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *GetEntitiesResponseDataObjectsInnerInventoryInner) SetEmails(v []string)`

SetEmails sets Emails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


