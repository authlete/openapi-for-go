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

// checks if the BackchannelAuthenticationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackchannelAuthenticationResponse{}

// BackchannelAuthenticationResponse struct for BackchannelAuthenticationResponse
type BackchannelAuthenticationResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take.
	Action *string `json:"action,omitempty"`
	// The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of `action` parameter.
	ResponseContent *string `json:"responseContent,omitempty"`
	// The client ID of the client application that has made the backchannel authentication request.
	ClientId *int64 `json:"clientId,omitempty"`
	// The client ID alias of the client application that has made the backchannel authentication request.
	ClientIdAlias *string `json:"clientIdAlias,omitempty"`
	// `true` if the value of the client_id request parameter included in the backchannel authentication request is the client ID alias. `false` if the value is the original numeric client ID.
	ClientIdAliasUsed *bool `json:"clientIdAliasUsed,omitempty"`
	// The name of the client application which has made the backchannel authentication request.
	ClientName *string `json:"clientName,omitempty"`
	// The scopes requested by the backchannel authentication request.  Basically, this property holds the value of the `scope` request parameter in the backchannel authentication request. However, because unregistered scopes are dropped on Authlete side, if the `scope` request parameter contains unknown scopes, the list returned by this property becomes different from the value of the `scope` request parameter.  Note that `description` property and `descriptions` property of each `scope` object in the array contained in this property is always null even if descriptions of the scopes are registered.
	Scopes []Scope `json:"scopes,omitempty"`
	// The names of the claims which were requested indirectly via some special scopes. See [5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims) in OpenID Connect Core 1.0 for details.
	ClaimNames []string `json:"claimNames,omitempty"`
	// The client notification token included in the backchannel authentication request.
	ClientNotificationToken *string `json:"clientNotificationToken,omitempty"`
	// The list of ACR values requested by the backchannel authentication request.  Basically, this property holds the value of the `acr_values` request parameter in the backchannel authentication request. However, because unsupported ACR values are dropped on Authlete side, if the `acr_values` request parameter contains unrecognized ACR values, the list returned by this property becomes different from the value of the `acr_values` request parameter.
	Acrs []string `json:"acrs,omitempty"`
	// The type of the hint for end-user identification which was included in the backchannel authentication request.
	HintType *string `json:"hintType,omitempty"`
	// The value of the hint for end-user identification.
	Hint *string `json:"hint,omitempty"`
	// The value of the `sub` claim contained in the ID token hint included in the backchannel authentication request.
	Sub *string `json:"sub,omitempty"`
	// The binding message included in the backchannel authentication request.
	BindingMessage *string `json:"bindingMessage,omitempty"`
	// The binding message included in the backchannel authentication request.
	UserCode *string `json:"userCode,omitempty"`
	// The flag which indicates whether a user code is required.  `true` when both the `backchannel_user_code_parameter` metadata of the client (= Client's `bcUserCodeRequired` property) and the `backchannel_user_code_parameter_supported` metadata of the service (= Service's `backchannelUserCodeParameterSupported` property) are `true`.
	UserCodeRequired *bool `json:"userCodeRequired,omitempty"`
	// The requested expiry for the authentication request ID (`auth_req_id`).
	RequestedExpiry *int32 `json:"requestedExpiry,omitempty"`
	// The request context of the backchannel authentication request.  It is the value of the request_context claim in the signed authentication request and its format is JSON. request_context is a new claim added by the FAPI-CIBA profile.
	RequestContext *string `json:"requestContext,omitempty"`
	// The warnings raised during processing the backchannel authentication request.
	Warnings []string `json:"warnings,omitempty"`
	// The ticket which is necessary to call Authlete's `/auth/token/fail` API or `/auth/token/issue` API.  This parameter has a value only if the value of `grant_type` request parameter is `password` and the token request is valid.
	Ticket *string `json:"ticket,omitempty"`
	// The resources specified by the `resource` request parameters or by the `resource` property in the request object. If both are given, the values in the request object should be set. See \"Resource Indicators for OAuth 2.0\" for details.
	Resources            []string      `json:"resources,omitempty"`
	AuthorizationDetails *AuthzDetails `json:"authorizationDetails,omitempty"`
	// The attributes of this service that the client application belongs to.
	ServiceAttributes []Pair `json:"serviceAttributes,omitempty"`
	// The attributes of the client.
	ClientAttributes []Pair `json:"clientAttributes,omitempty"`
	// The dynamic scopes which the client application requested by the scope request parameter.
	DynamicScopes []DynamicScope `json:"dynamicScopes,omitempty"`
	DeliveryMode  *DeliveryMode  `json:"deliveryMode,omitempty"`
	// The client authentication method that was performed.
	ClientAuthMethod *string                `json:"clientAuthMethod,omitempty"`
	GmAction         *GrantManagementAction `json:"gmAction,omitempty"`
	// the value of the `grant_id` request parameter of the device authorization request.  The `grant_id` request parameter is defined in [Grant Management for OAuth 2.0](https://openid.net/specs/fapi-grant-management.html) , which is supported by Authlete 2.3 and newer versions.
	GrantId *string `json:"grantId,omitempty"`
	Grant   *Grant  `json:"grant,omitempty"`
	// The subject identifying the user who has given the grant identified by the `grant_id` request parameter of the device authorization request.  Authlete 2.3 and newer versions support <a href= \"https://openid.net/specs/fapi-grant-management.html\">Grant Management for OAuth 2.0</a>. An authorization request may contain a `grant_id` request parameter which is defined in the specification. If the value of the request parameter is valid, {@link #getGrantSubject()} will return the subject of the user who has given the grant to the client application. Authorization server implementations may use the value returned from {@link #getGrantSubject()} in order to determine the user to authenticate.  The user your system will authenticate during the authorization process (or has already authenticated) may be different from the user of the grant. The first implementer's draft of \"Grant Management for OAuth 2.0\" does not mention anything about the case, so the behavior in the case is left to implementations. Authlete will not perform the grant management action when the `subject` passed to Authlete does not match the user of the grant.
	GrantSubject *string `json:"grantSubject,omitempty"`
	// The entity ID of the client.
	ClientEntityId *string `json:"clientEntityId,omitempty"`
	// Flag which indicates whether the entity ID of the client was used when the request for the access token was made.
	ClientEntityIdUsed *bool `json:"clientEntityIdUsed,omitempty"`
}

