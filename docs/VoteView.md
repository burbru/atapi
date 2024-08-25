# VoteView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Voices** | Pointer to **int64** |  | [optional] 
**Voter** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 

## Methods

### NewVoteView

`func NewVoteView() *VoteView`

NewVoteView instantiates a new VoteView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoteViewWithDefaults

`func NewVoteViewWithDefaults() *VoteView`

NewVoteViewWithDefaults instantiates a new VoteView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VoteView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VoteView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VoteView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VoteView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVoices

`func (o *VoteView) GetVoices() int64`

GetVoices returns the Voices field if non-nil, zero value otherwise.

### GetVoicesOk

`func (o *VoteView) GetVoicesOk() (*int64, bool)`

GetVoicesOk returns a tuple with the Voices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoices

`func (o *VoteView) SetVoices(v int64)`

SetVoices sets Voices field to given value.

### HasVoices

`func (o *VoteView) HasVoices() bool`

HasVoices returns a boolean if a field has been set.

### GetVoter

`func (o *VoteView) GetVoter() UsernameView`

GetVoter returns the Voter field if non-nil, zero value otherwise.

### GetVoterOk

`func (o *VoteView) GetVoterOk() (*UsernameView, bool)`

GetVoterOk returns a tuple with the Voter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoter

`func (o *VoteView) SetVoter(v UsernameView)`

SetVoter sets Voter field to given value.

### HasVoter

`func (o *VoteView) HasVoter() bool`

HasVoter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


