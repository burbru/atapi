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

// checks if the DividendPaymentView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DividendPaymentView{}

// DividendPaymentView struct for DividendPaymentView
type DividendPaymentView struct {
	Company *CompactCompanyView `json:"company,omitempty"`
	Id *string `json:"id,omitempty"`
	MaximalCashVolume *float32 `json:"maximalCashVolume,omitempty"`
	StartDate *int64 `json:"startDate,omitempty"`
}

// NewDividendPaymentView instantiates a new DividendPaymentView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDividendPaymentView() *DividendPaymentView {
	this := DividendPaymentView{}
	return &this
}

// NewDividendPaymentViewWithDefaults instantiates a new DividendPaymentView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDividendPaymentViewWithDefaults() *DividendPaymentView {
	this := DividendPaymentView{}
	return &this
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *DividendPaymentView) GetCompany() CompactCompanyView {
	if o == nil || IsNil(o.Company) {
		var ret CompactCompanyView
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DividendPaymentView) GetCompanyOk() (*CompactCompanyView, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *DividendPaymentView) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given CompactCompanyView and assigns it to the Company field.
func (o *DividendPaymentView) SetCompany(v CompactCompanyView) {
	o.Company = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DividendPaymentView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DividendPaymentView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DividendPaymentView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DividendPaymentView) SetId(v string) {
	o.Id = &v
}

// GetMaximalCashVolume returns the MaximalCashVolume field value if set, zero value otherwise.
func (o *DividendPaymentView) GetMaximalCashVolume() float32 {
	if o == nil || IsNil(o.MaximalCashVolume) {
		var ret float32
		return ret
	}
	return *o.MaximalCashVolume
}

// GetMaximalCashVolumeOk returns a tuple with the MaximalCashVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DividendPaymentView) GetMaximalCashVolumeOk() (*float32, bool) {
	if o == nil || IsNil(o.MaximalCashVolume) {
		return nil, false
	}
	return o.MaximalCashVolume, true
}

// HasMaximalCashVolume returns a boolean if a field has been set.
func (o *DividendPaymentView) HasMaximalCashVolume() bool {
	if o != nil && !IsNil(o.MaximalCashVolume) {
		return true
	}

	return false
}

// SetMaximalCashVolume gets a reference to the given float32 and assigns it to the MaximalCashVolume field.
func (o *DividendPaymentView) SetMaximalCashVolume(v float32) {
	o.MaximalCashVolume = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *DividendPaymentView) GetStartDate() int64 {
	if o == nil || IsNil(o.StartDate) {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DividendPaymentView) GetStartDateOk() (*int64, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *DividendPaymentView) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *DividendPaymentView) SetStartDate(v int64) {
	o.StartDate = &v
}

func (o DividendPaymentView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DividendPaymentView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MaximalCashVolume) {
		toSerialize["maximalCashVolume"] = o.MaximalCashVolume
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	return toSerialize, nil
}

type NullableDividendPaymentView struct {
	value *DividendPaymentView
	isSet bool
}

func (v NullableDividendPaymentView) Get() *DividendPaymentView {
	return v.value
}

func (v *NullableDividendPaymentView) Set(val *DividendPaymentView) {
	v.value = val
	v.isSet = true
}

func (v NullableDividendPaymentView) IsSet() bool {
	return v.isSet
}

func (v *NullableDividendPaymentView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDividendPaymentView(val *DividendPaymentView) *NullableDividendPaymentView {
	return &NullableDividendPaymentView{value: val, isSet: true}
}

func (v NullableDividendPaymentView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDividendPaymentView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