// NewBackchannelAuthenticationResponse instantiates a new BackchannelAuthenticationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackchannelAuthenticationResponse() *BackchannelAuthenticationResponse {
	this := BackchannelAuthenticationResponse{}
	return &this
}

// NewBackchannelAuthenticationResponseWithDefaults instantiates a new BackchannelAuthenticationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackchannelAuthenticationResponseWithDefaults() *BackchannelAuthenticationResponse {
	this := BackchannelAuthenticationResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetResultCode() string {
	if o == nil || IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *BackchannelAuthenticationResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetResultMessage() string {
	if o == nil || IsNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasResultMessage() bool {
	if o != nil && !IsNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *BackchannelAuthenticationResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *BackchannelAuthenticationResponse) SetAction(v string) {
	o.Action = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetResponseContent() string {
	if o == nil || IsNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasResponseContent() bool {
	if o != nil && !IsNil(o.ResponseContent) {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *BackchannelAuthenticationResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetClientId() int64 {
	if o == nil || IsNil(o.ClientId) {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetClientIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int64 and assigns it to the ClientId field.
func (o *BackchannelAuthenticationResponse) SetClientId(v int64) {
	o.ClientId = &v
}

// GetClientIdAlias returns the ClientIdAlias field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetClientIdAlias() string {
	if o == nil || IsNil(o.ClientIdAlias) {
		var ret string
		return ret
	}
	return *o.ClientIdAlias
}

// GetClientIdAliasOk returns a tuple with the ClientIdAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetClientIdAliasOk() (*string, bool) {
	if o == nil || IsNil(o.ClientIdAlias) {
		return nil, false
	}
	return o.ClientIdAlias, true
}

// HasClientIdAlias returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasClientIdAlias() bool {
	if o != nil && !IsNil(o.ClientIdAlias) {
		return true
	}

	return false
}

// SetClientIdAlias gets a reference to the given string and assigns it to the ClientIdAlias field.
func (o *BackchannelAuthenticationResponse) SetClientIdAlias(v string) {
	o.ClientIdAlias = &v
}

// GetClientIdAliasUsed returns the ClientIdAliasUsed field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetClientIdAliasUsed() bool {
	if o == nil || IsNil(o.ClientIdAliasUsed) {
		var ret bool
		return ret
	}
	return *o.ClientIdAliasUsed
}

// GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetClientIdAliasUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientIdAliasUsed) {
		return nil, false
	}
	return o.ClientIdAliasUsed, true
}

// HasClientIdAliasUsed returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasClientIdAliasUsed() bool {
	if o != nil && !IsNil(o.ClientIdAliasUsed) {
		return true
	}

	return false
}

// SetClientIdAliasUsed gets a reference to the given bool and assigns it to the ClientIdAliasUsed field.
func (o *BackchannelAuthenticationResponse) SetClientIdAliasUsed(v bool) {
	o.ClientIdAliasUsed = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetClientName() string {
	if o == nil || IsNil(o.ClientName) {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetClientNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClientName) {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasClientName() bool {
	if o != nil && !IsNil(o.ClientName) {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *BackchannelAuthenticationResponse) SetClientName(v string) {
	o.ClientName = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetScopes() []Scope {
	if o == nil || IsNil(o.Scopes) {
		var ret []Scope
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetScopesOk() ([]Scope, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []Scope and assigns it to the Scopes field.
func (o *BackchannelAuthenticationResponse) SetScopes(v []Scope) {
	o.Scopes = v
}

// GetClaimNames returns the ClaimNames field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetClaimNames() []string {
	if o == nil || IsNil(o.ClaimNames) {
		var ret []string
		return ret
	}
	return o.ClaimNames
}

// GetClaimNamesOk returns a tuple with the ClaimNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetClaimNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.ClaimNames) {
		return nil, false
	}
	return o.ClaimNames, true
}

// HasClaimNames returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasClaimNames() bool {
	if o != nil && !IsNil(o.ClaimNames) {
		return true
	}

	return false
}

// SetClaimNames gets a reference to the given []string and assigns it to the ClaimNames field.
func (o *BackchannelAuthenticationResponse) SetClaimNames(v []string) {
	o.ClaimNames = v
}

// GetClientNotificationToken returns the ClientNotificationToken field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetClientNotificationToken() string {
	if o == nil || IsNil(o.ClientNotificationToken) {
		var ret string
		return ret
	}
	return *o.ClientNotificationToken
}

// GetClientNotificationTokenOk returns a tuple with the ClientNotificationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetClientNotificationTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ClientNotificationToken) {
		return nil, false
	}
	return o.ClientNotificationToken, true
}

// HasClientNotificationToken returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasClientNotificationToken() bool {
	if o != nil && !IsNil(o.ClientNotificationToken) {
		return true
	}

	return false
}

// SetClientNotificationToken gets a reference to the given string and assigns it to the ClientNotificationToken field.
func (o *BackchannelAuthenticationResponse) SetClientNotificationToken(v string) {
	o.ClientNotificationToken = &v
}

// GetAcrs returns the Acrs field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetAcrs() []string {
	if o == nil || IsNil(o.Acrs) {
		var ret []string
		return ret
	}
	return o.Acrs
}

