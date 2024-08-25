# ListingWithTradingVolumeView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskPrice** | Pointer to **float32** |  | [optional] 
**AskSize** | Pointer to **int64** |  | [optional] 
**BidPrice** | Pointer to **float32** |  | [optional] 
**BidSize** | Pointer to **int64** |  | [optional] 
**Date** | Pointer to **int64** |  | [optional] 
**EndDate** | Pointer to **int64** |  | [optional] 
**LastPrice** | Pointer to [**SecurityPriceView**](SecurityPriceView.md) |  | [optional] 
**Listing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SecurityIdentifier** | Pointer to **string** |  | [optional] 
**SpreadAbs** | Pointer to **float32** |  | [optional] 
**SpreadPercent** | Pointer to **float32** |  | [optional] 
**StartDate** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to **float32** |  | [optional] 

## Methods

### NewListingWithTradingVolumeView

`func NewListingWithTradingVolumeView() *ListingWithTradingVolumeView`

NewListingWithTradingVolumeView instantiates a new ListingWithTradingVolumeView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingWithTradingVolumeViewWithDefaults

`func NewListingWithTradingVolumeViewWithDefaults() *ListingWithTradingVolumeView`

NewListingWithTradingVolumeViewWithDefaults instantiates a new ListingWithTradingVolumeView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskPrice

`func (o *ListingWithTradingVolumeView) GetAskPrice() float32`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *ListingWithTradingVolumeView) GetAskPriceOk() (*float32, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *ListingWithTradingVolumeView) SetAskPrice(v float32)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *ListingWithTradingVolumeView) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetAskSize

`func (o *ListingWithTradingVolumeView) GetAskSize() int64`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *ListingWithTradingVolumeView) GetAskSizeOk() (*int64, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *ListingWithTradingVolumeView) SetAskSize(v int64)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *ListingWithTradingVolumeView) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.

### GetBidPrice

`func (o *ListingWithTradingVolumeView) GetBidPrice() float32`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *ListingWithTradingVolumeView) GetBidPriceOk() (*float32, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *ListingWithTradingVolumeView) SetBidPrice(v float32)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *ListingWithTradingVolumeView) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetBidSize

`func (o *ListingWithTradingVolumeView) GetBidSize() int64`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *ListingWithTradingVolumeView) GetBidSizeOk() (*int64, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *ListingWithTradingVolumeView) SetBidSize(v int64)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *ListingWithTradingVolumeView) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetDate

`func (o *ListingWithTradingVolumeView) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ListingWithTradingVolumeView) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ListingWithTradingVolumeView) SetDate(v int64)`

SetDate sets Date field to given value.

### HasDate

`func (o *ListingWithTradingVolumeView) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListingWithTradingVolumeView) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListingWithTradingVolumeView) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListingWithTradingVolumeView) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListingWithTradingVolumeView) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetLastPrice

`func (o *ListingWithTradingVolumeView) GetLastPrice() SecurityPriceView`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *ListingWithTradingVolumeView) GetLastPriceOk() (*SecurityPriceView, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *ListingWithTradingVolumeView) SetLastPrice(v SecurityPriceView)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *ListingWithTradingVolumeView) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetListing

`func (o *ListingWithTradingVolumeView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *ListingWithTradingVolumeView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *ListingWithTradingVolumeView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *ListingWithTradingVolumeView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetName

`func (o *ListingWithTradingVolumeView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListingWithTradingVolumeView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListingWithTradingVolumeView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListingWithTradingVolumeView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecurityIdentifier

`func (o *ListingWithTradingVolumeView) GetSecurityIdentifier() string`

GetSecurityIdentifier returns the SecurityIdentifier field if non-nil, zero value otherwise.

### GetSecurityIdentifierOk

`func (o *ListingWithTradingVolumeView) GetSecurityIdentifierOk() (*string, bool)`

GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIdentifier

`func (o *ListingWithTradingVolumeView) SetSecurityIdentifier(v string)`

SetSecurityIdentifier sets SecurityIdentifier field to given value.

### HasSecurityIdentifier

`func (o *ListingWithTradingVolumeView) HasSecurityIdentifier() bool`

HasSecurityIdentifier returns a boolean if a field has been set.

### GetSpreadAbs

`func (o *ListingWithTradingVolumeView) GetSpreadAbs() float32`

GetSpreadAbs returns the SpreadAbs field if non-nil, zero value otherwise.

### GetSpreadAbsOk

`func (o *ListingWithTradingVolumeView) GetSpreadAbsOk() (*float32, bool)`

GetSpreadAbsOk returns a tuple with the SpreadAbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpreadAbs

`func (o *ListingWithTradingVolumeView) SetSpreadAbs(v float32)`

SetSpreadAbs sets SpreadAbs field to given value.

### HasSpreadAbs

`func (o *ListingWithTradingVolumeView) HasSpreadAbs() bool`

HasSpreadAbs returns a boolean if a field has been set.

### GetSpreadPercent

`func (o *ListingWithTradingVolumeView) GetSpreadPercent() float32`

GetSpreadPercent returns the SpreadPercent field if non-nil, zero value otherwise.

### GetSpreadPercentOk

`func (o *ListingWithTradingVolumeView) GetSpreadPercentOk() (*float32, bool)`

GetSpreadPercentOk returns a tuple with the SpreadPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpreadPercent

`func (o *ListingWithTradingVolumeView) SetSpreadPercent(v float32)`

SetSpreadPercent sets SpreadPercent field to given value.

### HasSpreadPercent

`func (o *ListingWithTradingVolumeView) HasSpreadPercent() bool`

HasSpreadPercent returns a boolean if a field has been set.

### GetStartDate

`func (o *ListingWithTradingVolumeView) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListingWithTradingVolumeView) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListingWithTradingVolumeView) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListingWithTradingVolumeView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetType

`func (o *ListingWithTradingVolumeView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListingWithTradingVolumeView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListingWithTradingVolumeView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListingWithTradingVolumeView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVolume

`func (o *ListingWithTradingVolumeView) GetVolume() float32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *ListingWithTradingVolumeView) GetVolumeOk() (*float32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *ListingWithTradingVolumeView) SetVolume(v float32)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *ListingWithTradingVolumeView) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


