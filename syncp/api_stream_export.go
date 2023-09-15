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

type StreamExportAPI interface {

	/*
		CreateStreamShares Create Stream Shares

		Creates Stream Shares

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param streamExportId
		@return ApiCreateStreamSharesRequest
	*/
	CreateStreamShares(ctx context.Context, streamExportId string) ApiCreateStreamSharesRequest

	// CreateStreamSharesExecute executes the request
	//  @return StreamShareViewResponse
	CreateStreamSharesExecute(r ApiCreateStreamSharesRequest) (*StreamShareViewResponse, *http.Response, error)

	/*
		DeleteStreamExport Delete Stream Export

		Delete Stream Export

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param streamExportId
		@return ApiDeleteStreamExportRequest
	*/
	DeleteStreamExport(ctx context.Context, streamExportId string) ApiDeleteStreamExportRequest

	// DeleteStreamExportExecute executes the request
	DeleteStreamExportExecute(r ApiDeleteStreamExportRequest) (*http.Response, error)

	/*
		DeleteStreamShare Delete Stream Share

		Revokes the share for provided Account NKey

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param streamExportId
		@param targetAccountNKeyPublic
		@return ApiDeleteStreamShareRequest
	*/
	DeleteStreamShare(ctx context.Context, streamExportId string, targetAccountNKeyPublic string) ApiDeleteStreamShareRequest

	// DeleteStreamShareExecute executes the request
	DeleteStreamShareExecute(r ApiDeleteStreamShareRequest) (*http.Response, error)

	/*
		GetStreamExport Get Stream Export

		Returns Stream Export info

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param streamExportId
		@return ApiGetStreamExportRequest
	*/
	GetStreamExport(ctx context.Context, streamExportId string) ApiGetStreamExportRequest

	// GetStreamExportExecute executes the request
	//  @return StreamExportViewResponse
	GetStreamExportExecute(r ApiGetStreamExportRequest) (*StreamExportViewResponse, *http.Response, error)

	/*
		ListStreamShares List Stream Shares

		List stream shares for this export

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param streamExportId
		@return ApiListStreamSharesRequest
	*/
	ListStreamShares(ctx context.Context, streamExportId string) ApiListStreamSharesRequest

	// ListStreamSharesExecute executes the request
	//  @return StreamShareListResponse
	ListStreamSharesExecute(r ApiListStreamSharesRequest) (*StreamShareListResponse, *http.Response, error)
}

// StreamExportAPIService StreamExportAPI service
type StreamExportAPIService service

type ApiCreateStreamSharesRequest struct {
	ctx                      context.Context
	ApiService               StreamExportAPI
	streamExportId           string
	streamShareCreateRequest *StreamShareCreateRequest
}

func (r ApiCreateStreamSharesRequest) StreamShareCreateRequest(streamShareCreateRequest StreamShareCreateRequest) ApiCreateStreamSharesRequest {
	r.streamShareCreateRequest = &streamShareCreateRequest
	return r
}

func (r ApiCreateStreamSharesRequest) Execute() (*StreamShareViewResponse, *http.Response, error) {
	return r.ApiService.CreateStreamSharesExecute(r)
}

/*
CreateStreamShares Create Stream Shares

Creates Stream Shares

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamExportId
	@return ApiCreateStreamSharesRequest
*/
func (a *StreamExportAPIService) CreateStreamShares(ctx context.Context, streamExportId string) ApiCreateStreamSharesRequest {
	return ApiCreateStreamSharesRequest{
		ApiService:     a,
		ctx:            ctx,
		streamExportId: streamExportId,
	}
}

