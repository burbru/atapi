# SecurityOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**CheckResult** | Pointer to [**SecurityOrderCheckResult**](SecurityOrderCheckResult.md) |  | [optional] 
**CommittedCash** | Pointer to **float32** |  | [optional] 
**CounterParty** | Pointer to **string** |  | [optional] 
**CounterPartyName** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to **int64** |  | [optional] 
**ExecutionPrice** | Pointer to **float32** |  | [optional] 
**ExecutionVolume** | Pointer to **float32** |  | [optional] 
**GoodAfterDate** | Pointer to **int64** |  | [optional] 
**GoodTillDate** | Pointer to **int64** |  | [optional] 
**HourlyChange** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Listing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**NextHourlyChangeDate** | Pointer to **int64** |  | [optional] 
**NumberOfShares** | **int64** |  | 
**Owner** | **string** |  | 
**OwnerName** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**PrivateCounterParty** | Pointer to **bool** |  | [optional] 
**PrivateOwner** | Pointer to **bool** |  | [optional] 
**SecurityIdentifier** | **string** |  | 
**Spread** | Pointer to [**PriceSpreadView**](PriceSpreadView.md) |  | [optional] 
**Type** | **string** |  | 
**UncommittedCash** | Pointer to **float32** |  | [optional] 
**UncommittedShares** | Pointer to **int64** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**Volume** | Pointer to **float32** |  | [optional] 

## Methods

### NewSecurityOrder

`func NewSecurityOrder(action string, numberOfShares int64, owner string, securityIdentifier string, type_ string, ) *SecurityOrder`

NewSecurityOrder instantiates a new SecurityOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityOrderWithDefaults

`func NewSecurityOrderWithDefaults() *SecurityOrder`

NewSecurityOrderWithDefaults instantiates a new SecurityOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SecurityOrder) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SecurityOrder) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SecurityOrder) SetAction(v string)`

SetAction sets Action field to given value.


### GetCheckResult

`func (o *SecurityOrder) GetCheckResult() SecurityOrderCheckResult`

GetCheckResult returns the CheckResult field if non-nil, zero value otherwise.

### GetCheckResultOk

`func (o *SecurityOrder) GetCheckResultOk() (*SecurityOrderCheckResult, bool)`

GetCheckResultOk returns a tuple with the CheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResult

`func (o *SecurityOrder) SetCheckResult(v SecurityOrderCheckResult)`

SetCheckResult sets CheckResult field to given value.

### HasCheckResult

`func (o *SecurityOrder) HasCheckResult() bool`

HasCheckResult returns a boolean if a field has been set.

### GetCommittedCash

`func (o *SecurityOrder) GetCommittedCash() float32`

GetCommittedCash returns the CommittedCash field if non-nil, zero value otherwise.

### GetCommittedCashOk

`func (o *SecurityOrder) GetCommittedCashOk() (*float32, bool)`

GetCommittedCashOk returns a tuple with the CommittedCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedCash

`func (o *SecurityOrder) SetCommittedCash(v float32)`

SetCommittedCash sets CommittedCash field to given value.

### HasCommittedCash

`func (o *SecurityOrder) HasCommittedCash() bool`

HasCommittedCash returns a boolean if a field has been set.

### GetCounterParty

`func (o *SecurityOrder) GetCounterParty() string`

GetCounterParty returns the CounterParty field if non-nil, zero value otherwise.

### GetCounterPartyOk

`func (o *SecurityOrder) GetCounterPartyOk() (*string, bool)`

GetCounterPartyOk returns a tuple with the CounterParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterParty

`func (o *SecurityOrder) SetCounterParty(v string)`

SetCounterParty sets CounterParty field to given value.

### HasCounterParty

`func (o *SecurityOrder) HasCounterParty() bool`

HasCounterParty returns a boolean if a field has been set.

### GetCounterPartyName

`func (o *SecurityOrder) GetCounterPartyName() string`

GetCounterPartyName returns the CounterPartyName field if non-nil, zero value otherwise.

### GetCounterPartyNameOk

`func (o *SecurityOrder) GetCounterPartyNameOk() (*string, bool)`

GetCounterPartyNameOk returns a tuple with the CounterPartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterPartyName

`func (o *SecurityOrder) SetCounterPartyName(v string)`

SetCounterPartyName sets CounterPartyName field to given value.

### HasCounterPartyName

`func (o *SecurityOrder) HasCounterPartyName() bool`

HasCounterPartyName returns a boolean if a field has been set.

### GetCreationDate

`func (o *SecurityOrder) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *SecurityOrder) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *SecurityOrder) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *SecurityOrder) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetExecutionPrice

