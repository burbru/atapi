# EmploymentAgreementSalaryPaymentCompactCompanyView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to [**CompactCompanyView**](CompactCompanyView.md) |  | [optional] 
**DailyWage** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastPayment** | Pointer to [**SalaryPaymentView**](SalaryPaymentView.md) |  | [optional] 
**PayAutomatically** | Pointer to **bool** |  | [optional] 
**StartDate** | Pointer to **int64** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewEmploymentAgreementSalaryPaymentCompactCompanyView

`func NewEmploymentAgreementSalaryPaymentCompactCompanyView() *EmploymentAgreementSalaryPaymentCompactCompanyView`

NewEmploymentAgreementSalaryPaymentCompactCompanyView instantiates a new EmploymentAgreementSalaryPaymentCompactCompanyView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmploymentAgreementSalaryPaymentCompactCompanyViewWithDefaults

`func NewEmploymentAgreementSalaryPaymentCompactCompanyViewWithDefaults() *EmploymentAgreementSalaryPaymentCompactCompanyView`

NewEmploymentAgreementSalaryPaymentCompactCompanyViewWithDefaults instantiates a new EmploymentAgreementSalaryPaymentCompactCompanyView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) GetCompany() CompactCompanyView`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) GetCompanyOk() (*CompactCompanyView, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) SetCompany(v CompactCompanyView)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetDailyWage

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) GetDailyWage() float32`

GetDailyWage returns the DailyWage field if non-nil, zero value otherwise.

### GetDailyWageOk

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) GetDailyWageOk() (*float32, bool)`

GetDailyWageOk returns a tuple with the DailyWage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyWage

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) SetDailyWage(v float32)`

SetDailyWage sets DailyWage field to given value.

### HasDailyWage

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) HasDailyWage() bool`

HasDailyWage returns a boolean if a field has been set.

### GetId

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastPayment

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) GetLastPayment() SalaryPaymentView`

GetLastPayment returns the LastPayment field if non-nil, zero value otherwise.

### GetLastPaymentOk

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) GetLastPaymentOk() (*SalaryPaymentView, bool)`

GetLastPaymentOk returns a tuple with the LastPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPayment

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) SetLastPayment(v SalaryPaymentView)`

SetLastPayment sets LastPayment field to given value.

### HasLastPayment

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) HasLastPayment() bool`

HasLastPayment returns a boolean if a field has been set.

### GetPayAutomatically

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) GetPayAutomatically() bool`

GetPayAutomatically returns the PayAutomatically field if non-nil, zero value otherwise.

### GetPayAutomaticallyOk

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) GetPayAutomaticallyOk() (*bool, bool)`

GetPayAutomaticallyOk returns a tuple with the PayAutomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayAutomatically

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) SetPayAutomatically(v bool)`

SetPayAutomatically sets PayAutomatically field to given value.

### HasPayAutomatically

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) HasPayAutomatically() bool`

HasPayAutomatically returns a boolean if a field has been set.

### GetStartDate

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetVersion

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EmploymentAgreementSalaryPaymentCompactCompanyView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


