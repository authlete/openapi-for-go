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

// checks if the DeviceVerificationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceVerificationResponse{}

// DeviceVerificationResponse struct for DeviceVerificationResponse
type DeviceVerificationResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take.
	Action *string `json:"action,omitempty"`
	// The client ID of the client application to which the user code has been issued.
	ClientId *int64 `json:"clientId,omitempty"`
	// The client ID alias of the client application to which the user code has been issued.
	ClientIdAlias *string `json:"clientIdAlias,omitempty"`
	// `true` if the value of the `client_id` request parameter included in the device authorization request is the client ID alias. `false` if the value is the original numeric client ID.
	ClientIdAliasUsed *bool `json:"clientIdAliasUsed,omitempty"`
	// The name of the client application to which the user code has been issued.
	ClientName *string `json:"clientName,omitempty"`
	// The scopes requested by the device authorization request.  Note that `description` property and `descriptions` property of each scope object in the array contained in this property is always null even if descriptions of the scopes are registered.
	Scopes []Scope `json:"scopes,omitempty"`
	// The names of the claims which were requested indirectly via some special scopes. See [5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims) in OpenID Connect Core 1.0 for details.  This property is always `null` if the `scope` request parameter of the device authorization request does not include the `openid` scope even if special scopes (such as `profile`) are included in the request (unless the openid scope is included in the default set of scopes which is used when the `scope` request parameter is omitted).
	ClaimNames []string `json:"claimNames,omitempty"`
	// The list of ACR values requested by the device authorization request.
	Acrs []string `json:"acrs,omitempty"`
	// The resources specified by the `resource` request parameters or by the `resource` property in the request object. If both are given, the values in the request object should be set. See \"Resource Indicators for OAuth 2.0\" for details.
	Resources            []string      `json:"resources,omitempty"`
	AuthorizationDetails *AuthzDetails `json:"authorizationDetails,omitempty"`
	// The attributes of this service that the client application belongs to.
	ServiceAttributes []Pair `json:"serviceAttributes,omitempty"`
	// The attributes of the client.
	ClientAttributes []Pair `json:"clientAttributes,omitempty"`
	// The dynamic scopes which the client application requested by the scope request parameter.
	DynamicScopes []DynamicScope `json:"dynamicScopes,omitempty"`
	// Get the date in milliseconds since the Unix epoch (1970-01-01) at which the user code will expire.
	ExpiresAt *int64                 `json:"expiresAt,omitempty"`
	GmAction  *GrantManagementAction `json:"gmAction,omitempty"`
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

// NewDeviceVerificationResponse instantiates a new DeviceVerificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceVerificationResponse() *DeviceVerificationResponse {
	this := DeviceVerificationResponse{}
	return &this
}

// NewDeviceVerificationResponseWithDefaults instantiates a new DeviceVerificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceVerificationResponseWithDefaults() *DeviceVerificationResponse {
	this := DeviceVerificationResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetResultCode() string {
	if o == nil || IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *DeviceVerificationResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetResultMessage() string {
	if o == nil || IsNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasResultMessage() bool {
	if o != nil && !IsNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *DeviceVerificationResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *DeviceVerificationResponse) SetAction(v string) {
	o.Action = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetClientId() int64 {
	if o == nil || IsNil(o.ClientId) {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetClientIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int64 and assigns it to the ClientId field.
func (o *DeviceVerificationResponse) SetClientId(v int64) {
	o.ClientId = &v
}

// GetClientIdAlias returns the ClientIdAlias field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetClientIdAlias() string {
	if o == nil || IsNil(o.ClientIdAlias) {
		var ret string
		return ret
	}
	return *o.ClientIdAlias
}

// GetClientIdAliasOk returns a tuple with the ClientIdAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetClientIdAliasOk() (*string, bool) {
	if o == nil || IsNil(o.ClientIdAlias) {
		return nil, false
	}
	return o.ClientIdAlias, true
}

// HasClientIdAlias returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasClientIdAlias() bool {
	if o != nil && !IsNil(o.ClientIdAlias) {
		return true
	}

	return false
}

// SetClientIdAlias gets a reference to the given string and assigns it to the ClientIdAlias field.
func (o *DeviceVerificationResponse) SetClientIdAlias(v string) {
	o.ClientIdAlias = &v
}

// GetClientIdAliasUsed returns the ClientIdAliasUsed field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetClientIdAliasUsed() bool {
	if o == nil || IsNil(o.ClientIdAliasUsed) {
		var ret bool
		return ret
	}
	return *o.ClientIdAliasUsed
}

// GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetClientIdAliasUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientIdAliasUsed) {
		return nil, false
	}
	return o.ClientIdAliasUsed, true
}

