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

// checks if the CompanyView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyView{}

// CompanyView struct for CompanyView
type CompanyView struct {
	AchievementCount *int64 `json:"achievementCount,omitempty"`
	AchievementTotal *int64 `json:"achievementTotal,omitempty"`
	BankAccount *BankAccountView `json:"bankAccount,omitempty"`
	Ceo *UsernameView `json:"ceo,omitempty"`
	Id *string `json:"id,omitempty"`
	Listing *ListingView `json:"listing,omitempty"`
	LogoUrl *string `json:"logoUrl,omitempty"`
	Name *string `json:"name,omitempty"`
	SecuritiesAccountId *string `json:"securitiesAccountId,omitempty"`
	SecurityIdentifier *string `json:"securityIdentifier,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewCompanyView instantiates a new CompanyView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyView() *CompanyView {
	this := CompanyView{}
	return &this
}

// NewCompanyViewWithDefaults instantiates a new CompanyView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyViewWithDefaults() *CompanyView {
	this := CompanyView{}
	return &this
}

// GetAchievementCount returns the AchievementCount field value if set, zero value otherwise.
func (o *CompanyView) GetAchievementCount() int64 {
	if o == nil || IsNil(o.AchievementCount) {
		var ret int64
		return ret
	}
	return *o.AchievementCount
}

// GetAchievementCountOk returns a tuple with the AchievementCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyView) GetAchievementCountOk() (*int64, bool) {
	if o == nil || IsNil(o.AchievementCount) {
		return nil, false
	}
	return o.AchievementCount, true
}

// HasAchievementCount returns a boolean if a field has been set.
func (o *CompanyView) HasAchievementCount() bool {
	if o != nil && !IsNil(o.AchievementCount) {
		return true
	}

	return false
}

// SetAchievementCount gets a reference to the given int64 and assigns it to the AchievementCount field.
func (o *CompanyView) SetAchievementCount(v int64) {
	o.AchievementCount = &v
}

// GetAchievementTotal returns the AchievementTotal field value if set, zero value otherwise.
func (o *CompanyView) GetAchievementTotal() int64 {
	if o == nil || IsNil(o.AchievementTotal) {
		var ret int64
		return ret
	}
	return *o.AchievementTotal
}

// GetAchievementTotalOk returns a tuple with the AchievementTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyView) GetAchievementTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.AchievementTotal) {
		return nil, false
	}
	return o.AchievementTotal, true
}

// HasAchievementTotal returns a boolean if a field has been set.
func (o *CompanyView) HasAchievementTotal() bool {
	if o != nil && !IsNil(o.AchievementTotal) {
		return true
	}

	return false
}

// SetAchievementTotal gets a reference to the given int64 and assigns it to the AchievementTotal field.
func (o *CompanyView) SetAchievementTotal(v int64) {
	o.AchievementTotal = &v
}

// GetBankAccount returns the BankAccount field value if set, zero value otherwise.
func (o *CompanyView) GetBankAccount() BankAccountView {
	if o == nil || IsNil(o.BankAccount) {
		var ret BankAccountView
		return ret
	}
	return *o.BankAccount
}

// GetBankAccountOk returns a tuple with the BankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyView) GetBankAccountOk() (*BankAccountView, bool) {
	if o == nil || IsNil(o.BankAccount) {
		return nil, false
	}
	return o.BankAccount, true
}

// HasBankAccount returns a boolean if a field has been set.
func (o *CompanyView) HasBankAccount() bool {
	if o != nil && !IsNil(o.BankAccount) {
		return true
	}

	return false
}

// SetBankAccount gets a reference to the given BankAccountView and assigns it to the BankAccount field.
func (o *CompanyView) SetBankAccount(v BankAccountView) {
	o.BankAccount = &v
}

// GetCeo returns the Ceo field value if set, zero value otherwise.
func (o *CompanyView) GetCeo() UsernameView {
	if o == nil || IsNil(o.Ceo) {
		var ret UsernameView
		return ret
	}
	return *o.Ceo
}

// GetCeoOk returns a tuple with the Ceo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyView) GetCeoOk() (*UsernameView, bool) {
	if o == nil || IsNil(o.Ceo) {
		return nil, false
	}
	return o.Ceo, true
}

// HasCeo returns a boolean if a field has been set.
func (o *CompanyView) HasCeo() bool {
	if o != nil && !IsNil(o.Ceo) {
		return true
	}

	return false
}

// SetCeo gets a reference to the given UsernameView and assigns it to the Ceo field.
func (o *CompanyView) SetCeo(v UsernameView) {
	o.Ceo = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CompanyView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CompanyView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CompanyView) SetId(v string) {
	o.Id = &v
}

// GetListing returns the Listing field value if set, zero value otherwise.
func (o *CompanyView) GetListing() ListingView {
	if o == nil || IsNil(o.Listing) {
		var ret ListingView
		return ret
	}
	return *o.Listing
}

// GetListingOk returns a tuple with the Listing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyView) GetListingOk() (*ListingView, bool) {
	if o == nil || IsNil(o.Listing) {
		return nil, false
	}
	return o.Listing, true
}

// HasListing returns a boolean if a field has been set.
func (o *CompanyView) HasListing() bool {
	if o != nil && !IsNil(o.Listing) {
		return true
	}

	return false
}

// SetListing gets a reference to the given ListingView and assigns it to the Listing field.
func (o *CompanyView) SetListing(v ListingView) {
	o.Listing = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *CompanyView) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyView) GetLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *CompanyView) HasLogoUrl() bool {
	if o != nil && !IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *CompanyView) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CompanyView) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyView) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CompanyView) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CompanyView) SetName(v string) {
	o.Name = &v
}

// GetSecuritiesAccountId returns the SecuritiesAccountId field value if set, zero value otherwise.
func (o *CompanyView) GetSecuritiesAccountId() string {
	if o == nil || IsNil(o.SecuritiesAccountId) {
		var ret string
		return ret
	}
	return *o.SecuritiesAccountId
}

// GetSecuritiesAccountIdOk returns a tuple with the SecuritiesAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyView) GetSecuritiesAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecuritiesAccountId) {
		return nil, false
	}
	return o.SecuritiesAccountId, true
}

// HasSecuritiesAccountId returns a boolean if a field has been set.
func (o *CompanyView) HasSecuritiesAccountId() bool {
	if o != nil && !IsNil(o.SecuritiesAccountId) {
		return true
	}

	return false
}

// SetSecuritiesAccountId gets a reference to the given string and assigns it to the SecuritiesAccountId field.
func (o *CompanyView) SetSecuritiesAccountId(v string) {
	o.SecuritiesAccountId = &v
}

// GetSecurityIdentifier returns the SecurityIdentifier field value if set, zero value otherwise.
func (o *CompanyView) GetSecurityIdentifier() string {
	if o == nil || IsNil(o.SecurityIdentifier) {
		var ret string
		return ret
	}
	return *o.SecurityIdentifier
}

// GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyView) GetSecurityIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityIdentifier) {
		return nil, false
	}
	return o.SecurityIdentifier, true
}

// HasSecurityIdentifier returns a boolean if a field has been set.
func (o *CompanyView) HasSecurityIdentifier() bool {
	if o != nil && !IsNil(o.SecurityIdentifier) {
		return true
	}

	return false
}

// SetSecurityIdentifier gets a reference to the given string and assigns it to the SecurityIdentifier field.
func (o *CompanyView) SetSecurityIdentifier(v string) {
	o.SecurityIdentifier = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CompanyView) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyView) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CompanyView) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *CompanyView) SetVersion(v int64) {
	o.Version = &v
}

func (o CompanyView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AchievementCount) {
		toSerialize["achievementCount"] = o.AchievementCount
	}
	if !IsNil(o.AchievementTotal) {
		toSerialize["achievementTotal"] = o.AchievementTotal
	}
	if !IsNil(o.BankAccount) {
		toSerialize["bankAccount"] = o.BankAccount
	}
	if !IsNil(o.Ceo) {
		toSerialize["ceo"] = o.Ceo
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Listing) {
		toSerialize["listing"] = o.Listing
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

type NullableCompanyView struct {
	value *CompanyView
	isSet bool
}

func (v NullableCompanyView) Get() *CompanyView {
	return v.value
}

func (v *NullableCompanyView) Set(val *CompanyView) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyView) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyView(val *CompanyView) *NullableCompanyView {
	return &NullableCompanyView{value: val, isSet: true}
}

func (v NullableCompanyView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


