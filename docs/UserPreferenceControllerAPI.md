# \UserPreferenceControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserPreferenceUsingPOST**](UserPreferenceControllerAPI.md#CreateUserPreferenceUsingPOST) | **Post** /api/v2/userpreferences | Creates user preference
[**DeleteUserPreferenceUsingDELETE**](UserPreferenceControllerAPI.md#DeleteUserPreferenceUsingDELETE) | **Delete** /api/v2/userpreferences/{prefId} | Delete user preference
[**EditUserPreferenceUsingPUT**](UserPreferenceControllerAPI.md#EditUserPreferenceUsingPUT) | **Put** /api/v2/userpreferences/{prefId} | Edit user preference
[**ListUserPreferencesUsingGET**](UserPreferenceControllerAPI.md#ListUserPreferencesUsingGET) | **Get** /api/v2/userpreferences | List user preferences



## CreateUserPreferenceUsingPOST

> map[string]interface{} CreateUserPreferenceUsingPOST(ctx).Content(content).Identifier(identifier).Type_(type_).Execute()

Creates user preference

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
	content := "content_example" // string | Content of the user preference
	identifier := "identifier_example" // string | Identifier of the user preference
	type_ := "type__example" // string | Type of the user preference

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserPreferenceControllerAPI.CreateUserPreferenceUsingPOST(context.Background()).Content(content).Identifier(identifier).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserPreferenceControllerAPI.CreateUserPreferenceUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserPreferenceUsingPOST`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserPreferenceControllerAPI.CreateUserPreferenceUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserPreferenceUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **content** | **string** | Content of the user preference | 
 **identifier** | **string** | Identifier of the user preference | 
 **type_** | **string** | Type of the user preference | 

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


## DeleteUserPreferenceUsingDELETE

> map[string]interface{} DeleteUserPreferenceUsingDELETE(ctx, prefId).Execute()

Delete user preference

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
	prefId := "prefId_example" // string | Preference id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserPreferenceControllerAPI.DeleteUserPreferenceUsingDELETE(context.Background(), prefId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserPreferenceControllerAPI.DeleteUserPreferenceUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserPreferenceUsingDELETE`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserPreferenceControllerAPI.DeleteUserPreferenceUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefId** | **string** | Preference id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserPreferenceUsingDELETERequest struct via the builder pattern


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


## EditUserPreferenceUsingPUT

> map[string]interface{} EditUserPreferenceUsingPUT(ctx, prefId).Content(content).Identifier(identifier).Type_(type_).Execute()

Edit user preference

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
	prefId := "prefId_example" // string | Preference id
	content := "content_example" // string | Content of the user preference (optional)
	identifier := "identifier_example" // string | Identifier of the user preference (optional)
	type_ := "type__example" // string | Type of the user preference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserPreferenceControllerAPI.EditUserPreferenceUsingPUT(context.Background(), prefId).Content(content).Identifier(identifier).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserPreferenceControllerAPI.EditUserPreferenceUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditUserPreferenceUsingPUT`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserPreferenceControllerAPI.EditUserPreferenceUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefId** | **string** | Preference id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditUserPreferenceUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **content** | **string** | Content of the user preference | 
 **identifier** | **string** | Identifier of the user preference | 
 **type_** | **string** | Type of the user preference | 

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


## ListUserPreferencesUsingGET

> PageUserPreferenceView ListUserPreferencesUsingGET(ctx).Identifier(identifier).Page(page).Size(size).Sort(sort).Type_(type_).Execute()

List user preferences

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
	identifier := "identifier_example" // string | Identifier of the user preference (optional)
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)
	type_ := "type__example" // string | Type of the user preference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserPreferenceControllerAPI.ListUserPreferencesUsingGET(context.Background()).Identifier(identifier).Page(page).Size(size).Sort(sort).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserPreferenceControllerAPI.ListUserPreferencesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserPreferencesUsingGET`: PageUserPreferenceView
	fmt.Fprintf(os.Stdout, "Response from `UserPreferenceControllerAPI.ListUserPreferencesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUserPreferencesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifier** | **string** | Identifier of the user preference | 
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **type_** | **string** | Type of the user preference | 

### Return type

[**PageUserPreferenceView**](PageUserPreferenceView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

