# MessageBoardWithDetailsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LatestPost** | Pointer to [**PostView**](PostView.md) |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumberOfMembers** | Pointer to **int64** |  | [optional] 
**NumberOfPosts** | Pointer to **int64** |  | [optional] 
**NumberOfSubboards** | Pointer to **int64** |  | [optional] 
**Owner** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**Parent** | Pointer to [**MessageBoardView**](MessageBoardView.md) |  | [optional] 
**PublicBoard** | Pointer to **bool** |  | [optional] 
**Root** | Pointer to [**MessageBoardView**](MessageBoardView.md) |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewMessageBoardWithDetailsView

`func NewMessageBoardWithDetailsView() *MessageBoardWithDetailsView`

NewMessageBoardWithDetailsView instantiates a new MessageBoardWithDetailsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageBoardWithDetailsViewWithDefaults

`func NewMessageBoardWithDetailsViewWithDefaults() *MessageBoardWithDetailsView`

NewMessageBoardWithDetailsViewWithDefaults instantiates a new MessageBoardWithDetailsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *MessageBoardWithDetailsView) GetDateCreated() int64`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *MessageBoardWithDetailsView) GetDateCreatedOk() (*int64, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *MessageBoardWithDetailsView) SetDateCreated(v int64)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *MessageBoardWithDetailsView) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDescription

`func (o *MessageBoardWithDetailsView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MessageBoardWithDetailsView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MessageBoardWithDetailsView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MessageBoardWithDetailsView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *MessageBoardWithDetailsView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageBoardWithDetailsView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageBoardWithDetailsView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MessageBoardWithDetailsView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatestPost

`func (o *MessageBoardWithDetailsView) GetLatestPost() PostView`

GetLatestPost returns the LatestPost field if non-nil, zero value otherwise.

### GetLatestPostOk

`func (o *MessageBoardWithDetailsView) GetLatestPostOk() (*PostView, bool)`

GetLatestPostOk returns a tuple with the LatestPost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestPost

`func (o *MessageBoardWithDetailsView) SetLatestPost(v PostView)`

SetLatestPost sets LatestPost field to given value.

### HasLatestPost

`func (o *MessageBoardWithDetailsView) HasLatestPost() bool`

HasLatestPost returns a boolean if a field has been set.

### GetLocale

`func (o *MessageBoardWithDetailsView) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *MessageBoardWithDetailsView) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *MessageBoardWithDetailsView) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *MessageBoardWithDetailsView) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetName

`func (o *MessageBoardWithDetailsView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MessageBoardWithDetailsView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MessageBoardWithDetailsView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MessageBoardWithDetailsView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfMembers

`func (o *MessageBoardWithDetailsView) GetNumberOfMembers() int64`

GetNumberOfMembers returns the NumberOfMembers field if non-nil, zero value otherwise.

### GetNumberOfMembersOk

`func (o *MessageBoardWithDetailsView) GetNumberOfMembersOk() (*int64, bool)`

GetNumberOfMembersOk returns a tuple with the NumberOfMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfMembers

`func (o *MessageBoardWithDetailsView) SetNumberOfMembers(v int64)`

SetNumberOfMembers sets NumberOfMembers field to given value.

### HasNumberOfMembers

`func (o *MessageBoardWithDetailsView) HasNumberOfMembers() bool`

HasNumberOfMembers returns a boolean if a field has been set.

### GetNumberOfPosts

`func (o *MessageBoardWithDetailsView) GetNumberOfPosts() int64`

GetNumberOfPosts returns the NumberOfPosts field if non-nil, zero value otherwise.

### GetNumberOfPostsOk

`func (o *MessageBoardWithDetailsView) GetNumberOfPostsOk() (*int64, bool)`

GetNumberOfPostsOk returns a tuple with the NumberOfPosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPosts

`func (o *MessageBoardWithDetailsView) SetNumberOfPosts(v int64)`

SetNumberOfPosts sets NumberOfPosts field to given value.

### HasNumberOfPosts

`func (o *MessageBoardWithDetailsView) HasNumberOfPosts() bool`

HasNumberOfPosts returns a boolean if a field has been set.

### GetNumberOfSubboards

`func (o *MessageBoardWithDetailsView) GetNumberOfSubboards() int64`

GetNumberOfSubboards returns the NumberOfSubboards field if non-nil, zero value otherwise.

### GetNumberOfSubboardsOk

`func (o *MessageBoardWithDetailsView) GetNumberOfSubboardsOk() (*int64, bool)`

GetNumberOfSubboardsOk returns a tuple with the NumberOfSubboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSubboards

`func (o *MessageBoardWithDetailsView) SetNumberOfSubboards(v int64)`

SetNumberOfSubboards sets NumberOfSubboards field to given value.

### HasNumberOfSubboards

`func (o *MessageBoardWithDetailsView) HasNumberOfSubboards() bool`

HasNumberOfSubboards returns a boolean if a field has been set.

### GetOwner

`func (o *MessageBoardWithDetailsView) GetOwner() UsernameView`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MessageBoardWithDetailsView) GetOwnerOk() (*UsernameView, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MessageBoardWithDetailsView) SetOwner(v UsernameView)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *MessageBoardWithDetailsView) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParent

`func (o *MessageBoardWithDetailsView) GetParent() MessageBoardView`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *MessageBoardWithDetailsView) GetParentOk() (*MessageBoardView, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *MessageBoardWithDetailsView) SetParent(v MessageBoardView)`

SetParent sets Parent field to given value.

### HasParent

`func (o *MessageBoardWithDetailsView) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPublicBoard

`func (o *MessageBoardWithDetailsView) GetPublicBoard() bool`

GetPublicBoard returns the PublicBoard field if non-nil, zero value otherwise.

### GetPublicBoardOk

`func (o *MessageBoardWithDetailsView) GetPublicBoardOk() (*bool, bool)`

GetPublicBoardOk returns a tuple with the PublicBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicBoard

`func (o *MessageBoardWithDetailsView) SetPublicBoard(v bool)`

SetPublicBoard sets PublicBoard field to given value.

### HasPublicBoard

`func (o *MessageBoardWithDetailsView) HasPublicBoard() bool`

HasPublicBoard returns a boolean if a field has been set.

### GetRoot

`func (o *MessageBoardWithDetailsView) GetRoot() MessageBoardView`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *MessageBoardWithDetailsView) GetRootOk() (*MessageBoardView, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *MessageBoardWithDetailsView) SetRoot(v MessageBoardView)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *MessageBoardWithDetailsView) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetVersion

`func (o *MessageBoardWithDetailsView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MessageBoardWithDetailsView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MessageBoardWithDetailsView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MessageBoardWithDetailsView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


