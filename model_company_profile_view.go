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

// checks if the CompanyProfileView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyProfileView{}

// CompanyProfileView struct for CompanyProfileView
type CompanyProfileView struct {
	AchievementCount *int64 `json:"achievementCount,omitempty"`
	AchievementTotal *int64 `json:"achievementTotal,omitempty"`
	BankAccount *BankAccountView `json:"bankAccount,omitempty"`
	Ceo *UsernameView `json:"ceo,omitempty"`
	CeoEmploymentAgreement *EmploymentAgreementView `json:"ceoEmploymentAgreement,omitempty"`
	CompanyCapabilities *CompanyCapabilitiesView `json:"companyCapabilities,omitempty"`
	CurrentSpread *PriceSpreadView `json:"currentSpread,omitempty"`
	DesignatedSponsors []SecuritySponsorshipView `json:"designatedSponsors,omitempty"`
	Id *string `json:"id,omitempty"`
	IssuedBonds []BondView `json:"issuedBonds,omitempty"`
	LastOrderLogEntry *SecurityOrderLogEntryView `json:"lastOrderLogEntry,omitempty"`
	LastPrice *SecurityPriceView `json:"lastPrice,omitempty"`
	LastTrades []SecurityOrderLogEntryView `json:"lastTrades,omitempty"`
	Listing *ListingView `json:"listing,omitempty"`
	LogoUrl *string `json:"logoUrl,omitempty"`
	MarketCap *float32 `json:"marketCap,omitempty"`
	Name *string `json:"name,omitempty"`
	OutstandingShares *int64 `json:"outstandingShares,omitempty"`
	Prices14d []SecurityPriceView `json:"prices14d,omitempty"`
	SecuritiesAccountId *string `json:"securitiesAccountId,omitempty"`
	SecurityIdentifier *string `json:"securityIdentifier,omitempty"`
	SponsoredListings []SecuritySponsorshipView `json:"sponsoredListings,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewCompanyProfileView instantiates a new CompanyProfileView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyProfileView() *CompanyProfileView {
	this := CompanyProfileView{}
	return &this
}

// NewCompanyProfileViewWithDefaults instantiates a new CompanyProfileView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyProfileViewWithDefaults() *CompanyProfileView {
	this := CompanyProfileView{}
	return &this
}

// GetAchievementCount returns the AchievementCount field value if set, zero value otherwise.
func (o *CompanyProfileView) GetAchievementCount() int64 {
	if o == nil || IsNil(o.AchievementCount) {
		var ret int64
		return ret
	}
	return *o.AchievementCount
}

// GetAchievementCountOk returns a tuple with the AchievementCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetAchievementCountOk() (*int64, bool) {
	if o == nil || IsNil(o.AchievementCount) {
		return nil, false
	}
	return o.AchievementCount, true
}

// HasAchievementCount returns a boolean if a field has been set.
func (o *CompanyProfileView) HasAchievementCount() bool {
	if o != nil && !IsNil(o.AchievementCount) {
		return true
	}

	return false
}

// SetAchievementCount gets a reference to the given int64 and assigns it to the AchievementCount field.
func (o *CompanyProfileView) SetAchievementCount(v int64) {
	o.AchievementCount = &v
}

// GetAchievementTotal returns the AchievementTotal field value if set, zero value otherwise.
func (o *CompanyProfileView) GetAchievementTotal() int64 {
	if o == nil || IsNil(o.AchievementTotal) {
		var ret int64
		return ret
	}
	return *o.AchievementTotal
}

// GetAchievementTotalOk returns a tuple with the AchievementTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetAchievementTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.AchievementTotal) {
		return nil, false
	}
	return o.AchievementTotal, true
}

// HasAchievementTotal returns a boolean if a field has been set.
func (o *CompanyProfileView) HasAchievementTotal() bool {
	if o != nil && !IsNil(o.AchievementTotal) {
		return true
	}

	return false
}

// SetAchievementTotal gets a reference to the given int64 and assigns it to the AchievementTotal field.
func (o *CompanyProfileView) SetAchievementTotal(v int64) {
	o.AchievementTotal = &v
}

// GetBankAccount returns the BankAccount field value if set, zero value otherwise.
func (o *CompanyProfileView) GetBankAccount() BankAccountView {
	if o == nil || IsNil(o.BankAccount) {
		var ret BankAccountView
		return ret
	}
	return *o.BankAccount
}

// GetBankAccountOk returns a tuple with the BankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetBankAccountOk() (*BankAccountView, bool) {
	if o == nil || IsNil(o.BankAccount) {
		return nil, false
	}
	return o.BankAccount, true
}

// HasBankAccount returns a boolean if a field has been set.
func (o *CompanyProfileView) HasBankAccount() bool {
	if o != nil && !IsNil(o.BankAccount) {
		return true
	}

	return false
}

// SetBankAccount gets a reference to the given BankAccountView and assigns it to the BankAccount field.
func (o *CompanyProfileView) SetBankAccount(v BankAccountView) {
	o.BankAccount = &v
}

// GetCeo returns the Ceo field value if set, zero value otherwise.
func (o *CompanyProfileView) GetCeo() UsernameView {
	if o == nil || IsNil(o.Ceo) {
		var ret UsernameView
		return ret
	}
	return *o.Ceo
}

// GetCeoOk returns a tuple with the Ceo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetCeoOk() (*UsernameView, bool) {
	if o == nil || IsNil(o.Ceo) {
		return nil, false
	}
	return o.Ceo, true
}

// HasCeo returns a boolean if a field has been set.
func (o *CompanyProfileView) HasCeo() bool {
	if o != nil && !IsNil(o.Ceo) {
		return true
	}

	return false
}

// SetCeo gets a reference to the given UsernameView and assigns it to the Ceo field.
func (o *CompanyProfileView) SetCeo(v UsernameView) {
	o.Ceo = &v
}

// GetCeoEmploymentAgreement returns the CeoEmploymentAgreement field value if set, zero value otherwise.
func (o *CompanyProfileView) GetCeoEmploymentAgreement() EmploymentAgreementView {
	if o == nil || IsNil(o.CeoEmploymentAgreement) {
		var ret EmploymentAgreementView
		return ret
	}
	return *o.CeoEmploymentAgreement
}

// GetCeoEmploymentAgreementOk returns a tuple with the CeoEmploymentAgreement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetCeoEmploymentAgreementOk() (*EmploymentAgreementView, bool) {
	if o == nil || IsNil(o.CeoEmploymentAgreement) {
		return nil, false
	}
	return o.CeoEmploymentAgreement, true
}

// HasCeoEmploymentAgreement returns a boolean if a field has been set.
func (o *CompanyProfileView) HasCeoEmploymentAgreement() bool {
	if o != nil && !IsNil(o.CeoEmploymentAgreement) {
		return true
	}

	return false
}

// SetCeoEmploymentAgreement gets a reference to the given EmploymentAgreementView and assigns it to the CeoEmploymentAgreement field.
func (o *CompanyProfileView) SetCeoEmploymentAgreement(v EmploymentAgreementView) {
	o.CeoEmploymentAgreement = &v
}

// GetCompanyCapabilities returns the CompanyCapabilities field value if set, zero value otherwise.
func (o *CompanyProfileView) GetCompanyCapabilities() CompanyCapabilitiesView {
	if o == nil || IsNil(o.CompanyCapabilities) {
		var ret CompanyCapabilitiesView
		return ret
	}
	return *o.CompanyCapabilities
}

// GetCompanyCapabilitiesOk returns a tuple with the CompanyCapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetCompanyCapabilitiesOk() (*CompanyCapabilitiesView, bool) {
	if o == nil || IsNil(o.CompanyCapabilities) {
		return nil, false
	}
	return o.CompanyCapabilities, true
}

// HasCompanyCapabilities returns a boolean if a field has been set.
func (o *CompanyProfileView) HasCompanyCapabilities() bool {
	if o != nil && !IsNil(o.CompanyCapabilities) {
		return true
	}

	return false
}

// SetCompanyCapabilities gets a reference to the given CompanyCapabilitiesView and assigns it to the CompanyCapabilities field.
func (o *CompanyProfileView) SetCompanyCapabilities(v CompanyCapabilitiesView) {
	o.CompanyCapabilities = &v
}

// GetCurrentSpread returns the CurrentSpread field value if set, zero value otherwise.
func (o *CompanyProfileView) GetCurrentSpread() PriceSpreadView {
	if o == nil || IsNil(o.CurrentSpread) {
		var ret PriceSpreadView
		return ret
	}
	return *o.CurrentSpread
}

// GetCurrentSpreadOk returns a tuple with the CurrentSpread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetCurrentSpreadOk() (*PriceSpreadView, bool) {
	if o == nil || IsNil(o.CurrentSpread) {
		return nil, false
	}
	return o.CurrentSpread, true
}

// HasCurrentSpread returns a boolean if a field has been set.
func (o *CompanyProfileView) HasCurrentSpread() bool {
	if o != nil && !IsNil(o.CurrentSpread) {
		return true
	}

	return false
}

// SetCurrentSpread gets a reference to the given PriceSpreadView and assigns it to the CurrentSpread field.
func (o *CompanyProfileView) SetCurrentSpread(v PriceSpreadView) {
	o.CurrentSpread = &v
}

// GetDesignatedSponsors returns the DesignatedSponsors field value if set, zero value otherwise.
func (o *CompanyProfileView) GetDesignatedSponsors() []SecuritySponsorshipView {
	if o == nil || IsNil(o.DesignatedSponsors) {
		var ret []SecuritySponsorshipView
		return ret
	}
	return o.DesignatedSponsors
}

// GetDesignatedSponsorsOk returns a tuple with the DesignatedSponsors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetDesignatedSponsorsOk() ([]SecuritySponsorshipView, bool) {
	if o == nil || IsNil(o.DesignatedSponsors) {
		return nil, false
	}
	return o.DesignatedSponsors, true
}

// HasDesignatedSponsors returns a boolean if a field has been set.
func (o *CompanyProfileView) HasDesignatedSponsors() bool {
	if o != nil && !IsNil(o.DesignatedSponsors) {
		return true
	}

	return false
}

// SetDesignatedSponsors gets a reference to the given []SecuritySponsorshipView and assigns it to the DesignatedSponsors field.
func (o *CompanyProfileView) SetDesignatedSponsors(v []SecuritySponsorshipView) {
	o.DesignatedSponsors = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CompanyProfileView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CompanyProfileView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CompanyProfileView) SetId(v string) {
	o.Id = &v
}

// GetIssuedBonds returns the IssuedBonds field value if set, zero value otherwise.
func (o *CompanyProfileView) GetIssuedBonds() []BondView {
	if o == nil || IsNil(o.IssuedBonds) {
		var ret []BondView
		return ret
	}
	return o.IssuedBonds
}

// GetIssuedBondsOk returns a tuple with the IssuedBonds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetIssuedBondsOk() ([]BondView, bool) {
	if o == nil || IsNil(o.IssuedBonds) {
		return nil, false
	}
	return o.IssuedBonds, true
}

// HasIssuedBonds returns a boolean if a field has been set.
func (o *CompanyProfileView) HasIssuedBonds() bool {
	if o != nil && !IsNil(o.IssuedBonds) {
		return true
	}

	return false
}

// SetIssuedBonds gets a reference to the given []BondView and assigns it to the IssuedBonds field.
func (o *CompanyProfileView) SetIssuedBonds(v []BondView) {
	o.IssuedBonds = v
}

// GetLastOrderLogEntry returns the LastOrderLogEntry field value if set, zero value otherwise.
func (o *CompanyProfileView) GetLastOrderLogEntry() SecurityOrderLogEntryView {
	if o == nil || IsNil(o.LastOrderLogEntry) {
		var ret SecurityOrderLogEntryView
		return ret
	}
	return *o.LastOrderLogEntry
}

// GetLastOrderLogEntryOk returns a tuple with the LastOrderLogEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetLastOrderLogEntryOk() (*SecurityOrderLogEntryView, bool) {
	if o == nil || IsNil(o.LastOrderLogEntry) {
		return nil, false
	}
	return o.LastOrderLogEntry, true
}

// HasLastOrderLogEntry returns a boolean if a field has been set.
func (o *CompanyProfileView) HasLastOrderLogEntry() bool {
	if o != nil && !IsNil(o.LastOrderLogEntry) {
		return true
	}

	return false
}

// SetLastOrderLogEntry gets a reference to the given SecurityOrderLogEntryView and assigns it to the LastOrderLogEntry field.
func (o *CompanyProfileView) SetLastOrderLogEntry(v SecurityOrderLogEntryView) {
	o.LastOrderLogEntry = &v
}

// GetLastPrice returns the LastPrice field value if set, zero value otherwise.
func (o *CompanyProfileView) GetLastPrice() SecurityPriceView {
	if o == nil || IsNil(o.LastPrice) {
		var ret SecurityPriceView
		return ret
	}
	return *o.LastPrice
}

// GetLastPriceOk returns a tuple with the LastPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetLastPriceOk() (*SecurityPriceView, bool) {
	if o == nil || IsNil(o.LastPrice) {
		return nil, false
	}
	return o.LastPrice, true
}

// HasLastPrice returns a boolean if a field has been set.
func (o *CompanyProfileView) HasLastPrice() bool {
	if o != nil && !IsNil(o.LastPrice) {
		return true
	}

	return false
}

// SetLastPrice gets a reference to the given SecurityPriceView and assigns it to the LastPrice field.
func (o *CompanyProfileView) SetLastPrice(v SecurityPriceView) {
	o.LastPrice = &v
}

// GetLastTrades returns the LastTrades field value if set, zero value otherwise.
func (o *CompanyProfileView) GetLastTrades() []SecurityOrderLogEntryView {
	if o == nil || IsNil(o.LastTrades) {
		var ret []SecurityOrderLogEntryView
		return ret
	}
	return o.LastTrades
}

// GetLastTradesOk returns a tuple with the LastTrades field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetLastTradesOk() ([]SecurityOrderLogEntryView, bool) {
	if o == nil || IsNil(o.LastTrades) {
		return nil, false
	}
	return o.LastTrades, true
}

// HasLastTrades returns a boolean if a field has been set.
func (o *CompanyProfileView) HasLastTrades() bool {
	if o != nil && !IsNil(o.LastTrades) {
		return true
	}

	return false
}

// SetLastTrades gets a reference to the given []SecurityOrderLogEntryView and assigns it to the LastTrades field.
func (o *CompanyProfileView) SetLastTrades(v []SecurityOrderLogEntryView) {
	o.LastTrades = v
}

// GetListing returns the Listing field value if set, zero value otherwise.
func (o *CompanyProfileView) GetListing() ListingView {
	if o == nil || IsNil(o.Listing) {
		var ret ListingView
		return ret
	}
	return *o.Listing
}

// GetListingOk returns a tuple with the Listing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetListingOk() (*ListingView, bool) {
	if o == nil || IsNil(o.Listing) {
		return nil, false
	}
	return o.Listing, true
}

// HasListing returns a boolean if a field has been set.
func (o *CompanyProfileView) HasListing() bool {
	if o != nil && !IsNil(o.Listing) {
		return true
	}

	return false
}

// SetListing gets a reference to the given ListingView and assigns it to the Listing field.
func (o *CompanyProfileView) SetListing(v ListingView) {
	o.Listing = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *CompanyProfileView) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *CompanyProfileView) HasLogoUrl() bool {
	if o != nil && !IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *CompanyProfileView) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetMarketCap returns the MarketCap field value if set, zero value otherwise.
func (o *CompanyProfileView) GetMarketCap() float32 {
	if o == nil || IsNil(o.MarketCap) {
		var ret float32
		return ret
	}
	return *o.MarketCap
}

// GetMarketCapOk returns a tuple with the MarketCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetMarketCapOk() (*float32, bool) {
	if o == nil || IsNil(o.MarketCap) {
		return nil, false
	}
	return o.MarketCap, true
}

// HasMarketCap returns a boolean if a field has been set.
func (o *CompanyProfileView) HasMarketCap() bool {
	if o != nil && !IsNil(o.MarketCap) {
		return true
	}

	return false
}

// SetMarketCap gets a reference to the given float32 and assigns it to the MarketCap field.
func (o *CompanyProfileView) SetMarketCap(v float32) {
	o.MarketCap = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CompanyProfileView) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CompanyProfileView) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CompanyProfileView) SetName(v string) {
	o.Name = &v
}

// GetOutstandingShares returns the OutstandingShares field value if set, zero value otherwise.
func (o *CompanyProfileView) GetOutstandingShares() int64 {
	if o == nil || IsNil(o.OutstandingShares) {
		var ret int64
		return ret
	}
	return *o.OutstandingShares
}

// GetOutstandingSharesOk returns a tuple with the OutstandingShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetOutstandingSharesOk() (*int64, bool) {
	if o == nil || IsNil(o.OutstandingShares) {
		return nil, false
	}
	return o.OutstandingShares, true
}

// HasOutstandingShares returns a boolean if a field has been set.
func (o *CompanyProfileView) HasOutstandingShares() bool {
	if o != nil && !IsNil(o.OutstandingShares) {
		return true
	}

	return false
}

// SetOutstandingShares gets a reference to the given int64 and assigns it to the OutstandingShares field.
func (o *CompanyProfileView) SetOutstandingShares(v int64) {
	o.OutstandingShares = &v
}

// GetPrices14d returns the Prices14d field value if set, zero value otherwise.
func (o *CompanyProfileView) GetPrices14d() []SecurityPriceView {
	if o == nil || IsNil(o.Prices14d) {
		var ret []SecurityPriceView
		return ret
	}
	return o.Prices14d
}

// GetPrices14dOk returns a tuple with the Prices14d field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetPrices14dOk() ([]SecurityPriceView, bool) {
	if o == nil || IsNil(o.Prices14d) {
		return nil, false
	}
	return o.Prices14d, true
}

// HasPrices14d returns a boolean if a field has been set.
func (o *CompanyProfileView) HasPrices14d() bool {
	if o != nil && !IsNil(o.Prices14d) {
		return true
	}

	return false
}

// SetPrices14d gets a reference to the given []SecurityPriceView and assigns it to the Prices14d field.
func (o *CompanyProfileView) SetPrices14d(v []SecurityPriceView) {
	o.Prices14d = v
}

// GetSecuritiesAccountId returns the SecuritiesAccountId field value if set, zero value otherwise.
func (o *CompanyProfileView) GetSecuritiesAccountId() string {
	if o == nil || IsNil(o.SecuritiesAccountId) {
		var ret string
		return ret
	}
	return *o.SecuritiesAccountId
}

// GetSecuritiesAccountIdOk returns a tuple with the SecuritiesAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetSecuritiesAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecuritiesAccountId) {
		return nil, false
	}
	return o.SecuritiesAccountId, true
}

// HasSecuritiesAccountId returns a boolean if a field has been set.
func (o *CompanyProfileView) HasSecuritiesAccountId() bool {
	if o != nil && !IsNil(o.SecuritiesAccountId) {
		return true
	}

	return false
}

// SetSecuritiesAccountId gets a reference to the given string and assigns it to the SecuritiesAccountId field.
func (o *CompanyProfileView) SetSecuritiesAccountId(v string) {
	o.SecuritiesAccountId = &v
}

// GetSecurityIdentifier returns the SecurityIdentifier field value if set, zero value otherwise.
func (o *CompanyProfileView) GetSecurityIdentifier() string {
	if o == nil || IsNil(o.SecurityIdentifier) {
		var ret string
		return ret
	}
	return *o.SecurityIdentifier
}

// GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetSecurityIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityIdentifier) {
		return nil, false
	}
	return o.SecurityIdentifier, true
}

// HasSecurityIdentifier returns a boolean if a field has been set.
func (o *CompanyProfileView) HasSecurityIdentifier() bool {
	if o != nil && !IsNil(o.SecurityIdentifier) {
		return true
	}

	return false
}

// SetSecurityIdentifier gets a reference to the given string and assigns it to the SecurityIdentifier field.
func (o *CompanyProfileView) SetSecurityIdentifier(v string) {
	o.SecurityIdentifier = &v
}

// GetSponsoredListings returns the SponsoredListings field value if set, zero value otherwise.
func (o *CompanyProfileView) GetSponsoredListings() []SecuritySponsorshipView {
	if o == nil || IsNil(o.SponsoredListings) {
		var ret []SecuritySponsorshipView
		return ret
	}
	return o.SponsoredListings
}

// GetSponsoredListingsOk returns a tuple with the SponsoredListings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetSponsoredListingsOk() ([]SecuritySponsorshipView, bool) {
	if o == nil || IsNil(o.SponsoredListings) {
		return nil, false
	}
	return o.SponsoredListings, true
}

// HasSponsoredListings returns a boolean if a field has been set.
func (o *CompanyProfileView) HasSponsoredListings() bool {
	if o != nil && !IsNil(o.SponsoredListings) {
		return true
	}

	return false
}

// SetSponsoredListings gets a reference to the given []SecuritySponsorshipView and assigns it to the SponsoredListings field.
func (o *CompanyProfileView) SetSponsoredListings(v []SecuritySponsorshipView) {
	o.SponsoredListings = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CompanyProfileView) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileView) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CompanyProfileView) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *CompanyProfileView) SetVersion(v int64) {
	o.Version = &v
}

func (o CompanyProfileView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyProfileView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AchievementCount) {
		toSerialize["achievementCount"] = o.AchievementCount
	}
	if !IsNil(o.AchievementTotal) {
		toSerialize["achievementTotal"] = o.AchievementTotal
	}
	if !IsNil(o.BankAccount) {
		toSerialize["bankAccount"] = o.BankAccount
	}
	if !IsNil(o.Ceo) {
		toSerialize["ceo"] = o.Ceo
	}
	if !IsNil(o.CeoEmploymentAgreement) {
		toSerialize["ceoEmploymentAgreement"] = o.CeoEmploymentAgreement
	}
	if !IsNil(o.CompanyCapabilities) {
		toSerialize["companyCapabilities"] = o.CompanyCapabilities
	}
	if !IsNil(o.CurrentSpread) {
		toSerialize["currentSpread"] = o.CurrentSpread
	}
	if !IsNil(o.DesignatedSponsors) {
		toSerialize["designatedSponsors"] = o.DesignatedSponsors
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IssuedBonds) {
		toSerialize["issuedBonds"] = o.IssuedBonds
	}
	if !IsNil(o.LastOrderLogEntry) {
		toSerialize["lastOrderLogEntry"] = o.LastOrderLogEntry
	}
	if !IsNil(o.LastPrice) {
		toSerialize["lastPrice"] = o.LastPrice
	}
	if !IsNil(o.LastTrades) {
		toSerialize["lastTrades"] = o.LastTrades
	}
	if !IsNil(o.Listing) {
		toSerialize["listing"] = o.Listing
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
	if !IsNil(o.SecuritiesAccountId) {
		toSerialize["securitiesAccountId"] = o.SecuritiesAccountId
	}
	if !IsNil(o.SecurityIdentifier) {
		toSerialize["securityIdentifier"] = o.SecurityIdentifier
	}
	if !IsNil(o.SponsoredListings) {
		toSerialize["sponsoredListings"] = o.SponsoredListings
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableCompanyProfileView struct {
	value *CompanyProfileView
	isSet bool
}

func (v NullableCompanyProfileView) Get() *CompanyProfileView {
	return v.value
}

func (v *NullableCompanyProfileView) Set(val *CompanyProfileView) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyProfileView) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyProfileView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyProfileView(val *CompanyProfileView) *NullableCompanyProfileView {
	return &NullableCompanyProfileView{value: val, isSet: true}
}

func (v NullableCompanyProfileView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyProfileView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


