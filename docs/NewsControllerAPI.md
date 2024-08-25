# \NewsControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNewsPostUsingGET**](NewsControllerAPI.md#GetNewsPostUsingGET) | **Get** /api/v2/news/{postId} | Returns news post
[**ListAlliancesNewsUsingGET**](NewsControllerAPI.md#ListAlliancesNewsUsingGET) | **Get** /api/v2/alliances/{allianceId}/news | Lists alliance&#39;s news
[**ListAuthorsNewsUsingGET**](NewsControllerAPI.md#ListAuthorsNewsUsingGET) | **Get** /api/v2/authors/{userId}/news | Lists author&#39;s news
[**ListCompanysNewsUsingGET**](NewsControllerAPI.md#ListCompanysNewsUsingGET) | **Get** /api/v2/companies/{companyId}/news | Lists company&#39;s news
[**ListHashtagsNewsUsingGET**](NewsControllerAPI.md#ListHashtagsNewsUsingGET) | **Get** /api/v2/hashtags/{hashTag}/news | Lists hashtag&#39;s news
[**ListNewsUsingGET**](NewsControllerAPI.md#ListNewsUsingGET) | **Get** /api/v2/news | Lists news
[**ListPostUsingGET**](NewsControllerAPI.md#ListPostUsingGET) | **Get** /api/v2/news/hot | Lists most interesting news



## GetNewsPostUsingGET

> PostView GetNewsPostUsingGET(ctx, postId).Execute()

Returns news post

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
	postId := "postId_example" // string | postId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsControllerAPI.GetNewsPostUsingGET(context.Background(), postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsControllerAPI.GetNewsPostUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNewsPostUsingGET`: PostView
	fmt.Fprintf(os.Stdout, "Response from `NewsControllerAPI.GetNewsPostUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** | postId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNewsPostUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostView**](PostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlliancesNewsUsingGET

> PagePostView ListAlliancesNewsUsingGET(ctx, allianceId).Page(page).Size(size).Sort(sort).Execute()

Lists alliance's news



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
	allianceId := "allianceId_example" // string | allianceId
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsControllerAPI.ListAlliancesNewsUsingGET(context.Background(), allianceId).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsControllerAPI.ListAlliancesNewsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlliancesNewsUsingGET`: PagePostView
	fmt.Fprintf(os.Stdout, "Response from `NewsControllerAPI.ListAlliancesNewsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allianceId** | **string** | allianceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAlliancesNewsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PagePostView**](PagePostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthorsNewsUsingGET

> PagePostView ListAuthorsNewsUsingGET(ctx, userId).Page(page).Size(size).Sort(sort).Execute()

Lists author's news



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
	userId := "userId_example" // string | userId
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsControllerAPI.ListAuthorsNewsUsingGET(context.Background(), userId).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsControllerAPI.ListAuthorsNewsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAuthorsNewsUsingGET`: PagePostView
	fmt.Fprintf(os.Stdout, "Response from `NewsControllerAPI.ListAuthorsNewsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | userId | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthorsNewsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PagePostView**](PagePostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompanysNewsUsingGET

> PagePostView ListCompanysNewsUsingGET(ctx, companyId).Page(page).Size(size).Sort(sort).Execute()

Lists company's news



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
	companyId := "companyId_example" // string | companyId
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsControllerAPI.ListCompanysNewsUsingGET(context.Background(), companyId).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsControllerAPI.ListCompanysNewsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompanysNewsUsingGET`: PagePostView
	fmt.Fprintf(os.Stdout, "Response from `NewsControllerAPI.ListCompanysNewsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompanysNewsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PagePostView**](PagePostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHashtagsNewsUsingGET

> PagePostView ListHashtagsNewsUsingGET(ctx, hashTag).Page(page).Size(size).Sort(sort).Execute()

Lists hashtag's news



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
	hashTag := "hashTag_example" // string | hashTag
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsControllerAPI.ListHashtagsNewsUsingGET(context.Background(), hashTag).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsControllerAPI.ListHashtagsNewsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHashtagsNewsUsingGET`: PagePostView
	fmt.Fprintf(os.Stdout, "Response from `NewsControllerAPI.ListHashtagsNewsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hashTag** | **string** | hashTag | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHashtagsNewsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PagePostView**](PagePostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNewsUsingGET

> PagePostView ListNewsUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Lists news



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
	resp, r, err := apiClient.NewsControllerAPI.ListNewsUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsControllerAPI.ListNewsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNewsUsingGET`: PagePostView
	fmt.Fprintf(os.Stdout, "Response from `NewsControllerAPI.ListNewsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNewsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PagePostView**](PagePostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPostUsingGET

> PostView ListPostUsingGET(ctx).Count(count).LastPostId(lastPostId).Execute()

Lists most interesting news

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
	count := int32(56) // int32 | Number of news (default: 20) (optional)
	lastPostId := "lastPostId_example" // string | ID of last seen post (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsControllerAPI.ListPostUsingGET(context.Background()).Count(count).LastPostId(lastPostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsControllerAPI.ListPostUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPostUsingGET`: PostView
	fmt.Fprintf(os.Stdout, "Response from `NewsControllerAPI.ListPostUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPostUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of news (default: 20) | 
 **lastPostId** | **string** | ID of last seen post | 

### Return type

[**PostView**](PostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

