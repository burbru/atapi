# CentralbankReservesView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankingLicense** | Pointer to [**BankingLicenseView**](BankingLicenseView.md) |  | [optional] 
**CashHolding** | Pointer to **float32** |  | [optional] 
**CoinsForNextBoost** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InterestRateBoost** | Pointer to **float32** |  | [optional] 
**MaxCentralBankLoans** | Pointer to **float32** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewCentralbankReservesView

`func NewCentralbankReservesView() *CentralbankReservesView`

NewCentralbankReservesView instantiates a new CentralbankReservesView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCentralbankReservesViewWithDefaults

`func NewCentralbankReservesViewWithDefaults() *CentralbankReservesView`

NewCentralbankReservesViewWithDefaults instantiates a new CentralbankReservesView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankingLicense

`func (o *CentralbankReservesView) GetBankingLicense() BankingLicenseView`

GetBankingLicense returns the BankingLicense field if non-nil, zero value otherwise.

### GetBankingLicenseOk

`func (o *CentralbankReservesView) GetBankingLicenseOk() (*BankingLicenseView, bool)`

GetBankingLicenseOk returns a tuple with the BankingLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankingLicense

`func (o *CentralbankReservesView) SetBankingLicense(v BankingLicenseView)`

SetBankingLicense sets BankingLicense field to given value.

### HasBankingLicense

`func (o *CentralbankReservesView) HasBankingLicense() bool`

HasBankingLicense returns a boolean if a field has been set.

### GetCashHolding

`func (o *CentralbankReservesView) GetCashHolding() float32`

GetCashHolding returns the CashHolding field if non-nil, zero value otherwise.

### GetCashHoldingOk

`func (o *CentralbankReservesView) GetCashHoldingOk() (*float32, bool)`

GetCashHoldingOk returns a tuple with the CashHolding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashHolding

`func (o *CentralbankReservesView) SetCashHolding(v float32)`

SetCashHolding sets CashHolding field to given value.

### HasCashHolding

`func (o *CentralbankReservesView) HasCashHolding() bool`

HasCashHolding returns a boolean if a field has been set.

### GetCoinsForNextBoost

`func (o *CentralbankReservesView) GetCoinsForNextBoost() int64`

GetCoinsForNextBoost returns the CoinsForNextBoost field if non-nil, zero value otherwise.

### GetCoinsForNextBoostOk

`func (o *CentralbankReservesView) GetCoinsForNextBoostOk() (*int64, bool)`

GetCoinsForNextBoostOk returns a tuple with the CoinsForNextBoost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinsForNextBoost

`func (o *CentralbankReservesView) SetCoinsForNextBoost(v int64)`

SetCoinsForNextBoost sets CoinsForNextBoost field to given value.

### HasCoinsForNextBoost

`func (o *CentralbankReservesView) HasCoinsForNextBoost() bool`

HasCoinsForNextBoost returns a boolean if a field has been set.

### GetId

`func (o *CentralbankReservesView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CentralbankReservesView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CentralbankReservesView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CentralbankReservesView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterestRateBoost

`func (o *CentralbankReservesView) GetInterestRateBoost() float32`

GetInterestRateBoost returns the InterestRateBoost field if non-nil, zero value otherwise.

### GetInterestRateBoostOk

`func (o *CentralbankReservesView) GetInterestRateBoostOk() (*float32, bool)`

GetInterestRateBoostOk returns a tuple with the InterestRateBoost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRateBoost

`func (o *CentralbankReservesView) SetInterestRateBoost(v float32)`

SetInterestRateBoost sets InterestRateBoost field to given value.

### HasInterestRateBoost

`func (o *CentralbankReservesView) HasInterestRateBoost() bool`

HasInterestRateBoost returns a boolean if a field has been set.

### GetMaxCentralBankLoans

`func (o *CentralbankReservesView) GetMaxCentralBankLoans() float32`

GetMaxCentralBankLoans returns the MaxCentralBankLoans field if non-nil, zero value otherwise.

### GetMaxCentralBankLoansOk

`func (o *CentralbankReservesView) GetMaxCentralBankLoansOk() (*float32, bool)`

GetMaxCentralBankLoansOk returns a tuple with the MaxCentralBankLoans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCentralBankLoans

`func (o *CentralbankReservesView) SetMaxCentralBankLoans(v float32)`

SetMaxCentralBankLoans sets MaxCentralBankLoans field to given value.

### HasMaxCentralBankLoans

`func (o *CentralbankReservesView) HasMaxCentralBankLoans() bool`

HasMaxCentralBankLoans returns a boolean if a field has been set.

### GetVersion

`func (o *CentralbankReservesView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CentralbankReservesView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CentralbankReservesView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CentralbankReservesView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


