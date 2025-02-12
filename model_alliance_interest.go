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

// checks if the AllianceInterest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllianceInterest{}

// AllianceInterest struct for AllianceInterest
type AllianceInterest struct {
	AllianceId *string `json:"allianceId,omitempty"`
	Id *string `json:"id,omitempty"`
	Interest *int64 `json:"interest,omitempty"`
	UserId *string `json:"userId,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewAllianceInterest instantiates a new AllianceInterest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllianceInterest() *AllianceInterest {
	this := AllianceInterest{}
	return &this
}

// NewAllianceInterestWithDefaults instantiates a new AllianceInterest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllianceInterestWithDefaults() *AllianceInterest {
	this := AllianceInterest{}
	return &this
}

// GetAllianceId returns the AllianceId field value if set, zero value otherwise.
func (o *AllianceInterest) GetAllianceId() string {
	if o == nil || IsNil(o.AllianceId) {
		var ret string
		return ret
	}
	return *o.AllianceId
}

// GetAllianceIdOk returns a tuple with the AllianceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllianceInterest) GetAllianceIdOk() (*string, bool) {
	if o == nil || IsNil(o.AllianceId) {
		return nil, false
	}
	return o.AllianceId, true
}

// HasAllianceId returns a boolean if a field has been set.
func (o *AllianceInterest) HasAllianceId() bool {
	if o != nil && !IsNil(o.AllianceId) {
		return true
	}

	return false
}

// SetAllianceId gets a reference to the given string and assigns it to the AllianceId field.
func (o *AllianceInterest) SetAllianceId(v string) {
	o.AllianceId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AllianceInterest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllianceInterest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AllianceInterest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AllianceInterest) SetId(v string) {
	o.Id = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *AllianceInterest) GetInterest() int64 {
	if o == nil || IsNil(o.Interest) {
		var ret int64
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllianceInterest) GetInterestOk() (*int64, bool) {
	if o == nil || IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *AllianceInterest) HasInterest() bool {
	if o != nil && !IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given int64 and assigns it to the Interest field.
func (o *AllianceInterest) SetInterest(v int64) {
	o.Interest = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AllianceInterest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllianceInterest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AllianceInterest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AllianceInterest) SetUserId(v string) {
	o.UserId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AllianceInterest) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllianceInterest) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AllianceInterest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *AllianceInterest) SetVersion(v int64) {
	o.Version = &v
}

func (o AllianceInterest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllianceInterest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllianceId) {
		toSerialize["allianceId"] = o.AllianceId
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

type NullableAllianceInterest struct {
	value *AllianceInterest
	isSet bool
}

func (v NullableAllianceInterest) Get() *AllianceInterest {
	return v.value
}

func (v *NullableAllianceInterest) Set(val *AllianceInterest) {
	v.value = val
	v.isSet = true
}

func (v NullableAllianceInterest) IsSet() bool {
	return v.isSet
}

func (v *NullableAllianceInterest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllianceInterest(val *AllianceInterest) *NullableAllianceInterest {
	return &NullableAllianceInterest{value: val, isSet: true}
}

func (v NullableAllianceInterest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllianceInterest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


