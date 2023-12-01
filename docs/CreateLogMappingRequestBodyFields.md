# CreateLogMappingRequestBodyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**Fields** | [**[]CreateLogMappingRequestBodyFieldsFieldsInner**](CreateLogMappingRequestBodyFieldsFieldsInner.md) |  | 
**SkippedValues** | Pointer to **[]string** |  | [optional] 
**UnstructuredFields** | Pointer to [**CreateLogMappingRequestBodyFieldsUnstructuredFields**](CreateLogMappingRequestBodyFieldsUnstructuredFields.md) |  | [optional] 
**StructuredFields** | Pointer to [**CreateLogMappingRequestBodyFieldsStructuredFields**](CreateLogMappingRequestBodyFieldsStructuredFields.md) |  | [optional] 
**StructuredInputs** | Pointer to [**[]CreateLogMappingRequestBodyFieldsStructuredFields**](CreateLogMappingRequestBodyFieldsStructuredFields.md) |  | [optional] 
**RecordType** | **string** |  | 
**ProductGuid** | **string** |  | 
**RelatesEntities** | Pointer to **bool** |  | [optional] 
**Enabled** | **bool** |  | 

## Methods

### NewCreateLogMappingRequestBodyFields

`func NewCreateLogMappingRequestBodyFields(name string, fields []CreateLogMappingRequestBodyFieldsFieldsInner, recordType string, productGuid string, enabled bool, ) *CreateLogMappingRequestBodyFields`

NewCreateLogMappingRequestBodyFields instantiates a new CreateLogMappingRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogMappingRequestBodyFieldsWithDefaults

`func NewCreateLogMappingRequestBodyFieldsWithDefaults() *CreateLogMappingRequestBodyFields`

NewCreateLogMappingRequestBodyFieldsWithDefaults instantiates a new CreateLogMappingRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateLogMappingRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLogMappingRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLogMappingRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *CreateLogMappingRequestBodyFields) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateLogMappingRequestBodyFields) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateLogMappingRequestBodyFields) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateLogMappingRequestBodyFields) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetFields

`func (o *CreateLogMappingRequestBodyFields) GetFields() []CreateLogMappingRequestBodyFieldsFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CreateLogMappingRequestBodyFields) GetFieldsOk() (*[]CreateLogMappingRequestBodyFieldsFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CreateLogMappingRequestBodyFields) SetFields(v []CreateLogMappingRequestBodyFieldsFieldsInner)`

SetFields sets Fields field to given value.


### GetSkippedValues

`func (o *CreateLogMappingRequestBodyFields) GetSkippedValues() []string`

GetSkippedValues returns the SkippedValues field if non-nil, zero value otherwise.

### GetSkippedValuesOk

`func (o *CreateLogMappingRequestBodyFields) GetSkippedValuesOk() (*[]string, bool)`

GetSkippedValuesOk returns a tuple with the SkippedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedValues

`func (o *CreateLogMappingRequestBodyFields) SetSkippedValues(v []string)`

SetSkippedValues sets SkippedValues field to given value.

### HasSkippedValues

`func (o *CreateLogMappingRequestBodyFields) HasSkippedValues() bool`

HasSkippedValues returns a boolean if a field has been set.

### GetUnstructuredFields

`func (o *CreateLogMappingRequestBodyFields) GetUnstructuredFields() CreateLogMappingRequestBodyFieldsUnstructuredFields`

GetUnstructuredFields returns the UnstructuredFields field if non-nil, zero value otherwise.

### GetUnstructuredFieldsOk

`func (o *CreateLogMappingRequestBodyFields) GetUnstructuredFieldsOk() (*CreateLogMappingRequestBodyFieldsUnstructuredFields, bool)`

GetUnstructuredFieldsOk returns a tuple with the UnstructuredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnstructuredFields

`func (o *CreateLogMappingRequestBodyFields) SetUnstructuredFields(v CreateLogMappingRequestBodyFieldsUnstructuredFields)`

SetUnstructuredFields sets UnstructuredFields field to given value.

### HasUnstructuredFields

`func (o *CreateLogMappingRequestBodyFields) HasUnstructuredFields() bool`

HasUnstructuredFields returns a boolean if a field has been set.

### GetStructuredFields

`func (o *CreateLogMappingRequestBodyFields) GetStructuredFields() CreateLogMappingRequestBodyFieldsStructuredFields`

GetStructuredFields returns the StructuredFields field if non-nil, zero value otherwise.

### GetStructuredFieldsOk

`func (o *CreateLogMappingRequestBodyFields) GetStructuredFieldsOk() (*CreateLogMappingRequestBodyFieldsStructuredFields, bool)`

GetStructuredFieldsOk returns a tuple with the StructuredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredFields

`func (o *CreateLogMappingRequestBodyFields) SetStructuredFields(v CreateLogMappingRequestBodyFieldsStructuredFields)`

SetStructuredFields sets StructuredFields field to given value.

### HasStructuredFields

`func (o *CreateLogMappingRequestBodyFields) HasStructuredFields() bool`

HasStructuredFields returns a boolean if a field has been set.

### GetStructuredInputs

`func (o *CreateLogMappingRequestBodyFields) GetStructuredInputs() []CreateLogMappingRequestBodyFieldsStructuredFields`

GetStructuredInputs returns the StructuredInputs field if non-nil, zero value otherwise.

### GetStructuredInputsOk

`func (o *CreateLogMappingRequestBodyFields) GetStructuredInputsOk() (*[]CreateLogMappingRequestBodyFieldsStructuredFields, bool)`

GetStructuredInputsOk returns a tuple with the StructuredInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredInputs

`func (o *CreateLogMappingRequestBodyFields) SetStructuredInputs(v []CreateLogMappingRequestBodyFieldsStructuredFields)`

SetStructuredInputs sets StructuredInputs field to given value.

### HasStructuredInputs

`func (o *CreateLogMappingRequestBodyFields) HasStructuredInputs() bool`

HasStructuredInputs returns a boolean if a field has been set.

### GetRecordType

`func (o *CreateLogMappingRequestBodyFields) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CreateLogMappingRequestBodyFields) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CreateLogMappingRequestBodyFields) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetProductGuid

`func (o *CreateLogMappingRequestBodyFields) GetProductGuid() string`

GetProductGuid returns the ProductGuid field if non-nil, zero value otherwise.

### GetProductGuidOk

`func (o *CreateLogMappingRequestBodyFields) GetProductGuidOk() (*string, bool)`

GetProductGuidOk returns a tuple with the ProductGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGuid

`func (o *CreateLogMappingRequestBodyFields) SetProductGuid(v string)`

SetProductGuid sets ProductGuid field to given value.


### GetRelatesEntities

`func (o *CreateLogMappingRequestBodyFields) GetRelatesEntities() bool`

GetRelatesEntities returns the RelatesEntities field if non-nil, zero value otherwise.

### GetRelatesEntitiesOk

`func (o *CreateLogMappingRequestBodyFields) GetRelatesEntitiesOk() (*bool, bool)`

GetRelatesEntitiesOk returns a tuple with the RelatesEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatesEntities

`func (o *CreateLogMappingRequestBodyFields) SetRelatesEntities(v bool)`

SetRelatesEntities sets RelatesEntities field to given value.

### HasRelatesEntities

`func (o *CreateLogMappingRequestBodyFields) HasRelatesEntities() bool`

HasRelatesEntities returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateLogMappingRequestBodyFields) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateLogMappingRequestBodyFields) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateLogMappingRequestBodyFields) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


