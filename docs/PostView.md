# PostView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alliance** | Pointer to [**AllianceView**](AllianceView.md) |  | [optional] 
**Author** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 
**Comment** | Pointer to **bool** |  | [optional] 
**Company** | Pointer to [**CompactCompanyView**](CompactCompanyView.md) |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **int64** |  | [optional] 
**DateDeleted** | Pointer to **int64** |  | [optional] 
**DateEdited** | Pointer to **int64** |  | [optional] 
**HashTags** | Pointer to [**[]HashTag**](HashTag.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Listing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**ListingPostType** | Pointer to **string** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**MessageBoard** | Pointer to [**MessageBoardView**](MessageBoardView.md) |  | [optional] 
**News** | Pointer to **bool** |  | [optional] 
**NewsPostOrComment** | Pointer to **bool** |  | [optional] 
**NumberOfComments** | Pointer to **int64** |  | [optional] 
**NumberOfDislikes** | Pointer to **int64** |  | [optional] 
**NumberOfLikes** | Pointer to **int64** |  | [optional] 
**Parent** | Pointer to **string** |  | [optional] 
**Root** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewPostView

`func NewPostView() *PostView`

NewPostView instantiates a new PostView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostViewWithDefaults

`func NewPostViewWithDefaults() *PostView`

NewPostViewWithDefaults instantiates a new PostView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlliance

`func (o *PostView) GetAlliance() AllianceView`

GetAlliance returns the Alliance field if non-nil, zero value otherwise.

### GetAllianceOk

`func (o *PostView) GetAllianceOk() (*AllianceView, bool)`

GetAllianceOk returns a tuple with the Alliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlliance

`func (o *PostView) SetAlliance(v AllianceView)`

SetAlliance sets Alliance field to given value.

### HasAlliance

`func (o *PostView) HasAlliance() bool`

HasAlliance returns a boolean if a field has been set.

### GetAuthor

`func (o *PostView) GetAuthor() UsernameView`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *PostView) GetAuthorOk() (*UsernameView, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *PostView) SetAuthor(v UsernameView)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *PostView) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetComment

`func (o *PostView) GetComment() bool`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PostView) GetCommentOk() (*bool, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PostView) SetComment(v bool)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PostView) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCompany

`func (o *PostView) GetCompany() CompactCompanyView`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PostView) GetCompanyOk() (*CompactCompanyView, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PostView) SetCompany(v CompactCompanyView)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PostView) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetContent

`func (o *PostView) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PostView) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PostView) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *PostView) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDateCreated

`func (o *PostView) GetDateCreated() int64`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *PostView) GetDateCreatedOk() (*int64, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *PostView) SetDateCreated(v int64)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *PostView) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateDeleted

`func (o *PostView) GetDateDeleted() int64`

GetDateDeleted returns the DateDeleted field if non-nil, zero value otherwise.

### GetDateDeletedOk

`func (o *PostView) GetDateDeletedOk() (*int64, bool)`

GetDateDeletedOk returns a tuple with the DateDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeleted

`func (o *PostView) SetDateDeleted(v int64)`

SetDateDeleted sets DateDeleted field to given value.

### HasDateDeleted

`func (o *PostView) HasDateDeleted() bool`

HasDateDeleted returns a boolean if a field has been set.

### GetDateEdited

`func (o *PostView) GetDateEdited() int64`

GetDateEdited returns the DateEdited field if non-nil, zero value otherwise.

### GetDateEditedOk

`func (o *PostView) GetDateEditedOk() (*int64, bool)`

GetDateEditedOk returns a tuple with the DateEdited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEdited

`func (o *PostView) SetDateEdited(v int64)`

SetDateEdited sets DateEdited field to given value.

### HasDateEdited

`func (o *PostView) HasDateEdited() bool`

HasDateEdited returns a boolean if a field has been set.

### GetHashTags

`func (o *PostView) GetHashTags() []HashTag`

GetHashTags returns the HashTags field if non-nil, zero value otherwise.

### GetHashTagsOk

`func (o *PostView) GetHashTagsOk() (*[]HashTag, bool)`

GetHashTagsOk returns a tuple with the HashTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashTags

`func (o *PostView) SetHashTags(v []HashTag)`

SetHashTags sets HashTags field to given value.

### HasHashTags

`func (o *PostView) HasHashTags() bool`

HasHashTags returns a boolean if a field has been set.

### GetId

`func (o *PostView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PostView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetListing

`func (o *PostView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *PostView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *PostView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *PostView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetListingPostType

`func (o *PostView) GetListingPostType() string`

GetListingPostType returns the ListingPostType field if non-nil, zero value otherwise.

### GetListingPostTypeOk

`func (o *PostView) GetListingPostTypeOk() (*string, bool)`

GetListingPostTypeOk returns a tuple with the ListingPostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingPostType

`func (o *PostView) SetListingPostType(v string)`

SetListingPostType sets ListingPostType field to given value.

### HasListingPostType

`func (o *PostView) HasListingPostType() bool`

HasListingPostType returns a boolean if a field has been set.

### GetLocale

`func (o *PostView) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *PostView) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *PostView) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *PostView) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetMessageBoard

