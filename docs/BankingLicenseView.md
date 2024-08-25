# BankingLicenseView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to [**CompactCompanyView**](CompactCompanyView.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **int64** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewBankingLicenseView

`func NewBankingLicenseView() *BankingLicenseView`

NewBankingLicenseView instantiates a new BankingLicenseView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankingLicenseViewWithDefaults

`func NewBankingLicenseViewWithDefaults() *BankingLicenseView`

NewBankingLicenseViewWithDefaults instantiates a new BankingLicenseView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *BankingLicenseView) GetCompany() CompactCompanyView`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *BankingLicenseView) GetCompanyOk() (*CompactCompanyView, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *BankingLicenseView) SetCompany(v CompactCompanyView)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *BankingLicenseView) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetId

`func (o *BankingLicenseView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankingLicenseView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankingLicenseView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BankingLicenseView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartDate

`func (o *BankingLicenseView) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BankingLicenseView) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BankingLicenseView) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BankingLicenseView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetVersion

`func (o *BankingLicenseView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BankingLicenseView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BankingLicenseView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BankingLicenseView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


