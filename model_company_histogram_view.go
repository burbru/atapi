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

// checks if the CompanyHistogramView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyHistogramView{}

// CompanyHistogramView struct for CompanyHistogramView
type CompanyHistogramView struct {
	BondsVolumeHistogram *HistogramView `json:"bondsVolumeHistogram,omitempty"`
	BookValueHistogram *HistogramView `json:"bookValueHistogram,omitempty"`
	CashFlowHistogram *HistogramView `json:"cashFlowHistogram,omitempty"`
	CashHistogram *HistogramView `json:"cashHistogram,omitempty"`
	CentralBankReservesHistogram *HistogramView `json:"centralBankReservesHistogram,omitempty"`
	ClosePriceHistogram *HistogramView `json:"closePriceHistogram,omitempty"`
	Company *CompactCompanyView `json:"company,omitempty"`
	NetCashHistogram *HistogramView `json:"netCashHistogram,omitempty"`
	OutstandingSharesHistogram *HistogramView `json:"outstandingSharesHistogram,omitempty"`
	ReposVolumeHistogram *HistogramView `json:"reposVolumeHistogram,omitempty"`
	SharesInBuysHistogram *HistogramView `json:"sharesInBuysHistogram,omitempty"`
	SharesInSellsHistogram *HistogramView `json:"sharesInSellsHistogram,omitempty"`
	SystemReposVolumeHistogram *HistogramView `json:"systemReposVolumeHistogram,omitempty"`
	TradeVolumeHistogram *HistogramView `json:"tradeVolumeHistogram,omitempty"`
}

// NewCompanyHistogramView instantiates a new CompanyHistogramView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyHistogramView() *CompanyHistogramView {
	this := CompanyHistogramView{}
	return &this
}

// NewCompanyHistogramViewWithDefaults instantiates a new CompanyHistogramView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyHistogramViewWithDefaults() *CompanyHistogramView {
	this := CompanyHistogramView{}
	return &this
}

// GetBondsVolumeHistogram returns the BondsVolumeHistogram field value if set, zero value otherwise.
func (o *CompanyHistogramView) GetBondsVolumeHistogram() HistogramView {
	if o == nil || IsNil(o.BondsVolumeHistogram) {
		var ret HistogramView
		return ret
	}
	return *o.BondsVolumeHistogram
}

// GetBondsVolumeHistogramOk returns a tuple with the BondsVolumeHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyHistogramView) GetBondsVolumeHistogramOk() (*HistogramView, bool) {
	if o == nil || IsNil(o.BondsVolumeHistogram) {
		return nil, false
	}
	return o.BondsVolumeHistogram, true
}

// HasBondsVolumeHistogram returns a boolean if a field has been set.
func (o *CompanyHistogramView) HasBondsVolumeHistogram() bool {
	if o != nil && !IsNil(o.BondsVolumeHistogram) {
		return true
	}

	return false
}

// SetBondsVolumeHistogram gets a reference to the given HistogramView and assigns it to the BondsVolumeHistogram field.
func (o *CompanyHistogramView) SetBondsVolumeHistogram(v HistogramView) {
	o.BondsVolumeHistogram = &v
}

// GetBookValueHistogram returns the BookValueHistogram field value if set, zero value otherwise.
func (o *CompanyHistogramView) GetBookValueHistogram() HistogramView {
	if o == nil || IsNil(o.BookValueHistogram) {
		var ret HistogramView
		return ret
	}
	return *o.BookValueHistogram
}

// GetBookValueHistogramOk returns a tuple with the BookValueHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyHistogramView) GetBookValueHistogramOk() (*HistogramView, bool) {
	if o == nil || IsNil(o.BookValueHistogram) {
		return nil, false
	}
	return o.BookValueHistogram, true
}

// HasBookValueHistogram returns a boolean if a field has been set.
func (o *CompanyHistogramView) HasBookValueHistogram() bool {
	if o != nil && !IsNil(o.BookValueHistogram) {
		return true
	}

	return false
}

