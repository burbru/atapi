# \PortfolioControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMyPortfolioUsingGET**](PortfolioControllerAPI.md#GetMyPortfolioUsingGET) | **Get** /api/v2/my/portfolio | Returns portfolio view of private securities account
[**GetPortfolioFixedIncomeSecuritiesUsingGET**](PortfolioControllerAPI.md#GetPortfolioFixedIncomeSecuritiesUsingGET) | **Get** /api/portfolios/fixedincome/{securitiesAccountId} | Returns fixed income securities in portfolio
[**GetPortfolioUsingGET**](PortfolioControllerAPI.md#GetPortfolioUsingGET) | **Get** /api/portfolios/{securitiesAccountId} | Returns portfolio view of securities account



## GetMyPortfolioUsingGET

> PortfolioView GetMyPortfolioUsingGET(ctx).Execute()

Returns portfolio view of private securities account

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
	resp, r, err := apiClient.PortfolioControllerAPI.GetMyPortfolioUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioControllerAPI.GetMyPortfolioUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyPortfolioUsingGET`: PortfolioView
	fmt.Fprintf(os.Stdout, "Response from `PortfolioControllerAPI.GetMyPortfolioUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyPortfolioUsingGETRequest struct via the builder pattern


### Return type

[**PortfolioView**](PortfolioView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortfolioFixedIncomeSecuritiesUsingGET

> FixedIncomeSecurityView GetPortfolioFixedIncomeSecuritiesUsingGET(ctx, securitiesAccountId).Execute()

Returns fixed income securities in portfolio

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
	securitiesAccountId := "securitiesAccountId_example" // string | Securities account id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioControllerAPI.GetPortfolioFixedIncomeSecuritiesUsingGET(context.Background(), securitiesAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioControllerAPI.GetPortfolioFixedIncomeSecuritiesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortfolioFixedIncomeSecuritiesUsingGET`: FixedIncomeSecurityView
	fmt.Fprintf(os.Stdout, "Response from `PortfolioControllerAPI.GetPortfolioFixedIncomeSecuritiesUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securitiesAccountId** | **string** | Securities account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortfolioFixedIncomeSecuritiesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FixedIncomeSecurityView**](FixedIncomeSecurityView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortfolioUsingGET

> PortfolioView GetPortfolioUsingGET(ctx, securitiesAccountId).Execute()

Returns portfolio view of securities account

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
	securitiesAccountId := "securitiesAccountId_example" // string | Securities account id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioControllerAPI.GetPortfolioUsingGET(context.Background(), securitiesAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioControllerAPI.GetPortfolioUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortfolioUsingGET`: PortfolioView
	fmt.Fprintf(os.Stdout, "Response from `PortfolioControllerAPI.GetPortfolioUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securitiesAccountId** | **string** | Securities account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortfolioUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PortfolioView**](PortfolioView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

