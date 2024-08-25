# PageHistorizedListingDataView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]HistorizedListingDataView**](HistorizedListingDataView.md) |  | [optional] 
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

### NewPageHistorizedListingDataView

`func NewPageHistorizedListingDataView() *PageHistorizedListingDataView`

NewPageHistorizedListingDataView instantiates a new PageHistorizedListingDataView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageHistorizedListingDataViewWithDefaults

`func NewPageHistorizedListingDataViewWithDefaults() *PageHistorizedListingDataView`

NewPageHistorizedListingDataViewWithDefaults instantiates a new PageHistorizedListingDataView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PageHistorizedListingDataView) GetContent() []HistorizedListingDataView`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PageHistorizedListingDataView) GetContentOk() (*[]HistorizedListingDataView, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PageHistorizedListingDataView) SetContent(v []HistorizedListingDataView)`

SetContent sets Content field to given value.

### HasContent

`func (o *PageHistorizedListingDataView) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEmpty

`func (o *PageHistorizedListingDataView) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *PageHistorizedListingDataView) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *PageHistorizedListingDataView) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *PageHistorizedListingDataView) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.

### GetFirst

`func (o *PageHistorizedListingDataView) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PageHistorizedListingDataView) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PageHistorizedListingDataView) SetFirst(v bool)`

SetFirst sets First field to given value.

### HasFirst

`func (o *PageHistorizedListingDataView) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *PageHistorizedListingDataView) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PageHistorizedListingDataView) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PageHistorizedListingDataView) SetLast(v bool)`

SetLast sets Last field to given value.

### HasLast

`func (o *PageHistorizedListingDataView) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNumber

`func (o *PageHistorizedListingDataView) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PageHistorizedListingDataView) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PageHistorizedListingDataView) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PageHistorizedListingDataView) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetNumberOfElements

`func (o *PageHistorizedListingDataView) GetNumberOfElements() int32`

GetNumberOfElements returns the NumberOfElements field if non-nil, zero value otherwise.

### GetNumberOfElementsOk

`func (o *PageHistorizedListingDataView) GetNumberOfElementsOk() (*int32, bool)`

GetNumberOfElementsOk returns a tuple with the NumberOfElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfElements

`func (o *PageHistorizedListingDataView) SetNumberOfElements(v int32)`

SetNumberOfElements sets NumberOfElements field to given value.

### HasNumberOfElements

`func (o *PageHistorizedListingDataView) HasNumberOfElements() bool`

HasNumberOfElements returns a boolean if a field has been set.

### GetSize

`func (o *PageHistorizedListingDataView) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageHistorizedListingDataView) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageHistorizedListingDataView) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PageHistorizedListingDataView) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSort

`func (o *PageHistorizedListingDataView) GetSort() Sort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PageHistorizedListingDataView) GetSortOk() (*Sort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PageHistorizedListingDataView) SetSort(v Sort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PageHistorizedListingDataView) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTotalElements

`func (o *PageHistorizedListingDataView) GetTotalElements() int64`

GetTotalElements returns the TotalElements field if non-nil, zero value otherwise.

### GetTotalElementsOk

`func (o *PageHistorizedListingDataView) GetTotalElementsOk() (*int64, bool)`

GetTotalElementsOk returns a tuple with the TotalElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElements

`func (o *PageHistorizedListingDataView) SetTotalElements(v int64)`

SetTotalElements sets TotalElements field to given value.

### HasTotalElements

`func (o *PageHistorizedListingDataView) HasTotalElements() bool`

HasTotalElements returns a boolean if a field has been set.

### GetTotalPages

`func (o *PageHistorizedListingDataView) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PageHistorizedListingDataView) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PageHistorizedListingDataView) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PageHistorizedListingDataView) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


