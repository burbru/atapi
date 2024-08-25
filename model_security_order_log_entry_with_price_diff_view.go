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

// checks if the SecurityOrderLogEntryWithPriceDiffView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityOrderLogEntryWithPriceDiffView{}

// SecurityOrderLogEntryWithPriceDiffView struct for SecurityOrderLogEntryWithPriceDiffView
type SecurityOrderLogEntryWithPriceDiffView struct {
	Listing *ListingView `json:"listing,omitempty"`
	Log *SecurityOrderLogEntryView `json:"log,omitempty"`
	PriceChangeAbs *float32 `json:"priceChangeAbs,omitempty"`
	PriceChangeRelInPercent *float32 `json:"priceChangeRelInPercent,omitempty"`
}

// NewSecurityOrderLogEntryWithPriceDiffView instantiates a new SecurityOrderLogEntryWithPriceDiffView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityOrderLogEntryWithPriceDiffView() *SecurityOrderLogEntryWithPriceDiffView {
	this := SecurityOrderLogEntryWithPriceDiffView{}
	return &this
}

// NewSecurityOrderLogEntryWithPriceDiffViewWithDefaults instantiates a new SecurityOrderLogEntryWithPriceDiffView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityOrderLogEntryWithPriceDiffViewWithDefaults() *SecurityOrderLogEntryWithPriceDiffView {
	this := SecurityOrderLogEntryWithPriceDiffView{}
	return &this
}

// GetListing returns the Listing field value if set, zero value otherwise.
func (o *SecurityOrderLogEntryWithPriceDiffView) GetListing() ListingView {
	if o == nil || IsNil(o.Listing) {
		var ret ListingView
		return ret
	}
	return *o.Listing
}

// GetListingOk returns a tuple with the Listing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityOrderLogEntryWithPriceDiffView) GetListingOk() (*ListingView, bool) {
	if o == nil || IsNil(o.Listing) {
		return nil, false
	}
	return o.Listing, true
}

// HasListing returns a boolean if a field has been set.
func (o *SecurityOrderLogEntryWithPriceDiffView) HasListing() bool {
	if o != nil && !IsNil(o.Listing) {
		return true
	}

	return false
}

// SetListing gets a reference to the given ListingView and assigns it to the Listing field.
func (o *SecurityOrderLogEntryWithPriceDiffView) SetListing(v ListingView) {
	o.Listing = &v
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *SecurityOrderLogEntryWithPriceDiffView) GetLog() SecurityOrderLogEntryView {
	if o == nil || IsNil(o.Log) {
		var ret SecurityOrderLogEntryView
		return ret
	}
	return *o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityOrderLogEntryWithPriceDiffView) GetLogOk() (*SecurityOrderLogEntryView, bool) {
	if o == nil || IsNil(o.Log) {
		return nil, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *SecurityOrderLogEntryWithPriceDiffView) HasLog() bool {
	if o != nil && !IsNil(o.Log) {
		return true
	}

	return false
}

// SetLog gets a reference to the given SecurityOrderLogEntryView and assigns it to the Log field.
func (o *SecurityOrderLogEntryWithPriceDiffView) SetLog(v SecurityOrderLogEntryView) {
	o.Log = &v
}

// GetPriceChangeAbs returns the PriceChangeAbs field value if set, zero value otherwise.
func (o *SecurityOrderLogEntryWithPriceDiffView) GetPriceChangeAbs() float32 {
	if o == nil || IsNil(o.PriceChangeAbs) {
		var ret float32
		return ret
	}
	return *o.PriceChangeAbs
}

// GetPriceChangeAbsOk returns a tuple with the PriceChangeAbs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityOrderLogEntryWithPriceDiffView) GetPriceChangeAbsOk() (*float32, bool) {
	if o == nil || IsNil(o.PriceChangeAbs) {
		return nil, false
	}
	return o.PriceChangeAbs, true
}

// HasPriceChangeAbs returns a boolean if a field has been set.
func (o *SecurityOrderLogEntryWithPriceDiffView) HasPriceChangeAbs() bool {
	if o != nil && !IsNil(o.PriceChangeAbs) {
		return true
	}

	return false
}

// SetPriceChangeAbs gets a reference to the given float32 and assigns it to the PriceChangeAbs field.
func (o *SecurityOrderLogEntryWithPriceDiffView) SetPriceChangeAbs(v float32) {
	o.PriceChangeAbs = &v
}

// GetPriceChangeRelInPercent returns the PriceChangeRelInPercent field value if set, zero value otherwise.
func (o *SecurityOrderLogEntryWithPriceDiffView) GetPriceChangeRelInPercent() float32 {
	if o == nil || IsNil(o.PriceChangeRelInPercent) {
		var ret float32
		return ret
	}
	return *o.PriceChangeRelInPercent
}

// GetPriceChangeRelInPercentOk returns a tuple with the PriceChangeRelInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityOrderLogEntryWithPriceDiffView) GetPriceChangeRelInPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.PriceChangeRelInPercent) {
		return nil, false
	}
	return o.PriceChangeRelInPercent, true
}

// HasPriceChangeRelInPercent returns a boolean if a field has been set.
func (o *SecurityOrderLogEntryWithPriceDiffView) HasPriceChangeRelInPercent() bool {
	if o != nil && !IsNil(o.PriceChangeRelInPercent) {
		return true
	}

	return false
}

// SetPriceChangeRelInPercent gets a reference to the given float32 and assigns it to the PriceChangeRelInPercent field.
func (o *SecurityOrderLogEntryWithPriceDiffView) SetPriceChangeRelInPercent(v float32) {
	o.PriceChangeRelInPercent = &v
}

func (o SecurityOrderLogEntryWithPriceDiffView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityOrderLogEntryWithPriceDiffView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Listing) {
		toSerialize["listing"] = o.Listing
	}
	if !IsNil(o.Log) {
		toSerialize["log"] = o.Log
	}
	if !IsNil(o.PriceChangeAbs) {
		toSerialize["priceChangeAbs"] = o.PriceChangeAbs
	}
	if !IsNil(o.PriceChangeRelInPercent) {
		toSerialize["priceChangeRelInPercent"] = o.PriceChangeRelInPercent
	}
	return toSerialize, nil
}

type NullableSecurityOrderLogEntryWithPriceDiffView struct {
	value *SecurityOrderLogEntryWithPriceDiffView
	isSet bool
}

func (v NullableSecurityOrderLogEntryWithPriceDiffView) Get() *SecurityOrderLogEntryWithPriceDiffView {
	return v.value
}

func (v *NullableSecurityOrderLogEntryWithPriceDiffView) Set(val *SecurityOrderLogEntryWithPriceDiffView) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityOrderLogEntryWithPriceDiffView) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityOrderLogEntryWithPriceDiffView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityOrderLogEntryWithPriceDiffView(val *SecurityOrderLogEntryWithPriceDiffView) *NullableSecurityOrderLogEntryWithPriceDiffView {
	return &NullableSecurityOrderLogEntryWithPriceDiffView{value: val, isSet: true}
}

func (v NullableSecurityOrderLogEntryWithPriceDiffView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityOrderLogEntryWithPriceDiffView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


