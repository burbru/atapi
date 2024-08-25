# CapitalIncreasePollView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbstentionRule** | Pointer to **string** |  | [optional] 
**ApprovalVotesPercentage** | Pointer to **float32** |  | [optional] 
**CapitalIncreaseEndDate** | Pointer to **int64** |  | [optional] 
**CapitalIncreaseStartDate** | Pointer to **int64** |  | [optional] 
**CapitalIncreaseSubscriptionPeriodEndDate** | Pointer to **int64** |  | [optional] 
**CapitalIncreaseType** | Pointer to **string** |  | [optional] 
**CastApprovalVotesPercentage** | Pointer to **float32** |  | [optional] 
**CastRefusalVotesPercentage** | Pointer to **float32** |  | [optional] 
**CastVotesPercentage** | Pointer to **float32** |  | [optional] 
**Company** | Pointer to [**CompactCompanyView**](CompactCompanyView.md) |  | [optional] 
**EndDate** | Pointer to **int64** |  | [optional] 
**Group** | Pointer to [**[]VoiceNumberView**](VoiceNumberView.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LeftVotesPercentage** | Pointer to **float32** |  | [optional] 
**MinimalCashVolume** | Pointer to **float32** |  | [optional] 
**Motion** | Pointer to **string** |  | [optional] 
**NumberOfShares** | Pointer to **int64** |  | [optional] 
**PollInitiator** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**RefusalVotesPercentage** | Pointer to **float32** |  | [optional] 
**ResultExpireDate** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **int64** |  | [optional] 
**SubscriptionFraction** | Pointer to **int64** |  | [optional] 
**TotalNumberOfCastVotes** | Pointer to **int64** |  | [optional] 
**TotalNumberOfVoices** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**Votes** | Pointer to [**[]VoteView**](VoteView.md) |  | [optional] 

## Methods

### NewCapitalIncreasePollView

`func NewCapitalIncreasePollView() *CapitalIncreasePollView`

NewCapitalIncreasePollView instantiates a new CapitalIncreasePollView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapitalIncreasePollViewWithDefaults

`func NewCapitalIncreasePollViewWithDefaults() *CapitalIncreasePollView`

NewCapitalIncreasePollViewWithDefaults instantiates a new CapitalIncreasePollView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbstentionRule

`func (o *CapitalIncreasePollView) GetAbstentionRule() string`

GetAbstentionRule returns the AbstentionRule field if non-nil, zero value otherwise.

### GetAbstentionRuleOk

`func (o *CapitalIncreasePollView) GetAbstentionRuleOk() (*string, bool)`

GetAbstentionRuleOk returns a tuple with the AbstentionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstentionRule

`func (o *CapitalIncreasePollView) SetAbstentionRule(v string)`

SetAbstentionRule sets AbstentionRule field to given value.

### HasAbstentionRule

`func (o *CapitalIncreasePollView) HasAbstentionRule() bool`

HasAbstentionRule returns a boolean if a field has been set.

### GetApprovalVotesPercentage

`func (o *CapitalIncreasePollView) GetApprovalVotesPercentage() float32`

GetApprovalVotesPercentage returns the ApprovalVotesPercentage field if non-nil, zero value otherwise.

### GetApprovalVotesPercentageOk

`func (o *CapitalIncreasePollView) GetApprovalVotesPercentageOk() (*float32, bool)`

GetApprovalVotesPercentageOk returns a tuple with the ApprovalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalVotesPercentage

`func (o *CapitalIncreasePollView) SetApprovalVotesPercentage(v float32)`

SetApprovalVotesPercentage sets ApprovalVotesPercentage field to given value.

### HasApprovalVotesPercentage

`func (o *CapitalIncreasePollView) HasApprovalVotesPercentage() bool`

HasApprovalVotesPercentage returns a boolean if a field has been set.

### GetCapitalIncreaseEndDate

`func (o *CapitalIncreasePollView) GetCapitalIncreaseEndDate() int64`

GetCapitalIncreaseEndDate returns the CapitalIncreaseEndDate field if non-nil, zero value otherwise.

### GetCapitalIncreaseEndDateOk

`func (o *CapitalIncreasePollView) GetCapitalIncreaseEndDateOk() (*int64, bool)`

GetCapitalIncreaseEndDateOk returns a tuple with the CapitalIncreaseEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalIncreaseEndDate

`func (o *CapitalIncreasePollView) SetCapitalIncreaseEndDate(v int64)`

SetCapitalIncreaseEndDate sets CapitalIncreaseEndDate field to given value.

### HasCapitalIncreaseEndDate

`func (o *CapitalIncreasePollView) HasCapitalIncreaseEndDate() bool`

HasCapitalIncreaseEndDate returns a boolean if a field has been set.

### GetCapitalIncreaseStartDate

`func (o *CapitalIncreasePollView) GetCapitalIncreaseStartDate() int64`

GetCapitalIncreaseStartDate returns the CapitalIncreaseStartDate field if non-nil, zero value otherwise.

### GetCapitalIncreaseStartDateOk

`func (o *CapitalIncreasePollView) GetCapitalIncreaseStartDateOk() (*int64, bool)`

GetCapitalIncreaseStartDateOk returns a tuple with the CapitalIncreaseStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalIncreaseStartDate

`func (o *CapitalIncreasePollView) SetCapitalIncreaseStartDate(v int64)`

SetCapitalIncreaseStartDate sets CapitalIncreaseStartDate field to given value.

### HasCapitalIncreaseStartDate

`func (o *CapitalIncreasePollView) HasCapitalIncreaseStartDate() bool`

HasCapitalIncreaseStartDate returns a boolean if a field has been set.

### GetCapitalIncreaseSubscriptionPeriodEndDate

`func (o *CapitalIncreasePollView) GetCapitalIncreaseSubscriptionPeriodEndDate() int64`

GetCapitalIncreaseSubscriptionPeriodEndDate returns the CapitalIncreaseSubscriptionPeriodEndDate field if non-nil, zero value otherwise.

### GetCapitalIncreaseSubscriptionPeriodEndDateOk

`func (o *CapitalIncreasePollView) GetCapitalIncreaseSubscriptionPeriodEndDateOk() (*int64, bool)`

GetCapitalIncreaseSubscriptionPeriodEndDateOk returns a tuple with the CapitalIncreaseSubscriptionPeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalIncreaseSubscriptionPeriodEndDate

`func (o *CapitalIncreasePollView) SetCapitalIncreaseSubscriptionPeriodEndDate(v int64)`

SetCapitalIncreaseSubscriptionPeriodEndDate sets CapitalIncreaseSubscriptionPeriodEndDate field to given value.

### HasCapitalIncreaseSubscriptionPeriodEndDate

`func (o *CapitalIncreasePollView) HasCapitalIncreaseSubscriptionPeriodEndDate() bool`

HasCapitalIncreaseSubscriptionPeriodEndDate returns a boolean if a field has been set.

### GetCapitalIncreaseType

`func (o *CapitalIncreasePollView) GetCapitalIncreaseType() string`

GetCapitalIncreaseType returns the CapitalIncreaseType field if non-nil, zero value otherwise.

### GetCapitalIncreaseTypeOk

`func (o *CapitalIncreasePollView) GetCapitalIncreaseTypeOk() (*string, bool)`

GetCapitalIncreaseTypeOk returns a tuple with the CapitalIncreaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalIncreaseType

`func (o *CapitalIncreasePollView) SetCapitalIncreaseType(v string)`

SetCapitalIncreaseType sets CapitalIncreaseType field to given value.

### HasCapitalIncreaseType

`func (o *CapitalIncreasePollView) HasCapitalIncreaseType() bool`

HasCapitalIncreaseType returns a boolean if a field has been set.

### GetCastApprovalVotesPercentage

`func (o *CapitalIncreasePollView) GetCastApprovalVotesPercentage() float32`

GetCastApprovalVotesPercentage returns the CastApprovalVotesPercentage field if non-nil, zero value otherwise.

### GetCastApprovalVotesPercentageOk

`func (o *CapitalIncreasePollView) GetCastApprovalVotesPercentageOk() (*float32, bool)`

GetCastApprovalVotesPercentageOk returns a tuple with the CastApprovalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastApprovalVotesPercentage

`func (o *CapitalIncreasePollView) SetCastApprovalVotesPercentage(v float32)`

SetCastApprovalVotesPercentage sets CastApprovalVotesPercentage field to given value.

### HasCastApprovalVotesPercentage

`func (o *CapitalIncreasePollView) HasCastApprovalVotesPercentage() bool`

HasCastApprovalVotesPercentage returns a boolean if a field has been set.

### GetCastRefusalVotesPercentage

`func (o *CapitalIncreasePollView) GetCastRefusalVotesPercentage() float32`

GetCastRefusalVotesPercentage returns the CastRefusalVotesPercentage field if non-nil, zero value otherwise.

### GetCastRefusalVotesPercentageOk

`func (o *CapitalIncreasePollView) GetCastRefusalVotesPercentageOk() (*float32, bool)`

GetCastRefusalVotesPercentageOk returns a tuple with the CastRefusalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastRefusalVotesPercentage

`func (o *CapitalIncreasePollView) SetCastRefusalVotesPercentage(v float32)`

SetCastRefusalVotesPercentage sets CastRefusalVotesPercentage field to given value.

### HasCastRefusalVotesPercentage

`func (o *CapitalIncreasePollView) HasCastRefusalVotesPercentage() bool`

HasCastRefusalVotesPercentage returns a boolean if a field has been set.

### GetCastVotesPercentage

`func (o *CapitalIncreasePollView) GetCastVotesPercentage() float32`

GetCastVotesPercentage returns the CastVotesPercentage field if non-nil, zero value otherwise.

### GetCastVotesPercentageOk

`func (o *CapitalIncreasePollView) GetCastVotesPercentageOk() (*float32, bool)`

GetCastVotesPercentageOk returns a tuple with the CastVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastVotesPercentage

`func (o *CapitalIncreasePollView) SetCastVotesPercentage(v float32)`

SetCastVotesPercentage sets CastVotesPercentage field to given value.

### HasCastVotesPercentage

`func (o *CapitalIncreasePollView) HasCastVotesPercentage() bool`

HasCastVotesPercentage returns a boolean if a field has been set.

### GetCompany

`func (o *CapitalIncreasePollView) GetCompany() CompactCompanyView`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CapitalIncreasePollView) GetCompanyOk() (*CompactCompanyView, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CapitalIncreasePollView) SetCompany(v CompactCompanyView)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CapitalIncreasePollView) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetEndDate

