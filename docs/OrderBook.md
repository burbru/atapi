# OrderBook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyEntries** | Pointer to [**[]OrderBookEntry**](OrderBookEntry.md) |  | [optional] 
**MaxBuySize** | Pointer to **int64** |  | [optional] 
**MaxSellSize** | Pointer to **int64** |  | [optional] 
**SellEntries** | Pointer to [**[]OrderBookEntry**](OrderBookEntry.md) |  | [optional] 

## Methods

### NewOrderBook

`func NewOrderBook() *OrderBook`

NewOrderBook instantiates a new OrderBook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderBookWithDefaults

`func NewOrderBookWithDefaults() *OrderBook`

NewOrderBookWithDefaults instantiates a new OrderBook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyEntries

`func (o *OrderBook) GetBuyEntries() []OrderBookEntry`

GetBuyEntries returns the BuyEntries field if non-nil, zero value otherwise.

### GetBuyEntriesOk

`func (o *OrderBook) GetBuyEntriesOk() (*[]OrderBookEntry, bool)`

GetBuyEntriesOk returns a tuple with the BuyEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyEntries

`func (o *OrderBook) SetBuyEntries(v []OrderBookEntry)`

SetBuyEntries sets BuyEntries field to given value.

### HasBuyEntries

`func (o *OrderBook) HasBuyEntries() bool`

HasBuyEntries returns a boolean if a field has been set.

### GetMaxBuySize

`func (o *OrderBook) GetMaxBuySize() int64`

GetMaxBuySize returns the MaxBuySize field if non-nil, zero value otherwise.

### GetMaxBuySizeOk

`func (o *OrderBook) GetMaxBuySizeOk() (*int64, bool)`

GetMaxBuySizeOk returns a tuple with the MaxBuySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBuySize

`func (o *OrderBook) SetMaxBuySize(v int64)`

SetMaxBuySize sets MaxBuySize field to given value.

### HasMaxBuySize

`func (o *OrderBook) HasMaxBuySize() bool`

HasMaxBuySize returns a boolean if a field has been set.

### GetMaxSellSize

`func (o *OrderBook) GetMaxSellSize() int64`

GetMaxSellSize returns the MaxSellSize field if non-nil, zero value otherwise.

### GetMaxSellSizeOk

`func (o *OrderBook) GetMaxSellSizeOk() (*int64, bool)`

GetMaxSellSizeOk returns a tuple with the MaxSellSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSellSize

`func (o *OrderBook) SetMaxSellSize(v int64)`

SetMaxSellSize sets MaxSellSize field to given value.

### HasMaxSellSize

`func (o *OrderBook) HasMaxSellSize() bool`

HasMaxSellSize returns a boolean if a field has been set.

### GetSellEntries

`func (o *OrderBook) GetSellEntries() []OrderBookEntry`

GetSellEntries returns the SellEntries field if non-nil, zero value otherwise.

### GetSellEntriesOk

`func (o *OrderBook) GetSellEntriesOk() (*[]OrderBookEntry, bool)`

GetSellEntriesOk returns a tuple with the SellEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellEntries

`func (o *OrderBook) SetSellEntries(v []OrderBookEntry)`

SetSellEntries sets SellEntries field to given value.

### HasSellEntries

`func (o *OrderBook) HasSellEntries() bool`

HasSellEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


