/*
Authlete API

Authlete API Document. 

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type ClientExtensionApi interface {

	/*
	ClientExtensionRequestablesScopesDeleteApi /api/{serviceId}/client/extension/requestable_scopes/delete/{clientId} API

	Delete a requestable scopes of a client


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@param clientId A client ID. 
	@return ApiClientExtensionRequestablesScopesDeleteApiRequest
	*/
	ClientExtensionRequestablesScopesDeleteApi(ctx context.Context, serviceId string, clientId string) ApiClientExtensionRequestablesScopesDeleteApiRequest

	// ClientExtensionRequestablesScopesDeleteApiExecute executes the request
	ClientExtensionRequestablesScopesDeleteApiExecute(r ApiClientExtensionRequestablesScopesDeleteApiRequest) (*http.Response, error)

	/*
	ClientExtensionRequestablesScopesGetApi /api/{serviceId}/client/extension/requestable_scopes/get/{clientId} API

	Get the requestable scopes per client


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@param clientId A client ID. 
	@return ApiClientExtensionRequestablesScopesGetApiRequest
	*/
	ClientExtensionRequestablesScopesGetApi(ctx context.Context, serviceId string, clientId string) ApiClientExtensionRequestablesScopesGetApiRequest

	// ClientExtensionRequestablesScopesGetApiExecute executes the request
	//  @return ClientExtensionRequestableScopesGetResponse
	ClientExtensionRequestablesScopesGetApiExecute(r ApiClientExtensionRequestablesScopesGetApiRequest) (*ClientExtensionRequestableScopesGetResponse, *http.Response, error)

	/*
	ClientExtensionRequestablesScopesUpdateApi /api/{serviceId}/client/extension/requestable_scopes/update/{clientId} API

	Update requestable scopes of a client


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@param clientId A client ID. 
	@return ApiClientExtensionRequestablesScopesUpdateApiRequest
	*/
	ClientExtensionRequestablesScopesUpdateApi(ctx context.Context, serviceId string, clientId string) ApiClientExtensionRequestablesScopesUpdateApiRequest

	// ClientExtensionRequestablesScopesUpdateApiExecute executes the request
	//  @return ClientExtensionRequestableScopesUpdateResponse
	ClientExtensionRequestablesScopesUpdateApiExecute(r ApiClientExtensionRequestablesScopesUpdateApiRequest) (*ClientExtensionRequestableScopesUpdateResponse, *http.Response, error)
}

// ClientExtensionApiService ClientExtensionApi service
type ClientExtensionApiService service

type ApiClientExtensionRequestablesScopesDeleteApiRequest struct {
	ctx context.Context
	ApiService ClientExtensionApi
	serviceId string
	clientId string
}

func (r ApiClientExtensionRequestablesScopesDeleteApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.ClientExtensionRequestablesScopesDeleteApiExecute(r)
}

/*
ClientExtensionRequestablesScopesDeleteApi /api/{serviceId}/client/extension/requestable_scopes/delete/{clientId} API

Delete a requestable scopes of a client


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId A service ID.
 @param clientId A client ID. 
 @return ApiClientExtensionRequestablesScopesDeleteApiRequest
*/
func (a *ClientExtensionApiService) ClientExtensionRequestablesScopesDeleteApi(ctx context.Context, serviceId string, clientId string) ApiClientExtensionRequestablesScopesDeleteApiRequest {
	return ApiClientExtensionRequestablesScopesDeleteApiRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
		clientId: clientId,
	}
}

// Execute executes the request
func (a *ClientExtensionApiService) ClientExtensionRequestablesScopesDeleteApiExecute(r ApiClientExtensionRequestablesScopesDeleteApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientExtensionApiService.ClientExtensionRequestablesScopesDeleteApi")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/extension/requestable_scopes/delete/{clientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v Result
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Result
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Result
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Result
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiClientExtensionRequestablesScopesGetApiRequest struct {
	ctx context.Context
	ApiService ClientExtensionApi
	serviceId string
	clientId string
}

func (r ApiClientExtensionRequestablesScopesGetApiRequest) Execute() (*ClientExtensionRequestableScopesGetResponse, *http.Response, error) {
	return r.ApiService.ClientExtensionRequestablesScopesGetApiExecute(r)
}

/*
ClientExtensionRequestablesScopesGetApi /api/{serviceId}/client/extension/requestable_scopes/get/{clientId} API

Get the requestable scopes per client


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId A service ID.
 @param clientId A client ID. 
 @return ApiClientExtensionRequestablesScopesGetApiRequest
*/
func (a *ClientExtensionApiService) ClientExtensionRequestablesScopesGetApi(ctx context.Context, serviceId string, clientId string) ApiClientExtensionRequestablesScopesGetApiRequest {
	return ApiClientExtensionRequestablesScopesGetApiRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
		clientId: clientId,
	}
}

// Execute executes the request
//  @return ClientExtensionRequestableScopesGetResponse
func (a *ClientExtensionApiService) ClientExtensionRequestablesScopesGetApiExecute(r ApiClientExtensionRequestablesScopesGetApiRequest) (*ClientExtensionRequestableScopesGetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ClientExtensionRequestableScopesGetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientExtensionApiService.ClientExtensionRequestablesScopesGetApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/extension/requestable_scopes/get/{clientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)

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

type ApiClientExtensionRequestablesScopesUpdateApiRequest struct {
	ctx context.Context
	ApiService ClientExtensionApi
	serviceId string
	clientId string
	clientExtensionRequestableScopesUpdateRequest *ClientExtensionRequestableScopesUpdateRequest
}

func (r ApiClientExtensionRequestablesScopesUpdateApiRequest) ClientExtensionRequestableScopesUpdateRequest(clientExtensionRequestableScopesUpdateRequest ClientExtensionRequestableScopesUpdateRequest) ApiClientExtensionRequestablesScopesUpdateApiRequest {
	r.clientExtensionRequestableScopesUpdateRequest = &clientExtensionRequestableScopesUpdateRequest
	return r
}

func (r ApiClientExtensionRequestablesScopesUpdateApiRequest) Execute() (*ClientExtensionRequestableScopesUpdateResponse, *http.Response, error) {
	return r.ApiService.ClientExtensionRequestablesScopesUpdateApiExecute(r)
}

/*
ClientExtensionRequestablesScopesUpdateApi /api/{serviceId}/client/extension/requestable_scopes/update/{clientId} API

Update requestable scopes of a client


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId A service ID.
 @param clientId A client ID. 
 @return ApiClientExtensionRequestablesScopesUpdateApiRequest
*/
func (a *ClientExtensionApiService) ClientExtensionRequestablesScopesUpdateApi(ctx context.Context, serviceId string, clientId string) ApiClientExtensionRequestablesScopesUpdateApiRequest {
	return ApiClientExtensionRequestablesScopesUpdateApiRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
		clientId: clientId,
	}
}

// Execute executes the request
//  @return ClientExtensionRequestableScopesUpdateResponse
func (a *ClientExtensionApiService) ClientExtensionRequestablesScopesUpdateApiExecute(r ApiClientExtensionRequestablesScopesUpdateApiRequest) (*ClientExtensionRequestableScopesUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ClientExtensionRequestableScopesUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientExtensionApiService.ClientExtensionRequestablesScopesUpdateApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/extension/requestable_scopes/update/{clientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientExtensionRequestableScopesUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("clientExtensionRequestableScopesUpdateRequest is required and must be specified")
	}

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
	localVarPostBody = r.clientExtensionRequestableScopesUpdateRequest
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
