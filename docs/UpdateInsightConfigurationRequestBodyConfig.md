# UpdateInsightConfigurationRequestBodyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LookbackDays** | Pointer to **float64** |  | [optional] 
**Threshold** | Pointer to **float64** |  | [optional] 
**GlobalSignalSuppressionWindow** | Pointer to **int32** | Number of hours between 24 and 72. | [optional] 

## Methods

### NewUpdateInsightConfigurationRequestBodyConfig

`func NewUpdateInsightConfigurationRequestBodyConfig() *UpdateInsightConfigurationRequestBodyConfig`

NewUpdateInsightConfigurationRequestBodyConfig instantiates a new UpdateInsightConfigurationRequestBodyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInsightConfigurationRequestBodyConfigWithDefaults

`func NewUpdateInsightConfigurationRequestBodyConfigWithDefaults() *UpdateInsightConfigurationRequestBodyConfig`

NewUpdateInsightConfigurationRequestBodyConfigWithDefaults instantiates a new UpdateInsightConfigurationRequestBodyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLookbackDays

`func (o *UpdateInsightConfigurationRequestBodyConfig) GetLookbackDays() float64`

GetLookbackDays returns the LookbackDays field if non-nil, zero value otherwise.

### GetLookbackDaysOk

`func (o *UpdateInsightConfigurationRequestBodyConfig) GetLookbackDaysOk() (*float64, bool)`

GetLookbackDaysOk returns a tuple with the LookbackDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookbackDays

`func (o *UpdateInsightConfigurationRequestBodyConfig) SetLookbackDays(v float64)`

SetLookbackDays sets LookbackDays field to given value.

### HasLookbackDays

`func (o *UpdateInsightConfigurationRequestBodyConfig) HasLookbackDays() bool`

HasLookbackDays returns a boolean if a field has been set.

### GetThreshold

`func (o *UpdateInsightConfigurationRequestBodyConfig) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *UpdateInsightConfigurationRequestBodyConfig) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *UpdateInsightConfigurationRequestBodyConfig) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *UpdateInsightConfigurationRequestBodyConfig) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetGlobalSignalSuppressionWindow

`func (o *UpdateInsightConfigurationRequestBodyConfig) GetGlobalSignalSuppressionWindow() int32`

GetGlobalSignalSuppressionWindow returns the GlobalSignalSuppressionWindow field if non-nil, zero value otherwise.

### GetGlobalSignalSuppressionWindowOk

`func (o *UpdateInsightConfigurationRequestBodyConfig) GetGlobalSignalSuppressionWindowOk() (*int32, bool)`

GetGlobalSignalSuppressionWindowOk returns a tuple with the GlobalSignalSuppressionWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSignalSuppressionWindow

`func (o *UpdateInsightConfigurationRequestBodyConfig) SetGlobalSignalSuppressionWindow(v int32)`

SetGlobalSignalSuppressionWindow sets GlobalSignalSuppressionWindow field to given value.

### HasGlobalSignalSuppressionWindow

`func (o *UpdateInsightConfigurationRequestBodyConfig) HasGlobalSignalSuppressionWindow() bool`

HasGlobalSignalSuppressionWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


