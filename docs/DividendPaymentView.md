# DividendPaymentView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to [**CompactCompanyView**](CompactCompanyView.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MaximalCashVolume** | Pointer to **float32** |  | [optional] 
**StartDate** | Pointer to **int64** |  | [optional] 

## Methods

### NewDividendPaymentView

`func NewDividendPaymentView() *DividendPaymentView`

NewDividendPaymentView instantiates a new DividendPaymentView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDividendPaymentViewWithDefaults

`func NewDividendPaymentViewWithDefaults() *DividendPaymentView`

NewDividendPaymentViewWithDefaults instantiates a new DividendPaymentView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *DividendPaymentView) GetCompany() CompactCompanyView`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *DividendPaymentView) GetCompanyOk() (*CompactCompanyView, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *DividendPaymentView) SetCompany(v CompactCompanyView)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *DividendPaymentView) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetId

`func (o *DividendPaymentView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DividendPaymentView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DividendPaymentView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DividendPaymentView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaximalCashVolume

`func (o *DividendPaymentView) GetMaximalCashVolume() float32`

GetMaximalCashVolume returns the MaximalCashVolume field if non-nil, zero value otherwise.

### GetMaximalCashVolumeOk

`func (o *DividendPaymentView) GetMaximalCashVolumeOk() (*float32, bool)`

GetMaximalCashVolumeOk returns a tuple with the MaximalCashVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximalCashVolume

`func (o *DividendPaymentView) SetMaximalCashVolume(v float32)`

SetMaximalCashVolume sets MaximalCashVolume field to given value.

### HasMaximalCashVolume

`func (o *DividendPaymentView) HasMaximalCashVolume() bool`

HasMaximalCashVolume returns a boolean if a field has been set.

### GetStartDate

`func (o *DividendPaymentView) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DividendPaymentView) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DividendPaymentView) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DividendPaymentView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


