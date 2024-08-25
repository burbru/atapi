# \EmpireControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompanyHistogramsUsingGET**](EmpireControllerAPI.md#GetCompanyHistogramsUsingGET) | **Get** /api/v2/companyhistograms/{companyId} | Returns a company&#39;s histograms
[**GetOwnCompaniesByEmpireShareUsingGET**](EmpireControllerAPI.md#GetOwnCompaniesByEmpireShareUsingGET) | **Get** /api/v2/my/companiesbyempireshare | Lists own companies in descending order according empire share in percent
[**GetSuggestionsUsingGET**](EmpireControllerAPI.md#GetSuggestionsUsingGET) | **Get** /api/v2/suggestions | Lists current logged in user&#39;s suggestions
[**GetTakeOverPossibilitiesUsingGET**](EmpireControllerAPI.md#GetTakeOverPossibilitiesUsingGET) | **Get** /api/v2/my/takeoverpossibilities | Lists takeover possibilities in descending order according to their share in percent
[**ListCompanyHistogramsUsingGET**](EmpireControllerAPI.md#ListCompanyHistogramsUsingGET) | **Get** /api/v2/my/companyhistograms | Lists companies&#39; histograms



## GetCompanyHistogramsUsingGET

> CompanyHistogramView GetCompanyHistogramsUsingGET(ctx, companyId).Execute()

Returns a company's histograms

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
	resp, r, err := apiClient.EmpireControllerAPI.GetCompanyHistogramsUsingGET(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmpireControllerAPI.GetCompanyHistogramsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyHistogramsUsingGET`: CompanyHistogramView
	fmt.Fprintf(os.Stdout, "Response from `EmpireControllerAPI.GetCompanyHistogramsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Company id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyHistogramsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyHistogramView**](CompanyHistogramView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwnCompaniesByEmpireShareUsingGET

> PageListingShareView GetOwnCompaniesByEmpireShareUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Lists own companies in descending order according empire share in percent

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
	resp, r, err := apiClient.EmpireControllerAPI.GetOwnCompaniesByEmpireShareUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmpireControllerAPI.GetOwnCompaniesByEmpireShareUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwnCompaniesByEmpireShareUsingGET`: PageListingShareView
	fmt.Fprintf(os.Stdout, "Response from `EmpireControllerAPI.GetOwnCompaniesByEmpireShareUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnCompaniesByEmpireShareUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageListingShareView**](PageListingShareView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuggestionsUsingGET

> PageSuggestion GetSuggestionsUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Lists current logged in user's suggestions

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
	resp, r, err := apiClient.EmpireControllerAPI.GetSuggestionsUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmpireControllerAPI.GetSuggestionsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSuggestionsUsingGET`: PageSuggestion
	fmt.Fprintf(os.Stdout, "Response from `EmpireControllerAPI.GetSuggestionsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSuggestionsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageSuggestion**](PageSuggestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTakeOverPossibilitiesUsingGET

> PageListingShareView GetTakeOverPossibilitiesUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Lists takeover possibilities in descending order according to their share in percent

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
	resp, r, err := apiClient.EmpireControllerAPI.GetTakeOverPossibilitiesUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmpireControllerAPI.GetTakeOverPossibilitiesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTakeOverPossibilitiesUsingGET`: PageListingShareView
	fmt.Fprintf(os.Stdout, "Response from `EmpireControllerAPI.GetTakeOverPossibilitiesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTakeOverPossibilitiesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageListingShareView**](PageListingShareView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompanyHistogramsUsingGET

> PageCompanyHistogramView ListCompanyHistogramsUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Lists companies' histograms

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
	resp, r, err := apiClient.EmpireControllerAPI.ListCompanyHistogramsUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmpireControllerAPI.ListCompanyHistogramsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompanyHistogramsUsingGET`: PageCompanyHistogramView
	fmt.Fprintf(os.Stdout, "Response from `EmpireControllerAPI.ListCompanyHistogramsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCompanyHistogramsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageCompanyHistogramView**](PageCompanyHistogramView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

