# SpreadFilterDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedValues** | Pointer to **[]string** |  | [optional] 
**DefaultFilterOperator** | Pointer to **string** |  | [optional] 
**DefaultParameters** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to [**MessagePrototype**](MessagePrototype.md) |  | [optional] 
**FieldName** | Pointer to **string** |  | [optional] 
**FilterOperators** | Pointer to **[]string** |  | [optional] 
**NextFilterDefinitions** | Pointer to [**[]FilterDefinition**](FilterDefinition.md) |  | [optional] 
**ParameterType** | Pointer to **string** |  | [optional] 
**PredicateOperators** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSpreadFilterDefinition

`func NewSpreadFilterDefinition() *SpreadFilterDefinition`

NewSpreadFilterDefinition instantiates a new SpreadFilterDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpreadFilterDefinitionWithDefaults

`func NewSpreadFilterDefinitionWithDefaults() *SpreadFilterDefinition`

NewSpreadFilterDefinitionWithDefaults instantiates a new SpreadFilterDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedValues

`func (o *SpreadFilterDefinition) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *SpreadFilterDefinition) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *SpreadFilterDefinition) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *SpreadFilterDefinition) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### GetDefaultFilterOperator

`func (o *SpreadFilterDefinition) GetDefaultFilterOperator() string`

GetDefaultFilterOperator returns the DefaultFilterOperator field if non-nil, zero value otherwise.

### GetDefaultFilterOperatorOk

`func (o *SpreadFilterDefinition) GetDefaultFilterOperatorOk() (*string, bool)`

GetDefaultFilterOperatorOk returns a tuple with the DefaultFilterOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFilterOperator

`func (o *SpreadFilterDefinition) SetDefaultFilterOperator(v string)`

SetDefaultFilterOperator sets DefaultFilterOperator field to given value.

### HasDefaultFilterOperator

`func (o *SpreadFilterDefinition) HasDefaultFilterOperator() bool`

HasDefaultFilterOperator returns a boolean if a field has been set.

### GetDefaultParameters

`func (o *SpreadFilterDefinition) GetDefaultParameters() []string`

GetDefaultParameters returns the DefaultParameters field if non-nil, zero value otherwise.

### GetDefaultParametersOk

`func (o *SpreadFilterDefinition) GetDefaultParametersOk() (*[]string, bool)`

GetDefaultParametersOk returns a tuple with the DefaultParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultParameters

`func (o *SpreadFilterDefinition) SetDefaultParameters(v []string)`

SetDefaultParameters sets DefaultParameters field to given value.

### HasDefaultParameters

`func (o *SpreadFilterDefinition) HasDefaultParameters() bool`

HasDefaultParameters returns a boolean if a field has been set.

### GetDescription

`func (o *SpreadFilterDefinition) GetDescription() MessagePrototype`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SpreadFilterDefinition) GetDescriptionOk() (*MessagePrototype, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SpreadFilterDefinition) SetDescription(v MessagePrototype)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SpreadFilterDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFieldName

`func (o *SpreadFilterDefinition) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *SpreadFilterDefinition) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *SpreadFilterDefinition) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *SpreadFilterDefinition) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetFilterOperators

`func (o *SpreadFilterDefinition) GetFilterOperators() []string`

GetFilterOperators returns the FilterOperators field if non-nil, zero value otherwise.

### GetFilterOperatorsOk

`func (o *SpreadFilterDefinition) GetFilterOperatorsOk() (*[]string, bool)`

GetFilterOperatorsOk returns a tuple with the FilterOperators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterOperators

`func (o *SpreadFilterDefinition) SetFilterOperators(v []string)`

SetFilterOperators sets FilterOperators field to given value.

### HasFilterOperators

`func (o *SpreadFilterDefinition) HasFilterOperators() bool`

HasFilterOperators returns a boolean if a field has been set.

### GetNextFilterDefinitions

`func (o *SpreadFilterDefinition) GetNextFilterDefinitions() []FilterDefinition`

GetNextFilterDefinitions returns the NextFilterDefinitions field if non-nil, zero value otherwise.

### GetNextFilterDefinitionsOk

`func (o *SpreadFilterDefinition) GetNextFilterDefinitionsOk() (*[]FilterDefinition, bool)`

GetNextFilterDefinitionsOk returns a tuple with the NextFilterDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFilterDefinitions

`func (o *SpreadFilterDefinition) SetNextFilterDefinitions(v []FilterDefinition)`

SetNextFilterDefinitions sets NextFilterDefinitions field to given value.

### HasNextFilterDefinitions

`func (o *SpreadFilterDefinition) HasNextFilterDefinitions() bool`

HasNextFilterDefinitions returns a boolean if a field has been set.

### GetParameterType

`func (o *SpreadFilterDefinition) GetParameterType() string`

GetParameterType returns the ParameterType field if non-nil, zero value otherwise.

### GetParameterTypeOk

`func (o *SpreadFilterDefinition) GetParameterTypeOk() (*string, bool)`

GetParameterTypeOk returns a tuple with the ParameterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterType

`func (o *SpreadFilterDefinition) SetParameterType(v string)`

SetParameterType sets ParameterType field to given value.

### HasParameterType

`func (o *SpreadFilterDefinition) HasParameterType() bool`

HasParameterType returns a boolean if a field has been set.

### GetPredicateOperators

`func (o *SpreadFilterDefinition) GetPredicateOperators() []string`

GetPredicateOperators returns the PredicateOperators field if non-nil, zero value otherwise.

### GetPredicateOperatorsOk

`func (o *SpreadFilterDefinition) GetPredicateOperatorsOk() (*[]string, bool)`

GetPredicateOperatorsOk returns a tuple with the PredicateOperators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicateOperators

`func (o *SpreadFilterDefinition) SetPredicateOperators(v []string)`

SetPredicateOperators sets PredicateOperators field to given value.

### HasPredicateOperators

`func (o *SpreadFilterDefinition) HasPredicateOperators() bool`

HasPredicateOperators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


