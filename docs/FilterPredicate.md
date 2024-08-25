# FilterPredicate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 
**Parameter** | Pointer to **string** |  | [optional] 

## Methods

### NewFilterPredicate

`func NewFilterPredicate() *FilterPredicate`

NewFilterPredicate instantiates a new FilterPredicate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterPredicateWithDefaults

`func NewFilterPredicateWithDefaults() *FilterPredicate`

NewFilterPredicateWithDefaults instantiates a new FilterPredicate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FilterPredicate) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FilterPredicate) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FilterPredicate) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *FilterPredicate) HasField() bool`

HasField returns a boolean if a field has been set.

### GetOperator

`func (o *FilterPredicate) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *FilterPredicate) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *FilterPredicate) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *FilterPredicate) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetParameter

`func (o *FilterPredicate) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *FilterPredicate) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *FilterPredicate) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *FilterPredicate) HasParameter() bool`

HasParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


