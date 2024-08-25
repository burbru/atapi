# \SearchControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultiSearchUsingGET**](SearchControllerAPI.md#MultiSearchUsingGET) | **Get** /api/multisearch/{searchString} | Search for multiple entities



## MultiSearchUsingGET

> SearchResult MultiSearchUsingGET(ctx, searchString).Execute()

Search for multiple entities

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
	searchString := "searchString_example" // string | Search string

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchControllerAPI.MultiSearchUsingGET(context.Background(), searchString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchControllerAPI.MultiSearchUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultiSearchUsingGET`: SearchResult
	fmt.Fprintf(os.Stdout, "Response from `SearchControllerAPI.MultiSearchUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchString** | **string** | Search string | 

### Other Parameters

Other parameters are passed through a pointer to a apiMultiSearchUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SearchResult**](SearchResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

