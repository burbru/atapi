# AllianceMembershipView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alliance** | Pointer to [**AllianceView**](AllianceView.md) |  | [optional] 
**DateJoined** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Member** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewAllianceMembershipView

`func NewAllianceMembershipView() *AllianceMembershipView`

NewAllianceMembershipView instantiates a new AllianceMembershipView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllianceMembershipViewWithDefaults

`func NewAllianceMembershipViewWithDefaults() *AllianceMembershipView`

NewAllianceMembershipViewWithDefaults instantiates a new AllianceMembershipView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlliance

`func (o *AllianceMembershipView) GetAlliance() AllianceView`

GetAlliance returns the Alliance field if non-nil, zero value otherwise.

### GetAllianceOk

`func (o *AllianceMembershipView) GetAllianceOk() (*AllianceView, bool)`

GetAllianceOk returns a tuple with the Alliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlliance

`func (o *AllianceMembershipView) SetAlliance(v AllianceView)`

SetAlliance sets Alliance field to given value.

### HasAlliance

`func (o *AllianceMembershipView) HasAlliance() bool`

HasAlliance returns a boolean if a field has been set.

### GetDateJoined

`func (o *AllianceMembershipView) GetDateJoined() int64`

GetDateJoined returns the DateJoined field if non-nil, zero value otherwise.

### GetDateJoinedOk

`func (o *AllianceMembershipView) GetDateJoinedOk() (*int64, bool)`

GetDateJoinedOk returns a tuple with the DateJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateJoined

`func (o *AllianceMembershipView) SetDateJoined(v int64)`

SetDateJoined sets DateJoined field to given value.

### HasDateJoined

`func (o *AllianceMembershipView) HasDateJoined() bool`

HasDateJoined returns a boolean if a field has been set.

### GetId

`func (o *AllianceMembershipView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllianceMembershipView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllianceMembershipView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AllianceMembershipView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMember

`func (o *AllianceMembershipView) GetMember() UsernameView`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *AllianceMembershipView) GetMemberOk() (*UsernameView, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *AllianceMembershipView) SetMember(v UsernameView)`

SetMember sets Member field to given value.

### HasMember

`func (o *AllianceMembershipView) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetOnline

`func (o *AllianceMembershipView) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *AllianceMembershipView) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *AllianceMembershipView) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *AllianceMembershipView) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetRole

`func (o *AllianceMembershipView) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AllianceMembershipView) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AllianceMembershipView) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AllianceMembershipView) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetVersion

`func (o *AllianceMembershipView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AllianceMembershipView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AllianceMembershipView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AllianceMembershipView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


