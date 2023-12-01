# UpdateLogMappingRequestBodyFields

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

## Methods

### NewUpdateLogMappingRequestBodyFields

`func NewUpdateLogMappingRequestBodyFields(name string, fields []CreateLogMappingRequestBodyFieldsFieldsInner, recordType string, productGuid string, ) *UpdateLogMappingRequestBodyFields`

NewUpdateLogMappingRequestBodyFields instantiates a new UpdateLogMappingRequestBodyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLogMappingRequestBodyFieldsWithDefaults

`func NewUpdateLogMappingRequestBodyFieldsWithDefaults() *UpdateLogMappingRequestBodyFields`

NewUpdateLogMappingRequestBodyFieldsWithDefaults instantiates a new UpdateLogMappingRequestBodyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateLogMappingRequestBodyFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLogMappingRequestBodyFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLogMappingRequestBodyFields) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *UpdateLogMappingRequestBodyFields) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *UpdateLogMappingRequestBodyFields) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *UpdateLogMappingRequestBodyFields) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *UpdateLogMappingRequestBodyFields) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetFields

`func (o *UpdateLogMappingRequestBodyFields) GetFields() []CreateLogMappingRequestBodyFieldsFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *UpdateLogMappingRequestBodyFields) GetFieldsOk() (*[]CreateLogMappingRequestBodyFieldsFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *UpdateLogMappingRequestBodyFields) SetFields(v []CreateLogMappingRequestBodyFieldsFieldsInner)`

SetFields sets Fields field to given value.


### GetSkippedValues

`func (o *UpdateLogMappingRequestBodyFields) GetSkippedValues() []string`

GetSkippedValues returns the SkippedValues field if non-nil, zero value otherwise.

### GetSkippedValuesOk

`func (o *UpdateLogMappingRequestBodyFields) GetSkippedValuesOk() (*[]string, bool)`

GetSkippedValuesOk returns a tuple with the SkippedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedValues

`func (o *UpdateLogMappingRequestBodyFields) SetSkippedValues(v []string)`

SetSkippedValues sets SkippedValues field to given value.

### HasSkippedValues

`func (o *UpdateLogMappingRequestBodyFields) HasSkippedValues() bool`

HasSkippedValues returns a boolean if a field has been set.

### GetUnstructuredFields

`func (o *UpdateLogMappingRequestBodyFields) GetUnstructuredFields() CreateLogMappingRequestBodyFieldsUnstructuredFields`

GetUnstructuredFields returns the UnstructuredFields field if non-nil, zero value otherwise.

### GetUnstructuredFieldsOk

`func (o *UpdateLogMappingRequestBodyFields) GetUnstructuredFieldsOk() (*CreateLogMappingRequestBodyFieldsUnstructuredFields, bool)`

GetUnstructuredFieldsOk returns a tuple with the UnstructuredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnstructuredFields

`func (o *UpdateLogMappingRequestBodyFields) SetUnstructuredFields(v CreateLogMappingRequestBodyFieldsUnstructuredFields)`

SetUnstructuredFields sets UnstructuredFields field to given value.

### HasUnstructuredFields

`func (o *UpdateLogMappingRequestBodyFields) HasUnstructuredFields() bool`

HasUnstructuredFields returns a boolean if a field has been set.

### GetStructuredFields

`func (o *UpdateLogMappingRequestBodyFields) GetStructuredFields() CreateLogMappingRequestBodyFieldsStructuredFields`

GetStructuredFields returns the StructuredFields field if non-nil, zero value otherwise.

### GetStructuredFieldsOk

`func (o *UpdateLogMappingRequestBodyFields) GetStructuredFieldsOk() (*CreateLogMappingRequestBodyFieldsStructuredFields, bool)`

GetStructuredFieldsOk returns a tuple with the StructuredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredFields

`func (o *UpdateLogMappingRequestBodyFields) SetStructuredFields(v CreateLogMappingRequestBodyFieldsStructuredFields)`

SetStructuredFields sets StructuredFields field to given value.

### HasStructuredFields

`func (o *UpdateLogMappingRequestBodyFields) HasStructuredFields() bool`

HasStructuredFields returns a boolean if a field has been set.

### GetStructuredInputs

`func (o *UpdateLogMappingRequestBodyFields) GetStructuredInputs() []CreateLogMappingRequestBodyFieldsStructuredFields`

GetStructuredInputs returns the StructuredInputs field if non-nil, zero value otherwise.

### GetStructuredInputsOk

`func (o *UpdateLogMappingRequestBodyFields) GetStructuredInputsOk() (*[]CreateLogMappingRequestBodyFieldsStructuredFields, bool)`

GetStructuredInputsOk returns a tuple with the StructuredInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredInputs

`func (o *UpdateLogMappingRequestBodyFields) SetStructuredInputs(v []CreateLogMappingRequestBodyFieldsStructuredFields)`

SetStructuredInputs sets StructuredInputs field to given value.

### HasStructuredInputs

`func (o *UpdateLogMappingRequestBodyFields) HasStructuredInputs() bool`

HasStructuredInputs returns a boolean if a field has been set.

### GetRecordType

`func (o *UpdateLogMappingRequestBodyFields) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *UpdateLogMappingRequestBodyFields) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *UpdateLogMappingRequestBodyFields) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetProductGuid

`func (o *UpdateLogMappingRequestBodyFields) GetProductGuid() string`

GetProductGuid returns the ProductGuid field if non-nil, zero value otherwise.

### GetProductGuidOk

`func (o *UpdateLogMappingRequestBodyFields) GetProductGuidOk() (*string, bool)`

GetProductGuidOk returns a tuple with the ProductGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGuid

`func (o *UpdateLogMappingRequestBodyFields) SetProductGuid(v string)`

SetProductGuid sets ProductGuid field to given value.


### GetRelatesEntities

`func (o *UpdateLogMappingRequestBodyFields) GetRelatesEntities() bool`

GetRelatesEntities returns the RelatesEntities field if non-nil, zero value otherwise.

### GetRelatesEntitiesOk

`func (o *UpdateLogMappingRequestBodyFields) GetRelatesEntitiesOk() (*bool, bool)`

GetRelatesEntitiesOk returns a tuple with the RelatesEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatesEntities

`func (o *UpdateLogMappingRequestBodyFields) SetRelatesEntities(v bool)`

SetRelatesEntities sets RelatesEntities field to given value.

### HasRelatesEntities

`func (o *UpdateLogMappingRequestBodyFields) HasRelatesEntities() bool`

HasRelatesEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


