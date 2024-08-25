# HistorizedCompanyDataView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BondsVolume** | Pointer to **float32** |  | [optional] 
**BookValue** | Pointer to **float32** |  | [optional] 
**BookValuePerShare** | Pointer to **float32** |  | [optional] 
**Cash** | Pointer to **float32** |  | [optional] 
**CashFlow** | Pointer to **float32** |  | [optional] 
**CentralBankReserves** | Pointer to **float32** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**FairValuePerShare** | Pointer to **float32** |  | [optional] 
**FreeFloatInPercent** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**NetCash** | Pointer to **float32** |  | [optional] 
**NetCashPerShare** | Pointer to **float32** |  | [optional] 
**ReposVolume** | Pointer to **float32** |  | [optional] 
**SystemReposVolume** | Pointer to **float32** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewHistorizedCompanyDataView

`func NewHistorizedCompanyDataView() *HistorizedCompanyDataView`

NewHistorizedCompanyDataView instantiates a new HistorizedCompanyDataView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistorizedCompanyDataViewWithDefaults

`func NewHistorizedCompanyDataViewWithDefaults() *HistorizedCompanyDataView`

NewHistorizedCompanyDataViewWithDefaults instantiates a new HistorizedCompanyDataView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBondsVolume

`func (o *HistorizedCompanyDataView) GetBondsVolume() float32`

GetBondsVolume returns the BondsVolume field if non-nil, zero value otherwise.

### GetBondsVolumeOk

`func (o *HistorizedCompanyDataView) GetBondsVolumeOk() (*float32, bool)`

GetBondsVolumeOk returns a tuple with the BondsVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondsVolume

`func (o *HistorizedCompanyDataView) SetBondsVolume(v float32)`

SetBondsVolume sets BondsVolume field to given value.

### HasBondsVolume

`func (o *HistorizedCompanyDataView) HasBondsVolume() bool`

HasBondsVolume returns a boolean if a field has been set.

### GetBookValue

`func (o *HistorizedCompanyDataView) GetBookValue() float32`

GetBookValue returns the BookValue field if non-nil, zero value otherwise.

### GetBookValueOk

`func (o *HistorizedCompanyDataView) GetBookValueOk() (*float32, bool)`

GetBookValueOk returns a tuple with the BookValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookValue

`func (o *HistorizedCompanyDataView) SetBookValue(v float32)`

SetBookValue sets BookValue field to given value.

### HasBookValue

`func (o *HistorizedCompanyDataView) HasBookValue() bool`

HasBookValue returns a boolean if a field has been set.

### GetBookValuePerShare

`func (o *HistorizedCompanyDataView) GetBookValuePerShare() float32`

GetBookValuePerShare returns the BookValuePerShare field if non-nil, zero value otherwise.

### GetBookValuePerShareOk

`func (o *HistorizedCompanyDataView) GetBookValuePerShareOk() (*float32, bool)`

GetBookValuePerShareOk returns a tuple with the BookValuePerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookValuePerShare

`func (o *HistorizedCompanyDataView) SetBookValuePerShare(v float32)`

SetBookValuePerShare sets BookValuePerShare field to given value.

### HasBookValuePerShare

`func (o *HistorizedCompanyDataView) HasBookValuePerShare() bool`

HasBookValuePerShare returns a boolean if a field has been set.

### GetCash

`func (o *HistorizedCompanyDataView) GetCash() float32`

GetCash returns the Cash field if non-nil, zero value otherwise.

### GetCashOk

`func (o *HistorizedCompanyDataView) GetCashOk() (*float32, bool)`

GetCashOk returns a tuple with the Cash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCash

`func (o *HistorizedCompanyDataView) SetCash(v float32)`

SetCash sets Cash field to given value.

### HasCash

`func (o *HistorizedCompanyDataView) HasCash() bool`

HasCash returns a boolean if a field has been set.

### GetCashFlow

`func (o *HistorizedCompanyDataView) GetCashFlow() float32`

GetCashFlow returns the CashFlow field if non-nil, zero value otherwise.

### GetCashFlowOk

`func (o *HistorizedCompanyDataView) GetCashFlowOk() (*float32, bool)`

GetCashFlowOk returns a tuple with the CashFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFlow

`func (o *HistorizedCompanyDataView) SetCashFlow(v float32)`

SetCashFlow sets CashFlow field to given value.

### HasCashFlow

`func (o *HistorizedCompanyDataView) HasCashFlow() bool`

HasCashFlow returns a boolean if a field has been set.

### GetCentralBankReserves

`func (o *HistorizedCompanyDataView) GetCentralBankReserves() float32`

GetCentralBankReserves returns the CentralBankReserves field if non-nil, zero value otherwise.

### GetCentralBankReservesOk

`func (o *HistorizedCompanyDataView) GetCentralBankReservesOk() (*float32, bool)`

GetCentralBankReservesOk returns a tuple with the CentralBankReserves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentralBankReserves

`func (o *HistorizedCompanyDataView) SetCentralBankReserves(v float32)`

SetCentralBankReserves sets CentralBankReserves field to given value.

