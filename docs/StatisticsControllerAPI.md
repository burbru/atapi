# \StatisticsControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrentMinimalStatsPublicUsingGET**](StatisticsControllerAPI.md#GetCurrentMinimalStatsPublicUsingGET) | **Get** /minimalstats | Current minimal market statistics
[**GetCurrentMinimalStatsPublicV2UsingGET**](StatisticsControllerAPI.md#GetCurrentMinimalStatsPublicV2UsingGET) | **Get** /api/v2/minimalstats | Current minimal market statistics
[**ListHourlyStatisticsUsingGET**](StatisticsControllerAPI.md#ListHourlyStatisticsUsingGET) | **Get** /api/marketstatistics/ | Lists latest hourly market statistics (2400 max)



## GetCurrentMinimalStatsPublicUsingGET

> MinimalMarketStatsView GetCurrentMinimalStatsPublicUsingGET(ctx).Execute()

Current minimal market statistics

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
	resp, r, err := apiClient.StatisticsControllerAPI.GetCurrentMinimalStatsPublicUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticsControllerAPI.GetCurrentMinimalStatsPublicUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentMinimalStatsPublicUsingGET`: MinimalMarketStatsView
	fmt.Fprintf(os.Stdout, "Response from `StatisticsControllerAPI.GetCurrentMinimalStatsPublicUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentMinimalStatsPublicUsingGETRequest struct via the builder pattern


### Return type

[**MinimalMarketStatsView**](MinimalMarketStatsView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentMinimalStatsPublicV2UsingGET

> MinimalMarketStatsView GetCurrentMinimalStatsPublicV2UsingGET(ctx).Execute()

Current minimal market statistics

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
	resp, r, err := apiClient.StatisticsControllerAPI.GetCurrentMinimalStatsPublicV2UsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticsControllerAPI.GetCurrentMinimalStatsPublicV2UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentMinimalStatsPublicV2UsingGET`: MinimalMarketStatsView
	fmt.Fprintf(os.Stdout, "Response from `StatisticsControllerAPI.GetCurrentMinimalStatsPublicV2UsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentMinimalStatsPublicV2UsingGETRequest struct via the builder pattern


### Return type

[**MinimalMarketStatsView**](MinimalMarketStatsView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHourlyStatisticsUsingGET

> MarketStatsSnapshot ListHourlyStatisticsUsingGET(ctx).Execute()

Lists latest hourly market statistics (2400 max)

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
	resp, r, err := apiClient.StatisticsControllerAPI.ListHourlyStatisticsUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticsControllerAPI.ListHourlyStatisticsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHourlyStatisticsUsingGET`: MarketStatsSnapshot
	fmt.Fprintf(os.Stdout, "Response from `StatisticsControllerAPI.ListHourlyStatisticsUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListHourlyStatisticsUsingGETRequest struct via the builder pattern


### Return type

[**MarketStatsSnapshot**](MarketStatsSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

