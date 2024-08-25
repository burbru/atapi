# \CashTransferLogControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLogsUsingGET**](CashTransferLogControllerAPI.md#ListLogsUsingGET) | **Get** /api/v2/cashtransferlogs/{bankAccountId} | Lists cash transfer logs
[**ListLogsV1UsingGET**](CashTransferLogControllerAPI.md#ListLogsV1UsingGET) | **Get** /api/cashtransferlogs | Lists cash transfer logs



## ListLogsUsingGET

> CashTransferLogEntryView ListLogsUsingGET(ctx, bankAccountId).Page(page).Search(search).Size(size).Sort(sort).Execute()

Lists cash transfer logs

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
	bankAccountId := "bankAccountId_example" // string | Bank account id
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	search := "search_example" // string | Fulltext search (be aware of non-localized subject and content and be aware of MessagePrototype's structure with substitutions; useful for security identifiers) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashTransferLogControllerAPI.ListLogsUsingGET(context.Background(), bankAccountId).Page(page).Search(search).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CashTransferLogControllerAPI.ListLogsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLogsUsingGET`: CashTransferLogEntryView
	fmt.Fprintf(os.Stdout, "Response from `CashTransferLogControllerAPI.ListLogsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankAccountId** | **string** | Bank account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **search** | **string** | Fulltext search (be aware of non-localized subject and content and be aware of MessagePrototype&#39;s structure with substitutions; useful for security identifiers) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**CashTransferLogEntryView**](CashTransferLogEntryView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogsV1UsingGET

> CashTransferLogEntryView ListLogsV1UsingGET(ctx).EndDate(endDate).ReceiverBankAccountId(receiverBankAccountId).SenderBankAccountId(senderBankAccountId).StartDate(startDate).Execute()

Lists cash transfer logs

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
	endDate := "endDate_example" // string | End date (optional)
	receiverBankAccountId := "receiverBankAccountId_example" // string | Receiver's bank account id (optional)
	senderBankAccountId := "senderBankAccountId_example" // string | Sender's bank account id (optional)
	startDate := "startDate_example" // string | Start date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashTransferLogControllerAPI.ListLogsV1UsingGET(context.Background()).EndDate(endDate).ReceiverBankAccountId(receiverBankAccountId).SenderBankAccountId(senderBankAccountId).StartDate(startDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CashTransferLogControllerAPI.ListLogsV1UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLogsV1UsingGET`: CashTransferLogEntryView
	fmt.Fprintf(os.Stdout, "Response from `CashTransferLogControllerAPI.ListLogsV1UsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogsV1UsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endDate** | **string** | End date | 
 **receiverBankAccountId** | **string** | Receiver&#39;s bank account id | 
 **senderBankAccountId** | **string** | Sender&#39;s bank account id | 
 **startDate** | **string** | Start date | 

### Return type

[**CashTransferLogEntryView**](CashTransferLogEntryView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

