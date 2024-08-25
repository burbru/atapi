# \UserAccountControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateLicenseUsingPATCH**](UserAccountControllerAPI.md#ActivateLicenseUsingPATCH) | **Patch** /api/v2/premiumlicenses/{licCode} | Activates premium license
[**AddLicenseUsingPUT**](UserAccountControllerAPI.md#AddLicenseUsingPUT) | **Put** /api/v2/premiumlicenses | Add premium license
[**ChangeUserUsingPATCH**](UserAccountControllerAPI.md#ChangeUserUsingPATCH) | **Patch** /api/v2/my/user | Changes user properties
[**CreateNewUserUsingPOST**](UserAccountControllerAPI.md#CreateNewUserUsingPOST) | **Post** /user/register | Registers a new user
[**CreateScreenshotUsingPUT**](UserAccountControllerAPI.md#CreateScreenshotUsingPUT) | **Put** /api/v2/screenshot/** | Creates screenshot for given url
[**DeleteUserUsingDELETE**](UserAccountControllerAPI.md#DeleteUserUsingDELETE) | **Delete** /api/v2/my/user | Removes user with deletion token, sends email link when deletion token is not set
[**FindUserUsingGET**](UserAccountControllerAPI.md#FindUserUsingGET) | **Get** /api/search/users/{namePart} | Finds users by part of their usernames
[**GetJwtTokenUsingPOST**](UserAccountControllerAPI.md#GetJwtTokenUsingPOST) | **Post** /user/token/ | Returns User&#39;s JWT Token
[**GetLastActivityUsingGET**](UserAccountControllerAPI.md#GetLastActivityUsingGET) | **Get** /api/v2/my/onlinetracking | Returns online tracking
[**GetLicenseUsingGET**](UserAccountControllerAPI.md#GetLicenseUsingGET) | **Get** /api/v2/premiumlicenses/{licCode} | Returns premium license
[**GetLoggedInUserAccountUsingGET**](UserAccountControllerAPI.md#GetLoggedInUserAccountUsingGET) | **Get** /api/user | Returns the logged in user
[**GetOpenGraphImageForUrlUsingGET**](UserAccountControllerAPI.md#GetOpenGraphImageForUrlUsingGET) | **Get** /user/ogimage/** | getOpenGraphImageForUrl
[**GetOpenGraphImageUsingGET**](UserAccountControllerAPI.md#GetOpenGraphImageUsingGET) | **Get** /user/ogimage | getOpenGraphImage
[**GetReferrerUsingGET**](UserAccountControllerAPI.md#GetReferrerUsingGET) | **Get** /api/v2/my/referrer | Gets referrer
[**GetScreenshotUsingGET**](UserAccountControllerAPI.md#GetScreenshotUsingGET) | **Get** /user/screenshot/{username}/** | Return screenshot for given url
[**GetUserAccountByRefIdUsingGET**](UserAccountControllerAPI.md#GetUserAccountByRefIdUsingGET) | **Get** /api/users/refId/{refId} | Returns user by his/her referral ID 
[**GetUserAccountByUsernameUsingGET**](UserAccountControllerAPI.md#GetUserAccountByUsernameUsingGET) | **Get** /api/users/username/{username} | Returns user
[**GetUserAccountUsingGET**](UserAccountControllerAPI.md#GetUserAccountUsingGET) | **Get** /api/users/{userId} | Returns the user
[**GetUserProfileUsingGET**](UserAccountControllerAPI.md#GetUserProfileUsingGET) | **Get** /api/userprofiles/{username} | Returns user profile
[**GetUsersUsingGET**](UserAccountControllerAPI.md#GetUsersUsingGET) | **Get** /api/users | Lists users
[**ListAllPremiumOrderEventsUsingGET**](UserAccountControllerAPI.md#ListAllPremiumOrderEventsUsingGET) | **Get** /api/v2/premiumorderevents | Lists all logged in user&#39;s premium order events
[**ListPossibleReferrersUsingGET**](UserAccountControllerAPI.md#ListPossibleReferrersUsingGET) | **Get** /api/v2/my/possibleferrers/{namePart} | Lists all possible referrers
[**ListReferredUsersUsingGET**](UserAccountControllerAPI.md#ListReferredUsersUsingGET) | **Get** /api/v2/my/referredusers | Lists all users referred by logged in user
[**ListTeamMembersUsingGET**](UserAccountControllerAPI.md#ListTeamMembersUsingGET) | **Get** /api/v2/teammember | Lists all team members
[**ListUsersUsingGET**](UserAccountControllerAPI.md#ListUsersUsingGET) | **Get** /api/v2/users | Lists all users
[**PaddleWebhookUsingPOST**](UserAccountControllerAPI.md#PaddleWebhookUsingPOST) | **Post** /user/paddlewebhook | paddleWebhook
[**ResetPasswordUsingPUT**](UserAccountControllerAPI.md#ResetPasswordUsingPUT) | **Put** /user/passwordreset | Sends email with new password
[**SetReferrerUsingPATCH**](UserAccountControllerAPI.md#SetReferrerUsingPATCH) | **Patch** /api/v2/my/referrer/{refId} | Sets referrer (has to be older than logged in user)
[**StripeWebhookUsingPOST**](UserAccountControllerAPI.md#StripeWebhookUsingPOST) | **Post** /user/stripewebhook | stripeWebhook



## ActivateLicenseUsingPATCH

> PremiumLicenseView ActivateLicenseUsingPATCH(ctx, licCode).Execute()

Activates premium license

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	licCode := "licCode_example" // string | Premium license code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.ActivateLicenseUsingPATCH(context.Background(), licCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.ActivateLicenseUsingPATCH``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateLicenseUsingPATCH`: PremiumLicenseView
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.ActivateLicenseUsingPATCH`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licCode** | **string** | Premium license code | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateLicenseUsingPATCHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PremiumLicenseView**](PremiumLicenseView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLicenseUsingPUT

> PremiumLicenseView AddLicenseUsingPUT(ctx).Days(days).Reference(reference).GoalId(goalId).ReceiverId(receiverId).UserId(userId).Execute()

Add premium license

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	days := int32(56) // int32 | Days
	reference := "reference_example" // string | Reference
	goalId := "goalId_example" // string | Goal ID (optional)
	receiverId := "receiverId_example" // string | Receiver ID (optional)
	userId := "userId_example" // string | User ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.AddLicenseUsingPUT(context.Background()).Days(days).Reference(reference).GoalId(goalId).ReceiverId(receiverId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.AddLicenseUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddLicenseUsingPUT`: PremiumLicenseView
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.AddLicenseUsingPUT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLicenseUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **int32** | Days | 
 **reference** | **string** | Reference | 
 **goalId** | **string** | Goal ID | 
 **receiverId** | **string** | Receiver ID | 
 **userId** | **string** | User ID | 

### Return type

[**PremiumLicenseView**](PremiumLicenseView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeUserUsingPATCH

> MessagePrototype ChangeUserUsingPATCH(ctx).EmailSubscriptionType(emailSubscriptionType).Emailaddress(emailaddress).Password(password).Username(username).Execute()

Changes user properties



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	emailSubscriptionType := "emailSubscriptionType_example" // string | Email subscription type (optional)
	emailaddress := "emailaddress_example" // string | Email address (optional)
	password := "password_example" // string | Password (optional)
	username := "username_example" // string | Username (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.ChangeUserUsingPATCH(context.Background()).EmailSubscriptionType(emailSubscriptionType).Emailaddress(emailaddress).Password(password).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.ChangeUserUsingPATCH``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeUserUsingPATCH`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.ChangeUserUsingPATCH`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeUserUsingPATCHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailSubscriptionType** | **string** | Email subscription type | 
 **emailaddress** | **string** | Email address | 
 **password** | **string** | Password | 
 **username** | **string** | Username | 

### Return type

[**MessagePrototype**](MessagePrototype.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewUserUsingPOST

> UserAccountView CreateNewUserUsingPOST(ctx).EmailAddress(emailAddress).Password(password).Username(username).Locale(locale).Execute()

Registers a new user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	emailAddress := "emailAddress_example" // string | Email address
	password := "password_example" // string | Password
	username := "username_example" // string | Username
	locale := "locale_example" // string | Locale (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.CreateNewUserUsingPOST(context.Background()).EmailAddress(emailAddress).Password(password).Username(username).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.CreateNewUserUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewUserUsingPOST`: UserAccountView
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.CreateNewUserUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewUserUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailAddress** | **string** | Email address | 
 **password** | **string** | Password | 
 **username** | **string** | Username | 
 **locale** | **string** | Locale | 

### Return type

[**UserAccountView**](UserAccountView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateScreenshotUsingPUT

> map[string]interface{} CreateScreenshotUsingPUT(ctx).Execute()

Creates screenshot for given url

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.CreateScreenshotUsingPUT(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.CreateScreenshotUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateScreenshotUsingPUT`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.CreateScreenshotUsingPUT`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateScreenshotUsingPUTRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserUsingDELETE

> MessagePrototype DeleteUserUsingDELETE(ctx).DeletionToken(deletionToken).Execute()

Removes user with deletion token, sends email link when deletion token is not set

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	deletionToken := "deletionToken_example" // string | deletionToken (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.DeleteUserUsingDELETE(context.Background()).DeletionToken(deletionToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.DeleteUserUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.DeleteUserUsingDELETE`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deletionToken** | **string** | deletionToken | 

### Return type

[**MessagePrototype**](MessagePrototype.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindUserUsingGET

> UsernameView FindUserUsingGET(ctx, namePart).Execute()

Finds users by part of their usernames

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	namePart := "namePart_example" // string | Part of user name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.FindUserUsingGET(context.Background(), namePart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.FindUserUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindUserUsingGET`: UsernameView
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.FindUserUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namePart** | **string** | Part of user name | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindUserUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsernameView**](UsernameView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJwtTokenUsingPOST

> MessagePrototype GetJwtTokenUsingPOST(ctx).Password(password).Username(username).PartnerId(partnerId).Execute()

Returns User's JWT Token

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	password := "password_example" // string | Password
	username := "username_example" // string | Username
	partnerId := "partnerId_example" // string | Partner ID (optional, can be submitted via X-Authorization too) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.GetJwtTokenUsingPOST(context.Background()).Password(password).Username(username).PartnerId(partnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.GetJwtTokenUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJwtTokenUsingPOST`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.GetJwtTokenUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJwtTokenUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **password** | **string** | Password | 
 **username** | **string** | Username | 
 **partnerId** | **string** | Partner ID (optional, can be submitted via X-Authorization too) | 

### Return type

[**MessagePrototype**](MessagePrototype.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastActivityUsingGET

> OnlineTrackingView GetLastActivityUsingGET(ctx).Execute()

Returns online tracking

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.GetLastActivityUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.GetLastActivityUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLastActivityUsingGET`: OnlineTrackingView
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.GetLastActivityUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastActivityUsingGETRequest struct via the builder pattern


### Return type

[**OnlineTrackingView**](OnlineTrackingView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseUsingGET

> PremiumLicenseView GetLicenseUsingGET(ctx, licCode).Execute()

Returns premium license

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	licCode := "licCode_example" // string | Premium license code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.GetLicenseUsingGET(context.Background(), licCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.GetLicenseUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicenseUsingGET`: PremiumLicenseView
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.GetLicenseUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licCode** | **string** | Premium license code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PremiumLicenseView**](PremiumLicenseView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoggedInUserAccountUsingGET

> UserAccountView GetLoggedInUserAccountUsingGET(ctx).Execute()

Returns the logged in user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.GetLoggedInUserAccountUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.GetLoggedInUserAccountUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoggedInUserAccountUsingGET`: UserAccountView
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.GetLoggedInUserAccountUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoggedInUserAccountUsingGETRequest struct via the builder pattern


### Return type

[**UserAccountView**](UserAccountView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenGraphImageForUrlUsingGET

> string GetOpenGraphImageForUrlUsingGET(ctx).Execute()

getOpenGraphImageForUrl

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.GetOpenGraphImageForUrlUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.GetOpenGraphImageForUrlUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpenGraphImageForUrlUsingGET`: string
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.GetOpenGraphImageForUrlUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenGraphImageForUrlUsingGETRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenGraphImageUsingGET

> string GetOpenGraphImageUsingGET(ctx).Referer(referer).Execute()

getOpenGraphImage

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	referer := "referer_example" // string | referer (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.GetOpenGraphImageUsingGET(context.Background()).Referer(referer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.GetOpenGraphImageUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpenGraphImageUsingGET`: string
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.GetOpenGraphImageUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenGraphImageUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **referer** | **string** | referer | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReferrerUsingGET

> MessagePrototype GetReferrerUsingGET(ctx).Execute()

Gets referrer

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.GetReferrerUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.GetReferrerUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReferrerUsingGET`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.GetReferrerUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReferrerUsingGETRequest struct via the builder pattern


### Return type

[**MessagePrototype**](MessagePrototype.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScreenshotUsingGET

> string GetScreenshotUsingGET(ctx, username).Execute()

Return screenshot for given url

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	username := "username_example" // string | username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.GetScreenshotUsingGET(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.GetScreenshotUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScreenshotUsingGET`: string
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.GetScreenshotUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | username | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScreenshotUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAccountByRefIdUsingGET

> UsernameView GetUserAccountByRefIdUsingGET(ctx, refId).Execute()

Returns user by his/her referral ID 

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	refId := "refId_example" // string | Ref ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.GetUserAccountByRefIdUsingGET(context.Background(), refId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.GetUserAccountByRefIdUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAccountByRefIdUsingGET`: UsernameView
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.GetUserAccountByRefIdUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refId** | **string** | Ref ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAccountByRefIdUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsernameView**](UsernameView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAccountByUsernameUsingGET

> UsernameView GetUserAccountByUsernameUsingGET(ctx, username).Execute()

Returns user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	username := "username_example" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.GetUserAccountByUsernameUsingGET(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.GetUserAccountByUsernameUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAccountByUsernameUsingGET`: UsernameView
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.GetUserAccountByUsernameUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAccountByUsernameUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsernameView**](UsernameView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAccountUsingGET

> UsernameView GetUserAccountUsingGET(ctx, userId).Execute()

Returns the user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	userId := "userId_example" // string | User id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.GetUserAccountUsingGET(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.GetUserAccountUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAccountUsingGET`: UsernameView
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.GetUserAccountUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAccountUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsernameView**](UsernameView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserProfileUsingGET

> UserProfileView GetUserProfileUsingGET(ctx, username).Execute()

Returns user profile

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	username := "username_example" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.GetUserProfileUsingGET(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.GetUserProfileUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserProfileUsingGET`: UserProfileView
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.GetUserProfileUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserProfileUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserProfileView**](UserProfileView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersUsingGET

> UsernameView GetUsersUsingGET(ctx).Execute()

Lists users

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.GetUsersUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.GetUsersUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersUsingGET`: UsernameView
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.GetUsersUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersUsingGETRequest struct via the builder pattern


### Return type

[**UsernameView**](UsernameView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllPremiumOrderEventsUsingGET

> PagePremiumOrderEventView ListAllPremiumOrderEventsUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Lists all logged in user's premium order events

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.ListAllPremiumOrderEventsUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.ListAllPremiumOrderEventsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllPremiumOrderEventsUsingGET`: PagePremiumOrderEventView
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.ListAllPremiumOrderEventsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllPremiumOrderEventsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PagePremiumOrderEventView**](PagePremiumOrderEventView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPossibleReferrersUsingGET

> UsernameView ListPossibleReferrersUsingGET(ctx, namePart).Page(page).Size(size).Sort(sort).Execute()

Lists all possible referrers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	namePart := "namePart_example" // string | Part of user name
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.ListPossibleReferrersUsingGET(context.Background(), namePart).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.ListPossibleReferrersUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPossibleReferrersUsingGET`: UsernameView
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.ListPossibleReferrersUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namePart** | **string** | Part of user name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPossibleReferrersUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**UsernameView**](UsernameView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReferredUsersUsingGET

> UsernameView ListReferredUsersUsingGET(ctx).Page(page).Search(search).Size(size).Sort(sort).Execute()

Lists all users referred by logged in user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	search := "search_example" // string | Fulltext username search (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.ListReferredUsersUsingGET(context.Background()).Page(page).Search(search).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.ListReferredUsersUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReferredUsersUsingGET`: UsernameView
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.ListReferredUsersUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReferredUsersUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **search** | **string** | Fulltext username search | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**UsernameView**](UsernameView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeamMembersUsingGET

> UsernameView ListTeamMembersUsingGET(ctx).Execute()

Lists all team members



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.ListTeamMembersUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.ListTeamMembersUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTeamMembersUsingGET`: UsernameView
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.ListTeamMembersUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamMembersUsingGETRequest struct via the builder pattern


### Return type

[**UsernameView**](UsernameView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsersUsingGET

> UsernameView ListUsersUsingGET(ctx).Page(page).Search(search).Size(size).Sort(sort).Execute()

Lists all users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	search := "search_example" // string | Fulltext username search (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.ListUsersUsingGET(context.Background()).Page(page).Search(search).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.ListUsersUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsersUsingGET`: UsernameView
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.ListUsersUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **search** | **string** | Fulltext username search | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**UsernameView**](UsernameView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaddleWebhookUsingPOST

> string PaddleWebhookUsingPOST(ctx).PaddleSignature(paddleSignature).Content(content).Execute()

paddleWebhook

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	paddleSignature := "paddleSignature_example" // string | Paddle-Signature
	content := "content_example" // string | content

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.PaddleWebhookUsingPOST(context.Background()).PaddleSignature(paddleSignature).Content(content).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.PaddleWebhookUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaddleWebhookUsingPOST`: string
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.PaddleWebhookUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaddleWebhookUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paddleSignature** | **string** | Paddle-Signature | 
 **content** | **string** | content | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPasswordUsingPUT

> MessagePrototype ResetPasswordUsingPUT(ctx).Emailaddress(emailaddress).Execute()

Sends email with new password

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	emailaddress := "emailaddress_example" // string | Email address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.ResetPasswordUsingPUT(context.Background()).Emailaddress(emailaddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.ResetPasswordUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetPasswordUsingPUT`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.ResetPasswordUsingPUT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailaddress** | **string** | Email address | 

### Return type

[**MessagePrototype**](MessagePrototype.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetReferrerUsingPATCH

> MessagePrototype SetReferrerUsingPATCH(ctx, refId).Execute()

Sets referrer (has to be older than logged in user)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	refId := "refId_example" // string | Referrer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.SetReferrerUsingPATCH(context.Background(), refId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.SetReferrerUsingPATCH``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetReferrerUsingPATCH`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.SetReferrerUsingPATCH`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refId** | **string** | Referrer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetReferrerUsingPATCHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessagePrototype**](MessagePrototype.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StripeWebhookUsingPOST

> string StripeWebhookUsingPOST(ctx).Content(content).Execute()

stripeWebhook

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	content := "content_example" // string | content

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAccountControllerAPI.StripeWebhookUsingPOST(context.Background()).Content(content).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAccountControllerAPI.StripeWebhookUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StripeWebhookUsingPOST`: string
	fmt.Fprintf(os.Stdout, "Response from `UserAccountControllerAPI.StripeWebhookUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStripeWebhookUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **content** | **string** | content | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

