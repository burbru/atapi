# SponsorView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hours** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewSponsorView

`func NewSponsorView() *SponsorView`

NewSponsorView instantiates a new SponsorView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSponsorViewWithDefaults

`func NewSponsorViewWithDefaults() *SponsorView`

NewSponsorViewWithDefaults instantiates a new SponsorView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHours

`func (o *SponsorView) GetHours() int64`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *SponsorView) GetHoursOk() (*int64, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *SponsorView) SetHours(v int64)`

SetHours sets Hours field to given value.

### HasHours

`func (o *SponsorView) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetId

`func (o *SponsorView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SponsorView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SponsorView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SponsorView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUser

`func (o *SponsorView) GetUser() UsernameView`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SponsorView) GetUserOk() (*UsernameView, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SponsorView) SetUser(v UsernameView)`

SetUser sets User field to given value.

### HasUser

`func (o *SponsorView) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVersion

`func (o *SponsorView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SponsorView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SponsorView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SponsorView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


