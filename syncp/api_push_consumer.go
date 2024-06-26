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

type PushConsumerAPI interface {

	/*
		DeletePushConsumer Delete Push Consumer

		Delete Push Consumer

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param consumerId
		@return ApiDeletePushConsumerRequest
	*/
	DeletePushConsumer(ctx context.Context, consumerId string) ApiDeletePushConsumerRequest

	// DeletePushConsumerExecute executes the request
	DeletePushConsumerExecute(r ApiDeletePushConsumerRequest) (*http.Response, error)

	/*
		GetPushConsumerInfo Get Push Consumer

		Returns JetStream push consumer info

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param consumerId
		@return ApiGetPushConsumerInfoRequest
	*/
	GetPushConsumerInfo(ctx context.Context, consumerId string) ApiGetPushConsumerInfoRequest

	// GetPushConsumerInfoExecute executes the request
	//  @return JSPushConsumerInfoResponse
	GetPushConsumerInfoExecute(r ApiGetPushConsumerInfoRequest) (*JSPushConsumerInfoResponse, *http.Response, error)

	/*
		UpdatePushConsumer Update Push Consumer

		Update Push Consumer

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param consumerId
		@return ApiUpdatePushConsumerRequest
	*/
	UpdatePushConsumer(ctx context.Context, consumerId string) ApiUpdatePushConsumerRequest

	// UpdatePushConsumerExecute executes the request
	//  @return JSPushConsumerInfoResponse
	UpdatePushConsumerExecute(r ApiUpdatePushConsumerRequest) (*JSPushConsumerInfoResponse, *http.Response, error)
}

// PushConsumerAPIService PushConsumerAPI service
type PushConsumerAPIService service

type ApiDeletePushConsumerRequest struct {
	ctx        context.Context
	ApiService PushConsumerAPI
	consumerId string
}

func (r ApiDeletePushConsumerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePushConsumerExecute(r)
}

/*
DeletePushConsumer Delete Push Consumer

Delete Push Consumer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param consumerId
	@return ApiDeletePushConsumerRequest
*/
func (a *PushConsumerAPIService) DeletePushConsumer(ctx context.Context, consumerId string) ApiDeletePushConsumerRequest {
	return ApiDeletePushConsumerRequest{
		ApiService: a,
		ctx:        ctx,
		consumerId: consumerId,
	}
}

// Execute executes the request
func (a *PushConsumerAPIService) DeletePushConsumerExecute(r ApiDeletePushConsumerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PushConsumerAPIService.DeletePushConsumer")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/consumers/push/{consumerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"consumerId"+"}", url.PathEscape(parameterValueToString(r.consumerId, "consumerId")), -1)

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

type ApiGetPushConsumerInfoRequest struct {
	ctx        context.Context
	ApiService PushConsumerAPI
	consumerId string
}

func (r ApiGetPushConsumerInfoRequest) Execute() (*JSPushConsumerInfoResponse, *http.Response, error) {
	return r.ApiService.GetPushConsumerInfoExecute(r)
}

/*
GetPushConsumerInfo Get Push Consumer

Returns JetStream push consumer info

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param consumerId
	@return ApiGetPushConsumerInfoRequest
*/
func (a *PushConsumerAPIService) GetPushConsumerInfo(ctx context.Context, consumerId string) ApiGetPushConsumerInfoRequest {
	return ApiGetPushConsumerInfoRequest{
		ApiService: a,
		ctx:        ctx,
		consumerId: consumerId,
	}
}

// Execute executes the request
//
//	@return JSPushConsumerInfoResponse
func (a *PushConsumerAPIService) GetPushConsumerInfoExecute(r ApiGetPushConsumerInfoRequest) (*JSPushConsumerInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *JSPushConsumerInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PushConsumerAPIService.GetPushConsumerInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/consumers/push/{consumerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"consumerId"+"}", url.PathEscape(parameterValueToString(r.consumerId, "consumerId")), -1)

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

type ApiUpdatePushConsumerRequest struct {
	ctx                         context.Context
	ApiService                  PushConsumerAPI
	consumerId                  string
	jSPushConsumerUpdateRequest *JSPushConsumerUpdateRequest
}

func (r ApiUpdatePushConsumerRequest) JSPushConsumerUpdateRequest(jSPushConsumerUpdateRequest JSPushConsumerUpdateRequest) ApiUpdatePushConsumerRequest {
	r.jSPushConsumerUpdateRequest = &jSPushConsumerUpdateRequest
	return r
}

func (r ApiUpdatePushConsumerRequest) Execute() (*JSPushConsumerInfoResponse, *http.Response, error) {
	return r.ApiService.UpdatePushConsumerExecute(r)
}

/*
UpdatePushConsumer Update Push Consumer

Update Push Consumer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param consumerId
	@return ApiUpdatePushConsumerRequest
*/
func (a *PushConsumerAPIService) UpdatePushConsumer(ctx context.Context, consumerId string) ApiUpdatePushConsumerRequest {
	return ApiUpdatePushConsumerRequest{
		ApiService: a,
		ctx:        ctx,
		consumerId: consumerId,
	}
}

// Execute executes the request
//
//	@return JSPushConsumerInfoResponse
func (a *PushConsumerAPIService) UpdatePushConsumerExecute(r ApiUpdatePushConsumerRequest) (*JSPushConsumerInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *JSPushConsumerInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PushConsumerAPIService.UpdatePushConsumer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/consumers/push/{consumerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"consumerId"+"}", url.PathEscape(parameterValueToString(r.consumerId, "consumerId")), -1)

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
	localVarPostBody = r.jSPushConsumerUpdateRequest
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