// GetAcrsOk returns a tuple with the Acrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetAcrsOk() ([]string, bool) {
	if o == nil || IsNil(o.Acrs) {
		return nil, false
	}
	return o.Acrs, true
}

// HasAcrs returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasAcrs() bool {
	if o != nil && !IsNil(o.Acrs) {
		return true
	}

	return false
}

// SetAcrs gets a reference to the given []string and assigns it to the Acrs field.
func (o *BackchannelAuthenticationResponse) SetAcrs(v []string) {
	o.Acrs = v
}

// GetHintType returns the HintType field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetHintType() string {
	if o == nil || IsNil(o.HintType) {
		var ret string
		return ret
	}
	return *o.HintType
}

// GetHintTypeOk returns a tuple with the HintType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetHintTypeOk() (*string, bool) {
	if o == nil || IsNil(o.HintType) {
		return nil, false
	}
	return o.HintType, true
}

// HasHintType returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasHintType() bool {
	if o != nil && !IsNil(o.HintType) {
		return true
	}

	return false
}

// SetHintType gets a reference to the given string and assigns it to the HintType field.
func (o *BackchannelAuthenticationResponse) SetHintType(v string) {
	o.HintType = &v
}

// GetHint returns the Hint field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetHint() string {
	if o == nil || IsNil(o.Hint) {
		var ret string
		return ret
	}
	return *o.Hint
}

