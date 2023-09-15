/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type AgentTokenAPI interface {

	/*
		DeleteAgentToken Delete Agent Token

		Delete an agent token

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param tokenId
		@return ApiDeleteAgentTokenRequest
	*/
	DeleteAgentToken(ctx context.Context, tokenId string) ApiDeleteAgentTokenRequest

	// DeleteAgentTokenExecute executes the request
	DeleteAgentTokenExecute(r ApiDeleteAgentTokenRequest) (*http.Response, error)
}

// AgentTokenAPIService AgentTokenAPI service
type AgentTokenAPIService service

type ApiDeleteAgentTokenRequest struct {
	ctx        context.Context
	ApiService AgentTokenAPI
	tokenId    string
}

func (r ApiDeleteAgentTokenRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAgentTokenExecute(r)
}

/*
DeleteAgentToken Delete Agent Token

Delete an agent token

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tokenId
	@return ApiDeleteAgentTokenRequest
*/
func (a *AgentTokenAPIService) DeleteAgentToken(ctx context.Context, tokenId string) ApiDeleteAgentTokenRequest {
	return ApiDeleteAgentTokenRequest{
		ApiService: a,
		ctx:        ctx,
		tokenId:    tokenId,
	}
}

// Execute executes the request
func (a *AgentTokenAPIService) DeleteAgentTokenExecute(r ApiDeleteAgentTokenRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentTokenAPIService.DeleteAgentToken")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/agent-tokens/{tokenId}"
	localVarPath = strings.Replace(localVarPath, "{"+"tokenId"+"}", url.PathEscape(parameterValueToString(r.tokenId, "tokenId")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}