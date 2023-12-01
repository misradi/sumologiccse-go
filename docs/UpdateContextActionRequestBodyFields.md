# UpdateContextActionRequestBodyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**UrlTemplate** | Pointer to **string** |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**IocTypes** | **[]string** |  | 
**EntityTypes** | Pointer to **[]string** |  | [optional] 
**RecordFields** | Pointer to **[]string** |  | [optional] 
**AllRecordFields** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateContextActionRequestBodyFields

`func NewUpdateContextActionRequestBodyFields(name string, iocTypes []string, ) *UpdateContextActionRequestBodyFields`

NewUpdateContextActionRequestBodyFields instantiates a new UpdateContextActionRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContextActionRequestBodyFieldsWithDefaults

`func NewUpdateContextActionRequestBodyFieldsWithDefaults() *UpdateContextActionRequestBodyFields`

NewUpdateContextActionRequestBodyFieldsWithDefaults instantiates a new UpdateContextActionRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateContextActionRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateContextActionRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateContextActionRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *UpdateContextActionRequestBodyFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateContextActionRequestBodyFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateContextActionRequestBodyFields) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateContextActionRequestBodyFields) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrlTemplate

`func (o *UpdateContextActionRequestBodyFields) GetUrlTemplate() string`

GetUrlTemplate returns the UrlTemplate field if non-nil, zero value otherwise.

### GetUrlTemplateOk

`func (o *UpdateContextActionRequestBodyFields) GetUrlTemplateOk() (*string, bool)`

GetUrlTemplateOk returns a tuple with the UrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlTemplate

`func (o *UpdateContextActionRequestBodyFields) SetUrlTemplate(v string)`

SetUrlTemplate sets UrlTemplate field to given value.

### HasUrlTemplate

`func (o *UpdateContextActionRequestBodyFields) HasUrlTemplate() bool`

HasUrlTemplate returns a boolean if a field has been set.

### GetTemplate

`func (o *UpdateContextActionRequestBodyFields) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *UpdateContextActionRequestBodyFields) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *UpdateContextActionRequestBodyFields) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *UpdateContextActionRequestBodyFields) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetIocTypes

`func (o *UpdateContextActionRequestBodyFields) GetIocTypes() []string`

GetIocTypes returns the IocTypes field if non-nil, zero value otherwise.

### GetIocTypesOk

`func (o *UpdateContextActionRequestBodyFields) GetIocTypesOk() (*[]string, bool)`

GetIocTypesOk returns a tuple with the IocTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIocTypes

`func (o *UpdateContextActionRequestBodyFields) SetIocTypes(v []string)`

SetIocTypes sets IocTypes field to given value.


### GetEntityTypes

`func (o *UpdateContextActionRequestBodyFields) GetEntityTypes() []string`

GetEntityTypes returns the EntityTypes field if non-nil, zero value otherwise.

### GetEntityTypesOk

`func (o *UpdateContextActionRequestBodyFields) GetEntityTypesOk() (*[]string, bool)`

GetEntityTypesOk returns a tuple with the EntityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypes

`func (o *UpdateContextActionRequestBodyFields) SetEntityTypes(v []string)`

SetEntityTypes sets EntityTypes field to given value.

### HasEntityTypes

`func (o *UpdateContextActionRequestBodyFields) HasEntityTypes() bool`

HasEntityTypes returns a boolean if a field has been set.

### GetRecordFields

`func (o *UpdateContextActionRequestBodyFields) GetRecordFields() []string`

GetRecordFields returns the RecordFields field if non-nil, zero value otherwise.

### GetRecordFieldsOk

`func (o *UpdateContextActionRequestBodyFields) GetRecordFieldsOk() (*[]string, bool)`

GetRecordFieldsOk returns a tuple with the RecordFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordFields

`func (o *UpdateContextActionRequestBodyFields) SetRecordFields(v []string)`

SetRecordFields sets RecordFields field to given value.

### HasRecordFields

`func (o *UpdateContextActionRequestBodyFields) HasRecordFields() bool`

HasRecordFields returns a boolean if a field has been set.

### GetAllRecordFields

`func (o *UpdateContextActionRequestBodyFields) GetAllRecordFields() bool`

GetAllRecordFields returns the AllRecordFields field if non-nil, zero value otherwise.

### GetAllRecordFieldsOk

`func (o *UpdateContextActionRequestBodyFields) GetAllRecordFieldsOk() (*bool, bool)`

GetAllRecordFieldsOk returns a tuple with the AllRecordFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRecordFields

`func (o *UpdateContextActionRequestBodyFields) SetAllRecordFields(v bool)`

SetAllRecordFields sets AllRecordFields field to given value.

### HasAllRecordFields

`func (o *UpdateContextActionRequestBodyFields) HasAllRecordFields() bool`

HasAllRecordFields returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateContextActionRequestBodyFields) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateContextActionRequestBodyFields) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateContextActionRequestBodyFields) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateContextActionRequestBodyFields) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


