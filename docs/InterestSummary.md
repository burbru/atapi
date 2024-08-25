# InterestSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllianceInterest** | Pointer to [**AllianceInterest**](AllianceInterest.md) |  | [optional] 
**AuthorInterest** | Pointer to [**AuthorInterest**](AuthorInterest.md) |  | [optional] 
**CommonControversy** | Pointer to [**Interest**](Interest.md) |  | [optional] 
**CommonGossip** | Pointer to [**Interest**](Interest.md) |  | [optional] 
**CommonPopularity** | Pointer to [**Interest**](Interest.md) |  | [optional] 
**CompanyInterest** | Pointer to [**CompanyInterest**](CompanyInterest.md) |  | [optional] 
**HashTagInterests** | Pointer to [**[]HashTagInterest**](HashTagInterest.md) |  | [optional] 
**InterestSum** | Pointer to [**Interest**](Interest.md) |  | [optional] 
**LocaleInterest** | Pointer to [**LocaleInterest**](LocaleInterest.md) |  | [optional] 
**MeanHashTagInterest** | Pointer to [**Interest**](Interest.md) |  | [optional] 
**PostInterest** | Pointer to [**PostInterest**](PostInterest.md) |  | [optional] 

## Methods

### NewInterestSummary

`func NewInterestSummary() *InterestSummary`

NewInterestSummary instantiates a new InterestSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterestSummaryWithDefaults

`func NewInterestSummaryWithDefaults() *InterestSummary`

NewInterestSummaryWithDefaults instantiates a new InterestSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllianceInterest

`func (o *InterestSummary) GetAllianceInterest() AllianceInterest`

GetAllianceInterest returns the AllianceInterest field if non-nil, zero value otherwise.

### GetAllianceInterestOk

`func (o *InterestSummary) GetAllianceInterestOk() (*AllianceInterest, bool)`

GetAllianceInterestOk returns a tuple with the AllianceInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllianceInterest

`func (o *InterestSummary) SetAllianceInterest(v AllianceInterest)`

SetAllianceInterest sets AllianceInterest field to given value.

### HasAllianceInterest

`func (o *InterestSummary) HasAllianceInterest() bool`

HasAllianceInterest returns a boolean if a field has been set.

### GetAuthorInterest

`func (o *InterestSummary) GetAuthorInterest() AuthorInterest`

GetAuthorInterest returns the AuthorInterest field if non-nil, zero value otherwise.

### GetAuthorInterestOk

`func (o *InterestSummary) GetAuthorInterestOk() (*AuthorInterest, bool)`

GetAuthorInterestOk returns a tuple with the AuthorInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorInterest

`func (o *InterestSummary) SetAuthorInterest(v AuthorInterest)`

SetAuthorInterest sets AuthorInterest field to given value.

### HasAuthorInterest

`func (o *InterestSummary) HasAuthorInterest() bool`

HasAuthorInterest returns a boolean if a field has been set.

### GetCommonControversy

`func (o *InterestSummary) GetCommonControversy() Interest`

GetCommonControversy returns the CommonControversy field if non-nil, zero value otherwise.

### GetCommonControversyOk

`func (o *InterestSummary) GetCommonControversyOk() (*Interest, bool)`

GetCommonControversyOk returns a tuple with the CommonControversy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonControversy

`func (o *InterestSummary) SetCommonControversy(v Interest)`

SetCommonControversy sets CommonControversy field to given value.

### HasCommonControversy

`func (o *InterestSummary) HasCommonControversy() bool`

HasCommonControversy returns a boolean if a field has been set.

### GetCommonGossip

`func (o *InterestSummary) GetCommonGossip() Interest`

GetCommonGossip returns the CommonGossip field if non-nil, zero value otherwise.

### GetCommonGossipOk

`func (o *InterestSummary) GetCommonGossipOk() (*Interest, bool)`

GetCommonGossipOk returns a tuple with the CommonGossip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonGossip

`func (o *InterestSummary) SetCommonGossip(v Interest)`

SetCommonGossip sets CommonGossip field to given value.

### HasCommonGossip

`func (o *InterestSummary) HasCommonGossip() bool`

HasCommonGossip returns a boolean if a field has been set.

### GetCommonPopularity

`func (o *InterestSummary) GetCommonPopularity() Interest`

GetCommonPopularity returns the CommonPopularity field if non-nil, zero value otherwise.

### GetCommonPopularityOk

`func (o *InterestSummary) GetCommonPopularityOk() (*Interest, bool)`

