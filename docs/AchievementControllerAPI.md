# \AchievementControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClaimAllianceAchievementUsingPUT**](AchievementControllerAPI.md#ClaimAllianceAchievementUsingPUT) | **Put** /api/v2/my/allianceachievementclaim/{achievementId} | Claim alliance&#39;s achievement
[**ClaimAllianceAchievementsUsingPUT**](AchievementControllerAPI.md#ClaimAllianceAchievementsUsingPUT) | **Put** /api/v2/my/allianceachievementclaim | Claim alliance&#39;s achievements
[**ClaimCompanyAchievementUsingPUT**](AchievementControllerAPI.md#ClaimCompanyAchievementUsingPUT) | **Put** /api/v2/my/companyachievementclaim/{achievementId} | Claim company&#39;s achievement
[**ClaimCompanyAchievementsUsingPUT**](AchievementControllerAPI.md#ClaimCompanyAchievementsUsingPUT) | **Put** /api/v2/my/companyachievementclaim | Claim company&#39;s achievements
[**ClaimUserAchievementUsingPUT**](AchievementControllerAPI.md#ClaimUserAchievementUsingPUT) | **Put** /api/v2/my/userachievementclaim/{achievementId} | Claim user&#39;s achievement
[**ClaimUserAchievementsUsingPUT**](AchievementControllerAPI.md#ClaimUserAchievementsUsingPUT) | **Put** /api/v2/my/userachievementclaim | Claim user&#39;s achievements
[**GetAllianceAchievementsUsingGET**](AchievementControllerAPI.md#GetAllianceAchievementsUsingGET) | **Get** /api/v2/allianceachievements/{allianceID} | Returns alliance&#39;s achievements
[**GetAllianceNotYetAchievementsUsingGET**](AchievementControllerAPI.md#GetAllianceNotYetAchievementsUsingGET) | **Get** /api/v2/my/notyetclaimedallianceachievements | Returns alliance&#39;s not yet claimed achievements
[**GetCompanyAchievementsUsingGET**](AchievementControllerAPI.md#GetCompanyAchievementsUsingGET) | **Get** /api/v2/companyachievements/{companyId} | Returns company&#39;s achievements
[**GetCompanyNotYetAchievementsUsingGET**](AchievementControllerAPI.md#GetCompanyNotYetAchievementsUsingGET) | **Get** /api/v2/my/notyetclaimedcompanyachievements | Returns company&#39;s not yet claimed achievements
[**GetOpenAchievementsUsingGET**](AchievementControllerAPI.md#GetOpenAchievementsUsingGET) | **Get** /api/v2/userachievementprogress/{username} | Returns user&#39;s progress of not yet reached achievements
[**GetOpenAllianceAchievementsUsingGET**](AchievementControllerAPI.md#GetOpenAllianceAchievementsUsingGET) | **Get** /api/v2/allianceachievementprogress/{allianceId} | Returns alliance&#39;s progress of not yet reached achievements
[**GetOpenCompanyAchievementsUsingGET**](AchievementControllerAPI.md#GetOpenCompanyAchievementsUsingGET) | **Get** /api/v2/companyachievementprogress/{companyId} | Returns company&#39;s progress of not yet reached achievements
[**GetUserAchievementsUsingGET**](AchievementControllerAPI.md#GetUserAchievementsUsingGET) | **Get** /api/v2/userachievements/{username} | Returns user&#39;s achievements
[**GetUserNotYetAchievementsUsingGET**](AchievementControllerAPI.md#GetUserNotYetAchievementsUsingGET) | **Get** /api/v2/my/notyetclaimeduserachievements | Returns user&#39;s not yet claimed achievements



## ClaimAllianceAchievementUsingPUT

> AllianceAchievementView ClaimAllianceAchievementUsingPUT(ctx, achievementId).Execute()

Claim alliance's achievement

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
	achievementId := "achievementId_example" // string | ID of alliance's achievement

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AchievementControllerAPI.ClaimAllianceAchievementUsingPUT(context.Background(), achievementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchievementControllerAPI.ClaimAllianceAchievementUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClaimAllianceAchievementUsingPUT`: AllianceAchievementView
	fmt.Fprintf(os.Stdout, "Response from `AchievementControllerAPI.ClaimAllianceAchievementUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**achievementId** | **string** | ID of alliance&#39;s achievement | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimAllianceAchievementUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AllianceAchievementView**](AllianceAchievementView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClaimAllianceAchievementsUsingPUT

> AllianceAchievementView ClaimAllianceAchievementsUsingPUT(ctx).Execute()

Claim alliance's achievements

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
	resp, r, err := apiClient.AchievementControllerAPI.ClaimAllianceAchievementsUsingPUT(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchievementControllerAPI.ClaimAllianceAchievementsUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClaimAllianceAchievementsUsingPUT`: AllianceAchievementView
	fmt.Fprintf(os.Stdout, "Response from `AchievementControllerAPI.ClaimAllianceAchievementsUsingPUT`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClaimAllianceAchievementsUsingPUTRequest struct via the builder pattern


### Return type

[**AllianceAchievementView**](AllianceAchievementView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClaimCompanyAchievementUsingPUT

> CompanyAchievementView ClaimCompanyAchievementUsingPUT(ctx, achievementId).Execute()

Claim company's achievement

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
	achievementId := "achievementId_example" // string | ID of company's achievement

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AchievementControllerAPI.ClaimCompanyAchievementUsingPUT(context.Background(), achievementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchievementControllerAPI.ClaimCompanyAchievementUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClaimCompanyAchievementUsingPUT`: CompanyAchievementView
	fmt.Fprintf(os.Stdout, "Response from `AchievementControllerAPI.ClaimCompanyAchievementUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**achievementId** | **string** | ID of company&#39;s achievement | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimCompanyAchievementUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyAchievementView**](CompanyAchievementView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClaimCompanyAchievementsUsingPUT

> CompanyAchievementView ClaimCompanyAchievementsUsingPUT(ctx).CompanyId(companyId).Execute()

Claim company's achievements

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
	companyId := "companyId_example" // string | Company Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AchievementControllerAPI.ClaimCompanyAchievementsUsingPUT(context.Background()).CompanyId(companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchievementControllerAPI.ClaimCompanyAchievementsUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClaimCompanyAchievementsUsingPUT`: CompanyAchievementView
	fmt.Fprintf(os.Stdout, "Response from `AchievementControllerAPI.ClaimCompanyAchievementsUsingPUT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClaimCompanyAchievementsUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company Id | 

### Return type

[**CompanyAchievementView**](CompanyAchievementView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClaimUserAchievementUsingPUT

> UserAchievementView ClaimUserAchievementUsingPUT(ctx, achievementId).Execute()

Claim user's achievement

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
	achievementId := "achievementId_example" // string | ID of user's achievement

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AchievementControllerAPI.ClaimUserAchievementUsingPUT(context.Background(), achievementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchievementControllerAPI.ClaimUserAchievementUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClaimUserAchievementUsingPUT`: UserAchievementView
	fmt.Fprintf(os.Stdout, "Response from `AchievementControllerAPI.ClaimUserAchievementUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**achievementId** | **string** | ID of user&#39;s achievement | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimUserAchievementUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserAchievementView**](UserAchievementView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClaimUserAchievementsUsingPUT

> UserAchievementView ClaimUserAchievementsUsingPUT(ctx).Execute()

Claim user's achievements

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
	resp, r, err := apiClient.AchievementControllerAPI.ClaimUserAchievementsUsingPUT(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchievementControllerAPI.ClaimUserAchievementsUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClaimUserAchievementsUsingPUT`: UserAchievementView
	fmt.Fprintf(os.Stdout, "Response from `AchievementControllerAPI.ClaimUserAchievementsUsingPUT`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClaimUserAchievementsUsingPUTRequest struct via the builder pattern


### Return type

[**UserAchievementView**](UserAchievementView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllianceAchievementsUsingGET

> AllianceAchievementView GetAllianceAchievementsUsingGET(ctx, allianceID).Execute()

Returns alliance's achievements

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
	allianceID := "allianceID_example" // string | Alliance ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AchievementControllerAPI.GetAllianceAchievementsUsingGET(context.Background(), allianceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchievementControllerAPI.GetAllianceAchievementsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllianceAchievementsUsingGET`: AllianceAchievementView
	fmt.Fprintf(os.Stdout, "Response from `AchievementControllerAPI.GetAllianceAchievementsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allianceID** | **string** | Alliance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllianceAchievementsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AllianceAchievementView**](AllianceAchievementView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllianceNotYetAchievementsUsingGET

> AllianceAchievementView GetAllianceNotYetAchievementsUsingGET(ctx).Execute()

Returns alliance's not yet claimed achievements

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
	resp, r, err := apiClient.AchievementControllerAPI.GetAllianceNotYetAchievementsUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchievementControllerAPI.GetAllianceNotYetAchievementsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllianceNotYetAchievementsUsingGET`: AllianceAchievementView
	fmt.Fprintf(os.Stdout, "Response from `AchievementControllerAPI.GetAllianceNotYetAchievementsUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllianceNotYetAchievementsUsingGETRequest struct via the builder pattern


### Return type

[**AllianceAchievementView**](AllianceAchievementView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyAchievementsUsingGET

> CompanyAchievementView GetCompanyAchievementsUsingGET(ctx, companyId).Execute()

Returns company's achievements

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
	companyId := "companyId_example" // string | Company ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AchievementControllerAPI.GetCompanyAchievementsUsingGET(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchievementControllerAPI.GetCompanyAchievementsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyAchievementsUsingGET`: CompanyAchievementView
	fmt.Fprintf(os.Stdout, "Response from `AchievementControllerAPI.GetCompanyAchievementsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Company ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyAchievementsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyAchievementView**](CompanyAchievementView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyNotYetAchievementsUsingGET

> CompanyAchievementView GetCompanyNotYetAchievementsUsingGET(ctx).CompanyId(companyId).Execute()

Returns company's not yet claimed achievements

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
	companyId := "companyId_example" // string | Company ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AchievementControllerAPI.GetCompanyNotYetAchievementsUsingGET(context.Background()).CompanyId(companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchievementControllerAPI.GetCompanyNotYetAchievementsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyNotYetAchievementsUsingGET`: CompanyAchievementView
	fmt.Fprintf(os.Stdout, "Response from `AchievementControllerAPI.GetCompanyNotYetAchievementsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyNotYetAchievementsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company ID | 

### Return type

[**CompanyAchievementView**](CompanyAchievementView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenAchievementsUsingGET

> UserAchievementProgressView GetOpenAchievementsUsingGET(ctx, username).Execute()

Returns user's progress of not yet reached achievements

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
	resp, r, err := apiClient.AchievementControllerAPI.GetOpenAchievementsUsingGET(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchievementControllerAPI.GetOpenAchievementsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpenAchievementsUsingGET`: UserAchievementProgressView
	fmt.Fprintf(os.Stdout, "Response from `AchievementControllerAPI.GetOpenAchievementsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenAchievementsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserAchievementProgressView**](UserAchievementProgressView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenAllianceAchievementsUsingGET

> AllianceAchievementProgressView GetOpenAllianceAchievementsUsingGET(ctx, allianceId).Execute()

Returns alliance's progress of not yet reached achievements

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
	allianceId := "allianceId_example" // string | Alliance ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AchievementControllerAPI.GetOpenAllianceAchievementsUsingGET(context.Background(), allianceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchievementControllerAPI.GetOpenAllianceAchievementsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpenAllianceAchievementsUsingGET`: AllianceAchievementProgressView
	fmt.Fprintf(os.Stdout, "Response from `AchievementControllerAPI.GetOpenAllianceAchievementsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allianceId** | **string** | Alliance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenAllianceAchievementsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AllianceAchievementProgressView**](AllianceAchievementProgressView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenCompanyAchievementsUsingGET

> CompanyAchievementProgressView GetOpenCompanyAchievementsUsingGET(ctx, companyId).Execute()

Returns company's progress of not yet reached achievements

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
	companyId := "companyId_example" // string | Company ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AchievementControllerAPI.GetOpenCompanyAchievementsUsingGET(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchievementControllerAPI.GetOpenCompanyAchievementsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpenCompanyAchievementsUsingGET`: CompanyAchievementProgressView
	fmt.Fprintf(os.Stdout, "Response from `AchievementControllerAPI.GetOpenCompanyAchievementsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Company ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenCompanyAchievementsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyAchievementProgressView**](CompanyAchievementProgressView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAchievementsUsingGET

> UserAchievementView GetUserAchievementsUsingGET(ctx, username).Execute()

Returns user's achievements

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
	resp, r, err := apiClient.AchievementControllerAPI.GetUserAchievementsUsingGET(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchievementControllerAPI.GetUserAchievementsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAchievementsUsingGET`: UserAchievementView
	fmt.Fprintf(os.Stdout, "Response from `AchievementControllerAPI.GetUserAchievementsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAchievementsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserAchievementView**](UserAchievementView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserNotYetAchievementsUsingGET

> UserAchievementView GetUserNotYetAchievementsUsingGET(ctx).Execute()

Returns user's not yet claimed achievements

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
	resp, r, err := apiClient.AchievementControllerAPI.GetUserNotYetAchievementsUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchievementControllerAPI.GetUserNotYetAchievementsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserNotYetAchievementsUsingGET`: UserAchievementView
	fmt.Fprintf(os.Stdout, "Response from `AchievementControllerAPI.GetUserNotYetAchievementsUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserNotYetAchievementsUsingGETRequest struct via the builder pattern


### Return type

[**UserAchievementView**](UserAchievementView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

