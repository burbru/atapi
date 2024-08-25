# SecuritySponsorshipView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesignatedSponsor** | Pointer to [**CompactCompanyView**](CompactCompanyView.md) |  | [optional] 
**Listing** | Pointer to [**ListingView**](ListingView.md) |  | [optional] 
**SponsorRating** | Pointer to [**DesignatedSponsorRatingView**](DesignatedSponsorRatingView.md) |  | [optional] 

## Methods

### NewSecuritySponsorshipView

`func NewSecuritySponsorshipView() *SecuritySponsorshipView`

NewSecuritySponsorshipView instantiates a new SecuritySponsorshipView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySponsorshipViewWithDefaults

`func NewSecuritySponsorshipViewWithDefaults() *SecuritySponsorshipView`

NewSecuritySponsorshipViewWithDefaults instantiates a new SecuritySponsorshipView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesignatedSponsor

`func (o *SecuritySponsorshipView) GetDesignatedSponsor() CompactCompanyView`

GetDesignatedSponsor returns the DesignatedSponsor field if non-nil, zero value otherwise.

### GetDesignatedSponsorOk

`func (o *SecuritySponsorshipView) GetDesignatedSponsorOk() (*CompactCompanyView, bool)`

GetDesignatedSponsorOk returns a tuple with the DesignatedSponsor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignatedSponsor

`func (o *SecuritySponsorshipView) SetDesignatedSponsor(v CompactCompanyView)`

SetDesignatedSponsor sets DesignatedSponsor field to given value.

### HasDesignatedSponsor

`func (o *SecuritySponsorshipView) HasDesignatedSponsor() bool`

HasDesignatedSponsor returns a boolean if a field has been set.

### GetListing

`func (o *SecuritySponsorshipView) GetListing() ListingView`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *SecuritySponsorshipView) GetListingOk() (*ListingView, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *SecuritySponsorshipView) SetListing(v ListingView)`

SetListing sets Listing field to given value.

### HasListing

`func (o *SecuritySponsorshipView) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetSponsorRating

`func (o *SecuritySponsorshipView) GetSponsorRating() DesignatedSponsorRatingView`

GetSponsorRating returns the SponsorRating field if non-nil, zero value otherwise.

### GetSponsorRatingOk

`func (o *SecuritySponsorshipView) GetSponsorRatingOk() (*DesignatedSponsorRatingView, bool)`

GetSponsorRatingOk returns a tuple with the SponsorRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorRating

`func (o *SecuritySponsorshipView) SetSponsorRating(v DesignatedSponsorRatingView)`

SetSponsorRating sets SponsorRating field to given value.

### HasSponsorRating

`func (o *SecuritySponsorshipView) HasSponsorRating() bool`

HasSponsorRating returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


