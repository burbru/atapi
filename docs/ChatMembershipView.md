# ChatMembershipView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Member** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewChatMembershipView

`func NewChatMembershipView() *ChatMembershipView`

NewChatMembershipView instantiates a new ChatMembershipView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMembershipViewWithDefaults

`func NewChatMembershipViewWithDefaults() *ChatMembershipView`

NewChatMembershipViewWithDefaults instantiates a new ChatMembershipView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *ChatMembershipView) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *ChatMembershipView) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *ChatMembershipView) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *ChatMembershipView) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetId

`func (o *ChatMembershipView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatMembershipView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatMembershipView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChatMembershipView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMember

`func (o *ChatMembershipView) GetMember() UsernameView`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ChatMembershipView) GetMemberOk() (*UsernameView, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ChatMembershipView) SetMember(v UsernameView)`

SetMember sets Member field to given value.

### HasMember

`func (o *ChatMembershipView) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetOnline

`func (o *ChatMembershipView) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ChatMembershipView) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ChatMembershipView) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *ChatMembershipView) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetRole

`func (o *ChatMembershipView) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ChatMembershipView) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ChatMembershipView) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ChatMembershipView) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetVersion

`func (o *ChatMembershipView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ChatMembershipView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ChatMembershipView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ChatMembershipView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


