# PageSponsorView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]SponsorView**](SponsorView.md) |  | [optional] 
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

### NewPageSponsorView

`func NewPageSponsorView() *PageSponsorView`

NewPageSponsorView instantiates a new PageSponsorView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageSponsorViewWithDefaults

`func NewPageSponsorViewWithDefaults() *PageSponsorView`

NewPageSponsorViewWithDefaults instantiates a new PageSponsorView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PageSponsorView) GetContent() []SponsorView`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PageSponsorView) GetContentOk() (*[]SponsorView, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PageSponsorView) SetContent(v []SponsorView)`

SetContent sets Content field to given value.

### HasContent

`func (o *PageSponsorView) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEmpty

`func (o *PageSponsorView) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *PageSponsorView) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *PageSponsorView) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *PageSponsorView) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.

### GetFirst

`func (o *PageSponsorView) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PageSponsorView) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PageSponsorView) SetFirst(v bool)`

SetFirst sets First field to given value.

### HasFirst

`func (o *PageSponsorView) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *PageSponsorView) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PageSponsorView) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PageSponsorView) SetLast(v bool)`

SetLast sets Last field to given value.

### HasLast

`func (o *PageSponsorView) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNumber

`func (o *PageSponsorView) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PageSponsorView) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PageSponsorView) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PageSponsorView) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetNumberOfElements

`func (o *PageSponsorView) GetNumberOfElements() int32`

GetNumberOfElements returns the NumberOfElements field if non-nil, zero value otherwise.

### GetNumberOfElementsOk

`func (o *PageSponsorView) GetNumberOfElementsOk() (*int32, bool)`

GetNumberOfElementsOk returns a tuple with the NumberOfElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfElements

`func (o *PageSponsorView) SetNumberOfElements(v int32)`

SetNumberOfElements sets NumberOfElements field to given value.

### HasNumberOfElements

`func (o *PageSponsorView) HasNumberOfElements() bool`

HasNumberOfElements returns a boolean if a field has been set.

### GetSize

`func (o *PageSponsorView) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageSponsorView) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageSponsorView) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PageSponsorView) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSort

`func (o *PageSponsorView) GetSort() Sort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PageSponsorView) GetSortOk() (*Sort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PageSponsorView) SetSort(v Sort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PageSponsorView) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTotalElements

`func (o *PageSponsorView) GetTotalElements() int64`

GetTotalElements returns the TotalElements field if non-nil, zero value otherwise.

### GetTotalElementsOk

`func (o *PageSponsorView) GetTotalElementsOk() (*int64, bool)`

GetTotalElementsOk returns a tuple with the TotalElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElements

`func (o *PageSponsorView) SetTotalElements(v int64)`

SetTotalElements sets TotalElements field to given value.

### HasTotalElements

`func (o *PageSponsorView) HasTotalElements() bool`

HasTotalElements returns a boolean if a field has been set.

### GetTotalPages

`func (o *PageSponsorView) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PageSponsorView) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PageSponsorView) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PageSponsorView) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


