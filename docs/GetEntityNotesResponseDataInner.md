# GetEntityNotesResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**LastUpdated** | **time.Time** |  | 
**Note** | **string** |  | 
**Author** | Pointer to [**GetEntityNotesResponseDataInnerAuthor**](GetEntityNotesResponseDataInnerAuthor.md) |  | [optional] 

## Methods

### NewGetEntityNotesResponseDataInner

`func NewGetEntityNotesResponseDataInner(id string, timestamp time.Time, lastUpdated time.Time, note string, ) *GetEntityNotesResponseDataInner`

NewGetEntityNotesResponseDataInner instantiates a new GetEntityNotesResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntityNotesResponseDataInnerWithDefaults

`func NewGetEntityNotesResponseDataInnerWithDefaults() *GetEntityNotesResponseDataInner`

NewGetEntityNotesResponseDataInnerWithDefaults instantiates a new GetEntityNotesResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetEntityNotesResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEntityNotesResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEntityNotesResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetTimestamp

`func (o *GetEntityNotesResponseDataInner) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetEntityNotesResponseDataInner) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetEntityNotesResponseDataInner) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetLastUpdated

`func (o *GetEntityNotesResponseDataInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetEntityNotesResponseDataInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetEntityNotesResponseDataInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetNote

`func (o *GetEntityNotesResponseDataInner) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *GetEntityNotesResponseDataInner) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *GetEntityNotesResponseDataInner) SetNote(v string)`

SetNote sets Note field to given value.


### GetAuthor

`func (o *GetEntityNotesResponseDataInner) GetAuthor() GetEntityNotesResponseDataInnerAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *GetEntityNotesResponseDataInner) GetAuthorOk() (*GetEntityNotesResponseDataInnerAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *GetEntityNotesResponseDataInner) SetAuthor(v GetEntityNotesResponseDataInnerAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *GetEntityNotesResponseDataInner) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