// GetHintOk returns a tuple with the Hint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetHintOk() (*string, bool) {
	if o == nil || IsNil(o.Hint) {
		return nil, false
	}
	return o.Hint, true
}

// HasHint returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasHint() bool {
	if o != nil && !IsNil(o.Hint) {
		return true
	}

	return false
}

// SetHint gets a reference to the given string and assigns it to the Hint field.
func (o *BackchannelAuthenticationResponse) SetHint(v string) {
	o.Hint = &v
}

// GetSub returns the Sub field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetSub() string {
	if o == nil || IsNil(o.Sub) {
		var ret string
		return ret
	}
	return *o.Sub
}

// GetSubOk returns a tuple with the Sub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetSubOk() (*string, bool) {
	if o == nil || IsNil(o.Sub) {
		return nil, false
	}
	return o.Sub, true
}

// HasSub returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasSub() bool {
	if o != nil && !IsNil(o.Sub) {
		return true
	}

	return false
}

// SetSub gets a reference to the given string and assigns it to the Sub field.
func (o *BackchannelAuthenticationResponse) SetSub(v string) {
	o.Sub = &v
}

// GetBindingMessage returns the BindingMessage field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetBindingMessage() string {
	if o == nil || IsNil(o.BindingMessage) {
		var ret string
		return ret
	}
	return *o.BindingMessage
}

// GetBindingMessageOk returns a tuple with the BindingMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetBindingMessageOk() (*string, bool) {
	if o == nil || IsNil(o.BindingMessage) {
		return nil, false
	}
	return o.BindingMessage, true
}

// HasBindingMessage returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasBindingMessage() bool {
	if o != nil && !IsNil(o.BindingMessage) {
		return true
	}

	return false
}

// SetBindingMessage gets a reference to the given string and assigns it to the BindingMessage field.
func (o *BackchannelAuthenticationResponse) SetBindingMessage(v string) {
	o.BindingMessage = &v
}

// GetUserCode returns the UserCode field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetUserCode() string {
	if o == nil || IsNil(o.UserCode) {
		var ret string
		return ret
	}
	return *o.UserCode
}

// GetUserCodeOk returns a tuple with the UserCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetUserCodeOk() (*string, bool) {
	if o == nil || IsNil(o.UserCode) {
		return nil, false
	}
	return o.UserCode, true
}

// HasUserCode returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasUserCode() bool {
	if o != nil && !IsNil(o.UserCode) {
		return true
	}

	return false
}

// SetUserCode gets a reference to the given string and assigns it to the UserCode field.
func (o *BackchannelAuthenticationResponse) SetUserCode(v string) {
	o.UserCode = &v
}

// GetUserCodeRequired returns the UserCodeRequired field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetUserCodeRequired() bool {
	if o == nil || IsNil(o.UserCodeRequired) {
		var ret bool
		return ret
	}
	return *o.UserCodeRequired
}

// GetUserCodeRequiredOk returns a tuple with the UserCodeRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetUserCodeRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.UserCodeRequired) {
		return nil, false
	}
	return o.UserCodeRequired, true
}

// HasUserCodeRequired returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasUserCodeRequired() bool {
	if o != nil && !IsNil(o.UserCodeRequired) {
		return true
	}

	return false
}

