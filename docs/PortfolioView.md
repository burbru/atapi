# PortfolioView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cash** | Pointer to **float32** |  | [optional] 
**CommittedCash** | Pointer to **float32** |  | [optional] 
**Positions** | Pointer to [**[]PortfolioPositionView**](PortfolioPositionView.md) |  | [optional] 
**SecuritiesAccountId** | Pointer to **string** |  | [optional] 

## Methods

### NewPortfolioView

`func NewPortfolioView() *PortfolioView`

NewPortfolioView instantiates a new PortfolioView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortfolioViewWithDefaults

`func NewPortfolioViewWithDefaults() *PortfolioView`

NewPortfolioViewWithDefaults instantiates a new PortfolioView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCash

`func (o *PortfolioView) GetCash() float32`

GetCash returns the Cash field if non-nil, zero value otherwise.

### GetCashOk

`func (o *PortfolioView) GetCashOk() (*float32, bool)`

GetCashOk returns a tuple with the Cash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCash

`func (o *PortfolioView) SetCash(v float32)`

SetCash sets Cash field to given value.

### HasCash

`func (o *PortfolioView) HasCash() bool`

HasCash returns a boolean if a field has been set.

### GetCommittedCash

`func (o *PortfolioView) GetCommittedCash() float32`

GetCommittedCash returns the CommittedCash field if non-nil, zero value otherwise.

### GetCommittedCashOk

`func (o *PortfolioView) GetCommittedCashOk() (*float32, bool)`

GetCommittedCashOk returns a tuple with the CommittedCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedCash

`func (o *PortfolioView) SetCommittedCash(v float32)`

SetCommittedCash sets CommittedCash field to given value.

### HasCommittedCash

`func (o *PortfolioView) HasCommittedCash() bool`

HasCommittedCash returns a boolean if a field has been set.

### GetPositions

`func (o *PortfolioView) GetPositions() []PortfolioPositionView`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *PortfolioView) GetPositionsOk() (*[]PortfolioPositionView, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *PortfolioView) SetPositions(v []PortfolioPositionView)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *PortfolioView) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetSecuritiesAccountId

`func (o *PortfolioView) GetSecuritiesAccountId() string`

GetSecuritiesAccountId returns the SecuritiesAccountId field if non-nil, zero value otherwise.

### GetSecuritiesAccountIdOk

`func (o *PortfolioView) GetSecuritiesAccountIdOk() (*string, bool)`

GetSecuritiesAccountIdOk returns a tuple with the SecuritiesAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritiesAccountId

`func (o *PortfolioView) SetSecuritiesAccountId(v string)`

SetSecuritiesAccountId sets SecuritiesAccountId field to given value.

### HasSecuritiesAccountId

`func (o *PortfolioView) HasSecuritiesAccountId() bool`

HasSecuritiesAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