### HasCentralBankReserves

`func (o *HistorizedCompanyDataView) HasCentralBankReserves() bool`

HasCentralBankReserves returns a boolean if a field has been set.

### GetDate

`func (o *HistorizedCompanyDataView) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *HistorizedCompanyDataView) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *HistorizedCompanyDataView) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *HistorizedCompanyDataView) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetFairValuePerShare

`func (o *HistorizedCompanyDataView) GetFairValuePerShare() float32`

GetFairValuePerShare returns the FairValuePerShare field if non-nil, zero value otherwise.

### GetFairValuePerShareOk

`func (o *HistorizedCompanyDataView) GetFairValuePerShareOk() (*float32, bool)`

GetFairValuePerShareOk returns a tuple with the FairValuePerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFairValuePerShare

`func (o *HistorizedCompanyDataView) SetFairValuePerShare(v float32)`

SetFairValuePerShare sets FairValuePerShare field to given value.

### HasFairValuePerShare

`func (o *HistorizedCompanyDataView) HasFairValuePerShare() bool`

HasFairValuePerShare returns a boolean if a field has been set.

### GetFreeFloatInPercent

`func (o *HistorizedCompanyDataView) GetFreeFloatInPercent() float32`

GetFreeFloatInPercent returns the FreeFloatInPercent field if non-nil, zero value otherwise.

### GetFreeFloatInPercentOk

`func (o *HistorizedCompanyDataView) GetFreeFloatInPercentOk() (*float32, bool)`

GetFreeFloatInPercentOk returns a tuple with the FreeFloatInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeFloatInPercent

`func (o *HistorizedCompanyDataView) SetFreeFloatInPercent(v float32)`

SetFreeFloatInPercent sets FreeFloatInPercent field to given value.

### HasFreeFloatInPercent

`func (o *HistorizedCompanyDataView) HasFreeFloatInPercent() bool`

HasFreeFloatInPercent returns a boolean if a field has been set.

### GetId

`func (o *HistorizedCompanyDataView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistorizedCompanyDataView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistorizedCompanyDataView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HistorizedCompanyDataView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetCash

`func (o *HistorizedCompanyDataView) GetNetCash() float32`

GetNetCash returns the NetCash field if non-nil, zero value otherwise.

### GetNetCashOk

`func (o *HistorizedCompanyDataView) GetNetCashOk() (*float32, bool)`

GetNetCashOk returns a tuple with the NetCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetCash

`func (o *HistorizedCompanyDataView) SetNetCash(v float32)`

SetNetCash sets NetCash field to given value.

### HasNetCash

`func (o *HistorizedCompanyDataView) HasNetCash() bool`

HasNetCash returns a boolean if a field has been set.

### GetNetCashPerShare

`func (o *HistorizedCompanyDataView) GetNetCashPerShare() float32`

GetNetCashPerShare returns the NetCashPerShare field if non-nil, zero value otherwise.

### GetNetCashPerShareOk

`func (o *HistorizedCompanyDataView) GetNetCashPerShareOk() (*float32, bool)`

GetNetCashPerShareOk returns a tuple with the NetCashPerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetCashPerShare

`func (o *HistorizedCompanyDataView) SetNetCashPerShare(v float32)`

SetNetCashPerShare sets NetCashPerShare field to given value.

### HasNetCashPerShare

`func (o *HistorizedCompanyDataView) HasNetCashPerShare() bool`

HasNetCashPerShare returns a boolean if a field has been set.

### GetReposVolume

`func (o *HistorizedCompanyDataView) GetReposVolume() float32`

GetReposVolume returns the ReposVolume field if non-nil, zero value otherwise.

### GetReposVolumeOk

`func (o *HistorizedCompanyDataView) GetReposVolumeOk() (*float32, bool)`

GetReposVolumeOk returns a tuple with the ReposVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReposVolume

`func (o *HistorizedCompanyDataView) SetReposVolume(v float32)`

SetReposVolume sets ReposVolume field to given value.

### HasReposVolume

`func (o *HistorizedCompanyDataView) HasReposVolume() bool`

HasReposVolume returns a boolean if a field has been set.

### GetSystemReposVolume

`func (o *HistorizedCompanyDataView) GetSystemReposVolume() float32`

GetSystemReposVolume returns the SystemReposVolume field if non-nil, zero value otherwise.

### GetSystemReposVolumeOk

`func (o *HistorizedCompanyDataView) GetSystemReposVolumeOk() (*float32, bool)`

GetSystemReposVolumeOk returns a tuple with the SystemReposVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemReposVolume

`func (o *HistorizedCompanyDataView) SetSystemReposVolume(v float32)`

SetSystemReposVolume sets SystemReposVolume field to given value.

### HasSystemReposVolume

`func (o *HistorizedCompanyDataView) HasSystemReposVolume() bool`

HasSystemReposVolume returns a boolean if a field has been set.

### GetVersion

`func (o *HistorizedCompanyDataView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HistorizedCompanyDataView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HistorizedCompanyDataView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HistorizedCompanyDataView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


