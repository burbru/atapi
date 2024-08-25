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

// checks if the ListingShareView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListingShareView{}

// ListingShareView struct for ListingShareView
type ListingShareView struct {
	Listing *ListingView `json:"listing,omitempty"`
	ShareInPercent *float32 `json:"shareInPercent,omitempty"`
}

// NewListingShareView instantiates a new ListingShareView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListingShareView() *ListingShareView {
	this := ListingShareView{}
	return &this
}

// NewListingShareViewWithDefaults instantiates a new ListingShareView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListingShareViewWithDefaults() *ListingShareView {
	this := ListingShareView{}
	return &this
}

// GetListing returns the Listing field value if set, zero value otherwise.
func (o *ListingShareView) GetListing() ListingView {
	if o == nil || IsNil(o.Listing) {
		var ret ListingView
		return ret
	}
	return *o.Listing
}

// GetListingOk returns a tuple with the Listing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingShareView) GetListingOk() (*ListingView, bool) {
	if o == nil || IsNil(o.Listing) {
		return nil, false
	}
	return o.Listing, true
}

// HasListing returns a boolean if a field has been set.
func (o *ListingShareView) HasListing() bool {
	if o != nil && !IsNil(o.Listing) {
		return true
	}

	return false
}

// SetListing gets a reference to the given ListingView and assigns it to the Listing field.
func (o *ListingShareView) SetListing(v ListingView) {
	o.Listing = &v
}

// GetShareInPercent returns the ShareInPercent field value if set, zero value otherwise.
func (o *ListingShareView) GetShareInPercent() float32 {
	if o == nil || IsNil(o.ShareInPercent) {
		var ret float32
		return ret
	}
	return *o.ShareInPercent
}

// GetShareInPercentOk returns a tuple with the ShareInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingShareView) GetShareInPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.ShareInPercent) {
		return nil, false
	}
	return o.ShareInPercent, true
}

// HasShareInPercent returns a boolean if a field has been set.
func (o *ListingShareView) HasShareInPercent() bool {
	if o != nil && !IsNil(o.ShareInPercent) {
		return true
	}

	return false
}

// SetShareInPercent gets a reference to the given float32 and assigns it to the ShareInPercent field.
func (o *ListingShareView) SetShareInPercent(v float32) {
	o.ShareInPercent = &v
}

func (o ListingShareView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListingShareView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Listing) {
		toSerialize["listing"] = o.Listing
	}
	if !IsNil(o.ShareInPercent) {
		toSerialize["shareInPercent"] = o.ShareInPercent
	}
	return toSerialize, nil
}

type NullableListingShareView struct {
	value *ListingShareView
	isSet bool
}

func (v NullableListingShareView) Get() *ListingShareView {
	return v.value
}

func (v *NullableListingShareView) Set(val *ListingShareView) {
	v.value = val
	v.isSet = true
}

func (v NullableListingShareView) IsSet() bool {
	return v.isSet
}

func (v *NullableListingShareView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListingShareView(val *ListingShareView) *NullableListingShareView {
	return &NullableListingShareView{value: val, isSet: true}
}

func (v NullableListingShareView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListingShareView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


