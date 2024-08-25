# SecurityOrderLogEntryView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerSecuritiesAccount** | Pointer to **string** |  | [optional] 
**BuyerSecuritiesAccountName** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**NumberOfShares** | Pointer to **int64** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**SecurityIdentifier** | Pointer to **string** |  | [optional] 
**SellerAverageBuyingPrice** | Pointer to **float32** |  | [optional] 
**SellerSecuritiesAccount** | Pointer to **string** |  | [optional] 
**SellerSecuritiesAccountName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**Volume** | Pointer to **float32** |  | [optional] 

## Methods

### NewSecurityOrderLogEntryView

`func NewSecurityOrderLogEntryView() *SecurityOrderLogEntryView`

NewSecurityOrderLogEntryView instantiates a new SecurityOrderLogEntryView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityOrderLogEntryViewWithDefaults

`func NewSecurityOrderLogEntryViewWithDefaults() *SecurityOrderLogEntryView`

NewSecurityOrderLogEntryViewWithDefaults instantiates a new SecurityOrderLogEntryView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerSecuritiesAccount

`func (o *SecurityOrderLogEntryView) GetBuyerSecuritiesAccount() string`

GetBuyerSecuritiesAccount returns the BuyerSecuritiesAccount field if non-nil, zero value otherwise.

### GetBuyerSecuritiesAccountOk

`func (o *SecurityOrderLogEntryView) GetBuyerSecuritiesAccountOk() (*string, bool)`

GetBuyerSecuritiesAccountOk returns a tuple with the BuyerSecuritiesAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerSecuritiesAccount

`func (o *SecurityOrderLogEntryView) SetBuyerSecuritiesAccount(v string)`

SetBuyerSecuritiesAccount sets BuyerSecuritiesAccount field to given value.

### HasBuyerSecuritiesAccount

`func (o *SecurityOrderLogEntryView) HasBuyerSecuritiesAccount() bool`

HasBuyerSecuritiesAccount returns a boolean if a field has been set.

### GetBuyerSecuritiesAccountName

`func (o *SecurityOrderLogEntryView) GetBuyerSecuritiesAccountName() string`

GetBuyerSecuritiesAccountName returns the BuyerSecuritiesAccountName field if non-nil, zero value otherwise.

### GetBuyerSecuritiesAccountNameOk

`func (o *SecurityOrderLogEntryView) GetBuyerSecuritiesAccountNameOk() (*string, bool)`

GetBuyerSecuritiesAccountNameOk returns a tuple with the BuyerSecuritiesAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerSecuritiesAccountName

`func (o *SecurityOrderLogEntryView) SetBuyerSecuritiesAccountName(v string)`

SetBuyerSecuritiesAccountName sets BuyerSecuritiesAccountName field to given value.

### HasBuyerSecuritiesAccountName

`func (o *SecurityOrderLogEntryView) HasBuyerSecuritiesAccountName() bool`

HasBuyerSecuritiesAccountName returns a boolean if a field has been set.

### GetDate

`func (o *SecurityOrderLogEntryView) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SecurityOrderLogEntryView) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SecurityOrderLogEntryView) SetDate(v int64)`

SetDate sets Date field to given value.

### HasDate