`func (o *PostView) GetMessageBoard() MessageBoardView`

GetMessageBoard returns the MessageBoard field if non-nil, zero value otherwise.

### GetMessageBoardOk

`func (o *PostView) GetMessageBoardOk() (*MessageBoardView, bool)`

GetMessageBoardOk returns a tuple with the MessageBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBoard

`func (o *PostView) SetMessageBoard(v MessageBoardView)`

SetMessageBoard sets MessageBoard field to given value.

### HasMessageBoard

`func (o *PostView) HasMessageBoard() bool`

HasMessageBoard returns a boolean if a field has been set.

### GetNews

`func (o *PostView) GetNews() bool`

GetNews returns the News field if non-nil, zero value otherwise.

### GetNewsOk

`func (o *PostView) GetNewsOk() (*bool, bool)`

GetNewsOk returns a tuple with the News field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNews

`func (o *PostView) SetNews(v bool)`

SetNews sets News field to given value.

### HasNews

`func (o *PostView) HasNews() bool`

HasNews returns a boolean if a field has been set.

### GetNewsPostOrComment

`func (o *PostView) GetNewsPostOrComment() bool`

GetNewsPostOrComment returns the NewsPostOrComment field if non-nil, zero value otherwise.

### GetNewsPostOrCommentOk

`func (o *PostView) GetNewsPostOrCommentOk() (*bool, bool)`

GetNewsPostOrCommentOk returns a tuple with the NewsPostOrComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewsPostOrComment

`func (o *PostView) SetNewsPostOrComment(v bool)`

SetNewsPostOrComment sets NewsPostOrComment field to given value.

### HasNewsPostOrComment

`func (o *PostView) HasNewsPostOrComment() bool`

HasNewsPostOrComment returns a boolean if a field has been set.

### GetNumberOfComments

`func (o *PostView) GetNumberOfComments() int64`

GetNumberOfComments returns the NumberOfComments field if non-nil, zero value otherwise.

### GetNumberOfCommentsOk

`func (o *PostView) GetNumberOfCommentsOk() (*int64, bool)`

GetNumberOfCommentsOk returns a tuple with the NumberOfComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfComments

`func (o *PostView) SetNumberOfComments(v int64)`

SetNumberOfComments sets NumberOfComments field to given value.

### HasNumberOfComments

`func (o *PostView) HasNumberOfComments() bool`

HasNumberOfComments returns a boolean if a field has been set.

### GetNumberOfDislikes

`func (o *PostView) GetNumberOfDislikes() int64`

GetNumberOfDislikes returns the NumberOfDislikes field if non-nil, zero value otherwise.

### GetNumberOfDislikesOk

`func (o *PostView) GetNumberOfDislikesOk() (*int64, bool)`

GetNumberOfDislikesOk returns a tuple with the NumberOfDislikes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDislikes

`func (o *PostView) SetNumberOfDislikes(v int64)`

SetNumberOfDislikes sets NumberOfDislikes field to given value.

### HasNumberOfDislikes

`func (o *PostView) HasNumberOfDislikes() bool`

HasNumberOfDislikes returns a boolean if a field has been set.

### GetNumberOfLikes

`func (o *PostView) GetNumberOfLikes() int64`

GetNumberOfLikes returns the NumberOfLikes field if non-nil, zero value otherwise.

### GetNumberOfLikesOk

`func (o *PostView) GetNumberOfLikesOk() (*int64, bool)`

GetNumberOfLikesOk returns a tuple with the NumberOfLikes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfLikes

`func (o *PostView) SetNumberOfLikes(v int64)`

SetNumberOfLikes sets NumberOfLikes field to given value.

### HasNumberOfLikes

`func (o *PostView) HasNumberOfLikes() bool`

HasNumberOfLikes returns a boolean if a field has been set.

### GetParent

`func (o *PostView) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PostView) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PostView) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PostView) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetRoot

`func (o *PostView) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *PostView) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *PostView) SetRoot(v string)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *PostView) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetTitle

`func (o *PostView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PostView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PostView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PostView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVersion

`func (o *PostView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PostView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PostView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PostView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


