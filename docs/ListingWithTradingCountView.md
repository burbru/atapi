# ListingWithTradingCountView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskPrice** | Pointer to **float32** |  | [optional] 
**AskSize** | Pointer to **int64** |  | [optional] 
**BidPrice** | Pointer to **float32** |  | [optional] 
**BidSize** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
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

### NewListingWithTradingCountView

`func NewListingWithTradingCountView() *ListingWithTradingCountView`

NewListingWithTradingCountView instantiates a new ListingWithTradingCountView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingWithTradingCountViewWithDefaults

`func NewListingWithTradingCountViewWithDefaults() *ListingWithTradingCountView`

NewListingWithTradingCountViewWithDefaults instantiates a new ListingWithTradingCountView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskPrice

`func (o *ListingWithTradingCountView) GetAskPrice() float32`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *ListingWithTradingCountView) GetAskPriceOk() (*float32, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *ListingWithTradingCountView) SetAskPrice(v float32)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *ListingWithTradingCountView) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetAskSize

`func (o *ListingWithTradingCountView) GetAskSize() int64`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *ListingWithTradingCountView) GetAskSizeOk() (*int64, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *ListingWithTradingCountView) SetAskSize(v int64)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *ListingWithTradingCountView) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.

### GetBidPrice

`func (o *ListingWithTradingCountView) GetBidPrice() float32`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *ListingWithTradingCountView) GetBidPriceOk() (*float32, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *ListingWithTradingCountView) SetBidPrice(v float32)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *ListingWithTradingCountView) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetBidSize

`func (o *ListingWithTradingCountView) GetBidSize() int64`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *ListingWithTradingCountView) GetBidSizeOk() (*int64, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *ListingWithTradingCountView) SetBidSize(v int64)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *ListingWithTradingCountView) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetCount

`func (o *ListingWithTradingCountView) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListingWithTradingCountView) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListingWithTradingCountView) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListingWithTradingCountView) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDate

`func (o *ListingWithTradingCountView) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ListingWithTradingCountView) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ListingWithTradingCountView) SetDate(v int64)`

SetDate sets Date field to given value.

### HasDate

`func (o *ListingWithTradingCountView) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListingWithTradingCountView) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListingWithTradingCountView) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListingWithTradingCountView) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListingWithTradingCountView) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetLastPrice

`func (o *ListingWithTradingCountView) GetLastPrice() SecurityPriceView`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *ListingWithTradingCountView) GetLastPriceOk() (*SecurityPriceView, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *ListingWithTradingCountView) SetLastPrice(v SecurityPriceView)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *ListingWithTradingCountView) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetListing

`func (o *ListingWithTradingCountView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *ListingWithTradingCountView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *ListingWithTradingCountView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *ListingWithTradingCountView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetName

`func (o *ListingWithTradingCountView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListingWithTradingCountView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListingWithTradingCountView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListingWithTradingCountView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecurityIdentifier

`func (o *ListingWithTradingCountView) GetSecurityIdentifier() string`

GetSecurityIdentifier returns the SecurityIdentifier field if non-nil, zero value otherwise.

### GetSecurityIdentifierOk

`func (o *ListingWithTradingCountView) GetSecurityIdentifierOk() (*string, bool)`

GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIdentifier

`func (o *ListingWithTradingCountView) SetSecurityIdentifier(v string)`

SetSecurityIdentifier sets SecurityIdentifier field to given value.

### HasSecurityIdentifier

`func (o *ListingWithTradingCountView) HasSecurityIdentifier() bool`

HasSecurityIdentifier returns a boolean if a field has been set.

### GetSpreadAbs

`func (o *ListingWithTradingCountView) GetSpreadAbs() float32`

GetSpreadAbs returns the SpreadAbs field if non-nil, zero value otherwise.

### GetSpreadAbsOk

`func (o *ListingWithTradingCountView) GetSpreadAbsOk() (*float32, bool)`

GetSpreadAbsOk returns a tuple with the SpreadAbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpreadAbs

`func (o *ListingWithTradingCountView) SetSpreadAbs(v float32)`

SetSpreadAbs sets SpreadAbs field to given value.

### HasSpreadAbs

`func (o *ListingWithTradingCountView) HasSpreadAbs() bool`

HasSpreadAbs returns a boolean if a field has been set.

### GetSpreadPercent

`func (o *ListingWithTradingCountView) GetSpreadPercent() float32`

GetSpreadPercent returns the SpreadPercent field if non-nil, zero value otherwise.

### GetSpreadPercentOk

`func (o *ListingWithTradingCountView) GetSpreadPercentOk() (*float32, bool)`

GetSpreadPercentOk returns a tuple with the SpreadPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpreadPercent

`func (o *ListingWithTradingCountView) SetSpreadPercent(v float32)`

SetSpreadPercent sets SpreadPercent field to given value.

### HasSpreadPercent

`func (o *ListingWithTradingCountView) HasSpreadPercent() bool`

HasSpreadPercent returns a boolean if a field has been set.

### GetStartDate

`func (o *ListingWithTradingCountView) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListingWithTradingCountView) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListingWithTradingCountView) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListingWithTradingCountView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetType

`func (o *ListingWithTradingCountView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListingWithTradingCountView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListingWithTradingCountView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListingWithTradingCountView) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


