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

// checks if the DeviceAuthorizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAuthorizationResponse{}

// DeviceAuthorizationResponse struct for DeviceAuthorizationResponse
type DeviceAuthorizationResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take.
	Action *string `json:"action,omitempty"`
	// The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of `action` parameter.
	ResponseContent *string `json:"responseContent,omitempty"`
	// The client ID of the client application that has made the device authorization request.
	ClientId *int64 `json:"clientId,omitempty"`
	// The client ID alias of the client application that has made the device authorization request.
	ClientIdAlias *string `json:"clientIdAlias,omitempty"`
	// `true` if the value of the client_id request parameter included in the device authorization request is the client ID alias. `false` if the value is the original numeric client ID.
	ClientIdAliasUsed *bool `json:"clientIdAliasUsed,omitempty"`
	// The name of the client application which has made the device authorization request.
	ClientName *string `json:"clientName,omitempty"`
	// The client authentication method that should be performed at the device authorization endpoint.
	ClientAuthMethod *string `json:"clientAuthMethod,omitempty"`
	// The scopes requested by the device authorization request.  Basically, this property holds the value of the scope request parameter in the device authorization request. However, because unregistered scopes are dropped on Authlete side, if the `scope` request parameter contains unknown scopes, the list returned by this property becomes different from the value of the `scope` request parameter.  Note that `description` property and `descriptions` property of each scope object in the array contained in this property is always `null` even if descriptions of the scopes are registered.
	Scopes []Scope `json:"scopes,omitempty"`
	// The names of the claims which were requested indirectly via some special scopes. See [5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims) in OpenID Connect Core 1.0 for details.
	ClaimNames []string `json:"claimNames,omitempty"`
	// The list of ACR values requested by the device authorization request.  Basically, this property holds the value of the `acr_values` request parameter in the device authorization request. However, because unsupported ACR values are dropped on Authlete side, if the `acr_values` request parameter contains unrecognized ACR values, the list returned by this property becomes different from the value of the `acr_values` request parameter.
	Acrs []string `json:"acrs,omitempty"`
	// The device verification code. This corresponds to the `device_code` property in the response to the client.
	DeviceCode *string `json:"deviceCode,omitempty"`
	// The end-user verification code. This corresponds to the `user_code` property in the response to the client.
	UserCode *string `json:"userCode,omitempty"`
	// The end-user verification URI. This corresponds to the `verification_uri` property in the response to the client.
	VerificationUri *string `json:"verificationUri,omitempty"`
	// The end-user verification URI that includes the end-user verification code. This corresponds to the `verification_uri_complete` property in the response to the client.
	VerificationUriComplete *string `json:"verificationUriComplete,omitempty"`
	// The duration of the device verification code in seconds. This corresponds to the `expires_in` property in the response to the client.
	ExpiresIn *int32 `json:"expiresIn,omitempty"`
	// The minimum amount of time in seconds that the client must wait for between polling requests to the token endpoint. This corresponds to the `interval` property in the response to the client.
	Interval *int32 `json:"interval,omitempty"`
	// The warnings raised during processing the backchannel authentication request.
	Warnings []string `json:"warnings,omitempty"`
	// The resources specified by the `resource` request parameters. See \"Resource Indicators for OAuth 2.0\" for details.
	Resources            []string      `json:"resources,omitempty"`
	AuthorizationDetails *AuthzDetails `json:"authorizationDetails,omitempty"`
	// The attributes of this service that the client application belongs to.
	ServiceAttributes []Pair `json:"serviceAttributes,omitempty"`
	// The attributes of the client.
	ClientAttributes []Pair `json:"clientAttributes,omitempty"`
	// The dynamic scopes which the client application requested by the scope request parameter.
	DynamicScopes []DynamicScope         `json:"dynamicScopes,omitempty"`
	GmAction      *GrantManagementAction `json:"gmAction,omitempty"`
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

// NewDeviceAuthorizationResponse instantiates a new DeviceAuthorizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthorizationResponse() *DeviceAuthorizationResponse {
	this := DeviceAuthorizationResponse{}
	return &this
}

