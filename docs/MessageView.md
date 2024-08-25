# MessageView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**DateSent** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**Read** | Pointer to **bool** |  | [optional] 
**Sender** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewMessageView

`func NewMessageView() *MessageView`

NewMessageView instantiates a new MessageView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageViewWithDefaults

`func NewMessageViewWithDefaults() *MessageView`

NewMessageViewWithDefaults instantiates a new MessageView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *MessageView) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *MessageView) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *MessageView) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *MessageView) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetContent

`func (o *MessageView) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MessageView) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MessageView) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *MessageView) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDateSent

`func (o *MessageView) GetDateSent() int64`

GetDateSent returns the DateSent field if non-nil, zero value otherwise.

### GetDateSentOk

`func (o *MessageView) GetDateSentOk() (*int64, bool)`

GetDateSentOk returns a tuple with the DateSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateSent

`func (o *MessageView) SetDateSent(v int64)`

SetDateSent sets DateSent field to given value.

### HasDateSent

`func (o *MessageView) HasDateSent() bool`

HasDateSent returns a boolean if a field has been set.

### GetId

`func (o *MessageView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MessageView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOnline

`func (o *MessageView) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *MessageView) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *MessageView) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *MessageView) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetRead

`func (o *MessageView) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *MessageView) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *MessageView) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *MessageView) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetSender

`func (o *MessageView) GetSender() UsernameView`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *MessageView) GetSenderOk() (*UsernameView, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *MessageView) SetSender(v UsernameView)`

SetSender sets Sender field to given value.

### HasSender

`func (o *MessageView) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetStatus

`func (o *MessageView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessageView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MessageView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MessageView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *MessageView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MessageView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MessageView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MessageView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


