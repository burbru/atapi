# CompanyView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchievementCount** | Pointer to **int64** |  | [optional] 
**AchievementTotal** | Pointer to **int64** |  | [optional] 
**BankAccount** | Pointer to [**BankAccountView**](BankAccountView.md) |  | [optional] 
**Ceo** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Listing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SecuritiesAccountId** | Pointer to **string** |  | [optional] 
**SecurityIdentifier** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewCompanyView

`func NewCompanyView() *CompanyView`

NewCompanyView instantiates a new CompanyView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyViewWithDefaults

`func NewCompanyViewWithDefaults() *CompanyView`

NewCompanyViewWithDefaults instantiates a new CompanyView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchievementCount

`func (o *CompanyView) GetAchievementCount() int64`

GetAchievementCount returns the AchievementCount field if non-nil, zero value otherwise.

### GetAchievementCountOk

`func (o *CompanyView) GetAchievementCountOk() (*int64, bool)`

GetAchievementCountOk returns a tuple with the AchievementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementCount

`func (o *CompanyView) SetAchievementCount(v int64)`

SetAchievementCount sets AchievementCount field to given value.

### HasAchievementCount

`func (o *CompanyView) HasAchievementCount() bool`

HasAchievementCount returns a boolean if a field has been set.

### GetAchievementTotal

`func (o *CompanyView) GetAchievementTotal() int64`

GetAchievementTotal returns the AchievementTotal field if non-nil, zero value otherwise.

### GetAchievementTotalOk

`func (o *CompanyView) GetAchievementTotalOk() (*int64, bool)`

GetAchievementTotalOk returns a tuple with the AchievementTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementTotal

`func (o *CompanyView) SetAchievementTotal(v int64)`

SetAchievementTotal sets AchievementTotal field to given value.

### HasAchievementTotal

`func (o *CompanyView) HasAchievementTotal() bool`

HasAchievementTotal returns a boolean if a field has been set.

### GetBankAccount

`func (o *CompanyView) GetBankAccount() BankAccountView`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *CompanyView) GetBankAccountOk() (*BankAccountView, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *CompanyView) SetBankAccount(v BankAccountView)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *CompanyView) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### GetCeo

`func (o *CompanyView) GetCeo() UsernameView`

GetCeo returns the Ceo field if non-nil, zero value otherwise.

### GetCeoOk

`func (o *CompanyView) GetCeoOk() (*UsernameView, bool)`

GetCeoOk returns a tuple with the Ceo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeo

`func (o *CompanyView) SetCeo(v UsernameView)`

SetCeo sets Ceo field to given value.

### HasCeo

`func (o *CompanyView) HasCeo() bool`

HasCeo returns a boolean if a field has been set.

### GetId

`func (o *CompanyView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetListing

`func (o *CompanyView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *CompanyView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *CompanyView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *CompanyView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetLogoUrl

`func (o *CompanyView) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *CompanyView) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *CompanyView) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *CompanyView) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetName

`func (o *CompanyView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecuritiesAccountId

`func (o *CompanyView) GetSecuritiesAccountId() string`

GetSecuritiesAccountId returns the SecuritiesAccountId field if non-nil, zero value otherwise.

### GetSecuritiesAccountIdOk

`func (o *CompanyView) GetSecuritiesAccountIdOk() (*string, bool)`

GetSecuritiesAccountIdOk returns a tuple with the SecuritiesAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritiesAccountId

`func (o *CompanyView) SetSecuritiesAccountId(v string)`

SetSecuritiesAccountId sets SecuritiesAccountId field to given value.

### HasSecuritiesAccountId

`func (o *CompanyView) HasSecuritiesAccountId() bool`

HasSecuritiesAccountId returns a boolean if a field has been set.

### GetSecurityIdentifier

`func (o *CompanyView) GetSecurityIdentifier() string`

GetSecurityIdentifier returns the SecurityIdentifier field if non-nil, zero value otherwise.

### GetSecurityIdentifierOk

`func (o *CompanyView) GetSecurityIdentifierOk() (*string, bool)`

GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIdentifier

`func (o *CompanyView) SetSecurityIdentifier(v string)`

SetSecurityIdentifier sets SecurityIdentifier field to given value.

### HasSecurityIdentifier

`func (o *CompanyView) HasSecurityIdentifier() bool`

HasSecurityIdentifier returns a boolean if a field has been set.

### GetVersion

`func (o *CompanyView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CompanyView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CompanyView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CompanyView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


