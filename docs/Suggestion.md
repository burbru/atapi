# Suggestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionData** | Pointer to **map[string]interface{}** |  | [optional] 
**Text** | Pointer to [**MessagePrototype**](MessagePrototype.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 

## Methods

### NewSuggestion

`func NewSuggestion() *Suggestion`

NewSuggestion instantiates a new Suggestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestionWithDefaults

`func NewSuggestionWithDefaults() *Suggestion`

NewSuggestionWithDefaults instantiates a new Suggestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionData

`func (o *Suggestion) GetActionData() map[string]interface{}`

GetActionData returns the ActionData field if non-nil, zero value otherwise.

### GetActionDataOk

`func (o *Suggestion) GetActionDataOk() (*map[string]interface{}, bool)`

GetActionDataOk returns a tuple with the ActionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionData

`func (o *Suggestion) SetActionData(v map[string]interface{})`

SetActionData sets ActionData field to given value.

### HasActionData

`func (o *Suggestion) HasActionData() bool`

HasActionData returns a boolean if a field has been set.

### GetText

`func (o *Suggestion) GetText() MessagePrototype`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Suggestion) GetTextOk() (*MessagePrototype, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Suggestion) SetText(v MessagePrototype)`

SetText sets Text field to given value.

### HasText

`func (o *Suggestion) HasText() bool`

HasText returns a boolean if a field has been set.

### GetType

`func (o *Suggestion) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Suggestion) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Suggestion) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Suggestion) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *Suggestion) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Suggestion) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Suggestion) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Suggestion) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


