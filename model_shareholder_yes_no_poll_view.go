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

// checks if the ShareholderYesNoPollView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShareholderYesNoPollView{}

// ShareholderYesNoPollView struct for ShareholderYesNoPollView
type ShareholderYesNoPollView struct {
	AbstentionRule *string `json:"abstentionRule,omitempty"`
	ApprovalVotesPercentage *float32 `json:"approvalVotesPercentage,omitempty"`
	CastApprovalVotesPercentage *float32 `json:"castApprovalVotesPercentage,omitempty"`
	CastRefusalVotesPercentage *float32 `json:"castRefusalVotesPercentage,omitempty"`
	CastVotesPercentage *float32 `json:"castVotesPercentage,omitempty"`
	Company *CompactCompanyView `json:"company,omitempty"`
	EndDate *int64 `json:"endDate,omitempty"`
	Group []VoiceNumberView `json:"group,omitempty"`
	Id *string `json:"id,omitempty"`
	LeftVotesPercentage *float32 `json:"leftVotesPercentage,omitempty"`
	Motion *string `json:"motion,omitempty"`
	PollInitiator *UsernameView `json:"pollInitiator,omitempty"`
	RefusalVotesPercentage *float32 `json:"refusalVotesPercentage,omitempty"`
	ResultExpireDate *int64 `json:"resultExpireDate,omitempty"`
	StartDate *int64 `json:"startDate,omitempty"`
	TotalNumberOfCastVotes *int64 `json:"totalNumberOfCastVotes,omitempty"`
	TotalNumberOfVoices *int64 `json:"totalNumberOfVoices,omitempty"`
	Type *string `json:"type,omitempty"`
	Version *int64 `json:"version,omitempty"`
	Votes []VoteView `json:"votes,omitempty"`
}

// NewShareholderYesNoPollView instantiates a new ShareholderYesNoPollView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShareholderYesNoPollView() *ShareholderYesNoPollView {
	this := ShareholderYesNoPollView{}
	return &this
}

// NewShareholderYesNoPollViewWithDefaults instantiates a new ShareholderYesNoPollView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShareholderYesNoPollViewWithDefaults() *ShareholderYesNoPollView {
	this := ShareholderYesNoPollView{}
	return &this
}

// GetAbstentionRule returns the AbstentionRule field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetAbstentionRule() string {
	if o == nil || IsNil(o.AbstentionRule) {
		var ret string
		return ret
	}
	return *o.AbstentionRule
}

// GetAbstentionRuleOk returns a tuple with the AbstentionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetAbstentionRuleOk() (*string, bool) {
	if o == nil || IsNil(o.AbstentionRule) {
		return nil, false
	}
	return o.AbstentionRule, true
}

// HasAbstentionRule returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasAbstentionRule() bool {
	if o != nil && !IsNil(o.AbstentionRule) {
		return true
	}

	return false
}

// SetAbstentionRule gets a reference to the given string and assigns it to the AbstentionRule field.
func (o *ShareholderYesNoPollView) SetAbstentionRule(v string) {
	o.AbstentionRule = &v
}

// GetApprovalVotesPercentage returns the ApprovalVotesPercentage field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetApprovalVotesPercentage() float32 {
	if o == nil || IsNil(o.ApprovalVotesPercentage) {
		var ret float32
		return ret
	}
	return *o.ApprovalVotesPercentage
}

// GetApprovalVotesPercentageOk returns a tuple with the ApprovalVotesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetApprovalVotesPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.ApprovalVotesPercentage) {
		return nil, false
	}
	return o.ApprovalVotesPercentage, true
}

// HasApprovalVotesPercentage returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasApprovalVotesPercentage() bool {
	if o != nil && !IsNil(o.ApprovalVotesPercentage) {
		return true
	}

	return false
}

// SetApprovalVotesPercentage gets a reference to the given float32 and assigns it to the ApprovalVotesPercentage field.
func (o *ShareholderYesNoPollView) SetApprovalVotesPercentage(v float32) {
	o.ApprovalVotesPercentage = &v
}

