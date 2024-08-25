# VoiceNumberView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupMember** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**NumberOfVoices** | Pointer to **int64** |  | [optional] 

## Methods

### NewVoiceNumberView

`func NewVoiceNumberView() *VoiceNumberView`

NewVoiceNumberView instantiates a new VoiceNumberView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceNumberViewWithDefaults

`func NewVoiceNumberViewWithDefaults() *VoiceNumberView`

NewVoiceNumberViewWithDefaults instantiates a new VoiceNumberView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupMember

`func (o *VoiceNumberView) GetGroupMember() UsernameView`

GetGroupMember returns the GroupMember field if non-nil, zero value otherwise.

### GetGroupMemberOk

`func (o *VoiceNumberView) GetGroupMemberOk() (*UsernameView, bool)`

GetGroupMemberOk returns a tuple with the GroupMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMember

`func (o *VoiceNumberView) SetGroupMember(v UsernameView)`

SetGroupMember sets GroupMember field to given value.

### HasGroupMember

`func (o *VoiceNumberView) HasGroupMember() bool`

HasGroupMember returns a boolean if a field has been set.

### GetNumberOfVoices

`func (o *VoiceNumberView) GetNumberOfVoices() int64`

GetNumberOfVoices returns the NumberOfVoices field if non-nil, zero value otherwise.

### GetNumberOfVoicesOk

`func (o *VoiceNumberView) GetNumberOfVoicesOk() (*int64, bool)`

GetNumberOfVoicesOk returns a tuple with the NumberOfVoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfVoices

`func (o *VoiceNumberView) SetNumberOfVoices(v int64)`

SetNumberOfVoices sets NumberOfVoices field to given value.

### HasNumberOfVoices

`func (o *VoiceNumberView) HasNumberOfVoices() bool`

HasNumberOfVoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


