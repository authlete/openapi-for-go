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

type CIBAAPI interface {

	/*
			BackchannelAuthenticationApi Process Backchannel Authentication Request

			This API parses request parameters of a [backchannel authentication request](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html#auth_request)
		and returns necessary data for the authorization server implementation to process the backchannel
		authentication request further.

		<br>
		<details>
		<summary>Description</summary>

		This API is supposed to be called from within the implementation of the [backchannel authentication
		endpoint](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html#auth_backchannel_endpoint)
		of the service. The endpoint implementation must extract the request parameters from the
		backchannel authentication request from the client application and pass them as the value of parameters
		request parameter for Authlete's `/backchannel/authentication` API.

		The value of parameters is the entire entity body (which is formatted in `application/x-www-form-urlencoded`)
		of the request from the client application.

		The following code snippet is an example in JAX-RS showing how to extract request parameters from
		the backchannel authentication request.

		```java
		@POST
		@Consumes(MediaType.APPLICATION_FORM_URLENCODED)
		public Response post(String parameters)
		{
		    // 'parameters' is the entity body of the backchannel authentication request.
		    ......
		}
		```

		The endpoint implementation does not have to parse the request parameters from the client application
		because Authlete's `/backchannel/authentication` API does it.

		The response from `/backchannel/authentication` API has various parameters. Among them, it is `action`
		parameter that the authorization server implementation should check first because it denotes the
		next action that the authorization server implementation should take. According to the value of
		`action`, the service implementation must take the steps described below.

		**INTERNAL_SERVER_ERROR**

		When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the request from the authorization
		server implementation was wrong or that an error occurred in Authlete.
		In either case, from the viewpoint of the client application, it is an error on the server side.
		Therefore, the service implementation should generate a response to the client application with
		HTTP status of "500 Internal Server Error" and `application/json`.

		The value of `responseContent` is a JSON string which describes the error, so it can be used
		as the entity body of the response.

		The following illustrates the response which the service implementation should generate and return
		to the client application.

		```
		HTTP/1.1 500 Internal Server Error
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		**BAD_REQUEST**

		When the value of `action` is `BAD_REQUEST`, it means that the request from the client application
		is invalid.

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

		When the value of `action` is `UNAUTHORIZED`, it means that client authentication of the backchannel
		authentication request failed. Note that client authentication is always required at the backchannel
		authentication endpoint. This implies that public clients are not allowed to use the backchannel
		authentication endpoint.

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

		**USER_IDENTIFICATION**

		When the value of `action` is `USER_IDENTIFICATION`, it means that the backchannel authentication
		request from the client application is valid. The authorization server implementation has to follow
		the steps below.

		[1] END-USER IDENTIFICATION

		The first step is to determine the subject (= unique identifier) of the end-user from whom the
		client application wants to get authorization.

		According to the CIBA specification, a backchannel authentication request contains one (and only
		one) of the `login_hint_token`, `id_token_hint` and `login_hint` request parameters as a hint
		by which the authorization server identifies the subject of an end-user.<br>
		The authorization server implementation can know which hint is included in the backchannel authentication
		request by the `hintType` parameter. For example, when the value of the parameter `LOGIN_HINT`,
		it means that the backchannel authentication request contains the `login_hint` request parameter
		as a hint.<br>

		The value of the `hint` parameter is the value of the hint. For example, when the value of the
		`hintType` parameter is `LOGIN_HINT`, The value of the `hint` parameter is the value of the `login_hint`
		request parameter.<br>

		It is up to the authorization server implementation how to determine the subject of the end-user
		from the hint. Only when the `id_token_hint` request parameter is used, authorization server
		implementation can use the sub response parameter, which holds the value of the sub claim in the
		`id_token_hint` request parameter.

		[2] END-USER IDENTIFICATION ERROR

		There are some cases where the authorization server implementation encounters an error during
		the user identification process. In any error case, the service implementation has to return an
		HTTP response with the error response parameter to the client application. The following is an
		example of such error responses.

		```
		HTTP/1.1 400 Bad Request
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{ "error":"unknown_user_id" }
		```

		Authlete provides `/backchannel/authentication/fail` API that builds the response body (JSON)
		of an error response. However, because it is easy to build an error response manually, you may
		choose not to call the API. One good thing in using the API is that the API call can trigger
		deletion of the ticket which has been issued from Authlete's `/backchannel/authentication` API.
		If you don't call `/backchannel/authentication/fail` API, the ticket will continue to exist in
		the database until it is cleaned up by the batch program after the ticket expires.<br>

		Possible error cases that the authorization server implementation itself has to handle are as
		follows. Other error cases have already been covered by `/backchannel/authentication` API.

		- <u>`expired_login_hint_token`</u><br>
		  The authorization server implementation detected that the hint presented by the `login_hint_token`
		  request parameter has expired.

		  Note that the format of `login_hint_token` is not described in the CIBA Core spec at all and
		  so there is no consensus on how to detect expiration of `login_hint_token`. Interpretation
		  of `login_hint_token` is left to each authorization server implementation.

		- <u>`unknown_user_id`</u><br>
		  The authorization server implementation could not determine the subject of the end-user by
		  the presented hint.

		- <u>`unauthorized_client`</u><br>
		  The authorization server implementation has custom rules to reject backchannel authentication
		  requests from some particular clients and found that the client which has made the backchannel
		  authentication request is one of the particular clients.

		  Note that `/backchannel/authentication` API does not return `action=USER_IDENTIFICATION` in
		  cases where the client does not exist or client authentication has failed. Therefore, the
		  authorization server implementation will never have to use the error code `unauthorized_client`
		  unless the server has intentionally implemented custom rules to reject backchannel authentication
		  requests based on clients.

		- <u>`missing_user_code`</u><br>
		  The authorization server implementation has custom rules to require that a backchannel authentication
		  request include a user code for some particular users and found that the user identified by
		  the hint is one of the particular users.

		  Note that `/backchannel/authentication` API does not return `action=USER_IDENTIFICATION` when
		  both the `backchannel_user_code_parameter_supported` metadata of the server and the
		  `backchannel_user_code_parameter` metadata of the client are true and the backchannel authentication
		  request does not include the user_code request parameter. In this case, `/backchannel/authentication`
		  API returns action=BAD_REQUEST with JSON containing `"error":"missing_user_code"`. Therefore,
		  the authorization server implementation will never have to use the error code `missing_user_code`
		  unless the server has intentionally implemented custom rules to require a user code based
		  on users even in the case where the `backchannel_user_code_parameter` metadata of the client
		  which has made the backchannel authentication request is `false`.

		- <u>`invalid_user_code`</u><br>
		  The authorization server implementation detected that the presented user code is invalid.

		  Note that the format of user_code is not described in the CIBA Core spec at all and so there
		  is no consensus on how to judge whether a user code is valid or not. It is up to each authorization
		  server implementation how to handle user codes.

		- <u>`invalid_binding_message`</u><br>
		  The authorization server implementation detected that the presented binding message is invalid.

		  Note that the format of binding_message is not described in the CIBA Core spec at all and
		  so there is no consensus on how to judge whether a binding message is valid or not. It is
		  up to each authorization server implementation how to handle binding messages.

		- <u>`invalid_target`</u><br>
		  The authorization server implementation rejects the requested target resources.

		  The error code invalid_target is from "Resource Indicators for OAuth 2.0". The specification
		  defines the resource request parameter. By using the parameter, client applications can request
		  target resources that should be bound to the access token being issued. If the authorization
		  server wants to reject the request, call `/backchannel/authentication/fail` API with `INVALID_TARGET`.

		- <u>`access_denined`</u><br>
		  The authorization server implementation has custom rules to reject backchannel authentication
		  requests without asking the end-user and respond to the client as if the end-user had rejected
		  the request in some particular cases and found that the backchannel authentication request
		  is one of the particular cases.

		  The authorization server implementation will never have to use the error code `access_denied`
		  at this timing unless the server has intentionally implemented custom rules to reject backchannel
		  authentication requests without asking the end-user and respond to the client as if the end-user
		  had rejected the request.

		[3] AUTH_REQ_ID ISSUE

		If the authorization server implementation has successfully determined the subject of the end-user,
		the next action is to return an HTTP response to the client application which contains `auth_req_id`.

		Authlete provides `/backchannel/authentication/issue` API which generates a JSON containing `auth_req_id`,
		so, your next action is (1) call the API, (2) receive the response from the API, (3) build a response
		to the client application using the content of the API response, and (4) return the response to
		the client application. See the description of `/backchannel/authentication/issue` API for details.

		[4] END-USER AUTHENTICATION AND AUTHORIZATION

		After sending a JSON containing `auth_req_id` back to the client application, the service implementation
		starts to communicate with an authentication device of the end-user. It is assumed that end-user
		authentication is performed on the authentication device and the end-user confirms the content of
		the backchannel authentication request and grants authorization to the client application if everything
		is okay. The authorization server implementation must be able to receive the result of the end-user
		authentication and authorization from the authentication device.

		How to communicate with an authentication device and achieve end-user authentication and authorization
		is up to each authorization server implementation, but the following request parameters of the backchannel
		authentication request should be taken into consideration in any implementation.

		- <u>`acr_values`</u><br>
		  A backchannel authentication request may contain an array of ACRs (Authentication Context Class
		  References) in preference order. If multiple authentication devices are registered for the end-user,
		  the authorization server implementation should take the ACRs into consideration when selecting
		  the best authentication device.

		- <u>`scope`</u><br>
		  A backchannel authentication request always contains a list of scopes. At least, `openid` is
		  included in the list (otherwise `/backchannel/authentication` API returns `action=BAD_REQUEST`).
		  It would be better to show the requested scopes to the end-user on the authentication device
		  or somewhere appropriate.

		  If the scope request parameter contains `address`, `email`, `phone` and/or `profile`, they are
		  interpreted as defined in "5.4. Requesting Claims using Scope Values of OpenID Connect Core 1.0".
		  That is, they are expanded into a list of claim names. The claimNames parameter returns the expanded
		  result.

		- <u>`binding_message`</u><br>
		  A backchannel authentication request may contain a binding message. It is a human readable identifier
		  or message intended to be displayed on both the consumption device (client application) and the
		  authentication device.

		- <u>`user_code`</u><br>
		  A backchannel authentication request may contain a user code. It is a secret code, such as password
		  or pin, known only to the end-user but verifiable by the authorization server. The user code should
		  be used to authorize sending a request to the authentication device.

		[5] END-USER AUTHENTICATION AND AUTHORIZATION COMPLETION

		After receiving the result of end-user authentication and authorization, the authorization server
		implementation must call Authlete's `/backchannel/authentication/complete` API to tell Authlete
		the result and pass necessary data so that Authlete can generate an ID token, an access token and
		optionally a refresh token. See the description of the API for details.

		[6] CLIENT NOTIFICATION

		When the backchannel token delivery mode is either `ping` or `push`, the authorization server implementation
		must send a notification to the pre-registered notification endpoint of the client after the end-user
		authentication and authorization. In this case, the `action` parameter in a response from `/backchannel/authentication/complete`
		API is `NOTIFICATION`. See the description of `/backchannel/authentication/complete` API for details.

		[7] TOKEN REQUEST

		When the backchannel token delivery mode is either `ping` or `poll`, the client application will make
		a token request to the token endpoint to get an ID token, an access token and optionally a refresh
		token.

		A token request that corresponds to a backchannel authentication request uses `urn:openid:params:grant-type:ciba`
		as the value of the `grant_type` request parameter. Authlete's `/auth/token` API recognizes the
		grant type automatically and behaves properly, so the existing token endpoint implementation does
		not have to be changed to support CIBA.
		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@return ApiBackchannelAuthenticationApiRequest
	*/
	BackchannelAuthenticationApi(ctx context.Context, serviceId string) ApiBackchannelAuthenticationApiRequest

	// BackchannelAuthenticationApiExecute executes the request
	//  @return BackchannelAuthenticationResponse
	BackchannelAuthenticationApiExecute(r ApiBackchannelAuthenticationApiRequest) (*BackchannelAuthenticationResponse, *http.Response, error)

	/*
			BackchannelAuthenticationCompleteApi Complete Backchannel Authentication

			This API returns information about what action the authorization server should take after it receives
		the result of end-user's decision about whether the end-user has approved or rejected a client application's
		request on the authentication device.

		<br>
		<details>
		<summary>Description</summary>

		After the implementation of the backchannel authentication endpoint returns JSON containing an
		`auth_req_id` to the client, the authorization server starts a background process that communicates
		with the authentication device of the end-user. On the authentication device, end-user authentication
		is performed and the end-user is asked whether they give authorization to the client or not. The
		authorization server will receive the result of end-user authentication and authorization from
		the authentication device.

		After the authorization server receives the result from the authentication device, or even in the
		case where the server gave up receiving a response from the authentication device for some reasons,
		the server should call the `/backchannel/authentication/complete` API to tell Authlete the result.

		When the end-user was authenticated and authorization was granted to the client by the end-user,
		the authorization server should call the API with `result=AUTHORIZED`. In this successful case,
		the `subject` request parameter is mandatory. If the token delivery mode is `push`, the API will generate
		an access token, an ID token and optionally a refresh token. On the other hand, if the token delivery
		mode is `poll` or `ping`, the API will just update the database record so that `/auth/token` API
		can generate tokens later.

		When the authorization server received the decision of the end-user from the authentication device
		and it indicates that the end-user has rejected to give authorization to the client, the authorization
		server should call the API with `result=ACCESS_DENIED`. In this case, if the token delivery mode
		is `push`, the API will generate an error response that contains the error response parameter and
		optionally the `error_description` and error_uri response parameters (if the `errorDescription`
		and `errorUri` request parameters have been given). On the other hand, if the token delivery mode
		is `poll` or `ping`, the API will just update the database record so that `/auth/token` API can
		generate an error response later. In any token delivery mode, the value of the error parameter will
		become `access_denied`.

		When the authorization server could not get the result of end-user authentication and authorization
		from the authentication device for some reasons, the authorization server should call the API with
		`result=TRANSACTION_FAILED`. In this error case, the API will behave in the same way as in the case
		of `ACCESS_DENIED`. The only difference is that `expired_token` is used as the value of the `error`
		parameter.

		The response from `/backchannel/authentication/complete` API has various parameters. Among them,
		it is `action` parameter that the authorization server implementation should check first because
		it denotes the next action that the authorization server implementation should take. According to
		the value of `action`, the service implementation must take the steps described below.

		**SERVER_ERROR**

		When the value of `action` is `SERVER_ERROR`, it means either (1) that the request from the authorization
		server to Authlete was wrong, or (2) that an error occurred on Authlete side.

		When the backchannel token delivery mode is `ping` or `push`, `SERVER_ERROR` is used only when
		an error is detected before the record of the ticket (which is included in the API call to `/backchannel/authentication/complete`)
		is retrieved from the database successfully. If an error is detected after the record of the ticket
		is retrieved from the database, `NOTIFICATION` is used instead of `SERVER_ERROR`.

		When the backchannel token delivery mode is `poll`, `SERVER_ERROR` is used regardless of whether
		it is before or after the record of the ticket is retrieved from the database.

		**NO_ACTION**

		When the value of `action` is `NO_ACTION`, it means that the authorization server does not have
		to take any immediate action.

		`NO_ACTION` is returned when the backchannel token delivery mode is `poll`. In this case, the client
		will receive the final result at the token endpoint.

		**NOTIFICATION**

		When the value of `action` is `NOTIFICATION`, it means that the authorization server must send a
		notification to the client notification endpoint.

		According to the CIBA Core specification, the notification is an HTTP POST request whose request
		body is JSON and whose `Authorization` header contains the client notification token, which was
		included in the backchannel authentication request as the value of the `client_notification_token`
		request parameter, as a bearer token.

		When the backchannel token delivery mode is `ping`, the request body of the notification is JSON
		which contains the `auth_req_id` property only. When the backchannel token delivery mode is `push`,
		the request body will additionally contain an access token, an ID token and other properties. Note
		that when the backchannel token delivery mode is `poll`, a notification does not have to be sent
		to the client notification endpoint.

		In error cases, in the ping mode, however, the content of a notification is not different from the
		content in successful cases. That is, the notification contains the `auth_req_id` property only.
		The client will know the error when it accesses the token endpoint. On the other hand, in the
		`push` mode, in error cases, the content of a notification will include the `error` property instead
		of an access token and an ID token. The client will know the error by detecting that error is included
		in the notification.

		In any case, the value of `responseContent` is JSON which can be used as the request body of the
		notification.

		The client notification endpoint that the notification should be sent to the value of the `clientNotificationEndpoint`
		parameter. Likewise, the client notification token that the notification should include as a bearer
		token is the `clientNotificationToken` parameter. With these methods, the notification can be built
		like the following.

		```
		POST {clientNotificationEndpoint} HTTP/1.1
		HOST: {The host of clientNotificationEndpoint}
		Authorization: Bearer {notificationToken}
		Content-Type: application/json

		{responseContent}
		```
		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@return ApiBackchannelAuthenticationCompleteApiRequest
	*/
	BackchannelAuthenticationCompleteApi(ctx context.Context, serviceId string) ApiBackchannelAuthenticationCompleteApiRequest

	// BackchannelAuthenticationCompleteApiExecute executes the request
	//  @return BackchannelAuthenticationCompleteResponse
	BackchannelAuthenticationCompleteApiExecute(r ApiBackchannelAuthenticationCompleteApiRequest) (*BackchannelAuthenticationCompleteResponse, *http.Response, error)

	/*
			BackchannelAuthenticationFailApi Fail Backchannel Authentication Request

			The API prepares JSON that contains an error. The JSON should be used as the response body of the
		response which is returned to the client from the [backchannel authentication endpoint](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html#auth_backchannel_endpoint).

		<br>
		<details>
		<summary>Description</summary>

		This API is supposed to be called from within the implementation of the [backchannel authentication
		endpoint](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html#auth_backchannel_endpoint)
		of the service in order to generate an error response to the client application.

		The response from `/backchannel/authentication/fails` API has some parameters. Among them, it is
		`action` parameter that the authorization server implementation should check first because it denotes
		the next action that the authorization server implementation should take. According to the value
		of `action`, the authorization server implementation must take the steps described below.

		**INTERNAL_SERVER_ERROR**

		When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that (1) the `reason` request parameter
		of the API call was `SERVER_ERROR`, (2) an error occurred on Authlete side, or (3) the request parameters
		of the API call were wrong. In this case, the authorization server implementation should return
		a "500 Internal Server Error" response to the client application. However, in most cases, commercial
		implementations prefer to use other HTTP status code than 5xx.

		**BAD_REQUEST**

		When the value of `action` is `BAD_REQUEST`, the authorization server implementation should return a
		"400 Bad Request" response to the client application.

		**FORBIDDEN**

		When the value of `action` is `FORBIDDEN`, it means that the `reason` request parameter of the API call
		was `ACCESS_DENIED`. In this case, the backchannel authentication endpoint of the authorization
		server implementation should return a "403 Forbidden" response to the client application.

		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@return ApiBackchannelAuthenticationFailApiRequest
	*/
	BackchannelAuthenticationFailApi(ctx context.Context, serviceId string) ApiBackchannelAuthenticationFailApiRequest

	// BackchannelAuthenticationFailApiExecute executes the request
	//  @return BackchannelAuthenticationFailResponse
	BackchannelAuthenticationFailApiExecute(r ApiBackchannelAuthenticationFailApiRequest) (*BackchannelAuthenticationFailResponse, *http.Response, error)

	/*
			BackchannelAuthenticationIssueApi Issue Backchannel Authentication Response

			This API prepares JSON that contains an `auth_req_id`. The JSON should be used as the response body
		of the response which is returned to the client from the [backchannel authentication endpoint](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html#auth_backchannel_endpoint)

		<br>
		<details>
		<summary>Description</summary>

		This API is supposed to be called from within the implementation of the backchannel authentication
		endpoint of the service in order to generate a successful response to the client application.

		The description of the `/backchannel/authentication` API describes the timing when this API should
		be called and the meaning of request parameters. See [AUTH_REQ_ID ISSUE] in `USER_IDENTIFICATION`.

		The response from `/backchannel/authentication/issue` API has some parameters. Among them, it is
		`action` parameter that the authorization server implementation should check first because it denotes
		the next `action` that the authorization server implementation should take. According to the value
		of `action`, the authorization server implementation must take the steps described below.

		```java
		@POST
		@Consumes(MediaType.APPLICATION_FORM_URLENCODED)
		public Response post(String parameters)
		{
		    // 'parameters' is the entity body of the backchannel authentication request.
		    ......
		}
		```

		The endpoint implementation does not have to parse the request parameters from the client application
		because Authlete's `/backchannel/authentication` API does it.

		The response from `/backchannel/authentication` API has various parameters. Among them, it is `action`
		parameter that the authorization server implementation should check first because it denotes the
		next action that the authorization server implementation should take. According to the value of
		`action`, the service implementation must take the steps described below.

		**INTERNAL_SERVER_ERROR**

		When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the request from the authorization
		server implementation was wrong or that an error occurred in Authlete.
		In either case, from the viewpoint of the client application, it is an error on the server side.
		Therefore, the service implementation should generate a response to the client application with
		HTTP status of "500 Internal Server Error" and `application/json`.

		The value of `responseContent` is a JSON string which describes the error, so it can be used as
		the entity body of the response.

		The following illustrates the response which the service implementation should generate and return
		to the client application.

		```
		HTTP/1.1 500 Internal Server Error
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		**INVALID_TICKET**

		When the value of `action` is `INVALID_TICKET`, it means that the ticket included in the API call
		was invalid. For example, it does not exist or has expired.

		From a viewpoint of the client application, this is an error on the server side. Therefore, the
		authorization server implementation should generate a response to the client application with
		"500 Internal Server Error" and `application/json`.

		You can build an error response in the same way as shown in the description for the case of `INTERNAL_SERVER_ERROR`.

		**OK**

		When the value of `action` is `OK`, it means that Authlete has succeeded in preparing JSON that
		contains an `auth_req_id`. The JSON should be used as the response body of the response that is
		returned to the client from the backchannel authentication endpoint. `responseContent` contains
		the JSON.

		The following illustrates the response which the authorization server implementation should generate
		and return to the client application.

		```
		HTTP/1.1 200 OK
		Content-Type: text/html;charset=UTF-8
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```
		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@return ApiBackchannelAuthenticationIssueApiRequest
	*/
	BackchannelAuthenticationIssueApi(ctx context.Context, serviceId string) ApiBackchannelAuthenticationIssueApiRequest

	// BackchannelAuthenticationIssueApiExecute executes the request
	//  @return BackchannelAuthenticationIssueResponse
	BackchannelAuthenticationIssueApiExecute(r ApiBackchannelAuthenticationIssueApiRequest) (*BackchannelAuthenticationIssueResponse, *http.Response, error)
}

