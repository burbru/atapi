# IndexView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseValue** | Pointer to **float32** |  | [optional] 
**ChainingFactor** | Pointer to **float32** |  | [optional] 
**Listing** | Pointer to [**Listing**](Listing.md) |  | [optional] 
**Members** | Pointer to [**[]IndexMemberValuesView**](IndexMemberValuesView.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NextChainingDate** | Pointer to **int64** |  | [optional] 
**Owner** | Pointer to [**UsernameView**](UsernameView.md) |  | [optional] 

## Methods

### NewIndexView

`func NewIndexView() *IndexView`

NewIndexView instantiates a new IndexView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexViewWithDefaults

`func NewIndexViewWithDefaults() *IndexView`

NewIndexViewWithDefaults instantiates a new IndexView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseValue

`func (o *IndexView) GetBaseValue() float32`

GetBaseValue returns the BaseValue field if non-nil, zero value otherwise.

### GetBaseValueOk

`func (o *IndexView) GetBaseValueOk() (*float32, bool)`

GetBaseValueOk returns a tuple with the BaseValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseValue

`func (o *IndexView) SetBaseValue(v float32)`

SetBaseValue sets BaseValue field to given value.

### HasBaseValue

`func (o *IndexView) HasBaseValue() bool`

HasBaseValue returns a boolean if a field has been set.

### GetChainingFactor

`func (o *IndexView) GetChainingFactor() float32`

GetChainingFactor returns the ChainingFactor field if non-nil, zero value otherwise.

### GetChainingFactorOk

`func (o *IndexView) GetChainingFactorOk() (*float32, bool)`

GetChainingFactorOk returns a tuple with the ChainingFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainingFactor

`func (o *IndexView) SetChainingFactor(v float32)`

SetChainingFactor sets ChainingFactor field to given value.

### HasChainingFactor

`func (o *IndexView) HasChainingFactor() bool`

HasChainingFactor returns a boolean if a field has been set.

### GetListing

`func (o *IndexView) GetListing() Listing`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *IndexView) GetListingOk() (*Listing, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *IndexView) SetListing(v Listing)`

SetListing sets Listing field to given value.

### HasListing

`func (o *IndexView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetMembers

`func (o *IndexView) GetMembers() []IndexMemberValuesView`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *IndexView) GetMembersOk() (*[]IndexMemberValuesView, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *IndexView) SetMembers(v []IndexMemberValuesView)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *IndexView) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *IndexView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IndexView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IndexView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IndexView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextChainingDate

`func (o *IndexView) GetNextChainingDate() int64`

GetNextChainingDate returns the NextChainingDate field if non-nil, zero value otherwise.

### GetNextChainingDateOk

`func (o *IndexView) GetNextChainingDateOk() (*int64, bool)`

GetNextChainingDateOk returns a tuple with the NextChainingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextChainingDate

`func (o *IndexView) SetNextChainingDate(v int64)`

SetNextChainingDate sets NextChainingDate field to given value.

### HasNextChainingDate

`func (o *IndexView) HasNextChainingDate() bool`

HasNextChainingDate returns a boolean if a field has been set.

### GetOwner

`func (o *IndexView) GetOwner() UsernameView`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IndexView) GetOwnerOk() (*UsernameView, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IndexView) SetOwner(v UsernameView)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IndexView) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