// GetCastApprovalVotesPercentage returns the CastApprovalVotesPercentage field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetCastApprovalVotesPercentage() float32 {
	if o == nil || IsNil(o.CastApprovalVotesPercentage) {
		var ret float32
		return ret
	}
	return *o.CastApprovalVotesPercentage
}

// GetCastApprovalVotesPercentageOk returns a tuple with the CastApprovalVotesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetCastApprovalVotesPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.CastApprovalVotesPercentage) {
		return nil, false
	}
	return o.CastApprovalVotesPercentage, true
}

// HasCastApprovalVotesPercentage returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasCastApprovalVotesPercentage() bool {
	if o != nil && !IsNil(o.CastApprovalVotesPercentage) {
		return true
	}

	return false
}

// SetCastApprovalVotesPercentage gets a reference to the given float32 and assigns it to the CastApprovalVotesPercentage field.
func (o *ShareholderYesNoPollView) SetCastApprovalVotesPercentage(v float32) {
	o.CastApprovalVotesPercentage = &v
}

// GetCastRefusalVotesPercentage returns the CastRefusalVotesPercentage field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetCastRefusalVotesPercentage() float32 {
	if o == nil || IsNil(o.CastRefusalVotesPercentage) {
		var ret float32
		return ret
	}
	return *o.CastRefusalVotesPercentage
}

// GetCastRefusalVotesPercentageOk returns a tuple with the CastRefusalVotesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetCastRefusalVotesPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.CastRefusalVotesPercentage) {
		return nil, false
	}
	return o.CastRefusalVotesPercentage, true
}

// HasCastRefusalVotesPercentage returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasCastRefusalVotesPercentage() bool {
	if o != nil && !IsNil(o.CastRefusalVotesPercentage) {
		return true
	}

	return false
}

// SetCastRefusalVotesPercentage gets a reference to the given float32 and assigns it to the CastRefusalVotesPercentage field.
func (o *ShareholderYesNoPollView) SetCastRefusalVotesPercentage(v float32) {
	o.CastRefusalVotesPercentage = &v
}

// GetCastVotesPercentage returns the CastVotesPercentage field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetCastVotesPercentage() float32 {
	if o == nil || IsNil(o.CastVotesPercentage) {
		var ret float32
		return ret
	}
	return *o.CastVotesPercentage
}

// GetCastVotesPercentageOk returns a tuple with the CastVotesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetCastVotesPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.CastVotesPercentage) {
		return nil, false
	}
	return o.CastVotesPercentage, true
}

// HasCastVotesPercentage returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasCastVotesPercentage() bool {
	if o != nil && !IsNil(o.CastVotesPercentage) {
		return true
	}

	return false
}

