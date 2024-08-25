# BankAccountView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cash** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewBankAccountView

`func NewBankAccountView() *BankAccountView`

NewBankAccountView instantiates a new BankAccountView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountViewWithDefaults

`func NewBankAccountViewWithDefaults() *BankAccountView`

NewBankAccountViewWithDefaults instantiates a new BankAccountView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCash

`func (o *BankAccountView) GetCash() float32`

GetCash returns the Cash field if non-nil, zero value otherwise.

### GetCashOk

`func (o *BankAccountView) GetCashOk() (*float32, bool)`

GetCashOk returns a tuple with the Cash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCash

`func (o *BankAccountView) SetCash(v float32)`

SetCash sets Cash field to given value.

### HasCash

`func (o *BankAccountView) HasCash() bool`

HasCash returns a boolean if a field has been set.

### GetId

`func (o *BankAccountView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankAccountView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankAccountView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BankAccountView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *BankAccountView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BankAccountView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BankAccountView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BankAccountView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


