# PriceSpreadView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskPrice** | Pointer to **float32** |  | [optional] 
**AskSize** | Pointer to **int64** |  | [optional] 
**BidPrice** | Pointer to **float32** |  | [optional] 
**BidSize** | Pointer to **int64** |  | [optional] 
**Date** | Pointer to **int64** |  | [optional] 
**LastPrice** | Pointer to [**SecurityPriceView**](SecurityPriceView.md) |  | [optional] 
**MaxBidPrice** | Pointer to **float32** |  | [optional] 
**MinAskPrice** | Pointer to **float32** |  | [optional] 
**SpreadAbs** | Pointer to **float32** |  | [optional] 
**SpreadPercent** | Pointer to **float32** |  | [optional] 

## Methods

### NewPriceSpreadView

`func NewPriceSpreadView() *PriceSpreadView`

NewPriceSpreadView instantiates a new PriceSpreadView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceSpreadViewWithDefaults

`func NewPriceSpreadViewWithDefaults() *PriceSpreadView`

NewPriceSpreadViewWithDefaults instantiates a new PriceSpreadView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskPrice

`func (o *PriceSpreadView) GetAskPrice() float32`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *PriceSpreadView) GetAskPriceOk() (*float32, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *PriceSpreadView) SetAskPrice(v float32)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *PriceSpreadView) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetAskSize

`func (o *PriceSpreadView) GetAskSize() int64`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *PriceSpreadView) GetAskSizeOk() (*int64, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *PriceSpreadView) SetAskSize(v int64)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *PriceSpreadView) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.

### GetBidPrice

`func (o *PriceSpreadView) GetBidPrice() float32`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *PriceSpreadView) GetBidPriceOk() (*float32, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *PriceSpreadView) SetBidPrice(v float32)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *PriceSpreadView) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetBidSize

`func (o *PriceSpreadView) GetBidSize() int64`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *PriceSpreadView) GetBidSizeOk() (*int64, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *PriceSpreadView) SetBidSize(v int64)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *PriceSpreadView) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetDate

`func (o *PriceSpreadView) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PriceSpreadView) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PriceSpreadView) SetDate(v int64)`

SetDate sets Date field to given value.

### HasDate

`func (o *PriceSpreadView) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetLastPrice

`func (o *PriceSpreadView) GetLastPrice() SecurityPriceView`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *PriceSpreadView) GetLastPriceOk() (*SecurityPriceView, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *PriceSpreadView) SetLastPrice(v SecurityPriceView)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *PriceSpreadView) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetMaxBidPrice

`func (o *PriceSpreadView) GetMaxBidPrice() float32`

GetMaxBidPrice returns the MaxBidPrice field if non-nil, zero value otherwise.

### GetMaxBidPriceOk

`func (o *PriceSpreadView) GetMaxBidPriceOk() (*float32, bool)`

GetMaxBidPriceOk returns a tuple with the MaxBidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBidPrice

`func (o *PriceSpreadView) SetMaxBidPrice(v float32)`

SetMaxBidPrice sets MaxBidPrice field to given value.

### HasMaxBidPrice

`func (o *PriceSpreadView) HasMaxBidPrice() bool`

HasMaxBidPrice returns a boolean if a field has been set.

### GetMinAskPrice

`func (o *PriceSpreadView) GetMinAskPrice() float32`

GetMinAskPrice returns the MinAskPrice field if non-nil, zero value otherwise.

### GetMinAskPriceOk

`func (o *PriceSpreadView) GetMinAskPriceOk() (*float32, bool)`

GetMinAskPriceOk returns a tuple with the MinAskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAskPrice

`func (o *PriceSpreadView) SetMinAskPrice(v float32)`

SetMinAskPrice sets MinAskPrice field to given value.

### HasMinAskPrice

`func (o *PriceSpreadView) HasMinAskPrice() bool`

HasMinAskPrice returns a boolean if a field has been set.

### GetSpreadAbs

`func (o *PriceSpreadView) GetSpreadAbs() float32`

GetSpreadAbs returns the SpreadAbs field if non-nil, zero value otherwise.

### GetSpreadAbsOk

`func (o *PriceSpreadView) GetSpreadAbsOk() (*float32, bool)`

GetSpreadAbsOk returns a tuple with the SpreadAbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpreadAbs

`func (o *PriceSpreadView) SetSpreadAbs(v float32)`

SetSpreadAbs sets SpreadAbs field to given value.

### HasSpreadAbs

`func (o *PriceSpreadView) HasSpreadAbs() bool`

HasSpreadAbs returns a boolean if a field has been set.

### GetSpreadPercent

`func (o *PriceSpreadView) GetSpreadPercent() float32`

GetSpreadPercent returns the SpreadPercent field if non-nil, zero value otherwise.

### GetSpreadPercentOk

`func (o *PriceSpreadView) GetSpreadPercentOk() (*float32, bool)`

GetSpreadPercentOk returns a tuple with the SpreadPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpreadPercent

`func (o *PriceSpreadView) SetSpreadPercent(v float32)`

SetSpreadPercent sets SpreadPercent field to given value.

### HasSpreadPercent

`func (o *PriceSpreadView) HasSpreadPercent() bool`

HasSpreadPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


