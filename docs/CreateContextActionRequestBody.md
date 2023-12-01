# CreateContextActionRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "URL"]
**Template** | Pointer to **string** |  | [optional] 
**IocTypes** | **[]string** |  | 
**EntityTypes** | Pointer to **[]string** |  | [optional] 
**RecordFields** | Pointer to **[]string** |  | [optional] 
**AllRecordFields** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateContextActionRequestBody

`func NewCreateContextActionRequestBody(name string, iocTypes []string, ) *CreateContextActionRequestBody`

NewCreateContextActionRequestBody instantiates a new CreateContextActionRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContextActionRequestBodyWithDefaults

`func NewCreateContextActionRequestBodyWithDefaults() *CreateContextActionRequestBody`

NewCreateContextActionRequestBodyWithDefaults instantiates a new CreateContextActionRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateContextActionRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateContextActionRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateContextActionRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateContextActionRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateContextActionRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateContextActionRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateContextActionRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTemplate

`func (o *CreateContextActionRequestBody) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CreateContextActionRequestBody) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CreateContextActionRequestBody) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *CreateContextActionRequestBody) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetIocTypes

`func (o *CreateContextActionRequestBody) GetIocTypes() []string`

GetIocTypes returns the IocTypes field if non-nil, zero value otherwise.

### GetIocTypesOk

`func (o *CreateContextActionRequestBody) GetIocTypesOk() (*[]string, bool)`

GetIocTypesOk returns a tuple with the IocTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIocTypes

`func (o *CreateContextActionRequestBody) SetIocTypes(v []string)`

SetIocTypes sets IocTypes field to given value.


### GetEntityTypes

`func (o *CreateContextActionRequestBody) GetEntityTypes() []string`

GetEntityTypes returns the EntityTypes field if non-nil, zero value otherwise.

### GetEntityTypesOk

`func (o *CreateContextActionRequestBody) GetEntityTypesOk() (*[]string, bool)`

GetEntityTypesOk returns a tuple with the EntityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypes

`func (o *CreateContextActionRequestBody) SetEntityTypes(v []string)`

SetEntityTypes sets EntityTypes field to given value.

### HasEntityTypes

`func (o *CreateContextActionRequestBody) HasEntityTypes() bool`

HasEntityTypes returns a boolean if a field has been set.

### GetRecordFields

`func (o *CreateContextActionRequestBody) GetRecordFields() []string`

GetRecordFields returns the RecordFields field if non-nil, zero value otherwise.

### GetRecordFieldsOk

`func (o *CreateContextActionRequestBody) GetRecordFieldsOk() (*[]string, bool)`

GetRecordFieldsOk returns a tuple with the RecordFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordFields

`func (o *CreateContextActionRequestBody) SetRecordFields(v []string)`

SetRecordFields sets RecordFields field to given value.

### HasRecordFields

`func (o *CreateContextActionRequestBody) HasRecordFields() bool`

HasRecordFields returns a boolean if a field has been set.

### GetAllRecordFields

`func (o *CreateContextActionRequestBody) GetAllRecordFields() bool`

GetAllRecordFields returns the AllRecordFields field if non-nil, zero value otherwise.

### GetAllRecordFieldsOk

`func (o *CreateContextActionRequestBody) GetAllRecordFieldsOk() (*bool, bool)`

GetAllRecordFieldsOk returns a tuple with the AllRecordFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRecordFields

`func (o *CreateContextActionRequestBody) SetAllRecordFields(v bool)`

SetAllRecordFields sets AllRecordFields field to given value.

### HasAllRecordFields

`func (o *CreateContextActionRequestBody) HasAllRecordFields() bool`

HasAllRecordFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


