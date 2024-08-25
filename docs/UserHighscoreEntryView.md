# UserHighscoreEntryView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **time.Time** |  | [optional] 
**HistoryDate** | Pointer to **int64** |  | [optional] 
**HistoryPosition** | Pointer to **int64** |  | [optional] 
**HistoryValue** | Pointer to **float32** |  | [optional] 
**User** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 

## Methods

### NewUserHighscoreEntryView

`func NewUserHighscoreEntryView() *UserHighscoreEntryView`

NewUserHighscoreEntryView instantiates a new UserHighscoreEntryView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserHighscoreEntryViewWithDefaults

`func NewUserHighscoreEntryViewWithDefaults() *UserHighscoreEntryView`

NewUserHighscoreEntryViewWithDefaults instantiates a new UserHighscoreEntryView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *UserHighscoreEntryView) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *UserHighscoreEntryView) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *UserHighscoreEntryView) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *UserHighscoreEntryView) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetHistoryDate

`func (o *UserHighscoreEntryView) GetHistoryDate() int64`

GetHistoryDate returns the HistoryDate field if non-nil, zero value otherwise.

### GetHistoryDateOk

`func (o *UserHighscoreEntryView) GetHistoryDateOk() (*int64, bool)`

GetHistoryDateOk returns a tuple with the HistoryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryDate

`func (o *UserHighscoreEntryView) SetHistoryDate(v int64)`

SetHistoryDate sets HistoryDate field to given value.

### HasHistoryDate

`func (o *UserHighscoreEntryView) HasHistoryDate() bool`

HasHistoryDate returns a boolean if a field has been set.

### GetHistoryPosition

`func (o *UserHighscoreEntryView) GetHistoryPosition() int64`

GetHistoryPosition returns the HistoryPosition field if non-nil, zero value otherwise.

### GetHistoryPositionOk

`func (o *UserHighscoreEntryView) GetHistoryPositionOk() (*int64, bool)`

GetHistoryPositionOk returns a tuple with the HistoryPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryPosition

`func (o *UserHighscoreEntryView) SetHistoryPosition(v int64)`

SetHistoryPosition sets HistoryPosition field to given value.

### HasHistoryPosition

`func (o *UserHighscoreEntryView) HasHistoryPosition() bool`

HasHistoryPosition returns a boolean if a field has been set.

### GetHistoryValue

`func (o *UserHighscoreEntryView) GetHistoryValue() float32`

GetHistoryValue returns the HistoryValue field if non-nil, zero value otherwise.

### GetHistoryValueOk

`func (o *UserHighscoreEntryView) GetHistoryValueOk() (*float32, bool)`

GetHistoryValueOk returns a tuple with the HistoryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryValue

`func (o *UserHighscoreEntryView) SetHistoryValue(v float32)`

SetHistoryValue sets HistoryValue field to given value.

### HasHistoryValue

`func (o *UserHighscoreEntryView) HasHistoryValue() bool`

HasHistoryValue returns a boolean if a field has been set.

### GetUser

`func (o *UserHighscoreEntryView) GetUser() UsernameView`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserHighscoreEntryView) GetUserOk() (*UsernameView, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserHighscoreEntryView) SetUser(v UsernameView)`

SetUser sets User field to given value.

### HasUser

`func (o *UserHighscoreEntryView) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetValue

`func (o *UserHighscoreEntryView) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UserHighscoreEntryView) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UserHighscoreEntryView) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *UserHighscoreEntryView) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


