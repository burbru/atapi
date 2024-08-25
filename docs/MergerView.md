# MergerView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcquiringCompany** | Pointer to [**CompactCompanyView**](CompactCompanyView.md) |  | [optional] 
**Company** | Pointer to [**CompactCompanyView**](CompactCompanyView.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MaximalCashVolume** | Pointer to **float32** |  | [optional] 
**StartDate** | Pointer to **int64** |  | [optional] 

## Methods

### NewMergerView

`func NewMergerView() *MergerView`

NewMergerView instantiates a new MergerView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergerViewWithDefaults

`func NewMergerViewWithDefaults() *MergerView`

NewMergerViewWithDefaults instantiates a new MergerView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcquiringCompany

`func (o *MergerView) GetAcquiringCompany() CompactCompanyView`

GetAcquiringCompany returns the AcquiringCompany field if non-nil, zero value otherwise.

### GetAcquiringCompanyOk

`func (o *MergerView) GetAcquiringCompanyOk() (*CompactCompanyView, bool)`

GetAcquiringCompanyOk returns a tuple with the AcquiringCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquiringCompany

`func (o *MergerView) SetAcquiringCompany(v CompactCompanyView)`

SetAcquiringCompany sets AcquiringCompany field to given value.

### HasAcquiringCompany

`func (o *MergerView) HasAcquiringCompany() bool`

HasAcquiringCompany returns a boolean if a field has been set.

### GetCompany

`func (o *MergerView) GetCompany() CompactCompanyView`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *MergerView) GetCompanyOk() (*CompactCompanyView, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *MergerView) SetCompany(v CompactCompanyView)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *MergerView) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetId

`func (o *MergerView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MergerView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MergerView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MergerView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaximalCashVolume

`func (o *MergerView) GetMaximalCashVolume() float32`

GetMaximalCashVolume returns the MaximalCashVolume field if non-nil, zero value otherwise.

### GetMaximalCashVolumeOk

`func (o *MergerView) GetMaximalCashVolumeOk() (*float32, bool)`

GetMaximalCashVolumeOk returns a tuple with the MaximalCashVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximalCashVolume

`func (o *MergerView) SetMaximalCashVolume(v float32)`

SetMaximalCashVolume sets MaximalCashVolume field to given value.

### HasMaximalCashVolume

`func (o *MergerView) HasMaximalCashVolume() bool`

HasMaximalCashVolume returns a boolean if a field has been set.

### GetStartDate

`func (o *MergerView) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *MergerView) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *MergerView) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *MergerView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


