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

// checks if the RangeViewbigdecimal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RangeViewbigdecimal{}

// RangeViewbigdecimal struct for RangeViewbigdecimal
type RangeViewbigdecimal struct {
	LowerBound *float32 `json:"lowerBound,omitempty"`
	UpperBound *float32 `json:"upperBound,omitempty"`
}

// NewRangeViewbigdecimal instantiates a new RangeViewbigdecimal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRangeViewbigdecimal() *RangeViewbigdecimal {
	this := RangeViewbigdecimal{}
	return &this
}

// NewRangeViewbigdecimalWithDefaults instantiates a new RangeViewbigdecimal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeViewbigdecimalWithDefaults() *RangeViewbigdecimal {
	this := RangeViewbigdecimal{}
	return &this
}

// GetLowerBound returns the LowerBound field value if set, zero value otherwise.
func (o *RangeViewbigdecimal) GetLowerBound() float32 {
	if o == nil || IsNil(o.LowerBound) {
		var ret float32
		return ret
	}
	return *o.LowerBound
}

// GetLowerBoundOk returns a tuple with the LowerBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeViewbigdecimal) GetLowerBoundOk() (*float32, bool) {
	if o == nil || IsNil(o.LowerBound) {
		return nil, false
	}
	return o.LowerBound, true
}

// HasLowerBound returns a boolean if a field has been set.
func (o *RangeViewbigdecimal) HasLowerBound() bool {
	if o != nil && !IsNil(o.LowerBound) {
		return true
	}

	return false
}

// SetLowerBound gets a reference to the given float32 and assigns it to the LowerBound field.
func (o *RangeViewbigdecimal) SetLowerBound(v float32) {
	o.LowerBound = &v
}

// GetUpperBound returns the UpperBound field value if set, zero value otherwise.
func (o *RangeViewbigdecimal) GetUpperBound() float32 {
	if o == nil || IsNil(o.UpperBound) {
		var ret float32
		return ret
	}
	return *o.UpperBound
}

// GetUpperBoundOk returns a tuple with the UpperBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeViewbigdecimal) GetUpperBoundOk() (*float32, bool) {
	if o == nil || IsNil(o.UpperBound) {
		return nil, false
	}
	return o.UpperBound, true
}

// HasUpperBound returns a boolean if a field has been set.
func (o *RangeViewbigdecimal) HasUpperBound() bool {
	if o != nil && !IsNil(o.UpperBound) {
		return true
	}

	return false
}

// SetUpperBound gets a reference to the given float32 and assigns it to the UpperBound field.
func (o *RangeViewbigdecimal) SetUpperBound(v float32) {
	o.UpperBound = &v
}

func (o RangeViewbigdecimal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RangeViewbigdecimal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LowerBound) {
		toSerialize["lowerBound"] = o.LowerBound
	}
	if !IsNil(o.UpperBound) {
		toSerialize["upperBound"] = o.UpperBound
	}
	return toSerialize, nil
}

type NullableRangeViewbigdecimal struct {
	value *RangeViewbigdecimal
	isSet bool
}

func (v NullableRangeViewbigdecimal) Get() *RangeViewbigdecimal {
	return v.value
}

func (v *NullableRangeViewbigdecimal) Set(val *RangeViewbigdecimal) {
	v.value = val
	v.isSet = true
}

func (v NullableRangeViewbigdecimal) IsSet() bool {
	return v.isSet
}

func (v *NullableRangeViewbigdecimal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRangeViewbigdecimal(val *RangeViewbigdecimal) *NullableRangeViewbigdecimal {
	return &NullableRangeViewbigdecimal{value: val, isSet: true}
}

func (v NullableRangeViewbigdecimal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRangeViewbigdecimal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


