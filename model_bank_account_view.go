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

// checks if the BankAccountView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BankAccountView{}

// BankAccountView struct for BankAccountView
type BankAccountView struct {
	Cash *float32 `json:"cash,omitempty"`
	Id *string `json:"id,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewBankAccountView instantiates a new BankAccountView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankAccountView() *BankAccountView {
	this := BankAccountView{}
	return &this
}

// NewBankAccountViewWithDefaults instantiates a new BankAccountView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankAccountViewWithDefaults() *BankAccountView {
	this := BankAccountView{}
	return &this
}

// GetCash returns the Cash field value if set, zero value otherwise.
func (o *BankAccountView) GetCash() float32 {
	if o == nil || IsNil(o.Cash) {
		var ret float32
		return ret
	}
	return *o.Cash
}

// GetCashOk returns a tuple with the Cash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountView) GetCashOk() (*float32, bool) {
	if o == nil || IsNil(o.Cash) {
		return nil, false
	}
	return o.Cash, true
}

// HasCash returns a boolean if a field has been set.
func (o *BankAccountView) HasCash() bool {
	if o != nil && !IsNil(o.Cash) {
		return true
	}

	return false
}

// SetCash gets a reference to the given float32 and assigns it to the Cash field.
func (o *BankAccountView) SetCash(v float32) {
	o.Cash = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BankAccountView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BankAccountView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BankAccountView) SetId(v string) {
	o.Id = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BankAccountView) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountView) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BankAccountView) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *BankAccountView) SetVersion(v int64) {
	o.Version = &v
}

func (o BankAccountView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankAccountView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cash) {
		toSerialize["cash"] = o.Cash
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableBankAccountView struct {
	value *BankAccountView
	isSet bool
}

func (v NullableBankAccountView) Get() *BankAccountView {
	return v.value
}

func (v *NullableBankAccountView) Set(val *BankAccountView) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccountView) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccountView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccountView(val *BankAccountView) *NullableBankAccountView {
	return &NullableBankAccountView{value: val, isSet: true}
}

func (v NullableBankAccountView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankAccountView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