`func (o *CapitalIncreasePollView) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CapitalIncreasePollView) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CapitalIncreasePollView) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CapitalIncreasePollView) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetGroup

`func (o *CapitalIncreasePollView) GetGroup() []VoiceNumberView`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CapitalIncreasePollView) GetGroupOk() (*[]VoiceNumberView, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CapitalIncreasePollView) SetGroup(v []VoiceNumberView)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CapitalIncreasePollView) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *CapitalIncreasePollView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CapitalIncreasePollView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CapitalIncreasePollView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CapitalIncreasePollView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLeftVotesPercentage

`func (o *CapitalIncreasePollView) GetLeftVotesPercentage() float32`

GetLeftVotesPercentage returns the LeftVotesPercentage field if non-nil, zero value otherwise.

### GetLeftVotesPercentageOk

`func (o *CapitalIncreasePollView) GetLeftVotesPercentageOk() (*float32, bool)`

GetLeftVotesPercentageOk returns a tuple with the LeftVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftVotesPercentage

`func (o *CapitalIncreasePollView) SetLeftVotesPercentage(v float32)`

SetLeftVotesPercentage sets LeftVotesPercentage field to given value.

### HasLeftVotesPercentage

`func (o *CapitalIncreasePollView) HasLeftVotesPercentage() bool`

HasLeftVotesPercentage returns a boolean if a field has been set.

### GetMinimalCashVolume

`func (o *CapitalIncreasePollView) GetMinimalCashVolume() float32`

GetMinimalCashVolume returns the MinimalCashVolume field if non-nil, zero value otherwise.

### GetMinimalCashVolumeOk

`func (o *CapitalIncreasePollView) GetMinimalCashVolumeOk() (*float32, bool)`

GetMinimalCashVolumeOk returns a tuple with the MinimalCashVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalCashVolume

`func (o *CapitalIncreasePollView) SetMinimalCashVolume(v float32)`

SetMinimalCashVolume sets MinimalCashVolume field to given value.

### HasMinimalCashVolume

`func (o *CapitalIncreasePollView) HasMinimalCashVolume() bool`

HasMinimalCashVolume returns a boolean if a field has been set.

### GetMotion

`func (o *CapitalIncreasePollView) GetMotion() string`

GetMotion returns the Motion field if non-nil, zero value otherwise.

### GetMotionOk

`func (o *CapitalIncreasePollView) GetMotionOk() (*string, bool)`

GetMotionOk returns a tuple with the Motion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotion

`func (o *CapitalIncreasePollView) SetMotion(v string)`

SetMotion sets Motion field to given value.

### HasMotion

`func (o *CapitalIncreasePollView) HasMotion() bool`

HasMotion returns a boolean if a field has been set.

### GetNumberOfShares

`func (o *CapitalIncreasePollView) GetNumberOfShares() int64`

GetNumberOfShares returns the NumberOfShares field if non-nil, zero value otherwise.

### GetNumberOfSharesOk

`func (o *CapitalIncreasePollView) GetNumberOfSharesOk() (*int64, bool)`

GetNumberOfSharesOk returns a tuple with the NumberOfShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfShares

`func (o *CapitalIncreasePollView) SetNumberOfShares(v int64)`

SetNumberOfShares sets NumberOfShares field to given value.

### HasNumberOfShares

`func (o *CapitalIncreasePollView) HasNumberOfShares() bool`

HasNumberOfShares returns a boolean if a field has been set.

### GetPollInitiator

`func (o *CapitalIncreasePollView) GetPollInitiator() UsernameView`

GetPollInitiator returns the PollInitiator field if non-nil, zero value otherwise.

### GetPollInitiatorOk

`func (o *CapitalIncreasePollView) GetPollInitiatorOk() (*UsernameView, bool)`

GetPollInitiatorOk returns a tuple with the PollInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollInitiator

`func (o *CapitalIncreasePollView) SetPollInitiator(v UsernameView)`

SetPollInitiator sets PollInitiator field to given value.

### HasPollInitiator

`func (o *CapitalIncreasePollView) HasPollInitiator() bool`

HasPollInitiator returns a boolean if a field has been set.

### GetRefusalVotesPercentage

`func (o *CapitalIncreasePollView) GetRefusalVotesPercentage() float32`

GetRefusalVotesPercentage returns the RefusalVotesPercentage field if non-nil, zero value otherwise.

### GetRefusalVotesPercentageOk

`func (o *CapitalIncreasePollView) GetRefusalVotesPercentageOk() (*float32, bool)`

GetRefusalVotesPercentageOk returns a tuple with the RefusalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalVotesPercentage

`func (o *CapitalIncreasePollView) SetRefusalVotesPercentage(v float32)`

SetRefusalVotesPercentage sets RefusalVotesPercentage field to given value.

### HasRefusalVotesPercentage

`func (o *CapitalIncreasePollView) HasRefusalVotesPercentage() bool`

HasRefusalVotesPercentage returns a boolean if a field has been set.

### GetResultExpireDate

`func (o *CapitalIncreasePollView) GetResultExpireDate() int64`

GetResultExpireDate returns the ResultExpireDate field if non-nil, zero value otherwise.

### GetResultExpireDateOk

`func (o *CapitalIncreasePollView) GetResultExpireDateOk() (*int64, bool)`

GetResultExpireDateOk returns a tuple with the ResultExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultExpireDate

`func (o *CapitalIncreasePollView) SetResultExpireDate(v int64)`

SetResultExpireDate sets ResultExpireDate field to given value.

### HasResultExpireDate

`func (o *CapitalIncreasePollView) HasResultExpireDate() bool`

HasResultExpireDate returns a boolean if a field has been set.

### GetStartDate

`func (o *CapitalIncreasePollView) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CapitalIncreasePollView) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CapitalIncreasePollView) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CapitalIncreasePollView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetSubscriptionFraction

