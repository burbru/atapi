# \AtSpielControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAtsxCompactInfosUsingGET**](AtSpielControllerAPI.md#GetAtsxCompactInfosUsingGET) | **Get** /ats/atsxcompact | Liefert ATSX-Daten für rechtes Menü
[**GetBiggestTradedStocksUsingGET**](AtSpielControllerAPI.md#GetBiggestTradedStocksUsingGET) | **Get** /ats/biggesttradedstocks | Listet Aktien mit dem größten Handelsvolumen (24h) für rechtes Menü
[**GetBondsInterestUsingGET**](AtSpielControllerAPI.md#GetBondsInterestUsingGET) | **Get** /ats/bondsinterest | Liefert Anleihenzins für rechtes Menü
[**GetCompanyOfTheDayUsingGET**](AtSpielControllerAPI.md#GetCompanyOfTheDayUsingGET) | **Get** /ats/companyoftheday | Liefert AG des Tages
[**GetIndexChatUsingGET**](AtSpielControllerAPI.md#GetIndexChatUsingGET) | **Get** /ats/indexchat | Liefert Index-Chat
[**GetLiveDataUsingGET**](AtSpielControllerAPI.md#GetLiveDataUsingGET) | **Get** /ats/livedata | Liefert Live-Daten für Live-Börse
[**GetLobbyChatDeUsingGET**](AtSpielControllerAPI.md#GetLobbyChatDeUsingGET) | **Get** /ats/lobbychat | Liefert Lobby-Chat
[**GetPriceLoserUsingGET**](AtSpielControllerAPI.md#GetPriceLoserUsingGET) | **Get** /ats/priceloser | Liefert Aktie mit größtem Kursabstieg im Depot
[**GetPriceWinnerUsingGET**](AtSpielControllerAPI.md#GetPriceWinnerUsingGET) | **Get** /ats/pricewinner | Liefert Aktie mit größtem Kursanstieg im Depot
[**GetTopPremiumUserUsingGET**](AtSpielControllerAPI.md#GetTopPremiumUserUsingGET) | **Get** /ats/toppremium | Liefert Top-Premium-Spieler für die Zeitung links oben
[**ListAnniversariesUsingGET**](AtSpielControllerAPI.md#ListAnniversariesUsingGET) | **Get** /ats/anniversaries | Listet Jubiläen für Zeitung
[**ListAtsCommentsUsingGET**](AtSpielControllerAPI.md#ListAtsCommentsUsingGET) | **Get** /ats/comments/{postId} | Listet Kommentare
[**ListAtsLikesUsingGET**](AtSpielControllerAPI.md#ListAtsLikesUsingGET) | **Get** /ats/likes/{postId} | Listet Likes und Dislikes von einem Artikel
[**ListAtsNewsUsingGET**](AtSpielControllerAPI.md#ListAtsNewsUsingGET) | **Get** /ats/news | Listet News für Newsticker oben und Zeitung
[**ListAtsTradeNewsUsingGET**](AtSpielControllerAPI.md#ListAtsTradeNewsUsingGET) | **Get** /ats/tradenews | Listet Trades für Newsticker oben
[**ListCompaniesUsingGET**](AtSpielControllerAPI.md#ListCompaniesUsingGET) | **Get** /ats/companylist | Listet AGs für AG-Liste auf
[**ListCompanyLogosUsingGET**](AtSpielControllerAPI.md#ListCompanyLogosUsingGET) | **Get** /ats/companylogos | Listet Aktien-Börsennotierungen mit Bild für AG-Wand auf
[**ListHotAtsNewsUsingGET**](AtSpielControllerAPI.md#ListHotAtsNewsUsingGET) | **Get** /ats/hotnews | Listet heißdiskutierte News Zeitung
[**ListNewStocksUsingGET**](AtSpielControllerAPI.md#ListNewStocksUsingGET) | **Get** /ats/newstocks | Listet neue Aktien für rechtes Menü
[**ListOnlineUsersUsingGET**](AtSpielControllerAPI.md#ListOnlineUsersUsingGET) | **Get** /ats/onlineusers | Listet online User für Menü oben



## GetAtsxCompactInfosUsingGET

> IndexComparisonView GetAtsxCompactInfosUsingGET(ctx).Execute()

Liefert ATSX-Daten für rechtes Menü

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
	resp, r, err := apiClient.AtSpielControllerAPI.GetAtsxCompactInfosUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.GetAtsxCompactInfosUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAtsxCompactInfosUsingGET`: IndexComparisonView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.GetAtsxCompactInfosUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAtsxCompactInfosUsingGETRequest struct via the builder pattern


### Return type

[**IndexComparisonView**](IndexComparisonView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBiggestTradedStocksUsingGET

> ListingWithTradingCountView GetBiggestTradedStocksUsingGET(ctx).Execute()

Listet Aktien mit dem größten Handelsvolumen (24h) für rechtes Menü

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
	resp, r, err := apiClient.AtSpielControllerAPI.GetBiggestTradedStocksUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.GetBiggestTradedStocksUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBiggestTradedStocksUsingGET`: ListingWithTradingCountView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.GetBiggestTradedStocksUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBiggestTradedStocksUsingGETRequest struct via the builder pattern


### Return type

[**ListingWithTradingCountView**](ListingWithTradingCountView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBondsInterestUsingGET

> AverageDailyBondInterestRateView GetBondsInterestUsingGET(ctx).Execute()

Liefert Anleihenzins für rechtes Menü

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
	resp, r, err := apiClient.AtSpielControllerAPI.GetBondsInterestUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.GetBondsInterestUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBondsInterestUsingGET`: AverageDailyBondInterestRateView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.GetBondsInterestUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBondsInterestUsingGETRequest struct via the builder pattern


### Return type

[**AverageDailyBondInterestRateView**](AverageDailyBondInterestRateView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyOfTheDayUsingGET

> CompactCompanyView GetCompanyOfTheDayUsingGET(ctx).Execute()

Liefert AG des Tages

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
	resp, r, err := apiClient.AtSpielControllerAPI.GetCompanyOfTheDayUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.GetCompanyOfTheDayUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyOfTheDayUsingGET`: CompactCompanyView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.GetCompanyOfTheDayUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyOfTheDayUsingGETRequest struct via the builder pattern


### Return type

[**CompactCompanyView**](CompactCompanyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexChatUsingGET

> ChatView GetIndexChatUsingGET(ctx).Execute()

Liefert Index-Chat

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
	resp, r, err := apiClient.AtSpielControllerAPI.GetIndexChatUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.GetIndexChatUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexChatUsingGET`: ChatView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.GetIndexChatUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexChatUsingGETRequest struct via the builder pattern


### Return type

[**ChatView**](ChatView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveDataUsingGET

> AtsLiveDataView GetLiveDataUsingGET(ctx).OrderLogStartDate(orderLogStartDate).OrderStartDate(orderStartDate).Execute()

Liefert Live-Daten für Live-Börse

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
	orderLogStartDate := int64(789) // int64 | Order Log Start Date
	orderStartDate := int64(789) // int64 | Order Start Date

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AtSpielControllerAPI.GetLiveDataUsingGET(context.Background()).OrderLogStartDate(orderLogStartDate).OrderStartDate(orderStartDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.GetLiveDataUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLiveDataUsingGET`: AtsLiveDataView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.GetLiveDataUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLiveDataUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderLogStartDate** | **int64** | Order Log Start Date | 
 **orderStartDate** | **int64** | Order Start Date | 

### Return type

[**AtsLiveDataView**](AtsLiveDataView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLobbyChatDeUsingGET

> ChatView GetLobbyChatDeUsingGET(ctx).Execute()

Liefert Lobby-Chat

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
	resp, r, err := apiClient.AtSpielControllerAPI.GetLobbyChatDeUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.GetLobbyChatDeUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLobbyChatDeUsingGET`: ChatView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.GetLobbyChatDeUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLobbyChatDeUsingGETRequest struct via the builder pattern


### Return type

[**ChatView**](ChatView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceLoserUsingGET

> ListingWithPriceChangeView GetPriceLoserUsingGET(ctx).CompanyId(companyId).Execute()

Liefert Aktie mit größtem Kursabstieg im Depot

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
	companyId := "companyId_example" // string | Company ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AtSpielControllerAPI.GetPriceLoserUsingGET(context.Background()).CompanyId(companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.GetPriceLoserUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriceLoserUsingGET`: ListingWithPriceChangeView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.GetPriceLoserUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceLoserUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company ID | 

### Return type

[**ListingWithPriceChangeView**](ListingWithPriceChangeView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceWinnerUsingGET

> ListingWithPriceChangeView GetPriceWinnerUsingGET(ctx).CompanyId(companyId).Execute()

Liefert Aktie mit größtem Kursanstieg im Depot

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
	companyId := "companyId_example" // string | Company ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AtSpielControllerAPI.GetPriceWinnerUsingGET(context.Background()).CompanyId(companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.GetPriceWinnerUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriceWinnerUsingGET`: ListingWithPriceChangeView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.GetPriceWinnerUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceWinnerUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company ID | 

### Return type

[**ListingWithPriceChangeView**](ListingWithPriceChangeView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopPremiumUserUsingGET

> UsernameView GetTopPremiumUserUsingGET(ctx).Execute()

Liefert Top-Premium-Spieler für die Zeitung links oben

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
	resp, r, err := apiClient.AtSpielControllerAPI.GetTopPremiumUserUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.GetTopPremiumUserUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopPremiumUserUsingGET`: UsernameView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.GetTopPremiumUserUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopPremiumUserUsingGETRequest struct via the builder pattern


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


## ListAnniversariesUsingGET

> []UsernameView ListAnniversariesUsingGET(ctx).Execute()

Listet Jubiläen für Zeitung

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
	resp, r, err := apiClient.AtSpielControllerAPI.ListAnniversariesUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.ListAnniversariesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAnniversariesUsingGET`: []UsernameView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.ListAnniversariesUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAnniversariesUsingGETRequest struct via the builder pattern


### Return type

[**[]UsernameView**](UsernameView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAtsCommentsUsingGET

> []PostView ListAtsCommentsUsingGET(ctx, postId).Execute()

Listet Kommentare

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
	postId := "postId_example" // string | postId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AtSpielControllerAPI.ListAtsCommentsUsingGET(context.Background(), postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.ListAtsCommentsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAtsCommentsUsingGET`: []PostView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.ListAtsCommentsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** | postId | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAtsCommentsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PostView**](PostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAtsLikesUsingGET

> PostLikeView ListAtsLikesUsingGET(ctx, postId).Execute()

Listet Likes und Dislikes von einem Artikel

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
	postId := "postId_example" // string | postId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AtSpielControllerAPI.ListAtsLikesUsingGET(context.Background(), postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.ListAtsLikesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAtsLikesUsingGET`: PostLikeView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.ListAtsLikesUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** | postId | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAtsLikesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostLikeView**](PostLikeView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAtsNewsUsingGET

> PagePostView ListAtsNewsUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Listet News für Newsticker oben und Zeitung

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AtSpielControllerAPI.ListAtsNewsUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.ListAtsNewsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAtsNewsUsingGET`: PagePostView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.ListAtsNewsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAtsNewsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PagePostView**](PagePostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAtsTradeNewsUsingGET

> PageSecurityOrderLogEntryWithPriceDiffView ListAtsTradeNewsUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Listet Trades für Newsticker oben

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AtSpielControllerAPI.ListAtsTradeNewsUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.ListAtsTradeNewsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAtsTradeNewsUsingGET`: PageSecurityOrderLogEntryWithPriceDiffView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.ListAtsTradeNewsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAtsTradeNewsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageSecurityOrderLogEntryWithPriceDiffView**](PageSecurityOrderLogEntryWithPriceDiffView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompaniesUsingGET

> []AsinWithNameView ListCompaniesUsingGET(ctx).Execute()

Listet AGs für AG-Liste auf

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
	resp, r, err := apiClient.AtSpielControllerAPI.ListCompaniesUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.ListCompaniesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompaniesUsingGET`: []AsinWithNameView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.ListCompaniesUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCompaniesUsingGETRequest struct via the builder pattern


### Return type

[**[]AsinWithNameView**](AsinWithNameView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompanyLogosUsingGET

> []AsinWithLogoUrlView ListCompanyLogosUsingGET(ctx).Execute()

Listet Aktien-Börsennotierungen mit Bild für AG-Wand auf

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
	resp, r, err := apiClient.AtSpielControllerAPI.ListCompanyLogosUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.ListCompanyLogosUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompanyLogosUsingGET`: []AsinWithLogoUrlView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.ListCompanyLogosUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCompanyLogosUsingGETRequest struct via the builder pattern


### Return type

[**[]AsinWithLogoUrlView**](AsinWithLogoUrlView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHotAtsNewsUsingGET

> PagePostView ListHotAtsNewsUsingGET(ctx).StartDate(startDate).Page(page).Size(size).Sort(sort).Execute()

Listet heißdiskutierte News Zeitung

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
	startDate := int64(789) // int64 | Startdatum
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AtSpielControllerAPI.ListHotAtsNewsUsingGET(context.Background()).StartDate(startDate).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.ListHotAtsNewsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHotAtsNewsUsingGET`: PagePostView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.ListHotAtsNewsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHotAtsNewsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **int64** | Startdatum | 
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PagePostView**](PagePostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNewStocksUsingGET

> []NewStockView ListNewStocksUsingGET(ctx).Execute()

Listet neue Aktien für rechtes Menü

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
	resp, r, err := apiClient.AtSpielControllerAPI.ListNewStocksUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.ListNewStocksUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNewStocksUsingGET`: []NewStockView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.ListNewStocksUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListNewStocksUsingGETRequest struct via the builder pattern


### Return type

[**[]NewStockView**](NewStockView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOnlineUsersUsingGET

> []UsernameView ListOnlineUsersUsingGET(ctx).Execute()

Listet online User für Menü oben

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
	resp, r, err := apiClient.AtSpielControllerAPI.ListOnlineUsersUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AtSpielControllerAPI.ListOnlineUsersUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOnlineUsersUsingGET`: []UsernameView
	fmt.Fprintf(os.Stdout, "Response from `AtSpielControllerAPI.ListOnlineUsersUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOnlineUsersUsingGETRequest struct via the builder pattern


### Return type

[**[]UsernameView**](UsernameView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

