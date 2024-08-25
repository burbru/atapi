# \SecurityPriceControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPricesUsingGET**](SecurityPriceControllerAPI.md#ListPricesUsingGET) | **Get** /api/securityPrices | Lists all prices for security identified by security identifier



## ListPricesUsingGET

> SecurityPriceView ListPricesUsingGET(ctx).SecurityIdentifier(securityIdentifier).EndDate(endDate).StartDate(startDate).Execute()

Lists all prices for security identified by security identifier

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
	securityIdentifier := "securityIdentifier_example" // string | Security identifier
	endDate := "endDate_example" // string | End date (optional)
	startDate := "startDate_example" // string | Start date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityPriceControllerAPI.ListPricesUsingGET(context.Background()).SecurityIdentifier(securityIdentifier).EndDate(endDate).StartDate(startDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityPriceControllerAPI.ListPricesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPricesUsingGET`: SecurityPriceView
	fmt.Fprintf(os.Stdout, "Response from `SecurityPriceControllerAPI.ListPricesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPricesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityIdentifier** | **string** | Security identifier | 
 **endDate** | **string** | End date | 
 **startDate** | **string** | Start date | 

### Return type

[**SecurityPriceView**](SecurityPriceView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

