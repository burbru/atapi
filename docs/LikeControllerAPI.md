# \LikeControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLikeUsingDELETE**](LikeControllerAPI.md#DeleteLikeUsingDELETE) | **Delete** /api/v2/my/likes/{postId} | Removes Like/dislike of currently logged in user
[**LikePostUsingPUT**](LikeControllerAPI.md#LikePostUsingPUT) | **Put** /api/v2/my/likes/{postId} | Likes/dislikes post by currently logged in user
[**ListLikesUsingGET**](LikeControllerAPI.md#ListLikesUsingGET) | **Get** /api/v2/likes/{postId} | Lists likes/dislikes of post



## DeleteLikeUsingDELETE

> MessagePrototype DeleteLikeUsingDELETE(ctx, postId).Execute()

Removes Like/dislike of currently logged in user

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
	resp, r, err := apiClient.LikeControllerAPI.DeleteLikeUsingDELETE(context.Background(), postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LikeControllerAPI.DeleteLikeUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLikeUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `LikeControllerAPI.DeleteLikeUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** | postId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLikeUsingDELETERequest struct via the builder pattern


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


## LikePostUsingPUT

> PostLikeView LikePostUsingPUT(ctx, postId).Type_(type_).Execute()

Likes/dislikes post by currently logged in user

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
	type_ := "type__example" // string | Type of like

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LikeControllerAPI.LikePostUsingPUT(context.Background(), postId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LikeControllerAPI.LikePostUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LikePostUsingPUT`: PostLikeView
	fmt.Fprintf(os.Stdout, "Response from `LikeControllerAPI.LikePostUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** | postId | 

### Other Parameters

Other parameters are passed through a pointer to a apiLikePostUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Type of like | 

### Return type

[**PostLikeView**](PostLikeView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLikesUsingGET

> PostLikeView ListLikesUsingGET(ctx, postId).Execute()

Lists likes/dislikes of post

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
	resp, r, err := apiClient.LikeControllerAPI.ListLikesUsingGET(context.Background(), postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LikeControllerAPI.ListLikesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLikesUsingGET`: PostLikeView
	fmt.Fprintf(os.Stdout, "Response from `LikeControllerAPI.ListLikesUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** | postId | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLikesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostLikeView**](PostLikeView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

