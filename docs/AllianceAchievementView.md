# AllianceAchievementView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchievedDate** | Pointer to **int64** |  | [optional] 
**Alliance** | Pointer to [**AllianceView**](AllianceView.md) |  | [optional] 
**Claimed** | Pointer to **bool** |  | [optional] 
**CoinReward** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewAllianceAchievementView

`func NewAllianceAchievementView() *AllianceAchievementView`

NewAllianceAchievementView instantiates a new AllianceAchievementView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllianceAchievementViewWithDefaults

`func NewAllianceAchievementViewWithDefaults() *AllianceAchievementView`

NewAllianceAchievementViewWithDefaults instantiates a new AllianceAchievementView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchievedDate

`func (o *AllianceAchievementView) GetAchievedDate() int64`

GetAchievedDate returns the AchievedDate field if non-nil, zero value otherwise.

### GetAchievedDateOk

`func (o *AllianceAchievementView) GetAchievedDateOk() (*int64, bool)`

GetAchievedDateOk returns a tuple with the AchievedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievedDate

`func (o *AllianceAchievementView) SetAchievedDate(v int64)`

SetAchievedDate sets AchievedDate field to given value.

### HasAchievedDate

`func (o *AllianceAchievementView) HasAchievedDate() bool`

HasAchievedDate returns a boolean if a field has been set.

### GetAlliance

`func (o *AllianceAchievementView) GetAlliance() AllianceView`

GetAlliance returns the Alliance field if non-nil, zero value otherwise.

### GetAllianceOk

`func (o *AllianceAchievementView) GetAllianceOk() (*AllianceView, bool)`

GetAllianceOk returns a tuple with the Alliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlliance

`func (o *AllianceAchievementView) SetAlliance(v AllianceView)`

SetAlliance sets Alliance field to given value.

### HasAlliance

`func (o *AllianceAchievementView) HasAlliance() bool`

HasAlliance returns a boolean if a field has been set.

### GetClaimed

`func (o *AllianceAchievementView) GetClaimed() bool`

GetClaimed returns the Claimed field if non-nil, zero value otherwise.

### GetClaimedOk

`func (o *AllianceAchievementView) GetClaimedOk() (*bool, bool)`

GetClaimedOk returns a tuple with the Claimed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimed

`func (o *AllianceAchievementView) SetClaimed(v bool)`

SetClaimed sets Claimed field to given value.

### HasClaimed

`func (o *AllianceAchievementView) HasClaimed() bool`

HasClaimed returns a boolean if a field has been set.

### GetCoinReward

`func (o *AllianceAchievementView) GetCoinReward() int64`

GetCoinReward returns the CoinReward field if non-nil, zero value otherwise.

### GetCoinRewardOk

`func (o *AllianceAchievementView) GetCoinRewardOk() (*int64, bool)`

GetCoinRewardOk returns a tuple with the CoinReward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinReward

`func (o *AllianceAchievementView) SetCoinReward(v int64)`

SetCoinReward sets CoinReward field to given value.

### HasCoinReward

`func (o *AllianceAchievementView) HasCoinReward() bool`

HasCoinReward returns a boolean if a field has been set.

### GetDescription

`func (o *AllianceAchievementView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AllianceAchievementView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AllianceAchievementView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AllianceAchievementView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndDate

`func (o *AllianceAchievementView) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AllianceAchievementView) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AllianceAchievementView) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AllianceAchievementView) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *AllianceAchievementView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllianceAchievementView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllianceAchievementView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AllianceAchievementView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AllianceAchievementView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AllianceAchievementView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AllianceAchievementView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AllianceAchievementView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *AllianceAchievementView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AllianceAchievementView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AllianceAchievementView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AllianceAchievementView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


