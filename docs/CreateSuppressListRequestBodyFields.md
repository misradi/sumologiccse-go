# CreateSuppressListRequestBodyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**TargetColumn** | **string** |  | 
**DefaultTtl** | Pointer to **int32** | The default time-to-live (in seconds) for new Items added to this List. This default is only used to default the field in the UI, and is not used at all when new Items are added via the API. | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateSuppressListRequestBodyFields

`func NewCreateSuppressListRequestBodyFields(name string, targetColumn string, ) *CreateSuppressListRequestBodyFields`

NewCreateSuppressListRequestBodyFields instantiates a new CreateSuppressListRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSuppressListRequestBodyFieldsWithDefaults

`func NewCreateSuppressListRequestBodyFieldsWithDefaults() *CreateSuppressListRequestBodyFields`

NewCreateSuppressListRequestBodyFieldsWithDefaults instantiates a new CreateSuppressListRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSuppressListRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSuppressListRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSuppressListRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateSuppressListRequestBodyFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSuppressListRequestBodyFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSuppressListRequestBodyFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSuppressListRequestBodyFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTargetColumn

`func (o *CreateSuppressListRequestBodyFields) GetTargetColumn() string`

GetTargetColumn returns the TargetColumn field if non-nil, zero value otherwise.

### GetTargetColumnOk

`func (o *CreateSuppressListRequestBodyFields) GetTargetColumnOk() (*string, bool)`

GetTargetColumnOk returns a tuple with the TargetColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetColumn

`func (o *CreateSuppressListRequestBodyFields) SetTargetColumn(v string)`

SetTargetColumn sets TargetColumn field to given value.


### GetDefaultTtl

`func (o *CreateSuppressListRequestBodyFields) GetDefaultTtl() int32`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *CreateSuppressListRequestBodyFields) GetDefaultTtlOk() (*int32, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *CreateSuppressListRequestBodyFields) SetDefaultTtl(v int32)`

SetDefaultTtl sets DefaultTtl field to given value.

### HasDefaultTtl

`func (o *CreateSuppressListRequestBodyFields) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.

### GetActive

`func (o *CreateSuppressListRequestBodyFields) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateSuppressListRequestBodyFields) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateSuppressListRequestBodyFields) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateSuppressListRequestBodyFields) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


