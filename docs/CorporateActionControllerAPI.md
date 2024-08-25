# \CorporateActionControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GeMergersUsingGET**](CorporateActionControllerAPI.md#GeMergersUsingGET) | **Get** /api/v2/mergers | Returns mergers
[**GetCapitalIncreasesUsingGET**](CorporateActionControllerAPI.md#GetCapitalIncreasesUsingGET) | **Get** /api/v2/capitalincreases | Returns capital increases
[**GetCapitalReductionsUsingGET**](CorporateActionControllerAPI.md#GetCapitalReductionsUsingGET) | **Get** /api/v2/capitalreductions | Returns capital reductions
[**GetCompanyCapitalIncreaseUsingGET**](CorporateActionControllerAPI.md#GetCompanyCapitalIncreaseUsingGET) | **Get** /api/v2/companies/{companyId}/capitalincrease | Returns company&#39;s current capital increase
[**GetCompanyCapitalReductionUsingGET**](CorporateActionControllerAPI.md#GetCompanyCapitalReductionUsingGET) | **Get** /api/v2/companies/{companyId}/capitalreduction | Returns company&#39;s current capital reduction
[**GetCompanyDividendPaymentUsingGET**](CorporateActionControllerAPI.md#GetCompanyDividendPaymentUsingGET) | **Get** /api/v2/companies/{companyId}/dividendpayment | Returns company&#39;s current dividend payment
[**GetDividendPaymentsUsingGET**](CorporateActionControllerAPI.md#GetDividendPaymentsUsingGET) | **Get** /api/v2/dividendpayments | Returns dividend payments
[**GetMergerUsingGET**](CorporateActionControllerAPI.md#GetMergerUsingGET) | **Get** /api/v2/companies/{companyId}/merger | Returns company&#39;s current merger



## GeMergersUsingGET

> MergerView GeMergersUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Returns mergers

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
	resp, r, err := apiClient.CorporateActionControllerAPI.GeMergersUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CorporateActionControllerAPI.GeMergersUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GeMergersUsingGET`: MergerView
	fmt.Fprintf(os.Stdout, "Response from `CorporateActionControllerAPI.GeMergersUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeMergersUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**MergerView**](MergerView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCapitalIncreasesUsingGET

> CapitalIncreaseView GetCapitalIncreasesUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Returns capital increases

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
	resp, r, err := apiClient.CorporateActionControllerAPI.GetCapitalIncreasesUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CorporateActionControllerAPI.GetCapitalIncreasesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCapitalIncreasesUsingGET`: CapitalIncreaseView
	fmt.Fprintf(os.Stdout, "Response from `CorporateActionControllerAPI.GetCapitalIncreasesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCapitalIncreasesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**CapitalIncreaseView**](CapitalIncreaseView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCapitalReductionsUsingGET

> CapitalReductionView GetCapitalReductionsUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Returns capital reductions

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
	resp, r, err := apiClient.CorporateActionControllerAPI.GetCapitalReductionsUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CorporateActionControllerAPI.GetCapitalReductionsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCapitalReductionsUsingGET`: CapitalReductionView
	fmt.Fprintf(os.Stdout, "Response from `CorporateActionControllerAPI.GetCapitalReductionsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCapitalReductionsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**CapitalReductionView**](CapitalReductionView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyCapitalIncreaseUsingGET

> CapitalIncreaseView GetCompanyCapitalIncreaseUsingGET(ctx, companyId).Execute()

Returns company's current capital increase

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
	resp, r, err := apiClient.CorporateActionControllerAPI.GetCompanyCapitalIncreaseUsingGET(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CorporateActionControllerAPI.GetCompanyCapitalIncreaseUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCapitalIncreaseUsingGET`: CapitalIncreaseView
	fmt.Fprintf(os.Stdout, "Response from `CorporateActionControllerAPI.GetCompanyCapitalIncreaseUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Company id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCapitalIncreaseUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CapitalIncreaseView**](CapitalIncreaseView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyCapitalReductionUsingGET

> CapitalReductionView GetCompanyCapitalReductionUsingGET(ctx, companyId).Execute()

Returns company's current capital reduction

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
	resp, r, err := apiClient.CorporateActionControllerAPI.GetCompanyCapitalReductionUsingGET(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CorporateActionControllerAPI.GetCompanyCapitalReductionUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCapitalReductionUsingGET`: CapitalReductionView
	fmt.Fprintf(os.Stdout, "Response from `CorporateActionControllerAPI.GetCompanyCapitalReductionUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Company id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCapitalReductionUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CapitalReductionView**](CapitalReductionView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyDividendPaymentUsingGET

> DividendPaymentView GetCompanyDividendPaymentUsingGET(ctx, companyId).Execute()

Returns company's current dividend payment

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
	resp, r, err := apiClient.CorporateActionControllerAPI.GetCompanyDividendPaymentUsingGET(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CorporateActionControllerAPI.GetCompanyDividendPaymentUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyDividendPaymentUsingGET`: DividendPaymentView
	fmt.Fprintf(os.Stdout, "Response from `CorporateActionControllerAPI.GetCompanyDividendPaymentUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Company id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyDividendPaymentUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DividendPaymentView**](DividendPaymentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDividendPaymentsUsingGET

> DividendPaymentView GetDividendPaymentsUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Returns dividend payments

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
	resp, r, err := apiClient.CorporateActionControllerAPI.GetDividendPaymentsUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CorporateActionControllerAPI.GetDividendPaymentsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDividendPaymentsUsingGET`: DividendPaymentView
	fmt.Fprintf(os.Stdout, "Response from `CorporateActionControllerAPI.GetDividendPaymentsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDividendPaymentsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**DividendPaymentView**](DividendPaymentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMergerUsingGET

> CapitalReductionView GetMergerUsingGET(ctx, companyId).Execute()

Returns company's current merger

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
	resp, r, err := apiClient.CorporateActionControllerAPI.GetMergerUsingGET(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CorporateActionControllerAPI.GetMergerUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMergerUsingGET`: CapitalReductionView
	fmt.Fprintf(os.Stdout, "Response from `CorporateActionControllerAPI.GetMergerUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Company id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMergerUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CapitalReductionView**](CapitalReductionView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