// NewDeviceAuthorizationResponseWithDefaults instantiates a new DeviceAuthorizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthorizationResponseWithDefaults() *DeviceAuthorizationResponse {
	this := DeviceAuthorizationResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetResultCode() string {
	if o == nil || IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *DeviceAuthorizationResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetResultMessage() string {
	if o == nil || IsNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasResultMessage() bool {
	if o != nil && !IsNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *DeviceAuthorizationResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *DeviceAuthorizationResponse) SetAction(v string) {
	o.Action = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetResponseContent() string {
	if o == nil || IsNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasResponseContent() bool {
	if o != nil && !IsNil(o.ResponseContent) {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *DeviceAuthorizationResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetClientId() int64 {
	if o == nil || IsNil(o.ClientId) {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetClientIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int64 and assigns it to the ClientId field.
func (o *DeviceAuthorizationResponse) SetClientId(v int64) {
	o.ClientId = &v
}

// GetClientIdAlias returns the ClientIdAlias field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetClientIdAlias() string {
	if o == nil || IsNil(o.ClientIdAlias) {
		var ret string
		return ret
	}
	return *o.ClientIdAlias
}

// GetClientIdAliasOk returns a tuple with the ClientIdAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetClientIdAliasOk() (*string, bool) {
	if o == nil || IsNil(o.ClientIdAlias) {
		return nil, false
	}
	return o.ClientIdAlias, true
}

// HasClientIdAlias returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasClientIdAlias() bool {
	if o != nil && !IsNil(o.ClientIdAlias) {
		return true
	}

	return false
}

// SetClientIdAlias gets a reference to the given string and assigns it to the ClientIdAlias field.
func (o *DeviceAuthorizationResponse) SetClientIdAlias(v string) {
	o.ClientIdAlias = &v
}

// GetClientIdAliasUsed returns the ClientIdAliasUsed field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetClientIdAliasUsed() bool {
	if o == nil || IsNil(o.ClientIdAliasUsed) {
		var ret bool
		return ret
	}
	return *o.ClientIdAliasUsed
}

// GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetClientIdAliasUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientIdAliasUsed) {
		return nil, false
	}
	return o.ClientIdAliasUsed, true
}

// HasClientIdAliasUsed returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasClientIdAliasUsed() bool {
	if o != nil && !IsNil(o.ClientIdAliasUsed) {
		return true
	}

	return false
}

// SetClientIdAliasUsed gets a reference to the given bool and assigns it to the ClientIdAliasUsed field.
func (o *DeviceAuthorizationResponse) SetClientIdAliasUsed(v bool) {
	o.ClientIdAliasUsed = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetClientName() string {
	if o == nil || IsNil(o.ClientName) {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetClientNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClientName) {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasClientName() bool {
	if o != nil && !IsNil(o.ClientName) {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *DeviceAuthorizationResponse) SetClientName(v string) {
	o.ClientName = &v
}

// GetClientAuthMethod returns the ClientAuthMethod field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetClientAuthMethod() string {
	if o == nil || IsNil(o.ClientAuthMethod) {
		var ret string
		return ret
	}
	return *o.ClientAuthMethod
}

// GetClientAuthMethodOk returns a tuple with the ClientAuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetClientAuthMethodOk() (*string, bool) {
	if o == nil || IsNil(o.ClientAuthMethod) {
		return nil, false
	}
	return o.ClientAuthMethod, true
}

// HasClientAuthMethod returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasClientAuthMethod() bool {
	if o != nil && !IsNil(o.ClientAuthMethod) {
		return true
	}

	return false
}

// SetClientAuthMethod gets a reference to the given string and assigns it to the ClientAuthMethod field.
func (o *DeviceAuthorizationResponse) SetClientAuthMethod(v string) {
	o.ClientAuthMethod = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetScopes() []Scope {
	if o == nil || IsNil(o.Scopes) {
		var ret []Scope
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetScopesOk() ([]Scope, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []Scope and assigns it to the Scopes field.
func (o *DeviceAuthorizationResponse) SetScopes(v []Scope) {
	o.Scopes = v
}

// GetClaimNames returns the ClaimNames field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetClaimNames() []string {
	if o == nil || IsNil(o.ClaimNames) {
		var ret []string
		return ret
	}
	return o.ClaimNames
}

// GetClaimNamesOk returns a tuple with the ClaimNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetClaimNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.ClaimNames) {
		return nil, false
	}
	return o.ClaimNames, true
}

// HasClaimNames returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasClaimNames() bool {
	if o != nil && !IsNil(o.ClaimNames) {
		return true
	}

	return false
}

// SetClaimNames gets a reference to the given []string and assigns it to the ClaimNames field.
func (o *DeviceAuthorizationResponse) SetClaimNames(v []string) {
	o.ClaimNames = v
}

// GetAcrs returns the Acrs field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetAcrs() []string {
	if o == nil || IsNil(o.Acrs) {
		var ret []string
		return ret
	}
	return o.Acrs
}

// GetAcrsOk returns a tuple with the Acrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetAcrsOk() ([]string, bool) {
	if o == nil || IsNil(o.Acrs) {
		return nil, false
	}
	return o.Acrs, true
}

