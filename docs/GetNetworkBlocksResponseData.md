# GetNetworkBlocksResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNextPage** | **bool** |  | 
**Total** | **int32** |  | 
**Objects** | [**[]NetworkBlock**](NetworkBlock.md) |  | 

## Methods

### NewGetNetworkBlocksResponseData

`func NewGetNetworkBlocksResponseData(hasNextPage bool, total int32, objects []NetworkBlock, ) *GetNetworkBlocksResponseData`

NewGetNetworkBlocksResponseData instantiates a new GetNetworkBlocksResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkBlocksResponseDataWithDefaults

`func NewGetNetworkBlocksResponseDataWithDefaults() *GetNetworkBlocksResponseData`

NewGetNetworkBlocksResponseDataWithDefaults instantiates a new GetNetworkBlocksResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNextPage

`func (o *GetNetworkBlocksResponseData) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *GetNetworkBlocksResponseData) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *GetNetworkBlocksResponseData) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetTotal

`func (o *GetNetworkBlocksResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetNetworkBlocksResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetNetworkBlocksResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetObjects

`func (o *GetNetworkBlocksResponseData) GetObjects() []NetworkBlock`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GetNetworkBlocksResponseData) GetObjectsOk() (*[]NetworkBlock, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GetNetworkBlocksResponseData) SetObjects(v []NetworkBlock)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


