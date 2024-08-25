# IndexMemberValuesView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCapitalisation** | Pointer to **float32** |  | [optional] 
**BasePrice** | Pointer to **float32** |  | [optional] 
**BaseShares** | Pointer to **int64** |  | [optional] 
**Capitalisation** | Pointer to **float32** |  | [optional] 
**LastAdjustmentDate** | Pointer to **int64** |  | [optional] 
**Listing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**PriceAdjustmentFactor** | Pointer to **float32** |  | [optional] 
**Shares** | Pointer to **int64** |  | [optional] 

## Methods

### NewIndexMemberValuesView

`func NewIndexMemberValuesView() *IndexMemberValuesView`

NewIndexMemberValuesView instantiates a new IndexMemberValuesView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexMemberValuesViewWithDefaults

`func NewIndexMemberValuesViewWithDefaults() *IndexMemberValuesView`

NewIndexMemberValuesViewWithDefaults instantiates a new IndexMemberValuesView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCapitalisation

`func (o *IndexMemberValuesView) GetBaseCapitalisation() float32`

GetBaseCapitalisation returns the BaseCapitalisation field if non-nil, zero value otherwise.

### GetBaseCapitalisationOk

`func (o *IndexMemberValuesView) GetBaseCapitalisationOk() (*float32, bool)`

GetBaseCapitalisationOk returns a tuple with the BaseCapitalisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCapitalisation

`func (o *IndexMemberValuesView) SetBaseCapitalisation(v float32)`

SetBaseCapitalisation sets BaseCapitalisation field to given value.

### HasBaseCapitalisation

`func (o *IndexMemberValuesView) HasBaseCapitalisation() bool`

HasBaseCapitalisation returns a boolean if a field has been set.

### GetBasePrice

`func (o *IndexMemberValuesView) GetBasePrice() float32`

GetBasePrice returns the BasePrice field if non-nil, zero value otherwise.

### GetBasePriceOk

`func (o *IndexMemberValuesView) GetBasePriceOk() (*float32, bool)`

GetBasePriceOk returns a tuple with the BasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePrice

`func (o *IndexMemberValuesView) SetBasePrice(v float32)`

SetBasePrice sets BasePrice field to given value.

### HasBasePrice

`func (o *IndexMemberValuesView) HasBasePrice() bool`

HasBasePrice returns a boolean if a field has been set.

### GetBaseShares

`func (o *IndexMemberValuesView) GetBaseShares() int64`

GetBaseShares returns the BaseShares field if non-nil, zero value otherwise.

### GetBaseSharesOk

`func (o *IndexMemberValuesView) GetBaseSharesOk() (*int64, bool)`

GetBaseSharesOk returns a tuple with the BaseShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseShares

`func (o *IndexMemberValuesView) SetBaseShares(v int64)`

SetBaseShares sets BaseShares field to given value.

### HasBaseShares

`func (o *IndexMemberValuesView) HasBaseShares() bool`

HasBaseShares returns a boolean if a field has been set.

### GetCapitalisation

`func (o *IndexMemberValuesView) GetCapitalisation() float32`

GetCapitalisation returns the Capitalisation field if non-nil, zero value otherwise.

### GetCapitalisationOk

`func (o *IndexMemberValuesView) GetCapitalisationOk() (*float32, bool)`

GetCapitalisationOk returns a tuple with the Capitalisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalisation

`func (o *IndexMemberValuesView) SetCapitalisation(v float32)`

SetCapitalisation sets Capitalisation field to given value.

### HasCapitalisation

`func (o *IndexMemberValuesView) HasCapitalisation() bool`

HasCapitalisation returns a boolean if a field has been set.

### GetLastAdjustmentDate

`func (o *IndexMemberValuesView) GetLastAdjustmentDate() int64`

GetLastAdjustmentDate returns the LastAdjustmentDate field if non-nil, zero value otherwise.

### GetLastAdjustmentDateOk

`func (o *IndexMemberValuesView) GetLastAdjustmentDateOk() (*int64, bool)`

GetLastAdjustmentDateOk returns a tuple with the LastAdjustmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAdjustmentDate

`func (o *IndexMemberValuesView) SetLastAdjustmentDate(v int64)`

SetLastAdjustmentDate sets LastAdjustmentDate field to given value.

### HasLastAdjustmentDate

`func (o *IndexMemberValuesView) HasLastAdjustmentDate() bool`

HasLastAdjustmentDate returns a boolean if a field has been set.

### GetListing

`func (o *IndexMemberValuesView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *IndexMemberValuesView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *IndexMemberValuesView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *IndexMemberValuesView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetPrice

`func (o *IndexMemberValuesView) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *IndexMemberValuesView) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *IndexMemberValuesView) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *IndexMemberValuesView) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceAdjustmentFactor

`func (o *IndexMemberValuesView) GetPriceAdjustmentFactor() float32`

GetPriceAdjustmentFactor returns the PriceAdjustmentFactor field if non-nil, zero value otherwise.

### GetPriceAdjustmentFactorOk

`func (o *IndexMemberValuesView) GetPriceAdjustmentFactorOk() (*float32, bool)`

GetPriceAdjustmentFactorOk returns a tuple with the PriceAdjustmentFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAdjustmentFactor

`func (o *IndexMemberValuesView) SetPriceAdjustmentFactor(v float32)`

SetPriceAdjustmentFactor sets PriceAdjustmentFactor field to given value.

### HasPriceAdjustmentFactor

`func (o *IndexMemberValuesView) HasPriceAdjustmentFactor() bool`

HasPriceAdjustmentFactor returns a boolean if a field has been set.

### GetShares

`func (o *IndexMemberValuesView) GetShares() int64`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *IndexMemberValuesView) GetSharesOk() (*int64, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *IndexMemberValuesView) SetShares(v int64)`

SetShares sets Shares field to given value.

### HasShares

`func (o *IndexMemberValuesView) HasShares() bool`

HasShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


