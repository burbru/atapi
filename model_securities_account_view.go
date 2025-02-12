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

// checks if the SecuritiesAccountView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecuritiesAccountView{}

// SecuritiesAccountView struct for SecuritiesAccountView
type SecuritiesAccountView struct {
	ClearingAccountId *string `json:"clearingAccountId,omitempty"`
	Id *string `json:"id,omitempty"`
	PrivateAccount *bool `json:"privateAccount,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewSecuritiesAccountView instantiates a new SecuritiesAccountView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecuritiesAccountView() *SecuritiesAccountView {
	this := SecuritiesAccountView{}
	return &this
}

// NewSecuritiesAccountViewWithDefaults instantiates a new SecuritiesAccountView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecuritiesAccountViewWithDefaults() *SecuritiesAccountView {
	this := SecuritiesAccountView{}
	return &this
}

// GetClearingAccountId returns the ClearingAccountId field value if set, zero value otherwise.
func (o *SecuritiesAccountView) GetClearingAccountId() string {
	if o == nil || IsNil(o.ClearingAccountId) {
		var ret string
		return ret
	}
	return *o.ClearingAccountId
}

// GetClearingAccountIdOk returns a tuple with the ClearingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritiesAccountView) GetClearingAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClearingAccountId) {
		return nil, false
	}
	return o.ClearingAccountId, true
}

// HasClearingAccountId returns a boolean if a field has been set.
func (o *SecuritiesAccountView) HasClearingAccountId() bool {
	if o != nil && !IsNil(o.ClearingAccountId) {
		return true
	}

	return false
}

// SetClearingAccountId gets a reference to the given string and assigns it to the ClearingAccountId field.
func (o *SecuritiesAccountView) SetClearingAccountId(v string) {
	o.ClearingAccountId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecuritiesAccountView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritiesAccountView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SecuritiesAccountView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SecuritiesAccountView) SetId(v string) {
	o.Id = &v
}

// GetPrivateAccount returns the PrivateAccount field value if set, zero value otherwise.
func (o *SecuritiesAccountView) GetPrivateAccount() bool {
	if o == nil || IsNil(o.PrivateAccount) {
		var ret bool
		return ret
	}
	return *o.PrivateAccount
}

// GetPrivateAccountOk returns a tuple with the PrivateAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritiesAccountView) GetPrivateAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.PrivateAccount) {
		return nil, false
	}
	return o.PrivateAccount, true
}

// HasPrivateAccount returns a boolean if a field has been set.
func (o *SecuritiesAccountView) HasPrivateAccount() bool {
	if o != nil && !IsNil(o.PrivateAccount) {
		return true
	}

	return false
}

// SetPrivateAccount gets a reference to the given bool and assigns it to the PrivateAccount field.
func (o *SecuritiesAccountView) SetPrivateAccount(v bool) {
	o.PrivateAccount = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SecuritiesAccountView) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritiesAccountView) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SecuritiesAccountView) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *SecuritiesAccountView) SetVersion(v int64) {
	o.Version = &v
}

func (o SecuritiesAccountView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecuritiesAccountView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClearingAccountId) {
		toSerialize["clearingAccountId"] = o.ClearingAccountId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PrivateAccount) {
		toSerialize["privateAccount"] = o.PrivateAccount
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableSecuritiesAccountView struct {
	value *SecuritiesAccountView
	isSet bool
}

func (v NullableSecuritiesAccountView) Get() *SecuritiesAccountView {
	return v.value
}

func (v *NullableSecuritiesAccountView) Set(val *SecuritiesAccountView) {
	v.value = val
	v.isSet = true
}

func (v NullableSecuritiesAccountView) IsSet() bool {
	return v.isSet
}

func (v *NullableSecuritiesAccountView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecuritiesAccountView(val *SecuritiesAccountView) *NullableSecuritiesAccountView {
	return &NullableSecuritiesAccountView{value: val, isSet: true}
}

func (v NullableSecuritiesAccountView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecuritiesAccountView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


