# SystemBondView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FaceValue** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InterestRate** | Pointer to **float32** |  | [optional] 
**IssueDate** | Pointer to **int64** |  | [optional] 
**Listing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**MaturityDate** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PriceSpread** | Pointer to [**PriceSpreadView**](PriceSpreadView.md) |  | [optional] 
**RepurchaseListing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**Volume** | Pointer to **float32** |  | [optional] 

## Methods

### NewSystemBondView

`func NewSystemBondView() *SystemBondView`

NewSystemBondView instantiates a new SystemBondView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemBondViewWithDefaults

`func NewSystemBondViewWithDefaults() *SystemBondView`

NewSystemBondViewWithDefaults instantiates a new SystemBondView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFaceValue

`func (o *SystemBondView) GetFaceValue() float32`

GetFaceValue returns the FaceValue field if non-nil, zero value otherwise.

### GetFaceValueOk

`func (o *SystemBondView) GetFaceValueOk() (*float32, bool)`

GetFaceValueOk returns a tuple with the FaceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValue

`func (o *SystemBondView) SetFaceValue(v float32)`

SetFaceValue sets FaceValue field to given value.

### HasFaceValue

`func (o *SystemBondView) HasFaceValue() bool`

HasFaceValue returns a boolean if a field has been set.

### GetId

`func (o *SystemBondView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemBondView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemBondView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SystemBondView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterestRate

`func (o *SystemBondView) GetInterestRate() float32`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *SystemBondView) GetInterestRateOk() (*float32, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *SystemBondView) SetInterestRate(v float32)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *SystemBondView) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetIssueDate

`func (o *SystemBondView) GetIssueDate() int64`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *SystemBondView) GetIssueDateOk() (*int64, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *SystemBondView) SetIssueDate(v int64)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *SystemBondView) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetListing

`func (o *SystemBondView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *SystemBondView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *SystemBondView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *SystemBondView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetMaturityDate

`func (o *SystemBondView) GetMaturityDate() int64`

GetMaturityDate returns the MaturityDate field if non-nil, zero value otherwise.

### GetMaturityDateOk

`func (o *SystemBondView) GetMaturityDateOk() (*int64, bool)`

GetMaturityDateOk returns a tuple with the MaturityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityDate

`func (o *SystemBondView) SetMaturityDate(v int64)`

SetMaturityDate sets MaturityDate field to given value.

### HasMaturityDate

`func (o *SystemBondView) HasMaturityDate() bool`

HasMaturityDate returns a boolean if a field has been set.

### GetName

`func (o *SystemBondView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemBondView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemBondView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemBondView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriceSpread

`func (o *SystemBondView) GetPriceSpread() PriceSpreadView`

GetPriceSpread returns the PriceSpread field if non-nil, zero value otherwise.

### GetPriceSpreadOk

`func (o *SystemBondView) GetPriceSpreadOk() (*PriceSpreadView, bool)`

GetPriceSpreadOk returns a tuple with the PriceSpread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSpread

`func (o *SystemBondView) SetPriceSpread(v PriceSpreadView)`

SetPriceSpread sets PriceSpread field to given value.

### HasPriceSpread

`func (o *SystemBondView) HasPriceSpread() bool`

HasPriceSpread returns a boolean if a field has been set.

### GetRepurchaseListing

`func (o *SystemBondView) GetRepurchaseListing() ListingView`

GetRepurchaseListing returns the RepurchaseListing field if non-nil, zero value otherwise.

### GetRepurchaseListingOk

`func (o *SystemBondView) GetRepurchaseListingOk() (*ListingView, bool)`

GetRepurchaseListingOk returns a tuple with the RepurchaseListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepurchaseListing

`func (o *SystemBondView) SetRepurchaseListing(v ListingView)`

SetRepurchaseListing sets RepurchaseListing field to given value.

### HasRepurchaseListing

`func (o *SystemBondView) HasRepurchaseListing() bool`

HasRepurchaseListing returns a boolean if a field has been set.

### GetVersion

`func (o *SystemBondView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SystemBondView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SystemBondView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SystemBondView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVolume

`func (o *SystemBondView) GetVolume() float32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *SystemBondView) GetVolumeOk() (*float32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *SystemBondView) SetVolume(v float32)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *SystemBondView) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


