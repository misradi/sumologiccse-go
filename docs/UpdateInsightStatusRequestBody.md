# UpdateInsightStatusRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status to update this Insight to. Default values are \&quot;new\&quot;, \&quot;inprogress\&quot;, and \&quot;closed\&quot;, but custom statuses can also be created in the UI. | [optional] 
**Resolution** | Pointer to **string** | The resolution reason for closing this Insight. | [optional] 

## Methods

### NewUpdateInsightStatusRequestBody

`func NewUpdateInsightStatusRequestBody() *UpdateInsightStatusRequestBody`

NewUpdateInsightStatusRequestBody instantiates a new UpdateInsightStatusRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInsightStatusRequestBodyWithDefaults

`func NewUpdateInsightStatusRequestBodyWithDefaults() *UpdateInsightStatusRequestBody`

NewUpdateInsightStatusRequestBodyWithDefaults instantiates a new UpdateInsightStatusRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateInsightStatusRequestBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateInsightStatusRequestBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateInsightStatusRequestBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateInsightStatusRequestBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResolution

`func (o *UpdateInsightStatusRequestBody) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *UpdateInsightStatusRequestBody) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *UpdateInsightStatusRequestBody) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *UpdateInsightStatusRequestBody) HasResolution() bool`

HasResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


