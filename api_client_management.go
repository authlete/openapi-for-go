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

type ClientManagementAPI interface {

	/*
			ClientAuthorizationDeleteApi Delete Client Tokens

			Delete all existing access tokens issued to a client application by an end-user.

		The subject parameter is required and can be provided either in the path or as a query parameter.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@param clientId A client ID.
			@param subject Unique user ID of an end-user.
			@return ApiClientAuthorizationDeleteApiRequest
	*/
	ClientAuthorizationDeleteApi(ctx context.Context, serviceId string, clientId string, subject string) ApiClientAuthorizationDeleteApiRequest

	// ClientAuthorizationDeleteApiExecute executes the request
	//  @return ClientAuthorizationDeleteResponse
	ClientAuthorizationDeleteApiExecute(r ApiClientAuthorizationDeleteApiRequest) (*ClientAuthorizationDeleteResponse, *http.Response, error)

	/*
			ClientAuthorizationGetListApi Get Authorized Applications

			Get a list of client applications that an end-user has authorized.

		The subject parameter is required and can be provided either in the path or as a query parameter.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@param subject Unique user ID of an end-user.
			@return ApiClientAuthorizationGetListApiRequest
	*/
	ClientAuthorizationGetListApi(ctx context.Context, serviceId string, subject string) ApiClientAuthorizationGetListApiRequest

	// ClientAuthorizationGetListApiExecute executes the request
	//  @return ClientAuthorizationGetListResponse
	ClientAuthorizationGetListApiExecute(r ApiClientAuthorizationGetListApiRequest) (*ClientAuthorizationGetListResponse, *http.Response, error)

	/*
		ClientAuthorizationUpdateApi Update Client Tokens

		Update attributes of all existing access tokens given to a client application.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param serviceId A service ID.
		@param clientId A client ID.
		@return ApiClientAuthorizationUpdateApiRequest
	*/
	ClientAuthorizationUpdateApi(ctx context.Context, serviceId string, clientId string) ApiClientAuthorizationUpdateApiRequest

	// ClientAuthorizationUpdateApiExecute executes the request
	//  @return ClientAuthorizationUpdateResponse
	ClientAuthorizationUpdateApiExecute(r ApiClientAuthorizationUpdateApiRequest) (*ClientAuthorizationUpdateResponse, *http.Response, error)

	/*
		ClientCreateApi Create Client

		Create a new client.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param serviceId A service ID.
		@return ApiClientCreateApiRequest
	*/
	ClientCreateApi(ctx context.Context, serviceId string) ApiClientCreateApiRequest

	// ClientCreateApiExecute executes the request
	//  @return Client
	ClientCreateApiExecute(r ApiClientCreateApiRequest) (*Client, *http.Response, error)

	/*
		ClientDeleteApi Delete Client ⚡

		Delete a client.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param serviceId A service ID.
		@param clientId The client ID.
		@return ApiClientDeleteApiRequest
	*/
	ClientDeleteApi(ctx context.Context, serviceId string, clientId string) ApiClientDeleteApiRequest

	// ClientDeleteApiExecute executes the request
	ClientDeleteApiExecute(r ApiClientDeleteApiRequest) (*http.Response, error)

	/*
		ClientExtensionRequestablesScopesDeleteApi Delete Requestable Scopes

		Delete requestable scopes of a client


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param serviceId A service ID.
		@param clientId A client ID.
		@return ApiClientExtensionRequestablesScopesDeleteApiRequest
	*/
	ClientExtensionRequestablesScopesDeleteApi(ctx context.Context, serviceId string, clientId string) ApiClientExtensionRequestablesScopesDeleteApiRequest

	// ClientExtensionRequestablesScopesDeleteApiExecute executes the request
	ClientExtensionRequestablesScopesDeleteApiExecute(r ApiClientExtensionRequestablesScopesDeleteApiRequest) (*http.Response, error)

	/*
		ClientExtensionRequestablesScopesGetApi Get Requestable Scopes

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
		ClientExtensionRequestablesScopesUpdateApi Update Requestable Scopes

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

	/*
		ClientFlagUpdateApi Update Client Lock

		Lock and unlock a client


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param serviceId A service ID.
		@param clientIdentifier A client ID.
		@return ApiClientFlagUpdateApiRequest
	*/
	ClientFlagUpdateApi(ctx context.Context, serviceId string, clientIdentifier string) ApiClientFlagUpdateApiRequest

	// ClientFlagUpdateApiExecute executes the request
	//  @return ClientFlagUpdateResponse
	ClientFlagUpdateApiExecute(r ApiClientFlagUpdateApiRequest) (*ClientFlagUpdateResponse, *http.Response, error)

	/*
		ClientGetApi Get Client

		Get a client.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param serviceId A service ID.
		@param clientId A client ID.
		@return ApiClientGetApiRequest
	*/
	ClientGetApi(ctx context.Context, serviceId string, clientId string) ApiClientGetApiRequest

	// ClientGetApiExecute executes the request
	//  @return Client
	ClientGetApiExecute(r ApiClientGetApiRequest) (*Client, *http.Response, error)

	/*
			ClientGetListApi List Clients

			Get a list of clients on a service.

		If the access token can view a full service (including an admin), all clients within the
		service are returned. Otherwise, only clients that the access token can view within the
		service are returned.
		- ViewClient: []


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@return ApiClientGetListApiRequest
	*/
	ClientGetListApi(ctx context.Context, serviceId string) ApiClientGetListApiRequest

	// ClientGetListApiExecute executes the request
	//  @return ClientGetListResponse
	ClientGetListApiExecute(r ApiClientGetListApiRequest) (*ClientGetListResponse, *http.Response, error)

	/*
			ClientGrantedScopesDeleteApi Delete Granted Scopes

			Delete the set of scopes that an end-user has granted to a client application.

		<br>
		<details>
		<summary>Description</summary>

		Even if records about granted scopes are deleted by calling this API, existing access tokens are
		not deleted and scopes of existing access tokens are not changed.
		</details>

		The subject parameter is required and can be provided either in the path or as a query parameter.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@param clientId A client ID.
			@param subject Unique user ID of an end-user.
			@return ApiClientGrantedScopesDeleteApiRequest
	*/
	ClientGrantedScopesDeleteApi(ctx context.Context, serviceId string, clientId string, subject string) ApiClientGrantedScopesDeleteApiRequest

	// ClientGrantedScopesDeleteApiExecute executes the request
	//  @return ClientGrantedScopesDeleteResponse
	ClientGrantedScopesDeleteApiExecute(r ApiClientGrantedScopesDeleteApiRequest) (*ClientGrantedScopesDeleteResponse, *http.Response, error)

	/*
			ClientGrantedScopesGetApi Get Granted Scopes

			Get the set of scopes that a user has granted to a client application.

		<br>
		<details>
		<summary>Description</summary>

		Possible values for `requestableScopes` parameter in the response from this API are as follows.

		**null**

		The user has not granted authorization to the client application in the past, or records about the
		combination of the user and the client application have been deleted from Authlete's DB.

		**An empty set**

		The user has granted authorization to the client application in the past, but no scopes are associated
		with the authorization.

		**A set with at least one element**

		The user has granted authorization to the client application in the past and some scopes are associated
		with the authorization. These scopes are returned.
		Example: `[ "profile", "email" ]`

		The subject parameter is required and can be provided either in the path or as a query parameter.
		</details>


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@param clientId A client ID.
			@param subject Unique user ID of an end-user.
			@return ApiClientGrantedScopesGetApiRequest
	*/
	ClientGrantedScopesGetApi(ctx context.Context, serviceId string, clientId string, subject string) ApiClientGrantedScopesGetApiRequest

	// ClientGrantedScopesGetApiExecute executes the request
	//  @return ClientAuthorizationDeleteResponse
	ClientGrantedScopesGetApiExecute(r ApiClientGrantedScopesGetApiRequest) (*ClientAuthorizationDeleteResponse, *http.Response, error)

	/*
			ClientSecretRefreshApi Rotate Client Secret

			Refresh the client secret of a client. A new value of the client secret will be generated by the
		Authlete server.

		If you want to specify a new value, use `/api/client/secret/update` API.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@param clientIdentifier The client ID or the client ID alias of a client.
			@return ApiClientSecretRefreshApiRequest
	*/
	ClientSecretRefreshApi(ctx context.Context, serviceId string, clientIdentifier string) ApiClientSecretRefreshApiRequest

	// ClientSecretRefreshApiExecute executes the request
	//  @return ClientSecretRefreshResponse
	ClientSecretRefreshApiExecute(r ApiClientSecretRefreshApiRequest) (*ClientSecretRefreshResponse, *http.Response, error)

	/*
			ClientSecretUpdateApi Update Client Secret

			Update the client secret of a client.

		If you want to have the Authlete server generate a new value of the client secret, use `/api/client/secret/refresh`
		API.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param serviceId A service ID.
			@param clientIdentifier The client ID or the client ID alias of a client.
			@return ApiClientSecretUpdateApiRequest
	*/
	ClientSecretUpdateApi(ctx context.Context, serviceId string, clientIdentifier string) ApiClientSecretUpdateApiRequest

	// ClientSecretUpdateApiExecute executes the request
	//  @return ClientSecretUpdateResponse
	ClientSecretUpdateApiExecute(r ApiClientSecretUpdateApiRequest) (*ClientSecretUpdateResponse, *http.Response, error)

	/*
		ClientUpdateApi Update Client

		Update a client.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param serviceId A service ID.
		@param clientId A client ID.
		@return ApiClientUpdateApiRequest
	*/
	ClientUpdateApi(ctx context.Context, serviceId string, clientId string) ApiClientUpdateApiRequest

	// ClientUpdateApiExecute executes the request
	//  @return Client
	ClientUpdateApiExecute(r ApiClientUpdateApiRequest) (*Client, *http.Response, error)
}

