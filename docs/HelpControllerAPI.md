# \HelpControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHelpDefinitionUsingGET**](HelpControllerAPI.md#GetHelpDefinitionUsingGET) | **Get** /api/v2/helpcomments | Get help comment for identifier and locale



## GetHelpDefinitionUsingGET

> HelpComment GetHelpDefinitionUsingGET(ctx).Identifier(identifier).Locale(locale).Execute()

Get help comment for identifier and locale

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
	identifier := "identifier_example" // string | Help comment identifier
	locale := "locale_example" // string | Locale

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelpControllerAPI.GetHelpDefinitionUsingGET(context.Background()).Identifier(identifier).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelpControllerAPI.GetHelpDefinitionUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHelpDefinitionUsingGET`: HelpComment
	fmt.Fprintf(os.Stdout, "Response from `HelpControllerAPI.GetHelpDefinitionUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHelpDefinitionUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifier** | **string** | Help comment identifier | 
 **locale** | **string** | Locale | 

### Return type

[**HelpComment**](HelpComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