// CIBAAPIService CIBAAPI service
type CIBAAPIService service

type ApiBackchannelAuthenticationApiRequest struct {
	ctx                              context.Context
	ApiService                       CIBAAPI
	serviceId                        string
	backchannelAuthenticationRequest *BackchannelAuthenticationRequest
}

func (r ApiBackchannelAuthenticationApiRequest) BackchannelAuthenticationRequest(backchannelAuthenticationRequest BackchannelAuthenticationRequest) ApiBackchannelAuthenticationApiRequest {
	r.backchannelAuthenticationRequest = &backchannelAuthenticationRequest
	return r
}

func (r ApiBackchannelAuthenticationApiRequest) Execute() (*BackchannelAuthenticationResponse, *http.Response, error) {
	return r.ApiService.BackchannelAuthenticationApiExecute(r)
}

/*
BackchannelAuthenticationApi Process Backchannel Authentication Request

This API parses request parameters of a [backchannel authentication request](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html#auth_request)
and returns necessary data for the authorization server implementation to process the backchannel
authentication request further.

<br>
<details>
<summary>Description</summary>

This API is supposed to be called from within the implementation of the [backchannel authentication
endpoint](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html#auth_backchannel_endpoint)
of the service. The endpoint implementation must extract the request parameters from the
backchannel authentication request from the client application and pass them as the value of parameters
request parameter for Authlete's `/backchannel/authentication` API.

The value of parameters is the entire entity body (which is formatted in `application/x-www-form-urlencoded`)
of the request from the client application.

The following code snippet is an example in JAX-RS showing how to extract request parameters from
the backchannel authentication request.

```java
@POST
@Consumes(MediaType.APPLICATION_FORM_URLENCODED)
public Response post(String parameters)

	{
	    // 'parameters' is the entity body of the backchannel authentication request.
	    ......
	}

```

The endpoint implementation does not have to parse the request parameters from the client application
because Authlete's `/backchannel/authentication` API does it.

The response from `/backchannel/authentication` API has various parameters. Among them, it is `action`
parameter that the authorization server implementation should check first because it denotes the
next action that the authorization server implementation should take. According to the value of
`action`, the service implementation must take the steps described below.

**INTERNAL_SERVER_ERROR**

When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the request from the authorization
server implementation was wrong or that an error occurred in Authlete.
In either case, from the viewpoint of the client application, it is an error on the server side.
Therefore, the service implementation should generate a response to the client application with
HTTP status of "500 Internal Server Error" and `application/json`.

The value of `responseContent` is a JSON string which describes the error, so it can be used
as the entity body of the response.

The following illustrates the response which the service implementation should generate and return
to the client application.

```
HTTP/1.1 500 Internal Server Error
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

**BAD_REQUEST**

When the value of `action` is `BAD_REQUEST`, it means that the request from the client application
is invalid.

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

When the value of `action` is `UNAUTHORIZED`, it means that client authentication of the backchannel
authentication request failed. Note that client authentication is always required at the backchannel
authentication endpoint. This implies that public clients are not allowed to use the backchannel
authentication endpoint.

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

**USER_IDENTIFICATION**

When the value of `action` is `USER_IDENTIFICATION`, it means that the backchannel authentication
request from the client application is valid. The authorization server implementation has to follow
the steps below.

[1] END-USER IDENTIFICATION

The first step is to determine the subject (= unique identifier) of the end-user from whom the
client application wants to get authorization.

According to the CIBA specification, a backchannel authentication request contains one (and only
one) of the `login_hint_token`, `id_token_hint` and `login_hint` request parameters as a hint
by which the authorization server identifies the subject of an end-user.<br>
The authorization server implementation can know which hint is included in the backchannel authentication
request by the `hintType` parameter. For example, when the value of the parameter `LOGIN_HINT`,
it means that the backchannel authentication request contains the `login_hint` request parameter
as a hint.<br>

The value of the `hint` parameter is the value of the hint. For example, when the value of the
`hintType` parameter is `LOGIN_HINT`, The value of the `hint` parameter is the value of the `login_hint`
request parameter.<br>

It is up to the authorization server implementation how to determine the subject of the end-user
from the hint. Only when the `id_token_hint` request parameter is used, authorization server
implementation can use the sub response parameter, which holds the value of the sub claim in the
`id_token_hint` request parameter.

[2] END-USER IDENTIFICATION ERROR

There are some cases where the authorization server implementation encounters an error during
the user identification process. In any error case, the service implementation has to return an
HTTP response with the error response parameter to the client application. The following is an
example of such error responses.

```
HTTP/1.1 400 Bad Request
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{ "error":"unknown_user_id" }
```

Authlete provides `/backchannel/authentication/fail` API that builds the response body (JSON)
of an error response. However, because it is easy to build an error response manually, you may
choose not to call the API. One good thing in using the API is that the API call can trigger
deletion of the ticket which has been issued from Authlete's `/backchannel/authentication` API.
If you don't call `/backchannel/authentication/fail` API, the ticket will continue to exist in
the database until it is cleaned up by the batch program after the ticket expires.<br>

Possible error cases that the authorization server implementation itself has to handle are as
follows. Other error cases have already been covered by `/backchannel/authentication` API.

  - <u>`expired_login_hint_token`</u><br>
    The authorization server implementation detected that the hint presented by the `login_hint_token`
    request parameter has expired.

    Note that the format of `login_hint_token` is not described in the CIBA Core spec at all and
    so there is no consensus on how to detect expiration of `login_hint_token`. Interpretation
    of `login_hint_token` is left to each authorization server implementation.

  - <u>`unknown_user_id`</u><br>
    The authorization server implementation could not determine the subject of the end-user by
    the presented hint.

  - <u>`unauthorized_client`</u><br>
    The authorization server implementation has custom rules to reject backchannel authentication
    requests from some particular clients and found that the client which has made the backchannel
    authentication request is one of the particular clients.

    Note that `/backchannel/authentication` API does not return `action=USER_IDENTIFICATION` in
    cases where the client does not exist or client authentication has failed. Therefore, the
    authorization server implementation will never have to use the error code `unauthorized_client`
    unless the server has intentionally implemented custom rules to reject backchannel authentication
    requests based on clients.

  - <u>`missing_user_code`</u><br>
    The authorization server implementation has custom rules to require that a backchannel authentication
    request include a user code for some particular users and found that the user identified by
    the hint is one of the particular users.

    Note that `/backchannel/authentication` API does not return `action=USER_IDENTIFICATION` when
    both the `backchannel_user_code_parameter_supported` metadata of the server and the
    `backchannel_user_code_parameter` metadata of the client are true and the backchannel authentication
    request does not include the user_code request parameter. In this case, `/backchannel/authentication`
    API returns action=BAD_REQUEST with JSON containing `"error":"missing_user_code"`. Therefore,
    the authorization server implementation will never have to use the error code `missing_user_code`
    unless the server has intentionally implemented custom rules to require a user code based
    on users even in the case where the `backchannel_user_code_parameter` metadata of the client
    which has made the backchannel authentication request is `false`.

  - <u>`invalid_user_code`</u><br>
    The authorization server implementation detected that the presented user code is invalid.

    Note that the format of user_code is not described in the CIBA Core spec at all and so there
    is no consensus on how to judge whether a user code is valid or not. It is up to each authorization
    server implementation how to handle user codes.

  - <u>`invalid_binding_message`</u><br>
    The authorization server implementation detected that the presented binding message is invalid.

    Note that the format of binding_message is not described in the CIBA Core spec at all and
    so there is no consensus on how to judge whether a binding message is valid or not. It is
    up to each authorization server implementation how to handle binding messages.

  - <u>`invalid_target`</u><br>
    The authorization server implementation rejects the requested target resources.

    The error code invalid_target is from "Resource Indicators for OAuth 2.0". The specification
    defines the resource request parameter. By using the parameter, client applications can request
    target resources that should be bound to the access token being issued. If the authorization
    server wants to reject the request, call `/backchannel/authentication/fail` API with `INVALID_TARGET`.

  - <u>`access_denined`</u><br>
    The authorization server implementation has custom rules to reject backchannel authentication
    requests without asking the end-user and respond to the client as if the end-user had rejected
    the request in some particular cases and found that the backchannel authentication request
    is one of the particular cases.

    The authorization server implementation will never have to use the error code `access_denied`
    at this timing unless the server has intentionally implemented custom rules to reject backchannel
    authentication requests without asking the end-user and respond to the client as if the end-user
    had rejected the request.

[3] AUTH_REQ_ID ISSUE

If the authorization server implementation has successfully determined the subject of the end-user,
the next action is to return an HTTP response to the client application which contains `auth_req_id`.

Authlete provides `/backchannel/authentication/issue` API which generates a JSON containing `auth_req_id`,
so, your next action is (1) call the API, (2) receive the response from the API, (3) build a response
to the client application using the content of the API response, and (4) return the response to
the client application. See the description of `/backchannel/authentication/issue` API for details.

[4] END-USER AUTHENTICATION AND AUTHORIZATION

After sending a JSON containing `auth_req_id` back to the client application, the service implementation
starts to communicate with an authentication device of the end-user. It is assumed that end-user
authentication is performed on the authentication device and the end-user confirms the content of
the backchannel authentication request and grants authorization to the client application if everything
is okay. The authorization server implementation must be able to receive the result of the end-user
authentication and authorization from the authentication device.

How to communicate with an authentication device and achieve end-user authentication and authorization
is up to each authorization server implementation, but the following request parameters of the backchannel
authentication request should be taken into consideration in any implementation.

  - <u>`acr_values`</u><br>
    A backchannel authentication request may contain an array of ACRs (Authentication Context Class
    References) in preference order. If multiple authentication devices are registered for the end-user,
    the authorization server implementation should take the ACRs into consideration when selecting
    the best authentication device.

  - <u>`scope`</u><br>
    A backchannel authentication request always contains a list of scopes. At least, `openid` is
    included in the list (otherwise `/backchannel/authentication` API returns `action=BAD_REQUEST`).
    It would be better to show the requested scopes to the end-user on the authentication device
    or somewhere appropriate.

    If the scope request parameter contains `address`, `email`, `phone` and/or `profile`, they are
    interpreted as defined in "5.4. Requesting Claims using Scope Values of OpenID Connect Core 1.0".
    That is, they are expanded into a list of claim names. The claimNames parameter returns the expanded
    result.

  - <u>`binding_message`</u><br>
    A backchannel authentication request may contain a binding message. It is a human readable identifier
    or message intended to be displayed on both the consumption device (client application) and the
    authentication device.

  - <u>`user_code`</u><br>
    A backchannel authentication request may contain a user code. It is a secret code, such as password
    or pin, known only to the end-user but verifiable by the authorization server. The user code should
    be used to authorize sending a request to the authentication device.

[5] END-USER AUTHENTICATION AND AUTHORIZATION COMPLETION

After receiving the result of end-user authentication and authorization, the authorization server
implementation must call Authlete's `/backchannel/authentication/complete` API to tell Authlete
the result and pass necessary data so that Authlete can generate an ID token, an access token and
optionally a refresh token. See the description of the API for details.

[6] CLIENT NOTIFICATION

When the backchannel token delivery mode is either `ping` or `push`, the authorization server implementation
must send a notification to the pre-registered notification endpoint of the client after the end-user
authentication and authorization. In this case, the `action` parameter in a response from `/backchannel/authentication/complete`
API is `NOTIFICATION`. See the description of `/backchannel/authentication/complete` API for details.

[7] TOKEN REQUEST

When the backchannel token delivery mode is either `ping` or `poll`, the client application will make
a token request to the token endpoint to get an ID token, an access token and optionally a refresh
token.

A token request that corresponds to a backchannel authentication request uses `urn:openid:params:grant-type:ciba`
as the value of the `grant_type` request parameter. Authlete's `/auth/token` API recognizes the
grant type automatically and behaves properly, so the existing token endpoint implementation does
not have to be changed to support CIBA.
</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@return ApiBackchannelAuthenticationApiRequest
*/
func (a *CIBAAPIService) BackchannelAuthenticationApi(ctx context.Context, serviceId string) ApiBackchannelAuthenticationApiRequest {
	return ApiBackchannelAuthenticationApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return BackchannelAuthenticationResponse
func (a *CIBAAPIService) BackchannelAuthenticationApiExecute(r ApiBackchannelAuthenticationApiRequest) (*BackchannelAuthenticationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BackchannelAuthenticationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CIBAAPIService.BackchannelAuthenticationApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/backchannel/authentication"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.backchannelAuthenticationRequest == nil {
		return localVarReturnValue, nil, reportError("backchannelAuthenticationRequest is required and must be specified")
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
	localVarPostBody = r.backchannelAuthenticationRequest
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

type ApiBackchannelAuthenticationCompleteApiRequest struct {
	ctx                                      context.Context
	ApiService                               CIBAAPI
	serviceId                                string
	backchannelAuthenticationCompleteRequest *BackchannelAuthenticationCompleteRequest
}

func (r ApiBackchannelAuthenticationCompleteApiRequest) BackchannelAuthenticationCompleteRequest(backchannelAuthenticationCompleteRequest BackchannelAuthenticationCompleteRequest) ApiBackchannelAuthenticationCompleteApiRequest {
	r.backchannelAuthenticationCompleteRequest = &backchannelAuthenticationCompleteRequest
	return r
}

func (r ApiBackchannelAuthenticationCompleteApiRequest) Execute() (*BackchannelAuthenticationCompleteResponse, *http.Response, error) {
	return r.ApiService.BackchannelAuthenticationCompleteApiExecute(r)
}

/*
BackchannelAuthenticationCompleteApi Complete Backchannel Authentication

This API returns information about what action the authorization server should take after it receives
the result of end-user's decision about whether the end-user has approved or rejected a client application's
request on the authentication device.

<br>
<details>
<summary>Description</summary>

After the implementation of the backchannel authentication endpoint returns JSON containing an
`auth_req_id` to the client, the authorization server starts a background process that communicates
with the authentication device of the end-user. On the authentication device, end-user authentication
is performed and the end-user is asked whether they give authorization to the client or not. The
authorization server will receive the result of end-user authentication and authorization from
the authentication device.

After the authorization server receives the result from the authentication device, or even in the
case where the server gave up receiving a response from the authentication device for some reasons,
the server should call the `/backchannel/authentication/complete` API to tell Authlete the result.

When the end-user was authenticated and authorization was granted to the client by the end-user,
the authorization server should call the API with `result=AUTHORIZED`. In this successful case,
the `subject` request parameter is mandatory. If the token delivery mode is `push`, the API will generate
an access token, an ID token and optionally a refresh token. On the other hand, if the token delivery
mode is `poll` or `ping`, the API will just update the database record so that `/auth/token` API
can generate tokens later.

When the authorization server received the decision of the end-user from the authentication device
and it indicates that the end-user has rejected to give authorization to the client, the authorization
server should call the API with `result=ACCESS_DENIED`. In this case, if the token delivery mode
is `push`, the API will generate an error response that contains the error response parameter and
optionally the `error_description` and error_uri response parameters (if the `errorDescription`
and `errorUri` request parameters have been given). On the other hand, if the token delivery mode
is `poll` or `ping`, the API will just update the database record so that `/auth/token` API can
generate an error response later. In any token delivery mode, the value of the error parameter will
become `access_denied`.

When the authorization server could not get the result of end-user authentication and authorization
from the authentication device for some reasons, the authorization server should call the API with
`result=TRANSACTION_FAILED`. In this error case, the API will behave in the same way as in the case
of `ACCESS_DENIED`. The only difference is that `expired_token` is used as the value of the `error`
parameter.

The response from `/backchannel/authentication/complete` API has various parameters. Among them,
it is `action` parameter that the authorization server implementation should check first because
it denotes the next action that the authorization server implementation should take. According to
the value of `action`, the service implementation must take the steps described below.

**SERVER_ERROR**

When the value of `action` is `SERVER_ERROR`, it means either (1) that the request from the authorization
server to Authlete was wrong, or (2) that an error occurred on Authlete side.

When the backchannel token delivery mode is `ping` or `push`, `SERVER_ERROR` is used only when
an error is detected before the record of the ticket (which is included in the API call to `/backchannel/authentication/complete`)
is retrieved from the database successfully. If an error is detected after the record of the ticket
is retrieved from the database, `NOTIFICATION` is used instead of `SERVER_ERROR`.

When the backchannel token delivery mode is `poll`, `SERVER_ERROR` is used regardless of whether
it is before or after the record of the ticket is retrieved from the database.

**NO_ACTION**

When the value of `action` is `NO_ACTION`, it means that the authorization server does not have
to take any immediate action.

`NO_ACTION` is returned when the backchannel token delivery mode is `poll`. In this case, the client
will receive the final result at the token endpoint.

**NOTIFICATION**

When the value of `action` is `NOTIFICATION`, it means that the authorization server must send a
notification to the client notification endpoint.

According to the CIBA Core specification, the notification is an HTTP POST request whose request
body is JSON and whose `Authorization` header contains the client notification token, which was
included in the backchannel authentication request as the value of the `client_notification_token`
request parameter, as a bearer token.

When the backchannel token delivery mode is `ping`, the request body of the notification is JSON
which contains the `auth_req_id` property only. When the backchannel token delivery mode is `push`,
the request body will additionally contain an access token, an ID token and other properties. Note
that when the backchannel token delivery mode is `poll`, a notification does not have to be sent
to the client notification endpoint.

In error cases, in the ping mode, however, the content of a notification is not different from the
content in successful cases. That is, the notification contains the `auth_req_id` property only.
The client will know the error when it accesses the token endpoint. On the other hand, in the
`push` mode, in error cases, the content of a notification will include the `error` property instead
of an access token and an ID token. The client will know the error by detecting that error is included
in the notification.

In any case, the value of `responseContent` is JSON which can be used as the request body of the
notification.

The client notification endpoint that the notification should be sent to the value of the `clientNotificationEndpoint`
parameter. Likewise, the client notification token that the notification should include as a bearer
token is the `clientNotificationToken` parameter. With these methods, the notification can be built
like the following.

```
POST {clientNotificationEndpoint} HTTP/1.1
HOST: {The host of clientNotificationEndpoint}
Authorization: Bearer {notificationToken}
Content-Type: application/json

{responseContent}
```
</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@return ApiBackchannelAuthenticationCompleteApiRequest
*/
func (a *CIBAAPIService) BackchannelAuthenticationCompleteApi(ctx context.Context, serviceId string) ApiBackchannelAuthenticationCompleteApiRequest {
	return ApiBackchannelAuthenticationCompleteApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return BackchannelAuthenticationCompleteResponse
func (a *CIBAAPIService) BackchannelAuthenticationCompleteApiExecute(r ApiBackchannelAuthenticationCompleteApiRequest) (*BackchannelAuthenticationCompleteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BackchannelAuthenticationCompleteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CIBAAPIService.BackchannelAuthenticationCompleteApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/backchannel/authentication/complete"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.backchannelAuthenticationCompleteRequest == nil {
		return localVarReturnValue, nil, reportError("backchannelAuthenticationCompleteRequest is required and must be specified")
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
	localVarPostBody = r.backchannelAuthenticationCompleteRequest
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

type ApiBackchannelAuthenticationFailApiRequest struct {
	ctx                                  context.Context
	ApiService                           CIBAAPI
	serviceId                            string
	backchannelAuthenticationFailRequest *BackchannelAuthenticationFailRequest
}

func (r ApiBackchannelAuthenticationFailApiRequest) BackchannelAuthenticationFailRequest(backchannelAuthenticationFailRequest BackchannelAuthenticationFailRequest) ApiBackchannelAuthenticationFailApiRequest {
	r.backchannelAuthenticationFailRequest = &backchannelAuthenticationFailRequest
	return r
}

func (r ApiBackchannelAuthenticationFailApiRequest) Execute() (*BackchannelAuthenticationFailResponse, *http.Response, error) {
	return r.ApiService.BackchannelAuthenticationFailApiExecute(r)
}

/*
BackchannelAuthenticationFailApi Fail Backchannel Authentication Request

The API prepares JSON that contains an error. The JSON should be used as the response body of the
response which is returned to the client from the [backchannel authentication endpoint](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html#auth_backchannel_endpoint).

<br>
<details>
<summary>Description</summary>

This API is supposed to be called from within the implementation of the [backchannel authentication
endpoint](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html#auth_backchannel_endpoint)
of the service in order to generate an error response to the client application.

The response from `/backchannel/authentication/fails` API has some parameters. Among them, it is
`action` parameter that the authorization server implementation should check first because it denotes
the next action that the authorization server implementation should take. According to the value
of `action`, the authorization server implementation must take the steps described below.

**INTERNAL_SERVER_ERROR**

When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that (1) the `reason` request parameter
of the API call was `SERVER_ERROR`, (2) an error occurred on Authlete side, or (3) the request parameters
of the API call were wrong. In this case, the authorization server implementation should return
a "500 Internal Server Error" response to the client application. However, in most cases, commercial
implementations prefer to use other HTTP status code than 5xx.

**BAD_REQUEST**

When the value of `action` is `BAD_REQUEST`, the authorization server implementation should return a
"400 Bad Request" response to the client application.

**FORBIDDEN**

When the value of `action` is `FORBIDDEN`, it means that the `reason` request parameter of the API call
was `ACCESS_DENIED`. In this case, the backchannel authentication endpoint of the authorization
server implementation should return a "403 Forbidden" response to the client application.

</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@return ApiBackchannelAuthenticationFailApiRequest
*/
func (a *CIBAAPIService) BackchannelAuthenticationFailApi(ctx context.Context, serviceId string) ApiBackchannelAuthenticationFailApiRequest {
	return ApiBackchannelAuthenticationFailApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return BackchannelAuthenticationFailResponse
func (a *CIBAAPIService) BackchannelAuthenticationFailApiExecute(r ApiBackchannelAuthenticationFailApiRequest) (*BackchannelAuthenticationFailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BackchannelAuthenticationFailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CIBAAPIService.BackchannelAuthenticationFailApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/backchannel/authentication/fail"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.backchannelAuthenticationFailRequest == nil {
		return localVarReturnValue, nil, reportError("backchannelAuthenticationFailRequest is required and must be specified")
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
	localVarPostBody = r.backchannelAuthenticationFailRequest
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

type ApiBackchannelAuthenticationIssueApiRequest struct {
	ctx                                   context.Context
	ApiService                            CIBAAPI
	serviceId                             string
	backchannelAuthenticationIssueRequest *BackchannelAuthenticationIssueRequest
}

func (r ApiBackchannelAuthenticationIssueApiRequest) BackchannelAuthenticationIssueRequest(backchannelAuthenticationIssueRequest BackchannelAuthenticationIssueRequest) ApiBackchannelAuthenticationIssueApiRequest {
	r.backchannelAuthenticationIssueRequest = &backchannelAuthenticationIssueRequest
	return r
}

func (r ApiBackchannelAuthenticationIssueApiRequest) Execute() (*BackchannelAuthenticationIssueResponse, *http.Response, error) {
	return r.ApiService.BackchannelAuthenticationIssueApiExecute(r)
}

/*
BackchannelAuthenticationIssueApi Issue Backchannel Authentication Response

This API prepares JSON that contains an `auth_req_id`. The JSON should be used as the response body
of the response which is returned to the client from the [backchannel authentication endpoint](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html#auth_backchannel_endpoint)

<br>
<details>
<summary>Description</summary>

This API is supposed to be called from within the implementation of the backchannel authentication
endpoint of the service in order to generate a successful response to the client application.

The description of the `/backchannel/authentication` API describes the timing when this API should
be called and the meaning of request parameters. See [AUTH_REQ_ID ISSUE] in `USER_IDENTIFICATION`.

The response from `/backchannel/authentication/issue` API has some parameters. Among them, it is
`action` parameter that the authorization server implementation should check first because it denotes
the next `action` that the authorization server implementation should take. According to the value
of `action`, the authorization server implementation must take the steps described below.

```java
@POST
@Consumes(MediaType.APPLICATION_FORM_URLENCODED)
public Response post(String parameters)

	{
	    // 'parameters' is the entity body of the backchannel authentication request.
	    ......
	}

```

The endpoint implementation does not have to parse the request parameters from the client application
because Authlete's `/backchannel/authentication` API does it.

The response from `/backchannel/authentication` API has various parameters. Among them, it is `action`
parameter that the authorization server implementation should check first because it denotes the
next action that the authorization server implementation should take. According to the value of
`action`, the service implementation must take the steps described below.

**INTERNAL_SERVER_ERROR**

When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the request from the authorization
server implementation was wrong or that an error occurred in Authlete.
In either case, from the viewpoint of the client application, it is an error on the server side.
Therefore, the service implementation should generate a response to the client application with
HTTP status of "500 Internal Server Error" and `application/json`.

The value of `responseContent` is a JSON string which describes the error, so it can be used as
the entity body of the response.

The following illustrates the response which the service implementation should generate and return
to the client application.

```
HTTP/1.1 500 Internal Server Error
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

**INVALID_TICKET**

When the value of `action` is `INVALID_TICKET`, it means that the ticket included in the API call
was invalid. For example, it does not exist or has expired.

From a viewpoint of the client application, this is an error on the server side. Therefore, the
authorization server implementation should generate a response to the client application with
"500 Internal Server Error" and `application/json`.

You can build an error response in the same way as shown in the description for the case of `INTERNAL_SERVER_ERROR`.

**OK**

When the value of `action` is `OK`, it means that Authlete has succeeded in preparing JSON that
contains an `auth_req_id`. The JSON should be used as the response body of the response that is
returned to the client from the backchannel authentication endpoint. `responseContent` contains
the JSON.

The following illustrates the response which the authorization server implementation should generate
and return to the client application.

```
HTTP/1.1 200 OK
Content-Type: text/html;charset=UTF-8
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```
</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@return ApiBackchannelAuthenticationIssueApiRequest
*/
func (a *CIBAAPIService) BackchannelAuthenticationIssueApi(ctx context.Context, serviceId string) ApiBackchannelAuthenticationIssueApiRequest {
	return ApiBackchannelAuthenticationIssueApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return BackchannelAuthenticationIssueResponse
func (a *CIBAAPIService) BackchannelAuthenticationIssueApiExecute(r ApiBackchannelAuthenticationIssueApiRequest) (*BackchannelAuthenticationIssueResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BackchannelAuthenticationIssueResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CIBAAPIService.BackchannelAuthenticationIssueApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/backchannel/authentication/issue"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.backchannelAuthenticationIssueRequest == nil {
		return localVarReturnValue, nil, reportError("backchannelAuthenticationIssueRequest is required and must be specified")
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
	localVarPostBody = r.backchannelAuthenticationIssueRequest
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
