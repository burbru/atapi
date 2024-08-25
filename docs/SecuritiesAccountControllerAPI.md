# \SecuritiesAccountControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindSecuritiesAccountDetailsUsingGET**](SecuritiesAccountControllerAPI.md#FindSecuritiesAccountDetailsUsingGET) | **Get** /api/v2/securitiesaccountdetails | Finds securities account details for OTC orders
[**GetSecuritiesAccountDetailsUsingGET**](SecuritiesAccountControllerAPI.md#GetSecuritiesAccountDetailsUsingGET) | **Get** /api/v2/securitiesaccountdetails/{securitiesAccountId} | Returns securities account details for OTC orders
[**GetSecuritiesAccountUsingGET**](SecuritiesAccountControllerAPI.md#GetSecuritiesAccountUsingGET) | **Get** /api/v2/my/securitiesaccount | Returns private securities account



## FindSecuritiesAccountDetailsUsingGET

> SecuritiesAccountDetailsView FindSecuritiesAccountDetailsUsingGET(ctx).Search(search).Execute()

Finds securities account details for OTC orders

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
	search := "search_example" // string | Fulltext search

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecuritiesAccountControllerAPI.FindSecuritiesAccountDetailsUsingGET(context.Background()).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecuritiesAccountControllerAPI.FindSecuritiesAccountDetailsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindSecuritiesAccountDetailsUsingGET`: SecuritiesAccountDetailsView
	fmt.Fprintf(os.Stdout, "Response from `SecuritiesAccountControllerAPI.FindSecuritiesAccountDetailsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindSecuritiesAccountDetailsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Fulltext search | 

### Return type

[**SecuritiesAccountDetailsView**](SecuritiesAccountDetailsView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecuritiesAccountDetailsUsingGET

> SecuritiesAccountDetailsView GetSecuritiesAccountDetailsUsingGET(ctx, securitiesAccountId).Execute()

Returns securities account details for OTC orders

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
	securitiesAccountId := "securitiesAccountId_example" // string | ID of securities account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecuritiesAccountControllerAPI.GetSecuritiesAccountDetailsUsingGET(context.Background(), securitiesAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecuritiesAccountControllerAPI.GetSecuritiesAccountDetailsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecuritiesAccountDetailsUsingGET`: SecuritiesAccountDetailsView
	fmt.Fprintf(os.Stdout, "Response from `SecuritiesAccountControllerAPI.GetSecuritiesAccountDetailsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securitiesAccountId** | **string** | ID of securities account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecuritiesAccountDetailsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecuritiesAccountDetailsView**](SecuritiesAccountDetailsView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecuritiesAccountUsingGET

> SecuritiesAccountView GetSecuritiesAccountUsingGET(ctx).Execute()

Returns private securities account

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
	resp, r, err := apiClient.SecuritiesAccountControllerAPI.GetSecuritiesAccountUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecuritiesAccountControllerAPI.GetSecuritiesAccountUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecuritiesAccountUsingGET`: SecuritiesAccountView
	fmt.Fprintf(os.Stdout, "Response from `SecuritiesAccountControllerAPI.GetSecuritiesAccountUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecuritiesAccountUsingGETRequest struct via the builder pattern


### Return type

[**SecuritiesAccountView**](SecuritiesAccountView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

