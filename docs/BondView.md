# BondView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FaceValue** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InterestRate** | Pointer to **float32** |  | [optional] 
**IssueDate** | Pointer to **int64** |  | [optional] 
**Issuer** | Pointer to [**CompanyNameView**](CompanyNameView.md) |  | [optional] 
**Listing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**MaturityDate** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PriceSpread** | Pointer to [**PriceSpreadView**](PriceSpreadView.md) |  | [optional] 
**RepurchaseListing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**Volume** | Pointer to **float32** |  | [optional] 

## Methods

### NewBondView

`func NewBondView() *BondView`

NewBondView instantiates a new BondView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBondViewWithDefaults

`func NewBondViewWithDefaults() *BondView`

NewBondViewWithDefaults instantiates a new BondView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFaceValue

`func (o *BondView) GetFaceValue() float32`

GetFaceValue returns the FaceValue field if non-nil, zero value otherwise.

### GetFaceValueOk

`func (o *BondView) GetFaceValueOk() (*float32, bool)`

GetFaceValueOk returns a tuple with the FaceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValue

`func (o *BondView) SetFaceValue(v float32)`

SetFaceValue sets FaceValue field to given value.

### HasFaceValue

`func (o *BondView) HasFaceValue() bool`

HasFaceValue returns a boolean if a field has been set.

### GetId

`func (o *BondView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BondView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BondView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BondView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterestRate

`func (o *BondView) GetInterestRate() float32`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *BondView) GetInterestRateOk() (*float32, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *BondView) SetInterestRate(v float32)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *BondView) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetIssueDate

`func (o *BondView) GetIssueDate() int64`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *BondView) GetIssueDateOk() (*int64, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *BondView) SetIssueDate(v int64)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *BondView) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetIssuer

`func (o *BondView) GetIssuer() CompanyNameView`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *BondView) GetIssuerOk() (*CompanyNameView, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *BondView) SetIssuer(v CompanyNameView)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *BondView) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetListing

`func (o *BondView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *BondView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *BondView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *BondView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetMaturityDate

`func (o *BondView) GetMaturityDate() int64`

GetMaturityDate returns the MaturityDate field if non-nil, zero value otherwise.

### GetMaturityDateOk

`func (o *BondView) GetMaturityDateOk() (*int64, bool)`

GetMaturityDateOk returns a tuple with the MaturityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityDate

`func (o *BondView) SetMaturityDate(v int64)`

SetMaturityDate sets MaturityDate field to given value.

### HasMaturityDate

`func (o *BondView) HasMaturityDate() bool`

HasMaturityDate returns a boolean if a field has been set.

### GetName

`func (o *BondView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BondView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BondView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BondView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriceSpread

`func (o *BondView) GetPriceSpread() PriceSpreadView`

GetPriceSpread returns the PriceSpread field if non-nil, zero value otherwise.

### GetPriceSpreadOk

`func (o *BondView) GetPriceSpreadOk() (*PriceSpreadView, bool)`

GetPriceSpreadOk returns a tuple with the PriceSpread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSpread

`func (o *BondView) SetPriceSpread(v PriceSpreadView)`

SetPriceSpread sets PriceSpread field to given value.

### HasPriceSpread

`func (o *BondView) HasPriceSpread() bool`

HasPriceSpread returns a boolean if a field has been set.

### GetRepurchaseListing

`func (o *BondView) GetRepurchaseListing() ListingView`

GetRepurchaseListing returns the RepurchaseListing field if non-nil, zero value otherwise.

### GetRepurchaseListingOk

`func (o *BondView) GetRepurchaseListingOk() (*ListingView, bool)`

GetRepurchaseListingOk returns a tuple with the RepurchaseListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepurchaseListing

`func (o *BondView) SetRepurchaseListing(v ListingView)`

SetRepurchaseListing sets RepurchaseListing field to given value.

### HasRepurchaseListing

`func (o *BondView) HasRepurchaseListing() bool`

HasRepurchaseListing returns a boolean if a field has been set.

### GetVersion

`func (o *BondView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BondView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BondView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BondView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVolume

`func (o *BondView) GetVolume() float32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *BondView) GetVolumeOk() (*float32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *BondView) SetVolume(v float32)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *BondView) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


