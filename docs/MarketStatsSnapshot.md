# MarketStatsSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlphaCoins** | Pointer to **int64** |  | [optional] 
**AverageBondDurationInDays** | Pointer to **float32** |  | [optional] 
**AverageBookValue** | Pointer to **float32** |  | [optional] 
**AverageDailyWage** | Pointer to **float32** |  | [optional] 
**AverageDesignatedSponsorRating** | Pointer to **string** |  | [optional] 
**AverageYieldToMaturity** | Pointer to **float32** |  | [optional] 
**BondFaceVolume** | Pointer to **float32** |  | [optional] 
**CentralBankReserves** | Pointer to **float32** |  | [optional] 
**CommittedCash** | Pointer to **float32** |  | [optional] 
**CorporateCash** | Pointer to **float32** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MainInterestRate** | Pointer to **float32** |  | [optional] 
**MarketCap** | Pointer to **float32** |  | [optional] 
**NumberOfActiveOtherListings** | Pointer to **int64** |  | [optional] 
**NumberOfBanks** | Pointer to **int64** |  | [optional] 
**NumberOfBondOrders** | Pointer to **int64** |  | [optional] 
**NumberOfCashoutPolls** | Pointer to **int64** |  | [optional] 
**NumberOfCommittedShares** | Pointer to **int64** |  | [optional] 
**NumberOfCompanies** | Pointer to **int64** |  | [optional] 
**NumberOfDesignatedSponsors** | Pointer to **int64** |  | [optional] 
**NumberOfLiquidationPolls** | Pointer to **int64** |  | [optional] 
**NumberOfOrders** | Pointer to **int64** |  | [optional] 
**NumberOfOrders24h** | Pointer to **int64** |  | [optional] 
**NumberOfOtcOrders** | Pointer to **int64** |  | [optional] 
**NumberOfOtherListings** | Pointer to **int64** |  | [optional] 
**NumberOfOtherOrders** | Pointer to **int64** |  | [optional] 
**NumberOfPartnerUsers** | Pointer to **int64** |  | [optional] 
**NumberOfPremiumUsers** | Pointer to **int64** |  | [optional] 
**NumberOfRepoOrders** | Pointer to **int64** |  | [optional] 
**NumberOfStockOrders** | Pointer to **int64** |  | [optional] 
**NumberOfSystemBondOrders** | Pointer to **int64** |  | [optional] 
**NumberOfSystemRepoOrders** | Pointer to **int64** |  | [optional] 
**NumberOfUsers** | Pointer to **int64** |  | [optional] 
**OnlineUsers** | Pointer to **int64** |  | [optional] 
**OrderVolume** | Pointer to **float32** |  | [optional] 
**OrderVolume24h** | Pointer to **float32** |  | [optional] 
**PrivateCash** | Pointer to **float32** |  | [optional] 
**SystemBondFaceVolume** | Pointer to **float32** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewMarketStatsSnapshot

`func NewMarketStatsSnapshot() *MarketStatsSnapshot`

NewMarketStatsSnapshot instantiates a new MarketStatsSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketStatsSnapshotWithDefaults

`func NewMarketStatsSnapshotWithDefaults() *MarketStatsSnapshot`

NewMarketStatsSnapshotWithDefaults instantiates a new MarketStatsSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlphaCoins

`func (o *MarketStatsSnapshot) GetAlphaCoins() int64`

GetAlphaCoins returns the AlphaCoins field if non-nil, zero value otherwise.

### GetAlphaCoinsOk

`func (o *MarketStatsSnapshot) GetAlphaCoinsOk() (*int64, bool)`

GetAlphaCoinsOk returns a tuple with the AlphaCoins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlphaCoins

`func (o *MarketStatsSnapshot) SetAlphaCoins(v int64)`

SetAlphaCoins sets AlphaCoins field to given value.

### HasAlphaCoins

`func (o *MarketStatsSnapshot) HasAlphaCoins() bool`

HasAlphaCoins returns a boolean if a field has been set.

### GetAverageBondDurationInDays

`func (o *MarketStatsSnapshot) GetAverageBondDurationInDays() float32`

GetAverageBondDurationInDays returns the AverageBondDurationInDays field if non-nil, zero value otherwise.

