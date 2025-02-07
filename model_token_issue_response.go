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

// checks if the TokenIssueResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenIssueResponse{}

// TokenIssueResponse struct for TokenIssueResponse
type TokenIssueResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take.
	Action *string `json:"action,omitempty"`
	// The content that the authorization server implementation is to return to the client application. Its format is JSON.
	ResponseContent *string `json:"responseContent,omitempty"`
	// The newly issued access token. This parameter is a non-null value only when the value of `action` parameter is `OK`.
	AccessToken *string `json:"accessToken,omitempty"`
	// The datetime at which the newly issued access token will expire. The value is represented in milliseconds since the Unix epoch (1970-01-01).
	AccessTokenExpiresAt *int64 `json:"accessTokenExpiresAt,omitempty"`
	// The duration of the newly issued access token in seconds.
	AccessTokenDuration *int64 `json:"accessTokenDuration,omitempty"`
	// The refresh token. This parameter is a non-null value only when `action` is `OK` and the service supports the refresh token flow. If `refreshTokenKept` is set to `false`, a new refresh token is issued and the old refresh token used in the refresh token flow is invalidated. On the contrary, if `refreshTokenKept` is set to `true`, the refresh token itself is not refreshed.
	RefreshToken *string `json:"refreshToken,omitempty"`
	// The datetime at which the newly issued refresh token will expire. The value is represented in milliseconds since the Unix epoch (1970-01-01).
	RefreshTokenExpiresAt *int64 `json:"refreshTokenExpiresAt,omitempty"`
	// The duration of the newly issued refresh token in seconds.
	RefreshTokenDuration *int64 `json:"refreshTokenDuration,omitempty"`
	// The client ID.
	ClientId *int64 `json:"clientId,omitempty"`
	// The client ID alias. If the client did not have an alias, this parameter is `null`.
	ClientIdAlias *string `json:"clientIdAlias,omitempty"`
	// The flag which indicates whether the client ID alias was used when the token request was made. `true` if the client ID alias was used when the token request was made.
	ClientIdAliasUsed *bool `json:"clientIdAliasUsed,omitempty"`
	// The subject (= resource owner's ID) of the access token. Even if an access token has been issued by calling `/api/auth/token` API, this parameter is `null` if the flow of the token request was [Client Credentials Flow](https://datatracker.ietf.org/doc/html/rfc6749#section-4.4) (`grant_type=client_credentials`) because it means the access token is not associated with any specific end-user.
	Subject *string `json:"subject,omitempty"`
	// The scopes covered by the access token.
	Scopes []string `json:"scopes,omitempty"`
	// The extra properties associated with the access token. This parameter is `null` when no extra property is associated with the issued access token.
	Properties []Property `json:"properties,omitempty"`
	// The newly issued access token in JWT format. If the authorization server is configured to issue JWT-based access tokens (= if the service's `accessTokenSignAlg` value is a non-null value), a JWT-based access token is issued along with the original random-string one.
	JwtAccessToken *string `json:"jwtAccessToken,omitempty"`
	// The target resources of the access token being issued. See \"Resource Indicators for OAuth 2.0\" for details.
	AccessTokenResources []string      `json:"accessTokenResources,omitempty"`
	AuthorizationDetails *AuthzDetails `json:"authorizationDetails,omitempty"`
	// The attributes of this service that the client application belongs to.
	ServiceAttributes []Pair `json:"serviceAttributes,omitempty"`
	// The attributes of the client.
	ClientAttributes []Pair `json:"clientAttributes,omitempty"`
	// The entity ID of the client.
	ClientEntityId *string `json:"clientEntityId,omitempty"`
	// Flag which indicates whether the entity ID of the client was used when the request for the access token was made.
	ClientEntityIdUsed *bool `json:"clientEntityIdUsed,omitempty"`
	// The scopes associated with the refresh token. May be null.
	RefreshTokenScopes []string `json:"refreshTokenScopes,omitempty"`
}

// NewTokenIssueResponse instantiates a new TokenIssueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenIssueResponse() *TokenIssueResponse {
	this := TokenIssueResponse{}
	return &this
}

