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

// checks if the MessageBoardWithDetailsView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageBoardWithDetailsView{}

// MessageBoardWithDetailsView struct for MessageBoardWithDetailsView
type MessageBoardWithDetailsView struct {
	DateCreated *int64 `json:"dateCreated,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	LatestPost *PostView `json:"latestPost,omitempty"`
	Locale *string `json:"locale,omitempty"`
	Name *string `json:"name,omitempty"`
	NumberOfMembers *int64 `json:"numberOfMembers,omitempty"`
	NumberOfPosts *int64 `json:"numberOfPosts,omitempty"`
	NumberOfSubboards *int64 `json:"numberOfSubboards,omitempty"`
	Owner *UsernameView `json:"owner,omitempty"`
	Parent *MessageBoardView `json:"parent,omitempty"`
	PublicBoard *bool `json:"publicBoard,omitempty"`
	Root *MessageBoardView `json:"root,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewMessageBoardWithDetailsView instantiates a new MessageBoardWithDetailsView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageBoardWithDetailsView() *MessageBoardWithDetailsView {
	this := MessageBoardWithDetailsView{}
	return &this
}

// NewMessageBoardWithDetailsViewWithDefaults instantiates a new MessageBoardWithDetailsView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageBoardWithDetailsViewWithDefaults() *MessageBoardWithDetailsView {
	this := MessageBoardWithDetailsView{}
	return &this
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *MessageBoardWithDetailsView) GetDateCreated() int64 {
	if o == nil || IsNil(o.DateCreated) {
		var ret int64
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageBoardWithDetailsView) GetDateCreatedOk() (*int64, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *MessageBoardWithDetailsView) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given int64 and assigns it to the DateCreated field.
func (o *MessageBoardWithDetailsView) SetDateCreated(v int64) {
	o.DateCreated = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MessageBoardWithDetailsView) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageBoardWithDetailsView) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MessageBoardWithDetailsView) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MessageBoardWithDetailsView) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MessageBoardWithDetailsView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageBoardWithDetailsView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MessageBoardWithDetailsView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MessageBoardWithDetailsView) SetId(v string) {
	o.Id = &v
}

// GetLatestPost returns the LatestPost field value if set, zero value otherwise.
func (o *MessageBoardWithDetailsView) GetLatestPost() PostView {
	if o == nil || IsNil(o.LatestPost) {
		var ret PostView
		return ret
	}
	return *o.LatestPost
}

// GetLatestPostOk returns a tuple with the LatestPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageBoardWithDetailsView) GetLatestPostOk() (*PostView, bool) {
	if o == nil || IsNil(o.LatestPost) {
		return nil, false
	}
	return o.LatestPost, true
}

// HasLatestPost returns a boolean if a field has been set.
func (o *MessageBoardWithDetailsView) HasLatestPost() bool {
	if o != nil && !IsNil(o.LatestPost) {
		return true
	}

	return false
}

// SetLatestPost gets a reference to the given PostView and assigns it to the LatestPost field.
func (o *MessageBoardWithDetailsView) SetLatestPost(v PostView) {
	o.LatestPost = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *MessageBoardWithDetailsView) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageBoardWithDetailsView) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *MessageBoardWithDetailsView) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *MessageBoardWithDetailsView) SetLocale(v string) {
	o.Locale = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MessageBoardWithDetailsView) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageBoardWithDetailsView) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MessageBoardWithDetailsView) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MessageBoardWithDetailsView) SetName(v string) {
	o.Name = &v
}

// GetNumberOfMembers returns the NumberOfMembers field value if set, zero value otherwise.
func (o *MessageBoardWithDetailsView) GetNumberOfMembers() int64 {
	if o == nil || IsNil(o.NumberOfMembers) {
		var ret int64
		return ret
	}
	return *o.NumberOfMembers
}

// GetNumberOfMembersOk returns a tuple with the NumberOfMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageBoardWithDetailsView) GetNumberOfMembersOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfMembers) {
		return nil, false
	}
	return o.NumberOfMembers, true
}

// HasNumberOfMembers returns a boolean if a field has been set.
func (o *MessageBoardWithDetailsView) HasNumberOfMembers() bool {
	if o != nil && !IsNil(o.NumberOfMembers) {
		return true
	}

	return false
}

// SetNumberOfMembers gets a reference to the given int64 and assigns it to the NumberOfMembers field.
func (o *MessageBoardWithDetailsView) SetNumberOfMembers(v int64) {
	o.NumberOfMembers = &v
}

// GetNumberOfPosts returns the NumberOfPosts field value if set, zero value otherwise.
func (o *MessageBoardWithDetailsView) GetNumberOfPosts() int64 {
	if o == nil || IsNil(o.NumberOfPosts) {
		var ret int64
		return ret
	}
	return *o.NumberOfPosts
}

// GetNumberOfPostsOk returns a tuple with the NumberOfPosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageBoardWithDetailsView) GetNumberOfPostsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfPosts) {
		return nil, false
	}
	return o.NumberOfPosts, true
}

// HasNumberOfPosts returns a boolean if a field has been set.
func (o *MessageBoardWithDetailsView) HasNumberOfPosts() bool {
	if o != nil && !IsNil(o.NumberOfPosts) {
		return true
	}

	return false
}

// SetNumberOfPosts gets a reference to the given int64 and assigns it to the NumberOfPosts field.
func (o *MessageBoardWithDetailsView) SetNumberOfPosts(v int64) {
	o.NumberOfPosts = &v
}

