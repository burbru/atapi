# ListingProfileView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bond** | Pointer to [**BondView**](BondView.md) |  | [optional] 
**Building** | Pointer to [**BuildingView**](BuildingView.md) |  | [optional] 
**Company** | Pointer to [**CompanyCompactProfileView**](CompanyCompactProfileView.md) |  | [optional] 
**CurrentSpread** | Pointer to [**PriceSpreadView**](PriceSpreadView.md) |  | [optional] 
**DesignatedSponsors** | Pointer to [**[]SecuritySponsorshipView**](SecuritySponsorshipView.md) |  | [optional] 
**EndDate** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastOrderLogEntry** | Pointer to [**SecurityOrderLogEntryView**](SecurityOrderLogEntryView.md) |  | [optional] 
**LastPrice** | Pointer to [**SecurityPriceView**](SecurityPriceView.md) |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**MarketCap** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OutstandingShares** | Pointer to **int64** |  | [optional] 
**Prices14d** | Pointer to [**[]SecurityPriceView**](SecurityPriceView.md) |  | [optional] 
**SecurityIdentifier** | Pointer to **string** |  | [optional] 
**ShareholderStake** | Pointer to [**ShareholderStake**](ShareholderStake.md) |  | [optional] 
**StartDate** | Pointer to **int64** |  | [optional] 
**SystemBond** | Pointer to [**SystemBondView**](SystemBondView.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewListingProfileView

`func NewListingProfileView() *ListingProfileView`

NewListingProfileView instantiates a new ListingProfileView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingProfileViewWithDefaults

`func NewListingProfileViewWithDefaults() *ListingProfileView`

NewListingProfileViewWithDefaults instantiates a new ListingProfileView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBond

`func (o *ListingProfileView) GetBond() BondView`

GetBond returns the Bond field if non-nil, zero value otherwise.

### GetBondOk

`func (o *ListingProfileView) GetBondOk() (*BondView, bool)`

GetBondOk returns a tuple with the Bond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBond

`func (o *ListingProfileView) SetBond(v BondView)`

SetBond sets Bond field to given value.

### HasBond

`func (o *ListingProfileView) HasBond() bool`

HasBond returns a boolean if a field has been set.

### GetBuilding

`func (o *ListingProfileView) GetBuilding() BuildingView`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *ListingProfileView) GetBuildingOk() (*BuildingView, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *ListingProfileView) SetBuilding(v BuildingView)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *ListingProfileView) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### GetCompany

`func (o *ListingProfileView) GetCompany() CompanyCompactProfileView`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ListingProfileView) GetCompanyOk() (*CompanyCompactProfileView, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ListingProfileView) SetCompany(v CompanyCompactProfileView)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ListingProfileView) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCurrentSpread

`func (o *ListingProfileView) GetCurrentSpread() PriceSpreadView`

GetCurrentSpread returns the CurrentSpread field if non-nil, zero value otherwise.

### GetCurrentSpreadOk

`func (o *ListingProfileView) GetCurrentSpreadOk() (*PriceSpreadView, bool)`

GetCurrentSpreadOk returns a tuple with the CurrentSpread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSpread

`func (o *ListingProfileView) SetCurrentSpread(v PriceSpreadView)`

SetCurrentSpread sets CurrentSpread field to given value.

### HasCurrentSpread

`func (o *ListingProfileView) HasCurrentSpread() bool`

HasCurrentSpread returns a boolean if a field has been set.

### GetDesignatedSponsors

`func (o *ListingProfileView) GetDesignatedSponsors() []SecuritySponsorshipView`

GetDesignatedSponsors returns the DesignatedSponsors field if non-nil, zero value otherwise.

### GetDesignatedSponsorsOk

`func (o *ListingProfileView) GetDesignatedSponsorsOk() (*[]SecuritySponsorshipView, bool)`

GetDesignatedSponsorsOk returns a tuple with the DesignatedSponsors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignatedSponsors

`func (o *ListingProfileView) SetDesignatedSponsors(v []SecuritySponsorshipView)`

SetDesignatedSponsors sets DesignatedSponsors field to given value.

### HasDesignatedSponsors

`func (o *ListingProfileView) HasDesignatedSponsors() bool`

HasDesignatedSponsors returns a boolean if a field has been set.

### GetEndDate

`func (o *ListingProfileView) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListingProfileView) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListingProfileView) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListingProfileView) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *ListingProfileView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListingProfileView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListingProfileView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListingProfileView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastOrderLogEntry

`func (o *ListingProfileView) GetLastOrderLogEntry() SecurityOrderLogEntryView`

GetLastOrderLogEntry returns the LastOrderLogEntry field if non-nil, zero value otherwise.

### GetLastOrderLogEntryOk

`func (o *ListingProfileView) GetLastOrderLogEntryOk() (*SecurityOrderLogEntryView, bool)`

GetLastOrderLogEntryOk returns a tuple with the LastOrderLogEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOrderLogEntry

`func (o *ListingProfileView) SetLastOrderLogEntry(v SecurityOrderLogEntryView)`

SetLastOrderLogEntry sets LastOrderLogEntry field to given value.

### HasLastOrderLogEntry

`func (o *ListingProfileView) HasLastOrderLogEntry() bool`

HasLastOrderLogEntry returns a boolean if a field has been set.

### GetLastPrice

`func (o *ListingProfileView) GetLastPrice() SecurityPriceView`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *ListingProfileView) GetLastPriceOk() (*SecurityPriceView, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *ListingProfileView) SetLastPrice(v SecurityPriceView)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *ListingProfileView) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLogoUrl

`func (o *ListingProfileView) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *ListingProfileView) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *ListingProfileView) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *ListingProfileView) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetMarketCap

`func (o *ListingProfileView) GetMarketCap() float32`

GetMarketCap returns the MarketCap field if non-nil, zero value otherwise.

### GetMarketCapOk

`func (o *ListingProfileView) GetMarketCapOk() (*float32, bool)`

GetMarketCapOk returns a tuple with the MarketCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCap

`func (o *ListingProfileView) SetMarketCap(v float32)`

SetMarketCap sets MarketCap field to given value.

### HasMarketCap

`func (o *ListingProfileView) HasMarketCap() bool`

HasMarketCap returns a boolean if a field has been set.

### GetName

`func (o *ListingProfileView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListingProfileView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListingProfileView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListingProfileView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutstandingShares

`func (o *ListingProfileView) GetOutstandingShares() int64`

GetOutstandingShares returns the OutstandingShares field if non-nil, zero value otherwise.

### GetOutstandingSharesOk

`func (o *ListingProfileView) GetOutstandingSharesOk() (*int64, bool)`

GetOutstandingSharesOk returns a tuple with the OutstandingShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutstandingShares

`func (o *ListingProfileView) SetOutstandingShares(v int64)`

SetOutstandingShares sets OutstandingShares field to given value.

### HasOutstandingShares

`func (o *ListingProfileView) HasOutstandingShares() bool`

HasOutstandingShares returns a boolean if a field has been set.

### GetPrices14d

`func (o *ListingProfileView) GetPrices14d() []SecurityPriceView`

GetPrices14d returns the Prices14d field if non-nil, zero value otherwise.

### GetPrices14dOk

`func (o *ListingProfileView) GetPrices14dOk() (*[]SecurityPriceView, bool)`

GetPrices14dOk returns a tuple with the Prices14d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices14d

`func (o *ListingProfileView) SetPrices14d(v []SecurityPriceView)`

SetPrices14d sets Prices14d field to given value.

### HasPrices14d

`func (o *ListingProfileView) HasPrices14d() bool`

HasPrices14d returns a boolean if a field has been set.

### GetSecurityIdentifier

`func (o *ListingProfileView) GetSecurityIdentifier() string`

GetSecurityIdentifier returns the SecurityIdentifier field if non-nil, zero value otherwise.

### GetSecurityIdentifierOk

`func (o *ListingProfileView) GetSecurityIdentifierOk() (*string, bool)`

GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIdentifier

`func (o *ListingProfileView) SetSecurityIdentifier(v string)`

SetSecurityIdentifier sets SecurityIdentifier field to given value.

### HasSecurityIdentifier

`func (o *ListingProfileView) HasSecurityIdentifier() bool`

HasSecurityIdentifier returns a boolean if a field has been set.

### GetShareholderStake

`func (o *ListingProfileView) GetShareholderStake() ShareholderStake`

GetShareholderStake returns the ShareholderStake field if non-nil, zero value otherwise.

### GetShareholderStakeOk

`func (o *ListingProfileView) GetShareholderStakeOk() (*ShareholderStake, bool)`

GetShareholderStakeOk returns a tuple with the ShareholderStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareholderStake

`func (o *ListingProfileView) SetShareholderStake(v ShareholderStake)`

SetShareholderStake sets ShareholderStake field to given value.

### HasShareholderStake

`func (o *ListingProfileView) HasShareholderStake() bool`

HasShareholderStake returns a boolean if a field has been set.

### GetStartDate

`func (o *ListingProfileView) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListingProfileView) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListingProfileView) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListingProfileView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetSystemBond

`func (o *ListingProfileView) GetSystemBond() SystemBondView`

GetSystemBond returns the SystemBond field if non-nil, zero value otherwise.

### GetSystemBondOk

`func (o *ListingProfileView) GetSystemBondOk() (*SystemBondView, bool)`

GetSystemBondOk returns a tuple with the SystemBond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemBond

`func (o *ListingProfileView) SetSystemBond(v SystemBondView)`

SetSystemBond sets SystemBond field to given value.

### HasSystemBond

`func (o *ListingProfileView) HasSystemBond() bool`

HasSystemBond returns a boolean if a field has been set.

### GetType

`func (o *ListingProfileView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListingProfileView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListingProfileView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListingProfileView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *ListingProfileView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListingProfileView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListingProfileView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ListingProfileView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


