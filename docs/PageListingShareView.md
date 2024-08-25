# PageListingShareView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]ListingShareView**](ListingShareView.md) |  | [optional] 
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

### NewPageListingShareView

`func NewPageListingShareView() *PageListingShareView`

NewPageListingShareView instantiates a new PageListingShareView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageListingShareViewWithDefaults

`func NewPageListingShareViewWithDefaults() *PageListingShareView`

NewPageListingShareViewWithDefaults instantiates a new PageListingShareView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PageListingShareView) GetContent() []ListingShareView`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PageListingShareView) GetContentOk() (*[]ListingShareView, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PageListingShareView) SetContent(v []ListingShareView)`

SetContent sets Content field to given value.

### HasContent

`func (o *PageListingShareView) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEmpty

`func (o *PageListingShareView) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *PageListingShareView) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *PageListingShareView) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *PageListingShareView) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.

### GetFirst

`func (o *PageListingShareView) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PageListingShareView) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PageListingShareView) SetFirst(v bool)`

SetFirst sets First field to given value.

### HasFirst

`func (o *PageListingShareView) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *PageListingShareView) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PageListingShareView) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PageListingShareView) SetLast(v bool)`

SetLast sets Last field to given value.

### HasLast

`func (o *PageListingShareView) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNumber

`func (o *PageListingShareView) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PageListingShareView) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PageListingShareView) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PageListingShareView) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetNumberOfElements

`func (o *PageListingShareView) GetNumberOfElements() int32`

GetNumberOfElements returns the NumberOfElements field if non-nil, zero value otherwise.

### GetNumberOfElementsOk

`func (o *PageListingShareView) GetNumberOfElementsOk() (*int32, bool)`

GetNumberOfElementsOk returns a tuple with the NumberOfElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfElements

`func (o *PageListingShareView) SetNumberOfElements(v int32)`

SetNumberOfElements sets NumberOfElements field to given value.

### HasNumberOfElements

`func (o *PageListingShareView) HasNumberOfElements() bool`

HasNumberOfElements returns a boolean if a field has been set.

### GetSize

`func (o *PageListingShareView) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageListingShareView) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageListingShareView) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PageListingShareView) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSort

`func (o *PageListingShareView) GetSort() Sort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PageListingShareView) GetSortOk() (*Sort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PageListingShareView) SetSort(v Sort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PageListingShareView) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTotalElements

`func (o *PageListingShareView) GetTotalElements() int64`

GetTotalElements returns the TotalElements field if non-nil, zero value otherwise.

### GetTotalElementsOk

`func (o *PageListingShareView) GetTotalElementsOk() (*int64, bool)`

GetTotalElementsOk returns a tuple with the TotalElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElements

`func (o *PageListingShareView) SetTotalElements(v int64)`

SetTotalElements sets TotalElements field to given value.

### HasTotalElements

`func (o *PageListingShareView) HasTotalElements() bool`

HasTotalElements returns a boolean if a field has been set.

### GetTotalPages

`func (o *PageListingShareView) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PageListingShareView) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PageListingShareView) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PageListingShareView) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


