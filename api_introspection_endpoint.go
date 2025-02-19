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

type IntrospectionEndpointApi interface {

	/*
			AuthIntrospectionApi /api/auth/introspection API

			This API gathers information about an access token.

		<br>
		<details>
		<summary>Description</summary>

		This API is supposed to be called from within the implementations of protected resource endpoints
		of the authorization server implementation in order to get information about the access token which
		was presented by the client application.

		In general, a client application accesses a protected resource endpoint of a service with an access
		token, and the implementation of the endpoint checks whether the presented access token has enough
		privileges (= scopes) to access the protected resource before returning the protected resource to
		the client application. To achieve this flow, the endpoint implementation has to know detailed
		information about the access token. Authlete `/auth/introspection` API can be used to get such information.

		The response from `/auth/introspection` API has some parameters. Among them, it is `action` parameter
		that the authorization server implementation should check first because it denotes the next action
		that the authorization server implementation should take. According to the value of `action`, the
		authorization server implementation must take the steps described below.

		**INTERNAL_SERVER_ERROR**

		When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the request from the authorization
		server implementation was wrong or that an error occurred in Authlete.
		In either case, from the viewpoint of the client application, it is an error on the server side.
		Therefore, the service implementation should generate a response to the client application with
		HTTP status of "500 Internal Server Error". Authlete recommends `application/json` as the content
		type although OAuth 2.0 specification does not mention the format of the error response when the
		redirect URI is not usable.

		The value of `responseContent` is a string which describes the error in the format of
		[RFC 6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage), so if
		the protected resource of the service implementation wants to return an error response to the client
		application in the way that complies with RFC 6750 (in other words, if `accessTokenType` configuration
		parameter of the service is Bearer), the value of `responseContent` can be used as the value of
		`WWW-Authenticate` header.

		The following is an example response which complies with RFC 6750.

		```
		HTTP/1.1 500 Internal Server Error
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		**BAD_REQUEST**

		When the value of `action` is `BAD_REQUEST`, it means that the request from the client application
		does not contain an access token (= the request from the authorization server implementation to
		Authlete does not contain `token` request parameter).

		A response with HTTP status of "400 Bad Request" must be returned to the client application and
		the content type must be `application/json`.


		The value of `responseContent` is a string which describes the error in the format of [RFC
		6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage), so if the
		protected resource of the service implementation wants to return an error response to the client
		application in the way that complies with RFC 6750 (in other words, if `accessTokenType` configuration
		parameter of the service is `Bearer`), the value of `responseContent` can be used as the value of
		`WWW-Authenticate` header.

		The following is an example response which complies with RFC 6750.

		```
		HTTP/1.1 400 Bad Request
		WWW-Authenticate: {responseContent}
		Cache-Control: no-store
		Pragma: no-cache
		```

		**UNAUTHORIZED**

		When the value of `action` is `UNAUTHORIZED`, it means that the access token does not exist or has
		expired.

		The value of `responseContent` is a string which describes the error in the format of RFC
		6750 (OAuth 2.0 Bearer Token Usage), so if the protected resource of the service implementation
		wants to return an error response to the client application in the way that complies with [RFC
		6750](https://datatracker.ietf.org/doc/html/rfc6750) (in other words, if `accessTokenType` configuration
		parameter of the service is `Bearer`), the value of `responseContent` can be used as the value of
		`WWW-Authenticate` header.

		The following is an example response which complies with RFC 6750.

		```
		HTTP/1.1 401 Unauthorized
		WWW-Authenticate: {responseContent}
		Cache-Control: no-store
		Pragma: no-cache
		```

		**FORBIDDEN**

		When the value of `action` is `FORBIDDEN`, it means that the access token does not cover the required
		scopes or that the subject associated with the access token is different from the subject contained
		in the request.

		A response with HTTP status of "400 Bad Request" must be returned to the client application and
		the content type must be `application/json`.

		The value of `responseContent` is a string which describes the error in the format of [RFC
		6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage), so if the
		protected resource of the service implementation wants to return an error response to the client
		application in the way that complies with RFC 6750 (in other words, if `accessTokenType` configuration
		parameter of the service is Bearer), the value of `responseContent` can be used as the value of
		`WWW-Authenticate` header.

		The following is an example response which complies with RFC 6750.

		```
		HTTP/1.1 403 Forbidden
		WWW-Authenticate: {responseContent}
		Cache-Control: no-store
		Pragma: no-cache
		```

		**OK**

		When the value of `action` is `OK`, it means that the access token which the client application
		presented is valid (= exists and has not expired).

		The implementation of the protected resource endpoint is supposed to return the protected resource
		to the client application.

		When action is `OK`, the value of `responseContent` is `"Bearer error=\"invalid_request\""`. This
		is the simplest string which can be used as the value of `WWW-Authenticate` header to indicate
		"400 Bad Request". The implementation of the protected resource endpoint may use this string to
		tell the client application that the request was bad (e.g. in case necessary request parameters
		for the protected resource endpoint are missing). However, in such a case, the implementation
		should generate a more informative error message to help developers of client applications.

		The following is an example error response which complies with RFC 6750.

		```
		HTTP/1.1 400 Bad Request
		WWW-Authenticate: {responseContent}
		Cache-Control: no-store
		Pragma: no-cache
		```

		Basically, The value of `responseContent` is a string which describes the error in the format of
		[RFC 6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage). So, if
		the service has selected `Bearer` as the value of `accessTokenType` configuration parameter, the
		value of `responseContent` can be used directly as the value of `WWW-Authenticate` header. However,
		if the service has selected another different token type, the service has to generate error messages
		for itself.

		_**JWT-based access token**_

		Since version 2.1, Authlete provides a feature to issue access tokens in JWT format. This feature
		can be enabled by setting a non-null value to the `accessTokenSignAlg` property of the service
		(see the description of the Service class for details). `/api/auth/introspection` API can accept
		access tokens in JWT format. However, note that the API does not return information contained in
		a given JWT-based access token but returns information stored in the database record which corresponds
		to the given JWT-based access token. Because attributes of the database record can be modified
		after the access token is issued (for example, by using `/api/auth/token/update` API), information
		returned by `/api/auth/introspection` API and information the given JWT-based access token holds
		may be different.

		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiAuthIntrospectionApiRequest
	*/
	AuthIntrospectionApi(ctx context.Context) ApiAuthIntrospectionApiRequest

	// AuthIntrospectionApiExecute executes the request
	//  @return IntrospectionResponse
	AuthIntrospectionApiExecute(r ApiAuthIntrospectionApiRequest) (*IntrospectionResponse, *http.Response, error)

	/*
			AuthIntrospectionStandardApi /api/auth/introspection/standard API

			This API exists to help your authorization server provide its own introspection API which complies
		with [RFC 7662](https://tools.ietf.org/html/rfc7662) (OAuth 2.0 Token Introspection).

		<br>
		<details>
		<summary>Description</summary>

		This API is supposed to be called from within the implementations of the introspection endpoint
		of your service. The authorization server implementation should retrieve the value of `action` from
		the response and take the following steps according to the value.

		In general, a client application accesses a protected resource endpoint of a service with an access
		token, and the implementation of the endpoint checks whether the presented access token has enough
		privileges (= scopes) to access the protected resource before returning the protected resource to
		the client application. To achieve this flow, the endpoint implementation has to know detailed
		information about the access token. Authlete `/auth/introspection` API can be used to get such information.

		The response from `/auth/introspection` API has some parameters. Among them, it is `action` parameter
		that the authorization server implementation should check first because it denotes the next action
		that the authorization server implementation should take. According to the value of `action`, the
		authorization server implementation must take the steps described below.

		**INTERNAL_SERVER_ERROR**

		When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the request from the authorization
		server implementation was wrong or that an error occurred in Authlete.

		In either case, from the viewpoint of the client application, it is an error on the server side.
		Therefore, the service implementation should generate a response to the client application with
		HTTP status of "500 Internal Server Error".

		The value of `responseContent` is a JSON string which describes the error, so it can be used
		as the entity body of the response if you want. Note that, however, [RFC 7662](https://datatracker.ietf.org/doc/html/rfc7662) does not mention anything about the response
		body of error responses.

		The following illustrates an example response which the introspection endpoint of the authorization
		server implementation generates and returns to the client application.

		```
		HTTP/1.1 500 Internal Server Error
		Content-Type: application/json

		{responseContent}
		```

		**BAD_REQUEST**

		When the value of `action` is `BAD_REQUEST`, it means that the request from the client application
		is invalid. This happens when the request from the client did not include the token request parameter.
		See "[2.1. Introspection Request](https://datatracker.ietf.org/doc/html/rfc7662#section-2.1)" in
		RFC 7662 for details about requirements for introspection requests.

		The HTTP status of the response returned to the client application should be "400 Bad Request".

		The value of `responseContent` is a JSON string which describes the error, so it can be used
		as the entity body of the response if you want. Note that, however, [RFC 7662](https://datatracker.ietf.org/doc/html/rfc7662)
		does not mention anything about the response body of error responses.

		The following illustrates an example response which the introspection endpoint of the authorization
		server implementation generates and returns to the client application.

		```
		HTTP/1.1 400 Bad Request
		Content-Type: application/json

		{responseContent}
		```

		**OK**

		When the value of `action` is `OK`, the request from the client application is valid.

		The HTTP status of the response returned to the client application must be "200 OK" and its content
		type must be `application/json`.

		The value of `responseContent` is a JSON string which complies with the introspection response
		defined in "2.2. Introspection Response" in RFC7662.

		The following illustrates the response which the introspection endpoint of your authorization server
		implementation should generate and return to the client application.

		```
		HTTP/1.1 200 OK
		Content-Type: application/json

		{responseContent}
		```

		Note that RFC 7662 says _"To prevent token scanning attacks, **the endpoint MUST also require some
		form of authorization to access this endpoint**"_. This means that you have to protect your introspection
		endpoint in some way or other. Authlete does not care about how your introspection endpoint is protected.
		In most cases, as mentioned in RFC 7662, "401 Unauthorized" is a proper response when an introspection
		request does not satisfy authorization requirements imposed by your introspection endpoint.

		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiAuthIntrospectionStandardApiRequest
	*/
	AuthIntrospectionStandardApi(ctx context.Context) ApiAuthIntrospectionStandardApiRequest

	// AuthIntrospectionStandardApiExecute executes the request
	//  @return StandardIntrospectionResponse
	AuthIntrospectionStandardApiExecute(r ApiAuthIntrospectionStandardApiRequest) (*StandardIntrospectionResponse, *http.Response, error)
}

