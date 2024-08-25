# \NotificationControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeNotificationUsingPUT**](NotificationControllerAPI.md#ChangeNotificationUsingPUT) | **Put** /api/v2/notifications/{notificationId} | Changes notification
[**ChangeNotificationsUsingPUT**](NotificationControllerAPI.md#ChangeNotificationsUsingPUT) | **Put** /api/v2/notifications | Changes notifications
[**DeleteNotificationUsingDELETE**](NotificationControllerAPI.md#DeleteNotificationUsingDELETE) | **Delete** /api/v2/notifications/{notificationId} | Deletes notification
[**DeleteNotificationsUsingDELETE**](NotificationControllerAPI.md#DeleteNotificationsUsingDELETE) | **Delete** /api/v2/notifications | Deletes notifications
[**GetMyNotificationsCountDeprecatedUsingGET**](NotificationControllerAPI.md#GetMyNotificationsCountDeprecatedUsingGET) | **Get** /api/v2/notifications/unread/count | Returns number of unread notifications
[**GetMyNotificationsCountUsingGET**](NotificationControllerAPI.md#GetMyNotificationsCountUsingGET) | **Get** /api/v2/notifications/count | Returns number of unread notifications
[**GetMyNotificationsUsingGET**](NotificationControllerAPI.md#GetMyNotificationsUsingGET) | **Get** /api/v2/notifications | Lists notifications for logged in user
[**GetMyNotificationsV1UsingGET**](NotificationControllerAPI.md#GetMyNotificationsV1UsingGET) | **Get** /api/notifications | Lists notifications for logged in user
[**GetMyUnreadNotificationsUsingGET**](NotificationControllerAPI.md#GetMyUnreadNotificationsUsingGET) | **Get** /api/notifications/unread/ | Lists unread notifications for logged in user
[**GetNotificationUsingGET**](NotificationControllerAPI.md#GetNotificationUsingGET) | **Get** /api/v2/notifications/{notificationId} | Returns notification
[**MarkNotificationsAsReadUsingPUT**](NotificationControllerAPI.md#MarkNotificationsAsReadUsingPUT) | **Put** /api/notifications/read/ | Marks all notifications for logged in user as read



## ChangeNotificationUsingPUT

> NotificationView ChangeNotificationUsingPUT(ctx, notificationId).IsRead(isRead).Execute()

Changes notification

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
	isRead := true // bool | Is read by logged in user
	notificationId := "notificationId_example" // string | Notification id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationControllerAPI.ChangeNotificationUsingPUT(context.Background(), notificationId).IsRead(isRead).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationControllerAPI.ChangeNotificationUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeNotificationUsingPUT`: NotificationView
	fmt.Fprintf(os.Stdout, "Response from `NotificationControllerAPI.ChangeNotificationUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | Notification id | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeNotificationUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isRead** | **bool** | Is read by logged in user | 


### Return type

[**NotificationView**](NotificationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeNotificationsUsingPUT

> NotificationView ChangeNotificationsUsingPUT(ctx).IsRead(isRead).NotificationIds(notificationIds).Search(search).Execute()

Changes notifications

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
	isRead := true // bool | Is read by logged in user
	notificationIds := []string{"Inner_example"} // []string | Notification ids (optional)
	search := "search_example" // string | Fulltext search (be aware of non-localized subject and content and be aware of MessagePrototype's structure with substitutions; useful for security identifiers) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationControllerAPI.ChangeNotificationsUsingPUT(context.Background()).IsRead(isRead).NotificationIds(notificationIds).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationControllerAPI.ChangeNotificationsUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeNotificationsUsingPUT`: NotificationView
	fmt.Fprintf(os.Stdout, "Response from `NotificationControllerAPI.ChangeNotificationsUsingPUT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeNotificationsUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isRead** | **bool** | Is read by logged in user | 
 **notificationIds** | **[]string** | Notification ids | 
 **search** | **string** | Fulltext search (be aware of non-localized subject and content and be aware of MessagePrototype&#39;s structure with substitutions; useful for security identifiers) | 

### Return type

[**NotificationView**](NotificationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationUsingDELETE

> MessagePrototype DeleteNotificationUsingDELETE(ctx, notificationId).Execute()

Deletes notification

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
	notificationId := "notificationId_example" // string | Notification id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationControllerAPI.DeleteNotificationUsingDELETE(context.Background(), notificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationControllerAPI.DeleteNotificationUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNotificationUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `NotificationControllerAPI.DeleteNotificationUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | Notification id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationUsingDELETERequest struct via the builder pattern


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


## DeleteNotificationsUsingDELETE

> MessagePrototype DeleteNotificationsUsingDELETE(ctx).NotificationIds(notificationIds).Search(search).Execute()

Deletes notifications

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
	notificationIds := []string{"Inner_example"} // []string | Notification ids (optional)
	search := "search_example" // string | Fulltext search (be aware of non-localized subject and content and be aware of MessagePrototype's structure with substitutions; useful for security identifiers) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationControllerAPI.DeleteNotificationsUsingDELETE(context.Background()).NotificationIds(notificationIds).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationControllerAPI.DeleteNotificationsUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNotificationsUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `NotificationControllerAPI.DeleteNotificationsUsingDELETE`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationsUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationIds** | **[]string** | Notification ids | 
 **search** | **string** | Fulltext search (be aware of non-localized subject and content and be aware of MessagePrototype&#39;s structure with substitutions; useful for security identifiers) | 

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


## GetMyNotificationsCountDeprecatedUsingGET

> int64 GetMyNotificationsCountDeprecatedUsingGET(ctx).Execute()

Returns number of unread notifications

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
	resp, r, err := apiClient.NotificationControllerAPI.GetMyNotificationsCountDeprecatedUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationControllerAPI.GetMyNotificationsCountDeprecatedUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyNotificationsCountDeprecatedUsingGET`: int64
	fmt.Fprintf(os.Stdout, "Response from `NotificationControllerAPI.GetMyNotificationsCountDeprecatedUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyNotificationsCountDeprecatedUsingGETRequest struct via the builder pattern


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


## GetMyNotificationsCountUsingGET

> int64 GetMyNotificationsCountUsingGET(ctx).IsRead(isRead).Search(search).Execute()

Returns number of unread notifications

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
	isRead := true // bool | Is read by logged in user (optional)
	search := "search_example" // string | Fulltext search (be aware of non-localized subject and content and be aware of MessagePrototype's structure with substitutions; useful for security identifiers) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationControllerAPI.GetMyNotificationsCountUsingGET(context.Background()).IsRead(isRead).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationControllerAPI.GetMyNotificationsCountUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyNotificationsCountUsingGET`: int64
	fmt.Fprintf(os.Stdout, "Response from `NotificationControllerAPI.GetMyNotificationsCountUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyNotificationsCountUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isRead** | **bool** | Is read by logged in user | 
 **search** | **string** | Fulltext search (be aware of non-localized subject and content and be aware of MessagePrototype&#39;s structure with substitutions; useful for security identifiers) | 

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


## GetMyNotificationsUsingGET

> NotificationView GetMyNotificationsUsingGET(ctx).IsRead(isRead).Page(page).Search(search).Size(size).Sort(sort).Execute()

Lists notifications for logged in user

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
	isRead := true // bool | Is read by logged in user (optional)
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	search := "search_example" // string | Fulltext search (be aware of non-localized subject and content and be aware of MessagePrototype's structure with substitutions; useful for security identifiers) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationControllerAPI.GetMyNotificationsUsingGET(context.Background()).IsRead(isRead).Page(page).Search(search).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationControllerAPI.GetMyNotificationsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyNotificationsUsingGET`: NotificationView
	fmt.Fprintf(os.Stdout, "Response from `NotificationControllerAPI.GetMyNotificationsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyNotificationsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isRead** | **bool** | Is read by logged in user | 
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **search** | **string** | Fulltext search (be aware of non-localized subject and content and be aware of MessagePrototype&#39;s structure with substitutions; useful for security identifiers) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**NotificationView**](NotificationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyNotificationsV1UsingGET

> NotificationView GetMyNotificationsV1UsingGET(ctx).Execute()

Lists notifications for logged in user

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
	resp, r, err := apiClient.NotificationControllerAPI.GetMyNotificationsV1UsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationControllerAPI.GetMyNotificationsV1UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyNotificationsV1UsingGET`: NotificationView
	fmt.Fprintf(os.Stdout, "Response from `NotificationControllerAPI.GetMyNotificationsV1UsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyNotificationsV1UsingGETRequest struct via the builder pattern


### Return type

[**NotificationView**](NotificationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyUnreadNotificationsUsingGET

> NotificationView GetMyUnreadNotificationsUsingGET(ctx).Execute()

Lists unread notifications for logged in user

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
	resp, r, err := apiClient.NotificationControllerAPI.GetMyUnreadNotificationsUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationControllerAPI.GetMyUnreadNotificationsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyUnreadNotificationsUsingGET`: NotificationView
	fmt.Fprintf(os.Stdout, "Response from `NotificationControllerAPI.GetMyUnreadNotificationsUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyUnreadNotificationsUsingGETRequest struct via the builder pattern


### Return type

[**NotificationView**](NotificationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationUsingGET

> NotificationView GetNotificationUsingGET(ctx, notificationId).Execute()

Returns notification

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
	notificationId := "notificationId_example" // string | Notification id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationControllerAPI.GetNotificationUsingGET(context.Background(), notificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationControllerAPI.GetNotificationUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationUsingGET`: NotificationView
	fmt.Fprintf(os.Stdout, "Response from `NotificationControllerAPI.GetNotificationUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | Notification id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationView**](NotificationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkNotificationsAsReadUsingPUT

> MessagePrototype MarkNotificationsAsReadUsingPUT(ctx).Execute()

Marks all notifications for logged in user as read

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
	resp, r, err := apiClient.NotificationControllerAPI.MarkNotificationsAsReadUsingPUT(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationControllerAPI.MarkNotificationsAsReadUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkNotificationsAsReadUsingPUT`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `NotificationControllerAPI.MarkNotificationsAsReadUsingPUT`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarkNotificationsAsReadUsingPUTRequest struct via the builder pattern


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

