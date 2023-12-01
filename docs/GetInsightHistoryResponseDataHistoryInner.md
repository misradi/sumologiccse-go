# GetInsightHistoryResponseDataHistoryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**GetEntityNotesResponseDataInnerAuthor**](GetEntityNotesResponseDataInnerAuthor.md) |  | [optional] 
**Description** | **string** |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewGetInsightHistoryResponseDataHistoryInner

`func NewGetInsightHistoryResponseDataHistoryInner(description string, timestamp time.Time, ) *GetInsightHistoryResponseDataHistoryInner`

NewGetInsightHistoryResponseDataHistoryInner instantiates a new GetInsightHistoryResponseDataHistoryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInsightHistoryResponseDataHistoryInnerWithDefaults

`func NewGetInsightHistoryResponseDataHistoryInnerWithDefaults() *GetInsightHistoryResponseDataHistoryInner`

NewGetInsightHistoryResponseDataHistoryInnerWithDefaults instantiates a new GetInsightHistoryResponseDataHistoryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *GetInsightHistoryResponseDataHistoryInner) GetUser() GetEntityNotesResponseDataInnerAuthor`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetInsightHistoryResponseDataHistoryInner) GetUserOk() (*GetEntityNotesResponseDataInnerAuthor, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetInsightHistoryResponseDataHistoryInner) SetUser(v GetEntityNotesResponseDataInnerAuthor)`

SetUser sets User field to given value.

### HasUser

`func (o *GetInsightHistoryResponseDataHistoryInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetDescription

`func (o *GetInsightHistoryResponseDataHistoryInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetInsightHistoryResponseDataHistoryInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetInsightHistoryResponseDataHistoryInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTimestamp

`func (o *GetInsightHistoryResponseDataHistoryInner) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetInsightHistoryResponseDataHistoryInner) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetInsightHistoryResponseDataHistoryInner) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


