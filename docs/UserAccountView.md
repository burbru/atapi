# UserAccountView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to **string** |  | [optional] 
**EmailSubscriptionType** | Pointer to **string** |  | [optional] 
**GravatarHash** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**JwtToken** | Pointer to **string** |  | [optional] 
**MyUser** | Pointer to **bool** |  | [optional] 
**RefId** | Pointer to **string** |  | [optional] 
**RegistrationDate** | Pointer to **int64** |  | [optional] 
**UserCapabilities** | Pointer to [**UserCapabilitiesWithPartnerIdView**](UserCapabilitiesWithPartnerIdView.md) |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewUserAccountView

`func NewUserAccountView() *UserAccountView`

NewUserAccountView instantiates a new UserAccountView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccountViewWithDefaults

`func NewUserAccountViewWithDefaults() *UserAccountView`

NewUserAccountViewWithDefaults instantiates a new UserAccountView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *UserAccountView) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *UserAccountView) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *UserAccountView) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *UserAccountView) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetEmailSubscriptionType

`func (o *UserAccountView) GetEmailSubscriptionType() string`

GetEmailSubscriptionType returns the EmailSubscriptionType field if non-nil, zero value otherwise.

### GetEmailSubscriptionTypeOk

`func (o *UserAccountView) GetEmailSubscriptionTypeOk() (*string, bool)`

GetEmailSubscriptionTypeOk returns a tuple with the EmailSubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubscriptionType

`func (o *UserAccountView) SetEmailSubscriptionType(v string)`

SetEmailSubscriptionType sets EmailSubscriptionType field to given value.

### HasEmailSubscriptionType

`func (o *UserAccountView) HasEmailSubscriptionType() bool`

HasEmailSubscriptionType returns a boolean if a field has been set.

### GetGravatarHash

`func (o *UserAccountView) GetGravatarHash() string`

GetGravatarHash returns the GravatarHash field if non-nil, zero value otherwise.

### GetGravatarHashOk

`func (o *UserAccountView) GetGravatarHashOk() (*string, bool)`

GetGravatarHashOk returns a tuple with the GravatarHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGravatarHash

`func (o *UserAccountView) SetGravatarHash(v string)`

SetGravatarHash sets GravatarHash field to given value.

### HasGravatarHash

`func (o *UserAccountView) HasGravatarHash() bool`

HasGravatarHash returns a boolean if a field has been set.

### GetId

`func (o *UserAccountView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserAccountView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserAccountView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserAccountView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJwtToken

`func (o *UserAccountView) GetJwtToken() string`

GetJwtToken returns the JwtToken field if non-nil, zero value otherwise.

### GetJwtTokenOk

`func (o *UserAccountView) GetJwtTokenOk() (*string, bool)`

GetJwtTokenOk returns a tuple with the JwtToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtToken

`func (o *UserAccountView) SetJwtToken(v string)`

SetJwtToken sets JwtToken field to given value.

### HasJwtToken

`func (o *UserAccountView) HasJwtToken() bool`

HasJwtToken returns a boolean if a field has been set.

### GetMyUser

`func (o *UserAccountView) GetMyUser() bool`

GetMyUser returns the MyUser field if non-nil, zero value otherwise.

### GetMyUserOk

`func (o *UserAccountView) GetMyUserOk() (*bool, bool)`

GetMyUserOk returns a tuple with the MyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyUser

`func (o *UserAccountView) SetMyUser(v bool)`

SetMyUser sets MyUser field to given value.

### HasMyUser

`func (o *UserAccountView) HasMyUser() bool`

HasMyUser returns a boolean if a field has been set.

### GetRefId

`func (o *UserAccountView) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *UserAccountView) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *UserAccountView) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *UserAccountView) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRegistrationDate

`func (o *UserAccountView) GetRegistrationDate() int64`

GetRegistrationDate returns the RegistrationDate field if non-nil, zero value otherwise.

### GetRegistrationDateOk

`func (o *UserAccountView) GetRegistrationDateOk() (*int64, bool)`

GetRegistrationDateOk returns a tuple with the RegistrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationDate

`func (o *UserAccountView) SetRegistrationDate(v int64)`

SetRegistrationDate sets RegistrationDate field to given value.

### HasRegistrationDate

`func (o *UserAccountView) HasRegistrationDate() bool`

HasRegistrationDate returns a boolean if a field has been set.

### GetUserCapabilities

`func (o *UserAccountView) GetUserCapabilities() UserCapabilitiesWithPartnerIdView`

GetUserCapabilities returns the UserCapabilities field if non-nil, zero value otherwise.

### GetUserCapabilitiesOk

`func (o *UserAccountView) GetUserCapabilitiesOk() (*UserCapabilitiesWithPartnerIdView, bool)`

GetUserCapabilitiesOk returns a tuple with the UserCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCapabilities

`func (o *UserAccountView) SetUserCapabilities(v UserCapabilitiesWithPartnerIdView)`

SetUserCapabilities sets UserCapabilities field to given value.

### HasUserCapabilities

`func (o *UserAccountView) HasUserCapabilities() bool`

HasUserCapabilities returns a boolean if a field has been set.

### GetUsername

`func (o *UserAccountView) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserAccountView) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserAccountView) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserAccountView) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVersion

`func (o *UserAccountView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UserAccountView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UserAccountView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UserAccountView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


