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


// ComplaintControllerAPIService ComplaintControllerAPI service
type ComplaintControllerAPIService service

type ApiCreateComplaintUsingPOSTRequest struct {
	ctx context.Context
	ApiService *ComplaintControllerAPIService
	subjectMatterId *string
	subjectMatterType *string
	text *string
}

// Subject Matter ID
func (r ApiCreateComplaintUsingPOSTRequest) SubjectMatterId(subjectMatterId string) ApiCreateComplaintUsingPOSTRequest {
	r.subjectMatterId = &subjectMatterId
	return r
}

// Subject Matter Type
func (r ApiCreateComplaintUsingPOSTRequest) SubjectMatterType(subjectMatterType string) ApiCreateComplaintUsingPOSTRequest {
	r.subjectMatterType = &subjectMatterType
	return r
}

// Text
func (r ApiCreateComplaintUsingPOSTRequest) Text(text string) ApiCreateComplaintUsingPOSTRequest {
	r.text = &text
	return r
}

func (r ApiCreateComplaintUsingPOSTRequest) Execute() (*ComplaintView, *http.Response, error) {
	return r.ApiService.CreateComplaintUsingPOSTExecute(r)
}

/*
CreateComplaintUsingPOST Opens new complaint

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateComplaintUsingPOSTRequest
*/
func (a *ComplaintControllerAPIService) CreateComplaintUsingPOST(ctx context.Context) ApiCreateComplaintUsingPOSTRequest {
	return ApiCreateComplaintUsingPOSTRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ComplaintView
func (a *ComplaintControllerAPIService) CreateComplaintUsingPOSTExecute(r ApiCreateComplaintUsingPOSTRequest) (*ComplaintView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ComplaintView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComplaintControllerAPIService.CreateComplaintUsingPOST")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/complaints"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subjectMatterId == nil {
		return localVarReturnValue, nil, reportError("subjectMatterId is required and must be specified")
	}
	if r.subjectMatterType == nil {
		return localVarReturnValue, nil, reportError("subjectMatterType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "subjectMatterId", r.subjectMatterId, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "subjectMatterType", r.subjectMatterType, "")
	if r.text != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "text", r.text, "")
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

type ApiGetComplaintUsingGETRequest struct {
	ctx context.Context
	ApiService *ComplaintControllerAPIService
	complaintId string
}

func (r ApiGetComplaintUsingGETRequest) Execute() (*ComplaintView, *http.Response, error) {
	return r.ApiService.GetComplaintUsingGETExecute(r)
}

/*
GetComplaintUsingGET Returns complaint

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param complaintId Complaint id
 @return ApiGetComplaintUsingGETRequest
*/
func (a *ComplaintControllerAPIService) GetComplaintUsingGET(ctx context.Context, complaintId string) ApiGetComplaintUsingGETRequest {
	return ApiGetComplaintUsingGETRequest{
		ApiService: a,
		ctx: ctx,
		complaintId: complaintId,
	}
}

// Execute executes the request
//  @return ComplaintView
func (a *ComplaintControllerAPIService) GetComplaintUsingGETExecute(r ApiGetComplaintUsingGETRequest) (*ComplaintView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ComplaintView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComplaintControllerAPIService.GetComplaintUsingGET")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/complaints/{complaintId}"
	localVarPath = strings.Replace(localVarPath, "{"+"complaintId"+"}", url.PathEscape(parameterValueToString(r.complaintId, "complaintId")), -1)

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
