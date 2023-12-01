# GetAllSignalsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** |  | [optional] 
**Objects** | [**[]Signal**](Signal.md) |  | 

## Methods

### NewGetAllSignalsResponseData

`func NewGetAllSignalsResponseData(objects []Signal, ) *GetAllSignalsResponseData`

NewGetAllSignalsResponseData instantiates a new GetAllSignalsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllSignalsResponseDataWithDefaults

`func NewGetAllSignalsResponseDataWithDefaults() *GetAllSignalsResponseData`

NewGetAllSignalsResponseDataWithDefaults instantiates a new GetAllSignalsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *GetAllSignalsResponseData) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *GetAllSignalsResponseData) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *GetAllSignalsResponseData) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *GetAllSignalsResponseData) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetObjects

`func (o *GetAllSignalsResponseData) GetObjects() []Signal`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GetAllSignalsResponseData) GetObjectsOk() (*[]Signal, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GetAllSignalsResponseData) SetObjects(v []Signal)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


