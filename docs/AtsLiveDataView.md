# AtsLiveDataView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to [**[]AtsLiveDataOrderView**](AtsLiveDataOrderView.md) |  | [optional] 
**OrdersLog** | Pointer to [**[]AtsLiveDataOrderView**](AtsLiveDataOrderView.md) |  | [optional] 

## Methods

### NewAtsLiveDataView

`func NewAtsLiveDataView() *AtsLiveDataView`

NewAtsLiveDataView instantiates a new AtsLiveDataView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtsLiveDataViewWithDefaults

`func NewAtsLiveDataViewWithDefaults() *AtsLiveDataView`

NewAtsLiveDataViewWithDefaults instantiates a new AtsLiveDataView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *AtsLiveDataView) GetOrders() []AtsLiveDataOrderView`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *AtsLiveDataView) GetOrdersOk() (*[]AtsLiveDataOrderView, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *AtsLiveDataView) SetOrders(v []AtsLiveDataOrderView)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *AtsLiveDataView) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetOrdersLog

`func (o *AtsLiveDataView) GetOrdersLog() []AtsLiveDataOrderView`

GetOrdersLog returns the OrdersLog field if non-nil, zero value otherwise.

### GetOrdersLogOk

`func (o *AtsLiveDataView) GetOrdersLogOk() (*[]AtsLiveDataOrderView, bool)`

GetOrdersLogOk returns a tuple with the OrdersLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersLog

`func (o *AtsLiveDataView) SetOrdersLog(v []AtsLiveDataOrderView)`

SetOrdersLog sets OrdersLog field to given value.

### HasOrdersLog

`func (o *AtsLiveDataView) HasOrdersLog() bool`

HasOrdersLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


