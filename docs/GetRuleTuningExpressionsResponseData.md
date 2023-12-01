# GetRuleTuningExpressionsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Objects** | [**[]RuleTuningExpression**](RuleTuningExpression.md) |  | 

## Methods

### NewGetRuleTuningExpressionsResponseData

`func NewGetRuleTuningExpressionsResponseData(total int32, objects []RuleTuningExpression, ) *GetRuleTuningExpressionsResponseData`

NewGetRuleTuningExpressionsResponseData instantiates a new GetRuleTuningExpressionsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRuleTuningExpressionsResponseDataWithDefaults

`func NewGetRuleTuningExpressionsResponseDataWithDefaults() *GetRuleTuningExpressionsResponseData`

NewGetRuleTuningExpressionsResponseDataWithDefaults instantiates a new GetRuleTuningExpressionsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetRuleTuningExpressionsResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetRuleTuningExpressionsResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetRuleTuningExpressionsResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetObjects

`func (o *GetRuleTuningExpressionsResponseData) GetObjects() []RuleTuningExpression`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GetRuleTuningExpressionsResponseData) GetObjectsOk() (*[]RuleTuningExpression, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GetRuleTuningExpressionsResponseData) SetObjects(v []RuleTuningExpression)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