// SetCastVotesPercentage gets a reference to the given float32 and assigns it to the CastVotesPercentage field.
func (o *ShareholderYesNoPollView) SetCastVotesPercentage(v float32) {
	o.CastVotesPercentage = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetCompany() CompactCompanyView {
	if o == nil || IsNil(o.Company) {
		var ret CompactCompanyView
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetCompanyOk() (*CompactCompanyView, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given CompactCompanyView and assigns it to the Company field.
func (o *ShareholderYesNoPollView) SetCompany(v CompactCompanyView) {
	o.Company = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetEndDate() int64 {
	if o == nil || IsNil(o.EndDate) {
		var ret int64
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetEndDateOk() (*int64, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int64 and assigns it to the EndDate field.
func (o *ShareholderYesNoPollView) SetEndDate(v int64) {
	o.EndDate = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetGroup() []VoiceNumberView {
	if o == nil || IsNil(o.Group) {
		var ret []VoiceNumberView
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetGroupOk() ([]VoiceNumberView, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given []VoiceNumberView and assigns it to the Group field.
func (o *ShareholderYesNoPollView) SetGroup(v []VoiceNumberView) {
	o.Group = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ShareholderYesNoPollView) SetId(v string) {
	o.Id = &v
}

// GetLeftVotesPercentage returns the LeftVotesPercentage field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetLeftVotesPercentage() float32 {
	if o == nil || IsNil(o.LeftVotesPercentage) {
		var ret float32
		return ret
	}
	return *o.LeftVotesPercentage
}

// GetLeftVotesPercentageOk returns a tuple with the LeftVotesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetLeftVotesPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.LeftVotesPercentage) {
		return nil, false
	}
	return o.LeftVotesPercentage, true
}

// HasLeftVotesPercentage returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasLeftVotesPercentage() bool {
	if o != nil && !IsNil(o.LeftVotesPercentage) {
		return true
	}

	return false
}

// SetLeftVotesPercentage gets a reference to the given float32 and assigns it to the LeftVotesPercentage field.
func (o *ShareholderYesNoPollView) SetLeftVotesPercentage(v float32) {
	o.LeftVotesPercentage = &v
}

// GetMotion returns the Motion field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetMotion() string {
	if o == nil || IsNil(o.Motion) {
		var ret string
		return ret
	}
	return *o.Motion
}

// GetMotionOk returns a tuple with the Motion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetMotionOk() (*string, bool) {
	if o == nil || IsNil(o.Motion) {
		return nil, false
	}
	return o.Motion, true
}

// HasMotion returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasMotion() bool {
	if o != nil && !IsNil(o.Motion) {
		return true
	}

	return false
}

// SetMotion gets a reference to the given string and assigns it to the Motion field.
func (o *ShareholderYesNoPollView) SetMotion(v string) {
	o.Motion = &v
}

// GetPollInitiator returns the PollInitiator field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetPollInitiator() UsernameView {
	if o == nil || IsNil(o.PollInitiator) {
		var ret UsernameView
		return ret
	}
	return *o.PollInitiator
}

// GetPollInitiatorOk returns a tuple with the PollInitiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetPollInitiatorOk() (*UsernameView, bool) {
	if o == nil || IsNil(o.PollInitiator) {
		return nil, false
	}
	return o.PollInitiator, true
}

// HasPollInitiator returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasPollInitiator() bool {
	if o != nil && !IsNil(o.PollInitiator) {
		return true
	}

	return false
}

// SetPollInitiator gets a reference to the given UsernameView and assigns it to the PollInitiator field.
func (o *ShareholderYesNoPollView) SetPollInitiator(v UsernameView) {
	o.PollInitiator = &v
}

// GetRefusalVotesPercentage returns the RefusalVotesPercentage field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetRefusalVotesPercentage() float32 {
	if o == nil || IsNil(o.RefusalVotesPercentage) {
		var ret float32
		return ret
	}
	return *o.RefusalVotesPercentage
}

// GetRefusalVotesPercentageOk returns a tuple with the RefusalVotesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetRefusalVotesPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.RefusalVotesPercentage) {
		return nil, false
	}
	return o.RefusalVotesPercentage, true
}

// HasRefusalVotesPercentage returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasRefusalVotesPercentage() bool {
	if o != nil && !IsNil(o.RefusalVotesPercentage) {
		return true
	}

	return false
}

// SetRefusalVotesPercentage gets a reference to the given float32 and assigns it to the RefusalVotesPercentage field.
func (o *ShareholderYesNoPollView) SetRefusalVotesPercentage(v float32) {
	o.RefusalVotesPercentage = &v
}

// GetResultExpireDate returns the ResultExpireDate field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetResultExpireDate() int64 {
	if o == nil || IsNil(o.ResultExpireDate) {
		var ret int64
		return ret
	}
	return *o.ResultExpireDate
}

// GetResultExpireDateOk returns a tuple with the ResultExpireDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetResultExpireDateOk() (*int64, bool) {
	if o == nil || IsNil(o.ResultExpireDate) {
		return nil, false
	}
	return o.ResultExpireDate, true
}

// HasResultExpireDate returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasResultExpireDate() bool {
	if o != nil && !IsNil(o.ResultExpireDate) {
		return true
	}

	return false
}

