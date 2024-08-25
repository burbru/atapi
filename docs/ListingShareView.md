# ListingShareView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Listing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**ShareInPercent** | Pointer to **float32** |  | [optional] 

## Methods

### NewListingShareView

`func NewListingShareView() *ListingShareView`

NewListingShareView instantiates a new ListingShareView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingShareViewWithDefaults

`func NewListingShareViewWithDefaults() *ListingShareView`

NewListingShareViewWithDefaults instantiates a new ListingShareView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListing

`func (o *ListingShareView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *ListingShareView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *ListingShareView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *ListingShareView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetShareInPercent

`func (o *ListingShareView) GetShareInPercent() float32`

GetShareInPercent returns the ShareInPercent field if non-nil, zero value otherwise.

### GetShareInPercentOk

`func (o *ListingShareView) GetShareInPercentOk() (*float32, bool)`

GetShareInPercentOk returns a tuple with the ShareInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareInPercent

`func (o *ListingShareView) SetShareInPercent(v float32)`

SetShareInPercent sets ShareInPercent field to given value.

### HasShareInPercent

`func (o *ListingShareView) HasShareInPercent() bool`

HasShareInPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


