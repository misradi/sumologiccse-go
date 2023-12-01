# CreateCustomEntityTypeRequestBodyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human friend and unique name. Examples: \&quot;Ip Address\&quot;, \&quot;Username\&quot;, \&quot;Mac Address\&quot;. | 
**Fields** | **[]string** | Record schema fields. Examples: \&quot;file_hash_md5\&quot;, \&quot;file_hash_sha1\&quot;. | 
**Identifier** | **string** | Machine friendly and unique identifier. Examples: \&quot;ip\&quot;, \&quot;username\&quot;, \&quot;mac\&quot;. | 

## Methods

### NewCreateCustomEntityTypeRequestBodyFields

`func NewCreateCustomEntityTypeRequestBodyFields(name string, fields []string, identifier string, ) *CreateCustomEntityTypeRequestBodyFields`

NewCreateCustomEntityTypeRequestBodyFields instantiates a new CreateCustomEntityTypeRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomEntityTypeRequestBodyFieldsWithDefaults

`func NewCreateCustomEntityTypeRequestBodyFieldsWithDefaults() *CreateCustomEntityTypeRequestBodyFields`

NewCreateCustomEntityTypeRequestBodyFieldsWithDefaults instantiates a new CreateCustomEntityTypeRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCustomEntityTypeRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCustomEntityTypeRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCustomEntityTypeRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.


### GetFields

`func (o *CreateCustomEntityTypeRequestBodyFields) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CreateCustomEntityTypeRequestBodyFields) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CreateCustomEntityTypeRequestBodyFields) SetFields(v []string)`

SetFields sets Fields field to given value.


### GetIdentifier

`func (o *CreateCustomEntityTypeRequestBodyFields) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CreateCustomEntityTypeRequestBodyFields) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CreateCustomEntityTypeRequestBodyFields) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


