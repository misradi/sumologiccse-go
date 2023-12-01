# UpdateSuppressListRequestBodyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**DefaultTtl** | Pointer to **int32** | The default time-to-live (in seconds) for new Items added to this List. This default is only used to default the field in the UI, and is not used at all when new Items are added via the API. | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateSuppressListRequestBodyFields

`func NewUpdateSuppressListRequestBodyFields(description string, ) *UpdateSuppressListRequestBodyFields`

NewUpdateSuppressListRequestBodyFields instantiates a new UpdateSuppressListRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSuppressListRequestBodyFieldsWithDefaults

`func NewUpdateSuppressListRequestBodyFieldsWithDefaults() *UpdateSuppressListRequestBodyFields`

NewUpdateSuppressListRequestBodyFieldsWithDefaults instantiates a new UpdateSuppressListRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateSuppressListRequestBodyFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSuppressListRequestBodyFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSuppressListRequestBodyFields) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDefaultTtl

`func (o *UpdateSuppressListRequestBodyFields) GetDefaultTtl() int32`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *UpdateSuppressListRequestBodyFields) GetDefaultTtlOk() (*int32, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *UpdateSuppressListRequestBodyFields) SetDefaultTtl(v int32)`

SetDefaultTtl sets DefaultTtl field to given value.

### HasDefaultTtl

`func (o *UpdateSuppressListRequestBodyFields) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.

### GetActive

`func (o *UpdateSuppressListRequestBodyFields) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateSuppressListRequestBodyFields) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateSuppressListRequestBodyFields) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateSuppressListRequestBodyFields) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


