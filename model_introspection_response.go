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

// checks if the IntrospectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntrospectionResponse{}

// IntrospectionResponse struct for IntrospectionResponse
type IntrospectionResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take.
	Action *string `json:"action,omitempty"`
	// The content that the authorization server implementation can use as the value of `WWW-Authenticate` header on errors.
	ResponseContent *string `json:"responseContent,omitempty"`
	// The client ID.
	ClientId *int64 `json:"clientId,omitempty"`
	// The client ID alias when the token request was made. If the client did not have an alias, this parameter is `null`. Also, if the token request was invalid and it failed to identify a client, this parameter is `null`.
	ClientIdAlias *string `json:"clientIdAlias,omitempty"`
	// The flag which indicates whether the client ID alias was used when the token request was made. `true` if the client ID alias was used when the token request was made.
	ClientIdAliasUsed *bool `json:"clientIdAliasUsed,omitempty"`
	// The time at which the access token expires. The value is represented in milliseconds since the Unix epoch (1970-01-01).
	ExpiresAt *int64 `json:"expiresAt,omitempty"`
	// The subject who is associated with the access token. The value of this property is `null` if the access token was issued using the flow of [Client Credentials Grant](tools.ietf.org/html/rfc6749#section-4.4).
	Subject *string `json:"subject,omitempty"`
	// The scopes covered by the access token.
	Scopes []string `json:"scopes,omitempty"`
	// `true` if the access token exists.
	Existent *bool `json:"existent,omitempty"`
	// true` if the access token is usable (= exists and has not expired).
	Usable *bool `json:"usable,omitempty"`
	// `true` if the access token exists.
	Sufficient *bool `json:"sufficient,omitempty"`
	// `true` if the access token can be refreshed using the associated refresh token which had been issued along with the access token. `false` if the refresh token for the access token has expired or the access token has no associated refresh token.
	Refreshable *bool `json:"refreshable,omitempty"`
	// The extra properties associated with the access token.
	Properties []Property `json:"properties,omitempty"`
	// The client certificate thumbprint used to validate the access token.
	CertificateThumbprint *string `json:"certificateThumbprint,omitempty"`
	// The target resources. This represents the resources specified by the `resource` request parameters or by the `resource` property in the request object.  See \"Resource Indicators for OAuth 2.0\" for details.
	Resources []string `json:"resources,omitempty"`
	// The target resources this proeprty holds may be the same as or different from the ones `resource` property holds.  In some flows, the initial request and the subsequent token request are sent to different endpoints. Example flows are the Authorization Code Flow, the Refresh Token Flow, the CIBA Ping Mode, the CIBA Poll Mode and the Device Flow. In these flows, not only the initial request but also the subsequent token request can include the `resource` request parameters. The purpose of the `resource` request parameters in the token request is to narrow the range of the target resources from the original set of target resources requested by the preceding initial request. If narrowing down is performed, the target resources holded by the `resource` proeprty and the ones holded by this property are different. This property holds the narrowed set of target resources.  See \"Resource Indicators for OAuth 2.0\" for details.
	AccessTokenResources []string      `json:"accessTokenResources,omitempty"`
	AuthorizationDetails *AuthzDetails `json:"authorizationDetails,omitempty"`
	// The attributes of this service that the client application belongs to.
	ServiceAttributes []Pair `json:"serviceAttributes,omitempty"`
	// The attributes of the client.
	ClientAttributes []Pair `json:"clientAttributes,omitempty"`
	// The scopes property of this class is a list of scope names. The property does not hold information about scope attributes. This scopeDetails property was newly created to convey information about scope attributes.
	ScopeDetails []Scope `json:"scopeDetails,omitempty"`
	// The value of the `grant_id` request parameter of the device authorization request.  The `grant_id` request parameter is defined in [Grant Management for OAuth 2.0](https://openid.net/specs/fapi-grant-management.html) , which is supported by Authlete 2.3 and newer versions.
	GrantId *string `json:"grantId,omitempty"`
	Grant   *Grant  `json:"grant,omitempty"`
	// the flag which indicates whether the access token is for an external attachment.
	ForExternalAttachment *bool `json:"forExternalAttachment,omitempty"`
	// the claims that the user has consented for the client application to know.
	ConsentedClaims []string   `json:"consentedClaims,omitempty"`
	GrantType       *GrantType `json:"grantType,omitempty"`
	// The Authentication Context Class Reference of the user authentication that the authorization server performed  during the course of issuing the access token.
	Acr *string `json:"acr,omitempty"`
	// The time when the user authentication was performed during the course of issuing the access token.
	AuthTime *int64 `json:"authTime,omitempty"`
	// The entity ID of the client.
	ClientEntityId *string `json:"clientEntityId,omitempty"`
	// Flag which indicates whether the entity ID of the client was used when the request for the access token was made.
	ClientEntityIdUsed *bool `json:"clientEntityIdUsed,omitempty"`
	// The flag indicating whether the token is for credential issuance.
	ForCredentialIssuance *bool `json:"forCredentialIssuance,omitempty"`
	// The c_nonce
	Cnonce *string `json:"cnonce,omitempty"`
	// The time at which the `c_nonce` expires.
	CnonceExpiresAt *int64 `json:"cnonceExpiresAt,omitempty"`
	// The credentials that can be obtained by presenting the access token.
	IssuableCredentials *string `json:"issuableCredentials,omitempty"`
	// The expected nonce value for DPoP proof JWT, which should be used as the value of the `DPoP-Nonce` HTTP header.
	DpopNonce *string `json:"dpopNonce,omitempty"`
}

