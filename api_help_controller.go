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
)


// HelpControllerAPIService HelpControllerAPI service
type HelpControllerAPIService service

type ApiGetHelpDefinitionUsingGETRequest struct {
	ctx context.Context
	ApiService *HelpControllerAPIService
	identifier *string
	locale *string
}

// Help comment identifier
func (r ApiGetHelpDefinitionUsingGETRequest) Identifier(identifier string) ApiGetHelpDefinitionUsingGETRequest {
	r.identifier = &identifier
	return r
}

// Locale
func (r ApiGetHelpDefinitionUsingGETRequest) Locale(locale string) ApiGetHelpDefinitionUsingGETRequest {
	r.locale = &locale
	return r
}

func (r ApiGetHelpDefinitionUsingGETRequest) Execute() (*HelpComment, *http.Response, error) {
	return r.ApiService.GetHelpDefinitionUsingGETExecute(r)
}

/*
GetHelpDefinitionUsingGET Get help comment for identifier and locale

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetHelpDefinitionUsingGETRequest
*/
func (a *HelpControllerAPIService) GetHelpDefinitionUsingGET(ctx context.Context) ApiGetHelpDefinitionUsingGETRequest {
	return ApiGetHelpDefinitionUsingGETRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HelpComment
func (a *HelpControllerAPIService) GetHelpDefinitionUsingGETExecute(r ApiGetHelpDefinitionUsingGETRequest) (*HelpComment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HelpComment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HelpControllerAPIService.GetHelpDefinitionUsingGET")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/helpcomments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.identifier == nil {
		return localVarReturnValue, nil, reportError("identifier is required and must be specified")
	}
	if r.locale == nil {
		return localVarReturnValue, nil, reportError("locale is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "identifier", r.identifier, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "locale", r.locale, "")
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
