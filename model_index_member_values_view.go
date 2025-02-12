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

// checks if the IndexMemberValuesView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndexMemberValuesView{}

// IndexMemberValuesView struct for IndexMemberValuesView
type IndexMemberValuesView struct {
	BaseCapitalisation *float32 `json:"baseCapitalisation,omitempty"`
	BasePrice *float32 `json:"basePrice,omitempty"`
	BaseShares *int64 `json:"baseShares,omitempty"`
	Capitalisation *float32 `json:"capitalisation,omitempty"`
	LastAdjustmentDate *int64 `json:"lastAdjustmentDate,omitempty"`
	Listing *ListingView `json:"listing,omitempty"`
	Price *float32 `json:"price,omitempty"`
	PriceAdjustmentFactor *float32 `json:"priceAdjustmentFactor,omitempty"`
	Shares *int64 `json:"shares,omitempty"`
}

// NewIndexMemberValuesView instantiates a new IndexMemberValuesView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexMemberValuesView() *IndexMemberValuesView {
	this := IndexMemberValuesView{}
	return &this
}

// NewIndexMemberValuesViewWithDefaults instantiates a new IndexMemberValuesView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexMemberValuesViewWithDefaults() *IndexMemberValuesView {
	this := IndexMemberValuesView{}
	return &this
}

// GetBaseCapitalisation returns the BaseCapitalisation field value if set, zero value otherwise.
func (o *IndexMemberValuesView) GetBaseCapitalisation() float32 {
	if o == nil || IsNil(o.BaseCapitalisation) {
		var ret float32
		return ret
	}
	return *o.BaseCapitalisation
}

// GetBaseCapitalisationOk returns a tuple with the BaseCapitalisation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexMemberValuesView) GetBaseCapitalisationOk() (*float32, bool) {
	if o == nil || IsNil(o.BaseCapitalisation) {
		return nil, false
	}
	return o.BaseCapitalisation, true
}

// HasBaseCapitalisation returns a boolean if a field has been set.
func (o *IndexMemberValuesView) HasBaseCapitalisation() bool {
	if o != nil && !IsNil(o.BaseCapitalisation) {
		return true
	}

	return false
}

// SetBaseCapitalisation gets a reference to the given float32 and assigns it to the BaseCapitalisation field.
func (o *IndexMemberValuesView) SetBaseCapitalisation(v float32) {
	o.BaseCapitalisation = &v
}

// GetBasePrice returns the BasePrice field value if set, zero value otherwise.
func (o *IndexMemberValuesView) GetBasePrice() float32 {
	if o == nil || IsNil(o.BasePrice) {
		var ret float32
		return ret
	}
	return *o.BasePrice
}

// GetBasePriceOk returns a tuple with the BasePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexMemberValuesView) GetBasePriceOk() (*float32, bool) {
	if o == nil || IsNil(o.BasePrice) {
		return nil, false
	}
	return o.BasePrice, true
}

// HasBasePrice returns a boolean if a field has been set.
func (o *IndexMemberValuesView) HasBasePrice() bool {
	if o != nil && !IsNil(o.BasePrice) {
		return true
	}

	return false
}

// SetBasePrice gets a reference to the given float32 and assigns it to the BasePrice field.
func (o *IndexMemberValuesView) SetBasePrice(v float32) {
	o.BasePrice = &v
}

// GetBaseShares returns the BaseShares field value if set, zero value otherwise.
func (o *IndexMemberValuesView) GetBaseShares() int64 {
	if o == nil || IsNil(o.BaseShares) {
		var ret int64
		return ret
	}
	return *o.BaseShares
}

// GetBaseSharesOk returns a tuple with the BaseShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexMemberValuesView) GetBaseSharesOk() (*int64, bool) {
	if o == nil || IsNil(o.BaseShares) {
		return nil, false
	}
	return o.BaseShares, true
}

// HasBaseShares returns a boolean if a field has been set.
func (o *IndexMemberValuesView) HasBaseShares() bool {
	if o != nil && !IsNil(o.BaseShares) {
		return true
	}

	return false
}

// SetBaseShares gets a reference to the given int64 and assigns it to the BaseShares field.
func (o *IndexMemberValuesView) SetBaseShares(v int64) {
	o.BaseShares = &v
}

// GetCapitalisation returns the Capitalisation field value if set, zero value otherwise.
func (o *IndexMemberValuesView) GetCapitalisation() float32 {
	if o == nil || IsNil(o.Capitalisation) {
		var ret float32
		return ret
	}
	return *o.Capitalisation
}

// GetCapitalisationOk returns a tuple with the Capitalisation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexMemberValuesView) GetCapitalisationOk() (*float32, bool) {
	if o == nil || IsNil(o.Capitalisation) {
		return nil, false
	}
	return o.Capitalisation, true
}

// HasCapitalisation returns a boolean if a field has been set.
func (o *IndexMemberValuesView) HasCapitalisation() bool {
	if o != nil && !IsNil(o.Capitalisation) {
		return true
	}

	return false
}

// SetCapitalisation gets a reference to the given float32 and assigns it to the Capitalisation field.
func (o *IndexMemberValuesView) SetCapitalisation(v float32) {
	o.Capitalisation = &v
}

// GetLastAdjustmentDate returns the LastAdjustmentDate field value if set, zero value otherwise.
func (o *IndexMemberValuesView) GetLastAdjustmentDate() int64 {
	if o == nil || IsNil(o.LastAdjustmentDate) {
		var ret int64
		return ret
	}
	return *o.LastAdjustmentDate
}

