# \SystemBondControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemBondBySecIdentUsingGET**](SystemBondControllerAPI.md#GetSystemBondBySecIdentUsingGET) | **Get** /api/systembonds/securityidentifier/{securityIdentifier} | Returns bond by its security identifier
[**GetSystemBondUsingGET**](SystemBondControllerAPI.md#GetSystemBondUsingGET) | **Get** /api/systembonds/{bondId} | Returns system bond
[**IssueSystemBondsUsingPOST**](SystemBondControllerAPI.md#IssueSystemBondsUsingPOST) | **Post** /api/systembonds | Issues new bonds to be bought by the central bank at the current main interest rate
[**ListsBondsUsingGET**](SystemBondControllerAPI.md#ListsBondsUsingGET) | **Get** /api/systembonds | Lists all bonds



## GetSystemBondBySecIdentUsingGET

> SystemBondView GetSystemBondBySecIdentUsingGET(ctx, securityIdentifier).Execute()

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
	resp, r, err := apiClient.SystemBondControllerAPI.GetSystemBondBySecIdentUsingGET(context.Background(), securityIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemBondControllerAPI.GetSystemBondBySecIdentUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemBondBySecIdentUsingGET`: SystemBondView
	fmt.Fprintf(os.Stdout, "Response from `SystemBondControllerAPI.GetSystemBondBySecIdentUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityIdentifier** | **string** | Security Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemBondBySecIdentUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SystemBondView**](SystemBondView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemBondUsingGET

> SystemBondView GetSystemBondUsingGET(ctx, bondId).Execute()

Returns system bond

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
	bondId := "bondId_example" // string | System bond id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemBondControllerAPI.GetSystemBondUsingGET(context.Background(), bondId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemBondControllerAPI.GetSystemBondUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemBondUsingGET`: SystemBondView
	fmt.Fprintf(os.Stdout, "Response from `SystemBondControllerAPI.GetSystemBondUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bondId** | **string** | System bond id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemBondUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SystemBondView**](SystemBondView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueSystemBondsUsingPOST

> SystemBondView IssueSystemBondsUsingPOST(ctx).CompanyId(companyId).NumberOfBonds(numberOfBonds).Execute()

Issues new bonds to be bought by the central bank at the current main interest rate

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
	numberOfBonds := "numberOfBonds_example" // string | Number of bonds to emit

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemBondControllerAPI.IssueSystemBondsUsingPOST(context.Background()).CompanyId(companyId).NumberOfBonds(numberOfBonds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemBondControllerAPI.IssueSystemBondsUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueSystemBondsUsingPOST`: SystemBondView
	fmt.Fprintf(os.Stdout, "Response from `SystemBondControllerAPI.IssueSystemBondsUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueSystemBondsUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company id | 
 **numberOfBonds** | **string** | Number of bonds to emit | 

### Return type

[**SystemBondView**](SystemBondView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsBondsUsingGET

> BondView ListsBondsUsingGET(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemBondControllerAPI.ListsBondsUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemBondControllerAPI.ListsBondsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListsBondsUsingGET`: BondView
	fmt.Fprintf(os.Stdout, "Response from `SystemBondControllerAPI.ListsBondsUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListsBondsUsingGETRequest struct via the builder pattern


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

