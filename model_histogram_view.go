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

// checks if the HistogramView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistogramView{}

// HistogramView struct for HistogramView
type HistogramView struct {
	HighlightRange *RangeViewbigdecimal `json:"highlightRange,omitempty"`
	HighlightValue *float32 `json:"highlightValue,omitempty"`
	Histogram *map[string]int32 `json:"histogram,omitempty"`
}

// NewHistogramView instantiates a new HistogramView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistogramView() *HistogramView {
	this := HistogramView{}
	return &this
}

// NewHistogramViewWithDefaults instantiates a new HistogramView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistogramViewWithDefaults() *HistogramView {
	this := HistogramView{}
	return &this
}

// GetHighlightRange returns the HighlightRange field value if set, zero value otherwise.
func (o *HistogramView) GetHighlightRange() RangeViewbigdecimal {
	if o == nil || IsNil(o.HighlightRange) {
		var ret RangeViewbigdecimal
		return ret
	}
	return *o.HighlightRange
}

// GetHighlightRangeOk returns a tuple with the HighlightRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistogramView) GetHighlightRangeOk() (*RangeViewbigdecimal, bool) {
	if o == nil || IsNil(o.HighlightRange) {
		return nil, false
	}
	return o.HighlightRange, true
}

// HasHighlightRange returns a boolean if a field has been set.
func (o *HistogramView) HasHighlightRange() bool {
	if o != nil && !IsNil(o.HighlightRange) {
		return true
	}

	return false
}

// SetHighlightRange gets a reference to the given RangeViewbigdecimal and assigns it to the HighlightRange field.
func (o *HistogramView) SetHighlightRange(v RangeViewbigdecimal) {
	o.HighlightRange = &v
}

// GetHighlightValue returns the HighlightValue field value if set, zero value otherwise.
func (o *HistogramView) GetHighlightValue() float32 {
	if o == nil || IsNil(o.HighlightValue) {
		var ret float32
		return ret
	}
	return *o.HighlightValue
}

// GetHighlightValueOk returns a tuple with the HighlightValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistogramView) GetHighlightValueOk() (*float32, bool) {
	if o == nil || IsNil(o.HighlightValue) {
		return nil, false
	}
	return o.HighlightValue, true
}

// HasHighlightValue returns a boolean if a field has been set.
func (o *HistogramView) HasHighlightValue() bool {
	if o != nil && !IsNil(o.HighlightValue) {
		return true
	}

	return false
}

// SetHighlightValue gets a reference to the given float32 and assigns it to the HighlightValue field.
func (o *HistogramView) SetHighlightValue(v float32) {
	o.HighlightValue = &v
}

// GetHistogram returns the Histogram field value if set, zero value otherwise.
func (o *HistogramView) GetHistogram() map[string]int32 {
	if o == nil || IsNil(o.Histogram) {
		var ret map[string]int32
		return ret
	}
	return *o.Histogram
}

// GetHistogramOk returns a tuple with the Histogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistogramView) GetHistogramOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.Histogram) {
		return nil, false
	}
	return o.Histogram, true
}

// HasHistogram returns a boolean if a field has been set.
func (o *HistogramView) HasHistogram() bool {
	if o != nil && !IsNil(o.Histogram) {
		return true
	}

	return false
}

// SetHistogram gets a reference to the given map[string]int32 and assigns it to the Histogram field.
func (o *HistogramView) SetHistogram(v map[string]int32) {
	o.Histogram = &v
}

func (o HistogramView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistogramView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HighlightRange) {
		toSerialize["highlightRange"] = o.HighlightRange
	}
	if !IsNil(o.HighlightValue) {
		toSerialize["highlightValue"] = o.HighlightValue
	}
	if !IsNil(o.Histogram) {
		toSerialize["histogram"] = o.Histogram
	}
	return toSerialize, nil
}

type NullableHistogramView struct {
	value *HistogramView
	isSet bool
}

func (v NullableHistogramView) Get() *HistogramView {
	return v.value
}

func (v *NullableHistogramView) Set(val *HistogramView) {
	v.value = val
	v.isSet = true
}

func (v NullableHistogramView) IsSet() bool {
	return v.isSet
}

func (v *NullableHistogramView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistogramView(val *HistogramView) *NullableHistogramView {
	return &NullableHistogramView{value: val, isSet: true}
}

func (v NullableHistogramView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistogramView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