// SetBookValueHistogram gets a reference to the given HistogramView and assigns it to the BookValueHistogram field.
func (o *CompanyHistogramView) SetBookValueHistogram(v HistogramView) {
	o.BookValueHistogram = &v
}

// GetCashFlowHistogram returns the CashFlowHistogram field value if set, zero value otherwise.
func (o *CompanyHistogramView) GetCashFlowHistogram() HistogramView {
	if o == nil || IsNil(o.CashFlowHistogram) {
		var ret HistogramView
		return ret
	}
	return *o.CashFlowHistogram
}

// GetCashFlowHistogramOk returns a tuple with the CashFlowHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyHistogramView) GetCashFlowHistogramOk() (*HistogramView, bool) {
	if o == nil || IsNil(o.CashFlowHistogram) {
		return nil, false
	}
	return o.CashFlowHistogram, true
}

// HasCashFlowHistogram returns a boolean if a field has been set.
func (o *CompanyHistogramView) HasCashFlowHistogram() bool {
	if o != nil && !IsNil(o.CashFlowHistogram) {
		return true
	}

	return false
}

// SetCashFlowHistogram gets a reference to the given HistogramView and assigns it to the CashFlowHistogram field.
func (o *CompanyHistogramView) SetCashFlowHistogram(v HistogramView) {
	o.CashFlowHistogram = &v
}

// GetCashHistogram returns the CashHistogram field value if set, zero value otherwise.
func (o *CompanyHistogramView) GetCashHistogram() HistogramView {
	if o == nil || IsNil(o.CashHistogram) {
		var ret HistogramView
		return ret
	}
	return *o.CashHistogram
}

// GetCashHistogramOk returns a tuple with the CashHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyHistogramView) GetCashHistogramOk() (*HistogramView, bool) {
	if o == nil || IsNil(o.CashHistogram) {
		return nil, false
	}
	return o.CashHistogram, true
}

// HasCashHistogram returns a boolean if a field has been set.
func (o *CompanyHistogramView) HasCashHistogram() bool {
	if o != nil && !IsNil(o.CashHistogram) {
		return true
	}

	return false
}

// SetCashHistogram gets a reference to the given HistogramView and assigns it to the CashHistogram field.
func (o *CompanyHistogramView) SetCashHistogram(v HistogramView) {
	o.CashHistogram = &v
}

// GetCentralBankReservesHistogram returns the CentralBankReservesHistogram field value if set, zero value otherwise.
func (o *CompanyHistogramView) GetCentralBankReservesHistogram() HistogramView {
	if o == nil || IsNil(o.CentralBankReservesHistogram) {
		var ret HistogramView
		return ret
	}
	return *o.CentralBankReservesHistogram
}

// GetCentralBankReservesHistogramOk returns a tuple with the CentralBankReservesHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyHistogramView) GetCentralBankReservesHistogramOk() (*HistogramView, bool) {
	if o == nil || IsNil(o.CentralBankReservesHistogram) {
		return nil, false
	}
	return o.CentralBankReservesHistogram, true
}

// HasCentralBankReservesHistogram returns a boolean if a field has been set.
func (o *CompanyHistogramView) HasCentralBankReservesHistogram() bool {
	if o != nil && !IsNil(o.CentralBankReservesHistogram) {
		return true
	}

	return false
}

// SetCentralBankReservesHistogram gets a reference to the given HistogramView and assigns it to the CentralBankReservesHistogram field.
func (o *CompanyHistogramView) SetCentralBankReservesHistogram(v HistogramView) {
	o.CentralBankReservesHistogram = &v
}

// GetClosePriceHistogram returns the ClosePriceHistogram field value if set, zero value otherwise.
func (o *CompanyHistogramView) GetClosePriceHistogram() HistogramView {
	if o == nil || IsNil(o.ClosePriceHistogram) {
		var ret HistogramView
		return ret
	}
	return *o.ClosePriceHistogram
}

