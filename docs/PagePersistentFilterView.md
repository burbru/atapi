# PagePersistentFilterView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]PersistentFilterView**](PersistentFilterView.md) |  | [optional] 
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

### NewPagePersistentFilterView

`func NewPagePersistentFilterView() *PagePersistentFilterView`

NewPagePersistentFilterView instantiates a new PagePersistentFilterView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagePersistentFilterViewWithDefaults

`func NewPagePersistentFilterViewWithDefaults() *PagePersistentFilterView`

NewPagePersistentFilterViewWithDefaults instantiates a new PagePersistentFilterView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PagePersistentFilterView) GetContent() []PersistentFilterView`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PagePersistentFilterView) GetContentOk() (*[]PersistentFilterView, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PagePersistentFilterView) SetContent(v []PersistentFilterView)`

SetContent sets Content field to given value.

### HasContent

`func (o *PagePersistentFilterView) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEmpty

`func (o *PagePersistentFilterView) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *PagePersistentFilterView) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *PagePersistentFilterView) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *PagePersistentFilterView) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.

### GetFirst

`func (o *PagePersistentFilterView) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PagePersistentFilterView) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PagePersistentFilterView) SetFirst(v bool)`

SetFirst sets First field to given value.

### HasFirst

`func (o *PagePersistentFilterView) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *PagePersistentFilterView) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PagePersistentFilterView) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PagePersistentFilterView) SetLast(v bool)`

SetLast sets Last field to given value.

### HasLast

`func (o *PagePersistentFilterView) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNumber

`func (o *PagePersistentFilterView) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PagePersistentFilterView) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PagePersistentFilterView) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PagePersistentFilterView) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetNumberOfElements

`func (o *PagePersistentFilterView) GetNumberOfElements() int32`

GetNumberOfElements returns the NumberOfElements field if non-nil, zero value otherwise.

### GetNumberOfElementsOk

`func (o *PagePersistentFilterView) GetNumberOfElementsOk() (*int32, bool)`

GetNumberOfElementsOk returns a tuple with the NumberOfElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfElements

`func (o *PagePersistentFilterView) SetNumberOfElements(v int32)`

SetNumberOfElements sets NumberOfElements field to given value.

### HasNumberOfElements

`func (o *PagePersistentFilterView) HasNumberOfElements() bool`

HasNumberOfElements returns a boolean if a field has been set.

### GetSize

`func (o *PagePersistentFilterView) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PagePersistentFilterView) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PagePersistentFilterView) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PagePersistentFilterView) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSort

`func (o *PagePersistentFilterView) GetSort() Sort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PagePersistentFilterView) GetSortOk() (*Sort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PagePersistentFilterView) SetSort(v Sort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PagePersistentFilterView) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTotalElements

`func (o *PagePersistentFilterView) GetTotalElements() int64`

GetTotalElements returns the TotalElements field if non-nil, zero value otherwise.

### GetTotalElementsOk

`func (o *PagePersistentFilterView) GetTotalElementsOk() (*int64, bool)`

GetTotalElementsOk returns a tuple with the TotalElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElements

`func (o *PagePersistentFilterView) SetTotalElements(v int64)`

SetTotalElements sets TotalElements field to given value.

### HasTotalElements

`func (o *PagePersistentFilterView) HasTotalElements() bool`

HasTotalElements returns a boolean if a field has been set.

### GetTotalPages

`func (o *PagePersistentFilterView) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PagePersistentFilterView) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PagePersistentFilterView) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PagePersistentFilterView) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


