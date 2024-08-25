# AbstractPollView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbstentionRule** | Pointer to **string** |  | [optional] 
**CastVotesPercentage** | Pointer to **float32** |  | [optional] 
**EndDate** | Pointer to **int64** |  | [optional] 
**Group** | Pointer to [**[]VoiceNumberView**](VoiceNumberView.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Motion** | Pointer to **string** |  | [optional] 
**PollInitiator** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**ResultExpireDate** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **int64** |  | [optional] 
**TotalNumberOfCastVotes** | Pointer to **int64** |  | [optional] 
**TotalNumberOfVoices** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**Votes** | Pointer to [**[]VoteView**](VoteView.md) |  | [optional] 

## Methods

### NewAbstractPollView

`func NewAbstractPollView() *AbstractPollView`

NewAbstractPollView instantiates a new AbstractPollView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractPollViewWithDefaults

`func NewAbstractPollViewWithDefaults() *AbstractPollView`

NewAbstractPollViewWithDefaults instantiates a new AbstractPollView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbstentionRule

`func (o *AbstractPollView) GetAbstentionRule() string`

GetAbstentionRule returns the AbstentionRule field if non-nil, zero value otherwise.

### GetAbstentionRuleOk

`func (o *AbstractPollView) GetAbstentionRuleOk() (*string, bool)`

GetAbstentionRuleOk returns a tuple with the AbstentionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstentionRule

`func (o *AbstractPollView) SetAbstentionRule(v string)`

SetAbstentionRule sets AbstentionRule field to given value.

### HasAbstentionRule

`func (o *AbstractPollView) HasAbstentionRule() bool`

HasAbstentionRule returns a boolean if a field has been set.

### GetCastVotesPercentage

`func (o *AbstractPollView) GetCastVotesPercentage() float32`

GetCastVotesPercentage returns the CastVotesPercentage field if non-nil, zero value otherwise.

### GetCastVotesPercentageOk

`func (o *AbstractPollView) GetCastVotesPercentageOk() (*float32, bool)`

GetCastVotesPercentageOk returns a tuple with the CastVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastVotesPercentage

`func (o *AbstractPollView) SetCastVotesPercentage(v float32)`

SetCastVotesPercentage sets CastVotesPercentage field to given value.

### HasCastVotesPercentage

`func (o *AbstractPollView) HasCastVotesPercentage() bool`

HasCastVotesPercentage returns a boolean if a field has been set.

### GetEndDate

`func (o *AbstractPollView) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AbstractPollView) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AbstractPollView) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AbstractPollView) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetGroup

`func (o *AbstractPollView) GetGroup() []VoiceNumberView`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AbstractPollView) GetGroupOk() (*[]VoiceNumberView, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AbstractPollView) SetGroup(v []VoiceNumberView)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AbstractPollView) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *AbstractPollView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AbstractPollView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AbstractPollView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AbstractPollView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMotion

`func (o *AbstractPollView) GetMotion() string`

GetMotion returns the Motion field if non-nil, zero value otherwise.

### GetMotionOk

`func (o *AbstractPollView) GetMotionOk() (*string, bool)`

GetMotionOk returns a tuple with the Motion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotion

`func (o *AbstractPollView) SetMotion(v string)`

SetMotion sets Motion field to given value.

### HasMotion

`func (o *AbstractPollView) HasMotion() bool`

HasMotion returns a boolean if a field has been set.

### GetPollInitiator

`func (o *AbstractPollView) GetPollInitiator() UsernameView`

GetPollInitiator returns the PollInitiator field if non-nil, zero value otherwise.

### GetPollInitiatorOk

`func (o *AbstractPollView) GetPollInitiatorOk() (*UsernameView, bool)`

GetPollInitiatorOk returns a tuple with the PollInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollInitiator

`func (o *AbstractPollView) SetPollInitiator(v UsernameView)`

SetPollInitiator sets PollInitiator field to given value.

### HasPollInitiator

`func (o *AbstractPollView) HasPollInitiator() bool`

HasPollInitiator returns a boolean if a field has been set.

### GetResultExpireDate

`func (o *AbstractPollView) GetResultExpireDate() int64`

GetResultExpireDate returns the ResultExpireDate field if non-nil, zero value otherwise.

### GetResultExpireDateOk

`func (o *AbstractPollView) GetResultExpireDateOk() (*int64, bool)`

GetResultExpireDateOk returns a tuple with the ResultExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultExpireDate

`func (o *AbstractPollView) SetResultExpireDate(v int64)`

SetResultExpireDate sets ResultExpireDate field to given value.

### HasResultExpireDate

`func (o *AbstractPollView) HasResultExpireDate() bool`

HasResultExpireDate returns a boolean if a field has been set.

### GetStartDate

`func (o *AbstractPollView) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AbstractPollView) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AbstractPollView) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AbstractPollView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTotalNumberOfCastVotes

`func (o *AbstractPollView) GetTotalNumberOfCastVotes() int64`

GetTotalNumberOfCastVotes returns the TotalNumberOfCastVotes field if non-nil, zero value otherwise.

### GetTotalNumberOfCastVotesOk

`func (o *AbstractPollView) GetTotalNumberOfCastVotesOk() (*int64, bool)`

GetTotalNumberOfCastVotesOk returns a tuple with the TotalNumberOfCastVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfCastVotes

`func (o *AbstractPollView) SetTotalNumberOfCastVotes(v int64)`

SetTotalNumberOfCastVotes sets TotalNumberOfCastVotes field to given value.

### HasTotalNumberOfCastVotes

`func (o *AbstractPollView) HasTotalNumberOfCastVotes() bool`

HasTotalNumberOfCastVotes returns a boolean if a field has been set.

### GetTotalNumberOfVoices

`func (o *AbstractPollView) GetTotalNumberOfVoices() int64`

GetTotalNumberOfVoices returns the TotalNumberOfVoices field if non-nil, zero value otherwise.

### GetTotalNumberOfVoicesOk

`func (o *AbstractPollView) GetTotalNumberOfVoicesOk() (*int64, bool)`

GetTotalNumberOfVoicesOk returns a tuple with the TotalNumberOfVoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfVoices

`func (o *AbstractPollView) SetTotalNumberOfVoices(v int64)`

SetTotalNumberOfVoices sets TotalNumberOfVoices field to given value.

### HasTotalNumberOfVoices

`func (o *AbstractPollView) HasTotalNumberOfVoices() bool`

HasTotalNumberOfVoices returns a boolean if a field has been set.

### GetType

`func (o *AbstractPollView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AbstractPollView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AbstractPollView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AbstractPollView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *AbstractPollView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AbstractPollView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AbstractPollView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AbstractPollView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVotes

`func (o *AbstractPollView) GetVotes() []VoteView`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *AbstractPollView) GetVotesOk() (*[]VoteView, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *AbstractPollView) SetVotes(v []VoteView)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *AbstractPollView) HasVotes() bool`

HasVotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


