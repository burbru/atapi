# \DerivativeControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWarrantUsingPOST**](DerivativeControllerAPI.md#CreateWarrantUsingPOST) | **Post** /api/v2/warrants | Creates warrant
[**GetWarrantUsingGET**](DerivativeControllerAPI.md#GetWarrantUsingGET) | **Get** /api/v2/warrants/{warrantId} | Returns warrant
[**ListWarrantsUsingGET**](DerivativeControllerAPI.md#ListWarrantsUsingGET) | **Get** /api/v2/warrants | Lists all warrants for underlying



## CreateWarrantUsingPOST

> AllianceView CreateWarrantUsingPOST(ctx).CashDeposit(cashDeposit).CompanyId(companyId).Type_(type_).UnderlyingAsin(underlyingAsin).Execute()

Creates warrant

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
	cashDeposit := float32(8.14) // float32 | Cash Deposit
	companyId := "companyId_example" // string | Company ID
	type_ := "type__example" // string | Warrant Type
	underlyingAsin := "underlyingAsin_example" // string | Underlying ASIN

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DerivativeControllerAPI.CreateWarrantUsingPOST(context.Background()).CashDeposit(cashDeposit).CompanyId(companyId).Type_(type_).UnderlyingAsin(underlyingAsin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DerivativeControllerAPI.CreateWarrantUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWarrantUsingPOST`: AllianceView
	fmt.Fprintf(os.Stdout, "Response from `DerivativeControllerAPI.CreateWarrantUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWarrantUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cashDeposit** | **float32** | Cash Deposit | 
 **companyId** | **string** | Company ID | 
 **type_** | **string** | Warrant Type | 
 **underlyingAsin** | **string** | Underlying ASIN | 

### Return type

[**AllianceView**](AllianceView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWarrantUsingGET

> WarrantView GetWarrantUsingGET(ctx, warrantId).Execute()

Returns warrant

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
	warrantId := "warrantId_example" // string | Warrant Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DerivativeControllerAPI.GetWarrantUsingGET(context.Background(), warrantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DerivativeControllerAPI.GetWarrantUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWarrantUsingGET`: WarrantView
	fmt.Fprintf(os.Stdout, "Response from `DerivativeControllerAPI.GetWarrantUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warrantId** | **string** | Warrant Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWarrantUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WarrantView**](WarrantView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWarrantsUsingGET

> PageWarrantView ListWarrantsUsingGET(ctx).UnderlyingAsin(underlyingAsin).Page(page).Size(size).Sort(sort).Execute()

Lists all warrants for underlying



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
	underlyingAsin := "underlyingAsin_example" // string | Underlying ASIN
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DerivativeControllerAPI.ListWarrantsUsingGET(context.Background()).UnderlyingAsin(underlyingAsin).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DerivativeControllerAPI.ListWarrantsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWarrantsUsingGET`: PageWarrantView
	fmt.Fprintf(os.Stdout, "Response from `DerivativeControllerAPI.ListWarrantsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWarrantsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **underlyingAsin** | **string** | Underlying ASIN | 
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageWarrantView**](PageWarrantView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

