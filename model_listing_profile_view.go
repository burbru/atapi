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

// checks if the ListingProfileView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListingProfileView{}

// ListingProfileView struct for ListingProfileView
type ListingProfileView struct {
	Bond *BondView `json:"bond,omitempty"`
	Building *BuildingView `json:"building,omitempty"`
	Company *CompanyCompactProfileView `json:"company,omitempty"`
	CurrentSpread *PriceSpreadView `json:"currentSpread,omitempty"`
	DesignatedSponsors []SecuritySponsorshipView `json:"designatedSponsors,omitempty"`
	EndDate *int64 `json:"endDate,omitempty"`
	Id *string `json:"id,omitempty"`
	LastOrderLogEntry *SecurityOrderLogEntryView `json:"lastOrderLogEntry,omitempty"`
	LastPrice *SecurityPriceView `json:"lastPrice,omitempty"`
	LogoUrl *string `json:"logoUrl,omitempty"`
	MarketCap *float32 `json:"marketCap,omitempty"`
	Name *string `json:"name,omitempty"`
	OutstandingShares *int64 `json:"outstandingShares,omitempty"`
	Prices14d []SecurityPriceView `json:"prices14d,omitempty"`
	SecurityIdentifier *string `json:"securityIdentifier,omitempty"`
	ShareholderStake *ShareholderStake `json:"shareholderStake,omitempty"`
	StartDate *int64 `json:"startDate,omitempty"`
	SystemBond *SystemBondView `json:"systemBond,omitempty"`
	Type *string `json:"type,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewListingProfileView instantiates a new ListingProfileView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListingProfileView() *ListingProfileView {
	this := ListingProfileView{}
	return &this
}

// NewListingProfileViewWithDefaults instantiates a new ListingProfileView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListingProfileViewWithDefaults() *ListingProfileView {
	this := ListingProfileView{}
	return &this
}

// GetBond returns the Bond field value if set, zero value otherwise.
func (o *ListingProfileView) GetBond() BondView {
	if o == nil || IsNil(o.Bond) {
		var ret BondView
		return ret
	}
	return *o.Bond
}

// GetBondOk returns a tuple with the Bond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetBondOk() (*BondView, bool) {
	if o == nil || IsNil(o.Bond) {
		return nil, false
	}
	return o.Bond, true
}

// HasBond returns a boolean if a field has been set.
func (o *ListingProfileView) HasBond() bool {
	if o != nil && !IsNil(o.Bond) {
		return true
	}

	return false
}

// SetBond gets a reference to the given BondView and assigns it to the Bond field.
func (o *ListingProfileView) SetBond(v BondView) {
	o.Bond = &v
}

// GetBuilding returns the Building field value if set, zero value otherwise.
func (o *ListingProfileView) GetBuilding() BuildingView {
	if o == nil || IsNil(o.Building) {
		var ret BuildingView
		return ret
	}
	return *o.Building
}

// GetBuildingOk returns a tuple with the Building field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetBuildingOk() (*BuildingView, bool) {
	if o == nil || IsNil(o.Building) {
		return nil, false
	}
	return o.Building, true
}

// HasBuilding returns a boolean if a field has been set.
func (o *ListingProfileView) HasBuilding() bool {
	if o != nil && !IsNil(o.Building) {
		return true
	}

	return false
}

// SetBuilding gets a reference to the given BuildingView and assigns it to the Building field.
func (o *ListingProfileView) SetBuilding(v BuildingView) {
	o.Building = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *ListingProfileView) GetCompany() CompanyCompactProfileView {
	if o == nil || IsNil(o.Company) {
		var ret CompanyCompactProfileView
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetCompanyOk() (*CompanyCompactProfileView, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *ListingProfileView) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given CompanyCompactProfileView and assigns it to the Company field.
func (o *ListingProfileView) SetCompany(v CompanyCompactProfileView) {
	o.Company = &v
}

// GetCurrentSpread returns the CurrentSpread field value if set, zero value otherwise.
func (o *ListingProfileView) GetCurrentSpread() PriceSpreadView {
	if o == nil || IsNil(o.CurrentSpread) {
		var ret PriceSpreadView
		return ret
	}
	return *o.CurrentSpread
}

// GetCurrentSpreadOk returns a tuple with the CurrentSpread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetCurrentSpreadOk() (*PriceSpreadView, bool) {
	if o == nil || IsNil(o.CurrentSpread) {
		return nil, false
	}
	return o.CurrentSpread, true
}

// HasCurrentSpread returns a boolean if a field has been set.
func (o *ListingProfileView) HasCurrentSpread() bool {
	if o != nil && !IsNil(o.CurrentSpread) {
		return true
	}

	return false
}

// SetCurrentSpread gets a reference to the given PriceSpreadView and assigns it to the CurrentSpread field.
func (o *ListingProfileView) SetCurrentSpread(v PriceSpreadView) {
	o.CurrentSpread = &v
}

// GetDesignatedSponsors returns the DesignatedSponsors field value if set, zero value otherwise.
func (o *ListingProfileView) GetDesignatedSponsors() []SecuritySponsorshipView {
	if o == nil || IsNil(o.DesignatedSponsors) {
		var ret []SecuritySponsorshipView
		return ret
	}
	return o.DesignatedSponsors
}

// GetDesignatedSponsorsOk returns a tuple with the DesignatedSponsors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetDesignatedSponsorsOk() ([]SecuritySponsorshipView, bool) {
	if o == nil || IsNil(o.DesignatedSponsors) {
		return nil, false
	}
	return o.DesignatedSponsors, true
}

// HasDesignatedSponsors returns a boolean if a field has been set.
func (o *ListingProfileView) HasDesignatedSponsors() bool {
	if o != nil && !IsNil(o.DesignatedSponsors) {
		return true
	}

	return false
}

// SetDesignatedSponsors gets a reference to the given []SecuritySponsorshipView and assigns it to the DesignatedSponsors field.
func (o *ListingProfileView) SetDesignatedSponsors(v []SecuritySponsorshipView) {
	o.DesignatedSponsors = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ListingProfileView) GetEndDate() int64 {
	if o == nil || IsNil(o.EndDate) {
		var ret int64
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetEndDateOk() (*int64, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ListingProfileView) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int64 and assigns it to the EndDate field.
func (o *ListingProfileView) SetEndDate(v int64) {
	o.EndDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListingProfileView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListingProfileView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListingProfileView) SetId(v string) {
	o.Id = &v
}

// GetLastOrderLogEntry returns the LastOrderLogEntry field value if set, zero value otherwise.
func (o *ListingProfileView) GetLastOrderLogEntry() SecurityOrderLogEntryView {
	if o == nil || IsNil(o.LastOrderLogEntry) {
		var ret SecurityOrderLogEntryView
		return ret
	}
	return *o.LastOrderLogEntry
}

// GetLastOrderLogEntryOk returns a tuple with the LastOrderLogEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetLastOrderLogEntryOk() (*SecurityOrderLogEntryView, bool) {
	if o == nil || IsNil(o.LastOrderLogEntry) {
		return nil, false
	}
	return o.LastOrderLogEntry, true
}

// HasLastOrderLogEntry returns a boolean if a field has been set.
func (o *ListingProfileView) HasLastOrderLogEntry() bool {
	if o != nil && !IsNil(o.LastOrderLogEntry) {
		return true
	}

	return false
}

// SetLastOrderLogEntry gets a reference to the given SecurityOrderLogEntryView and assigns it to the LastOrderLogEntry field.
func (o *ListingProfileView) SetLastOrderLogEntry(v SecurityOrderLogEntryView) {
	o.LastOrderLogEntry = &v
}

// GetLastPrice returns the LastPrice field value if set, zero value otherwise.
func (o *ListingProfileView) GetLastPrice() SecurityPriceView {
	if o == nil || IsNil(o.LastPrice) {
		var ret SecurityPriceView
		return ret
	}
	return *o.LastPrice
}

// GetLastPriceOk returns a tuple with the LastPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetLastPriceOk() (*SecurityPriceView, bool) {
	if o == nil || IsNil(o.LastPrice) {
		return nil, false
	}
	return o.LastPrice, true
}

// HasLastPrice returns a boolean if a field has been set.
func (o *ListingProfileView) HasLastPrice() bool {
	if o != nil && !IsNil(o.LastPrice) {
		return true
	}

	return false
}

// SetLastPrice gets a reference to the given SecurityPriceView and assigns it to the LastPrice field.
func (o *ListingProfileView) SetLastPrice(v SecurityPriceView) {
	o.LastPrice = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *ListingProfileView) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *ListingProfileView) HasLogoUrl() bool {
	if o != nil && !IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *ListingProfileView) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetMarketCap returns the MarketCap field value if set, zero value otherwise.
func (o *ListingProfileView) GetMarketCap() float32 {
	if o == nil || IsNil(o.MarketCap) {
		var ret float32
		return ret
	}
	return *o.MarketCap
}

// GetMarketCapOk returns a tuple with the MarketCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetMarketCapOk() (*float32, bool) {
	if o == nil || IsNil(o.MarketCap) {
		return nil, false
	}
	return o.MarketCap, true
}

// HasMarketCap returns a boolean if a field has been set.
func (o *ListingProfileView) HasMarketCap() bool {
	if o != nil && !IsNil(o.MarketCap) {
		return true
	}

	return false
}

// SetMarketCap gets a reference to the given float32 and assigns it to the MarketCap field.
func (o *ListingProfileView) SetMarketCap(v float32) {
	o.MarketCap = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListingProfileView) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListingProfileView) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListingProfileView) SetName(v string) {
	o.Name = &v
}

// GetOutstandingShares returns the OutstandingShares field value if set, zero value otherwise.
func (o *ListingProfileView) GetOutstandingShares() int64 {
	if o == nil || IsNil(o.OutstandingShares) {
		var ret int64
		return ret
	}
	return *o.OutstandingShares
}

// GetOutstandingSharesOk returns a tuple with the OutstandingShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetOutstandingSharesOk() (*int64, bool) {
	if o == nil || IsNil(o.OutstandingShares) {
		return nil, false
	}
	return o.OutstandingShares, true
}

// HasOutstandingShares returns a boolean if a field has been set.
func (o *ListingProfileView) HasOutstandingShares() bool {
	if o != nil && !IsNil(o.OutstandingShares) {
		return true
	}

	return false
}

// SetOutstandingShares gets a reference to the given int64 and assigns it to the OutstandingShares field.
func (o *ListingProfileView) SetOutstandingShares(v int64) {
	o.OutstandingShares = &v
}

// GetPrices14d returns the Prices14d field value if set, zero value otherwise.
func (o *ListingProfileView) GetPrices14d() []SecurityPriceView {
	if o == nil || IsNil(o.Prices14d) {
		var ret []SecurityPriceView
		return ret
	}
	return o.Prices14d
}

// GetPrices14dOk returns a tuple with the Prices14d field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetPrices14dOk() ([]SecurityPriceView, bool) {
	if o == nil || IsNil(o.Prices14d) {
		return nil, false
	}
	return o.Prices14d, true
}

// HasPrices14d returns a boolean if a field has been set.
func (o *ListingProfileView) HasPrices14d() bool {
	if o != nil && !IsNil(o.Prices14d) {
		return true
	}

	return false
}

// SetPrices14d gets a reference to the given []SecurityPriceView and assigns it to the Prices14d field.
func (o *ListingProfileView) SetPrices14d(v []SecurityPriceView) {
	o.Prices14d = v
}

// GetSecurityIdentifier returns the SecurityIdentifier field value if set, zero value otherwise.
func (o *ListingProfileView) GetSecurityIdentifier() string {
	if o == nil || IsNil(o.SecurityIdentifier) {
		var ret string
		return ret
	}
	return *o.SecurityIdentifier
}

// GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetSecurityIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityIdentifier) {
		return nil, false
	}
	return o.SecurityIdentifier, true
}

// HasSecurityIdentifier returns a boolean if a field has been set.
func (o *ListingProfileView) HasSecurityIdentifier() bool {
	if o != nil && !IsNil(o.SecurityIdentifier) {
		return true
	}

	return false
}

// SetSecurityIdentifier gets a reference to the given string and assigns it to the SecurityIdentifier field.
func (o *ListingProfileView) SetSecurityIdentifier(v string) {
	o.SecurityIdentifier = &v
}

// GetShareholderStake returns the ShareholderStake field value if set, zero value otherwise.
func (o *ListingProfileView) GetShareholderStake() ShareholderStake {
	if o == nil || IsNil(o.ShareholderStake) {
		var ret ShareholderStake
		return ret
	}
	return *o.ShareholderStake
}

// GetShareholderStakeOk returns a tuple with the ShareholderStake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetShareholderStakeOk() (*ShareholderStake, bool) {
	if o == nil || IsNil(o.ShareholderStake) {
		return nil, false
	}
	return o.ShareholderStake, true
}

// HasShareholderStake returns a boolean if a field has been set.
func (o *ListingProfileView) HasShareholderStake() bool {
	if o != nil && !IsNil(o.ShareholderStake) {
		return true
	}

	return false
}

// SetShareholderStake gets a reference to the given ShareholderStake and assigns it to the ShareholderStake field.
func (o *ListingProfileView) SetShareholderStake(v ShareholderStake) {
	o.ShareholderStake = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ListingProfileView) GetStartDate() int64 {
	if o == nil || IsNil(o.StartDate) {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetStartDateOk() (*int64, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ListingProfileView) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *ListingProfileView) SetStartDate(v int64) {
	o.StartDate = &v
}

// GetSystemBond returns the SystemBond field value if set, zero value otherwise.
func (o *ListingProfileView) GetSystemBond() SystemBondView {
	if o == nil || IsNil(o.SystemBond) {
		var ret SystemBondView
		return ret
	}
	return *o.SystemBond
}

// GetSystemBondOk returns a tuple with the SystemBond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetSystemBondOk() (*SystemBondView, bool) {
	if o == nil || IsNil(o.SystemBond) {
		return nil, false
	}
	return o.SystemBond, true
}

// HasSystemBond returns a boolean if a field has been set.
func (o *ListingProfileView) HasSystemBond() bool {
	if o != nil && !IsNil(o.SystemBond) {
		return true
	}

	return false
}

// SetSystemBond gets a reference to the given SystemBondView and assigns it to the SystemBond field.
func (o *ListingProfileView) SetSystemBond(v SystemBondView) {
	o.SystemBond = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListingProfileView) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ListingProfileView) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ListingProfileView) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ListingProfileView) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingProfileView) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ListingProfileView) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *ListingProfileView) SetVersion(v int64) {
	o.Version = &v
}

func (o ListingProfileView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListingProfileView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bond) {
		toSerialize["bond"] = o.Bond
	}
	if !IsNil(o.Building) {
		toSerialize["building"] = o.Building
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.CurrentSpread) {
		toSerialize["currentSpread"] = o.CurrentSpread
	}
	if !IsNil(o.DesignatedSponsors) {
		toSerialize["designatedSponsors"] = o.DesignatedSponsors
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastOrderLogEntry) {
		toSerialize["lastOrderLogEntry"] = o.LastOrderLogEntry
	}
	if !IsNil(o.LastPrice) {
		toSerialize["lastPrice"] = o.LastPrice
	}
	if !IsNil(o.LogoUrl) {
		toSerialize["logoUrl"] = o.LogoUrl
	}
	if !IsNil(o.MarketCap) {
		toSerialize["marketCap"] = o.MarketCap
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OutstandingShares) {
		toSerialize["outstandingShares"] = o.OutstandingShares
	}
	if !IsNil(o.Prices14d) {
		toSerialize["prices14d"] = o.Prices14d
	}
	if !IsNil(o.SecurityIdentifier) {
		toSerialize["securityIdentifier"] = o.SecurityIdentifier
	}
	if !IsNil(o.ShareholderStake) {
		toSerialize["shareholderStake"] = o.ShareholderStake
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.SystemBond) {
		toSerialize["systemBond"] = o.SystemBond
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableListingProfileView struct {
	value *ListingProfileView
	isSet bool
}

func (v NullableListingProfileView) Get() *ListingProfileView {
	return v.value
}

func (v *NullableListingProfileView) Set(val *ListingProfileView) {
	v.value = val
	v.isSet = true
}

func (v NullableListingProfileView) IsSet() bool {
	return v.isSet
}

func (v *NullableListingProfileView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListingProfileView(val *ListingProfileView) *NullableListingProfileView {
	return &NullableListingProfileView{value: val, isSet: true}
}

func (v NullableListingProfileView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListingProfileView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


