# MessageBoardMembershipView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Member** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**MessageBoard** | Pointer to [**MessageBoardView**](MessageBoardView.md) |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewMessageBoardMembershipView

`func NewMessageBoardMembershipView() *MessageBoardMembershipView`

NewMessageBoardMembershipView instantiates a new MessageBoardMembershipView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageBoardMembershipViewWithDefaults

`func NewMessageBoardMembershipViewWithDefaults() *MessageBoardMembershipView`

NewMessageBoardMembershipViewWithDefaults instantiates a new MessageBoardMembershipView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageBoardMembershipView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageBoardMembershipView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageBoardMembershipView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MessageBoardMembershipView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMember

`func (o *MessageBoardMembershipView) GetMember() UsernameView`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *MessageBoardMembershipView) GetMemberOk() (*UsernameView, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *MessageBoardMembershipView) SetMember(v UsernameView)`

SetMember sets Member field to given value.

### HasMember

`func (o *MessageBoardMembershipView) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetMessageBoard

`func (o *MessageBoardMembershipView) GetMessageBoard() MessageBoardView`

GetMessageBoard returns the MessageBoard field if non-nil, zero value otherwise.

### GetMessageBoardOk

`func (o *MessageBoardMembershipView) GetMessageBoardOk() (*MessageBoardView, bool)`

GetMessageBoardOk returns a tuple with the MessageBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBoard

`func (o *MessageBoardMembershipView) SetMessageBoard(v MessageBoardView)`

SetMessageBoard sets MessageBoard field to given value.

### HasMessageBoard

`func (o *MessageBoardMembershipView) HasMessageBoard() bool`

HasMessageBoard returns a boolean if a field has been set.

### GetOnline

`func (o *MessageBoardMembershipView) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *MessageBoardMembershipView) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *MessageBoardMembershipView) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *MessageBoardMembershipView) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetRole

`func (o *MessageBoardMembershipView) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MessageBoardMembershipView) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MessageBoardMembershipView) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *MessageBoardMembershipView) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetVersion

`func (o *MessageBoardMembershipView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MessageBoardMembershipView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MessageBoardMembershipView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MessageBoardMembershipView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


