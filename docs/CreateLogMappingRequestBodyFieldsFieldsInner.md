# CreateLogMappingRequestBodyFieldsFieldsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | Pointer to **string** |  | [optional] 
**ValueType** | Pointer to **string** |  | [optional] 
**SkippedValues** | Pointer to **[]string** |  | [optional] 
**Lookup** | Pointer to [**[]CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner**](CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner.md) |  | [optional] 
**DefaultValue** | Pointer to **string** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**CaseInsensitive** | Pointer to **bool** |  | [optional] 
**AlternateValues** | Pointer to **[]string** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**SplitDelimiter** | Pointer to **string** |  | [optional] 
**SplitIndex** | Pointer to **string** |  | [optional] 
**FieldJoin** | Pointer to **[]string** |  | [optional] 
**JoinDelimiter** | Pointer to **string** |  | [optional] 
**FormatParameters** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateLogMappingRequestBodyFieldsFieldsInner

`func NewCreateLogMappingRequestBodyFieldsFieldsInner(name string, ) *CreateLogMappingRequestBodyFieldsFieldsInner`

NewCreateLogMappingRequestBodyFieldsFieldsInner instantiates a new CreateLogMappingRequestBodyFieldsFieldsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogMappingRequestBodyFieldsFieldsInnerWithDefaults

`func NewCreateLogMappingRequestBodyFieldsFieldsInnerWithDefaults() *CreateLogMappingRequestBodyFieldsFieldsInner`

NewCreateLogMappingRequestBodyFieldsFieldsInnerWithDefaults instantiates a new CreateLogMappingRequestBodyFieldsFieldsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueType

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetSkippedValues

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetSkippedValues() []string`

GetSkippedValues returns the SkippedValues field if non-nil, zero value otherwise.

### GetSkippedValuesOk

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetSkippedValuesOk() (*[]string, bool)`

GetSkippedValuesOk returns a tuple with the SkippedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedValues

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetSkippedValues(v []string)`

SetSkippedValues sets SkippedValues field to given value.

### HasSkippedValues

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasSkippedValues() bool`

HasSkippedValues returns a boolean if a field has been set.

### GetLookup

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetLookup() []CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner`

GetLookup returns the Lookup field if non-nil, zero value otherwise.

### GetLookupOk

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetLookupOk() (*[]CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner, bool)`

GetLookupOk returns a tuple with the Lookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookup

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetLookup(v []CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner)`

SetLookup sets Lookup field to given value.

### HasLookup

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasLookup() bool`

HasLookup returns a boolean if a field has been set.

### GetDefaultValue

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetFormat

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetCaseInsensitive

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetCaseInsensitive() bool`

GetCaseInsensitive returns the CaseInsensitive field if non-nil, zero value otherwise.

### GetCaseInsensitiveOk

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetCaseInsensitiveOk() (*bool, bool)`

GetCaseInsensitiveOk returns a tuple with the CaseInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInsensitive

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetCaseInsensitive(v bool)`

SetCaseInsensitive sets CaseInsensitive field to given value.

### HasCaseInsensitive

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasCaseInsensitive() bool`

HasCaseInsensitive returns a boolean if a field has been set.

### GetAlternateValues

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetAlternateValues() []string`

GetAlternateValues returns the AlternateValues field if non-nil, zero value otherwise.

### GetAlternateValuesOk

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetAlternateValuesOk() (*[]string, bool)`

GetAlternateValuesOk returns a tuple with the AlternateValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateValues

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetAlternateValues(v []string)`

SetAlternateValues sets AlternateValues field to given value.

### HasAlternateValues

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasAlternateValues() bool`

HasAlternateValues returns a boolean if a field has been set.

### GetTimeZone

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetSplitDelimiter

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetSplitDelimiter() string`

GetSplitDelimiter returns the SplitDelimiter field if non-nil, zero value otherwise.

### GetSplitDelimiterOk

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetSplitDelimiterOk() (*string, bool)`

GetSplitDelimiterOk returns a tuple with the SplitDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitDelimiter

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetSplitDelimiter(v string)`

SetSplitDelimiter sets SplitDelimiter field to given value.

### HasSplitDelimiter

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasSplitDelimiter() bool`

HasSplitDelimiter returns a boolean if a field has been set.

### GetSplitIndex

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetSplitIndex() string`

GetSplitIndex returns the SplitIndex field if non-nil, zero value otherwise.

### GetSplitIndexOk

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetSplitIndexOk() (*string, bool)`

GetSplitIndexOk returns a tuple with the SplitIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitIndex

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetSplitIndex(v string)`

SetSplitIndex sets SplitIndex field to given value.

### HasSplitIndex

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasSplitIndex() bool`

HasSplitIndex returns a boolean if a field has been set.

### GetFieldJoin

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetFieldJoin() []string`

GetFieldJoin returns the FieldJoin field if non-nil, zero value otherwise.

### GetFieldJoinOk

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetFieldJoinOk() (*[]string, bool)`

GetFieldJoinOk returns a tuple with the FieldJoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldJoin

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetFieldJoin(v []string)`

SetFieldJoin sets FieldJoin field to given value.

### HasFieldJoin

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasFieldJoin() bool`

HasFieldJoin returns a boolean if a field has been set.

### GetJoinDelimiter

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetJoinDelimiter() string`

GetJoinDelimiter returns the JoinDelimiter field if non-nil, zero value otherwise.

### GetJoinDelimiterOk

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetJoinDelimiterOk() (*string, bool)`

GetJoinDelimiterOk returns a tuple with the JoinDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinDelimiter

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetJoinDelimiter(v string)`

SetJoinDelimiter sets JoinDelimiter field to given value.

### HasJoinDelimiter

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasJoinDelimiter() bool`

HasJoinDelimiter returns a boolean if a field has been set.

### GetFormatParameters

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetFormatParameters() []string`

GetFormatParameters returns the FormatParameters field if non-nil, zero value otherwise.

### GetFormatParametersOk

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetFormatParametersOk() (*[]string, bool)`

GetFormatParametersOk returns a tuple with the FormatParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatParameters

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetFormatParameters(v []string)`

SetFormatParameters sets FormatParameters field to given value.

### HasFormatParameters

`func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasFormatParameters() bool`

HasFormatParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


