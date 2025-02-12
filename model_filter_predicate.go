/*
Api Documentation

Api Documentation

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package atapi

import (
	"encoding/json"
)

// checks if the FilterPredicate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilterPredicate{}

// FilterPredicate struct for FilterPredicate
type FilterPredicate struct {
	Field *string `json:"field,omitempty"`
	Operator *string `json:"operator,omitempty"`
	Parameter *string `json:"parameter,omitempty"`
}

// NewFilterPredicate instantiates a new FilterPredicate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterPredicate() *FilterPredicate {
	this := FilterPredicate{}
	return &this
}

// NewFilterPredicateWithDefaults instantiates a new FilterPredicate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterPredicateWithDefaults() *FilterPredicate {
	this := FilterPredicate{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *FilterPredicate) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterPredicate) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *FilterPredicate) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *FilterPredicate) SetField(v string) {
	o.Field = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *FilterPredicate) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterPredicate) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *FilterPredicate) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *FilterPredicate) SetOperator(v string) {
	o.Operator = &v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *FilterPredicate) GetParameter() string {
	if o == nil || IsNil(o.Parameter) {
		var ret string
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterPredicate) GetParameterOk() (*string, bool) {
	if o == nil || IsNil(o.Parameter) {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *FilterPredicate) HasParameter() bool {
	if o != nil && !IsNil(o.Parameter) {
		return true
	}

	return false
}

// SetParameter gets a reference to the given string and assigns it to the Parameter field.
func (o *FilterPredicate) SetParameter(v string) {
	o.Parameter = &v
}

func (o FilterPredicate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilterPredicate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Parameter) {
		toSerialize["parameter"] = o.Parameter
	}
	return toSerialize, nil
}

type NullableFilterPredicate struct {
	value *FilterPredicate
	isSet bool
}

func (v NullableFilterPredicate) Get() *FilterPredicate {
	return v.value
}

func (v *NullableFilterPredicate) Set(val *FilterPredicate) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterPredicate) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterPredicate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterPredicate(val *FilterPredicate) *NullableFilterPredicate {
	return &NullableFilterPredicate{value: val, isSet: true}
}

func (v NullableFilterPredicate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterPredicate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


