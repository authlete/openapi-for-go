/*
Authlete API Explorer

<div class=\"min-h-screen bg-gray-100 dark:bg-gray-900 text-gray-900 dark:text-gray-100 p-6\">   <div class=\"flex justify-end mb-4\">     <label for=\"theme-toggle\" class=\"flex items-center cursor-pointer\">       <div class=\"relative\">Dark mode:         <input type=\"checkbox\" id=\"theme-toggle\" class=\"sr-only\" onchange=\"toggleTheme()\">         <div class=\"block bg-gray-600 w-14 h-8 rounded-full\"></div>         <div class=\"dot absolute left-1 top-1 bg-white w-6 h-6 rounded-full transition\"></div>       </div>     </label>   </div>   <header class=\"bg-green-500 dark:bg-green-700 p-4 rounded-lg text-white text-center\">     <p>       Welcome to the <strong>Authlete API documentation</strong>. Authlete is an <strong>API-first service</strong>       where every aspect of the platform is configurable via API. This explorer provides a convenient way to       authenticate and interact with the API, allowing you to see Authlete in action quickly. 🚀     </p>     <p>       At a high level, the Authlete API is grouped into two categories:     </p>     <ul class=\"list-disc list-inside\">       <li><strong>Management APIs</strong>: Enable you to manage services and clients. 🔧</li>       <li><strong>Runtime APIs</strong>: Allow you to build your own Authorization Servers or Verifiable Credential (VC)         issuers. 🔐</li>     </ul>     <p>All API endpoints are secured using access tokens issued by Authlete's Identity Provider (IdP). If you already       have an Authlete account, simply use the <em>Get Token</em> option on the Authentication page to log in and obtain       an access token for API usage. If you don't have an account yet, <a href=\"https://console.authlete.com/register\">sign up         here</a> to get started.</p>   </header>   <main>     <section id=\"api-servers\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🌐 API Servers</h2>       <p>Authlete is a global service with clusters available in multiple regions across the world.</p>       <p>Currently, our service is available in the following regions:</p>       <div class=\"grid grid-cols-2 gap-4\">         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇺🇸 US</p>         </div>         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇯🇵 JP</p>         </div>         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇪🇺 EU</p>         </div>         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇧🇷 Brazil</p>         </div>       </div>       <p>Our customers can host their data in the region that best meets their requirements.</p>       <a href=\"#servers\" class=\"block mt-4 text-green-500 dark:text-green-300 hover:underline text-center\">Select your         preferred server</a>     </section>     <section id=\"authentication\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🔑 Authentication</h2>       <p>The API Explorer requires an access token to call the API.</p>       <p>You can create the access token from the <a href=\"https://console.authlete.com\">Authlete Management Console</a> and set it in the HTTP Bearer section of Authentication page.</p>       <p>Alternatively, if you have an Authlete account, the API Explorer can log you in with your Authlete account and         automatically acquire the required access token.</p>       <div class=\"theme-admonition theme-admonition-warning admonition_o5H7 alert alert--warning\">         <div class=\"admonitionContent_Knsx\">           <p>⚠️ <strong>Important Note:</strong> When the API Explorer acquires the token after login, the access tokens             will have the same permissions as the user who logs in as part of this flow.</p>         </div>       </div>       <a href=\"#auth\" class=\"block mt-4 text-green-500 dark:text-green-300 hover:underline text-center\">Setup your         access token</a>     </section>     <section id=\"tutorials\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🎓 Tutorials</h2>       <p>If you have successfully tested the API from the API Console and want to take the next step of integrating the         API into your application, or if you want to see a sample using Authlete APIs, follow the links below. These         resources will help you understand key concepts and how to integrate Authlete API into your applications.</p>       <div class=\"mt-4\">         <a href=\"https://www.authlete.com/developers/getting_started/\"           class=\"block text-green-500 dark:text-green-300 font-bold hover:underline mb-2\">🚀 Getting Started with           Authlete</a>           </br>         <a href=\"https://www.authlete.com/developers/tutorial/signup/\"           class=\"block text-green-500 dark:text-green-300 font-bold hover:underline\">🔑 From Sign-Up to the First API           Request</a>       </div>     </section>     <section id=\"support\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🛠 Contact Us</h2>       <p>If you have any questions or need assistance, our team is here to help.</p>       <a href=\"https://www.authlete.com/contact/\"         class=\"block mt-4 text-green-500 dark:text-green-300 font-bold hover:underline\">Contact Page</a>     </section>   </main> </div>

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
)

// checks if the AuthorizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationResponse{}

// AuthorizationResponse struct for AuthorizationResponse
type AuthorizationResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take.
	Action  *string                     `json:"action,omitempty"`
	Client  *ClientLimitedAuthorization `json:"client,omitempty"`
	Display *Display                    `json:"display,omitempty"`
	// The maximum authentication age. This value comes from `max_age` request parameter, or `defaultMaxAge` configuration parameter of the client application when the authorization request does not contain `max_age` request parameter.  See \"[OpenID Connect Core 1.0, 3.1.2.1. Authentication Request](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest), max_age\" for `max_age` request parameter, and see \"[OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata), default_max_age\" for `defaultMaxAge` configuration parameter.
	MaxAge  *int32   `json:"maxAge,omitempty"`
	Service *Service `json:"service,omitempty"`
	// The scopes that the client application requests. This value comes from `scope` request parameter. If the request does not contain `scope` parameter, this parameter is a list of scopes which are registered as default. If the authorization request does not have `scope` request parameter and the service has not registered any default scope, the value of this parameter is `null`.  It is ensured that scopes listed by this parameters are contained in the list of supported scopes which are specified by `supportedScopes` configuration parameter of the service. Unsupported scopes in the authorization request do not cause an error and are just ignored.  OpenID Connect defines some scope names which need to be treated specially. The table below lists the special scope names.  | Name | Description | | --- | --- | | `openid` | This scope must be contained in `scope` request parameter to promote an OAuth 2.0 authorization request to an OpenID Connect request. It is described in \"[OpenID Connect Core 1.0, 3.1.2.1. Authentication Request](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest), scope\". | | `profile` | This scope is used to request some claims to be embedded in the ID token. The claims are `name`, `family_name`, `given_name`, `middle_name`, `nickname`, `preferred_username`, `profile`, `picture`, `website`, `gender`, `birthdate`, `zoneinfo`, `locale`, and `updated_at`. It is described in [OpenID Connect Core 1.0, 5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims). | | `email` | This scope is used to request some claims to be embedded in the ID token. The claims are `email` and `email_verified`. It is described in [OpenID Connect Core 1.0, 5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims). | | `address` |  This scope is used to request `address` claim to be embedded in the ID token. It is described in [OpenID Connect Core 1.0, 5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims).<br><br> The format of `address` claim is not a simple string. It is described in [OpenID Connect Core 1.0, 5.1.1. Address Claim](https://openid.net/specs/openid-connect-core-1_0.html#AddressClaim).  | | `phone` | This scope is used to request some claims to be embedded in the ID token. The claims are `phone_number` and `phone_number_verified`. It is described in [OpenID Connect Core 1.0, 5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims).  | | `offline_access` | The following is an excerpt about this scope from [OpenID Connect Core 1.0, 11. Offline Access](https://openid.net/specs/openid-connect-core-1_0.html#OfflineAccess). <blockquote>This scope value requests that an OAuth 2.0 Refresh Token be issued that can be used to obtain an Access Token that grants access to the end-user's userinfo endpoint even when the end-user is not present (not logged in).</blockquote>  |  Note that, if `response_type` request parameter does not contain code, `offline_acccess` scope is removed from this list even when scope request parameter contains `offline_access`. This behavior is a requirement written in [OpenID Connect Core 1.0, 11. Offline Access](https://openid.net/specs/openid-connect-core-1_0.html#OfflineAccess).
	Scopes []Scope `json:"scopes,omitempty"`
	// The locales that the client application presented as candidates to be used for UI. This value comes from `ui_locales` request parameter. The format of `ui_locales` is a space-separated list of language tag values defined in [RFC5646](https://datatracker.ietf.org/doc/html/rfc5646). See \"[OpenID Connect Core 1.0, 3.1.2.1. Authentication Request](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest), ui_locales\" for details.  It is ensured that locales listed by this parameters are contained in the list of supported UI locales which are specified by `supportedUiLocales` configuration parameter of the service. Unsupported UI locales in the authorization request do not cause an error and are just ignored.
	UiLocales []string `json:"uiLocales,omitempty"`
	// End-user's preferred languages and scripts for claims. This value comes from `claims_locales` request parameter. The format of `claims_locales` is a space-separated list of language tag values defined in [RFC5646](https://datatracker.ietf.org/doc/html/rfc5646). See \"[OpenID Connect Core 1.0, 5.2. Claims Languages and Scripts](https://openid.net/specs/openid-connect-core-1_0.html#ClaimsLanguagesAndScripts)\" for details.  It is ensured that locales listed by this parameters are contained in the list of supported claim locales which are specified by `supportedClaimsLocales` configuration parameter of the service. Unsupported claim locales in the authorization request do not cause an error and are just ignored.
	ClaimsLocales []string `json:"claimsLocales,omitempty"`
	// The list of claims that the client application requests to be embedded in the ID token. The value comes from (1) `id_token` in `claims` request parameter [1] and/or (2) special scopes (`profile`, `email`, `address` and `phone`) which are expanded to claims.  See [OpenID Connect Core 1.0, 5.5. Requesting Claims using the \"claims\" Request Parameter](https://openid.net/specs/openid-connect-core-1_0.html#ClaimsParameter) for `claims` request parameter, and see [OpenID Connect Core 1.0, 5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims) for the special scopes.
	Claims []string `json:"claims,omitempty"`
	// This boolean value indicates whether the authentication of the end-user must be one of the ACRs (Authentication Context Class References) listed in `acrs` parameter. This parameter becomes `true` only when (1) the authorization request contains `claims` request parameter and (2) `acr` claim is in it, and (3) `essential` property of the `acr` claim is `true`. See [OpenID Connect Core 1.0, 5.5.1.1. Requesting the \"acr\" Claim](https://openid.net/specs/openid-connect-core-1_0.html#acrSemantics) for details.
	AcrEssential *bool `json:"acrEssential,omitempty"`
	// `true` if the value of the `client_id` request parameter included in the authorization request is the client ID alias. `false` if the value is the original numeric client ID.
	ClientIdAliasUsed *bool `json:"clientIdAliasUsed,omitempty"`
	// The list of ACRs (Authentication Context Class References) one of which the client application requests to be satisfied for the authentication of the end-user. This value comes from `acr_values` request parameter or `defaultAcrs` configuration parameter of the client application.  See \"[OpenID Connect Core 1.0, 3.1.2.1. Authentication Request](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest), acr_values\" for `acr_values` request parameter, and see \"[OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata), default_acr_values\" for `defaultAcrs` configuration parameter.
	Acrs []string `json:"acrs,omitempty"`
	// The subject (= unique user ID managed by the authorization server implementation) that the client application expects to grant authorization. The value comes from `sub` claim in `claims` request parameter.
	Subject *string `json:"subject,omitempty"`
	// A hint about the login identifier of the end-user. The value comes from `login_hint` request parameter.
	LoginHint *string `json:"loginHint,omitempty"`
	// The list of values of prompt request parameter. See \"[OpenID Connect Core 1.0, 3.1.2.1. Authentication Request](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest), prompt\" for prompt request parameter.
	Prompts      []Prompt `json:"prompts,omitempty"`
	LowestPrompt *Prompt  `json:"lowestPrompt,omitempty"`
	// The payload part of the request object. The value of this proprty is `null` if the authorization request does not include a request object.
	RequestObjectPayload *string `json:"requestObjectPayload,omitempty"`
	// The value of the `id_token` property in the claims request parameter or in the claims property in a request object.  A client application may request certain claims be embedded in an ID token or in a response from the userInfo endpoint. There are several ways. Including the `claims` request parameter and including the `claims` property in a request object are such examples. In both the cases, the value of the `claims` parameter/property is JSON. Its format is described in [5.5. Requesting Claims using the \"claims\" Request Parameter](https://openid.net/specs/openid-connect-core-1_0.html#ClaimsParameter).  The following is an excerpt from the specification. You can find `userinfo` and `id_token` are top-level properties.  ```json {   \"userinfo\":   {     \"given_name\": { \"essential\": true },     \"nickname\": null,     \"email\": { \"essential\": true },     \"email_verified\": { \"essential\": true },     \"picture\": null,     \"http://example.info/claims/groups\": null   },   \"id_token\":   {     \"auth_time\": { \"essential\": true },     \"acr\": { \"values\": [ \"urn:mace:incommon:iap:silver\" ] }   } } ```  This value of this property is the value of the `id_token` property in JSON format. For example, if the JSON above is included in an authorization request, this property holds JSON equivalent to the following.  ```json {   \"auth_time\": { \"essential\": true },   \"acr\": { \"values\": [ \"urn:mace:incommon:iap:silver\" ] } } ```  Note that if a request object is given and it contains the `claims` property and if the `claims` request parameter is also given, this property holds the former value.
	IdTokenClaims *string `json:"idTokenClaims,omitempty"`
	// The value of the `userinfo` property in the `claims` request parameter or in the `claims` property in a request object.  A client application may request certain claims be embedded in an ID token or in a response from the userInfo endpoint. There are several ways. Including the `claims` request parameter and including the `claims` property in a request object are such examples. In both the cases, the value of the `claims` parameter/property is JSON. Its format is described in [5.5. Requesting Claims using the \"claims\" Request Parameter](https://openid.net/specs/openid-connect-core-1_0.html#ClaimsParameter).  The following is an excerpt from the specification. You can find `userinfo` and `id_token` are top-level properties.  ```json {   \"userinfo\":   {     \"given_name\": { \"essential\": true },     \"nickname\": null,     \"email\": { \"essential\": true },     \"email_verified\": { \"essential\": true },     \"picture\": null,     \"http://example.info/claims/groups\": null   },   \"id_token\":   {     \"auth_time\": { \"essential\": true },     \"acr\": { \"values\": [ \"urn:mace:incommon:iap:silver\" ] }   } } ````  The value of this property is the value of the `userinfo` property in JSON format. For example, if the JSON above is included in an authorization request, this property holds JSON equivalent to the following.  ```json {   \"given_name\": { \"essential\": true },   \"nickname\": null,   \"email\": { \"essential\": true },   \"email_verified\": { \"essential\": true },   \"picture\": null,   \"http://example.info/claims/groups\": null } ```  Note that if a request object is given and it contains the `claims` property and if the `claims` request parameter is also given, the value of this property holds the former value.
	UserInfoClaims *string `json:"userInfoClaims,omitempty"`
	// The resources specified by the `resource` request parameters or by the `resource` property in the request object. If both are given, the values in the request object should be set. See \"Resource Indicators for OAuth 2.0\" for details.
	Resources            []string      `json:"resources,omitempty"`
	AuthorizationDetails *AuthzDetails `json:"authorizationDetails,omitempty"`
	// The `purpose` request parameter is defined in [9. Transaction-specific Purpose](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html#name-transaction-specific-purpos) of [OpenID Connect for Identity Assurance 1.0](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html) as follows:  > purpose: OPTIONAL. String describing the purpose for obtaining certain user data from the OP. The purpose MUST NOT be shorter than 3 characters and MUST NOT be longer than 300 characters. If these rules are violated, the authentication request MUST fail and the OP returns an error invalid_request to the RP.
	Purpose *string `json:"purpose,omitempty"`
	// The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of `action` parameter.
	ResponseContent *string `json:"responseContent,omitempty"`
	// A ticket issued by Authlete to the service implementation. This is needed when the service implementation calls either `/auth/authorization/fail` API or `/auth/authorization/issue` API.
	Ticket *string `json:"ticket,omitempty"`
	// The dynamic scopes which the client application requested by the scope request parameter.
	DynamicScopes []DynamicScope         `json:"dynamicScopes,omitempty"`
	GmAction      *GrantManagementAction `json:"gmAction,omitempty"`
	// the value of the `grant_id` request parameter of the device authorization request.  The `grant_id` request parameter is defined in [Grant Management for OAuth 2.0](https://openid.net/specs/fapi-grant-management.html) , which is supported by Authlete 2.3 and newer versions.
	GrantId *string `json:"grantId,omitempty"`
	Grant   *Grant  `json:"grant,omitempty"`
	// The subject identifying the user who has given the grant identified by the `grant_id` request parameter of the device authorization request.  Authlete 2.3 and newer versions support <a href= \"https://openid.net/specs/fapi-grant-management.html\">Grant Management for OAuth 2.0</a>. An authorization request may contain a `grant_id` request parameter which is defined in the specification. If the value of the request parameter is valid, {@link #getGrantSubject()} will return the subject of the user who has given the grant to the client application. Authorization server implementations may use the value returned from {@link #getGrantSubject()} in order to determine the user to authenticate.  The user your system will authenticate during the authorization process (or has already authenticated) may be different from the user of the grant. The first implementer's draft of \"Grant Management for OAuth 2.0\" does not mention anything about the case, so the behavior in the case is left to implementations. Authlete will not perform the grant management action when the `subject` passed to Authlete does not match the user of the grant.
	GrantSubject *string `json:"grantSubject,omitempty"`
	// Get names of claims that are requested indirectly by <i>\"transformed claims\"</i>.  <p> A client application can request <i>\"transformed claims\"</i> by adding names of transformed claims in the `claims` request parameter. The following is an example of the `claims` request parameter that requests a predefined transformed claim named `18_or_over` and a transformed claim named `nationality_usa` to be embedded in the response from the userinfo endpoint. </p>  ```json {   \"transformed_claims\": {     \"nationality_usa\": {       \"claim\": \"nationalities\",       \"fn\": [         [ \"eq\", \"USA\" ],         \"any\"       ]     }   },   \"userinfo\": {     \"::18_or_over\": null,     \":nationality_usa\": null   } } ```  The example above assumes that a transformed claim named `18_or_over` is predefined by the authorization server like below.  ```json {   \"18_or_over\": {     \"claim\": \"birthdate\",     \"fn\": [       \"years_ago\",       [ \"gte\", 18 ]     ]   } } ```  In the example, the `nationalities` claim is requested indirectly by the `nationality_usa` transformed claim. Likewise, the `birthdate` claim is requested indirectly by the `18_or_over` transformed claim.  When the `claims` request parameter of an authorization request is like the example above, this `requestedClaimsForTx` property will hold the following value.  ```json [ \"birthdate\", \"nationalities\" ] ```  It is expected that the authorization server implementation prepares values of the listed claims and passes them as the value of the `claimsForTx` request parameter when it calls the `/api/auth/userinfo/issue` API. The following is an example of the value of the `claimsForTx` request parameter.  ```json {   \"birthdate\": \"1970-01-23\",   \"nationalities\": [ \"DEU\", \"USA\" ] } ```
	RequestedClaimsForTx []string `json:"requestedClaimsForTx,omitempty"`
	// Names of verified claims that will be referenced when transformed claims are computed.
	RequestedVerifiedClaimsForTx [][]string `json:"requestedVerifiedClaimsForTx,omitempty"`
	// the value of the `transformed_claims` property in the `claims` request parameter of an authorization request or in the `claims` property in a request object.
	TransformedClaims *string `json:"transformedClaims,omitempty"`
	// Flag which indicates whether the entity ID of the client was used when the request for the access token was made.
	ClientEntityIdUsed *bool `json:"clientEntityIdUsed,omitempty"`
	// Get the list of claims that the client application requests to be embedded in userinfo responses. The value comes from the `\"scope\"` and `\"claims\"` request parameters of the original authorization request.
	ClaimsAtUserInfo    []string             `json:"claimsAtUserInfo,omitempty"`
	CredentialOfferInfo *CredentialOfferInfo `json:"credentialOfferInfo,omitempty"`
	// Get the information about the <b>issuable credentials</b> that can be obtained by presenting the access token that will be issued as a result of the authorization request.
	IssuableCredentials *string `json:"issuableCredentials,omitempty"`
}

// NewAuthorizationResponse instantiates a new AuthorizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationResponse() *AuthorizationResponse {
	this := AuthorizationResponse{}
	return &this
}

// NewAuthorizationResponseWithDefaults instantiates a new AuthorizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationResponseWithDefaults() *AuthorizationResponse {
	this := AuthorizationResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetResultCode() string {
	if o == nil || IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *AuthorizationResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetResultMessage() string {
	if o == nil || IsNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasResultMessage() bool {
	if o != nil && !IsNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *AuthorizationResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *AuthorizationResponse) SetAction(v string) {
	o.Action = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetClient() ClientLimitedAuthorization {
	if o == nil || IsNil(o.Client) {
		var ret ClientLimitedAuthorization
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetClientOk() (*ClientLimitedAuthorization, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given ClientLimitedAuthorization and assigns it to the Client field.
func (o *AuthorizationResponse) SetClient(v ClientLimitedAuthorization) {
	o.Client = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetDisplay() Display {
	if o == nil || IsNil(o.Display) {
		var ret Display
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetDisplayOk() (*Display, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given Display and assigns it to the Display field.
func (o *AuthorizationResponse) SetDisplay(v Display) {
	o.Display = &v
}

// GetMaxAge returns the MaxAge field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetMaxAge() int32 {
	if o == nil || IsNil(o.MaxAge) {
		var ret int32
		return ret
	}
	return *o.MaxAge
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetMaxAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAge) {
		return nil, false
	}
	return o.MaxAge, true
}

// HasMaxAge returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasMaxAge() bool {
	if o != nil && !IsNil(o.MaxAge) {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given int32 and assigns it to the MaxAge field.
func (o *AuthorizationResponse) SetMaxAge(v int32) {
	o.MaxAge = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetService() Service {
	if o == nil || IsNil(o.Service) {
		var ret Service
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetServiceOk() (*Service, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given Service and assigns it to the Service field.
func (o *AuthorizationResponse) SetService(v Service) {
	o.Service = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetScopes() []Scope {
	if o == nil || IsNil(o.Scopes) {
		var ret []Scope
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetScopesOk() ([]Scope, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []Scope and assigns it to the Scopes field.
func (o *AuthorizationResponse) SetScopes(v []Scope) {
	o.Scopes = v
}

// GetUiLocales returns the UiLocales field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetUiLocales() []string {
	if o == nil || IsNil(o.UiLocales) {
		var ret []string
		return ret
	}
	return o.UiLocales
}

// GetUiLocalesOk returns a tuple with the UiLocales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetUiLocalesOk() ([]string, bool) {
	if o == nil || IsNil(o.UiLocales) {
		return nil, false
	}
	return o.UiLocales, true
}

// HasUiLocales returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasUiLocales() bool {
	if o != nil && !IsNil(o.UiLocales) {
		return true
	}

	return false
}

// SetUiLocales gets a reference to the given []string and assigns it to the UiLocales field.
func (o *AuthorizationResponse) SetUiLocales(v []string) {
	o.UiLocales = v
}

// GetClaimsLocales returns the ClaimsLocales field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetClaimsLocales() []string {
	if o == nil || IsNil(o.ClaimsLocales) {
		var ret []string
		return ret
	}
	return o.ClaimsLocales
}

// GetClaimsLocalesOk returns a tuple with the ClaimsLocales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetClaimsLocalesOk() ([]string, bool) {
	if o == nil || IsNil(o.ClaimsLocales) {
		return nil, false
	}
	return o.ClaimsLocales, true
}

// HasClaimsLocales returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasClaimsLocales() bool {
	if o != nil && !IsNil(o.ClaimsLocales) {
		return true
	}

	return false
}

// SetClaimsLocales gets a reference to the given []string and assigns it to the ClaimsLocales field.
func (o *AuthorizationResponse) SetClaimsLocales(v []string) {
	o.ClaimsLocales = v
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetClaims() []string {
	if o == nil || IsNil(o.Claims) {
		var ret []string
		return ret
	}
	return o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetClaimsOk() ([]string, bool) {
	if o == nil || IsNil(o.Claims) {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasClaims() bool {
	if o != nil && !IsNil(o.Claims) {
		return true
	}

	return false
}

// SetClaims gets a reference to the given []string and assigns it to the Claims field.
func (o *AuthorizationResponse) SetClaims(v []string) {
	o.Claims = v
}

// GetAcrEssential returns the AcrEssential field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetAcrEssential() bool {
	if o == nil || IsNil(o.AcrEssential) {
		var ret bool
		return ret
	}
	return *o.AcrEssential
}

// GetAcrEssentialOk returns a tuple with the AcrEssential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetAcrEssentialOk() (*bool, bool) {
	if o == nil || IsNil(o.AcrEssential) {
		return nil, false
	}
	return o.AcrEssential, true
}

// HasAcrEssential returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasAcrEssential() bool {
	if o != nil && !IsNil(o.AcrEssential) {
		return true
	}

	return false
}

// SetAcrEssential gets a reference to the given bool and assigns it to the AcrEssential field.
func (o *AuthorizationResponse) SetAcrEssential(v bool) {
	o.AcrEssential = &v
}

// GetClientIdAliasUsed returns the ClientIdAliasUsed field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetClientIdAliasUsed() bool {
	if o == nil || IsNil(o.ClientIdAliasUsed) {
		var ret bool
		return ret
	}
	return *o.ClientIdAliasUsed
}

// GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetClientIdAliasUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientIdAliasUsed) {
		return nil, false
	}
	return o.ClientIdAliasUsed, true
}

// HasClientIdAliasUsed returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasClientIdAliasUsed() bool {
	if o != nil && !IsNil(o.ClientIdAliasUsed) {
		return true
	}

	return false
}

// SetClientIdAliasUsed gets a reference to the given bool and assigns it to the ClientIdAliasUsed field.
func (o *AuthorizationResponse) SetClientIdAliasUsed(v bool) {
	o.ClientIdAliasUsed = &v
}

// GetAcrs returns the Acrs field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetAcrs() []string {
	if o == nil || IsNil(o.Acrs) {
		var ret []string
		return ret
	}
	return o.Acrs
}

// GetAcrsOk returns a tuple with the Acrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetAcrsOk() ([]string, bool) {
	if o == nil || IsNil(o.Acrs) {
		return nil, false
	}
	return o.Acrs, true
}

// HasAcrs returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasAcrs() bool {
	if o != nil && !IsNil(o.Acrs) {
		return true
	}

	return false
}

// SetAcrs gets a reference to the given []string and assigns it to the Acrs field.
func (o *AuthorizationResponse) SetAcrs(v []string) {
	o.Acrs = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *AuthorizationResponse) SetSubject(v string) {
	o.Subject = &v
}

// GetLoginHint returns the LoginHint field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetLoginHint() string {
	if o == nil || IsNil(o.LoginHint) {
		var ret string
		return ret
	}
	return *o.LoginHint
}

// GetLoginHintOk returns a tuple with the LoginHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetLoginHintOk() (*string, bool) {
	if o == nil || IsNil(o.LoginHint) {
		return nil, false
	}
	return o.LoginHint, true
}

// HasLoginHint returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasLoginHint() bool {
	if o != nil && !IsNil(o.LoginHint) {
		return true
	}

	return false
}

// SetLoginHint gets a reference to the given string and assigns it to the LoginHint field.
func (o *AuthorizationResponse) SetLoginHint(v string) {
	o.LoginHint = &v
}

// GetPrompts returns the Prompts field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetPrompts() []Prompt {
	if o == nil || IsNil(o.Prompts) {
		var ret []Prompt
		return ret
	}
	return o.Prompts
}

// GetPromptsOk returns a tuple with the Prompts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetPromptsOk() ([]Prompt, bool) {
	if o == nil || IsNil(o.Prompts) {
		return nil, false
	}
	return o.Prompts, true
}

// HasPrompts returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasPrompts() bool {
	if o != nil && !IsNil(o.Prompts) {
		return true
	}

	return false
}

// SetPrompts gets a reference to the given []Prompt and assigns it to the Prompts field.
func (o *AuthorizationResponse) SetPrompts(v []Prompt) {
	o.Prompts = v
}

// GetLowestPrompt returns the LowestPrompt field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetLowestPrompt() Prompt {
	if o == nil || IsNil(o.LowestPrompt) {
		var ret Prompt
		return ret
	}
	return *o.LowestPrompt
}

// GetLowestPromptOk returns a tuple with the LowestPrompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetLowestPromptOk() (*Prompt, bool) {
	if o == nil || IsNil(o.LowestPrompt) {
		return nil, false
	}
	return o.LowestPrompt, true
}

// HasLowestPrompt returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasLowestPrompt() bool {
	if o != nil && !IsNil(o.LowestPrompt) {
		return true
	}

	return false
}

// SetLowestPrompt gets a reference to the given Prompt and assigns it to the LowestPrompt field.
func (o *AuthorizationResponse) SetLowestPrompt(v Prompt) {
	o.LowestPrompt = &v
}

// GetRequestObjectPayload returns the RequestObjectPayload field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetRequestObjectPayload() string {
	if o == nil || IsNil(o.RequestObjectPayload) {
		var ret string
		return ret
	}
	return *o.RequestObjectPayload
}

// GetRequestObjectPayloadOk returns a tuple with the RequestObjectPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetRequestObjectPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.RequestObjectPayload) {
		return nil, false
	}
	return o.RequestObjectPayload, true
}

// HasRequestObjectPayload returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasRequestObjectPayload() bool {
	if o != nil && !IsNil(o.RequestObjectPayload) {
		return true
	}

	return false
}

// SetRequestObjectPayload gets a reference to the given string and assigns it to the RequestObjectPayload field.
func (o *AuthorizationResponse) SetRequestObjectPayload(v string) {
	o.RequestObjectPayload = &v
}

// GetIdTokenClaims returns the IdTokenClaims field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetIdTokenClaims() string {
	if o == nil || IsNil(o.IdTokenClaims) {
		var ret string
		return ret
	}
	return *o.IdTokenClaims
}

// GetIdTokenClaimsOk returns a tuple with the IdTokenClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetIdTokenClaimsOk() (*string, bool) {
	if o == nil || IsNil(o.IdTokenClaims) {
		return nil, false
	}
	return o.IdTokenClaims, true
}

// HasIdTokenClaims returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasIdTokenClaims() bool {
	if o != nil && !IsNil(o.IdTokenClaims) {
		return true
	}

	return false
}

// SetIdTokenClaims gets a reference to the given string and assigns it to the IdTokenClaims field.
func (o *AuthorizationResponse) SetIdTokenClaims(v string) {
	o.IdTokenClaims = &v
}

// GetUserInfoClaims returns the UserInfoClaims field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetUserInfoClaims() string {
	if o == nil || IsNil(o.UserInfoClaims) {
		var ret string
		return ret
	}
	return *o.UserInfoClaims
}

// GetUserInfoClaimsOk returns a tuple with the UserInfoClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetUserInfoClaimsOk() (*string, bool) {
	if o == nil || IsNil(o.UserInfoClaims) {
		return nil, false
	}
	return o.UserInfoClaims, true
}

// HasUserInfoClaims returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasUserInfoClaims() bool {
	if o != nil && !IsNil(o.UserInfoClaims) {
		return true
	}

	return false
}

// SetUserInfoClaims gets a reference to the given string and assigns it to the UserInfoClaims field.
func (o *AuthorizationResponse) SetUserInfoClaims(v string) {
	o.UserInfoClaims = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetResources() []string {
	if o == nil || IsNil(o.Resources) {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetResourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *AuthorizationResponse) SetResources(v []string) {
	o.Resources = v
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetAuthorizationDetails() AuthzDetails {
	if o == nil || IsNil(o.AuthorizationDetails) {
		var ret AuthzDetails
		return ret
	}
	return *o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetAuthorizationDetailsOk() (*AuthzDetails, bool) {
	if o == nil || IsNil(o.AuthorizationDetails) {
		return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasAuthorizationDetails() bool {
	if o != nil && !IsNil(o.AuthorizationDetails) {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given AuthzDetails and assigns it to the AuthorizationDetails field.
func (o *AuthorizationResponse) SetAuthorizationDetails(v AuthzDetails) {
	o.AuthorizationDetails = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetPurpose() string {
	if o == nil || IsNil(o.Purpose) {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetPurposeOk() (*string, bool) {
	if o == nil || IsNil(o.Purpose) {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasPurpose() bool {
	if o != nil && !IsNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *AuthorizationResponse) SetPurpose(v string) {
	o.Purpose = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetResponseContent() string {
	if o == nil || IsNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasResponseContent() bool {
	if o != nil && !IsNil(o.ResponseContent) {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *AuthorizationResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetTicket returns the Ticket field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetTicket() string {
	if o == nil || IsNil(o.Ticket) {
		var ret string
		return ret
	}
	return *o.Ticket
}

// GetTicketOk returns a tuple with the Ticket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetTicketOk() (*string, bool) {
	if o == nil || IsNil(o.Ticket) {
		return nil, false
	}
	return o.Ticket, true
}

// HasTicket returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasTicket() bool {
	if o != nil && !IsNil(o.Ticket) {
		return true
	}

	return false
}

// SetTicket gets a reference to the given string and assigns it to the Ticket field.
func (o *AuthorizationResponse) SetTicket(v string) {
	o.Ticket = &v
}

// GetDynamicScopes returns the DynamicScopes field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetDynamicScopes() []DynamicScope {
	if o == nil || IsNil(o.DynamicScopes) {
		var ret []DynamicScope
		return ret
	}
	return o.DynamicScopes
}

// GetDynamicScopesOk returns a tuple with the DynamicScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetDynamicScopesOk() ([]DynamicScope, bool) {
	if o == nil || IsNil(o.DynamicScopes) {
		return nil, false
	}
	return o.DynamicScopes, true
}

// HasDynamicScopes returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasDynamicScopes() bool {
	if o != nil && !IsNil(o.DynamicScopes) {
		return true
	}

	return false
}

// SetDynamicScopes gets a reference to the given []DynamicScope and assigns it to the DynamicScopes field.
func (o *AuthorizationResponse) SetDynamicScopes(v []DynamicScope) {
	o.DynamicScopes = v
}

// GetGmAction returns the GmAction field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetGmAction() GrantManagementAction {
	if o == nil || IsNil(o.GmAction) {
		var ret GrantManagementAction
		return ret
	}
	return *o.GmAction
}

// GetGmActionOk returns a tuple with the GmAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetGmActionOk() (*GrantManagementAction, bool) {
	if o == nil || IsNil(o.GmAction) {
		return nil, false
	}
	return o.GmAction, true
}

// HasGmAction returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasGmAction() bool {
	if o != nil && !IsNil(o.GmAction) {
		return true
	}

	return false
}

// SetGmAction gets a reference to the given GrantManagementAction and assigns it to the GmAction field.
func (o *AuthorizationResponse) SetGmAction(v GrantManagementAction) {
	o.GmAction = &v
}

// GetGrantId returns the GrantId field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetGrantId() string {
	if o == nil || IsNil(o.GrantId) {
		var ret string
		return ret
	}
	return *o.GrantId
}

// GetGrantIdOk returns a tuple with the GrantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetGrantIdOk() (*string, bool) {
	if o == nil || IsNil(o.GrantId) {
		return nil, false
	}
	return o.GrantId, true
}

// HasGrantId returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasGrantId() bool {
	if o != nil && !IsNil(o.GrantId) {
		return true
	}

	return false
}

// SetGrantId gets a reference to the given string and assigns it to the GrantId field.
func (o *AuthorizationResponse) SetGrantId(v string) {
	o.GrantId = &v
}

// GetGrant returns the Grant field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetGrant() Grant {
	if o == nil || IsNil(o.Grant) {
		var ret Grant
		return ret
	}
	return *o.Grant
}

// GetGrantOk returns a tuple with the Grant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetGrantOk() (*Grant, bool) {
	if o == nil || IsNil(o.Grant) {
		return nil, false
	}
	return o.Grant, true
}

// HasGrant returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasGrant() bool {
	if o != nil && !IsNil(o.Grant) {
		return true
	}

	return false
}

// SetGrant gets a reference to the given Grant and assigns it to the Grant field.
func (o *AuthorizationResponse) SetGrant(v Grant) {
	o.Grant = &v
}

// GetGrantSubject returns the GrantSubject field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetGrantSubject() string {
	if o == nil || IsNil(o.GrantSubject) {
		var ret string
		return ret
	}
	return *o.GrantSubject
}

// GetGrantSubjectOk returns a tuple with the GrantSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetGrantSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.GrantSubject) {
		return nil, false
	}
	return o.GrantSubject, true
}

// HasGrantSubject returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasGrantSubject() bool {
	if o != nil && !IsNil(o.GrantSubject) {
		return true
	}

	return false
}

// SetGrantSubject gets a reference to the given string and assigns it to the GrantSubject field.
func (o *AuthorizationResponse) SetGrantSubject(v string) {
	o.GrantSubject = &v
}

// GetRequestedClaimsForTx returns the RequestedClaimsForTx field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetRequestedClaimsForTx() []string {
	if o == nil || IsNil(o.RequestedClaimsForTx) {
		var ret []string
		return ret
	}
	return o.RequestedClaimsForTx
}

// GetRequestedClaimsForTxOk returns a tuple with the RequestedClaimsForTx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetRequestedClaimsForTxOk() ([]string, bool) {
	if o == nil || IsNil(o.RequestedClaimsForTx) {
		return nil, false
	}
	return o.RequestedClaimsForTx, true
}

// HasRequestedClaimsForTx returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasRequestedClaimsForTx() bool {
	if o != nil && !IsNil(o.RequestedClaimsForTx) {
		return true
	}

	return false
}

// SetRequestedClaimsForTx gets a reference to the given []string and assigns it to the RequestedClaimsForTx field.
func (o *AuthorizationResponse) SetRequestedClaimsForTx(v []string) {
	o.RequestedClaimsForTx = v
}

// GetRequestedVerifiedClaimsForTx returns the RequestedVerifiedClaimsForTx field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetRequestedVerifiedClaimsForTx() [][]string {
	if o == nil || IsNil(o.RequestedVerifiedClaimsForTx) {
		var ret [][]string
		return ret
	}
	return o.RequestedVerifiedClaimsForTx
}

// GetRequestedVerifiedClaimsForTxOk returns a tuple with the RequestedVerifiedClaimsForTx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetRequestedVerifiedClaimsForTxOk() ([][]string, bool) {
	if o == nil || IsNil(o.RequestedVerifiedClaimsForTx) {
		return nil, false
	}
	return o.RequestedVerifiedClaimsForTx, true
}

// HasRequestedVerifiedClaimsForTx returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasRequestedVerifiedClaimsForTx() bool {
	if o != nil && !IsNil(o.RequestedVerifiedClaimsForTx) {
		return true
	}

	return false
}

// SetRequestedVerifiedClaimsForTx gets a reference to the given [][]string and assigns it to the RequestedVerifiedClaimsForTx field.
func (o *AuthorizationResponse) SetRequestedVerifiedClaimsForTx(v [][]string) {
	o.RequestedVerifiedClaimsForTx = v
}

// GetTransformedClaims returns the TransformedClaims field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetTransformedClaims() string {
	if o == nil || IsNil(o.TransformedClaims) {
		var ret string
		return ret
	}
	return *o.TransformedClaims
}

// GetTransformedClaimsOk returns a tuple with the TransformedClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetTransformedClaimsOk() (*string, bool) {
	if o == nil || IsNil(o.TransformedClaims) {
		return nil, false
	}
	return o.TransformedClaims, true
}

// HasTransformedClaims returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasTransformedClaims() bool {
	if o != nil && !IsNil(o.TransformedClaims) {
		return true
	}

	return false
}

// SetTransformedClaims gets a reference to the given string and assigns it to the TransformedClaims field.
func (o *AuthorizationResponse) SetTransformedClaims(v string) {
	o.TransformedClaims = &v
}

// GetClientEntityIdUsed returns the ClientEntityIdUsed field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetClientEntityIdUsed() bool {
	if o == nil || IsNil(o.ClientEntityIdUsed) {
		var ret bool
		return ret
	}
	return *o.ClientEntityIdUsed
}

// GetClientEntityIdUsedOk returns a tuple with the ClientEntityIdUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetClientEntityIdUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientEntityIdUsed) {
		return nil, false
	}
	return o.ClientEntityIdUsed, true
}

// HasClientEntityIdUsed returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasClientEntityIdUsed() bool {
	if o != nil && !IsNil(o.ClientEntityIdUsed) {
		return true
	}

	return false
}

// SetClientEntityIdUsed gets a reference to the given bool and assigns it to the ClientEntityIdUsed field.
func (o *AuthorizationResponse) SetClientEntityIdUsed(v bool) {
	o.ClientEntityIdUsed = &v
}

// GetClaimsAtUserInfo returns the ClaimsAtUserInfo field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetClaimsAtUserInfo() []string {
	if o == nil || IsNil(o.ClaimsAtUserInfo) {
		var ret []string
		return ret
	}
	return o.ClaimsAtUserInfo
}

// GetClaimsAtUserInfoOk returns a tuple with the ClaimsAtUserInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetClaimsAtUserInfoOk() ([]string, bool) {
	if o == nil || IsNil(o.ClaimsAtUserInfo) {
		return nil, false
	}
	return o.ClaimsAtUserInfo, true
}

// HasClaimsAtUserInfo returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasClaimsAtUserInfo() bool {
	if o != nil && !IsNil(o.ClaimsAtUserInfo) {
		return true
	}

	return false
}

// SetClaimsAtUserInfo gets a reference to the given []string and assigns it to the ClaimsAtUserInfo field.
func (o *AuthorizationResponse) SetClaimsAtUserInfo(v []string) {
	o.ClaimsAtUserInfo = v
}

// GetCredentialOfferInfo returns the CredentialOfferInfo field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetCredentialOfferInfo() CredentialOfferInfo {
	if o == nil || IsNil(o.CredentialOfferInfo) {
		var ret CredentialOfferInfo
		return ret
	}
	return *o.CredentialOfferInfo
}

// GetCredentialOfferInfoOk returns a tuple with the CredentialOfferInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetCredentialOfferInfoOk() (*CredentialOfferInfo, bool) {
	if o == nil || IsNil(o.CredentialOfferInfo) {
		return nil, false
	}
	return o.CredentialOfferInfo, true
}

// HasCredentialOfferInfo returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasCredentialOfferInfo() bool {
	if o != nil && !IsNil(o.CredentialOfferInfo) {
		return true
	}

	return false
}

// SetCredentialOfferInfo gets a reference to the given CredentialOfferInfo and assigns it to the CredentialOfferInfo field.
func (o *AuthorizationResponse) SetCredentialOfferInfo(v CredentialOfferInfo) {
	o.CredentialOfferInfo = &v
}

// GetIssuableCredentials returns the IssuableCredentials field value if set, zero value otherwise.
func (o *AuthorizationResponse) GetIssuableCredentials() string {
	if o == nil || IsNil(o.IssuableCredentials) {
		var ret string
		return ret
	}
	return *o.IssuableCredentials
}

// GetIssuableCredentialsOk returns a tuple with the IssuableCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetIssuableCredentialsOk() (*string, bool) {
	if o == nil || IsNil(o.IssuableCredentials) {
		return nil, false
	}
	return o.IssuableCredentials, true
}

// HasIssuableCredentials returns a boolean if a field has been set.
func (o *AuthorizationResponse) HasIssuableCredentials() bool {
	if o != nil && !IsNil(o.IssuableCredentials) {
		return true
	}

	return false
}

// SetIssuableCredentials gets a reference to the given string and assigns it to the IssuableCredentials field.
func (o *AuthorizationResponse) SetIssuableCredentials(v string) {
	o.IssuableCredentials = &v
}

func (o AuthorizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResultCode) {
		toSerialize["resultCode"] = o.ResultCode
	}
	if !IsNil(o.ResultMessage) {
		toSerialize["resultMessage"] = o.ResultMessage
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	if !IsNil(o.MaxAge) {
		toSerialize["maxAge"] = o.MaxAge
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.UiLocales) {
		toSerialize["uiLocales"] = o.UiLocales
	}
	if !IsNil(o.ClaimsLocales) {
		toSerialize["claimsLocales"] = o.ClaimsLocales
	}
	if !IsNil(o.Claims) {
		toSerialize["claims"] = o.Claims
	}
	if !IsNil(o.AcrEssential) {
		toSerialize["acrEssential"] = o.AcrEssential
	}
	if !IsNil(o.ClientIdAliasUsed) {
		toSerialize["clientIdAliasUsed"] = o.ClientIdAliasUsed
	}
	if !IsNil(o.Acrs) {
		toSerialize["acrs"] = o.Acrs
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.LoginHint) {
		toSerialize["loginHint"] = o.LoginHint
	}
	if !IsNil(o.Prompts) {
		toSerialize["prompts"] = o.Prompts
	}
	if !IsNil(o.LowestPrompt) {
		toSerialize["lowestPrompt"] = o.LowestPrompt
	}
	if !IsNil(o.RequestObjectPayload) {
		toSerialize["requestObjectPayload"] = o.RequestObjectPayload
	}
	if !IsNil(o.IdTokenClaims) {
		toSerialize["idTokenClaims"] = o.IdTokenClaims
	}
	if !IsNil(o.UserInfoClaims) {
		toSerialize["userInfoClaims"] = o.UserInfoClaims
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.AuthorizationDetails) {
		toSerialize["authorizationDetails"] = o.AuthorizationDetails
	}
	if !IsNil(o.Purpose) {
		toSerialize["purpose"] = o.Purpose
	}
	if !IsNil(o.ResponseContent) {
		toSerialize["responseContent"] = o.ResponseContent
	}
	if !IsNil(o.Ticket) {
		toSerialize["ticket"] = o.Ticket
	}
	if !IsNil(o.DynamicScopes) {
		toSerialize["dynamicScopes"] = o.DynamicScopes
	}
	if !IsNil(o.GmAction) {
		toSerialize["gmAction"] = o.GmAction
	}
	if !IsNil(o.GrantId) {
		toSerialize["grantId"] = o.GrantId
	}
	if !IsNil(o.Grant) {
		toSerialize["grant"] = o.Grant
	}
	if !IsNil(o.GrantSubject) {
		toSerialize["grantSubject"] = o.GrantSubject
	}
	if !IsNil(o.RequestedClaimsForTx) {
		toSerialize["requestedClaimsForTx"] = o.RequestedClaimsForTx
	}
	if !IsNil(o.RequestedVerifiedClaimsForTx) {
		toSerialize["requestedVerifiedClaimsForTx"] = o.RequestedVerifiedClaimsForTx
	}
	if !IsNil(o.TransformedClaims) {
		toSerialize["transformedClaims"] = o.TransformedClaims
	}
	if !IsNil(o.ClientEntityIdUsed) {
		toSerialize["clientEntityIdUsed"] = o.ClientEntityIdUsed
	}
	if !IsNil(o.ClaimsAtUserInfo) {
		toSerialize["claimsAtUserInfo"] = o.ClaimsAtUserInfo
	}
	if !IsNil(o.CredentialOfferInfo) {
		toSerialize["credentialOfferInfo"] = o.CredentialOfferInfo
	}
	if !IsNil(o.IssuableCredentials) {
		toSerialize["issuableCredentials"] = o.IssuableCredentials
	}
	return toSerialize, nil
}

type NullableAuthorizationResponse struct {
	value *AuthorizationResponse
	isSet bool
}

func (v NullableAuthorizationResponse) Get() *AuthorizationResponse {
	return v.value
}

func (v *NullableAuthorizationResponse) Set(val *AuthorizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationResponse(val *AuthorizationResponse) *NullableAuthorizationResponse {
	return &NullableAuthorizationResponse{value: val, isSet: true}
}

func (v NullableAuthorizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
