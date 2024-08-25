# PageNotificationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]NotificationView**](NotificationView.md) |  | [optional] 
**Empty** | Pointer to **bool** |  | [optional] 
**First** | Pointer to **bool** |  | [optional] 
**Last** | Pointer to **bool** |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**NumberOfElements** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Sort** | Pointer to [**Sort**](Sort.md) |  | [optional] 
**TotalElements** | Pointer to **int64** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] 

## Methods

### NewPageNotificationView

`func NewPageNotificationView() *PageNotificationView`

NewPageNotificationView instantiates a new PageNotificationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageNotificationViewWithDefaults

`func NewPageNotificationViewWithDefaults() *PageNotificationView`

NewPageNotificationViewWithDefaults instantiates a new PageNotificationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PageNotificationView) GetContent() []NotificationView`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PageNotificationView) GetContentOk() (*[]NotificationView, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PageNotificationView) SetContent(v []NotificationView)`

SetContent sets Content field to given value.

### HasContent

`func (o *PageNotificationView) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEmpty

`func (o *PageNotificationView) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *PageNotificationView) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *PageNotificationView) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *PageNotificationView) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.

### GetFirst

`func (o *PageNotificationView) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PageNotificationView) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PageNotificationView) SetFirst(v bool)`

SetFirst sets First field to given value.

### HasFirst

`func (o *PageNotificationView) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *PageNotificationView) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PageNotificationView) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PageNotificationView) SetLast(v bool)`

SetLast sets Last field to given value.

### HasLast

`func (o *PageNotificationView) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNumber

`func (o *PageNotificationView) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PageNotificationView) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PageNotificationView) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PageNotificationView) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetNumberOfElements

`func (o *PageNotificationView) GetNumberOfElements() int32`

GetNumberOfElements returns the NumberOfElements field if non-nil, zero value otherwise.

### GetNumberOfElementsOk

`func (o *PageNotificationView) GetNumberOfElementsOk() (*int32, bool)`

GetNumberOfElementsOk returns a tuple with the NumberOfElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfElements

`func (o *PageNotificationView) SetNumberOfElements(v int32)`

SetNumberOfElements sets NumberOfElements field to given value.

### HasNumberOfElements

`func (o *PageNotificationView) HasNumberOfElements() bool`

HasNumberOfElements returns a boolean if a field has been set.

### GetSize

`func (o *PageNotificationView) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageNotificationView) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageNotificationView) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PageNotificationView) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSort

`func (o *PageNotificationView) GetSort() Sort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PageNotificationView) GetSortOk() (*Sort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PageNotificationView) SetSort(v Sort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PageNotificationView) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTotalElements

`func (o *PageNotificationView) GetTotalElements() int64`

GetTotalElements returns the TotalElements field if non-nil, zero value otherwise.

### GetTotalElementsOk

`func (o *PageNotificationView) GetTotalElementsOk() (*int64, bool)`

GetTotalElementsOk returns a tuple with the TotalElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElements

`func (o *PageNotificationView) SetTotalElements(v int64)`

SetTotalElements sets TotalElements field to given value.

### HasTotalElements

`func (o *PageNotificationView) HasTotalElements() bool`

HasTotalElements returns a boolean if a field has been set.

### GetTotalPages

`func (o *PageNotificationView) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PageNotificationView) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PageNotificationView) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PageNotificationView) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


