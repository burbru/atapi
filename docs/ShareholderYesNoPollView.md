# ShareholderYesNoPollView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbstentionRule** | Pointer to **string** |  | [optional] 
**ApprovalVotesPercentage** | Pointer to **float32** |  | [optional] 
**CastApprovalVotesPercentage** | Pointer to **float32** |  | [optional] 
**CastRefusalVotesPercentage** | Pointer to **float32** |  | [optional] 
**CastVotesPercentage** | Pointer to **float32** |  | [optional] 
**Company** | Pointer to [**CompactCompanyView**](CompactCompanyView.md) |  | [optional] 
**EndDate** | Pointer to **int64** |  | [optional] 
**Group** | Pointer to [**[]VoiceNumberView**](VoiceNumberView.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LeftVotesPercentage** | Pointer to **float32** |  | [optional] 
**Motion** | Pointer to **string** |  | [optional] 
**PollInitiator** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**RefusalVotesPercentage** | Pointer to **float32** |  | [optional] 
**ResultExpireDate** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **int64** |  | [optional] 
**TotalNumberOfCastVotes** | Pointer to **int64** |  | [optional] 
**TotalNumberOfVoices** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**Votes** | Pointer to [**[]VoteView**](VoteView.md) |  | [optional] 

## Methods

### NewShareholderYesNoPollView

`func NewShareholderYesNoPollView() *ShareholderYesNoPollView`

NewShareholderYesNoPollView instantiates a new ShareholderYesNoPollView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareholderYesNoPollViewWithDefaults

`func NewShareholderYesNoPollViewWithDefaults() *ShareholderYesNoPollView`

NewShareholderYesNoPollViewWithDefaults instantiates a new ShareholderYesNoPollView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbstentionRule

`func (o *ShareholderYesNoPollView) GetAbstentionRule() string`

GetAbstentionRule returns the AbstentionRule field if non-nil, zero value otherwise.

### GetAbstentionRuleOk

`func (o *ShareholderYesNoPollView) GetAbstentionRuleOk() (*string, bool)`

GetAbstentionRuleOk returns a tuple with the AbstentionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstentionRule

`func (o *ShareholderYesNoPollView) SetAbstentionRule(v string)`

SetAbstentionRule sets AbstentionRule field to given value.

### HasAbstentionRule

`func (o *ShareholderYesNoPollView) HasAbstentionRule() bool`

HasAbstentionRule returns a boolean if a field has been set.

### GetApprovalVotesPercentage

`func (o *ShareholderYesNoPollView) GetApprovalVotesPercentage() float32`

GetApprovalVotesPercentage returns the ApprovalVotesPercentage field if non-nil, zero value otherwise.

### GetApprovalVotesPercentageOk

`func (o *ShareholderYesNoPollView) GetApprovalVotesPercentageOk() (*float32, bool)`

GetApprovalVotesPercentageOk returns a tuple with the ApprovalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalVotesPercentage

`func (o *ShareholderYesNoPollView) SetApprovalVotesPercentage(v float32)`

SetApprovalVotesPercentage sets ApprovalVotesPercentage field to given value.

### HasApprovalVotesPercentage

`func (o *ShareholderYesNoPollView) HasApprovalVotesPercentage() bool`

HasApprovalVotesPercentage returns a boolean if a field has been set.

### GetCastApprovalVotesPercentage

`func (o *ShareholderYesNoPollView) GetCastApprovalVotesPercentage() float32`

GetCastApprovalVotesPercentage returns the CastApprovalVotesPercentage field if non-nil, zero value otherwise.

### GetCastApprovalVotesPercentageOk

`func (o *ShareholderYesNoPollView) GetCastApprovalVotesPercentageOk() (*float32, bool)`

GetCastApprovalVotesPercentageOk returns a tuple with the CastApprovalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastApprovalVotesPercentage

`func (o *ShareholderYesNoPollView) SetCastApprovalVotesPercentage(v float32)`

SetCastApprovalVotesPercentage sets CastApprovalVotesPercentage field to given value.

### HasCastApprovalVotesPercentage

`func (o *ShareholderYesNoPollView) HasCastApprovalVotesPercentage() bool`

HasCastApprovalVotesPercentage returns a boolean if a field has been set.

### GetCastRefusalVotesPercentage

`func (o *ShareholderYesNoPollView) GetCastRefusalVotesPercentage() float32`

GetCastRefusalVotesPercentage returns the CastRefusalVotesPercentage field if non-nil, zero value otherwise.

### GetCastRefusalVotesPercentageOk

`func (o *ShareholderYesNoPollView) GetCastRefusalVotesPercentageOk() (*float32, bool)`

GetCastRefusalVotesPercentageOk returns a tuple with the CastRefusalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastRefusalVotesPercentage

`func (o *ShareholderYesNoPollView) SetCastRefusalVotesPercentage(v float32)`

SetCastRefusalVotesPercentage sets CastRefusalVotesPercentage field to given value.

### HasCastRefusalVotesPercentage

`func (o *ShareholderYesNoPollView) HasCastRefusalVotesPercentage() bool`

HasCastRefusalVotesPercentage returns a boolean if a field has been set.

### GetCastVotesPercentage

`func (o *ShareholderYesNoPollView) GetCastVotesPercentage() float32`

GetCastVotesPercentage returns the CastVotesPercentage field if non-nil, zero value otherwise.

### GetCastVotesPercentageOk

`func (o *ShareholderYesNoPollView) GetCastVotesPercentageOk() (*float32, bool)`

GetCastVotesPercentageOk returns a tuple with the CastVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastVotesPercentage

`func (o *ShareholderYesNoPollView) SetCastVotesPercentage(v float32)`

SetCastVotesPercentage sets CastVotesPercentage field to given value.

### HasCastVotesPercentage

`func (o *ShareholderYesNoPollView) HasCastVotesPercentage() bool`

HasCastVotesPercentage returns a boolean if a field has been set.

### GetCompany

`func (o *ShareholderYesNoPollView) GetCompany() CompactCompanyView`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ShareholderYesNoPollView) GetCompanyOk() (*CompactCompanyView, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ShareholderYesNoPollView) SetCompany(v CompactCompanyView)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ShareholderYesNoPollView) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetEndDate

