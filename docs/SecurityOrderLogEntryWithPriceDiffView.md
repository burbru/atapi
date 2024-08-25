# SecurityOrderLogEntryWithPriceDiffView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Listing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**Log** | Pointer to [**SecurityOrderLogEntryView**](SecurityOrderLogEntryView.md) |  | [optional] 
**PriceChangeAbs** | Pointer to **float32** |  | [optional] 
**PriceChangeRelInPercent** | Pointer to **float32** |  | [optional] 

## Methods

### NewSecurityOrderLogEntryWithPriceDiffView

`func NewSecurityOrderLogEntryWithPriceDiffView() *SecurityOrderLogEntryWithPriceDiffView`

NewSecurityOrderLogEntryWithPriceDiffView instantiates a new SecurityOrderLogEntryWithPriceDiffView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityOrderLogEntryWithPriceDiffViewWithDefaults

`func NewSecurityOrderLogEntryWithPriceDiffViewWithDefaults() *SecurityOrderLogEntryWithPriceDiffView`

NewSecurityOrderLogEntryWithPriceDiffViewWithDefaults instantiates a new SecurityOrderLogEntryWithPriceDiffView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListing

`func (o *SecurityOrderLogEntryWithPriceDiffView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *SecurityOrderLogEntryWithPriceDiffView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *SecurityOrderLogEntryWithPriceDiffView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *SecurityOrderLogEntryWithPriceDiffView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetLog

`func (o *SecurityOrderLogEntryWithPriceDiffView) GetLog() SecurityOrderLogEntryView`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *SecurityOrderLogEntryWithPriceDiffView) GetLogOk() (*SecurityOrderLogEntryView, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *SecurityOrderLogEntryWithPriceDiffView) SetLog(v SecurityOrderLogEntryView)`

SetLog sets Log field to given value.

### HasLog

`func (o *SecurityOrderLogEntryWithPriceDiffView) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetPriceChangeAbs

`func (o *SecurityOrderLogEntryWithPriceDiffView) GetPriceChangeAbs() float32`

GetPriceChangeAbs returns the PriceChangeAbs field if non-nil, zero value otherwise.

### GetPriceChangeAbsOk

`func (o *SecurityOrderLogEntryWithPriceDiffView) GetPriceChangeAbsOk() (*float32, bool)`

GetPriceChangeAbsOk returns a tuple with the PriceChangeAbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChangeAbs

`func (o *SecurityOrderLogEntryWithPriceDiffView) SetPriceChangeAbs(v float32)`

SetPriceChangeAbs sets PriceChangeAbs field to given value.

### HasPriceChangeAbs

`func (o *SecurityOrderLogEntryWithPriceDiffView) HasPriceChangeAbs() bool`

HasPriceChangeAbs returns a boolean if a field has been set.

### GetPriceChangeRelInPercent

`func (o *SecurityOrderLogEntryWithPriceDiffView) GetPriceChangeRelInPercent() float32`

GetPriceChangeRelInPercent returns the PriceChangeRelInPercent field if non-nil, zero value otherwise.

### GetPriceChangeRelInPercentOk

`func (o *SecurityOrderLogEntryWithPriceDiffView) GetPriceChangeRelInPercentOk() (*float32, bool)`

GetPriceChangeRelInPercentOk returns a tuple with the PriceChangeRelInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChangeRelInPercent

`func (o *SecurityOrderLogEntryWithPriceDiffView) SetPriceChangeRelInPercent(v float32)`

SetPriceChangeRelInPercent sets PriceChangeRelInPercent field to given value.

### HasPriceChangeRelInPercent

`func (o *SecurityOrderLogEntryWithPriceDiffView) HasPriceChangeRelInPercent() bool`

HasPriceChangeRelInPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


