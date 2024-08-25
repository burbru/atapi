# Filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextFilters** | Pointer to [**[]Filter**](Filter.md) |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 
**Predicate** | Pointer to [**FilterPredicate**](FilterPredicate.md) |  | [optional] 

## Methods

### NewFilter

`func NewFilter() *Filter`

NewFilter instantiates a new Filter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterWithDefaults

`func NewFilterWithDefaults() *Filter`

NewFilterWithDefaults instantiates a new Filter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextFilters

`func (o *Filter) GetNextFilters() []Filter`

GetNextFilters returns the NextFilters field if non-nil, zero value otherwise.

### GetNextFiltersOk

`func (o *Filter) GetNextFiltersOk() (*[]Filter, bool)`

GetNextFiltersOk returns a tuple with the NextFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFilters

`func (o *Filter) SetNextFilters(v []Filter)`

SetNextFilters sets NextFilters field to given value.

### HasNextFilters

`func (o *Filter) HasNextFilters() bool`

HasNextFilters returns a boolean if a field has been set.

### GetOperator

`func (o *Filter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Filter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Filter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *Filter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetPredicate

`func (o *Filter) GetPredicate() FilterPredicate`

GetPredicate returns the Predicate field if non-nil, zero value otherwise.

### GetPredicateOk

`func (o *Filter) GetPredicateOk() (*FilterPredicate, bool)`

GetPredicateOk returns a tuple with the Predicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicate

`func (o *Filter) SetPredicate(v FilterPredicate)`

SetPredicate sets Predicate field to given value.

### HasPredicate

`func (o *Filter) HasPredicate() bool`

HasPredicate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