`func (o *SecurityOrderLogEntryView) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetId

`func (o *SecurityOrderLogEntryView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityOrderLogEntryView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityOrderLogEntryView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityOrderLogEntryView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumberOfShares

`func (o *SecurityOrderLogEntryView) GetNumberOfShares() int64`

GetNumberOfShares returns the NumberOfShares field if non-nil, zero value otherwise.

### GetNumberOfSharesOk

`func (o *SecurityOrderLogEntryView) GetNumberOfSharesOk() (*int64, bool)`

GetNumberOfSharesOk returns a tuple with the NumberOfShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfShares

`func (o *SecurityOrderLogEntryView) SetNumberOfShares(v int64)`

SetNumberOfShares sets NumberOfShares field to given value.

### HasNumberOfShares

`func (o *SecurityOrderLogEntryView) HasNumberOfShares() bool`

HasNumberOfShares returns a boolean if a field has been set.

### GetPrice

`func (o *SecurityOrderLogEntryView) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SecurityOrderLogEntryView) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SecurityOrderLogEntryView) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SecurityOrderLogEntryView) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSecurityIdentifier

`func (o *SecurityOrderLogEntryView) GetSecurityIdentifier() string`

GetSecurityIdentifier returns the SecurityIdentifier field if non-nil, zero value otherwise.

### GetSecurityIdentifierOk

`func (o *SecurityOrderLogEntryView) GetSecurityIdentifierOk() (*string, bool)`

GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIdentifier

`func (o *SecurityOrderLogEntryView) SetSecurityIdentifier(v string)`

SetSecurityIdentifier sets SecurityIdentifier field to given value.

### HasSecurityIdentifier

`func (o *SecurityOrderLogEntryView) HasSecurityIdentifier() bool`

HasSecurityIdentifier returns a boolean if a field has been set.

### GetSellerAverageBuyingPrice

`func (o *SecurityOrderLogEntryView) GetSellerAverageBuyingPrice() float32`

GetSellerAverageBuyingPrice returns the SellerAverageBuyingPrice field if non-nil, zero value otherwise.

### GetSellerAverageBuyingPriceOk

`func (o *SecurityOrderLogEntryView) GetSellerAverageBuyingPriceOk() (*float32, bool)`

GetSellerAverageBuyingPriceOk returns a tuple with the SellerAverageBuyingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerAverageBuyingPrice

`func (o *SecurityOrderLogEntryView) SetSellerAverageBuyingPrice(v float32)`

SetSellerAverageBuyingPrice sets SellerAverageBuyingPrice field to given value.

### HasSellerAverageBuyingPrice

`func (o *SecurityOrderLogEntryView) HasSellerAverageBuyingPrice() bool`

HasSellerAverageBuyingPrice returns a boolean if a field has been set.

### GetSellerSecuritiesAccount

`func (o *SecurityOrderLogEntryView) GetSellerSecuritiesAccount() string`

GetSellerSecuritiesAccount returns the SellerSecuritiesAccount field if non-nil, zero value otherwise.

### GetSellerSecuritiesAccountOk

`func (o *SecurityOrderLogEntryView) GetSellerSecuritiesAccountOk() (*string, bool)`

GetSellerSecuritiesAccountOk returns a tuple with the SellerSecuritiesAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerSecuritiesAccount

`func (o *SecurityOrderLogEntryView) SetSellerSecuritiesAccount(v string)`

SetSellerSecuritiesAccount sets SellerSecuritiesAccount field to given value.

### HasSellerSecuritiesAccount

`func (o *SecurityOrderLogEntryView) HasSellerSecuritiesAccount() bool`

HasSellerSecuritiesAccount returns a boolean if a field has been set.

### GetSellerSecuritiesAccountName

`func (o *SecurityOrderLogEntryView) GetSellerSecuritiesAccountName() string`

GetSellerSecuritiesAccountName returns the SellerSecuritiesAccountName field if non-nil, zero value otherwise.

### GetSellerSecuritiesAccountNameOk

`func (o *SecurityOrderLogEntryView) GetSellerSecuritiesAccountNameOk() (*string, bool)`

GetSellerSecuritiesAccountNameOk returns a tuple with the SellerSecuritiesAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerSecuritiesAccountName

`func (o *SecurityOrderLogEntryView) SetSellerSecuritiesAccountName(v string)`

SetSellerSecuritiesAccountName sets SellerSecuritiesAccountName field to given value.

### HasSellerSecuritiesAccountName

`func (o *SecurityOrderLogEntryView) HasSellerSecuritiesAccountName() bool`

HasSellerSecuritiesAccountName returns a boolean if a field has been set.

### GetVersion

`func (o *SecurityOrderLogEntryView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SecurityOrderLogEntryView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SecurityOrderLogEntryView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SecurityOrderLogEntryView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVolume

`func (o *SecurityOrderLogEntryView) GetVolume() float32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *SecurityOrderLogEntryView) GetVolumeOk() (*float32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *SecurityOrderLogEntryView) SetVolume(v float32)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *SecurityOrderLogEntryView) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


