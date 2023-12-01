# InsightResolution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Source** | **string** |  | 
**Parent** | Pointer to [**InsightResolutionParent**](InsightResolutionParent.md) |  | [optional] 

## Methods

### NewInsightResolution

`func NewInsightResolution(id int32, name string, description string, source string, ) *InsightResolution`

NewInsightResolution instantiates a new InsightResolution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightResolutionWithDefaults

`func NewInsightResolutionWithDefaults() *InsightResolution`

NewInsightResolutionWithDefaults instantiates a new InsightResolution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InsightResolution) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InsightResolution) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InsightResolution) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *InsightResolution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InsightResolution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InsightResolution) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InsightResolution) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InsightResolution) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InsightResolution) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSource

`func (o *InsightResolution) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InsightResolution) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InsightResolution) SetSource(v string)`

SetSource sets Source field to given value.


### GetParent

`func (o *InsightResolution) GetParent() InsightResolutionParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *InsightResolution) GetParentOk() (*InsightResolutionParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *InsightResolution) SetParent(v InsightResolutionParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *InsightResolution) HasParent() bool`

HasParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


