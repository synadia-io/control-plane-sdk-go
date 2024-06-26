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

type KvBucketAPI interface {

	/*
		CreateKvPullConsumer Create Pull Consumer

		Creates Pull Consumer

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param streamId
		@return ApiCreateKvPullConsumerRequest
	*/
	CreateKvPullConsumer(ctx context.Context, streamId string) ApiCreateKvPullConsumerRequest

	// CreateKvPullConsumerExecute executes the request
	//  @return JSPullConsumerInfoResponse
	CreateKvPullConsumerExecute(r ApiCreateKvPullConsumerRequest) (*JSPullConsumerInfoResponse, *http.Response, error)

	/*
		CreateKvPushConsumer Create Push Consumer

		Creates Push Consumer

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param streamId
		@return ApiCreateKvPushConsumerRequest
	*/
	CreateKvPushConsumer(ctx context.Context, streamId string) ApiCreateKvPushConsumerRequest

	// CreateKvPushConsumerExecute executes the request
	//  @return JSPushConsumerInfoResponse
	CreateKvPushConsumerExecute(r ApiCreateKvPushConsumerRequest) (*JSPushConsumerInfoResponse, *http.Response, error)

	/*
		DeleteKvBucket Delete KV Bucket

		Deletes KV Bucket

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param streamId
		@return ApiDeleteKvBucketRequest
	*/
	DeleteKvBucket(ctx context.Context, streamId string) ApiDeleteKvBucketRequest

	// DeleteKvBucketExecute executes the request
	DeleteKvBucketExecute(r ApiDeleteKvBucketRequest) (*http.Response, error)

	/*
		GetKvBucket Get KV Bucket

		Returns JetStream KV bucket info

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param streamId
		@return ApiGetKvBucketRequest
	*/
	GetKvBucket(ctx context.Context, streamId string) ApiGetKvBucketRequest

	// GetKvBucketExecute executes the request
	//  @return JSKVBucketViewResponse
	GetKvBucketExecute(r ApiGetKvBucketRequest) (*JSKVBucketViewResponse, *http.Response, error)

	/*
		ListKvConsumers List Consumers

		List Consumers

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param streamId
		@return ApiListKvConsumersRequest
	*/
	ListKvConsumers(ctx context.Context, streamId string) ApiListKvConsumersRequest

	// ListKvConsumersExecute executes the request
	//  @return JSConsumerInfoListResponse
	ListKvConsumersExecute(r ApiListKvConsumersRequest) (*JSConsumerInfoListResponse, *http.Response, error)

	/*
		PurgeKvBucket Purge KV Bucket

		Purges KV Bucket

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param streamId
		@return ApiPurgeKvBucketRequest
	*/
	PurgeKvBucket(ctx context.Context, streamId string) ApiPurgeKvBucketRequest

	// PurgeKvBucketExecute executes the request
	PurgeKvBucketExecute(r ApiPurgeKvBucketRequest) (*http.Response, error)

	/*
		UpdateKvBucket Update KV Bucket

		Update KV Bucket

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param streamId
		@return ApiUpdateKvBucketRequest
	*/
	UpdateKvBucket(ctx context.Context, streamId string) ApiUpdateKvBucketRequest

	// UpdateKvBucketExecute executes the request
	//  @return JSKVBucketViewResponse
	UpdateKvBucketExecute(r ApiUpdateKvBucketRequest) (*JSKVBucketViewResponse, *http.Response, error)
}

// KvBucketAPIService KvBucketAPI service
type KvBucketAPIService service

type ApiCreateKvPullConsumerRequest struct {
	ctx                         context.Context
	ApiService                  KvBucketAPI
	streamId                    string
	jSPullConsumerConfigRequest *JSPullConsumerConfigRequest
}

func (r ApiCreateKvPullConsumerRequest) JSPullConsumerConfigRequest(jSPullConsumerConfigRequest JSPullConsumerConfigRequest) ApiCreateKvPullConsumerRequest {
	r.jSPullConsumerConfigRequest = &jSPullConsumerConfigRequest
	return r
}