`func (o *CapitalIncreasePollView) GetSubscriptionFraction() int64`

GetSubscriptionFraction returns the SubscriptionFraction field if non-nil, zero value otherwise.

### GetSubscriptionFractionOk

`func (o *CapitalIncreasePollView) GetSubscriptionFractionOk() (*int64, bool)`

GetSubscriptionFractionOk returns a tuple with the SubscriptionFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionFraction

`func (o *CapitalIncreasePollView) SetSubscriptionFraction(v int64)`

SetSubscriptionFraction sets SubscriptionFraction field to given value.

### HasSubscriptionFraction

`func (o *CapitalIncreasePollView) HasSubscriptionFraction() bool`

HasSubscriptionFraction returns a boolean if a field has been set.

### GetTotalNumberOfCastVotes

`func (o *CapitalIncreasePollView) GetTotalNumberOfCastVotes() int64`

GetTotalNumberOfCastVotes returns the TotalNumberOfCastVotes field if non-nil, zero value otherwise.

### GetTotalNumberOfCastVotesOk

`func (o *CapitalIncreasePollView) GetTotalNumberOfCastVotesOk() (*int64, bool)`

GetTotalNumberOfCastVotesOk returns a tuple with the TotalNumberOfCastVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfCastVotes

`func (o *CapitalIncreasePollView) SetTotalNumberOfCastVotes(v int64)`

