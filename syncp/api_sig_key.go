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

type SigKeyAPI interface {

	/*
		DeleteAccountSk Delete Account Signing

		Deletes Account Signing

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param keyId
		@return ApiDeleteAccountSkRequest
	*/
	DeleteAccountSk(ctx context.Context, keyId string) ApiDeleteAccountSkRequest

	// DeleteAccountSkExecute executes the request
	DeleteAccountSkExecute(r ApiDeleteAccountSkRequest) (*http.Response, error)

	/*
		GetAccountSk Get Account Signing

		Get Account Signing

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param keyId
		@return ApiGetAccountSkRequest
	*/
	GetAccountSk(ctx context.Context, keyId string) ApiGetAccountSkRequest

	// GetAccountSkExecute executes the request
	//  @return SigningKeyViewResponse
	GetAccountSkExecute(r ApiGetAccountSkRequest) (*SigningKeyViewResponse, *http.Response, error)

	/*
		UpdateAccountSk Update Account Signing

		Update Account Signing

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param keyId
		@return ApiUpdateAccountSkRequest
	*/
	UpdateAccountSk(ctx context.Context, keyId string) ApiUpdateAccountSkRequest

	// UpdateAccountSkExecute executes the request
	//  @return SigningKeyViewResponse
	UpdateAccountSkExecute(r ApiUpdateAccountSkRequest) (*SigningKeyViewResponse, *http.Response, error)
}

// SigKeyAPIService SigKeyAPI service
type SigKeyAPIService service

type ApiDeleteAccountSkRequest struct {
	ctx        context.Context
	ApiService SigKeyAPI
	keyId      string
}

func (r ApiDeleteAccountSkRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAccountSkExecute(r)
}

/*
DeleteAccountSk Delete Account Signing

Deletes Account Signing

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param keyId
	@return ApiDeleteAccountSkRequest
*/
func (a *SigKeyAPIService) DeleteAccountSk(ctx context.Context, keyId string) ApiDeleteAccountSkRequest {
	return ApiDeleteAccountSkRequest{
		ApiService: a,
		ctx:        ctx,
		keyId:      keyId,
	}
}

// Execute executes the request
func (a *SigKeyAPIService) DeleteAccountSkExecute(r ApiDeleteAccountSkRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SigKeyAPIService.DeleteAccountSk")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account-sks/{keyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"keyId"+"}", url.PathEscape(parameterValueToString(r.keyId, "keyId")), -1)

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

type ApiGetAccountSkRequest struct {
	ctx        context.Context
	ApiService SigKeyAPI
	keyId      string
}

func (r ApiGetAccountSkRequest) Execute() (*SigningKeyViewResponse, *http.Response, error) {
	return r.ApiService.GetAccountSkExecute(r)
}

/*
GetAccountSk Get Account Signing

Get Account Signing

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param keyId
	@return ApiGetAccountSkRequest
*/
func (a *SigKeyAPIService) GetAccountSk(ctx context.Context, keyId string) ApiGetAccountSkRequest {
	return ApiGetAccountSkRequest{
		ApiService: a,
		ctx:        ctx,
		keyId:      keyId,
	}
}

// Execute executes the request
//
//	@return SigningKeyViewResponse
func (a *SigKeyAPIService) GetAccountSkExecute(r ApiGetAccountSkRequest) (*SigningKeyViewResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SigningKeyViewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SigKeyAPIService.GetAccountSk")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account-sks/{keyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"keyId"+"}", url.PathEscape(parameterValueToString(r.keyId, "keyId")), -1)

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

type ApiUpdateAccountSkRequest struct {
	ctx                     context.Context
	ApiService              SigKeyAPI
	keyId                   string
	signingKeyUpdateRequest *SigningKeyUpdateRequest
}

func (r ApiUpdateAccountSkRequest) SigningKeyUpdateRequest(signingKeyUpdateRequest SigningKeyUpdateRequest) ApiUpdateAccountSkRequest {
	r.signingKeyUpdateRequest = &signingKeyUpdateRequest
	return r
}

func (r ApiUpdateAccountSkRequest) Execute() (*SigningKeyViewResponse, *http.Response, error) {
	return r.ApiService.UpdateAccountSkExecute(r)
}

/*
UpdateAccountSk Update Account Signing

Update Account Signing

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param keyId
	@return ApiUpdateAccountSkRequest
*/
func (a *SigKeyAPIService) UpdateAccountSk(ctx context.Context, keyId string) ApiUpdateAccountSkRequest {
	return ApiUpdateAccountSkRequest{
		ApiService: a,
		ctx:        ctx,
		keyId:      keyId,
	}
}

// Execute executes the request
//
//	@return SigningKeyViewResponse
func (a *SigKeyAPIService) UpdateAccountSkExecute(r ApiUpdateAccountSkRequest) (*SigningKeyViewResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SigningKeyViewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SigKeyAPIService.UpdateAccountSk")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account-sks/{keyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"keyId"+"}", url.PathEscape(parameterValueToString(r.keyId, "keyId")), -1)

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
	localVarPostBody = r.signingKeyUpdateRequest
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