// HasAcrs returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasAcrs() bool {
	if o != nil && !IsNil(o.Acrs) {
		return true
	}

	return false
}

// SetAcrs gets a reference to the given []string and assigns it to the Acrs field.
func (o *DeviceAuthorizationResponse) SetAcrs(v []string) {
	o.Acrs = v
}

// GetDeviceCode returns the DeviceCode field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetDeviceCode() string {
	if o == nil || IsNil(o.DeviceCode) {
		var ret string
		return ret
	}
	return *o.DeviceCode
}

// GetDeviceCodeOk returns a tuple with the DeviceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetDeviceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceCode) {
		return nil, false
	}
	return o.DeviceCode, true
}

// HasDeviceCode returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasDeviceCode() bool {
	if o != nil && !IsNil(o.DeviceCode) {
		return true
	}

	return false
}

// SetDeviceCode gets a reference to the given string and assigns it to the DeviceCode field.
func (o *DeviceAuthorizationResponse) SetDeviceCode(v string) {
	o.DeviceCode = &v
}

// GetUserCode returns the UserCode field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetUserCode() string {
	if o == nil || IsNil(o.UserCode) {
		var ret string
		return ret
	}
	return *o.UserCode
}

// GetUserCodeOk returns a tuple with the UserCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetUserCodeOk() (*string, bool) {
	if o == nil || IsNil(o.UserCode) {
		return nil, false
	}
	return o.UserCode, true
}

// HasUserCode returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasUserCode() bool {
	if o != nil && !IsNil(o.UserCode) {
		return true
	}

	return false
}

// SetUserCode gets a reference to the given string and assigns it to the UserCode field.
func (o *DeviceAuthorizationResponse) SetUserCode(v string) {
	o.UserCode = &v
}

// GetVerificationUri returns the VerificationUri field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetVerificationUri() string {
	if o == nil || IsNil(o.VerificationUri) {
		var ret string
		return ret
	}
	return *o.VerificationUri
}

// GetVerificationUriOk returns a tuple with the VerificationUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetVerificationUriOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationUri) {
		return nil, false
	}
	return o.VerificationUri, true
}

// HasVerificationUri returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasVerificationUri() bool {
	if o != nil && !IsNil(o.VerificationUri) {
		return true
	}

	return false
}

// SetVerificationUri gets a reference to the given string and assigns it to the VerificationUri field.
func (o *DeviceAuthorizationResponse) SetVerificationUri(v string) {
	o.VerificationUri = &v
}

// GetVerificationUriComplete returns the VerificationUriComplete field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetVerificationUriComplete() string {
	if o == nil || IsNil(o.VerificationUriComplete) {
		var ret string
		return ret
	}
	return *o.VerificationUriComplete
}

// GetVerificationUriCompleteOk returns a tuple with the VerificationUriComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetVerificationUriCompleteOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationUriComplete) {
		return nil, false
	}
	return o.VerificationUriComplete, true
}

// HasVerificationUriComplete returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasVerificationUriComplete() bool {
	if o != nil && !IsNil(o.VerificationUriComplete) {
		return true
	}

	return false
}