// NewIntrospectionResponse instantiates a new IntrospectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntrospectionResponse() *IntrospectionResponse {
	this := IntrospectionResponse{}
	return &this
}

// NewIntrospectionResponseWithDefaults instantiates a new IntrospectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntrospectionResponseWithDefaults() *IntrospectionResponse {
	this := IntrospectionResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetResultCode() string {
	if o == nil || IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *IntrospectionResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetResultMessage() string {
	if o == nil || IsNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasResultMessage() bool {
	if o != nil && !IsNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *IntrospectionResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *IntrospectionResponse) SetAction(v string) {
	o.Action = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetResponseContent() string {
	if o == nil || IsNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasResponseContent() bool {
	if o != nil && !IsNil(o.ResponseContent) {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *IntrospectionResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetClientId() int64 {
	if o == nil || IsNil(o.ClientId) {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetClientIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int64 and assigns it to the ClientId field.
func (o *IntrospectionResponse) SetClientId(v int64) {
	o.ClientId = &v
}

// GetClientIdAlias returns the ClientIdAlias field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetClientIdAlias() string {
	if o == nil || IsNil(o.ClientIdAlias) {
		var ret string
		return ret
	}
	return *o.ClientIdAlias
}

// GetClientIdAliasOk returns a tuple with the ClientIdAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetClientIdAliasOk() (*string, bool) {
	if o == nil || IsNil(o.ClientIdAlias) {
		return nil, false
	}
	return o.ClientIdAlias, true
}

// HasClientIdAlias returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasClientIdAlias() bool {
	if o != nil && !IsNil(o.ClientIdAlias) {
		return true
	}

	return false
}

// SetClientIdAlias gets a reference to the given string and assigns it to the ClientIdAlias field.
func (o *IntrospectionResponse) SetClientIdAlias(v string) {
	o.ClientIdAlias = &v
}

// GetClientIdAliasUsed returns the ClientIdAliasUsed field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetClientIdAliasUsed() bool {
	if o == nil || IsNil(o.ClientIdAliasUsed) {
		var ret bool
		return ret
	}
	return *o.ClientIdAliasUsed
}

// GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetClientIdAliasUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientIdAliasUsed) {
		return nil, false
	}
	return o.ClientIdAliasUsed, true
}

// HasClientIdAliasUsed returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasClientIdAliasUsed() bool {
	if o != nil && !IsNil(o.ClientIdAliasUsed) {
		return true
	}

	return false
}

// SetClientIdAliasUsed gets a reference to the given bool and assigns it to the ClientIdAliasUsed field.
func (o *IntrospectionResponse) SetClientIdAliasUsed(v bool) {
	o.ClientIdAliasUsed = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *IntrospectionResponse) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *IntrospectionResponse) SetSubject(v string) {
	o.Subject = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *IntrospectionResponse) SetScopes(v []string) {
	o.Scopes = v
}

// GetExistent returns the Existent field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetExistent() bool {
	if o == nil || IsNil(o.Existent) {
		var ret bool
		return ret
	}
	return *o.Existent
}

// GetExistentOk returns a tuple with the Existent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetExistentOk() (*bool, bool) {
	if o == nil || IsNil(o.Existent) {
		return nil, false
	}
	return o.Existent, true
}