// GetNumberOfSubboards returns the NumberOfSubboards field value if set, zero value otherwise.
func (o *MessageBoardWithDetailsView) GetNumberOfSubboards() int64 {
	if o == nil || IsNil(o.NumberOfSubboards) {
		var ret int64
		return ret
	}
	return *o.NumberOfSubboards
}

// GetNumberOfSubboardsOk returns a tuple with the NumberOfSubboards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageBoardWithDetailsView) GetNumberOfSubboardsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfSubboards) {
		return nil, false
	}
	return o.NumberOfSubboards, true
}

// HasNumberOfSubboards returns a boolean if a field has been set.
func (o *MessageBoardWithDetailsView) HasNumberOfSubboards() bool {
	if o != nil && !IsNil(o.NumberOfSubboards) {
		return true
	}

	return false
}

// SetNumberOfSubboards gets a reference to the given int64 and assigns it to the NumberOfSubboards field.
func (o *MessageBoardWithDetailsView) SetNumberOfSubboards(v int64) {
	o.NumberOfSubboards = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *MessageBoardWithDetailsView) GetOwner() UsernameView {
	if o == nil || IsNil(o.Owner) {
		var ret UsernameView
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageBoardWithDetailsView) GetOwnerOk() (*UsernameView, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *MessageBoardWithDetailsView) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given UsernameView and assigns it to the Owner field.
func (o *MessageBoardWithDetailsView) SetOwner(v UsernameView) {
	o.Owner = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *MessageBoardWithDetailsView) GetParent() MessageBoardView {
	if o == nil || IsNil(o.Parent) {
		var ret MessageBoardView
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageBoardWithDetailsView) GetParentOk() (*MessageBoardView, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *MessageBoardWithDetailsView) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given MessageBoardView and assigns it to the Parent field.
func (o *MessageBoardWithDetailsView) SetParent(v MessageBoardView) {
	o.Parent = &v
}

// GetPublicBoard returns the PublicBoard field value if set, zero value otherwise.
func (o *MessageBoardWithDetailsView) GetPublicBoard() bool {
	if o == nil || IsNil(o.PublicBoard) {
		var ret bool
		return ret
	}
	return *o.PublicBoard
}

// GetPublicBoardOk returns a tuple with the PublicBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageBoardWithDetailsView) GetPublicBoardOk() (*bool, bool) {
	if o == nil || IsNil(o.PublicBoard) {
		return nil, false
	}
	return o.PublicBoard, true
}

// HasPublicBoard returns a boolean if a field has been set.
func (o *MessageBoardWithDetailsView) HasPublicBoard() bool {
	if o != nil && !IsNil(o.PublicBoard) {
		return true
	}

	return false
}

// SetPublicBoard gets a reference to the given bool and assigns it to the PublicBoard field.
func (o *MessageBoardWithDetailsView) SetPublicBoard(v bool) {
	o.PublicBoard = &v
}

// GetRoot returns the Root field value if set, zero value otherwise.
func (o *MessageBoardWithDetailsView) GetRoot() MessageBoardView {
	if o == nil || IsNil(o.Root) {
		var ret MessageBoardView
		return ret
	}
	return *o.Root
}

// GetRootOk returns a tuple with the Root field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageBoardWithDetailsView) GetRootOk() (*MessageBoardView, bool) {
	if o == nil || IsNil(o.Root) {
		return nil, false
	}
	return o.Root, true
}

// HasRoot returns a boolean if a field has been set.
func (o *MessageBoardWithDetailsView) HasRoot() bool {
	if o != nil && !IsNil(o.Root) {
		return true
	}

	return false
}

// SetRoot gets a reference to the given MessageBoardView and assigns it to the Root field.
func (o *MessageBoardWithDetailsView) SetRoot(v MessageBoardView) {
	o.Root = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MessageBoardWithDetailsView) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageBoardWithDetailsView) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MessageBoardWithDetailsView) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *MessageBoardWithDetailsView) SetVersion(v int64) {
	o.Version = &v
}

func (o MessageBoardWithDetailsView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageBoardWithDetailsView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LatestPost) {
		toSerialize["latestPost"] = o.LatestPost
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NumberOfMembers) {
		toSerialize["numberOfMembers"] = o.NumberOfMembers
	}
	if !IsNil(o.NumberOfPosts) {
		toSerialize["numberOfPosts"] = o.NumberOfPosts
	}
	if !IsNil(o.NumberOfSubboards) {
		toSerialize["numberOfSubboards"] = o.NumberOfSubboards
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.PublicBoard) {
		toSerialize["publicBoard"] = o.PublicBoard
	}
	if !IsNil(o.Root) {
		toSerialize["root"] = o.Root
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableMessageBoardWithDetailsView struct {
	value *MessageBoardWithDetailsView
	isSet bool
}

func (v NullableMessageBoardWithDetailsView) Get() *MessageBoardWithDetailsView {
	return v.value
}

func (v *NullableMessageBoardWithDetailsView) Set(val *MessageBoardWithDetailsView) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageBoardWithDetailsView) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageBoardWithDetailsView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageBoardWithDetailsView(val *MessageBoardWithDetailsView) *NullableMessageBoardWithDetailsView {
	return &NullableMessageBoardWithDetailsView{value: val, isSet: true}
}

func (v NullableMessageBoardWithDetailsView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageBoardWithDetailsView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