// GetClosePriceHistogramOk returns a tuple with the ClosePriceHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyHistogramView) GetClosePriceHistogramOk() (*HistogramView, bool) {
	if o == nil || IsNil(o.ClosePriceHistogram) {
		return nil, false
	}
	return o.ClosePriceHistogram, true
}

// HasClosePriceHistogram returns a boolean if a field has been set.
func (o *CompanyHistogramView) HasClosePriceHistogram() bool {
	if o != nil && !IsNil(o.ClosePriceHistogram) {
		return true
	}

	return false
}

// SetClosePriceHistogram gets a reference to the given HistogramView and assigns it to the ClosePriceHistogram field.
func (o *CompanyHistogramView) SetClosePriceHistogram(v HistogramView) {
	o.ClosePriceHistogram = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *CompanyHistogramView) GetCompany() CompactCompanyView {
	if o == nil || IsNil(o.Company) {
		var ret CompactCompanyView
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyHistogramView) GetCompanyOk() (*CompactCompanyView, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *CompanyHistogramView) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given CompactCompanyView and assigns it to the Company field.
func (o *CompanyHistogramView) SetCompany(v CompactCompanyView) {
	o.Company = &v
}

// GetNetCashHistogram returns the NetCashHistogram field value if set, zero value otherwise.
func (o *CompanyHistogramView) GetNetCashHistogram() HistogramView {
	if o == nil || IsNil(o.NetCashHistogram) {
		var ret HistogramView
		return ret
	}
	return *o.NetCashHistogram
}

// GetNetCashHistogramOk returns a tuple with the NetCashHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyHistogramView) GetNetCashHistogramOk() (*HistogramView, bool) {
	if o == nil || IsNil(o.NetCashHistogram) {
		return nil, false
	}
	return o.NetCashHistogram, true
}

// HasNetCashHistogram returns a boolean if a field has been set.
func (o *CompanyHistogramView) HasNetCashHistogram() bool {
	if o != nil && !IsNil(o.NetCashHistogram) {
		return true
	}

	return false
}

// SetNetCashHistogram gets a reference to the given HistogramView and assigns it to the NetCashHistogram field.
func (o *CompanyHistogramView) SetNetCashHistogram(v HistogramView) {
	o.NetCashHistogram = &v
}

// GetOutstandingSharesHistogram returns the OutstandingSharesHistogram field value if set, zero value otherwise.
func (o *CompanyHistogramView) GetOutstandingSharesHistogram() HistogramView {
	if o == nil || IsNil(o.OutstandingSharesHistogram) {
		var ret HistogramView
		return ret
	}
	return *o.OutstandingSharesHistogram
}

// GetOutstandingSharesHistogramOk returns a tuple with the OutstandingSharesHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyHistogramView) GetOutstandingSharesHistogramOk() (*HistogramView, bool) {
	if o == nil || IsNil(o.OutstandingSharesHistogram) {
		return nil, false
	}
	return o.OutstandingSharesHistogram, true
}

// HasOutstandingSharesHistogram returns a boolean if a field has been set.
func (o *CompanyHistogramView) HasOutstandingSharesHistogram() bool {
	if o != nil && !IsNil(o.OutstandingSharesHistogram) {
		return true
	}

	return false
}

// SetOutstandingSharesHistogram gets a reference to the given HistogramView and assigns it to the OutstandingSharesHistogram field.
func (o *CompanyHistogramView) SetOutstandingSharesHistogram(v HistogramView) {
	o.OutstandingSharesHistogram = &v
}

// GetReposVolumeHistogram returns the ReposVolumeHistogram field value if set, zero value otherwise.
func (o *CompanyHistogramView) GetReposVolumeHistogram() HistogramView {
	if o == nil || IsNil(o.ReposVolumeHistogram) {
		var ret HistogramView
		return ret
	}
	return *o.ReposVolumeHistogram
}

// GetReposVolumeHistogramOk returns a tuple with the ReposVolumeHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyHistogramView) GetReposVolumeHistogramOk() (*HistogramView, bool) {
	if o == nil || IsNil(o.ReposVolumeHistogram) {
		return nil, false
	}
	return o.ReposVolumeHistogram, true
}

