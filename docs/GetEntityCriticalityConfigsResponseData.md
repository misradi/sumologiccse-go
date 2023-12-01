# GetEntityCriticalityConfigsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNextPage** | **bool** |  | 
**Total** | **int32** |  | 
**Objects** | [**[]EntityCriticalityConfig**](EntityCriticalityConfig.md) |  | 

## Methods

### NewGetEntityCriticalityConfigsResponseData

`func NewGetEntityCriticalityConfigsResponseData(hasNextPage bool, total int32, objects []EntityCriticalityConfig, ) *GetEntityCriticalityConfigsResponseData`

NewGetEntityCriticalityConfigsResponseData instantiates a new GetEntityCriticalityConfigsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntityCriticalityConfigsResponseDataWithDefaults

`func NewGetEntityCriticalityConfigsResponseDataWithDefaults() *GetEntityCriticalityConfigsResponseData`

NewGetEntityCriticalityConfigsResponseDataWithDefaults instantiates a new GetEntityCriticalityConfigsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNextPage

`func (o *GetEntityCriticalityConfigsResponseData) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *GetEntityCriticalityConfigsResponseData) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *GetEntityCriticalityConfigsResponseData) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetTotal

`func (o *GetEntityCriticalityConfigsResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetEntityCriticalityConfigsResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetEntityCriticalityConfigsResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetObjects

`func (o *GetEntityCriticalityConfigsResponseData) GetObjects() []EntityCriticalityConfig`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GetEntityCriticalityConfigsResponseData) GetObjectsOk() (*[]EntityCriticalityConfig, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GetEntityCriticalityConfigsResponseData) SetObjects(v []EntityCriticalityConfig)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


