# GetAllEntitiesResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** |  | [optional] 
**Objects** | [**[]GetEntitiesResponseDataObjectsInner**](GetEntitiesResponseDataObjectsInner.md) |  | 

## Methods

### NewGetAllEntitiesResponseData

`func NewGetAllEntitiesResponseData(objects []GetEntitiesResponseDataObjectsInner, ) *GetAllEntitiesResponseData`

NewGetAllEntitiesResponseData instantiates a new GetAllEntitiesResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllEntitiesResponseDataWithDefaults

`func NewGetAllEntitiesResponseDataWithDefaults() *GetAllEntitiesResponseData`

NewGetAllEntitiesResponseDataWithDefaults instantiates a new GetAllEntitiesResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *GetAllEntitiesResponseData) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *GetAllEntitiesResponseData) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *GetAllEntitiesResponseData) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *GetAllEntitiesResponseData) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetObjects

`func (o *GetAllEntitiesResponseData) GetObjects() []GetEntitiesResponseDataObjectsInner`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GetAllEntitiesResponseData) GetObjectsOk() (*[]GetEntitiesResponseDataObjectsInner, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GetAllEntitiesResponseData) SetObjects(v []GetEntitiesResponseDataObjectsInner)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


