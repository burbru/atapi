# \InterestControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthorInterestUsingGET**](InterestControllerAPI.md#GetAuthorInterestUsingGET) | **Get** /api/v2/my/interests/authors/{userId} | Returns personal interest for author
[**GetCompanyInterestUsingGET**](InterestControllerAPI.md#GetCompanyInterestUsingGET) | **Get** /api/v2/my/interests/companies/{companyId} | Returns personal interest for company
[**GetHashTagInterestUsingGET**](InterestControllerAPI.md#GetHashTagInterestUsingGET) | **Get** /api/v2/my/interests/hashtags/{hashTag} | Returns personal interest for hash tag
[**PutAuthorInterestUsingPUT**](InterestControllerAPI.md#PutAuthorInterestUsingPUT) | **Put** /api/v2/my/interests/authors/{userId} | Sets personal interest for author
[**PutCompanyInterestUsingPUT**](InterestControllerAPI.md#PutCompanyInterestUsingPUT) | **Put** /api/v2/my/interests/companies/{companyId} | Sets personal interest for company
[**PutHashTagInterestUsingPUT**](InterestControllerAPI.md#PutHashTagInterestUsingPUT) | **Put** /api/v2/my/interests/hashtags/{hashTag} | Sets personal interest for hashtag



## GetAuthorInterestUsingGET

> AuthorInterest GetAuthorInterestUsingGET(ctx, userId).Execute()

Returns personal interest for author

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
	userId := "userId_example" // string | userId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InterestControllerAPI.GetAuthorInterestUsingGET(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterestControllerAPI.GetAuthorInterestUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthorInterestUsingGET`: AuthorInterest
	fmt.Fprintf(os.Stdout, "Response from `InterestControllerAPI.GetAuthorInterestUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | userId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorInterestUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthorInterest**](AuthorInterest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyInterestUsingGET

> CompanyInterest GetCompanyInterestUsingGET(ctx, companyId).Execute()

Returns personal interest for company

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
	companyId := "companyId_example" // string | companyId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InterestControllerAPI.GetCompanyInterestUsingGET(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterestControllerAPI.GetCompanyInterestUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyInterestUsingGET`: CompanyInterest
	fmt.Fprintf(os.Stdout, "Response from `InterestControllerAPI.GetCompanyInterestUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyInterestUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyInterest**](CompanyInterest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHashTagInterestUsingGET

> HashTagInterest GetHashTagInterestUsingGET(ctx, hashTag).Execute()

Returns personal interest for hash tag

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
	hashTag := "test" // string | hashTag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InterestControllerAPI.GetHashTagInterestUsingGET(context.Background(), hashTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterestControllerAPI.GetHashTagInterestUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHashTagInterestUsingGET`: HashTagInterest
	fmt.Fprintf(os.Stdout, "Response from `InterestControllerAPI.GetHashTagInterestUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hashTag** | **string** | hashTag | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHashTagInterestUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HashTagInterest**](HashTagInterest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAuthorInterestUsingPUT

> AuthorInterest PutAuthorInterestUsingPUT(ctx, userId).Interest(interest).Execute()

Sets personal interest for author

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
	interest := "23" // string | Interest
	userId := "userId_example" // string | userId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InterestControllerAPI.PutAuthorInterestUsingPUT(context.Background(), userId).Interest(interest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterestControllerAPI.PutAuthorInterestUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAuthorInterestUsingPUT`: AuthorInterest
	fmt.Fprintf(os.Stdout, "Response from `InterestControllerAPI.PutAuthorInterestUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | userId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAuthorInterestUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interest** | **string** | Interest | 


### Return type

[**AuthorInterest**](AuthorInterest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyInterestUsingPUT

> CompanyInterest PutCompanyInterestUsingPUT(ctx, companyId).Interest(interest).Execute()

Sets personal interest for company

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
	companyId := "companyId_example" // string | companyId
	interest := "23" // string | Interest

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InterestControllerAPI.PutCompanyInterestUsingPUT(context.Background(), companyId).Interest(interest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterestControllerAPI.PutCompanyInterestUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyInterestUsingPUT`: CompanyInterest
	fmt.Fprintf(os.Stdout, "Response from `InterestControllerAPI.PutCompanyInterestUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyInterestUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **interest** | **string** | Interest | 

### Return type

[**CompanyInterest**](CompanyInterest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutHashTagInterestUsingPUT

> HashTagInterest PutHashTagInterestUsingPUT(ctx, hashTag).Interest(interest).Execute()

Sets personal interest for hashtag

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
	hashTag := "test" // string | hashTag
	interest := "23" // string | Interest

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InterestControllerAPI.PutHashTagInterestUsingPUT(context.Background(), hashTag).Interest(interest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterestControllerAPI.PutHashTagInterestUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutHashTagInterestUsingPUT`: HashTagInterest
	fmt.Fprintf(os.Stdout, "Response from `InterestControllerAPI.PutHashTagInterestUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hashTag** | **string** | hashTag | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutHashTagInterestUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **interest** | **string** | Interest | 

### Return type

[**HashTagInterest**](HashTagInterest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