// HasExistent returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasExistent() bool {
	if o != nil && !IsNil(o.Existent) {
		return true
	}

	return false
}

// SetExistent gets a reference to the given bool and assigns it to the Existent field.
func (o *IntrospectionResponse) SetExistent(v bool) {
	o.Existent = &v
}

// GetUsable returns the Usable field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetUsable() bool {
	if o == nil || IsNil(o.Usable) {
		var ret bool
		return ret
	}
	return *o.Usable
}

// GetUsableOk returns a tuple with the Usable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetUsableOk() (*bool, bool) {
	if o == nil || IsNil(o.Usable) {
		return nil, false
	}
	return o.Usable, true
}

// HasUsable returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasUsable() bool {
	if o != nil && !IsNil(o.Usable) {
		return true
	}

	return false
}

// SetUsable gets a reference to the given bool and assigns it to the Usable field.
func (o *IntrospectionResponse) SetUsable(v bool) {
	o.Usable = &v
}

// GetSufficient returns the Sufficient field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetSufficient() bool {
	if o == nil || IsNil(o.Sufficient) {
		var ret bool
		return ret
	}
	return *o.Sufficient
}

// GetSufficientOk returns a tuple with the Sufficient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetSufficientOk() (*bool, bool) {
	if o == nil || IsNil(o.Sufficient) {
		return nil, false
	}
	return o.Sufficient, true
}

// HasSufficient returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasSufficient() bool {
	if o != nil && !IsNil(o.Sufficient) {
		return true
	}

	return false
}

// SetSufficient gets a reference to the given bool and assigns it to the Sufficient field.
func (o *IntrospectionResponse) SetSufficient(v bool) {
	o.Sufficient = &v
}

// GetRefreshable returns the Refreshable field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetRefreshable() bool {
	if o == nil || IsNil(o.Refreshable) {
		var ret bool
		return ret
	}
	return *o.Refreshable
}

// GetRefreshableOk returns a tuple with the Refreshable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetRefreshableOk() (*bool, bool) {
	if o == nil || IsNil(o.Refreshable) {
		return nil, false
	}
	return o.Refreshable, true
}

// HasRefreshable returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasRefreshable() bool {
	if o != nil && !IsNil(o.Refreshable) {
		return true
	}

	return false
}

// SetRefreshable gets a reference to the given bool and assigns it to the Refreshable field.
func (o *IntrospectionResponse) SetRefreshable(v bool) {
	o.Refreshable = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetProperties() []Property {
	if o == nil || IsNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetPropertiesOk() ([]Property, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *IntrospectionResponse) SetProperties(v []Property) {
	o.Properties = v
}

// GetCertificateThumbprint returns the CertificateThumbprint field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetCertificateThumbprint() string {
	if o == nil || IsNil(o.CertificateThumbprint) {
		var ret string
		return ret
	}
	return *o.CertificateThumbprint
}

// GetCertificateThumbprintOk returns a tuple with the CertificateThumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetCertificateThumbprintOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateThumbprint) {
		return nil, false
	}
	return o.CertificateThumbprint, true
}

// HasCertificateThumbprint returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasCertificateThumbprint() bool {
	if o != nil && !IsNil(o.CertificateThumbprint) {
		return true
	}

	return false
}

// SetCertificateThumbprint gets a reference to the given string and assigns it to the CertificateThumbprint field.
func (o *IntrospectionResponse) SetCertificateThumbprint(v string) {
	o.CertificateThumbprint = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetResources() []string {
	if o == nil || IsNil(o.Resources) {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetResourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *IntrospectionResponse) SetResources(v []string) {
	o.Resources = v
}

// GetAccessTokenResources returns the AccessTokenResources field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetAccessTokenResources() []string {
	if o == nil || IsNil(o.AccessTokenResources) {
		var ret []string
		return ret
	}
	return o.AccessTokenResources
}

// GetAccessTokenResourcesOk returns a tuple with the AccessTokenResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetAccessTokenResourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.AccessTokenResources) {
		return nil, false
	}
	return o.AccessTokenResources, true
}

// HasAccessTokenResources returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasAccessTokenResources() bool {
	if o != nil && !IsNil(o.AccessTokenResources) {
		return true
	}

	return false
}