// SetUserCodeRequired gets a reference to the given bool and assigns it to the UserCodeRequired field.
func (o *BackchannelAuthenticationResponse) SetUserCodeRequired(v bool) {
	o.UserCodeRequired = &v
}

// GetRequestedExpiry returns the RequestedExpiry field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetRequestedExpiry() int32 {
	if o == nil || IsNil(o.RequestedExpiry) {
		var ret int32
		return ret
	}
	return *o.RequestedExpiry
}

// GetRequestedExpiryOk returns a tuple with the RequestedExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetRequestedExpiryOk() (*int32, bool) {
	if o == nil || IsNil(o.RequestedExpiry) {
		return nil, false
	}
	return o.RequestedExpiry, true
}

// HasRequestedExpiry returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasRequestedExpiry() bool {
	if o != nil && !IsNil(o.RequestedExpiry) {
		return true
	}

	return false
}

// SetRequestedExpiry gets a reference to the given int32 and assigns it to the RequestedExpiry field.
func (o *BackchannelAuthenticationResponse) SetRequestedExpiry(v int32) {
	o.RequestedExpiry = &v
}

// GetRequestContext returns the RequestContext field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetRequestContext() string {
	if o == nil || IsNil(o.RequestContext) {
		var ret string
		return ret
	}
	return *o.RequestContext
}

// GetRequestContextOk returns a tuple with the RequestContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetRequestContextOk() (*string, bool) {
	if o == nil || IsNil(o.RequestContext) {
		return nil, false
	}
	return o.RequestContext, true
}

// HasRequestContext returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasRequestContext() bool {
	if o != nil && !IsNil(o.RequestContext) {
		return true
	}

	return false
}

// SetRequestContext gets a reference to the given string and assigns it to the RequestContext field.
func (o *BackchannelAuthenticationResponse) SetRequestContext(v string) {
	o.RequestContext = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetWarnings() []string {
	if o == nil || IsNil(o.Warnings) {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetWarningsOk() ([]string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *BackchannelAuthenticationResponse) SetWarnings(v []string) {
	o.Warnings = v
}

// GetTicket returns the Ticket field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetTicket() string {
	if o == nil || IsNil(o.Ticket) {
		var ret string
		return ret
	}
	return *o.Ticket
}

// GetTicketOk returns a tuple with the Ticket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetTicketOk() (*string, bool) {
	if o == nil || IsNil(o.Ticket) {
		return nil, false
	}
	return o.Ticket, true
}

// HasTicket returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasTicket() bool {
	if o != nil && !IsNil(o.Ticket) {
		return true
	}

	return false
}

