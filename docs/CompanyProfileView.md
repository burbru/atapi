# CompanyProfileView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchievementCount** | Pointer to **int64** |  | [optional] 
**AchievementTotal** | Pointer to **int64** |  | [optional] 
**BankAccount** | Pointer to [**BankAccountView**](BankAccountView.md) |  | [optional] 
**Ceo** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**CeoEmploymentAgreement** | Pointer to [**EmploymentAgreementView**](EmploymentAgreementView.md) |  | [optional] 
**CompanyCapabilities** | Pointer to [**CompanyCapabilitiesView**](CompanyCapabilitiesView.md) |  | [optional] 
**CurrentSpread** | Pointer to [**PriceSpreadView**](PriceSpreadView.md) |  | [optional] 
**DesignatedSponsors** | Pointer to [**[]SecuritySponsorshipView**](SecuritySponsorshipView.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IssuedBonds** | Pointer to [**[]BondView**](BondView.md) |  | [optional] 
**LastOrderLogEntry** | Pointer to [**SecurityOrderLogEntryView**](SecurityOrderLogEntryView.md) |  | [optional] 
**LastPrice** | Pointer to [**SecurityPriceView**](SecurityPriceView.md) |  | [optional] 
**LastTrades** | Pointer to [**[]SecurityOrderLogEntryView**](SecurityOrderLogEntryView.md) |  | [optional] 
**Listing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**MarketCap** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OutstandingShares** | Pointer to **int64** |  | [optional] 
**Prices14d** | Pointer to [**[]SecurityPriceView**](SecurityPriceView.md) |  | [optional] 
**SecuritiesAccountId** | Pointer to **string** |  | [optional] 
**SecurityIdentifier** | Pointer to **string** |  | [optional] 
**SponsoredListings** | Pointer to [**[]SecuritySponsorshipView**](SecuritySponsorshipView.md) |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewCompanyProfileView

`func NewCompanyProfileView() *CompanyProfileView`

NewCompanyProfileView instantiates a new CompanyProfileView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyProfileViewWithDefaults

`func NewCompanyProfileViewWithDefaults() *CompanyProfileView`

NewCompanyProfileViewWithDefaults instantiates a new CompanyProfileView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchievementCount

`func (o *CompanyProfileView) GetAchievementCount() int64`

GetAchievementCount returns the AchievementCount field if non-nil, zero value otherwise.

### GetAchievementCountOk

`func (o *CompanyProfileView) GetAchievementCountOk() (*int64, bool)`

GetAchievementCountOk returns a tuple with the AchievementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementCount

`func (o *CompanyProfileView) SetAchievementCount(v int64)`

SetAchievementCount sets AchievementCount field to given value.

### HasAchievementCount

`func (o *CompanyProfileView) HasAchievementCount() bool`

HasAchievementCount returns a boolean if a field has been set.

### GetAchievementTotal

`func (o *CompanyProfileView) GetAchievementTotal() int64`

GetAchievementTotal returns the AchievementTotal field if non-nil, zero value otherwise.

### GetAchievementTotalOk

`func (o *CompanyProfileView) GetAchievementTotalOk() (*int64, bool)`

GetAchievementTotalOk returns a tuple with the AchievementTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementTotal

`func (o *CompanyProfileView) SetAchievementTotal(v int64)`

SetAchievementTotal sets AchievementTotal field to given value.

### HasAchievementTotal

`func (o *CompanyProfileView) HasAchievementTotal() bool`

HasAchievementTotal returns a boolean if a field has been set.

### GetBankAccount

`func (o *CompanyProfileView) GetBankAccount() BankAccountView`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *CompanyProfileView) GetBankAccountOk() (*BankAccountView, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *CompanyProfileView) SetBankAccount(v BankAccountView)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *CompanyProfileView) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### GetCeo

`func (o *CompanyProfileView) GetCeo() UsernameView`

GetCeo returns the Ceo field if non-nil, zero value otherwise.

### GetCeoOk

`func (o *CompanyProfileView) GetCeoOk() (*UsernameView, bool)`

GetCeoOk returns a tuple with the Ceo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeo

`func (o *CompanyProfileView) SetCeo(v UsernameView)`

SetCeo sets Ceo field to given value.

### HasCeo

`func (o *CompanyProfileView) HasCeo() bool`

HasCeo returns a boolean if a field has been set.

### GetCeoEmploymentAgreement

`func (o *CompanyProfileView) GetCeoEmploymentAgreement() EmploymentAgreementView`

GetCeoEmploymentAgreement returns the CeoEmploymentAgreement field if non-nil, zero value otherwise.

### GetCeoEmploymentAgreementOk

`func (o *CompanyProfileView) GetCeoEmploymentAgreementOk() (*EmploymentAgreementView, bool)`

GetCeoEmploymentAgreementOk returns a tuple with the CeoEmploymentAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeoEmploymentAgreement

`func (o *CompanyProfileView) SetCeoEmploymentAgreement(v EmploymentAgreementView)`

