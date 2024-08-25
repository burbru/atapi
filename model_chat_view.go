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

// checks if the ChatView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatView{}

// ChatView struct for ChatView
type ChatView struct {
	ChatName *string `json:"chatName,omitempty"`
	DateCreated *int64 `json:"dateCreated,omitempty"`
	GroupChat *bool `json:"groupChat,omitempty"`
	Id *string `json:"id,omitempty"`
	LastMessage *MessageView `json:"lastMessage,omitempty"`
	NumOfUnreadMessages *int64 `json:"numOfUnreadMessages,omitempty"`
	Owner *UsernameView `json:"owner,omitempty"`
	Participants []UsernameView `json:"participants,omitempty"`
	PublicChat *bool `json:"publicChat,omitempty"`
	Readonly *bool `json:"readonly,omitempty"`
	Status *string `json:"status,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewChatView instantiates a new ChatView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatView() *ChatView {
	this := ChatView{}
	return &this
}

// NewChatViewWithDefaults instantiates a new ChatView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatViewWithDefaults() *ChatView {
	this := ChatView{}
	return &this
}

// GetChatName returns the ChatName field value if set, zero value otherwise.
func (o *ChatView) GetChatName() string {
	if o == nil || IsNil(o.ChatName) {
		var ret string
		return ret
	}
	return *o.ChatName
}

// GetChatNameOk returns a tuple with the ChatName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatView) GetChatNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChatName) {
		return nil, false
	}
	return o.ChatName, true
}

// HasChatName returns a boolean if a field has been set.
func (o *ChatView) HasChatName() bool {
	if o != nil && !IsNil(o.ChatName) {
		return true
	}

	return false
}

// SetChatName gets a reference to the given string and assigns it to the ChatName field.
func (o *ChatView) SetChatName(v string) {
	o.ChatName = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ChatView) GetDateCreated() int64 {
	if o == nil || IsNil(o.DateCreated) {
		var ret int64
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatView) GetDateCreatedOk() (*int64, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ChatView) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given int64 and assigns it to the DateCreated field.
func (o *ChatView) SetDateCreated(v int64) {
	o.DateCreated = &v
}

// GetGroupChat returns the GroupChat field value if set, zero value otherwise.
func (o *ChatView) GetGroupChat() bool {
	if o == nil || IsNil(o.GroupChat) {
		var ret bool
		return ret
	}
	return *o.GroupChat
}

// GetGroupChatOk returns a tuple with the GroupChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatView) GetGroupChatOk() (*bool, bool) {
	if o == nil || IsNil(o.GroupChat) {
		return nil, false
	}
	return o.GroupChat, true
}

// HasGroupChat returns a boolean if a field has been set.
func (o *ChatView) HasGroupChat() bool {
	if o != nil && !IsNil(o.GroupChat) {
		return true
	}

	return false
}

// SetGroupChat gets a reference to the given bool and assigns it to the GroupChat field.
func (o *ChatView) SetGroupChat(v bool) {
	o.GroupChat = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChatView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChatView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChatView) SetId(v string) {
	o.Id = &v
}

// GetLastMessage returns the LastMessage field value if set, zero value otherwise.
func (o *ChatView) GetLastMessage() MessageView {
	if o == nil || IsNil(o.LastMessage) {
		var ret MessageView
		return ret
	}
	return *o.LastMessage
}

// GetLastMessageOk returns a tuple with the LastMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatView) GetLastMessageOk() (*MessageView, bool) {
	if o == nil || IsNil(o.LastMessage) {
		return nil, false
	}
	return o.LastMessage, true
}

// HasLastMessage returns a boolean if a field has been set.
func (o *ChatView) HasLastMessage() bool {
	if o != nil && !IsNil(o.LastMessage) {
		return true
	}

	return false
}

// SetLastMessage gets a reference to the given MessageView and assigns it to the LastMessage field.
func (o *ChatView) SetLastMessage(v MessageView) {
	o.LastMessage = &v
}

// GetNumOfUnreadMessages returns the NumOfUnreadMessages field value if set, zero value otherwise.
func (o *ChatView) GetNumOfUnreadMessages() int64 {
	if o == nil || IsNil(o.NumOfUnreadMessages) {
		var ret int64
		return ret
	}
	return *o.NumOfUnreadMessages
}

// GetNumOfUnreadMessagesOk returns a tuple with the NumOfUnreadMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatView) GetNumOfUnreadMessagesOk() (*int64, bool) {
	if o == nil || IsNil(o.NumOfUnreadMessages) {
		return nil, false
	}
	return o.NumOfUnreadMessages, true
}

// HasNumOfUnreadMessages returns a boolean if a field has been set.
func (o *ChatView) HasNumOfUnreadMessages() bool {
	if o != nil && !IsNil(o.NumOfUnreadMessages) {
		return true
	}

	return false
}

// SetNumOfUnreadMessages gets a reference to the given int64 and assigns it to the NumOfUnreadMessages field.
func (o *ChatView) SetNumOfUnreadMessages(v int64) {
	o.NumOfUnreadMessages = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ChatView) GetOwner() UsernameView {
	if o == nil || IsNil(o.Owner) {
		var ret UsernameView
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatView) GetOwnerOk() (*UsernameView, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ChatView) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given UsernameView and assigns it to the Owner field.
func (o *ChatView) SetOwner(v UsernameView) {
	o.Owner = &v
}

// GetParticipants returns the Participants field value if set, zero value otherwise.
func (o *ChatView) GetParticipants() []UsernameView {
	if o == nil || IsNil(o.Participants) {
		var ret []UsernameView
		return ret
	}
	return o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatView) GetParticipantsOk() ([]UsernameView, bool) {
	if o == nil || IsNil(o.Participants) {
		return nil, false
	}
	return o.Participants, true
}

// HasParticipants returns a boolean if a field has been set.
func (o *ChatView) HasParticipants() bool {
	if o != nil && !IsNil(o.Participants) {
		return true
	}

	return false
}

// SetParticipants gets a reference to the given []UsernameView and assigns it to the Participants field.
func (o *ChatView) SetParticipants(v []UsernameView) {
	o.Participants = v
}

// GetPublicChat returns the PublicChat field value if set, zero value otherwise.
func (o *ChatView) GetPublicChat() bool {
	if o == nil || IsNil(o.PublicChat) {
		var ret bool
		return ret
	}
	return *o.PublicChat
}

// GetPublicChatOk returns a tuple with the PublicChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatView) GetPublicChatOk() (*bool, bool) {
	if o == nil || IsNil(o.PublicChat) {
		return nil, false
	}
	return o.PublicChat, true
}

// HasPublicChat returns a boolean if a field has been set.
func (o *ChatView) HasPublicChat() bool {
	if o != nil && !IsNil(o.PublicChat) {
		return true
	}

	return false
}

// SetPublicChat gets a reference to the given bool and assigns it to the PublicChat field.
func (o *ChatView) SetPublicChat(v bool) {
	o.PublicChat = &v
}

// GetReadonly returns the Readonly field value if set, zero value otherwise.
func (o *ChatView) GetReadonly() bool {
	if o == nil || IsNil(o.Readonly) {
		var ret bool
		return ret
	}
	return *o.Readonly
}

// GetReadonlyOk returns a tuple with the Readonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatView) GetReadonlyOk() (*bool, bool) {
	if o == nil || IsNil(o.Readonly) {
		return nil, false
	}
	return o.Readonly, true
}

// HasReadonly returns a boolean if a field has been set.
func (o *ChatView) HasReadonly() bool {
	if o != nil && !IsNil(o.Readonly) {
		return true
	}

	return false
}

// SetReadonly gets a reference to the given bool and assigns it to the Readonly field.
func (o *ChatView) SetReadonly(v bool) {
	o.Readonly = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ChatView) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatView) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ChatView) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ChatView) SetStatus(v string) {
	o.Status = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ChatView) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatView) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ChatView) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *ChatView) SetVersion(v int64) {
	o.Version = &v
}

func (o ChatView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChatName) {
		toSerialize["chatName"] = o.ChatName
	}
	if !IsNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !IsNil(o.GroupChat) {
		toSerialize["groupChat"] = o.GroupChat
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastMessage) {
		toSerialize["lastMessage"] = o.LastMessage
	}
	if !IsNil(o.NumOfUnreadMessages) {
		toSerialize["numOfUnreadMessages"] = o.NumOfUnreadMessages
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Participants) {
		toSerialize["participants"] = o.Participants
	}
	if !IsNil(o.PublicChat) {
		toSerialize["publicChat"] = o.PublicChat
	}
	if !IsNil(o.Readonly) {
		toSerialize["readonly"] = o.Readonly
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableChatView struct {
	value *ChatView
	isSet bool
}

func (v NullableChatView) Get() *ChatView {
	return v.value
}

func (v *NullableChatView) Set(val *ChatView) {
	v.value = val
	v.isSet = true
}

func (v NullableChatView) IsSet() bool {
	return v.isSet
}

func (v *NullableChatView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatView(val *ChatView) *NullableChatView {
	return &NullableChatView{value: val, isSet: true}
}

func (v NullableChatView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


