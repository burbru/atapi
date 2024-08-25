# ShareholderView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to [**CompactCompanyView**](CompactCompanyView.md) |  | [optional] 
**NumberOfShares** | Pointer to **int64** |  | [optional] 
**OutstandingShares** | Pointer to **int64** |  | [optional] 
**SecurityIdentifier** | Pointer to **string** |  | [optional] 
**ShareInPercent** | Pointer to **float32** |  | [optional] 
**User** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 

## Methods

### NewShareholderView

`func NewShareholderView() *ShareholderView`

NewShareholderView instantiates a new ShareholderView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareholderViewWithDefaults

`func NewShareholderViewWithDefaults() *ShareholderView`

NewShareholderViewWithDefaults instantiates a new ShareholderView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *ShareholderView) GetCompany() CompactCompanyView`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ShareholderView) GetCompanyOk() (*CompactCompanyView, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ShareholderView) SetCompany(v CompactCompanyView)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ShareholderView) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetNumberOfShares

`func (o *ShareholderView) GetNumberOfShares() int64`

GetNumberOfShares returns the NumberOfShares field if non-nil, zero value otherwise.

### GetNumberOfSharesOk

`func (o *ShareholderView) GetNumberOfSharesOk() (*int64, bool)`

GetNumberOfSharesOk returns a tuple with the NumberOfShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfShares

`func (o *ShareholderView) SetNumberOfShares(v int64)`

SetNumberOfShares sets NumberOfShares field to given value.

### HasNumberOfShares

`func (o *ShareholderView) HasNumberOfShares() bool`

HasNumberOfShares returns a boolean if a field has been set.

### GetOutstandingShares

`func (o *ShareholderView) GetOutstandingShares() int64`

GetOutstandingShares returns the OutstandingShares field if non-nil, zero value otherwise.

### GetOutstandingSharesOk

`func (o *ShareholderView) GetOutstandingSharesOk() (*int64, bool)`

GetOutstandingSharesOk returns a tuple with the OutstandingShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutstandingShares

`func (o *ShareholderView) SetOutstandingShares(v int64)`

SetOutstandingShares sets OutstandingShares field to given value.

### HasOutstandingShares

`func (o *ShareholderView) HasOutstandingShares() bool`

HasOutstandingShares returns a boolean if a field has been set.

### GetSecurityIdentifier

`func (o *ShareholderView) GetSecurityIdentifier() string`

GetSecurityIdentifier returns the SecurityIdentifier field if non-nil, zero value otherwise.

### GetSecurityIdentifierOk

`func (o *ShareholderView) GetSecurityIdentifierOk() (*string, bool)`

GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIdentifier

`func (o *ShareholderView) SetSecurityIdentifier(v string)`

SetSecurityIdentifier sets SecurityIdentifier field to given value.

### HasSecurityIdentifier

`func (o *ShareholderView) HasSecurityIdentifier() bool`

HasSecurityIdentifier returns a boolean if a field has been set.

### GetShareInPercent

`func (o *ShareholderView) GetShareInPercent() float32`

GetShareInPercent returns the ShareInPercent field if non-nil, zero value otherwise.

### GetShareInPercentOk

`func (o *ShareholderView) GetShareInPercentOk() (*float32, bool)`

GetShareInPercentOk returns a tuple with the ShareInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareInPercent

`func (o *ShareholderView) SetShareInPercent(v float32)`

SetShareInPercent sets ShareInPercent field to given value.

### HasShareInPercent

`func (o *ShareholderView) HasShareInPercent() bool`

HasShareInPercent returns a boolean if a field has been set.

### GetUser

`func (o *ShareholderView) GetUser() UsernameView`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ShareholderView) GetUserOk() (*UsernameView, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ShareholderView) SetUser(v UsernameView)`

SetUser sets User field to given value.

### HasUser

`func (o *ShareholderView) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