`func (o *ShareholderYesNoPollView) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ShareholderYesNoPollView) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ShareholderYesNoPollView) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ShareholderYesNoPollView) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetGroup

`func (o *ShareholderYesNoPollView) GetGroup() []VoiceNumberView`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ShareholderYesNoPollView) GetGroupOk() (*[]VoiceNumberView, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ShareholderYesNoPollView) SetGroup(v []VoiceNumberView)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ShareholderYesNoPollView) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *ShareholderYesNoPollView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShareholderYesNoPollView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShareholderYesNoPollView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShareholderYesNoPollView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLeftVotesPercentage

`func (o *ShareholderYesNoPollView) GetLeftVotesPercentage() float32`

GetLeftVotesPercentage returns the LeftVotesPercentage field if non-nil, zero value otherwise.

### GetLeftVotesPercentageOk

`func (o *ShareholderYesNoPollView) GetLeftVotesPercentageOk() (*float32, bool)`

GetLeftVotesPercentageOk returns a tuple with the LeftVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftVotesPercentage

`func (o *ShareholderYesNoPollView) SetLeftVotesPercentage(v float32)`

SetLeftVotesPercentage sets LeftVotesPercentage field to given value.

### HasLeftVotesPercentage

`func (o *ShareholderYesNoPollView) HasLeftVotesPercentage() bool`

HasLeftVotesPercentage returns a boolean if a field has been set.

### GetMotion

`func (o *ShareholderYesNoPollView) GetMotion() string`

GetMotion returns the Motion field if non-nil, zero value otherwise.

### GetMotionOk

`func (o *ShareholderYesNoPollView) GetMotionOk() (*string, bool)`

GetMotionOk returns a tuple with the Motion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotion

`func (o *ShareholderYesNoPollView) SetMotion(v string)`

SetMotion sets Motion field to given value.

### HasMotion

`func (o *ShareholderYesNoPollView) HasMotion() bool`

HasMotion returns a boolean if a field has been set.

### GetPollInitiator

`func (o *ShareholderYesNoPollView) GetPollInitiator() UsernameView`

GetPollInitiator returns the PollInitiator field if non-nil, zero value otherwise.

### GetPollInitiatorOk

`func (o *ShareholderYesNoPollView) GetPollInitiatorOk() (*UsernameView, bool)`

GetPollInitiatorOk returns a tuple with the PollInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollInitiator

`func (o *ShareholderYesNoPollView) SetPollInitiator(v UsernameView)`

SetPollInitiator sets PollInitiator field to given value.

### HasPollInitiator

`func (o *ShareholderYesNoPollView) HasPollInitiator() bool`

HasPollInitiator returns a boolean if a field has been set.

### GetRefusalVotesPercentage

`func (o *ShareholderYesNoPollView) GetRefusalVotesPercentage() float32`

GetRefusalVotesPercentage returns the RefusalVotesPercentage field if non-nil, zero value otherwise.

### GetRefusalVotesPercentageOk

`func (o *ShareholderYesNoPollView) GetRefusalVotesPercentageOk() (*float32, bool)`

GetRefusalVotesPercentageOk returns a tuple with the RefusalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalVotesPercentage

`func (o *ShareholderYesNoPollView) SetRefusalVotesPercentage(v float32)`

SetRefusalVotesPercentage sets RefusalVotesPercentage field to given value.

### HasRefusalVotesPercentage

`func (o *ShareholderYesNoPollView) HasRefusalVotesPercentage() bool`

HasRefusalVotesPercentage returns a boolean if a field has been set.

### GetResultExpireDate

`func (o *ShareholderYesNoPollView) GetResultExpireDate() int64`

GetResultExpireDate returns the ResultExpireDate field if non-nil, zero value otherwise.

### GetResultExpireDateOk

`func (o *ShareholderYesNoPollView) GetResultExpireDateOk() (*int64, bool)`

GetResultExpireDateOk returns a tuple with the ResultExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultExpireDate

`func (o *ShareholderYesNoPollView) SetResultExpireDate(v int64)`

SetResultExpireDate sets ResultExpireDate field to given value.

### HasResultExpireDate

`func (o *ShareholderYesNoPollView) HasResultExpireDate() bool`

HasResultExpireDate returns a boolean if a field has been set.

### GetStartDate

`func (o *ShareholderYesNoPollView) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ShareholderYesNoPollView) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ShareholderYesNoPollView) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ShareholderYesNoPollView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTotalNumberOfCastVotes

