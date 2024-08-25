# ListingFilterDefinition

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

### NewListingFilterDefinition

`func NewListingFilterDefinition() *ListingFilterDefinition`

NewListingFilterDefinition instantiates a new ListingFilterDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingFilterDefinitionWithDefaults

`func NewListingFilterDefinitionWithDefaults() *ListingFilterDefinition`

NewListingFilterDefinitionWithDefaults instantiates a new ListingFilterDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedValues

`func (o *ListingFilterDefinition) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *ListingFilterDefinition) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *ListingFilterDefinition) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *ListingFilterDefinition) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### GetDefaultFilterOperator

`func (o *ListingFilterDefinition) GetDefaultFilterOperator() string`

GetDefaultFilterOperator returns the DefaultFilterOperator field if non-nil, zero value otherwise.

### GetDefaultFilterOperatorOk

`func (o *ListingFilterDefinition) GetDefaultFilterOperatorOk() (*string, bool)`

GetDefaultFilterOperatorOk returns a tuple with the DefaultFilterOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFilterOperator

`func (o *ListingFilterDefinition) SetDefaultFilterOperator(v string)`

SetDefaultFilterOperator sets DefaultFilterOperator field to given value.

### HasDefaultFilterOperator

`func (o *ListingFilterDefinition) HasDefaultFilterOperator() bool`

HasDefaultFilterOperator returns a boolean if a field has been set.

### GetDefaultParameters

`func (o *ListingFilterDefinition) GetDefaultParameters() []string`

GetDefaultParameters returns the DefaultParameters field if non-nil, zero value otherwise.

### GetDefaultParametersOk

`func (o *ListingFilterDefinition) GetDefaultParametersOk() (*[]string, bool)`

GetDefaultParametersOk returns a tuple with the DefaultParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultParameters

`func (o *ListingFilterDefinition) SetDefaultParameters(v []string)`

SetDefaultParameters sets DefaultParameters field to given value.

### HasDefaultParameters

`func (o *ListingFilterDefinition) HasDefaultParameters() bool`

HasDefaultParameters returns a boolean if a field has been set.

### GetDescription

`func (o *ListingFilterDefinition) GetDescription() MessagePrototype`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListingFilterDefinition) GetDescriptionOk() (*MessagePrototype, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListingFilterDefinition) SetDescription(v MessagePrototype)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListingFilterDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFieldName

`func (o *ListingFilterDefinition) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ListingFilterDefinition) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ListingFilterDefinition) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *ListingFilterDefinition) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetFilterOperators

`func (o *ListingFilterDefinition) GetFilterOperators() []string`

GetFilterOperators returns the FilterOperators field if non-nil, zero value otherwise.

### GetFilterOperatorsOk

`func (o *ListingFilterDefinition) GetFilterOperatorsOk() (*[]string, bool)`

GetFilterOperatorsOk returns a tuple with the FilterOperators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterOperators

`func (o *ListingFilterDefinition) SetFilterOperators(v []string)`

SetFilterOperators sets FilterOperators field to given value.

### HasFilterOperators

`func (o *ListingFilterDefinition) HasFilterOperators() bool`

HasFilterOperators returns a boolean if a field has been set.

### GetNextFilterDefinitions

`func (o *ListingFilterDefinition) GetNextFilterDefinitions() []FilterDefinition`

GetNextFilterDefinitions returns the NextFilterDefinitions field if non-nil, zero value otherwise.

### GetNextFilterDefinitionsOk

`func (o *ListingFilterDefinition) GetNextFilterDefinitionsOk() (*[]FilterDefinition, bool)`

GetNextFilterDefinitionsOk returns a tuple with the NextFilterDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFilterDefinitions

`func (o *ListingFilterDefinition) SetNextFilterDefinitions(v []FilterDefinition)`

SetNextFilterDefinitions sets NextFilterDefinitions field to given value.

### HasNextFilterDefinitions

`func (o *ListingFilterDefinition) HasNextFilterDefinitions() bool`

HasNextFilterDefinitions returns a boolean if a field has been set.

### GetParameterType

`func (o *ListingFilterDefinition) GetParameterType() string`

GetParameterType returns the ParameterType field if non-nil, zero value otherwise.

### GetParameterTypeOk

`func (o *ListingFilterDefinition) GetParameterTypeOk() (*string, bool)`

GetParameterTypeOk returns a tuple with the ParameterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterType

`func (o *ListingFilterDefinition) SetParameterType(v string)`

SetParameterType sets ParameterType field to given value.

### HasParameterType

`func (o *ListingFilterDefinition) HasParameterType() bool`

HasParameterType returns a boolean if a field has been set.

### GetPredicateOperators

`func (o *ListingFilterDefinition) GetPredicateOperators() []string`

GetPredicateOperators returns the PredicateOperators field if non-nil, zero value otherwise.

### GetPredicateOperatorsOk

`func (o *ListingFilterDefinition) GetPredicateOperatorsOk() (*[]string, bool)`

GetPredicateOperatorsOk returns a tuple with the PredicateOperators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicateOperators

`func (o *ListingFilterDefinition) SetPredicateOperators(v []string)`

SetPredicateOperators sets PredicateOperators field to given value.

### HasPredicateOperators

`func (o *ListingFilterDefinition) HasPredicateOperators() bool`

HasPredicateOperators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


