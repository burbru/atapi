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

// checks if the ComplaintView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComplaintView{}

// ComplaintView struct for ComplaintView
type ComplaintView struct {
	CloseDate *int64 `json:"closeDate,omitempty"`
	Comment *string `json:"comment,omitempty"`
	Id *string `json:"id,omitempty"`
	OpenDate *int64 `json:"openDate,omitempty"`
	ReviewDate *int64 `json:"reviewDate,omitempty"`
	Status *string `json:"status,omitempty"`
	SubjectMatter *string `json:"subjectMatter,omitempty"`
	SubjectMatterType *string `json:"subjectMatterType,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewComplaintView instantiates a new ComplaintView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComplaintView() *ComplaintView {
	this := ComplaintView{}
	return &this
}

// NewComplaintViewWithDefaults instantiates a new ComplaintView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComplaintViewWithDefaults() *ComplaintView {
	this := ComplaintView{}
	return &this
}

// GetCloseDate returns the CloseDate field value if set, zero value otherwise.
func (o *ComplaintView) GetCloseDate() int64 {
	if o == nil || IsNil(o.CloseDate) {
		var ret int64
		return ret
	}
	return *o.CloseDate
}

// GetCloseDateOk returns a tuple with the CloseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplaintView) GetCloseDateOk() (*int64, bool) {
	if o == nil || IsNil(o.CloseDate) {
		return nil, false
	}
	return o.CloseDate, true
}

// HasCloseDate returns a boolean if a field has been set.
func (o *ComplaintView) HasCloseDate() bool {
	if o != nil && !IsNil(o.CloseDate) {
		return true
	}

	return false
}

// SetCloseDate gets a reference to the given int64 and assigns it to the CloseDate field.
func (o *ComplaintView) SetCloseDate(v int64) {
	o.CloseDate = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ComplaintView) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplaintView) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ComplaintView) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ComplaintView) SetComment(v string) {
	o.Comment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComplaintView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplaintView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComplaintView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ComplaintView) SetId(v string) {
	o.Id = &v
}

// GetOpenDate returns the OpenDate field value if set, zero value otherwise.
func (o *ComplaintView) GetOpenDate() int64 {
	if o == nil || IsNil(o.OpenDate) {
		var ret int64
		return ret
	}
	return *o.OpenDate
}

// GetOpenDateOk returns a tuple with the OpenDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplaintView) GetOpenDateOk() (*int64, bool) {
	if o == nil || IsNil(o.OpenDate) {
		return nil, false
	}
	return o.OpenDate, true
}

// HasOpenDate returns a boolean if a field has been set.
func (o *ComplaintView) HasOpenDate() bool {
	if o != nil && !IsNil(o.OpenDate) {
		return true
	}

	return false
}

// SetOpenDate gets a reference to the given int64 and assigns it to the OpenDate field.
func (o *ComplaintView) SetOpenDate(v int64) {
	o.OpenDate = &v
}

// GetReviewDate returns the ReviewDate field value if set, zero value otherwise.
func (o *ComplaintView) GetReviewDate() int64 {
	if o == nil || IsNil(o.ReviewDate) {
		var ret int64
		return ret
	}
	return *o.ReviewDate
}

// GetReviewDateOk returns a tuple with the ReviewDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplaintView) GetReviewDateOk() (*int64, bool) {
	if o == nil || IsNil(o.ReviewDate) {
		return nil, false
	}
	return o.ReviewDate, true
}

// HasReviewDate returns a boolean if a field has been set.
func (o *ComplaintView) HasReviewDate() bool {
	if o != nil && !IsNil(o.ReviewDate) {
		return true
	}

	return false
}

// SetReviewDate gets a reference to the given int64 and assigns it to the ReviewDate field.
func (o *ComplaintView) SetReviewDate(v int64) {
	o.ReviewDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ComplaintView) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplaintView) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ComplaintView) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ComplaintView) SetStatus(v string) {
	o.Status = &v
}

// GetSubjectMatter returns the SubjectMatter field value if set, zero value otherwise.
func (o *ComplaintView) GetSubjectMatter() string {
	if o == nil || IsNil(o.SubjectMatter) {
		var ret string
		return ret
	}
	return *o.SubjectMatter
}

// GetSubjectMatterOk returns a tuple with the SubjectMatter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplaintView) GetSubjectMatterOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectMatter) {
		return nil, false
	}
	return o.SubjectMatter, true
}

// HasSubjectMatter returns a boolean if a field has been set.
func (o *ComplaintView) HasSubjectMatter() bool {
	if o != nil && !IsNil(o.SubjectMatter) {
		return true
	}

	return false
}

// SetSubjectMatter gets a reference to the given string and assigns it to the SubjectMatter field.
func (o *ComplaintView) SetSubjectMatter(v string) {
	o.SubjectMatter = &v
}

// GetSubjectMatterType returns the SubjectMatterType field value if set, zero value otherwise.
func (o *ComplaintView) GetSubjectMatterType() string {
	if o == nil || IsNil(o.SubjectMatterType) {
		var ret string
		return ret
	}
	return *o.SubjectMatterType
}

// GetSubjectMatterTypeOk returns a tuple with the SubjectMatterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplaintView) GetSubjectMatterTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectMatterType) {
		return nil, false
	}
	return o.SubjectMatterType, true
}

// HasSubjectMatterType returns a boolean if a field has been set.
func (o *ComplaintView) HasSubjectMatterType() bool {
	if o != nil && !IsNil(o.SubjectMatterType) {
		return true
	}

	return false
}

// SetSubjectMatterType gets a reference to the given string and assigns it to the SubjectMatterType field.
func (o *ComplaintView) SetSubjectMatterType(v string) {
	o.SubjectMatterType = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ComplaintView) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplaintView) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ComplaintView) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *ComplaintView) SetVersion(v int64) {
	o.Version = &v
}

func (o ComplaintView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComplaintView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloseDate) {
		toSerialize["closeDate"] = o.CloseDate
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.OpenDate) {
		toSerialize["openDate"] = o.OpenDate
	}
	if !IsNil(o.ReviewDate) {
		toSerialize["reviewDate"] = o.ReviewDate
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubjectMatter) {
		toSerialize["subjectMatter"] = o.SubjectMatter
	}
	if !IsNil(o.SubjectMatterType) {
		toSerialize["subjectMatterType"] = o.SubjectMatterType
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableComplaintView struct {
	value *ComplaintView
	isSet bool
}

func (v NullableComplaintView) Get() *ComplaintView {
	return v.value
}

func (v *NullableComplaintView) Set(val *ComplaintView) {
	v.value = val
	v.isSet = true
}

func (v NullableComplaintView) IsSet() bool {
	return v.isSet
}

func (v *NullableComplaintView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComplaintView(val *ComplaintView) *NullableComplaintView {
	return &NullableComplaintView{value: val, isSet: true}
}

func (v NullableComplaintView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComplaintView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


