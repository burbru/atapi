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

// checks if the PlainChatView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlainChatView{}

// PlainChatView struct for PlainChatView
type PlainChatView struct {
	ChatName *string `json:"chatName,omitempty"`
	DateCreated *int64 `json:"dateCreated,omitempty"`
	Id *string `json:"id,omitempty"`
	PublicChat *bool `json:"publicChat,omitempty"`
	Readonly *bool `json:"readonly,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewPlainChatView instantiates a new PlainChatView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlainChatView() *PlainChatView {
	this := PlainChatView{}
	return &this
}

// NewPlainChatViewWithDefaults instantiates a new PlainChatView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlainChatViewWithDefaults() *PlainChatView {
	this := PlainChatView{}
	return &this
}

// GetChatName returns the ChatName field value if set, zero value otherwise.
func (o *PlainChatView) GetChatName() string {
	if o == nil || IsNil(o.ChatName) {
		var ret string
		return ret
	}
	return *o.ChatName
}

// GetChatNameOk returns a tuple with the ChatName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlainChatView) GetChatNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChatName) {
		return nil, false
	}
	return o.ChatName, true
}

// HasChatName returns a boolean if a field has been set.
func (o *PlainChatView) HasChatName() bool {
	if o != nil && !IsNil(o.ChatName) {
		return true
	}

	return false
}

// SetChatName gets a reference to the given string and assigns it to the ChatName field.
func (o *PlainChatView) SetChatName(v string) {
	o.ChatName = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *PlainChatView) GetDateCreated() int64 {
	if o == nil || IsNil(o.DateCreated) {
		var ret int64
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlainChatView) GetDateCreatedOk() (*int64, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *PlainChatView) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given int64 and assigns it to the DateCreated field.
func (o *PlainChatView) SetDateCreated(v int64) {
	o.DateCreated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PlainChatView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlainChatView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PlainChatView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PlainChatView) SetId(v string) {
	o.Id = &v
}

// GetPublicChat returns the PublicChat field value if set, zero value otherwise.
func (o *PlainChatView) GetPublicChat() bool {
	if o == nil || IsNil(o.PublicChat) {
		var ret bool
		return ret
	}
	return *o.PublicChat
}

// GetPublicChatOk returns a tuple with the PublicChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlainChatView) GetPublicChatOk() (*bool, bool) {
	if o == nil || IsNil(o.PublicChat) {
		return nil, false
	}
	return o.PublicChat, true
}

// HasPublicChat returns a boolean if a field has been set.
func (o *PlainChatView) HasPublicChat() bool {
	if o != nil && !IsNil(o.PublicChat) {
		return true
	}

	return false
}

// SetPublicChat gets a reference to the given bool and assigns it to the PublicChat field.
func (o *PlainChatView) SetPublicChat(v bool) {
	o.PublicChat = &v
}

// GetReadonly returns the Readonly field value if set, zero value otherwise.
func (o *PlainChatView) GetReadonly() bool {
	if o == nil || IsNil(o.Readonly) {
		var ret bool
		return ret
	}
	return *o.Readonly
}

// GetReadonlyOk returns a tuple with the Readonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlainChatView) GetReadonlyOk() (*bool, bool) {
	if o == nil || IsNil(o.Readonly) {
		return nil, false
	}
	return o.Readonly, true
}

// HasReadonly returns a boolean if a field has been set.
func (o *PlainChatView) HasReadonly() bool {
	if o != nil && !IsNil(o.Readonly) {
		return true
	}

	return false
}

// SetReadonly gets a reference to the given bool and assigns it to the Readonly field.
func (o *PlainChatView) SetReadonly(v bool) {
	o.Readonly = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PlainChatView) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlainChatView) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PlainChatView) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *PlainChatView) SetVersion(v int64) {
	o.Version = &v
}

func (o PlainChatView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlainChatView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChatName) {
		toSerialize["chatName"] = o.ChatName
	}
	if !IsNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PublicChat) {
		toSerialize["publicChat"] = o.PublicChat
	}
	if !IsNil(o.Readonly) {
		toSerialize["readonly"] = o.Readonly
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullablePlainChatView struct {
	value *PlainChatView
	isSet bool
}

func (v NullablePlainChatView) Get() *PlainChatView {
	return v.value
}

func (v *NullablePlainChatView) Set(val *PlainChatView) {
	v.value = val
	v.isSet = true
}

func (v NullablePlainChatView) IsSet() bool {
	return v.isSet
}

func (v *NullablePlainChatView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlainChatView(val *PlainChatView) *NullablePlainChatView {
	return &NullablePlainChatView{value: val, isSet: true}
}

func (v NullablePlainChatView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlainChatView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


