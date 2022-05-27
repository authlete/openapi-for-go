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
	PushedAuthApi /api/pushed_auth_req API

	This API creates a pushed request authorization. It authenticates the client and creates a authorization_uri to be returned by the authorization server.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPushedAuthApiRequest
	*/
	PushedAuthApi(ctx context.Context) ApiPushedAuthApiRequest

	// PushedAuthApiExecute executes the request
	//  @return PushedAuthorizationResponse
	PushedAuthApiExecute(r ApiPushedAuthApiRequest) (*PushedAuthorizationResponse, *http.Response, error)
}

// PushedAuthorizationEndpointApiService PushedAuthorizationEndpointApi service
type PushedAuthorizationEndpointApiService service

type ApiPushedAuthApiRequest struct {
	ctx context.Context
	ApiService PushedAuthorizationEndpointApi
	pushedAuthorizationRequest *PushedAuthorizationRequest
}

func (r ApiPushedAuthApiRequest) PushedAuthorizationRequest(pushedAuthorizationRequest PushedAuthorizationRequest) ApiPushedAuthApiRequest {
	r.pushedAuthorizationRequest = &pushedAuthorizationRequest
	return r
}

func (r ApiPushedAuthApiRequest) Execute() (*PushedAuthorizationResponse, *http.Response, error) {
	return r.ApiService.PushedAuthApiExecute(r)
}

/*
PushedAuthApi /api/pushed_auth_req API

This API creates a pushed request authorization. It authenticates the client and creates a authorization_uri to be returned by the authorization server.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPushedAuthApiRequest
*/
func (a *PushedAuthorizationEndpointApiService) PushedAuthApi(ctx context.Context) ApiPushedAuthApiRequest {
	return ApiPushedAuthApiRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PushedAuthorizationResponse
func (a *PushedAuthorizationEndpointApiService) PushedAuthApiExecute(r ApiPushedAuthApiRequest) (*PushedAuthorizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PushedAuthorizationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PushedAuthorizationEndpointApiService.PushedAuthApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pushed_auth_req"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pushedAuthorizationRequest == nil {
		return localVarReturnValue, nil, reportError("pushedAuthorizationRequest is required and must be specified")
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
	localVarPostBody = r.pushedAuthorizationRequest
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