SetTotalNumberOfCastVotes sets TotalNumberOfCastVotes field to given value.

### HasTotalNumberOfCastVotes

`func (o *CapitalIncreasePollView) HasTotalNumberOfCastVotes() bool`

HasTotalNumberOfCastVotes returns a boolean if a field has been set.

### GetTotalNumberOfVoices

`func (o *CapitalIncreasePollView) GetTotalNumberOfVoices() int64`

GetTotalNumberOfVoices returns the TotalNumberOfVoices field if non-nil, zero value otherwise.

### GetTotalNumberOfVoicesOk

`func (o *CapitalIncreasePollView) GetTotalNumberOfVoicesOk() (*int64, bool)`

GetTotalNumberOfVoicesOk returns a tuple with the TotalNumberOfVoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfVoices

`func (o *CapitalIncreasePollView) SetTotalNumberOfVoices(v int64)`

SetTotalNumberOfVoices sets TotalNumberOfVoices field to given value.

### HasTotalNumberOfVoices

`func (o *CapitalIncreasePollView) HasTotalNumberOfVoices() bool`

HasTotalNumberOfVoices returns a boolean if a field has been set.

### GetType

`func (o *CapitalIncreasePollView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CapitalIncreasePollView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CapitalIncreasePollView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CapitalIncreasePollView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *CapitalIncreasePollView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CapitalIncreasePollView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CapitalIncreasePollView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CapitalIncreasePollView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVotes

`func (o *CapitalIncreasePollView) GetVotes() []VoteView`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *CapitalIncreasePollView) GetVotesOk() (*[]VoteView, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *CapitalIncreasePollView) SetVotes(v []VoteView)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *CapitalIncreasePollView) HasVotes() bool`

HasVotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


