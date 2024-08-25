# \SubscriptionControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscribeToAuthorUsingPUT**](SubscriptionControllerAPI.md#SubscribeToAuthorUsingPUT) | **Put** /api/v2/my/subscriptions/authors/{userId} | Subscribes to/Ignores news from an author
[**SubscribeToCompanyUsingPUT**](SubscriptionControllerAPI.md#SubscribeToCompanyUsingPUT) | **Put** /api/v2/my/subscriptions/companies/{companyId} | Subscribes to/Ignores news from a company
[**SubscribeToHashtagUsingPUT**](SubscriptionControllerAPI.md#SubscribeToHashtagUsingPUT) | **Put** /api/v2/my/subscriptions/hashtags/{hashtag} | Subscribes to/Ignores news from a hashtag
[**UnsubscribeFromAuthorUsingDELETE**](SubscriptionControllerAPI.md#UnsubscribeFromAuthorUsingDELETE) | **Delete** /api/v2/my/subscriptions/authors/{userId} | Unsubscribes news from an author
[**UnsubscribeFromCompanyUsingDELETE**](SubscriptionControllerAPI.md#UnsubscribeFromCompanyUsingDELETE) | **Delete** /api/v2/my/subscriptions/companies/{companyId} | Unsubscribes news from a company
[**UnsubscribeFromHashtagUsingDELETE**](SubscriptionControllerAPI.md#UnsubscribeFromHashtagUsingDELETE) | **Delete** /api/v2/my/subscriptions/hashtags/{hashtag} | Unsubscribes news from a hashtag



## SubscribeToAuthorUsingPUT

> AuthorInterest SubscribeToAuthorUsingPUT(ctx, userId).Action(action).Execute()

Subscribes to/Ignores news from an author

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
	action := "action_example" // string | Action
	userId := "userId_example" // string | userId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionControllerAPI.SubscribeToAuthorUsingPUT(context.Background(), userId).Action(action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionControllerAPI.SubscribeToAuthorUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscribeToAuthorUsingPUT`: AuthorInterest
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionControllerAPI.SubscribeToAuthorUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | userId | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeToAuthorUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | Action | 


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


## SubscribeToCompanyUsingPUT

> CompanyInterest SubscribeToCompanyUsingPUT(ctx, companyId).Action(action).Execute()

Subscribes to/Ignores news from a company

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
	action := "action_example" // string | Action
	companyId := "companyId_example" // string | companyId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionControllerAPI.SubscribeToCompanyUsingPUT(context.Background(), companyId).Action(action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionControllerAPI.SubscribeToCompanyUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscribeToCompanyUsingPUT`: CompanyInterest
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionControllerAPI.SubscribeToCompanyUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeToCompanyUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | Action | 


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


## SubscribeToHashtagUsingPUT

> HashTagInterest SubscribeToHashtagUsingPUT(ctx, hashtag).Action(action).Execute()

Subscribes to/Ignores news from a hashtag

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
	action := "action_example" // string | Action
	hashtag := "hashtag_example" // string | hashtag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionControllerAPI.SubscribeToHashtagUsingPUT(context.Background(), hashtag).Action(action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionControllerAPI.SubscribeToHashtagUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscribeToHashtagUsingPUT`: HashTagInterest
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionControllerAPI.SubscribeToHashtagUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hashtag** | **string** | hashtag | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeToHashtagUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | Action | 


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


## UnsubscribeFromAuthorUsingDELETE

> MessagePrototype UnsubscribeFromAuthorUsingDELETE(ctx, userId).Execute()

Unsubscribes news from an author

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
	resp, r, err := apiClient.SubscriptionControllerAPI.UnsubscribeFromAuthorUsingDELETE(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionControllerAPI.UnsubscribeFromAuthorUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsubscribeFromAuthorUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionControllerAPI.UnsubscribeFromAuthorUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | userId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeFromAuthorUsingDELETERequest struct via the builder pattern


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


## UnsubscribeFromCompanyUsingDELETE

> MessagePrototype UnsubscribeFromCompanyUsingDELETE(ctx, companyId).Execute()

Unsubscribes news from a company

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
	resp, r, err := apiClient.SubscriptionControllerAPI.UnsubscribeFromCompanyUsingDELETE(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionControllerAPI.UnsubscribeFromCompanyUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsubscribeFromCompanyUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionControllerAPI.UnsubscribeFromCompanyUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeFromCompanyUsingDELETERequest struct via the builder pattern


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


## UnsubscribeFromHashtagUsingDELETE

> MessagePrototype UnsubscribeFromHashtagUsingDELETE(ctx, hashtag).Execute()

Unsubscribes news from a hashtag

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
	hashtag := "hashtag_example" // string | hashtag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionControllerAPI.UnsubscribeFromHashtagUsingDELETE(context.Background(), hashtag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionControllerAPI.UnsubscribeFromHashtagUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsubscribeFromHashtagUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionControllerAPI.UnsubscribeFromHashtagUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hashtag** | **string** | hashtag | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeFromHashtagUsingDELETERequest struct via the builder pattern


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

