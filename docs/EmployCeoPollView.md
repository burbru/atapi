# EmployCeoPollView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbstentionRule** | Pointer to **string** |  | [optional] 
**Applicant** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**ApprovalVotesPercentage** | Pointer to **float32** |  | [optional] 
**CastApprovalVotesPercentage** | Pointer to **float32** |  | [optional] 
**CastRefusalVotesPercentage** | Pointer to **float32** |  | [optional] 
**CastVotesPercentage** | Pointer to **float32** |  | [optional] 
**Company** | Pointer to [**CompactCompanyView**](CompactCompanyView.md) |  | [optional] 
**DailyWage** | Pointer to **float32** |  | [optional] 
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

### NewEmployCeoPollView

`func NewEmployCeoPollView() *EmployCeoPollView`

NewEmployCeoPollView instantiates a new EmployCeoPollView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployCeoPollViewWithDefaults

`func NewEmployCeoPollViewWithDefaults() *EmployCeoPollView`

NewEmployCeoPollViewWithDefaults instantiates a new EmployCeoPollView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbstentionRule

`func (o *EmployCeoPollView) GetAbstentionRule() string`

GetAbstentionRule returns the AbstentionRule field if non-nil, zero value otherwise.

### GetAbstentionRuleOk

`func (o *EmployCeoPollView) GetAbstentionRuleOk() (*string, bool)`

GetAbstentionRuleOk returns a tuple with the AbstentionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstentionRule

`func (o *EmployCeoPollView) SetAbstentionRule(v string)`

SetAbstentionRule sets AbstentionRule field to given value.

### HasAbstentionRule

`func (o *EmployCeoPollView) HasAbstentionRule() bool`

HasAbstentionRule returns a boolean if a field has been set.

### GetApplicant

`func (o *EmployCeoPollView) GetApplicant() UsernameView`

GetApplicant returns the Applicant field if non-nil, zero value otherwise.

### GetApplicantOk

`func (o *EmployCeoPollView) GetApplicantOk() (*UsernameView, bool)`

GetApplicantOk returns a tuple with the Applicant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicant

`func (o *EmployCeoPollView) SetApplicant(v UsernameView)`

SetApplicant sets Applicant field to given value.

### HasApplicant

`func (o *EmployCeoPollView) HasApplicant() bool`

HasApplicant returns a boolean if a field has been set.

### GetApprovalVotesPercentage

`func (o *EmployCeoPollView) GetApprovalVotesPercentage() float32`

GetApprovalVotesPercentage returns the ApprovalVotesPercentage field if non-nil, zero value otherwise.

### GetApprovalVotesPercentageOk

`func (o *EmployCeoPollView) GetApprovalVotesPercentageOk() (*float32, bool)`

GetApprovalVotesPercentageOk returns a tuple with the ApprovalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalVotesPercentage

`func (o *EmployCeoPollView) SetApprovalVotesPercentage(v float32)`

SetApprovalVotesPercentage sets ApprovalVotesPercentage field to given value.

### HasApprovalVotesPercentage

`func (o *EmployCeoPollView) HasApprovalVotesPercentage() bool`

HasApprovalVotesPercentage returns a boolean if a field has been set.

### GetCastApprovalVotesPercentage

`func (o *EmployCeoPollView) GetCastApprovalVotesPercentage() float32`

GetCastApprovalVotesPercentage returns the CastApprovalVotesPercentage field if non-nil, zero value otherwise.

### GetCastApprovalVotesPercentageOk

`func (o *EmployCeoPollView) GetCastApprovalVotesPercentageOk() (*float32, bool)`

GetCastApprovalVotesPercentageOk returns a tuple with the CastApprovalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastApprovalVotesPercentage

`func (o *EmployCeoPollView) SetCastApprovalVotesPercentage(v float32)`

SetCastApprovalVotesPercentage sets CastApprovalVotesPercentage field to given value.

### HasCastApprovalVotesPercentage

`func (o *EmployCeoPollView) HasCastApprovalVotesPercentage() bool`

HasCastApprovalVotesPercentage returns a boolean if a field has been set.

### GetCastRefusalVotesPercentage

`func (o *EmployCeoPollView) GetCastRefusalVotesPercentage() float32`

GetCastRefusalVotesPercentage returns the CastRefusalVotesPercentage field if non-nil, zero value otherwise.

### GetCastRefusalVotesPercentageOk

