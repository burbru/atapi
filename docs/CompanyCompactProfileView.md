# CompanyCompactProfileView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchievementCount** | Pointer to **int64** |  | [optional] 
**AchievementTotal** | Pointer to **int64** |  | [optional] 
**BankAccount** | Pointer to [**BankAccountView**](BankAccountView.md) |  | [optional] 
**Ceo** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**CeoEmploymentAgreement** | Pointer to [**EmploymentAgreementView**](EmploymentAgreementView.md) |  | [optional] 
**CompanyCapabilities** | Pointer to [**CompanyCapabilitiesView**](CompanyCapabilitiesView.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IssuedBonds** | Pointer to [**[]BondView**](BondView.md) |  | [optional] 
**LastTrades** | Pointer to [**[]SecurityOrderLogEntryView**](SecurityOrderLogEntryView.md) |  | [optional] 
**Listing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SecuritiesAccountId** | Pointer to **string** |  | [optional] 
**SecurityIdentifier** | Pointer to **string** |  | [optional] 
**SponsoredListings** | Pointer to [**[]SecuritySponsorshipView**](SecuritySponsorshipView.md) |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewCompanyCompactProfileView

`func NewCompanyCompactProfileView() *CompanyCompactProfileView`

NewCompanyCompactProfileView instantiates a new CompanyCompactProfileView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyCompactProfileViewWithDefaults

`func NewCompanyCompactProfileViewWithDefaults() *CompanyCompactProfileView`

NewCompanyCompactProfileViewWithDefaults instantiates a new CompanyCompactProfileView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchievementCount

`func (o *CompanyCompactProfileView) GetAchievementCount() int64`

GetAchievementCount returns the AchievementCount field if non-nil, zero value otherwise.

### GetAchievementCountOk

`func (o *CompanyCompactProfileView) GetAchievementCountOk() (*int64, bool)`

GetAchievementCountOk returns a tuple with the AchievementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementCount

`func (o *CompanyCompactProfileView) SetAchievementCount(v int64)`

SetAchievementCount sets AchievementCount field to given value.

### HasAchievementCount

`func (o *CompanyCompactProfileView) HasAchievementCount() bool`

HasAchievementCount returns a boolean if a field has been set.

### GetAchievementTotal

`func (o *CompanyCompactProfileView) GetAchievementTotal() int64`

GetAchievementTotal returns the AchievementTotal field if non-nil, zero value otherwise.

### GetAchievementTotalOk

`func (o *CompanyCompactProfileView) GetAchievementTotalOk() (*int64, bool)`

GetAchievementTotalOk returns a tuple with the AchievementTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementTotal

`func (o *CompanyCompactProfileView) SetAchievementTotal(v int64)`

SetAchievementTotal sets AchievementTotal field to given value.

### HasAchievementTotal

`func (o *CompanyCompactProfileView) HasAchievementTotal() bool`

HasAchievementTotal returns a boolean if a field has been set.

### GetBankAccount

`func (o *CompanyCompactProfileView) GetBankAccount() BankAccountView`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *CompanyCompactProfileView) GetBankAccountOk() (*BankAccountView, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *CompanyCompactProfileView) SetBankAccount(v BankAccountView)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *CompanyCompactProfileView) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### GetCeo

`func (o *CompanyCompactProfileView) GetCeo() UsernameView`

GetCeo returns the Ceo field if non-nil, zero value otherwise.

### GetCeoOk

`func (o *CompanyCompactProfileView) GetCeoOk() (*UsernameView, bool)`

GetCeoOk returns a tuple with the Ceo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeo

`func (o *CompanyCompactProfileView) SetCeo(v UsernameView)`

SetCeo sets Ceo field to given value.

### HasCeo

`func (o *CompanyCompactProfileView) HasCeo() bool`

HasCeo returns a boolean if a field has been set.

### GetCeoEmploymentAgreement

`func (o *CompanyCompactProfileView) GetCeoEmploymentAgreement() EmploymentAgreementView`

GetCeoEmploymentAgreement returns the CeoEmploymentAgreement field if non-nil, zero value otherwise.

### GetCeoEmploymentAgreementOk

`func (o *CompanyCompactProfileView) GetCeoEmploymentAgreementOk() (*EmploymentAgreementView, bool)`

GetCeoEmploymentAgreementOk returns a tuple with the CeoEmploymentAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeoEmploymentAgreement

`func (o *CompanyCompactProfileView) SetCeoEmploymentAgreement(v EmploymentAgreementView)`

SetCeoEmploymentAgreement sets CeoEmploymentAgreement field to given value.

### HasCeoEmploymentAgreement

`func (o *CompanyCompactProfileView) HasCeoEmploymentAgreement() bool`

HasCeoEmploymentAgreement returns a boolean if a field has been set.

### GetCompanyCapabilities

`func (o *CompanyCompactProfileView) GetCompanyCapabilities() CompanyCapabilitiesView`

GetCompanyCapabilities returns the CompanyCapabilities field if non-nil, zero value otherwise.

### GetCompanyCapabilitiesOk

