# LogMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Source** | **string** |  | 
**SkippedValues** | **[]string** |  | 
**Fields** | [**[]CreateLogMappingRequestBodyFieldsFieldsInner**](CreateLogMappingRequestBodyFieldsFieldsInner.md) |  | 
**Enabled** | **bool** |  | 
**RelatesEntities** | **bool** |  | 
**UnstructuredFields** | Pointer to [**CreateLogMappingRequestBodyFieldsUnstructuredFields**](CreateLogMappingRequestBodyFieldsUnstructuredFields.md) |  | [optional] 
**StructuredInputs** | [**[]LogMappingStructuredInputsInner**](LogMappingStructuredInputsInner.md) |  | 
**RecordType** | **string** |  | 
**ProductGuid** | Pointer to **string** |  | [optional] 
**Input** | Pointer to **map[string]interface{}** |  | [optional] 
**Output** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewLogMapping

`func NewLogMapping(id string, name string, source string, skippedValues []string, fields []CreateLogMappingRequestBodyFieldsFieldsInner, enabled bool, relatesEntities bool, structuredInputs []LogMappingStructuredInputsInner, recordType string, ) *LogMapping`

NewLogMapping instantiates a new LogMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogMappingWithDefaults

`func NewLogMappingWithDefaults() *LogMapping`

NewLogMappingWithDefaults instantiates a new LogMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogMapping) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogMapping) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogMapping) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LogMapping) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogMapping) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogMapping) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *LogMapping) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LogMapping) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LogMapping) SetSource(v string)`

SetSource sets Source field to given value.


### GetSkippedValues

`func (o *LogMapping) GetSkippedValues() []string`

GetSkippedValues returns the SkippedValues field if non-nil, zero value otherwise.

### GetSkippedValuesOk

`func (o *LogMapping) GetSkippedValuesOk() (*[]string, bool)`

GetSkippedValuesOk returns a tuple with the SkippedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedValues

`func (o *LogMapping) SetSkippedValues(v []string)`

SetSkippedValues sets SkippedValues field to given value.


### GetFields

`func (o *LogMapping) GetFields() []CreateLogMappingRequestBodyFieldsFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *LogMapping) GetFieldsOk() (*[]CreateLogMappingRequestBodyFieldsFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *LogMapping) SetFields(v []CreateLogMappingRequestBodyFieldsFieldsInner)`

SetFields sets Fields field to given value.


### GetEnabled

`func (o *LogMapping) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LogMapping) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LogMapping) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRelatesEntities

`func (o *LogMapping) GetRelatesEntities() bool`

GetRelatesEntities returns the RelatesEntities field if non-nil, zero value otherwise.

### GetRelatesEntitiesOk

`func (o *LogMapping) GetRelatesEntitiesOk() (*bool, bool)`

GetRelatesEntitiesOk returns a tuple with the RelatesEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatesEntities

`func (o *LogMapping) SetRelatesEntities(v bool)`

SetRelatesEntities sets RelatesEntities field to given value.


### GetUnstructuredFields

`func (o *LogMapping) GetUnstructuredFields() CreateLogMappingRequestBodyFieldsUnstructuredFields`

GetUnstructuredFields returns the UnstructuredFields field if non-nil, zero value otherwise.

### GetUnstructuredFieldsOk

`func (o *LogMapping) GetUnstructuredFieldsOk() (*CreateLogMappingRequestBodyFieldsUnstructuredFields, bool)`

GetUnstructuredFieldsOk returns a tuple with the UnstructuredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnstructuredFields

`func (o *LogMapping) SetUnstructuredFields(v CreateLogMappingRequestBodyFieldsUnstructuredFields)`

SetUnstructuredFields sets UnstructuredFields field to given value.

### HasUnstructuredFields

`func (o *LogMapping) HasUnstructuredFields() bool`

HasUnstructuredFields returns a boolean if a field has been set.

### GetStructuredInputs

`func (o *LogMapping) GetStructuredInputs() []LogMappingStructuredInputsInner`

GetStructuredInputs returns the StructuredInputs field if non-nil, zero value otherwise.

### GetStructuredInputsOk

`func (o *LogMapping) GetStructuredInputsOk() (*[]LogMappingStructuredInputsInner, bool)`

GetStructuredInputsOk returns a tuple with the StructuredInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredInputs

`func (o *LogMapping) SetStructuredInputs(v []LogMappingStructuredInputsInner)`

SetStructuredInputs sets StructuredInputs field to given value.


### GetRecordType

`func (o *LogMapping) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *LogMapping) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *LogMapping) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetProductGuid

`func (o *LogMapping) GetProductGuid() string`

GetProductGuid returns the ProductGuid field if non-nil, zero value otherwise.

### GetProductGuidOk

`func (o *LogMapping) GetProductGuidOk() (*string, bool)`

GetProductGuidOk returns a tuple with the ProductGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGuid

`func (o *LogMapping) SetProductGuid(v string)`

SetProductGuid sets ProductGuid field to given value.

### HasProductGuid

`func (o *LogMapping) HasProductGuid() bool`

HasProductGuid returns a boolean if a field has been set.

### GetInput

`func (o *LogMapping) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *LogMapping) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *LogMapping) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *LogMapping) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *LogMapping) GetOutput() map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *LogMapping) GetOutputOk() (*map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *LogMapping) SetOutput(v map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *LogMapping) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


