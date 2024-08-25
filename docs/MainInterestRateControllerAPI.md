# \MainInterestRateControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLatestRateUsingGET**](MainInterestRateControllerAPI.md#GetLatestRateUsingGET) | **Get** /api/maininterestrate/latest/ | Returns latest main interest rate in percent
[**ListAverageRatesUsingGET**](MainInterestRateControllerAPI.md#ListAverageRatesUsingGET) | **Get** /api/v2/averagebondinterestrate | Lists average daily bond interest rates in percent
[**ListRatesUsingGET**](MainInterestRateControllerAPI.md#ListRatesUsingGET) | **Get** /api/maininterestrate/ | Lists main interest rates  in percent
[**ListRatesV2UsingGET**](MainInterestRateControllerAPI.md#ListRatesV2UsingGET) | **Get** /api/v2/maininterestrates | Lists main interest rates  in percent



## GetLatestRateUsingGET

> MainInterestRateView GetLatestRateUsingGET(ctx).Execute()

Returns latest main interest rate in percent

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
	resp, r, err := apiClient.MainInterestRateControllerAPI.GetLatestRateUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MainInterestRateControllerAPI.GetLatestRateUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestRateUsingGET`: MainInterestRateView
	fmt.Fprintf(os.Stdout, "Response from `MainInterestRateControllerAPI.GetLatestRateUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestRateUsingGETRequest struct via the builder pattern


### Return type

[**MainInterestRateView**](MainInterestRateView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAverageRatesUsingGET

> AverageDailyBondInterestRateView ListAverageRatesUsingGET(ctx).Execute()

Lists average daily bond interest rates in percent

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
	resp, r, err := apiClient.MainInterestRateControllerAPI.ListAverageRatesUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MainInterestRateControllerAPI.ListAverageRatesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAverageRatesUsingGET`: AverageDailyBondInterestRateView
	fmt.Fprintf(os.Stdout, "Response from `MainInterestRateControllerAPI.ListAverageRatesUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAverageRatesUsingGETRequest struct via the builder pattern


### Return type

[**AverageDailyBondInterestRateView**](AverageDailyBondInterestRateView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRatesUsingGET

> MainInterestRateView ListRatesUsingGET(ctx).Execute()

Lists main interest rates  in percent

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
	resp, r, err := apiClient.MainInterestRateControllerAPI.ListRatesUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MainInterestRateControllerAPI.ListRatesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRatesUsingGET`: MainInterestRateView
	fmt.Fprintf(os.Stdout, "Response from `MainInterestRateControllerAPI.ListRatesUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRatesUsingGETRequest struct via the builder pattern


### Return type

[**MainInterestRateView**](MainInterestRateView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRatesV2UsingGET

> MainInterestRateView ListRatesV2UsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Lists main interest rates  in percent

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
	resp, r, err := apiClient.MainInterestRateControllerAPI.ListRatesV2UsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MainInterestRateControllerAPI.ListRatesV2UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRatesV2UsingGET`: MainInterestRateView
	fmt.Fprintf(os.Stdout, "Response from `MainInterestRateControllerAPI.ListRatesV2UsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRatesV2UsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**MainInterestRateView**](MainInterestRateView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

