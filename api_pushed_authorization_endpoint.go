/*
Authlete API

Authlete API Document. 

API version: 2.2.25
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


type PushedAuthorizationEndpointApi interface {

	/*
	PushedAuthReqApi /api/pushed_auth_req API

	This API creates a pushed request authorization. It authenticates the client and creates a authorization_uri to be returned by the authorization server.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPushedAuthReqApiRequest
	*/
	PushedAuthReqApi(ctx context.Context) ApiPushedAuthReqApiRequest

	// PushedAuthReqApiExecute executes the request
	//  @return PushedAuthReqResponse
	PushedAuthReqApiExecute(r ApiPushedAuthReqApiRequest) (*PushedAuthReqResponse, *http.Response, error)
}

// PushedAuthorizationEndpointApiService PushedAuthorizationEndpointApi service
type PushedAuthorizationEndpointApiService service

type ApiPushedAuthReqApiRequest struct {
	ctx context.Context
	ApiService PushedAuthorizationEndpointApi
	pushedAuthReqRequest *PushedAuthReqRequest
}

func (r ApiPushedAuthReqApiRequest) PushedAuthReqRequest(pushedAuthReqRequest PushedAuthReqRequest) ApiPushedAuthReqApiRequest {
	r.pushedAuthReqRequest = &pushedAuthReqRequest
	return r
}

func (r ApiPushedAuthReqApiRequest) Execute() (*PushedAuthReqResponse, *http.Response, error) {
	return r.ApiService.PushedAuthReqApiExecute(r)
}

/*
PushedAuthReqApi /api/pushed_auth_req API

This API creates a pushed request authorization. It authenticates the client and creates a authorization_uri to be returned by the authorization server.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPushedAuthReqApiRequest
*/
func (a *PushedAuthorizationEndpointApiService) PushedAuthReqApi(ctx context.Context) ApiPushedAuthReqApiRequest {
	return ApiPushedAuthReqApiRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PushedAuthReqResponse
func (a *PushedAuthorizationEndpointApiService) PushedAuthReqApiExecute(r ApiPushedAuthReqApiRequest) (*PushedAuthReqResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PushedAuthReqResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PushedAuthorizationEndpointApiService.PushedAuthReqApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pushed_auth_req"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pushedAuthReqRequest == nil {
		return localVarReturnValue, nil, reportError("pushedAuthReqRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

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
	localVarPostBody = r.pushedAuthReqRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Result
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Result
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Result
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Result
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
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
