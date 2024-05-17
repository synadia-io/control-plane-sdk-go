/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

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

type StreamAPI interface {

	/*
		CreatePullConsumer Create Pull Consumer

		Creates Pull Consumer

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param streamId
		@return ApiCreatePullConsumerRequest
	*/
	CreatePullConsumer(ctx context.Context, streamId string) ApiCreatePullConsumerRequest

	// CreatePullConsumerExecute executes the request
	//  @return JSPullConsumerInfoResponse
	CreatePullConsumerExecute(r ApiCreatePullConsumerRequest) (*JSPullConsumerInfoResponse, *http.Response, error)

	/*
		CreatePushConsumer Create Push Consumer

		Create Push Consumer

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param streamId
		@return ApiCreatePushConsumerRequest
	*/
	CreatePushConsumer(ctx context.Context, streamId string) ApiCreatePushConsumerRequest

	// CreatePushConsumerExecute executes the request
	//  @return JSPushConsumerInfoResponse
	CreatePushConsumerExecute(r ApiCreatePushConsumerRequest) (*JSPushConsumerInfoResponse, *http.Response, error)

	/*
		DeleteStream Delete Stream

		Deletes Stream

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param streamId
		@return ApiDeleteStreamRequest
	*/
	DeleteStream(ctx context.Context, streamId string) ApiDeleteStreamRequest

	// DeleteStreamExecute executes the request
	DeleteStreamExecute(r ApiDeleteStreamRequest) (*http.Response, error)

	/*
		GetStreamInfo Get Stream

		Get Stream

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param streamId
		@return ApiGetStreamInfoRequest
	*/
	GetStreamInfo(ctx context.Context, streamId string) ApiGetStreamInfoRequest

	// GetStreamInfoExecute executes the request
	//  @return JSStreamInfoResponse
	GetStreamInfoExecute(r ApiGetStreamInfoRequest) (*JSStreamInfoResponse, *http.Response, error)

	/*
		ListConsumers List Consumers

		List Consumers

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param streamId
		@return ApiListConsumersRequest
	*/
	ListConsumers(ctx context.Context, streamId string) ApiListConsumersRequest

	// ListConsumersExecute executes the request
	//  @return JSConsumerInfoListResponse
	ListConsumersExecute(r ApiListConsumersRequest) (*JSConsumerInfoListResponse, *http.Response, error)

	/*
		UpdateStream Update Stream

		Updates Stream

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param streamId
		@return ApiUpdateStreamRequest
	*/
	UpdateStream(ctx context.Context, streamId string) ApiUpdateStreamRequest

	// UpdateStreamExecute executes the request
	//  @return JSStreamInfoResponse
	UpdateStreamExecute(r ApiUpdateStreamRequest) (*JSStreamInfoResponse, *http.Response, error)
}

// StreamAPIService StreamAPI service
type StreamAPIService service

type ApiCreatePullConsumerRequest struct {
	ctx                         context.Context
	ApiService                  StreamAPI
	streamId                    string
	jSPullConsumerConfigRequest *JSPullConsumerConfigRequest
}

func (r ApiCreatePullConsumerRequest) JSPullConsumerConfigRequest(jSPullConsumerConfigRequest JSPullConsumerConfigRequest) ApiCreatePullConsumerRequest {
	r.jSPullConsumerConfigRequest = &jSPullConsumerConfigRequest
	return r
}

func (r ApiCreatePullConsumerRequest) Execute() (*JSPullConsumerInfoResponse, *http.Response, error) {
	return r.ApiService.CreatePullConsumerExecute(r)
}

/*
CreatePullConsumer Create Pull Consumer

Creates Pull Consumer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamId
	@return ApiCreatePullConsumerRequest
*/
func (a *StreamAPIService) CreatePullConsumer(ctx context.Context, streamId string) ApiCreatePullConsumerRequest {
	return ApiCreatePullConsumerRequest{
		ApiService: a,
		ctx:        ctx,
		streamId:   streamId,
	}
}