// SetAccessTokenResources gets a reference to the given []string and assigns it to the AccessTokenResources field.
func (o *IntrospectionResponse) SetAccessTokenResources(v []string) {
	o.AccessTokenResources = v
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetAuthorizationDetails() AuthzDetails {
	if o == nil || IsNil(o.AuthorizationDetails) {
		var ret AuthzDetails
		return ret
	}
	return *o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetAuthorizationDetailsOk() (*AuthzDetails, bool) {
	if o == nil || IsNil(o.AuthorizationDetails) {
		return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasAuthorizationDetails() bool {
	if o != nil && !IsNil(o.AuthorizationDetails) {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given AuthzDetails and assigns it to the AuthorizationDetails field.
func (o *IntrospectionResponse) SetAuthorizationDetails(v AuthzDetails) {
	o.AuthorizationDetails = &v
}

// GetServiceAttributes returns the ServiceAttributes field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetServiceAttributes() []Pair {
	if o == nil || IsNil(o.ServiceAttributes) {
		var ret []Pair
		return ret
	}
	return o.ServiceAttributes
}

// GetServiceAttributesOk returns a tuple with the ServiceAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetServiceAttributesOk() ([]Pair, bool) {
	if o == nil || IsNil(o.ServiceAttributes) {
		return nil, false
	}
	return o.ServiceAttributes, true
}

// HasServiceAttributes returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasServiceAttributes() bool {
	if o != nil && !IsNil(o.ServiceAttributes) {
		return true
	}

	return false
}

// SetServiceAttributes gets a reference to the given []Pair and assigns it to the ServiceAttributes field.
func (o *IntrospectionResponse) SetServiceAttributes(v []Pair) {
	o.ServiceAttributes = v
}

// GetClientAttributes returns the ClientAttributes field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetClientAttributes() []Pair {
	if o == nil || IsNil(o.ClientAttributes) {
		var ret []Pair
		return ret
	}
	return o.ClientAttributes
}

// GetClientAttributesOk returns a tuple with the ClientAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetClientAttributesOk() ([]Pair, bool) {
	if o == nil || IsNil(o.ClientAttributes) {
		return nil, false
	}
	return o.ClientAttributes, true
}

// HasClientAttributes returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasClientAttributes() bool {
	if o != nil && !IsNil(o.ClientAttributes) {
		return true
	}

	return false
}

// SetClientAttributes gets a reference to the given []Pair and assigns it to the ClientAttributes field.
func (o *IntrospectionResponse) SetClientAttributes(v []Pair) {
	o.ClientAttributes = v
}

// GetScopeDetails returns the ScopeDetails field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetScopeDetails() []Scope {
	if o == nil || IsNil(o.ScopeDetails) {
		var ret []Scope
		return ret
	}
	return o.ScopeDetails
}

// GetScopeDetailsOk returns a tuple with the ScopeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetScopeDetailsOk() ([]Scope, bool) {
	if o == nil || IsNil(o.ScopeDetails) {
		return nil, false
	}
	return o.ScopeDetails, true
}

// HasScopeDetails returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasScopeDetails() bool {
	if o != nil && !IsNil(o.ScopeDetails) {
		return true
	}

	return false
}

// SetScopeDetails gets a reference to the given []Scope and assigns it to the ScopeDetails field.
func (o *IntrospectionResponse) SetScopeDetails(v []Scope) {
	o.ScopeDetails = v
}

// GetGrantId returns the GrantId field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetGrantId() string {
	if o == nil || IsNil(o.GrantId) {
		var ret string
		return ret
	}
	return *o.GrantId
}

// GetGrantIdOk returns a tuple with the GrantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetGrantIdOk() (*string, bool) {
	if o == nil || IsNil(o.GrantId) {
		return nil, false
	}
	return o.GrantId, true
}

// HasGrantId returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasGrantId() bool {
	if o != nil && !IsNil(o.GrantId) {
		return true
	}

	return false
}

// SetGrantId gets a reference to the given string and assigns it to the GrantId field.
func (o *IntrospectionResponse) SetGrantId(v string) {
	o.GrantId = &v
}

// GetGrant returns the Grant field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetGrant() Grant {
	if o == nil || IsNil(o.Grant) {
		var ret Grant
		return ret
	}
	return *o.Grant
}

// GetGrantOk returns a tuple with the Grant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetGrantOk() (*Grant, bool) {
	if o == nil || IsNil(o.Grant) {
		return nil, false
	}
	return o.Grant, true
}

