# PriceSpreadListingView

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

## Methods

### NewPriceSpreadListingView

`func NewPriceSpreadListingView() *PriceSpreadListingView`

NewPriceSpreadListingView instantiates a new PriceSpreadListingView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceSpreadListingViewWithDefaults

`func NewPriceSpreadListingViewWithDefaults() *PriceSpreadListingView`

NewPriceSpreadListingViewWithDefaults instantiates a new PriceSpreadListingView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskPrice

`func (o *PriceSpreadListingView) GetAskPrice() float32`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *PriceSpreadListingView) GetAskPriceOk() (*float32, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *PriceSpreadListingView) SetAskPrice(v float32)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *PriceSpreadListingView) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetAskSize

`func (o *PriceSpreadListingView) GetAskSize() int64`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *PriceSpreadListingView) GetAskSizeOk() (*int64, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *PriceSpreadListingView) SetAskSize(v int64)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *PriceSpreadListingView) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.

### GetBidPrice

`func (o *PriceSpreadListingView) GetBidPrice() float32`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *PriceSpreadListingView) GetBidPriceOk() (*float32, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *PriceSpreadListingView) SetBidPrice(v float32)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *PriceSpreadListingView) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetBidSize

`func (o *PriceSpreadListingView) GetBidSize() int64`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *PriceSpreadListingView) GetBidSizeOk() (*int64, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *PriceSpreadListingView) SetBidSize(v int64)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *PriceSpreadListingView) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetDate

`func (o *PriceSpreadListingView) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PriceSpreadListingView) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PriceSpreadListingView) SetDate(v int64)`

SetDate sets Date field to given value.

### HasDate

`func (o *PriceSpreadListingView) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetEndDate

`func (o *PriceSpreadListingView) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PriceSpreadListingView) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PriceSpreadListingView) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PriceSpreadListingView) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetLastPrice

`func (o *PriceSpreadListingView) GetLastPrice() SecurityPriceView`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *PriceSpreadListingView) GetLastPriceOk() (*SecurityPriceView, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *PriceSpreadListingView) SetLastPrice(v SecurityPriceView)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *PriceSpreadListingView) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetListing

`func (o *PriceSpreadListingView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *PriceSpreadListingView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *PriceSpreadListingView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *PriceSpreadListingView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetName

`func (o *PriceSpreadListingView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PriceSpreadListingView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PriceSpreadListingView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PriceSpreadListingView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecurityIdentifier

`func (o *PriceSpreadListingView) GetSecurityIdentifier() string`

GetSecurityIdentifier returns the SecurityIdentifier field if non-nil, zero value otherwise.

### GetSecurityIdentifierOk

`func (o *PriceSpreadListingView) GetSecurityIdentifierOk() (*string, bool)`

GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIdentifier

`func (o *PriceSpreadListingView) SetSecurityIdentifier(v string)`

SetSecurityIdentifier sets SecurityIdentifier field to given value.

### HasSecurityIdentifier

`func (o *PriceSpreadListingView) HasSecurityIdentifier() bool`

HasSecurityIdentifier returns a boolean if a field has been set.

### GetSpreadAbs

`func (o *PriceSpreadListingView) GetSpreadAbs() float32`

GetSpreadAbs returns the SpreadAbs field if non-nil, zero value otherwise.

### GetSpreadAbsOk

`func (o *PriceSpreadListingView) GetSpreadAbsOk() (*float32, bool)`

GetSpreadAbsOk returns a tuple with the SpreadAbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpreadAbs

`func (o *PriceSpreadListingView) SetSpreadAbs(v float32)`

SetSpreadAbs sets SpreadAbs field to given value.

### HasSpreadAbs

`func (o *PriceSpreadListingView) HasSpreadAbs() bool`

HasSpreadAbs returns a boolean if a field has been set.

### GetSpreadPercent

`func (o *PriceSpreadListingView) GetSpreadPercent() float32`

GetSpreadPercent returns the SpreadPercent field if non-nil, zero value otherwise.

### GetSpreadPercentOk

`func (o *PriceSpreadListingView) GetSpreadPercentOk() (*float32, bool)`

GetSpreadPercentOk returns a tuple with the SpreadPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpreadPercent

`func (o *PriceSpreadListingView) SetSpreadPercent(v float32)`

SetSpreadPercent sets SpreadPercent field to given value.

### HasSpreadPercent

`func (o *PriceSpreadListingView) HasSpreadPercent() bool`

HasSpreadPercent returns a boolean if a field has been set.

### GetStartDate

`func (o *PriceSpreadListingView) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PriceSpreadListingView) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PriceSpreadListingView) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PriceSpreadListingView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetType

`func (o *PriceSpreadListingView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceSpreadListingView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceSpreadListingView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PriceSpreadListingView) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


