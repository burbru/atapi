# \HistoryControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistoryEntriesUsingGET**](HistoryControllerAPI.md#GetHistoryEntriesUsingGET) | **Get** /api/v2/history | List entity&#39;s history entries



## GetHistoryEntriesUsingGET

> PageHistoryEntryView GetHistoryEntriesUsingGET(ctx).EntityId(entityId).HistoryType(historyType).Page(page).Size(size).Sort(sort).Execute()

List entity's history entries

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
	entityId := "entityId_example" // string | Entity ID
	historyType := "historyType_example" // string | History Entry Type (optional)
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoryControllerAPI.GetHistoryEntriesUsingGET(context.Background()).EntityId(entityId).HistoryType(historyType).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoryControllerAPI.GetHistoryEntriesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoryEntriesUsingGET`: PageHistoryEntryView
	fmt.Fprintf(os.Stdout, "Response from `HistoryControllerAPI.GetHistoryEntriesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryEntriesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | Entity ID | 
 **historyType** | **string** | History Entry Type | 
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageHistoryEntryView**](PageHistoryEntryView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

