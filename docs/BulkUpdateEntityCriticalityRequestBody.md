# BulkUpdateEntityCriticalityRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criticality** | Pointer to **string** |  | [optional] 
**EntityIds** | **[]string** |  | 

## Methods

### NewBulkUpdateEntityCriticalityRequestBody

`func NewBulkUpdateEntityCriticalityRequestBody(entityIds []string, ) *BulkUpdateEntityCriticalityRequestBody`

NewBulkUpdateEntityCriticalityRequestBody instantiates a new BulkUpdateEntityCriticalityRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateEntityCriticalityRequestBodyWithDefaults

`func NewBulkUpdateEntityCriticalityRequestBodyWithDefaults() *BulkUpdateEntityCriticalityRequestBody`

NewBulkUpdateEntityCriticalityRequestBodyWithDefaults instantiates a new BulkUpdateEntityCriticalityRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriticality

`func (o *BulkUpdateEntityCriticalityRequestBody) GetCriticality() string`

GetCriticality returns the Criticality field if non-nil, zero value otherwise.

### GetCriticalityOk

`func (o *BulkUpdateEntityCriticalityRequestBody) GetCriticalityOk() (*string, bool)`

GetCriticalityOk returns a tuple with the Criticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticality

`func (o *BulkUpdateEntityCriticalityRequestBody) SetCriticality(v string)`

SetCriticality sets Criticality field to given value.

### HasCriticality

`func (o *BulkUpdateEntityCriticalityRequestBody) HasCriticality() bool`

HasCriticality returns a boolean if a field has been set.

### GetEntityIds

`func (o *BulkUpdateEntityCriticalityRequestBody) GetEntityIds() []string`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *BulkUpdateEntityCriticalityRequestBody) GetEntityIdsOk() (*[]string, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *BulkUpdateEntityCriticalityRequestBody) SetEntityIds(v []string)`

SetEntityIds sets EntityIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


