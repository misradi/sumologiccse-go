# UpdateCustomEntityTypeRequestBodyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human friend and unique name. Examples: \&quot;Ip Address\&quot;, \&quot;Username\&quot;, \&quot;Mac Address\&quot;. | 
**Fields** | **[]string** | Record schema fields. Examples: \&quot;file_hash_md5\&quot;, \&quot;file_hash_sha1\&quot;. | 

## Methods

### NewUpdateCustomEntityTypeRequestBodyFields

`func NewUpdateCustomEntityTypeRequestBodyFields(name string, fields []string, ) *UpdateCustomEntityTypeRequestBodyFields`

NewUpdateCustomEntityTypeRequestBodyFields instantiates a new UpdateCustomEntityTypeRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomEntityTypeRequestBodyFieldsWithDefaults

`func NewUpdateCustomEntityTypeRequestBodyFieldsWithDefaults() *UpdateCustomEntityTypeRequestBodyFields`

NewUpdateCustomEntityTypeRequestBodyFieldsWithDefaults instantiates a new UpdateCustomEntityTypeRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCustomEntityTypeRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCustomEntityTypeRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCustomEntityTypeRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.


### GetFields

`func (o *UpdateCustomEntityTypeRequestBodyFields) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *UpdateCustomEntityTypeRequestBodyFields) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *UpdateCustomEntityTypeRequestBodyFields) SetFields(v []string)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


