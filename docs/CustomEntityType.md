# CustomEntityType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Identifier** | **string** | Machine friendly and unique identifier. Examples: \&quot;ip\&quot;, \&quot;username\&quot;, \&quot;mac\&quot;. | 
**Name** | **string** | Human friend and unique name. Examples: \&quot;Ip Address\&quot;, \&quot;Username\&quot;, \&quot;Mac Address\&quot;. | 
**Fields** | **[]string** | Record schema fields. Examples: \&quot;file_hash_md5\&quot;, \&quot;file_hash_sha1\&quot;. | 

## Methods

### NewCustomEntityType

`func NewCustomEntityType(id string, identifier string, name string, fields []string, ) *CustomEntityType`

NewCustomEntityType instantiates a new CustomEntityType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEntityTypeWithDefaults

`func NewCustomEntityTypeWithDefaults() *CustomEntityType`

NewCustomEntityTypeWithDefaults instantiates a new CustomEntityType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomEntityType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomEntityType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomEntityType) SetId(v string)`

SetId sets Id field to given value.


### GetIdentifier

`func (o *CustomEntityType) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CustomEntityType) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CustomEntityType) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *CustomEntityType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEntityType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEntityType) SetName(v string)`

SetName sets Name field to given value.


### GetFields

`func (o *CustomEntityType) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CustomEntityType) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CustomEntityType) SetFields(v []string)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


