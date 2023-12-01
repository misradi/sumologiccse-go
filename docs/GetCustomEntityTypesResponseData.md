# GetCustomEntityTypesResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNextPage** | **bool** |  | 
**Total** | **int32** |  | 
**Objects** | [**[]CustomEntityType**](CustomEntityType.md) |  | 

## Methods

### NewGetCustomEntityTypesResponseData

`func NewGetCustomEntityTypesResponseData(hasNextPage bool, total int32, objects []CustomEntityType, ) *GetCustomEntityTypesResponseData`

NewGetCustomEntityTypesResponseData instantiates a new GetCustomEntityTypesResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomEntityTypesResponseDataWithDefaults

`func NewGetCustomEntityTypesResponseDataWithDefaults() *GetCustomEntityTypesResponseData`

NewGetCustomEntityTypesResponseDataWithDefaults instantiates a new GetCustomEntityTypesResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNextPage

`func (o *GetCustomEntityTypesResponseData) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *GetCustomEntityTypesResponseData) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *GetCustomEntityTypesResponseData) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetTotal

`func (o *GetCustomEntityTypesResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetCustomEntityTypesResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetCustomEntityTypesResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetObjects

`func (o *GetCustomEntityTypesResponseData) GetObjects() []CustomEntityType`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GetCustomEntityTypesResponseData) GetObjectsOk() (*[]CustomEntityType, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GetCustomEntityTypesResponseData) SetObjects(v []CustomEntityType)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


