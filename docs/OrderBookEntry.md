# OrderBookEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**PriceLimit** | Pointer to **float32** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 

## Methods

### NewOrderBookEntry

`func NewOrderBookEntry() *OrderBookEntry`

NewOrderBookEntry instantiates a new OrderBookEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderBookEntryWithDefaults

`func NewOrderBookEntryWithDefaults() *OrderBookEntry`

NewOrderBookEntryWithDefaults instantiates a new OrderBookEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *OrderBookEntry) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OrderBookEntry) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OrderBookEntry) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *OrderBookEntry) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPriceLimit

`func (o *OrderBookEntry) GetPriceLimit() float32`

GetPriceLimit returns the PriceLimit field if non-nil, zero value otherwise.

### GetPriceLimitOk

`func (o *OrderBookEntry) GetPriceLimitOk() (*float32, bool)`

GetPriceLimitOk returns a tuple with the PriceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLimit

`func (o *OrderBookEntry) SetPriceLimit(v float32)`

SetPriceLimit sets PriceLimit field to given value.

### HasPriceLimit

`func (o *OrderBookEntry) HasPriceLimit() bool`

HasPriceLimit returns a boolean if a field has been set.

### GetSize

`func (o *OrderBookEntry) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *OrderBookEntry) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *OrderBookEntry) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *OrderBookEntry) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


