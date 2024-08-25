# CapitalReductionPollView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbstentionRule** | Pointer to **string** |  | [optional] 
**ApprovalVotesPercentage** | Pointer to **float32** |  | [optional] 
**CapitalReductionEndDate** | Pointer to **int64** |  | [optional] 
**CapitalReductionStartDate** | Pointer to **int64** |  | [optional] 
**CastApprovalVotesPercentage** | Pointer to **float32** |  | [optional] 
**CastRefusalVotesPercentage** | Pointer to **float32** |  | [optional] 
**CastVotesPercentage** | Pointer to **float32** |  | [optional] 
**Company** | Pointer to [**CompactCompanyView**](CompactCompanyView.md) |  | [optional] 
**EndDate** | Pointer to **int64** |  | [optional] 
**Group** | Pointer to [**[]VoiceNumberView**](VoiceNumberView.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LeftVotesPercentage** | Pointer to **float32** |  | [optional] 
**MaximalCashVolume** | Pointer to **float32** |  | [optional] 
**Motion** | Pointer to **string** |  | [optional] 
**NumberOfShares** | Pointer to **int64** |  | [optional] 
**PollInitiator** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**RefusalVotesPercentage** | Pointer to **float32** |  | [optional] 
**ResultExpireDate** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **int64** |  | [optional] 
**TotalNumberOfCastVotes** | Pointer to **int64** |  | [optional] 
**TotalNumberOfVoices** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**Votes** | Pointer to [**[]VoteView**](VoteView.md) |  | [optional] 

## Methods

### NewCapitalReductionPollView

`func NewCapitalReductionPollView() *CapitalReductionPollView`

NewCapitalReductionPollView instantiates a new CapitalReductionPollView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapitalReductionPollViewWithDefaults

`func NewCapitalReductionPollViewWithDefaults() *CapitalReductionPollView`

NewCapitalReductionPollViewWithDefaults instantiates a new CapitalReductionPollView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbstentionRule

`func (o *CapitalReductionPollView) GetAbstentionRule() string`

GetAbstentionRule returns the AbstentionRule field if non-nil, zero value otherwise.

### GetAbstentionRuleOk

`func (o *CapitalReductionPollView) GetAbstentionRuleOk() (*string, bool)`

GetAbstentionRuleOk returns a tuple with the AbstentionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstentionRule

`func (o *CapitalReductionPollView) SetAbstentionRule(v string)`

SetAbstentionRule sets AbstentionRule field to given value.

### HasAbstentionRule

`func (o *CapitalReductionPollView) HasAbstentionRule() bool`

HasAbstentionRule returns a boolean if a field has been set.

### GetApprovalVotesPercentage

`func (o *CapitalReductionPollView) GetApprovalVotesPercentage() float32`

GetApprovalVotesPercentage returns the ApprovalVotesPercentage field if non-nil, zero value otherwise.

### GetApprovalVotesPercentageOk

`func (o *CapitalReductionPollView) GetApprovalVotesPercentageOk() (*float32, bool)`

GetApprovalVotesPercentageOk returns a tuple with the ApprovalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalVotesPercentage

`func (o *CapitalReductionPollView) SetApprovalVotesPercentage(v float32)`

SetApprovalVotesPercentage sets ApprovalVotesPercentage field to given value.

### HasApprovalVotesPercentage

`func (o *CapitalReductionPollView) HasApprovalVotesPercentage() bool`

HasApprovalVotesPercentage returns a boolean if a field has been set.

### GetCapitalReductionEndDate

`func (o *CapitalReductionPollView) GetCapitalReductionEndDate() int64`

GetCapitalReductionEndDate returns the CapitalReductionEndDate field if non-nil, zero value otherwise.

### GetCapitalReductionEndDateOk

`func (o *CapitalReductionPollView) GetCapitalReductionEndDateOk() (*int64, bool)`

GetCapitalReductionEndDateOk returns a tuple with the CapitalReductionEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalReductionEndDate

`func (o *CapitalReductionPollView) SetCapitalReductionEndDate(v int64)`

SetCapitalReductionEndDate sets CapitalReductionEndDate field to given value.

### HasCapitalReductionEndDate

`func (o *CapitalReductionPollView) HasCapitalReductionEndDate() bool`

HasCapitalReductionEndDate returns a boolean if a field has been set.

### GetCapitalReductionStartDate

`func (o *CapitalReductionPollView) GetCapitalReductionStartDate() int64`

GetCapitalReductionStartDate returns the CapitalReductionStartDate field if non-nil, zero value otherwise.

### GetCapitalReductionStartDateOk

`func (o *CapitalReductionPollView) GetCapitalReductionStartDateOk() (*int64, bool)`

GetCapitalReductionStartDateOk returns a tuple with the CapitalReductionStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalReductionStartDate

`func (o *CapitalReductionPollView) SetCapitalReductionStartDate(v int64)`

SetCapitalReductionStartDate sets CapitalReductionStartDate field to given value.

### HasCapitalReductionStartDate

`func (o *CapitalReductionPollView) HasCapitalReductionStartDate() bool`

HasCapitalReductionStartDate returns a boolean if a field has been set.

### GetCastApprovalVotesPercentage

`func (o *CapitalReductionPollView) GetCastApprovalVotesPercentage() float32`

GetCastApprovalVotesPercentage returns the CastApprovalVotesPercentage field if non-nil, zero value otherwise.

### GetCastApprovalVotesPercentageOk

`func (o *CapitalReductionPollView) GetCastApprovalVotesPercentageOk() (*float32, bool)`

GetCastApprovalVotesPercentageOk returns a tuple with the CastApprovalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastApprovalVotesPercentage

`func (o *CapitalReductionPollView) SetCastApprovalVotesPercentage(v float32)`

SetCastApprovalVotesPercentage sets CastApprovalVotesPercentage field to given value.

### HasCastApprovalVotesPercentage

`func (o *CapitalReductionPollView) HasCastApprovalVotesPercentage() bool`

HasCastApprovalVotesPercentage returns a boolean if a field has been set.

### GetCastRefusalVotesPercentage

`func (o *CapitalReductionPollView) GetCastRefusalVotesPercentage() float32`

GetCastRefusalVotesPercentage returns the CastRefusalVotesPercentage field if non-nil, zero value otherwise.

### GetCastRefusalVotesPercentageOk

`func (o *CapitalReductionPollView) GetCastRefusalVotesPercentageOk() (*float32, bool)`

GetCastRefusalVotesPercentageOk returns a tuple with the CastRefusalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastRefusalVotesPercentage

`func (o *CapitalReductionPollView) SetCastRefusalVotesPercentage(v float32)`

SetCastRefusalVotesPercentage sets CastRefusalVotesPercentage field to given value.

### HasCastRefusalVotesPercentage

`func (o *CapitalReductionPollView) HasCastRefusalVotesPercentage() bool`

HasCastRefusalVotesPercentage returns a boolean if a field has been set.

### GetCastVotesPercentage

`func (o *CapitalReductionPollView) GetCastVotesPercentage() float32`

GetCastVotesPercentage returns the CastVotesPercentage field if non-nil, zero value otherwise.

### GetCastVotesPercentageOk

`func (o *CapitalReductionPollView) GetCastVotesPercentageOk() (*float32, bool)`

GetCastVotesPercentageOk returns a tuple with the CastVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastVotesPercentage

`func (o *CapitalReductionPollView) SetCastVotesPercentage(v float32)`

SetCastVotesPercentage sets CastVotesPercentage field to given value.

### HasCastVotesPercentage

`func (o *CapitalReductionPollView) HasCastVotesPercentage() bool`

HasCastVotesPercentage returns a boolean if a field has been set.

### GetCompany

`func (o *CapitalReductionPollView) GetCompany() CompactCompanyView`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CapitalReductionPollView) GetCompanyOk() (*CompactCompanyView, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CapitalReductionPollView) SetCompany(v CompactCompanyView)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CapitalReductionPollView) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetEndDate

`func (o *CapitalReductionPollView) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CapitalReductionPollView) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CapitalReductionPollView) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CapitalReductionPollView) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetGroup

`func (o *CapitalReductionPollView) GetGroup() []VoiceNumberView`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CapitalReductionPollView) GetGroupOk() (*[]VoiceNumberView, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CapitalReductionPollView) SetGroup(v []VoiceNumberView)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CapitalReductionPollView) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *CapitalReductionPollView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CapitalReductionPollView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CapitalReductionPollView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CapitalReductionPollView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLeftVotesPercentage

`func (o *CapitalReductionPollView) GetLeftVotesPercentage() float32`

GetLeftVotesPercentage returns the LeftVotesPercentage field if non-nil, zero value otherwise.

### GetLeftVotesPercentageOk

`func (o *CapitalReductionPollView) GetLeftVotesPercentageOk() (*float32, bool)`

GetLeftVotesPercentageOk returns a tuple with the LeftVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftVotesPercentage

`func (o *CapitalReductionPollView) SetLeftVotesPercentage(v float32)`

SetLeftVotesPercentage sets LeftVotesPercentage field to given value.

### HasLeftVotesPercentage

`func (o *CapitalReductionPollView) HasLeftVotesPercentage() bool`

HasLeftVotesPercentage returns a boolean if a field has been set.

### GetMaximalCashVolume

`func (o *CapitalReductionPollView) GetMaximalCashVolume() float32`

GetMaximalCashVolume returns the MaximalCashVolume field if non-nil, zero value otherwise.

### GetMaximalCashVolumeOk

`func (o *CapitalReductionPollView) GetMaximalCashVolumeOk() (*float32, bool)`

GetMaximalCashVolumeOk returns a tuple with the MaximalCashVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximalCashVolume

`func (o *CapitalReductionPollView) SetMaximalCashVolume(v float32)`

SetMaximalCashVolume sets MaximalCashVolume field to given value.

### HasMaximalCashVolume

`func (o *CapitalReductionPollView) HasMaximalCashVolume() bool`

HasMaximalCashVolume returns a boolean if a field has been set.

### GetMotion

`func (o *CapitalReductionPollView) GetMotion() string`

GetMotion returns the Motion field if non-nil, zero value otherwise.

### GetMotionOk

`func (o *CapitalReductionPollView) GetMotionOk() (*string, bool)`

GetMotionOk returns a tuple with the Motion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotion

`func (o *CapitalReductionPollView) SetMotion(v string)`

SetMotion sets Motion field to given value.

### HasMotion

`func (o *CapitalReductionPollView) HasMotion() bool`

HasMotion returns a boolean if a field has been set.

### GetNumberOfShares

`func (o *CapitalReductionPollView) GetNumberOfShares() int64`

GetNumberOfShares returns the NumberOfShares field if non-nil, zero value otherwise.

### GetNumberOfSharesOk

`func (o *CapitalReductionPollView) GetNumberOfSharesOk() (*int64, bool)`

GetNumberOfSharesOk returns a tuple with the NumberOfShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfShares

`func (o *CapitalReductionPollView) SetNumberOfShares(v int64)`

SetNumberOfShares sets NumberOfShares field to given value.

### HasNumberOfShares

`func (o *CapitalReductionPollView) HasNumberOfShares() bool`

HasNumberOfShares returns a boolean if a field has been set.

### GetPollInitiator

`func (o *CapitalReductionPollView) GetPollInitiator() UsernameView`

GetPollInitiator returns the PollInitiator field if non-nil, zero value otherwise.

### GetPollInitiatorOk

`func (o *CapitalReductionPollView) GetPollInitiatorOk() (*UsernameView, bool)`

GetPollInitiatorOk returns a tuple with the PollInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollInitiator

`func (o *CapitalReductionPollView) SetPollInitiator(v UsernameView)`

SetPollInitiator sets PollInitiator field to given value.

### HasPollInitiator

`func (o *CapitalReductionPollView) HasPollInitiator() bool`

HasPollInitiator returns a boolean if a field has been set.

### GetPrice

`func (o *CapitalReductionPollView) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CapitalReductionPollView) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CapitalReductionPollView) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CapitalReductionPollView) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetRefusalVotesPercentage

`func (o *CapitalReductionPollView) GetRefusalVotesPercentage() float32`

GetRefusalVotesPercentage returns the RefusalVotesPercentage field if non-nil, zero value otherwise.

### GetRefusalVotesPercentageOk

`func (o *CapitalReductionPollView) GetRefusalVotesPercentageOk() (*float32, bool)`

GetRefusalVotesPercentageOk returns a tuple with the RefusalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalVotesPercentage

`func (o *CapitalReductionPollView) SetRefusalVotesPercentage(v float32)`

SetRefusalVotesPercentage sets RefusalVotesPercentage field to given value.

### HasRefusalVotesPercentage

`func (o *CapitalReductionPollView) HasRefusalVotesPercentage() bool`

HasRefusalVotesPercentage returns a boolean if a field has been set.

### GetResultExpireDate

`func (o *CapitalReductionPollView) GetResultExpireDate() int64`

GetResultExpireDate returns the ResultExpireDate field if non-nil, zero value otherwise.

### GetResultExpireDateOk

`func (o *CapitalReductionPollView) GetResultExpireDateOk() (*int64, bool)`

GetResultExpireDateOk returns a tuple with the ResultExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultExpireDate

`func (o *CapitalReductionPollView) SetResultExpireDate(v int64)`

SetResultExpireDate sets ResultExpireDate field to given value.