// HasReposVolumeHistogram returns a boolean if a field has been set.
func (o *CompanyHistogramView) HasReposVolumeHistogram() bool {
	if o != nil && !IsNil(o.ReposVolumeHistogram) {
		return true
	}

	return false
}

// SetReposVolumeHistogram gets a reference to the given HistogramView and assigns it to the ReposVolumeHistogram field.
func (o *CompanyHistogramView) SetReposVolumeHistogram(v HistogramView) {
	o.ReposVolumeHistogram = &v
}

// GetSharesInBuysHistogram returns the SharesInBuysHistogram field value if set, zero value otherwise.
func (o *CompanyHistogramView) GetSharesInBuysHistogram() HistogramView {
	if o == nil || IsNil(o.SharesInBuysHistogram) {
		var ret HistogramView
		return ret
	}
	return *o.SharesInBuysHistogram
}

// GetSharesInBuysHistogramOk returns a tuple with the SharesInBuysHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyHistogramView) GetSharesInBuysHistogramOk() (*HistogramView, bool) {
	if o == nil || IsNil(o.SharesInBuysHistogram) {
		return nil, false
	}
	return o.SharesInBuysHistogram, true
}

// HasSharesInBuysHistogram returns a boolean if a field has been set.
func (o *CompanyHistogramView) HasSharesInBuysHistogram() bool {
	if o != nil && !IsNil(o.SharesInBuysHistogram) {
		return true
	}

	return false
}

// SetSharesInBuysHistogram gets a reference to the given HistogramView and assigns it to the SharesInBuysHistogram field.
func (o *CompanyHistogramView) SetSharesInBuysHistogram(v HistogramView) {
	o.SharesInBuysHistogram = &v
}

// GetSharesInSellsHistogram returns the SharesInSellsHistogram field value if set, zero value otherwise.
func (o *CompanyHistogramView) GetSharesInSellsHistogram() HistogramView {
	if o == nil || IsNil(o.SharesInSellsHistogram) {
		var ret HistogramView
		return ret
	}
	return *o.SharesInSellsHistogram
}

// GetSharesInSellsHistogramOk returns a tuple with the SharesInSellsHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyHistogramView) GetSharesInSellsHistogramOk() (*HistogramView, bool) {
	if o == nil || IsNil(o.SharesInSellsHistogram) {
		return nil, false
	}
	return o.SharesInSellsHistogram, true
}

// HasSharesInSellsHistogram returns a boolean if a field has been set.
func (o *CompanyHistogramView) HasSharesInSellsHistogram() bool {
	if o != nil && !IsNil(o.SharesInSellsHistogram) {
		return true
	}

	return false
}

// SetSharesInSellsHistogram gets a reference to the given HistogramView and assigns it to the SharesInSellsHistogram field.
func (o *CompanyHistogramView) SetSharesInSellsHistogram(v HistogramView) {
	o.SharesInSellsHistogram = &v
}

// GetSystemReposVolumeHistogram returns the SystemReposVolumeHistogram field value if set, zero value otherwise.
func (o *CompanyHistogramView) GetSystemReposVolumeHistogram() HistogramView {
	if o == nil || IsNil(o.SystemReposVolumeHistogram) {
		var ret HistogramView
		return ret
	}
	return *o.SystemReposVolumeHistogram
}

// GetSystemReposVolumeHistogramOk returns a tuple with the SystemReposVolumeHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyHistogramView) GetSystemReposVolumeHistogramOk() (*HistogramView, bool) {
	if o == nil || IsNil(o.SystemReposVolumeHistogram) {
		return nil, false
	}
	return o.SystemReposVolumeHistogram, true
}

// HasSystemReposVolumeHistogram returns a boolean if a field has been set.
func (o *CompanyHistogramView) HasSystemReposVolumeHistogram() bool {
	if o != nil && !IsNil(o.SystemReposVolumeHistogram) {
		return true
	}

	return false
}