`func (o *SecurityOrder) GetExecutionPrice() float32`

GetExecutionPrice returns the ExecutionPrice field if non-nil, zero value otherwise.

### GetExecutionPriceOk

`func (o *SecurityOrder) GetExecutionPriceOk() (*float32, bool)`

GetExecutionPriceOk returns a tuple with the ExecutionPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionPrice

`func (o *SecurityOrder) SetExecutionPrice(v float32)`

SetExecutionPrice sets ExecutionPrice field to given value.

### HasExecutionPrice

`func (o *SecurityOrder) HasExecutionPrice() bool`

HasExecutionPrice returns a boolean if a field has been set.

### GetExecutionVolume

`func (o *SecurityOrder) GetExecutionVolume() float32`

GetExecutionVolume returns the ExecutionVolume field if non-nil, zero value otherwise.

### GetExecutionVolumeOk

`func (o *SecurityOrder) GetExecutionVolumeOk() (*float32, bool)`

GetExecutionVolumeOk returns a tuple with the ExecutionVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionVolume

`func (o *SecurityOrder) SetExecutionVolume(v float32)`

SetExecutionVolume sets ExecutionVolume field to given value.

### HasExecutionVolume

`func (o *SecurityOrder) HasExecutionVolume() bool`

HasExecutionVolume returns a boolean if a field has been set.

### GetGoodAfterDate

`func (o *SecurityOrder) GetGoodAfterDate() int64`

GetGoodAfterDate returns the GoodAfterDate field if non-nil, zero value otherwise.

### GetGoodAfterDateOk

`func (o *SecurityOrder) GetGoodAfterDateOk() (*int64, bool)`

GetGoodAfterDateOk returns a tuple with the GoodAfterDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodAfterDate

`func (o *SecurityOrder) SetGoodAfterDate(v int64)`

SetGoodAfterDate sets GoodAfterDate field to given value.

### HasGoodAfterDate

`func (o *SecurityOrder) HasGoodAfterDate() bool`

HasGoodAfterDate returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *SecurityOrder) GetGoodTillDate() int64`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *SecurityOrder) GetGoodTillDateOk() (*int64, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *SecurityOrder) SetGoodTillDate(v int64)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *SecurityOrder) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.

### GetHourlyChange

`func (o *SecurityOrder) GetHourlyChange() float32`

GetHourlyChange returns the HourlyChange field if non-nil, zero value otherwise.

### GetHourlyChangeOk

`func (o *SecurityOrder) GetHourlyChangeOk() (*float32, bool)`

GetHourlyChangeOk returns a tuple with the HourlyChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyChange

`func (o *SecurityOrder) SetHourlyChange(v float32)`

SetHourlyChange sets HourlyChange field to given value.

### HasHourlyChange

`func (o *SecurityOrder) HasHourlyChange() bool`

HasHourlyChange returns a boolean if a field has been set.

### GetId

`func (o *SecurityOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetListing

`func (o *SecurityOrder) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *SecurityOrder) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *SecurityOrder) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *SecurityOrder) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetNextHourlyChangeDate

`func (o *SecurityOrder) GetNextHourlyChangeDate() int64`

GetNextHourlyChangeDate returns the NextHourlyChangeDate field if non-nil, zero value otherwise.

### GetNextHourlyChangeDateOk

`func (o *SecurityOrder) GetNextHourlyChangeDateOk() (*int64, bool)`

GetNextHourlyChangeDateOk returns a tuple with the NextHourlyChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHourlyChangeDate

`func (o *SecurityOrder) SetNextHourlyChangeDate(v int64)`

SetNextHourlyChangeDate sets NextHourlyChangeDate field to given value.

### HasNextHourlyChangeDate

`func (o *SecurityOrder) HasNextHourlyChangeDate() bool`

HasNextHourlyChangeDate returns a boolean if a field has been set.

### GetNumberOfShares

`func (o *SecurityOrder) GetNumberOfShares() int64`

GetNumberOfShares returns the NumberOfShares field if non-nil, zero value otherwise.

### GetNumberOfSharesOk

`func (o *SecurityOrder) GetNumberOfSharesOk() (*int64, bool)`

GetNumberOfSharesOk returns a tuple with the NumberOfShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfShares

`func (o *SecurityOrder) SetNumberOfShares(v int64)`

SetNumberOfShares sets NumberOfShares field to given value.


### GetOwner

`func (o *SecurityOrder) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SecurityOrder) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SecurityOrder) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetOwnerName

