# OrderCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckResult** | Pointer to [**SecurityOrderCheckResult**](SecurityOrderCheckResult.md) |  | [optional] 
**ExecutionPrice** | Pointer to **float32** |  | [optional] 
**ExecutionVolume** | Pointer to **float32** |  | [optional] 
**NumberOfShares** | Pointer to **int64** |  | [optional] 
**Spread** | Pointer to [**PriceSpreadView**](PriceSpreadView.md) |  | [optional] 
**UncommittedCash** | Pointer to **float32** |  | [optional] 

## Methods

### NewOrderCheck

`func NewOrderCheck() *OrderCheck`

NewOrderCheck instantiates a new OrderCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCheckWithDefaults

`func NewOrderCheckWithDefaults() *OrderCheck`

NewOrderCheckWithDefaults instantiates a new OrderCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckResult

`func (o *OrderCheck) GetCheckResult() SecurityOrderCheckResult`

GetCheckResult returns the CheckResult field if non-nil, zero value otherwise.

### GetCheckResultOk

`func (o *OrderCheck) GetCheckResultOk() (*SecurityOrderCheckResult, bool)`

GetCheckResultOk returns a tuple with the CheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResult

`func (o *OrderCheck) SetCheckResult(v SecurityOrderCheckResult)`

SetCheckResult sets CheckResult field to given value.

### HasCheckResult

`func (o *OrderCheck) HasCheckResult() bool`

HasCheckResult returns a boolean if a field has been set.

### GetExecutionPrice

`func (o *OrderCheck) GetExecutionPrice() float32`

GetExecutionPrice returns the ExecutionPrice field if non-nil, zero value otherwise.

### GetExecutionPriceOk

`func (o *OrderCheck) GetExecutionPriceOk() (*float32, bool)`

GetExecutionPriceOk returns a tuple with the ExecutionPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionPrice

`func (o *OrderCheck) SetExecutionPrice(v float32)`

SetExecutionPrice sets ExecutionPrice field to given value.

### HasExecutionPrice

`func (o *OrderCheck) HasExecutionPrice() bool`

HasExecutionPrice returns a boolean if a field has been set.

### GetExecutionVolume

`func (o *OrderCheck) GetExecutionVolume() float32`

GetExecutionVolume returns the ExecutionVolume field if non-nil, zero value otherwise.

### GetExecutionVolumeOk

`func (o *OrderCheck) GetExecutionVolumeOk() (*float32, bool)`

GetExecutionVolumeOk returns a tuple with the ExecutionVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionVolume

`func (o *OrderCheck) SetExecutionVolume(v float32)`

SetExecutionVolume sets ExecutionVolume field to given value.

### HasExecutionVolume

`func (o *OrderCheck) HasExecutionVolume() bool`

HasExecutionVolume returns a boolean if a field has been set.

### GetNumberOfShares

`func (o *OrderCheck) GetNumberOfShares() int64`

GetNumberOfShares returns the NumberOfShares field if non-nil, zero value otherwise.

### GetNumberOfSharesOk

`func (o *OrderCheck) GetNumberOfSharesOk() (*int64, bool)`

GetNumberOfSharesOk returns a tuple with the NumberOfShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfShares

`func (o *OrderCheck) SetNumberOfShares(v int64)`

SetNumberOfShares sets NumberOfShares field to given value.

### HasNumberOfShares

`func (o *OrderCheck) HasNumberOfShares() bool`

HasNumberOfShares returns a boolean if a field has been set.

### GetSpread

`func (o *OrderCheck) GetSpread() PriceSpreadView`

GetSpread returns the Spread field if non-nil, zero value otherwise.

### GetSpreadOk

`func (o *OrderCheck) GetSpreadOk() (*PriceSpreadView, bool)`

GetSpreadOk returns a tuple with the Spread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpread

`func (o *OrderCheck) SetSpread(v PriceSpreadView)`

SetSpread sets Spread field to given value.

### HasSpread

`func (o *OrderCheck) HasSpread() bool`

HasSpread returns a boolean if a field has been set.

### GetUncommittedCash

`func (o *OrderCheck) GetUncommittedCash() float32`

GetUncommittedCash returns the UncommittedCash field if non-nil, zero value otherwise.

### GetUncommittedCashOk

`func (o *OrderCheck) GetUncommittedCashOk() (*float32, bool)`

GetUncommittedCashOk returns a tuple with the UncommittedCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncommittedCash

`func (o *OrderCheck) SetUncommittedCash(v float32)`

SetUncommittedCash sets UncommittedCash field to given value.

### HasUncommittedCash

`func (o *OrderCheck) HasUncommittedCash() bool`

HasUncommittedCash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