// HasClientIdAliasUsed returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasClientIdAliasUsed() bool {
	if o != nil && !IsNil(o.ClientIdAliasUsed) {
		return true
	}

	return false
}

// SetClientIdAliasUsed gets a reference to the given bool and assigns it to the ClientIdAliasUsed field.
func (o *DeviceVerificationResponse) SetClientIdAliasUsed(v bool) {
	o.ClientIdAliasUsed = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetClientName() string {
	if o == nil || IsNil(o.ClientName) {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetClientNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClientName) {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasClientName() bool {
	if o != nil && !IsNil(o.ClientName) {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *DeviceVerificationResponse) SetClientName(v string) {
	o.ClientName = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetScopes() []Scope {
	if o == nil || IsNil(o.Scopes) {
		var ret []Scope
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetScopesOk() ([]Scope, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []Scope and assigns it to the Scopes field.
func (o *DeviceVerificationResponse) SetScopes(v []Scope) {
	o.Scopes = v
}

// GetClaimNames returns the ClaimNames field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetClaimNames() []string {
	if o == nil || IsNil(o.ClaimNames) {
		var ret []string
		return ret
	}
	return o.ClaimNames
}

// GetClaimNamesOk returns a tuple with the ClaimNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetClaimNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.ClaimNames) {
		return nil, false
	}
	return o.ClaimNames, true
}

// HasClaimNames returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasClaimNames() bool {
	if o != nil && !IsNil(o.ClaimNames) {
		return true
	}

	return false
}

// SetClaimNames gets a reference to the given []string and assigns it to the ClaimNames field.
func (o *DeviceVerificationResponse) SetClaimNames(v []string) {
	o.ClaimNames = v
}

// GetAcrs returns the Acrs field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetAcrs() []string {
	if o == nil || IsNil(o.Acrs) {
		var ret []string
		return ret
	}
	return o.Acrs
}

// GetAcrsOk returns a tuple with the Acrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetAcrsOk() ([]string, bool) {
	if o == nil || IsNil(o.Acrs) {
		return nil, false
	}
	return o.Acrs, true
}

// HasAcrs returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasAcrs() bool {
	if o != nil && !IsNil(o.Acrs) {
		return true
	}

	return false
}

// SetAcrs gets a reference to the given []string and assigns it to the Acrs field.
func (o *DeviceVerificationResponse) SetAcrs(v []string) {
	o.Acrs = v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetResources() []string {
	if o == nil || IsNil(o.Resources) {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetResourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *DeviceVerificationResponse) SetResources(v []string) {
	o.Resources = v
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetAuthorizationDetails() AuthzDetails {
	if o == nil || IsNil(o.AuthorizationDetails) {
		var ret AuthzDetails
		return ret
	}
	return *o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetAuthorizationDetailsOk() (*AuthzDetails, bool) {
	if o == nil || IsNil(o.AuthorizationDetails) {
		return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasAuthorizationDetails() bool {
	if o != nil && !IsNil(o.AuthorizationDetails) {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given AuthzDetails and assigns it to the AuthorizationDetails field.
func (o *DeviceVerificationResponse) SetAuthorizationDetails(v AuthzDetails) {
	o.AuthorizationDetails = &v
}

// GetServiceAttributes returns the ServiceAttributes field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetServiceAttributes() []Pair {
	if o == nil || IsNil(o.ServiceAttributes) {
		var ret []Pair
		return ret
	}
	return o.ServiceAttributes
}

// GetServiceAttributesOk returns a tuple with the ServiceAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetServiceAttributesOk() ([]Pair, bool) {
	if o == nil || IsNil(o.ServiceAttributes) {
		return nil, false
	}
	return o.ServiceAttributes, true
}

// HasServiceAttributes returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasServiceAttributes() bool {
	if o != nil && !IsNil(o.ServiceAttributes) {
		return true
	}

	return false
}

// SetServiceAttributes gets a reference to the given []Pair and assigns it to the ServiceAttributes field.
func (o *DeviceVerificationResponse) SetServiceAttributes(v []Pair) {
	o.ServiceAttributes = v
}

// GetClientAttributes returns the ClientAttributes field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetClientAttributes() []Pair {
	if o == nil || IsNil(o.ClientAttributes) {
		var ret []Pair
		return ret
	}
	return o.ClientAttributes
}

// GetClientAttributesOk returns a tuple with the ClientAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetClientAttributesOk() ([]Pair, bool) {
	if o == nil || IsNil(o.ClientAttributes) {
		return nil, false
	}
	return o.ClientAttributes, true
}

// HasClientAttributes returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasClientAttributes() bool {
	if o != nil && !IsNil(o.ClientAttributes) {
		return true
	}

	return false
}

// SetClientAttributes gets a reference to the given []Pair and assigns it to the ClientAttributes field.
func (o *DeviceVerificationResponse) SetClientAttributes(v []Pair) {
	o.ClientAttributes = v
}

// GetDynamicScopes returns the DynamicScopes field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetDynamicScopes() []DynamicScope {
	if o == nil || IsNil(o.DynamicScopes) {
		var ret []DynamicScope
		return ret
	}
	return o.DynamicScopes
}

// GetDynamicScopesOk returns a tuple with the DynamicScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetDynamicScopesOk() ([]DynamicScope, bool) {
	if o == nil || IsNil(o.DynamicScopes) {
		return nil, false
	}
	return o.DynamicScopes, true
}

// HasDynamicScopes returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasDynamicScopes() bool {
	if o != nil && !IsNil(o.DynamicScopes) {
		return true
	}

	return false
}

