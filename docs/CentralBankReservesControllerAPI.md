# \CentralBankReservesControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompanyReservesUsingGET**](CentralBankReservesControllerAPI.md#GetCompanyReservesUsingGET) | **Get** /api/centralbankreserves | Returns company&#39;s reserves
[**GetLastReservesPaymentUsingGET**](CentralBankReservesControllerAPI.md#GetLastReservesPaymentUsingGET) | **Get** /api/v2/lastcentralbankreservespayment | Returns latest central bank reserves payment
[**GetReservesUsingGET**](CentralBankReservesControllerAPI.md#GetReservesUsingGET) | **Get** /api/centralbankreserves/{reservesId} | Returns reserves
[**IncreaseInterestRateBoostUsingPUT**](CentralBankReservesControllerAPI.md#IncreaseInterestRateBoostUsingPUT) | **Put** /api/v2/centralbankreserves/{reservesId} | Increase central bank reserves interest rate boost by 0.01%
[**IncreaseReservesUsingPUT**](CentralBankReservesControllerAPI.md#IncreaseReservesUsingPUT) | **Put** /api/centralbankreserves | Increase central bank reserves



## GetCompanyReservesUsingGET

> CentralbankReservesView GetCompanyReservesUsingGET(ctx).CompanyId(companyId).Execute()

Returns company's reserves

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
	resp, r, err := apiClient.CentralBankReservesControllerAPI.GetCompanyReservesUsingGET(context.Background()).CompanyId(companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CentralBankReservesControllerAPI.GetCompanyReservesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyReservesUsingGET`: CentralbankReservesView
	fmt.Fprintf(os.Stdout, "Response from `CentralBankReservesControllerAPI.GetCompanyReservesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyReservesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company id | 

### Return type

[**CentralbankReservesView**](CentralbankReservesView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastReservesPaymentUsingGET

> CentralBankReservesInterestPaymentView GetLastReservesPaymentUsingGET(ctx).Execute()

Returns latest central bank reserves payment

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
	resp, r, err := apiClient.CentralBankReservesControllerAPI.GetLastReservesPaymentUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CentralBankReservesControllerAPI.GetLastReservesPaymentUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLastReservesPaymentUsingGET`: CentralBankReservesInterestPaymentView
	fmt.Fprintf(os.Stdout, "Response from `CentralBankReservesControllerAPI.GetLastReservesPaymentUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastReservesPaymentUsingGETRequest struct via the builder pattern


### Return type

[**CentralBankReservesInterestPaymentView**](CentralBankReservesInterestPaymentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReservesUsingGET

> CentralbankReservesView GetReservesUsingGET(ctx, reservesId).Execute()

Returns reserves

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
	reservesId := "reservesId_example" // string | Reserves id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CentralBankReservesControllerAPI.GetReservesUsingGET(context.Background(), reservesId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CentralBankReservesControllerAPI.GetReservesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReservesUsingGET`: CentralbankReservesView
	fmt.Fprintf(os.Stdout, "Response from `CentralBankReservesControllerAPI.GetReservesUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservesId** | **string** | Reserves id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReservesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CentralbankReservesView**](CentralbankReservesView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncreaseInterestRateBoostUsingPUT

> CentralbankReservesView IncreaseInterestRateBoostUsingPUT(ctx, reservesId).IncreaseInterestRateBoost(increaseInterestRateBoost).Multiplier(multiplier).Execute()

Increase central bank reserves interest rate boost by 0.01%

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
	increaseInterestRateBoost := true // bool | Increase interest boost rate
	reservesId := "reservesId_example" // string | Reserves id
	multiplier := int64(789) // int64 | Number of boosts (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CentralBankReservesControllerAPI.IncreaseInterestRateBoostUsingPUT(context.Background(), reservesId).IncreaseInterestRateBoost(increaseInterestRateBoost).Multiplier(multiplier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CentralBankReservesControllerAPI.IncreaseInterestRateBoostUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IncreaseInterestRateBoostUsingPUT`: CentralbankReservesView
	fmt.Fprintf(os.Stdout, "Response from `CentralBankReservesControllerAPI.IncreaseInterestRateBoostUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservesId** | **string** | Reserves id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIncreaseInterestRateBoostUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **increaseInterestRateBoost** | **bool** | Increase interest boost rate | 

 **multiplier** | **int64** | Number of boosts | 

### Return type

[**CentralbankReservesView**](CentralbankReservesView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncreaseReservesUsingPUT

> CentralbankReservesView IncreaseReservesUsingPUT(ctx).CashAmount(cashAmount).CompanyId(companyId).Execute()

Increase central bank reserves

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
	cashAmount := "cashAmount_example" // string | Cash amount to increase cash holding
	companyId := "companyId_example" // string | Company id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CentralBankReservesControllerAPI.IncreaseReservesUsingPUT(context.Background()).CashAmount(cashAmount).CompanyId(companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CentralBankReservesControllerAPI.IncreaseReservesUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IncreaseReservesUsingPUT`: CentralbankReservesView
	fmt.Fprintf(os.Stdout, "Response from `CentralBankReservesControllerAPI.IncreaseReservesUsingPUT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIncreaseReservesUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cashAmount** | **string** | Cash amount to increase cash holding | 
 **companyId** | **string** | Company id | 

### Return type

[**CentralbankReservesView**](CentralbankReservesView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

