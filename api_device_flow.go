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

type DeviceFlowApi interface {

	/*
			DeviceAuthorizationApi /api/device/authorization API

			This API parses request parameters of a [device authorization request](https://datatracker.ietf.org/doc/html/rfc8628#section-3.1)
		and returns necessary data for the authorization server implementation to process the device authorization
		request further.

		<br>
		<details>
		<summary>Description</summary>

		This API is supposed to be called from the within the implementation of the device authorization
		endpoint of the service. The service implementation should retrieve the value of `action` from the
		response and take the following steps according to the value.

		**INTERNAL_SERVER_ERROR**

		When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the API call from the authorization
		server implementation was wrong or that an error occurred in Authlete.

		In either case, from a viewpoint of the client application, it is an error on the server side.
		Therefore, the authorization server implementation should generate a response to the client application
		with "500 Internal Server Error"s and `application/json`.

		The value of `responseContent` is a JSON string which describes t he error, so it can be
		used as the entity body of the response.

		The following illustrates the response which the authorization server implementation should generate
		and return to the client application.

		```
		HTTP/1.1 500 Internal Server Error
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		**BAD_REQUEST**

		When the value of `action` is `BAD_REQUEST`, it means that the request from the client application
		is wrong.

		The authorization server implementation should generate a response to the client application with
		"400 Bad Request" and `application/json`.

		The value of `responseContent` is a JSON string which describes the error, so it can be used as
		the entity body of the response.

		The following illustrates the response which the service implementation should generate and return
		to the client application.

		```
		HTTP/1.1 400 Bad Request
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		**UNAUTHORIZED**

		When the value of `action` is `UNAUTHORIZED`, it means that client authentication of the device authorization
		request failed.

		The authorization server implementation should generate a response to the client application with
		"401 Unauthorized" and `application/json`.

		The value of `responseContent` is a JSON string which describes the error, so it can be used as
		the entity body of the response.

		The following illustrates the response which the service implementation must generate and return
		to the client application.

		```
		HTTP/1.1 401 Unauthorized
		WWW-Authenticate: (challenge)
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		**OK**

		When the value of `action` is `OK`, it means that the device authorization request from the client
		application is valid.

		The authorization server implementation should generate a response to the client application with
		"200 OK" and `application/json`.

		The `responseContent` is a JSON string which can be used as the entity body of the response.

		The following illustrates the response which the authorization server implementation should generate
		and return to the client application.
		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiDeviceAuthorizationApiRequest
	*/
	DeviceAuthorizationApi(ctx context.Context) ApiDeviceAuthorizationApiRequest

	// DeviceAuthorizationApiExecute executes the request
	//  @return DeviceAuthorizationResponse
	DeviceAuthorizationApiExecute(r ApiDeviceAuthorizationApiRequest) (*DeviceAuthorizationResponse, *http.Response, error)

	/*
			DeviceCompleteApi /api/device/complete API

			This API returns information about what action the authorization server should take after it receives
		the result of end-user's decision about whether the end-user has approved or rejected a client
		application's request.

		<br>
		<details>
		<summary>Description</summary>

		In the device flow, an end-user accesses the verification endpoint of the authorization server where
		she interacts with the verification endpoint and inputs a user code. The verification endpoint checks
		if the user code is valid and then asks the end-user whether she approves or rejects the authorization
		request which the user code represents.

		After the authorization server receives the decision of the end-user, it should call Authlete's
		`/device/complete` API to tell Authlete the decision.

		When the end-user was authenticated and authorization was granted to the client by the end-user,
		the authorization server should call the API with `result=AUTHORIZED`. In this successful case,
		the subject request parameter is mandatory. The API will update the database record so that `/auth/token`
		API can generate an access token later.

		If the `scope` parameter of the device authorization request included the openid scope, an ID token
		is generated. In this case, `sub`, `authTime`, `acr` and `claims` request parameters in the API
		call to `/device/complete` affect the ID token.

		When the authorization server receives the decision of the end-user and it indicates that she has
		rejected to give authorization to the client, the authorization server should call the API with
		`result=ACCESS_DENIED`. In this case, the API will update the database record so that the `/auth/token`
		API can generate an error response later. If `errorDescription` and `errorUri` request parameters
		are given to the `/device/complete` API, they will be used as the values of `error_description`
		and `error_uri` response parameters in the error response from the token endpoint.

		When the authorization server could not get decision from the end-user for some reasons, the authorization
		server should call the API with `result=TRANSACTION_FAILED`. In this error case, the API will behave
		in the same way as in the case of `ACCESS_DENIED`. The only difference is that `expired_token` is
		used as the value of the `error` response parameter instead of `access_denied`.

		After receiving a response from the `/device/complete` API, the implementation of the authorization
		server should retrieve the value of `action` from the response and take the following steps according
		to the value.

		**SERVER_ERROR**

		When the value of `action` is `SERVER_ERROR`, it means that an error occurred on Authlete side. The
		authorization server implementation should tell the end-user that something wrong happened and
		urge her to re-initiate a device flow.

		**USER_CODE_NOT_EXIST**

		When the value of `action` is `USER_CODE_NOT_EXIST`, it means that the user code included in the API
		call does not exist. The authorization server implementation should tell the end-user that the user
		code has been invalidated and urge her to re-initiate a device flow.

		**USER_CODE_EXPIRED**

		When the value of `action` is `USER_CODE_EXPIRED`,  it means that the user code included in the API
		call has expired. The authorization server implementation should tell the end-user that the user
		code has expired and urge her to re-initiate a device flow.

		**INVALID_REQUEST**

		When the value of `action` is `INVALID_REQUEST`, it means that the API call is invalid. Probably,
		the authorization server implementation has some bugs.

		**SUCCESS**

		When the value of `action` is `SUCCESS`, it means that the API call has been processed successfully.
		The authorization server should return a successful response to the web browser the end-user is
		using.
		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiDeviceCompleteApiRequest
	*/
	DeviceCompleteApi(ctx context.Context) ApiDeviceCompleteApiRequest

	// DeviceCompleteApiExecute executes the request
	//  @return DeviceCompleteResponse
	DeviceCompleteApiExecute(r ApiDeviceCompleteApiRequest) (*DeviceCompleteResponse, *http.Response, error)

	/*
			DeviceVerificationApi /api/device/verification API

			The API returns information associated with a user code.

		<br>
		<details>
		<summary>Description</summary>

		After receiving a response from the device authorization endpoint of the authorization server,
		the client application shows the end-user the user code and the verification URI which are included
		in the device authorization response. Then, the end-user will access the verification URI using
		a web browser on another device (typically, a smart phone). In normal implementations, the verification
		endpoint will return an HTML page with an input form where the end-user inputs a user code. The
		authorization server will receive a user code from the form.

		After receiving a user code, the authorization server should call Authlete's `/device/verification`
		API with the user code. And then, the authorization server implementation should retrieve the value
		of `action` parameter from the API response and take the following steps according to the value.

		**SERVER_ERROR**

		When the value of `action` is `SERVER_ERROR`, it means that an error occurred on Authlete side. The
		authorization server implementation should tell the end-user that something wrong happened and
		urge her to re-initiate a device flow.

		**NOT_EXIST**

		When the value of `action` is `NOT_EXIST`, it means that the user code does not exist. The authorization
		server implementation should tell the end-user that the user code is invalid and urge her to retry
		to input a valid user code.

		**EXPIRED**

		When the value of `action` is `EXPIRED`, it means that the user code has expired. The authorization
		server implementation should tell the end-user that the user code has expired and urge her to
		re-initiate a device flow.

		**VALID**

		When the value of `action` is `VALID`, it means that the user code exists, has not expired, and
		belongs to the service. The authorization server implementation should interact with the end-user
		to ask whether she approves or rejects the authorization request from the device.
		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiDeviceVerificationApiRequest
	*/
	DeviceVerificationApi(ctx context.Context) ApiDeviceVerificationApiRequest

	// DeviceVerificationApiExecute executes the request
	//  @return DeviceVerificationResponse
	DeviceVerificationApiExecute(r ApiDeviceVerificationApiRequest) (*DeviceVerificationResponse, *http.Response, error)
}

