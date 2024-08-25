# \HighscoreControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllianceHighscoresUsingGET**](HighscoreControllerAPI.md#GetAllianceHighscoresUsingGET) | **Get** /api/v2/alliancehighscores | List alliance highscore entries
[**GetBestAlliancesUsingGET**](HighscoreControllerAPI.md#GetBestAlliancesUsingGET) | **Get** /api/v2/bestalliances | List best alliances (combining all alliance highscores)
[**GetBestCompaniesUsingGET**](HighscoreControllerAPI.md#GetBestCompaniesUsingGET) | **Get** /api/v2/bestcompanies | List best companies (combining all company highscores)
[**GetBestUsersUsingGET**](HighscoreControllerAPI.md#GetBestUsersUsingGET) | **Get** /api/v2/bestusers | List best users (combining all user highscores)
[**GetCompanyHighscoresUsingGET**](HighscoreControllerAPI.md#GetCompanyHighscoresUsingGET) | **Get** /api/v2/companyhighscores | List company highscore entries
[**GetUserHighscoresUsingGET**](HighscoreControllerAPI.md#GetUserHighscoresUsingGET) | **Get** /api/v2/userhighscores | List user highscore entries
[**ListHighscoreHistoryUsingGET**](HighscoreControllerAPI.md#ListHighscoreHistoryUsingGET) | **Get** /api/v2/highscorehistoryentries | List highscore history entries



## GetAllianceHighscoresUsingGET

> PageAllianceHighscoreEntryView GetAllianceHighscoresUsingGET(ctx).HighscoreType(highscoreType).Page(page).Size(size).Sort(sort).Execute()

List alliance highscore entries

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
	highscoreType := "highscoreType_example" // string | Highscore Type
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HighscoreControllerAPI.GetAllianceHighscoresUsingGET(context.Background()).HighscoreType(highscoreType).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighscoreControllerAPI.GetAllianceHighscoresUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllianceHighscoresUsingGET`: PageAllianceHighscoreEntryView
	fmt.Fprintf(os.Stdout, "Response from `HighscoreControllerAPI.GetAllianceHighscoresUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllianceHighscoresUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **highscoreType** | **string** | Highscore Type | 
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageAllianceHighscoreEntryView**](PageAllianceHighscoreEntryView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBestAlliancesUsingGET

> PageAllianceView GetBestAlliancesUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

List best alliances (combining all alliance highscores)

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
	resp, r, err := apiClient.HighscoreControllerAPI.GetBestAlliancesUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighscoreControllerAPI.GetBestAlliancesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBestAlliancesUsingGET`: PageAllianceView
	fmt.Fprintf(os.Stdout, "Response from `HighscoreControllerAPI.GetBestAlliancesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBestAlliancesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageAllianceView**](PageAllianceView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBestCompaniesUsingGET

> PageCompanyView GetBestCompaniesUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

List best companies (combining all company highscores)

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
	resp, r, err := apiClient.HighscoreControllerAPI.GetBestCompaniesUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighscoreControllerAPI.GetBestCompaniesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBestCompaniesUsingGET`: PageCompanyView
	fmt.Fprintf(os.Stdout, "Response from `HighscoreControllerAPI.GetBestCompaniesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBestCompaniesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageCompanyView**](PageCompanyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBestUsersUsingGET

> PageUsernameView GetBestUsersUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

List best users (combining all user highscores)

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
	resp, r, err := apiClient.HighscoreControllerAPI.GetBestUsersUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighscoreControllerAPI.GetBestUsersUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBestUsersUsingGET`: PageUsernameView
	fmt.Fprintf(os.Stdout, "Response from `HighscoreControllerAPI.GetBestUsersUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBestUsersUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageUsernameView**](PageUsernameView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyHighscoresUsingGET

> PageCompanyHighscoreEntryView GetCompanyHighscoresUsingGET(ctx).HighscoreType(highscoreType).Page(page).Size(size).Sort(sort).Execute()

List company highscore entries

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
	highscoreType := "highscoreType_example" // string | Highscore Type
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HighscoreControllerAPI.GetCompanyHighscoresUsingGET(context.Background()).HighscoreType(highscoreType).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighscoreControllerAPI.GetCompanyHighscoresUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyHighscoresUsingGET`: PageCompanyHighscoreEntryView
	fmt.Fprintf(os.Stdout, "Response from `HighscoreControllerAPI.GetCompanyHighscoresUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyHighscoresUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **highscoreType** | **string** | Highscore Type | 
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageCompanyHighscoreEntryView**](PageCompanyHighscoreEntryView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserHighscoresUsingGET

> PageUserHighscoreEntryView GetUserHighscoresUsingGET(ctx).HighscoreType(highscoreType).Page(page).Size(size).Sort(sort).Execute()

List user highscore entries

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
	highscoreType := "highscoreType_example" // string | Highscore Type
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HighscoreControllerAPI.GetUserHighscoresUsingGET(context.Background()).HighscoreType(highscoreType).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighscoreControllerAPI.GetUserHighscoresUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserHighscoresUsingGET`: PageUserHighscoreEntryView
	fmt.Fprintf(os.Stdout, "Response from `HighscoreControllerAPI.GetUserHighscoresUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserHighscoresUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **highscoreType** | **string** | Highscore Type | 
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageUserHighscoreEntryView**](PageUserHighscoreEntryView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHighscoreHistoryUsingGET

> PageHighscoreHistoryEntryView ListHighscoreHistoryUsingGET(ctx).EntityId(entityId).HighscoreType(highscoreType).Page(page).Size(size).Sort(sort).Execute()

List highscore history entries

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
	entityId := "entityId_example" // string | Entity ID
	highscoreType := "highscoreType_example" // string | Highscore Type
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HighscoreControllerAPI.ListHighscoreHistoryUsingGET(context.Background()).EntityId(entityId).HighscoreType(highscoreType).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighscoreControllerAPI.ListHighscoreHistoryUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHighscoreHistoryUsingGET`: PageHighscoreHistoryEntryView
	fmt.Fprintf(os.Stdout, "Response from `HighscoreControllerAPI.ListHighscoreHistoryUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHighscoreHistoryUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | Entity ID | 
 **highscoreType** | **string** | Highscore Type | 
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageHighscoreHistoryEntryView**](PageHighscoreHistoryEntryView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

