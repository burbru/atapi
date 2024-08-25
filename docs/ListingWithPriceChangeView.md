# ListingWithPriceChangeView

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
**PriceChangeInPercent** | Pointer to **float32** |  | [optional] 
**SecurityIdentifier** | Pointer to **string** |  | [optional] 
**SpreadAbs** | Pointer to **float32** |  | [optional] 
**SpreadPercent** | Pointer to **float32** |  | [optional] 
**StartDate** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewListingWithPriceChangeView

`func NewListingWithPriceChangeView() *ListingWithPriceChangeView`

NewListingWithPriceChangeView instantiates a new ListingWithPriceChangeView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingWithPriceChangeViewWithDefaults

`func NewListingWithPriceChangeViewWithDefaults() *ListingWithPriceChangeView`

NewListingWithPriceChangeViewWithDefaults instantiates a new ListingWithPriceChangeView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskPrice

`func (o *ListingWithPriceChangeView) GetAskPrice() float32`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *ListingWithPriceChangeView) GetAskPriceOk() (*float32, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *ListingWithPriceChangeView) SetAskPrice(v float32)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *ListingWithPriceChangeView) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetAskSize

`func (o *ListingWithPriceChangeView) GetAskSize() int64`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *ListingWithPriceChangeView) GetAskSizeOk() (*int64, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *ListingWithPriceChangeView) SetAskSize(v int64)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *ListingWithPriceChangeView) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.

### GetBidPrice

`func (o *ListingWithPriceChangeView) GetBidPrice() float32`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *ListingWithPriceChangeView) GetBidPriceOk() (*float32, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *ListingWithPriceChangeView) SetBidPrice(v float32)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *ListingWithPriceChangeView) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetBidSize

`func (o *ListingWithPriceChangeView) GetBidSize() int64`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *ListingWithPriceChangeView) GetBidSizeOk() (*int64, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *ListingWithPriceChangeView) SetBidSize(v int64)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *ListingWithPriceChangeView) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetDate

`func (o *ListingWithPriceChangeView) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ListingWithPriceChangeView) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ListingWithPriceChangeView) SetDate(v int64)`

SetDate sets Date field to given value.

### HasDate

`func (o *ListingWithPriceChangeView) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListingWithPriceChangeView) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListingWithPriceChangeView) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListingWithPriceChangeView) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListingWithPriceChangeView) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetLastPrice

`func (o *ListingWithPriceChangeView) GetLastPrice() SecurityPriceView`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *ListingWithPriceChangeView) GetLastPriceOk() (*SecurityPriceView, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *ListingWithPriceChangeView) SetLastPrice(v SecurityPriceView)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *ListingWithPriceChangeView) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetListing

`func (o *ListingWithPriceChangeView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *ListingWithPriceChangeView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *ListingWithPriceChangeView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *ListingWithPriceChangeView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetName

`func (o *ListingWithPriceChangeView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListingWithPriceChangeView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListingWithPriceChangeView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListingWithPriceChangeView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriceChangeInPercent

`func (o *ListingWithPriceChangeView) GetPriceChangeInPercent() float32`

GetPriceChangeInPercent returns the PriceChangeInPercent field if non-nil, zero value otherwise.

### GetPriceChangeInPercentOk

`func (o *ListingWithPriceChangeView) GetPriceChangeInPercentOk() (*float32, bool)`

GetPriceChangeInPercentOk returns a tuple with the PriceChangeInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChangeInPercent

`func (o *ListingWithPriceChangeView) SetPriceChangeInPercent(v float32)`

SetPriceChangeInPercent sets PriceChangeInPercent field to given value.

### HasPriceChangeInPercent

`func (o *ListingWithPriceChangeView) HasPriceChangeInPercent() bool`

HasPriceChangeInPercent returns a boolean if a field has been set.

### GetSecurityIdentifier

`func (o *ListingWithPriceChangeView) GetSecurityIdentifier() string`

GetSecurityIdentifier returns the SecurityIdentifier field if non-nil, zero value otherwise.

### GetSecurityIdentifierOk

`func (o *ListingWithPriceChangeView) GetSecurityIdentifierOk() (*string, bool)`

GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIdentifier

`func (o *ListingWithPriceChangeView) SetSecurityIdentifier(v string)`

SetSecurityIdentifier sets SecurityIdentifier field to given value.

### HasSecurityIdentifier

`func (o *ListingWithPriceChangeView) HasSecurityIdentifier() bool`

HasSecurityIdentifier returns a boolean if a field has been set.

### GetSpreadAbs

`func (o *ListingWithPriceChangeView) GetSpreadAbs() float32`

GetSpreadAbs returns the SpreadAbs field if non-nil, zero value otherwise.

### GetSpreadAbsOk

`func (o *ListingWithPriceChangeView) GetSpreadAbsOk() (*float32, bool)`

GetSpreadAbsOk returns a tuple with the SpreadAbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpreadAbs

`func (o *ListingWithPriceChangeView) SetSpreadAbs(v float32)`

SetSpreadAbs sets SpreadAbs field to given value.

### HasSpreadAbs

`func (o *ListingWithPriceChangeView) HasSpreadAbs() bool`

HasSpreadAbs returns a boolean if a field has been set.

### GetSpreadPercent

`func (o *ListingWithPriceChangeView) GetSpreadPercent() float32`

GetSpreadPercent returns the SpreadPercent field if non-nil, zero value otherwise.

### GetSpreadPercentOk

`func (o *ListingWithPriceChangeView) GetSpreadPercentOk() (*float32, bool)`

GetSpreadPercentOk returns a tuple with the SpreadPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpreadPercent

`func (o *ListingWithPriceChangeView) SetSpreadPercent(v float32)`

SetSpreadPercent sets SpreadPercent field to given value.

### HasSpreadPercent

`func (o *ListingWithPriceChangeView) HasSpreadPercent() bool`

HasSpreadPercent returns a boolean if a field has been set.

### GetStartDate

`func (o *ListingWithPriceChangeView) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListingWithPriceChangeView) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListingWithPriceChangeView) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListingWithPriceChangeView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetType

`func (o *ListingWithPriceChangeView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListingWithPriceChangeView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListingWithPriceChangeView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListingWithPriceChangeView) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


