# \LocaleControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLocalesUsingGET**](LocaleControllerAPI.md#GetLocalesUsingGET) | **Get** /api/locales | Lists available locales
[**GetMyLocaleUsingGET**](LocaleControllerAPI.md#GetMyLocaleUsingGET) | **Get** /api/locale | Returns logged in user&#39;s locale
[**SetMyLocaleUsingPUT**](LocaleControllerAPI.md#SetMyLocaleUsingPUT) | **Put** /api/locale | Sets logged in user&#39;s locale



## GetLocalesUsingGET

> []string GetLocalesUsingGET(ctx).Execute()

Lists available locales

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
	resp, r, err := apiClient.LocaleControllerAPI.GetLocalesUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocaleControllerAPI.GetLocalesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalesUsingGET`: []string
	fmt.Fprintf(os.Stdout, "Response from `LocaleControllerAPI.GetLocalesUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalesUsingGETRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyLocaleUsingGET

> MessagePrototype GetMyLocaleUsingGET(ctx).Execute()

Returns logged in user's locale

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
	resp, r, err := apiClient.LocaleControllerAPI.GetMyLocaleUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocaleControllerAPI.GetMyLocaleUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyLocaleUsingGET`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `LocaleControllerAPI.GetMyLocaleUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyLocaleUsingGETRequest struct via the builder pattern


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


## SetMyLocaleUsingPUT

> MessagePrototype SetMyLocaleUsingPUT(ctx).Locale(locale).Execute()

Sets logged in user's locale

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
	locale := "locale_example" // string | Locale

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocaleControllerAPI.SetMyLocaleUsingPUT(context.Background()).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocaleControllerAPI.SetMyLocaleUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetMyLocaleUsingPUT`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `LocaleControllerAPI.SetMyLocaleUsingPUT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetMyLocaleUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locale** | **string** | Locale | 

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