`func (o *SecurityOrder) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *SecurityOrder) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *SecurityOrder) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *SecurityOrder) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetPrice

`func (o *SecurityOrder) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SecurityOrder) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SecurityOrder) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SecurityOrder) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPrivateCounterParty

`func (o *SecurityOrder) GetPrivateCounterParty() bool`

GetPrivateCounterParty returns the PrivateCounterParty field if non-nil, zero value otherwise.

### GetPrivateCounterPartyOk

`func (o *SecurityOrder) GetPrivateCounterPartyOk() (*bool, bool)`

GetPrivateCounterPartyOk returns a tuple with the PrivateCounterParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateCounterParty

`func (o *SecurityOrder) SetPrivateCounterParty(v bool)`

SetPrivateCounterParty sets PrivateCounterParty field to given value.

### HasPrivateCounterParty

`func (o *SecurityOrder) HasPrivateCounterParty() bool`

HasPrivateCounterParty returns a boolean if a field has been set.

### GetPrivateOwner

`func (o *SecurityOrder) GetPrivateOwner() bool`

GetPrivateOwner returns the PrivateOwner field if non-nil, zero value otherwise.

### GetPrivateOwnerOk

`func (o *SecurityOrder) GetPrivateOwnerOk() (*bool, bool)`

GetPrivateOwnerOk returns a tuple with the PrivateOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateOwner

`func (o *SecurityOrder) SetPrivateOwner(v bool)`

SetPrivateOwner sets PrivateOwner field to given value.

### HasPrivateOwner

`func (o *SecurityOrder) HasPrivateOwner() bool`

HasPrivateOwner returns a boolean if a field has been set.

### GetSecurityIdentifier

`func (o *SecurityOrder) GetSecurityIdentifier() string`

GetSecurityIdentifier returns the SecurityIdentifier field if non-nil, zero value otherwise.

### GetSecurityIdentifierOk

`func (o *SecurityOrder) GetSecurityIdentifierOk() (*string, bool)`

GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIdentifier

`func (o *SecurityOrder) SetSecurityIdentifier(v string)`

SetSecurityIdentifier sets SecurityIdentifier field to given value.


### GetSpread

`func (o *SecurityOrder) GetSpread() PriceSpreadView`

GetSpread returns the Spread field if non-nil, zero value otherwise.

### GetSpreadOk

`func (o *SecurityOrder) GetSpreadOk() (*PriceSpreadView, bool)`

GetSpreadOk returns a tuple with the Spread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpread

`func (o *SecurityOrder) SetSpread(v PriceSpreadView)`

SetSpread sets Spread field to given value.

### HasSpread

`func (o *SecurityOrder) HasSpread() bool`

HasSpread returns a boolean if a field has been set.

### GetType

`func (o *SecurityOrder) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityOrder) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityOrder) SetType(v string)`

SetType sets Type field to given value.


### GetUncommittedCash

`func (o *SecurityOrder) GetUncommittedCash() float32`

GetUncommittedCash returns the UncommittedCash field if non-nil, zero value otherwise.

### GetUncommittedCashOk

`func (o *SecurityOrder) GetUncommittedCashOk() (*float32, bool)`

GetUncommittedCashOk returns a tuple with the UncommittedCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncommittedCash

`func (o *SecurityOrder) SetUncommittedCash(v float32)`

SetUncommittedCash sets UncommittedCash field to given value.

### HasUncommittedCash

`func (o *SecurityOrder) HasUncommittedCash() bool`

HasUncommittedCash returns a boolean if a field has been set.

### GetUncommittedShares

`func (o *SecurityOrder) GetUncommittedShares() int64`

GetUncommittedShares returns the UncommittedShares field if non-nil, zero value otherwise.

### GetUncommittedSharesOk

`func (o *SecurityOrder) GetUncommittedSharesOk() (*int64, bool)`

GetUncommittedSharesOk returns a tuple with the UncommittedShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncommittedShares

`func (o *SecurityOrder) SetUncommittedShares(v int64)`

SetUncommittedShares sets UncommittedShares field to given value.

### HasUncommittedShares

`func (o *SecurityOrder) HasUncommittedShares() bool`

HasUncommittedShares returns a boolean if a field has been set.

### GetVersion

`func (o *SecurityOrder) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SecurityOrder) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SecurityOrder) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SecurityOrder) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVolume

`func (o *SecurityOrder) GetVolume() float32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *SecurityOrder) GetVolumeOk() (*float32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *SecurityOrder) SetVolume(v float32)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *SecurityOrder) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


