# \CompanyControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLogoUsingPUT**](CompanyControllerAPI.md#AddLogoUsingPUT) | **Put** /api/companies/logo/{companyId} | Adds logo to company
[**CreateCompanyUsingPOST**](CompanyControllerAPI.md#CreateCompanyUsingPOST) | **Post** /api/companies | Creates listed company
[**GetCompanyBySecAccIdUsingGET**](CompanyControllerAPI.md#GetCompanyBySecAccIdUsingGET) | **Get** /api/companies/securitiesaccount/{securitiesAccountId} | Returns company by its securities account&#39;s id
[**GetCompanyBySecurityIdentifierUsingGET**](CompanyControllerAPI.md#GetCompanyBySecurityIdentifierUsingGET) | **Get** /api/companies/securityIdentifier/{securityIdentifier} | Returns company by securityIdentifier
[**GetCompanyCapsUsingGET**](CompanyControllerAPI.md#GetCompanyCapsUsingGET) | **Get** /api/v2/companycaps/{companyId} | Returns company capabilities
[**GetCompanyProfileUsingGET**](CompanyControllerAPI.md#GetCompanyProfileUsingGET) | **Get** /api/companyprofiles/{companyId} | Returns company profile
[**GetCompanyUsingGET**](CompanyControllerAPI.md#GetCompanyUsingGET) | **Get** /api/companies/{companyId} | Returns company
[**ListAllCompaniesUsingGET**](CompanyControllerAPI.md#ListAllCompaniesUsingGET) | **Get** /api/v2/companies | Lists all companies
[**ListAllCompaniesV1UsingGET**](CompanyControllerAPI.md#ListAllCompaniesV1UsingGET) | **Get** /api/companies/all/ | Lists all companies
[**ListCompaniesByUserIdUsingGET**](CompanyControllerAPI.md#ListCompaniesByUserIdUsingGET) | **Get** /api/companies/ceo/userid/{userId} | Lists companies controlled by ceo with user id
[**ListCompaniesByUserNameUsingGET**](CompanyControllerAPI.md#ListCompaniesByUserNameUsingGET) | **Get** /api/companies/ceo/username/{username} | Lists companies controlled by ceo with username
[**ListMyCompaniesUsingGET**](CompanyControllerAPI.md#ListMyCompaniesUsingGET) | **Get** /api/v2/my/companies | Lists companies controlled by logged in user
[**ListMyCompaniesV1UsingGET**](CompanyControllerAPI.md#ListMyCompaniesV1UsingGET) | **Get** /api/companies | Lists companies controlled by logged in user
[**RemoveLogoUsingDELETE**](CompanyControllerAPI.md#RemoveLogoUsingDELETE) | **Delete** /api/companies/logo/{companyId} | Removes logo
[**SearchCompaniesUsingGET**](CompanyControllerAPI.md#SearchCompaniesUsingGET) | **Get** /api/search/companies/{namePart} | Finds companies by name part



## AddLogoUsingPUT

> CompanyView AddLogoUsingPUT(ctx, companyId).LogoUrl(logoUrl).Execute()

Adds logo to company

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
	logoUrl := "logoUrl_example" // string | Logo URL

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyControllerAPI.AddLogoUsingPUT(context.Background(), companyId).LogoUrl(logoUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyControllerAPI.AddLogoUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddLogoUsingPUT`: CompanyView
	fmt.Fprintf(os.Stdout, "Response from `CompanyControllerAPI.AddLogoUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Company id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLogoUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logoUrl** | **string** | Logo URL | 

### Return type

[**CompanyView**](CompanyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCompanyUsingPOST

> CompanyView CreateCompanyUsingPOST(ctx).CashDeposit(cashDeposit).Name(name).CustomAsin(customAsin).CustomNumberOfShares(customNumberOfShares).Execute()

Creates listed company

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
	cashDeposit := "cashDeposit_example" // string | Cash deposit amount
	name := "name_example" // string | Name
	customAsin := "customAsin_example" // string | Custom ASIN (premium feature) (optional)
	customNumberOfShares := int64(789) // int64 | Custom number of shares (premium feature) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyControllerAPI.CreateCompanyUsingPOST(context.Background()).CashDeposit(cashDeposit).Name(name).CustomAsin(customAsin).CustomNumberOfShares(customNumberOfShares).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyControllerAPI.CreateCompanyUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCompanyUsingPOST`: CompanyView
	fmt.Fprintf(os.Stdout, "Response from `CompanyControllerAPI.CreateCompanyUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCompanyUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cashDeposit** | **string** | Cash deposit amount | 
 **name** | **string** | Name | 
 **customAsin** | **string** | Custom ASIN (premium feature) | 
 **customNumberOfShares** | **int64** | Custom number of shares (premium feature) | 

### Return type

[**CompanyView**](CompanyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyBySecAccIdUsingGET

> CompanyView GetCompanyBySecAccIdUsingGET(ctx, securitiesAccountId).Execute()

Returns company by its securities account's id

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyControllerAPI.GetCompanyBySecAccIdUsingGET(context.Background(), securitiesAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyControllerAPI.GetCompanyBySecAccIdUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyBySecAccIdUsingGET`: CompanyView
	fmt.Fprintf(os.Stdout, "Response from `CompanyControllerAPI.GetCompanyBySecAccIdUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securitiesAccountId** | **string** | Securities account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyBySecAccIdUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyView**](CompanyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyBySecurityIdentifierUsingGET

> CompanyView GetCompanyBySecurityIdentifierUsingGET(ctx, securityIdentifier).Execute()

Returns company by securityIdentifier

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
	resp, r, err := apiClient.CompanyControllerAPI.GetCompanyBySecurityIdentifierUsingGET(context.Background(), securityIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyControllerAPI.GetCompanyBySecurityIdentifierUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyBySecurityIdentifierUsingGET`: CompanyView
	fmt.Fprintf(os.Stdout, "Response from `CompanyControllerAPI.GetCompanyBySecurityIdentifierUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityIdentifier** | **string** | Security Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyBySecurityIdentifierUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyView**](CompanyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyCapsUsingGET

> CompanyCapabilitiesView GetCompanyCapsUsingGET(ctx, companyId).Execute()

Returns company capabilities

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
	resp, r, err := apiClient.CompanyControllerAPI.GetCompanyCapsUsingGET(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyControllerAPI.GetCompanyCapsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCapsUsingGET`: CompanyCapabilitiesView
	fmt.Fprintf(os.Stdout, "Response from `CompanyControllerAPI.GetCompanyCapsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Company id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCapsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyCapabilitiesView**](CompanyCapabilitiesView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyProfileUsingGET

> CompanyProfileView GetCompanyProfileUsingGET(ctx, companyId).Execute()

Returns company profile

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
	resp, r, err := apiClient.CompanyControllerAPI.GetCompanyProfileUsingGET(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyControllerAPI.GetCompanyProfileUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyProfileUsingGET`: CompanyProfileView
	fmt.Fprintf(os.Stdout, "Response from `CompanyControllerAPI.GetCompanyProfileUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Company id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyProfileUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyProfileView**](CompanyProfileView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyUsingGET

> CompanyView GetCompanyUsingGET(ctx, companyId).Execute()

Returns company

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
	resp, r, err := apiClient.CompanyControllerAPI.GetCompanyUsingGET(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyControllerAPI.GetCompanyUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyUsingGET`: CompanyView
	fmt.Fprintf(os.Stdout, "Response from `CompanyControllerAPI.GetCompanyUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Company id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyView**](CompanyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllCompaniesUsingGET

> PageCompanyView ListAllCompaniesUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Lists all companies



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
	resp, r, err := apiClient.CompanyControllerAPI.ListAllCompaniesUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyControllerAPI.ListAllCompaniesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllCompaniesUsingGET`: PageCompanyView
	fmt.Fprintf(os.Stdout, "Response from `CompanyControllerAPI.ListAllCompaniesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllCompaniesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageCompanyView**](PageCompanyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllCompaniesV1UsingGET

> CompanyView ListAllCompaniesV1UsingGET(ctx).Execute()

Lists all companies

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
	resp, r, err := apiClient.CompanyControllerAPI.ListAllCompaniesV1UsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyControllerAPI.ListAllCompaniesV1UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllCompaniesV1UsingGET`: CompanyView
	fmt.Fprintf(os.Stdout, "Response from `CompanyControllerAPI.ListAllCompaniesV1UsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAllCompaniesV1UsingGETRequest struct via the builder pattern


### Return type

[**CompanyView**](CompanyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompaniesByUserIdUsingGET

> CompanyView ListCompaniesByUserIdUsingGET(ctx, userId).Execute()

Lists companies controlled by ceo with user id

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
	userId := "userId_example" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyControllerAPI.ListCompaniesByUserIdUsingGET(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyControllerAPI.ListCompaniesByUserIdUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompaniesByUserIdUsingGET`: CompanyView
	fmt.Fprintf(os.Stdout, "Response from `CompanyControllerAPI.ListCompaniesByUserIdUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompaniesByUserIdUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyView**](CompanyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompaniesByUserNameUsingGET

> CompanyView ListCompaniesByUserNameUsingGET(ctx, username).Execute()

Lists companies controlled by ceo with username

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
	username := "username_example" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyControllerAPI.ListCompaniesByUserNameUsingGET(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyControllerAPI.ListCompaniesByUserNameUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompaniesByUserNameUsingGET`: CompanyView
	fmt.Fprintf(os.Stdout, "Response from `CompanyControllerAPI.ListCompaniesByUserNameUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompaniesByUserNameUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyView**](CompanyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMyCompaniesUsingGET

> PageCompanyView ListMyCompaniesUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Lists companies controlled by logged in user



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
	resp, r, err := apiClient.CompanyControllerAPI.ListMyCompaniesUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyControllerAPI.ListMyCompaniesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyCompaniesUsingGET`: PageCompanyView
	fmt.Fprintf(os.Stdout, "Response from `CompanyControllerAPI.ListMyCompaniesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMyCompaniesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageCompanyView**](PageCompanyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMyCompaniesV1UsingGET

> CompanyView ListMyCompaniesV1UsingGET(ctx).Execute()

Lists companies controlled by logged in user

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
	resp, r, err := apiClient.CompanyControllerAPI.ListMyCompaniesV1UsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyControllerAPI.ListMyCompaniesV1UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyCompaniesV1UsingGET`: CompanyView
	fmt.Fprintf(os.Stdout, "Response from `CompanyControllerAPI.ListMyCompaniesV1UsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMyCompaniesV1UsingGETRequest struct via the builder pattern


### Return type

[**CompanyView**](CompanyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLogoUsingDELETE

> CompanyView RemoveLogoUsingDELETE(ctx, companyId).Execute()

Removes logo

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
	resp, r, err := apiClient.CompanyControllerAPI.RemoveLogoUsingDELETE(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyControllerAPI.RemoveLogoUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveLogoUsingDELETE`: CompanyView
	fmt.Fprintf(os.Stdout, "Response from `CompanyControllerAPI.RemoveLogoUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Company id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLogoUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyView**](CompanyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCompaniesUsingGET

> CompactCompanyView SearchCompaniesUsingGET(ctx, namePart).Execute()

Finds companies by name part

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
	namePart := "namePart_example" // string | Name part

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyControllerAPI.SearchCompaniesUsingGET(context.Background(), namePart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyControllerAPI.SearchCompaniesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCompaniesUsingGET`: CompactCompanyView
	fmt.Fprintf(os.Stdout, "Response from `CompanyControllerAPI.SearchCompaniesUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namePart** | **string** | Name part | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCompaniesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompactCompanyView**](CompactCompanyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

