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

// checks if the HelpComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelpComment{}

// HelpComment struct for HelpComment
type HelpComment struct {
	Text *string `json:"text,omitempty"`
}

// NewHelpComment instantiates a new HelpComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelpComment() *HelpComment {
	this := HelpComment{}
	return &this
}

// NewHelpCommentWithDefaults instantiates a new HelpComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelpCommentWithDefaults() *HelpComment {
	this := HelpComment{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *HelpComment) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelpComment) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *HelpComment) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *HelpComment) SetText(v string) {
	o.Text = &v
}

func (o HelpComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelpComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableHelpComment struct {
	value *HelpComment
	isSet bool
}

func (v NullableHelpComment) Get() *HelpComment {
	return v.value
}

func (v *NullableHelpComment) Set(val *HelpComment) {
	v.value = val
	v.isSet = true
}

func (v NullableHelpComment) IsSet() bool {
	return v.isSet
}

func (v *NullableHelpComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelpComment(val *HelpComment) *NullableHelpComment {
	return &NullableHelpComment{value: val, isSet: true}
}

func (v NullableHelpComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelpComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


