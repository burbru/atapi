# DividendPaymentPollView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbstentionRule** | Pointer to **string** |  | [optional] 
**ApprovalVotesPercentage** | Pointer to **float32** |  | [optional] 
**CastApprovalVotesPercentage** | Pointer to **float32** |  | [optional] 
**CastRefusalVotesPercentage** | Pointer to **float32** |  | [optional] 
**CastVotesPercentage** | Pointer to **float32** |  | [optional] 
**Company** | Pointer to [**CompactCompanyView**](CompactCompanyView.md) |  | [optional] 
**DividendPaymentStartDate** | Pointer to **int64** |  | [optional] 
**EndDate** | Pointer to **int64** |  | [optional] 
**Group** | Pointer to [**[]VoiceNumberView**](VoiceNumberView.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LeftVotesPercentage** | Pointer to **float32** |  | [optional] 
**MaximalCashVolume** | Pointer to **float32** |  | [optional] 
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

### NewDividendPaymentPollView

`func NewDividendPaymentPollView() *DividendPaymentPollView`

NewDividendPaymentPollView instantiates a new DividendPaymentPollView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDividendPaymentPollViewWithDefaults

`func NewDividendPaymentPollViewWithDefaults() *DividendPaymentPollView`

NewDividendPaymentPollViewWithDefaults instantiates a new DividendPaymentPollView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbstentionRule

`func (o *DividendPaymentPollView) GetAbstentionRule() string`

GetAbstentionRule returns the AbstentionRule field if non-nil, zero value otherwise.

### GetAbstentionRuleOk

`func (o *DividendPaymentPollView) GetAbstentionRuleOk() (*string, bool)`

GetAbstentionRuleOk returns a tuple with the AbstentionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstentionRule

`func (o *DividendPaymentPollView) SetAbstentionRule(v string)`

SetAbstentionRule sets AbstentionRule field to given value.

### HasAbstentionRule

`func (o *DividendPaymentPollView) HasAbstentionRule() bool`

HasAbstentionRule returns a boolean if a field has been set.

### GetApprovalVotesPercentage

`func (o *DividendPaymentPollView) GetApprovalVotesPercentage() float32`

GetApprovalVotesPercentage returns the ApprovalVotesPercentage field if non-nil, zero value otherwise.

### GetApprovalVotesPercentageOk

`func (o *DividendPaymentPollView) GetApprovalVotesPercentageOk() (*float32, bool)`

GetApprovalVotesPercentageOk returns a tuple with the ApprovalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalVotesPercentage

`func (o *DividendPaymentPollView) SetApprovalVotesPercentage(v float32)`

SetApprovalVotesPercentage sets ApprovalVotesPercentage field to given value.

### HasApprovalVotesPercentage

`func (o *DividendPaymentPollView) HasApprovalVotesPercentage() bool`

HasApprovalVotesPercentage returns a boolean if a field has been set.

### GetCastApprovalVotesPercentage

`func (o *DividendPaymentPollView) GetCastApprovalVotesPercentage() float32`

GetCastApprovalVotesPercentage returns the CastApprovalVotesPercentage field if non-nil, zero value otherwise.

### GetCastApprovalVotesPercentageOk

`func (o *DividendPaymentPollView) GetCastApprovalVotesPercentageOk() (*float32, bool)`

GetCastApprovalVotesPercentageOk returns a tuple with the CastApprovalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastApprovalVotesPercentage

`func (o *DividendPaymentPollView) SetCastApprovalVotesPercentage(v float32)`

SetCastApprovalVotesPercentage sets CastApprovalVotesPercentage field to given value.

### HasCastApprovalVotesPercentage

`func (o *DividendPaymentPollView) HasCastApprovalVotesPercentage() bool`

HasCastApprovalVotesPercentage returns a boolean if a field has been set.

### GetCastRefusalVotesPercentage

`func (o *DividendPaymentPollView) GetCastRefusalVotesPercentage() float32`

GetCastRefusalVotesPercentage returns the CastRefusalVotesPercentage field if non-nil, zero value otherwise.

### GetCastRefusalVotesPercentageOk

`func (o *DividendPaymentPollView) GetCastRefusalVotesPercentageOk() (*float32, bool)`

GetCastRefusalVotesPercentageOk returns a tuple with the CastRefusalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastRefusalVotesPercentage

`func (o *DividendPaymentPollView) SetCastRefusalVotesPercentage(v float32)`

SetCastRefusalVotesPercentage sets CastRefusalVotesPercentage field to given value.

### HasCastRefusalVotesPercentage

`func (o *DividendPaymentPollView) HasCastRefusalVotesPercentage() bool`

HasCastRefusalVotesPercentage returns a boolean if a field has been set.

### GetCastVotesPercentage

`func (o *DividendPaymentPollView) GetCastVotesPercentage() float32`

GetCastVotesPercentage returns the CastVotesPercentage field if non-nil, zero value otherwise.

### GetCastVotesPercentageOk

`func (o *DividendPaymentPollView) GetCastVotesPercentageOk() (*float32, bool)`

GetCastVotesPercentageOk returns a tuple with the CastVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastVotesPercentage

`func (o *DividendPaymentPollView) SetCastVotesPercentage(v float32)`

SetCastVotesPercentage sets CastVotesPercentage field to given value.

### HasCastVotesPercentage

`func (o *DividendPaymentPollView) HasCastVotesPercentage() bool`

HasCastVotesPercentage returns a boolean if a field has been set.

### GetCompany

`func (o *DividendPaymentPollView) GetCompany() CompactCompanyView`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *DividendPaymentPollView) GetCompanyOk() (*CompactCompanyView, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *DividendPaymentPollView) SetCompany(v CompactCompanyView)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *DividendPaymentPollView) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetDividendPaymentStartDate

