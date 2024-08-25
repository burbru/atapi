# \ComplaintControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComplaintUsingPOST**](ComplaintControllerAPI.md#CreateComplaintUsingPOST) | **Post** /api/complaints | Opens new complaint
[**GetComplaintUsingGET**](ComplaintControllerAPI.md#GetComplaintUsingGET) | **Get** /api/complaints/{complaintId} | Returns complaint



## CreateComplaintUsingPOST

> ComplaintView CreateComplaintUsingPOST(ctx).SubjectMatterId(subjectMatterId).SubjectMatterType(subjectMatterType).Text(text).Execute()

Opens new complaint

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
	subjectMatterId := "subjectMatterId_example" // string | Subject Matter ID
	subjectMatterType := "subjectMatterType_example" // string | Subject Matter Type
	text := "text_example" // string | Text (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplaintControllerAPI.CreateComplaintUsingPOST(context.Background()).SubjectMatterId(subjectMatterId).SubjectMatterType(subjectMatterType).Text(text).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplaintControllerAPI.CreateComplaintUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateComplaintUsingPOST`: ComplaintView
	fmt.Fprintf(os.Stdout, "Response from `ComplaintControllerAPI.CreateComplaintUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateComplaintUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subjectMatterId** | **string** | Subject Matter ID | 
 **subjectMatterType** | **string** | Subject Matter Type | 
 **text** | **string** | Text | 

### Return type

[**ComplaintView**](ComplaintView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplaintUsingGET

> ComplaintView GetComplaintUsingGET(ctx, complaintId).Execute()

Returns complaint

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
	complaintId := "complaintId_example" // string | Complaint id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplaintControllerAPI.GetComplaintUsingGET(context.Background(), complaintId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplaintControllerAPI.GetComplaintUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComplaintUsingGET`: ComplaintView
	fmt.Fprintf(os.Stdout, "Response from `ComplaintControllerAPI.GetComplaintUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**complaintId** | **string** | Complaint id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComplaintUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComplaintView**](ComplaintView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

