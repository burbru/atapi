# \InterestTenderControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrentInterestTenderUsingGET**](InterestTenderControllerAPI.md#GetCurrentInterestTenderUsingGET) | **Get** /api/v2/interesttenders | Returns current interest tender



## GetCurrentInterestTenderUsingGET

> InterestTenderView GetCurrentInterestTenderUsingGET(ctx).Execute()

Returns current interest tender

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
	resp, r, err := apiClient.InterestTenderControllerAPI.GetCurrentInterestTenderUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterestTenderControllerAPI.GetCurrentInterestTenderUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentInterestTenderUsingGET`: InterestTenderView
	fmt.Fprintf(os.Stdout, "Response from `InterestTenderControllerAPI.GetCurrentInterestTenderUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentInterestTenderUsingGETRequest struct via the builder pattern


### Return type

[**InterestTenderView**](InterestTenderView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

