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

// checks if the HashTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HashTag{}

// HashTag struct for HashTag
type HashTag struct {
	Id *string `json:"id,omitempty"`
	Tag *string `json:"tag,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewHashTag instantiates a new HashTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHashTag() *HashTag {
	this := HashTag{}
	return &this
}

// NewHashTagWithDefaults instantiates a new HashTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHashTagWithDefaults() *HashTag {
	this := HashTag{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HashTag) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashTag) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HashTag) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HashTag) SetId(v string) {
	o.Id = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *HashTag) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashTag) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *HashTag) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *HashTag) SetTag(v string) {
	o.Tag = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HashTag) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashTag) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HashTag) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *HashTag) SetVersion(v int64) {
	o.Version = &v
}

func (o HashTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HashTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableHashTag struct {
	value *HashTag
	isSet bool
}

func (v NullableHashTag) Get() *HashTag {
	return v.value
}

func (v *NullableHashTag) Set(val *HashTag) {
	v.value = val
	v.isSet = true
}

func (v NullableHashTag) IsSet() bool {
	return v.isSet
}

func (v *NullableHashTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHashTag(val *HashTag) *NullableHashTag {
	return &NullableHashTag{value: val, isSet: true}
}

func (v NullableHashTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHashTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


