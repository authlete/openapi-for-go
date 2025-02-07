/*
Authlete API Explorer

<div class=\"min-h-screen bg-gray-100 dark:bg-gray-900 text-gray-900 dark:text-gray-100 p-6\">   <div class=\"flex justify-end mb-4\">     <label for=\"theme-toggle\" class=\"flex items-center cursor-pointer\">       <div class=\"relative\">Dark mode:         <input type=\"checkbox\" id=\"theme-toggle\" class=\"sr-only\" onchange=\"toggleTheme()\">         <div class=\"block bg-gray-600 w-14 h-8 rounded-full\"></div>         <div class=\"dot absolute left-1 top-1 bg-white w-6 h-6 rounded-full transition\"></div>       </div>     </label>   </div>   <header class=\"bg-green-500 dark:bg-green-700 p-4 rounded-lg text-white text-center\">     <p>       Welcome to the <strong>Authlete API documentation</strong>. Authlete is an <strong>API-first service</strong>       where every aspect of the platform is configurable via API. This explorer provides a convenient way to       authenticate and interact with the API, allowing you to see Authlete in action quickly. 🚀     </p>     <p>       At a high level, the Authlete API is grouped into two categories:     </p>     <ul class=\"list-disc list-inside\">       <li><strong>Management APIs</strong>: Enable you to manage services and clients. 🔧</li>       <li><strong>Runtime APIs</strong>: Allow you to build your own Authorization Servers or Verifiable Credential (VC)         issuers. 🔐</li>     </ul>     <p>All API endpoints are secured using access tokens issued by Authlete's Identity Provider (IdP). If you already       have an Authlete account, simply use the <em>Get Token</em> option on the Authentication page to log in and obtain       an access token for API usage. If you don't have an account yet, <a href=\"https://console.authlete.com/register\">sign up         here</a> to get started.</p>   </header>   <main>     <section id=\"api-servers\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🌐 API Servers</h2>       <p>Authlete is a global service with clusters available in multiple regions across the world.</p>       <p>Currently, our service is available in the following regions:</p>       <div class=\"grid grid-cols-2 gap-4\">         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇺🇸 US</p>         </div>         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇯🇵 JP</p>         </div>         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇪🇺 EU</p>         </div>         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇧🇷 Brazil</p>         </div>       </div>       <p>Our customers can host their data in the region that best meets their requirements.</p>       <a href=\"#servers\" class=\"block mt-4 text-green-500 dark:text-green-300 hover:underline text-center\">Select your         preferred server</a>     </section>     <section id=\"authentication\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🔑 Authentication</h2>       <p>The API Explorer requires an access token to call the API.</p>       <p>You can create the access token from the <a href=\"https://console.authlete.com\">Authlete Management Console</a> and set it in the HTTP Bearer section of Authentication page.</p>       <p>Alternatively, if you have an Authlete account, the API Explorer can log you in with your Authlete account and         automatically acquire the required access token.</p>       <div class=\"theme-admonition theme-admonition-warning admonition_o5H7 alert alert--warning\">         <div class=\"admonitionContent_Knsx\">           <p>⚠️ <strong>Important Note:</strong> When the API Explorer acquires the token after login, the access tokens             will have the same permissions as the user who logs in as part of this flow.</p>         </div>       </div>       <a href=\"#auth\" class=\"block mt-4 text-green-500 dark:text-green-300 hover:underline text-center\">Setup your         access token</a>     </section>     <section id=\"tutorials\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🎓 Tutorials</h2>       <p>If you have successfully tested the API from the API Console and want to take the next step of integrating the         API into your application, or if you want to see a sample using Authlete APIs, follow the links below. These         resources will help you understand key concepts and how to integrate Authlete API into your applications.</p>       <div class=\"mt-4\">         <a href=\"https://www.authlete.com/developers/getting_started/\"           class=\"block text-green-500 dark:text-green-300 font-bold hover:underline mb-2\">🚀 Getting Started with           Authlete</a>           </br>         <a href=\"https://www.authlete.com/developers/tutorial/signup/\"           class=\"block text-green-500 dark:text-green-300 font-bold hover:underline\">🔑 From Sign-Up to the First API           Request</a>       </div>     </section>     <section id=\"support\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🛠 Contact Us</h2>       <p>If you have any questions or need assistance, our team is here to help.</p>       <a href=\"https://www.authlete.com/contact/\"         class=\"block mt-4 text-green-500 dark:text-green-300 font-bold hover:underline\">Contact Page</a>     </section>   </main> </div>

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type DynamicClientRegistrationAPI interface {

	/*
			ClientRegistrationApi Register Client

			Register a client. This API is supposed to be used to implement a client registration endpoint that
		complies with [RFC 7591](https://datatracker.ietf.org/doc/html/rfc7591) (OAuth 2.0 Dynamic Client
		Registration Protocol).

		<br>
		<details>
		<summary>Description</summary>

		This API is supposed to be called from the within the implementation of the client registration
		endpoint of the authorization server. The authorization server implementation should retrieve
		the value of `action` from the response and take the following steps according to the value.

		**INTERNAL_SERVER_ERROR**

		When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the API call from the authorization
		server implementation was wrong or that an error occurred in Authlete.

		In either case, from a viewpoint of the client or developer, it is an error on the server side.
		Therefore, the authorization server implementation should generate a response with "500 Internal
		Server Error"s and `application/json`.

		The value of `responseContent` is a JSON string which describes the error, so it can be used as
		the entity body of the response.

		The following illustrates the response which the authorization server implementation should generate
		and return to the client or developer.

		```
		HTTP/1.1 500 Internal Server Error
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		The endpoint implementation may return another different response to the client or developer since
		"500 Internal Server Error" is not required by the specification.

		**BAD_REQUEST**

		When the value of `action` is `BAD_REQUEST`, it means that the request from the client or developer
		was wrong.

		The authorization server implementation should generate a response with "400 Bad Request" and `application/json`.

		The value of `responseContent` is a JSON string which describes the error, so it can be used
		as the entity body of the response.

		The following illustrates the response which the authorization server implementation should generate
		and return to the client or developer.

		```
		HTTP/1.1 400 Bad Request
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		**CREATED**

		When the value of `action` is `CREATED`, it means that the request from the client or developer is
		valid.

		The authorization server implementation should generate a response to the client or developer with
		"201 CREATED" and `application/json`.

		The `responseContent` a JSON string which can be used as the entity body of the response.

		The following illustrates the response which the authorization server implementation should generate
		and return to the client or developer.

		```
		HTTP/1.1 201 CREATED
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```
		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@return ApiClientRegistrationApiRequest
	*/
	ClientRegistrationApi(ctx context.Context, serviceId string) ApiClientRegistrationApiRequest

	// ClientRegistrationApiExecute executes the request
	//  @return ClientRegistrationResponse
	ClientRegistrationApiExecute(r ApiClientRegistrationApiRequest) (*ClientRegistrationResponse, *http.Response, error)

	/*
			ClientRegistrationDeleteApi Delete Client

			Delete a dynamically registered client. This API is supposed to be used to implement a client
		registration management endpoint that complies with [RFC 7592](https://datatracker.ietf.org/doc/html/rfc7592)
		(OAuth 2.0 Dynamic Registration Management).

		<br>
		<details>
		<summary>Description</summary>

		This API is supposed to be called from the within the implementation of the client registration
		management endpoint of the authorization server. The authorization server implementation should
		retrieve the value of `action` from the response and take the following steps according to the value.

		**INTERNAL_SERVER_ERROR**

		When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the API call from the authorization
		server implementation was wrong or that an error occurred in Authlete.

		In either case, from a viewpoint of the client or developer, it is an error on the server side.
		Therefore, the authorization server implementation should generate a response with "500 Internal
		Server Error"s and `application/json`.

		The value of `responseContent` is a JSON string which describes the error, so it can be used as
		the entity body of the response.

		The following illustrates the response which the authorization server implementation should generate
		and return to the client or developer.

		```
		HTTP/1.1 500 Internal Server Error
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		The endpoint implementation may return another different response to the client or developer since
		"500 Internal Server Error" is not required by the specification.

		**BAD_REQUEST**

		When the value of `action` is `BAD_REQUEST`, it means that the request from the client or developer
		was wrong.

		The authorization server implementation should generate a response with "400 Bad Request" and `application/json`.

		The value of `responseContent` is a JSON string which describes the error, so it can be used as
		the entity body of the response.

		The following illustrates the response which the authorization server implementation should generate
		and return to the client or developer.

		```
		HTTP/1.1 400 Bad Request
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		**UNAUTHORIZED**

		When the value of `action` is `UNAUTHORIZED`, it means that the registration access token used by
		the client configuration request (RFC 7592) is invalid, or the client application which the token
		is tied to does not exist any longer or is invalid.

		The HTTP status of the response returned to the client application must be "401 Unauthorized" and
		the content type must be `application/json`.

		The value of `responseContent` is a JSON string which describes the error, so it can be used as
		the entity body of the response.

		The following illustrates the response which the endpoint implementation should generate and return
		to the client application.

		```
		HTTP/1.1 401 Unauthorized
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		NOTE: The `UNAUTHORIZED` value was added in October, 2021. See the description of
		`Service.unauthorizedOnClientConfigSupported` for details.

		**DELETED**

		When the value of `action` is `DELETED`, it means that the request from the client or developer is
		valid.

		The authorization server implementation should generate a response to the client or developer with
		"204 No Content".

		The following illustrates the response which the authorization server implementation should generate
		and return to the client or developer.

		```
		HTTP/1.1 204 No Content
		Cache-Control: no-store
		Pragma: no-cache
		```
		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@return ApiClientRegistrationDeleteApiRequest
	*/
	ClientRegistrationDeleteApi(ctx context.Context, serviceId string) ApiClientRegistrationDeleteApiRequest

	// ClientRegistrationDeleteApiExecute executes the request
	//  @return ClientRegistrationDeleteResponse
	ClientRegistrationDeleteApiExecute(r ApiClientRegistrationDeleteApiRequest) (*ClientRegistrationDeleteResponse, *http.Response, error)

	/*
			ClientRegistrationGetApi Get Client

			Get a dynamically registered client. This API is supposed to be used to implement a client registration
		management endpoint that complies with [RFC 7592](https://datatracker.ietf.org/doc/html/rfc7592)
		(OAuth 2.0 Dynamic Registration Management).

		<br>
		<details>
		<summary>Description</summary>

		This API is supposed to be called from the within the implementation of the client registration
		management endpoint of the authorization server. The authorization server implementation should
		retrieve the value of `action` from the response and take the following steps according to the value.

		**INTERNAL_SERVER_ERROR**

		When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the API call from the authorization
		server implementation was wrong or that an error occurred in Authlete.

		In either case, from a viewpoint of the client or developer, it is an error on the server side.
		Therefore, the authorization server implementation should generate a response to the client or developer
		with "500 Internal Server Error"s and `application/json`.

		The value of `responseContent` is a JSON string which describes the error, so it can be used as
		the entity body of the response.

		The following illustrates the response which the authorization server implementation should generate
		and return to the client or developer.

		```
		HTTP/1.1 500 Internal Server Error
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		The endpoint implementation may return another different response to the client or developer since
		"500 Internal Server Error" is not required by the specification.

		**BAD_REQUEST**

		When the value of `action` is `BAD_REQUEST`, it means that the request from the client or developer
		was wrong.

		The authorization server implementation should generate a response to the client or developer with
		"400 Bad Request" and `application/json`.

		The value of `responseContent` is a JSON string which describes the error, so it can be used as
		the entity body of the response.

		The following illustrates the response which the authorization server implementation should generate
		and return to the client or developer.

		```
		HTTP/1.1 400 Bad Request
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		**UNAUTHORIZED**

		When the value of `action` is `UNAUTHORIZED`, it means that the registration access token used by
		the client configuration request (RFC 7592) is invalid, or the client application which the token
		is tied to does not exist any longer or is invalid.

		The HTTP status of the response returned to the client application must be "401 Unauthorized" and
		the content type must be `application/json`.

		The value of `responseContent` is a JSON string which describes the error, so it can be used as
		the entity body of the response.

		The following illustrates the response which the endpoint implementation should generate and return
		to the client application.

		```
		HTTP/1.1 401 Unauthorized
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		NOTE: The `UNAUTHORIZED` value was added in October, 2021. See the description of
		`Service.unauthorizedOnClientConfigSupported` for details.

		**OK**

		When the value of `action` is `OK`, it means that the request from the client or developer is valid.

		The authorization server implementation should generate a response to the client or developer with
		"200 OK" and `application/json`.

		The `responseContent` a JSON string which can be used as the entity body of the response.

		The following illustrates the response which the authorization server implementation should generate
		and return to the client or developer.

		```
		HTTP/1.1 200 OK
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```
		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@return ApiClientRegistrationGetApiRequest
	*/
	ClientRegistrationGetApi(ctx context.Context, serviceId string) ApiClientRegistrationGetApiRequest

	// ClientRegistrationGetApiExecute executes the request
	//  @return ClientRegistrationResponse
	ClientRegistrationGetApiExecute(r ApiClientRegistrationGetApiRequest) (*ClientRegistrationResponse, *http.Response, error)

	/*
			ClientRegistrationUpdateApi Update Client

			Update a dynamically registered client. This API is supposed to be used to implement a client
		registration management endpoint that complies with [RFC 7592](https://datatracker.ietf.org/doc/html/rfc7592)
		(OAuth 2.0 Dynamic Registration Management).

		<br>
		<details>
		<summary>Description</summary>

		This API is supposed to be called from the within the implementation of the client registration
		management endpoint of the authorization server. The authorization server implementation should
		retrieve the value of `action` from the response and take the following steps according to the value.

		**INTERNAL_SERVER_ERROR**

		When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the API call from the authorization
		server implementation was wrong or that an error occurred in Authlete.

		In either case, from a viewpoint of the client or developer, it is an error on the server side.
		Therefore, the authorization server implementation should generate a response with "500 Internal
		Server Error"s and `application/json`.

		The value of `responseContent` is a JSON string which describes the error, so it can be used as
		the entity body of the response.

		The following illustrates the response which the authorization server implementation should generate
		and return to the client or developer.

		```
		HTTP/1.1 500 Internal Server Error
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		The endpoint implementation may return another different response to the client or developer since
		"500 Internal Server Error" is not required by the specification.

		**BAD_REQUEST**

		When the value of `action` is `BAD_REQUEST`, it means that the request from the client or developer
		was wrong.

		The authorization server implementation should generate a response with "400 Bad Request" and `application/json`.

		The value of `responseContent` is a JSON string which describes the error, so it can be used as
		the entity body of the response.

		The following illustrates the response which the authorization server implementation should generate
		and return to the client or developer.

		```
		HTTP/1.1 400 Bad Request
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		**UNAUTHORIZED**

		When the value of `action` is `UNAUTHORIZED`, it means that the registration access token used by
		the client configuration request (RFC 7592) is invalid, or the client application which the token
		is tied to does not exist any longer or is invalid.

		The HTTP status of the response returned to the client application must be "401 Unauthorized" and
		the content type must be `application/json`.

		The value of `responseContent` is a JSON string which describes the error, so it can be used as
		the entity body of the response.

		The following illustrates the response which the endpoint implementation should generate and return
		to the client application.

		```
		HTTP/1.1 401 Unauthorized
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		NOTE: The `UNAUTHORIZED` value was added in October, 2021. See the description of
		`Service.unauthorizedOnClientConfigSupported` for details.

		**UPDATED**

		When the value of `action` is `UPDATED`, it means that the request from the client or developer is
		valid.

		The authorization server implementation should generate a response to the client or developer with
		"200 OK" and `application/json`.

		The `responseContent` a JSON string which can be used as the entity body of the response.

		The following illustrates the response which the authorization server implementation should generate
		and return to the client or developer.

		```
		HTTP/1.1 200 OK
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```
		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@return ApiClientRegistrationUpdateApiRequest
	*/
	ClientRegistrationUpdateApi(ctx context.Context, serviceId string) ApiClientRegistrationUpdateApiRequest

	// ClientRegistrationUpdateApiExecute executes the request
	//  @return ClientRegistrationUpdateResponse
	ClientRegistrationUpdateApiExecute(r ApiClientRegistrationUpdateApiRequest) (*ClientRegistrationUpdateResponse, *http.Response, error)
}

