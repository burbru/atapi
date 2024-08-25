# HistorizedListingDataView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskPrice** | Pointer to **float32** |  | [optional] 
**BidPrice** | Pointer to **float32** |  | [optional] 
**ClosePrice** | Pointer to **float32** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**HighPrice** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LowPrice** | Pointer to **float32** |  | [optional] 
**OpenPrice** | Pointer to **float32** |  | [optional] 
**OutstandingShares** | Pointer to **int64** |  | [optional] 
**SharesInBuys** | Pointer to **int64** |  | [optional] 
**SharesInSells** | Pointer to **int64** |  | [optional] 
**TradeVolume** | Pointer to **float32** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewHistorizedListingDataView

`func NewHistorizedListingDataView() *HistorizedListingDataView`

NewHistorizedListingDataView instantiates a new HistorizedListingDataView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistorizedListingDataViewWithDefaults

`func NewHistorizedListingDataViewWithDefaults() *HistorizedListingDataView`

NewHistorizedListingDataViewWithDefaults instantiates a new HistorizedListingDataView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskPrice

`func (o *HistorizedListingDataView) GetAskPrice() float32`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *HistorizedListingDataView) GetAskPriceOk() (*float32, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *HistorizedListingDataView) SetAskPrice(v float32)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *HistorizedListingDataView) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetBidPrice

`func (o *HistorizedListingDataView) GetBidPrice() float32`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *HistorizedListingDataView) GetBidPriceOk() (*float32, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *HistorizedListingDataView) SetBidPrice(v float32)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *HistorizedListingDataView) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetClosePrice

`func (o *HistorizedListingDataView) GetClosePrice() float32`

GetClosePrice returns the ClosePrice field if non-nil, zero value otherwise.

### GetClosePriceOk

`func (o *HistorizedListingDataView) GetClosePriceOk() (*float32, bool)`

GetClosePriceOk returns a tuple with the ClosePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePrice

`func (o *HistorizedListingDataView) SetClosePrice(v float32)`

SetClosePrice sets ClosePrice field to given value.

### HasClosePrice

`func (o *HistorizedListingDataView) HasClosePrice() bool`

HasClosePrice returns a boolean if a field has been set.

### GetDate

`func (o *HistorizedListingDataView) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *HistorizedListingDataView) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *HistorizedListingDataView) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *HistorizedListingDataView) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetHighPrice

`func (o *HistorizedListingDataView) GetHighPrice() float32`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *HistorizedListingDataView) GetHighPriceOk() (*float32, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *HistorizedListingDataView) SetHighPrice(v float32)`

SetHighPrice sets HighPrice field to given value.

### HasHighPrice

`func (o *HistorizedListingDataView) HasHighPrice() bool`

HasHighPrice returns a boolean if a field has been set.

### GetId

`func (o *HistorizedListingDataView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistorizedListingDataView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistorizedListingDataView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HistorizedListingDataView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLowPrice

`func (o *HistorizedListingDataView) GetLowPrice() float32`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *HistorizedListingDataView) GetLowPriceOk() (*float32, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *HistorizedListingDataView) SetLowPrice(v float32)`

SetLowPrice sets LowPrice field to given value.

### HasLowPrice

`func (o *HistorizedListingDataView) HasLowPrice() bool`

HasLowPrice returns a boolean if a field has been set.

### GetOpenPrice

`func (o *HistorizedListingDataView) GetOpenPrice() float32`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *HistorizedListingDataView) GetOpenPriceOk() (*float32, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *HistorizedListingDataView) SetOpenPrice(v float32)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *HistorizedListingDataView) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetOutstandingShares

`func (o *HistorizedListingDataView) GetOutstandingShares() int64`

GetOutstandingShares returns the OutstandingShares field if non-nil, zero value otherwise.

### GetOutstandingSharesOk

`func (o *HistorizedListingDataView) GetOutstandingSharesOk() (*int64, bool)`

GetOutstandingSharesOk returns a tuple with the OutstandingShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutstandingShares

`func (o *HistorizedListingDataView) SetOutstandingShares(v int64)`

SetOutstandingShares sets OutstandingShares field to given value.

### HasOutstandingShares

`func (o *HistorizedListingDataView) HasOutstandingShares() bool`

HasOutstandingShares returns a boolean if a field has been set.

### GetSharesInBuys

`func (o *HistorizedListingDataView) GetSharesInBuys() int64`

GetSharesInBuys returns the SharesInBuys field if non-nil, zero value otherwise.

### GetSharesInBuysOk

`func (o *HistorizedListingDataView) GetSharesInBuysOk() (*int64, bool)`

GetSharesInBuysOk returns a tuple with the SharesInBuys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesInBuys

`func (o *HistorizedListingDataView) SetSharesInBuys(v int64)`

SetSharesInBuys sets SharesInBuys field to given value.

### HasSharesInBuys

`func (o *HistorizedListingDataView) HasSharesInBuys() bool`

HasSharesInBuys returns a boolean if a field has been set.

### GetSharesInSells

`func (o *HistorizedListingDataView) GetSharesInSells() int64`

GetSharesInSells returns the SharesInSells field if non-nil, zero value otherwise.

### GetSharesInSellsOk

`func (o *HistorizedListingDataView) GetSharesInSellsOk() (*int64, bool)`

GetSharesInSellsOk returns a tuple with the SharesInSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesInSells

`func (o *HistorizedListingDataView) SetSharesInSells(v int64)`

SetSharesInSells sets SharesInSells field to given value.

### HasSharesInSells

`func (o *HistorizedListingDataView) HasSharesInSells() bool`

HasSharesInSells returns a boolean if a field has been set.

### GetTradeVolume

`func (o *HistorizedListingDataView) GetTradeVolume() float32`

GetTradeVolume returns the TradeVolume field if non-nil, zero value otherwise.

### GetTradeVolumeOk

`func (o *HistorizedListingDataView) GetTradeVolumeOk() (*float32, bool)`

GetTradeVolumeOk returns a tuple with the TradeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeVolume

`func (o *HistorizedListingDataView) SetTradeVolume(v float32)`

SetTradeVolume sets TradeVolume field to given value.

### HasTradeVolume

`func (o *HistorizedListingDataView) HasTradeVolume() bool`

HasTradeVolume returns a boolean if a field has been set.

### GetVersion

`func (o *HistorizedListingDataView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HistorizedListingDataView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HistorizedListingDataView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HistorizedListingDataView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


