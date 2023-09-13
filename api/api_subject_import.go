/*
Synadia Control Plane

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type SubjectImportAPI interface {

	/*
		DeleteSubjectImport Delete Subject Import

		Delete Subject Import

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param subjectImportId
		@return ApiDeleteSubjectImportRequest
	*/
	DeleteSubjectImport(ctx context.Context, subjectImportId string) ApiDeleteSubjectImportRequest

	// DeleteSubjectImportExecute executes the request
	DeleteSubjectImportExecute(r ApiDeleteSubjectImportRequest) (*http.Response, error)

	/*
		GetSubjectImport Get Subject Import

		Returns Subject Import

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param subjectImportId
		@return ApiGetSubjectImportRequest
	*/
	GetSubjectImport(ctx context.Context, subjectImportId string) ApiGetSubjectImportRequest

	// GetSubjectImportExecute executes the request
	//  @return SubjectImportViewResponse
	GetSubjectImportExecute(r ApiGetSubjectImportRequest) (*SubjectImportViewResponse, *http.Response, error)

	/*
		UpdateSubjectImport Update Subject Import

		Updates Subject Import

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param subjectImportId
		@return ApiUpdateSubjectImportRequest
	*/
	UpdateSubjectImport(ctx context.Context, subjectImportId string) ApiUpdateSubjectImportRequest

	// UpdateSubjectImportExecute executes the request
	//  @return SubjectImportViewResponse
	UpdateSubjectImportExecute(r ApiUpdateSubjectImportRequest) (*SubjectImportViewResponse, *http.Response, error)
}

// SubjectImportAPIService SubjectImportAPI service
type SubjectImportAPIService service

type ApiDeleteSubjectImportRequest struct {
	ctx             context.Context
	ApiService      SubjectImportAPI
	subjectImportId string
}

func (r ApiDeleteSubjectImportRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSubjectImportExecute(r)
}

/*
DeleteSubjectImport Delete Subject Import

Delete Subject Import

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param subjectImportId
	@return ApiDeleteSubjectImportRequest
*/
func (a *SubjectImportAPIService) DeleteSubjectImport(ctx context.Context, subjectImportId string) ApiDeleteSubjectImportRequest {
	return ApiDeleteSubjectImportRequest{
		ApiService:      a,
		ctx:             ctx,
		subjectImportId: subjectImportId,
	}
}

// Execute executes the request
func (a *SubjectImportAPIService) DeleteSubjectImportExecute(r ApiDeleteSubjectImportRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubjectImportAPIService.DeleteSubjectImport")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subject-imports/{subjectImportId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subjectImportId"+"}", url.PathEscape(parameterValueToString(r.subjectImportId, "subjectImportId")), -1)

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

type ApiGetSubjectImportRequest struct {
	ctx             context.Context
	ApiService      SubjectImportAPI
	subjectImportId string
}

func (r ApiGetSubjectImportRequest) Execute() (*SubjectImportViewResponse, *http.Response, error) {
	return r.ApiService.GetSubjectImportExecute(r)
}

/*
GetSubjectImport Get Subject Import

Returns Subject Import

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param subjectImportId
	@return ApiGetSubjectImportRequest
*/
func (a *SubjectImportAPIService) GetSubjectImport(ctx context.Context, subjectImportId string) ApiGetSubjectImportRequest {
	return ApiGetSubjectImportRequest{
		ApiService:      a,
		ctx:             ctx,
		subjectImportId: subjectImportId,
	}
}

// Execute executes the request
//
//	@return SubjectImportViewResponse
func (a *SubjectImportAPIService) GetSubjectImportExecute(r ApiGetSubjectImportRequest) (*SubjectImportViewResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SubjectImportViewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubjectImportAPIService.GetSubjectImport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subject-imports/{subjectImportId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subjectImportId"+"}", url.PathEscape(parameterValueToString(r.subjectImportId, "subjectImportId")), -1)

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

type ApiUpdateSubjectImportRequest struct {
	ctx                        context.Context
	ApiService                 SubjectImportAPI
	subjectImportId            string
	subjectImportUpdateRequest *SubjectImportUpdateRequest
}

func (r ApiUpdateSubjectImportRequest) SubjectImportUpdateRequest(subjectImportUpdateRequest SubjectImportUpdateRequest) ApiUpdateSubjectImportRequest {
	r.subjectImportUpdateRequest = &subjectImportUpdateRequest
	return r
}

func (r ApiUpdateSubjectImportRequest) Execute() (*SubjectImportViewResponse, *http.Response, error) {
	return r.ApiService.UpdateSubjectImportExecute(r)
}

/*
UpdateSubjectImport Update Subject Import

Updates Subject Import

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param subjectImportId
	@return ApiUpdateSubjectImportRequest
*/
func (a *SubjectImportAPIService) UpdateSubjectImport(ctx context.Context, subjectImportId string) ApiUpdateSubjectImportRequest {
	return ApiUpdateSubjectImportRequest{
		ApiService:      a,
		ctx:             ctx,
		subjectImportId: subjectImportId,
	}
}

// Execute executes the request
//
//	@return SubjectImportViewResponse
func (a *SubjectImportAPIService) UpdateSubjectImportExecute(r ApiUpdateSubjectImportRequest) (*SubjectImportViewResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SubjectImportViewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubjectImportAPIService.UpdateSubjectImport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subject-imports/{subjectImportId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subjectImportId"+"}", url.PathEscape(parameterValueToString(r.subjectImportId, "subjectImportId")), -1)

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
	localVarPostBody = r.subjectImportUpdateRequest
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
