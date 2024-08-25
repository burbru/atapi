# FixedIncomeSecurityView

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

### NewFixedIncomeSecurityView

`func NewFixedIncomeSecurityView() *FixedIncomeSecurityView`

NewFixedIncomeSecurityView instantiates a new FixedIncomeSecurityView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedIncomeSecurityViewWithDefaults

`func NewFixedIncomeSecurityViewWithDefaults() *FixedIncomeSecurityView`

NewFixedIncomeSecurityViewWithDefaults instantiates a new FixedIncomeSecurityView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFaceValue

`func (o *FixedIncomeSecurityView) GetFaceValue() float32`

GetFaceValue returns the FaceValue field if non-nil, zero value otherwise.

### GetFaceValueOk

`func (o *FixedIncomeSecurityView) GetFaceValueOk() (*float32, bool)`

GetFaceValueOk returns a tuple with the FaceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValue

`func (o *FixedIncomeSecurityView) SetFaceValue(v float32)`

SetFaceValue sets FaceValue field to given value.

### HasFaceValue

`func (o *FixedIncomeSecurityView) HasFaceValue() bool`

HasFaceValue returns a boolean if a field has been set.

### GetId

`func (o *FixedIncomeSecurityView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FixedIncomeSecurityView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FixedIncomeSecurityView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FixedIncomeSecurityView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterestRate

`func (o *FixedIncomeSecurityView) GetInterestRate() float32`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *FixedIncomeSecurityView) GetInterestRateOk() (*float32, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *FixedIncomeSecurityView) SetInterestRate(v float32)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *FixedIncomeSecurityView) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetIssueDate

`func (o *FixedIncomeSecurityView) GetIssueDate() int64`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *FixedIncomeSecurityView) GetIssueDateOk() (*int64, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *FixedIncomeSecurityView) SetIssueDate(v int64)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *FixedIncomeSecurityView) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetListing

`func (o *FixedIncomeSecurityView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *FixedIncomeSecurityView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *FixedIncomeSecurityView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *FixedIncomeSecurityView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetMaturityDate

`func (o *FixedIncomeSecurityView) GetMaturityDate() int64`

GetMaturityDate returns the MaturityDate field if non-nil, zero value otherwise.

### GetMaturityDateOk

`func (o *FixedIncomeSecurityView) GetMaturityDateOk() (*int64, bool)`

GetMaturityDateOk returns a tuple with the MaturityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityDate

`func (o *FixedIncomeSecurityView) SetMaturityDate(v int64)`

SetMaturityDate sets MaturityDate field to given value.

### HasMaturityDate

`func (o *FixedIncomeSecurityView) HasMaturityDate() bool`

HasMaturityDate returns a boolean if a field has been set.

### GetName

`func (o *FixedIncomeSecurityView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FixedIncomeSecurityView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FixedIncomeSecurityView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FixedIncomeSecurityView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriceSpread

`func (o *FixedIncomeSecurityView) GetPriceSpread() PriceSpreadView`

GetPriceSpread returns the PriceSpread field if non-nil, zero value otherwise.

### GetPriceSpreadOk

`func (o *FixedIncomeSecurityView) GetPriceSpreadOk() (*PriceSpreadView, bool)`

GetPriceSpreadOk returns a tuple with the PriceSpread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSpread

`func (o *FixedIncomeSecurityView) SetPriceSpread(v PriceSpreadView)`

SetPriceSpread sets PriceSpread field to given value.

### HasPriceSpread

`func (o *FixedIncomeSecurityView) HasPriceSpread() bool`

HasPriceSpread returns a boolean if a field has been set.

### GetRepurchaseListing

`func (o *FixedIncomeSecurityView) GetRepurchaseListing() ListingView`

GetRepurchaseListing returns the RepurchaseListing field if non-nil, zero value otherwise.

### GetRepurchaseListingOk

`func (o *FixedIncomeSecurityView) GetRepurchaseListingOk() (*ListingView, bool)`

GetRepurchaseListingOk returns a tuple with the RepurchaseListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepurchaseListing

`func (o *FixedIncomeSecurityView) SetRepurchaseListing(v ListingView)`

SetRepurchaseListing sets RepurchaseListing field to given value.

### HasRepurchaseListing

`func (o *FixedIncomeSecurityView) HasRepurchaseListing() bool`

HasRepurchaseListing returns a boolean if a field has been set.

### GetVersion

`func (o *FixedIncomeSecurityView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FixedIncomeSecurityView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FixedIncomeSecurityView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FixedIncomeSecurityView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVolume

`func (o *FixedIncomeSecurityView) GetVolume() float32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *FixedIncomeSecurityView) GetVolumeOk() (*float32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *FixedIncomeSecurityView) SetVolume(v float32)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *FixedIncomeSecurityView) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


