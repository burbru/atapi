# ListingMarketFilterResultView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Listing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**Price** | Pointer to [**PriceSpreadMininalView**](PriceSpreadMininalView.md) |  | [optional] 

## Methods

### NewListingMarketFilterResultView

`func NewListingMarketFilterResultView() *ListingMarketFilterResultView`

NewListingMarketFilterResultView instantiates a new ListingMarketFilterResultView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingMarketFilterResultViewWithDefaults

`func NewListingMarketFilterResultViewWithDefaults() *ListingMarketFilterResultView`

NewListingMarketFilterResultViewWithDefaults instantiates a new ListingMarketFilterResultView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListing

`func (o *ListingMarketFilterResultView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *ListingMarketFilterResultView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *ListingMarketFilterResultView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *ListingMarketFilterResultView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetPrice

`func (o *ListingMarketFilterResultView) GetPrice() PriceSpreadMininalView`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ListingMarketFilterResultView) GetPriceOk() (*PriceSpreadMininalView, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ListingMarketFilterResultView) SetPrice(v PriceSpreadMininalView)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ListingMarketFilterResultView) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