### HasResultExpireDate

`func (o *CapitalReductionPollView) HasResultExpireDate() bool`

HasResultExpireDate returns a boolean if a field has been set.

### GetStartDate

`func (o *CapitalReductionPollView) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CapitalReductionPollView) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CapitalReductionPollView) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CapitalReductionPollView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTotalNumberOfCastVotes

`func (o *CapitalReductionPollView) GetTotalNumberOfCastVotes() int64`

GetTotalNumberOfCastVotes returns the TotalNumberOfCastVotes field if non-nil, zero value otherwise.

### GetTotalNumberOfCastVotesOk

`func (o *CapitalReductionPollView) GetTotalNumberOfCastVotesOk() (*int64, bool)`

GetTotalNumberOfCastVotesOk returns a tuple with the TotalNumberOfCastVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfCastVotes

`func (o *CapitalReductionPollView) SetTotalNumberOfCastVotes(v int64)`

SetTotalNumberOfCastVotes sets TotalNumberOfCastVotes field to given value.

### HasTotalNumberOfCastVotes

`func (o *CapitalReductionPollView) HasTotalNumberOfCastVotes() bool`

HasTotalNumberOfCastVotes returns a boolean if a field has been set.

### GetTotalNumberOfVoices

`func (o *CapitalReductionPollView) GetTotalNumberOfVoices() int64`

GetTotalNumberOfVoices returns the TotalNumberOfVoices field if non-nil, zero value otherwise.

### GetTotalNumberOfVoicesOk

`func (o *CapitalReductionPollView) GetTotalNumberOfVoicesOk() (*int64, bool)`

GetTotalNumberOfVoicesOk returns a tuple with the TotalNumberOfVoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfVoices

`func (o *CapitalReductionPollView) SetTotalNumberOfVoices(v int64)`

SetTotalNumberOfVoices sets TotalNumberOfVoices field to given value.

### HasTotalNumberOfVoices

`func (o *CapitalReductionPollView) HasTotalNumberOfVoices() bool`

HasTotalNumberOfVoices returns a boolean if a field has been set.

### GetType

`func (o *CapitalReductionPollView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CapitalReductionPollView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CapitalReductionPollView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CapitalReductionPollView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *CapitalReductionPollView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CapitalReductionPollView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CapitalReductionPollView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CapitalReductionPollView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVotes

`func (o *CapitalReductionPollView) GetVotes() []VoteView`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *CapitalReductionPollView) GetVotesOk() (*[]VoteView, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *CapitalReductionPollView) SetVotes(v []VoteView)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *CapitalReductionPollView) HasVotes() bool`

HasVotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


