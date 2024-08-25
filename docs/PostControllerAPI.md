# \PostControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommentUsingPOST**](PostControllerAPI.md#CreateCommentUsingPOST) | **Post** /api/v2/posts/{postId}/comments | Creates comment
[**CreatePostUsingPOST**](PostControllerAPI.md#CreatePostUsingPOST) | **Post** /api/v2/posts | Creates post
[**DeletePostUsingDELETE**](PostControllerAPI.md#DeletePostUsingDELETE) | **Delete** /api/v2/posts/{postId} | Removes post
[**EditPostUsingPUT**](PostControllerAPI.md#EditPostUsingPUT) | **Put** /api/v2/posts/{postId} | Edits post
[**GetListingDescriptionUsingGET**](PostControllerAPI.md#GetListingDescriptionUsingGET) | **Get** /api/v2/listingposts | Returns listing description
[**GetPersonalInterestInPostUsingGET**](PostControllerAPI.md#GetPersonalInterestInPostUsingGET) | **Get** /api/v2/my/interests/posts/{postId} | Returns personal interest in post
[**GetPostUsingGET**](PostControllerAPI.md#GetPostUsingGET) | **Get** /api/v2/posts/{postId} | Returns post
[**ListCommentsUsingGET**](PostControllerAPI.md#ListCommentsUsingGET) | **Get** /api/v2/posts/{postId}/comments | Lists post&#39;s comments. Shows all comments for root posts and the first level comments for comment
[**ListListingPostsUsingGET**](PostControllerAPI.md#ListListingPostsUsingGET) | **Get** /api/v2/posts | List all postings regarding listing



## CreateCommentUsingPOST

> PostView CreateCommentUsingPOST(ctx, postId).Title(title).Content(content).AllianceId(allianceId).CompanyId(companyId).Locale(locale).Execute()

Creates comment

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
	title := "title_example" // string | Title
	content := "content_example" // string | content
	allianceId := "allianceId_example" // string | Alliance id (optional)
	companyId := "companyId_example" // string | Company id (optional)
	locale := "locale_example" // string | Locale (default: logged in user's locale) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostControllerAPI.CreateCommentUsingPOST(context.Background(), postId).Title(title).Content(content).AllianceId(allianceId).CompanyId(companyId).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostControllerAPI.CreateCommentUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCommentUsingPOST`: PostView
	fmt.Fprintf(os.Stdout, "Response from `PostControllerAPI.CreateCommentUsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** | postId | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommentUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **title** | **string** | Title | 
 **content** | **string** | content | 
 **allianceId** | **string** | Alliance id | 
 **companyId** | **string** | Company id | 
 **locale** | **string** | Locale (default: logged in user&#39;s locale) | 

### Return type

[**PostView**](PostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/html, application/json, text/plain
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePostUsingPOST

> PostView CreatePostUsingPOST(ctx).Title(title).Content(content).AllianceId(allianceId).Asin(asin).CompanyId(companyId).ListingPostType(listingPostType).Locale(locale).MessageBoardId(messageBoardId).Execute()

Creates post

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
	title := "title_example" // string | Title
	content := "content_example" // string | content
	allianceId := "allianceId_example" // string | Alliance id (optional)
	asin := "asin_example" // string | ASIN (optional)
	companyId := "companyId_example" // string | Company id (optional)
	listingPostType := "listingPostType_example" // string | Listing post type (optional)
	locale := "locale_example" // string | Locale (default: logged in user's locale) (optional)
	messageBoardId := "messageBoardId_example" // string | Message board id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostControllerAPI.CreatePostUsingPOST(context.Background()).Title(title).Content(content).AllianceId(allianceId).Asin(asin).CompanyId(companyId).ListingPostType(listingPostType).Locale(locale).MessageBoardId(messageBoardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostControllerAPI.CreatePostUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePostUsingPOST`: PostView
	fmt.Fprintf(os.Stdout, "Response from `PostControllerAPI.CreatePostUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePostUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **string** | Title | 
 **content** | **string** | content | 
 **allianceId** | **string** | Alliance id | 
 **asin** | **string** | ASIN | 
 **companyId** | **string** | Company id | 
 **listingPostType** | **string** | Listing post type | 
 **locale** | **string** | Locale (default: logged in user&#39;s locale) | 
 **messageBoardId** | **string** | Message board id | 

### Return type

[**PostView**](PostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/html, application/json, text/plain
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePostUsingDELETE

> MessagePrototype DeletePostUsingDELETE(ctx, postId).Execute()

Removes post



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
	resp, r, err := apiClient.PostControllerAPI.DeletePostUsingDELETE(context.Background(), postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostControllerAPI.DeletePostUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePostUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `PostControllerAPI.DeletePostUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** | postId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostUsingDELETERequest struct via the builder pattern


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


## EditPostUsingPUT

> PostView EditPostUsingPUT(ctx, postId).Title(title).Content(content).AllianceId(allianceId).Asin(asin).CompanyId(companyId).ListingPostType(listingPostType).Locale(locale).MessageBoardId(messageBoardId).Execute()

Edits post

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
	title := "title_example" // string | Title
	content := "content_example" // string | content
	allianceId := "allianceId_example" // string | Alliance id (optional)
	asin := "asin_example" // string | ASIN (optional)
	companyId := "companyId_example" // string | Company id (optional)
	listingPostType := "listingPostType_example" // string | Listing post type (optional)
	locale := "locale_example" // string | Locale (default: logged in user's locale) (optional)
	messageBoardId := "messageBoardId_example" // string | Message board id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostControllerAPI.EditPostUsingPUT(context.Background(), postId).Title(title).Content(content).AllianceId(allianceId).Asin(asin).CompanyId(companyId).ListingPostType(listingPostType).Locale(locale).MessageBoardId(messageBoardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostControllerAPI.EditPostUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditPostUsingPUT`: PostView
	fmt.Fprintf(os.Stdout, "Response from `PostControllerAPI.EditPostUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** | postId | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPostUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **title** | **string** | Title | 
 **content** | **string** | content | 
 **allianceId** | **string** | Alliance id | 
 **asin** | **string** | ASIN | 
 **companyId** | **string** | Company id | 
 **listingPostType** | **string** | Listing post type | 
 **locale** | **string** | Locale (default: logged in user&#39;s locale) | 
 **messageBoardId** | **string** | Message board id | 

### Return type

[**PostView**](PostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/html, application/json, text/plain
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListingDescriptionUsingGET

> PostView GetListingDescriptionUsingGET(ctx).Asin(asin).ListingPostType(listingPostType).Execute()

Returns listing description

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
	asin := "asin_example" // string | ASIN (optional)
	listingPostType := "listingPostType_example" // string | Listing post type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostControllerAPI.GetListingDescriptionUsingGET(context.Background()).Asin(asin).ListingPostType(listingPostType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostControllerAPI.GetListingDescriptionUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListingDescriptionUsingGET`: PostView
	fmt.Fprintf(os.Stdout, "Response from `PostControllerAPI.GetListingDescriptionUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetListingDescriptionUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asin** | **string** | ASIN | 
 **listingPostType** | **string** | Listing post type | 

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


## GetPersonalInterestInPostUsingGET

> InterestSummary GetPersonalInterestInPostUsingGET(ctx, postId).Execute()

Returns personal interest in post

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
	resp, r, err := apiClient.PostControllerAPI.GetPersonalInterestInPostUsingGET(context.Background(), postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostControllerAPI.GetPersonalInterestInPostUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPersonalInterestInPostUsingGET`: InterestSummary
	fmt.Fprintf(os.Stdout, "Response from `PostControllerAPI.GetPersonalInterestInPostUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** | postId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonalInterestInPostUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InterestSummary**](InterestSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostUsingGET

> PostView GetPostUsingGET(ctx, postId).Execute()

Returns post

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
	resp, r, err := apiClient.PostControllerAPI.GetPostUsingGET(context.Background(), postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostControllerAPI.GetPostUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPostUsingGET`: PostView
	fmt.Fprintf(os.Stdout, "Response from `PostControllerAPI.GetPostUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** | postId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostUsingGETRequest struct via the builder pattern


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


## ListCommentsUsingGET

> []PostView ListCommentsUsingGET(ctx, postId).Execute()

Lists post's comments. Shows all comments for root posts and the first level comments for comment

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
	resp, r, err := apiClient.PostControllerAPI.ListCommentsUsingGET(context.Background(), postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostControllerAPI.ListCommentsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCommentsUsingGET`: []PostView
	fmt.Fprintf(os.Stdout, "Response from `PostControllerAPI.ListCommentsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** | postId | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCommentsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PostView**](PostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListListingPostsUsingGET

> PagePostView ListListingPostsUsingGET(ctx).Asin(asin).Page(page).Size(size).Sort(sort).Execute()

List all postings regarding listing

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
	asin := "asin_example" // string | ASIN (optional)
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostControllerAPI.ListListingPostsUsingGET(context.Background()).Asin(asin).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostControllerAPI.ListListingPostsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListListingPostsUsingGET`: PagePostView
	fmt.Fprintf(os.Stdout, "Response from `PostControllerAPI.ListListingPostsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListListingPostsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asin** | **string** | ASIN | 
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

