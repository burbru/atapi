# SecurityOrderCheckResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConcerningParams** | Pointer to **[]string** |  | [optional] 
**Failed** | Pointer to **bool** |  | [optional] 
**Msg** | Pointer to [**MessagePrototype**](MessagePrototype.md) |  | [optional] 
**Ok** | Pointer to **bool** |  | [optional] 

## Methods

### NewSecurityOrderCheckResult

`func NewSecurityOrderCheckResult() *SecurityOrderCheckResult`

NewSecurityOrderCheckResult instantiates a new SecurityOrderCheckResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityOrderCheckResultWithDefaults

`func NewSecurityOrderCheckResultWithDefaults() *SecurityOrderCheckResult`

NewSecurityOrderCheckResultWithDefaults instantiates a new SecurityOrderCheckResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConcerningParams

`func (o *SecurityOrderCheckResult) GetConcerningParams() []string`

GetConcerningParams returns the ConcerningParams field if non-nil, zero value otherwise.

### GetConcerningParamsOk

`func (o *SecurityOrderCheckResult) GetConcerningParamsOk() (*[]string, bool)`

GetConcerningParamsOk returns a tuple with the ConcerningParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcerningParams

`func (o *SecurityOrderCheckResult) SetConcerningParams(v []string)`

SetConcerningParams sets ConcerningParams field to given value.

### HasConcerningParams

`func (o *SecurityOrderCheckResult) HasConcerningParams() bool`

HasConcerningParams returns a boolean if a field has been set.

### GetFailed

`func (o *SecurityOrderCheckResult) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *SecurityOrderCheckResult) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *SecurityOrderCheckResult) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *SecurityOrderCheckResult) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetMsg

`func (o *SecurityOrderCheckResult) GetMsg() MessagePrototype`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *SecurityOrderCheckResult) GetMsgOk() (*MessagePrototype, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *SecurityOrderCheckResult) SetMsg(v MessagePrototype)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *SecurityOrderCheckResult) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetOk

`func (o *SecurityOrderCheckResult) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *SecurityOrderCheckResult) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *SecurityOrderCheckResult) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *SecurityOrderCheckResult) HasOk() bool`

HasOk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


