# UpdateInsightAssigneeRequestBodyAssignee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the assignee, either USER or TEAM. | 
**Value** | **string** | The username or team name of the user/team to be assigned. | 

## Methods

### NewUpdateInsightAssigneeRequestBodyAssignee

`func NewUpdateInsightAssigneeRequestBodyAssignee(type_ string, value string, ) *UpdateInsightAssigneeRequestBodyAssignee`

NewUpdateInsightAssigneeRequestBodyAssignee instantiates a new UpdateInsightAssigneeRequestBodyAssignee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInsightAssigneeRequestBodyAssigneeWithDefaults

`func NewUpdateInsightAssigneeRequestBodyAssigneeWithDefaults() *UpdateInsightAssigneeRequestBodyAssignee`

NewUpdateInsightAssigneeRequestBodyAssigneeWithDefaults instantiates a new UpdateInsightAssigneeRequestBodyAssignee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateInsightAssigneeRequestBodyAssignee) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateInsightAssigneeRequestBodyAssignee) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateInsightAssigneeRequestBodyAssignee) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *UpdateInsightAssigneeRequestBodyAssignee) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateInsightAssigneeRequestBodyAssignee) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateInsightAssigneeRequestBodyAssignee) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


