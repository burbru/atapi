# UserProfileView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccount** | Pointer to [**BankAccountView**](BankAccountView.md) |  | [optional] 
**CashTransferLogs** | Pointer to [**[]CashTransferLogEntryView**](CashTransferLogEntryView.md) |  | [optional] 
**Employments** | Pointer to [**[]EmploymentAgreementCompactCompanyView**](EmploymentAgreementCompactCompanyView.md) |  | [optional] 
**InitiatedPolls** | Pointer to [**[]AbstractPollView**](AbstractPollView.md) |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**Polls** | Pointer to [**[]AbstractPollView**](AbstractPollView.md) |  | [optional] 
**SalaryPayments** | Pointer to [**[]SalaryPaymentView**](SalaryPaymentView.md) |  | [optional] 
**User** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 

## Methods

### NewUserProfileView

`func NewUserProfileView() *UserProfileView`

NewUserProfileView instantiates a new UserProfileView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileViewWithDefaults

`func NewUserProfileViewWithDefaults() *UserProfileView`

NewUserProfileViewWithDefaults instantiates a new UserProfileView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAccount

`func (o *UserProfileView) GetBankAccount() BankAccountView`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *UserProfileView) GetBankAccountOk() (*BankAccountView, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *UserProfileView) SetBankAccount(v BankAccountView)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *UserProfileView) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### GetCashTransferLogs

`func (o *UserProfileView) GetCashTransferLogs() []CashTransferLogEntryView`

GetCashTransferLogs returns the CashTransferLogs field if non-nil, zero value otherwise.

### GetCashTransferLogsOk

`func (o *UserProfileView) GetCashTransferLogsOk() (*[]CashTransferLogEntryView, bool)`

GetCashTransferLogsOk returns a tuple with the CashTransferLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashTransferLogs

`func (o *UserProfileView) SetCashTransferLogs(v []CashTransferLogEntryView)`

SetCashTransferLogs sets CashTransferLogs field to given value.

### HasCashTransferLogs

`func (o *UserProfileView) HasCashTransferLogs() bool`

HasCashTransferLogs returns a boolean if a field has been set.

### GetEmployments

`func (o *UserProfileView) GetEmployments() []EmploymentAgreementCompactCompanyView`

GetEmployments returns the Employments field if non-nil, zero value otherwise.

### GetEmploymentsOk

`func (o *UserProfileView) GetEmploymentsOk() (*[]EmploymentAgreementCompactCompanyView, bool)`

GetEmploymentsOk returns a tuple with the Employments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployments

`func (o *UserProfileView) SetEmployments(v []EmploymentAgreementCompactCompanyView)`

SetEmployments sets Employments field to given value.

### HasEmployments

`func (o *UserProfileView) HasEmployments() bool`

HasEmployments returns a boolean if a field has been set.

### GetInitiatedPolls

`func (o *UserProfileView) GetInitiatedPolls() []AbstractPollView`

GetInitiatedPolls returns the InitiatedPolls field if non-nil, zero value otherwise.

### GetInitiatedPollsOk

`func (o *UserProfileView) GetInitiatedPollsOk() (*[]AbstractPollView, bool)`

GetInitiatedPollsOk returns a tuple with the InitiatedPolls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedPolls

`func (o *UserProfileView) SetInitiatedPolls(v []AbstractPollView)`

SetInitiatedPolls sets InitiatedPolls field to given value.

### HasInitiatedPolls

`func (o *UserProfileView) HasInitiatedPolls() bool`

HasInitiatedPolls returns a boolean if a field has been set.

### GetLocale

`func (o *UserProfileView) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UserProfileView) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UserProfileView) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *UserProfileView) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetPolls

`func (o *UserProfileView) GetPolls() []AbstractPollView`

GetPolls returns the Polls field if non-nil, zero value otherwise.

### GetPollsOk

`func (o *UserProfileView) GetPollsOk() (*[]AbstractPollView, bool)`

GetPollsOk returns a tuple with the Polls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolls

`func (o *UserProfileView) SetPolls(v []AbstractPollView)`

SetPolls sets Polls field to given value.

### HasPolls

`func (o *UserProfileView) HasPolls() bool`

HasPolls returns a boolean if a field has been set.

### GetSalaryPayments

`func (o *UserProfileView) GetSalaryPayments() []SalaryPaymentView`

GetSalaryPayments returns the SalaryPayments field if non-nil, zero value otherwise.

### GetSalaryPaymentsOk

`func (o *UserProfileView) GetSalaryPaymentsOk() (*[]SalaryPaymentView, bool)`

GetSalaryPaymentsOk returns a tuple with the SalaryPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalaryPayments

`func (o *UserProfileView) SetSalaryPayments(v []SalaryPaymentView)`

SetSalaryPayments sets SalaryPayments field to given value.

### HasSalaryPayments

`func (o *UserProfileView) HasSalaryPayments() bool`

HasSalaryPayments returns a boolean if a field has been set.

### GetUser

`func (o *UserProfileView) GetUser() UsernameView`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserProfileView) GetUserOk() (*UsernameView, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserProfileView) SetUser(v UsernameView)`

SetUser sets User field to given value.

### HasUser

`func (o *UserProfileView) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


