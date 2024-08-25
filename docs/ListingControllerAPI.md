# \ListingControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindListingsByAsinPartUsingGET**](ListingControllerAPI.md#FindListingsByAsinPartUsingGET) | **Get** /api/search/listings/{securityIdentifierPart} | Finds listing by part of their security identifiers
[**FindListingsByAsinUsingGET**](ListingControllerAPI.md#FindListingsByAsinUsingGET) | **Get** /api/search/listings/ | Finds listing by part of their security identifiers
[**GetListingProfileUsingGET**](ListingControllerAPI.md#GetListingProfileUsingGET) | **Get** /api/v2/listingprofiles | Returns listing profile by security identifier
[**GetListingUsingGET**](ListingControllerAPI.md#GetListingUsingGET) | **Get** /api/listings/{securityIdentifier} | Returns listing
[**GetOutstandingSharesUsingGET**](ListingControllerAPI.md#GetOutstandingSharesUsingGET) | **Get** /api/listings/outstandingshares/{securityIdentifier} | Returns Outstanding Shares
[**GetProfileV1UsingGET**](ListingControllerAPI.md#GetProfileV1UsingGET) | **Get** /api/listingprofiles/{securityIdentifier} | Gets listing profile
[**ListHistorizedCompanyDataUsingGET**](ListingControllerAPI.md#ListHistorizedCompanyDataUsingGET) | **Get** /api/v2/historizedcompanydata/{securityIdentifier} | Lists all historized company data
[**ListHistorizedListingDataUsingGET**](ListingControllerAPI.md#ListHistorizedListingDataUsingGET) | **Get** /api/v2/historizedlistingdata/{securityIdentifier} | Lists all historized listing data
[**ListListingsUsingGET**](ListingControllerAPI.md#ListListingsUsingGET) | **Get** /api/v2/listings | Lists all listings without end date
[**ListListingsV1UsingGET**](ListingControllerAPI.md#ListListingsV1UsingGET) | **Get** /api/listings/ | Lists all listings without end date (200 max.)
[**SearchListingsUsingGET**](ListingControllerAPI.md#SearchListingsUsingGET) | **Get** /api/shareholders/{securityIdentifier} | Lists Shareholders



## FindListingsByAsinPartUsingGET

> ListingView FindListingsByAsinPartUsingGET(ctx, securityIdentifierPart).Execute()

Finds listing by part of their security identifiers

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
	securityIdentifierPart := "securityIdentifierPart_example" // string | Part of security identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListingControllerAPI.FindListingsByAsinPartUsingGET(context.Background(), securityIdentifierPart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingControllerAPI.FindListingsByAsinPartUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindListingsByAsinPartUsingGET`: ListingView
	fmt.Fprintf(os.Stdout, "Response from `ListingControllerAPI.FindListingsByAsinPartUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityIdentifierPart** | **string** | Part of security identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindListingsByAsinPartUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListingView**](ListingView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindListingsByAsinUsingGET

> ListingView FindListingsByAsinUsingGET(ctx).Execute()

Finds listing by part of their security identifiers

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
	resp, r, err := apiClient.ListingControllerAPI.FindListingsByAsinUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingControllerAPI.FindListingsByAsinUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindListingsByAsinUsingGET`: ListingView
	fmt.Fprintf(os.Stdout, "Response from `ListingControllerAPI.FindListingsByAsinUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindListingsByAsinUsingGETRequest struct via the builder pattern


### Return type

[**ListingView**](ListingView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListingProfileUsingGET

> ListingProfileView GetListingProfileUsingGET(ctx).SecurityIdentifier(securityIdentifier).Execute()

Returns listing profile by security identifier

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
	securityIdentifier := "securityIdentifier_example" // string | Security identifier (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListingControllerAPI.GetListingProfileUsingGET(context.Background()).SecurityIdentifier(securityIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingControllerAPI.GetListingProfileUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListingProfileUsingGET`: ListingProfileView
	fmt.Fprintf(os.Stdout, "Response from `ListingControllerAPI.GetListingProfileUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetListingProfileUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityIdentifier** | **string** | Security identifier | 

### Return type

[**ListingProfileView**](ListingProfileView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListingUsingGET

> ListingView GetListingUsingGET(ctx, securityIdentifier).Execute()

Returns listing

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
	securityIdentifier := "securityIdentifier_example" // string | Security Identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListingControllerAPI.GetListingUsingGET(context.Background(), securityIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingControllerAPI.GetListingUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListingUsingGET`: ListingView
	fmt.Fprintf(os.Stdout, "Response from `ListingControllerAPI.GetListingUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityIdentifier** | **string** | Security Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListingView**](ListingView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutstandingSharesUsingGET

> int64 GetOutstandingSharesUsingGET(ctx, securityIdentifier).Execute()

Returns Outstanding Shares

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
	securityIdentifier := "securityIdentifier_example" // string | Security Identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListingControllerAPI.GetOutstandingSharesUsingGET(context.Background(), securityIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingControllerAPI.GetOutstandingSharesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutstandingSharesUsingGET`: int64
	fmt.Fprintf(os.Stdout, "Response from `ListingControllerAPI.GetOutstandingSharesUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityIdentifier** | **string** | Security Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutstandingSharesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfileV1UsingGET

> ListingProfileView GetProfileV1UsingGET(ctx, securityIdentifier).SecuritiesAccountId(securitiesAccountId).Execute()

Gets listing profile

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
	securitiesAccountId := "securitiesAccountId_example" // string | Shareholder's securities account id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListingControllerAPI.GetProfileV1UsingGET(context.Background(), securityIdentifier).SecuritiesAccountId(securitiesAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingControllerAPI.GetProfileV1UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfileV1UsingGET`: ListingProfileView
	fmt.Fprintf(os.Stdout, "Response from `ListingControllerAPI.GetProfileV1UsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityIdentifier** | **string** | Security identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileV1UsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **securitiesAccountId** | **string** | Shareholder&#39;s securities account id | 

### Return type

[**ListingProfileView**](ListingProfileView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHistorizedCompanyDataUsingGET

> PageHistorizedCompanyDataView ListHistorizedCompanyDataUsingGET(ctx, securityIdentifier).Page(page).Size(size).Sort(sort).Execute()

Lists all historized company data



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
	securityIdentifier := "securityIdentifier_example" // string | Security Identifier
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListingControllerAPI.ListHistorizedCompanyDataUsingGET(context.Background(), securityIdentifier).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingControllerAPI.ListHistorizedCompanyDataUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHistorizedCompanyDataUsingGET`: PageHistorizedCompanyDataView
	fmt.Fprintf(os.Stdout, "Response from `ListingControllerAPI.ListHistorizedCompanyDataUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityIdentifier** | **string** | Security Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHistorizedCompanyDataUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageHistorizedCompanyDataView**](PageHistorizedCompanyDataView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHistorizedListingDataUsingGET

> PageHistorizedListingDataView ListHistorizedListingDataUsingGET(ctx, securityIdentifier).Page(page).Size(size).Sort(sort).Execute()

Lists all historized listing data



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
	securityIdentifier := "securityIdentifier_example" // string | Security Identifier
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListingControllerAPI.ListHistorizedListingDataUsingGET(context.Background(), securityIdentifier).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingControllerAPI.ListHistorizedListingDataUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHistorizedListingDataUsingGET`: PageHistorizedListingDataView
	fmt.Fprintf(os.Stdout, "Response from `ListingControllerAPI.ListHistorizedListingDataUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityIdentifier** | **string** | Security Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHistorizedListingDataUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageHistorizedListingDataView**](PageHistorizedListingDataView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListListingsUsingGET

> PageListingView ListListingsUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Lists all listings without end date



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
	resp, r, err := apiClient.ListingControllerAPI.ListListingsUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingControllerAPI.ListListingsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListListingsUsingGET`: PageListingView
	fmt.Fprintf(os.Stdout, "Response from `ListingControllerAPI.ListListingsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListListingsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageListingView**](PageListingView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListListingsV1UsingGET

> ListingView ListListingsV1UsingGET(ctx).Execute()

Lists all listings without end date (200 max.)

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
	resp, r, err := apiClient.ListingControllerAPI.ListListingsV1UsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingControllerAPI.ListListingsV1UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListListingsV1UsingGET`: ListingView
	fmt.Fprintf(os.Stdout, "Response from `ListingControllerAPI.ListListingsV1UsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListListingsV1UsingGETRequest struct via the builder pattern


### Return type

[**ListingView**](ListingView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchListingsUsingGET

> ShareholderView SearchListingsUsingGET(ctx, securityIdentifier).Execute()

Lists Shareholders

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
	securityIdentifier := "securityIdentifier_example" // string | Security Identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListingControllerAPI.SearchListingsUsingGET(context.Background(), securityIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingControllerAPI.SearchListingsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchListingsUsingGET`: ShareholderView
	fmt.Fprintf(os.Stdout, "Response from `ListingControllerAPI.SearchListingsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityIdentifier** | **string** | Security Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchListingsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShareholderView**](ShareholderView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

