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

type TokenEndpointAPI interface {

	/*
			AuthTokenApi Process Token Request

			This API parses request parameters of an authorization request and returns necessary data for the
		authorization server implementation to process the authorization request further.

		<br>
		<details>
		<summary>Description</summary>

		This API is supposed to be called from with the implementation of the token endpoint of the service.
		The endpoint implementation must extract the request parameters from the token request from the
		client application and pass them as the value of parameters request parameter to Authlete's `/auth/token` API.
		The value of parameters is the entire entity body (which is formatted in `application/x-www-form-urlencoded`)
		of the token request.

		In addition, if the token endpoint of the authorization server implementation supports basic authentication
		as a means of [client authentication](https://datatracker.ietf.org/doc/html/rfc6749#section-2.3),
		the client credentials must be extracted from `Authorization` header and they must be passed as
		`clientId` request parameter and `clientSecret` request parameter to Authlete's `/auth/token` API.

		The following code snippet is an example in JAX-RS showing how to extract request parameters from
		the token request and client credentials from Authorization header.

		```java
		@POST
		@Consumes(MediaType.APPLICATION_FORM_URLENCODED)
		public Response post(
		        @HeaderParam(HttpHeaders.AUTHORIZATION) String auth,
		        String parameters)
		{
		    // Convert the value of Authorization header (credentials of
		    // the client application), if any, into BasicCredentials.
		    BasicCredentials credentials = BasicCredentials.parse(auth);

		    // The credentials of the client application extracted from
		    // 'Authorization' header. These may be null.
		    String clientId     = credentials == null ? null
		                        : credentials.getUserId();
		    String clientSecret = credentials == null ? null
		                        : credentials.getPassword();

		    // Process the given parameters.
		    return process(parameters, clientId, clientSecret);
		}
		```

		The response from `/auth/token` API has some parameters. Among them, it is action parameter that
		the service implementation should check first because it denotes the next action that the authorization
		server implementation should take. According to the value of action, the authorization server
		implementation must take the steps described below.

		**INTERNAL_SERVER_ERROR**

		When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the request from the authorization
		server implementation was wrong or that an error occurred in Authlete.
		In either case, from the viewpoint of the client application, it is an error on the server side.
		Therefore, the service implementation should generate a response to the client application with
		HTTP status of "500 Internal Server Error". Authlete recommends `application/json` as the content
		type although OAuth 2.0 specification does not mention the format of the error response when the
		redirect URI is not usable.

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

		The endpoint implementation may return another different response to the client application
		since "500 Internal Server Error" is not required by OAuth 2.0.


		**INVALID_CLIENT**

		When the value of `action` is `INVALID_CLIENT`, it means that authentication of the client failed.
		In this case, the HTTP status of the response to the client application is either "400 Bad Request"
		or "401 Unauthorized". This requirement comes from [RFC 6749, 5.2. Error Response](https://datatracker.ietf.org/doc/html/rfc6749#section-5.2).
		The description about `invalid_client` shown below is an excerpt from RFC 6749.

		Client authentication failed (e.g., unknown client, no client authentication included, or unsupported
		authentication method). The authorization server MAY return an HTTP 401 (Unauthorized) status code
		to indicate which HTTP authentication schemes are supported. If the client attempted to authenticate
		via the `Authorization` request header field, the authorization server MUST respond with an HTTP
		401 (Unauthorized) status code and include the `WWW-Authenticate` response header field matching
		the authentication scheme used by the client.

		In either case, the value of `responseContent` is a JSON string which can be used as the entity
		body of the response to the client application.

		The following illustrate responses which the service implementation must generate and return to
		the client application.

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

		A response with HTTP status of "400 Bad Request" must be returned to the client application and
		the content type must be `application/json`.

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
		**PASSWORD**

		When the value of `"action"` is `"PASSWORD"`, it means that
		the request from the client application is valid and `grant_type`
		is `"password"`. That is, the flow is
		<a href="https://www.rfc-editor.org/rfc/rfc6749.html#section-4.3">"Resource Owner
		Password Credentials"</a>.

		In this case, {@link #getUsername()} returns the value of `"username"`
		request parameter and {@link #getPassword()} returns the value of {@code
		"password"} request parameter which were contained in the token request
		from the client application. The service implementation must validate the
		credentials of the resource owner (= end-user) and take either of the
		actions below according to the validation result.

		1. When the credentials are valid, call Authlete's /auth/token/issue} API to generate an access token for the client
		    application. The API requires `"ticket"` request parameter and
		    `"subject"` request parameter.
		    Use the value returned from {@link #getTicket()} method as the value
		    for `"ticket"` parameter.
		2. The response from `/auth/token/issue` API ({@link
		    TokenIssueResponse}) contains data (an access token and others)
		    which should be returned to the client application. Use the data
		    to generate a response to the client application.
		3. When the credentials are invalid</b>, call Authlete's {@code
		    /auth/token/fail} API with `reason=`{@link
		    TokenFailRequest.Reason#INVALID_RESOURCE_OWNER_CREDENTIALS
		    INVALID_RESOURCE_OWNER_CREDENTIALS} to generate an error response
		    for the client application. The API requires `"ticket"`
		    request parameter. Use the value returned from {@link #getTicket()}
		    method as the value for `"ticket"` parameter.
		4. The response from `/auth/token/fail` API ({@link
		    TokenFailResponse}) contains error information which should be
		    returned to the client application. Use it to generate a response
		    to the client application.

		**OK**

		When the value of `action` is `OK`, it means that the request from the client application is valid
		and an access token, and optionally an ID token, is ready to be issued.

		The HTTP status of the response returned to the client application must be "200 OK" and the content
		type must be `application/json`.

		The value of `responseContent` is a JSON string which contains an access token (and optionally
		an ID token), so it can be used as the entity body of the response.

		The following illustrates the response which the service implementation must generate and return
		to the client application.

		```
		HTTP/1.1 200 OK
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```


		**TOKEN_EXCHANGE (Authlete 2.3 onwards)**

		When the value of `"action"` is `"TOKEN_EXCHANGE"`, it means
		that the request from the client application is a valid token exchange
		request (cf. <a href="https://www.rfc-editor.org/rfc/rfc8693.html">RFC
		8693 OAuth 2.0 Token Exchange</a>) and that the request has already passed
		the following validation steps.


		1. Confirm that the value of the `requested_token_type` request parameter
		is one of the registered token type identifiers if the request parameter is
		given and its value is not empty.
		2. Confirm that the `subject_token` request parameter is given and its
		value is not empty.
		3. Confirm that the `subject_token_type` request parameter is given and
		its value is one of the registered token type identifiers.
		4. Confirm that the `actor_token_type` request parameter is given and
		its value is one of the registered token type identifiers if the
		`actor_token` request parameter is given and its value is not empty.
		5. Confirm that the `actor_token_type` request parameter is not given
		or its value is empty when the `actor_token` request parameter is
		not given or its value is empty.


		Furthermore, Authlete performs additional validation on the tokens specified
		by the `subject_token` request parameter and the `actor_token`
		request parameter according to their respective token types as shown below.

		**Token Validation Steps**

		*Token Type: `urn:ietf:params:oauth:token-type:jwt`*

		1. Confirm that the format conforms to the JWT specification [RFC 7519][https://www.rfc-editor.org/rfc/rfc7519.html].
		2. Check if the JWT is encrypted and if it is encrypted, then (a) reject
		        the token exchange request when the {@link
		        Service#isTokenExchangeEncryptedJwtRejected()
		        tokenExchangeEncryptedJwtRejected} flag of the service is `true`
		        or (b) skip remaining validation steps when the flag is `false`.
		        Note that Authlete does not verify an encrypted JWT because there is
		        no standard way to obtain the key to decrypt the JWT with. This means
		        that you must verify an encrypted JWT by yourself when one is used as
		        an input token with the token type
		        { @code "urn:ietf:params:oauth:token-type:jwt" }.
		3. Confirm that the current time has not reached the time indicated by
		        the `exp` claim if the JWT contains the claim.
		4. Confirm that the current time is equal to or after the time indicated
		        by the `iat` claim if the JWT contains the claim.
		5.Confirm that the current time is equal to or after the time indicated
		        by the `nbf` claim if the JWT contains the claim.
		6. Check if the JWT is signed and if it is not signed, then (a) reject
		        the token exchange request when the {@link
		        Service#isTokenExchangeUnsignedJwtRejected()
		        tokenExchangeUnsignedJwtRejected} flag of the service is `true`
		        or (b) finish validation on the input token. Note that Authlete does
		        not verify the signature of the JWT because there is no standard way
		        to obtain the key to verify the signature of a JWT with. This means
		        that you must verify the signature by yourself when a signed JWT is
		        used as an input token with the token type
		        `"urn:ietf:params:oauth:token-type:jwt"`.

		*Token Type: `urn:ietf:params:oauth:token-type:access_token`*

		1. Confirm that the token is an access token that has been issued by
		        the Authlete server of your service. This implies that access
		        tokens issued by other systems cannot be used as a subject token
		        or an actor token with the token type
		        <code>urn:ietf:params:oauth:token-type:access_token</code>.
		2. Confirm that the access token has not expired.
		3. Confirm that the access token belongs to the service.

		*Token Type: `urn:ietf:params:oauth:token-type:refresh_token`*

		1. Confirm that the token is a refresh token that has been issued by
		        the Authlete server of your service. This implies that refresh
		        tokens issued by other systems cannot be used as a subject token
		        or an actor token with the token type
		        <code>urn:ietf:params:oauth:token-type:refresh_token</code>.
		2. Confirm that the refresh token has not expired.
		3. Confirm that the refresh token belongs to the service.

		*Token Type: `urn:ietf:params:oauth:token-type:id_token`*

		1. Confirm that the format conforms to the JWT specification (<a href=
		        "https://www.rfc-editor.org/rfc/rfc7519.html">RFC 7519</a>).
		2. Check if the ID Token is encrypted and if it is encrypted, then (a)
		        reject the token exchange request when the {@link
		        Service#isTokenExchangeEncryptedJwtRejected()
		        tokenExchangeEncryptedJwtRejected} flag of the service is `true`
		        or (b) skip remaining validation steps when the flag is `false`.
		        Note that Authlete does not verify an encrypted ID Token because
		        there is no standard way to obtain the key to decrypt the ID Token
		        with in the context of token exchange where the client ID for the
		        encrypted ID Token cannot be determined. This means that you must
		        verify an encrypted ID Token by yourself when one is used as an
		        input token with the token type
		        `"urn:ietf:params:oauth:token-type:id_token"`.
		3. Confirm that the ID Token contains the `exp` claim and the
		        current time has not reached the time indicated by the claim.
		4. Confirm that the ID Token contains the `iat` claim and the
		        current time is equal to or after the time indicated by the claim.
		5. Confirm that the current time is equal to or after the time indicated
		        by the `nbf` claim if the ID Token contains the claim.
		6. Confirm that the ID Token contains the `iss` claim and the
		        value is a valid URI. In addition, confirm that the URI has the
		        `https` scheme, no query component and no fragment component.
		7. Confirm that the ID Token contains the `aud` claim and its
		        value is a JSON string or an array of JSON strings.
		8. Confirm that the value of the `nonce` claim is a JSON string
		        if the ID Token contains the claim.
		9. Check if the ID Token is signed and if it is not signed, then (a)
		        reject the token exchange request when the {@link
		        Service#isTokenExchangeUnsignedJwtRejected()
		        tokenExchangeUnsignedJwtRejected} flag of the service is `true`
		        or (b) finish validation on the input token.
		10. Confirm that the signature algorithm is asymmetric. This implies that
		        ID Tokens whose signature algorithm is symmetric (`HS256`,
		        `HS384` or `HS512`) cannot be used as a subject token or
		        an actor token with the token type
		        `urn:ietf:params:oauth:token-type:id_token`.
		11. Verify the signature of the ID Token. Signature verification is
		        performed even in the case where the issuer of the ID Token is not
		        your service. But in that case, the issuer must support the discovery
		        endpoint defined in <a href=
		        "https://openid.net/specs/openid-connect-discovery-1_0.html">OpenID
		        Connect Discovery 1.0</a>. Otherwise, signature verification fails.

		*Token Type: `urn:ietf:params:oauth:token-type:saml1`*

		(Authlete does not perform any validation for this token type.)

		*Token Type: `urn:ietf:params:oauth:token-type:saml2`*

		(Authlete does not perform any validation for this token type.)

		The specification of Token Exchange (<a href=
		"https://www.rfc-editor.org/rfc/rfc8693.html">RFC 8693</a>) is very
		flexible. In other words, the specification has abandoned the task of
		determining details. Therefore, for secure token exchange, you have
		to complement the specification with your own rules. For that purpose,
		Authlete provides some configuration options as listed below.
		Authorization server implementers may utilize them and/or implement
		their own rules.


		In the case of {@link Action#TOKEN_EXCHANGE TOKEN_EXCHANGE}, the {@link
		#getResponseContent()} method returns `null`. You have to construct
		the token response by yourself.

		For example, you may generate an access token by calling Authlete's
		`/api/auth/token/create` API and construct a token response like
		below.

		```
		HTTP/1.1 401 Unauthorized
		WWW-Authenticate: {challenge}
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		```
		HTTP/1.1 200 OK
		Content-Type: application/json
		Cache-Control: no-cache, no-store
		{
		    "access_token": "{@link TokenCreateResponse#getAccessToken()}",
		    "issued_token_type": "urn:ietf:params:oauth:token-type:access_token",
		    "token_type": "Bearer",
		    "expires_in": { @link TokenCreateResponse#getExpiresIn() },
		    "scope": "String.join(" ", {@link TokenCreateResponse#getScopes()})"
		}
		```


		**JWT_BEARER JWT_BEARER (Authlete 2.3 onwards)**

		When the value of `"action"` is `"JWT_BEARER"`, it means that
		the request from the client application is a valid token request with the
		grant type `"urn:ietf:params:oauth:grant-type:jwt-bearer"` (<a href=
		"https://www.rfc-editor.org/rfc/rfc7523.html">RFC 7523 JSON Web Token (JWT)
		Profile for OAuth 2.0 Client Authentication and Authorization Grants</a>)
		and that the request has already passed the following validation steps.

		1. Confirm that the `assertion` request parameter is given and its value
		  is not empty.
		2. Confirm that the format of the assertion conforms to the JWT specification
		  (<a href="https://www.rfc-editor.org/rfc/rfc7519.html">RFC 7519</a>).

		3. Check if the JWT is encrypted and if it is encrypted, then (a) reject the
		  token request when the {@link Service#isJwtGrantEncryptedJwtRejected()
		  jwtGrantEncryptedJwtRejected} flag of the service is `true` or (b)
		  skip remaining validation steps when the flag is `false`. Note that
		  Authlete does not verify an encrypted JWT because there is no standard way
		  to obtain the key to decrypt the JWT with. This means that you must verify
		  an encrypted JWT by yourself.
		4. Confirm that the JWT contains the `iss` claim and its value is a
		  JSON string.
		5. Confirm that the JWT contains the `sub` claim and its value is a
		  JSON string.
		6. Confirm that the JWT contains the `aud` claim and its value is
		  either a JSON string or an array of JSON strings.
		7. Confirm that the issuer identifier of the service (cf. {@link Service#getIssuer()})
		  or the URL of the token endpoint (cf. {@link Service#getTokenEndpoint()})
		  is listed as audience in the `aud` claim.
		8. Confirm that the JWT contains the `exp` claim and the current time
		  has not reached the time indicated by the claim.
		9. Confirm that the current time is equal to or after the time indicated by
		  by the `iat` claim if the JWT contains the claim.
		10. Confirm that the current time is equal to or after the time indicated by
		  by the `nbf` claim if the JWT contains the claim.
		11. Check if the JWT is signed and if it is not signed, then (a) reject the
		  token request when the {@link Service#isJwtGrantUnsignedJwtRejected()
		  jwtGrantUnsignedJwtRejected} flag of the service is `true` or (b)
		  finish validation on the JWT. Note that Authlete does not verify the
		  signature of the JWT because there is no standard way to obtain the key
		  to verify the signature of a JWT with. This means that you must verify
		  the signature by yourself.

		Authlete provides some configuration options for the grant type as listed
		below. Authorization server implementers may utilize them and/or implement
		their own rules.


		```
		HTTP/1.1 200 OK
		Content-Type: application/json
		Cache-Control: no-cache, no-store
		{
		      "access_token": "{@link TokenCreateResponse#getAccessToken()}",
		      "token_type":   "Bearer",
		      "expires_in":   {@link TokenCreateResponse#getExpiresIn()},
		      "scope":        "String.join(" ", {@link TokenCreateResponse#getScopes()})"
		                          }
		```

		  Finally, note again that Authlete does not verify the signature of the JWT
		  specified by the `assertion` request parameter. You must verify the
		  signature by yourself.

		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@return ApiAuthTokenApiRequest
	*/
	AuthTokenApi(ctx context.Context, serviceId string) ApiAuthTokenApiRequest

	// AuthTokenApiExecute executes the request
	//  @return TokenResponse
	AuthTokenApiExecute(r ApiAuthTokenApiRequest) (*TokenResponse, *http.Response, error)

	/*
			AuthTokenFailApi Fail Token Request

			This API generates a content of an error token response that the authorization server implementation
		returns to the client application.

		<br>
		<details>
		<summary>Description</summary>

		This API is supposed to be called from within the implementation of the token endpoint of the service
		in order to generate an error response to the client application.

		The description of the `/auth/token` API describes the timing when this API should be called. See
		the description for the case of `action=PASSWORD`.

		The response from `/auth/token/fail` API has some parameters. Among them, it is `action` parameter
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

		The endpoint implementation may return another different response to the client application
		since "500 Internal Server Error" is not required by OAuth 2.0.

		**BAD_REQUEST**

		When the value of `action` is `BAD_REQUEST`, it means that Authlete's `/auth/token/fail` API successfully
		generated an error response for the client application.

		The HTTP status of the response returned to the client application must be "400 Bad Request" and
		the content type must be `application/json`.

		The value of `responseContent` is a JSON string which describes the error, so it can be used
		as the entity body of the response.

		The following illustrates the response which the service implementation should generate and return
		to the client application.

		```
		HTTP/1.1 400 Bad Request
		Content-Type: application/json
		Cache-Control: no-store
		Pragma: no-cache

		{responseContent}
		```

		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@return ApiAuthTokenFailApiRequest
	*/
	AuthTokenFailApi(ctx context.Context, serviceId string) ApiAuthTokenFailApiRequest

	// AuthTokenFailApiExecute executes the request
	//  @return TokenFailResponse
	AuthTokenFailApiExecute(r ApiAuthTokenFailApiRequest) (*TokenFailResponse, *http.Response, error)

	/*
			AuthTokenIssueApi Issue Token Response

			This API generates a content of a successful token response that the authorization server implementation
		returns to the client application.

		<br>
		<details>
		<summary>Description</summary>

		This API is supposed to be called from within the implementation of the token endpoint of the service
		in order to generate a successful response to the client application.

		The description of the `/auth/token` API describes the timing when this API should be called. See
		the description for the case of `action=PASSWORD`.

		The response from `/auth/token/issue` API has some parameters. Among them, it is `action` parameter
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

		The endpoint implementation may return another different response to the client application
		since "500 Internal Server Error" is not required by OAuth 2.0.

		**OK**

		When the value of `action` is `OK`, it means that Authlete's `/auth/token/issue` API successfully
		generated an access token.

		The HTTP status of the response returned to the client application must be "200 OK" and the content
		type must be`application/json`.

		The value of `responseContent` is a JSON string which contains an access token, so it can be used
		as the entity body of the response.

		The following illustrates the response which the service implementation must generate and return
		to the client application.

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
			@return ApiAuthTokenIssueApiRequest
	*/
	AuthTokenIssueApi(ctx context.Context, serviceId string) ApiAuthTokenIssueApiRequest

	// AuthTokenIssueApiExecute executes the request
	//  @return TokenIssueResponse
	AuthTokenIssueApiExecute(r ApiAuthTokenIssueApiRequest) (*TokenIssueResponse, *http.Response, error)

	/*
			IdtokenReissueApi Reissue ID Token

			The API is expected to be called only when the value of the `action`
		parameter in a response from the `/auth/token` API is [ID_TOKEN_REISSUABLE](https://authlete.github.io/authlete-java-common/com/authlete/common/dto/TokenResponse.Action.html#ID_TOKEN_REISSUABLE). The purpose
		of the `/idtoken/reissue` API is to generate a token response that
		includes a new ID token together with a new access token and a refresh
		token.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@return ApiIdtokenReissueApiRequest
	*/
	IdtokenReissueApi(ctx context.Context, serviceId string) ApiIdtokenReissueApiRequest

	// IdtokenReissueApiExecute executes the request
	//  @return IdtokenReissueResponse
	IdtokenReissueApiExecute(r ApiIdtokenReissueApiRequest) (*IdtokenReissueResponse, *http.Response, error)
}