// GetLastAdjustmentDateOk returns a tuple with the LastAdjustmentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexMemberValuesView) GetLastAdjustmentDateOk() (*int64, bool) {
	if o == nil || IsNil(o.LastAdjustmentDate) {
		return nil, false
	}
	return o.LastAdjustmentDate, true
}

// HasLastAdjustmentDate returns a boolean if a field has been set.
func (o *IndexMemberValuesView) HasLastAdjustmentDate() bool {
	if o != nil && !IsNil(o.LastAdjustmentDate) {
		return true
	}

	return false
}

// SetLastAdjustmentDate gets a reference to the given int64 and assigns it to the LastAdjustmentDate field.
func (o *IndexMemberValuesView) SetLastAdjustmentDate(v int64) {
	o.LastAdjustmentDate = &v
}

// GetListing returns the Listing field value if set, zero value otherwise.
func (o *IndexMemberValuesView) GetListing() ListingView {
	if o == nil || IsNil(o.Listing) {
		var ret ListingView
		return ret
	}
	return *o.Listing
}

// GetListingOk returns a tuple with the Listing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexMemberValuesView) GetListingOk() (*ListingView, bool) {
	if o == nil || IsNil(o.Listing) {
		return nil, false
	}
	return o.Listing, true
}

// HasListing returns a boolean if a field has been set.
func (o *IndexMemberValuesView) HasListing() bool {
	if o != nil && !IsNil(o.Listing) {
		return true
	}

	return false
}

// SetListing gets a reference to the given ListingView and assigns it to the Listing field.
func (o *IndexMemberValuesView) SetListing(v ListingView) {
	o.Listing = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *IndexMemberValuesView) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexMemberValuesView) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *IndexMemberValuesView) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *IndexMemberValuesView) SetPrice(v float32) {
	o.Price = &v
}

// GetPriceAdjustmentFactor returns the PriceAdjustmentFactor field value if set, zero value otherwise.
func (o *IndexMemberValuesView) GetPriceAdjustmentFactor() float32 {
	if o == nil || IsNil(o.PriceAdjustmentFactor) {
		var ret float32
		return ret
	}
	return *o.PriceAdjustmentFactor
}

// GetPriceAdjustmentFactorOk returns a tuple with the PriceAdjustmentFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexMemberValuesView) GetPriceAdjustmentFactorOk() (*float32, bool) {
	if o == nil || IsNil(o.PriceAdjustmentFactor) {
		return nil, false
	}
	return o.PriceAdjustmentFactor, true
}

// HasPriceAdjustmentFactor returns a boolean if a field has been set.
func (o *IndexMemberValuesView) HasPriceAdjustmentFactor() bool {
	if o != nil && !IsNil(o.PriceAdjustmentFactor) {
		return true
	}

	return false
}

// SetPriceAdjustmentFactor gets a reference to the given float32 and assigns it to the PriceAdjustmentFactor field.
func (o *IndexMemberValuesView) SetPriceAdjustmentFactor(v float32) {
	o.PriceAdjustmentFactor = &v
}

// GetShares returns the Shares field value if set, zero value otherwise.
func (o *IndexMemberValuesView) GetShares() int64 {
	if o == nil || IsNil(o.Shares) {
		var ret int64
		return ret
	}
	return *o.Shares
}

// GetSharesOk returns a tuple with the Shares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexMemberValuesView) GetSharesOk() (*int64, bool) {
	if o == nil || IsNil(o.Shares) {
		return nil, false
	}
	return o.Shares, true
}

// HasShares returns a boolean if a field has been set.
func (o *IndexMemberValuesView) HasShares() bool {
	if o != nil && !IsNil(o.Shares) {
		return true
	}

	return false
}

// SetShares gets a reference to the given int64 and assigns it to the Shares field.
func (o *IndexMemberValuesView) SetShares(v int64) {
	o.Shares = &v
}

func (o IndexMemberValuesView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexMemberValuesView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BaseCapitalisation) {
		toSerialize["baseCapitalisation"] = o.BaseCapitalisation
	}
	if !IsNil(o.BasePrice) {
		toSerialize["basePrice"] = o.BasePrice
	}
	if !IsNil(o.BaseShares) {
		toSerialize["baseShares"] = o.BaseShares
	}
	if !IsNil(o.Capitalisation) {
		toSerialize["capitalisation"] = o.Capitalisation
	}
	if !IsNil(o.LastAdjustmentDate) {
		toSerialize["lastAdjustmentDate"] = o.LastAdjustmentDate
	}
	if !IsNil(o.Listing) {
		toSerialize["listing"] = o.Listing
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.PriceAdjustmentFactor) {
		toSerialize["priceAdjustmentFactor"] = o.PriceAdjustmentFactor
	}
	if !IsNil(o.Shares) {
		toSerialize["shares"] = o.Shares
	}
	return toSerialize, nil
}

type NullableIndexMemberValuesView struct {
	value *IndexMemberValuesView
	isSet bool
}

func (v NullableIndexMemberValuesView) Get() *IndexMemberValuesView {
	return v.value
}

func (v *NullableIndexMemberValuesView) Set(val *IndexMemberValuesView) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexMemberValuesView) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexMemberValuesView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexMemberValuesView(val *IndexMemberValuesView) *NullableIndexMemberValuesView {
	return &NullableIndexMemberValuesView{value: val, isSet: true}
}

func (v NullableIndexMemberValuesView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexMemberValuesView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


