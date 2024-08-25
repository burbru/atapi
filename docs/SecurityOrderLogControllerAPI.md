# \SecurityOrderLogControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculateAverageBuyingPriceUsingGET**](SecurityOrderLogControllerAPI.md#CalculateAverageBuyingPriceUsingGET) | **Get** /api/v2/averagebuyingprice | Returns the average buying price based on security log entries
[**LiLoSecuIdentUsingGET**](SecurityOrderLogControllerAPI.md#LiLoSecuIdentUsingGET) | **Get** /api/securityorderlogs | Lists security order log entries for trades of securities identified by security identifier
[**ListOrderLogsByAsinUsingGET**](SecurityOrderLogControllerAPI.md#ListOrderLogsByAsinUsingGET) | **Get** /api/v2/securityorderlogs/by-asin/{asin} | Lists security order log entries for trades of securities identified by security identifier
[**ListOrderLogsUsingGET**](SecurityOrderLogControllerAPI.md#ListOrderLogsUsingGET) | **Get** /api/v2/securityorderlogs | Lists security order log entries for trades of securities identified by security identifier



## CalculateAverageBuyingPriceUsingGET

> float32 CalculateAverageBuyingPriceUsingGET(ctx).SecuritiesAccountId(securitiesAccountId).SecurityIdentifier(securityIdentifier).Execute()

Returns the average buying price based on security log entries

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
	securityIdentifier := "securityIdentifier_example" // string | Security Identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityOrderLogControllerAPI.CalculateAverageBuyingPriceUsingGET(context.Background()).SecuritiesAccountId(securitiesAccountId).SecurityIdentifier(securityIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityOrderLogControllerAPI.CalculateAverageBuyingPriceUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalculateAverageBuyingPriceUsingGET`: float32
	fmt.Fprintf(os.Stdout, "Response from `SecurityOrderLogControllerAPI.CalculateAverageBuyingPriceUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCalculateAverageBuyingPriceUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securitiesAccountId** | **string** | Securities account id | 
 **securityIdentifier** | **string** | Security Identifier | 

### Return type

**float32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LiLoSecuIdentUsingGET

> SecurityOrderLogEntryView LiLoSecuIdentUsingGET(ctx).BuyerSecuritiesAccountId(buyerSecuritiesAccountId).EndDate(endDate).SecurityIdentifier(securityIdentifier).SellerSecuritiesAccountId(sellerSecuritiesAccountId).StartDate(startDate).Execute()

Lists security order log entries for trades of securities identified by security identifier

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
	buyerSecuritiesAccountId := "buyerSecuritiesAccountId_example" // string | Buyer's securities account id (optional)
	endDate := "endDate_example" // string | End date (optional)
	securityIdentifier := "securityIdentifier_example" // string | Security identifier (optional)
	sellerSecuritiesAccountId := "sellerSecuritiesAccountId_example" // string | Seller's securities account id (optional)
	startDate := "startDate_example" // string | Start date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityOrderLogControllerAPI.LiLoSecuIdentUsingGET(context.Background()).BuyerSecuritiesAccountId(buyerSecuritiesAccountId).EndDate(endDate).SecurityIdentifier(securityIdentifier).SellerSecuritiesAccountId(sellerSecuritiesAccountId).StartDate(startDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityOrderLogControllerAPI.LiLoSecuIdentUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LiLoSecuIdentUsingGET`: SecurityOrderLogEntryView
	fmt.Fprintf(os.Stdout, "Response from `SecurityOrderLogControllerAPI.LiLoSecuIdentUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLiLoSecuIdentUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buyerSecuritiesAccountId** | **string** | Buyer&#39;s securities account id | 
 **endDate** | **string** | End date | 
 **securityIdentifier** | **string** | Security identifier | 
 **sellerSecuritiesAccountId** | **string** | Seller&#39;s securities account id | 
 **startDate** | **string** | Start date | 

### Return type

[**SecurityOrderLogEntryView**](SecurityOrderLogEntryView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrderLogsByAsinUsingGET

> SecurityOrderLogEntryView ListOrderLogsByAsinUsingGET(ctx, asin).Page(page).Size(size).Sort(sort).Execute()

Lists security order log entries for trades of securities identified by security identifier

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
	asin := "asin_example" // string | ASIN
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityOrderLogControllerAPI.ListOrderLogsByAsinUsingGET(context.Background(), asin).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityOrderLogControllerAPI.ListOrderLogsByAsinUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrderLogsByAsinUsingGET`: SecurityOrderLogEntryView
	fmt.Fprintf(os.Stdout, "Response from `SecurityOrderLogControllerAPI.ListOrderLogsByAsinUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**asin** | **string** | ASIN | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrderLogsByAsinUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**SecurityOrderLogEntryView**](SecurityOrderLogEntryView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrderLogsUsingGET

> SecurityOrderLogEntryView ListOrderLogsUsingGET(ctx).SecuritiesAccountId(securitiesAccountId).Page(page).Search(search).Size(size).Sort(sort).Execute()

Lists security order log entries for trades of securities identified by security identifier

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
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	search := "search_example" // string | Fulltext search (be aware of non-localized subject and content and be aware of MessagePrototype's structure with substitutions; useful for security identifiers) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityOrderLogControllerAPI.ListOrderLogsUsingGET(context.Background()).SecuritiesAccountId(securitiesAccountId).Page(page).Search(search).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityOrderLogControllerAPI.ListOrderLogsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrderLogsUsingGET`: SecurityOrderLogEntryView
	fmt.Fprintf(os.Stdout, "Response from `SecurityOrderLogControllerAPI.ListOrderLogsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrderLogsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securitiesAccountId** | **string** | Securities account id | 
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **search** | **string** | Fulltext search (be aware of non-localized subject and content and be aware of MessagePrototype&#39;s structure with substitutions; useful for security identifiers) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**SecurityOrderLogEntryView**](SecurityOrderLogEntryView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

