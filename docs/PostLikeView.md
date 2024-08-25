# PostLikeView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**CompactUsernameView**](CompactUsernameView.md) |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewPostLikeView

`func NewPostLikeView() *PostLikeView`

NewPostLikeView instantiates a new PostLikeView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLikeViewWithDefaults

`func NewPostLikeViewWithDefaults() *PostLikeView`

NewPostLikeViewWithDefaults instantiates a new PostLikeView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostLikeView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostLikeView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostLikeView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PostLikeView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostLikeView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostLikeView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostLikeView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostLikeView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *PostLikeView) GetUser() CompactUsernameView`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PostLikeView) GetUserOk() (*CompactUsernameView, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PostLikeView) SetUser(v CompactUsernameView)`

SetUser sets User field to given value.

### HasUser

`func (o *PostLikeView) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVersion

`func (o *PostLikeView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PostLikeView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PostLikeView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PostLikeView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