// SetDynamicScopes gets a reference to the given []DynamicScope and assigns it to the DynamicScopes field.
func (o *DeviceVerificationResponse) SetDynamicScopes(v []DynamicScope) {
	o.DynamicScopes = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *DeviceVerificationResponse) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

// GetGmAction returns the GmAction field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetGmAction() GrantManagementAction {
	if o == nil || IsNil(o.GmAction) {
		var ret GrantManagementAction
		return ret
	}
	return *o.GmAction
}

// GetGmActionOk returns a tuple with the GmAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetGmActionOk() (*GrantManagementAction, bool) {
	if o == nil || IsNil(o.GmAction) {
		return nil, false
	}
	return o.GmAction, true
}

// HasGmAction returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasGmAction() bool {
	if o != nil && !IsNil(o.GmAction) {
		return true
	}

	return false
}

// SetGmAction gets a reference to the given GrantManagementAction and assigns it to the GmAction field.
func (o *DeviceVerificationResponse) SetGmAction(v GrantManagementAction) {
	o.GmAction = &v
}

// GetGrantId returns the GrantId field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetGrantId() string {
	if o == nil || IsNil(o.GrantId) {
		var ret string
		return ret
	}
	return *o.GrantId
}

// GetGrantIdOk returns a tuple with the GrantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetGrantIdOk() (*string, bool) {
	if o == nil || IsNil(o.GrantId) {
		return nil, false
	}
	return o.GrantId, true
}

// HasGrantId returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasGrantId() bool {
	if o != nil && !IsNil(o.GrantId) {
		return true
	}

	return false
}

