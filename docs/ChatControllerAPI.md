# \ChatControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddToChatByUsernameUsingPUT**](ChatControllerAPI.md#AddToChatByUsernameUsingPUT) | **Put** /api/chats/adduser/username | Adds user to chat
[**AddToChatUsingPOST**](ChatControllerAPI.md#AddToChatUsingPOST) | **Post** /api/v2/chatmemberships | Adds user to chat
[**AddToChatV1UsingPUT**](ChatControllerAPI.md#AddToChatV1UsingPUT) | **Put** /api/chats/adduser/userid | Adds user to chat
[**ChangeChatMembershipUsingPUT**](ChatControllerAPI.md#ChangeChatMembershipUsingPUT) | **Put** /api/v2/chatmemberships/{membershipId} | Change chat membership
[**ChangeChatUsingPUT**](ChatControllerAPI.md#ChangeChatUsingPUT) | **Put** /api/v2/chats/{chatId} | Changes chat properties
[**ChangeChatV1UsingPUT**](ChatControllerAPI.md#ChangeChatV1UsingPUT) | **Put** /api/chats/{chatId} | Changes chat properties
[**CreateChatByIdsUsingPOST**](ChatControllerAPI.md#CreateChatByIdsUsingPOST) | **Post** /api/chats/userids | Creates a new chat
[**CreateChatByNamesUsingPOST**](ChatControllerAPI.md#CreateChatByNamesUsingPOST) | **Post** /api/chats/usernames | Creates a new chat
[**CreateChatUsingPOST**](ChatControllerAPI.md#CreateChatUsingPOST) | **Post** /api/v2/chats | Creates a new chat
[**DeleteChatUsingDELETE**](ChatControllerAPI.md#DeleteChatUsingDELETE) | **Delete** /api/v2/chats/{chatId} | Deletes chat
[**DeleteChatV1UsingDELETE**](ChatControllerAPI.md#DeleteChatV1UsingDELETE) | **Delete** /api/chats/{chatId} | Deletes chat
[**GetChatMembershipUsingGET**](ChatControllerAPI.md#GetChatMembershipUsingGET) | **Get** /api/v2/chatmemberships/{membershipId} | Returns chat membership
[**GetChatMembershipsUsingGET**](ChatControllerAPI.md#GetChatMembershipsUsingGET) | **Get** /api/v2/chatmemberships | Returns chat memberships of chat
[**GetChatUsingGET**](ChatControllerAPI.md#GetChatUsingGET) | **Get** /api/v2/chats/{chatId} | Returns chat
[**GetChatV1UsingGET**](ChatControllerAPI.md#GetChatV1UsingGET) | **Get** /api/chats/{chatId} | Returns chat
[**GetChatsUsingGET**](ChatControllerAPI.md#GetChatsUsingGET) | **Get** /api/v2/my/chats/unread | Returns chats with unread messages
[**GetChatsV1UsingGET**](ChatControllerAPI.md#GetChatsV1UsingGET) | **Get** /api/chats/unread | Returns chats with unread messages
[**GetMyChatsUsingGET**](ChatControllerAPI.md#GetMyChatsUsingGET) | **Get** /api/v2/my/chats | Lists chats in which logged in user participates
[**GetMyChatsV1UsingGET**](ChatControllerAPI.md#GetMyChatsV1UsingGET) | **Get** /api/chats | Lists chats in which logged in user participates
[**GetUnreadChatsCountUsingGET**](ChatControllerAPI.md#GetUnreadChatsCountUsingGET) | **Get** /api/v2/my/chats/unread/count | Returns number of chats with unread messages
[**ListPublicChatsUsingGET**](ChatControllerAPI.md#ListPublicChatsUsingGET) | **Get** /api/chatrooms | Lists public chatrooms
[**MarkChatAsReadUsingPUT**](ChatControllerAPI.md#MarkChatAsReadUsingPUT) | **Put** /api/v2/my/chats/read | Sets a chat as read by logged in user
[**MarkChatAsReadV1UsingPUT**](ChatControllerAPI.md#MarkChatAsReadV1UsingPUT) | **Put** /api/chats/read | Sets a chat as read by logged in user
[**QuitChatUsingPUT**](ChatControllerAPI.md#QuitChatUsingPUT) | **Put** /api/chats/quitchat/{chatId} | Removes logged in user from chat
[**RemoveChatMembershipUsingDELETE**](ChatControllerAPI.md#RemoveChatMembershipUsingDELETE) | **Delete** /api/v2/chatmemberships/{membershipId} | Removes chat membership
[**RemoveFromChatUsingPUT**](ChatControllerAPI.md#RemoveFromChatUsingPUT) | **Put** /api/chats/removeuser | Removes user from chat
[**RemoveMyChatMembershipUsingDELETE**](ChatControllerAPI.md#RemoveMyChatMembershipUsingDELETE) | **Delete** /api/v2/my/chatmemberships/{chatId} | Removes logged in user&#39;s chat membership



## AddToChatByUsernameUsingPUT

> ChatView AddToChatByUsernameUsingPUT(ctx).ChatId(chatId).Username(username).Execute()

Adds user to chat

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
	username := "username_example" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.AddToChatByUsernameUsingPUT(context.Background()).ChatId(chatId).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.AddToChatByUsernameUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddToChatByUsernameUsingPUT`: ChatView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.AddToChatByUsernameUsingPUT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddToChatByUsernameUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatId** | **string** | Chat id | 
 **username** | **string** | Username | 

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


## AddToChatUsingPOST

> ChatMembershipView AddToChatUsingPOST(ctx).ChatId(chatId).UserId(userId).Role(role).Execute()

Adds user to chat

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
	userId := "userId_example" // string | User id
	role := "role_example" // string | Role (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.AddToChatUsingPOST(context.Background()).ChatId(chatId).UserId(userId).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.AddToChatUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddToChatUsingPOST`: ChatMembershipView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.AddToChatUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddToChatUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatId** | **string** | Chat id | 
 **userId** | **string** | User id | 
 **role** | **string** | Role | 

### Return type

[**ChatMembershipView**](ChatMembershipView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddToChatV1UsingPUT

> ChatView AddToChatV1UsingPUT(ctx).ChatId(chatId).UserId(userId).Execute()

Adds user to chat

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
	userId := "userId_example" // string | User id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.AddToChatV1UsingPUT(context.Background()).ChatId(chatId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.AddToChatV1UsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddToChatV1UsingPUT`: ChatView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.AddToChatV1UsingPUT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddToChatV1UsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatId** | **string** | Chat id | 
 **userId** | **string** | User id | 

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


## ChangeChatMembershipUsingPUT

> ChatMembershipView ChangeChatMembershipUsingPUT(ctx, membershipId).Role(role).Execute()

Change chat membership

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
	membershipId := "membershipId_example" // string | Membership id
	role := "role_example" // string | Role (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.ChangeChatMembershipUsingPUT(context.Background(), membershipId).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.ChangeChatMembershipUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeChatMembershipUsingPUT`: ChatMembershipView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.ChangeChatMembershipUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**membershipId** | **string** | Membership id | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeChatMembershipUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | **string** | Role | 

### Return type

[**ChatMembershipView**](ChatMembershipView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeChatUsingPUT

> PlainChatView ChangeChatUsingPUT(ctx, chatId).ChatName(chatName).Public(public).Readonly(readonly).Execute()

Changes chat properties

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
	chatName := "chatName_example" // string | Chat name (optional)
	public := true // bool | Public (optional)
	readonly := true // bool | Readonly (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.ChangeChatUsingPUT(context.Background(), chatId).ChatName(chatName).Public(public).Readonly(readonly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.ChangeChatUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeChatUsingPUT`: PlainChatView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.ChangeChatUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | Chat id | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeChatUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatName** | **string** | Chat name | 
 **public** | **bool** | Public | 
 **readonly** | **bool** | Readonly | 

### Return type

[**PlainChatView**](PlainChatView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeChatV1UsingPUT

> ChatView ChangeChatV1UsingPUT(ctx, chatId).ChatName(chatName).Readonly(readonly).Execute()

Changes chat properties

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
	chatName := "chatName_example" // string | Chat name (optional)
	readonly := true // bool | Readonly (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.ChangeChatV1UsingPUT(context.Background(), chatId).ChatName(chatName).Readonly(readonly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.ChangeChatV1UsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeChatV1UsingPUT`: ChatView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.ChangeChatV1UsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | Chat id | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeChatV1UsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatName** | **string** | Chat name | 
 **readonly** | **bool** | Readonly | 

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


## CreateChatByIdsUsingPOST

> ChatView CreateChatByIdsUsingPOST(ctx).Participants(participants).ChatName(chatName).Readonly(readonly).Execute()

Creates a new chat

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
	participants := []string{"Inner_example"} // []string | Participants user ids
	chatName := "chatName_example" // string | Chat name (optional)
	readonly := true // bool | Readonly (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.CreateChatByIdsUsingPOST(context.Background()).Participants(participants).ChatName(chatName).Readonly(readonly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.CreateChatByIdsUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChatByIdsUsingPOST`: ChatView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.CreateChatByIdsUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateChatByIdsUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **participants** | **[]string** | Participants user ids | 
 **chatName** | **string** | Chat name | 
 **readonly** | **bool** | Readonly | 

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


## CreateChatByNamesUsingPOST

> ChatView CreateChatByNamesUsingPOST(ctx).Participants(participants).ChatName(chatName).Readonly(readonly).Execute()

Creates a new chat

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
	participants := []string{"Inner_example"} // []string | Participants usernames
	chatName := "chatName_example" // string | Chat name (optional)
	readonly := true // bool | Readonly (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.CreateChatByNamesUsingPOST(context.Background()).Participants(participants).ChatName(chatName).Readonly(readonly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.CreateChatByNamesUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChatByNamesUsingPOST`: ChatView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.CreateChatByNamesUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateChatByNamesUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **participants** | **[]string** | Participants usernames | 
 **chatName** | **string** | Chat name | 
 **readonly** | **bool** | Readonly | 

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


## CreateChatUsingPOST

> ChatView CreateChatUsingPOST(ctx).ChatName(chatName).Public(public).Readonly(readonly).Execute()

Creates a new chat

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
	chatName := "chatName_example" // string | Chat name (optional)
	public := true // bool | Public (optional)
	readonly := true // bool | Readonly (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.CreateChatUsingPOST(context.Background()).ChatName(chatName).Public(public).Readonly(readonly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.CreateChatUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChatUsingPOST`: ChatView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.CreateChatUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateChatUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatName** | **string** | Chat name | 
 **public** | **bool** | Public | 
 **readonly** | **bool** | Readonly | 

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


## DeleteChatUsingDELETE

> MessagePrototype DeleteChatUsingDELETE(ctx, chatId).Execute()

Deletes chat

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.DeleteChatUsingDELETE(context.Background(), chatId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.DeleteChatUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteChatUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.DeleteChatUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | Chat id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChatUsingDELETERequest struct via the builder pattern


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


## DeleteChatV1UsingDELETE

> map[string]interface{} DeleteChatV1UsingDELETE(ctx, chatId).Execute()

Deletes chat

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.DeleteChatV1UsingDELETE(context.Background(), chatId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.DeleteChatV1UsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteChatV1UsingDELETE`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.DeleteChatV1UsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | Chat id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChatV1UsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChatMembershipUsingGET

> ChatMembershipView GetChatMembershipUsingGET(ctx, membershipId).Execute()

Returns chat membership

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
	membershipId := "membershipId_example" // string | Membership id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.GetChatMembershipUsingGET(context.Background(), membershipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.GetChatMembershipUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatMembershipUsingGET`: ChatMembershipView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.GetChatMembershipUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**membershipId** | **string** | Membership id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatMembershipUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatMembershipView**](ChatMembershipView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChatMembershipsUsingGET

> ChatMembershipView GetChatMembershipsUsingGET(ctx).ChatId(chatId).Execute()

Returns chat memberships of chat

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.GetChatMembershipsUsingGET(context.Background()).ChatId(chatId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.GetChatMembershipsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatMembershipsUsingGET`: ChatMembershipView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.GetChatMembershipsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChatMembershipsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatId** | **string** | Chat id | 

### Return type

[**ChatMembershipView**](ChatMembershipView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChatUsingGET

> PlainChatView GetChatUsingGET(ctx, chatId).Execute()

Returns chat

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.GetChatUsingGET(context.Background(), chatId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.GetChatUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatUsingGET`: PlainChatView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.GetChatUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | Chat id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlainChatView**](PlainChatView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChatV1UsingGET

> ChatView GetChatV1UsingGET(ctx, chatId).Execute()

Returns chat

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.GetChatV1UsingGET(context.Background(), chatId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.GetChatV1UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatV1UsingGET`: ChatView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.GetChatV1UsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | Chat id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatV1UsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetChatsUsingGET

> CompactChatView GetChatsUsingGET(ctx).Execute()

Returns chats with unread messages

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
	resp, r, err := apiClient.ChatControllerAPI.GetChatsUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.GetChatsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatsUsingGET`: CompactChatView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.GetChatsUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatsUsingGETRequest struct via the builder pattern


### Return type

[**CompactChatView**](CompactChatView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChatsV1UsingGET

> CompactChatView GetChatsV1UsingGET(ctx).Execute()

Returns chats with unread messages

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
	resp, r, err := apiClient.ChatControllerAPI.GetChatsV1UsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.GetChatsV1UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatsV1UsingGET`: CompactChatView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.GetChatsV1UsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatsV1UsingGETRequest struct via the builder pattern


### Return type

[**CompactChatView**](CompactChatView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyChatsUsingGET

> ChatView GetMyChatsUsingGET(ctx).Execute()

Lists chats in which logged in user participates

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
	resp, r, err := apiClient.ChatControllerAPI.GetMyChatsUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.GetMyChatsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyChatsUsingGET`: ChatView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.GetMyChatsUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyChatsUsingGETRequest struct via the builder pattern


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


## GetMyChatsV1UsingGET

> ChatView GetMyChatsV1UsingGET(ctx).Execute()

Lists chats in which logged in user participates

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
	resp, r, err := apiClient.ChatControllerAPI.GetMyChatsV1UsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.GetMyChatsV1UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyChatsV1UsingGET`: ChatView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.GetMyChatsV1UsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyChatsV1UsingGETRequest struct via the builder pattern


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


## GetUnreadChatsCountUsingGET

> int64 GetUnreadChatsCountUsingGET(ctx).Execute()

Returns number of chats with unread messages

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
	resp, r, err := apiClient.ChatControllerAPI.GetUnreadChatsCountUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.GetUnreadChatsCountUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUnreadChatsCountUsingGET`: int64
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.GetUnreadChatsCountUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnreadChatsCountUsingGETRequest struct via the builder pattern


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


## ListPublicChatsUsingGET

> ChatRoomView ListPublicChatsUsingGET(ctx).MemberId(memberId).Execute()

Lists public chatrooms

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
	memberId := "memberId_example" // string | User ID of chat member (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.ListPublicChatsUsingGET(context.Background()).MemberId(memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.ListPublicChatsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPublicChatsUsingGET`: ChatRoomView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.ListPublicChatsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPublicChatsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **memberId** | **string** | User ID of chat member | 

### Return type

[**ChatRoomView**](ChatRoomView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkChatAsReadUsingPUT

> ChatView MarkChatAsReadUsingPUT(ctx).ChatId(chatId).Execute()

Sets a chat as read by logged in user

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
	chatId := "chatId_example" // string | Chat ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.MarkChatAsReadUsingPUT(context.Background()).ChatId(chatId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.MarkChatAsReadUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkChatAsReadUsingPUT`: ChatView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.MarkChatAsReadUsingPUT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkChatAsReadUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatId** | **string** | Chat ID | 

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


## MarkChatAsReadV1UsingPUT

> ChatView MarkChatAsReadV1UsingPUT(ctx).ChatId(chatId).Execute()

Sets a chat as read by logged in user

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
	chatId := "chatId_example" // string | Chat ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.MarkChatAsReadV1UsingPUT(context.Background()).ChatId(chatId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.MarkChatAsReadV1UsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkChatAsReadV1UsingPUT`: ChatView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.MarkChatAsReadV1UsingPUT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkChatAsReadV1UsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatId** | **string** | Chat ID | 

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


## QuitChatUsingPUT

> ChatView QuitChatUsingPUT(ctx, chatId).Execute()

Removes logged in user from chat

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.QuitChatUsingPUT(context.Background(), chatId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.QuitChatUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuitChatUsingPUT`: ChatView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.QuitChatUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | Chat id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuitChatUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## RemoveChatMembershipUsingDELETE

> MessagePrototype RemoveChatMembershipUsingDELETE(ctx, membershipId).Execute()

Removes chat membership

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
	membershipId := "membershipId_example" // string | Membership id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.RemoveChatMembershipUsingDELETE(context.Background(), membershipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.RemoveChatMembershipUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveChatMembershipUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.RemoveChatMembershipUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**membershipId** | **string** | Membership id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveChatMembershipUsingDELETERequest struct via the builder pattern


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


## RemoveFromChatUsingPUT

> ChatView RemoveFromChatUsingPUT(ctx).ChatId(chatId).UserId(userId).Execute()

Removes user from chat

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
	userId := "userId_example" // string | User id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.RemoveFromChatUsingPUT(context.Background()).ChatId(chatId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.RemoveFromChatUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveFromChatUsingPUT`: ChatView
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.RemoveFromChatUsingPUT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFromChatUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatId** | **string** | Chat id | 
 **userId** | **string** | User id | 

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


## RemoveMyChatMembershipUsingDELETE

> MessagePrototype RemoveMyChatMembershipUsingDELETE(ctx, chatId).Execute()

Removes logged in user's chat membership

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatControllerAPI.RemoveMyChatMembershipUsingDELETE(context.Background(), chatId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatControllerAPI.RemoveMyChatMembershipUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveMyChatMembershipUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `ChatControllerAPI.RemoveMyChatMembershipUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | Chat id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMyChatMembershipUsingDELETERequest struct via the builder pattern


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

