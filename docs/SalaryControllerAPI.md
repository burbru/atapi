# \SalaryControllerAPI

All URIs are relative to *https://stable.alpha-trader.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeAgreementUsingPATCH**](SalaryControllerAPI.md#ChangeAgreementUsingPATCH) | **Patch** /api/v2/employmentagreements/company/{companyId} | Changes automatic payment
[**DeleteAgreementV1UsingDELETE**](SalaryControllerAPI.md#DeleteAgreementV1UsingDELETE) | **Delete** /api/v2/employmentagreements/{agreementId} | Delete employment agreement
[**GetAgreementV1UsingGET**](SalaryControllerAPI.md#GetAgreementV1UsingGET) | **Get** /api/employmentagreements/company/{companyId} | Returns employment agreement
[**GetAgreementsUsingGET**](SalaryControllerAPI.md#GetAgreementsUsingGET) | **Get** /api/v2/employmentagreements | Returns employment agreements
[**GetAgreementsV1UsingGET**](SalaryControllerAPI.md#GetAgreementsV1UsingGET) | **Get** /api/employmentagreements/ | Lists logged in user&#39;s employment agreements
[**GetMyAgreementsUsingGET**](SalaryControllerAPI.md#GetMyAgreementsUsingGET) | **Get** /api/v2/my/employmentagreements | Lists logged in user&#39;s employment agreements
[**GetMyPossibleDailySalaryUsingGET**](SalaryControllerAPI.md#GetMyPossibleDailySalaryUsingGET) | **Get** /api/v2/my/possibledailysalary | Returns logged in user&#39;s possible daily salary
[**GetPaymentUsingGET**](SalaryControllerAPI.md#GetPaymentUsingGET) | **Get** /api/salarypayments/{paymentId} | Returns payment
[**GetPossibleDailySalaryUsingGET**](SalaryControllerAPI.md#GetPossibleDailySalaryUsingGET) | **Get** /api/v2/possibledailysalary/{userId} | Returns user&#39;s possible daily salary
[**PayCeoUsingPOST**](SalaryControllerAPI.md#PayCeoUsingPOST) | **Post** /api/salarypayments/company/{companyId} | Pays CEO
[**PaySalariesUsingPUT**](SalaryControllerAPI.md#PaySalariesUsingPUT) | **Put** /api/v2/my/salarypayments | Pay all possible salaries



## ChangeAgreementUsingPATCH

> EmploymentAgreementView ChangeAgreementUsingPATCH(ctx, companyId).PayAutomatically(payAutomatically).Execute()

Changes automatic payment

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
	payAutomatically := true // bool | Pay automatically

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalaryControllerAPI.ChangeAgreementUsingPATCH(context.Background(), companyId).PayAutomatically(payAutomatically).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalaryControllerAPI.ChangeAgreementUsingPATCH``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeAgreementUsingPATCH`: EmploymentAgreementView
	fmt.Fprintf(os.Stdout, "Response from `SalaryControllerAPI.ChangeAgreementUsingPATCH`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Company id | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeAgreementUsingPATCHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payAutomatically** | **bool** | Pay automatically | 

### Return type

[**EmploymentAgreementView**](EmploymentAgreementView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgreementV1UsingDELETE

> MessagePrototype DeleteAgreementV1UsingDELETE(ctx, agreementId).Execute()

Delete employment agreement

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
	agreementId := "agreementId_example" // string | Eployment agreement id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalaryControllerAPI.DeleteAgreementV1UsingDELETE(context.Background(), agreementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalaryControllerAPI.DeleteAgreementV1UsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAgreementV1UsingDELETE`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `SalaryControllerAPI.DeleteAgreementV1UsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | Eployment agreement id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgreementV1UsingDELETERequest struct via the builder pattern


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


## GetAgreementV1UsingGET

> EmploymentAgreementView GetAgreementV1UsingGET(ctx, companyId).Execute()

Returns employment agreement

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
	resp, r, err := apiClient.SalaryControllerAPI.GetAgreementV1UsingGET(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalaryControllerAPI.GetAgreementV1UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgreementV1UsingGET`: EmploymentAgreementView
	fmt.Fprintf(os.Stdout, "Response from `SalaryControllerAPI.GetAgreementV1UsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Company id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgreementV1UsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmploymentAgreementView**](EmploymentAgreementView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgreementsUsingGET

> EmploymentAgreementSalaryPaymentCompactCompanyView GetAgreementsUsingGET(ctx).CompanyId(companyId).Page(page).Size(size).Sort(sort).UserId(userId).Execute()

Returns employment agreements

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
	companyId := "companyId_example" // string | Company id (optional)
	page := int32(56) // int32 | Results page you want to retrieve (0..N) (optional)
	size := int32(56) // int32 | Number of records per page (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)
	userId := "userId_example" // string | User id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalaryControllerAPI.GetAgreementsUsingGET(context.Background()).CompanyId(companyId).Page(page).Size(size).Sort(sort).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalaryControllerAPI.GetAgreementsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgreementsUsingGET`: EmploymentAgreementSalaryPaymentCompactCompanyView
	fmt.Fprintf(os.Stdout, "Response from `SalaryControllerAPI.GetAgreementsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgreementsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Company id | 
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **userId** | **string** | User id | 

### Return type

[**EmploymentAgreementSalaryPaymentCompactCompanyView**](EmploymentAgreementSalaryPaymentCompactCompanyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgreementsV1UsingGET

> CompanyEmploymentAgreementView GetAgreementsV1UsingGET(ctx).Execute()

Lists logged in user's employment agreements

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
	resp, r, err := apiClient.SalaryControllerAPI.GetAgreementsV1UsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalaryControllerAPI.GetAgreementsV1UsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgreementsV1UsingGET`: CompanyEmploymentAgreementView
	fmt.Fprintf(os.Stdout, "Response from `SalaryControllerAPI.GetAgreementsV1UsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgreementsV1UsingGETRequest struct via the builder pattern


### Return type

[**CompanyEmploymentAgreementView**](CompanyEmploymentAgreementView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyAgreementsUsingGET

> EmploymentAgreementSalaryPaymentCompactCompanyView GetMyAgreementsUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Lists logged in user's employment agreements

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
	resp, r, err := apiClient.SalaryControllerAPI.GetMyAgreementsUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalaryControllerAPI.GetMyAgreementsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyAgreementsUsingGET`: EmploymentAgreementSalaryPaymentCompactCompanyView
	fmt.Fprintf(os.Stdout, "Response from `SalaryControllerAPI.GetMyAgreementsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyAgreementsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results page you want to retrieve (0..N) | 
 **size** | **int32** | Number of records per page | 
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**EmploymentAgreementSalaryPaymentCompactCompanyView**](EmploymentAgreementSalaryPaymentCompactCompanyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyPossibleDailySalaryUsingGET

> PossibleDailySalaryView GetMyPossibleDailySalaryUsingGET(ctx).Execute()

Returns logged in user's possible daily salary

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
	resp, r, err := apiClient.SalaryControllerAPI.GetMyPossibleDailySalaryUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalaryControllerAPI.GetMyPossibleDailySalaryUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyPossibleDailySalaryUsingGET`: PossibleDailySalaryView
	fmt.Fprintf(os.Stdout, "Response from `SalaryControllerAPI.GetMyPossibleDailySalaryUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyPossibleDailySalaryUsingGETRequest struct via the builder pattern


### Return type

[**PossibleDailySalaryView**](PossibleDailySalaryView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentUsingGET

> SalaryPaymentView GetPaymentUsingGET(ctx, paymentId).Execute()

Returns payment

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
	paymentId := "paymentId_example" // string | Payment id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalaryControllerAPI.GetPaymentUsingGET(context.Background(), paymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalaryControllerAPI.GetPaymentUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentUsingGET`: SalaryPaymentView
	fmt.Fprintf(os.Stdout, "Response from `SalaryControllerAPI.GetPaymentUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | **string** | Payment id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SalaryPaymentView**](SalaryPaymentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPossibleDailySalaryUsingGET

> PossibleDailySalaryView GetPossibleDailySalaryUsingGET(ctx, userId).Execute()

Returns user's possible daily salary

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
	userId := "userId_example" // string | User id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalaryControllerAPI.GetPossibleDailySalaryUsingGET(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalaryControllerAPI.GetPossibleDailySalaryUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPossibleDailySalaryUsingGET`: PossibleDailySalaryView
	fmt.Fprintf(os.Stdout, "Response from `SalaryControllerAPI.GetPossibleDailySalaryUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPossibleDailySalaryUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PossibleDailySalaryView**](PossibleDailySalaryView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayCeoUsingPOST

> SalaryPaymentView PayCeoUsingPOST(ctx, companyId).Execute()

Pays CEO

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
	resp, r, err := apiClient.SalaryControllerAPI.PayCeoUsingPOST(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalaryControllerAPI.PayCeoUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PayCeoUsingPOST`: SalaryPaymentView
	fmt.Fprintf(os.Stdout, "Response from `SalaryControllerAPI.PayCeoUsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Company id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPayCeoUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SalaryPaymentView**](SalaryPaymentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaySalariesUsingPUT

> MessagePrototype PaySalariesUsingPUT(ctx).Execute()

Pay all possible salaries

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
	resp, r, err := apiClient.SalaryControllerAPI.PaySalariesUsingPUT(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalaryControllerAPI.PaySalariesUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaySalariesUsingPUT`: MessagePrototype
	fmt.Fprintf(os.Stdout, "Response from `SalaryControllerAPI.PaySalariesUsingPUT`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaySalariesUsingPUTRequest struct via the builder pattern


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