// HasGrant returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasGrant() bool {
	if o != nil && !IsNil(o.Grant) {
		return true
	}

	return false
}

// SetGrant gets a reference to the given Grant and assigns it to the Grant field.
func (o *IntrospectionResponse) SetGrant(v Grant) {
	o.Grant = &v
}

// GetForExternalAttachment returns the ForExternalAttachment field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetForExternalAttachment() bool {
	if o == nil || IsNil(o.ForExternalAttachment) {
		var ret bool
		return ret
	}
	return *o.ForExternalAttachment
}

// GetForExternalAttachmentOk returns a tuple with the ForExternalAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetForExternalAttachmentOk() (*bool, bool) {
	if o == nil || IsNil(o.ForExternalAttachment) {
		return nil, false
	}
	return o.ForExternalAttachment, true
}

// HasForExternalAttachment returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasForExternalAttachment() bool {
	if o != nil && !IsNil(o.ForExternalAttachment) {
		return true
	}

	return false
}

// SetForExternalAttachment gets a reference to the given bool and assigns it to the ForExternalAttachment field.
func (o *IntrospectionResponse) SetForExternalAttachment(v bool) {
	o.ForExternalAttachment = &v
}

// GetConsentedClaims returns the ConsentedClaims field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetConsentedClaims() []string {
	if o == nil || IsNil(o.ConsentedClaims) {
		var ret []string
		return ret
	}
	return o.ConsentedClaims
}

// GetConsentedClaimsOk returns a tuple with the ConsentedClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetConsentedClaimsOk() ([]string, bool) {
	if o == nil || IsNil(o.ConsentedClaims) {
		return nil, false
	}
	return o.ConsentedClaims, true
}

// HasConsentedClaims returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasConsentedClaims() bool {
	if o != nil && !IsNil(o.ConsentedClaims) {
		return true
	}

	return false
}

// SetConsentedClaims gets a reference to the given []string and assigns it to the ConsentedClaims field.
func (o *IntrospectionResponse) SetConsentedClaims(v []string) {
	o.ConsentedClaims = v
}

// GetGrantType returns the GrantType field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetGrantType() GrantType {
	if o == nil || IsNil(o.GrantType) {
		var ret GrantType
		return ret
	}
	return *o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetGrantTypeOk() (*GrantType, bool) {
	if o == nil || IsNil(o.GrantType) {
		return nil, false
	}
	return o.GrantType, true
}

// HasGrantType returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasGrantType() bool {
	if o != nil && !IsNil(o.GrantType) {
		return true
	}

	return false
}

// SetGrantType gets a reference to the given GrantType and assigns it to the GrantType field.
func (o *IntrospectionResponse) SetGrantType(v GrantType) {
	o.GrantType = &v
}

// GetAcr returns the Acr field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetAcr() string {
	if o == nil || IsNil(o.Acr) {
		var ret string
		return ret
	}
	return *o.Acr
}

// GetAcrOk returns a tuple with the Acr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetAcrOk() (*string, bool) {
	if o == nil || IsNil(o.Acr) {
		return nil, false
	}
	return o.Acr, true
}

// HasAcr returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasAcr() bool {
	if o != nil && !IsNil(o.Acr) {
		return true
	}

	return false
}

// SetAcr gets a reference to the given string and assigns it to the Acr field.
func (o *IntrospectionResponse) SetAcr(v string) {
	o.Acr = &v
}

// GetAuthTime returns the AuthTime field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetAuthTime() int64 {
	if o == nil || IsNil(o.AuthTime) {
		var ret int64
		return ret
	}
	return *o.AuthTime
}

// GetAuthTimeOk returns a tuple with the AuthTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetAuthTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.AuthTime) {
		return nil, false
	}
	return o.AuthTime, true
}

// HasAuthTime returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasAuthTime() bool {
	if o != nil && !IsNil(o.AuthTime) {
		return true
	}

	return false
}

// SetAuthTime gets a reference to the given int64 and assigns it to the AuthTime field.
func (o *IntrospectionResponse) SetAuthTime(v int64) {
	o.AuthTime = &v
}

// GetClientEntityId returns the ClientEntityId field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetClientEntityId() string {
	if o == nil || IsNil(o.ClientEntityId) {
		var ret string
		return ret
	}
	return *o.ClientEntityId
}

