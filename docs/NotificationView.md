# NotificationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**MessagePrototype**](MessagePrototype.md) |  | [optional] 
**Date** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ReadByReceiver** | Pointer to **bool** |  | [optional] 
**Receiver** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**Subject** | Pointer to [**MessagePrototype**](MessagePrototype.md) |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewNotificationView

`func NewNotificationView() *NotificationView`

NewNotificationView instantiates a new NotificationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationViewWithDefaults

`func NewNotificationViewWithDefaults() *NotificationView`

NewNotificationViewWithDefaults instantiates a new NotificationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *NotificationView) GetContent() MessagePrototype`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NotificationView) GetContentOk() (*MessagePrototype, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NotificationView) SetContent(v MessagePrototype)`

SetContent sets Content field to given value.

### HasContent

`func (o *NotificationView) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDate

`func (o *NotificationView) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *NotificationView) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *NotificationView) SetDate(v int64)`

SetDate sets Date field to given value.

### HasDate

`func (o *NotificationView) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetId

`func (o *NotificationView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReadByReceiver

`func (o *NotificationView) GetReadByReceiver() bool`

GetReadByReceiver returns the ReadByReceiver field if non-nil, zero value otherwise.

### GetReadByReceiverOk

`func (o *NotificationView) GetReadByReceiverOk() (*bool, bool)`

GetReadByReceiverOk returns a tuple with the ReadByReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadByReceiver

`func (o *NotificationView) SetReadByReceiver(v bool)`

SetReadByReceiver sets ReadByReceiver field to given value.

### HasReadByReceiver

`func (o *NotificationView) HasReadByReceiver() bool`

HasReadByReceiver returns a boolean if a field has been set.

### GetReceiver

`func (o *NotificationView) GetReceiver() UsernameView`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *NotificationView) GetReceiverOk() (*UsernameView, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *NotificationView) SetReceiver(v UsernameView)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *NotificationView) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetSubject

`func (o *NotificationView) GetSubject() MessagePrototype`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *NotificationView) GetSubjectOk() (*MessagePrototype, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *NotificationView) SetSubject(v MessagePrototype)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *NotificationView) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetVersion

`func (o *NotificationView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NotificationView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NotificationView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NotificationView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


