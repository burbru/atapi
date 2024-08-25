# \SharePositionControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPosUsingGET**](SharePositionControllerAPI.md#ListPosUsingGET) | **Get** /api/v2/sharepositions | Lists all positions, from empire without securitiesAccountId



## ListPosUsingGET

> []SharePositionView ListPosUsingGET(ctx).SecurityIdentifier(securityIdentifier).SecuritiesAccountId(securitiesAccountId).Execute()

Lists all positions, from empire without securitiesAccountId

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
	securityIdentifier := "securityIdentifier_example" // string | securityIdentifier
	securitiesAccountId := "securitiesAccountId_example" // string | securitiesAccountId (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharePositionControllerAPI.ListPosUsingGET(context.Background()).SecurityIdentifier(securityIdentifier).SecuritiesAccountId(securitiesAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharePositionControllerAPI.ListPosUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPosUsingGET`: []SharePositionView
	fmt.Fprintf(os.Stdout, "Response from `SharePositionControllerAPI.ListPosUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPosUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityIdentifier** | **string** | securityIdentifier | 
 **securitiesAccountId** | **string** | securitiesAccountId | 

### Return type

[**[]SharePositionView**](SharePositionView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