// DeviceFlowApiService DeviceFlowApi service
type DeviceFlowApiService service

type ApiDeviceAuthorizationApiRequest struct {
	ctx                        context.Context
	ApiService                 DeviceFlowApi
	deviceAuthorizationRequest *DeviceAuthorizationRequest
}

func (r ApiDeviceAuthorizationApiRequest) DeviceAuthorizationRequest(deviceAuthorizationRequest DeviceAuthorizationRequest) ApiDeviceAuthorizationApiRequest {
	r.deviceAuthorizationRequest = &deviceAuthorizationRequest
	return r
}

func (r ApiDeviceAuthorizationApiRequest) Execute() (*DeviceAuthorizationResponse, *http.Response, error) {
	return r.ApiService.DeviceAuthorizationApiExecute(r)
}

/*
DeviceAuthorizationApi /api/device/authorization API

This API parses request parameters of a [device authorization request](https://datatracker.ietf.org/doc/html/rfc8628#section-3.1)
and returns necessary data for the authorization server implementation to process the device authorization
request further.

<br>
<details>
<summary>Description</summary>

This API is supposed to be called from the within the implementation of the device authorization
endpoint of the service. The service implementation should retrieve the value of `action` from the
response and take the following steps according to the value.

**INTERNAL_SERVER_ERROR**

When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the API call from the authorization
server implementation was wrong or that an error occurred in Authlete.

In either case, from a viewpoint of the client application, it is an error on the server side.
Therefore, the authorization server implementation should generate a response to the client application
with "500 Internal Server Error"s and `application/json`.

The value of `responseContent` is a JSON string which describes t he error, so it can be
used as the entity body of the response.

The following illustrates the response which the authorization server implementation should generate
and return to the client application.

```
HTTP/1.1 500 Internal Server Error
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

**BAD_REQUEST**

When the value of `action` is `BAD_REQUEST`, it means that the request from the client application
is wrong.

The authorization server implementation should generate a response to the client application with
"400 Bad Request" and `application/json`.

The value of `responseContent` is a JSON string which describes the error, so it can be used as
the entity body of the response.

The following illustrates the response which the service implementation should generate and return
to the client application.

```
HTTP/1.1 400 Bad Request
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

**UNAUTHORIZED**

When the value of `action` is `UNAUTHORIZED`, it means that client authentication of the device authorization
request failed.

The authorization server implementation should generate a response to the client application with
"401 Unauthorized" and `application/json`.

The value of `responseContent` is a JSON string which describes the error, so it can be used as
the entity body of the response.

The following illustrates the response which the service implementation must generate and return
to the client application.

```
HTTP/1.1 401 Unauthorized
WWW-Authenticate: (challenge)
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

**OK**

When the value of `action` is `OK`, it means that the device authorization request from the client
application is valid.

The authorization server implementation should generate a response to the client application with
"200 OK" and `application/json`.

The `responseContent` is a JSON string which can be used as the entity body of the response.

The following illustrates the response which the authorization server implementation should generate
and return to the client application.
</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDeviceAuthorizationApiRequest
*/
func (a *DeviceFlowApiService) DeviceAuthorizationApi(ctx context.Context) ApiDeviceAuthorizationApiRequest {
	return ApiDeviceAuthorizationApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DeviceAuthorizationResponse
func (a *DeviceFlowApiService) DeviceAuthorizationApiExecute(r ApiDeviceAuthorizationApiRequest) (*DeviceAuthorizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeviceAuthorizationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceFlowApiService.DeviceAuthorizationApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/device/authorization"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deviceAuthorizationRequest == nil {
		return localVarReturnValue, nil, reportError("deviceAuthorizationRequest is required and must be specified")
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
	localVarPostBody = r.deviceAuthorizationRequest
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

type ApiDeviceCompleteApiRequest struct {
	ctx                   context.Context
	ApiService            DeviceFlowApi
	deviceCompleteRequest *DeviceCompleteRequest
}

func (r ApiDeviceCompleteApiRequest) DeviceCompleteRequest(deviceCompleteRequest DeviceCompleteRequest) ApiDeviceCompleteApiRequest {
	r.deviceCompleteRequest = &deviceCompleteRequest
	return r
}

func (r ApiDeviceCompleteApiRequest) Execute() (*DeviceCompleteResponse, *http.Response, error) {
	return r.ApiService.DeviceCompleteApiExecute(r)
}

/*
DeviceCompleteApi /api/device/complete API

This API returns information about what action the authorization server should take after it receives
the result of end-user's decision about whether the end-user has approved or rejected a client
application's request.

<br>
<details>
<summary>Description</summary>

In the device flow, an end-user accesses the verification endpoint of the authorization server where
she interacts with the verification endpoint and inputs a user code. The verification endpoint checks
if the user code is valid and then asks the end-user whether she approves or rejects the authorization
request which the user code represents.

After the authorization server receives the decision of the end-user, it should call Authlete's
`/device/complete` API to tell Authlete the decision.

When the end-user was authenticated and authorization was granted to the client by the end-user,
the authorization server should call the API with `result=AUTHORIZED`. In this successful case,
the subject request parameter is mandatory. The API will update the database record so that `/auth/token`
API can generate an access token later.

If the `scope` parameter of the device authorization request included the openid scope, an ID token
is generated. In this case, `sub`, `authTime`, `acr` and `claims` request parameters in the API
call to `/device/complete` affect the ID token.

When the authorization server receives the decision of the end-user and it indicates that she has
rejected to give authorization to the client, the authorization server should call the API with
`result=ACCESS_DENIED`. In this case, the API will update the database record so that the `/auth/token`
API can generate an error response later. If `errorDescription` and `errorUri` request parameters
are given to the `/device/complete` API, they will be used as the values of `error_description`
and `error_uri` response parameters in the error response from the token endpoint.

When the authorization server could not get decision from the end-user for some reasons, the authorization
server should call the API with `result=TRANSACTION_FAILED`. In this error case, the API will behave
in the same way as in the case of `ACCESS_DENIED`. The only difference is that `expired_token` is
used as the value of the `error` response parameter instead of `access_denied`.

After receiving a response from the `/device/complete` API, the implementation of the authorization
server should retrieve the value of `action` from the response and take the following steps according
to the value.

**SERVER_ERROR**

When the value of `action` is `SERVER_ERROR`, it means that an error occurred on Authlete side. The
authorization server implementation should tell the end-user that something wrong happened and
urge her to re-initiate a device flow.

**USER_CODE_NOT_EXIST**

When the value of `action` is `USER_CODE_NOT_EXIST`, it means that the user code included in the API
call does not exist. The authorization server implementation should tell the end-user that the user
code has been invalidated and urge her to re-initiate a device flow.

**USER_CODE_EXPIRED**

When the value of `action` is `USER_CODE_EXPIRED`,  it means that the user code included in the API
call has expired. The authorization server implementation should tell the end-user that the user
code has expired and urge her to re-initiate a device flow.

**INVALID_REQUEST**

When the value of `action` is `INVALID_REQUEST`, it means that the API call is invalid. Probably,
the authorization server implementation has some bugs.

**SUCCESS**

When the value of `action` is `SUCCESS`, it means that the API call has been processed successfully.
The authorization server should return a successful response to the web browser the end-user is
using.
</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDeviceCompleteApiRequest
*/
func (a *DeviceFlowApiService) DeviceCompleteApi(ctx context.Context) ApiDeviceCompleteApiRequest {
	return ApiDeviceCompleteApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DeviceCompleteResponse
func (a *DeviceFlowApiService) DeviceCompleteApiExecute(r ApiDeviceCompleteApiRequest) (*DeviceCompleteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeviceCompleteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceFlowApiService.DeviceCompleteApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/device/complete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deviceCompleteRequest == nil {
		return localVarReturnValue, nil, reportError("deviceCompleteRequest is required and must be specified")
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
	localVarPostBody = r.deviceCompleteRequest
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

type ApiDeviceVerificationApiRequest struct {
	ctx                       context.Context
	ApiService                DeviceFlowApi
	deviceVerificationRequest *DeviceVerificationRequest
}

func (r ApiDeviceVerificationApiRequest) DeviceVerificationRequest(deviceVerificationRequest DeviceVerificationRequest) ApiDeviceVerificationApiRequest {
	r.deviceVerificationRequest = &deviceVerificationRequest
	return r
}

func (r ApiDeviceVerificationApiRequest) Execute() (*DeviceVerificationResponse, *http.Response, error) {
	return r.ApiService.DeviceVerificationApiExecute(r)
}

/*
DeviceVerificationApi /api/device/verification API

The API returns information associated with a user code.

<br>
<details>
<summary>Description</summary>

After receiving a response from the device authorization endpoint of the authorization server,
the client application shows the end-user the user code and the verification URI which are included
in the device authorization response. Then, the end-user will access the verification URI using
a web browser on another device (typically, a smart phone). In normal implementations, the verification
endpoint will return an HTML page with an input form where the end-user inputs a user code. The
authorization server will receive a user code from the form.

After receiving a user code, the authorization server should call Authlete's `/device/verification`
API with the user code. And then, the authorization server implementation should retrieve the value
of `action` parameter from the API response and take the following steps according to the value.

**SERVER_ERROR**

When the value of `action` is `SERVER_ERROR`, it means that an error occurred on Authlete side. The
authorization server implementation should tell the end-user that something wrong happened and
urge her to re-initiate a device flow.

**NOT_EXIST**

When the value of `action` is `NOT_EXIST`, it means that the user code does not exist. The authorization
server implementation should tell the end-user that the user code is invalid and urge her to retry
to input a valid user code.

**EXPIRED**

When the value of `action` is `EXPIRED`, it means that the user code has expired. The authorization
server implementation should tell the end-user that the user code has expired and urge her to
re-initiate a device flow.

**VALID**

When the value of `action` is `VALID`, it means that the user code exists, has not expired, and
belongs to the service. The authorization server implementation should interact with the end-user
to ask whether she approves or rejects the authorization request from the device.
</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDeviceVerificationApiRequest
*/
func (a *DeviceFlowApiService) DeviceVerificationApi(ctx context.Context) ApiDeviceVerificationApiRequest {
	return ApiDeviceVerificationApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DeviceVerificationResponse
func (a *DeviceFlowApiService) DeviceVerificationApiExecute(r ApiDeviceVerificationApiRequest) (*DeviceVerificationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeviceVerificationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceFlowApiService.DeviceVerificationApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/device/verification"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deviceVerificationRequest == nil {
		return localVarReturnValue, nil, reportError("deviceVerificationRequest is required and must be specified")
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
	localVarPostBody = r.deviceVerificationRequest
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