`func (o *ShareholderYesNoPollView) GetTotalNumberOfCastVotes() int64`

GetTotalNumberOfCastVotes returns the TotalNumberOfCastVotes field if non-nil, zero value otherwise.

### GetTotalNumberOfCastVotesOk

`func (o *ShareholderYesNoPollView) GetTotalNumberOfCastVotesOk() (*int64, bool)`

GetTotalNumberOfCastVotesOk returns a tuple with the TotalNumberOfCastVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfCastVotes

`func (o *ShareholderYesNoPollView) SetTotalNumberOfCastVotes(v int64)`

SetTotalNumberOfCastVotes sets TotalNumberOfCastVotes field to given value.

### HasTotalNumberOfCastVotes

`func (o *ShareholderYesNoPollView) HasTotalNumberOfCastVotes() bool`

HasTotalNumberOfCastVotes returns a boolean if a field has been set.

### GetTotalNumberOfVoices

`func (o *ShareholderYesNoPollView) GetTotalNumberOfVoices() int64`

GetTotalNumberOfVoices returns the TotalNumberOfVoices field if non-nil, zero value otherwise.

### GetTotalNumberOfVoicesOk

`func (o *ShareholderYesNoPollView) GetTotalNumberOfVoicesOk() (*int64, bool)`

GetTotalNumberOfVoicesOk returns a tuple with the TotalNumberOfVoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfVoices

`func (o *ShareholderYesNoPollView) SetTotalNumberOfVoices(v int64)`

SetTotalNumberOfVoices sets TotalNumberOfVoices field to given value.

### HasTotalNumberOfVoices

`func (o *ShareholderYesNoPollView) HasTotalNumberOfVoices() bool`

HasTotalNumberOfVoices returns a boolean if a field has been set.

### GetType

`func (o *ShareholderYesNoPollView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShareholderYesNoPollView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShareholderYesNoPollView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ShareholderYesNoPollView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *ShareholderYesNoPollView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ShareholderYesNoPollView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ShareholderYesNoPollView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ShareholderYesNoPollView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVotes

`func (o *ShareholderYesNoPollView) GetVotes() []VoteView`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *ShareholderYesNoPollView) GetVotesOk() (*[]VoteView, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *ShareholderYesNoPollView) SetVotes(v []VoteView)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *ShareholderYesNoPollView) HasVotes() bool`

HasVotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


