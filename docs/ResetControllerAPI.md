# \ResetControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CleanusersUsingPUT**](ResetControllerAPI.md#CleanusersUsingPUT) | **Put** /api/cleanusers | Cleans users
[**DeleteTestusersUsingDELETE**](ResetControllerAPI.md#DeleteTestusersUsingDELETE) | **Delete** /api/testusers | Deletes test users
[**DeleteUserAdminUsingDELETE**](ResetControllerAPI.md#DeleteUserAdminUsingDELETE) | **Delete** /api/v2/user/{username} | Removes user
[**GenCodesUsingPOST**](ResetControllerAPI.md#GenCodesUsingPOST) | **Post** /api/codes | Generate codes
[**GetJwtTokenUsingGET**](ResetControllerAPI.md#GetJwtTokenUsingGET) | **Get** /api/reset/usertoken | Returns User&#39;s JWT Token
[**RemoveFromChatInBatchUsingPUT**](ResetControllerAPI.md#RemoveFromChatInBatchUsingPUT) | **Put** /api/reset/removeusersfromchat | Removes users from chat, 300 in one run max
[**ResetUsingPUT**](ResetControllerAPI.md#ResetUsingPUT) | **Put** /api/reset/ | Resets the market
[**RetryPaddleEventUsingPOST**](ResetControllerAPI.md#RetryPaddleEventUsingPOST) | **Post** /api/v2/paddleretry | Retries paddle event
[**SendMailUsingPOST**](ResetControllerAPI.md#SendMailUsingPOST) | **Post** /api/mail | Sends mail
[**SendTestMailUsingPOST**](ResetControllerAPI.md#SendTestMailUsingPOST) | **Post** /api/testmail | Sends test mail
[**TestmailUsingPUT**](ResetControllerAPI.md#TestmailUsingPUT) | **Put** /api/testmail | Sends test mail
[**UnsubscribeEmailUsingPUT**](ResetControllerAPI.md#UnsubscribeEmailUsingPUT) | **Put** /api/unsubscribe | Unsubscribe email



## CleanusersUsingPUT

> MessagePrototype CleanusersUsingPUT(ctx).Execute()

Cleans users

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
	resp, r, err := apiClient.ResetControllerAPI.CleanusersUsingPUT(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResetControllerAPI.CleanusersUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CleanusersUsingPUT`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `ResetControllerAPI.CleanusersUsingPUT`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCleanusersUsingPUTRequest struct via the builder pattern


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


## DeleteTestusersUsingDELETE

> MessagePrototype DeleteTestusersUsingDELETE(ctx).NumberOfUsers(numberOfUsers).Execute()

Deletes test users

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
	numberOfUsers := int32(56) // int32 | numberOfUsers

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResetControllerAPI.DeleteTestusersUsingDELETE(context.Background()).NumberOfUsers(numberOfUsers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResetControllerAPI.DeleteTestusersUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTestusersUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `ResetControllerAPI.DeleteTestusersUsingDELETE`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTestusersUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **numberOfUsers** | **int32** | numberOfUsers | 

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


## DeleteUserAdminUsingDELETE

> MessagePrototype DeleteUserAdminUsingDELETE(ctx, username).Execute()

Removes user

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
	resp, r, err := apiClient.ResetControllerAPI.DeleteUserAdminUsingDELETE(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResetControllerAPI.DeleteUserAdminUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserAdminUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `ResetControllerAPI.DeleteUserAdminUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserAdminUsingDELETERequest struct via the builder pattern


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


## GenCodesUsingPOST

> MessagePrototype GenCodesUsingPOST(ctx).Days(days).Number(number).Execute()

Generate codes

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
	days := int32(56) // int32 | days
	number := int32(56) // int32 | number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResetControllerAPI.GenCodesUsingPOST(context.Background()).Days(days).Number(number).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResetControllerAPI.GenCodesUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenCodesUsingPOST`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `ResetControllerAPI.GenCodesUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenCodesUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **int32** | days | 
 **number** | **int32** | number | 

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


## GetJwtTokenUsingGET

> MessagePrototype GetJwtTokenUsingGET(ctx).Username(username).Execute()

Returns User's JWT Token

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
	resp, r, err := apiClient.ResetControllerAPI.GetJwtTokenUsingGET(context.Background()).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResetControllerAPI.GetJwtTokenUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJwtTokenUsingGET`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `ResetControllerAPI.GetJwtTokenUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJwtTokenUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** | Username | 

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


## RemoveFromChatInBatchUsingPUT

> ChatView RemoveFromChatInBatchUsingPUT(ctx).ChatId(chatId).Namepart(namepart).Execute()

Removes users from chat, 300 in one run max

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
	chatId := "chatId_example" // string | Chat id
	namepart := "namepart_example" // string | Namepart

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResetControllerAPI.RemoveFromChatInBatchUsingPUT(context.Background()).ChatId(chatId).Namepart(namepart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResetControllerAPI.RemoveFromChatInBatchUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveFromChatInBatchUsingPUT`: ChatView
	fmt.Fprintf(os.Stdout, "Response from `ResetControllerAPI.RemoveFromChatInBatchUsingPUT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFromChatInBatchUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatId** | **string** | Chat id | 
 **namepart** | **string** | Namepart | 

### Return type

[**ChatView**](ChatView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetUsingPUT

> MessagePrototype ResetUsingPUT(ctx).Execute()

Resets the market

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
	resp, r, err := apiClient.ResetControllerAPI.ResetUsingPUT(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResetControllerAPI.ResetUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetUsingPUT`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `ResetControllerAPI.ResetUsingPUT`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResetUsingPUTRequest struct via the builder pattern


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


## RetryPaddleEventUsingPOST

> MessagePrototype RetryPaddleEventUsingPOST(ctx).RawContent(rawContent).Execute()

Retries paddle event

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
	rawContent := "rawContent_example" // string | Raw content

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResetControllerAPI.RetryPaddleEventUsingPOST(context.Background()).RawContent(rawContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResetControllerAPI.RetryPaddleEventUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryPaddleEventUsingPOST`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `ResetControllerAPI.RetryPaddleEventUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetryPaddleEventUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rawContent** | **string** | Raw content | 

### Return type

[**MessagePrototype**](MessagePrototype.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMailUsingPOST

> MessagePrototype SendMailUsingPOST(ctx).Locale(locale).Subject(subject).HtmlContent(htmlContent).Execute()

Sends mail

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
	locale := "locale_example" // string | locale
	subject := "subject_example" // string | subject
	htmlContent := "htmlContent_example" // string | htmlContent

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResetControllerAPI.SendMailUsingPOST(context.Background()).Locale(locale).Subject(subject).HtmlContent(htmlContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResetControllerAPI.SendMailUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendMailUsingPOST`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `ResetControllerAPI.SendMailUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendMailUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locale** | **string** | locale | 
 **subject** | **string** | subject | 
 **htmlContent** | **string** | htmlContent | 

### Return type

[**MessagePrototype**](MessagePrototype.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTestMailUsingPOST

> MessagePrototype SendTestMailUsingPOST(ctx).EmailAddress(emailAddress).Subject(subject).HtmlContent(htmlContent).Execute()

Sends test mail

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
	emailAddress := "emailAddress_example" // string | emailAddress
	subject := "subject_example" // string | subject
	htmlContent := "htmlContent_example" // string | htmlContent

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResetControllerAPI.SendTestMailUsingPOST(context.Background()).EmailAddress(emailAddress).Subject(subject).HtmlContent(htmlContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResetControllerAPI.SendTestMailUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendTestMailUsingPOST`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `ResetControllerAPI.SendTestMailUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendTestMailUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailAddress** | **string** | emailAddress | 
 **subject** | **string** | subject | 
 **htmlContent** | **string** | htmlContent | 

### Return type

[**MessagePrototype**](MessagePrototype.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestmailUsingPUT

> MessagePrototype TestmailUsingPUT(ctx).Execute()

Sends test mail

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
	resp, r, err := apiClient.ResetControllerAPI.TestmailUsingPUT(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResetControllerAPI.TestmailUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestmailUsingPUT`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `ResetControllerAPI.TestmailUsingPUT`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestmailUsingPUTRequest struct via the builder pattern


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


## UnsubscribeEmailUsingPUT

> MessagePrototype UnsubscribeEmailUsingPUT(ctx).EmailAdress(emailAdress).Execute()

Unsubscribe email

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
	emailAdress := "emailAdress_example" // string | emailAddress

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResetControllerAPI.UnsubscribeEmailUsingPUT(context.Background()).EmailAdress(emailAdress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResetControllerAPI.UnsubscribeEmailUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsubscribeEmailUsingPUT`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `ResetControllerAPI.UnsubscribeEmailUsingPUT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeEmailUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailAdress** | **string** | emailAddress | 

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

