/*
Api Documentation

Api Documentation

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package atapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// SubscriptionControllerAPIService SubscriptionControllerAPI service
type SubscriptionControllerAPIService service

type ApiSubscribeToAuthorUsingPUTRequest struct {
	ctx context.Context
	ApiService *SubscriptionControllerAPIService
	action *string
	userId string
}

// Action
func (r ApiSubscribeToAuthorUsingPUTRequest) Action(action string) ApiSubscribeToAuthorUsingPUTRequest {
	r.action = &action
	return r
}

func (r ApiSubscribeToAuthorUsingPUTRequest) Execute() (*AuthorInterest, *http.Response, error) {
	return r.ApiService.SubscribeToAuthorUsingPUTExecute(r)
}

/*
SubscribeToAuthorUsingPUT Subscribes to/Ignores news from an author

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId userId
 @return ApiSubscribeToAuthorUsingPUTRequest
*/
func (a *SubscriptionControllerAPIService) SubscribeToAuthorUsingPUT(ctx context.Context, userId string) ApiSubscribeToAuthorUsingPUTRequest {
	return ApiSubscribeToAuthorUsingPUTRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return AuthorInterest
func (a *SubscriptionControllerAPIService) SubscribeToAuthorUsingPUTExecute(r ApiSubscribeToAuthorUsingPUTRequest) (*AuthorInterest, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthorInterest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionControllerAPIService.SubscribeToAuthorUsingPUT")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/my/subscriptions/authors/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.action == nil {
		return localVarReturnValue, nil, reportError("action is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "action", r.action, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSubscribeToCompanyUsingPUTRequest struct {
	ctx context.Context
	ApiService *SubscriptionControllerAPIService
	action *string
	companyId string
}

// Action
func (r ApiSubscribeToCompanyUsingPUTRequest) Action(action string) ApiSubscribeToCompanyUsingPUTRequest {
	r.action = &action
	return r
}

func (r ApiSubscribeToCompanyUsingPUTRequest) Execute() (*CompanyInterest, *http.Response, error) {
	return r.ApiService.SubscribeToCompanyUsingPUTExecute(r)
}

/*
SubscribeToCompanyUsingPUT Subscribes to/Ignores news from a company

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId companyId
 @return ApiSubscribeToCompanyUsingPUTRequest
*/
func (a *SubscriptionControllerAPIService) SubscribeToCompanyUsingPUT(ctx context.Context, companyId string) ApiSubscribeToCompanyUsingPUTRequest {
	return ApiSubscribeToCompanyUsingPUTRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return CompanyInterest
func (a *SubscriptionControllerAPIService) SubscribeToCompanyUsingPUTExecute(r ApiSubscribeToCompanyUsingPUTRequest) (*CompanyInterest, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CompanyInterest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionControllerAPIService.SubscribeToCompanyUsingPUT")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/my/subscriptions/companies/{companyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"companyId"+"}", url.PathEscape(parameterValueToString(r.companyId, "companyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.action == nil {
		return localVarReturnValue, nil, reportError("action is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "action", r.action, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSubscribeToHashtagUsingPUTRequest struct {
	ctx context.Context
	ApiService *SubscriptionControllerAPIService
	action *string
	hashtag string
}

// Action
func (r ApiSubscribeToHashtagUsingPUTRequest) Action(action string) ApiSubscribeToHashtagUsingPUTRequest {
	r.action = &action
	return r
}

func (r ApiSubscribeToHashtagUsingPUTRequest) Execute() (*HashTagInterest, *http.Response, error) {
	return r.ApiService.SubscribeToHashtagUsingPUTExecute(r)
}

/*
SubscribeToHashtagUsingPUT Subscribes to/Ignores news from a hashtag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param hashtag hashtag
 @return ApiSubscribeToHashtagUsingPUTRequest
*/
func (a *SubscriptionControllerAPIService) SubscribeToHashtagUsingPUT(ctx context.Context, hashtag string) ApiSubscribeToHashtagUsingPUTRequest {
	return ApiSubscribeToHashtagUsingPUTRequest{
		ApiService: a,
		ctx: ctx,
		hashtag: hashtag,
	}
}

// Execute executes the request
//  @return HashTagInterest
func (a *SubscriptionControllerAPIService) SubscribeToHashtagUsingPUTExecute(r ApiSubscribeToHashtagUsingPUTRequest) (*HashTagInterest, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HashTagInterest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionControllerAPIService.SubscribeToHashtagUsingPUT")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/my/subscriptions/hashtags/{hashtag}"
	localVarPath = strings.Replace(localVarPath, "{"+"hashtag"+"}", url.PathEscape(parameterValueToString(r.hashtag, "hashtag")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.action == nil {
		return localVarReturnValue, nil, reportError("action is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "action", r.action, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUnsubscribeFromAuthorUsingDELETERequest struct {
	ctx context.Context
	ApiService *SubscriptionControllerAPIService
	userId string
}

func (r ApiUnsubscribeFromAuthorUsingDELETERequest) Execute() (*MessagePrototype, *http.Response, error) {
	return r.ApiService.UnsubscribeFromAuthorUsingDELETEExecute(r)
}

/*
UnsubscribeFromAuthorUsingDELETE Unsubscribes news from an author

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId userId
 @return ApiUnsubscribeFromAuthorUsingDELETERequest
*/
func (a *SubscriptionControllerAPIService) UnsubscribeFromAuthorUsingDELETE(ctx context.Context, userId string) ApiUnsubscribeFromAuthorUsingDELETERequest {
	return ApiUnsubscribeFromAuthorUsingDELETERequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return MessagePrototype
func (a *SubscriptionControllerAPIService) UnsubscribeFromAuthorUsingDELETEExecute(r ApiUnsubscribeFromAuthorUsingDELETERequest) (*MessagePrototype, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MessagePrototype
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionControllerAPIService.UnsubscribeFromAuthorUsingDELETE")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/my/subscriptions/authors/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUnsubscribeFromCompanyUsingDELETERequest struct {
	ctx context.Context
	ApiService *SubscriptionControllerAPIService
	companyId string
}

func (r ApiUnsubscribeFromCompanyUsingDELETERequest) Execute() (*MessagePrototype, *http.Response, error) {
	return r.ApiService.UnsubscribeFromCompanyUsingDELETEExecute(r)
}

/*
UnsubscribeFromCompanyUsingDELETE Unsubscribes news from a company

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId companyId
 @return ApiUnsubscribeFromCompanyUsingDELETERequest
*/
func (a *SubscriptionControllerAPIService) UnsubscribeFromCompanyUsingDELETE(ctx context.Context, companyId string) ApiUnsubscribeFromCompanyUsingDELETERequest {
	return ApiUnsubscribeFromCompanyUsingDELETERequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return MessagePrototype
func (a *SubscriptionControllerAPIService) UnsubscribeFromCompanyUsingDELETEExecute(r ApiUnsubscribeFromCompanyUsingDELETERequest) (*MessagePrototype, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MessagePrototype
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionControllerAPIService.UnsubscribeFromCompanyUsingDELETE")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/my/subscriptions/companies/{companyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"companyId"+"}", url.PathEscape(parameterValueToString(r.companyId, "companyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUnsubscribeFromHashtagUsingDELETERequest struct {
	ctx context.Context
	ApiService *SubscriptionControllerAPIService
	hashtag string
}

func (r ApiUnsubscribeFromHashtagUsingDELETERequest) Execute() (*MessagePrototype, *http.Response, error) {
	return r.ApiService.UnsubscribeFromHashtagUsingDELETEExecute(r)
}

/*
UnsubscribeFromHashtagUsingDELETE Unsubscribes news from a hashtag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param hashtag hashtag
 @return ApiUnsubscribeFromHashtagUsingDELETERequest
*/
func (a *SubscriptionControllerAPIService) UnsubscribeFromHashtagUsingDELETE(ctx context.Context, hashtag string) ApiUnsubscribeFromHashtagUsingDELETERequest {
	return ApiUnsubscribeFromHashtagUsingDELETERequest{
		ApiService: a,
		ctx: ctx,
		hashtag: hashtag,
	}
}

// Execute executes the request
//  @return MessagePrototype
func (a *SubscriptionControllerAPIService) UnsubscribeFromHashtagUsingDELETEExecute(r ApiUnsubscribeFromHashtagUsingDELETERequest) (*MessagePrototype, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MessagePrototype
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionControllerAPIService.UnsubscribeFromHashtagUsingDELETE")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/my/subscriptions/hashtags/{hashtag}"
	localVarPath = strings.Replace(localVarPath, "{"+"hashtag"+"}", url.PathEscape(parameterValueToString(r.hashtag, "hashtag")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