// SetGrantId gets a reference to the given string and assigns it to the GrantId field.
func (o *DeviceVerificationResponse) SetGrantId(v string) {
	o.GrantId = &v
}

// GetGrant returns the Grant field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetGrant() Grant {
	if o == nil || IsNil(o.Grant) {
		var ret Grant
		return ret
	}
	return *o.Grant
}

// GetGrantOk returns a tuple with the Grant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetGrantOk() (*Grant, bool) {
	if o == nil || IsNil(o.Grant) {
		return nil, false
	}
	return o.Grant, true
}

// HasGrant returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasGrant() bool {
	if o != nil && !IsNil(o.Grant) {
		return true
	}

	return false
}

// SetGrant gets a reference to the given Grant and assigns it to the Grant field.
func (o *DeviceVerificationResponse) SetGrant(v Grant) {
	o.Grant = &v
}

// GetGrantSubject returns the GrantSubject field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetGrantSubject() string {
	if o == nil || IsNil(o.GrantSubject) {
		var ret string
		return ret
	}
	return *o.GrantSubject
}

// GetGrantSubjectOk returns a tuple with the GrantSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetGrantSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.GrantSubject) {
		return nil, false
	}
	return o.GrantSubject, true
}

// HasGrantSubject returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasGrantSubject() bool {
	if o != nil && !IsNil(o.GrantSubject) {
		return true
	}

	return false
}

// SetGrantSubject gets a reference to the given string and assigns it to the GrantSubject field.
func (o *DeviceVerificationResponse) SetGrantSubject(v string) {
	o.GrantSubject = &v
}

// GetClientEntityId returns the ClientEntityId field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetClientEntityId() string {
	if o == nil || IsNil(o.ClientEntityId) {
		var ret string
		return ret
	}
	return *o.ClientEntityId
}

// GetClientEntityIdOk returns a tuple with the ClientEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetClientEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientEntityId) {
		return nil, false
	}
	return o.ClientEntityId, true
}

// HasClientEntityId returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasClientEntityId() bool {
	if o != nil && !IsNil(o.ClientEntityId) {
		return true
	}

	return false
}

// SetClientEntityId gets a reference to the given string and assigns it to the ClientEntityId field.
func (o *DeviceVerificationResponse) SetClientEntityId(v string) {
	o.ClientEntityId = &v
}

// GetClientEntityIdUsed returns the ClientEntityIdUsed field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetClientEntityIdUsed() bool {
	if o == nil || IsNil(o.ClientEntityIdUsed) {
		var ret bool
		return ret
	}
	return *o.ClientEntityIdUsed
}

// GetClientEntityIdUsedOk returns a tuple with the ClientEntityIdUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetClientEntityIdUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientEntityIdUsed) {
		return nil, false
	}
	return o.ClientEntityIdUsed, true
}

// HasClientEntityIdUsed returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasClientEntityIdUsed() bool {
	if o != nil && !IsNil(o.ClientEntityIdUsed) {
		return true
	}

	return false
}

// SetClientEntityIdUsed gets a reference to the given bool and assigns it to the ClientEntityIdUsed field.
func (o *DeviceVerificationResponse) SetClientEntityIdUsed(v bool) {
	o.ClientEntityIdUsed = &v
}

func (o DeviceVerificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceVerificationResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Acrs) {
		toSerialize["acrs"] = o.Acrs
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
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
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

type NullableDeviceVerificationResponse struct {
	value *DeviceVerificationResponse
	isSet bool
}

func (v NullableDeviceVerificationResponse) Get() *DeviceVerificationResponse {
	return v.value
}

func (v *NullableDeviceVerificationResponse) Set(val *DeviceVerificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceVerificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceVerificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceVerificationResponse(val *DeviceVerificationResponse) *NullableDeviceVerificationResponse {
	return &NullableDeviceVerificationResponse{value: val, isSet: true}
}

func (v NullableDeviceVerificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceVerificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
