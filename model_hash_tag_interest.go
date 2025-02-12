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

// checks if the HashTagInterest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HashTagInterest{}

// HashTagInterest struct for HashTagInterest
type HashTagInterest struct {
	HashTag *HashTag `json:"hashTag,omitempty"`
	Id *string `json:"id,omitempty"`
	Interest *int64 `json:"interest,omitempty"`
	UserId *string `json:"userId,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewHashTagInterest instantiates a new HashTagInterest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHashTagInterest() *HashTagInterest {
	this := HashTagInterest{}
	return &this
}

// NewHashTagInterestWithDefaults instantiates a new HashTagInterest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHashTagInterestWithDefaults() *HashTagInterest {
	this := HashTagInterest{}
	return &this
}

// GetHashTag returns the HashTag field value if set, zero value otherwise.
func (o *HashTagInterest) GetHashTag() HashTag {
	if o == nil || IsNil(o.HashTag) {
		var ret HashTag
		return ret
	}
	return *o.HashTag
}

// GetHashTagOk returns a tuple with the HashTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashTagInterest) GetHashTagOk() (*HashTag, bool) {
	if o == nil || IsNil(o.HashTag) {
		return nil, false
	}
	return o.HashTag, true
}

// HasHashTag returns a boolean if a field has been set.
func (o *HashTagInterest) HasHashTag() bool {
	if o != nil && !IsNil(o.HashTag) {
		return true
	}

	return false
}

// SetHashTag gets a reference to the given HashTag and assigns it to the HashTag field.
func (o *HashTagInterest) SetHashTag(v HashTag) {
	o.HashTag = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HashTagInterest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashTagInterest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HashTagInterest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HashTagInterest) SetId(v string) {
	o.Id = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *HashTagInterest) GetInterest() int64 {
	if o == nil || IsNil(o.Interest) {
		var ret int64
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashTagInterest) GetInterestOk() (*int64, bool) {
	if o == nil || IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *HashTagInterest) HasInterest() bool {
	if o != nil && !IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given int64 and assigns it to the Interest field.
func (o *HashTagInterest) SetInterest(v int64) {
	o.Interest = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *HashTagInterest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashTagInterest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *HashTagInterest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *HashTagInterest) SetUserId(v string) {
	o.UserId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HashTagInterest) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashTagInterest) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HashTagInterest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *HashTagInterest) SetVersion(v int64) {
	o.Version = &v
}

func (o HashTagInterest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HashTagInterest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HashTag) {
		toSerialize["hashTag"] = o.HashTag
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableHashTagInterest struct {
	value *HashTagInterest
	isSet bool
}

func (v NullableHashTagInterest) Get() *HashTagInterest {
	return v.value
}

func (v *NullableHashTagInterest) Set(val *HashTagInterest) {
	v.value = val
	v.isSet = true
}

func (v NullableHashTagInterest) IsSet() bool {
	return v.isSet
}

func (v *NullableHashTagInterest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHashTagInterest(val *HashTagInterest) *NullableHashTagInterest {
	return &NullableHashTagInterest{value: val, isSet: true}
}

func (v NullableHashTagInterest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHashTagInterest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


