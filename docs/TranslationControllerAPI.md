# \TranslationControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TranslateUsingGET**](TranslationControllerAPI.md#TranslateUsingGET) | **Get** /translations/v2/{param}/{locale}.json | Returns translation strings



## TranslateUsingGET

> map[string]string TranslateUsingGET(ctx, locale, param).Execute()

Returns translation strings

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
	locale := "locale_example" // string | Locale (e.g. de_DE, de-DE, de, en_US)
	param := "param_example" // string | Parameter for caching purposes, will be ignored

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationControllerAPI.TranslateUsingGET(context.Background(), locale, param).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationControllerAPI.TranslateUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TranslateUsingGET`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `TranslationControllerAPI.TranslateUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locale** | **string** | Locale (e.g. de_DE, de-DE, de, en_US) | 
**param** | **string** | Parameter for caching purposes, will be ignored | 

### Other Parameters

Other parameters are passed through a pointer to a apiTranslateUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

