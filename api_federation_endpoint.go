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

type FederationEndpointAPI interface {

	/*
			FederationConfigurationApi Process Entity Configuration Request

			This API gathers the federation configuration about a service.

		The authorization server implementation should
		retrieve the value of the <code>action</code>
		response parameter from the API response and take the following steps
		according to the value.

		<h3><code>OK</code></h3>

		When the value of the <code> action</code> response
		parameter is <code>OK</code>, it means that Authlete
		could prepare an entity configuration successfully.

		In this case, the implementation of the entity configuration endpoint of the
		authorization server should return an HTTP response to the client application
		with the HTTP status code "`200 OK`" and the content type
		"`application/entity-statement+jwt`". The message body (= an entity
		configuration in the JWT format) of the response has been prepared by
		Authlete's `/federation/configuration` API and it is available as the
		<code>responseContent</code> response parameter.

		The implementation of the entity configuration endpoint can construct an
		HTTP response by doing like below.

		<pre style="border: solid 1px black; padding: 0.5em;">
		200 OK
		Content-Type: application/entity-statement+jwt
		(Other HTTP headers)

		<i>(the value of the responseContent response parameter)</i></pre>

		<h3><code>NOT_FOUND</code></h3>

		When the value of the <code> action</code> response
		parameter is <code>NOT_FOUND</code>, it means that
		the service configuration has not enabled the feature of <a href=
		"https://openid.net/specs/openid-connect-federation-1_0.html">OpenID Connect
		Federation 1.0</a> and so the client application should have not access the
		entity configuration endpoint.

		In this case, the implementation of the entity configuration endpoint of the
		authorization server should return an HTTP response to the client application
		with the HTTP status code "`404 Not Found`" and the content type
		"`application/json`". The message body (= error information in the JSON
		format) of the response has been prepared by Authlete's
		`/federation/configuration` API and it is available as the
		<code>responseContent</code> response parameter.

		The implementation of the entity configuration endpoint can construct an
		HTTP response by doing like below.

		<pre style="border: solid 1px black; padding: 0.5em;">
		404 Not Found
		Content-Type: application/json
		(Other HTTP headers)

		<i>(the value of the responseContent response parameter)</i></pre>

		<h3><code>INTERNAL_SERVER_ERROR</code></h3>

		could prepare an entity configuration successfully.

		In this case, the implementation of the entity configuration endpoint of the
		authorization server should return an HTTP response to the client application
		with the HTTP status code "`200 OK`" and the content type
		"`application/entity-statement+jwt`". The message body (= an entity
		configuration in the JWT format) of the response has been prepared by
		Authlete's `/federation/configuration` API and it is available as the
		<code>responseContent</code> response parameter.

		The implementation of the entity configuration endpoint can construct an
		HTTP response by doing like below.

		<pre style="border: solid 1px black; padding: 0.5em;">
		200 OK
		Content-Type: application/entity-statement+jwt
		(Other HTTP headers)

		<i>(the value of the responseContent response parameter)</i></pre>


		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@return ApiFederationConfigurationApiRequest
	*/
	FederationConfigurationApi(ctx context.Context, serviceId string) ApiFederationConfigurationApiRequest

	// FederationConfigurationApiExecute executes the request
	//  @return FederationConfigurationResponse
	FederationConfigurationApiExecute(r ApiFederationConfigurationApiRequest) (*FederationConfigurationResponse, *http.Response, error)

	/*
			FederationRegistrationApi Process Federation Registration Request

			The Authlete API is for implementations of the <b>federation registration
		endpoint</b> that accepts "explicit client registration". Its details are
		defined in <a href="https://openid.net/specs/openid-connect-federation-1_0.html"
		>OpenID Connect Federation 1.0</a>.
		</p>

		<p>
		The endpoint accepts `POST` requests whose `Content-Type`
		is either of the following.
		</p>

		<ol>
		  <li>`application/entity-statement+jwt`
		  <li>`application/trust-chain+json`
		</ol>

		<p>
		When the `Content-Type` of a request is
		`application/entity-statement+jwt`, the content of the request is
		the entity configuration of a relying party that is to be registered.
		In this case, the implementation of the federation registration endpoint
		should call Authlete's `/federation/registration` API with the
		entity configuration set to the `entityConfiguration` request
		parameter.
		</p>

		<p>
		On the other hand, when the `Content-Type` of a request is
		`application/trust-chain+json`, the content of the request is a
		JSON array that contains entity statements in JWT format. The sequence
		of the entity statements composes the trust chain of a relying party
		that is to be registered. In this case, the implementation of the
		federation registration endpoint should call Authlete's
		`/federation/registration` API with the trust chain set to the
		`trustChain` request parameter.
		</p>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@return ApiFederationRegistrationApiRequest
	*/
	FederationRegistrationApi(ctx context.Context, serviceId string) ApiFederationRegistrationApiRequest

	// FederationRegistrationApiExecute executes the request
	//  @return FederationRegistrationResponse
	FederationRegistrationApiExecute(r ApiFederationRegistrationApiRequest) (*FederationRegistrationResponse, *http.Response, error)
}