### GetAverageBondDurationInDaysOk

`func (o *MarketStatsSnapshot) GetAverageBondDurationInDaysOk() (*float32, bool)`

GetAverageBondDurationInDaysOk returns a tuple with the AverageBondDurationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageBondDurationInDays

`func (o *MarketStatsSnapshot) SetAverageBondDurationInDays(v float32)`

SetAverageBondDurationInDays sets AverageBondDurationInDays field to given value.

### HasAverageBondDurationInDays

`func (o *MarketStatsSnapshot) HasAverageBondDurationInDays() bool`

HasAverageBondDurationInDays returns a boolean if a field has been set.

### GetAverageBookValue

`func (o *MarketStatsSnapshot) GetAverageBookValue() float32`

GetAverageBookValue returns the AverageBookValue field if non-nil, zero value otherwise.

### GetAverageBookValueOk

`func (o *MarketStatsSnapshot) GetAverageBookValueOk() (*float32, bool)`

GetAverageBookValueOk returns a tuple with the AverageBookValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageBookValue

`func (o *MarketStatsSnapshot) SetAverageBookValue(v float32)`

SetAverageBookValue sets AverageBookValue field to given value.

### HasAverageBookValue

`func (o *MarketStatsSnapshot) HasAverageBookValue() bool`

HasAverageBookValue returns a boolean if a field has been set.

### GetAverageDailyWage

`func (o *MarketStatsSnapshot) GetAverageDailyWage() float32`

GetAverageDailyWage returns the AverageDailyWage field if non-nil, zero value otherwise.

### GetAverageDailyWageOk

`func (o *MarketStatsSnapshot) GetAverageDailyWageOk() (*float32, bool)`

GetAverageDailyWageOk returns a tuple with the AverageDailyWage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDailyWage

`func (o *MarketStatsSnapshot) SetAverageDailyWage(v float32)`

SetAverageDailyWage sets AverageDailyWage field to given value.

### HasAverageDailyWage

`func (o *MarketStatsSnapshot) HasAverageDailyWage() bool`

HasAverageDailyWage returns a boolean if a field has been set.

### GetAverageDesignatedSponsorRating

`func (o *MarketStatsSnapshot) GetAverageDesignatedSponsorRating() string`

GetAverageDesignatedSponsorRating returns the AverageDesignatedSponsorRating field if non-nil, zero value otherwise.

### GetAverageDesignatedSponsorRatingOk

`func (o *MarketStatsSnapshot) GetAverageDesignatedSponsorRatingOk() (*string, bool)`

GetAverageDesignatedSponsorRatingOk returns a tuple with the AverageDesignatedSponsorRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDesignatedSponsorRating

`func (o *MarketStatsSnapshot) SetAverageDesignatedSponsorRating(v string)`

SetAverageDesignatedSponsorRating sets AverageDesignatedSponsorRating field to given value.

### HasAverageDesignatedSponsorRating

`func (o *MarketStatsSnapshot) HasAverageDesignatedSponsorRating() bool`

HasAverageDesignatedSponsorRating returns a boolean if a field has been set.

### GetAverageYieldToMaturity

`func (o *MarketStatsSnapshot) GetAverageYieldToMaturity() float32`

GetAverageYieldToMaturity returns the AverageYieldToMaturity field if non-nil, zero value otherwise.

### GetAverageYieldToMaturityOk

`func (o *MarketStatsSnapshot) GetAverageYieldToMaturityOk() (*float32, bool)`

GetAverageYieldToMaturityOk returns a tuple with the AverageYieldToMaturity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageYieldToMaturity

`func (o *MarketStatsSnapshot) SetAverageYieldToMaturity(v float32)`

SetAverageYieldToMaturity sets AverageYieldToMaturity field to given value.

### HasAverageYieldToMaturity

`func (o *MarketStatsSnapshot) HasAverageYieldToMaturity() bool`

HasAverageYieldToMaturity returns a boolean if a field has been set.

### GetBondFaceVolume

`func (o *MarketStatsSnapshot) GetBondFaceVolume() float32`

GetBondFaceVolume returns the BondFaceVolume field if non-nil, zero value otherwise.

### GetBondFaceVolumeOk