// SetSystemReposVolumeHistogram gets a reference to the given HistogramView and assigns it to the SystemReposVolumeHistogram field.
func (o *CompanyHistogramView) SetSystemReposVolumeHistogram(v HistogramView) {
	o.SystemReposVolumeHistogram = &v
}

// GetTradeVolumeHistogram returns the TradeVolumeHistogram field value if set, zero value otherwise.
func (o *CompanyHistogramView) GetTradeVolumeHistogram() HistogramView {
	if o == nil || IsNil(o.TradeVolumeHistogram) {
		var ret HistogramView
		return ret
	}
	return *o.TradeVolumeHistogram
}

// GetTradeVolumeHistogramOk returns a tuple with the TradeVolumeHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyHistogramView) GetTradeVolumeHistogramOk() (*HistogramView, bool) {
	if o == nil || IsNil(o.TradeVolumeHistogram) {
		return nil, false
	}
	return o.TradeVolumeHistogram, true
}

// HasTradeVolumeHistogram returns a boolean if a field has been set.
func (o *CompanyHistogramView) HasTradeVolumeHistogram() bool {
	if o != nil && !IsNil(o.TradeVolumeHistogram) {
		return true
	}

	return false
}

// SetTradeVolumeHistogram gets a reference to the given HistogramView and assigns it to the TradeVolumeHistogram field.
func (o *CompanyHistogramView) SetTradeVolumeHistogram(v HistogramView) {
	o.TradeVolumeHistogram = &v
}

func (o CompanyHistogramView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyHistogramView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BondsVolumeHistogram) {
		toSerialize["bondsVolumeHistogram"] = o.BondsVolumeHistogram
	}
	if !IsNil(o.BookValueHistogram) {
		toSerialize["bookValueHistogram"] = o.BookValueHistogram
	}
	if !IsNil(o.CashFlowHistogram) {
		toSerialize["cashFlowHistogram"] = o.CashFlowHistogram
	}
	if !IsNil(o.CashHistogram) {
		toSerialize["cashHistogram"] = o.CashHistogram
	}
	if !IsNil(o.CentralBankReservesHistogram) {
		toSerialize["centralBankReservesHistogram"] = o.CentralBankReservesHistogram
	}
	if !IsNil(o.ClosePriceHistogram) {
		toSerialize["closePriceHistogram"] = o.ClosePriceHistogram
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.NetCashHistogram) {
		toSerialize["netCashHistogram"] = o.NetCashHistogram
	}
	if !IsNil(o.OutstandingSharesHistogram) {
		toSerialize["outstandingSharesHistogram"] = o.OutstandingSharesHistogram
	}
	if !IsNil(o.ReposVolumeHistogram) {
		toSerialize["reposVolumeHistogram"] = o.ReposVolumeHistogram
	}
	if !IsNil(o.SharesInBuysHistogram) {
		toSerialize["sharesInBuysHistogram"] = o.SharesInBuysHistogram
	}
	if !IsNil(o.SharesInSellsHistogram) {
		toSerialize["sharesInSellsHistogram"] = o.SharesInSellsHistogram
	}
	if !IsNil(o.SystemReposVolumeHistogram) {
		toSerialize["systemReposVolumeHistogram"] = o.SystemReposVolumeHistogram
	}
	if !IsNil(o.TradeVolumeHistogram) {
		toSerialize["tradeVolumeHistogram"] = o.TradeVolumeHistogram
	}
	return toSerialize, nil
}

type NullableCompanyHistogramView struct {
	value *CompanyHistogramView
	isSet bool
}

func (v NullableCompanyHistogramView) Get() *CompanyHistogramView {
	return v.value
}

func (v *NullableCompanyHistogramView) Set(val *CompanyHistogramView) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyHistogramView) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyHistogramView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyHistogramView(val *CompanyHistogramView) *NullableCompanyHistogramView {
	return &NullableCompanyHistogramView{value: val, isSet: true}
}

func (v NullableCompanyHistogramView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyHistogramView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


