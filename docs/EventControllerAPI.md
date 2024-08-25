# \EventControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindEventsByEventTypeUsingGET**](EventControllerAPI.md#FindEventsByEventTypeUsingGET) | **Get** /api/search/events/type/{eventType} | Finds nonpersistent events by event type  (optional: after date, realms)
[**FindEventsByFullTextUsingGET**](EventControllerAPI.md#FindEventsByFullTextUsingGET) | **Get** /api/search/events/{fullTextPart} | Finds nonpersistent events by full-text search (optional: after date)
[**FindEventsByRealmsUsingGET**](EventControllerAPI.md#FindEventsByRealmsUsingGET) | **Get** /api/search/events/ | Finds nonpersistent events by realms (optional: after date)
[**GetAllUsingGET**](EventControllerAPI.md#GetAllUsingGET) | **Get** /api/events/ | Returns all nonpersistent events (20000 max, optional: after date)
[**GetMyEventsUsingGET**](EventControllerAPI.md#GetMyEventsUsingGET) | **Get** /api/events/user/ | Returns all nonpersistent events fore current user (20000 max, optional: after date)



## FindEventsByEventTypeUsingGET

> Event FindEventsByEventTypeUsingGET(ctx, eventType).AfterDate(afterDate).Realms(realms).Execute()

Finds nonpersistent events by event type  (optional: after date, realms)

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
	eventType := "eventType_example" // string | Event type
	afterDate := int64(789) // int64 | After date (optional)
	realms := []string{"Inner_example"} // []string | Realms (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventControllerAPI.FindEventsByEventTypeUsingGET(context.Background(), eventType).AfterDate(afterDate).Realms(realms).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventControllerAPI.FindEventsByEventTypeUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindEventsByEventTypeUsingGET`: Event
	fmt.Fprintf(os.Stdout, "Response from `EventControllerAPI.FindEventsByEventTypeUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventType** | **string** | Event type | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindEventsByEventTypeUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afterDate** | **int64** | After date | 
 **realms** | **[]string** | Realms | 

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindEventsByFullTextUsingGET

> Event FindEventsByFullTextUsingGET(ctx, fullTextPart).AfterDate(afterDate).Execute()

Finds nonpersistent events by full-text search (optional: after date)

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
	fullTextPart := "fullTextPart_example" // string | Part of the event full-text representation
	afterDate := int64(789) // int64 | After date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventControllerAPI.FindEventsByFullTextUsingGET(context.Background(), fullTextPart).AfterDate(afterDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventControllerAPI.FindEventsByFullTextUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindEventsByFullTextUsingGET`: Event
	fmt.Fprintf(os.Stdout, "Response from `EventControllerAPI.FindEventsByFullTextUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fullTextPart** | **string** | Part of the event full-text representation | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindEventsByFullTextUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afterDate** | **int64** | After date | 

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindEventsByRealmsUsingGET

> Event FindEventsByRealmsUsingGET(ctx).Realms(realms).AfterDate(afterDate).Execute()

Finds nonpersistent events by realms (optional: after date)

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
	realms := []string{"Inner_example"} // []string | Realms
	afterDate := int64(789) // int64 | After date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventControllerAPI.FindEventsByRealmsUsingGET(context.Background()).Realms(realms).AfterDate(afterDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventControllerAPI.FindEventsByRealmsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindEventsByRealmsUsingGET`: Event
	fmt.Fprintf(os.Stdout, "Response from `EventControllerAPI.FindEventsByRealmsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindEventsByRealmsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **realms** | **[]string** | Realms | 
 **afterDate** | **int64** | After date | 

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllUsingGET

> Event GetAllUsingGET(ctx).AfterDate(afterDate).Execute()

Returns all nonpersistent events (20000 max, optional: after date)

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
	afterDate := int64(789) // int64 | After date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventControllerAPI.GetAllUsingGET(context.Background()).AfterDate(afterDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventControllerAPI.GetAllUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllUsingGET`: Event
	fmt.Fprintf(os.Stdout, "Response from `EventControllerAPI.GetAllUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afterDate** | **int64** | After date | 

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyEventsUsingGET

> Event GetMyEventsUsingGET(ctx).AfterDate(afterDate).Execute()

Returns all nonpersistent events fore current user (20000 max, optional: after date)

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
	afterDate := int64(789) // int64 | After date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventControllerAPI.GetMyEventsUsingGET(context.Background()).AfterDate(afterDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventControllerAPI.GetMyEventsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyEventsUsingGET`: Event
	fmt.Fprintf(os.Stdout, "Response from `EventControllerAPI.GetMyEventsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyEventsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afterDate** | **int64** | After date | 

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

