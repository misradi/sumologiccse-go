# GetInsightCommentsResponseDataCommentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Author** | Pointer to [**GetEntityNotesResponseDataInnerAuthor**](GetEntityNotesResponseDataInnerAuthor.md) |  | [optional] 
**Timestamp** | **time.Time** |  | 
**Body** | **string** |  | 

## Methods

### NewGetInsightCommentsResponseDataCommentsInner

`func NewGetInsightCommentsResponseDataCommentsInner(id string, timestamp time.Time, body string, ) *GetInsightCommentsResponseDataCommentsInner`

NewGetInsightCommentsResponseDataCommentsInner instantiates a new GetInsightCommentsResponseDataCommentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInsightCommentsResponseDataCommentsInnerWithDefaults

`func NewGetInsightCommentsResponseDataCommentsInnerWithDefaults() *GetInsightCommentsResponseDataCommentsInner`

NewGetInsightCommentsResponseDataCommentsInnerWithDefaults instantiates a new GetInsightCommentsResponseDataCommentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInsightCommentsResponseDataCommentsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInsightCommentsResponseDataCommentsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInsightCommentsResponseDataCommentsInner) SetId(v string)`

SetId sets Id field to given value.


### GetAuthor

`func (o *GetInsightCommentsResponseDataCommentsInner) GetAuthor() GetEntityNotesResponseDataInnerAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *GetInsightCommentsResponseDataCommentsInner) GetAuthorOk() (*GetEntityNotesResponseDataInnerAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *GetInsightCommentsResponseDataCommentsInner) SetAuthor(v GetEntityNotesResponseDataInnerAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *GetInsightCommentsResponseDataCommentsInner) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetTimestamp

`func (o *GetInsightCommentsResponseDataCommentsInner) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetInsightCommentsResponseDataCommentsInner) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetInsightCommentsResponseDataCommentsInner) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetBody

`func (o *GetInsightCommentsResponseDataCommentsInner) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GetInsightCommentsResponseDataCommentsInner) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GetInsightCommentsResponseDataCommentsInner) SetBody(v string)`

SetBody sets Body field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


