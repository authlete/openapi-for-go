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

type UserInfoEndpointAPI interface {

	/*
			AuthUserinfoApi Process UserInfo Request

			This API gathers information about a user.

		<br>
		<details>
		<summary>Description</summary>

		This API is supposed to be called from within the implementation of the [userinfo endpoint](https://openid.net/specs/openid-connect-core-1_0.html#UserInfo)
		of the authorization server in order to get information about the user that is associated with
		an access token.

		The response from `/auth/userinfo` API has various parameters. Among them, it is `action` parameter
		that the authorization server implementation should check first because it denotes the next action
		that the authorization server implementation should take. According to the value of `action`, the
		service implementation must take the steps described below.

		**INTERNAL_SERVER_ERROR**

		When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the request from the authorization
		server implementation was wrong or that an error occurred in Authlete. In either case, from the
		viewpoint of the client application, it is an error on the server side. Therefore, the service
		implementation should generate a response to the client application with HTTP status of "500 Internal
		Server Error".

		The value of `responseContent` is a string which describes the error in the format of [RFC 6750](https://datatracker.ietf.org/doc/html/rfc6750)
		(OAuth 2.0 Bearer Token Usage) so the userinfo endpoint implementation can use the value of `responseContent`
		as the value of`WWW-Authenticate` header.

		The following is an example response which complies with RFC 6750. Note that OpenID Connect Core
		1.0 requires that an error response from userinfo endpoint comply with RFC 6750. See [5.3.3. UserInfo
		Response](https://openid.net/specs/openid-connect-core-1_0.html#UserInfoError) for details.

		```
		HTTP/1.1 500 Internal Server Error
		WWW-Authenticate: {responseContent}
		Cache-Control: no-store
		Pragma: no-cache
		```

		**BAD_REQUEST**

		When the value of `action` is `BAD_REQUEST`, it means that the request from the client application
		does not contain an access token (= the request from the authorization server implementation to
		Authlete does not contain `token` parameter).

		The value of `responseContent` is a string which describes the error in the format
		of [RFC 6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage) so the
		userinfo endpoint implementation can use the value of `responseContent` as the value of`WWW-Authenticate`
		header.

		The following is an example response which complies with RFC 6750. Note that OpenID Connect Core
		1.0 requires that an error response from userinfo endpoint comply with RFC 6750. See [5.3.3. UserInfo
		Response](https://openid.net/specs/openid-connect-core-1_0.html#UserInfoError) for details.

		```
		HTTP/1.1 400 Bad Request
		WWW-Authenticate: {responseContent}
		Cache-Control: no-store
		Pragma: no-cache
		```

		**UNAUTHORIZED**

		When the value of `action` is `UNAUTHORIZED`, it means that the access token does not exist, has
		expired, or is not associated with any subject (= any user account).

		The value of `responseContent` is a string which describes the error in the format of [RFC
		6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage) so the userinfo
		endpoint implementation can use the value of `responseContent` as the value of`WWW-Authenticate`
		header.

		The following is an example response which complies with RFC 6750. Note that OpenID Connect Core
		1.0 requires that an error response from userinfo endpoint comply with RFC 6750. See [5.3.3. UserInfo
		Response](https://openid.net/specs/openid-connect-core-1_0.html#UserInfoError) for details.

		```
		HTTP/1.1 401 Unauthorized
		WWW-Authenticate: {responseContent}
		Cache-Control: no-store
		Pragma: no-cache
		```

		**FORBIDDEN**

		When the value of `action` is `FORBIDDEN`, it means that the access token does not include the
		`openid` scope.

		The value of `responseContent` is a string which describes the error in the format of [RFC 6750](https://datatracker.ietf.org/doc/html/rfc6750)
		(OAuth 2.0 Bearer Token Usage) so the userinfo endpoint implementation can use the value of `responseContent`
		as the value of`WWW-Authenticate` header.

		The following is an example response which complies with RFC 6750. Note that OpenID Connect Core
		1.0 requires that an error response from userinfo endpoint comply with RFC 6750. See [5.3.3. UserInfo
		Response](https://openid.net/specs/openid-connect-core-1_0.html#UserInfoError) for details.

		```
		HTTP/1.1 403 Forbidden
		WWW-Authenticate: {responseContent}
		Cache-Control: no-store
		Pragma: no-cache
		```

		**OK**

		When the value of `action` is `OK`, it means that the access token which the client application
		presented is valid. To be concrete, it means that the access token exists, has not expired, includes
		the openid scope, and is associated with a subject (= a user account).

		What the userinfo endpoint implementation should do next is to collect information about the subject
		(user) from your database. The value of the `subject` is contained in the subject parameter in the
		response from this API and the names of data, i.e., the claims names are contained in the claims
		parameter in the response. For example, if the `subject` parameter is `joe123` and the claims
		parameter is `[ "given_name", "email" ]`, you need to extract information about joe123's given name
		and email from your database.

		Then, call Authlete's `/auth/userinfo/issue` API with the collected information and the access token
		in order to make Authlete generate an ID token.

		If an error occurred during the above steps, generate an error response to the client. The response
		should comply with [RFC 6750](https://datatracker.ietf.org/doc/html/rfc6750). For example, if the
		subject associated with the access token does not exist in your database any longer, you may feel
		like generating a response like below.

		```
		HTTP/1.1 400 Bad Request
		WWW-Authenticate: Bearer error="invalid_token",
		 error_description="The subject associated with the access token does not exist."
		Cache-Control: no-store
		Pragma: no-cache
		```

		Also, an error might occur on database access. If you treat the error as an internal server error,
		then the response would be like the following.

		```
		HTTP/1.1 500 Internal Server Error
		WWW-Authenticate: Bearer error="server_error",
		 error_description="Failed to extract information about the subject from the database."
		Cache-Control: no-store
		Pragma: no-cache
		```
		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@return ApiAuthUserinfoApiRequest
	*/
	AuthUserinfoApi(ctx context.Context, serviceId string) ApiAuthUserinfoApiRequest

	// AuthUserinfoApiExecute executes the request
	//  @return UserinfoResponse
	AuthUserinfoApiExecute(r ApiAuthUserinfoApiRequest) (*UserinfoResponse, *http.Response, error)

	/*
			AuthUserinfoIssueApi Issue UserInfo Response

			This API generates an ID token.

		<br>
		<details>
		<summary>Description</summary>

		This API is supposed to be called from within the implementation of the [userinfo endpoint](https://openid.net/specs/openid-connect-core-1_0.html#UserInfo)
		of the authorization server in order to generate an ID token. Before calling this API, a valid
		response from `/auth/userinfo` API must be obtained. Then, call this API with the access token
		contained in the response and the claims values of the user (subject) associated with the access
		token. See **OK** written in the description of `/auth/userinfo` API for details.

		The response from `/auth/userinfo/issue` API has various parameters. Among them, it is `action`
		parameter that the authorization server implementation should check first because it denotes the
		next action that the authorization server implementation should take. According to the value of
		`action`, the service implementation must take the steps described below.

		**INTERNAL_SERVER_ERROR**

		When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the request from the authorization
		server implementation was wrong or that an error occurred in Authlete. In either case, from the
		viewpoint of the client application, it is an error on the server side. Therefore, the service
		implementation should generate a response to the client application with HTTP status of "500 Internal
		Server Error".

		The parameter `responseContent` returns a string which describes the error in the format of [RFC
		6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage) so the userinfo
		endpoint implementation can use the value of `responseContent` as the value of`WWW-Authenticate`
		header.

		The following is an example response which complies with RFC 6750. Note that OpenID Connect Core
		1.0 requires that an error response from userinfo endpoint comply with RFC 6750. See [5.3.3. UserInfo
		Response](https://openid.net/specs/openid-connect-core-1_0.html#UserInfoError) for details.

		```
		HTTP/1.1 500 Internal Server Error
		WWW-Authenticate: {responseContent}
		Cache-Control: no-store
		Pragma: no-cache
		```

		**BAD_REQUEST**

		When the value of `action` is `BAD_REQUEST`, it means that the request from the client application
		does not contain an access token (= the request from the authorization server implementation to
		Authlete does not contain `token` parameter).

		The parameter `responseContent` returns a string which describes the error in the format of [RFC
		6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage) so the userinfo
		endpoint implementation can use the value of `responseContent` as the value of`WWW-Authenticate`
		header.

		The following is an example response which complies with RFC 6750. Note that OpenID Connect Core
		1.0 requires that an error response from userinfo endpoint comply with RFC 6750. See [5.3.3. UserInfo
		Response](https://openid.net/specs/openid-connect-core-1_0.html#UserInfoError) for details.

		```
		HTTP/1.1 400 Bad Request
		WWW-Authenticate: {responseContent}
		Cache-Control: no-store
		Pragma: no-cache
		```

		**UNAUTHORIZED**

		When the value of `action` is `UNAUTHORIZED`, it means that the access token does not exist, has
		expired, or is not associated with any subject (= any user account).

		The parameter `responseContent` returns a string which describes the error in the format of [RFC
		6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage) so the userinfo
		endpoint implementation can use the value of `responseContent` as the value of`WWW-Authenticate`
		header.

		The following is an example response which complies with RFC 6750. Note that OpenID Connect Core
		1.0 requires that an error response from userinfo endpoint comply with RFC 6750. See [5.3.3. UserInfo
		Response](https://openid.net/specs/openid-connect-core-1_0.html#UserInfoError) for details.

		```
		HTTP/1.1 401 Unauthorized
		WWW-Authenticate: {responseContent}
		Cache-Control: no-store
		Pragma: no-cache
		```

		**FORBIDDEN**

		When the value of `action` is `FORBIDDEN`, it means that the access token does not include the
		`openid` scope.

		The parameter `responseContent` returns a string which describes the error in the format of [RFC
		6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage) so the userinfo
		endpoint implementation can use the value of `responseContent` as the value of`WWW-Authenticate`
		header.

		The following is an example response which complies with RFC 6750. Note that OpenID Connect Core
		1.0 requires that an error response from userinfo endpoint comply with RFC 6750. See [5.3.3. UserInfo
		Response](https://openid.net/specs/openid-connect-core-1_0.html#UserInfoError) for details.

		```
		HTTP/1.1 403 Forbidden
		WWW-Authenticate: {responseContent}
		Cache-Control: no-store
		Pragma: no-cache
		```

		**JSON**

		When the value of `action` is `JSON`, it means that the access token which the client application
		presented is valid and an ID token was successfully generated in the format of JSON.

		The userinfo endpoint implementation is expected to generate a response to the client application.
		The content type of the response must be `application/json` and the response body must be an ID
		token in JSON format.

		The value of `responseContent` is the ID token in JSON format when `action` is `JSON`, so
		a response to the client can be built like below.

		```
		HTTP/1.1 200 OK
		Cache-Control: no-store
		Pragma: no-cache
		Content-Type: application/json;charset=UTF-8

		{responseContent}
		```

		**JWT**

		When the value of `action` is `JWT`, it means that the access token which the client application
		presented is valid and an ID token was successfully generated in the format of JWT (JSON Web Token)
		([RFC 7519](https://datatracker.ietf.org/doc/html/rfc7519)).

		The userinfo endpoint implementation is expected to generate a response to the client application.
		The content type of the response must be `application/jwt` and the response body must be an ID
		token in JWT format.

		The value of `responseContent` is the ID token in JSON format when `action` is `JWT`, so a response
		to the client can be built like below.

		```
		HTTP/1.1 200 OK
		Cache-Control: no-store
		Pragma: no-cache
		Content-Type: application/jwt

		{responseContent}
		```

		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@return ApiAuthUserinfoIssueApiRequest
	*/
	AuthUserinfoIssueApi(ctx context.Context, serviceId string) ApiAuthUserinfoIssueApiRequest

	// AuthUserinfoIssueApiExecute executes the request
	//  @return UserinfoIssueResponse
	AuthUserinfoIssueApiExecute(r ApiAuthUserinfoIssueApiRequest) (*UserinfoIssueResponse, *http.Response, error)
}

