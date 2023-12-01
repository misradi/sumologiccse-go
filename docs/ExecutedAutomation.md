# ExecutedAutomation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | **string** |  | 
**CseResourceType** | **string** |  | 
**CseResourceId** | **string** |  | 
**Status** | **string** |  | 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**PlaybookId** | **string** |  | 
**PlaybookName** | **string** |  | 
**PlaybookDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewExecutedAutomation

`func NewExecutedAutomation(executionId string, cseResourceType string, cseResourceId string, status string, playbookId string, playbookName string, ) *ExecutedAutomation`

NewExecutedAutomation instantiates a new ExecutedAutomation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutedAutomationWithDefaults

`func NewExecutedAutomationWithDefaults() *ExecutedAutomation`

NewExecutedAutomationWithDefaults instantiates a new ExecutedAutomation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *ExecutedAutomation) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ExecutedAutomation) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ExecutedAutomation) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetCseResourceType

`func (o *ExecutedAutomation) GetCseResourceType() string`

GetCseResourceType returns the CseResourceType field if non-nil, zero value otherwise.

### GetCseResourceTypeOk

`func (o *ExecutedAutomation) GetCseResourceTypeOk() (*string, bool)`

GetCseResourceTypeOk returns a tuple with the CseResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCseResourceType

`func (o *ExecutedAutomation) SetCseResourceType(v string)`

SetCseResourceType sets CseResourceType field to given value.


### GetCseResourceId

`func (o *ExecutedAutomation) GetCseResourceId() string`

GetCseResourceId returns the CseResourceId field if non-nil, zero value otherwise.

### GetCseResourceIdOk

`func (o *ExecutedAutomation) GetCseResourceIdOk() (*string, bool)`

GetCseResourceIdOk returns a tuple with the CseResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCseResourceId

`func (o *ExecutedAutomation) SetCseResourceId(v string)`

SetCseResourceId sets CseResourceId field to given value.


### GetStatus

`func (o *ExecutedAutomation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecutedAutomation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecutedAutomation) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartDate

`func (o *ExecutedAutomation) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ExecutedAutomation) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ExecutedAutomation) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ExecutedAutomation) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ExecutedAutomation) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ExecutedAutomation) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ExecutedAutomation) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ExecutedAutomation) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPlaybookId

`func (o *ExecutedAutomation) GetPlaybookId() string`

GetPlaybookId returns the PlaybookId field if non-nil, zero value otherwise.

### GetPlaybookIdOk

`func (o *ExecutedAutomation) GetPlaybookIdOk() (*string, bool)`

GetPlaybookIdOk returns a tuple with the PlaybookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookId

`func (o *ExecutedAutomation) SetPlaybookId(v string)`

SetPlaybookId sets PlaybookId field to given value.


### GetPlaybookName

`func (o *ExecutedAutomation) GetPlaybookName() string`

GetPlaybookName returns the PlaybookName field if non-nil, zero value otherwise.

### GetPlaybookNameOk

`func (o *ExecutedAutomation) GetPlaybookNameOk() (*string, bool)`

GetPlaybookNameOk returns a tuple with the PlaybookName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookName

`func (o *ExecutedAutomation) SetPlaybookName(v string)`

SetPlaybookName sets PlaybookName field to given value.


### GetPlaybookDescription

`func (o *ExecutedAutomation) GetPlaybookDescription() string`

GetPlaybookDescription returns the PlaybookDescription field if non-nil, zero value otherwise.

### GetPlaybookDescriptionOk

`func (o *ExecutedAutomation) GetPlaybookDescriptionOk() (*string, bool)`

GetPlaybookDescriptionOk returns a tuple with the PlaybookDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookDescription

`func (o *ExecutedAutomation) SetPlaybookDescription(v string)`

SetPlaybookDescription sets PlaybookDescription field to given value.

### HasPlaybookDescription

`func (o *ExecutedAutomation) HasPlaybookDescription() bool`

HasPlaybookDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


