# GetContextActionsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**HasNextPage** | **bool** |  | 
**Objects** | [**[]ContextAction**](ContextAction.md) |  | 

## Methods

### NewGetContextActionsResponseData

`func NewGetContextActionsResponseData(total int32, hasNextPage bool, objects []ContextAction, ) *GetContextActionsResponseData`

NewGetContextActionsResponseData instantiates a new GetContextActionsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContextActionsResponseDataWithDefaults

`func NewGetContextActionsResponseDataWithDefaults() *GetContextActionsResponseData`

NewGetContextActionsResponseDataWithDefaults instantiates a new GetContextActionsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetContextActionsResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetContextActionsResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetContextActionsResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetHasNextPage

`func (o *GetContextActionsResponseData) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *GetContextActionsResponseData) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *GetContextActionsResponseData) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetObjects

`func (o *GetContextActionsResponseData) GetObjects() []ContextAction`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GetContextActionsResponseData) GetObjectsOk() (*[]ContextAction, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GetContextActionsResponseData) SetObjects(v []ContextAction)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


