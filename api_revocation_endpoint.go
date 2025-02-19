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

type RevocationEndpointApi interface {

	/*
			AuthRevocationApi /api/auth/revocation API

			This API revokes access tokens and refresh tokens.

		<br>
		<details>
		<summary>Description</summary>

		This API is supposed to be called from within the implementation of the revocation endpoint ([RFC
		7009](tools.ietf.org/html/rfc7009)) of the authorization server implementation in order to revoke
		access tokens and refresh tokens.

		The response from `/auth/revocation` API has some parameters. Among them, it is `action` parameter
		that the authorization server implementation should check first because it denotes the next action
		that the authorization server implementation should take. According to the value of `action`, the
		authorization server implementation must take the steps described below.

		**INTERNAL_SERVER_ERROR**

		When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the request from the authorization
		server implementation was wrong or that an error occurred in Authlete.
		In either case, from the viewpoint of the client application, it is an error on the server side.
		Therefore, the service implementation should generate a response to the client application with
		HTTP status of "500 Internal Server Error".

		The value of `responseContent` is a JSON string which describes the error, so it can be
		used as the entity body of the response.

		The following illustrates the response which the service implementation should generate and return
		to the client application.

		```
		HTTP/1.1 500 Internal Server Error
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		**INVALID_CLIENT**

		When the value of `action` is `INVALID_CLIENT`, it means that authentication of the client failed.
		In this case, the HTTP status of the response to the client application is either "400 Bad Request"
		or "401 Unauthorized". The description about `invalid_client` shown below is an excerpt from [RFC
		6749](https://datatracker.ietf.org/doc/html/rfc6749).

		`invalid_client`

		> Client authentication failed (e.g., unknown client, no client authentication included, or unsupported
		authentication method). The authorization server MAY return an HTTP 401 (Unauthorized) status code
		to indicate which HTTP authentication schemes are supported. If the client attempted to authenticate
		via the `Authorization` request header field, the authorization server MUST respond with an HTTP
		401 (Unauthorized) status code and include the `WWW-Authenticate` response header field matching
		the authentication scheme used by the client.

		In either case, the value of `responseContent` is a JSON string which can be used as the entity
		body of the response to the client application.

		The following illustrates the response which the service implementation should generate and return
		to the client application.

		```
		HTTP/1.1 400 Bad Request
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		<br>

		```
		HTTP/1.1 401 Unauthorized
		WWW-Authenticate: {challenge}
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		**BAD_REQUEST**

		When the value of `action` is `BAD_REQUEST`, it means that the request from the client application
		is invalid.

		The HTTP status of the response returned to the client application must be "400 Bad Request" and
		the content type must be `application/json`. [RFC 7009](https://datatracker.ietf.org/doc/html/rfc7009),
		[2.2.1. Error Respons](https://datatracker.ietf.org/doc/html/rfc7009#section-2.2.1) states "The
		error presentation conforms to the definition in [Section 5.2](https://datatracker.ietf.org/doc/html/rfc6749#section-5.2)
		of [[RFC 6749](https://datatracker.ietf.org/doc/html/rfc6749)]."

		The value of `responseContent` is a JSON string which describes the error, so it can be used
		as the entity body of the response.

		The following illustrates the response which the authorization server implementation should generate
		and return to the client application.

		```
		HTTP/1.1 400 Bad Request
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		**OK**

		When the value of `action` is `OK`, it means that the request from the client application is valid
		and the presented token has been revoked successfully or if the client submitted an invalid token.
		Note that invalid tokens do not cause an error. See [2.2. Revocation Response](https://datatracker.ietf.org/doc/html/rfc7009#section-2.2) for details.

		The HTTP status of the response returned to the client application must be 200 OK.

		If the original request from the client application contains callback request parameter and its
		value is not empty, the content type should be `application/javascript` and the content should be
		a JavaScript snippet for JSONP.

		The value of `responseContent` is JavaScript snippet if the original request from the client application
		contains callback request parameter and its value is not empty. Otherwise, the value of `responseContent`
		is `null`.

		```
		HTTP/1.1 200 OK
		Content-Type: application/javascript
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```
		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiAuthRevocationApiRequest
	*/
	AuthRevocationApi(ctx context.Context) ApiAuthRevocationApiRequest

	// AuthRevocationApiExecute executes the request
	//  @return RevocationResponse
	AuthRevocationApiExecute(r ApiAuthRevocationApiRequest) (*RevocationResponse, *http.Response, error)
}

// RevocationEndpointApiService RevocationEndpointApi service
type RevocationEndpointApiService service

type ApiAuthRevocationApiRequest struct {
	ctx               context.Context
	ApiService        RevocationEndpointApi
	revocationRequest *RevocationRequest
}

