# GetEntityGroupsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNextPage** | **bool** |  | 
**Total** | **int32** |  | 
**Objects** | [**[]GetEntityGroupsResponseDataObjectsInner**](GetEntityGroupsResponseDataObjectsInner.md) |  | 

## Methods

### NewGetEntityGroupsResponseData

`func NewGetEntityGroupsResponseData(hasNextPage bool, total int32, objects []GetEntityGroupsResponseDataObjectsInner, ) *GetEntityGroupsResponseData`

NewGetEntityGroupsResponseData instantiates a new GetEntityGroupsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntityGroupsResponseDataWithDefaults

`func NewGetEntityGroupsResponseDataWithDefaults() *GetEntityGroupsResponseData`

NewGetEntityGroupsResponseDataWithDefaults instantiates a new GetEntityGroupsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNextPage

`func (o *GetEntityGroupsResponseData) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *GetEntityGroupsResponseData) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *GetEntityGroupsResponseData) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetTotal

`func (o *GetEntityGroupsResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetEntityGroupsResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetEntityGroupsResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetObjects

`func (o *GetEntityGroupsResponseData) GetObjects() []GetEntityGroupsResponseDataObjectsInner`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GetEntityGroupsResponseData) GetObjectsOk() (*[]GetEntityGroupsResponseDataObjectsInner, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GetEntityGroupsResponseData) SetObjects(v []GetEntityGroupsResponseDataObjectsInner)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


