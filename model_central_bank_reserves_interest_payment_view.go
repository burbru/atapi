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

// checks if the CentralBankReservesInterestPaymentView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CentralBankReservesInterestPaymentView{}

// CentralBankReservesInterestPaymentView struct for CentralBankReservesInterestPaymentView
type CentralBankReservesInterestPaymentView struct {
	Id *string `json:"id,omitempty"`
	NextPaymentDate *int64 `json:"nextPaymentDate,omitempty"`
	PaidInterest *float32 `json:"paidInterest,omitempty"`
	PaymentDate *int64 `json:"paymentDate,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewCentralBankReservesInterestPaymentView instantiates a new CentralBankReservesInterestPaymentView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCentralBankReservesInterestPaymentView() *CentralBankReservesInterestPaymentView {
	this := CentralBankReservesInterestPaymentView{}
	return &this
}

// NewCentralBankReservesInterestPaymentViewWithDefaults instantiates a new CentralBankReservesInterestPaymentView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCentralBankReservesInterestPaymentViewWithDefaults() *CentralBankReservesInterestPaymentView {
	this := CentralBankReservesInterestPaymentView{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CentralBankReservesInterestPaymentView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CentralBankReservesInterestPaymentView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CentralBankReservesInterestPaymentView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CentralBankReservesInterestPaymentView) SetId(v string) {
	o.Id = &v
}

// GetNextPaymentDate returns the NextPaymentDate field value if set, zero value otherwise.
func (o *CentralBankReservesInterestPaymentView) GetNextPaymentDate() int64 {
	if o == nil || IsNil(o.NextPaymentDate) {
		var ret int64
		return ret
	}
	return *o.NextPaymentDate
}

// GetNextPaymentDateOk returns a tuple with the NextPaymentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CentralBankReservesInterestPaymentView) GetNextPaymentDateOk() (*int64, bool) {
	if o == nil || IsNil(o.NextPaymentDate) {
		return nil, false
	}
	return o.NextPaymentDate, true
}

// HasNextPaymentDate returns a boolean if a field has been set.
func (o *CentralBankReservesInterestPaymentView) HasNextPaymentDate() bool {
	if o != nil && !IsNil(o.NextPaymentDate) {
		return true
	}

	return false
}

// SetNextPaymentDate gets a reference to the given int64 and assigns it to the NextPaymentDate field.
func (o *CentralBankReservesInterestPaymentView) SetNextPaymentDate(v int64) {
	o.NextPaymentDate = &v
}

// GetPaidInterest returns the PaidInterest field value if set, zero value otherwise.
func (o *CentralBankReservesInterestPaymentView) GetPaidInterest() float32 {
	if o == nil || IsNil(o.PaidInterest) {
		var ret float32
		return ret
	}
	return *o.PaidInterest
}

// GetPaidInterestOk returns a tuple with the PaidInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CentralBankReservesInterestPaymentView) GetPaidInterestOk() (*float32, bool) {
	if o == nil || IsNil(o.PaidInterest) {
		return nil, false
	}
	return o.PaidInterest, true
}

// HasPaidInterest returns a boolean if a field has been set.
func (o *CentralBankReservesInterestPaymentView) HasPaidInterest() bool {
	if o != nil && !IsNil(o.PaidInterest) {
		return true
	}

	return false
}

// SetPaidInterest gets a reference to the given float32 and assigns it to the PaidInterest field.
func (o *CentralBankReservesInterestPaymentView) SetPaidInterest(v float32) {
	o.PaidInterest = &v
}

// GetPaymentDate returns the PaymentDate field value if set, zero value otherwise.
func (o *CentralBankReservesInterestPaymentView) GetPaymentDate() int64 {
	if o == nil || IsNil(o.PaymentDate) {
		var ret int64
		return ret
	}
	return *o.PaymentDate
}

// GetPaymentDateOk returns a tuple with the PaymentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CentralBankReservesInterestPaymentView) GetPaymentDateOk() (*int64, bool) {
	if o == nil || IsNil(o.PaymentDate) {
		return nil, false
	}
	return o.PaymentDate, true
}

// HasPaymentDate returns a boolean if a field has been set.
func (o *CentralBankReservesInterestPaymentView) HasPaymentDate() bool {
	if o != nil && !IsNil(o.PaymentDate) {
		return true
	}

	return false
}

// SetPaymentDate gets a reference to the given int64 and assigns it to the PaymentDate field.
func (o *CentralBankReservesInterestPaymentView) SetPaymentDate(v int64) {
	o.PaymentDate = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CentralBankReservesInterestPaymentView) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CentralBankReservesInterestPaymentView) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CentralBankReservesInterestPaymentView) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *CentralBankReservesInterestPaymentView) SetVersion(v int64) {
	o.Version = &v
}

func (o CentralBankReservesInterestPaymentView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CentralBankReservesInterestPaymentView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.NextPaymentDate) {
		toSerialize["nextPaymentDate"] = o.NextPaymentDate
	}
	if !IsNil(o.PaidInterest) {
		toSerialize["paidInterest"] = o.PaidInterest
	}
	if !IsNil(o.PaymentDate) {
		toSerialize["paymentDate"] = o.PaymentDate
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableCentralBankReservesInterestPaymentView struct {
	value *CentralBankReservesInterestPaymentView
	isSet bool
}

func (v NullableCentralBankReservesInterestPaymentView) Get() *CentralBankReservesInterestPaymentView {
	return v.value
}

func (v *NullableCentralBankReservesInterestPaymentView) Set(val *CentralBankReservesInterestPaymentView) {
	v.value = val
	v.isSet = true
}

func (v NullableCentralBankReservesInterestPaymentView) IsSet() bool {
	return v.isSet
}

func (v *NullableCentralBankReservesInterestPaymentView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCentralBankReservesInterestPaymentView(val *CentralBankReservesInterestPaymentView) *NullableCentralBankReservesInterestPaymentView {
	return &NullableCentralBankReservesInterestPaymentView{value: val, isSet: true}
}

func (v NullableCentralBankReservesInterestPaymentView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCentralBankReservesInterestPaymentView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