// GetClientEntityIdOk returns a tuple with the ClientEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetClientEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientEntityId) {
		return nil, false
	}
	return o.ClientEntityId, true
}

// HasClientEntityId returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasClientEntityId() bool {
	if o != nil && !IsNil(o.ClientEntityId) {
		return true
	}

	return false
}

// SetClientEntityId gets a reference to the given string and assigns it to the ClientEntityId field.
func (o *IntrospectionResponse) SetClientEntityId(v string) {
	o.ClientEntityId = &v
}

// GetClientEntityIdUsed returns the ClientEntityIdUsed field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetClientEntityIdUsed() bool {
	if o == nil || IsNil(o.ClientEntityIdUsed) {
		var ret bool
		return ret
	}
	return *o.ClientEntityIdUsed
}

// GetClientEntityIdUsedOk returns a tuple with the ClientEntityIdUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetClientEntityIdUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientEntityIdUsed) {
		return nil, false
	}
	return o.ClientEntityIdUsed, true
}

// HasClientEntityIdUsed returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasClientEntityIdUsed() bool {
	if o != nil && !IsNil(o.ClientEntityIdUsed) {
		return true
	}

	return false
}

// SetClientEntityIdUsed gets a reference to the given bool and assigns it to the ClientEntityIdUsed field.
func (o *IntrospectionResponse) SetClientEntityIdUsed(v bool) {
	o.ClientEntityIdUsed = &v
}

// GetForCredentialIssuance returns the ForCredentialIssuance field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetForCredentialIssuance() bool {
	if o == nil || IsNil(o.ForCredentialIssuance) {
		var ret bool
		return ret
	}
	return *o.ForCredentialIssuance
}

// GetForCredentialIssuanceOk returns a tuple with the ForCredentialIssuance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetForCredentialIssuanceOk() (*bool, bool) {
	if o == nil || IsNil(o.ForCredentialIssuance) {
		return nil, false
	}
	return o.ForCredentialIssuance, true
}

// HasForCredentialIssuance returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasForCredentialIssuance() bool {
	if o != nil && !IsNil(o.ForCredentialIssuance) {
		return true
	}

	return false
}

// SetForCredentialIssuance gets a reference to the given bool and assigns it to the ForCredentialIssuance field.
func (o *IntrospectionResponse) SetForCredentialIssuance(v bool) {
	o.ForCredentialIssuance = &v
}

// GetCnonce returns the Cnonce field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetCnonce() string {
	if o == nil || IsNil(o.Cnonce) {
		var ret string
		return ret
	}
	return *o.Cnonce
}

// GetCnonceOk returns a tuple with the Cnonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetCnonceOk() (*string, bool) {
	if o == nil || IsNil(o.Cnonce) {
		return nil, false
	}
	return o.Cnonce, true
}

// HasCnonce returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasCnonce() bool {
	if o != nil && !IsNil(o.Cnonce) {
		return true
	}

	return false
}

// SetCnonce gets a reference to the given string and assigns it to the Cnonce field.
func (o *IntrospectionResponse) SetCnonce(v string) {
	o.Cnonce = &v
}

// GetCnonceExpiresAt returns the CnonceExpiresAt field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetCnonceExpiresAt() int64 {
	if o == nil || IsNil(o.CnonceExpiresAt) {
		var ret int64
		return ret
	}
	return *o.CnonceExpiresAt
}

// GetCnonceExpiresAtOk returns a tuple with the CnonceExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetCnonceExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CnonceExpiresAt) {
		return nil, false
	}
	return o.CnonceExpiresAt, true
}

// HasCnonceExpiresAt returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasCnonceExpiresAt() bool {
	if o != nil && !IsNil(o.CnonceExpiresAt) {
		return true
	}

	return false
}

// SetCnonceExpiresAt gets a reference to the given int64 and assigns it to the CnonceExpiresAt field.
func (o *IntrospectionResponse) SetCnonceExpiresAt(v int64) {
	o.CnonceExpiresAt = &v
}

// GetIssuableCredentials returns the IssuableCredentials field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetIssuableCredentials() string {
	if o == nil || IsNil(o.IssuableCredentials) {
		var ret string
		return ret
	}
	return *o.IssuableCredentials
}

// GetIssuableCredentialsOk returns a tuple with the IssuableCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetIssuableCredentialsOk() (*string, bool) {
	if o == nil || IsNil(o.IssuableCredentials) {
		return nil, false
	}
	return o.IssuableCredentials, true
}