// Execute executes the request
//
//	@return StreamShareViewResponse
func (a *StreamExportAPIService) CreateStreamSharesExecute(r ApiCreateStreamSharesRequest) (*StreamShareViewResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StreamShareViewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamExportAPIService.CreateStreamShares")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stream-exports/{streamExportId}/shares"
	localVarPath = strings.Replace(localVarPath, "{"+"streamExportId"+"}", url.PathEscape(parameterValueToString(r.streamExportId, "streamExportId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.streamShareCreateRequest
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

type ApiDeleteStreamExportRequest struct {
	ctx            context.Context
	ApiService     StreamExportAPI
	streamExportId string
}

func (r ApiDeleteStreamExportRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteStreamExportExecute(r)
}

/*
DeleteStreamExport Delete Stream Export

Delete Stream Export

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamExportId
	@return ApiDeleteStreamExportRequest
*/
func (a *StreamExportAPIService) DeleteStreamExport(ctx context.Context, streamExportId string) ApiDeleteStreamExportRequest {
	return ApiDeleteStreamExportRequest{
		ApiService:     a,
		ctx:            ctx,
		streamExportId: streamExportId,
	}
}

// Execute executes the request
func (a *StreamExportAPIService) DeleteStreamExportExecute(r ApiDeleteStreamExportRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamExportAPIService.DeleteStreamExport")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stream-exports/{streamExportId}"
	localVarPath = strings.Replace(localVarPath, "{"+"streamExportId"+"}", url.PathEscape(parameterValueToString(r.streamExportId, "streamExportId")), -1)

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

type ApiDeleteStreamShareRequest struct {
	ctx                     context.Context
	ApiService              StreamExportAPI
	streamExportId          string
	targetAccountNKeyPublic string
}

func (r ApiDeleteStreamShareRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteStreamShareExecute(r)
}

/*
DeleteStreamShare Delete Stream Share

Revokes the share for provided Account NKey

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamExportId
	@param targetAccountNKeyPublic
	@return ApiDeleteStreamShareRequest
*/
func (a *StreamExportAPIService) DeleteStreamShare(ctx context.Context, streamExportId string, targetAccountNKeyPublic string) ApiDeleteStreamShareRequest {
	return ApiDeleteStreamShareRequest{
		ApiService:              a,
		ctx:                     ctx,
		streamExportId:          streamExportId,
		targetAccountNKeyPublic: targetAccountNKeyPublic,
	}
}

// Execute executes the request
func (a *StreamExportAPIService) DeleteStreamShareExecute(r ApiDeleteStreamShareRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamExportAPIService.DeleteStreamShare")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stream-exports/{streamExportId}/shares/{targetAccountNKeyPublic}"
	localVarPath = strings.Replace(localVarPath, "{"+"streamExportId"+"}", url.PathEscape(parameterValueToString(r.streamExportId, "streamExportId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"targetAccountNKeyPublic"+"}", url.PathEscape(parameterValueToString(r.targetAccountNKeyPublic, "targetAccountNKeyPublic")), -1)

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

type ApiGetStreamExportRequest struct {
	ctx            context.Context
	ApiService     StreamExportAPI
	streamExportId string
}

func (r ApiGetStreamExportRequest) Execute() (*StreamExportViewResponse, *http.Response, error) {
	return r.ApiService.GetStreamExportExecute(r)
}

/*
GetStreamExport Get Stream Export

Returns Stream Export info

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamExportId
	@return ApiGetStreamExportRequest
*/
func (a *StreamExportAPIService) GetStreamExport(ctx context.Context, streamExportId string) ApiGetStreamExportRequest {
	return ApiGetStreamExportRequest{
		ApiService:     a,
		ctx:            ctx,
		streamExportId: streamExportId,
	}
}

// Execute executes the request
//
//	@return StreamExportViewResponse
func (a *StreamExportAPIService) GetStreamExportExecute(r ApiGetStreamExportRequest) (*StreamExportViewResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StreamExportViewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamExportAPIService.GetStreamExport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stream-exports/{streamExportId}"
	localVarPath = strings.Replace(localVarPath, "{"+"streamExportId"+"}", url.PathEscape(parameterValueToString(r.streamExportId, "streamExportId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type ApiListStreamSharesRequest struct {
	ctx            context.Context
	ApiService     StreamExportAPI
	streamExportId string
}

func (r ApiListStreamSharesRequest) Execute() (*StreamShareListResponse, *http.Response, error) {
	return r.ApiService.ListStreamSharesExecute(r)
}

/*
ListStreamShares List Stream Shares

List stream shares for this export

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamExportId
	@return ApiListStreamSharesRequest
*/
func (a *StreamExportAPIService) ListStreamShares(ctx context.Context, streamExportId string) ApiListStreamSharesRequest {
	return ApiListStreamSharesRequest{
		ApiService:     a,
		ctx:            ctx,
		streamExportId: streamExportId,
	}
}

// Execute executes the request
//
//	@return StreamShareListResponse
func (a *StreamExportAPIService) ListStreamSharesExecute(r ApiListStreamSharesRequest) (*StreamShareListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StreamShareListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamExportAPIService.ListStreamShares")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stream-exports/{streamExportId}/shares"
	localVarPath = strings.Replace(localVarPath, "{"+"streamExportId"+"}", url.PathEscape(parameterValueToString(r.streamExportId, "streamExportId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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