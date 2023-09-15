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

type AppUserAPI interface {

	/*
		DeleteAppUser Delete App User

		Delete App User

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param appUserId
		@return ApiDeleteAppUserRequest
	*/
	DeleteAppUser(ctx context.Context, appUserId string) ApiDeleteAppUserRequest

	// DeleteAppUserExecute executes the request
	DeleteAppUserExecute(r ApiDeleteAppUserRequest) (*http.Response, error)

	/*
		GetAppUser Get App User

		Get App User

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param appUserId
		@return ApiGetAppUserRequest
	*/
	GetAppUser(ctx context.Context, appUserId string) ApiGetAppUserRequest

	// GetAppUserExecute executes the request
	//  @return AppUserViewResponse
	GetAppUserExecute(r ApiGetAppUserRequest) (*AppUserViewResponse, *http.Response, error)

	/*
		ListAppUserRoles List Roles

		List an App User's Roles

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param appUserId
		@return ApiListAppUserRolesRequest
	*/
	ListAppUserRoles(ctx context.Context, appUserId string) ApiListAppUserRolesRequest

	// ListAppUserRolesExecute executes the request
	//  @return AppUserAssignListResponse
	ListAppUserRolesExecute(r ApiListAppUserRolesRequest) (*AppUserAssignListResponse, *http.Response, error)

	/*
		UpdateAppUser Update App User

		Update App User

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param appUserId
		@return ApiUpdateAppUserRequest
	*/
	UpdateAppUser(ctx context.Context, appUserId string) ApiUpdateAppUserRequest

	// UpdateAppUserExecute executes the request
	//  @return AppUserUpdateResponse
	UpdateAppUserExecute(r ApiUpdateAppUserRequest) (*AppUserUpdateResponse, *http.Response, error)
}

// AppUserAPIService AppUserAPI service
type AppUserAPIService service

type ApiDeleteAppUserRequest struct {
	ctx        context.Context
	ApiService AppUserAPI
	appUserId  string
}

func (r ApiDeleteAppUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAppUserExecute(r)
}

/*
DeleteAppUser Delete App User

Delete App User

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param appUserId
	@return ApiDeleteAppUserRequest
*/
func (a *AppUserAPIService) DeleteAppUser(ctx context.Context, appUserId string) ApiDeleteAppUserRequest {
	return ApiDeleteAppUserRequest{
		ApiService: a,
		ctx:        ctx,
		appUserId:  appUserId,
	}
}

// Execute executes the request
func (a *AppUserAPIService) DeleteAppUserExecute(r ApiDeleteAppUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppUserAPIService.DeleteAppUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/app-users/{appUserId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appUserId"+"}", url.PathEscape(parameterValueToString(r.appUserId, "appUserId")), -1)

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

type ApiGetAppUserRequest struct {
	ctx        context.Context
	ApiService AppUserAPI
	appUserId  string
}

func (r ApiGetAppUserRequest) Execute() (*AppUserViewResponse, *http.Response, error) {
	return r.ApiService.GetAppUserExecute(r)
}

/*
GetAppUser Get App User

Get App User

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param appUserId
	@return ApiGetAppUserRequest
*/
func (a *AppUserAPIService) GetAppUser(ctx context.Context, appUserId string) ApiGetAppUserRequest {
	return ApiGetAppUserRequest{
		ApiService: a,
		ctx:        ctx,
		appUserId:  appUserId,
	}
}

// Execute executes the request
//
//	@return AppUserViewResponse
func (a *AppUserAPIService) GetAppUserExecute(r ApiGetAppUserRequest) (*AppUserViewResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppUserViewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppUserAPIService.GetAppUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/app-users/{appUserId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appUserId"+"}", url.PathEscape(parameterValueToString(r.appUserId, "appUserId")), -1)

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

type ApiListAppUserRolesRequest struct {
	ctx        context.Context
	ApiService AppUserAPI
	appUserId  string
}

func (r ApiListAppUserRolesRequest) Execute() (*AppUserAssignListResponse, *http.Response, error) {
	return r.ApiService.ListAppUserRolesExecute(r)
}

/*
ListAppUserRoles List Roles

List an App User's Roles

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param appUserId
	@return ApiListAppUserRolesRequest
*/
func (a *AppUserAPIService) ListAppUserRoles(ctx context.Context, appUserId string) ApiListAppUserRolesRequest {
	return ApiListAppUserRolesRequest{
		ApiService: a,
		ctx:        ctx,
		appUserId:  appUserId,
	}
}

// Execute executes the request
//
//	@return AppUserAssignListResponse
func (a *AppUserAPIService) ListAppUserRolesExecute(r ApiListAppUserRolesRequest) (*AppUserAssignListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppUserAssignListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppUserAPIService.ListAppUserRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/app-users/{appUserId}/roles"
	localVarPath = strings.Replace(localVarPath, "{"+"appUserId"+"}", url.PathEscape(parameterValueToString(r.appUserId, "appUserId")), -1)

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

type ApiUpdateAppUserRequest struct {
	ctx                  context.Context
	ApiService           AppUserAPI
	appUserId            string
	appUserUpdateRequest *AppUserUpdateRequest
}

func (r ApiUpdateAppUserRequest) AppUserUpdateRequest(appUserUpdateRequest AppUserUpdateRequest) ApiUpdateAppUserRequest {
	r.appUserUpdateRequest = &appUserUpdateRequest
	return r
}

func (r ApiUpdateAppUserRequest) Execute() (*AppUserUpdateResponse, *http.Response, error) {
	return r.ApiService.UpdateAppUserExecute(r)
}

/*
UpdateAppUser Update App User

Update App User

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param appUserId
	@return ApiUpdateAppUserRequest
*/
func (a *AppUserAPIService) UpdateAppUser(ctx context.Context, appUserId string) ApiUpdateAppUserRequest {
	return ApiUpdateAppUserRequest{
		ApiService: a,
		ctx:        ctx,
		appUserId:  appUserId,
	}
}

// Execute executes the request
//
//	@return AppUserUpdateResponse
func (a *AppUserAPIService) UpdateAppUserExecute(r ApiUpdateAppUserRequest) (*AppUserUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppUserUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppUserAPIService.UpdateAppUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/app-users/{appUserId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appUserId"+"}", url.PathEscape(parameterValueToString(r.appUserId, "appUserId")), -1)

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
	localVarPostBody = r.appUserUpdateRequest
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