func (r ApiCreateKvPullConsumerRequest) Execute() (*JSPullConsumerInfoResponse, *http.Response, error) {
	return r.ApiService.CreateKvPullConsumerExecute(r)
}

/*
CreateKvPullConsumer Create Pull Consumer

Creates Pull Consumer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamId
	@return ApiCreateKvPullConsumerRequest
*/
func (a *KvBucketAPIService) CreateKvPullConsumer(ctx context.Context, streamId string) ApiCreateKvPullConsumerRequest {
	return ApiCreateKvPullConsumerRequest{
		ApiService: a,
		ctx:        ctx,
		streamId:   streamId,
	}
}

// Execute executes the request
//
//	@return JSPullConsumerInfoResponse
func (a *KvBucketAPIService) CreateKvPullConsumerExecute(r ApiCreateKvPullConsumerRequest) (*JSPullConsumerInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *JSPullConsumerInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KvBucketAPIService.CreateKvPullConsumer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jetstream/kv-bucket/{streamId}/consumers/pull"
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

type ApiCreateKvPushConsumerRequest struct {
	ctx                         context.Context
	ApiService                  KvBucketAPI
	streamId                    string
	jSPushConsumerConfigRequest *JSPushConsumerConfigRequest
}

func (r ApiCreateKvPushConsumerRequest) JSPushConsumerConfigRequest(jSPushConsumerConfigRequest JSPushConsumerConfigRequest) ApiCreateKvPushConsumerRequest {
	r.jSPushConsumerConfigRequest = &jSPushConsumerConfigRequest
	return r
}

func (r ApiCreateKvPushConsumerRequest) Execute() (*JSPushConsumerInfoResponse, *http.Response, error) {
	return r.ApiService.CreateKvPushConsumerExecute(r)
}

/*
CreateKvPushConsumer Create Push Consumer

Creates Push Consumer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamId
	@return ApiCreateKvPushConsumerRequest
*/
func (a *KvBucketAPIService) CreateKvPushConsumer(ctx context.Context, streamId string) ApiCreateKvPushConsumerRequest {
	return ApiCreateKvPushConsumerRequest{
		ApiService: a,
		ctx:        ctx,
		streamId:   streamId,
	}
}

// Execute executes the request
//
//	@return JSPushConsumerInfoResponse
func (a *KvBucketAPIService) CreateKvPushConsumerExecute(r ApiCreateKvPushConsumerRequest) (*JSPushConsumerInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *JSPushConsumerInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KvBucketAPIService.CreateKvPushConsumer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jetstream/kv-bucket/{streamId}/consumers/push"
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

type ApiDeleteKvBucketRequest struct {
	ctx        context.Context
	ApiService KvBucketAPI
	streamId   string
}

func (r ApiDeleteKvBucketRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteKvBucketExecute(r)
}

/*
DeleteKvBucket Delete KV Bucket

Deletes KV Bucket

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamId
	@return ApiDeleteKvBucketRequest
*/
func (a *KvBucketAPIService) DeleteKvBucket(ctx context.Context, streamId string) ApiDeleteKvBucketRequest {
	return ApiDeleteKvBucketRequest{
		ApiService: a,
		ctx:        ctx,
		streamId:   streamId,
	}
}

// Execute executes the request
func (a *KvBucketAPIService) DeleteKvBucketExecute(r ApiDeleteKvBucketRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KvBucketAPIService.DeleteKvBucket")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jetstream/kv-bucket/{streamId}"
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

type ApiGetKvBucketRequest struct {
	ctx        context.Context
	ApiService KvBucketAPI
	streamId   string
}

func (r ApiGetKvBucketRequest) Execute() (*JSKVBucketViewResponse, *http.Response, error) {
	return r.ApiService.GetKvBucketExecute(r)
}

/*
GetKvBucket Get KV Bucket

Returns JetStream KV bucket info

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamId
	@return ApiGetKvBucketRequest
*/
func (a *KvBucketAPIService) GetKvBucket(ctx context.Context, streamId string) ApiGetKvBucketRequest {
	return ApiGetKvBucketRequest{
		ApiService: a,
		ctx:        ctx,
		streamId:   streamId,
	}
}

// Execute executes the request
//
//	@return JSKVBucketViewResponse
func (a *KvBucketAPIService) GetKvBucketExecute(r ApiGetKvBucketRequest) (*JSKVBucketViewResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *JSKVBucketViewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KvBucketAPIService.GetKvBucket")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jetstream/kv-bucket/{streamId}"
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

type ApiListKvConsumersRequest struct {
	ctx        context.Context
	ApiService KvBucketAPI
	streamId   string
}

func (r ApiListKvConsumersRequest) Execute() (*JSConsumerInfoListResponse, *http.Response, error) {
	return r.ApiService.ListKvConsumersExecute(r)
}

/*
ListKvConsumers List Consumers

List Consumers

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamId
	@return ApiListKvConsumersRequest
*/
func (a *KvBucketAPIService) ListKvConsumers(ctx context.Context, streamId string) ApiListKvConsumersRequest {
	return ApiListKvConsumersRequest{
		ApiService: a,
		ctx:        ctx,
		streamId:   streamId,
	}
}

// Execute executes the request
//
//	@return JSConsumerInfoListResponse
func (a *KvBucketAPIService) ListKvConsumersExecute(r ApiListKvConsumersRequest) (*JSConsumerInfoListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *JSConsumerInfoListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KvBucketAPIService.ListKvConsumers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jetstream/kv-bucket/{streamId}/consumers"
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

type ApiPurgeKvBucketRequest struct {
	ctx        context.Context
	ApiService KvBucketAPI
	streamId   string
}

func (r ApiPurgeKvBucketRequest) Execute() (*http.Response, error) {
	return r.ApiService.PurgeKvBucketExecute(r)
}

/*
PurgeKvBucket Purge KV Bucket

Purges KV Bucket

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamId
	@return ApiPurgeKvBucketRequest
*/
func (a *KvBucketAPIService) PurgeKvBucket(ctx context.Context, streamId string) ApiPurgeKvBucketRequest {
	return ApiPurgeKvBucketRequest{
		ApiService: a,
		ctx:        ctx,
		streamId:   streamId,
	}
}

// Execute executes the request
func (a *KvBucketAPIService) PurgeKvBucketExecute(r ApiPurgeKvBucketRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KvBucketAPIService.PurgeKvBucket")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jetstream/kv-bucket/{streamId}/purge"
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

type ApiUpdateKvBucketRequest struct {
	ctx                     context.Context
	ApiService              KvBucketAPI
	streamId                string
	jSKVBucketUpdateRequest *JSKVBucketUpdateRequest
}

func (r ApiUpdateKvBucketRequest) JSKVBucketUpdateRequest(jSKVBucketUpdateRequest JSKVBucketUpdateRequest) ApiUpdateKvBucketRequest {
	r.jSKVBucketUpdateRequest = &jSKVBucketUpdateRequest
	return r
}

func (r ApiUpdateKvBucketRequest) Execute() (*JSKVBucketViewResponse, *http.Response, error) {
	return r.ApiService.UpdateKvBucketExecute(r)
}

/*
UpdateKvBucket Update KV Bucket

Update KV Bucket

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamId
	@return ApiUpdateKvBucketRequest
*/
func (a *KvBucketAPIService) UpdateKvBucket(ctx context.Context, streamId string) ApiUpdateKvBucketRequest {
	return ApiUpdateKvBucketRequest{
		ApiService: a,
		ctx:        ctx,
		streamId:   streamId,
	}
}

// Execute executes the request
//
//	@return JSKVBucketViewResponse
func (a *KvBucketAPIService) UpdateKvBucketExecute(r ApiUpdateKvBucketRequest) (*JSKVBucketViewResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *JSKVBucketViewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KvBucketAPIService.UpdateKvBucket")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jetstream/kv-bucket/{streamId}"
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
	localVarPostBody = r.jSKVBucketUpdateRequest
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
