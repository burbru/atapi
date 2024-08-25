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
	"reflect"
)


// EventControllerAPIService EventControllerAPI service
type EventControllerAPIService service

type ApiFindEventsByEventTypeUsingGETRequest struct {
	ctx context.Context
	ApiService *EventControllerAPIService
	eventType string
	afterDate *int64
	realms *[]string
}

// After date
func (r ApiFindEventsByEventTypeUsingGETRequest) AfterDate(afterDate int64) ApiFindEventsByEventTypeUsingGETRequest {
	r.afterDate = &afterDate
	return r
}

// Realms
func (r ApiFindEventsByEventTypeUsingGETRequest) Realms(realms []string) ApiFindEventsByEventTypeUsingGETRequest {
	r.realms = &realms
	return r
}

func (r ApiFindEventsByEventTypeUsingGETRequest) Execute() (*Event, *http.Response, error) {
	return r.ApiService.FindEventsByEventTypeUsingGETExecute(r)
}

/*
FindEventsByEventTypeUsingGET Finds nonpersistent events by event type  (optional: after date, realms)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param eventType Event type
 @return ApiFindEventsByEventTypeUsingGETRequest
*/
func (a *EventControllerAPIService) FindEventsByEventTypeUsingGET(ctx context.Context, eventType string) ApiFindEventsByEventTypeUsingGETRequest {
	return ApiFindEventsByEventTypeUsingGETRequest{
		ApiService: a,
		ctx: ctx,
		eventType: eventType,
	}
}

// Execute executes the request
//  @return Event
func (a *EventControllerAPIService) FindEventsByEventTypeUsingGETExecute(r ApiFindEventsByEventTypeUsingGETRequest) (*Event, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Event
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventControllerAPIService.FindEventsByEventTypeUsingGET")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/search/events/type/{eventType}"
	localVarPath = strings.Replace(localVarPath, "{"+"eventType"+"}", url.PathEscape(parameterValueToString(r.eventType, "eventType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.afterDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "afterDate", r.afterDate, "")
	}
	if r.realms != nil {
		t := *r.realms
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "realms[]", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "realms[]", t, "multi")
		}
	}
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

type ApiFindEventsByFullTextUsingGETRequest struct {
	ctx context.Context
	ApiService *EventControllerAPIService
	fullTextPart string
	afterDate *int64
}

// After date
func (r ApiFindEventsByFullTextUsingGETRequest) AfterDate(afterDate int64) ApiFindEventsByFullTextUsingGETRequest {
	r.afterDate = &afterDate
	return r
}

func (r ApiFindEventsByFullTextUsingGETRequest) Execute() (*Event, *http.Response, error) {
	return r.ApiService.FindEventsByFullTextUsingGETExecute(r)
}

/*
FindEventsByFullTextUsingGET Finds nonpersistent events by full-text search (optional: after date)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fullTextPart Part of the event full-text representation
 @return ApiFindEventsByFullTextUsingGETRequest
*/
func (a *EventControllerAPIService) FindEventsByFullTextUsingGET(ctx context.Context, fullTextPart string) ApiFindEventsByFullTextUsingGETRequest {
	return ApiFindEventsByFullTextUsingGETRequest{
		ApiService: a,
		ctx: ctx,
		fullTextPart: fullTextPart,
	}
}

// Execute executes the request
//  @return Event
func (a *EventControllerAPIService) FindEventsByFullTextUsingGETExecute(r ApiFindEventsByFullTextUsingGETRequest) (*Event, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Event
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventControllerAPIService.FindEventsByFullTextUsingGET")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/search/events/{fullTextPart}"
	localVarPath = strings.Replace(localVarPath, "{"+"fullTextPart"+"}", url.PathEscape(parameterValueToString(r.fullTextPart, "fullTextPart")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.afterDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "afterDate", r.afterDate, "")
	}
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

type ApiFindEventsByRealmsUsingGETRequest struct {
	ctx context.Context
	ApiService *EventControllerAPIService
	realms *[]string
	afterDate *int64
}

// Realms
func (r ApiFindEventsByRealmsUsingGETRequest) Realms(realms []string) ApiFindEventsByRealmsUsingGETRequest {
	r.realms = &realms
	return r
}

// After date
func (r ApiFindEventsByRealmsUsingGETRequest) AfterDate(afterDate int64) ApiFindEventsByRealmsUsingGETRequest {
	r.afterDate = &afterDate
	return r
}

func (r ApiFindEventsByRealmsUsingGETRequest) Execute() (*Event, *http.Response, error) {
	return r.ApiService.FindEventsByRealmsUsingGETExecute(r)
}

/*
FindEventsByRealmsUsingGET Finds nonpersistent events by realms (optional: after date)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFindEventsByRealmsUsingGETRequest
*/
func (a *EventControllerAPIService) FindEventsByRealmsUsingGET(ctx context.Context) ApiFindEventsByRealmsUsingGETRequest {
	return ApiFindEventsByRealmsUsingGETRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Event
func (a *EventControllerAPIService) FindEventsByRealmsUsingGETExecute(r ApiFindEventsByRealmsUsingGETRequest) (*Event, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Event
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventControllerAPIService.FindEventsByRealmsUsingGET")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/search/events/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.realms == nil {
		return localVarReturnValue, nil, reportError("realms is required and must be specified")
	}

	if r.afterDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "afterDate", r.afterDate, "")
	}
	{
		t := *r.realms
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "realms[]", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "realms[]", t, "multi")
		}
	}
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

type ApiGetAllUsingGETRequest struct {
	ctx context.Context
	ApiService *EventControllerAPIService
	afterDate *int64
}

// After date
func (r ApiGetAllUsingGETRequest) AfterDate(afterDate int64) ApiGetAllUsingGETRequest {
	r.afterDate = &afterDate
	return r
}

func (r ApiGetAllUsingGETRequest) Execute() (*Event, *http.Response, error) {
	return r.ApiService.GetAllUsingGETExecute(r)
}

/*
GetAllUsingGET Returns all nonpersistent events (20000 max, optional: after date)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllUsingGETRequest
*/
func (a *EventControllerAPIService) GetAllUsingGET(ctx context.Context) ApiGetAllUsingGETRequest {
	return ApiGetAllUsingGETRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Event
func (a *EventControllerAPIService) GetAllUsingGETExecute(r ApiGetAllUsingGETRequest) (*Event, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Event
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventControllerAPIService.GetAllUsingGET")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/events/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.afterDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "afterDate", r.afterDate, "")
	}
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

type ApiGetMyEventsUsingGETRequest struct {
	ctx context.Context
	ApiService *EventControllerAPIService
	afterDate *int64
}

// After date
func (r ApiGetMyEventsUsingGETRequest) AfterDate(afterDate int64) ApiGetMyEventsUsingGETRequest {
	r.afterDate = &afterDate
	return r
}

func (r ApiGetMyEventsUsingGETRequest) Execute() (*Event, *http.Response, error) {
	return r.ApiService.GetMyEventsUsingGETExecute(r)
}

/*
GetMyEventsUsingGET Returns all nonpersistent events fore current user (20000 max, optional: after date)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMyEventsUsingGETRequest
*/
func (a *EventControllerAPIService) GetMyEventsUsingGET(ctx context.Context) ApiGetMyEventsUsingGETRequest {
	return ApiGetMyEventsUsingGETRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Event
func (a *EventControllerAPIService) GetMyEventsUsingGETExecute(r ApiGetMyEventsUsingGETRequest) (*Event, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Event
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventControllerAPIService.GetMyEventsUsingGET")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/events/user/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.afterDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "afterDate", r.afterDate, "")
	}
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
