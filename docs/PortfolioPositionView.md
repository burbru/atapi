# PortfolioPositionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageBuyingPrice** | Pointer to **float32** |  | [optional] 
**CommittedShares** | Pointer to **int64** |  | [optional] 
**CurrentAskPrice** | Pointer to **float32** |  | [optional] 
**CurrentAskSize** | Pointer to **int64** |  | [optional] 
**CurrentBidPrice** | Pointer to **float32** |  | [optional] 
**CurrentBidSize** | Pointer to **int64** |  | [optional] 
**LastBuyingPrice** | Pointer to **float32** |  | [optional] 
**LastPrice** | Pointer to [**FixedIncomeEnabledSecurityPriceView**](FixedIncomeEnabledSecurityPriceView.md) |  | [optional] 
**LastPriceUpdate** | Pointer to **int64** |  | [optional] 
**Listing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**NumberOfShares** | Pointer to **int64** |  | [optional] 
**SecurityIdentifier** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to **float32** |  | [optional] 

## Methods

### NewPortfolioPositionView

`func NewPortfolioPositionView() *PortfolioPositionView`

NewPortfolioPositionView instantiates a new PortfolioPositionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortfolioPositionViewWithDefaults

`func NewPortfolioPositionViewWithDefaults() *PortfolioPositionView`

NewPortfolioPositionViewWithDefaults instantiates a new PortfolioPositionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageBuyingPrice

`func (o *PortfolioPositionView) GetAverageBuyingPrice() float32`

GetAverageBuyingPrice returns the AverageBuyingPrice field if non-nil, zero value otherwise.

### GetAverageBuyingPriceOk

`func (o *PortfolioPositionView) GetAverageBuyingPriceOk() (*float32, bool)`

GetAverageBuyingPriceOk returns a tuple with the AverageBuyingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageBuyingPrice

`func (o *PortfolioPositionView) SetAverageBuyingPrice(v float32)`

SetAverageBuyingPrice sets AverageBuyingPrice field to given value.

### HasAverageBuyingPrice

`func (o *PortfolioPositionView) HasAverageBuyingPrice() bool`

HasAverageBuyingPrice returns a boolean if a field has been set.

### GetCommittedShares

`func (o *PortfolioPositionView) GetCommittedShares() int64`

GetCommittedShares returns the CommittedShares field if non-nil, zero value otherwise.

### GetCommittedSharesOk

`func (o *PortfolioPositionView) GetCommittedSharesOk() (*int64, bool)`

GetCommittedSharesOk returns a tuple with the CommittedShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedShares

`func (o *PortfolioPositionView) SetCommittedShares(v int64)`

SetCommittedShares sets CommittedShares field to given value.

### HasCommittedShares

`func (o *PortfolioPositionView) HasCommittedShares() bool`

HasCommittedShares returns a boolean if a field has been set.

### GetCurrentAskPrice

`func (o *PortfolioPositionView) GetCurrentAskPrice() float32`

GetCurrentAskPrice returns the CurrentAskPrice field if non-nil, zero value otherwise.

### GetCurrentAskPriceOk

`func (o *PortfolioPositionView) GetCurrentAskPriceOk() (*float32, bool)`

GetCurrentAskPriceOk returns a tuple with the CurrentAskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAskPrice

`func (o *PortfolioPositionView) SetCurrentAskPrice(v float32)`

SetCurrentAskPrice sets CurrentAskPrice field to given value.

### HasCurrentAskPrice

`func (o *PortfolioPositionView) HasCurrentAskPrice() bool`

HasCurrentAskPrice returns a boolean if a field has been set.

### GetCurrentAskSize

`func (o *PortfolioPositionView) GetCurrentAskSize() int64`

GetCurrentAskSize returns the CurrentAskSize field if non-nil, zero value otherwise.

### GetCurrentAskSizeOk

`func (o *PortfolioPositionView) GetCurrentAskSizeOk() (*int64, bool)`

GetCurrentAskSizeOk returns a tuple with the CurrentAskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAskSize

`func (o *PortfolioPositionView) SetCurrentAskSize(v int64)`

SetCurrentAskSize sets CurrentAskSize field to given value.

### HasCurrentAskSize

`func (o *PortfolioPositionView) HasCurrentAskSize() bool`

HasCurrentAskSize returns a boolean if a field has been set.

### GetCurrentBidPrice

`func (o *PortfolioPositionView) GetCurrentBidPrice() float32`

GetCurrentBidPrice returns the CurrentBidPrice field if non-nil, zero value otherwise.

### GetCurrentBidPriceOk

`func (o *PortfolioPositionView) GetCurrentBidPriceOk() (*float32, bool)`

GetCurrentBidPriceOk returns a tuple with the CurrentBidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBidPrice

