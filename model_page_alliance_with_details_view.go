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

// checks if the PageAllianceWithDetailsView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageAllianceWithDetailsView{}

// PageAllianceWithDetailsView struct for PageAllianceWithDetailsView
type PageAllianceWithDetailsView struct {
	Content []AllianceWithDetailsView `json:"content,omitempty"`
	Empty *bool `json:"empty,omitempty"`
	First *bool `json:"first,omitempty"`
	Last *bool `json:"last,omitempty"`
	Number *int32 `json:"number,omitempty"`
	NumberOfElements *int32 `json:"numberOfElements,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Sort *Sort `json:"sort,omitempty"`
	TotalElements *int64 `json:"totalElements,omitempty"`
	TotalPages *int32 `json:"totalPages,omitempty"`
}

// NewPageAllianceWithDetailsView instantiates a new PageAllianceWithDetailsView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageAllianceWithDetailsView() *PageAllianceWithDetailsView {
	this := PageAllianceWithDetailsView{}
	return &this
}

// NewPageAllianceWithDetailsViewWithDefaults instantiates a new PageAllianceWithDetailsView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageAllianceWithDetailsViewWithDefaults() *PageAllianceWithDetailsView {
	this := PageAllianceWithDetailsView{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *PageAllianceWithDetailsView) GetContent() []AllianceWithDetailsView {
	if o == nil || IsNil(o.Content) {
		var ret []AllianceWithDetailsView
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageAllianceWithDetailsView) GetContentOk() ([]AllianceWithDetailsView, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *PageAllianceWithDetailsView) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given []AllianceWithDetailsView and assigns it to the Content field.
func (o *PageAllianceWithDetailsView) SetContent(v []AllianceWithDetailsView) {
	o.Content = v
}

// GetEmpty returns the Empty field value if set, zero value otherwise.
func (o *PageAllianceWithDetailsView) GetEmpty() bool {
	if o == nil || IsNil(o.Empty) {
		var ret bool
		return ret
	}
	return *o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageAllianceWithDetailsView) GetEmptyOk() (*bool, bool) {
	if o == nil || IsNil(o.Empty) {
		return nil, false
	}
	return o.Empty, true
}

// HasEmpty returns a boolean if a field has been set.
func (o *PageAllianceWithDetailsView) HasEmpty() bool {
	if o != nil && !IsNil(o.Empty) {
		return true
	}

	return false
}

// SetEmpty gets a reference to the given bool and assigns it to the Empty field.
func (o *PageAllianceWithDetailsView) SetEmpty(v bool) {
	o.Empty = &v
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *PageAllianceWithDetailsView) GetFirst() bool {
	if o == nil || IsNil(o.First) {
		var ret bool
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageAllianceWithDetailsView) GetFirstOk() (*bool, bool) {
	if o == nil || IsNil(o.First) {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *PageAllianceWithDetailsView) HasFirst() bool {
	if o != nil && !IsNil(o.First) {
		return true
	}

	return false
}

// SetFirst gets a reference to the given bool and assigns it to the First field.
func (o *PageAllianceWithDetailsView) SetFirst(v bool) {
	o.First = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *PageAllianceWithDetailsView) GetLast() bool {
	if o == nil || IsNil(o.Last) {
		var ret bool
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageAllianceWithDetailsView) GetLastOk() (*bool, bool) {
	if o == nil || IsNil(o.Last) {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *PageAllianceWithDetailsView) HasLast() bool {
	if o != nil && !IsNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given bool and assigns it to the Last field.
func (o *PageAllianceWithDetailsView) SetLast(v bool) {
	o.Last = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *PageAllianceWithDetailsView) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageAllianceWithDetailsView) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *PageAllianceWithDetailsView) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *PageAllianceWithDetailsView) SetNumber(v int32) {
	o.Number = &v
}

// GetNumberOfElements returns the NumberOfElements field value if set, zero value otherwise.
func (o *PageAllianceWithDetailsView) GetNumberOfElements() int32 {
	if o == nil || IsNil(o.NumberOfElements) {
		var ret int32
		return ret
	}
	return *o.NumberOfElements
}

// GetNumberOfElementsOk returns a tuple with the NumberOfElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageAllianceWithDetailsView) GetNumberOfElementsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfElements) {
		return nil, false
	}
	return o.NumberOfElements, true
}

// HasNumberOfElements returns a boolean if a field has been set.
func (o *PageAllianceWithDetailsView) HasNumberOfElements() bool {
	if o != nil && !IsNil(o.NumberOfElements) {
		return true
	}

	return false
}

// SetNumberOfElements gets a reference to the given int32 and assigns it to the NumberOfElements field.
func (o *PageAllianceWithDetailsView) SetNumberOfElements(v int32) {
	o.NumberOfElements = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *PageAllianceWithDetailsView) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageAllianceWithDetailsView) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *PageAllianceWithDetailsView) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *PageAllianceWithDetailsView) SetSize(v int32) {
	o.Size = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *PageAllianceWithDetailsView) GetSort() Sort {
	if o == nil || IsNil(o.Sort) {
		var ret Sort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageAllianceWithDetailsView) GetSortOk() (*Sort, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *PageAllianceWithDetailsView) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given Sort and assigns it to the Sort field.
func (o *PageAllianceWithDetailsView) SetSort(v Sort) {
	o.Sort = &v
}

// GetTotalElements returns the TotalElements field value if set, zero value otherwise.
func (o *PageAllianceWithDetailsView) GetTotalElements() int64 {
	if o == nil || IsNil(o.TotalElements) {
		var ret int64
		return ret
	}
	return *o.TotalElements
}

// GetTotalElementsOk returns a tuple with the TotalElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageAllianceWithDetailsView) GetTotalElementsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalElements) {
		return nil, false
	}
	return o.TotalElements, true
}

// HasTotalElements returns a boolean if a field has been set.
func (o *PageAllianceWithDetailsView) HasTotalElements() bool {
	if o != nil && !IsNil(o.TotalElements) {
		return true
	}

	return false
}

// SetTotalElements gets a reference to the given int64 and assigns it to the TotalElements field.
func (o *PageAllianceWithDetailsView) SetTotalElements(v int64) {
	o.TotalElements = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *PageAllianceWithDetailsView) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageAllianceWithDetailsView) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *PageAllianceWithDetailsView) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *PageAllianceWithDetailsView) SetTotalPages(v int32) {
	o.TotalPages = &v
}

func (o PageAllianceWithDetailsView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageAllianceWithDetailsView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Empty) {
		toSerialize["empty"] = o.Empty
	}
	if !IsNil(o.First) {
		toSerialize["first"] = o.First
	}
	if !IsNil(o.Last) {
		toSerialize["last"] = o.Last
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.NumberOfElements) {
		toSerialize["numberOfElements"] = o.NumberOfElements
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !IsNil(o.TotalElements) {
		toSerialize["totalElements"] = o.TotalElements
	}
	if !IsNil(o.TotalPages) {
		toSerialize["totalPages"] = o.TotalPages
	}
	return toSerialize, nil
}

type NullablePageAllianceWithDetailsView struct {
	value *PageAllianceWithDetailsView
	isSet bool
}

func (v NullablePageAllianceWithDetailsView) Get() *PageAllianceWithDetailsView {
	return v.value
}

func (v *NullablePageAllianceWithDetailsView) Set(val *PageAllianceWithDetailsView) {
	v.value = val
	v.isSet = true
}

func (v NullablePageAllianceWithDetailsView) IsSet() bool {
	return v.isSet
}

func (v *NullablePageAllianceWithDetailsView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageAllianceWithDetailsView(val *PageAllianceWithDetailsView) *NullablePageAllianceWithDetailsView {
	return &NullablePageAllianceWithDetailsView{value: val, isSet: true}
}

func (v NullablePageAllianceWithDetailsView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageAllianceWithDetailsView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


