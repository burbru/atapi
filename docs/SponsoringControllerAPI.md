# \SponsoringControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListGoalsUsingGET**](SponsoringControllerAPI.md#ListGoalsUsingGET) | **Get** /api/v2/sponsoringgoals | Lists all sponsoring goals
[**ListSponsorsUsingGET**](SponsoringControllerAPI.md#ListSponsorsUsingGET) | **Get** /api/v2/sponsors | Lists all sponsors



## ListGoalsUsingGET

> PageSponsoringGoalView ListGoalsUsingGET(ctx).Page(page).Size(size).Sort(sort).Type_(type_).Execute()

Lists all sponsoring goals



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
	type_ := "type__example" // string | type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsoringControllerAPI.ListGoalsUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsoringControllerAPI.ListGoalsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGoalsUsingGET`: PageSponsoringGoalView
	fmt.Fprintf(os.Stdout, "Response from `SponsoringControllerAPI.ListGoalsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGoalsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **type_** | **string** | type | 

### Return type

[**PageSponsoringGoalView**](PageSponsoringGoalView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSponsorsUsingGET

> PageSponsorView ListSponsorsUsingGET(ctx).AfterDate(afterDate).Page(page).Size(size).Sort(sort).Execute()

Lists all sponsors



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
	afterDate := int64(789) // int64 | Only lists sponsored hours after Date (optional)
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsoringControllerAPI.ListSponsorsUsingGET(context.Background()).AfterDate(afterDate).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsoringControllerAPI.ListSponsorsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSponsorsUsingGET`: PageSponsorView
	fmt.Fprintf(os.Stdout, "Response from `SponsoringControllerAPI.ListSponsorsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSponsorsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afterDate** | **int64** | Only lists sponsored hours after Date | 
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageSponsorView**](PageSponsorView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

