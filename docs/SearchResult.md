# SearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bond** | Pointer to [**BondView**](BondView.md) |  | [optional] 
**Company** | Pointer to [**CompanyNameView**](CompanyNameView.md) |  | [optional] 
**Listing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**NaturalIdentifier** | Pointer to **string** |  | [optional] 
**SystemBond** | Pointer to [**SystemBondView**](SystemBondView.md) |  | [optional] 
**UserAccount** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 

## Methods

### NewSearchResult

`func NewSearchResult() *SearchResult`

NewSearchResult instantiates a new SearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultWithDefaults

`func NewSearchResultWithDefaults() *SearchResult`

NewSearchResultWithDefaults instantiates a new SearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBond

`func (o *SearchResult) GetBond() BondView`

GetBond returns the Bond field if non-nil, zero value otherwise.

### GetBondOk

`func (o *SearchResult) GetBondOk() (*BondView, bool)`

GetBondOk returns a tuple with the Bond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBond

`func (o *SearchResult) SetBond(v BondView)`

SetBond sets Bond field to given value.

### HasBond

`func (o *SearchResult) HasBond() bool`

HasBond returns a boolean if a field has been set.

### GetCompany

`func (o *SearchResult) GetCompany() CompanyNameView`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *SearchResult) GetCompanyOk() (*CompanyNameView, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *SearchResult) SetCompany(v CompanyNameView)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *SearchResult) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetListing

`func (o *SearchResult) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *SearchResult) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *SearchResult) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *SearchResult) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetNaturalIdentifier

`func (o *SearchResult) GetNaturalIdentifier() string`

GetNaturalIdentifier returns the NaturalIdentifier field if non-nil, zero value otherwise.

### GetNaturalIdentifierOk

`func (o *SearchResult) GetNaturalIdentifierOk() (*string, bool)`

GetNaturalIdentifierOk returns a tuple with the NaturalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalIdentifier

`func (o *SearchResult) SetNaturalIdentifier(v string)`

SetNaturalIdentifier sets NaturalIdentifier field to given value.

### HasNaturalIdentifier

`func (o *SearchResult) HasNaturalIdentifier() bool`

HasNaturalIdentifier returns a boolean if a field has been set.

### GetSystemBond

`func (o *SearchResult) GetSystemBond() SystemBondView`

GetSystemBond returns the SystemBond field if non-nil, zero value otherwise.

### GetSystemBondOk

`func (o *SearchResult) GetSystemBondOk() (*SystemBondView, bool)`

GetSystemBondOk returns a tuple with the SystemBond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemBond

`func (o *SearchResult) SetSystemBond(v SystemBondView)`

SetSystemBond sets SystemBond field to given value.

### HasSystemBond

`func (o *SearchResult) HasSystemBond() bool`

HasSystemBond returns a boolean if a field has been set.

### GetUserAccount

`func (o *SearchResult) GetUserAccount() UsernameView`

GetUserAccount returns the UserAccount field if non-nil, zero value otherwise.

### GetUserAccountOk

`func (o *SearchResult) GetUserAccountOk() (*UsernameView, bool)`

GetUserAccountOk returns a tuple with the UserAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccount

`func (o *SearchResult) SetUserAccount(v UsernameView)`

SetUserAccount sets UserAccount field to given value.

### HasUserAccount

`func (o *SearchResult) HasUserAccount() bool`

HasUserAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