// UserInfoEndpointAPIService UserInfoEndpointAPI service
type UserInfoEndpointAPIService service

type ApiAuthUserinfoApiRequest struct {
	ctx             context.Context
	ApiService      UserInfoEndpointAPI
	serviceId       string
	userinfoRequest *UserinfoRequest
}

func (r ApiAuthUserinfoApiRequest) UserinfoRequest(userinfoRequest UserinfoRequest) ApiAuthUserinfoApiRequest {
	r.userinfoRequest = &userinfoRequest
	return r
}

func (r ApiAuthUserinfoApiRequest) Execute() (*UserinfoResponse, *http.Response, error) {
	return r.ApiService.AuthUserinfoApiExecute(r)
}

/*
AuthUserinfoApi Process UserInfo Request

This API gathers information about a user.

<br>
<details>
<summary>Description</summary>

This API is supposed to be called from within the implementation of the [userinfo endpoint](https://openid.net/specs/openid-connect-core-1_0.html#UserInfo)
of the authorization server in order to get information about the user that is associated with
an access token.

The response from `/auth/userinfo` API has various parameters. Among them, it is `action` parameter
that the authorization server implementation should check first because it denotes the next action
that the authorization server implementation should take. According to the value of `action`, the
service implementation must take the steps described below.

**INTERNAL_SERVER_ERROR**

When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the request from the authorization
server implementation was wrong or that an error occurred in Authlete. In either case, from the
viewpoint of the client application, it is an error on the server side. Therefore, the service
implementation should generate a response to the client application with HTTP status of "500 Internal
Server Error".

The value of `responseContent` is a string which describes the error in the format of [RFC 6750](https://datatracker.ietf.org/doc/html/rfc6750)
(OAuth 2.0 Bearer Token Usage) so the userinfo endpoint implementation can use the value of `responseContent`
as the value of`WWW-Authenticate` header.

The following is an example response which complies with RFC 6750. Note that OpenID Connect Core
1.0 requires that an error response from userinfo endpoint comply with RFC 6750. See [5.3.3. UserInfo
Response](https://openid.net/specs/openid-connect-core-1_0.html#UserInfoError) for details.

```
HTTP/1.1 500 Internal Server Error
WWW-Authenticate: {responseContent}
Cache-Control: no-store
Pragma: no-cache
```

**BAD_REQUEST**

When the value of `action` is `BAD_REQUEST`, it means that the request from the client application
does not contain an access token (= the request from the authorization server implementation to
Authlete does not contain `token` parameter).

The value of `responseContent` is a string which describes the error in the format
of [RFC 6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage) so the
userinfo endpoint implementation can use the value of `responseContent` as the value of`WWW-Authenticate`
header.

The following is an example response which complies with RFC 6750. Note that OpenID Connect Core
1.0 requires that an error response from userinfo endpoint comply with RFC 6750. See [5.3.3. UserInfo
Response](https://openid.net/specs/openid-connect-core-1_0.html#UserInfoError) for details.

```
HTTP/1.1 400 Bad Request
WWW-Authenticate: {responseContent}
Cache-Control: no-store
Pragma: no-cache
```

**UNAUTHORIZED**

When the value of `action` is `UNAUTHORIZED`, it means that the access token does not exist, has
expired, or is not associated with any subject (= any user account).

The value of `responseContent` is a string which describes the error in the format of [RFC
6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage) so the userinfo
endpoint implementation can use the value of `responseContent` as the value of`WWW-Authenticate`
header.

The following is an example response which complies with RFC 6750. Note that OpenID Connect Core
1.0 requires that an error response from userinfo endpoint comply with RFC 6750. See [5.3.3. UserInfo
Response](https://openid.net/specs/openid-connect-core-1_0.html#UserInfoError) for details.

```
HTTP/1.1 401 Unauthorized
WWW-Authenticate: {responseContent}
Cache-Control: no-store
Pragma: no-cache
```

**FORBIDDEN**

When the value of `action` is `FORBIDDEN`, it means that the access token does not include the
`openid` scope.

The value of `responseContent` is a string which describes the error in the format of [RFC 6750](https://datatracker.ietf.org/doc/html/rfc6750)
(OAuth 2.0 Bearer Token Usage) so the userinfo endpoint implementation can use the value of `responseContent`
as the value of`WWW-Authenticate` header.

The following is an example response which complies with RFC 6750. Note that OpenID Connect Core
1.0 requires that an error response from userinfo endpoint comply with RFC 6750. See [5.3.3. UserInfo
Response](https://openid.net/specs/openid-connect-core-1_0.html#UserInfoError) for details.

```
HTTP/1.1 403 Forbidden
WWW-Authenticate: {responseContent}
Cache-Control: no-store
Pragma: no-cache
```

**OK**

When the value of `action` is `OK`, it means that the access token which the client application
presented is valid. To be concrete, it means that the access token exists, has not expired, includes
the openid scope, and is associated with a subject (= a user account).

What the userinfo endpoint implementation should do next is to collect information about the subject
(user) from your database. The value of the `subject` is contained in the subject parameter in the
response from this API and the names of data, i.e., the claims names are contained in the claims
parameter in the response. For example, if the `subject` parameter is `joe123` and the claims
parameter is `[ "given_name", "email" ]`, you need to extract information about joe123's given name
and email from your database.

Then, call Authlete's `/auth/userinfo/issue` API with the collected information and the access token
in order to make Authlete generate an ID token.

If an error occurred during the above steps, generate an error response to the client. The response
should comply with [RFC 6750](https://datatracker.ietf.org/doc/html/rfc6750). For example, if the
subject associated with the access token does not exist in your database any longer, you may feel
like generating a response like below.

```
HTTP/1.1 400 Bad Request
WWW-Authenticate: Bearer error="invalid_token",

	error_description="The subject associated with the access token does not exist."

Cache-Control: no-store
Pragma: no-cache
```

Also, an error might occur on database access. If you treat the error as an internal server error,
then the response would be like the following.

```
HTTP/1.1 500 Internal Server Error
WWW-Authenticate: Bearer error="server_error",

	error_description="Failed to extract information about the subject from the database."

Cache-Control: no-store
Pragma: no-cache
```
</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@return ApiAuthUserinfoApiRequest
*/
func (a *UserInfoEndpointAPIService) AuthUserinfoApi(ctx context.Context, serviceId string) ApiAuthUserinfoApiRequest {
	return ApiAuthUserinfoApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return UserinfoResponse
func (a *UserInfoEndpointAPIService) AuthUserinfoApiExecute(r ApiAuthUserinfoApiRequest) (*UserinfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UserinfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserInfoEndpointAPIService.AuthUserinfoApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/auth/userinfo"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userinfoRequest == nil {
		return localVarReturnValue, nil, reportError("userinfoRequest is required and must be specified")
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
	localVarPostBody = r.userinfoRequest
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

type ApiAuthUserinfoIssueApiRequest struct {
	ctx                  context.Context
	ApiService           UserInfoEndpointAPI
	serviceId            string
	userinfoIssueRequest *UserinfoIssueRequest
}

func (r ApiAuthUserinfoIssueApiRequest) UserinfoIssueRequest(userinfoIssueRequest UserinfoIssueRequest) ApiAuthUserinfoIssueApiRequest {
	r.userinfoIssueRequest = &userinfoIssueRequest
	return r
}

func (r ApiAuthUserinfoIssueApiRequest) Execute() (*UserinfoIssueResponse, *http.Response, error) {
	return r.ApiService.AuthUserinfoIssueApiExecute(r)
}

/*
AuthUserinfoIssueApi Issue UserInfo Response

This API generates an ID token.

<br>
<details>
<summary>Description</summary>

This API is supposed to be called from within the implementation of the [userinfo endpoint](https://openid.net/specs/openid-connect-core-1_0.html#UserInfo)
of the authorization server in order to generate an ID token. Before calling this API, a valid
response from `/auth/userinfo` API must be obtained. Then, call this API with the access token
contained in the response and the claims values of the user (subject) associated with the access
token. See **OK** written in the description of `/auth/userinfo` API for details.

The response from `/auth/userinfo/issue` API has various parameters. Among them, it is `action`
parameter that the authorization server implementation should check first because it denotes the
next action that the authorization server implementation should take. According to the value of
`action`, the service implementation must take the steps described below.

**INTERNAL_SERVER_ERROR**

When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the request from the authorization
server implementation was wrong or that an error occurred in Authlete. In either case, from the
viewpoint of the client application, it is an error on the server side. Therefore, the service
implementation should generate a response to the client application with HTTP status of "500 Internal
Server Error".

The parameter `responseContent` returns a string which describes the error in the format of [RFC
6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage) so the userinfo
endpoint implementation can use the value of `responseContent` as the value of`WWW-Authenticate`
header.

The following is an example response which complies with RFC 6750. Note that OpenID Connect Core
1.0 requires that an error response from userinfo endpoint comply with RFC 6750. See [5.3.3. UserInfo
Response](https://openid.net/specs/openid-connect-core-1_0.html#UserInfoError) for details.

```
HTTP/1.1 500 Internal Server Error
WWW-Authenticate: {responseContent}
Cache-Control: no-store
Pragma: no-cache
```

**BAD_REQUEST**

When the value of `action` is `BAD_REQUEST`, it means that the request from the client application
does not contain an access token (= the request from the authorization server implementation to
Authlete does not contain `token` parameter).

The parameter `responseContent` returns a string which describes the error in the format of [RFC
6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage) so the userinfo
endpoint implementation can use the value of `responseContent` as the value of`WWW-Authenticate`
header.

The following is an example response which complies with RFC 6750. Note that OpenID Connect Core
1.0 requires that an error response from userinfo endpoint comply with RFC 6750. See [5.3.3. UserInfo
Response](https://openid.net/specs/openid-connect-core-1_0.html#UserInfoError) for details.

```
HTTP/1.1 400 Bad Request
WWW-Authenticate: {responseContent}
Cache-Control: no-store
Pragma: no-cache
```

**UNAUTHORIZED**

When the value of `action` is `UNAUTHORIZED`, it means that the access token does not exist, has
expired, or is not associated with any subject (= any user account).

The parameter `responseContent` returns a string which describes the error in the format of [RFC
6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage) so the userinfo
endpoint implementation can use the value of `responseContent` as the value of`WWW-Authenticate`
header.

The following is an example response which complies with RFC 6750. Note that OpenID Connect Core
1.0 requires that an error response from userinfo endpoint comply with RFC 6750. See [5.3.3. UserInfo
Response](https://openid.net/specs/openid-connect-core-1_0.html#UserInfoError) for details.

```
HTTP/1.1 401 Unauthorized
WWW-Authenticate: {responseContent}
Cache-Control: no-store
Pragma: no-cache
```

**FORBIDDEN**

When the value of `action` is `FORBIDDEN`, it means that the access token does not include the
`openid` scope.

The parameter `responseContent` returns a string which describes the error in the format of [RFC
6750](https://datatracker.ietf.org/doc/html/rfc6750) (OAuth 2.0 Bearer Token Usage) so the userinfo
endpoint implementation can use the value of `responseContent` as the value of`WWW-Authenticate`
header.

The following is an example response which complies with RFC 6750. Note that OpenID Connect Core
1.0 requires that an error response from userinfo endpoint comply with RFC 6750. See [5.3.3. UserInfo
Response](https://openid.net/specs/openid-connect-core-1_0.html#UserInfoError) for details.

```
HTTP/1.1 403 Forbidden
WWW-Authenticate: {responseContent}
Cache-Control: no-store
Pragma: no-cache
```

**JSON**

When the value of `action` is `JSON`, it means that the access token which the client application
presented is valid and an ID token was successfully generated in the format of JSON.

The userinfo endpoint implementation is expected to generate a response to the client application.
The content type of the response must be `application/json` and the response body must be an ID
token in JSON format.

The value of `responseContent` is the ID token in JSON format when `action` is `JSON`, so
a response to the client can be built like below.

```
HTTP/1.1 200 OK
Cache-Control: no-store
Pragma: no-cache
Content-Type: application/json;charset=UTF-8

{responseContent}
```

**JWT**

When the value of `action` is `JWT`, it means that the access token which the client application
presented is valid and an ID token was successfully generated in the format of JWT (JSON Web Token)
([RFC 7519](https://datatracker.ietf.org/doc/html/rfc7519)).

The userinfo endpoint implementation is expected to generate a response to the client application.
The content type of the response must be `application/jwt` and the response body must be an ID
token in JWT format.

The value of `responseContent` is the ID token in JSON format when `action` is `JWT`, so a response
to the client can be built like below.

```
HTTP/1.1 200 OK
Cache-Control: no-store
Pragma: no-cache
Content-Type: application/jwt

{responseContent}
```

</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@return ApiAuthUserinfoIssueApiRequest
*/
func (a *UserInfoEndpointAPIService) AuthUserinfoIssueApi(ctx context.Context, serviceId string) ApiAuthUserinfoIssueApiRequest {
	return ApiAuthUserinfoIssueApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return UserinfoIssueResponse
func (a *UserInfoEndpointAPIService) AuthUserinfoIssueApiExecute(r ApiAuthUserinfoIssueApiRequest) (*UserinfoIssueResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UserinfoIssueResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserInfoEndpointAPIService.AuthUserinfoIssueApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/auth/userinfo/issue"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userinfoIssueRequest == nil {
		return localVarReturnValue, nil, reportError("userinfoIssueRequest is required and must be specified")
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
	localVarPostBody = r.userinfoIssueRequest
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