// NewTokenIssueResponseWithDefaults instantiates a new TokenIssueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenIssueResponseWithDefaults() *TokenIssueResponse {
	this := TokenIssueResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetResultCode() string {
	if o == nil || IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *TokenIssueResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetResultMessage() string {
	if o == nil || IsNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasResultMessage() bool {
	if o != nil && !IsNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *TokenIssueResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *TokenIssueResponse) SetAction(v string) {
	o.Action = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetResponseContent() string {
	if o == nil || IsNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasResponseContent() bool {
	if o != nil && !IsNil(o.ResponseContent) {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *TokenIssueResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *TokenIssueResponse) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetAccessTokenExpiresAt returns the AccessTokenExpiresAt field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetAccessTokenExpiresAt() int64 {
	if o == nil || IsNil(o.AccessTokenExpiresAt) {
		var ret int64
		return ret
	}
	return *o.AccessTokenExpiresAt
}

// GetAccessTokenExpiresAtOk returns a tuple with the AccessTokenExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetAccessTokenExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.AccessTokenExpiresAt) {
		return nil, false
	}
	return o.AccessTokenExpiresAt, true
}

// HasAccessTokenExpiresAt returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasAccessTokenExpiresAt() bool {
	if o != nil && !IsNil(o.AccessTokenExpiresAt) {
		return true
	}

	return false
}

// SetAccessTokenExpiresAt gets a reference to the given int64 and assigns it to the AccessTokenExpiresAt field.
func (o *TokenIssueResponse) SetAccessTokenExpiresAt(v int64) {
	o.AccessTokenExpiresAt = &v
}

// GetAccessTokenDuration returns the AccessTokenDuration field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetAccessTokenDuration() int64 {
	if o == nil || IsNil(o.AccessTokenDuration) {
		var ret int64
		return ret
	}
	return *o.AccessTokenDuration
}

// GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetAccessTokenDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.AccessTokenDuration) {
		return nil, false
	}
	return o.AccessTokenDuration, true
}

// HasAccessTokenDuration returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasAccessTokenDuration() bool {
	if o != nil && !IsNil(o.AccessTokenDuration) {
		return true
	}

	return false
}

// SetAccessTokenDuration gets a reference to the given int64 and assigns it to the AccessTokenDuration field.
func (o *TokenIssueResponse) SetAccessTokenDuration(v int64) {
	o.AccessTokenDuration = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *TokenIssueResponse) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetRefreshTokenExpiresAt returns the RefreshTokenExpiresAt field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetRefreshTokenExpiresAt() int64 {
	if o == nil || IsNil(o.RefreshTokenExpiresAt) {
		var ret int64
		return ret
	}
	return *o.RefreshTokenExpiresAt
}

// GetRefreshTokenExpiresAtOk returns a tuple with the RefreshTokenExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetRefreshTokenExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.RefreshTokenExpiresAt) {
		return nil, false
	}
	return o.RefreshTokenExpiresAt, true
}

// HasRefreshTokenExpiresAt returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasRefreshTokenExpiresAt() bool {
	if o != nil && !IsNil(o.RefreshTokenExpiresAt) {
		return true
	}

	return false
}

// SetRefreshTokenExpiresAt gets a reference to the given int64 and assigns it to the RefreshTokenExpiresAt field.
func (o *TokenIssueResponse) SetRefreshTokenExpiresAt(v int64) {
	o.RefreshTokenExpiresAt = &v
}

// GetRefreshTokenDuration returns the RefreshTokenDuration field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetRefreshTokenDuration() int64 {
	if o == nil || IsNil(o.RefreshTokenDuration) {
		var ret int64
		return ret
	}
	return *o.RefreshTokenDuration
}

// GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetRefreshTokenDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.RefreshTokenDuration) {
		return nil, false
	}
	return o.RefreshTokenDuration, true
}

// HasRefreshTokenDuration returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasRefreshTokenDuration() bool {
	if o != nil && !IsNil(o.RefreshTokenDuration) {
		return true
	}

	return false
}

// SetRefreshTokenDuration gets a reference to the given int64 and assigns it to the RefreshTokenDuration field.
func (o *TokenIssueResponse) SetRefreshTokenDuration(v int64) {
	o.RefreshTokenDuration = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetClientId() int64 {
	if o == nil || IsNil(o.ClientId) {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetClientIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int64 and assigns it to the ClientId field.
func (o *TokenIssueResponse) SetClientId(v int64) {
	o.ClientId = &v
}

// GetClientIdAlias returns the ClientIdAlias field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetClientIdAlias() string {
	if o == nil || IsNil(o.ClientIdAlias) {
		var ret string
		return ret
	}
	return *o.ClientIdAlias
}

// GetClientIdAliasOk returns a tuple with the ClientIdAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetClientIdAliasOk() (*string, bool) {
	if o == nil || IsNil(o.ClientIdAlias) {
		return nil, false
	}
	return o.ClientIdAlias, true
}

// HasClientIdAlias returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasClientIdAlias() bool {
	if o != nil && !IsNil(o.ClientIdAlias) {
		return true
	}

	return false
}

// SetClientIdAlias gets a reference to the given string and assigns it to the ClientIdAlias field.
func (o *TokenIssueResponse) SetClientIdAlias(v string) {
	o.ClientIdAlias = &v
}