// SetVerificationUriComplete gets a reference to the given string and assigns it to the VerificationUriComplete field.
func (o *DeviceAuthorizationResponse) SetVerificationUriComplete(v string) {
	o.VerificationUriComplete = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetExpiresIn() int32 {
	if o == nil || IsNil(o.ExpiresIn) {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetExpiresInOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasExpiresIn() bool {
	if o != nil && !IsNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *DeviceAuthorizationResponse) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetInterval() int32 {
	if o == nil || IsNil(o.Interval) {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *DeviceAuthorizationResponse) SetInterval(v int32) {
	o.Interval = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetWarnings() []string {
	if o == nil || IsNil(o.Warnings) {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetWarningsOk() ([]string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *DeviceAuthorizationResponse) SetWarnings(v []string) {
	o.Warnings = v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetResources() []string {
	if o == nil || IsNil(o.Resources) {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetResourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *DeviceAuthorizationResponse) SetResources(v []string) {
	o.Resources = v
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetAuthorizationDetails() AuthzDetails {
	if o == nil || IsNil(o.AuthorizationDetails) {
		var ret AuthzDetails
		return ret
	}
	return *o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetAuthorizationDetailsOk() (*AuthzDetails, bool) {
	if o == nil || IsNil(o.AuthorizationDetails) {
		return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasAuthorizationDetails() bool {
	if o != nil && !IsNil(o.AuthorizationDetails) {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given AuthzDetails and assigns it to the AuthorizationDetails field.
func (o *DeviceAuthorizationResponse) SetAuthorizationDetails(v AuthzDetails) {
	o.AuthorizationDetails = &v
}

// GetServiceAttributes returns the ServiceAttributes field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetServiceAttributes() []Pair {
	if o == nil || IsNil(o.ServiceAttributes) {
		var ret []Pair
		return ret
	}
	return o.ServiceAttributes
}

// GetServiceAttributesOk returns a tuple with the ServiceAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetServiceAttributesOk() ([]Pair, bool) {
	if o == nil || IsNil(o.ServiceAttributes) {
		return nil, false
	}
	return o.ServiceAttributes, true
}

// HasServiceAttributes returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasServiceAttributes() bool {
	if o != nil && !IsNil(o.ServiceAttributes) {
		return true
	}

	return false
}

// SetServiceAttributes gets a reference to the given []Pair and assigns it to the ServiceAttributes field.
func (o *DeviceAuthorizationResponse) SetServiceAttributes(v []Pair) {
	o.ServiceAttributes = v
}

// GetClientAttributes returns the ClientAttributes field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetClientAttributes() []Pair {
	if o == nil || IsNil(o.ClientAttributes) {
		var ret []Pair
		return ret
	}
	return o.ClientAttributes
}

// GetClientAttributesOk returns a tuple with the ClientAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetClientAttributesOk() ([]Pair, bool) {
	if o == nil || IsNil(o.ClientAttributes) {
		return nil, false
	}
	return o.ClientAttributes, true
}

// HasClientAttributes returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasClientAttributes() bool {
	if o != nil && !IsNil(o.ClientAttributes) {
		return true
	}

	return false
}

// SetClientAttributes gets a reference to the given []Pair and assigns it to the ClientAttributes field.
func (o *DeviceAuthorizationResponse) SetClientAttributes(v []Pair) {
	o.ClientAttributes = v
}

// GetDynamicScopes returns the DynamicScopes field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetDynamicScopes() []DynamicScope {
	if o == nil || IsNil(o.DynamicScopes) {
		var ret []DynamicScope
		return ret
	}
	return o.DynamicScopes
}

// GetDynamicScopesOk returns a tuple with the DynamicScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetDynamicScopesOk() ([]DynamicScope, bool) {
	if o == nil || IsNil(o.DynamicScopes) {
		return nil, false
	}
	return o.DynamicScopes, true
}

// HasDynamicScopes returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasDynamicScopes() bool {
	if o != nil && !IsNil(o.DynamicScopes) {
		return true
	}

	return false
}

// SetDynamicScopes gets a reference to the given []DynamicScope and assigns it to the DynamicScopes field.
func (o *DeviceAuthorizationResponse) SetDynamicScopes(v []DynamicScope) {
	o.DynamicScopes = v
}

// GetGmAction returns the GmAction field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetGmAction() GrantManagementAction {
	if o == nil || IsNil(o.GmAction) {
		var ret GrantManagementAction
		return ret
	}
	return *o.GmAction
}

// GetGmActionOk returns a tuple with the GmAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetGmActionOk() (*GrantManagementAction, bool) {
	if o == nil || IsNil(o.GmAction) {
		return nil, false
	}
	return o.GmAction, true
}

// HasGmAction returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasGmAction() bool {
	if o != nil && !IsNil(o.GmAction) {
		return true
	}

	return false
}

// SetGmAction gets a reference to the given GrantManagementAction and assigns it to the GmAction field.
func (o *DeviceAuthorizationResponse) SetGmAction(v GrantManagementAction) {
	o.GmAction = &v
}

// GetGrantId returns the GrantId field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetGrantId() string {
	if o == nil || IsNil(o.GrantId) {
		var ret string
		return ret
	}
	return *o.GrantId
}

// GetGrantIdOk returns a tuple with the GrantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetGrantIdOk() (*string, bool) {
	if o == nil || IsNil(o.GrantId) {
		return nil, false
	}
	return o.GrantId, true
}

// HasGrantId returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasGrantId() bool {
	if o != nil && !IsNil(o.GrantId) {
		return true
	}

	return false
}

// SetGrantId gets a reference to the given string and assigns it to the GrantId field.
func (o *DeviceAuthorizationResponse) SetGrantId(v string) {
	o.GrantId = &v
}