// TokenEndpointAPIService TokenEndpointAPI service
type TokenEndpointAPIService service

type ApiAuthTokenApiRequest struct {
	ctx          context.Context
	ApiService   TokenEndpointAPI
	serviceId    string
	tokenRequest *TokenRequest
}

func (r ApiAuthTokenApiRequest) TokenRequest(tokenRequest TokenRequest) ApiAuthTokenApiRequest {
	r.tokenRequest = &tokenRequest
	return r
}

func (r ApiAuthTokenApiRequest) Execute() (*TokenResponse, *http.Response, error) {
	return r.ApiService.AuthTokenApiExecute(r)
}

/*
AuthTokenApi Process Token Request

This API parses request parameters of an authorization request and returns necessary data for the
authorization server implementation to process the authorization request further.

<br>
<details>
<summary>Description</summary>

This API is supposed to be called from with the implementation of the token endpoint of the service.
The endpoint implementation must extract the request parameters from the token request from the
client application and pass them as the value of parameters request parameter to Authlete's `/auth/token` API.
The value of parameters is the entire entity body (which is formatted in `application/x-www-form-urlencoded`)
of the token request.

In addition, if the token endpoint of the authorization server implementation supports basic authentication
as a means of [client authentication](https://datatracker.ietf.org/doc/html/rfc6749#section-2.3),
the client credentials must be extracted from `Authorization` header and they must be passed as
`clientId` request parameter and `clientSecret` request parameter to Authlete's `/auth/token` API.

The following code snippet is an example in JAX-RS showing how to extract request parameters from
the token request and client credentials from Authorization header.

```java
@POST
@Consumes(MediaType.APPLICATION_FORM_URLENCODED)
public Response post(

	@HeaderParam(HttpHeaders.AUTHORIZATION) String auth,
	String parameters)

	{
	    // Convert the value of Authorization header (credentials of
	    // the client application), if any, into BasicCredentials.
	    BasicCredentials credentials = BasicCredentials.parse(auth);

	    // The credentials of the client application extracted from
	    // 'Authorization' header. These may be null.
	    String clientId     = credentials == null ? null
	                        : credentials.getUserId();
	    String clientSecret = credentials == null ? null
	                        : credentials.getPassword();

	    // Process the given parameters.
	    return process(parameters, clientId, clientSecret);
	}

```

The response from `/auth/token` API has some parameters. Among them, it is action parameter that
the service implementation should check first because it denotes the next action that the authorization
server implementation should take. According to the value of action, the authorization server
implementation must take the steps described below.

**INTERNAL_SERVER_ERROR**

When the value of `action` is `INTERNAL_SERVER_ERROR`, it means that the request from the authorization
server implementation was wrong or that an error occurred in Authlete.
In either case, from the viewpoint of the client application, it is an error on the server side.
Therefore, the service implementation should generate a response to the client application with
HTTP status of "500 Internal Server Error". Authlete recommends `application/json` as the content
type although OAuth 2.0 specification does not mention the format of the error response when the
redirect URI is not usable.

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

The endpoint implementation may return another different response to the client application
since "500 Internal Server Error" is not required by OAuth 2.0.

**INVALID_CLIENT**

When the value of `action` is `INVALID_CLIENT`, it means that authentication of the client failed.
In this case, the HTTP status of the response to the client application is either "400 Bad Request"
or "401 Unauthorized". This requirement comes from [RFC 6749, 5.2. Error Response](https://datatracker.ietf.org/doc/html/rfc6749#section-5.2).
The description about `invalid_client` shown below is an excerpt from RFC 6749.

Client authentication failed (e.g., unknown client, no client authentication included, or unsupported
authentication method). The authorization server MAY return an HTTP 401 (Unauthorized) status code
to indicate which HTTP authentication schemes are supported. If the client attempted to authenticate
via the `Authorization` request header field, the authorization server MUST respond with an HTTP
401 (Unauthorized) status code and include the `WWW-Authenticate` response header field matching
the authentication scheme used by the client.

In either case, the value of `responseContent` is a JSON string which can be used as the entity
body of the response to the client application.

The following illustrate responses which the service implementation must generate and return to
the client application.

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

A response with HTTP status of "400 Bad Request" must be returned to the client application and
the content type must be `application/json`.

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
**PASSWORD**

When the value of `"action"` is `"PASSWORD"`, it means that
the request from the client application is valid and `grant_type`
is `"password"`. That is, the flow is
<a href="https://www.rfc-editor.org/rfc/rfc6749.html#section-4.3">"Resource Owner
Password Credentials"</a>.

In this case, {@link #getUsername()} returns the value of `"username"`
request parameter and {@link #getPassword()} returns the value of {@code
"password"} request parameter which were contained in the token request
from the client application. The service implementation must validate the
credentials of the resource owner (= end-user) and take either of the
actions below according to the validation result.

 1. When the credentials are valid, call Authlete's /auth/token/issue} API to generate an access token for the client
    application. The API requires `"ticket"` request parameter and
    `"subject"` request parameter.
    Use the value returned from {@link #getTicket()} method as the value
    for `"ticket"` parameter.
 2. The response from `/auth/token/issue` API ({@link
    TokenIssueResponse}) contains data (an access token and others)
    which should be returned to the client application. Use the data
    to generate a response to the client application.
 3. When the credentials are invalid</b>, call Authlete's {@code
    /auth/token/fail} API with `reason=`{@link
    TokenFailRequest.Reason#INVALID_RESOURCE_OWNER_CREDENTIALS
    INVALID_RESOURCE_OWNER_CREDENTIALS} to generate an error response
    for the client application. The API requires `"ticket"`
    request parameter. Use the value returned from {@link #getTicket()}
    method as the value for `"ticket"` parameter.
 4. The response from `/auth/token/fail` API ({@link
    TokenFailResponse}) contains error information which should be
    returned to the client application. Use it to generate a response
    to the client application.

**OK**

When the value of `action` is `OK`, it means that the request from the client application is valid
and an access token, and optionally an ID token, is ready to be issued.

The HTTP status of the response returned to the client application must be "200 OK" and the content
type must be `application/json`.

The value of `responseContent` is a JSON string which contains an access token (and optionally
an ID token), so it can be used as the entity body of the response.

The following illustrates the response which the service implementation must generate and return
to the client application.

```
HTTP/1.1 200 OK
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

**TOKEN_EXCHANGE (Authlete 2.3 onwards)**

When the value of `"action"` is `"TOKEN_EXCHANGE"`, it means
that the request from the client application is a valid token exchange
request (cf. <a href="https://www.rfc-editor.org/rfc/rfc8693.html">RFC
8693 OAuth 2.0 Token Exchange</a>) and that the request has already passed
the following validation steps.

1. Confirm that the value of the `requested_token_type` request parameter
is one of the registered token type identifiers if the request parameter is
given and its value is not empty.
2. Confirm that the `subject_token` request parameter is given and its
value is not empty.
3. Confirm that the `subject_token_type` request parameter is given and
its value is one of the registered token type identifiers.
4. Confirm that the `actor_token_type` request parameter is given and
its value is one of the registered token type identifiers if the
`actor_token` request parameter is given and its value is not empty.
5. Confirm that the `actor_token_type` request parameter is not given
or its value is empty when the `actor_token` request parameter is
not given or its value is empty.

Furthermore, Authlete performs additional validation on the tokens specified
by the `subject_token` request parameter and the `actor_token`
request parameter according to their respective token types as shown below.

**Token Validation Steps**

*Token Type: `urn:ietf:params:oauth:token-type:jwt`*

 1. Confirm that the format conforms to the JWT specification [RFC 7519][https://www.rfc-editor.org/rfc/rfc7519.html].
 2. Check if the JWT is encrypted and if it is encrypted, then (a) reject
    the token exchange request when the {@link
    Service#isTokenExchangeEncryptedJwtRejected()
    tokenExchangeEncryptedJwtRejected} flag of the service is `true`
    or (b) skip remaining validation steps when the flag is `false`.
    Note that Authlete does not verify an encrypted JWT because there is
    no standard way to obtain the key to decrypt the JWT with. This means
    that you must verify an encrypted JWT by yourself when one is used as
    an input token with the token type
    { @code "urn:ietf:params:oauth:token-type:jwt" }.
 3. Confirm that the current time has not reached the time indicated by
    the `exp` claim if the JWT contains the claim.
 4. Confirm that the current time is equal to or after the time indicated
    by the `iat` claim if the JWT contains the claim.

5.Confirm that the current time is equal to or after the time indicated

		by the `nbf` claim if the JWT contains the claim.
	 6. Check if the JWT is signed and if it is not signed, then (a) reject
	    the token exchange request when the {@link
	    Service#isTokenExchangeUnsignedJwtRejected()
	    tokenExchangeUnsignedJwtRejected} flag of the service is `true`
	    or (b) finish validation on the input token. Note that Authlete does
	    not verify the signature of the JWT because there is no standard way
	    to obtain the key to verify the signature of a JWT with. This means
	    that you must verify the signature by yourself when a signed JWT is
	    used as an input token with the token type
	    `"urn:ietf:params:oauth:token-type:jwt"`.

*Token Type: `urn:ietf:params:oauth:token-type:access_token`*

 1. Confirm that the token is an access token that has been issued by
    the Authlete server of your service. This implies that access
    tokens issued by other systems cannot be used as a subject token
    or an actor token with the token type
    <code>urn:ietf:params:oauth:token-type:access_token</code>.
 2. Confirm that the access token has not expired.
 3. Confirm that the access token belongs to the service.

*Token Type: `urn:ietf:params:oauth:token-type:refresh_token`*

 1. Confirm that the token is a refresh token that has been issued by
    the Authlete server of your service. This implies that refresh
    tokens issued by other systems cannot be used as a subject token
    or an actor token with the token type
    <code>urn:ietf:params:oauth:token-type:refresh_token</code>.
 2. Confirm that the refresh token has not expired.
 3. Confirm that the refresh token belongs to the service.

*Token Type: `urn:ietf:params:oauth:token-type:id_token`*

 1. Confirm that the format conforms to the JWT specification (<a href=
    "https://www.rfc-editor.org/rfc/rfc7519.html">RFC 7519</a>).
 2. Check if the ID Token is encrypted and if it is encrypted, then (a)
    reject the token exchange request when the {@link
    Service#isTokenExchangeEncryptedJwtRejected()
    tokenExchangeEncryptedJwtRejected} flag of the service is `true`
    or (b) skip remaining validation steps when the flag is `false`.
    Note that Authlete does not verify an encrypted ID Token because
    there is no standard way to obtain the key to decrypt the ID Token
    with in the context of token exchange where the client ID for the
    encrypted ID Token cannot be determined. This means that you must
    verify an encrypted ID Token by yourself when one is used as an
    input token with the token type
    `"urn:ietf:params:oauth:token-type:id_token"`.
 3. Confirm that the ID Token contains the `exp` claim and the
    current time has not reached the time indicated by the claim.
 4. Confirm that the ID Token contains the `iat` claim and the
    current time is equal to or after the time indicated by the claim.
 5. Confirm that the current time is equal to or after the time indicated
    by the `nbf` claim if the ID Token contains the claim.
 6. Confirm that the ID Token contains the `iss` claim and the
    value is a valid URI. In addition, confirm that the URI has the
    `https` scheme, no query component and no fragment component.
 7. Confirm that the ID Token contains the `aud` claim and its
    value is a JSON string or an array of JSON strings.
 8. Confirm that the value of the `nonce` claim is a JSON string
    if the ID Token contains the claim.
 9. Check if the ID Token is signed and if it is not signed, then (a)
    reject the token exchange request when the {@link
    Service#isTokenExchangeUnsignedJwtRejected()
    tokenExchangeUnsignedJwtRejected} flag of the service is `true`
    or (b) finish validation on the input token.
 10. Confirm that the signature algorithm is asymmetric. This implies that
    ID Tokens whose signature algorithm is symmetric (`HS256`,
    `HS384` or `HS512`) cannot be used as a subject token or
    an actor token with the token type
    `urn:ietf:params:oauth:token-type:id_token`.
 11. Verify the signature of the ID Token. Signature verification is
    performed even in the case where the issuer of the ID Token is not
    your service. But in that case, the issuer must support the discovery
    endpoint defined in <a href=
    "https://openid.net/specs/openid-connect-discovery-1_0.html">OpenID
    Connect Discovery 1.0</a>. Otherwise, signature verification fails.

*Token Type: `urn:ietf:params:oauth:token-type:saml1`*

(Authlete does not perform any validation for this token type.)

*Token Type: `urn:ietf:params:oauth:token-type:saml2`*

(Authlete does not perform any validation for this token type.)

The specification of Token Exchange (<a href=
"https://www.rfc-editor.org/rfc/rfc8693.html">RFC 8693</a>) is very
flexible. In other words, the specification has abandoned the task of
determining details. Therefore, for secure token exchange, you have
to complement the specification with your own rules. For that purpose,
Authlete provides some configuration options as listed below.
Authorization server implementers may utilize them and/or implement
their own rules.

In the case of {@link Action#TOKEN_EXCHANGE TOKEN_EXCHANGE}, the {@link
#getResponseContent()} method returns `null`. You have to construct
the token response by yourself.

For example, you may generate an access token by calling Authlete's
`/api/auth/token/create` API and construct a token response like
below.

```
HTTP/1.1 401 Unauthorized
WWW-Authenticate: {challenge}
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

```
HTTP/1.1 200 OK
Content-Type: application/json
Cache-Control: no-cache, no-store

	{
	    "access_token": "{@link TokenCreateResponse#getAccessToken()}",
	    "issued_token_type": "urn:ietf:params:oauth:token-type:access_token",
	    "token_type": "Bearer",
	    "expires_in": { @link TokenCreateResponse#getExpiresIn() },
	    "scope": "String.join(" ", {@link TokenCreateResponse#getScopes()})"
	}

```

**JWT_BEARER JWT_BEARER (Authlete 2.3 onwards)**

When the value of `"action"` is `"JWT_BEARER"`, it means that
the request from the client application is a valid token request with the
grant type `"urn:ietf:params:oauth:grant-type:jwt-bearer"` (<a href=
"https://www.rfc-editor.org/rfc/rfc7523.html">RFC 7523 JSON Web Token (JWT)
Profile for OAuth 2.0 Client Authentication and Authorization Grants</a>)
and that the request has already passed the following validation steps.

 1. Confirm that the `assertion` request parameter is given and its value
    is not empty.

 2. Confirm that the format of the assertion conforms to the JWT specification
    (<a href="https://www.rfc-editor.org/rfc/rfc7519.html">RFC 7519</a>).

 3. Check if the JWT is encrypted and if it is encrypted, then (a) reject the
    token request when the {@link Service#isJwtGrantEncryptedJwtRejected()
    jwtGrantEncryptedJwtRejected} flag of the service is `true` or (b)
    skip remaining validation steps when the flag is `false`. Note that
    Authlete does not verify an encrypted JWT because there is no standard way
    to obtain the key to decrypt the JWT with. This means that you must verify
    an encrypted JWT by yourself.

 4. Confirm that the JWT contains the `iss` claim and its value is a
    JSON string.

 5. Confirm that the JWT contains the `sub` claim and its value is a
    JSON string.

 6. Confirm that the JWT contains the `aud` claim and its value is
    either a JSON string or an array of JSON strings.

 7. Confirm that the issuer identifier of the service (cf. {@link Service#getIssuer()})
    or the URL of the token endpoint (cf. {@link Service#getTokenEndpoint()})
    is listed as audience in the `aud` claim.

 8. Confirm that the JWT contains the `exp` claim and the current time
    has not reached the time indicated by the claim.

 9. Confirm that the current time is equal to or after the time indicated by
    by the `iat` claim if the JWT contains the claim.

 10. Confirm that the current time is equal to or after the time indicated by
    by the `nbf` claim if the JWT contains the claim.

 11. Check if the JWT is signed and if it is not signed, then (a) reject the
    token request when the {@link Service#isJwtGrantUnsignedJwtRejected()
    jwtGrantUnsignedJwtRejected} flag of the service is `true` or (b)
    finish validation on the JWT. Note that Authlete does not verify the
    signature of the JWT because there is no standard way to obtain the key
    to verify the signature of a JWT with. This means that you must verify
    the signature by yourself.

Authlete provides some configuration options for the grant type as listed
below. Authorization server implementers may utilize them and/or implement
their own rules.

```
HTTP/1.1 200 OK
Content-Type: application/json
Cache-Control: no-cache, no-store

	{
	      "access_token": "{@link TokenCreateResponse#getAccessToken()}",
	      "token_type":   "Bearer",
	      "expires_in":   {@link TokenCreateResponse#getExpiresIn()},
	      "scope":        "String.join(" ", {@link TokenCreateResponse#getScopes()})"
	                          }

```

	Finally, note again that Authlete does not verify the signature of the JWT
	specified by the `assertion` request parameter. You must verify the
	signature by yourself.

</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@return ApiAuthTokenApiRequest
*/
func (a *TokenEndpointAPIService) AuthTokenApi(ctx context.Context, serviceId string) ApiAuthTokenApiRequest {
	return ApiAuthTokenApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return TokenResponse
func (a *TokenEndpointAPIService) AuthTokenApiExecute(r ApiAuthTokenApiRequest) (*TokenResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokenEndpointAPIService.AuthTokenApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/auth/token"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tokenRequest == nil {
		return localVarReturnValue, nil, reportError("tokenRequest is required and must be specified")
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
	localVarPostBody = r.tokenRequest
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

type ApiAuthTokenFailApiRequest struct {
	ctx              context.Context
	ApiService       TokenEndpointAPI
	serviceId        string
	tokenFailRequest *TokenFailRequest
}

func (r ApiAuthTokenFailApiRequest) TokenFailRequest(tokenFailRequest TokenFailRequest) ApiAuthTokenFailApiRequest {
	r.tokenFailRequest = &tokenFailRequest
	return r
}

func (r ApiAuthTokenFailApiRequest) Execute() (*TokenFailResponse, *http.Response, error) {
	return r.ApiService.AuthTokenFailApiExecute(r)
}

/*
AuthTokenFailApi Fail Token Request

This API generates a content of an error token response that the authorization server implementation
returns to the client application.

<br>
<details>
<summary>Description</summary>

This API is supposed to be called from within the implementation of the token endpoint of the service
in order to generate an error response to the client application.

The description of the `/auth/token` API describes the timing when this API should be called. See
the description for the case of `action=PASSWORD`.

The response from `/auth/token/fail` API has some parameters. Among them, it is `action` parameter
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

The endpoint implementation may return another different response to the client application
since "500 Internal Server Error" is not required by OAuth 2.0.

**BAD_REQUEST**

When the value of `action` is `BAD_REQUEST`, it means that Authlete's `/auth/token/fail` API successfully
generated an error response for the client application.

The HTTP status of the response returned to the client application must be "400 Bad Request" and
the content type must be `application/json`.

The value of `responseContent` is a JSON string which describes the error, so it can be used
as the entity body of the response.

The following illustrates the response which the service implementation should generate and return
to the client application.

```
HTTP/1.1 400 Bad Request
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{responseContent}
```

</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@return ApiAuthTokenFailApiRequest
*/
func (a *TokenEndpointAPIService) AuthTokenFailApi(ctx context.Context, serviceId string) ApiAuthTokenFailApiRequest {
	return ApiAuthTokenFailApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return TokenFailResponse
func (a *TokenEndpointAPIService) AuthTokenFailApiExecute(r ApiAuthTokenFailApiRequest) (*TokenFailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TokenFailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokenEndpointAPIService.AuthTokenFailApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/auth/token/fail"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tokenFailRequest == nil {
		return localVarReturnValue, nil, reportError("tokenFailRequest is required and must be specified")
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
	localVarPostBody = r.tokenFailRequest
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

type ApiAuthTokenIssueApiRequest struct {
	ctx               context.Context
	ApiService        TokenEndpointAPI
	serviceId         string
	tokenIssueRequest *TokenIssueRequest
}

func (r ApiAuthTokenIssueApiRequest) TokenIssueRequest(tokenIssueRequest TokenIssueRequest) ApiAuthTokenIssueApiRequest {
	r.tokenIssueRequest = &tokenIssueRequest
	return r
}

func (r ApiAuthTokenIssueApiRequest) Execute() (*TokenIssueResponse, *http.Response, error) {
	return r.ApiService.AuthTokenIssueApiExecute(r)
}

/*
AuthTokenIssueApi Issue Token Response

This API generates a content of a successful token response that the authorization server implementation
returns to the client application.

<br>
<details>
<summary>Description</summary>

This API is supposed to be called from within the implementation of the token endpoint of the service
in order to generate a successful response to the client application.

The description of the `/auth/token` API describes the timing when this API should be called. See
the description for the case of `action=PASSWORD`.

The response from `/auth/token/issue` API has some parameters. Among them, it is `action` parameter
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

The endpoint implementation may return another different response to the client application
since "500 Internal Server Error" is not required by OAuth 2.0.

**OK**

When the value of `action` is `OK`, it means that Authlete's `/auth/token/issue` API successfully
generated an access token.

The HTTP status of the response returned to the client application must be "200 OK" and the content
type must be`application/json`.

The value of `responseContent` is a JSON string which contains an access token, so it can be used
as the entity body of the response.

The following illustrates the response which the service implementation must generate and return
to the client application.

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
	@return ApiAuthTokenIssueApiRequest
*/
func (a *TokenEndpointAPIService) AuthTokenIssueApi(ctx context.Context, serviceId string) ApiAuthTokenIssueApiRequest {
	return ApiAuthTokenIssueApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return TokenIssueResponse
func (a *TokenEndpointAPIService) AuthTokenIssueApiExecute(r ApiAuthTokenIssueApiRequest) (*TokenIssueResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TokenIssueResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokenEndpointAPIService.AuthTokenIssueApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/auth/token/issue"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tokenIssueRequest == nil {
		return localVarReturnValue, nil, reportError("tokenIssueRequest is required and must be specified")
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
	localVarPostBody = r.tokenIssueRequest
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

type ApiIdtokenReissueApiRequest struct {
	ctx                   context.Context
	ApiService            TokenEndpointAPI
	serviceId             string
	idtokenReissueRequest *IdtokenReissueRequest
}

func (r ApiIdtokenReissueApiRequest) IdtokenReissueRequest(idtokenReissueRequest IdtokenReissueRequest) ApiIdtokenReissueApiRequest {
	r.idtokenReissueRequest = &idtokenReissueRequest
	return r
}

func (r ApiIdtokenReissueApiRequest) Execute() (*IdtokenReissueResponse, *http.Response, error) {
	return r.ApiService.IdtokenReissueApiExecute(r)
}

/*
IdtokenReissueApi Reissue ID Token

The API is expected to be called only when the value of the `action`
parameter in a response from the `/auth/token` API is [ID_TOKEN_REISSUABLE](https://authlete.github.io/authlete-java-common/com/authlete/common/dto/TokenResponse.Action.html#ID_TOKEN_REISSUABLE). The purpose
of the `/idtoken/reissue` API is to generate a token response that
includes a new ID token together with a new access token and a refresh
token.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@return ApiIdtokenReissueApiRequest
*/
func (a *TokenEndpointAPIService) IdtokenReissueApi(ctx context.Context, serviceId string) ApiIdtokenReissueApiRequest {
	return ApiIdtokenReissueApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return IdtokenReissueResponse
func (a *TokenEndpointAPIService) IdtokenReissueApiExecute(r ApiIdtokenReissueApiRequest) (*IdtokenReissueResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IdtokenReissueResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokenEndpointAPIService.IdtokenReissueApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/idtoken/reissue"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.idtokenReissueRequest
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