// GetClientIdAliasUsed returns the ClientIdAliasUsed field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetClientIdAliasUsed() bool {
	if o == nil || IsNil(o.ClientIdAliasUsed) {
		var ret bool
		return ret
	}
	return *o.ClientIdAliasUsed
}

// GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetClientIdAliasUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientIdAliasUsed) {
		return nil, false
	}
	return o.ClientIdAliasUsed, true
}

// HasClientIdAliasUsed returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasClientIdAliasUsed() bool {
	if o != nil && !IsNil(o.ClientIdAliasUsed) {
		return true
	}

	return false
}

// SetClientIdAliasUsed gets a reference to the given bool and assigns it to the ClientIdAliasUsed field.
func (o *TokenIssueResponse) SetClientIdAliasUsed(v bool) {
	o.ClientIdAliasUsed = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *TokenIssueResponse) SetSubject(v string) {
	o.Subject = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *TokenIssueResponse) SetScopes(v []string) {
	o.Scopes = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetProperties() []Property {
	if o == nil || IsNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetPropertiesOk() ([]Property, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *TokenIssueResponse) SetProperties(v []Property) {
	o.Properties = v
}

// GetJwtAccessToken returns the JwtAccessToken field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetJwtAccessToken() string {
	if o == nil || IsNil(o.JwtAccessToken) {
		var ret string
		return ret
	}
	return *o.JwtAccessToken
}

// GetJwtAccessTokenOk returns a tuple with the JwtAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetJwtAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.JwtAccessToken) {
		return nil, false
	}
	return o.JwtAccessToken, true
}

// HasJwtAccessToken returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasJwtAccessToken() bool {
	if o != nil && !IsNil(o.JwtAccessToken) {
		return true
	}

	return false
}

// SetJwtAccessToken gets a reference to the given string and assigns it to the JwtAccessToken field.
func (o *TokenIssueResponse) SetJwtAccessToken(v string) {
	o.JwtAccessToken = &v
}

// GetAccessTokenResources returns the AccessTokenResources field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetAccessTokenResources() []string {
	if o == nil || IsNil(o.AccessTokenResources) {
		var ret []string
		return ret
	}
	return o.AccessTokenResources
}

// GetAccessTokenResourcesOk returns a tuple with the AccessTokenResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetAccessTokenResourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.AccessTokenResources) {
		return nil, false
	}
	return o.AccessTokenResources, true
}

// HasAccessTokenResources returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasAccessTokenResources() bool {
	if o != nil && !IsNil(o.AccessTokenResources) {
		return true
	}

	return false
}

// SetAccessTokenResources gets a reference to the given []string and assigns it to the AccessTokenResources field.
func (o *TokenIssueResponse) SetAccessTokenResources(v []string) {
	o.AccessTokenResources = v
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetAuthorizationDetails() AuthzDetails {
	if o == nil || IsNil(o.AuthorizationDetails) {
		var ret AuthzDetails
		return ret
	}
	return *o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetAuthorizationDetailsOk() (*AuthzDetails, bool) {
	if o == nil || IsNil(o.AuthorizationDetails) {
		return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasAuthorizationDetails() bool {
	if o != nil && !IsNil(o.AuthorizationDetails) {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given AuthzDetails and assigns it to the AuthorizationDetails field.
func (o *TokenIssueResponse) SetAuthorizationDetails(v AuthzDetails) {
	o.AuthorizationDetails = &v
}

// GetServiceAttributes returns the ServiceAttributes field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetServiceAttributes() []Pair {
	if o == nil || IsNil(o.ServiceAttributes) {
		var ret []Pair
		return ret
	}
	return o.ServiceAttributes
}

// GetServiceAttributesOk returns a tuple with the ServiceAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetServiceAttributesOk() ([]Pair, bool) {
	if o == nil || IsNil(o.ServiceAttributes) {
		return nil, false
	}
	return o.ServiceAttributes, true
}

// HasServiceAttributes returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasServiceAttributes() bool {
	if o != nil && !IsNil(o.ServiceAttributes) {
		return true
	}

	return false
}

// SetServiceAttributes gets a reference to the given []Pair and assigns it to the ServiceAttributes field.
func (o *TokenIssueResponse) SetServiceAttributes(v []Pair) {
	o.ServiceAttributes = v
}

// GetClientAttributes returns the ClientAttributes field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetClientAttributes() []Pair {
	if o == nil || IsNil(o.ClientAttributes) {
		var ret []Pair
		return ret
	}
	return o.ClientAttributes
}

// GetClientAttributesOk returns a tuple with the ClientAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetClientAttributesOk() ([]Pair, bool) {
	if o == nil || IsNil(o.ClientAttributes) {
		return nil, false
	}
	return o.ClientAttributes, true
}

// HasClientAttributes returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasClientAttributes() bool {
	if o != nil && !IsNil(o.ClientAttributes) {
		return true
	}

	return false
}

