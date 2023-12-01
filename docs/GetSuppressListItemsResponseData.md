# GetSuppressListItemsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNextPage** | **bool** |  | 
**Total** | **int32** |  | 
**Objects** | [**[]SuppressListItem**](SuppressListItem.md) |  | 

## Methods

### NewGetSuppressListItemsResponseData

`func NewGetSuppressListItemsResponseData(hasNextPage bool, total int32, objects []SuppressListItem, ) *GetSuppressListItemsResponseData`

NewGetSuppressListItemsResponseData instantiates a new GetSuppressListItemsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSuppressListItemsResponseDataWithDefaults

`func NewGetSuppressListItemsResponseDataWithDefaults() *GetSuppressListItemsResponseData`

NewGetSuppressListItemsResponseDataWithDefaults instantiates a new GetSuppressListItemsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNextPage

`func (o *GetSuppressListItemsResponseData) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *GetSuppressListItemsResponseData) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *GetSuppressListItemsResponseData) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetTotal

`func (o *GetSuppressListItemsResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetSuppressListItemsResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetSuppressListItemsResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetObjects

`func (o *GetSuppressListItemsResponseData) GetObjects() []SuppressListItem`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GetSuppressListItemsResponseData) GetObjectsOk() (*[]SuppressListItem, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GetSuppressListItemsResponseData) SetObjects(v []SuppressListItem)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


