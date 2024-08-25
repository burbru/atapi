# HistoryEntryView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**MessagePrototype**](MessagePrototype.md) |  | [optional] 
**Date** | Pointer to **int64** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewHistoryEntryView

`func NewHistoryEntryView() *HistoryEntryView`

NewHistoryEntryView instantiates a new HistoryEntryView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryEntryViewWithDefaults

`func NewHistoryEntryViewWithDefaults() *HistoryEntryView`

NewHistoryEntryViewWithDefaults instantiates a new HistoryEntryView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *HistoryEntryView) GetContent() MessagePrototype`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *HistoryEntryView) GetContentOk() (*MessagePrototype, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *HistoryEntryView) SetContent(v MessagePrototype)`

SetContent sets Content field to given value.

### HasContent

`func (o *HistoryEntryView) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDate

`func (o *HistoryEntryView) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *HistoryEntryView) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *HistoryEntryView) SetDate(v int64)`

SetDate sets Date field to given value.

### HasDate

`func (o *HistoryEntryView) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetEntityId

`func (o *HistoryEntryView) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *HistoryEntryView) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *HistoryEntryView) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *HistoryEntryView) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetId

`func (o *HistoryEntryView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryEntryView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryEntryView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HistoryEntryView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *HistoryEntryView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HistoryEntryView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HistoryEntryView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HistoryEntryView) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