`func (o *EmployCeoPollView) GetCastRefusalVotesPercentageOk() (*float32, bool)`

GetCastRefusalVotesPercentageOk returns a tuple with the CastRefusalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastRefusalVotesPercentage

`func (o *EmployCeoPollView) SetCastRefusalVotesPercentage(v float32)`

SetCastRefusalVotesPercentage sets CastRefusalVotesPercentage field to given value.

### HasCastRefusalVotesPercentage

`func (o *EmployCeoPollView) HasCastRefusalVotesPercentage() bool`

HasCastRefusalVotesPercentage returns a boolean if a field has been set.

### GetCastVotesPercentage

`func (o *EmployCeoPollView) GetCastVotesPercentage() float32`

GetCastVotesPercentage returns the CastVotesPercentage field if non-nil, zero value otherwise.

### GetCastVotesPercentageOk

`func (o *EmployCeoPollView) GetCastVotesPercentageOk() (*float32, bool)`

GetCastVotesPercentageOk returns a tuple with the CastVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastVotesPercentage

`func (o *EmployCeoPollView) SetCastVotesPercentage(v float32)`

SetCastVotesPercentage sets CastVotesPercentage field to given value.

### HasCastVotesPercentage

`func (o *EmployCeoPollView) HasCastVotesPercentage() bool`

HasCastVotesPercentage returns a boolean if a field has been set.

### GetCompany

`func (o *EmployCeoPollView) GetCompany() CompactCompanyView`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *EmployCeoPollView) GetCompanyOk() (*CompactCompanyView, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *EmployCeoPollView) SetCompany(v CompactCompanyView)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *EmployCeoPollView) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetDailyWage

`func (o *EmployCeoPollView) GetDailyWage() float32`

GetDailyWage returns the DailyWage field if non-nil, zero value otherwise.

### GetDailyWageOk

`func (o *EmployCeoPollView) GetDailyWageOk() (*float32, bool)`

GetDailyWageOk returns a tuple with the DailyWage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyWage

`func (o *EmployCeoPollView) SetDailyWage(v float32)`

SetDailyWage sets DailyWage field to given value.

### HasDailyWage

`func (o *EmployCeoPollView) HasDailyWage() bool`

HasDailyWage returns a boolean if a field has been set.

### GetEndDate