// SetTicket gets a reference to the given string and assigns it to the Ticket field.
func (o *BackchannelAuthenticationResponse) SetTicket(v string) {
	o.Ticket = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetResources() []string {
	if o == nil || IsNil(o.Resources) {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetResourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *BackchannelAuthenticationResponse) SetResources(v []string) {
	o.Resources = v
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetAuthorizationDetails() AuthzDetails {
	if o == nil || IsNil(o.AuthorizationDetails) {
		var ret AuthzDetails
		return ret
	}
	return *o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetAuthorizationDetailsOk() (*AuthzDetails, bool) {
	if o == nil || IsNil(o.AuthorizationDetails) {
		return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasAuthorizationDetails() bool {
	if o != nil && !IsNil(o.AuthorizationDetails) {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given AuthzDetails and assigns it to the AuthorizationDetails field.
func (o *BackchannelAuthenticationResponse) SetAuthorizationDetails(v AuthzDetails) {
	o.AuthorizationDetails = &v
}

// GetServiceAttributes returns the ServiceAttributes field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetServiceAttributes() []Pair {
	if o == nil || IsNil(o.ServiceAttributes) {
		var ret []Pair
		return ret
	}
	return o.ServiceAttributes
}

// GetServiceAttributesOk returns a tuple with the ServiceAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetServiceAttributesOk() ([]Pair, bool) {
	if o == nil || IsNil(o.ServiceAttributes) {
		return nil, false
	}
	return o.ServiceAttributes, true
}

// HasServiceAttributes returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasServiceAttributes() bool {
	if o != nil && !IsNil(o.ServiceAttributes) {
		return true
	}

	return false
}

// SetServiceAttributes gets a reference to the given []Pair and assigns it to the ServiceAttributes field.
func (o *BackchannelAuthenticationResponse) SetServiceAttributes(v []Pair) {
	o.ServiceAttributes = v
}

// GetClientAttributes returns the ClientAttributes field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetClientAttributes() []Pair {
	if o == nil || IsNil(o.ClientAttributes) {
		var ret []Pair
		return ret
	}
	return o.ClientAttributes
}

// GetClientAttributesOk returns a tuple with the ClientAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetClientAttributesOk() ([]Pair, bool) {
	if o == nil || IsNil(o.ClientAttributes) {
		return nil, false
	}
	return o.ClientAttributes, true
}

// HasClientAttributes returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasClientAttributes() bool {
	if o != nil && !IsNil(o.ClientAttributes) {
		return true
	}

	return false
}

// SetClientAttributes gets a reference to the given []Pair and assigns it to the ClientAttributes field.
func (o *BackchannelAuthenticationResponse) SetClientAttributes(v []Pair) {
	o.ClientAttributes = v
}

// GetDynamicScopes returns the DynamicScopes field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetDynamicScopes() []DynamicScope {
	if o == nil || IsNil(o.DynamicScopes) {
		var ret []DynamicScope
		return ret
	}
	return o.DynamicScopes
}

// GetDynamicScopesOk returns a tuple with the DynamicScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetDynamicScopesOk() ([]DynamicScope, bool) {
	if o == nil || IsNil(o.DynamicScopes) {
		return nil, false
	}
	return o.DynamicScopes, true
}

// HasDynamicScopes returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasDynamicScopes() bool {
	if o != nil && !IsNil(o.DynamicScopes) {
		return true
	}

	return false
}

// SetDynamicScopes gets a reference to the given []DynamicScope and assigns it to the DynamicScopes field.
func (o *BackchannelAuthenticationResponse) SetDynamicScopes(v []DynamicScope) {
	o.DynamicScopes = v
}

// GetDeliveryMode returns the DeliveryMode field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetDeliveryMode() DeliveryMode {
	if o == nil || IsNil(o.DeliveryMode) {
		var ret DeliveryMode
		return ret
	}
	return *o.DeliveryMode
}

// GetDeliveryModeOk returns a tuple with the DeliveryMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetDeliveryModeOk() (*DeliveryMode, bool) {
	if o == nil || IsNil(o.DeliveryMode) {
		return nil, false
	}
	return o.DeliveryMode, true
}

// HasDeliveryMode returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasDeliveryMode() bool {
	if o != nil && !IsNil(o.DeliveryMode) {
		return true
	}

	return false
}

// SetDeliveryMode gets a reference to the given DeliveryMode and assigns it to the DeliveryMode field.
func (o *BackchannelAuthenticationResponse) SetDeliveryMode(v DeliveryMode) {
	o.DeliveryMode = &v
}

// GetClientAuthMethod returns the ClientAuthMethod field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetClientAuthMethod() string {
	if o == nil || IsNil(o.ClientAuthMethod) {
		var ret string
		return ret
	}
	return *o.ClientAuthMethod
}

// GetClientAuthMethodOk returns a tuple with the ClientAuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetClientAuthMethodOk() (*string, bool) {
	if o == nil || IsNil(o.ClientAuthMethod) {
		return nil, false
	}
	return o.ClientAuthMethod, true
}

// HasClientAuthMethod returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasClientAuthMethod() bool {
	if o != nil && !IsNil(o.ClientAuthMethod) {
		return true
	}

	return false
}

// SetClientAuthMethod gets a reference to the given string and assigns it to the ClientAuthMethod field.
func (o *BackchannelAuthenticationResponse) SetClientAuthMethod(v string) {
	o.ClientAuthMethod = &v
}

// GetGmAction returns the GmAction field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetGmAction() GrantManagementAction {
	if o == nil || IsNil(o.GmAction) {
		var ret GrantManagementAction
		return ret
	}
	return *o.GmAction
}

