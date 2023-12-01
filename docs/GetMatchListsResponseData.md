# GetMatchListsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNextPage** | **bool** |  | 
**Total** | **int32** |  | 
**Objects** | [**[]MatchList**](MatchList.md) |  | 

## Methods

### NewGetMatchListsResponseData

`func NewGetMatchListsResponseData(hasNextPage bool, total int32, objects []MatchList, ) *GetMatchListsResponseData`

NewGetMatchListsResponseData instantiates a new GetMatchListsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMatchListsResponseDataWithDefaults

`func NewGetMatchListsResponseDataWithDefaults() *GetMatchListsResponseData`

NewGetMatchListsResponseDataWithDefaults instantiates a new GetMatchListsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNextPage

`func (o *GetMatchListsResponseData) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *GetMatchListsResponseData) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *GetMatchListsResponseData) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetTotal

`func (o *GetMatchListsResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetMatchListsResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetMatchListsResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetObjects

`func (o *GetMatchListsResponseData) GetObjects() []MatchList`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GetMatchListsResponseData) GetObjectsOk() (*[]MatchList, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GetMatchListsResponseData) SetObjects(v []MatchList)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


