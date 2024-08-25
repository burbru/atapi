# \VotingControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CastAllVotesUsingPOST**](VotingControllerAPI.md#CastAllVotesUsingPOST) | **Post** /api/v2/polls | Casts a vote
[**CastVoteUsingPOST**](VotingControllerAPI.md#CastVoteUsingPOST) | **Post** /api/v2/polls/{pollId} | Casts a vote
[**CastVoteV1UsingPOST**](VotingControllerAPI.md#CastVoteV1UsingPOST) | **Post** /api/polls/{pollId} | Casts a vote
[**DeletePollUsingDELETE**](VotingControllerAPI.md#DeletePollUsingDELETE) | **Delete** /api/v2/polls/{pollId} | Deletes poll
[**DeletePollV1UsingDELETE**](VotingControllerAPI.md#DeletePollV1UsingDELETE) | **Delete** /api/polls/{pollId} | Deletes poll
[**ExecuteVoteResultUsingPUT**](VotingControllerAPI.md#ExecuteVoteResultUsingPUT) | **Put** /api/v2/polls/{pollId} | Executes poll result
[**ExecuteVoteResultV1UsingPOST**](VotingControllerAPI.md#ExecuteVoteResultV1UsingPOST) | **Post** /api/polls/execute/{pollId} | Executes poll result
[**GetAllPollsUsingGET**](VotingControllerAPI.md#GetAllPollsUsingGET) | **Get** /api/polls | Returns all polls
[**GetCompanyPollsUsingGET**](VotingControllerAPI.md#GetCompanyPollsUsingGET) | **Get** /api/v2/companies/{companyId}/polls | Returns company&#39;s polls
[**GetMyPollsPagedUsingGET**](VotingControllerAPI.md#GetMyPollsPagedUsingGET) | **Get** /api/v2/my/polls | Returns logged in user&#39;s polls
[**GetPollUsingGET**](VotingControllerAPI.md#GetPollUsingGET) | **Get** /api/v2/polls/{pollId} | Returns poll
[**GetPollV1UsingGET**](VotingControllerAPI.md#GetPollV1UsingGET) | **Get** /api/polls/{pollId} | Returns poll
[**GetPollsV1UsingGET**](VotingControllerAPI.md#GetPollsV1UsingGET) | **Get** /api/initiatedpolls/ | Returns logged in user&#39;s initiated polls
[**InitiateCapitalIncreasePollUsingPOST**](VotingControllerAPI.md#InitiateCapitalIncreasePollUsingPOST) | **Post** /api/v2/capitalincreasepolls | Initiate capital increase poll
[**InitiateCapitalReductionPollUsingPOST**](VotingControllerAPI.md#InitiateCapitalReductionPollUsingPOST) | **Post** /api/v2/capitalreductionpolls | Initiate capital reduction poll
[**InitiateCashOutPollUsingPOST**](VotingControllerAPI.md#InitiateCashOutPollUsingPOST) | **Post** /api/polls/cashout | Initiate cash out poll
[**InitiateCashOutPollUsingPOST1**](VotingControllerAPI.md#InitiateCashOutPollUsingPOST1) | **Post** /api/v2/cashoutpolls | Initiate cash out poll
[**InitiateChangeCompanyNamePollUsingPOST**](VotingControllerAPI.md#InitiateChangeCompanyNamePollUsingPOST) | **Post** /api/v2/changecompanynamepolls | Initiate change company name poll
[**InitiateDividendPaymentPollUsingPOST**](VotingControllerAPI.md#InitiateDividendPaymentPollUsingPOST) | **Post** /api/v2/dividendpaymentpolls | Initiate dividend payment poll
[**InitiateEmployCeoPollUsingPOST**](VotingControllerAPI.md#InitiateEmployCeoPollUsingPOST) | **Post** /api/v2/employceopolls | Initiate employ CEO poll
[**InitiateEmployCeoPollV1UsingPOST**](VotingControllerAPI.md#InitiateEmployCeoPollV1UsingPOST) | **Post** /api/polls/employceo | Initiate employ CEO poll
[**InitiateLiquidationPollUsingPOST**](VotingControllerAPI.md#InitiateLiquidationPollUsingPOST) | **Post** /api/v2/liquidationpolls | Initiate liquidation poll
[**InitiateLiquidationPollV1UsingPOST**](VotingControllerAPI.md#InitiateLiquidationPollV1UsingPOST) | **Post** /api/polls/liquidation | Initiate liquidation poll
[**InitiateMergerPollUsingPOST**](VotingControllerAPI.md#InitiateMergerPollUsingPOST) | **Post** /api/v2/mergerpolls | Initiate merger poll



## CastAllVotesUsingPOST

> AbstractPollView CastAllVotesUsingPOST(ctx).OnlyHarmless(onlyHarmless).VotingType(votingType).Execute()

Casts a vote

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
	onlyHarmless := true // bool | onlyHarmless
	votingType := "votingType_example" // string | Voting type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.CastAllVotesUsingPOST(context.Background()).OnlyHarmless(onlyHarmless).VotingType(votingType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.CastAllVotesUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CastAllVotesUsingPOST`: AbstractPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.CastAllVotesUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCastAllVotesUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onlyHarmless** | **bool** | onlyHarmless | 
 **votingType** | **string** | Voting type | 

### Return type

[**AbstractPollView**](AbstractPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CastVoteUsingPOST

> AbstractPollView CastVoteUsingPOST(ctx, pollId).Voices(voices).VotingType(votingType).Execute()

Casts a vote

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
	pollId := "pollId_example" // string | Poll id
	voices := int64(789) // int64 | Number of voices
	votingType := "votingType_example" // string | Voting type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.CastVoteUsingPOST(context.Background(), pollId).Voices(voices).VotingType(votingType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.CastVoteUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CastVoteUsingPOST`: AbstractPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.CastVoteUsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pollId** | **string** | Poll id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCastVoteUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **voices** | **int64** | Number of voices | 
 **votingType** | **string** | Voting type | 

### Return type

[**AbstractPollView**](AbstractPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CastVoteV1UsingPOST

> AbstractPollView CastVoteV1UsingPOST(ctx, pollId).Voices(voices).VotingType(votingType).Execute()

Casts a vote

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
	pollId := "pollId_example" // string | Poll id
	voices := int64(789) // int64 | Number of voices
	votingType := "votingType_example" // string | Voting type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.CastVoteV1UsingPOST(context.Background(), pollId).Voices(voices).VotingType(votingType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.CastVoteV1UsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CastVoteV1UsingPOST`: AbstractPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.CastVoteV1UsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pollId** | **string** | Poll id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCastVoteV1UsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **voices** | **int64** | Number of voices | 
 **votingType** | **string** | Voting type | 

### Return type

[**AbstractPollView**](AbstractPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePollUsingDELETE

> MessagePrototype DeletePollUsingDELETE(ctx, pollId).Execute()

Deletes poll

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
	pollId := "pollId_example" // string | Poll id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.DeletePollUsingDELETE(context.Background(), pollId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.DeletePollUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePollUsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.DeletePollUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pollId** | **string** | Poll id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePollUsingDELETERequest struct via the builder pattern


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


## DeletePollV1UsingDELETE

> MessagePrototype DeletePollV1UsingDELETE(ctx, pollId).Execute()

Deletes poll

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
	pollId := "pollId_example" // string | Poll id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.DeletePollV1UsingDELETE(context.Background(), pollId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.DeletePollV1UsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePollV1UsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.DeletePollV1UsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pollId** | **string** | Poll id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePollV1UsingDELETERequest struct via the builder pattern


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


## ExecuteVoteResultUsingPUT

> MessagePrototype ExecuteVoteResultUsingPUT(ctx, pollId).Execute()

Executes poll result

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
	pollId := "pollId_example" // string | Poll id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.ExecuteVoteResultUsingPUT(context.Background(), pollId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.ExecuteVoteResultUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteVoteResultUsingPUT`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.ExecuteVoteResultUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pollId** | **string** | Poll id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteVoteResultUsingPUTRequest struct via the builder pattern


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


## ExecuteVoteResultV1UsingPOST

> MessagePrototype ExecuteVoteResultV1UsingPOST(ctx, pollId).Execute()

Executes poll result

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
	pollId := "pollId_example" // string | Poll id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.ExecuteVoteResultV1UsingPOST(context.Background(), pollId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.ExecuteVoteResultV1UsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteVoteResultV1UsingPOST`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.ExecuteVoteResultV1UsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pollId** | **string** | Poll id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteVoteResultV1UsingPOSTRequest struct via the builder pattern


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


## GetAllPollsUsingGET

> AbstractPollView GetAllPollsUsingGET(ctx).Execute()

Returns all polls

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
	resp, r, err := apiClient.VotingControllerAPI.GetAllPollsUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.GetAllPollsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllPollsUsingGET`: AbstractPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.GetAllPollsUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPollsUsingGETRequest struct via the builder pattern


### Return type

[**AbstractPollView**](AbstractPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyPollsUsingGET

> AbstractPollView GetCompanyPollsUsingGET(ctx, companyId).Execute()

Returns company's polls

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
	companyId := "companyId_example" // string | Company id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.GetCompanyPollsUsingGET(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.GetCompanyPollsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPollsUsingGET`: AbstractPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.GetCompanyPollsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Company id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPollsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AbstractPollView**](AbstractPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyPollsPagedUsingGET

> AbstractPollView GetMyPollsPagedUsingGET(ctx).Page(page).SelfInitiated(selfInitiated).Size(size).Sort(sort).VotingStatus(votingStatus).Execute()

Returns logged in user's polls

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
	selfInitiated := true // bool | Initiated by logged in user (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)
	votingStatus := "votingStatus_example" // string | Status of own voting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.GetMyPollsPagedUsingGET(context.Background()).Page(page).SelfInitiated(selfInitiated).Size(size).Sort(sort).VotingStatus(votingStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.GetMyPollsPagedUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyPollsPagedUsingGET`: AbstractPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.GetMyPollsPagedUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyPollsPagedUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **selfInitiated** | **bool** | Initiated by logged in user | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **votingStatus** | **string** | Status of own voting | 

### Return type

[**AbstractPollView**](AbstractPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPollUsingGET

> AbstractPollView GetPollUsingGET(ctx, pollId).Execute()

Returns poll

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
	pollId := "pollId_example" // string | Poll id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.GetPollUsingGET(context.Background(), pollId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.GetPollUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPollUsingGET`: AbstractPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.GetPollUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pollId** | **string** | Poll id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPollUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AbstractPollView**](AbstractPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPollV1UsingGET

> AbstractPollView GetPollV1UsingGET(ctx, pollId).Execute()

Returns poll

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
	pollId := "pollId_example" // string | Poll id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.GetPollV1UsingGET(context.Background(), pollId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.GetPollV1UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPollV1UsingGET`: AbstractPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.GetPollV1UsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pollId** | **string** | Poll id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPollV1UsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AbstractPollView**](AbstractPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPollsV1UsingGET

> AbstractPollView GetPollsV1UsingGET(ctx).Execute()

Returns logged in user's initiated polls

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
	resp, r, err := apiClient.VotingControllerAPI.GetPollsV1UsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.GetPollsV1UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPollsV1UsingGET`: AbstractPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.GetPollsV1UsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPollsV1UsingGETRequest struct via the builder pattern


### Return type

[**AbstractPollView**](AbstractPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateCapitalIncreasePollUsingPOST

> CapitalIncreasePollView InitiateCapitalIncreasePollUsingPOST(ctx).CapitalIncreaseType(capitalIncreaseType).CompanyId(companyId).MinimalCashVolume(minimalCashVolume).Price(price).Execute()

Initiate capital increase poll

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
	capitalIncreaseType := "capitalIncreaseType_example" // string | Capital increase type
	companyId := "companyId_example" // string | Company id
	minimalCashVolume := "minimalCashVolume_example" // string | Minimal cash volume
	price := "price_example" // string | Price

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.InitiateCapitalIncreasePollUsingPOST(context.Background()).CapitalIncreaseType(capitalIncreaseType).CompanyId(companyId).MinimalCashVolume(minimalCashVolume).Price(price).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.InitiateCapitalIncreasePollUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateCapitalIncreasePollUsingPOST`: CapitalIncreasePollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.InitiateCapitalIncreasePollUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateCapitalIncreasePollUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **capitalIncreaseType** | **string** | Capital increase type | 
 **companyId** | **string** | Company id | 
 **minimalCashVolume** | **string** | Minimal cash volume | 
 **price** | **string** | Price | 

### Return type

[**CapitalIncreasePollView**](CapitalIncreasePollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateCapitalReductionPollUsingPOST

> CapitalReductionPollView InitiateCapitalReductionPollUsingPOST(ctx).CompanyId(companyId).NumberOfShares(numberOfShares).Price(price).Execute()

Initiate capital reduction poll

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
	companyId := "companyId_example" // string | Company id
	numberOfShares := int64(789) // int64 | Number of shares
	price := "price_example" // string | Price

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.InitiateCapitalReductionPollUsingPOST(context.Background()).CompanyId(companyId).NumberOfShares(numberOfShares).Price(price).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.InitiateCapitalReductionPollUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateCapitalReductionPollUsingPOST`: CapitalReductionPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.InitiateCapitalReductionPollUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateCapitalReductionPollUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company id | 
 **numberOfShares** | **int64** | Number of shares | 
 **price** | **string** | Price | 

### Return type

[**CapitalReductionPollView**](CapitalReductionPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateCashOutPollUsingPOST

> ShareholderYesNoPollView InitiateCashOutPollUsingPOST(ctx).CompanyId(companyId).Execute()

Initiate cash out poll

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
	companyId := "companyId_example" // string | Company id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.InitiateCashOutPollUsingPOST(context.Background()).CompanyId(companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.InitiateCashOutPollUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateCashOutPollUsingPOST`: ShareholderYesNoPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.InitiateCashOutPollUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateCashOutPollUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company id | 

### Return type

[**ShareholderYesNoPollView**](ShareholderYesNoPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateCashOutPollUsingPOST1

> ShareholderYesNoPollView InitiateCashOutPollUsingPOST1(ctx).CompanyId(companyId).Execute()

Initiate cash out poll

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
	companyId := "companyId_example" // string | Company id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.InitiateCashOutPollUsingPOST1(context.Background()).CompanyId(companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.InitiateCashOutPollUsingPOST1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateCashOutPollUsingPOST1`: ShareholderYesNoPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.InitiateCashOutPollUsingPOST1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateCashOutPollUsingPOST1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company id | 

### Return type

[**ShareholderYesNoPollView**](ShareholderYesNoPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateChangeCompanyNamePollUsingPOST

> EmployCeoPollView InitiateChangeCompanyNamePollUsingPOST(ctx).CompanyId(companyId).NewName(newName).Execute()

Initiate change company name poll

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
	companyId := "companyId_example" // string | Company id
	newName := "newName_example" // string | New Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.InitiateChangeCompanyNamePollUsingPOST(context.Background()).CompanyId(companyId).NewName(newName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.InitiateChangeCompanyNamePollUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateChangeCompanyNamePollUsingPOST`: EmployCeoPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.InitiateChangeCompanyNamePollUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateChangeCompanyNamePollUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company id | 
 **newName** | **string** | New Name | 

### Return type

[**EmployCeoPollView**](EmployCeoPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateDividendPaymentPollUsingPOST

> DividendPaymentPollView InitiateDividendPaymentPollUsingPOST(ctx).CompanyId(companyId).MaximalCashVolume(maximalCashVolume).Execute()

Initiate dividend payment poll

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
	companyId := "companyId_example" // string | Company id
	maximalCashVolume := "maximalCashVolume_example" // string | Maximal cash volume

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.InitiateDividendPaymentPollUsingPOST(context.Background()).CompanyId(companyId).MaximalCashVolume(maximalCashVolume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.InitiateDividendPaymentPollUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateDividendPaymentPollUsingPOST`: DividendPaymentPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.InitiateDividendPaymentPollUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateDividendPaymentPollUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company id | 
 **maximalCashVolume** | **string** | Maximal cash volume | 

### Return type

[**DividendPaymentPollView**](DividendPaymentPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateEmployCeoPollUsingPOST

> EmployCeoPollView InitiateEmployCeoPollUsingPOST(ctx).CompanyId(companyId).DailyWage(dailyWage).Execute()

Initiate employ CEO poll

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
	companyId := "companyId_example" // string | Company id
	dailyWage := "dailyWage_example" // string | Requested daily wage

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.InitiateEmployCeoPollUsingPOST(context.Background()).CompanyId(companyId).DailyWage(dailyWage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.InitiateEmployCeoPollUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateEmployCeoPollUsingPOST`: EmployCeoPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.InitiateEmployCeoPollUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateEmployCeoPollUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company id | 
 **dailyWage** | **string** | Requested daily wage | 

### Return type

[**EmployCeoPollView**](EmployCeoPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateEmployCeoPollV1UsingPOST

> EmployCeoPollView InitiateEmployCeoPollV1UsingPOST(ctx).CompanyId(companyId).DailyWage(dailyWage).Execute()

Initiate employ CEO poll

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
	companyId := "companyId_example" // string | Company id
	dailyWage := "dailyWage_example" // string | Requested daily wage

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.InitiateEmployCeoPollV1UsingPOST(context.Background()).CompanyId(companyId).DailyWage(dailyWage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.InitiateEmployCeoPollV1UsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateEmployCeoPollV1UsingPOST`: EmployCeoPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.InitiateEmployCeoPollV1UsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateEmployCeoPollV1UsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company id | 
 **dailyWage** | **string** | Requested daily wage | 

### Return type

[**EmployCeoPollView**](EmployCeoPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateLiquidationPollUsingPOST

> ShareholderYesNoPollView InitiateLiquidationPollUsingPOST(ctx).CompanyId(companyId).Execute()

Initiate liquidation poll

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
	companyId := "companyId_example" // string | Company id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.InitiateLiquidationPollUsingPOST(context.Background()).CompanyId(companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.InitiateLiquidationPollUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateLiquidationPollUsingPOST`: ShareholderYesNoPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.InitiateLiquidationPollUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateLiquidationPollUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company id | 

### Return type

[**ShareholderYesNoPollView**](ShareholderYesNoPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateLiquidationPollV1UsingPOST

> ShareholderYesNoPollView InitiateLiquidationPollV1UsingPOST(ctx).CompanyId(companyId).Execute()

Initiate liquidation poll

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
	companyId := "companyId_example" // string | Company id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.InitiateLiquidationPollV1UsingPOST(context.Background()).CompanyId(companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.InitiateLiquidationPollV1UsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateLiquidationPollV1UsingPOST`: ShareholderYesNoPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.InitiateLiquidationPollV1UsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateLiquidationPollV1UsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company id | 

### Return type

[**ShareholderYesNoPollView**](ShareholderYesNoPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateMergerPollUsingPOST

> DividendPaymentPollView InitiateMergerPollUsingPOST(ctx).AcquiringCompanyId(acquiringCompanyId).CompanyId(companyId).MaximalCashVolume(maximalCashVolume).Execute()

Initiate merger poll

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
	acquiringCompanyId := "acquiringCompanyId_example" // string | Acquiring company id
	companyId := "companyId_example" // string | Company id
	maximalCashVolume := "maximalCashVolume_example" // string | Maximal cash volume

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VotingControllerAPI.InitiateMergerPollUsingPOST(context.Background()).AcquiringCompanyId(acquiringCompanyId).CompanyId(companyId).MaximalCashVolume(maximalCashVolume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VotingControllerAPI.InitiateMergerPollUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateMergerPollUsingPOST`: DividendPaymentPollView
	fmt.Fprintf(os.Stdout, "Response from `VotingControllerAPI.InitiateMergerPollUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateMergerPollUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acquiringCompanyId** | **string** | Acquiring company id | 
 **companyId** | **string** | Company id | 
 **maximalCashVolume** | **string** | Maximal cash volume | 

### Return type

[**DividendPaymentPollView**](DividendPaymentPollView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

