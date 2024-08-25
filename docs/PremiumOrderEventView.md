# PremiumOrderEventView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventCreated** | Pointer to **int64** |  | [optional] 
**EventReceived** | Pointer to **int64** |  | [optional] 
**Live** | Pointer to **bool** |  | [optional] 
**ProcessedByFastSpring** | Pointer to **bool** |  | [optional] 
**ProcessedByVendor** | Pointer to **bool** |  | [optional] 
**ProcessingState** | Pointer to **string** |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewPremiumOrderEventView

`func NewPremiumOrderEventView() *PremiumOrderEventView`

NewPremiumOrderEventView instantiates a new PremiumOrderEventView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPremiumOrderEventViewWithDefaults

`func NewPremiumOrderEventViewWithDefaults() *PremiumOrderEventView`

NewPremiumOrderEventViewWithDefaults instantiates a new PremiumOrderEventView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventCreated

`func (o *PremiumOrderEventView) GetEventCreated() int64`

GetEventCreated returns the EventCreated field if non-nil, zero value otherwise.

### GetEventCreatedOk

`func (o *PremiumOrderEventView) GetEventCreatedOk() (*int64, bool)`

GetEventCreatedOk returns a tuple with the EventCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCreated

`func (o *PremiumOrderEventView) SetEventCreated(v int64)`

SetEventCreated sets EventCreated field to given value.

### HasEventCreated

`func (o *PremiumOrderEventView) HasEventCreated() bool`

HasEventCreated returns a boolean if a field has been set.

### GetEventReceived

`func (o *PremiumOrderEventView) GetEventReceived() int64`

GetEventReceived returns the EventReceived field if non-nil, zero value otherwise.

### GetEventReceivedOk

`func (o *PremiumOrderEventView) GetEventReceivedOk() (*int64, bool)`

GetEventReceivedOk returns a tuple with the EventReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReceived

`func (o *PremiumOrderEventView) SetEventReceived(v int64)`

SetEventReceived sets EventReceived field to given value.

### HasEventReceived

`func (o *PremiumOrderEventView) HasEventReceived() bool`

HasEventReceived returns a boolean if a field has been set.

### GetLive

`func (o *PremiumOrderEventView) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *PremiumOrderEventView) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *PremiumOrderEventView) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *PremiumOrderEventView) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetProcessedByFastSpring

`func (o *PremiumOrderEventView) GetProcessedByFastSpring() bool`

GetProcessedByFastSpring returns the ProcessedByFastSpring field if non-nil, zero value otherwise.

### GetProcessedByFastSpringOk

`func (o *PremiumOrderEventView) GetProcessedByFastSpringOk() (*bool, bool)`

GetProcessedByFastSpringOk returns a tuple with the ProcessedByFastSpring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedByFastSpring

`func (o *PremiumOrderEventView) SetProcessedByFastSpring(v bool)`

SetProcessedByFastSpring sets ProcessedByFastSpring field to given value.

### HasProcessedByFastSpring

`func (o *PremiumOrderEventView) HasProcessedByFastSpring() bool`

HasProcessedByFastSpring returns a boolean if a field has been set.

### GetProcessedByVendor

`func (o *PremiumOrderEventView) GetProcessedByVendor() bool`

GetProcessedByVendor returns the ProcessedByVendor field if non-nil, zero value otherwise.

### GetProcessedByVendorOk

`func (o *PremiumOrderEventView) GetProcessedByVendorOk() (*bool, bool)`

GetProcessedByVendorOk returns a tuple with the ProcessedByVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedByVendor

`func (o *PremiumOrderEventView) SetProcessedByVendor(v bool)`

SetProcessedByVendor sets ProcessedByVendor field to given value.

### HasProcessedByVendor

`func (o *PremiumOrderEventView) HasProcessedByVendor() bool`

HasProcessedByVendor returns a boolean if a field has been set.

### GetProcessingState

`func (o *PremiumOrderEventView) GetProcessingState() string`

GetProcessingState returns the ProcessingState field if non-nil, zero value otherwise.

### GetProcessingStateOk

`func (o *PremiumOrderEventView) GetProcessingStateOk() (*string, bool)`

GetProcessingStateOk returns a tuple with the ProcessingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingState

`func (o *PremiumOrderEventView) SetProcessingState(v string)`

SetProcessingState sets ProcessingState field to given value.

### HasProcessingState

`func (o *PremiumOrderEventView) HasProcessingState() bool`

HasProcessingState returns a boolean if a field has been set.

### GetReference

`func (o *PremiumOrderEventView) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PremiumOrderEventView) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PremiumOrderEventView) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PremiumOrderEventView) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetType

`func (o *PremiumOrderEventView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PremiumOrderEventView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PremiumOrderEventView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PremiumOrderEventView) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


