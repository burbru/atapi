# ChatView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatName** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **int64** |  | [optional] 
**GroupChat** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastMessage** | Pointer to [**MessageView**](MessageView.md) |  | [optional] 
**NumOfUnreadMessages** | Pointer to **int64** |  | [optional] 
**Owner** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**Participants** | Pointer to [**[]UsernameView**](UsernameView.md) |  | [optional] 
**PublicChat** | Pointer to **bool** |  | [optional] 
**Readonly** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewChatView

`func NewChatView() *ChatView`

NewChatView instantiates a new ChatView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatViewWithDefaults

`func NewChatViewWithDefaults() *ChatView`

NewChatViewWithDefaults instantiates a new ChatView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatName

`func (o *ChatView) GetChatName() string`

GetChatName returns the ChatName field if non-nil, zero value otherwise.

### GetChatNameOk

`func (o *ChatView) GetChatNameOk() (*string, bool)`

GetChatNameOk returns a tuple with the ChatName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatName

`func (o *ChatView) SetChatName(v string)`

SetChatName sets ChatName field to given value.

### HasChatName

`func (o *ChatView) HasChatName() bool`

HasChatName returns a boolean if a field has been set.

### GetDateCreated

`func (o *ChatView) GetDateCreated() int64`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ChatView) GetDateCreatedOk() (*int64, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ChatView) SetDateCreated(v int64)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ChatView) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetGroupChat

`func (o *ChatView) GetGroupChat() bool`

GetGroupChat returns the GroupChat field if non-nil, zero value otherwise.

### GetGroupChatOk

`func (o *ChatView) GetGroupChatOk() (*bool, bool)`

GetGroupChatOk returns a tuple with the GroupChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupChat

`func (o *ChatView) SetGroupChat(v bool)`

SetGroupChat sets GroupChat field to given value.

### HasGroupChat

`func (o *ChatView) HasGroupChat() bool`

HasGroupChat returns a boolean if a field has been set.

### GetId

`func (o *ChatView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChatView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastMessage

`func (o *ChatView) GetLastMessage() MessageView`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *ChatView) GetLastMessageOk() (*MessageView, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *ChatView) SetLastMessage(v MessageView)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *ChatView) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### GetNumOfUnreadMessages

`func (o *ChatView) GetNumOfUnreadMessages() int64`

GetNumOfUnreadMessages returns the NumOfUnreadMessages field if non-nil, zero value otherwise.

### GetNumOfUnreadMessagesOk

`func (o *ChatView) GetNumOfUnreadMessagesOk() (*int64, bool)`

GetNumOfUnreadMessagesOk returns a tuple with the NumOfUnreadMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfUnreadMessages

`func (o *ChatView) SetNumOfUnreadMessages(v int64)`

SetNumOfUnreadMessages sets NumOfUnreadMessages field to given value.

### HasNumOfUnreadMessages

`func (o *ChatView) HasNumOfUnreadMessages() bool`

HasNumOfUnreadMessages returns a boolean if a field has been set.

### GetOwner

`func (o *ChatView) GetOwner() UsernameView`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ChatView) GetOwnerOk() (*UsernameView, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ChatView) SetOwner(v UsernameView)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ChatView) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParticipants

`func (o *ChatView) GetParticipants() []UsernameView`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *ChatView) GetParticipantsOk() (*[]UsernameView, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *ChatView) SetParticipants(v []UsernameView)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *ChatView) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetPublicChat

`func (o *ChatView) GetPublicChat() bool`

GetPublicChat returns the PublicChat field if non-nil, zero value otherwise.

### GetPublicChatOk

`func (o *ChatView) GetPublicChatOk() (*bool, bool)`

GetPublicChatOk returns a tuple with the PublicChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicChat

`func (o *ChatView) SetPublicChat(v bool)`

SetPublicChat sets PublicChat field to given value.

### HasPublicChat

`func (o *ChatView) HasPublicChat() bool`

HasPublicChat returns a boolean if a field has been set.

### GetReadonly

`func (o *ChatView) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *ChatView) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *ChatView) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *ChatView) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetStatus

`func (o *ChatView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChatView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChatView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChatView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *ChatView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ChatView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ChatView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ChatView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