// FederationEndpointAPIService FederationEndpointAPI service
type FederationEndpointAPIService service

type ApiFederationConfigurationApiRequest struct {
	ctx        context.Context
	ApiService FederationEndpointAPI
	serviceId  string
	body       *map[string]interface{}
}

func (r ApiFederationConfigurationApiRequest) Body(body map[string]interface{}) ApiFederationConfigurationApiRequest {
	r.body = &body
	return r
}

func (r ApiFederationConfigurationApiRequest) Execute() (*FederationConfigurationResponse, *http.Response, error) {
	return r.ApiService.FederationConfigurationApiExecute(r)
}

/*
FederationConfigurationApi Process Entity Configuration Request

This API gathers the federation configuration about a service.

The authorization server implementation should
retrieve the value of the <code>action</code>
response parameter from the API response and take the following steps
according to the value.

<h3><code>OK</code></h3>

When the value of the <code> action</code> response
parameter is <code>OK</code>, it means that Authlete
could prepare an entity configuration successfully.

In this case, the implementation of the entity configuration endpoint of the
authorization server should return an HTTP response to the client application
with the HTTP status code "`200 OK`" and the content type
"`application/entity-statement+jwt`". The message body (= an entity
configuration in the JWT format) of the response has been prepared by
Authlete's `/federation/configuration` API and it is available as the
<code>responseContent</code> response parameter.

The implementation of the entity configuration endpoint can construct an
HTTP response by doing like below.

<pre style="border: solid 1px black; padding: 0.5em;">
200 OK
Content-Type: application/entity-statement+jwt
(Other HTTP headers)

<i>(the value of the responseContent response parameter)</i></pre>

<h3><code>NOT_FOUND</code></h3>

When the value of the <code> action</code> response
parameter is <code>NOT_FOUND</code>, it means that
the service configuration has not enabled the feature of <a href=
"https://openid.net/specs/openid-connect-federation-1_0.html">OpenID Connect
Federation 1.0</a> and so the client application should have not access the
entity configuration endpoint.

In this case, the implementation of the entity configuration endpoint of the
authorization server should return an HTTP response to the client application
with the HTTP status code "`404 Not Found`" and the content type
"`application/json`". The message body (= error information in the JSON
format) of the response has been prepared by Authlete's
`/federation/configuration` API and it is available as the
<code>responseContent</code> response parameter.

The implementation of the entity configuration endpoint can construct an
HTTP response by doing like below.

<pre style="border: solid 1px black; padding: 0.5em;">
404 Not Found
Content-Type: application/json
(Other HTTP headers)

<i>(the value of the responseContent response parameter)</i></pre>

<h3><code>INTERNAL_SERVER_ERROR</code></h3>

could prepare an entity configuration successfully.

In this case, the implementation of the entity configuration endpoint of the
authorization server should return an HTTP response to the client application
with the HTTP status code "`200 OK`" and the content type
"`application/entity-statement+jwt`". The message body (= an entity
configuration in the JWT format) of the response has been prepared by
Authlete's `/federation/configuration` API and it is available as the
<code>responseContent</code> response parameter.

The implementation of the entity configuration endpoint can construct an
HTTP response by doing like below.

<pre style="border: solid 1px black; padding: 0.5em;">
200 OK
Content-Type: application/entity-statement+jwt
(Other HTTP headers)

<i>(the value of the responseContent response parameter)</i></pre>

</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@return ApiFederationConfigurationApiRequest
*/
func (a *FederationEndpointAPIService) FederationConfigurationApi(ctx context.Context, serviceId string) ApiFederationConfigurationApiRequest {
	return ApiFederationConfigurationApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return FederationConfigurationResponse
func (a *FederationEndpointAPIService) FederationConfigurationApiExecute(r ApiFederationConfigurationApiRequest) (*FederationConfigurationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FederationConfigurationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FederationEndpointAPIService.FederationConfigurationApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/federation/configuration"
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
	localVarPostBody = r.body
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

type ApiFederationRegistrationApiRequest struct {
	ctx                           context.Context
	ApiService                    FederationEndpointAPI
	serviceId                     string
	federationRegistrationRequest *FederationRegistrationRequest
}

func (r ApiFederationRegistrationApiRequest) FederationRegistrationRequest(federationRegistrationRequest FederationRegistrationRequest) ApiFederationRegistrationApiRequest {
	r.federationRegistrationRequest = &federationRegistrationRequest
	return r
}

func (r ApiFederationRegistrationApiRequest) Execute() (*FederationRegistrationResponse, *http.Response, error) {
	return r.ApiService.FederationRegistrationApiExecute(r)
}

/*
FederationRegistrationApi Process Federation Registration Request

The Authlete API is for implementations of the <b>federation registration
endpoint</b> that accepts "explicit client registration". Its details are
defined in <a href="https://openid.net/specs/openid-connect-federation-1_0.html"
>OpenID Connect Federation 1.0</a>.
</p>

<p>
The endpoint accepts `POST` requests whose `Content-Type`
is either of the following.
</p>

<ol>

	<li>`application/entity-statement+jwt`
	<li>`application/trust-chain+json`

</ol>

<p>
When the `Content-Type` of a request is
`application/entity-statement+jwt`, the content of the request is
the entity configuration of a relying party that is to be registered.
In this case, the implementation of the federation registration endpoint
should call Authlete's `/federation/registration` API with the
entity configuration set to the `entityConfiguration` request
parameter.
</p>

<p>
On the other hand, when the `Content-Type` of a request is
`application/trust-chain+json`, the content of the request is a
JSON array that contains entity statements in JWT format. The sequence
of the entity statements composes the trust chain of a relying party
that is to be registered. In this case, the implementation of the
federation registration endpoint should call Authlete's
`/federation/registration` API with the trust chain set to the
`trustChain` request parameter.
</p>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@return ApiFederationRegistrationApiRequest
*/
func (a *FederationEndpointAPIService) FederationRegistrationApi(ctx context.Context, serviceId string) ApiFederationRegistrationApiRequest {
	return ApiFederationRegistrationApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return FederationRegistrationResponse
func (a *FederationEndpointAPIService) FederationRegistrationApiExecute(r ApiFederationRegistrationApiRequest) (*FederationRegistrationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FederationRegistrationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FederationEndpointAPIService.FederationRegistrationApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/federation/registration"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.federationRegistrationRequest == nil {
		return localVarReturnValue, nil, reportError("federationRegistrationRequest is required and must be specified")
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
	localVarPostBody = r.federationRegistrationRequest
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