// GetGmActionOk returns a tuple with the GmAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetGmActionOk() (*GrantManagementAction, bool) {
	if o == nil || IsNil(o.GmAction) {
		return nil, false
	}
	return o.GmAction, true
}

// HasGmAction returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasGmAction() bool {
	if o != nil && !IsNil(o.GmAction) {
		return true
	}

	return false
}

// SetGmAction gets a reference to the given GrantManagementAction and assigns it to the GmAction field.
func (o *BackchannelAuthenticationResponse) SetGmAction(v GrantManagementAction) {
	o.GmAction = &v
}

// GetGrantId returns the GrantId field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetGrantId() string {
	if o == nil || IsNil(o.GrantId) {
		var ret string
		return ret
	}
	return *o.GrantId
}

// GetGrantIdOk returns a tuple with the GrantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetGrantIdOk() (*string, bool) {
	if o == nil || IsNil(o.GrantId) {
		return nil, false
	}
	return o.GrantId, true
}

// HasGrantId returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasGrantId() bool {
	if o != nil && !IsNil(o.GrantId) {
		return true
	}

	return false
}

// SetGrantId gets a reference to the given string and assigns it to the GrantId field.
func (o *BackchannelAuthenticationResponse) SetGrantId(v string) {
	o.GrantId = &v
}

// GetGrant returns the Grant field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetGrant() Grant {
	if o == nil || IsNil(o.Grant) {
		var ret Grant
		return ret
	}
	return *o.Grant
}

// GetGrantOk returns a tuple with the Grant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetGrantOk() (*Grant, bool) {
	if o == nil || IsNil(o.Grant) {
		return nil, false
	}
	return o.Grant, true
}

// HasGrant returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasGrant() bool {
	if o != nil && !IsNil(o.Grant) {
		return true
	}

	return false
}

// SetGrant gets a reference to the given Grant and assigns it to the Grant field.
func (o *BackchannelAuthenticationResponse) SetGrant(v Grant) {
	o.Grant = &v
}

// GetGrantSubject returns the GrantSubject field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetGrantSubject() string {
	if o == nil || IsNil(o.GrantSubject) {
		var ret string
		return ret
	}
	return *o.GrantSubject
}

// GetGrantSubjectOk returns a tuple with the GrantSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetGrantSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.GrantSubject) {
		return nil, false
	}
	return o.GrantSubject, true
}

// HasGrantSubject returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasGrantSubject() bool {
	if o != nil && !IsNil(o.GrantSubject) {
		return true
	}

	return false
}

// SetGrantSubject gets a reference to the given string and assigns it to the GrantSubject field.
func (o *BackchannelAuthenticationResponse) SetGrantSubject(v string) {
	o.GrantSubject = &v
}

// GetClientEntityId returns the ClientEntityId field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetClientEntityId() string {
	if o == nil || IsNil(o.ClientEntityId) {
		var ret string
		return ret
	}
	return *o.ClientEntityId
}

// GetClientEntityIdOk returns a tuple with the ClientEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetClientEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientEntityId) {
		return nil, false
	}
	return o.ClientEntityId, true
}

// HasClientEntityId returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasClientEntityId() bool {
	if o != nil && !IsNil(o.ClientEntityId) {
		return true
	}

	return false
}

// SetClientEntityId gets a reference to the given string and assigns it to the ClientEntityId field.
func (o *BackchannelAuthenticationResponse) SetClientEntityId(v string) {
	o.ClientEntityId = &v
}

// GetClientEntityIdUsed returns the ClientEntityIdUsed field value if set, zero value otherwise.
func (o *BackchannelAuthenticationResponse) GetClientEntityIdUsed() bool {
	if o == nil || IsNil(o.ClientEntityIdUsed) {
		var ret bool
		return ret
	}
	return *o.ClientEntityIdUsed
}

// GetClientEntityIdUsedOk returns a tuple with the ClientEntityIdUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationResponse) GetClientEntityIdUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientEntityIdUsed) {
		return nil, false
	}
	return o.ClientEntityIdUsed, true
}

