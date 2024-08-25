# NewStockView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Listing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**PriceDiffInPercent** | Pointer to **float32** |  | [optional] 

## Methods

### NewNewStockView

`func NewNewStockView() *NewStockView`

NewNewStockView instantiates a new NewStockView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewStockViewWithDefaults

`func NewNewStockViewWithDefaults() *NewStockView`

NewNewStockViewWithDefaults instantiates a new NewStockView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListing

`func (o *NewStockView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *NewStockView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *NewStockView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *NewStockView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetPriceDiffInPercent

`func (o *NewStockView) GetPriceDiffInPercent() float32`

GetPriceDiffInPercent returns the PriceDiffInPercent field if non-nil, zero value otherwise.

### GetPriceDiffInPercentOk

`func (o *NewStockView) GetPriceDiffInPercentOk() (*float32, bool)`

GetPriceDiffInPercentOk returns a tuple with the PriceDiffInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDiffInPercent

`func (o *NewStockView) SetPriceDiffInPercent(v float32)`

SetPriceDiffInPercent sets PriceDiffInPercent field to given value.

### HasPriceDiffInPercent

`func (o *NewStockView) HasPriceDiffInPercent() bool`

HasPriceDiffInPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