// ClientManagementAPIService ClientManagementAPI service
type ClientManagementAPIService service

type ApiClientAuthorizationDeleteApiRequest struct {
	ctx        context.Context
	ApiService ClientManagementAPI
	serviceId  string
	clientId   string
	subject    string
	subject2   *string
}

// Unique user ID of an end-user.
func (r ApiClientAuthorizationDeleteApiRequest) Subject2(subject2 string) ApiClientAuthorizationDeleteApiRequest {
	r.subject2 = &subject2
	return r
}

func (r ApiClientAuthorizationDeleteApiRequest) Execute() (*ClientAuthorizationDeleteResponse, *http.Response, error) {
	return r.ApiService.ClientAuthorizationDeleteApiExecute(r)
}

/*
ClientAuthorizationDeleteApi Delete Client Tokens

Delete all existing access tokens issued to a client application by an end-user.

The subject parameter is required and can be provided either in the path or as a query parameter.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@param clientId A client ID.
	@param subject Unique user ID of an end-user.
	@return ApiClientAuthorizationDeleteApiRequest
*/
func (a *ClientManagementAPIService) ClientAuthorizationDeleteApi(ctx context.Context, serviceId string, clientId string, subject string) ApiClientAuthorizationDeleteApiRequest {
	return ApiClientAuthorizationDeleteApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		clientId:   clientId,
		subject:    subject,
	}
}

