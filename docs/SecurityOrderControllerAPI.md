# \SecurityOrderControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrderUsingPOST**](SecurityOrderControllerAPI.md#AddOrderUsingPOST) | **Post** /api/securityorders | Adds a new security order
[**CheckOrderUsingGET**](SecurityOrderControllerAPI.md#CheckOrderUsingGET) | **Get** /api/securityorders/check/ | Checks Security Order for Acceptance (DEPRECATED: Use addOrder with checkOrder &#x3D; true)
[**DeleteOrderUsingDELETE**](SecurityOrderControllerAPI.md#DeleteOrderUsingDELETE) | **Delete** /api/securityorders/{orderId} | Deletes a single order
[**DeleteOrdersUsingDELETE**](SecurityOrderControllerAPI.md#DeleteOrdersUsingDELETE) | **Delete** /api/securityorders | Deletes all orders
[**GetCounterOtcOrdersUsingGET**](SecurityOrderControllerAPI.md#GetCounterOtcOrdersUsingGET) | **Get** /api/securityorders/counterparty/{securitiesAccountId} | Lists unfilled counter otc orders for securities account
[**GetOrderUsingGET**](SecurityOrderControllerAPI.md#GetOrderUsingGET) | **Get** /api/securityorders/{orderId} | Returns a single order
[**GetOrderbookUsingGET**](SecurityOrderControllerAPI.md#GetOrderbookUsingGET) | **Get** /api/orderbook/{securityIdentifier} | Returns orderbook of security
[**GetOwnedOrdersUsingGET**](SecurityOrderControllerAPI.md#GetOwnedOrdersUsingGET) | **Get** /api/securityorders/securitiesaccount/{securitiesAccountId} | Lists unfilled orders owned by securities account
[**GetOwnedOrdersV2UsingGET**](SecurityOrderControllerAPI.md#GetOwnedOrdersV2UsingGET) | **Get** /api/v2/securityorders | Lists unfilled orders



## AddOrderUsingPOST

> SecurityOrder AddOrderUsingPOST(ctx).Action(action).NumberOfShares(numberOfShares).Owner(owner).SecurityIdentifier(securityIdentifier).Type_(type_).CheckOrderOnly(checkOrderOnly).Counterparty(counterparty).GoodAfterDate(goodAfterDate).GoodTillDate(goodTillDate).HourlyChange(hourlyChange).Price(price).Execute()

Adds a new security order



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
	action := "action_example" // string | SecurityOrder action
	numberOfShares := int64(789) // int64 | Number of shares
	owner := "owner_example" // string | ID of securities account
	securityIdentifier := "securityIdentifier_example" // string | Security identifier
	type_ := "type__example" // string | SecurityOrder type
	checkOrderOnly := true // bool | Check Order without execution (optional)
	counterparty := "counterparty_example" // string | OTC counterparty's securities account (optional)
	goodAfterDate := int64(789) // int64 | Good after date (premium feature) (optional)
	goodTillDate := int64(789) // int64 | Good till date (premium feature) (optional)
	hourlyChange := "hourlyChange_example" // string | Hourly change (optional)
	price := "price_example" // string | Price (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityOrderControllerAPI.AddOrderUsingPOST(context.Background()).Action(action).NumberOfShares(numberOfShares).Owner(owner).SecurityIdentifier(securityIdentifier).Type_(type_).CheckOrderOnly(checkOrderOnly).Counterparty(counterparty).GoodAfterDate(goodAfterDate).GoodTillDate(goodTillDate).HourlyChange(hourlyChange).Price(price).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityOrderControllerAPI.AddOrderUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddOrderUsingPOST`: SecurityOrder
	fmt.Fprintf(os.Stdout, "Response from `SecurityOrderControllerAPI.AddOrderUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddOrderUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | SecurityOrder action | 
 **numberOfShares** | **int64** | Number of shares | 
 **owner** | **string** | ID of securities account | 
 **securityIdentifier** | **string** | Security identifier | 
 **type_** | **string** | SecurityOrder type | 
 **checkOrderOnly** | **bool** | Check Order without execution | 
 **counterparty** | **string** | OTC counterparty&#39;s securities account | 
 **goodAfterDate** | **int64** | Good after date (premium feature) | 
 **goodTillDate** | **int64** | Good till date (premium feature) | 
 **hourlyChange** | **string** | Hourly change | 
 **price** | **string** | Price | 

### Return type

[**SecurityOrder**](SecurityOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckOrderUsingGET

> OrderCheck CheckOrderUsingGET(ctx).Action(action).Owner(owner).SecurityIdentifier(securityIdentifier).Type_(type_).NumberOfShares(numberOfShares).Price(price).Execute()

Checks Security Order for Acceptance (DEPRECATED: Use addOrder with checkOrder = true)

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
	action := "action_example" // string | SecurityOrder action
	owner := "owner_example" // string | ID of securities account
	securityIdentifier := "securityIdentifier_example" // string | Security identifier
	type_ := "type__example" // string | SecurityOrder type
	numberOfShares := int64(789) // int64 | Number of shares (optional)
	price := "price_example" // string | Price (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityOrderControllerAPI.CheckOrderUsingGET(context.Background()).Action(action).Owner(owner).SecurityIdentifier(securityIdentifier).Type_(type_).NumberOfShares(numberOfShares).Price(price).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityOrderControllerAPI.CheckOrderUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckOrderUsingGET`: OrderCheck
	fmt.Fprintf(os.Stdout, "Response from `SecurityOrderControllerAPI.CheckOrderUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckOrderUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | SecurityOrder action | 
 **owner** | **string** | ID of securities account | 
 **securityIdentifier** | **string** | Security identifier | 
 **type_** | **string** | SecurityOrder type | 
 **numberOfShares** | **int64** | Number of shares | 
 **price** | **string** | Price | 

### Return type

[**OrderCheck**](OrderCheck.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrderUsingDELETE

> MessagePrototype DeleteOrderUsingDELETE(ctx, orderId).Execute()

Deletes a single order

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
	orderId := "orderId_example" // string | SecurityOrder id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityOrderControllerAPI.DeleteOrderUsingDELETE(context.Background(), orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityOrderControllerAPI.DeleteOrderUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrderUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `SecurityOrderControllerAPI.DeleteOrderUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | SecurityOrder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrderUsingDELETERequest struct via the builder pattern


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


## DeleteOrdersUsingDELETE

> MessagePrototype DeleteOrdersUsingDELETE(ctx).Owner(owner).CreatedBefore(createdBefore).Execute()

Deletes all orders

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
	owner := "owner_example" // string | ID of securities account
	createdBefore := "createdBefore_example" // string | Orders created before (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityOrderControllerAPI.DeleteOrdersUsingDELETE(context.Background()).Owner(owner).CreatedBefore(createdBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityOrderControllerAPI.DeleteOrdersUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrdersUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `SecurityOrderControllerAPI.DeleteOrdersUsingDELETE`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrdersUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | ID of securities account | 
 **createdBefore** | **string** | Orders created before | 

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


## GetCounterOtcOrdersUsingGET

> []SecurityOrder GetCounterOtcOrdersUsingGET(ctx, securitiesAccountId).Execute()

Lists unfilled counter otc orders for securities account

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
	securitiesAccountId := "securitiesAccountId_example" // string | Securities account id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityOrderControllerAPI.GetCounterOtcOrdersUsingGET(context.Background(), securitiesAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityOrderControllerAPI.GetCounterOtcOrdersUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCounterOtcOrdersUsingGET`: []SecurityOrder
	fmt.Fprintf(os.Stdout, "Response from `SecurityOrderControllerAPI.GetCounterOtcOrdersUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securitiesAccountId** | **string** | Securities account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCounterOtcOrdersUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SecurityOrder**](SecurityOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderUsingGET

> SecurityOrder GetOrderUsingGET(ctx, orderId).Execute()

Returns a single order

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
	orderId := "orderId_example" // string | SecurityOrder id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityOrderControllerAPI.GetOrderUsingGET(context.Background(), orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityOrderControllerAPI.GetOrderUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderUsingGET`: SecurityOrder
	fmt.Fprintf(os.Stdout, "Response from `SecurityOrderControllerAPI.GetOrderUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | SecurityOrder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityOrder**](SecurityOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderbookUsingGET

> OrderBook GetOrderbookUsingGET(ctx, securityIdentifier).Execute()

Returns orderbook of security

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
	securityIdentifier := "securityIdentifier_example" // string | Security identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityOrderControllerAPI.GetOrderbookUsingGET(context.Background(), securityIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityOrderControllerAPI.GetOrderbookUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderbookUsingGET`: OrderBook
	fmt.Fprintf(os.Stdout, "Response from `SecurityOrderControllerAPI.GetOrderbookUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityIdentifier** | **string** | Security identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderbookUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrderBook**](OrderBook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwnedOrdersUsingGET

> []SecurityOrder GetOwnedOrdersUsingGET(ctx, securitiesAccountId).Execute()

Lists unfilled orders owned by securities account

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
	securitiesAccountId := "securitiesAccountId_example" // string | Securities account id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityOrderControllerAPI.GetOwnedOrdersUsingGET(context.Background(), securitiesAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityOrderControllerAPI.GetOwnedOrdersUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwnedOrdersUsingGET`: []SecurityOrder
	fmt.Fprintf(os.Stdout, "Response from `SecurityOrderControllerAPI.GetOwnedOrdersUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securitiesAccountId** | **string** | Securities account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnedOrdersUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SecurityOrder**](SecurityOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwnedOrdersV2UsingGET

> PageSecurityOrder GetOwnedOrdersV2UsingGET(ctx).Counterparty(counterparty).Page(page).SecuritiesAccountId(securitiesAccountId).Size(size).Sort(sort).Execute()

Lists unfilled orders

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
	counterparty := "counterparty_example" // string | Counterparty securities account id (optional)
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	securitiesAccountId := "securitiesAccountId_example" // string | Owning securities account id (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityOrderControllerAPI.GetOwnedOrdersV2UsingGET(context.Background()).Counterparty(counterparty).Page(page).SecuritiesAccountId(securitiesAccountId).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityOrderControllerAPI.GetOwnedOrdersV2UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwnedOrdersV2UsingGET`: PageSecurityOrder
	fmt.Fprintf(os.Stdout, "Response from `SecurityOrderControllerAPI.GetOwnedOrdersV2UsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnedOrdersV2UsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **counterparty** | **string** | Counterparty securities account id | 
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **securitiesAccountId** | **string** | Owning securities account id | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageSecurityOrder**](PageSecurityOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

