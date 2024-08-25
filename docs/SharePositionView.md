# SharePositionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageBuyingPrice** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastPrice** | Pointer to **float32** |  | [optional] 
**LastPriceUpdate** | Pointer to **int64** |  | [optional] 
**Listing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**NumberOfShares** | Pointer to **int64** |  | [optional] 
**SecuritiesAccount** | Pointer to [**SecuritiesAccountCompactView**](SecuritiesAccountCompactView.md) |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewSharePositionView

`func NewSharePositionView() *SharePositionView`

NewSharePositionView instantiates a new SharePositionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharePositionViewWithDefaults

`func NewSharePositionViewWithDefaults() *SharePositionView`

NewSharePositionViewWithDefaults instantiates a new SharePositionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageBuyingPrice

`func (o *SharePositionView) GetAverageBuyingPrice() float32`

GetAverageBuyingPrice returns the AverageBuyingPrice field if non-nil, zero value otherwise.

### GetAverageBuyingPriceOk

`func (o *SharePositionView) GetAverageBuyingPriceOk() (*float32, bool)`

GetAverageBuyingPriceOk returns a tuple with the AverageBuyingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageBuyingPrice

`func (o *SharePositionView) SetAverageBuyingPrice(v float32)`

SetAverageBuyingPrice sets AverageBuyingPrice field to given value.

### HasAverageBuyingPrice

`func (o *SharePositionView) HasAverageBuyingPrice() bool`

HasAverageBuyingPrice returns a boolean if a field has been set.

### GetId

`func (o *SharePositionView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharePositionView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharePositionView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SharePositionView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastPrice

`func (o *SharePositionView) GetLastPrice() float32`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *SharePositionView) GetLastPriceOk() (*float32, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *SharePositionView) SetLastPrice(v float32)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *SharePositionView) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLastPriceUpdate

`func (o *SharePositionView) GetLastPriceUpdate() int64`

GetLastPriceUpdate returns the LastPriceUpdate field if non-nil, zero value otherwise.

### GetLastPriceUpdateOk

`func (o *SharePositionView) GetLastPriceUpdateOk() (*int64, bool)`

GetLastPriceUpdateOk returns a tuple with the LastPriceUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPriceUpdate

`func (o *SharePositionView) SetLastPriceUpdate(v int64)`

SetLastPriceUpdate sets LastPriceUpdate field to given value.

### HasLastPriceUpdate

`func (o *SharePositionView) HasLastPriceUpdate() bool`

HasLastPriceUpdate returns a boolean if a field has been set.

### GetListing

`func (o *SharePositionView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *SharePositionView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *SharePositionView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *SharePositionView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetNumberOfShares

`func (o *SharePositionView) GetNumberOfShares() int64`

GetNumberOfShares returns the NumberOfShares field if non-nil, zero value otherwise.

### GetNumberOfSharesOk

`func (o *SharePositionView) GetNumberOfSharesOk() (*int64, bool)`

GetNumberOfSharesOk returns a tuple with the NumberOfShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfShares

`func (o *SharePositionView) SetNumberOfShares(v int64)`

SetNumberOfShares sets NumberOfShares field to given value.

### HasNumberOfShares

`func (o *SharePositionView) HasNumberOfShares() bool`

HasNumberOfShares returns a boolean if a field has been set.

### GetSecuritiesAccount

`func (o *SharePositionView) GetSecuritiesAccount() SecuritiesAccountCompactView`

GetSecuritiesAccount returns the SecuritiesAccount field if non-nil, zero value otherwise.

### GetSecuritiesAccountOk

`func (o *SharePositionView) GetSecuritiesAccountOk() (*SecuritiesAccountCompactView, bool)`

GetSecuritiesAccountOk returns a tuple with the SecuritiesAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritiesAccount

`func (o *SharePositionView) SetSecuritiesAccount(v SecuritiesAccountCompactView)`

SetSecuritiesAccount sets SecuritiesAccount field to given value.

### HasSecuritiesAccount

`func (o *SharePositionView) HasSecuritiesAccount() bool`

HasSecuritiesAccount returns a boolean if a field has been set.

### GetVersion

`func (o *SharePositionView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SharePositionView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SharePositionView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SharePositionView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