// HasIssuableCredentials returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasIssuableCredentials() bool {
	if o != nil && !IsNil(o.IssuableCredentials) {
		return true
	}

	return false
}

// SetIssuableCredentials gets a reference to the given string and assigns it to the IssuableCredentials field.
func (o *IntrospectionResponse) SetIssuableCredentials(v string) {
	o.IssuableCredentials = &v
}

// GetDpopNonce returns the DpopNonce field value if set, zero value otherwise.
func (o *IntrospectionResponse) GetDpopNonce() string {
	if o == nil || IsNil(o.DpopNonce) {
		var ret string
		return ret
	}
	return *o.DpopNonce
}

// GetDpopNonceOk returns a tuple with the DpopNonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionResponse) GetDpopNonceOk() (*string, bool) {
	if o == nil || IsNil(o.DpopNonce) {
		return nil, false
	}
	return o.DpopNonce, true
}

// HasDpopNonce returns a boolean if a field has been set.
func (o *IntrospectionResponse) HasDpopNonce() bool {
	if o != nil && !IsNil(o.DpopNonce) {
		return true
	}

	return false
}

// SetDpopNonce gets a reference to the given string and assigns it to the DpopNonce field.
func (o *IntrospectionResponse) SetDpopNonce(v string) {
	o.DpopNonce = &v
}

func (o IntrospectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntrospectionResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Existent) {
		toSerialize["existent"] = o.Existent
	}
	if !IsNil(o.Usable) {
		toSerialize["usable"] = o.Usable
	}
	if !IsNil(o.Sufficient) {
		toSerialize["sufficient"] = o.Sufficient
	}
	if !IsNil(o.Refreshable) {
		toSerialize["refreshable"] = o.Refreshable
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.CertificateThumbprint) {
		toSerialize["certificateThumbprint"] = o.CertificateThumbprint
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.AccessTokenResources) {
		toSerialize["accessTokenResources"] = o.AccessTokenResources
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
	if !IsNil(o.ScopeDetails) {
		toSerialize["scopeDetails"] = o.ScopeDetails
	}
	if !IsNil(o.GrantId) {
		toSerialize["grantId"] = o.GrantId
	}
	if !IsNil(o.Grant) {
		toSerialize["grant"] = o.Grant
	}
	if !IsNil(o.ForExternalAttachment) {
		toSerialize["forExternalAttachment"] = o.ForExternalAttachment
	}
	if !IsNil(o.ConsentedClaims) {
		toSerialize["consentedClaims"] = o.ConsentedClaims
	}
	if !IsNil(o.GrantType) {
		toSerialize["grantType"] = o.GrantType
	}
	if !IsNil(o.Acr) {
		toSerialize["acr"] = o.Acr
	}
	if !IsNil(o.AuthTime) {
		toSerialize["authTime"] = o.AuthTime
	}
	if !IsNil(o.ClientEntityId) {
		toSerialize["clientEntityId"] = o.ClientEntityId
	}
	if !IsNil(o.ClientEntityIdUsed) {
		toSerialize["clientEntityIdUsed"] = o.ClientEntityIdUsed
	}
	if !IsNil(o.ForCredentialIssuance) {
		toSerialize["forCredentialIssuance"] = o.ForCredentialIssuance
	}
	if !IsNil(o.Cnonce) {
		toSerialize["cnonce"] = o.Cnonce
	}
	if !IsNil(o.CnonceExpiresAt) {
		toSerialize["cnonceExpiresAt"] = o.CnonceExpiresAt
	}
	if !IsNil(o.IssuableCredentials) {
		toSerialize["issuableCredentials"] = o.IssuableCredentials
	}
	if !IsNil(o.DpopNonce) {
		toSerialize["dpopNonce"] = o.DpopNonce
	}
	return toSerialize, nil
}

type NullableIntrospectionResponse struct {
	value *IntrospectionResponse
	isSet bool
}

func (v NullableIntrospectionResponse) Get() *IntrospectionResponse {
	return v.value
}

func (v *NullableIntrospectionResponse) Set(val *IntrospectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIntrospectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIntrospectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntrospectionResponse(val *IntrospectionResponse) *NullableIntrospectionResponse {
	return &NullableIntrospectionResponse{value: val, isSet: true}
}

func (v NullableIntrospectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntrospectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