// HasClientEntityIdUsed returns a boolean if a field has been set.
func (o *BackchannelAuthenticationResponse) HasClientEntityIdUsed() bool {
	if o != nil && !IsNil(o.ClientEntityIdUsed) {
		return true
	}

	return false
}

// SetClientEntityIdUsed gets a reference to the given bool and assigns it to the ClientEntityIdUsed field.
func (o *BackchannelAuthenticationResponse) SetClientEntityIdUsed(v bool) {
	o.ClientEntityIdUsed = &v
}

func (o BackchannelAuthenticationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackchannelAuthenticationResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ResponseContent) {
		toSerialize["responseContent"] = o.ResponseContent
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.ClientIdAlias) {
		toSerialize["clientIdAlias"] = o.ClientIdAlias
	}
	if !IsNil(o.ClientIdAliasUsed) {
		toSerialize["clientIdAliasUsed"] = o.ClientIdAliasUsed
	}
	if !IsNil(o.ClientName) {
		toSerialize["clientName"] = o.ClientName
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.ClaimNames) {
		toSerialize["claimNames"] = o.ClaimNames
	}
	if !IsNil(o.ClientNotificationToken) {
		toSerialize["clientNotificationToken"] = o.ClientNotificationToken
	}
	if !IsNil(o.Acrs) {
		toSerialize["acrs"] = o.Acrs
	}
	if !IsNil(o.HintType) {
		toSerialize["hintType"] = o.HintType
	}
	if !IsNil(o.Hint) {
		toSerialize["hint"] = o.Hint
	}
	if !IsNil(o.Sub) {
		toSerialize["sub"] = o.Sub
	}
	if !IsNil(o.BindingMessage) {
		toSerialize["bindingMessage"] = o.BindingMessage
	}
	if !IsNil(o.UserCode) {
		toSerialize["userCode"] = o.UserCode
	}
	if !IsNil(o.UserCodeRequired) {
		toSerialize["userCodeRequired"] = o.UserCodeRequired
	}
	if !IsNil(o.RequestedExpiry) {
		toSerialize["requestedExpiry"] = o.RequestedExpiry
	}
	if !IsNil(o.RequestContext) {
		toSerialize["requestContext"] = o.RequestContext
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.Ticket) {
		toSerialize["ticket"] = o.Ticket
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.AuthorizationDetails) {
		toSerialize["authorizationDetails"] = o.AuthorizationDetails
	}
	if !IsNil(o.ServiceAttributes) {
		toSerialize["serviceAttributes"] = o.ServiceAttributes
	}
	if !IsNil(o.ClientAttributes) {
		toSerialize["clientAttributes"] = o.ClientAttributes
	}
	if !IsNil(o.DynamicScopes) {
		toSerialize["dynamicScopes"] = o.DynamicScopes
	}
	if !IsNil(o.DeliveryMode) {
		toSerialize["deliveryMode"] = o.DeliveryMode
	}
	if !IsNil(o.ClientAuthMethod) {
		toSerialize["clientAuthMethod"] = o.ClientAuthMethod
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
	if !IsNil(o.ClientEntityId) {
		toSerialize["clientEntityId"] = o.ClientEntityId
	}
	if !IsNil(o.ClientEntityIdUsed) {
		toSerialize["clientEntityIdUsed"] = o.ClientEntityIdUsed
	}
	return toSerialize, nil
}

type NullableBackchannelAuthenticationResponse struct {
	value *BackchannelAuthenticationResponse
	isSet bool
}

func (v NullableBackchannelAuthenticationResponse) Get() *BackchannelAuthenticationResponse {
	return v.value
}

func (v *NullableBackchannelAuthenticationResponse) Set(val *BackchannelAuthenticationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBackchannelAuthenticationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBackchannelAuthenticationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackchannelAuthenticationResponse(val *BackchannelAuthenticationResponse) *NullableBackchannelAuthenticationResponse {
	return &NullableBackchannelAuthenticationResponse{value: val, isSet: true}
}

func (v NullableBackchannelAuthenticationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackchannelAuthenticationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
