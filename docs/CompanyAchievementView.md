# CompanyAchievementView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchievedDate** | Pointer to **int64** |  | [optional] 
**Claimed** | Pointer to **bool** |  | [optional] 
**CoinReward** | Pointer to **int64** |  | [optional] 
**Company** | Pointer to [**CompactCompanyView**](CompactCompanyView.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewCompanyAchievementView

`func NewCompanyAchievementView() *CompanyAchievementView`

NewCompanyAchievementView instantiates a new CompanyAchievementView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyAchievementViewWithDefaults

`func NewCompanyAchievementViewWithDefaults() *CompanyAchievementView`

NewCompanyAchievementViewWithDefaults instantiates a new CompanyAchievementView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchievedDate

`func (o *CompanyAchievementView) GetAchievedDate() int64`

GetAchievedDate returns the AchievedDate field if non-nil, zero value otherwise.

### GetAchievedDateOk

`func (o *CompanyAchievementView) GetAchievedDateOk() (*int64, bool)`

GetAchievedDateOk returns a tuple with the AchievedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievedDate

`func (o *CompanyAchievementView) SetAchievedDate(v int64)`

SetAchievedDate sets AchievedDate field to given value.

### HasAchievedDate

`func (o *CompanyAchievementView) HasAchievedDate() bool`

HasAchievedDate returns a boolean if a field has been set.

### GetClaimed

`func (o *CompanyAchievementView) GetClaimed() bool`

GetClaimed returns the Claimed field if non-nil, zero value otherwise.

### GetClaimedOk

`func (o *CompanyAchievementView) GetClaimedOk() (*bool, bool)`

GetClaimedOk returns a tuple with the Claimed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimed

`func (o *CompanyAchievementView) SetClaimed(v bool)`

SetClaimed sets Claimed field to given value.

### HasClaimed

`func (o *CompanyAchievementView) HasClaimed() bool`

HasClaimed returns a boolean if a field has been set.

### GetCoinReward

`func (o *CompanyAchievementView) GetCoinReward() int64`

GetCoinReward returns the CoinReward field if non-nil, zero value otherwise.

### GetCoinRewardOk

`func (o *CompanyAchievementView) GetCoinRewardOk() (*int64, bool)`

GetCoinRewardOk returns a tuple with the CoinReward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinReward

`func (o *CompanyAchievementView) SetCoinReward(v int64)`

SetCoinReward sets CoinReward field to given value.

### HasCoinReward

`func (o *CompanyAchievementView) HasCoinReward() bool`

HasCoinReward returns a boolean if a field has been set.

### GetCompany

`func (o *CompanyAchievementView) GetCompany() CompactCompanyView`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CompanyAchievementView) GetCompanyOk() (*CompactCompanyView, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CompanyAchievementView) SetCompany(v CompactCompanyView)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CompanyAchievementView) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetDescription

`func (o *CompanyAchievementView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CompanyAchievementView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CompanyAchievementView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CompanyAchievementView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndDate

`func (o *CompanyAchievementView) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CompanyAchievementView) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CompanyAchievementView) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CompanyAchievementView) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *CompanyAchievementView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyAchievementView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyAchievementView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyAchievementView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CompanyAchievementView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompanyAchievementView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompanyAchievementView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CompanyAchievementView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *CompanyAchievementView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CompanyAchievementView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CompanyAchievementView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CompanyAchievementView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


