# UpdateTagSchemaRequestBodyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Label** | **string** |  | 
**ContentTypes** | Pointer to **[]string** |  | [optional] 
**Freeform** | **bool** |  | 
**ValueOptions** | Pointer to [**[]UpdateTagSchemaRequestBodyFieldsValueOptionsInner**](UpdateTagSchemaRequestBodyFieldsValueOptionsInner.md) |  | [optional] 

## Methods

### NewUpdateTagSchemaRequestBodyFields

`func NewUpdateTagSchemaRequestBodyFields(key string, label string, freeform bool, ) *UpdateTagSchemaRequestBodyFields`

NewUpdateTagSchemaRequestBodyFields instantiates a new UpdateTagSchemaRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTagSchemaRequestBodyFieldsWithDefaults

`func NewUpdateTagSchemaRequestBodyFieldsWithDefaults() *UpdateTagSchemaRequestBodyFields`

NewUpdateTagSchemaRequestBodyFieldsWithDefaults instantiates a new UpdateTagSchemaRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *UpdateTagSchemaRequestBodyFields) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateTagSchemaRequestBodyFields) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateTagSchemaRequestBodyFields) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *UpdateTagSchemaRequestBodyFields) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateTagSchemaRequestBodyFields) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateTagSchemaRequestBodyFields) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetContentTypes

`func (o *UpdateTagSchemaRequestBodyFields) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *UpdateTagSchemaRequestBodyFields) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *UpdateTagSchemaRequestBodyFields) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *UpdateTagSchemaRequestBodyFields) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetFreeform

`func (o *UpdateTagSchemaRequestBodyFields) GetFreeform() bool`

GetFreeform returns the Freeform field if non-nil, zero value otherwise.

### GetFreeformOk

`func (o *UpdateTagSchemaRequestBodyFields) GetFreeformOk() (*bool, bool)`

GetFreeformOk returns a tuple with the Freeform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeform

`func (o *UpdateTagSchemaRequestBodyFields) SetFreeform(v bool)`

SetFreeform sets Freeform field to given value.


### GetValueOptions

`func (o *UpdateTagSchemaRequestBodyFields) GetValueOptions() []UpdateTagSchemaRequestBodyFieldsValueOptionsInner`

GetValueOptions returns the ValueOptions field if non-nil, zero value otherwise.

### GetValueOptionsOk

`func (o *UpdateTagSchemaRequestBodyFields) GetValueOptionsOk() (*[]UpdateTagSchemaRequestBodyFieldsValueOptionsInner, bool)`

GetValueOptionsOk returns a tuple with the ValueOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueOptions

`func (o *UpdateTagSchemaRequestBodyFields) SetValueOptions(v []UpdateTagSchemaRequestBodyFieldsValueOptionsInner)`

SetValueOptions sets ValueOptions field to given value.

### HasValueOptions

`func (o *UpdateTagSchemaRequestBodyFields) HasValueOptions() bool`

HasValueOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