// SetClientAttributes gets a reference to the given []Pair and assigns it to the ClientAttributes field.
func (o *TokenIssueResponse) SetClientAttributes(v []Pair) {
	o.ClientAttributes = v
}

// GetClientEntityId returns the ClientEntityId field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetClientEntityId() string {
	if o == nil || IsNil(o.ClientEntityId) {
		var ret string
		return ret
	}
	return *o.ClientEntityId
}

// GetClientEntityIdOk returns a tuple with the ClientEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetClientEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientEntityId) {
		return nil, false
	}
	return o.ClientEntityId, true
}

// HasClientEntityId returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasClientEntityId() bool {
	if o != nil && !IsNil(o.ClientEntityId) {
		return true
	}

	return false
}

// SetClientEntityId gets a reference to the given string and assigns it to the ClientEntityId field.
func (o *TokenIssueResponse) SetClientEntityId(v string) {
	o.ClientEntityId = &v
}

// GetClientEntityIdUsed returns the ClientEntityIdUsed field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetClientEntityIdUsed() bool {
	if o == nil || IsNil(o.ClientEntityIdUsed) {
		var ret bool
		return ret
	}
	return *o.ClientEntityIdUsed
}

// GetClientEntityIdUsedOk returns a tuple with the ClientEntityIdUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetClientEntityIdUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientEntityIdUsed) {
		return nil, false
	}
	return o.ClientEntityIdUsed, true
}

// HasClientEntityIdUsed returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasClientEntityIdUsed() bool {
	if o != nil && !IsNil(o.ClientEntityIdUsed) {
		return true
	}

	return false
}

// SetClientEntityIdUsed gets a reference to the given bool and assigns it to the ClientEntityIdUsed field.
func (o *TokenIssueResponse) SetClientEntityIdUsed(v bool) {
	o.ClientEntityIdUsed = &v
}

// GetRefreshTokenScopes returns the RefreshTokenScopes field value if set, zero value otherwise.
func (o *TokenIssueResponse) GetRefreshTokenScopes() []string {
	if o == nil || IsNil(o.RefreshTokenScopes) {
		var ret []string
		return ret
	}
	return o.RefreshTokenScopes
}

// GetRefreshTokenScopesOk returns a tuple with the RefreshTokenScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueResponse) GetRefreshTokenScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.RefreshTokenScopes) {
		return nil, false
	}
	return o.RefreshTokenScopes, true
}

// HasRefreshTokenScopes returns a boolean if a field has been set.
func (o *TokenIssueResponse) HasRefreshTokenScopes() bool {
	if o != nil && !IsNil(o.RefreshTokenScopes) {
		return true
	}

	return false
}

// SetRefreshTokenScopes gets a reference to the given []string and assigns it to the RefreshTokenScopes field.
func (o *TokenIssueResponse) SetRefreshTokenScopes(v []string) {
	o.RefreshTokenScopes = v
}

func (o TokenIssueResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenIssueResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	if !IsNil(o.AccessTokenExpiresAt) {
		toSerialize["accessTokenExpiresAt"] = o.AccessTokenExpiresAt
	}
	if !IsNil(o.AccessTokenDuration) {
		toSerialize["accessTokenDuration"] = o.AccessTokenDuration
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refreshToken"] = o.RefreshToken
	}
	if !IsNil(o.RefreshTokenExpiresAt) {
		toSerialize["refreshTokenExpiresAt"] = o.RefreshTokenExpiresAt
	}
	if !IsNil(o.RefreshTokenDuration) {
		toSerialize["refreshTokenDuration"] = o.RefreshTokenDuration
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
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.JwtAccessToken) {
		toSerialize["jwtAccessToken"] = o.JwtAccessToken
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
	if !IsNil(o.ClientEntityId) {
		toSerialize["clientEntityId"] = o.ClientEntityId
	}
	if !IsNil(o.ClientEntityIdUsed) {
		toSerialize["clientEntityIdUsed"] = o.ClientEntityIdUsed
	}
	if !IsNil(o.RefreshTokenScopes) {
		toSerialize["refreshTokenScopes"] = o.RefreshTokenScopes
	}
	return toSerialize, nil
}

type NullableTokenIssueResponse struct {
	value *TokenIssueResponse
	isSet bool
}

func (v NullableTokenIssueResponse) Get() *TokenIssueResponse {
	return v.value
}

func (v *NullableTokenIssueResponse) Set(val *TokenIssueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenIssueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenIssueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenIssueResponse(val *TokenIssueResponse) *NullableTokenIssueResponse {
	return &NullableTokenIssueResponse{value: val, isSet: true}
}

func (v NullableTokenIssueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenIssueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