// DynamicClientRegistrationAPIService DynamicClientRegistrationAPI service
type DynamicClientRegistrationAPIService service

type ApiClientRegistrationApiRequest struct {
	ctx                       context.Context
	ApiService                DynamicClientRegistrationAPI
	serviceId                 string
	clientRegistrationRequest *ClientRegistrationRequest
}

func (r ApiClientRegistrationApiRequest) ClientRegistrationRequest(clientRegistrationRequest ClientRegistrationRequest) ApiClientRegistrationApiRequest {
	r.clientRegistrationRequest = &clientRegistrationRequest
	return r
}

func (r ApiClientRegistrationApiRequest) Execute() (*ClientRegistrationResponse, *http.Response, error) {
	return r.ApiService.ClientRegistrationApiExecute(r)
}

/*
ClientRegistrationApi Register Client

Register a client. This API is supposed to be used to implement a client registration endpoint that
complies with [RFC 7591](https://datatracker.ietf.org/doc/html/rfc7591) (OAuth 2.0 Dynamic Client
Registration Protocol).

<br>
<details>
<summary>Description</summary>

This API is supposed to be called from the within the implementation of the client registration
endpoint of the authorization server. The authorization server implementation should retrieve
the value of `action` from the response and take the following steps according to the value.

**INTERNAL_SERVER_ERROR**

When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the API call from the authorization
server implementation was wrong or that an error occurred in Authlete.

In either case, from a viewpoint of the client or developer, it is an error on the server side.
Therefore, the authorization server implementation should generate a response with "500 Internal
Server Error"s and `application/json`.

The value of `responseContent` is a JSON string which describes the error, so it can be used as
the entity body of the response.

The following illustrates the response which the authorization server implementation should generate
and return to the client or developer.

```
HTTP/1.1 500 Internal Server Error
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

The endpoint implementation may return another different response to the client or developer since
"500 Internal Server Error" is not required by the specification.

**BAD_REQUEST**

When the value of `action` is `BAD_REQUEST`, it means that the request from the client or developer
was wrong.

The authorization server implementation should generate a response with "400 Bad Request" and `application/json`.

The value of `responseContent` is a JSON string which describes the error, so it can be used
as the entity body of the response.

The following illustrates the response which the authorization server implementation should generate
and return to the client or developer.

```
HTTP/1.1 400 Bad Request
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

**CREATED**

When the value of `action` is `CREATED`, it means that the request from the client or developer is
valid.

The authorization server implementation should generate a response to the client or developer with
"201 CREATED" and `application/json`.

The `responseContent` a JSON string which can be used as the entity body of the response.

The following illustrates the response which the authorization server implementation should generate
and return to the client or developer.

```
HTTP/1.1 201 CREATED
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```
</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@return ApiClientRegistrationApiRequest
*/
func (a *DynamicClientRegistrationAPIService) ClientRegistrationApi(ctx context.Context, serviceId string) ApiClientRegistrationApiRequest {
	return ApiClientRegistrationApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return ClientRegistrationResponse
func (a *DynamicClientRegistrationAPIService) ClientRegistrationApiExecute(r ApiClientRegistrationApiRequest) (*ClientRegistrationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientRegistrationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DynamicClientRegistrationAPIService.ClientRegistrationApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/registration"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientRegistrationRequest == nil {
		return localVarReturnValue, nil, reportError("clientRegistrationRequest is required and must be specified")
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
	localVarPostBody = r.clientRegistrationRequest
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

type ApiClientRegistrationDeleteApiRequest struct {
	ctx                             context.Context
	ApiService                      DynamicClientRegistrationAPI
	serviceId                       string
	clientRegistrationDeleteRequest *ClientRegistrationDeleteRequest
}

func (r ApiClientRegistrationDeleteApiRequest) ClientRegistrationDeleteRequest(clientRegistrationDeleteRequest ClientRegistrationDeleteRequest) ApiClientRegistrationDeleteApiRequest {
	r.clientRegistrationDeleteRequest = &clientRegistrationDeleteRequest
	return r
}

func (r ApiClientRegistrationDeleteApiRequest) Execute() (*ClientRegistrationDeleteResponse, *http.Response, error) {
	return r.ApiService.ClientRegistrationDeleteApiExecute(r)
}

/*
ClientRegistrationDeleteApi Delete Client

Delete a dynamically registered client. This API is supposed to be used to implement a client
registration management endpoint that complies with [RFC 7592](https://datatracker.ietf.org/doc/html/rfc7592)
(OAuth 2.0 Dynamic Registration Management).

<br>
<details>
<summary>Description</summary>

This API is supposed to be called from the within the implementation of the client registration
management endpoint of the authorization server. The authorization server implementation should
retrieve the value of `action` from the response and take the following steps according to the value.

**INTERNAL_SERVER_ERROR**

When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the API call from the authorization
server implementation was wrong or that an error occurred in Authlete.

In either case, from a viewpoint of the client or developer, it is an error on the server side.
Therefore, the authorization server implementation should generate a response with "500 Internal
Server Error"s and `application/json`.

The value of `responseContent` is a JSON string which describes the error, so it can be used as
the entity body of the response.

The following illustrates the response which the authorization server implementation should generate
and return to the client or developer.

```
HTTP/1.1 500 Internal Server Error
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

The endpoint implementation may return another different response to the client or developer since
"500 Internal Server Error" is not required by the specification.

**BAD_REQUEST**

When the value of `action` is `BAD_REQUEST`, it means that the request from the client or developer
was wrong.

The authorization server implementation should generate a response with "400 Bad Request" and `application/json`.

The value of `responseContent` is a JSON string which describes the error, so it can be used as
the entity body of the response.

The following illustrates the response which the authorization server implementation should generate
and return to the client or developer.

```
HTTP/1.1 400 Bad Request
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

**UNAUTHORIZED**

When the value of `action` is `UNAUTHORIZED`, it means that the registration access token used by
the client configuration request (RFC 7592) is invalid, or the client application which the token
is tied to does not exist any longer or is invalid.

The HTTP status of the response returned to the client application must be "401 Unauthorized" and
the content type must be `application/json`.

The value of `responseContent` is a JSON string which describes the error, so it can be used as
the entity body of the response.

The following illustrates the response which the endpoint implementation should generate and return
to the client application.

```
HTTP/1.1 401 Unauthorized
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

NOTE: The `UNAUTHORIZED` value was added in October, 2021. See the description of
`Service.unauthorizedOnClientConfigSupported` for details.

**DELETED**

When the value of `action` is `DELETED`, it means that the request from the client or developer is
valid.

The authorization server implementation should generate a response to the client or developer with
"204 No Content".

The following illustrates the response which the authorization server implementation should generate
and return to the client or developer.

```
HTTP/1.1 204 No Content
Cache-Control: no-store
Pragma: no-cache
```
</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@return ApiClientRegistrationDeleteApiRequest
*/
func (a *DynamicClientRegistrationAPIService) ClientRegistrationDeleteApi(ctx context.Context, serviceId string) ApiClientRegistrationDeleteApiRequest {
	return ApiClientRegistrationDeleteApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return ClientRegistrationDeleteResponse
func (a *DynamicClientRegistrationAPIService) ClientRegistrationDeleteApiExecute(r ApiClientRegistrationDeleteApiRequest) (*ClientRegistrationDeleteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientRegistrationDeleteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DynamicClientRegistrationAPIService.ClientRegistrationDeleteApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/registration/delete"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientRegistrationDeleteRequest == nil {
		return localVarReturnValue, nil, reportError("clientRegistrationDeleteRequest is required and must be specified")
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
	localVarPostBody = r.clientRegistrationDeleteRequest
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

type ApiClientRegistrationGetApiRequest struct {
	ctx                       context.Context
	ApiService                DynamicClientRegistrationAPI
	serviceId                 string
	clientRegistrationRequest *ClientRegistrationRequest
}

func (r ApiClientRegistrationGetApiRequest) ClientRegistrationRequest(clientRegistrationRequest ClientRegistrationRequest) ApiClientRegistrationGetApiRequest {
	r.clientRegistrationRequest = &clientRegistrationRequest
	return r
}

func (r ApiClientRegistrationGetApiRequest) Execute() (*ClientRegistrationResponse, *http.Response, error) {
	return r.ApiService.ClientRegistrationGetApiExecute(r)
}

/*
ClientRegistrationGetApi Get Client

Get a dynamically registered client. This API is supposed to be used to implement a client registration
management endpoint that complies with [RFC 7592](https://datatracker.ietf.org/doc/html/rfc7592)
(OAuth 2.0 Dynamic Registration Management).

<br>
<details>
<summary>Description</summary>

This API is supposed to be called from the within the implementation of the client registration
management endpoint of the authorization server. The authorization server implementation should
retrieve the value of `action` from the response and take the following steps according to the value.

**INTERNAL_SERVER_ERROR**

When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the API call from the authorization
server implementation was wrong or that an error occurred in Authlete.

In either case, from a viewpoint of the client or developer, it is an error on the server side.
Therefore, the authorization server implementation should generate a response to the client or developer
with "500 Internal Server Error"s and `application/json`.

The value of `responseContent` is a JSON string which describes the error, so it can be used as
the entity body of the response.

The following illustrates the response which the authorization server implementation should generate
and return to the client or developer.

```
HTTP/1.1 500 Internal Server Error
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

The endpoint implementation may return another different response to the client or developer since
"500 Internal Server Error" is not required by the specification.

**BAD_REQUEST**

When the value of `action` is `BAD_REQUEST`, it means that the request from the client or developer
was wrong.

The authorization server implementation should generate a response to the client or developer with
"400 Bad Request" and `application/json`.

The value of `responseContent` is a JSON string which describes the error, so it can be used as
the entity body of the response.

The following illustrates the response which the authorization server implementation should generate
and return to the client or developer.

```
HTTP/1.1 400 Bad Request
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

**UNAUTHORIZED**

When the value of `action` is `UNAUTHORIZED`, it means that the registration access token used by
the client configuration request (RFC 7592) is invalid, or the client application which the token
is tied to does not exist any longer or is invalid.

The HTTP status of the response returned to the client application must be "401 Unauthorized" and
the content type must be `application/json`.

The value of `responseContent` is a JSON string which describes the error, so it can be used as
the entity body of the response.

The following illustrates the response which the endpoint implementation should generate and return
to the client application.

```
HTTP/1.1 401 Unauthorized
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

NOTE: The `UNAUTHORIZED` value was added in October, 2021. See the description of
`Service.unauthorizedOnClientConfigSupported` for details.

**OK**

When the value of `action` is `OK`, it means that the request from the client or developer is valid.

The authorization server implementation should generate a response to the client or developer with
"200 OK" and `application/json`.

The `responseContent` a JSON string which can be used as the entity body of the response.

The following illustrates the response which the authorization server implementation should generate
and return to the client or developer.

```
HTTP/1.1 200 OK
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```
</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@return ApiClientRegistrationGetApiRequest
*/
func (a *DynamicClientRegistrationAPIService) ClientRegistrationGetApi(ctx context.Context, serviceId string) ApiClientRegistrationGetApiRequest {
	return ApiClientRegistrationGetApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return ClientRegistrationResponse
func (a *DynamicClientRegistrationAPIService) ClientRegistrationGetApiExecute(r ApiClientRegistrationGetApiRequest) (*ClientRegistrationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientRegistrationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DynamicClientRegistrationAPIService.ClientRegistrationGetApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/registration/get"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientRegistrationRequest == nil {
		return localVarReturnValue, nil, reportError("clientRegistrationRequest is required and must be specified")
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
	localVarPostBody = r.clientRegistrationRequest
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

type ApiClientRegistrationUpdateApiRequest struct {
	ctx                             context.Context
	ApiService                      DynamicClientRegistrationAPI
	serviceId                       string
	clientRegistrationUpdateRequest *ClientRegistrationUpdateRequest
}

func (r ApiClientRegistrationUpdateApiRequest) ClientRegistrationUpdateRequest(clientRegistrationUpdateRequest ClientRegistrationUpdateRequest) ApiClientRegistrationUpdateApiRequest {
	r.clientRegistrationUpdateRequest = &clientRegistrationUpdateRequest
	return r
}

func (r ApiClientRegistrationUpdateApiRequest) Execute() (*ClientRegistrationUpdateResponse, *http.Response, error) {
	return r.ApiService.ClientRegistrationUpdateApiExecute(r)
}

/*
ClientRegistrationUpdateApi Update Client

Update a dynamically registered client. This API is supposed to be used to implement a client
registration management endpoint that complies with [RFC 7592](https://datatracker.ietf.org/doc/html/rfc7592)
(OAuth 2.0 Dynamic Registration Management).

<br>
<details>
<summary>Description</summary>

This API is supposed to be called from the within the implementation of the client registration
management endpoint of the authorization server. The authorization server implementation should
retrieve the value of `action` from the response and take the following steps according to the value.

**INTERNAL_SERVER_ERROR**

When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the API call from the authorization
server implementation was wrong or that an error occurred in Authlete.

In either case, from a viewpoint of the client or developer, it is an error on the server side.
Therefore, the authorization server implementation should generate a response with "500 Internal
Server Error"s and `application/json`.

The value of `responseContent` is a JSON string which describes the error, so it can be used as
the entity body of the response.

The following illustrates the response which the authorization server implementation should generate
and return to the client or developer.

```
HTTP/1.1 500 Internal Server Error
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

The endpoint implementation may return another different response to the client or developer since
"500 Internal Server Error" is not required by the specification.

**BAD_REQUEST**

When the value of `action` is `BAD_REQUEST`, it means that the request from the client or developer
was wrong.

The authorization server implementation should generate a response with "400 Bad Request" and `application/json`.

The value of `responseContent` is a JSON string which describes the error, so it can be used as
the entity body of the response.

The following illustrates the response which the authorization server implementation should generate
and return to the client or developer.

```
HTTP/1.1 400 Bad Request
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

**UNAUTHORIZED**

When the value of `action` is `UNAUTHORIZED`, it means that the registration access token used by
the client configuration request (RFC 7592) is invalid, or the client application which the token
is tied to does not exist any longer or is invalid.

The HTTP status of the response returned to the client application must be "401 Unauthorized" and
the content type must be `application/json`.

The value of `responseContent` is a JSON string which describes the error, so it can be used as
the entity body of the response.

The following illustrates the response which the endpoint implementation should generate and return
to the client application.

```
HTTP/1.1 401 Unauthorized
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

NOTE: The `UNAUTHORIZED` value was added in October, 2021. See the description of
`Service.unauthorizedOnClientConfigSupported` for details.

**UPDATED**

When the value of `action` is `UPDATED`, it means that the request from the client or developer is
valid.

The authorization server implementation should generate a response to the client or developer with
"200 OK" and `application/json`.

The `responseContent` a JSON string which can be used as the entity body of the response.

The following illustrates the response which the authorization server implementation should generate
and return to the client or developer.

```
HTTP/1.1 200 OK
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```
</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@return ApiClientRegistrationUpdateApiRequest
*/
func (a *DynamicClientRegistrationAPIService) ClientRegistrationUpdateApi(ctx context.Context, serviceId string) ApiClientRegistrationUpdateApiRequest {
	return ApiClientRegistrationUpdateApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return ClientRegistrationUpdateResponse
func (a *DynamicClientRegistrationAPIService) ClientRegistrationUpdateApiExecute(r ApiClientRegistrationUpdateApiRequest) (*ClientRegistrationUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientRegistrationUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DynamicClientRegistrationAPIService.ClientRegistrationUpdateApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/registration/update"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientRegistrationUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("clientRegistrationUpdateRequest is required and must be specified")
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
	localVarPostBody = r.clientRegistrationUpdateRequest
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