`func (o *MarketStatsSnapshot) GetBondFaceVolumeOk() (*float32, bool)`

GetBondFaceVolumeOk returns a tuple with the BondFaceVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondFaceVolume

`func (o *MarketStatsSnapshot) SetBondFaceVolume(v float32)`

SetBondFaceVolume sets BondFaceVolume field to given value.

### HasBondFaceVolume

`func (o *MarketStatsSnapshot) HasBondFaceVolume() bool`

HasBondFaceVolume returns a boolean if a field has been set.

### GetCentralBankReserves

`func (o *MarketStatsSnapshot) GetCentralBankReserves() float32`

GetCentralBankReserves returns the CentralBankReserves field if non-nil, zero value otherwise.

### GetCentralBankReservesOk

`func (o *MarketStatsSnapshot) GetCentralBankReservesOk() (*float32, bool)`

GetCentralBankReservesOk returns a tuple with the CentralBankReserves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentralBankReserves

`func (o *MarketStatsSnapshot) SetCentralBankReserves(v float32)`

SetCentralBankReserves sets CentralBankReserves field to given value.

### HasCentralBankReserves

`func (o *MarketStatsSnapshot) HasCentralBankReserves() bool`

HasCentralBankReserves returns a boolean if a field has been set.

### GetCommittedCash

`func (o *MarketStatsSnapshot) GetCommittedCash() float32`

GetCommittedCash returns the CommittedCash field if non-nil, zero value otherwise.

### GetCommittedCashOk

`func (o *MarketStatsSnapshot) GetCommittedCashOk() (*float32, bool)`

GetCommittedCashOk returns a tuple with the CommittedCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedCash

`func (o *MarketStatsSnapshot) SetCommittedCash(v float32)`

SetCommittedCash sets CommittedCash field to given value.

### HasCommittedCash

`func (o *MarketStatsSnapshot) HasCommittedCash() bool`

HasCommittedCash returns a boolean if a field has been set.

### GetCorporateCash

`func (o *MarketStatsSnapshot) GetCorporateCash() float32`

GetCorporateCash returns the CorporateCash field if non-nil, zero value otherwise.

### GetCorporateCashOk

`func (o *MarketStatsSnapshot) GetCorporateCashOk() (*float32, bool)`

GetCorporateCashOk returns a tuple with the CorporateCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateCash

`func (o *MarketStatsSnapshot) SetCorporateCash(v float32)`

SetCorporateCash sets CorporateCash field to given value.

### HasCorporateCash

`func (o *MarketStatsSnapshot) HasCorporateCash() bool`

HasCorporateCash returns a boolean if a field has been set.

### GetDate

