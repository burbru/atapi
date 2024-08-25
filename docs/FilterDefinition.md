# FilterDefinition

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

### NewFilterDefinition

`func NewFilterDefinition() *FilterDefinition`

NewFilterDefinition instantiates a new FilterDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterDefinitionWithDefaults

`func NewFilterDefinitionWithDefaults() *FilterDefinition`

NewFilterDefinitionWithDefaults instantiates a new FilterDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedValues

`func (o *FilterDefinition) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *FilterDefinition) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *FilterDefinition) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *FilterDefinition) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### GetDefaultFilterOperator

`func (o *FilterDefinition) GetDefaultFilterOperator() string`

GetDefaultFilterOperator returns the DefaultFilterOperator field if non-nil, zero value otherwise.

### GetDefaultFilterOperatorOk

`func (o *FilterDefinition) GetDefaultFilterOperatorOk() (*string, bool)`

GetDefaultFilterOperatorOk returns a tuple with the DefaultFilterOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFilterOperator

`func (o *FilterDefinition) SetDefaultFilterOperator(v string)`

SetDefaultFilterOperator sets DefaultFilterOperator field to given value.

### HasDefaultFilterOperator

`func (o *FilterDefinition) HasDefaultFilterOperator() bool`

HasDefaultFilterOperator returns a boolean if a field has been set.

### GetDefaultParameters

`func (o *FilterDefinition) GetDefaultParameters() []string`

GetDefaultParameters returns the DefaultParameters field if non-nil, zero value otherwise.

### GetDefaultParametersOk

`func (o *FilterDefinition) GetDefaultParametersOk() (*[]string, bool)`

GetDefaultParametersOk returns a tuple with the DefaultParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultParameters

`func (o *FilterDefinition) SetDefaultParameters(v []string)`

SetDefaultParameters sets DefaultParameters field to given value.

### HasDefaultParameters

`func (o *FilterDefinition) HasDefaultParameters() bool`

HasDefaultParameters returns a boolean if a field has been set.

### GetDescription

`func (o *FilterDefinition) GetDescription() MessagePrototype`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FilterDefinition) GetDescriptionOk() (*MessagePrototype, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FilterDefinition) SetDescription(v MessagePrototype)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FilterDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFieldName

`func (o *FilterDefinition) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *FilterDefinition) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *FilterDefinition) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *FilterDefinition) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetFilterOperators

`func (o *FilterDefinition) GetFilterOperators() []string`

GetFilterOperators returns the FilterOperators field if non-nil, zero value otherwise.

### GetFilterOperatorsOk

`func (o *FilterDefinition) GetFilterOperatorsOk() (*[]string, bool)`

GetFilterOperatorsOk returns a tuple with the FilterOperators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterOperators

`func (o *FilterDefinition) SetFilterOperators(v []string)`

SetFilterOperators sets FilterOperators field to given value.

### HasFilterOperators

`func (o *FilterDefinition) HasFilterOperators() bool`

HasFilterOperators returns a boolean if a field has been set.

### GetNextFilterDefinitions

`func (o *FilterDefinition) GetNextFilterDefinitions() []FilterDefinition`

GetNextFilterDefinitions returns the NextFilterDefinitions field if non-nil, zero value otherwise.

### GetNextFilterDefinitionsOk

`func (o *FilterDefinition) GetNextFilterDefinitionsOk() (*[]FilterDefinition, bool)`

GetNextFilterDefinitionsOk returns a tuple with the NextFilterDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFilterDefinitions

`func (o *FilterDefinition) SetNextFilterDefinitions(v []FilterDefinition)`

SetNextFilterDefinitions sets NextFilterDefinitions field to given value.

### HasNextFilterDefinitions

`func (o *FilterDefinition) HasNextFilterDefinitions() bool`

HasNextFilterDefinitions returns a boolean if a field has been set.

### GetParameterType

`func (o *FilterDefinition) GetParameterType() string`

GetParameterType returns the ParameterType field if non-nil, zero value otherwise.

### GetParameterTypeOk

`func (o *FilterDefinition) GetParameterTypeOk() (*string, bool)`

GetParameterTypeOk returns a tuple with the ParameterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterType

`func (o *FilterDefinition) SetParameterType(v string)`

SetParameterType sets ParameterType field to given value.

### HasParameterType

`func (o *FilterDefinition) HasParameterType() bool`

HasParameterType returns a boolean if a field has been set.

### GetPredicateOperators

`func (o *FilterDefinition) GetPredicateOperators() []string`

GetPredicateOperators returns the PredicateOperators field if non-nil, zero value otherwise.

### GetPredicateOperatorsOk

`func (o *FilterDefinition) GetPredicateOperatorsOk() (*[]string, bool)`

GetPredicateOperatorsOk returns a tuple with the PredicateOperators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicateOperators

`func (o *FilterDefinition) SetPredicateOperators(v []string)`

SetPredicateOperators sets PredicateOperators field to given value.

### HasPredicateOperators

`func (o *FilterDefinition) HasPredicateOperators() bool`

HasPredicateOperators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