`func (o *DividendPaymentPollView) GetDividendPaymentStartDate() int64`

GetDividendPaymentStartDate returns the DividendPaymentStartDate field if non-nil, zero value otherwise.

### GetDividendPaymentStartDateOk

`func (o *DividendPaymentPollView) GetDividendPaymentStartDateOk() (*int64, bool)`

GetDividendPaymentStartDateOk returns a tuple with the DividendPaymentStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendPaymentStartDate

`func (o *DividendPaymentPollView) SetDividendPaymentStartDate(v int64)`

SetDividendPaymentStartDate sets DividendPaymentStartDate field to given value.

### HasDividendPaymentStartDate

`func (o *DividendPaymentPollView) HasDividendPaymentStartDate() bool`

HasDividendPaymentStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *DividendPaymentPollView) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DividendPaymentPollView) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DividendPaymentPollView) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DividendPaymentPollView) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetGroup

`func (o *DividendPaymentPollView) GetGroup() []VoiceNumberView`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DividendPaymentPollView) GetGroupOk() (*[]VoiceNumberView, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DividendPaymentPollView) SetGroup(v []VoiceNumberView)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DividendPaymentPollView) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *DividendPaymentPollView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DividendPaymentPollView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DividendPaymentPollView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DividendPaymentPollView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLeftVotesPercentage

`func (o *DividendPaymentPollView) GetLeftVotesPercentage() float32`

GetLeftVotesPercentage returns the LeftVotesPercentage field if non-nil, zero value otherwise.

### GetLeftVotesPercentageOk

`func (o *DividendPaymentPollView) GetLeftVotesPercentageOk() (*float32, bool)`

GetLeftVotesPercentageOk returns a tuple with the LeftVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftVotesPercentage

`func (o *DividendPaymentPollView) SetLeftVotesPercentage(v float32)`

SetLeftVotesPercentage sets LeftVotesPercentage field to given value.

### HasLeftVotesPercentage

`func (o *DividendPaymentPollView) HasLeftVotesPercentage() bool`

HasLeftVotesPercentage returns a boolean if a field has been set.

### GetMaximalCashVolume

`func (o *DividendPaymentPollView) GetMaximalCashVolume() float32`

GetMaximalCashVolume returns the MaximalCashVolume field if non-nil, zero value otherwise.

### GetMaximalCashVolumeOk

`func (o *DividendPaymentPollView) GetMaximalCashVolumeOk() (*float32, bool)`

GetMaximalCashVolumeOk returns a tuple with the MaximalCashVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximalCashVolume

`func (o *DividendPaymentPollView) SetMaximalCashVolume(v float32)`

SetMaximalCashVolume sets MaximalCashVolume field to given value.

### HasMaximalCashVolume

`func (o *DividendPaymentPollView) HasMaximalCashVolume() bool`

HasMaximalCashVolume returns a boolean if a field has been set.

### GetMotion

`func (o *DividendPaymentPollView) GetMotion() string`

GetMotion returns the Motion field if non-nil, zero value otherwise.

### GetMotionOk

`func (o *DividendPaymentPollView) GetMotionOk() (*string, bool)`

GetMotionOk returns a tuple with the Motion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotion

`func (o *DividendPaymentPollView) SetMotion(v string)`

SetMotion sets Motion field to given value.

### HasMotion

`func (o *DividendPaymentPollView) HasMotion() bool`

HasMotion returns a boolean if a field has been set.

### GetPollInitiator

