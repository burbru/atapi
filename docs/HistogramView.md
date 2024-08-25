# HistogramView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HighlightRange** | Pointer to [**RangeViewbigdecimal**](RangeViewbigdecimal.md) |  | [optional] 
**HighlightValue** | Pointer to **float32** |  | [optional] 
**Histogram** | Pointer to **map[string]int32** |  | [optional] 

## Methods

### NewHistogramView

`func NewHistogramView() *HistogramView`

NewHistogramView instantiates a new HistogramView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistogramViewWithDefaults

`func NewHistogramViewWithDefaults() *HistogramView`

NewHistogramViewWithDefaults instantiates a new HistogramView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHighlightRange

`func (o *HistogramView) GetHighlightRange() RangeViewbigdecimal`

GetHighlightRange returns the HighlightRange field if non-nil, zero value otherwise.

### GetHighlightRangeOk

`func (o *HistogramView) GetHighlightRangeOk() (*RangeViewbigdecimal, bool)`

GetHighlightRangeOk returns a tuple with the HighlightRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightRange

`func (o *HistogramView) SetHighlightRange(v RangeViewbigdecimal)`

SetHighlightRange sets HighlightRange field to given value.

### HasHighlightRange

`func (o *HistogramView) HasHighlightRange() bool`

HasHighlightRange returns a boolean if a field has been set.

### GetHighlightValue

`func (o *HistogramView) GetHighlightValue() float32`

GetHighlightValue returns the HighlightValue field if non-nil, zero value otherwise.

### GetHighlightValueOk

`func (o *HistogramView) GetHighlightValueOk() (*float32, bool)`

GetHighlightValueOk returns a tuple with the HighlightValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightValue

`func (o *HistogramView) SetHighlightValue(v float32)`

SetHighlightValue sets HighlightValue field to given value.

### HasHighlightValue

`func (o *HistogramView) HasHighlightValue() bool`

HasHighlightValue returns a boolean if a field has been set.

### GetHistogram

`func (o *HistogramView) GetHistogram() map[string]int32`

GetHistogram returns the Histogram field if non-nil, zero value otherwise.

### GetHistogramOk

`func (o *HistogramView) GetHistogramOk() (*map[string]int32, bool)`

GetHistogramOk returns a tuple with the Histogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogram

`func (o *HistogramView) SetHistogram(v map[string]int32)`

SetHistogram sets Histogram field to given value.

### HasHistogram

`func (o *HistogramView) HasHistogram() bool`

HasHistogram returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


