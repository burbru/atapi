/*
Api Documentation

Api Documentation

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package atapi

import (
	"encoding/json"
)

// checks if the CapitalIncreasePollView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapitalIncreasePollView{}

// CapitalIncreasePollView struct for CapitalIncreasePollView
type CapitalIncreasePollView struct {
	AbstentionRule *string `json:"abstentionRule,omitempty"`
	ApprovalVotesPercentage *float32 `json:"approvalVotesPercentage,omitempty"`
	CapitalIncreaseEndDate *int64 `json:"capitalIncreaseEndDate,omitempty"`
	CapitalIncreaseStartDate *int64 `json:"capitalIncreaseStartDate,omitempty"`
	CapitalIncreaseSubscriptionPeriodEndDate *int64 `json:"capitalIncreaseSubscriptionPeriodEndDate,omitempty"`
	CapitalIncreaseType *string `json:"capitalIncreaseType,omitempty"`
	CastApprovalVotesPercentage *float32 `json:"castApprovalVotesPercentage,omitempty"`
	CastRefusalVotesPercentage *float32 `json:"castRefusalVotesPercentage,omitempty"`
	CastVotesPercentage *float32 `json:"castVotesPercentage,omitempty"`
	Company *CompactCompanyView `json:"company,omitempty"`
	EndDate *int64 `json:"endDate,omitempty"`
	Group []VoiceNumberView `json:"group,omitempty"`
	Id *string `json:"id,omitempty"`
	LeftVotesPercentage *float32 `json:"leftVotesPercentage,omitempty"`
	MinimalCashVolume *float32 `json:"minimalCashVolume,omitempty"`
	Motion *string `json:"motion,omitempty"`
	NumberOfShares *int64 `json:"numberOfShares,omitempty"`
	PollInitiator *UsernameView `json:"pollInitiator,omitempty"`
	RefusalVotesPercentage *float32 `json:"refusalVotesPercentage,omitempty"`
	ResultExpireDate *int64 `json:"resultExpireDate,omitempty"`
	StartDate *int64 `json:"startDate,omitempty"`
	SubscriptionFraction *int64 `json:"subscriptionFraction,omitempty"`
	TotalNumberOfCastVotes *int64 `json:"totalNumberOfCastVotes,omitempty"`
	TotalNumberOfVoices *int64 `json:"totalNumberOfVoices,omitempty"`
	Type *string `json:"type,omitempty"`
	Version *int64 `json:"version,omitempty"`
	Votes []VoteView `json:"votes,omitempty"`
}

// NewCapitalIncreasePollView instantiates a new CapitalIncreasePollView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapitalIncreasePollView() *CapitalIncreasePollView {
	this := CapitalIncreasePollView{}
	return &this
}

// NewCapitalIncreasePollViewWithDefaults instantiates a new CapitalIncreasePollView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapitalIncreasePollViewWithDefaults() *CapitalIncreasePollView {
	this := CapitalIncreasePollView{}
	return &this
}

// GetAbstentionRule returns the AbstentionRule field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetAbstentionRule() string {
	if o == nil || IsNil(o.AbstentionRule) {
		var ret string
		return ret
	}
	return *o.AbstentionRule
}

// GetAbstentionRuleOk returns a tuple with the AbstentionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetAbstentionRuleOk() (*string, bool) {
	if o == nil || IsNil(o.AbstentionRule) {
		return nil, false
	}
	return o.AbstentionRule, true
}

// HasAbstentionRule returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasAbstentionRule() bool {
	if o != nil && !IsNil(o.AbstentionRule) {
		return true
	}

	return false
}

// SetAbstentionRule gets a reference to the given string and assigns it to the AbstentionRule field.
func (o *CapitalIncreasePollView) SetAbstentionRule(v string) {
	o.AbstentionRule = &v
}

// GetApprovalVotesPercentage returns the ApprovalVotesPercentage field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetApprovalVotesPercentage() float32 {
	if o == nil || IsNil(o.ApprovalVotesPercentage) {
		var ret float32
		return ret
	}
	return *o.ApprovalVotesPercentage
}

// GetApprovalVotesPercentageOk returns a tuple with the ApprovalVotesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetApprovalVotesPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.ApprovalVotesPercentage) {
		return nil, false
	}
	return o.ApprovalVotesPercentage, true
}

// HasApprovalVotesPercentage returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasApprovalVotesPercentage() bool {
	if o != nil && !IsNil(o.ApprovalVotesPercentage) {
		return true
	}

	return false
}

// SetApprovalVotesPercentage gets a reference to the given float32 and assigns it to the ApprovalVotesPercentage field.
func (o *CapitalIncreasePollView) SetApprovalVotesPercentage(v float32) {
	o.ApprovalVotesPercentage = &v
}

// GetCapitalIncreaseEndDate returns the CapitalIncreaseEndDate field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetCapitalIncreaseEndDate() int64 {
	if o == nil || IsNil(o.CapitalIncreaseEndDate) {
		var ret int64
		return ret
	}
	return *o.CapitalIncreaseEndDate
}

// GetCapitalIncreaseEndDateOk returns a tuple with the CapitalIncreaseEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetCapitalIncreaseEndDateOk() (*int64, bool) {
	if o == nil || IsNil(o.CapitalIncreaseEndDate) {
		return nil, false
	}
	return o.CapitalIncreaseEndDate, true
}

// HasCapitalIncreaseEndDate returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasCapitalIncreaseEndDate() bool {
	if o != nil && !IsNil(o.CapitalIncreaseEndDate) {
		return true
	}

	return false
}

// SetCapitalIncreaseEndDate gets a reference to the given int64 and assigns it to the CapitalIncreaseEndDate field.
func (o *CapitalIncreasePollView) SetCapitalIncreaseEndDate(v int64) {
	o.CapitalIncreaseEndDate = &v
}

// GetCapitalIncreaseStartDate returns the CapitalIncreaseStartDate field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetCapitalIncreaseStartDate() int64 {
	if o == nil || IsNil(o.CapitalIncreaseStartDate) {
		var ret int64
		return ret
	}
	return *o.CapitalIncreaseStartDate
}

// GetCapitalIncreaseStartDateOk returns a tuple with the CapitalIncreaseStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetCapitalIncreaseStartDateOk() (*int64, bool) {
	if o == nil || IsNil(o.CapitalIncreaseStartDate) {
		return nil, false
	}
	return o.CapitalIncreaseStartDate, true
}

// HasCapitalIncreaseStartDate returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasCapitalIncreaseStartDate() bool {
	if o != nil && !IsNil(o.CapitalIncreaseStartDate) {
		return true
	}

	return false
}

// SetCapitalIncreaseStartDate gets a reference to the given int64 and assigns it to the CapitalIncreaseStartDate field.
func (o *CapitalIncreasePollView) SetCapitalIncreaseStartDate(v int64) {
	o.CapitalIncreaseStartDate = &v
}

// GetCapitalIncreaseSubscriptionPeriodEndDate returns the CapitalIncreaseSubscriptionPeriodEndDate field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetCapitalIncreaseSubscriptionPeriodEndDate() int64 {
	if o == nil || IsNil(o.CapitalIncreaseSubscriptionPeriodEndDate) {
		var ret int64
		return ret
	}
	return *o.CapitalIncreaseSubscriptionPeriodEndDate
}

// GetCapitalIncreaseSubscriptionPeriodEndDateOk returns a tuple with the CapitalIncreaseSubscriptionPeriodEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetCapitalIncreaseSubscriptionPeriodEndDateOk() (*int64, bool) {
	if o == nil || IsNil(o.CapitalIncreaseSubscriptionPeriodEndDate) {
		return nil, false
	}
	return o.CapitalIncreaseSubscriptionPeriodEndDate, true
}

// HasCapitalIncreaseSubscriptionPeriodEndDate returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasCapitalIncreaseSubscriptionPeriodEndDate() bool {
	if o != nil && !IsNil(o.CapitalIncreaseSubscriptionPeriodEndDate) {
		return true
	}

	return false
}

// SetCapitalIncreaseSubscriptionPeriodEndDate gets a reference to the given int64 and assigns it to the CapitalIncreaseSubscriptionPeriodEndDate field.
func (o *CapitalIncreasePollView) SetCapitalIncreaseSubscriptionPeriodEndDate(v int64) {
	o.CapitalIncreaseSubscriptionPeriodEndDate = &v
}

// GetCapitalIncreaseType returns the CapitalIncreaseType field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetCapitalIncreaseType() string {
	if o == nil || IsNil(o.CapitalIncreaseType) {
		var ret string
		return ret
	}
	return *o.CapitalIncreaseType
}

// GetCapitalIncreaseTypeOk returns a tuple with the CapitalIncreaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetCapitalIncreaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CapitalIncreaseType) {
		return nil, false
	}
	return o.CapitalIncreaseType, true
}

// HasCapitalIncreaseType returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasCapitalIncreaseType() bool {
	if o != nil && !IsNil(o.CapitalIncreaseType) {
		return true
	}

	return false
}

// SetCapitalIncreaseType gets a reference to the given string and assigns it to the CapitalIncreaseType field.
func (o *CapitalIncreasePollView) SetCapitalIncreaseType(v string) {
	o.CapitalIncreaseType = &v
}

// GetCastApprovalVotesPercentage returns the CastApprovalVotesPercentage field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetCastApprovalVotesPercentage() float32 {
	if o == nil || IsNil(o.CastApprovalVotesPercentage) {
		var ret float32
		return ret
	}
	return *o.CastApprovalVotesPercentage
}

// GetCastApprovalVotesPercentageOk returns a tuple with the CastApprovalVotesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetCastApprovalVotesPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.CastApprovalVotesPercentage) {
		return nil, false
	}
	return o.CastApprovalVotesPercentage, true
}

// HasCastApprovalVotesPercentage returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasCastApprovalVotesPercentage() bool {
	if o != nil && !IsNil(o.CastApprovalVotesPercentage) {
		return true
	}

	return false
}

// SetCastApprovalVotesPercentage gets a reference to the given float32 and assigns it to the CastApprovalVotesPercentage field.
func (o *CapitalIncreasePollView) SetCastApprovalVotesPercentage(v float32) {
	o.CastApprovalVotesPercentage = &v
}

// GetCastRefusalVotesPercentage returns the CastRefusalVotesPercentage field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetCastRefusalVotesPercentage() float32 {
	if o == nil || IsNil(o.CastRefusalVotesPercentage) {
		var ret float32
		return ret
	}
	return *o.CastRefusalVotesPercentage
}

// GetCastRefusalVotesPercentageOk returns a tuple with the CastRefusalVotesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetCastRefusalVotesPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.CastRefusalVotesPercentage) {
		return nil, false
	}
	return o.CastRefusalVotesPercentage, true
}

// HasCastRefusalVotesPercentage returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasCastRefusalVotesPercentage() bool {
	if o != nil && !IsNil(o.CastRefusalVotesPercentage) {
		return true
	}

	return false
}

// SetCastRefusalVotesPercentage gets a reference to the given float32 and assigns it to the CastRefusalVotesPercentage field.
func (o *CapitalIncreasePollView) SetCastRefusalVotesPercentage(v float32) {
	o.CastRefusalVotesPercentage = &v
}

// GetCastVotesPercentage returns the CastVotesPercentage field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetCastVotesPercentage() float32 {
	if o == nil || IsNil(o.CastVotesPercentage) {
		var ret float32
		return ret
	}
	return *o.CastVotesPercentage
}

// GetCastVotesPercentageOk returns a tuple with the CastVotesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetCastVotesPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.CastVotesPercentage) {
		return nil, false
	}
	return o.CastVotesPercentage, true
}

// HasCastVotesPercentage returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasCastVotesPercentage() bool {
	if o != nil && !IsNil(o.CastVotesPercentage) {
		return true
	}

	return false
}

// SetCastVotesPercentage gets a reference to the given float32 and assigns it to the CastVotesPercentage field.
func (o *CapitalIncreasePollView) SetCastVotesPercentage(v float32) {
	o.CastVotesPercentage = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetCompany() CompactCompanyView {
	if o == nil || IsNil(o.Company) {
		var ret CompactCompanyView
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetCompanyOk() (*CompactCompanyView, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given CompactCompanyView and assigns it to the Company field.
func (o *CapitalIncreasePollView) SetCompany(v CompactCompanyView) {
	o.Company = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetEndDate() int64 {
	if o == nil || IsNil(o.EndDate) {
		var ret int64
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetEndDateOk() (*int64, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int64 and assigns it to the EndDate field.
func (o *CapitalIncreasePollView) SetEndDate(v int64) {
	o.EndDate = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetGroup() []VoiceNumberView {
	if o == nil || IsNil(o.Group) {
		var ret []VoiceNumberView
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetGroupOk() ([]VoiceNumberView, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given []VoiceNumberView and assigns it to the Group field.
func (o *CapitalIncreasePollView) SetGroup(v []VoiceNumberView) {
	o.Group = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CapitalIncreasePollView) SetId(v string) {
	o.Id = &v
}

// GetLeftVotesPercentage returns the LeftVotesPercentage field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetLeftVotesPercentage() float32 {
	if o == nil || IsNil(o.LeftVotesPercentage) {
		var ret float32
		return ret
	}
	return *o.LeftVotesPercentage
}

// GetLeftVotesPercentageOk returns a tuple with the LeftVotesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetLeftVotesPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.LeftVotesPercentage) {
		return nil, false
	}
	return o.LeftVotesPercentage, true
}

// HasLeftVotesPercentage returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasLeftVotesPercentage() bool {
	if o != nil && !IsNil(o.LeftVotesPercentage) {
		return true
	}

	return false
}

// SetLeftVotesPercentage gets a reference to the given float32 and assigns it to the LeftVotesPercentage field.
func (o *CapitalIncreasePollView) SetLeftVotesPercentage(v float32) {
	o.LeftVotesPercentage = &v
}

// GetMinimalCashVolume returns the MinimalCashVolume field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetMinimalCashVolume() float32 {
	if o == nil || IsNil(o.MinimalCashVolume) {
		var ret float32
		return ret
	}
	return *o.MinimalCashVolume
}

// GetMinimalCashVolumeOk returns a tuple with the MinimalCashVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetMinimalCashVolumeOk() (*float32, bool) {
	if o == nil || IsNil(o.MinimalCashVolume) {
		return nil, false
	}
	return o.MinimalCashVolume, true
}

// HasMinimalCashVolume returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasMinimalCashVolume() bool {
	if o != nil && !IsNil(o.MinimalCashVolume) {
		return true
	}

	return false
}

// SetMinimalCashVolume gets a reference to the given float32 and assigns it to the MinimalCashVolume field.
func (o *CapitalIncreasePollView) SetMinimalCashVolume(v float32) {
	o.MinimalCashVolume = &v
}

// GetMotion returns the Motion field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetMotion() string {
	if o == nil || IsNil(o.Motion) {
		var ret string
		return ret
	}
	return *o.Motion
}

// GetMotionOk returns a tuple with the Motion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetMotionOk() (*string, bool) {
	if o == nil || IsNil(o.Motion) {
		return nil, false
	}
	return o.Motion, true
}

// HasMotion returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasMotion() bool {
	if o != nil && !IsNil(o.Motion) {
		return true
	}

	return false
}

// SetMotion gets a reference to the given string and assigns it to the Motion field.
func (o *CapitalIncreasePollView) SetMotion(v string) {
	o.Motion = &v
}

// GetNumberOfShares returns the NumberOfShares field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetNumberOfShares() int64 {
	if o == nil || IsNil(o.NumberOfShares) {
		var ret int64
		return ret
	}
	return *o.NumberOfShares
}

// GetNumberOfSharesOk returns a tuple with the NumberOfShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetNumberOfSharesOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfShares) {
		return nil, false
	}
	return o.NumberOfShares, true
}

// HasNumberOfShares returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasNumberOfShares() bool {
	if o != nil && !IsNil(o.NumberOfShares) {
		return true
	}

	return false
}

// SetNumberOfShares gets a reference to the given int64 and assigns it to the NumberOfShares field.
func (o *CapitalIncreasePollView) SetNumberOfShares(v int64) {
	o.NumberOfShares = &v
}

// GetPollInitiator returns the PollInitiator field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetPollInitiator() UsernameView {
	if o == nil || IsNil(o.PollInitiator) {
		var ret UsernameView
		return ret
	}
	return *o.PollInitiator
}

// GetPollInitiatorOk returns a tuple with the PollInitiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetPollInitiatorOk() (*UsernameView, bool) {
	if o == nil || IsNil(o.PollInitiator) {
		return nil, false
	}
	return o.PollInitiator, true
}

// HasPollInitiator returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasPollInitiator() bool {
	if o != nil && !IsNil(o.PollInitiator) {
		return true
	}

	return false
}

// SetPollInitiator gets a reference to the given UsernameView and assigns it to the PollInitiator field.
func (o *CapitalIncreasePollView) SetPollInitiator(v UsernameView) {
	o.PollInitiator = &v
}

// GetRefusalVotesPercentage returns the RefusalVotesPercentage field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetRefusalVotesPercentage() float32 {
	if o == nil || IsNil(o.RefusalVotesPercentage) {
		var ret float32
		return ret
	}
	return *o.RefusalVotesPercentage
}

// GetRefusalVotesPercentageOk returns a tuple with the RefusalVotesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetRefusalVotesPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.RefusalVotesPercentage) {
		return nil, false
	}
	return o.RefusalVotesPercentage, true
}

// HasRefusalVotesPercentage returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasRefusalVotesPercentage() bool {
	if o != nil && !IsNil(o.RefusalVotesPercentage) {
		return true
	}

	return false
}

// SetRefusalVotesPercentage gets a reference to the given float32 and assigns it to the RefusalVotesPercentage field.
func (o *CapitalIncreasePollView) SetRefusalVotesPercentage(v float32) {
	o.RefusalVotesPercentage = &v
}

// GetResultExpireDate returns the ResultExpireDate field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetResultExpireDate() int64 {
	if o == nil || IsNil(o.ResultExpireDate) {
		var ret int64
		return ret
	}
	return *o.ResultExpireDate
}

// GetResultExpireDateOk returns a tuple with the ResultExpireDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetResultExpireDateOk() (*int64, bool) {
	if o == nil || IsNil(o.ResultExpireDate) {
		return nil, false
	}
	return o.ResultExpireDate, true
}

// HasResultExpireDate returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasResultExpireDate() bool {
	if o != nil && !IsNil(o.ResultExpireDate) {
		return true
	}

	return false
}

// SetResultExpireDate gets a reference to the given int64 and assigns it to the ResultExpireDate field.
func (o *CapitalIncreasePollView) SetResultExpireDate(v int64) {
	o.ResultExpireDate = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetStartDate() int64 {
	if o == nil || IsNil(o.StartDate) {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetStartDateOk() (*int64, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *CapitalIncreasePollView) SetStartDate(v int64) {
	o.StartDate = &v
}

// GetSubscriptionFraction returns the SubscriptionFraction field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetSubscriptionFraction() int64 {
	if o == nil || IsNil(o.SubscriptionFraction) {
		var ret int64
		return ret
	}
	return *o.SubscriptionFraction
}

// GetSubscriptionFractionOk returns a tuple with the SubscriptionFraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetSubscriptionFractionOk() (*int64, bool) {
	if o == nil || IsNil(o.SubscriptionFraction) {
		return nil, false
	}
	return o.SubscriptionFraction, true
}

// HasSubscriptionFraction returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasSubscriptionFraction() bool {
	if o != nil && !IsNil(o.SubscriptionFraction) {
		return true
	}

	return false
}

// SetSubscriptionFraction gets a reference to the given int64 and assigns it to the SubscriptionFraction field.
func (o *CapitalIncreasePollView) SetSubscriptionFraction(v int64) {
	o.SubscriptionFraction = &v
}

// GetTotalNumberOfCastVotes returns the TotalNumberOfCastVotes field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetTotalNumberOfCastVotes() int64 {
	if o == nil || IsNil(o.TotalNumberOfCastVotes) {
		var ret int64
		return ret
	}
	return *o.TotalNumberOfCastVotes
}

// GetTotalNumberOfCastVotesOk returns a tuple with the TotalNumberOfCastVotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetTotalNumberOfCastVotesOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalNumberOfCastVotes) {
		return nil, false
	}
	return o.TotalNumberOfCastVotes, true
}

// HasTotalNumberOfCastVotes returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasTotalNumberOfCastVotes() bool {
	if o != nil && !IsNil(o.TotalNumberOfCastVotes) {
		return true
	}

	return false
}

// SetTotalNumberOfCastVotes gets a reference to the given int64 and assigns it to the TotalNumberOfCastVotes field.
func (o *CapitalIncreasePollView) SetTotalNumberOfCastVotes(v int64) {
	o.TotalNumberOfCastVotes = &v
}

// GetTotalNumberOfVoices returns the TotalNumberOfVoices field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetTotalNumberOfVoices() int64 {
	if o == nil || IsNil(o.TotalNumberOfVoices) {
		var ret int64
		return ret
	}
	return *o.TotalNumberOfVoices
}

// GetTotalNumberOfVoicesOk returns a tuple with the TotalNumberOfVoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetTotalNumberOfVoicesOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalNumberOfVoices) {
		return nil, false
	}
	return o.TotalNumberOfVoices, true
}

// HasTotalNumberOfVoices returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasTotalNumberOfVoices() bool {
	if o != nil && !IsNil(o.TotalNumberOfVoices) {
		return true
	}

	return false
}

// SetTotalNumberOfVoices gets a reference to the given int64 and assigns it to the TotalNumberOfVoices field.
func (o *CapitalIncreasePollView) SetTotalNumberOfVoices(v int64) {
	o.TotalNumberOfVoices = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CapitalIncreasePollView) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *CapitalIncreasePollView) SetVersion(v int64) {
	o.Version = &v
}

// GetVotes returns the Votes field value if set, zero value otherwise.
func (o *CapitalIncreasePollView) GetVotes() []VoteView {
	if o == nil || IsNil(o.Votes) {
		var ret []VoteView
		return ret
	}
	return o.Votes
}

// GetVotesOk returns a tuple with the Votes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalIncreasePollView) GetVotesOk() ([]VoteView, bool) {
	if o == nil || IsNil(o.Votes) {
		return nil, false
	}
	return o.Votes, true
}

// HasVotes returns a boolean if a field has been set.
func (o *CapitalIncreasePollView) HasVotes() bool {
	if o != nil && !IsNil(o.Votes) {
		return true
	}

	return false
}

// SetVotes gets a reference to the given []VoteView and assigns it to the Votes field.
func (o *CapitalIncreasePollView) SetVotes(v []VoteView) {
	o.Votes = v
}

func (o CapitalIncreasePollView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapitalIncreasePollView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AbstentionRule) {
		toSerialize["abstentionRule"] = o.AbstentionRule
	}
	if !IsNil(o.ApprovalVotesPercentage) {
		toSerialize["approvalVotesPercentage"] = o.ApprovalVotesPercentage
	}
	if !IsNil(o.CapitalIncreaseEndDate) {
		toSerialize["capitalIncreaseEndDate"] = o.CapitalIncreaseEndDate
	}
	if !IsNil(o.CapitalIncreaseStartDate) {
		toSerialize["capitalIncreaseStartDate"] = o.CapitalIncreaseStartDate
	}
	if !IsNil(o.CapitalIncreaseSubscriptionPeriodEndDate) {
		toSerialize["capitalIncreaseSubscriptionPeriodEndDate"] = o.CapitalIncreaseSubscriptionPeriodEndDate
	}
	if !IsNil(o.CapitalIncreaseType) {
		toSerialize["capitalIncreaseType"] = o.CapitalIncreaseType
	}
	if !IsNil(o.CastApprovalVotesPercentage) {
		toSerialize["castApprovalVotesPercentage"] = o.CastApprovalVotesPercentage
	}
	if !IsNil(o.CastRefusalVotesPercentage) {
		toSerialize["castRefusalVotesPercentage"] = o.CastRefusalVotesPercentage
	}
	if !IsNil(o.CastVotesPercentage) {
		toSerialize["castVotesPercentage"] = o.CastVotesPercentage
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LeftVotesPercentage) {
		toSerialize["leftVotesPercentage"] = o.LeftVotesPercentage
	}
	if !IsNil(o.MinimalCashVolume) {
		toSerialize["minimalCashVolume"] = o.MinimalCashVolume
	}
	if !IsNil(o.Motion) {
		toSerialize["motion"] = o.Motion
	}
	if !IsNil(o.NumberOfShares) {
		toSerialize["numberOfShares"] = o.NumberOfShares
	}
	if !IsNil(o.PollInitiator) {
		toSerialize["pollInitiator"] = o.PollInitiator
	}
	if !IsNil(o.RefusalVotesPercentage) {
		toSerialize["refusalVotesPercentage"] = o.RefusalVotesPercentage
	}
	if !IsNil(o.ResultExpireDate) {
		toSerialize["resultExpireDate"] = o.ResultExpireDate
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.SubscriptionFraction) {
		toSerialize["subscriptionFraction"] = o.SubscriptionFraction
	}
	if !IsNil(o.TotalNumberOfCastVotes) {
		toSerialize["totalNumberOfCastVotes"] = o.TotalNumberOfCastVotes
	}
	if !IsNil(o.TotalNumberOfVoices) {
		toSerialize["totalNumberOfVoices"] = o.TotalNumberOfVoices
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Votes) {
		toSerialize["votes"] = o.Votes
	}
	return toSerialize, nil
}

type NullableCapitalIncreasePollView struct {
	value *CapitalIncreasePollView
	isSet bool
}

func (v NullableCapitalIncreasePollView) Get() *CapitalIncreasePollView {
	return v.value
}

func (v *NullableCapitalIncreasePollView) Set(val *CapitalIncreasePollView) {
	v.value = val
	v.isSet = true
}

func (v NullableCapitalIncreasePollView) IsSet() bool {
	return v.isSet
}

func (v *NullableCapitalIncreasePollView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapitalIncreasePollView(val *CapitalIncreasePollView) *NullableCapitalIncreasePollView {
	return &NullableCapitalIncreasePollView{value: val, isSet: true}
}

func (v NullableCapitalIncreasePollView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapitalIncreasePollView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