`func (o *CompanyCompactProfileView) GetCompanyCapabilitiesOk() (*CompanyCapabilitiesView, bool)`

GetCompanyCapabilitiesOk returns a tuple with the CompanyCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyCapabilities

`func (o *CompanyCompactProfileView) SetCompanyCapabilities(v CompanyCapabilitiesView)`

SetCompanyCapabilities sets CompanyCapabilities field to given value.

### HasCompanyCapabilities

`func (o *CompanyCompactProfileView) HasCompanyCapabilities() bool`

HasCompanyCapabilities returns a boolean if a field has been set.

### GetId

`func (o *CompanyCompactProfileView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyCompactProfileView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyCompactProfileView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyCompactProfileView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuedBonds

`func (o *CompanyCompactProfileView) GetIssuedBonds() []BondView`

GetIssuedBonds returns the IssuedBonds field if non-nil, zero value otherwise.

### GetIssuedBondsOk

`func (o *CompanyCompactProfileView) GetIssuedBondsOk() (*[]BondView, bool)`

GetIssuedBondsOk returns a tuple with the IssuedBonds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedBonds

`func (o *CompanyCompactProfileView) SetIssuedBonds(v []BondView)`

SetIssuedBonds sets IssuedBonds field to given value.

### HasIssuedBonds

`func (o *CompanyCompactProfileView) HasIssuedBonds() bool`

HasIssuedBonds returns a boolean if a field has been set.

### GetLastTrades

`func (o *CompanyCompactProfileView) GetLastTrades() []SecurityOrderLogEntryView`

GetLastTrades returns the LastTrades field if non-nil, zero value otherwise.

### GetLastTradesOk

`func (o *CompanyCompactProfileView) GetLastTradesOk() (*[]SecurityOrderLogEntryView, bool)`

GetLastTradesOk returns a tuple with the LastTrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrades

`func (o *CompanyCompactProfileView) SetLastTrades(v []SecurityOrderLogEntryView)`

SetLastTrades sets LastTrades field to given value.

### HasLastTrades

`func (o *CompanyCompactProfileView) HasLastTrades() bool`

HasLastTrades returns a boolean if a field has been set.

### GetListing

`func (o *CompanyCompactProfileView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *CompanyCompactProfileView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *CompanyCompactProfileView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *CompanyCompactProfileView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetLogoUrl

`func (o *CompanyCompactProfileView) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *CompanyCompactProfileView) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *CompanyCompactProfileView) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *CompanyCompactProfileView) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetName

`func (o *CompanyCompactProfileView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyCompactProfileView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyCompactProfileView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyCompactProfileView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecuritiesAccountId

`func (o *CompanyCompactProfileView) GetSecuritiesAccountId() string`

GetSecuritiesAccountId returns the SecuritiesAccountId field if non-nil, zero value otherwise.

### GetSecuritiesAccountIdOk

`func (o *CompanyCompactProfileView) GetSecuritiesAccountIdOk() (*string, bool)`

GetSecuritiesAccountIdOk returns a tuple with the SecuritiesAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritiesAccountId

`func (o *CompanyCompactProfileView) SetSecuritiesAccountId(v string)`

SetSecuritiesAccountId sets SecuritiesAccountId field to given value.

### HasSecuritiesAccountId

`func (o *CompanyCompactProfileView) HasSecuritiesAccountId() bool`

HasSecuritiesAccountId returns a boolean if a field has been set.

### GetSecurityIdentifier

`func (o *CompanyCompactProfileView) GetSecurityIdentifier() string`

GetSecurityIdentifier returns the SecurityIdentifier field if non-nil, zero value otherwise.

### GetSecurityIdentifierOk

`func (o *CompanyCompactProfileView) GetSecurityIdentifierOk() (*string, bool)`

GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIdentifier

`func (o *CompanyCompactProfileView) SetSecurityIdentifier(v string)`

SetSecurityIdentifier sets SecurityIdentifier field to given value.

### HasSecurityIdentifier

`func (o *CompanyCompactProfileView) HasSecurityIdentifier() bool`

HasSecurityIdentifier returns a boolean if a field has been set.

### GetSponsoredListings

`func (o *CompanyCompactProfileView) GetSponsoredListings() []SecuritySponsorshipView`

GetSponsoredListings returns the SponsoredListings field if non-nil, zero value otherwise.

### GetSponsoredListingsOk

`func (o *CompanyCompactProfileView) GetSponsoredListingsOk() (*[]SecuritySponsorshipView, bool)`

GetSponsoredListingsOk returns a tuple with the SponsoredListings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsoredListings

`func (o *CompanyCompactProfileView) SetSponsoredListings(v []SecuritySponsorshipView)`

SetSponsoredListings sets SponsoredListings field to given value.

### HasSponsoredListings

`func (o *CompanyCompactProfileView) HasSponsoredListings() bool`

HasSponsoredListings returns a boolean if a field has been set.

### GetVersion

`func (o *CompanyCompactProfileView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CompanyCompactProfileView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CompanyCompactProfileView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CompanyCompactProfileView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


