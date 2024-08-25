# MessagePrototype

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilledString** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Substitutions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMessagePrototype

`func NewMessagePrototype() *MessagePrototype`

NewMessagePrototype instantiates a new MessagePrototype object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagePrototypeWithDefaults

`func NewMessagePrototypeWithDefaults() *MessagePrototype`

NewMessagePrototypeWithDefaults instantiates a new MessagePrototype object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilledString

`func (o *MessagePrototype) GetFilledString() string`

GetFilledString returns the FilledString field if non-nil, zero value otherwise.

### GetFilledStringOk

`func (o *MessagePrototype) GetFilledStringOk() (*string, bool)`

GetFilledStringOk returns a tuple with the FilledString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledString

`func (o *MessagePrototype) SetFilledString(v string)`

SetFilledString sets FilledString field to given value.

### HasFilledString

`func (o *MessagePrototype) HasFilledString() bool`

HasFilledString returns a boolean if a field has been set.

### GetMessage

`func (o *MessagePrototype) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MessagePrototype) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MessagePrototype) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MessagePrototype) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSubstitutions

`func (o *MessagePrototype) GetSubstitutions() []string`

GetSubstitutions returns the Substitutions field if non-nil, zero value otherwise.

### GetSubstitutionsOk

`func (o *MessagePrototype) GetSubstitutionsOk() (*[]string, bool)`

GetSubstitutionsOk returns a tuple with the Substitutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstitutions

`func (o *MessagePrototype) SetSubstitutions(v []string)`

SetSubstitutions sets Substitutions field to given value.

### HasSubstitutions

`func (o *MessagePrototype) HasSubstitutions() bool`

HasSubstitutions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


