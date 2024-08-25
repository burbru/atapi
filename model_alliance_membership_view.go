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

// checks if the AllianceMembershipView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllianceMembershipView{}

// AllianceMembershipView struct for AllianceMembershipView
type AllianceMembershipView struct {
	Alliance *AllianceView `json:"alliance,omitempty"`
	DateJoined *int64 `json:"dateJoined,omitempty"`
	Id *string `json:"id,omitempty"`
	Member *UsernameView `json:"member,omitempty"`
	Online *bool `json:"online,omitempty"`
	Role *string `json:"role,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewAllianceMembershipView instantiates a new AllianceMembershipView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllianceMembershipView() *AllianceMembershipView {
	this := AllianceMembershipView{}
	return &this
}

// NewAllianceMembershipViewWithDefaults instantiates a new AllianceMembershipView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllianceMembershipViewWithDefaults() *AllianceMembershipView {
	this := AllianceMembershipView{}
	return &this
}

// GetAlliance returns the Alliance field value if set, zero value otherwise.
func (o *AllianceMembershipView) GetAlliance() AllianceView {
	if o == nil || IsNil(o.Alliance) {
		var ret AllianceView
		return ret
	}
	return *o.Alliance
}

// GetAllianceOk returns a tuple with the Alliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllianceMembershipView) GetAllianceOk() (*AllianceView, bool) {
	if o == nil || IsNil(o.Alliance) {
		return nil, false
	}
	return o.Alliance, true
}

// HasAlliance returns a boolean if a field has been set.
func (o *AllianceMembershipView) HasAlliance() bool {
	if o != nil && !IsNil(o.Alliance) {
		return true
	}

	return false
}

// SetAlliance gets a reference to the given AllianceView and assigns it to the Alliance field.
func (o *AllianceMembershipView) SetAlliance(v AllianceView) {
	o.Alliance = &v
}

// GetDateJoined returns the DateJoined field value if set, zero value otherwise.
func (o *AllianceMembershipView) GetDateJoined() int64 {
	if o == nil || IsNil(o.DateJoined) {
		var ret int64
		return ret
	}
	return *o.DateJoined
}

// GetDateJoinedOk returns a tuple with the DateJoined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllianceMembershipView) GetDateJoinedOk() (*int64, bool) {
	if o == nil || IsNil(o.DateJoined) {
		return nil, false
	}
	return o.DateJoined, true
}

// HasDateJoined returns a boolean if a field has been set.
func (o *AllianceMembershipView) HasDateJoined() bool {
	if o != nil && !IsNil(o.DateJoined) {
		return true
	}

	return false
}

// SetDateJoined gets a reference to the given int64 and assigns it to the DateJoined field.
func (o *AllianceMembershipView) SetDateJoined(v int64) {
	o.DateJoined = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AllianceMembershipView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllianceMembershipView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AllianceMembershipView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AllianceMembershipView) SetId(v string) {
	o.Id = &v
}

// GetMember returns the Member field value if set, zero value otherwise.
func (o *AllianceMembershipView) GetMember() UsernameView {
	if o == nil || IsNil(o.Member) {
		var ret UsernameView
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllianceMembershipView) GetMemberOk() (*UsernameView, bool) {
	if o == nil || IsNil(o.Member) {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *AllianceMembershipView) HasMember() bool {
	if o != nil && !IsNil(o.Member) {
		return true
	}

	return false
}

// SetMember gets a reference to the given UsernameView and assigns it to the Member field.
func (o *AllianceMembershipView) SetMember(v UsernameView) {
	o.Member = &v
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *AllianceMembershipView) GetOnline() bool {
	if o == nil || IsNil(o.Online) {
		var ret bool
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllianceMembershipView) GetOnlineOk() (*bool, bool) {
	if o == nil || IsNil(o.Online) {
		return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *AllianceMembershipView) HasOnline() bool {
	if o != nil && !IsNil(o.Online) {
		return true
	}

	return false
}

// SetOnline gets a reference to the given bool and assigns it to the Online field.
func (o *AllianceMembershipView) SetOnline(v bool) {
	o.Online = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AllianceMembershipView) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllianceMembershipView) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AllianceMembershipView) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *AllianceMembershipView) SetRole(v string) {
	o.Role = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AllianceMembershipView) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllianceMembershipView) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AllianceMembershipView) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *AllianceMembershipView) SetVersion(v int64) {
	o.Version = &v
}

func (o AllianceMembershipView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllianceMembershipView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alliance) {
		toSerialize["alliance"] = o.Alliance
	}
	if !IsNil(o.DateJoined) {
		toSerialize["dateJoined"] = o.DateJoined
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Member) {
		toSerialize["member"] = o.Member
	}
	if !IsNil(o.Online) {
		toSerialize["online"] = o.Online
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableAllianceMembershipView struct {
	value *AllianceMembershipView
	isSet bool
}

func (v NullableAllianceMembershipView) Get() *AllianceMembershipView {
	return v.value
}

func (v *NullableAllianceMembershipView) Set(val *AllianceMembershipView) {
	v.value = val
	v.isSet = true
}

func (v NullableAllianceMembershipView) IsSet() bool {
	return v.isSet
}

func (v *NullableAllianceMembershipView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllianceMembershipView(val *AllianceMembershipView) *NullableAllianceMembershipView {
	return &NullableAllianceMembershipView{value: val, isSet: true}
}

func (v NullableAllianceMembershipView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllianceMembershipView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


