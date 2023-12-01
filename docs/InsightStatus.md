# InsightStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**DisplayName** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 

## Methods

### NewInsightStatus

`func NewInsightStatus(id string, name string, displayName string, ) *InsightStatus`

NewInsightStatus instantiates a new InsightStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightStatusWithDefaults

`func NewInsightStatusWithDefaults() *InsightStatus`

NewInsightStatusWithDefaults instantiates a new InsightStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InsightStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InsightStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InsightStatus) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *InsightStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InsightStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InsightStatus) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *InsightStatus) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InsightStatus) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InsightStatus) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *InsightStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InsightStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InsightStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InsightStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetColor

`func (o *InsightStatus) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *InsightStatus) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *InsightStatus) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *InsightStatus) HasColor() bool`

HasColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