`func (o *MarketStatsSnapshot) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *MarketStatsSnapshot) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *MarketStatsSnapshot) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *MarketStatsSnapshot) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetId

`func (o *MarketStatsSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketStatsSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketStatsSnapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketStatsSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMainInterestRate

`func (o *MarketStatsSnapshot) GetMainInterestRate() float32`

GetMainInterestRate returns the MainInterestRate field if non-nil, zero value otherwise.

### GetMainInterestRateOk

`func (o *MarketStatsSnapshot) GetMainInterestRateOk() (*float32, bool)`

GetMainInterestRateOk returns a tuple with the MainInterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainInterestRate

`func (o *MarketStatsSnapshot) SetMainInterestRate(v float32)`

SetMainInterestRate sets MainInterestRate field to given value.

### HasMainInterestRate

`func (o *MarketStatsSnapshot) HasMainInterestRate() bool`

HasMainInterestRate returns a boolean if a field has been set.

### GetMarketCap

`func (o *MarketStatsSnapshot) GetMarketCap() float32`

GetMarketCap returns the MarketCap field if non-nil, zero value otherwise.

### GetMarketCapOk

`func (o *MarketStatsSnapshot) GetMarketCapOk() (*float32, bool)`

GetMarketCapOk returns a tuple with the MarketCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCap

`func (o *MarketStatsSnapshot) SetMarketCap(v float32)`

SetMarketCap sets MarketCap field to given value.

### HasMarketCap

`func (o *MarketStatsSnapshot) HasMarketCap() bool`

HasMarketCap returns a boolean if a field has been set.

### GetNumberOfActiveOtherListings

`func (o *MarketStatsSnapshot) GetNumberOfActiveOtherListings() int64`

GetNumberOfActiveOtherListings returns the NumberOfActiveOtherListings field if non-nil, zero value otherwise.

### GetNumberOfActiveOtherListingsOk

`func (o *MarketStatsSnapshot) GetNumberOfActiveOtherListingsOk() (*int64, bool)`

GetNumberOfActiveOtherListingsOk returns a tuple with the NumberOfActiveOtherListings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfActiveOtherListings

`func (o *MarketStatsSnapshot) SetNumberOfActiveOtherListings(v int64)`

SetNumberOfActiveOtherListings sets NumberOfActiveOtherListings field to given value.

### HasNumberOfActiveOtherListings

`func (o *MarketStatsSnapshot) HasNumberOfActiveOtherListings() bool`

HasNumberOfActiveOtherListings returns a boolean if a field has been set.

### GetNumberOfBanks

`func (o *MarketStatsSnapshot) GetNumberOfBanks() int64`

GetNumberOfBanks returns the NumberOfBanks field if non-nil, zero value otherwise.

### GetNumberOfBanksOk

`func (o *MarketStatsSnapshot) GetNumberOfBanksOk() (*int64, bool)`

GetNumberOfBanksOk returns a tuple with the NumberOfBanks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfBanks

`func (o *MarketStatsSnapshot) SetNumberOfBanks(v int64)`

SetNumberOfBanks sets NumberOfBanks field to given value.

### HasNumberOfBanks

`func (o *MarketStatsSnapshot) HasNumberOfBanks() bool`

HasNumberOfBanks returns a boolean if a field has been set.

### GetNumberOfBondOrders

`func (o *MarketStatsSnapshot) GetNumberOfBondOrders() int64`

GetNumberOfBondOrders returns the NumberOfBondOrders field if non-nil, zero value otherwise.

### GetNumberOfBondOrdersOk

`func (o *MarketStatsSnapshot) GetNumberOfBondOrdersOk() (*int64, bool)`

GetNumberOfBondOrdersOk returns a tuple with the NumberOfBondOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfBondOrders

`func (o *MarketStatsSnapshot) SetNumberOfBondOrders(v int64)`

SetNumberOfBondOrders sets NumberOfBondOrders field to given value.

### HasNumberOfBondOrders

`func (o *MarketStatsSnapshot) HasNumberOfBondOrders() bool`

HasNumberOfBondOrders returns a boolean if a field has been set.

### GetNumberOfCashoutPolls

`func (o *MarketStatsSnapshot) GetNumberOfCashoutPolls() int64`

GetNumberOfCashoutPolls returns the NumberOfCashoutPolls field if non-nil, zero value otherwise.

### GetNumberOfCashoutPollsOk

`func (o *MarketStatsSnapshot) GetNumberOfCashoutPollsOk() (*int64, bool)`

GetNumberOfCashoutPollsOk returns a tuple with the NumberOfCashoutPolls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCashoutPolls

`func (o *MarketStatsSnapshot) SetNumberOfCashoutPolls(v int64)`

SetNumberOfCashoutPolls sets NumberOfCashoutPolls field to given value.

### HasNumberOfCashoutPolls

`func (o *MarketStatsSnapshot) HasNumberOfCashoutPolls() bool`

HasNumberOfCashoutPolls returns a boolean if a field has been set.

### GetNumberOfCommittedShares

`func (o *MarketStatsSnapshot) GetNumberOfCommittedShares() int64`

GetNumberOfCommittedShares returns the NumberOfCommittedShares field if non-nil, zero value otherwise.

### GetNumberOfCommittedSharesOk

`func (o *MarketStatsSnapshot) GetNumberOfCommittedSharesOk() (*int64, bool)`

GetNumberOfCommittedSharesOk returns a tuple with the NumberOfCommittedShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCommittedShares

`func (o *MarketStatsSnapshot) SetNumberOfCommittedShares(v int64)`

SetNumberOfCommittedShares sets NumberOfCommittedShares field to given value.

### HasNumberOfCommittedShares

`func (o *MarketStatsSnapshot) HasNumberOfCommittedShares() bool`

HasNumberOfCommittedShares returns a boolean if a field has been set.

### GetNumberOfCompanies

`func (o *MarketStatsSnapshot) GetNumberOfCompanies() int64`

GetNumberOfCompanies returns the NumberOfCompanies field if non-nil, zero value otherwise.

### GetNumberOfCompaniesOk

`func (o *MarketStatsSnapshot) GetNumberOfCompaniesOk() (*int64, bool)`

GetNumberOfCompaniesOk returns a tuple with the NumberOfCompanies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCompanies

`func (o *MarketStatsSnapshot) SetNumberOfCompanies(v int64)`

SetNumberOfCompanies sets NumberOfCompanies field to given value.

### HasNumberOfCompanies

`func (o *MarketStatsSnapshot) HasNumberOfCompanies() bool`

HasNumberOfCompanies returns a boolean if a field has been set.

### GetNumberOfDesignatedSponsors

`func (o *MarketStatsSnapshot) GetNumberOfDesignatedSponsors() int64`

GetNumberOfDesignatedSponsors returns the NumberOfDesignatedSponsors field if non-nil, zero value otherwise.

### GetNumberOfDesignatedSponsorsOk

`func (o *MarketStatsSnapshot) GetNumberOfDesignatedSponsorsOk() (*int64, bool)`

GetNumberOfDesignatedSponsorsOk returns a tuple with the NumberOfDesignatedSponsors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDesignatedSponsors

`func (o *MarketStatsSnapshot) SetNumberOfDesignatedSponsors(v int64)`

SetNumberOfDesignatedSponsors sets NumberOfDesignatedSponsors field to given value.

### HasNumberOfDesignatedSponsors

`func (o *MarketStatsSnapshot) HasNumberOfDesignatedSponsors() bool`

HasNumberOfDesignatedSponsors returns a boolean if a field has been set.

### GetNumberOfLiquidationPolls

`func (o *MarketStatsSnapshot) GetNumberOfLiquidationPolls() int64`

GetNumberOfLiquidationPolls returns the NumberOfLiquidationPolls field if non-nil, zero value otherwise.

### GetNumberOfLiquidationPollsOk

`func (o *MarketStatsSnapshot) GetNumberOfLiquidationPollsOk() (*int64, bool)`

GetNumberOfLiquidationPollsOk returns a tuple with the NumberOfLiquidationPolls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfLiquidationPolls

`func (o *MarketStatsSnapshot) SetNumberOfLiquidationPolls(v int64)`

SetNumberOfLiquidationPolls sets NumberOfLiquidationPolls field to given value.

### HasNumberOfLiquidationPolls

`func (o *MarketStatsSnapshot) HasNumberOfLiquidationPolls() bool`

HasNumberOfLiquidationPolls returns a boolean if a field has been set.

### GetNumberOfOrders

`func (o *MarketStatsSnapshot) GetNumberOfOrders() int64`

GetNumberOfOrders returns the NumberOfOrders field if non-nil, zero value otherwise.

### GetNumberOfOrdersOk

`func (o *MarketStatsSnapshot) GetNumberOfOrdersOk() (*int64, bool)`

GetNumberOfOrdersOk returns a tuple with the NumberOfOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfOrders

`func (o *MarketStatsSnapshot) SetNumberOfOrders(v int64)`

SetNumberOfOrders sets NumberOfOrders field to given value.

### HasNumberOfOrders

`func (o *MarketStatsSnapshot) HasNumberOfOrders() bool`

HasNumberOfOrders returns a boolean if a field has been set.

### GetNumberOfOrders24h

`func (o *MarketStatsSnapshot) GetNumberOfOrders24h() int64`

GetNumberOfOrders24h returns the NumberOfOrders24h field if non-nil, zero value otherwise.

### GetNumberOfOrders24hOk

`func (o *MarketStatsSnapshot) GetNumberOfOrders24hOk() (*int64, bool)`

GetNumberOfOrders24hOk returns a tuple with the NumberOfOrders24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfOrders24h

`func (o *MarketStatsSnapshot) SetNumberOfOrders24h(v int64)`

SetNumberOfOrders24h sets NumberOfOrders24h field to given value.

### HasNumberOfOrders24h

`func (o *MarketStatsSnapshot) HasNumberOfOrders24h() bool`

HasNumberOfOrders24h returns a boolean if a field has been set.

### GetNumberOfOtcOrders

`func (o *MarketStatsSnapshot) GetNumberOfOtcOrders() int64`

GetNumberOfOtcOrders returns the NumberOfOtcOrders field if non-nil, zero value otherwise.

### GetNumberOfOtcOrdersOk

`func (o *MarketStatsSnapshot) GetNumberOfOtcOrdersOk() (*int64, bool)`

GetNumberOfOtcOrdersOk returns a tuple with the NumberOfOtcOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfOtcOrders

`func (o *MarketStatsSnapshot) SetNumberOfOtcOrders(v int64)`

SetNumberOfOtcOrders sets NumberOfOtcOrders field to given value.

### HasNumberOfOtcOrders

`func (o *MarketStatsSnapshot) HasNumberOfOtcOrders() bool`

HasNumberOfOtcOrders returns a boolean if a field has been set.

### GetNumberOfOtherListings

`func (o *MarketStatsSnapshot) GetNumberOfOtherListings() int64`

GetNumberOfOtherListings returns the NumberOfOtherListings field if non-nil, zero value otherwise.

### GetNumberOfOtherListingsOk

`func (o *MarketStatsSnapshot) GetNumberOfOtherListingsOk() (*int64, bool)`

GetNumberOfOtherListingsOk returns a tuple with the NumberOfOtherListings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfOtherListings

`func (o *MarketStatsSnapshot) SetNumberOfOtherListings(v int64)`

SetNumberOfOtherListings sets NumberOfOtherListings field to given value.

### HasNumberOfOtherListings

`func (o *MarketStatsSnapshot) HasNumberOfOtherListings() bool`

HasNumberOfOtherListings returns a boolean if a field has been set.

### GetNumberOfOtherOrders

`func (o *MarketStatsSnapshot) GetNumberOfOtherOrders() int64`

GetNumberOfOtherOrders returns the NumberOfOtherOrders field if non-nil, zero value otherwise.

### GetNumberOfOtherOrdersOk

`func (o *MarketStatsSnapshot) GetNumberOfOtherOrdersOk() (*int64, bool)`

GetNumberOfOtherOrdersOk returns a tuple with the NumberOfOtherOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfOtherOrders

`func (o *MarketStatsSnapshot) SetNumberOfOtherOrders(v int64)`

SetNumberOfOtherOrders sets NumberOfOtherOrders field to given value.

### HasNumberOfOtherOrders

`func (o *MarketStatsSnapshot) HasNumberOfOtherOrders() bool`

HasNumberOfOtherOrders returns a boolean if a field has been set.

### GetNumberOfPartnerUsers

`func (o *MarketStatsSnapshot) GetNumberOfPartnerUsers() int64`

GetNumberOfPartnerUsers returns the NumberOfPartnerUsers field if non-nil, zero value otherwise.

### GetNumberOfPartnerUsersOk

`func (o *MarketStatsSnapshot) GetNumberOfPartnerUsersOk() (*int64, bool)`

GetNumberOfPartnerUsersOk returns a tuple with the NumberOfPartnerUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPartnerUsers

`func (o *MarketStatsSnapshot) SetNumberOfPartnerUsers(v int64)`

SetNumberOfPartnerUsers sets NumberOfPartnerUsers field to given value.

### HasNumberOfPartnerUsers

`func (o *MarketStatsSnapshot) HasNumberOfPartnerUsers() bool`

HasNumberOfPartnerUsers returns a boolean if a field has been set.

### GetNumberOfPremiumUsers

`func (o *MarketStatsSnapshot) GetNumberOfPremiumUsers() int64`

GetNumberOfPremiumUsers returns the NumberOfPremiumUsers field if non-nil, zero value otherwise.

### GetNumberOfPremiumUsersOk

`func (o *MarketStatsSnapshot) GetNumberOfPremiumUsersOk() (*int64, bool)`

GetNumberOfPremiumUsersOk returns a tuple with the NumberOfPremiumUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPremiumUsers

`func (o *MarketStatsSnapshot) SetNumberOfPremiumUsers(v int64)`

SetNumberOfPremiumUsers sets NumberOfPremiumUsers field to given value.

### HasNumberOfPremiumUsers

`func (o *MarketStatsSnapshot) HasNumberOfPremiumUsers() bool`

HasNumberOfPremiumUsers returns a boolean if a field has been set.

### GetNumberOfRepoOrders

`func (o *MarketStatsSnapshot) GetNumberOfRepoOrders() int64`

GetNumberOfRepoOrders returns the NumberOfRepoOrders field if non-nil, zero value otherwise.

### GetNumberOfRepoOrdersOk

`func (o *MarketStatsSnapshot) GetNumberOfRepoOrdersOk() (*int64, bool)`

GetNumberOfRepoOrdersOk returns a tuple with the NumberOfRepoOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfRepoOrders

`func (o *MarketStatsSnapshot) SetNumberOfRepoOrders(v int64)`

SetNumberOfRepoOrders sets NumberOfRepoOrders field to given value.

### HasNumberOfRepoOrders

`func (o *MarketStatsSnapshot) HasNumberOfRepoOrders() bool`

HasNumberOfRepoOrders returns a boolean if a field has been set.

### GetNumberOfStockOrders

`func (o *MarketStatsSnapshot) GetNumberOfStockOrders() int64`

GetNumberOfStockOrders returns the NumberOfStockOrders field if non-nil, zero value otherwise.

### GetNumberOfStockOrdersOk

`func (o *MarketStatsSnapshot) GetNumberOfStockOrdersOk() (*int64, bool)`

GetNumberOfStockOrdersOk returns a tuple with the NumberOfStockOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfStockOrders

`func (o *MarketStatsSnapshot) SetNumberOfStockOrders(v int64)`

SetNumberOfStockOrders sets NumberOfStockOrders field to given value.

### HasNumberOfStockOrders

`func (o *MarketStatsSnapshot) HasNumberOfStockOrders() bool`

HasNumberOfStockOrders returns a boolean if a field has been set.

### GetNumberOfSystemBondOrders

`func (o *MarketStatsSnapshot) GetNumberOfSystemBondOrders() int64`

GetNumberOfSystemBondOrders returns the NumberOfSystemBondOrders field if non-nil, zero value otherwise.

### GetNumberOfSystemBondOrdersOk

`func (o *MarketStatsSnapshot) GetNumberOfSystemBondOrdersOk() (*int64, bool)`

GetNumberOfSystemBondOrdersOk returns a tuple with the NumberOfSystemBondOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSystemBondOrders

`func (o *MarketStatsSnapshot) SetNumberOfSystemBondOrders(v int64)`

SetNumberOfSystemBondOrders sets NumberOfSystemBondOrders field to given value.

### HasNumberOfSystemBondOrders

`func (o *MarketStatsSnapshot) HasNumberOfSystemBondOrders() bool`

HasNumberOfSystemBondOrders returns a boolean if a field has been set.

### GetNumberOfSystemRepoOrders

`func (o *MarketStatsSnapshot) GetNumberOfSystemRepoOrders() int64`

GetNumberOfSystemRepoOrders returns the NumberOfSystemRepoOrders field if non-nil, zero value otherwise.

### GetNumberOfSystemRepoOrdersOk

`func (o *MarketStatsSnapshot) GetNumberOfSystemRepoOrdersOk() (*int64, bool)`

GetNumberOfSystemRepoOrdersOk returns a tuple with the NumberOfSystemRepoOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSystemRepoOrders

`func (o *MarketStatsSnapshot) SetNumberOfSystemRepoOrders(v int64)`

SetNumberOfSystemRepoOrders sets NumberOfSystemRepoOrders field to given value.

### HasNumberOfSystemRepoOrders

`func (o *MarketStatsSnapshot) HasNumberOfSystemRepoOrders() bool`

HasNumberOfSystemRepoOrders returns a boolean if a field has been set.

### GetNumberOfUsers

`func (o *MarketStatsSnapshot) GetNumberOfUsers() int64`

GetNumberOfUsers returns the NumberOfUsers field if non-nil, zero value otherwise.

### GetNumberOfUsersOk

`func (o *MarketStatsSnapshot) GetNumberOfUsersOk() (*int64, bool)`

GetNumberOfUsersOk returns a tuple with the NumberOfUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfUsers

`func (o *MarketStatsSnapshot) SetNumberOfUsers(v int64)`

SetNumberOfUsers sets NumberOfUsers field to given value.

### HasNumberOfUsers

`func (o *MarketStatsSnapshot) HasNumberOfUsers() bool`

HasNumberOfUsers returns a boolean if a field has been set.

### GetOnlineUsers

`func (o *MarketStatsSnapshot) GetOnlineUsers() int64`

GetOnlineUsers returns the OnlineUsers field if non-nil, zero value otherwise.

### GetOnlineUsersOk

`func (o *MarketStatsSnapshot) GetOnlineUsersOk() (*int64, bool)`

GetOnlineUsersOk returns a tuple with the OnlineUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineUsers

`func (o *MarketStatsSnapshot) SetOnlineUsers(v int64)`

SetOnlineUsers sets OnlineUsers field to given value.

### HasOnlineUsers

`func (o *MarketStatsSnapshot) HasOnlineUsers() bool`

HasOnlineUsers returns a boolean if a field has been set.

### GetOrderVolume

`func (o *MarketStatsSnapshot) GetOrderVolume() float32`

GetOrderVolume returns the OrderVolume field if non-nil, zero value otherwise.

### GetOrderVolumeOk

`func (o *MarketStatsSnapshot) GetOrderVolumeOk() (*float32, bool)`

GetOrderVolumeOk returns a tuple with the OrderVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderVolume

`func (o *MarketStatsSnapshot) SetOrderVolume(v float32)`

SetOrderVolume sets OrderVolume field to given value.

### HasOrderVolume

`func (o *MarketStatsSnapshot) HasOrderVolume() bool`

HasOrderVolume returns a boolean if a field has been set.

### GetOrderVolume24h

`func (o *MarketStatsSnapshot) GetOrderVolume24h() float32`

GetOrderVolume24h returns the OrderVolume24h field if non-nil, zero value otherwise.

### GetOrderVolume24hOk

`func (o *MarketStatsSnapshot) GetOrderVolume24hOk() (*float32, bool)`

GetOrderVolume24hOk returns a tuple with the OrderVolume24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderVolume24h

`func (o *MarketStatsSnapshot) SetOrderVolume24h(v float32)`

SetOrderVolume24h sets OrderVolume24h field to given value.

### HasOrderVolume24h

`func (o *MarketStatsSnapshot) HasOrderVolume24h() bool`

HasOrderVolume24h returns a boolean if a field has been set.

### GetPrivateCash

`func (o *MarketStatsSnapshot) GetPrivateCash() float32`

GetPrivateCash returns the PrivateCash field if non-nil, zero value otherwise.

### GetPrivateCashOk

`func (o *MarketStatsSnapshot) GetPrivateCashOk() (*float32, bool)`

GetPrivateCashOk returns a tuple with the PrivateCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateCash

`func (o *MarketStatsSnapshot) SetPrivateCash(v float32)`

SetPrivateCash sets PrivateCash field to given value.

### HasPrivateCash

`func (o *MarketStatsSnapshot) HasPrivateCash() bool`

HasPrivateCash returns a boolean if a field has been set.

### GetSystemBondFaceVolume

`func (o *MarketStatsSnapshot) GetSystemBondFaceVolume() float32`

GetSystemBondFaceVolume returns the SystemBondFaceVolume field if non-nil, zero value otherwise.

### GetSystemBondFaceVolumeOk

`func (o *MarketStatsSnapshot) GetSystemBondFaceVolumeOk() (*float32, bool)`

GetSystemBondFaceVolumeOk returns a tuple with the SystemBondFaceVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemBondFaceVolume

`func (o *MarketStatsSnapshot) SetSystemBondFaceVolume(v float32)`

SetSystemBondFaceVolume sets SystemBondFaceVolume field to given value.

### HasSystemBondFaceVolume

`func (o *MarketStatsSnapshot) HasSystemBondFaceVolume() bool`

HasSystemBondFaceVolume returns a boolean if a field has been set.

### GetVersion

`func (o *MarketStatsSnapshot) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MarketStatsSnapshot) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MarketStatsSnapshot) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MarketStatsSnapshot) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