SetCeoEmploymentAgreement sets CeoEmploymentAgreement field to given value.

### HasCeoEmploymentAgreement

`func (o *CompanyProfileView) HasCeoEmploymentAgreement() bool`

HasCeoEmploymentAgreement returns a boolean if a field has been set.

### GetCompanyCapabilities

`func (o *CompanyProfileView) GetCompanyCapabilities() CompanyCapabilitiesView`

GetCompanyCapabilities returns the CompanyCapabilities field if non-nil, zero value otherwise.

### GetCompanyCapabilitiesOk

`func (o *CompanyProfileView) GetCompanyCapabilitiesOk() (*CompanyCapabilitiesView, bool)`

GetCompanyCapabilitiesOk returns a tuple with the CompanyCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyCapabilities

`func (o *CompanyProfileView) SetCompanyCapabilities(v CompanyCapabilitiesView)`

SetCompanyCapabilities sets CompanyCapabilities field to given value.

### HasCompanyCapabilities

`func (o *CompanyProfileView) HasCompanyCapabilities() bool`

HasCompanyCapabilities returns a boolean if a field has been set.

### GetCurrentSpread

`func (o *CompanyProfileView) GetCurrentSpread() PriceSpreadView`

GetCurrentSpread returns the CurrentSpread field if non-nil, zero value otherwise.

### GetCurrentSpreadOk

`func (o *CompanyProfileView) GetCurrentSpreadOk() (*PriceSpreadView, bool)`

GetCurrentSpreadOk returns a tuple with the CurrentSpread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSpread

`func (o *CompanyProfileView) SetCurrentSpread(v PriceSpreadView)`

SetCurrentSpread sets CurrentSpread field to given value.

### HasCurrentSpread

`func (o *CompanyProfileView) HasCurrentSpread() bool`

HasCurrentSpread returns a boolean if a field has been set.

### GetDesignatedSponsors

`func (o *CompanyProfileView) GetDesignatedSponsors() []SecuritySponsorshipView`

GetDesignatedSponsors returns the DesignatedSponsors field if non-nil, zero value otherwise.

### GetDesignatedSponsorsOk

`func (o *CompanyProfileView) GetDesignatedSponsorsOk() (*[]SecuritySponsorshipView, bool)`

GetDesignatedSponsorsOk returns a tuple with the DesignatedSponsors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignatedSponsors

`func (o *CompanyProfileView) SetDesignatedSponsors(v []SecuritySponsorshipView)`

SetDesignatedSponsors sets DesignatedSponsors field to given value.

### HasDesignatedSponsors

`func (o *CompanyProfileView) HasDesignatedSponsors() bool`

HasDesignatedSponsors returns a boolean if a field has been set.

### GetId

