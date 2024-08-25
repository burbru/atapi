# MinerView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoinsPerHour** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MaximumCapacity** | Pointer to **float32** |  | [optional] 
**NextLevelCoinsPerHour** | Pointer to **float32** |  | [optional] 
**NextLevelCosts** | Pointer to **float32** |  | [optional] 
**Owner** | Pointer to [**SecuritiesAccountView**](SecuritiesAccountView.md) |  | [optional] 
**Storage** | Pointer to **float32** |  | [optional] 
**TransferableCoins** | Pointer to **int64** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewMinerView

`func NewMinerView() *MinerView`

NewMinerView instantiates a new MinerView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinerViewWithDefaults

`func NewMinerViewWithDefaults() *MinerView`

NewMinerViewWithDefaults instantiates a new MinerView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoinsPerHour

`func (o *MinerView) GetCoinsPerHour() float32`

GetCoinsPerHour returns the CoinsPerHour field if non-nil, zero value otherwise.

### GetCoinsPerHourOk

`func (o *MinerView) GetCoinsPerHourOk() (*float32, bool)`

GetCoinsPerHourOk returns a tuple with the CoinsPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinsPerHour

`func (o *MinerView) SetCoinsPerHour(v float32)`

SetCoinsPerHour sets CoinsPerHour field to given value.

### HasCoinsPerHour

`func (o *MinerView) HasCoinsPerHour() bool`

HasCoinsPerHour returns a boolean if a field has been set.

### GetId

`func (o *MinerView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MinerView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MinerView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MinerView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaximumCapacity

`func (o *MinerView) GetMaximumCapacity() float32`

GetMaximumCapacity returns the MaximumCapacity field if non-nil, zero value otherwise.

### GetMaximumCapacityOk

`func (o *MinerView) GetMaximumCapacityOk() (*float32, bool)`

GetMaximumCapacityOk returns a tuple with the MaximumCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCapacity

`func (o *MinerView) SetMaximumCapacity(v float32)`

SetMaximumCapacity sets MaximumCapacity field to given value.

### HasMaximumCapacity

`func (o *MinerView) HasMaximumCapacity() bool`

HasMaximumCapacity returns a boolean if a field has been set.

### GetNextLevelCoinsPerHour

`func (o *MinerView) GetNextLevelCoinsPerHour() float32`

GetNextLevelCoinsPerHour returns the NextLevelCoinsPerHour field if non-nil, zero value otherwise.

### GetNextLevelCoinsPerHourOk

`func (o *MinerView) GetNextLevelCoinsPerHourOk() (*float32, bool)`

GetNextLevelCoinsPerHourOk returns a tuple with the NextLevelCoinsPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLevelCoinsPerHour

`func (o *MinerView) SetNextLevelCoinsPerHour(v float32)`

SetNextLevelCoinsPerHour sets NextLevelCoinsPerHour field to given value.

### HasNextLevelCoinsPerHour

`func (o *MinerView) HasNextLevelCoinsPerHour() bool`

HasNextLevelCoinsPerHour returns a boolean if a field has been set.

### GetNextLevelCosts

`func (o *MinerView) GetNextLevelCosts() float32`

GetNextLevelCosts returns the NextLevelCosts field if non-nil, zero value otherwise.

### GetNextLevelCostsOk

`func (o *MinerView) GetNextLevelCostsOk() (*float32, bool)`

GetNextLevelCostsOk returns a tuple with the NextLevelCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLevelCosts

`func (o *MinerView) SetNextLevelCosts(v float32)`

SetNextLevelCosts sets NextLevelCosts field to given value.

### HasNextLevelCosts

`func (o *MinerView) HasNextLevelCosts() bool`

HasNextLevelCosts returns a boolean if a field has been set.

### GetOwner

`func (o *MinerView) GetOwner() SecuritiesAccountView`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MinerView) GetOwnerOk() (*SecuritiesAccountView, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MinerView) SetOwner(v SecuritiesAccountView)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *MinerView) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetStorage

`func (o *MinerView) GetStorage() float32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *MinerView) GetStorageOk() (*float32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *MinerView) SetStorage(v float32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *MinerView) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetTransferableCoins

`func (o *MinerView) GetTransferableCoins() int64`

GetTransferableCoins returns the TransferableCoins field if non-nil, zero value otherwise.

### GetTransferableCoinsOk

`func (o *MinerView) GetTransferableCoinsOk() (*int64, bool)`

GetTransferableCoinsOk returns a tuple with the TransferableCoins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferableCoins

`func (o *MinerView) SetTransferableCoins(v int64)`

SetTransferableCoins sets TransferableCoins field to given value.

### HasTransferableCoins

`func (o *MinerView) HasTransferableCoins() bool`

HasTransferableCoins returns a boolean if a field has been set.

### GetVersion

`func (o *MinerView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MinerView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MinerView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MinerView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


