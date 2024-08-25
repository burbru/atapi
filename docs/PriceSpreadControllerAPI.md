# \PriceSpreadControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFilterUsingDELETE**](PriceSpreadControllerAPI.md#DeleteFilterUsingDELETE) | **Delete** /api/v2/filter/pricespreads/{filterId} | Deletes filter
[**FilterSpreadsUsingPOST**](PriceSpreadControllerAPI.md#FilterSpreadsUsingPOST) | **Post** /api/v2/filter/pricespreads | Filters price spreads
[**GetBiggestTradedSecuritiesUsingGET**](PriceSpreadControllerAPI.md#GetBiggestTradedSecuritiesUsingGET) | **Get** /api/v2/biggesttradedsecurities | Lists traded securities with the biggest trading volume of the last 24h
[**GetFilterDefinitionUsingGET**](PriceSpreadControllerAPI.md#GetFilterDefinitionUsingGET) | **Get** /api/v2/filterdefinition/pricespreads | Price spread filter definition
[**GetFilterUsingGET**](PriceSpreadControllerAPI.md#GetFilterUsingGET) | **Get** /api/v2/filter/pricespreads/{filterId} | Returns filter
[**GetMostFrequentlyTradedSecuritiesUsingGET**](PriceSpreadControllerAPI.md#GetMostFrequentlyTradedSecuritiesUsingGET) | **Get** /api/v2/mostfrequentlytradedsecurities | Lists most frequently traded securities of the last 24h
[**GetSecuritiesWithBigPriceChangesUsingGET**](PriceSpreadControllerAPI.md#GetSecuritiesWithBigPriceChangesUsingGET) | **Get** /api/v2/securitieswithbigpricechanges | Lists securities sorted by their price change between the last two closing prices
[**GetSpreadUsingGET**](PriceSpreadControllerAPI.md#GetSpreadUsingGET) | **Get** /api/pricespreads/{securityIdentifier} | Returns price spread
[**ListSpreadsUsingGET**](PriceSpreadControllerAPI.md#ListSpreadsUsingGET) | **Get** /api/v2/pricespreads | Lists all price spreads
[**ListSpreadsV1UsingGET**](PriceSpreadControllerAPI.md#ListSpreadsV1UsingGET) | **Get** /api/pricespreads/ | Lists all price spreads
[**ListsFilterUsingGET**](PriceSpreadControllerAPI.md#ListsFilterUsingGET) | **Get** /api/v2/filter/pricespreads | Lists current user&#39;s filters



## DeleteFilterUsingDELETE

> MessagePrototype DeleteFilterUsingDELETE(ctx, filterId).Execute()

Deletes filter

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
	filterId := "filterId_example" // string | Id of persistent filter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceSpreadControllerAPI.DeleteFilterUsingDELETE(context.Background(), filterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceSpreadControllerAPI.DeleteFilterUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFilterUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `PriceSpreadControllerAPI.DeleteFilterUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterId** | **string** | Id of persistent filter | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFilterUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessagePrototype**](MessagePrototype.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilterSpreadsUsingPOST

> ListingMarketFilterResultView FilterSpreadsUsingPOST(ctx).FilterId(filterId).Filter(filter).Execute()

Filters price spreads



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
	filterId := "filterId_example" // string | Filter ID (duplicates the filter, can be used for other user's filter too) (optional)
	filter := *openapiclient.NewPriceSpreadListingViewFilter() // PriceSpreadListingViewFilter | Filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceSpreadControllerAPI.FilterSpreadsUsingPOST(context.Background()).FilterId(filterId).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceSpreadControllerAPI.FilterSpreadsUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilterSpreadsUsingPOST`: ListingMarketFilterResultView
	fmt.Fprintf(os.Stdout, "Response from `PriceSpreadControllerAPI.FilterSpreadsUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilterSpreadsUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterId** | **string** | Filter ID (duplicates the filter, can be used for other user&#39;s filter too) | 
 **filter** | [**PriceSpreadListingViewFilter**](PriceSpreadListingViewFilter.md) | Filter | 

### Return type

[**ListingMarketFilterResultView**](ListingMarketFilterResultView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBiggestTradedSecuritiesUsingGET

> ListingWithTradingCountView GetBiggestTradedSecuritiesUsingGET(ctx).Page(page).Size(size).Sort(sort).Type_(type_).Execute()

Lists traded securities with the biggest trading volume of the last 24h

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
	type_ := "type__example" // string | Listing type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceSpreadControllerAPI.GetBiggestTradedSecuritiesUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceSpreadControllerAPI.GetBiggestTradedSecuritiesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBiggestTradedSecuritiesUsingGET`: ListingWithTradingCountView
	fmt.Fprintf(os.Stdout, "Response from `PriceSpreadControllerAPI.GetBiggestTradedSecuritiesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBiggestTradedSecuritiesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **type_** | **string** | Listing type | 

### Return type

[**ListingWithTradingCountView**](ListingWithTradingCountView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilterDefinitionUsingGET

> PriceSpreadListingFilterDefinition GetFilterDefinitionUsingGET(ctx).Execute()

Price spread filter definition

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
	resp, r, err := apiClient.PriceSpreadControllerAPI.GetFilterDefinitionUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceSpreadControllerAPI.GetFilterDefinitionUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilterDefinitionUsingGET`: PriceSpreadListingFilterDefinition
	fmt.Fprintf(os.Stdout, "Response from `PriceSpreadControllerAPI.GetFilterDefinitionUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilterDefinitionUsingGETRequest struct via the builder pattern


### Return type

[**PriceSpreadListingFilterDefinition**](PriceSpreadListingFilterDefinition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilterUsingGET

> PriceSpreadListingViewFilter GetFilterUsingGET(ctx, filterId).Execute()

Returns filter

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
	filterId := "filterId_example" // string | Id of persistent filter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceSpreadControllerAPI.GetFilterUsingGET(context.Background(), filterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceSpreadControllerAPI.GetFilterUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilterUsingGET`: PriceSpreadListingViewFilter
	fmt.Fprintf(os.Stdout, "Response from `PriceSpreadControllerAPI.GetFilterUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterId** | **string** | Id of persistent filter | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilterUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PriceSpreadListingViewFilter**](PriceSpreadListingViewFilter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMostFrequentlyTradedSecuritiesUsingGET

> ListingWithTradingCountView GetMostFrequentlyTradedSecuritiesUsingGET(ctx).Page(page).Size(size).Sort(sort).Type_(type_).Execute()

Lists most frequently traded securities of the last 24h

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
	type_ := "type__example" // string | Listing type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceSpreadControllerAPI.GetMostFrequentlyTradedSecuritiesUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceSpreadControllerAPI.GetMostFrequentlyTradedSecuritiesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMostFrequentlyTradedSecuritiesUsingGET`: ListingWithTradingCountView
	fmt.Fprintf(os.Stdout, "Response from `PriceSpreadControllerAPI.GetMostFrequentlyTradedSecuritiesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMostFrequentlyTradedSecuritiesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **type_** | **string** | Listing type | 

### Return type

[**ListingWithTradingCountView**](ListingWithTradingCountView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecuritiesWithBigPriceChangesUsingGET

> ListingWithPriceChangeView GetSecuritiesWithBigPriceChangesUsingGET(ctx).LosersFirst(losersFirst).Page(page).Size(size).Sort(sort).Execute()

Lists securities sorted by their price change between the last two closing prices

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
	losersFirst := true // bool | List losers first (optional)
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceSpreadControllerAPI.GetSecuritiesWithBigPriceChangesUsingGET(context.Background()).LosersFirst(losersFirst).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceSpreadControllerAPI.GetSecuritiesWithBigPriceChangesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecuritiesWithBigPriceChangesUsingGET`: ListingWithPriceChangeView
	fmt.Fprintf(os.Stdout, "Response from `PriceSpreadControllerAPI.GetSecuritiesWithBigPriceChangesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecuritiesWithBigPriceChangesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **losersFirst** | **bool** | List losers first | 
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**ListingWithPriceChangeView**](ListingWithPriceChangeView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpreadUsingGET

> PriceSpreadListingView GetSpreadUsingGET(ctx, securityIdentifier).Execute()

Returns price spread

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceSpreadControllerAPI.GetSpreadUsingGET(context.Background(), securityIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceSpreadControllerAPI.GetSpreadUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpreadUsingGET`: PriceSpreadListingView
	fmt.Fprintf(os.Stdout, "Response from `PriceSpreadControllerAPI.GetSpreadUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityIdentifier** | **string** | Security identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpreadUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PriceSpreadListingView**](PriceSpreadListingView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpreadsUsingGET

> PagePriceSpreadListingView ListSpreadsUsingGET(ctx).Page(page).Search(search).Size(size).Sort(sort).Execute()

Lists all price spreads



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
	search := "search_example" // string | Fulltext search in listing's name and security identifier (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceSpreadControllerAPI.ListSpreadsUsingGET(context.Background()).Page(page).Search(search).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceSpreadControllerAPI.ListSpreadsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSpreadsUsingGET`: PagePriceSpreadListingView
	fmt.Fprintf(os.Stdout, "Response from `PriceSpreadControllerAPI.ListSpreadsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSpreadsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **search** | **string** | Fulltext search in listing&#39;s name and security identifier | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PagePriceSpreadListingView**](PagePriceSpreadListingView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpreadsV1UsingGET

> PriceSpreadListingView ListSpreadsV1UsingGET(ctx).Execute()

Lists all price spreads

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
	resp, r, err := apiClient.PriceSpreadControllerAPI.ListSpreadsV1UsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceSpreadControllerAPI.ListSpreadsV1UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSpreadsV1UsingGET`: PriceSpreadListingView
	fmt.Fprintf(os.Stdout, "Response from `PriceSpreadControllerAPI.ListSpreadsV1UsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSpreadsV1UsingGETRequest struct via the builder pattern


### Return type

[**PriceSpreadListingView**](PriceSpreadListingView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsFilterUsingGET

> PagePersistentFilterView ListsFilterUsingGET(ctx).Page(page).Search(search).Size(size).Sort(sort).Execute()

Lists current user's filters

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
	search := "search_example" // string | Search for name parts (optional) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceSpreadControllerAPI.ListsFilterUsingGET(context.Background()).Page(page).Search(search).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceSpreadControllerAPI.ListsFilterUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListsFilterUsingGET`: PagePersistentFilterView
	fmt.Fprintf(os.Stdout, "Response from `PriceSpreadControllerAPI.ListsFilterUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListsFilterUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **search** | **string** | Search for name parts (optional) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PagePersistentFilterView**](PagePersistentFilterView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