// IntrospectionEndpointApiService IntrospectionEndpointApi service
type IntrospectionEndpointApiService service

type ApiAuthIntrospectionApiRequest struct {
	ctx                  context.Context
	ApiService           IntrospectionEndpointApi
	introspectionRequest *IntrospectionRequest
}

func (r ApiAuthIntrospectionApiRequest) IntrospectionRequest(introspectionRequest IntrospectionRequest) ApiAuthIntrospectionApiRequest {
	r.introspectionRequest = &introspectionRequest
	return r
}

func (r ApiAuthIntrospectionApiRequest) Execute() (*IntrospectionResponse, *http.Response, error) {
	return r.ApiService.AuthIntrospectionApiExecute(r)
}

/*
AuthIntrospectionApi /api/auth/introspection API

This API gathers information about an access token.

<br>
<details>
<summary>Description</summary>

This API is supposed to be called from within the implementations of protected resource endpoints
of the authorization server implementation in order to get information about the access token which
was presented by the client application.

In general, a client application accesses a protected resource endpoint of a service with an access
token, and the implementation of the endpoint checks whether the presented access token has enough
privileges (= scopes) to access the protected resource before returning the protected resource to
the client application. To achieve this flow, the endpoint implementation has to know detailed
information about the access token. Authlete `/auth/introspection` API can be used to get such information.

The response from `/auth/introspection` API has some parameters. Among them, it is `action` parameter
that the authorization server implementation should check first because it denotes the next action
that the authorization server implementation should take. According to the value of `action`, the
authorization server implementation must take the steps described below.

**INTERNAL_SERVER_ERROR**

When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the request from the authorization
server implementation was wrong or that an error occurred in Authlete.
In either case, from the viewpoint of the client application, it is an error on the server side.
Therefore, the service implementation should generate a response to the client application with
HTTP status of "500 Internal Server Error". Authlete recommends `application/json` as the content
type although OAuth 2.0 specification does not mention the format of the error response when the
redirect URI is not usable.

The value of `responseContent` is a string which describes the error in the format of
[RFC 6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage), so if
the protected resource of the service implementation wants to return an error response to the client
application in the way that complies with RFC 6750 (in other words, if `accessTokenType` configuration
parameter of the service is Bearer), the value of `responseContent` can be used as the value of
`WWW-Authenticate` header.

The following is an example response which complies with RFC 6750.

```
HTTP/1.1 500 Internal Server Error
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

**BAD_REQUEST**

When the value of `action` is `BAD_REQUEST`, it means that the request from the client application
does not contain an access token (= the request from the authorization server implementation to
Authlete does not contain `token` request parameter).

A response with HTTP status of "400 Bad Request" must be returned to the client application and
the content type must be `application/json`.

The value of `responseContent` is a string which describes the error in the format of [RFC
6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage), so if the
protected resource of the service implementation wants to return an error response to the client
application in the way that complies with RFC 6750 (in other words, if `accessTokenType` configuration
parameter of the service is `Bearer`), the value of `responseContent` can be used as the value of
`WWW-Authenticate` header.

The following is an example response which complies with RFC 6750.

```
HTTP/1.1 400 Bad Request
WWW-Authenticate: {responseContent}
Cache-Control: no-store
Pragma: no-cache
```

**UNAUTHORIZED**

When the value of `action` is `UNAUTHORIZED`, it means that the access token does not exist or has
expired.

The value of `responseContent` is a string which describes the error in the format of RFC
6750 (OAuth 2.0 Bearer Token Usage), so if the protected resource of the service implementation
wants to return an error response to the client application in the way that complies with [RFC
6750](https://datatracker.ietf.org/doc/html/rfc6750) (in other words, if `accessTokenType` configuration
parameter of the service is `Bearer`), the value of `responseContent` can be used as the value of
`WWW-Authenticate` header.

The following is an example response which complies with RFC 6750.

```
HTTP/1.1 401 Unauthorized
WWW-Authenticate: {responseContent}
Cache-Control: no-store
Pragma: no-cache
```

**FORBIDDEN**

When the value of `action` is `FORBIDDEN`, it means that the access token does not cover the required
scopes or that the subject associated with the access token is different from the subject contained
in the request.

A response with HTTP status of "400 Bad Request" must be returned to the client application and
the content type must be `application/json`.

The value of `responseContent` is a string which describes the error in the format of [RFC
6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage), so if the
protected resource of the service implementation wants to return an error response to the client
application in the way that complies with RFC 6750 (in other words, if `accessTokenType` configuration
parameter of the service is Bearer), the value of `responseContent` can be used as the value of
`WWW-Authenticate` header.

The following is an example response which complies with RFC 6750.

```
HTTP/1.1 403 Forbidden
WWW-Authenticate: {responseContent}
Cache-Control: no-store
Pragma: no-cache
```

**OK**

When the value of `action` is `OK`, it means that the access token which the client application
presented is valid (= exists and has not expired).

The implementation of the protected resource endpoint is supposed to return the protected resource
to the client application.

When action is `OK`, the value of `responseContent` is `"Bearer error=\"invalid_request\""`. This
is the simplest string which can be used as the value of `WWW-Authenticate` header to indicate
"400 Bad Request". The implementation of the protected resource endpoint may use this string to
tell the client application that the request was bad (e.g. in case necessary request parameters
for the protected resource endpoint are missing). However, in such a case, the implementation
should generate a more informative error message to help developers of client applications.

The following is an example error response which complies with RFC 6750.

```
HTTP/1.1 400 Bad Request
WWW-Authenticate: {responseContent}
Cache-Control: no-store
Pragma: no-cache
```

Basically, The value of `responseContent` is a string which describes the error in the format of
[RFC 6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage). So, if
the service has selected `Bearer` as the value of `accessTokenType` configuration parameter, the
value of `responseContent` can be used directly as the value of `WWW-Authenticate` header. However,
if the service has selected another different token type, the service has to generate error messages
for itself.

_**JWT-based access token**_

Since version 2.1, Authlete provides a feature to issue access tokens in JWT format. This feature
can be enabled by setting a non-null value to the `accessTokenSignAlg` property of the service
(see the description of the Service class for details). `/api/auth/introspection` API can accept
access tokens in JWT format. However, note that the API does not return information contained in
a given JWT-based access token but returns information stored in the database record which corresponds
to the given JWT-based access token. Because attributes of the database record can be modified
after the access token is issued (for example, by using `/api/auth/token/update` API), information
returned by `/api/auth/introspection` API and information the given JWT-based access token holds
may be different.

</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAuthIntrospectionApiRequest
*/
func (a *IntrospectionEndpointApiService) AuthIntrospectionApi(ctx context.Context) ApiAuthIntrospectionApiRequest {
	return ApiAuthIntrospectionApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IntrospectionResponse
func (a *IntrospectionEndpointApiService) AuthIntrospectionApiExecute(r ApiAuthIntrospectionApiRequest) (*IntrospectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IntrospectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntrospectionEndpointApiService.AuthIntrospectionApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/auth/introspection"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.introspectionRequest == nil {
		return localVarReturnValue, nil, reportError("introspectionRequest is required and must be specified")
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
	localVarPostBody = r.introspectionRequest
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

type ApiAuthIntrospectionStandardApiRequest struct {
	ctx                          context.Context
	ApiService                   IntrospectionEndpointApi
	standardIntrospectionRequest *StandardIntrospectionRequest
}

func (r ApiAuthIntrospectionStandardApiRequest) StandardIntrospectionRequest(standardIntrospectionRequest StandardIntrospectionRequest) ApiAuthIntrospectionStandardApiRequest {
	r.standardIntrospectionRequest = &standardIntrospectionRequest
	return r
}

func (r ApiAuthIntrospectionStandardApiRequest) Execute() (*StandardIntrospectionResponse, *http.Response, error) {
	return r.ApiService.AuthIntrospectionStandardApiExecute(r)
}

/*
AuthIntrospectionStandardApi /api/auth/introspection/standard API

This API exists to help your authorization server provide its own introspection API which complies
with [RFC 7662](https://tools.ietf.org/html/rfc7662) (OAuth 2.0 Token Introspection).

<br>
<details>
<summary>Description</summary>

This API is supposed to be called from within the implementations of the introspection endpoint
of your service. The authorization server implementation should retrieve the value of `action` from
the response and take the following steps according to the value.

In general, a client application accesses a protected resource endpoint of a service with an access
token, and the implementation of the endpoint checks whether the presented access token has enough
privileges (= scopes) to access the protected resource before returning the protected resource to
the client application. To achieve this flow, the endpoint implementation has to know detailed
information about the access token. Authlete `/auth/introspection` API can be used to get such information.

The response from `/auth/introspection` API has some parameters. Among them, it is `action` parameter
that the authorization server implementation should check first because it denotes the next action
that the authorization server implementation should take. According to the value of `action`, the
authorization server implementation must take the steps described below.

**INTERNAL_SERVER_ERROR**

When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the request from the authorization
server implementation was wrong or that an error occurred in Authlete.

In either case, from the viewpoint of the client application, it is an error on the server side.
Therefore, the service implementation should generate a response to the client application with
HTTP status of "500 Internal Server Error".

The value of `responseContent` is a JSON string which describes the error, so it can be used
as the entity body of the response if you want. Note that, however, [RFC 7662](https://datatracker.ietf.org/doc/html/rfc7662) does not mention anything about the response
body of error responses.

The following illustrates an example response which the introspection endpoint of the authorization
server implementation generates and returns to the client application.

```
HTTP/1.1 500 Internal Server Error
Content-Type: application/json

{responseContent}
```

**BAD_REQUEST**

When the value of `action` is `BAD_REQUEST`, it means that the request from the client application
is invalid. This happens when the request from the client did not include the token request parameter.
See "[2.1. Introspection Request](https://datatracker.ietf.org/doc/html/rfc7662#section-2.1)" in
RFC 7662 for details about requirements for introspection requests.

The HTTP status of the response returned to the client application should be "400 Bad Request".

The value of `responseContent` is a JSON string which describes the error, so it can be used
as the entity body of the response if you want. Note that, however, [RFC 7662](https://datatracker.ietf.org/doc/html/rfc7662)
does not mention anything about the response body of error responses.

The following illustrates an example response which the introspection endpoint of the authorization
server implementation generates and returns to the client application.

```
HTTP/1.1 400 Bad Request
Content-Type: application/json

{responseContent}
```

**OK**

When the value of `action` is `OK`, the request from the client application is valid.

The HTTP status of the response returned to the client application must be "200 OK" and its content
type must be `application/json`.

The value of `responseContent` is a JSON string which complies with the introspection response
defined in "2.2. Introspection Response" in RFC7662.

The following illustrates the response which the introspection endpoint of your authorization server
implementation should generate and return to the client application.

```
HTTP/1.1 200 OK
Content-Type: application/json

{responseContent}
```

Note that RFC 7662 says _"To prevent token scanning attacks, **the endpoint MUST also require some
form of authorization to access this endpoint**"_. This means that you have to protect your introspection
endpoint in some way or other. Authlete does not care about how your introspection endpoint is protected.
In most cases, as mentioned in RFC 7662, "401 Unauthorized" is a proper response when an introspection
request does not satisfy authorization requirements imposed by your introspection endpoint.

</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAuthIntrospectionStandardApiRequest
*/
func (a *IntrospectionEndpointApiService) AuthIntrospectionStandardApi(ctx context.Context) ApiAuthIntrospectionStandardApiRequest {
	return ApiAuthIntrospectionStandardApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StandardIntrospectionResponse
func (a *IntrospectionEndpointApiService) AuthIntrospectionStandardApiExecute(r ApiAuthIntrospectionStandardApiRequest) (*StandardIntrospectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StandardIntrospectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntrospectionEndpointApiService.AuthIntrospectionStandardApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/auth/introspection/standard"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.standardIntrospectionRequest == nil {
		return localVarReturnValue, nil, reportError("standardIntrospectionRequest is required and must be specified")
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
	localVarPostBody = r.standardIntrospectionRequest
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
