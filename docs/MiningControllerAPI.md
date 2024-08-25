# \MiningControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMyMinerUsingGET**](MiningControllerAPI.md#GetMyMinerUsingGET) | **Get** /api/v2/my/miner | Return private miner
[**TransferUsingPUT**](MiningControllerAPI.md#TransferUsingPUT) | **Put** /api/v2/my/cointransfer | Transfer privately mined coins to own portfolio
[**UpgradeUsingPUT**](MiningControllerAPI.md#UpgradeUsingPUT) | **Put** /api/v2/my/minerupgrade | Upgrade private miner



## GetMyMinerUsingGET

> MinerView GetMyMinerUsingGET(ctx).Execute()

Return private miner

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
	resp, r, err := apiClient.MiningControllerAPI.GetMyMinerUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiningControllerAPI.GetMyMinerUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyMinerUsingGET`: MinerView
	fmt.Fprintf(os.Stdout, "Response from `MiningControllerAPI.GetMyMinerUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyMinerUsingGETRequest struct via the builder pattern


### Return type

[**MinerView**](MinerView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferUsingPUT

> MinerView TransferUsingPUT(ctx).Execute()

Transfer privately mined coins to own portfolio

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
	resp, r, err := apiClient.MiningControllerAPI.TransferUsingPUT(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiningControllerAPI.TransferUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferUsingPUT`: MinerView
	fmt.Fprintf(os.Stdout, "Response from `MiningControllerAPI.TransferUsingPUT`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTransferUsingPUTRequest struct via the builder pattern


### Return type

[**MinerView**](MinerView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeUsingPUT

> MinerView UpgradeUsingPUT(ctx).Execute()

Upgrade private miner

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
	resp, r, err := apiClient.MiningControllerAPI.UpgradeUsingPUT(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiningControllerAPI.UpgradeUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeUsingPUT`: MinerView
	fmt.Fprintf(os.Stdout, "Response from `MiningControllerAPI.UpgradeUsingPUT`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeUsingPUTRequest struct via the builder pattern


### Return type

[**MinerView**](MinerView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

