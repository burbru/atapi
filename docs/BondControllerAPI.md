# \BondControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBondBySecIdentUsingGET**](BondControllerAPI.md#GetBondBySecIdentUsingGET) | **Get** /api/bonds/securityidentifier/{securityIdentifier} | Returns bond by its security identifier
[**GetBondUsingGET**](BondControllerAPI.md#GetBondUsingGET) | **Get** /api/bonds/{bondId} | Returns bond
[**IssueBondsUsingPOST**](BondControllerAPI.md#IssueBondsUsingPOST) | **Post** /api/bonds | Issues new corporate bonds
[**ListBondsUsingGET**](BondControllerAPI.md#ListBondsUsingGET) | **Get** /api/v2/bonds | Lists all bonds
[**ListBondsV1UsingGET**](BondControllerAPI.md#ListBondsV1UsingGET) | **Get** /api/bonds | Lists all bonds (200 max.)



## GetBondBySecIdentUsingGET

> BondView GetBondBySecIdentUsingGET(ctx, securityIdentifier).Execute()

Returns bond by its security identifier

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
	resp, r, err := apiClient.BondControllerAPI.GetBondBySecIdentUsingGET(context.Background(), securityIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BondControllerAPI.GetBondBySecIdentUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBondBySecIdentUsingGET`: BondView
	fmt.Fprintf(os.Stdout, "Response from `BondControllerAPI.GetBondBySecIdentUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityIdentifier** | **string** | Security Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBondBySecIdentUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BondView**](BondView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBondUsingGET

> BondView GetBondUsingGET(ctx, bondId).Execute()

Returns bond

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
	bondId := "bondId_example" // string | Bond id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BondControllerAPI.GetBondUsingGET(context.Background(), bondId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BondControllerAPI.GetBondUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBondUsingGET`: BondView
	fmt.Fprintf(os.Stdout, "Response from `BondControllerAPI.GetBondUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bondId** | **string** | Bond id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBondUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BondView**](BondView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueBondsUsingPOST

> BondView IssueBondsUsingPOST(ctx).CompanyId(companyId).FaceValue(faceValue).InterestRate(interestRate).MaturityDate(maturityDate).NumberOfBonds(numberOfBonds).Execute()

Issues new corporate bonds

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
	faceValue := "faceValue_example" // string | Face value
	interestRate := "interestRate_example" // string | Interest Rate in percent to pay at maturity date
	maturityDate := "maturityDate_example" // string | Maturity date
	numberOfBonds := "numberOfBonds_example" // string | Number of bonds to emit

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BondControllerAPI.IssueBondsUsingPOST(context.Background()).CompanyId(companyId).FaceValue(faceValue).InterestRate(interestRate).MaturityDate(maturityDate).NumberOfBonds(numberOfBonds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BondControllerAPI.IssueBondsUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueBondsUsingPOST`: BondView
	fmt.Fprintf(os.Stdout, "Response from `BondControllerAPI.IssueBondsUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueBondsUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company id | 
 **faceValue** | **string** | Face value | 
 **interestRate** | **string** | Interest Rate in percent to pay at maturity date | 
 **maturityDate** | **string** | Maturity date | 
 **numberOfBonds** | **string** | Number of bonds to emit | 

### Return type

[**BondView**](BondView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBondsUsingGET

> BondView ListBondsUsingGET(ctx).Page(page).Search(search).Size(size).Sort(sort).Execute()

Lists all bonds



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
	search := "search_example" // string | Fulltext search (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BondControllerAPI.ListBondsUsingGET(context.Background()).Page(page).Search(search).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BondControllerAPI.ListBondsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBondsUsingGET`: BondView
	fmt.Fprintf(os.Stdout, "Response from `BondControllerAPI.ListBondsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBondsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **search** | **string** | Fulltext search | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**BondView**](BondView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBondsV1UsingGET

> BondView ListBondsV1UsingGET(ctx).Execute()

Lists all bonds (200 max.)

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
	resp, r, err := apiClient.BondControllerAPI.ListBondsV1UsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BondControllerAPI.ListBondsV1UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBondsV1UsingGET`: BondView
	fmt.Fprintf(os.Stdout, "Response from `BondControllerAPI.ListBondsV1UsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBondsV1UsingGETRequest struct via the builder pattern


### Return type

[**BondView**](BondView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