// GetGrant returns the Grant field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetGrant() Grant {
	if o == nil || IsNil(o.Grant) {
		var ret Grant
		return ret
	}
	return *o.Grant
}

// GetGrantOk returns a tuple with the Grant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetGrantOk() (*Grant, bool) {
	if o == nil || IsNil(o.Grant) {
		return nil, false
	}
	return o.Grant, true
}

// HasGrant returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasGrant() bool {
	if o != nil && !IsNil(o.Grant) {
		return true
	}

	return false
}

// SetGrant gets a reference to the given Grant and assigns it to the Grant field.
func (o *DeviceAuthorizationResponse) SetGrant(v Grant) {
	o.Grant = &v
}

// GetGrantSubject returns the GrantSubject field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetGrantSubject() string {
	if o == nil || IsNil(o.GrantSubject) {
		var ret string
		return ret
	}
	return *o.GrantSubject
}

// GetGrantSubjectOk returns a tuple with the GrantSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetGrantSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.GrantSubject) {
		return nil, false
	}
	return o.GrantSubject, true
}

// HasGrantSubject returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasGrantSubject() bool {
	if o != nil && !IsNil(o.GrantSubject) {
		return true
	}

	return false
}

// SetGrantSubject gets a reference to the given string and assigns it to the GrantSubject field.
func (o *DeviceAuthorizationResponse) SetGrantSubject(v string) {
	o.GrantSubject = &v
}

// GetClientEntityId returns the ClientEntityId field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetClientEntityId() string {
	if o == nil || IsNil(o.ClientEntityId) {
		var ret string
		return ret
	}
	return *o.ClientEntityId
}

// GetClientEntityIdOk returns a tuple with the ClientEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetClientEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientEntityId) {
		return nil, false
	}
	return o.ClientEntityId, true
}

// HasClientEntityId returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasClientEntityId() bool {
	if o != nil && !IsNil(o.ClientEntityId) {
		return true
	}

	return false
}

// SetClientEntityId gets a reference to the given string and assigns it to the ClientEntityId field.
func (o *DeviceAuthorizationResponse) SetClientEntityId(v string) {
	o.ClientEntityId = &v
}

// GetClientEntityIdUsed returns the ClientEntityIdUsed field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetClientEntityIdUsed() bool {
	if o == nil || IsNil(o.ClientEntityIdUsed) {
		var ret bool
		return ret
	}
	return *o.ClientEntityIdUsed
}

// GetClientEntityIdUsedOk returns a tuple with the ClientEntityIdUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetClientEntityIdUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientEntityIdUsed) {
		return nil, false
	}
	return o.ClientEntityIdUsed, true
}

// HasClientEntityIdUsed returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasClientEntityIdUsed() bool {
	if o != nil && !IsNil(o.ClientEntityIdUsed) {
		return true
	}

	return false
}

// SetClientEntityIdUsed gets a reference to the given bool and assigns it to the ClientEntityIdUsed field.
func (o *DeviceAuthorizationResponse) SetClientEntityIdUsed(v bool) {
	o.ClientEntityIdUsed = &v
}

func (o DeviceAuthorizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAuthorizationResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ClientAuthMethod) {
		toSerialize["clientAuthMethod"] = o.ClientAuthMethod
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.ClaimNames) {
		toSerialize["claimNames"] = o.ClaimNames
	}
	if !IsNil(o.Acrs) {
		toSerialize["acrs"] = o.Acrs
	}
	if !IsNil(o.DeviceCode) {
		toSerialize["deviceCode"] = o.DeviceCode
	}
	if !IsNil(o.UserCode) {
		toSerialize["userCode"] = o.UserCode
	}
	if !IsNil(o.VerificationUri) {
		toSerialize["verificationUri"] = o.VerificationUri
	}
	if !IsNil(o.VerificationUriComplete) {
		toSerialize["verificationUriComplete"] = o.VerificationUriComplete
	}
	if !IsNil(o.ExpiresIn) {
		toSerialize["expiresIn"] = o.ExpiresIn
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
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

type NullableDeviceAuthorizationResponse struct {
	value *DeviceAuthorizationResponse
	isSet bool
}

func (v NullableDeviceAuthorizationResponse) Get() *DeviceAuthorizationResponse {
	return v.value
}

func (v *NullableDeviceAuthorizationResponse) Set(val *DeviceAuthorizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthorizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthorizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthorizationResponse(val *DeviceAuthorizationResponse) *NullableDeviceAuthorizationResponse {
	return &NullableDeviceAuthorizationResponse{value: val, isSet: true}
}

func (v NullableDeviceAuthorizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthorizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
