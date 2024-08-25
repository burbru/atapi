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

// checks if the PageCapitalIncreaseView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageCapitalIncreaseView{}

// PageCapitalIncreaseView struct for PageCapitalIncreaseView
type PageCapitalIncreaseView struct {
	Content []CapitalIncreaseView `json:"content,omitempty"`
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

// NewPageCapitalIncreaseView instantiates a new PageCapitalIncreaseView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageCapitalIncreaseView() *PageCapitalIncreaseView {
	this := PageCapitalIncreaseView{}
	return &this
}

// NewPageCapitalIncreaseViewWithDefaults instantiates a new PageCapitalIncreaseView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageCapitalIncreaseViewWithDefaults() *PageCapitalIncreaseView {
	this := PageCapitalIncreaseView{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *PageCapitalIncreaseView) GetContent() []CapitalIncreaseView {
	if o == nil || IsNil(o.Content) {
		var ret []CapitalIncreaseView
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageCapitalIncreaseView) GetContentOk() ([]CapitalIncreaseView, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *PageCapitalIncreaseView) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given []CapitalIncreaseView and assigns it to the Content field.
func (o *PageCapitalIncreaseView) SetContent(v []CapitalIncreaseView) {
	o.Content = v
}

// GetEmpty returns the Empty field value if set, zero value otherwise.
func (o *PageCapitalIncreaseView) GetEmpty() bool {
	if o == nil || IsNil(o.Empty) {
		var ret bool
		return ret
	}
	return *o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageCapitalIncreaseView) GetEmptyOk() (*bool, bool) {
	if o == nil || IsNil(o.Empty) {
		return nil, false
	}
	return o.Empty, true
}

// HasEmpty returns a boolean if a field has been set.
func (o *PageCapitalIncreaseView) HasEmpty() bool {
	if o != nil && !IsNil(o.Empty) {
		return true
	}

	return false
}

// SetEmpty gets a reference to the given bool and assigns it to the Empty field.
func (o *PageCapitalIncreaseView) SetEmpty(v bool) {
	o.Empty = &v
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *PageCapitalIncreaseView) GetFirst() bool {
	if o == nil || IsNil(o.First) {
		var ret bool
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageCapitalIncreaseView) GetFirstOk() (*bool, bool) {
	if o == nil || IsNil(o.First) {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *PageCapitalIncreaseView) HasFirst() bool {
	if o != nil && !IsNil(o.First) {
		return true
	}

	return false
}

// SetFirst gets a reference to the given bool and assigns it to the First field.
func (o *PageCapitalIncreaseView) SetFirst(v bool) {
	o.First = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *PageCapitalIncreaseView) GetLast() bool {
	if o == nil || IsNil(o.Last) {
		var ret bool
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageCapitalIncreaseView) GetLastOk() (*bool, bool) {
	if o == nil || IsNil(o.Last) {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *PageCapitalIncreaseView) HasLast() bool {
	if o != nil && !IsNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given bool and assigns it to the Last field.
func (o *PageCapitalIncreaseView) SetLast(v bool) {
	o.Last = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *PageCapitalIncreaseView) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageCapitalIncreaseView) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *PageCapitalIncreaseView) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *PageCapitalIncreaseView) SetNumber(v int32) {
	o.Number = &v
}

// GetNumberOfElements returns the NumberOfElements field value if set, zero value otherwise.
func (o *PageCapitalIncreaseView) GetNumberOfElements() int32 {
	if o == nil || IsNil(o.NumberOfElements) {
		var ret int32
		return ret
	}
	return *o.NumberOfElements
}

// GetNumberOfElementsOk returns a tuple with the NumberOfElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageCapitalIncreaseView) GetNumberOfElementsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfElements) {
		return nil, false
	}
	return o.NumberOfElements, true
}

// HasNumberOfElements returns a boolean if a field has been set.
func (o *PageCapitalIncreaseView) HasNumberOfElements() bool {
	if o != nil && !IsNil(o.NumberOfElements) {
		return true
	}

	return false
}

// SetNumberOfElements gets a reference to the given int32 and assigns it to the NumberOfElements field.
func (o *PageCapitalIncreaseView) SetNumberOfElements(v int32) {
	o.NumberOfElements = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *PageCapitalIncreaseView) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageCapitalIncreaseView) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *PageCapitalIncreaseView) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *PageCapitalIncreaseView) SetSize(v int32) {
	o.Size = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *PageCapitalIncreaseView) GetSort() Sort {
	if o == nil || IsNil(o.Sort) {
		var ret Sort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageCapitalIncreaseView) GetSortOk() (*Sort, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *PageCapitalIncreaseView) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given Sort and assigns it to the Sort field.
func (o *PageCapitalIncreaseView) SetSort(v Sort) {
	o.Sort = &v
}

// GetTotalElements returns the TotalElements field value if set, zero value otherwise.
func (o *PageCapitalIncreaseView) GetTotalElements() int64 {
	if o == nil || IsNil(o.TotalElements) {
		var ret int64
		return ret
	}
	return *o.TotalElements
}

// GetTotalElementsOk returns a tuple with the TotalElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageCapitalIncreaseView) GetTotalElementsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalElements) {
		return nil, false
	}
	return o.TotalElements, true
}

// HasTotalElements returns a boolean if a field has been set.
func (o *PageCapitalIncreaseView) HasTotalElements() bool {
	if o != nil && !IsNil(o.TotalElements) {
		return true
	}

	return false
}

// SetTotalElements gets a reference to the given int64 and assigns it to the TotalElements field.
func (o *PageCapitalIncreaseView) SetTotalElements(v int64) {
	o.TotalElements = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *PageCapitalIncreaseView) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageCapitalIncreaseView) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *PageCapitalIncreaseView) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *PageCapitalIncreaseView) SetTotalPages(v int32) {
	o.TotalPages = &v
}

func (o PageCapitalIncreaseView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageCapitalIncreaseView) ToMap() (map[string]interface{}, error) {
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

type NullablePageCapitalIncreaseView struct {
	value *PageCapitalIncreaseView
	isSet bool
}

func (v NullablePageCapitalIncreaseView) Get() *PageCapitalIncreaseView {
	return v.value
}

func (v *NullablePageCapitalIncreaseView) Set(val *PageCapitalIncreaseView) {
	v.value = val
	v.isSet = true
}

func (v NullablePageCapitalIncreaseView) IsSet() bool {
	return v.isSet
}

func (v *NullablePageCapitalIncreaseView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageCapitalIncreaseView(val *PageCapitalIncreaseView) *NullablePageCapitalIncreaseView {
	return &NullablePageCapitalIncreaseView{value: val, isSet: true}
}

func (v NullablePageCapitalIncreaseView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageCapitalIncreaseView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