func (r ApiAuthRevocationApiRequest) RevocationRequest(revocationRequest RevocationRequest) ApiAuthRevocationApiRequest {
	r.revocationRequest = &revocationRequest
	return r
}

func (r ApiAuthRevocationApiRequest) Execute() (*RevocationResponse, *http.Response, error) {
	return r.ApiService.AuthRevocationApiExecute(r)
}

/*
AuthRevocationApi /api/auth/revocation API

This API revokes access tokens and refresh tokens.

<br>
<details>
<summary>Description</summary>

This API is supposed to be called from within the implementation of the revocation endpoint ([RFC
7009](tools.ietf.org/html/rfc7009)) of the authorization server implementation in order to revoke
access tokens and refresh tokens.

The response from `/auth/revocation` API has some parameters. Among them, it is `action` parameter
that the authorization server implementation should check first because it denotes the next action
that the authorization server implementation should take. According to the value of `action`, the
authorization server implementation must take the steps described below.

**INTERNAL_SERVER_ERROR**

When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the request from the authorization
server implementation was wrong or that an error occurred in Authlete.
In either case, from the viewpoint of the client application, it is an error on the server side.
Therefore, the service implementation should generate a response to the client application with
HTTP status of "500 Internal Server Error".

The value of `responseContent` is a JSON string which describes the error, so it can be
used as the entity body of the response.

The following illustrates the response which the service implementation should generate and return
to the client application.

```
HTTP/1.1 500 Internal Server Error
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

**INVALID_CLIENT**

When the value of `action` is `INVALID_CLIENT`, it means that authentication of the client failed.
In this case, the HTTP status of the response to the client application is either "400 Bad Request"
or "401 Unauthorized". The description about `invalid_client` shown below is an excerpt from [RFC
6749](https://datatracker.ietf.org/doc/html/rfc6749).

`invalid_client`

> Client authentication failed (e.g., unknown client, no client authentication included, or unsupported
authentication method). The authorization server MAY return an HTTP 401 (Unauthorized) status code
to indicate which HTTP authentication schemes are supported. If the client attempted to authenticate
via the `Authorization` request header field, the authorization server MUST respond with an HTTP
401 (Unauthorized) status code and include the `WWW-Authenticate` response header field matching
the authentication scheme used by the client.

In either case, the value of `responseContent` is a JSON string which can be used as the entity
body of the response to the client application.

The following illustrates the response which the service implementation should generate and return
to the client application.

```
HTTP/1.1 400 Bad Request
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

<br>

```
HTTP/1.1 401 Unauthorized
WWW-Authenticate: {challenge}
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

**BAD_REQUEST**

When the value of `action` is `BAD_REQUEST`, it means that the request from the client application
is invalid.

The HTTP status of the response returned to the client application must be "400 Bad Request" and
the content type must be `application/json`. [RFC 7009](https://datatracker.ietf.org/doc/html/rfc7009),
[2.2.1. Error Respons](https://datatracker.ietf.org/doc/html/rfc7009#section-2.2.1) states "The
error presentation conforms to the definition in [Section 5.2](https://datatracker.ietf.org/doc/html/rfc6749#section-5.2)
of [[RFC 6749](https://datatracker.ietf.org/doc/html/rfc6749)]."

The value of `responseContent` is a JSON string which describes the error, so it can be used
as the entity body of the response.

The following illustrates the response which the authorization server implementation should generate
and return to the client application.

```
HTTP/1.1 400 Bad Request
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

**OK**

When the value of `action` is `OK`, it means that the request from the client application is valid
and the presented token has been revoked successfully or if the client submitted an invalid token.
Note that invalid tokens do not cause an error. See [2.2. Revocation Response](https://datatracker.ietf.org/doc/html/rfc7009#section-2.2) for details.

The HTTP status of the response returned to the client application must be 200 OK.

If the original request from the client application contains callback request parameter and its
value is not empty, the content type should be `application/javascript` and the content should be
a JavaScript snippet for JSONP.

The value of `responseContent` is JavaScript snippet if the original request from the client application
contains callback request parameter and its value is not empty. Otherwise, the value of `responseContent`
is `null`.

```
HTTP/1.1 200 OK
Content-Type: application/javascript
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```
</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAuthRevocationApiRequest
*/
func (a *RevocationEndpointApiService) AuthRevocationApi(ctx context.Context) ApiAuthRevocationApiRequest {
	return ApiAuthRevocationApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RevocationResponse
func (a *RevocationEndpointApiService) AuthRevocationApiExecute(r ApiAuthRevocationApiRequest) (*RevocationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RevocationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RevocationEndpointApiService.AuthRevocationApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/auth/revocation"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.revocationRequest == nil {
		return localVarReturnValue, nil, reportError("revocationRequest is required and must be specified")
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
	localVarPostBody = r.revocationRequest
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
