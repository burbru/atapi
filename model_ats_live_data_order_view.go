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

// checks if the AtsLiveDataOrderView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AtsLiveDataOrderView{}

// AtsLiveDataOrderView struct for AtsLiveDataOrderView
type AtsLiveDataOrderView struct {
	Action *string `json:"action,omitempty"`
	Date *int64 `json:"date,omitempty"`
	Diff *float32 `json:"diff,omitempty"`
	Id *string `json:"id,omitempty"`
	Limit *float32 `json:"limit,omitempty"`
	Listing *ListingView `json:"listing,omitempty"`
	NumberOfShares *int64 `json:"numberOfShares,omitempty"`
	Volume *float32 `json:"volume,omitempty"`
}

// NewAtsLiveDataOrderView instantiates a new AtsLiveDataOrderView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtsLiveDataOrderView() *AtsLiveDataOrderView {
	this := AtsLiveDataOrderView{}
	return &this
}

// NewAtsLiveDataOrderViewWithDefaults instantiates a new AtsLiveDataOrderView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtsLiveDataOrderViewWithDefaults() *AtsLiveDataOrderView {
	this := AtsLiveDataOrderView{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AtsLiveDataOrderView) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtsLiveDataOrderView) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AtsLiveDataOrderView) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *AtsLiveDataOrderView) SetAction(v string) {
	o.Action = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *AtsLiveDataOrderView) GetDate() int64 {
	if o == nil || IsNil(o.Date) {
		var ret int64
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtsLiveDataOrderView) GetDateOk() (*int64, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *AtsLiveDataOrderView) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given int64 and assigns it to the Date field.
func (o *AtsLiveDataOrderView) SetDate(v int64) {
	o.Date = &v
}

// GetDiff returns the Diff field value if set, zero value otherwise.
func (o *AtsLiveDataOrderView) GetDiff() float32 {
	if o == nil || IsNil(o.Diff) {
		var ret float32
		return ret
	}
	return *o.Diff
}

// GetDiffOk returns a tuple with the Diff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtsLiveDataOrderView) GetDiffOk() (*float32, bool) {
	if o == nil || IsNil(o.Diff) {
		return nil, false
	}
	return o.Diff, true
}

// HasDiff returns a boolean if a field has been set.
func (o *AtsLiveDataOrderView) HasDiff() bool {
	if o != nil && !IsNil(o.Diff) {
		return true
	}

	return false
}

// SetDiff gets a reference to the given float32 and assigns it to the Diff field.
func (o *AtsLiveDataOrderView) SetDiff(v float32) {
	o.Diff = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AtsLiveDataOrderView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtsLiveDataOrderView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AtsLiveDataOrderView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AtsLiveDataOrderView) SetId(v string) {
	o.Id = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *AtsLiveDataOrderView) GetLimit() float32 {
	if o == nil || IsNil(o.Limit) {
		var ret float32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtsLiveDataOrderView) GetLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *AtsLiveDataOrderView) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given float32 and assigns it to the Limit field.
func (o *AtsLiveDataOrderView) SetLimit(v float32) {
	o.Limit = &v
}

// GetListing returns the Listing field value if set, zero value otherwise.
func (o *AtsLiveDataOrderView) GetListing() ListingView {
	if o == nil || IsNil(o.Listing) {
		var ret ListingView
		return ret
	}
	return *o.Listing
}

// GetListingOk returns a tuple with the Listing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtsLiveDataOrderView) GetListingOk() (*ListingView, bool) {
	if o == nil || IsNil(o.Listing) {
		return nil, false
	}
	return o.Listing, true
}

// HasListing returns a boolean if a field has been set.
func (o *AtsLiveDataOrderView) HasListing() bool {
	if o != nil && !IsNil(o.Listing) {
		return true
	}

	return false
}

// SetListing gets a reference to the given ListingView and assigns it to the Listing field.
func (o *AtsLiveDataOrderView) SetListing(v ListingView) {
	o.Listing = &v
}

// GetNumberOfShares returns the NumberOfShares field value if set, zero value otherwise.
func (o *AtsLiveDataOrderView) GetNumberOfShares() int64 {
	if o == nil || IsNil(o.NumberOfShares) {
		var ret int64
		return ret
	}
	return *o.NumberOfShares
}

// GetNumberOfSharesOk returns a tuple with the NumberOfShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtsLiveDataOrderView) GetNumberOfSharesOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfShares) {
		return nil, false
	}
	return o.NumberOfShares, true
}

// HasNumberOfShares returns a boolean if a field has been set.
func (o *AtsLiveDataOrderView) HasNumberOfShares() bool {
	if o != nil && !IsNil(o.NumberOfShares) {
		return true
	}

	return false
}

// SetNumberOfShares gets a reference to the given int64 and assigns it to the NumberOfShares field.
func (o *AtsLiveDataOrderView) SetNumberOfShares(v int64) {
	o.NumberOfShares = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *AtsLiveDataOrderView) GetVolume() float32 {
	if o == nil || IsNil(o.Volume) {
		var ret float32
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtsLiveDataOrderView) GetVolumeOk() (*float32, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *AtsLiveDataOrderView) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given float32 and assigns it to the Volume field.
func (o *AtsLiveDataOrderView) SetVolume(v float32) {
	o.Volume = &v
}

func (o AtsLiveDataOrderView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AtsLiveDataOrderView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Diff) {
		toSerialize["diff"] = o.Diff
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Listing) {
		toSerialize["listing"] = o.Listing
	}
	if !IsNil(o.NumberOfShares) {
		toSerialize["numberOfShares"] = o.NumberOfShares
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	return toSerialize, nil
}

type NullableAtsLiveDataOrderView struct {
	value *AtsLiveDataOrderView
	isSet bool
}

func (v NullableAtsLiveDataOrderView) Get() *AtsLiveDataOrderView {
	return v.value
}

func (v *NullableAtsLiveDataOrderView) Set(val *AtsLiveDataOrderView) {
	v.value = val
	v.isSet = true
}

func (v NullableAtsLiveDataOrderView) IsSet() bool {
	return v.isSet
}

func (v *NullableAtsLiveDataOrderView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAtsLiveDataOrderView(val *AtsLiveDataOrderView) *NullableAtsLiveDataOrderView {
	return &NullableAtsLiveDataOrderView{value: val, isSet: true}
}

func (v NullableAtsLiveDataOrderView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAtsLiveDataOrderView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


