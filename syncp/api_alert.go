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

type AlertAPI interface {

	/*
		AcknowledgeAlert Acknowledge Alert

		Acknowledge Alert

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param alertId
		@return ApiAcknowledgeAlertRequest
	*/
	AcknowledgeAlert(ctx context.Context, alertId string) ApiAcknowledgeAlertRequest

	// AcknowledgeAlertExecute executes the request
	AcknowledgeAlertExecute(r ApiAcknowledgeAlertRequest) (*http.Response, error)

	/*
		GetAlert Get Alert

		Get Alert

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param alertId
		@return ApiGetAlertRequest
	*/
	GetAlert(ctx context.Context, alertId string) ApiGetAlertRequest

	// GetAlertExecute executes the request
	//  @return AlertViewResponse
	GetAlertExecute(r ApiGetAlertRequest) (*AlertViewResponse, *http.Response, error)
}

// AlertAPIService AlertAPI service
type AlertAPIService service

type ApiAcknowledgeAlertRequest struct {
	ctx        context.Context
	ApiService AlertAPI
	alertId    string
}

func (r ApiAcknowledgeAlertRequest) Execute() (*http.Response, error) {
	return r.ApiService.AcknowledgeAlertExecute(r)
}

/*
AcknowledgeAlert Acknowledge Alert

Acknowledge Alert

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param alertId
	@return ApiAcknowledgeAlertRequest
*/
func (a *AlertAPIService) AcknowledgeAlert(ctx context.Context, alertId string) ApiAcknowledgeAlertRequest {
	return ApiAcknowledgeAlertRequest{
		ApiService: a,
		ctx:        ctx,
		alertId:    alertId,
	}
}

// Execute executes the request
func (a *AlertAPIService) AcknowledgeAlertExecute(r ApiAcknowledgeAlertRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertAPIService.AcknowledgeAlert")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/alerts/{alertId}/acknowledge"
	localVarPath = strings.Replace(localVarPath, "{"+"alertId"+"}", url.PathEscape(parameterValueToString(r.alertId, "alertId")), -1)

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

type ApiGetAlertRequest struct {
	ctx        context.Context
	ApiService AlertAPI
	alertId    string
}

func (r ApiGetAlertRequest) Execute() (*AlertViewResponse, *http.Response, error) {
	return r.ApiService.GetAlertExecute(r)
}

/*
GetAlert Get Alert

Get Alert

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param alertId
	@return ApiGetAlertRequest
*/
func (a *AlertAPIService) GetAlert(ctx context.Context, alertId string) ApiGetAlertRequest {
	return ApiGetAlertRequest{
		ApiService: a,
		ctx:        ctx,
		alertId:    alertId,
	}
}

// Execute executes the request
//
//	@return AlertViewResponse
func (a *AlertAPIService) GetAlertExecute(r ApiGetAlertRequest) (*AlertViewResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlertViewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertAPIService.GetAlert")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/alerts/{alertId}"
	localVarPath = strings.Replace(localVarPath, "{"+"alertId"+"}", url.PathEscape(parameterValueToString(r.alertId, "alertId")), -1)

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