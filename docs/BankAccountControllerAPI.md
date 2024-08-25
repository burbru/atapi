# \BankAccountControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBankAccountUsingGET**](BankAccountControllerAPI.md#GetBankAccountUsingGET) | **Get** /api/v2/bankaccounts/{bankAccountId} | Returns bank account
[**GetMyBankAccountUsingGET**](BankAccountControllerAPI.md#GetMyBankAccountUsingGET) | **Get** /api/v2/my/bankaccounts/ | Returns user&#39;s bank accounts
[**GetMyBankAccountV1UsingGET**](BankAccountControllerAPI.md#GetMyBankAccountV1UsingGET) | **Get** /api/bankaccounts/ | Returns user&#39;s bank account
[**SendMoneyUsingPUT**](BankAccountControllerAPI.md#SendMoneyUsingPUT) | **Put** /api/v2/banktransfer/{senderBankAccountId} | Bank transfer



## GetBankAccountUsingGET

> BankAccountView GetBankAccountUsingGET(ctx, bankAccountId).Execute()

Returns bank account

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankAccountControllerAPI.GetBankAccountUsingGET(context.Background(), bankAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankAccountControllerAPI.GetBankAccountUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankAccountUsingGET`: BankAccountView
	fmt.Fprintf(os.Stdout, "Response from `BankAccountControllerAPI.GetBankAccountUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankAccountId** | **string** | Bank account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankAccountUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BankAccountView**](BankAccountView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyBankAccountUsingGET

> []BankAccountView GetMyBankAccountUsingGET(ctx).Execute()

Returns user's bank accounts

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
	resp, r, err := apiClient.BankAccountControllerAPI.GetMyBankAccountUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankAccountControllerAPI.GetMyBankAccountUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyBankAccountUsingGET`: []BankAccountView
	fmt.Fprintf(os.Stdout, "Response from `BankAccountControllerAPI.GetMyBankAccountUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyBankAccountUsingGETRequest struct via the builder pattern


### Return type

[**[]BankAccountView**](BankAccountView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyBankAccountV1UsingGET

> BankAccountView GetMyBankAccountV1UsingGET(ctx).Execute()

Returns user's bank account

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
	resp, r, err := apiClient.BankAccountControllerAPI.GetMyBankAccountV1UsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankAccountControllerAPI.GetMyBankAccountV1UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyBankAccountV1UsingGET`: BankAccountView
	fmt.Fprintf(os.Stdout, "Response from `BankAccountControllerAPI.GetMyBankAccountV1UsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyBankAccountV1UsingGETRequest struct via the builder pattern


### Return type

[**BankAccountView**](BankAccountView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMoneyUsingPUT

> MessagePrototype SendMoneyUsingPUT(ctx, senderBankAccountId).CashAmount(cashAmount).ReceiverBankAccountId(receiverBankAccountId).Execute()

Bank transfer

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
	cashAmount := "cashAmount_example" // string | Cash amount
	receiverBankAccountId := "receiverBankAccountId_example" // string | Receiver bank account id
	senderBankAccountId := "senderBankAccountId_example" // string | Sender bank account id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankAccountControllerAPI.SendMoneyUsingPUT(context.Background(), senderBankAccountId).CashAmount(cashAmount).ReceiverBankAccountId(receiverBankAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankAccountControllerAPI.SendMoneyUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendMoneyUsingPUT`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `BankAccountControllerAPI.SendMoneyUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**senderBankAccountId** | **string** | Sender bank account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendMoneyUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cashAmount** | **string** | Cash amount | 
 **receiverBankAccountId** | **string** | Receiver bank account id | 


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