`func (o *DividendPaymentPollView) GetPollInitiator() UsernameView`

GetPollInitiator returns the PollInitiator field if non-nil, zero value otherwise.

### GetPollInitiatorOk

`func (o *DividendPaymentPollView) GetPollInitiatorOk() (*UsernameView, bool)`

GetPollInitiatorOk returns a tuple with the PollInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollInitiator

`func (o *DividendPaymentPollView) SetPollInitiator(v UsernameView)`

SetPollInitiator sets PollInitiator field to given value.

### HasPollInitiator

`func (o *DividendPaymentPollView) HasPollInitiator() bool`

HasPollInitiator returns a boolean if a field has been set.

### GetRefusalVotesPercentage

`func (o *DividendPaymentPollView) GetRefusalVotesPercentage() float32`

GetRefusalVotesPercentage returns the RefusalVotesPercentage field if non-nil, zero value otherwise.

### GetRefusalVotesPercentageOk

`func (o *DividendPaymentPollView) GetRefusalVotesPercentageOk() (*float32, bool)`

GetRefusalVotesPercentageOk returns a tuple with the RefusalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalVotesPercentage

`func (o *DividendPaymentPollView) SetRefusalVotesPercentage(v float32)`

SetRefusalVotesPercentage sets RefusalVotesPercentage field to given value.

### HasRefusalVotesPercentage

`func (o *DividendPaymentPollView) HasRefusalVotesPercentage() bool`

HasRefusalVotesPercentage returns a boolean if a field has been set.

### GetResultExpireDate

`func (o *DividendPaymentPollView) GetResultExpireDate() int64`

GetResultExpireDate returns the ResultExpireDate field if non-nil, zero value otherwise.

### GetResultExpireDateOk

`func (o *DividendPaymentPollView) GetResultExpireDateOk() (*int64, bool)`

GetResultExpireDateOk returns a tuple with the ResultExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultExpireDate

`func (o *DividendPaymentPollView) SetResultExpireDate(v int64)`

SetResultExpireDate sets ResultExpireDate field to given value.

### HasResultExpireDate

`func (o *DividendPaymentPollView) HasResultExpireDate() bool`

HasResultExpireDate returns a boolean if a field has been set.

### GetStartDate

`func (o *DividendPaymentPollView) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DividendPaymentPollView) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DividendPaymentPollView) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DividendPaymentPollView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTotalNumberOfCastVotes

`func (o *DividendPaymentPollView) GetTotalNumberOfCastVotes() int64`

GetTotalNumberOfCastVotes returns the TotalNumberOfCastVotes field if non-nil, zero value otherwise.

### GetTotalNumberOfCastVotesOk

`func (o *DividendPaymentPollView) GetTotalNumberOfCastVotesOk() (*int64, bool)`

GetTotalNumberOfCastVotesOk returns a tuple with the TotalNumberOfCastVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfCastVotes

`func (o *DividendPaymentPollView) SetTotalNumberOfCastVotes(v int64)`

SetTotalNumberOfCastVotes sets TotalNumberOfCastVotes field to given value.

### HasTotalNumberOfCastVotes

`func (o *DividendPaymentPollView) HasTotalNumberOfCastVotes() bool`

HasTotalNumberOfCastVotes returns a boolean if a field has been set.

### GetTotalNumberOfVoices

`func (o *DividendPaymentPollView) GetTotalNumberOfVoices() int64`

GetTotalNumberOfVoices returns the TotalNumberOfVoices field if non-nil, zero value otherwise.

### GetTotalNumberOfVoicesOk

`func (o *DividendPaymentPollView) GetTotalNumberOfVoicesOk() (*int64, bool)`

GetTotalNumberOfVoicesOk returns a tuple with the TotalNumberOfVoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfVoices

`func (o *DividendPaymentPollView) SetTotalNumberOfVoices(v int64)`

SetTotalNumberOfVoices sets TotalNumberOfVoices field to given value.

### HasTotalNumberOfVoices

`func (o *DividendPaymentPollView) HasTotalNumberOfVoices() bool`

HasTotalNumberOfVoices returns a boolean if a field has been set.

### GetType

`func (o *DividendPaymentPollView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DividendPaymentPollView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DividendPaymentPollView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DividendPaymentPollView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *DividendPaymentPollView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DividendPaymentPollView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DividendPaymentPollView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DividendPaymentPollView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVotes

`func (o *DividendPaymentPollView) GetVotes() []VoteView`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *DividendPaymentPollView) GetVotesOk() (*[]VoteView, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *DividendPaymentPollView) SetVotes(v []VoteView)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *DividendPaymentPollView) HasVotes() bool`

HasVotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