`func (o *PortfolioPositionView) SetCurrentBidPrice(v float32)`

SetCurrentBidPrice sets CurrentBidPrice field to given value.

### HasCurrentBidPrice

`func (o *PortfolioPositionView) HasCurrentBidPrice() bool`

HasCurrentBidPrice returns a boolean if a field has been set.

### GetCurrentBidSize

`func (o *PortfolioPositionView) GetCurrentBidSize() int64`

GetCurrentBidSize returns the CurrentBidSize field if non-nil, zero value otherwise.

### GetCurrentBidSizeOk

`func (o *PortfolioPositionView) GetCurrentBidSizeOk() (*int64, bool)`

GetCurrentBidSizeOk returns a tuple with the CurrentBidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBidSize

`func (o *PortfolioPositionView) SetCurrentBidSize(v int64)`

SetCurrentBidSize sets CurrentBidSize field to given value.

### HasCurrentBidSize

`func (o *PortfolioPositionView) HasCurrentBidSize() bool`

HasCurrentBidSize returns a boolean if a field has been set.

### GetLastBuyingPrice

`func (o *PortfolioPositionView) GetLastBuyingPrice() float32`

GetLastBuyingPrice returns the LastBuyingPrice field if non-nil, zero value otherwise.

### GetLastBuyingPriceOk

`func (o *PortfolioPositionView) GetLastBuyingPriceOk() (*float32, bool)`

GetLastBuyingPriceOk returns a tuple with the LastBuyingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBuyingPrice

`func (o *PortfolioPositionView) SetLastBuyingPrice(v float32)`

SetLastBuyingPrice sets LastBuyingPrice field to given value.

### HasLastBuyingPrice

`func (o *PortfolioPositionView) HasLastBuyingPrice() bool`

HasLastBuyingPrice returns a boolean if a field has been set.

### GetLastPrice

`func (o *PortfolioPositionView) GetLastPrice() FixedIncomeEnabledSecurityPriceView`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *PortfolioPositionView) GetLastPriceOk() (*FixedIncomeEnabledSecurityPriceView, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *PortfolioPositionView) SetLastPrice(v FixedIncomeEnabledSecurityPriceView)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *PortfolioPositionView) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLastPriceUpdate

`func (o *PortfolioPositionView) GetLastPriceUpdate() int64`

GetLastPriceUpdate returns the LastPriceUpdate field if non-nil, zero value otherwise.

### GetLastPriceUpdateOk

`func (o *PortfolioPositionView) GetLastPriceUpdateOk() (*int64, bool)`

GetLastPriceUpdateOk returns a tuple with the LastPriceUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPriceUpdate

`func (o *PortfolioPositionView) SetLastPriceUpdate(v int64)`

SetLastPriceUpdate sets LastPriceUpdate field to given value.

### HasLastPriceUpdate

`func (o *PortfolioPositionView) HasLastPriceUpdate() bool`

HasLastPriceUpdate returns a boolean if a field has been set.

### GetListing

`func (o *PortfolioPositionView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *PortfolioPositionView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *PortfolioPositionView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *PortfolioPositionView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetNumberOfShares

`func (o *PortfolioPositionView) GetNumberOfShares() int64`

GetNumberOfShares returns the NumberOfShares field if non-nil, zero value otherwise.

### GetNumberOfSharesOk

`func (o *PortfolioPositionView) GetNumberOfSharesOk() (*int64, bool)`

GetNumberOfSharesOk returns a tuple with the NumberOfShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfShares

`func (o *PortfolioPositionView) SetNumberOfShares(v int64)`

SetNumberOfShares sets NumberOfShares field to given value.

### HasNumberOfShares

`func (o *PortfolioPositionView) HasNumberOfShares() bool`

HasNumberOfShares returns a boolean if a field has been set.

### GetSecurityIdentifier

`func (o *PortfolioPositionView) GetSecurityIdentifier() string`

GetSecurityIdentifier returns the SecurityIdentifier field if non-nil, zero value otherwise.

### GetSecurityIdentifierOk

`func (o *PortfolioPositionView) GetSecurityIdentifierOk() (*string, bool)`

GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIdentifier

`func (o *PortfolioPositionView) SetSecurityIdentifier(v string)`

SetSecurityIdentifier sets SecurityIdentifier field to given value.

### HasSecurityIdentifier

`func (o *PortfolioPositionView) HasSecurityIdentifier() bool`

HasSecurityIdentifier returns a boolean if a field has been set.

### GetType

`func (o *PortfolioPositionView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PortfolioPositionView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PortfolioPositionView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PortfolioPositionView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVolume

`func (o *PortfolioPositionView) GetVolume() float32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *PortfolioPositionView) GetVolumeOk() (*float32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *PortfolioPositionView) SetVolume(v float32)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *PortfolioPositionView) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


