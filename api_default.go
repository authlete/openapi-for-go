/*
Authlete API

Authlete API Document.

API version: 2.3.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

type DefaultApi interface {

	/*
		MiscEchoApi /api/misc/echo API

		Echo test endpoint. Will return all path parameters in the request


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiMiscEchoApiRequest
	*/
	MiscEchoApi(ctx context.Context) ApiMiscEchoApiRequest

	// MiscEchoApiExecute executes the request
	MiscEchoApiExecute(r ApiMiscEchoApiRequest) (*http.Response, error)
}

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiMiscEchoApiRequest struct {
	ctx        context.Context
	ApiService DefaultApi
}

func (r ApiMiscEchoApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.MiscEchoApiExecute(r)
}

/*
MiscEchoApi /api/misc/echo API

Echo test endpoint. Will return all path parameters in the request

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiMiscEchoApiRequest
*/
func (a *DefaultApiService) MiscEchoApi(ctx context.Context) ApiMiscEchoApiRequest {
	return ApiMiscEchoApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *DefaultApiService) MiscEchoApiExecute(r ApiMiscEchoApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.MiscEchoApi")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/misc/echo"

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
