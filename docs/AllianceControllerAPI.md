# \AllianceControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddToAllianceUsingPOST**](AllianceControllerAPI.md#AddToAllianceUsingPOST) | **Post** /api/v2/alliancememberships | Adds user to alliance
[**ChangeAllianceMembershipUsingPUT**](AllianceControllerAPI.md#ChangeAllianceMembershipUsingPUT) | **Put** /api/v2/alliancememberships/{membershipId} | Change alliance membership
[**CreateAllianceUsingPOST**](AllianceControllerAPI.md#CreateAllianceUsingPOST) | **Post** /api/v2/alliances | Creates alliance
[**DeleteAllianceUsingDELETE**](AllianceControllerAPI.md#DeleteAllianceUsingDELETE) | **Delete** /api/v2/alliances/{allianceId} | Deletes alliance
[**DeleteMyAllianceMembershipUsingDELETE**](AllianceControllerAPI.md#DeleteMyAllianceMembershipUsingDELETE) | **Delete** /api/v2/my/alliancemembership | Delete logged in user&#39;s alliance membership
[**EditAllianceUsingPUT**](AllianceControllerAPI.md#EditAllianceUsingPUT) | **Put** /api/v2/alliances/{allianceId} | Changes alliance
[**GetAllianceMembershipsUsingGET**](AllianceControllerAPI.md#GetAllianceMembershipsUsingGET) | **Get** /api/v2/alliancememberships | Lists memberships of alliance
[**GetAllianceUsingGET**](AllianceControllerAPI.md#GetAllianceUsingGET) | **Get** /api/v2/alliances/{allianceId} | Returns alliance
[**GetMyAllianceMembershipUsingGET**](AllianceControllerAPI.md#GetMyAllianceMembershipUsingGET) | **Get** /api/v2/my/alliancemembership | Returns logged in user&#39;s alliance membership
[**GetUserAllianceMembershipUsingGET**](AllianceControllerAPI.md#GetUserAllianceMembershipUsingGET) | **Get** /api/v2/alliancememberships/username | Returns user&#39;s alliance membership
[**ListAlliancesUsingGET**](AllianceControllerAPI.md#ListAlliancesUsingGET) | **Get** /api/v2/alliances | Lists alliances
[**RemoveAllianceMembershipOfAllianceUsingDELETE**](AllianceControllerAPI.md#RemoveAllianceMembershipOfAllianceUsingDELETE) | **Delete** /api/v2/alliancememberships | Removes alliance membership
[**RemoveAllianceMembershipUsingDELETE**](AllianceControllerAPI.md#RemoveAllianceMembershipUsingDELETE) | **Delete** /api/v2/alliancememberships/{membershipId} | Removes alliance membership
[**ReturnAllianceMembershipUsingGET**](AllianceControllerAPI.md#ReturnAllianceMembershipUsingGET) | **Get** /api/v2/alliancememberships/{membershipId} | Return alliance membership



## AddToAllianceUsingPOST

> AllianceMembershipView AddToAllianceUsingPOST(ctx).AllianceId(allianceId).Role(role).UserId(userId).Execute()

Adds user to alliance

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
	allianceId := "allianceId_example" // string | Alliance id
	role := "role_example" // string | Role
	userId := "userId_example" // string | User id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllianceControllerAPI.AddToAllianceUsingPOST(context.Background()).AllianceId(allianceId).Role(role).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllianceControllerAPI.AddToAllianceUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddToAllianceUsingPOST`: AllianceMembershipView
	fmt.Fprintf(os.Stdout, "Response from `AllianceControllerAPI.AddToAllianceUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddToAllianceUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allianceId** | **string** | Alliance id | 
 **role** | **string** | Role | 
 **userId** | **string** | User id | 

### Return type

[**AllianceMembershipView**](AllianceMembershipView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeAllianceMembershipUsingPUT

> AllianceMembershipView ChangeAllianceMembershipUsingPUT(ctx, membershipId).Role(role).Execute()

Change alliance membership

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
	membershipId := "membershipId_example" // string | Membership id
	role := "role_example" // string | Role (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllianceControllerAPI.ChangeAllianceMembershipUsingPUT(context.Background(), membershipId).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllianceControllerAPI.ChangeAllianceMembershipUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeAllianceMembershipUsingPUT`: AllianceMembershipView
	fmt.Fprintf(os.Stdout, "Response from `AllianceControllerAPI.ChangeAllianceMembershipUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**membershipId** | **string** | Membership id | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeAllianceMembershipUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | **string** | Role | 

### Return type

[**AllianceMembershipView**](AllianceMembershipView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAllianceUsingPOST

> AllianceView CreateAllianceUsingPOST(ctx).Description(description).Locale(locale).Name(name).LogoUrl(logoUrl).Execute()

Creates alliance

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
	description := "description_example" // string | Description
	locale := "locale_example" // string | Locale
	name := "name_example" // string | Name
	logoUrl := "logoUrl_example" // string | Logo URL (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllianceControllerAPI.CreateAllianceUsingPOST(context.Background()).Description(description).Locale(locale).Name(name).LogoUrl(logoUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllianceControllerAPI.CreateAllianceUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAllianceUsingPOST`: AllianceView
	fmt.Fprintf(os.Stdout, "Response from `AllianceControllerAPI.CreateAllianceUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAllianceUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | Description | 
 **locale** | **string** | Locale | 
 **name** | **string** | Name | 
 **logoUrl** | **string** | Logo URL | 

### Return type

[**AllianceView**](AllianceView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllianceUsingDELETE

> string DeleteAllianceUsingDELETE(ctx, allianceId).Execute()

Deletes alliance

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
	allianceId := "allianceId_example" // string | Alliance id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllianceControllerAPI.DeleteAllianceUsingDELETE(context.Background(), allianceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllianceControllerAPI.DeleteAllianceUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllianceUsingDELETE`: string
	fmt.Fprintf(os.Stdout, "Response from `AllianceControllerAPI.DeleteAllianceUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allianceId** | **string** | Alliance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllianceUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMyAllianceMembershipUsingDELETE

> AllianceMembershipView DeleteMyAllianceMembershipUsingDELETE(ctx).Execute()

Delete logged in user's alliance membership

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
	resp, r, err := apiClient.AllianceControllerAPI.DeleteMyAllianceMembershipUsingDELETE(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllianceControllerAPI.DeleteMyAllianceMembershipUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMyAllianceMembershipUsingDELETE`: AllianceMembershipView
	fmt.Fprintf(os.Stdout, "Response from `AllianceControllerAPI.DeleteMyAllianceMembershipUsingDELETE`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMyAllianceMembershipUsingDELETERequest struct via the builder pattern


### Return type

[**AllianceMembershipView**](AllianceMembershipView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditAllianceUsingPUT

> AllianceView EditAllianceUsingPUT(ctx, allianceId).Locale(locale).Name(name).Description(description).LogoUrl(logoUrl).Execute()

Changes alliance

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
	allianceId := "allianceId_example" // string | Alliance id
	locale := "locale_example" // string | Locale
	name := "name_example" // string | Name
	description := "description_example" // string | Description (optional)
	logoUrl := "logoUrl_example" // string | Logo URL (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllianceControllerAPI.EditAllianceUsingPUT(context.Background(), allianceId).Locale(locale).Name(name).Description(description).LogoUrl(logoUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllianceControllerAPI.EditAllianceUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAllianceUsingPUT`: AllianceView
	fmt.Fprintf(os.Stdout, "Response from `AllianceControllerAPI.EditAllianceUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allianceId** | **string** | Alliance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAllianceUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locale** | **string** | Locale | 
 **name** | **string** | Name | 
 **description** | **string** | Description | 
 **logoUrl** | **string** | Logo URL | 

### Return type

[**AllianceView**](AllianceView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllianceMembershipsUsingGET

> AllianceMembershipView GetAllianceMembershipsUsingGET(ctx).AllianceId(allianceId).Execute()

Lists memberships of alliance

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
	allianceId := "allianceId_example" // string | Alliance id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllianceControllerAPI.GetAllianceMembershipsUsingGET(context.Background()).AllianceId(allianceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllianceControllerAPI.GetAllianceMembershipsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllianceMembershipsUsingGET`: AllianceMembershipView
	fmt.Fprintf(os.Stdout, "Response from `AllianceControllerAPI.GetAllianceMembershipsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllianceMembershipsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allianceId** | **string** | Alliance id | 

### Return type

[**AllianceMembershipView**](AllianceMembershipView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllianceUsingGET

> string GetAllianceUsingGET(ctx, allianceId).Execute()

Returns alliance

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
	allianceId := "allianceId_example" // string | Message alliance id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllianceControllerAPI.GetAllianceUsingGET(context.Background(), allianceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllianceControllerAPI.GetAllianceUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllianceUsingGET`: string
	fmt.Fprintf(os.Stdout, "Response from `AllianceControllerAPI.GetAllianceUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allianceId** | **string** | Message alliance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllianceUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyAllianceMembershipUsingGET

> AllianceMembershipView GetMyAllianceMembershipUsingGET(ctx).Execute()

Returns logged in user's alliance membership

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
	resp, r, err := apiClient.AllianceControllerAPI.GetMyAllianceMembershipUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllianceControllerAPI.GetMyAllianceMembershipUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyAllianceMembershipUsingGET`: AllianceMembershipView
	fmt.Fprintf(os.Stdout, "Response from `AllianceControllerAPI.GetMyAllianceMembershipUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyAllianceMembershipUsingGETRequest struct via the builder pattern


### Return type

[**AllianceMembershipView**](AllianceMembershipView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAllianceMembershipUsingGET

> AllianceMembershipView GetUserAllianceMembershipUsingGET(ctx).Username(username).Execute()

Returns user's alliance membership

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
	resp, r, err := apiClient.AllianceControllerAPI.GetUserAllianceMembershipUsingGET(context.Background()).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllianceControllerAPI.GetUserAllianceMembershipUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAllianceMembershipUsingGET`: AllianceMembershipView
	fmt.Fprintf(os.Stdout, "Response from `AllianceControllerAPI.GetUserAllianceMembershipUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAllianceMembershipUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** | Username | 

### Return type

[**AllianceMembershipView**](AllianceMembershipView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlliancesUsingGET

> PageAllianceWithDetailsView ListAlliancesUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Lists alliances

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
	resp, r, err := apiClient.AllianceControllerAPI.ListAlliancesUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllianceControllerAPI.ListAlliancesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlliancesUsingGET`: PageAllianceWithDetailsView
	fmt.Fprintf(os.Stdout, "Response from `AllianceControllerAPI.ListAlliancesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAlliancesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageAllianceWithDetailsView**](PageAllianceWithDetailsView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAllianceMembershipOfAllianceUsingDELETE

> MessagePrototype RemoveAllianceMembershipOfAllianceUsingDELETE(ctx).AllianceId(allianceId).Execute()

Removes alliance membership

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
	allianceId := "allianceId_example" // string | Alliance Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllianceControllerAPI.RemoveAllianceMembershipOfAllianceUsingDELETE(context.Background()).AllianceId(allianceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllianceControllerAPI.RemoveAllianceMembershipOfAllianceUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveAllianceMembershipOfAllianceUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `AllianceControllerAPI.RemoveAllianceMembershipOfAllianceUsingDELETE`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAllianceMembershipOfAllianceUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allianceId** | **string** | Alliance Id | 

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


## RemoveAllianceMembershipUsingDELETE

> MessagePrototype RemoveAllianceMembershipUsingDELETE(ctx, membershipId).Execute()

Removes alliance membership

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
	membershipId := "membershipId_example" // string | Membership id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllianceControllerAPI.RemoveAllianceMembershipUsingDELETE(context.Background(), membershipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllianceControllerAPI.RemoveAllianceMembershipUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveAllianceMembershipUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `AllianceControllerAPI.RemoveAllianceMembershipUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**membershipId** | **string** | Membership id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAllianceMembershipUsingDELETERequest struct via the builder pattern


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


## ReturnAllianceMembershipUsingGET

> AllianceMembershipView ReturnAllianceMembershipUsingGET(ctx, membershipId).Execute()

Return alliance membership

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
	membershipId := "membershipId_example" // string | Membership id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllianceControllerAPI.ReturnAllianceMembershipUsingGET(context.Background(), membershipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllianceControllerAPI.ReturnAllianceMembershipUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnAllianceMembershipUsingGET`: AllianceMembershipView
	fmt.Fprintf(os.Stdout, "Response from `AllianceControllerAPI.ReturnAllianceMembershipUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**membershipId** | **string** | Membership id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllianceMembershipUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AllianceMembershipView**](AllianceMembershipView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