`func (o *CompanyProfileView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyProfileView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyProfileView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyProfileView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuedBonds

`func (o *CompanyProfileView) GetIssuedBonds() []BondView`

GetIssuedBonds returns the IssuedBonds field if non-nil, zero value otherwise.

### GetIssuedBondsOk

`func (o *CompanyProfileView) GetIssuedBondsOk() (*[]BondView, bool)`

GetIssuedBondsOk returns a tuple with the IssuedBonds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedBonds

`func (o *CompanyProfileView) SetIssuedBonds(v []BondView)`

SetIssuedBonds sets IssuedBonds field to given value.

### HasIssuedBonds

`func (o *CompanyProfileView) HasIssuedBonds() bool`

HasIssuedBonds returns a boolean if a field has been set.

### GetLastOrderLogEntry

`func (o *CompanyProfileView) GetLastOrderLogEntry() SecurityOrderLogEntryView`

GetLastOrderLogEntry returns the LastOrderLogEntry field if non-nil, zero value otherwise.

### GetLastOrderLogEntryOk

`func (o *CompanyProfileView) GetLastOrderLogEntryOk() (*SecurityOrderLogEntryView, bool)`

GetLastOrderLogEntryOk returns a tuple with the LastOrderLogEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOrderLogEntry

`func (o *CompanyProfileView) SetLastOrderLogEntry(v SecurityOrderLogEntryView)`

SetLastOrderLogEntry sets LastOrderLogEntry field to given value.

### HasLastOrderLogEntry

`func (o *CompanyProfileView) HasLastOrderLogEntry() bool`

HasLastOrderLogEntry returns a boolean if a field has been set.

### GetLastPrice

`func (o *CompanyProfileView) GetLastPrice() SecurityPriceView`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *CompanyProfileView) GetLastPriceOk() (*SecurityPriceView, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *CompanyProfileView) SetLastPrice(v SecurityPriceView)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *CompanyProfileView) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLastTrades

`func (o *CompanyProfileView) GetLastTrades() []SecurityOrderLogEntryView`

GetLastTrades returns the LastTrades field if non-nil, zero value otherwise.

### GetLastTradesOk

`func (o *CompanyProfileView) GetLastTradesOk() (*[]SecurityOrderLogEntryView, bool)`

GetLastTradesOk returns a tuple with the LastTrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrades

`func (o *CompanyProfileView) SetLastTrades(v []SecurityOrderLogEntryView)`

SetLastTrades sets LastTrades field to given value.

### HasLastTrades

`func (o *CompanyProfileView) HasLastTrades() bool`

HasLastTrades returns a boolean if a field has been set.

### GetListing

`func (o *CompanyProfileView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *CompanyProfileView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *CompanyProfileView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *CompanyProfileView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetLogoUrl

`func (o *CompanyProfileView) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *CompanyProfileView) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *CompanyProfileView) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *CompanyProfileView) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetMarketCap

`func (o *CompanyProfileView) GetMarketCap() float32`

GetMarketCap returns the MarketCap field if non-nil, zero value otherwise.

### GetMarketCapOk

`func (o *CompanyProfileView) GetMarketCapOk() (*float32, bool)`

GetMarketCapOk returns a tuple with the MarketCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCap

`func (o *CompanyProfileView) SetMarketCap(v float32)`

SetMarketCap sets MarketCap field to given value.

### HasMarketCap

`func (o *CompanyProfileView) HasMarketCap() bool`

HasMarketCap returns a boolean if a field has been set.

### GetName

`func (o *CompanyProfileView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyProfileView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyProfileView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyProfileView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutstandingShares

`func (o *CompanyProfileView) GetOutstandingShares() int64`

GetOutstandingShares returns the OutstandingShares field if non-nil, zero value otherwise.

### GetOutstandingSharesOk

`func (o *CompanyProfileView) GetOutstandingSharesOk() (*int64, bool)`

GetOutstandingSharesOk returns a tuple with the OutstandingShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutstandingShares

`func (o *CompanyProfileView) SetOutstandingShares(v int64)`

SetOutstandingShares sets OutstandingShares field to given value.

### HasOutstandingShares

`func (o *CompanyProfileView) HasOutstandingShares() bool`

HasOutstandingShares returns a boolean if a field has been set.

### GetPrices14d

`func (o *CompanyProfileView) GetPrices14d() []SecurityPriceView`

GetPrices14d returns the Prices14d field if non-nil, zero value otherwise.

### GetPrices14dOk

`func (o *CompanyProfileView) GetPrices14dOk() (*[]SecurityPriceView, bool)`

GetPrices14dOk returns a tuple with the Prices14d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices14d

`func (o *CompanyProfileView) SetPrices14d(v []SecurityPriceView)`

SetPrices14d sets Prices14d field to given value.

### HasPrices14d

`func (o *CompanyProfileView) HasPrices14d() bool`

HasPrices14d returns a boolean if a field has been set.

### GetSecuritiesAccountId

`func (o *CompanyProfileView) GetSecuritiesAccountId() string`

GetSecuritiesAccountId returns the SecuritiesAccountId field if non-nil, zero value otherwise.

### GetSecuritiesAccountIdOk

`func (o *CompanyProfileView) GetSecuritiesAccountIdOk() (*string, bool)`

GetSecuritiesAccountIdOk returns a tuple with the SecuritiesAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritiesAccountId

`func (o *CompanyProfileView) SetSecuritiesAccountId(v string)`

SetSecuritiesAccountId sets SecuritiesAccountId field to given value.

### HasSecuritiesAccountId

`func (o *CompanyProfileView) HasSecuritiesAccountId() bool`

HasSecuritiesAccountId returns a boolean if a field has been set.

### GetSecurityIdentifier

`func (o *CompanyProfileView) GetSecurityIdentifier() string`

GetSecurityIdentifier returns the SecurityIdentifier field if non-nil, zero value otherwise.

### GetSecurityIdentifierOk

`func (o *CompanyProfileView) GetSecurityIdentifierOk() (*string, bool)`

GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIdentifier

`func (o *CompanyProfileView) SetSecurityIdentifier(v string)`

SetSecurityIdentifier sets SecurityIdentifier field to given value.

### HasSecurityIdentifier

`func (o *CompanyProfileView) HasSecurityIdentifier() bool`

HasSecurityIdentifier returns a boolean if a field has been set.

### GetSponsoredListings

`func (o *CompanyProfileView) GetSponsoredListings() []SecuritySponsorshipView`

GetSponsoredListings returns the SponsoredListings field if non-nil, zero value otherwise.

### GetSponsoredListingsOk

`func (o *CompanyProfileView) GetSponsoredListingsOk() (*[]SecuritySponsorshipView, bool)`

GetSponsoredListingsOk returns a tuple with the SponsoredListings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsoredListings

`func (o *CompanyProfileView) SetSponsoredListings(v []SecuritySponsorshipView)`

SetSponsoredListings sets SponsoredListings field to given value.

### HasSponsoredListings

`func (o *CompanyProfileView) HasSponsoredListings() bool`

HasSponsoredListings returns a boolean if a field has been set.

### GetVersion

`func (o *CompanyProfileView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CompanyProfileView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CompanyProfileView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CompanyProfileView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