// Execute executes the request
//
//	@return JSPullConsumerInfoResponse
func (a *StreamAPIService) CreatePullConsumerExecute(r ApiCreatePullConsumerRequest) (*JSPullConsumerInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *JSPullConsumerInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamAPIService.CreatePullConsumer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jetstream/stream/{streamId}/consumers/pull"
	localVarPath = strings.Replace(localVarPath, "{"+"streamId"+"}", url.PathEscape(parameterValueToString(r.streamId, "streamId")), -1)

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
	localVarPostBody = r.jSPullConsumerConfigRequest
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
			code:  localVarHTTPResponse.StatusCode,
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			code:  localVarHTTPResponse.StatusCode,
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreatePushConsumerRequest struct {
	ctx                         context.Context
	ApiService                  StreamAPI
	streamId                    string
	jSPushConsumerConfigRequest *JSPushConsumerConfigRequest
}

func (r ApiCreatePushConsumerRequest) JSPushConsumerConfigRequest(jSPushConsumerConfigRequest JSPushConsumerConfigRequest) ApiCreatePushConsumerRequest {
	r.jSPushConsumerConfigRequest = &jSPushConsumerConfigRequest
	return r
}

func (r ApiCreatePushConsumerRequest) Execute() (*JSPushConsumerInfoResponse, *http.Response, error) {
	return r.ApiService.CreatePushConsumerExecute(r)
}

/*
CreatePushConsumer Create Push Consumer

Create Push Consumer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamId
	@return ApiCreatePushConsumerRequest
*/
func (a *StreamAPIService) CreatePushConsumer(ctx context.Context, streamId string) ApiCreatePushConsumerRequest {
	return ApiCreatePushConsumerRequest{
		ApiService: a,
		ctx:        ctx,
		streamId:   streamId,
	}
}

// Execute executes the request
//
//	@return JSPushConsumerInfoResponse
func (a *StreamAPIService) CreatePushConsumerExecute(r ApiCreatePushConsumerRequest) (*JSPushConsumerInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *JSPushConsumerInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamAPIService.CreatePushConsumer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jetstream/stream/{streamId}/consumers/push"
	localVarPath = strings.Replace(localVarPath, "{"+"streamId"+"}", url.PathEscape(parameterValueToString(r.streamId, "streamId")), -1)

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
	localVarPostBody = r.jSPushConsumerConfigRequest
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
			code:  localVarHTTPResponse.StatusCode,
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			code:  localVarHTTPResponse.StatusCode,
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteStreamRequest struct {
	ctx        context.Context
	ApiService StreamAPI
	streamId   string
}

func (r ApiDeleteStreamRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteStreamExecute(r)
}

/*
DeleteStream Delete Stream

Deletes Stream

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamId
	@return ApiDeleteStreamRequest
*/
func (a *StreamAPIService) DeleteStream(ctx context.Context, streamId string) ApiDeleteStreamRequest {
	return ApiDeleteStreamRequest{
		ApiService: a,
		ctx:        ctx,
		streamId:   streamId,
	}
}

// Execute executes the request
func (a *StreamAPIService) DeleteStreamExecute(r ApiDeleteStreamRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamAPIService.DeleteStream")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jetstream/stream/{streamId}"
	localVarPath = strings.Replace(localVarPath, "{"+"streamId"+"}", url.PathEscape(parameterValueToString(r.streamId, "streamId")), -1)

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
			code:  localVarHTTPResponse.StatusCode,
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetStreamInfoRequest struct {
	ctx        context.Context
	ApiService StreamAPI
	streamId   string
}

func (r ApiGetStreamInfoRequest) Execute() (*JSStreamInfoResponse, *http.Response, error) {
	return r.ApiService.GetStreamInfoExecute(r)
}

/*
GetStreamInfo Get Stream

Get Stream

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamId
	@return ApiGetStreamInfoRequest
*/
func (a *StreamAPIService) GetStreamInfo(ctx context.Context, streamId string) ApiGetStreamInfoRequest {
	return ApiGetStreamInfoRequest{
		ApiService: a,
		ctx:        ctx,
		streamId:   streamId,
	}
}

// Execute executes the request
//
//	@return JSStreamInfoResponse
func (a *StreamAPIService) GetStreamInfoExecute(r ApiGetStreamInfoRequest) (*JSStreamInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *JSStreamInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamAPIService.GetStreamInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jetstream/stream/{streamId}"
	localVarPath = strings.Replace(localVarPath, "{"+"streamId"+"}", url.PathEscape(parameterValueToString(r.streamId, "streamId")), -1)

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
			code:  localVarHTTPResponse.StatusCode,
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			code:  localVarHTTPResponse.StatusCode,
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListConsumersRequest struct {
	ctx        context.Context
	ApiService StreamAPI
	streamId   string
}

func (r ApiListConsumersRequest) Execute() (*JSConsumerInfoListResponse, *http.Response, error) {
	return r.ApiService.ListConsumersExecute(r)
}

/*
ListConsumers List Consumers

List Consumers

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamId
	@return ApiListConsumersRequest
*/
func (a *StreamAPIService) ListConsumers(ctx context.Context, streamId string) ApiListConsumersRequest {
	return ApiListConsumersRequest{
		ApiService: a,
		ctx:        ctx,
		streamId:   streamId,
	}
}

// Execute executes the request
//
//	@return JSConsumerInfoListResponse
func (a *StreamAPIService) ListConsumersExecute(r ApiListConsumersRequest) (*JSConsumerInfoListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *JSConsumerInfoListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamAPIService.ListConsumers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jetstream/stream/{streamId}/consumers"
	localVarPath = strings.Replace(localVarPath, "{"+"streamId"+"}", url.PathEscape(parameterValueToString(r.streamId, "streamId")), -1)

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
			code:  localVarHTTPResponse.StatusCode,
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			code:  localVarHTTPResponse.StatusCode,
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateStreamRequest struct {
	ctx                   context.Context
	ApiService            StreamAPI
	streamId              string
	jSStreamConfigRequest *JSStreamConfigRequest
}

func (r ApiUpdateStreamRequest) JSStreamConfigRequest(jSStreamConfigRequest JSStreamConfigRequest) ApiUpdateStreamRequest {
	r.jSStreamConfigRequest = &jSStreamConfigRequest
	return r
}

func (r ApiUpdateStreamRequest) Execute() (*JSStreamInfoResponse, *http.Response, error) {
	return r.ApiService.UpdateStreamExecute(r)
}

/*
UpdateStream Update Stream

Updates Stream

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamId
	@return ApiUpdateStreamRequest
*/
func (a *StreamAPIService) UpdateStream(ctx context.Context, streamId string) ApiUpdateStreamRequest {
	return ApiUpdateStreamRequest{
		ApiService: a,
		ctx:        ctx,
		streamId:   streamId,
	}
}

// Execute executes the request
//
//	@return JSStreamInfoResponse
func (a *StreamAPIService) UpdateStreamExecute(r ApiUpdateStreamRequest) (*JSStreamInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *JSStreamInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamAPIService.UpdateStream")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jetstream/stream/{streamId}"
	localVarPath = strings.Replace(localVarPath, "{"+"streamId"+"}", url.PathEscape(parameterValueToString(r.streamId, "streamId")), -1)

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
	localVarPostBody = r.jSStreamConfigRequest
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
			code:  localVarHTTPResponse.StatusCode,
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			code:  localVarHTTPResponse.StatusCode,
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
