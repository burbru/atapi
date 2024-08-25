# \BankingLicenseControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBankingLicenseUsingPOST**](BankingLicenseControllerAPI.md#CreateBankingLicenseUsingPOST) | **Post** /api/bankinglicense | Create banking license
[**GetBankingLicenseUsingGET**](BankingLicenseControllerAPI.md#GetBankingLicenseUsingGET) | **Get** /api/bankinglicense/{bankingLicenseId} | Returns banking license
[**GetCompanyBankingLicenseUsingGET**](BankingLicenseControllerAPI.md#GetCompanyBankingLicenseUsingGET) | **Get** /api/bankinglicense | Returns company&#39;s banking license



## CreateBankingLicenseUsingPOST

> BankingLicenseView CreateBankingLicenseUsingPOST(ctx).CompanyId(companyId).Execute()

Create banking license

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
	companyId := "companyId_example" // string | Company id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingLicenseControllerAPI.CreateBankingLicenseUsingPOST(context.Background()).CompanyId(companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingLicenseControllerAPI.CreateBankingLicenseUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBankingLicenseUsingPOST`: BankingLicenseView
	fmt.Fprintf(os.Stdout, "Response from `BankingLicenseControllerAPI.CreateBankingLicenseUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBankingLicenseUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company id | 

### Return type

[**BankingLicenseView**](BankingLicenseView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankingLicenseUsingGET

> BankingLicenseView GetBankingLicenseUsingGET(ctx, bankingLicenseId).Execute()

Returns banking license

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
	bankingLicenseId := "bankingLicenseId_example" // string | Banking License id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingLicenseControllerAPI.GetBankingLicenseUsingGET(context.Background(), bankingLicenseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingLicenseControllerAPI.GetBankingLicenseUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankingLicenseUsingGET`: BankingLicenseView
	fmt.Fprintf(os.Stdout, "Response from `BankingLicenseControllerAPI.GetBankingLicenseUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankingLicenseId** | **string** | Banking License id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankingLicenseUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BankingLicenseView**](BankingLicenseView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyBankingLicenseUsingGET

> BankingLicenseView GetCompanyBankingLicenseUsingGET(ctx).CompanyId(companyId).Execute()

Returns company's banking license

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
	companyId := "companyId_example" // string | Company id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingLicenseControllerAPI.GetCompanyBankingLicenseUsingGET(context.Background()).CompanyId(companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingLicenseControllerAPI.GetCompanyBankingLicenseUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyBankingLicenseUsingGET`: BankingLicenseView
	fmt.Fprintf(os.Stdout, "Response from `BankingLicenseControllerAPI.GetCompanyBankingLicenseUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyBankingLicenseUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company id | 

### Return type

[**BankingLicenseView**](BankingLicenseView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

