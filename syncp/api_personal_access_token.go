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

type PersonalAccessTokenAPI interface {

	/*
		DeletePersonalAccessToken Delete Personal Access Token

		Delete Personal Access Token

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param tokenId
		@return ApiDeletePersonalAccessTokenRequest
	*/
	DeletePersonalAccessToken(ctx context.Context, tokenId string) ApiDeletePersonalAccessTokenRequest

	// DeletePersonalAccessTokenExecute executes the request
	DeletePersonalAccessTokenExecute(r ApiDeletePersonalAccessTokenRequest) (*http.Response, error)

	/*
		GetPersonalAccessToken Get Personal Access Token

		Get Personal Access Token

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param tokenId
		@return ApiGetPersonalAccessTokenRequest
	*/
	GetPersonalAccessToken(ctx context.Context, tokenId string) ApiGetPersonalAccessTokenRequest

	// GetPersonalAccessTokenExecute executes the request
	//  @return AppUserAccessTokenViewResponse
	GetPersonalAccessTokenExecute(r ApiGetPersonalAccessTokenRequest) (*AppUserAccessTokenViewResponse, *http.Response, error)

	/*
		UpdatePersonalAccessToken Update Personal Access Token

		Update Personal Access Token

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param tokenId
		@return ApiUpdatePersonalAccessTokenRequest
	*/
	UpdatePersonalAccessToken(ctx context.Context, tokenId string) ApiUpdatePersonalAccessTokenRequest

	// UpdatePersonalAccessTokenExecute executes the request
	//  @return AppUserAccessTokenViewResponse
	UpdatePersonalAccessTokenExecute(r ApiUpdatePersonalAccessTokenRequest) (*AppUserAccessTokenViewResponse, *http.Response, error)
}

// PersonalAccessTokenAPIService PersonalAccessTokenAPI service
type PersonalAccessTokenAPIService service

type ApiDeletePersonalAccessTokenRequest struct {
	ctx        context.Context
	ApiService PersonalAccessTokenAPI
	tokenId    string
}

func (r ApiDeletePersonalAccessTokenRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePersonalAccessTokenExecute(r)
}

/*
DeletePersonalAccessToken Delete Personal Access Token

Delete Personal Access Token

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tokenId
	@return ApiDeletePersonalAccessTokenRequest
*/
func (a *PersonalAccessTokenAPIService) DeletePersonalAccessToken(ctx context.Context, tokenId string) ApiDeletePersonalAccessTokenRequest {
	return ApiDeletePersonalAccessTokenRequest{
		ApiService: a,
		ctx:        ctx,
		tokenId:    tokenId,
	}
}

// Execute executes the request
func (a *PersonalAccessTokenAPIService) DeletePersonalAccessTokenExecute(r ApiDeletePersonalAccessTokenRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PersonalAccessTokenAPIService.DeletePersonalAccessToken")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/personal-access-tokens/{tokenId}"
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

type ApiGetPersonalAccessTokenRequest struct {
	ctx        context.Context
	ApiService PersonalAccessTokenAPI
	tokenId    string
}

func (r ApiGetPersonalAccessTokenRequest) Execute() (*AppUserAccessTokenViewResponse, *http.Response, error) {
	return r.ApiService.GetPersonalAccessTokenExecute(r)
}

/*
GetPersonalAccessToken Get Personal Access Token

Get Personal Access Token

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tokenId
	@return ApiGetPersonalAccessTokenRequest
*/
func (a *PersonalAccessTokenAPIService) GetPersonalAccessToken(ctx context.Context, tokenId string) ApiGetPersonalAccessTokenRequest {
	return ApiGetPersonalAccessTokenRequest{
		ApiService: a,
		ctx:        ctx,
		tokenId:    tokenId,
	}
}

// Execute executes the request
//
//	@return AppUserAccessTokenViewResponse
func (a *PersonalAccessTokenAPIService) GetPersonalAccessTokenExecute(r ApiGetPersonalAccessTokenRequest) (*AppUserAccessTokenViewResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppUserAccessTokenViewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PersonalAccessTokenAPIService.GetPersonalAccessToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/personal-access-tokens/{tokenId}"
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

type ApiUpdatePersonalAccessTokenRequest struct {
	ctx                             context.Context
	ApiService                      PersonalAccessTokenAPI
	tokenId                         string
	appUserAccessTokenUpdateRequest *AppUserAccessTokenUpdateRequest
}

func (r ApiUpdatePersonalAccessTokenRequest) AppUserAccessTokenUpdateRequest(appUserAccessTokenUpdateRequest AppUserAccessTokenUpdateRequest) ApiUpdatePersonalAccessTokenRequest {
	r.appUserAccessTokenUpdateRequest = &appUserAccessTokenUpdateRequest
	return r
}

func (r ApiUpdatePersonalAccessTokenRequest) Execute() (*AppUserAccessTokenViewResponse, *http.Response, error) {
	return r.ApiService.UpdatePersonalAccessTokenExecute(r)
}

/*
UpdatePersonalAccessToken Update Personal Access Token

Update Personal Access Token

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tokenId
	@return ApiUpdatePersonalAccessTokenRequest
*/
func (a *PersonalAccessTokenAPIService) UpdatePersonalAccessToken(ctx context.Context, tokenId string) ApiUpdatePersonalAccessTokenRequest {
	return ApiUpdatePersonalAccessTokenRequest{
		ApiService: a,
		ctx:        ctx,
		tokenId:    tokenId,
	}
}

// Execute executes the request
//
//	@return AppUserAccessTokenViewResponse
func (a *PersonalAccessTokenAPIService) UpdatePersonalAccessTokenExecute(r ApiUpdatePersonalAccessTokenRequest) (*AppUserAccessTokenViewResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppUserAccessTokenViewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PersonalAccessTokenAPIService.UpdatePersonalAccessToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/personal-access-tokens/{tokenId}"
	localVarPath = strings.Replace(localVarPath, "{"+"tokenId"+"}", url.PathEscape(parameterValueToString(r.tokenId, "tokenId")), -1)

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
	localVarPostBody = r.appUserAccessTokenUpdateRequest
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