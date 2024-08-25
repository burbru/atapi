# CashTransferLogEntryView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**Date** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Message** | Pointer to [**MessagePrototype**](MessagePrototype.md) |  | [optional] 
**ReceiverBankAccount** | Pointer to **string** |  | [optional] 
**SenderBankAccount** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewCashTransferLogEntryView

`func NewCashTransferLogEntryView() *CashTransferLogEntryView`

NewCashTransferLogEntryView instantiates a new CashTransferLogEntryView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashTransferLogEntryViewWithDefaults

`func NewCashTransferLogEntryViewWithDefaults() *CashTransferLogEntryView`

NewCashTransferLogEntryViewWithDefaults instantiates a new CashTransferLogEntryView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CashTransferLogEntryView) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CashTransferLogEntryView) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CashTransferLogEntryView) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CashTransferLogEntryView) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDate

`func (o *CashTransferLogEntryView) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CashTransferLogEntryView) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CashTransferLogEntryView) SetDate(v int64)`

SetDate sets Date field to given value.

### HasDate

`func (o *CashTransferLogEntryView) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetId

`func (o *CashTransferLogEntryView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CashTransferLogEntryView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CashTransferLogEntryView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CashTransferLogEntryView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *CashTransferLogEntryView) GetMessage() MessagePrototype`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CashTransferLogEntryView) GetMessageOk() (*MessagePrototype, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CashTransferLogEntryView) SetMessage(v MessagePrototype)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CashTransferLogEntryView) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReceiverBankAccount

`func (o *CashTransferLogEntryView) GetReceiverBankAccount() string`

GetReceiverBankAccount returns the ReceiverBankAccount field if non-nil, zero value otherwise.

### GetReceiverBankAccountOk

`func (o *CashTransferLogEntryView) GetReceiverBankAccountOk() (*string, bool)`

GetReceiverBankAccountOk returns a tuple with the ReceiverBankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverBankAccount

`func (o *CashTransferLogEntryView) SetReceiverBankAccount(v string)`

SetReceiverBankAccount sets ReceiverBankAccount field to given value.

### HasReceiverBankAccount

`func (o *CashTransferLogEntryView) HasReceiverBankAccount() bool`

HasReceiverBankAccount returns a boolean if a field has been set.

### GetSenderBankAccount

`func (o *CashTransferLogEntryView) GetSenderBankAccount() string`

GetSenderBankAccount returns the SenderBankAccount field if non-nil, zero value otherwise.

### GetSenderBankAccountOk

`func (o *CashTransferLogEntryView) GetSenderBankAccountOk() (*string, bool)`

GetSenderBankAccountOk returns a tuple with the SenderBankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderBankAccount

`func (o *CashTransferLogEntryView) SetSenderBankAccount(v string)`

SetSenderBankAccount sets SenderBankAccount field to given value.

### HasSenderBankAccount

`func (o *CashTransferLogEntryView) HasSenderBankAccount() bool`

HasSenderBankAccount returns a boolean if a field has been set.

### GetVersion

`func (o *CashTransferLogEntryView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CashTransferLogEntryView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CashTransferLogEntryView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CashTransferLogEntryView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


