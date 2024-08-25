# WarrantView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to [**CompanyView**](CompanyView.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Listing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**Ratio** | Pointer to **float32** |  | [optional] 
**SubscriptionPeriodDate** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Underlying** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**UnderlyingCapValue** | Pointer to **float32** |  | [optional] 
**UnderlyingValue** | Pointer to **float32** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewWarrantView

`func NewWarrantView() *WarrantView`

NewWarrantView instantiates a new WarrantView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarrantViewWithDefaults

`func NewWarrantViewWithDefaults() *WarrantView`

NewWarrantViewWithDefaults instantiates a new WarrantView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *WarrantView) GetCompany() CompanyView`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *WarrantView) GetCompanyOk() (*CompanyView, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *WarrantView) SetCompany(v CompanyView)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *WarrantView) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetId

`func (o *WarrantView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WarrantView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WarrantView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WarrantView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetListing

`func (o *WarrantView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *WarrantView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *WarrantView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *WarrantView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetRatio

`func (o *WarrantView) GetRatio() float32`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *WarrantView) GetRatioOk() (*float32, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *WarrantView) SetRatio(v float32)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *WarrantView) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetSubscriptionPeriodDate

`func (o *WarrantView) GetSubscriptionPeriodDate() int64`

GetSubscriptionPeriodDate returns the SubscriptionPeriodDate field if non-nil, zero value otherwise.

### GetSubscriptionPeriodDateOk

`func (o *WarrantView) GetSubscriptionPeriodDateOk() (*int64, bool)`

GetSubscriptionPeriodDateOk returns a tuple with the SubscriptionPeriodDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriodDate

`func (o *WarrantView) SetSubscriptionPeriodDate(v int64)`

SetSubscriptionPeriodDate sets SubscriptionPeriodDate field to given value.

### HasSubscriptionPeriodDate

`func (o *WarrantView) HasSubscriptionPeriodDate() bool`

HasSubscriptionPeriodDate returns a boolean if a field has been set.

### GetType

`func (o *WarrantView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WarrantView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WarrantView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WarrantView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnderlying

`func (o *WarrantView) GetUnderlying() ListingView`

GetUnderlying returns the Underlying field if non-nil, zero value otherwise.

### GetUnderlyingOk

`func (o *WarrantView) GetUnderlyingOk() (*ListingView, bool)`

GetUnderlyingOk returns a tuple with the Underlying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlying

`func (o *WarrantView) SetUnderlying(v ListingView)`

SetUnderlying sets Underlying field to given value.

### HasUnderlying

`func (o *WarrantView) HasUnderlying() bool`

HasUnderlying returns a boolean if a field has been set.

### GetUnderlyingCapValue

`func (o *WarrantView) GetUnderlyingCapValue() float32`

GetUnderlyingCapValue returns the UnderlyingCapValue field if non-nil, zero value otherwise.

### GetUnderlyingCapValueOk

`func (o *WarrantView) GetUnderlyingCapValueOk() (*float32, bool)`

GetUnderlyingCapValueOk returns a tuple with the UnderlyingCapValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingCapValue

`func (o *WarrantView) SetUnderlyingCapValue(v float32)`

SetUnderlyingCapValue sets UnderlyingCapValue field to given value.

### HasUnderlyingCapValue

`func (o *WarrantView) HasUnderlyingCapValue() bool`

HasUnderlyingCapValue returns a boolean if a field has been set.

### GetUnderlyingValue

`func (o *WarrantView) GetUnderlyingValue() float32`

GetUnderlyingValue returns the UnderlyingValue field if non-nil, zero value otherwise.

### GetUnderlyingValueOk

`func (o *WarrantView) GetUnderlyingValueOk() (*float32, bool)`

GetUnderlyingValueOk returns a tuple with the UnderlyingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingValue

`func (o *WarrantView) SetUnderlyingValue(v float32)`

SetUnderlyingValue sets UnderlyingValue field to given value.

### HasUnderlyingValue

`func (o *WarrantView) HasUnderlyingValue() bool`

HasUnderlyingValue returns a boolean if a field has been set.

### GetVersion

`func (o *WarrantView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WarrantView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WarrantView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WarrantView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