GetCommonPopularityOk returns a tuple with the CommonPopularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonPopularity

`func (o *InterestSummary) SetCommonPopularity(v Interest)`

SetCommonPopularity sets CommonPopularity field to given value.

### HasCommonPopularity

`func (o *InterestSummary) HasCommonPopularity() bool`

HasCommonPopularity returns a boolean if a field has been set.

### GetCompanyInterest

`func (o *InterestSummary) GetCompanyInterest() CompanyInterest`

GetCompanyInterest returns the CompanyInterest field if non-nil, zero value otherwise.

### GetCompanyInterestOk

`func (o *InterestSummary) GetCompanyInterestOk() (*CompanyInterest, bool)`

GetCompanyInterestOk returns a tuple with the CompanyInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyInterest

`func (o *InterestSummary) SetCompanyInterest(v CompanyInterest)`

SetCompanyInterest sets CompanyInterest field to given value.

### HasCompanyInterest

`func (o *InterestSummary) HasCompanyInterest() bool`

HasCompanyInterest returns a boolean if a field has been set.

### GetHashTagInterests

`func (o *InterestSummary) GetHashTagInterests() []HashTagInterest`

GetHashTagInterests returns the HashTagInterests field if non-nil, zero value otherwise.

### GetHashTagInterestsOk

`func (o *InterestSummary) GetHashTagInterestsOk() (*[]HashTagInterest, bool)`

GetHashTagInterestsOk returns a tuple with the HashTagInterests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashTagInterests

`func (o *InterestSummary) SetHashTagInterests(v []HashTagInterest)`

SetHashTagInterests sets HashTagInterests field to given value.

### HasHashTagInterests

`func (o *InterestSummary) HasHashTagInterests() bool`

HasHashTagInterests returns a boolean if a field has been set.

### GetInterestSum

`func (o *InterestSummary) GetInterestSum() Interest`

GetInterestSum returns the InterestSum field if non-nil, zero value otherwise.

### GetInterestSumOk

`func (o *InterestSummary) GetInterestSumOk() (*Interest, bool)`

GetInterestSumOk returns a tuple with the InterestSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestSum

`func (o *InterestSummary) SetInterestSum(v Interest)`

SetInterestSum sets InterestSum field to given value.

### HasInterestSum

`func (o *InterestSummary) HasInterestSum() bool`

HasInterestSum returns a boolean if a field has been set.

### GetLocaleInterest

`func (o *InterestSummary) GetLocaleInterest() LocaleInterest`

GetLocaleInterest returns the LocaleInterest field if non-nil, zero value otherwise.

### GetLocaleInterestOk

`func (o *InterestSummary) GetLocaleInterestOk() (*LocaleInterest, bool)`

GetLocaleInterestOk returns a tuple with the LocaleInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocaleInterest

`func (o *InterestSummary) SetLocaleInterest(v LocaleInterest)`

SetLocaleInterest sets LocaleInterest field to given value.

### HasLocaleInterest

`func (o *InterestSummary) HasLocaleInterest() bool`

HasLocaleInterest returns a boolean if a field has been set.

### GetMeanHashTagInterest

`func (o *InterestSummary) GetMeanHashTagInterest() Interest`

GetMeanHashTagInterest returns the MeanHashTagInterest field if non-nil, zero value otherwise.

### GetMeanHashTagInterestOk

`func (o *InterestSummary) GetMeanHashTagInterestOk() (*Interest, bool)`

GetMeanHashTagInterestOk returns a tuple with the MeanHashTagInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeanHashTagInterest

`func (o *InterestSummary) SetMeanHashTagInterest(v Interest)`

SetMeanHashTagInterest sets MeanHashTagInterest field to given value.

### HasMeanHashTagInterest

`func (o *InterestSummary) HasMeanHashTagInterest() bool`

HasMeanHashTagInterest returns a boolean if a field has been set.

### GetPostInterest

`func (o *InterestSummary) GetPostInterest() PostInterest`

GetPostInterest returns the PostInterest field if non-nil, zero value otherwise.

### GetPostInterestOk

`func (o *InterestSummary) GetPostInterestOk() (*PostInterest, bool)`

GetPostInterestOk returns a tuple with the PostInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostInterest

`func (o *InterestSummary) SetPostInterest(v PostInterest)`

SetPostInterest sets PostInterest field to given value.

### HasPostInterest

`func (o *InterestSummary) HasPostInterest() bool`

HasPostInterest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


