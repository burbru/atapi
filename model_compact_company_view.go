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

// checks if the CompactCompanyView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompactCompanyView{}

// CompactCompanyView struct for CompactCompanyView
type CompactCompanyView struct {
	AchievementCount *int64 `json:"achievementCount,omitempty"`
	AchievementTotal *int64 `json:"achievementTotal,omitempty"`
	Ceo *UsernameView `json:"ceo,omitempty"`
	Id *string `json:"id,omitempty"`
	LogoUrl *string `json:"logoUrl,omitempty"`
	Name *string `json:"name,omitempty"`
	SecuritiesAccountId *string `json:"securitiesAccountId,omitempty"`
	SecurityIdentifier *string `json:"securityIdentifier,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewCompactCompanyView instantiates a new CompactCompanyView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompactCompanyView() *CompactCompanyView {
	this := CompactCompanyView{}
	return &this
}

// NewCompactCompanyViewWithDefaults instantiates a new CompactCompanyView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompactCompanyViewWithDefaults() *CompactCompanyView {
	this := CompactCompanyView{}
	return &this
}

// GetAchievementCount returns the AchievementCount field value if set, zero value otherwise.
func (o *CompactCompanyView) GetAchievementCount() int64 {
	if o == nil || IsNil(o.AchievementCount) {
		var ret int64
		return ret
	}
	return *o.AchievementCount
}

// GetAchievementCountOk returns a tuple with the AchievementCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompactCompanyView) GetAchievementCountOk() (*int64, bool) {
	if o == nil || IsNil(o.AchievementCount) {
		return nil, false
	}
	return o.AchievementCount, true
}

// HasAchievementCount returns a boolean if a field has been set.
func (o *CompactCompanyView) HasAchievementCount() bool {
	if o != nil && !IsNil(o.AchievementCount) {
		return true
	}

	return false
}

// SetAchievementCount gets a reference to the given int64 and assigns it to the AchievementCount field.
func (o *CompactCompanyView) SetAchievementCount(v int64) {
	o.AchievementCount = &v
}

// GetAchievementTotal returns the AchievementTotal field value if set, zero value otherwise.
func (o *CompactCompanyView) GetAchievementTotal() int64 {
	if o == nil || IsNil(o.AchievementTotal) {
		var ret int64
		return ret
	}
	return *o.AchievementTotal
}

// GetAchievementTotalOk returns a tuple with the AchievementTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompactCompanyView) GetAchievementTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.AchievementTotal) {
		return nil, false
	}
	return o.AchievementTotal, true
}

// HasAchievementTotal returns a boolean if a field has been set.
func (o *CompactCompanyView) HasAchievementTotal() bool {
	if o != nil && !IsNil(o.AchievementTotal) {
		return true
	}

	return false
}

// SetAchievementTotal gets a reference to the given int64 and assigns it to the AchievementTotal field.
func (o *CompactCompanyView) SetAchievementTotal(v int64) {
	o.AchievementTotal = &v
}

// GetCeo returns the Ceo field value if set, zero value otherwise.
func (o *CompactCompanyView) GetCeo() UsernameView {
	if o == nil || IsNil(o.Ceo) {
		var ret UsernameView
		return ret
	}
	return *o.Ceo
}

// GetCeoOk returns a tuple with the Ceo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompactCompanyView) GetCeoOk() (*UsernameView, bool) {
	if o == nil || IsNil(o.Ceo) {
		return nil, false
	}
	return o.Ceo, true
}

// HasCeo returns a boolean if a field has been set.
func (o *CompactCompanyView) HasCeo() bool {
	if o != nil && !IsNil(o.Ceo) {
		return true
	}

	return false
}

// SetCeo gets a reference to the given UsernameView and assigns it to the Ceo field.
func (o *CompactCompanyView) SetCeo(v UsernameView) {
	o.Ceo = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CompactCompanyView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompactCompanyView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CompactCompanyView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CompactCompanyView) SetId(v string) {
	o.Id = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *CompactCompanyView) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompactCompanyView) GetLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *CompactCompanyView) HasLogoUrl() bool {
	if o != nil && !IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *CompactCompanyView) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CompactCompanyView) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompactCompanyView) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CompactCompanyView) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CompactCompanyView) SetName(v string) {
	o.Name = &v
}

// GetSecuritiesAccountId returns the SecuritiesAccountId field value if set, zero value otherwise.
func (o *CompactCompanyView) GetSecuritiesAccountId() string {
	if o == nil || IsNil(o.SecuritiesAccountId) {
		var ret string
		return ret
	}
	return *o.SecuritiesAccountId
}

// GetSecuritiesAccountIdOk returns a tuple with the SecuritiesAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompactCompanyView) GetSecuritiesAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecuritiesAccountId) {
		return nil, false
	}
	return o.SecuritiesAccountId, true
}

// HasSecuritiesAccountId returns a boolean if a field has been set.
func (o *CompactCompanyView) HasSecuritiesAccountId() bool {
	if o != nil && !IsNil(o.SecuritiesAccountId) {
		return true
	}

	return false
}

// SetSecuritiesAccountId gets a reference to the given string and assigns it to the SecuritiesAccountId field.
func (o *CompactCompanyView) SetSecuritiesAccountId(v string) {
	o.SecuritiesAccountId = &v
}

// GetSecurityIdentifier returns the SecurityIdentifier field value if set, zero value otherwise.
func (o *CompactCompanyView) GetSecurityIdentifier() string {
	if o == nil || IsNil(o.SecurityIdentifier) {
		var ret string
		return ret
	}
	return *o.SecurityIdentifier
}

// GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompactCompanyView) GetSecurityIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityIdentifier) {
		return nil, false
	}
	return o.SecurityIdentifier, true
}

// HasSecurityIdentifier returns a boolean if a field has been set.
func (o *CompactCompanyView) HasSecurityIdentifier() bool {
	if o != nil && !IsNil(o.SecurityIdentifier) {
		return true
	}

	return false
}

// SetSecurityIdentifier gets a reference to the given string and assigns it to the SecurityIdentifier field.
func (o *CompactCompanyView) SetSecurityIdentifier(v string) {
	o.SecurityIdentifier = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CompactCompanyView) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompactCompanyView) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CompactCompanyView) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *CompactCompanyView) SetVersion(v int64) {
	o.Version = &v
}

func (o CompactCompanyView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompactCompanyView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AchievementCount) {
		toSerialize["achievementCount"] = o.AchievementCount
	}
	if !IsNil(o.AchievementTotal) {
		toSerialize["achievementTotal"] = o.AchievementTotal
	}
	if !IsNil(o.Ceo) {
		toSerialize["ceo"] = o.Ceo
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LogoUrl) {
		toSerialize["logoUrl"] = o.LogoUrl
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SecuritiesAccountId) {
		toSerialize["securitiesAccountId"] = o.SecuritiesAccountId
	}
	if !IsNil(o.SecurityIdentifier) {
		toSerialize["securityIdentifier"] = o.SecurityIdentifier
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableCompactCompanyView struct {
	value *CompactCompanyView
	isSet bool
}

func (v NullableCompactCompanyView) Get() *CompactCompanyView {
	return v.value
}

func (v *NullableCompactCompanyView) Set(val *CompactCompanyView) {
	v.value = val
	v.isSet = true
}

func (v NullableCompactCompanyView) IsSet() bool {
	return v.isSet
}

func (v *NullableCompactCompanyView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompactCompanyView(val *CompactCompanyView) *NullableCompactCompanyView {
	return &NullableCompactCompanyView{value: val, isSet: true}
}

func (v NullableCompactCompanyView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompactCompanyView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


