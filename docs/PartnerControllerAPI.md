# \PartnerControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPartnersUsingGET**](PartnerControllerAPI.md#ListPartnersUsingGET) | **Get** /api/partners/ | Lists partners
[**RegisterAsPartnerUsingPUT**](PartnerControllerAPI.md#RegisterAsPartnerUsingPUT) | **Put** /api/partners/register | Registers logged in user as partner with new partner id



## ListPartnersUsingGET

> UsernameView ListPartnersUsingGET(ctx).Execute()

Lists partners

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
	resp, r, err := apiClient.PartnerControllerAPI.ListPartnersUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartnerControllerAPI.ListPartnersUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPartnersUsingGET`: UsernameView
	fmt.Fprintf(os.Stdout, "Response from `PartnerControllerAPI.ListPartnersUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPartnersUsingGETRequest struct via the builder pattern


### Return type

[**UsernameView**](UsernameView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterAsPartnerUsingPUT

> UserAccountView RegisterAsPartnerUsingPUT(ctx).Execute()

Registers logged in user as partner with new partner id

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
	resp, r, err := apiClient.PartnerControllerAPI.RegisterAsPartnerUsingPUT(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartnerControllerAPI.RegisterAsPartnerUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterAsPartnerUsingPUT`: UserAccountView
	fmt.Fprintf(os.Stdout, "Response from `PartnerControllerAPI.RegisterAsPartnerUsingPUT`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAsPartnerUsingPUTRequest struct via the builder pattern


### Return type

[**UserAccountView**](UserAccountView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