// SetResultExpireDate gets a reference to the given int64 and assigns it to the ResultExpireDate field.
func (o *ShareholderYesNoPollView) SetResultExpireDate(v int64) {
	o.ResultExpireDate = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetStartDate() int64 {
	if o == nil || IsNil(o.StartDate) {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetStartDateOk() (*int64, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *ShareholderYesNoPollView) SetStartDate(v int64) {
	o.StartDate = &v
}

// GetTotalNumberOfCastVotes returns the TotalNumberOfCastVotes field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetTotalNumberOfCastVotes() int64 {
	if o == nil || IsNil(o.TotalNumberOfCastVotes) {
		var ret int64
		return ret
	}
	return *o.TotalNumberOfCastVotes
}

// GetTotalNumberOfCastVotesOk returns a tuple with the TotalNumberOfCastVotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetTotalNumberOfCastVotesOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalNumberOfCastVotes) {
		return nil, false
	}
	return o.TotalNumberOfCastVotes, true
}

// HasTotalNumberOfCastVotes returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasTotalNumberOfCastVotes() bool {
	if o != nil && !IsNil(o.TotalNumberOfCastVotes) {
		return true
	}

	return false
}

// SetTotalNumberOfCastVotes gets a reference to the given int64 and assigns it to the TotalNumberOfCastVotes field.
func (o *ShareholderYesNoPollView) SetTotalNumberOfCastVotes(v int64) {
	o.TotalNumberOfCastVotes = &v
}

// GetTotalNumberOfVoices returns the TotalNumberOfVoices field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetTotalNumberOfVoices() int64 {
	if o == nil || IsNil(o.TotalNumberOfVoices) {
		var ret int64
		return ret
	}
	return *o.TotalNumberOfVoices
}

// GetTotalNumberOfVoicesOk returns a tuple with the TotalNumberOfVoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetTotalNumberOfVoicesOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalNumberOfVoices) {
		return nil, false
	}
	return o.TotalNumberOfVoices, true
}

// HasTotalNumberOfVoices returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasTotalNumberOfVoices() bool {
	if o != nil && !IsNil(o.TotalNumberOfVoices) {
		return true
	}

	return false
}

// SetTotalNumberOfVoices gets a reference to the given int64 and assigns it to the TotalNumberOfVoices field.
func (o *ShareholderYesNoPollView) SetTotalNumberOfVoices(v int64) {
	o.TotalNumberOfVoices = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ShareholderYesNoPollView) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *ShareholderYesNoPollView) SetVersion(v int64) {
	o.Version = &v
}

// GetVotes returns the Votes field value if set, zero value otherwise.
func (o *ShareholderYesNoPollView) GetVotes() []VoteView {
	if o == nil || IsNil(o.Votes) {
		var ret []VoteView
		return ret
	}
	return o.Votes
}

// GetVotesOk returns a tuple with the Votes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareholderYesNoPollView) GetVotesOk() ([]VoteView, bool) {
	if o == nil || IsNil(o.Votes) {
		return nil, false
	}
	return o.Votes, true
}

// HasVotes returns a boolean if a field has been set.
func (o *ShareholderYesNoPollView) HasVotes() bool {
	if o != nil && !IsNil(o.Votes) {
		return true
	}

	return false
}

// SetVotes gets a reference to the given []VoteView and assigns it to the Votes field.
func (o *ShareholderYesNoPollView) SetVotes(v []VoteView) {
	o.Votes = v
}

func (o ShareholderYesNoPollView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShareholderYesNoPollView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AbstentionRule) {
		toSerialize["abstentionRule"] = o.AbstentionRule
	}
	if !IsNil(o.ApprovalVotesPercentage) {
		toSerialize["approvalVotesPercentage"] = o.ApprovalVotesPercentage
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
	if !IsNil(o.Motion) {
		toSerialize["motion"] = o.Motion
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

type NullableShareholderYesNoPollView struct {
	value *ShareholderYesNoPollView
	isSet bool
}

func (v NullableShareholderYesNoPollView) Get() *ShareholderYesNoPollView {
	return v.value
}

func (v *NullableShareholderYesNoPollView) Set(val *ShareholderYesNoPollView) {
	v.value = val
	v.isSet = true
}

func (v NullableShareholderYesNoPollView) IsSet() bool {
	return v.isSet
}

func (v *NullableShareholderYesNoPollView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShareholderYesNoPollView(val *ShareholderYesNoPollView) *NullableShareholderYesNoPollView {
	return &NullableShareholderYesNoPollView{value: val, isSet: true}
}

func (v NullableShareholderYesNoPollView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShareholderYesNoPollView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


