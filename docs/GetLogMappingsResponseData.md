# GetLogMappingsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNextPage** | **bool** |  | 
**Total** | **int32** |  | 
**Objects** | [**[]LogMapping**](LogMapping.md) |  | 

## Methods

### NewGetLogMappingsResponseData

`func NewGetLogMappingsResponseData(hasNextPage bool, total int32, objects []LogMapping, ) *GetLogMappingsResponseData`

NewGetLogMappingsResponseData instantiates a new GetLogMappingsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLogMappingsResponseDataWithDefaults

`func NewGetLogMappingsResponseDataWithDefaults() *GetLogMappingsResponseData`

NewGetLogMappingsResponseDataWithDefaults instantiates a new GetLogMappingsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNextPage

`func (o *GetLogMappingsResponseData) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *GetLogMappingsResponseData) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *GetLogMappingsResponseData) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetTotal

`func (o *GetLogMappingsResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetLogMappingsResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetLogMappingsResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetObjects

`func (o *GetLogMappingsResponseData) GetObjects() []LogMapping`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GetLogMappingsResponseData) GetObjectsOk() (*[]LogMapping, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GetLogMappingsResponseData) SetObjects(v []LogMapping)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