`func (o *EmployCeoPollView) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *EmployCeoPollView) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *EmployCeoPollView) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *EmployCeoPollView) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetGroup

`func (o *EmployCeoPollView) GetGroup() []VoiceNumberView`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *EmployCeoPollView) GetGroupOk() (*[]VoiceNumberView, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *EmployCeoPollView) SetGroup(v []VoiceNumberView)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *EmployCeoPollView) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *EmployCeoPollView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployCeoPollView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployCeoPollView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmployCeoPollView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLeftVotesPercentage

`func (o *EmployCeoPollView) GetLeftVotesPercentage() float32`

GetLeftVotesPercentage returns the LeftVotesPercentage field if non-nil, zero value otherwise.

### GetLeftVotesPercentageOk

`func (o *EmployCeoPollView) GetLeftVotesPercentageOk() (*float32, bool)`

GetLeftVotesPercentageOk returns a tuple with the LeftVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftVotesPercentage

`func (o *EmployCeoPollView) SetLeftVotesPercentage(v float32)`

SetLeftVotesPercentage sets LeftVotesPercentage field to given value.

### HasLeftVotesPercentage

`func (o *EmployCeoPollView) HasLeftVotesPercentage() bool`

HasLeftVotesPercentage returns a boolean if a field has been set.

### GetMotion

`func (o *EmployCeoPollView) GetMotion() string`

GetMotion returns the Motion field if non-nil, zero value otherwise.

### GetMotionOk

`func (o *EmployCeoPollView) GetMotionOk() (*string, bool)`

GetMotionOk returns a tuple with the Motion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotion

`func (o *EmployCeoPollView) SetMotion(v string)`

SetMotion sets Motion field to given value.

### HasMotion

`func (o *EmployCeoPollView) HasMotion() bool`

HasMotion returns a boolean if a field has been set.

### GetPollInitiator

`func (o *EmployCeoPollView) GetPollInitiator() UsernameView`

GetPollInitiator returns the PollInitiator field if non-nil, zero value otherwise.

### GetPollInitiatorOk

`func (o *EmployCeoPollView) GetPollInitiatorOk() (*UsernameView, bool)`

GetPollInitiatorOk returns a tuple with the PollInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollInitiator

`func (o *EmployCeoPollView) SetPollInitiator(v UsernameView)`

SetPollInitiator sets PollInitiator field to given value.

### HasPollInitiator

`func (o *EmployCeoPollView) HasPollInitiator() bool`

HasPollInitiator returns a boolean if a field has been set.

### GetRefusalVotesPercentage

`func (o *EmployCeoPollView) GetRefusalVotesPercentage() float32`

GetRefusalVotesPercentage returns the RefusalVotesPercentage field if non-nil, zero value otherwise.

### GetRefusalVotesPercentageOk

`func (o *EmployCeoPollView) GetRefusalVotesPercentageOk() (*float32, bool)`

GetRefusalVotesPercentageOk returns a tuple with the RefusalVotesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalVotesPercentage

`func (o *EmployCeoPollView) SetRefusalVotesPercentage(v float32)`

SetRefusalVotesPercentage sets RefusalVotesPercentage field to given value.

### HasRefusalVotesPercentage

`func (o *EmployCeoPollView) HasRefusalVotesPercentage() bool`

HasRefusalVotesPercentage returns a boolean if a field has been set.

### GetResultExpireDate

`func (o *EmployCeoPollView) GetResultExpireDate() int64`

GetResultExpireDate returns the ResultExpireDate field if non-nil, zero value otherwise.

### GetResultExpireDateOk

`func (o *EmployCeoPollView) GetResultExpireDateOk() (*int64, bool)`

GetResultExpireDateOk returns a tuple with the ResultExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultExpireDate

`func (o *EmployCeoPollView) SetResultExpireDate(v int64)`

SetResultExpireDate sets ResultExpireDate field to given value.

### HasResultExpireDate

`func (o *EmployCeoPollView) HasResultExpireDate() bool`

HasResultExpireDate returns a boolean if a field has been set.

### GetStartDate

`func (o *EmployCeoPollView) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EmployCeoPollView) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EmployCeoPollView) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *EmployCeoPollView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTotalNumberOfCastVotes

`func (o *EmployCeoPollView) GetTotalNumberOfCastVotes() int64`

GetTotalNumberOfCastVotes returns the TotalNumberOfCastVotes field if non-nil, zero value otherwise.

### GetTotalNumberOfCastVotesOk

`func (o *EmployCeoPollView) GetTotalNumberOfCastVotesOk() (*int64, bool)`

GetTotalNumberOfCastVotesOk returns a tuple with the TotalNumberOfCastVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfCastVotes

`func (o *EmployCeoPollView) SetTotalNumberOfCastVotes(v int64)`

SetTotalNumberOfCastVotes sets TotalNumberOfCastVotes field to given value.

### HasTotalNumberOfCastVotes

`func (o *EmployCeoPollView) HasTotalNumberOfCastVotes() bool`

HasTotalNumberOfCastVotes returns a boolean if a field has been set.

### GetTotalNumberOfVoices

`func (o *EmployCeoPollView) GetTotalNumberOfVoices() int64`

GetTotalNumberOfVoices returns the TotalNumberOfVoices field if non-nil, zero value otherwise.

### GetTotalNumberOfVoicesOk

`func (o *EmployCeoPollView) GetTotalNumberOfVoicesOk() (*int64, bool)`

GetTotalNumberOfVoicesOk returns a tuple with the TotalNumberOfVoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfVoices

`func (o *EmployCeoPollView) SetTotalNumberOfVoices(v int64)`

SetTotalNumberOfVoices sets TotalNumberOfVoices field to given value.

### HasTotalNumberOfVoices

`func (o *EmployCeoPollView) HasTotalNumberOfVoices() bool`

HasTotalNumberOfVoices returns a boolean if a field has been set.

### GetType

`func (o *EmployCeoPollView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmployCeoPollView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmployCeoPollView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EmployCeoPollView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *EmployCeoPollView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EmployCeoPollView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EmployCeoPollView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EmployCeoPollView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVotes

`func (o *EmployCeoPollView) GetVotes() []VoteView`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *EmployCeoPollView) GetVotesOk() (*[]VoteView, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *EmployCeoPollView) SetVotes(v []VoteView)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *EmployCeoPollView) HasVotes() bool`

HasVotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


