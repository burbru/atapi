# \MessageBoardControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddToBoardUsingPOST**](MessageBoardControllerAPI.md#AddToBoardUsingPOST) | **Post** /api/v2/messageboardmemberships | Adds user to messageboard
[**ChangeBoardMembershipUsingPUT**](MessageBoardControllerAPI.md#ChangeBoardMembershipUsingPUT) | **Put** /api/v2/messageboardmemberships/{membershipId} | Change messageboard membership
[**CreateMessageBoardUsingPOST**](MessageBoardControllerAPI.md#CreateMessageBoardUsingPOST) | **Post** /api/v2/messageboards | Creates message board
[**CreateSubboardUsingPOST**](MessageBoardControllerAPI.md#CreateSubboardUsingPOST) | **Post** /api/v2/messageboards/{parentId}/subboards | Creates subboard
[**DeleteMessageBoardUsingDELETE**](MessageBoardControllerAPI.md#DeleteMessageBoardUsingDELETE) | **Delete** /api/v2/messageboards/{boardId} | Deletes message (sub)board and its subboards, memberships and postings
[**EditMessageBoardUsingPUT**](MessageBoardControllerAPI.md#EditMessageBoardUsingPUT) | **Put** /api/v2/messageboards/{boardId} | Changes message board
[**GetBoardMembershipsUsingGET**](MessageBoardControllerAPI.md#GetBoardMembershipsUsingGET) | **Get** /api/v2/messageboardmemberships | Lists memberships of message board
[**GetMembershipUsingGET**](MessageBoardControllerAPI.md#GetMembershipUsingGET) | **Get** /api/v2/messageboardmemberships/{membershipId} | Returns messageboard membership
[**GetMessageBoardUsingGET**](MessageBoardControllerAPI.md#GetMessageBoardUsingGET) | **Get** /api/v2/messageboards/{boardId} | Returns message (sub)board
[**ListMessageBoardPostsUsingGET**](MessageBoardControllerAPI.md#ListMessageBoardPostsUsingGET) | **Get** /api/v2/messageboards/{boardId}/posts | Lists message board&#39;s posts
[**ListMyBoardsUsingGET**](MessageBoardControllerAPI.md#ListMyBoardsUsingGET) | **Get** /api/v2/my/messageboards | Lists message boards in which the logged in user is a member
[**ListPostsUsingGET**](MessageBoardControllerAPI.md#ListPostsUsingGET) | **Get** /api/v2/my/boardnews | Lists most interesting news
[**ListPublicBoardsUsingGET**](MessageBoardControllerAPI.md#ListPublicBoardsUsingGET) | **Get** /api/v2/messageboards | Lists public message boards
[**ListsReachableBoardsUsingGET**](MessageBoardControllerAPI.md#ListsReachableBoardsUsingGET) | **Get** /api/v2/messageboards/{boardId}/branch | Lists all (sub)boards of same message board tree branch
[**ListsSubboardsUsingGET**](MessageBoardControllerAPI.md#ListsSubboardsUsingGET) | **Get** /api/v2/messageboards/{parentId}/subboards | Lists subboards
[**RemoveBoardMembershipOfBoardUsingDELETE**](MessageBoardControllerAPI.md#RemoveBoardMembershipOfBoardUsingDELETE) | **Delete** /api/v2/messageboardmemberships | Removes message board membership
[**RemoveBoardMembershipUsingDELETE**](MessageBoardControllerAPI.md#RemoveBoardMembershipUsingDELETE) | **Delete** /api/v2/messageboardmemberships/{membershipId} | Removes message board membership



## AddToBoardUsingPOST

> MessageBoardMembershipView AddToBoardUsingPOST(ctx).MessageBoardId(messageBoardId).UserId(userId).Role(role).Execute()

Adds user to messageboard

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
	messageBoardId := "messageBoardId_example" // string | Message board id
	userId := "userId_example" // string | User id
	role := "role_example" // string | Role (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageBoardControllerAPI.AddToBoardUsingPOST(context.Background()).MessageBoardId(messageBoardId).UserId(userId).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageBoardControllerAPI.AddToBoardUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddToBoardUsingPOST`: MessageBoardMembershipView
	fmt.Fprintf(os.Stdout, "Response from `MessageBoardControllerAPI.AddToBoardUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddToBoardUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageBoardId** | **string** | Message board id | 
 **userId** | **string** | User id | 
 **role** | **string** | Role | 

### Return type

[**MessageBoardMembershipView**](MessageBoardMembershipView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeBoardMembershipUsingPUT

> MessageBoardMembershipView ChangeBoardMembershipUsingPUT(ctx, membershipId).Role(role).Execute()

Change messageboard membership

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
	resp, r, err := apiClient.MessageBoardControllerAPI.ChangeBoardMembershipUsingPUT(context.Background(), membershipId).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageBoardControllerAPI.ChangeBoardMembershipUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeBoardMembershipUsingPUT`: MessageBoardMembershipView
	fmt.Fprintf(os.Stdout, "Response from `MessageBoardControllerAPI.ChangeBoardMembershipUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**membershipId** | **string** | Membership id | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeBoardMembershipUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | **string** | Role | 

### Return type

[**MessageBoardMembershipView**](MessageBoardMembershipView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessageBoardUsingPOST

> MessageBoardView CreateMessageBoardUsingPOST(ctx).Locale(locale).Name(name).Description(description).Public(public).Execute()

Creates message board

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
	locale := "locale_example" // string | Locale
	name := "name_example" // string | Name
	description := "description_example" // string | Description (optional)
	public := true // bool | Public (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageBoardControllerAPI.CreateMessageBoardUsingPOST(context.Background()).Locale(locale).Name(name).Description(description).Public(public).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageBoardControllerAPI.CreateMessageBoardUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMessageBoardUsingPOST`: MessageBoardView
	fmt.Fprintf(os.Stdout, "Response from `MessageBoardControllerAPI.CreateMessageBoardUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessageBoardUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locale** | **string** | Locale | 
 **name** | **string** | Name | 
 **description** | **string** | Description | 
 **public** | **bool** | Public | 

### Return type

[**MessageBoardView**](MessageBoardView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubboardUsingPOST

> MessageBoardView CreateSubboardUsingPOST(ctx, parentId).Name(name).Description(description).Execute()

Creates subboard

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
	name := "name_example" // string | Name
	parentId := "parentId_example" // string | Parent board id
	description := "description_example" // string | Description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageBoardControllerAPI.CreateSubboardUsingPOST(context.Background(), parentId).Name(name).Description(description).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageBoardControllerAPI.CreateSubboardUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubboardUsingPOST`: MessageBoardView
	fmt.Fprintf(os.Stdout, "Response from `MessageBoardControllerAPI.CreateSubboardUsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **string** | Parent board id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubboardUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name | 

 **description** | **string** | Description | 

### Return type

[**MessageBoardView**](MessageBoardView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMessageBoardUsingDELETE

> string DeleteMessageBoardUsingDELETE(ctx, boardId).Execute()

Deletes message (sub)board and its subboards, memberships and postings

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
	boardId := "boardId_example" // string | Message board id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageBoardControllerAPI.DeleteMessageBoardUsingDELETE(context.Background(), boardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageBoardControllerAPI.DeleteMessageBoardUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMessageBoardUsingDELETE`: string
	fmt.Fprintf(os.Stdout, "Response from `MessageBoardControllerAPI.DeleteMessageBoardUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **string** | Message board id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMessageBoardUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditMessageBoardUsingPUT

> MessageBoardView EditMessageBoardUsingPUT(ctx, boardId).Locale(locale).Name(name).Description(description).Public(public).Execute()

Changes message board

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
	boardId := "boardId_example" // string | Message board id
	locale := "locale_example" // string | Locale
	name := "name_example" // string | Name
	description := "description_example" // string | Description (optional)
	public := true // bool | Public (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageBoardControllerAPI.EditMessageBoardUsingPUT(context.Background(), boardId).Locale(locale).Name(name).Description(description).Public(public).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageBoardControllerAPI.EditMessageBoardUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditMessageBoardUsingPUT`: MessageBoardView
	fmt.Fprintf(os.Stdout, "Response from `MessageBoardControllerAPI.EditMessageBoardUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **string** | Message board id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditMessageBoardUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locale** | **string** | Locale | 
 **name** | **string** | Name | 
 **description** | **string** | Description | 
 **public** | **bool** | Public | 

### Return type

[**MessageBoardView**](MessageBoardView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoardMembershipsUsingGET

> MessageBoardMembershipView GetBoardMembershipsUsingGET(ctx).MessageBoardId(messageBoardId).Execute()

Lists memberships of message board

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
	messageBoardId := "messageBoardId_example" // string | Message board id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageBoardControllerAPI.GetBoardMembershipsUsingGET(context.Background()).MessageBoardId(messageBoardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageBoardControllerAPI.GetBoardMembershipsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBoardMembershipsUsingGET`: MessageBoardMembershipView
	fmt.Fprintf(os.Stdout, "Response from `MessageBoardControllerAPI.GetBoardMembershipsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBoardMembershipsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageBoardId** | **string** | Message board id | 

### Return type

[**MessageBoardMembershipView**](MessageBoardMembershipView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMembershipUsingGET

> MessageBoardMembershipView GetMembershipUsingGET(ctx, membershipId).Execute()

Returns messageboard membership

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
	resp, r, err := apiClient.MessageBoardControllerAPI.GetMembershipUsingGET(context.Background(), membershipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageBoardControllerAPI.GetMembershipUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMembershipUsingGET`: MessageBoardMembershipView
	fmt.Fprintf(os.Stdout, "Response from `MessageBoardControllerAPI.GetMembershipUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**membershipId** | **string** | Membership id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMembershipUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageBoardMembershipView**](MessageBoardMembershipView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageBoardUsingGET

> MessageBoardWithDetailsView GetMessageBoardUsingGET(ctx, boardId).Execute()

Returns message (sub)board

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
	boardId := "boardId_example" // string | Message board id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageBoardControllerAPI.GetMessageBoardUsingGET(context.Background(), boardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageBoardControllerAPI.GetMessageBoardUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessageBoardUsingGET`: MessageBoardWithDetailsView
	fmt.Fprintf(os.Stdout, "Response from `MessageBoardControllerAPI.GetMessageBoardUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **string** | Message board id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageBoardUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageBoardWithDetailsView**](MessageBoardWithDetailsView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessageBoardPostsUsingGET

> PagePostView ListMessageBoardPostsUsingGET(ctx, boardId).Page(page).Search(search).Size(size).Sort(sort).Execute()

Lists message board's posts



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
	boardId := "boardId_example" // string | Message Board ID
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	search := "search_example" // string | Fulltext search in title and content (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageBoardControllerAPI.ListMessageBoardPostsUsingGET(context.Background(), boardId).Page(page).Search(search).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageBoardControllerAPI.ListMessageBoardPostsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMessageBoardPostsUsingGET`: PagePostView
	fmt.Fprintf(os.Stdout, "Response from `MessageBoardControllerAPI.ListMessageBoardPostsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **string** | Message Board ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMessageBoardPostsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **search** | **string** | Fulltext search in title and content | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PagePostView**](PagePostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMyBoardsUsingGET

> PageMessageBoardWithDetailsView ListMyBoardsUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Lists message boards in which the logged in user is a member

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
	resp, r, err := apiClient.MessageBoardControllerAPI.ListMyBoardsUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageBoardControllerAPI.ListMyBoardsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyBoardsUsingGET`: PageMessageBoardWithDetailsView
	fmt.Fprintf(os.Stdout, "Response from `MessageBoardControllerAPI.ListMyBoardsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMyBoardsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageMessageBoardWithDetailsView**](PageMessageBoardWithDetailsView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPostsUsingGET

> PostView ListPostsUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Lists most interesting news

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
	resp, r, err := apiClient.MessageBoardControllerAPI.ListPostsUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageBoardControllerAPI.ListPostsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPostsUsingGET`: PostView
	fmt.Fprintf(os.Stdout, "Response from `MessageBoardControllerAPI.ListPostsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPostsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PostView**](PostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPublicBoardsUsingGET

> PageMessageBoardWithDetailsView ListPublicBoardsUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Lists public message boards

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
	resp, r, err := apiClient.MessageBoardControllerAPI.ListPublicBoardsUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageBoardControllerAPI.ListPublicBoardsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPublicBoardsUsingGET`: PageMessageBoardWithDetailsView
	fmt.Fprintf(os.Stdout, "Response from `MessageBoardControllerAPI.ListPublicBoardsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPublicBoardsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageMessageBoardWithDetailsView**](PageMessageBoardWithDetailsView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsReachableBoardsUsingGET

> PageMessageBoardWithDetailsView ListsReachableBoardsUsingGET(ctx, boardId).Page(page).Size(size).Sort(sort).Execute()

Lists all (sub)boards of same message board tree branch

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
	boardId := "boardId_example" // string | Board id
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageBoardControllerAPI.ListsReachableBoardsUsingGET(context.Background(), boardId).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageBoardControllerAPI.ListsReachableBoardsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListsReachableBoardsUsingGET`: PageMessageBoardWithDetailsView
	fmt.Fprintf(os.Stdout, "Response from `MessageBoardControllerAPI.ListsReachableBoardsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **string** | Board id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsReachableBoardsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageMessageBoardWithDetailsView**](PageMessageBoardWithDetailsView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsSubboardsUsingGET

> PageMessageBoardWithDetailsView ListsSubboardsUsingGET(ctx, parentId).Page(page).Size(size).Sort(sort).Execute()

Lists subboards

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
	parentId := "parentId_example" // string | Parent board id
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageBoardControllerAPI.ListsSubboardsUsingGET(context.Background(), parentId).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageBoardControllerAPI.ListsSubboardsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListsSubboardsUsingGET`: PageMessageBoardWithDetailsView
	fmt.Fprintf(os.Stdout, "Response from `MessageBoardControllerAPI.ListsSubboardsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **string** | Parent board id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsSubboardsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageMessageBoardWithDetailsView**](PageMessageBoardWithDetailsView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveBoardMembershipOfBoardUsingDELETE

> MessagePrototype RemoveBoardMembershipOfBoardUsingDELETE(ctx).BoardId(boardId).Execute()

Removes message board membership

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
	boardId := "boardId_example" // string | Message board id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageBoardControllerAPI.RemoveBoardMembershipOfBoardUsingDELETE(context.Background()).BoardId(boardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageBoardControllerAPI.RemoveBoardMembershipOfBoardUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveBoardMembershipOfBoardUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `MessageBoardControllerAPI.RemoveBoardMembershipOfBoardUsingDELETE`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveBoardMembershipOfBoardUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **boardId** | **string** | Message board id | 

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


## RemoveBoardMembershipUsingDELETE

> MessagePrototype RemoveBoardMembershipUsingDELETE(ctx, membershipId).Execute()

Removes message board membership

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
	resp, r, err := apiClient.MessageBoardControllerAPI.RemoveBoardMembershipUsingDELETE(context.Background(), membershipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageBoardControllerAPI.RemoveBoardMembershipUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveBoardMembershipUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `MessageBoardControllerAPI.RemoveBoardMembershipUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**membershipId** | **string** | Membership id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveBoardMembershipUsingDELETERequest struct via the builder pattern


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