// Execute executes the request
//
//	@return ClientAuthorizationDeleteResponse
func (a *ClientManagementAPIService) ClientAuthorizationDeleteApiExecute(r ApiClientAuthorizationDeleteApiRequest) (*ClientAuthorizationDeleteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientAuthorizationDeleteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientManagementAPIService.ClientAuthorizationDeleteApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/authorization/delete/{clientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subject"+"}", url.PathEscape(parameterValueToString(r.subject, "subject")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subject2 == nil {
		return localVarReturnValue, nil, reportError("subject2 is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "subject", r.subject2, "", "")
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

type ApiClientAuthorizationGetListApiRequest struct {
	ctx        context.Context
	ApiService ClientManagementAPI
	serviceId  string
	subject    string
	subject2   *string
	developer  *string
	start      *int32
	end        *int32
}

// Unique user ID of an end-user.
func (r ApiClientAuthorizationGetListApiRequest) Subject2(subject2 string) ApiClientAuthorizationGetListApiRequest {
	r.subject2 = &subject2
	return r
}

// Unique ID of a client developer.
func (r ApiClientAuthorizationGetListApiRequest) Developer(developer string) ApiClientAuthorizationGetListApiRequest {
	r.developer = &developer
	return r
}

// Start index of search results (inclusive). The default value is 0.
func (r ApiClientAuthorizationGetListApiRequest) Start(start int32) ApiClientAuthorizationGetListApiRequest {
	r.start = &start
	return r
}

// End index of search results (exclusive). The default value is 5.
func (r ApiClientAuthorizationGetListApiRequest) End(end int32) ApiClientAuthorizationGetListApiRequest {
	r.end = &end
	return r
}

func (r ApiClientAuthorizationGetListApiRequest) Execute() (*ClientAuthorizationGetListResponse, *http.Response, error) {
	return r.ApiService.ClientAuthorizationGetListApiExecute(r)
}

/*
ClientAuthorizationGetListApi Get Authorized Applications

Get a list of client applications that an end-user has authorized.

The subject parameter is required and can be provided either in the path or as a query parameter.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@param subject Unique user ID of an end-user.
	@return ApiClientAuthorizationGetListApiRequest
*/
func (a *ClientManagementAPIService) ClientAuthorizationGetListApi(ctx context.Context, serviceId string, subject string) ApiClientAuthorizationGetListApiRequest {
	return ApiClientAuthorizationGetListApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		subject:    subject,
	}
}

// Execute executes the request
//
//	@return ClientAuthorizationGetListResponse
func (a *ClientManagementAPIService) ClientAuthorizationGetListApiExecute(r ApiClientAuthorizationGetListApiRequest) (*ClientAuthorizationGetListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientAuthorizationGetListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientManagementAPIService.ClientAuthorizationGetListApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/authorization/get/list"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subject"+"}", url.PathEscape(parameterValueToString(r.subject, "subject")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subject2 == nil {
		return localVarReturnValue, nil, reportError("subject2 is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "subject", r.subject2, "", "")
	if r.developer != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "developer", r.developer, "", "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "", "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "", "")
	}
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

type ApiClientAuthorizationUpdateApiRequest struct {
	ctx                              context.Context
	ApiService                       ClientManagementAPI
	serviceId                        string
	clientId                         string
	clientAuthorizationUpdateRequest *ClientAuthorizationUpdateRequest
}

func (r ApiClientAuthorizationUpdateApiRequest) ClientAuthorizationUpdateRequest(clientAuthorizationUpdateRequest ClientAuthorizationUpdateRequest) ApiClientAuthorizationUpdateApiRequest {
	r.clientAuthorizationUpdateRequest = &clientAuthorizationUpdateRequest
	return r
}

func (r ApiClientAuthorizationUpdateApiRequest) Execute() (*ClientAuthorizationUpdateResponse, *http.Response, error) {
	return r.ApiService.ClientAuthorizationUpdateApiExecute(r)
}

/*
ClientAuthorizationUpdateApi Update Client Tokens

Update attributes of all existing access tokens given to a client application.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@param clientId A client ID.
	@return ApiClientAuthorizationUpdateApiRequest
*/
func (a *ClientManagementAPIService) ClientAuthorizationUpdateApi(ctx context.Context, serviceId string, clientId string) ApiClientAuthorizationUpdateApiRequest {
	return ApiClientAuthorizationUpdateApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		clientId:   clientId,
	}
}

// Execute executes the request
//
//	@return ClientAuthorizationUpdateResponse
func (a *ClientManagementAPIService) ClientAuthorizationUpdateApiExecute(r ApiClientAuthorizationUpdateApiRequest) (*ClientAuthorizationUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientAuthorizationUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientManagementAPIService.ClientAuthorizationUpdateApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/authorization/update/{clientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.clientAuthorizationUpdateRequest
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

type ApiClientCreateApiRequest struct {
	ctx        context.Context
	ApiService ClientManagementAPI
	serviceId  string
	client     *Client
}

func (r ApiClientCreateApiRequest) Client(client Client) ApiClientCreateApiRequest {
	r.client = &client
	return r
}

func (r ApiClientCreateApiRequest) Execute() (*Client, *http.Response, error) {
	return r.ApiService.ClientCreateApiExecute(r)
}

/*
ClientCreateApi Create Client

Create a new client.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@return ApiClientCreateApiRequest
*/
func (a *ClientManagementAPIService) ClientCreateApi(ctx context.Context, serviceId string) ApiClientCreateApiRequest {
	return ApiClientCreateApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return Client
func (a *ClientManagementAPIService) ClientCreateApiExecute(r ApiClientCreateApiRequest) (*Client, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientManagementAPIService.ClientCreateApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/create"
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
	localVarPostBody = r.client
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

type ApiClientDeleteApiRequest struct {
	ctx        context.Context
	ApiService ClientManagementAPI
	serviceId  string
	clientId   string
}

func (r ApiClientDeleteApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.ClientDeleteApiExecute(r)
}

/*
ClientDeleteApi Delete Client ⚡

Delete a client.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@param clientId The client ID.
	@return ApiClientDeleteApiRequest
*/
func (a *ClientManagementAPIService) ClientDeleteApi(ctx context.Context, serviceId string, clientId string) ApiClientDeleteApiRequest {
	return ApiClientDeleteApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		clientId:   clientId,
	}
}

// Execute executes the request
func (a *ClientManagementAPIService) ClientDeleteApiExecute(r ApiClientDeleteApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientManagementAPIService.ClientDeleteApi")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/delete/{clientId}"
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiClientExtensionRequestablesScopesDeleteApiRequest struct {
	ctx        context.Context
	ApiService ClientManagementAPI
	serviceId  string
	clientId   string
}

func (r ApiClientExtensionRequestablesScopesDeleteApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.ClientExtensionRequestablesScopesDeleteApiExecute(r)
}

/*
ClientExtensionRequestablesScopesDeleteApi Delete Requestable Scopes

# Delete requestable scopes of a client

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@param clientId A client ID.
	@return ApiClientExtensionRequestablesScopesDeleteApiRequest
*/
func (a *ClientManagementAPIService) ClientExtensionRequestablesScopesDeleteApi(ctx context.Context, serviceId string, clientId string) ApiClientExtensionRequestablesScopesDeleteApiRequest {
	return ApiClientExtensionRequestablesScopesDeleteApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		clientId:   clientId,
	}
}

// Execute executes the request
func (a *ClientManagementAPIService) ClientExtensionRequestablesScopesDeleteApiExecute(r ApiClientExtensionRequestablesScopesDeleteApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientManagementAPIService.ClientExtensionRequestablesScopesDeleteApi")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx        context.Context
	ApiService ClientManagementAPI
	serviceId  string
	clientId   string
}

func (r ApiClientExtensionRequestablesScopesGetApiRequest) Execute() (*ClientExtensionRequestableScopesGetResponse, *http.Response, error) {
	return r.ApiService.ClientExtensionRequestablesScopesGetApiExecute(r)
}

/*
ClientExtensionRequestablesScopesGetApi Get Requestable Scopes

# Get the requestable scopes per client

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@param clientId A client ID.
	@return ApiClientExtensionRequestablesScopesGetApiRequest
*/
func (a *ClientManagementAPIService) ClientExtensionRequestablesScopesGetApi(ctx context.Context, serviceId string, clientId string) ApiClientExtensionRequestablesScopesGetApiRequest {
	return ApiClientExtensionRequestablesScopesGetApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		clientId:   clientId,
	}
}

// Execute executes the request
//
//	@return ClientExtensionRequestableScopesGetResponse
func (a *ClientManagementAPIService) ClientExtensionRequestablesScopesGetApiExecute(r ApiClientExtensionRequestablesScopesGetApiRequest) (*ClientExtensionRequestableScopesGetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientExtensionRequestableScopesGetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientManagementAPIService.ClientExtensionRequestablesScopesGetApi")
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

type ApiClientExtensionRequestablesScopesUpdateApiRequest struct {
	ctx                                           context.Context
	ApiService                                    ClientManagementAPI
	serviceId                                     string
	clientId                                      string
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
ClientExtensionRequestablesScopesUpdateApi Update Requestable Scopes

# Update requestable scopes of a client

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@param clientId A client ID.
	@return ApiClientExtensionRequestablesScopesUpdateApiRequest
*/
func (a *ClientManagementAPIService) ClientExtensionRequestablesScopesUpdateApi(ctx context.Context, serviceId string, clientId string) ApiClientExtensionRequestablesScopesUpdateApiRequest {
	return ApiClientExtensionRequestablesScopesUpdateApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		clientId:   clientId,
	}
}

// Execute executes the request
//
//	@return ClientExtensionRequestableScopesUpdateResponse
func (a *ClientManagementAPIService) ClientExtensionRequestablesScopesUpdateApiExecute(r ApiClientExtensionRequestablesScopesUpdateApiRequest) (*ClientExtensionRequestableScopesUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientExtensionRequestableScopesUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientManagementAPIService.ClientExtensionRequestablesScopesUpdateApi")
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

type ApiClientFlagUpdateApiRequest struct {
	ctx                     context.Context
	ApiService              ClientManagementAPI
	serviceId               string
	clientIdentifier        string
	clientFlagUpdateRequest *ClientFlagUpdateRequest
}

func (r ApiClientFlagUpdateApiRequest) ClientFlagUpdateRequest(clientFlagUpdateRequest ClientFlagUpdateRequest) ApiClientFlagUpdateApiRequest {
	r.clientFlagUpdateRequest = &clientFlagUpdateRequest
	return r
}

func (r ApiClientFlagUpdateApiRequest) Execute() (*ClientFlagUpdateResponse, *http.Response, error) {
	return r.ApiService.ClientFlagUpdateApiExecute(r)
}

/*
ClientFlagUpdateApi Update Client Lock

# Lock and unlock a client

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@param clientIdentifier A client ID.
	@return ApiClientFlagUpdateApiRequest
*/
func (a *ClientManagementAPIService) ClientFlagUpdateApi(ctx context.Context, serviceId string, clientIdentifier string) ApiClientFlagUpdateApiRequest {
	return ApiClientFlagUpdateApiRequest{
		ApiService:       a,
		ctx:              ctx,
		serviceId:        serviceId,
		clientIdentifier: clientIdentifier,
	}
}

// Execute executes the request
//
//	@return ClientFlagUpdateResponse
func (a *ClientManagementAPIService) ClientFlagUpdateApiExecute(r ApiClientFlagUpdateApiRequest) (*ClientFlagUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientFlagUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientManagementAPIService.ClientFlagUpdateApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/lock_flag/update/{clientIdentifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientIdentifier"+"}", url.PathEscape(parameterValueToString(r.clientIdentifier, "clientIdentifier")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.clientFlagUpdateRequest
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

type ApiClientGetApiRequest struct {
	ctx        context.Context
	ApiService ClientManagementAPI
	serviceId  string
	clientId   string
}

func (r ApiClientGetApiRequest) Execute() (*Client, *http.Response, error) {
	return r.ApiService.ClientGetApiExecute(r)
}

/*
ClientGetApi Get Client

Get a client.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@param clientId A client ID.
	@return ApiClientGetApiRequest
*/
func (a *ClientManagementAPIService) ClientGetApi(ctx context.Context, serviceId string, clientId string) ApiClientGetApiRequest {
	return ApiClientGetApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		clientId:   clientId,
	}
}

// Execute executes the request
//
//	@return Client
func (a *ClientManagementAPIService) ClientGetApiExecute(r ApiClientGetApiRequest) (*Client, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientManagementAPIService.ClientGetApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/get/{clientId}"
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

type ApiClientGetListApiRequest struct {
	ctx        context.Context
	ApiService ClientManagementAPI
	serviceId  string
	developer  *string
	start      *int32
	end        *int32
}

// The developer of client applications. The default value is null. If this parameter is not set to &#x60;null&#x60;, client application of the specified developer are returned. Otherwise, all client applications that belong to the service are returned.
func (r ApiClientGetListApiRequest) Developer(developer string) ApiClientGetListApiRequest {
	r.developer = &developer
	return r
}

// Start index (inclusive) of the result set. The default value is 0. Must not be a negative number.
func (r ApiClientGetListApiRequest) Start(start int32) ApiClientGetListApiRequest {
	r.start = &start
	return r
}

// End index (exclusive) of the result set. The default value is 5. Must not be a negative number.
func (r ApiClientGetListApiRequest) End(end int32) ApiClientGetListApiRequest {
	r.end = &end
	return r
}

func (r ApiClientGetListApiRequest) Execute() (*ClientGetListResponse, *http.Response, error) {
	return r.ApiService.ClientGetListApiExecute(r)
}

/*
ClientGetListApi List Clients

Get a list of clients on a service.

If the access token can view a full service (including an admin), all clients within the
service are returned. Otherwise, only clients that the access token can view within the
service are returned.
- ViewClient: []

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@return ApiClientGetListApiRequest
*/
func (a *ClientManagementAPIService) ClientGetListApi(ctx context.Context, serviceId string) ApiClientGetListApiRequest {
	return ApiClientGetListApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return ClientGetListResponse
func (a *ClientManagementAPIService) ClientGetListApiExecute(r ApiClientGetListApiRequest) (*ClientGetListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientGetListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientManagementAPIService.ClientGetListApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/get/list"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.developer != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "developer", r.developer, "", "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "", "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "", "")
	}
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

type ApiClientGrantedScopesDeleteApiRequest struct {
	ctx        context.Context
	ApiService ClientManagementAPI
	serviceId  string
	clientId   string
	subject    string
	subject2   *string
}

// Unique user ID of an end-user.
func (r ApiClientGrantedScopesDeleteApiRequest) Subject2(subject2 string) ApiClientGrantedScopesDeleteApiRequest {
	r.subject2 = &subject2
	return r
}

func (r ApiClientGrantedScopesDeleteApiRequest) Execute() (*ClientGrantedScopesDeleteResponse, *http.Response, error) {
	return r.ApiService.ClientGrantedScopesDeleteApiExecute(r)
}

/*
ClientGrantedScopesDeleteApi Delete Granted Scopes

Delete the set of scopes that an end-user has granted to a client application.

<br>
<details>
<summary>Description</summary>

Even if records about granted scopes are deleted by calling this API, existing access tokens are
not deleted and scopes of existing access tokens are not changed.
</details>

The subject parameter is required and can be provided either in the path or as a query parameter.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@param clientId A client ID.
	@param subject Unique user ID of an end-user.
	@return ApiClientGrantedScopesDeleteApiRequest
*/
func (a *ClientManagementAPIService) ClientGrantedScopesDeleteApi(ctx context.Context, serviceId string, clientId string, subject string) ApiClientGrantedScopesDeleteApiRequest {
	return ApiClientGrantedScopesDeleteApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		clientId:   clientId,
		subject:    subject,
	}
}

// Execute executes the request
//
//	@return ClientGrantedScopesDeleteResponse
func (a *ClientManagementAPIService) ClientGrantedScopesDeleteApiExecute(r ApiClientGrantedScopesDeleteApiRequest) (*ClientGrantedScopesDeleteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientGrantedScopesDeleteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientManagementAPIService.ClientGrantedScopesDeleteApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/granted_scopes/delete/{clientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subject"+"}", url.PathEscape(parameterValueToString(r.subject, "subject")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subject2 == nil {
		return localVarReturnValue, nil, reportError("subject2 is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "subject", r.subject2, "", "")
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

type ApiClientGrantedScopesGetApiRequest struct {
	ctx        context.Context
	ApiService ClientManagementAPI
	serviceId  string
	clientId   string
	subject    string
	subject2   *string
}

// Unique user ID of an end-user.
func (r ApiClientGrantedScopesGetApiRequest) Subject2(subject2 string) ApiClientGrantedScopesGetApiRequest {
	r.subject2 = &subject2
	return r
}

func (r ApiClientGrantedScopesGetApiRequest) Execute() (*ClientAuthorizationDeleteResponse, *http.Response, error) {
	return r.ApiService.ClientGrantedScopesGetApiExecute(r)
}

/*
ClientGrantedScopesGetApi Get Granted Scopes

Get the set of scopes that a user has granted to a client application.

<br>
<details>
<summary>Description</summary>

Possible values for `requestableScopes` parameter in the response from this API are as follows.

**null**

The user has not granted authorization to the client application in the past, or records about the
combination of the user and the client application have been deleted from Authlete's DB.

**An empty set**

The user has granted authorization to the client application in the past, but no scopes are associated
with the authorization.

**A set with at least one element**

The user has granted authorization to the client application in the past and some scopes are associated
with the authorization. These scopes are returned.
Example: `[ "profile", "email" ]`

The subject parameter is required and can be provided either in the path or as a query parameter.
</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@param clientId A client ID.
	@param subject Unique user ID of an end-user.
	@return ApiClientGrantedScopesGetApiRequest
*/
func (a *ClientManagementAPIService) ClientGrantedScopesGetApi(ctx context.Context, serviceId string, clientId string, subject string) ApiClientGrantedScopesGetApiRequest {
	return ApiClientGrantedScopesGetApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		clientId:   clientId,
		subject:    subject,
	}
}

// Execute executes the request
//
//	@return ClientAuthorizationDeleteResponse
func (a *ClientManagementAPIService) ClientGrantedScopesGetApiExecute(r ApiClientGrantedScopesGetApiRequest) (*ClientAuthorizationDeleteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientAuthorizationDeleteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientManagementAPIService.ClientGrantedScopesGetApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/granted_scopes/get/{clientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subject"+"}", url.PathEscape(parameterValueToString(r.subject, "subject")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subject2 == nil {
		return localVarReturnValue, nil, reportError("subject2 is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "subject", r.subject2, "", "")
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

type ApiClientSecretRefreshApiRequest struct {
	ctx              context.Context
	ApiService       ClientManagementAPI
	serviceId        string
	clientIdentifier string
}

func (r ApiClientSecretRefreshApiRequest) Execute() (*ClientSecretRefreshResponse, *http.Response, error) {
	return r.ApiService.ClientSecretRefreshApiExecute(r)
}

/*
ClientSecretRefreshApi Rotate Client Secret

Refresh the client secret of a client. A new value of the client secret will be generated by the
Authlete server.

If you want to specify a new value, use `/api/client/secret/update` API.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@param clientIdentifier The client ID or the client ID alias of a client.
	@return ApiClientSecretRefreshApiRequest
*/
func (a *ClientManagementAPIService) ClientSecretRefreshApi(ctx context.Context, serviceId string, clientIdentifier string) ApiClientSecretRefreshApiRequest {
	return ApiClientSecretRefreshApiRequest{
		ApiService:       a,
		ctx:              ctx,
		serviceId:        serviceId,
		clientIdentifier: clientIdentifier,
	}
}

// Execute executes the request
//
//	@return ClientSecretRefreshResponse
func (a *ClientManagementAPIService) ClientSecretRefreshApiExecute(r ApiClientSecretRefreshApiRequest) (*ClientSecretRefreshResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientSecretRefreshResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientManagementAPIService.ClientSecretRefreshApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/secret/refresh/{clientIdentifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientIdentifier"+"}", url.PathEscape(parameterValueToString(r.clientIdentifier, "clientIdentifier")), -1)

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

type ApiClientSecretUpdateApiRequest struct {
	ctx                       context.Context
	ApiService                ClientManagementAPI
	serviceId                 string
	clientIdentifier          string
	clientSecretUpdateRequest *ClientSecretUpdateRequest
}

func (r ApiClientSecretUpdateApiRequest) ClientSecretUpdateRequest(clientSecretUpdateRequest ClientSecretUpdateRequest) ApiClientSecretUpdateApiRequest {
	r.clientSecretUpdateRequest = &clientSecretUpdateRequest
	return r
}

func (r ApiClientSecretUpdateApiRequest) Execute() (*ClientSecretUpdateResponse, *http.Response, error) {
	return r.ApiService.ClientSecretUpdateApiExecute(r)
}

/*
ClientSecretUpdateApi Update Client Secret

Update the client secret of a client.

If you want to have the Authlete server generate a new value of the client secret, use `/api/client/secret/refresh`
API.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@param clientIdentifier The client ID or the client ID alias of a client.
	@return ApiClientSecretUpdateApiRequest
*/
func (a *ClientManagementAPIService) ClientSecretUpdateApi(ctx context.Context, serviceId string, clientIdentifier string) ApiClientSecretUpdateApiRequest {
	return ApiClientSecretUpdateApiRequest{
		ApiService:       a,
		ctx:              ctx,
		serviceId:        serviceId,
		clientIdentifier: clientIdentifier,
	}
}

// Execute executes the request
//
//	@return ClientSecretUpdateResponse
func (a *ClientManagementAPIService) ClientSecretUpdateApiExecute(r ApiClientSecretUpdateApiRequest) (*ClientSecretUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientSecretUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientManagementAPIService.ClientSecretUpdateApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/secret/update/{clientIdentifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientIdentifier"+"}", url.PathEscape(parameterValueToString(r.clientIdentifier, "clientIdentifier")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientSecretUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("clientSecretUpdateRequest is required and must be specified")
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
	localVarPostBody = r.clientSecretUpdateRequest
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

type ApiClientUpdateApiRequest struct {
	ctx        context.Context
	ApiService ClientManagementAPI
	serviceId  string
	clientId   string
	client     *Client
}

func (r ApiClientUpdateApiRequest) Client(client Client) ApiClientUpdateApiRequest {
	r.client = &client
	return r
}

func (r ApiClientUpdateApiRequest) Execute() (*Client, *http.Response, error) {
	return r.ApiService.ClientUpdateApiExecute(r)
}

/*
ClientUpdateApi Update Client

Update a client.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId A service ID.
	@param clientId A client ID.
	@return ApiClientUpdateApiRequest
*/
func (a *ClientManagementAPIService) ClientUpdateApi(ctx context.Context, serviceId string, clientId string) ApiClientUpdateApiRequest {
	return ApiClientUpdateApiRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
		clientId:   clientId,
	}
}

// Execute executes the request
//
//	@return Client
func (a *ClientManagementAPIService) ClientUpdateApiExecute(r ApiClientUpdateApiRequest) (*Client, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientManagementAPIService.ClientUpdateApi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/{serviceId}/client/update/{clientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.client
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
