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

// checks if the TokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenResponse{}

// TokenResponse struct for TokenResponse
type TokenResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take.
	Action *string `json:"action,omitempty"`
	// The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of `action` parameter.
	ResponseContent *string `json:"responseContent,omitempty"`
	// The value of `username` request parameter in the token request. The client application must specify username when it uses [Resource Owner Password Grant](https://datatracker.ietf.org/doc/html/rfc6749#section-4.3). In other words, when the value of `grant_type` request parameter is `password`, `username` request parameter must come along.  This parameter has a value only if the value of `grant_type` request parameter is `password` and the token request is valid.
	Username *string `json:"username,omitempty"`
	// The value of `password` request parameter in the token request. The client application must specify password when it uses [Resource Owner Password Grant](https://datatracker.ietf.org/doc/html/rfc6749#section-4.3). In other words, when the value of `grant_type` request parameter is `password`, `password` request parameter must come along.  This parameter has a value only if the value of `grant_type` request parameter is `password` and the token request is valid.
	Password *string `json:"password,omitempty"`
	// The ticket which is necessary to call Authlete's `/auth/token/fail` API or `/auth/token/issue` API.  This parameter has a value only if the value of `grant_type` request parameter is `password` and the token request is valid.
	Ticket *string `json:"ticket,omitempty"`
	// The newly issued access token.
	AccessToken *string `json:"accessToken,omitempty"`
	// The datetime at which the newly issued access token will expire. The value is represented in milliseconds since the Unix epoch (1970-01-01).
	AccessTokenExpiresAt *int64 `json:"accessTokenExpiresAt,omitempty"`
	// The duration of the newly issued access token in seconds.
	AccessTokenDuration *int64 `json:"accessTokenDuration,omitempty"`
	// The newly issued refresh token.
	RefreshToken *string `json:"refreshToken,omitempty"`
	// The datetime at which the newly issued refresh token will expire. The value is represented in milliseconds since the Unix epoch (1970-01-01).
	RefreshTokenExpiresAt *int64 `json:"refreshTokenExpiresAt,omitempty"`
	// The duration of the newly issued refresh token in seconds.
	RefreshTokenDuration *int64 `json:"refreshTokenDuration,omitempty"`
	// The newly issued ID token. Note that an ID token is issued from a token endpoint only when the `response_type` request parameter of the authorization request to an authorization endpoint has contained `code` and the `scope` request parameter has contained `openid`.
	IdToken *string `json:"idToken,omitempty"`
	// The grant type of the token request.
	GrantType *string `json:"grantType,omitempty"`
	// The client ID.
	ClientId *int64 `json:"clientId,omitempty"`
	// The client ID alias when the token request was made. If the client did not have an alias, this parameter is `null`. Also, if the token request was invalid and it failed to identify a client, this parameter is `null`.
	ClientIdAlias *string `json:"clientIdAlias,omitempty"`
	// The flag which indicates whether the client ID alias was used when the token request was made. `true` if the client ID alias was used when the token request was made.
	ClientIdAliasUsed *bool `json:"clientIdAliasUsed,omitempty"`
	// The subject (= resource owner's ID) of the access token. Even if an access token has been issued by the call of `/api/auth/token` API, this parameter is `null` if the flow of the token request was [Client Credentials Flow](https://datatracker.ietf.org/doc/html/rfc6749#section-4.4) (`grant_type=client_credentials`) because it means the access token is not associated with any specific end-user.
	Subject *string `json:"subject,omitempty"`
	// The scopes covered by the access token.
	Scopes []string `json:"scopes,omitempty"`
	// The extra properties associated with the access token. This parameter is `null` when no extra property is associated with the issued access token.
	Properties []Property `json:"properties,omitempty"`
	// The newly issued access token in JWT format. If the authorization server is configured to issue JWT-based access tokens (= if the service's `accessTokenSignAlg` value is a non-null value), a JWT-based access token is issued along with the original random-string one.
	JwtAccessToken *string `json:"jwtAccessToken,omitempty"`
	// The resources specified by the `resource` request parameters in the token request. See \"Resource Indicators for OAuth 2.0\" for details.
	Resources []string `json:"resources,omitempty"`
	// The target resources of the access token being issued. See \"Resource Indicators for OAuth 2.0\" for details.
	AccessTokenResources []string      `json:"accessTokenResources,omitempty"`
	AuthorizationDetails *AuthzDetails `json:"authorizationDetails,omitempty"`
	// The attributes of this service that the client application belongs to.
	ServiceAttributes []Pair `json:"serviceAttributes,omitempty"`
	// The attributes of the client.
	ClientAttributes []Pair `json:"clientAttributes,omitempty"`
	// The client authentication method that was performed at the token endpoint.
	ClientAuthMethod *string `json:"clientAuthMethod,omitempty"`
	// the value of the `grant_id` request parameter of the device authorization request.  The `grant_id` request parameter is defined in [Grant Management for OAuth 2.0](https://openid.net/specs/fapi-grant-management.html) , which is supported by Authlete 2.3 and newer versions.
	GrantId *string `json:"grantId,omitempty"`
	// The audiences on the token exchange request
	Audiences          []string   `json:"audiences,omitempty"`
	RequestedTokenType *TokenType `json:"requestedTokenType,omitempty"`
	SubjectToken       *string    `json:"subjectToken,omitempty"`
	SubjectTokenType   *TokenType `json:"subjectTokenType,omitempty"`
	SubjectTokenInfo   *TokenInfo `json:"subjectTokenInfo,omitempty"`
	ActorToken         *string    `json:"actorToken,omitempty"`
	ActorTokenType     *TokenType `json:"actorTokenType,omitempty"`
	ActorTokenInfo     *TokenInfo `json:"actorTokenInfo,omitempty"`
	// For RFC 7523 JSON Web Token (JWT) Profile for OAuth 2.0 Client Authentication and Authorization Grants
	Assertion *string `json:"assertion,omitempty"`
	// Indicate whether the previous refresh token that had been kept in the database for a short time was used
	PreviousRefreshTokenUsed *bool `json:"previousRefreshTokenUsed,omitempty"`
	// The entity ID of the client.
	ClientEntityId *string `json:"clientEntityId,omitempty"`
	// Flag which indicates whether the entity ID of the client was used when the request for the access token was made.
	ClientEntityIdUsed *bool `json:"clientEntityIdUsed,omitempty"`
	// Duration of the `c_nonce` in seconds.
	CnonceDuration *int64 `json:"cnonceDuration,omitempty"`
	// Get the expected nonce value for DPoP proof JWT, which should be used as the value of the `DPoP-Nonce` HTTP header.
	DpopNonce *string `json:"dpopNonce,omitempty"`
	// Get the `c_nonce`.
	Cnonce *string `json:"cnonce,omitempty"`
	// Get the time at which the `c_nonce` expires in milliseconds since the Unix epoch (1970-01-01).
	CnonceExpiresAt *int64 `json:"cnonceExpiresAt,omitempty"`
	// Get the names of the claims that the authorization request (which resulted in generation of the access token) requested to be embedded in ID tokens.
	RequestedIdTokenClaims []string `json:"requestedIdTokenClaims,omitempty"`
	// Scopes associated with the refresh token.
	RefreshTokenScopes []string `json:"refreshTokenScopes,omitempty"`
}

// NewTokenResponse instantiates a new TokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenResponse() *TokenResponse {
	this := TokenResponse{}
	return &this
}

// NewTokenResponseWithDefaults instantiates a new TokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenResponseWithDefaults() *TokenResponse {
	this := TokenResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *TokenResponse) GetResultCode() string {
	if o == nil || IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *TokenResponse) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *TokenResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *TokenResponse) GetResultMessage() string {
	if o == nil || IsNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *TokenResponse) HasResultMessage() bool {
	if o != nil && !IsNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *TokenResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *TokenResponse) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *TokenResponse) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *TokenResponse) SetAction(v string) {
	o.Action = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *TokenResponse) GetResponseContent() string {
	if o == nil || IsNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *TokenResponse) HasResponseContent() bool {
	if o != nil && !IsNil(o.ResponseContent) {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *TokenResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *TokenResponse) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *TokenResponse) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *TokenResponse) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *TokenResponse) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *TokenResponse) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *TokenResponse) SetPassword(v string) {
	o.Password = &v
}

// GetTicket returns the Ticket field value if set, zero value otherwise.
func (o *TokenResponse) GetTicket() string {
	if o == nil || IsNil(o.Ticket) {
		var ret string
		return ret
	}
	return *o.Ticket
}

// GetTicketOk returns a tuple with the Ticket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetTicketOk() (*string, bool) {
	if o == nil || IsNil(o.Ticket) {
		return nil, false
	}
	return o.Ticket, true
}

// HasTicket returns a boolean if a field has been set.
func (o *TokenResponse) HasTicket() bool {
	if o != nil && !IsNil(o.Ticket) {
		return true
	}

	return false
}

// SetTicket gets a reference to the given string and assigns it to the Ticket field.
func (o *TokenResponse) SetTicket(v string) {
	o.Ticket = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *TokenResponse) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *TokenResponse) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *TokenResponse) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetAccessTokenExpiresAt returns the AccessTokenExpiresAt field value if set, zero value otherwise.
func (o *TokenResponse) GetAccessTokenExpiresAt() int64 {
	if o == nil || IsNil(o.AccessTokenExpiresAt) {
		var ret int64
		return ret
	}
	return *o.AccessTokenExpiresAt
}

// GetAccessTokenExpiresAtOk returns a tuple with the AccessTokenExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetAccessTokenExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.AccessTokenExpiresAt) {
		return nil, false
	}
	return o.AccessTokenExpiresAt, true
}

// HasAccessTokenExpiresAt returns a boolean if a field has been set.
func (o *TokenResponse) HasAccessTokenExpiresAt() bool {
	if o != nil && !IsNil(o.AccessTokenExpiresAt) {
		return true
	}

	return false
}

// SetAccessTokenExpiresAt gets a reference to the given int64 and assigns it to the AccessTokenExpiresAt field.
func (o *TokenResponse) SetAccessTokenExpiresAt(v int64) {
	o.AccessTokenExpiresAt = &v
}

// GetAccessTokenDuration returns the AccessTokenDuration field value if set, zero value otherwise.
func (o *TokenResponse) GetAccessTokenDuration() int64 {
	if o == nil || IsNil(o.AccessTokenDuration) {
		var ret int64
		return ret
	}
	return *o.AccessTokenDuration
}

// GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetAccessTokenDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.AccessTokenDuration) {
		return nil, false
	}
	return o.AccessTokenDuration, true
}

// HasAccessTokenDuration returns a boolean if a field has been set.
func (o *TokenResponse) HasAccessTokenDuration() bool {
	if o != nil && !IsNil(o.AccessTokenDuration) {
		return true
	}

	return false
}

// SetAccessTokenDuration gets a reference to the given int64 and assigns it to the AccessTokenDuration field.
func (o *TokenResponse) SetAccessTokenDuration(v int64) {
	o.AccessTokenDuration = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *TokenResponse) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *TokenResponse) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *TokenResponse) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetRefreshTokenExpiresAt returns the RefreshTokenExpiresAt field value if set, zero value otherwise.
func (o *TokenResponse) GetRefreshTokenExpiresAt() int64 {
	if o == nil || IsNil(o.RefreshTokenExpiresAt) {
		var ret int64
		return ret
	}
	return *o.RefreshTokenExpiresAt
}

// GetRefreshTokenExpiresAtOk returns a tuple with the RefreshTokenExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetRefreshTokenExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.RefreshTokenExpiresAt) {
		return nil, false
	}
	return o.RefreshTokenExpiresAt, true
}

// HasRefreshTokenExpiresAt returns a boolean if a field has been set.
func (o *TokenResponse) HasRefreshTokenExpiresAt() bool {
	if o != nil && !IsNil(o.RefreshTokenExpiresAt) {
		return true
	}

	return false
}

// SetRefreshTokenExpiresAt gets a reference to the given int64 and assigns it to the RefreshTokenExpiresAt field.
func (o *TokenResponse) SetRefreshTokenExpiresAt(v int64) {
	o.RefreshTokenExpiresAt = &v
}

// GetRefreshTokenDuration returns the RefreshTokenDuration field value if set, zero value otherwise.
func (o *TokenResponse) GetRefreshTokenDuration() int64 {
	if o == nil || IsNil(o.RefreshTokenDuration) {
		var ret int64
		return ret
	}
	return *o.RefreshTokenDuration
}

// GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetRefreshTokenDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.RefreshTokenDuration) {
		return nil, false
	}
	return o.RefreshTokenDuration, true
}

// HasRefreshTokenDuration returns a boolean if a field has been set.
func (o *TokenResponse) HasRefreshTokenDuration() bool {
	if o != nil && !IsNil(o.RefreshTokenDuration) {
		return true
	}

	return false
}

// SetRefreshTokenDuration gets a reference to the given int64 and assigns it to the RefreshTokenDuration field.
func (o *TokenResponse) SetRefreshTokenDuration(v int64) {
	o.RefreshTokenDuration = &v
}

// GetIdToken returns the IdToken field value if set, zero value otherwise.
func (o *TokenResponse) GetIdToken() string {
	if o == nil || IsNil(o.IdToken) {
		var ret string
		return ret
	}
	return *o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetIdTokenOk() (*string, bool) {
	if o == nil || IsNil(o.IdToken) {
		return nil, false
	}
	return o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *TokenResponse) HasIdToken() bool {
	if o != nil && !IsNil(o.IdToken) {
		return true
	}

	return false
}

// SetIdToken gets a reference to the given string and assigns it to the IdToken field.
func (o *TokenResponse) SetIdToken(v string) {
	o.IdToken = &v
}

// GetGrantType returns the GrantType field value if set, zero value otherwise.
func (o *TokenResponse) GetGrantType() string {
	if o == nil || IsNil(o.GrantType) {
		var ret string
		return ret
	}
	return *o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetGrantTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GrantType) {
		return nil, false
	}
	return o.GrantType, true
}

// HasGrantType returns a boolean if a field has been set.
func (o *TokenResponse) HasGrantType() bool {
	if o != nil && !IsNil(o.GrantType) {
		return true
	}

	return false
}

// SetGrantType gets a reference to the given string and assigns it to the GrantType field.
func (o *TokenResponse) SetGrantType(v string) {
	o.GrantType = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TokenResponse) GetClientId() int64 {
	if o == nil || IsNil(o.ClientId) {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetClientIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TokenResponse) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int64 and assigns it to the ClientId field.
func (o *TokenResponse) SetClientId(v int64) {
	o.ClientId = &v
}

// GetClientIdAlias returns the ClientIdAlias field value if set, zero value otherwise.
func (o *TokenResponse) GetClientIdAlias() string {
	if o == nil || IsNil(o.ClientIdAlias) {
		var ret string
		return ret
	}
	return *o.ClientIdAlias
}

// GetClientIdAliasOk returns a tuple with the ClientIdAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetClientIdAliasOk() (*string, bool) {
	if o == nil || IsNil(o.ClientIdAlias) {
		return nil, false
	}
	return o.ClientIdAlias, true
}

// HasClientIdAlias returns a boolean if a field has been set.
func (o *TokenResponse) HasClientIdAlias() bool {
	if o != nil && !IsNil(o.ClientIdAlias) {
		return true
	}

	return false
}

// SetClientIdAlias gets a reference to the given string and assigns it to the ClientIdAlias field.
func (o *TokenResponse) SetClientIdAlias(v string) {
	o.ClientIdAlias = &v
}

// GetClientIdAliasUsed returns the ClientIdAliasUsed field value if set, zero value otherwise.
func (o *TokenResponse) GetClientIdAliasUsed() bool {
	if o == nil || IsNil(o.ClientIdAliasUsed) {
		var ret bool
		return ret
	}
	return *o.ClientIdAliasUsed
}

// GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetClientIdAliasUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientIdAliasUsed) {
		return nil, false
	}
	return o.ClientIdAliasUsed, true
}

// HasClientIdAliasUsed returns a boolean if a field has been set.
func (o *TokenResponse) HasClientIdAliasUsed() bool {
	if o != nil && !IsNil(o.ClientIdAliasUsed) {
		return true
	}

	return false
}

// SetClientIdAliasUsed gets a reference to the given bool and assigns it to the ClientIdAliasUsed field.
func (o *TokenResponse) SetClientIdAliasUsed(v bool) {
	o.ClientIdAliasUsed = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *TokenResponse) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *TokenResponse) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *TokenResponse) SetSubject(v string) {
	o.Subject = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *TokenResponse) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *TokenResponse) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *TokenResponse) SetScopes(v []string) {
	o.Scopes = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *TokenResponse) GetProperties() []Property {
	if o == nil || IsNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetPropertiesOk() ([]Property, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *TokenResponse) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *TokenResponse) SetProperties(v []Property) {
	o.Properties = v
}

// GetJwtAccessToken returns the JwtAccessToken field value if set, zero value otherwise.
func (o *TokenResponse) GetJwtAccessToken() string {
	if o == nil || IsNil(o.JwtAccessToken) {
		var ret string
		return ret
	}
	return *o.JwtAccessToken
}

// GetJwtAccessTokenOk returns a tuple with the JwtAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetJwtAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.JwtAccessToken) {
		return nil, false
	}
	return o.JwtAccessToken, true
}

// HasJwtAccessToken returns a boolean if a field has been set.
func (o *TokenResponse) HasJwtAccessToken() bool {
	if o != nil && !IsNil(o.JwtAccessToken) {
		return true
	}

	return false
}

// SetJwtAccessToken gets a reference to the given string and assigns it to the JwtAccessToken field.
func (o *TokenResponse) SetJwtAccessToken(v string) {
	o.JwtAccessToken = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *TokenResponse) GetResources() []string {
	if o == nil || IsNil(o.Resources) {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetResourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *TokenResponse) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *TokenResponse) SetResources(v []string) {
	o.Resources = v
}

// GetAccessTokenResources returns the AccessTokenResources field value if set, zero value otherwise.
func (o *TokenResponse) GetAccessTokenResources() []string {
	if o == nil || IsNil(o.AccessTokenResources) {
		var ret []string
		return ret
	}
	return o.AccessTokenResources
}

// GetAccessTokenResourcesOk returns a tuple with the AccessTokenResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetAccessTokenResourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.AccessTokenResources) {
		return nil, false
	}
	return o.AccessTokenResources, true
}

// HasAccessTokenResources returns a boolean if a field has been set.
func (o *TokenResponse) HasAccessTokenResources() bool {
	if o != nil && !IsNil(o.AccessTokenResources) {
		return true
	}

	return false
}

// SetAccessTokenResources gets a reference to the given []string and assigns it to the AccessTokenResources field.
func (o *TokenResponse) SetAccessTokenResources(v []string) {
	o.AccessTokenResources = v
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *TokenResponse) GetAuthorizationDetails() AuthzDetails {
	if o == nil || IsNil(o.AuthorizationDetails) {
		var ret AuthzDetails
		return ret
	}
	return *o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetAuthorizationDetailsOk() (*AuthzDetails, bool) {
	if o == nil || IsNil(o.AuthorizationDetails) {
		return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *TokenResponse) HasAuthorizationDetails() bool {
	if o != nil && !IsNil(o.AuthorizationDetails) {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given AuthzDetails and assigns it to the AuthorizationDetails field.
func (o *TokenResponse) SetAuthorizationDetails(v AuthzDetails) {
	o.AuthorizationDetails = &v
}

// GetServiceAttributes returns the ServiceAttributes field value if set, zero value otherwise.
func (o *TokenResponse) GetServiceAttributes() []Pair {
	if o == nil || IsNil(o.ServiceAttributes) {
		var ret []Pair
		return ret
	}
	return o.ServiceAttributes
}

// GetServiceAttributesOk returns a tuple with the ServiceAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetServiceAttributesOk() ([]Pair, bool) {
	if o == nil || IsNil(o.ServiceAttributes) {
		return nil, false
	}
	return o.ServiceAttributes, true
}

// HasServiceAttributes returns a boolean if a field has been set.
func (o *TokenResponse) HasServiceAttributes() bool {
	if o != nil && !IsNil(o.ServiceAttributes) {
		return true
	}

	return false
}

// SetServiceAttributes gets a reference to the given []Pair and assigns it to the ServiceAttributes field.
func (o *TokenResponse) SetServiceAttributes(v []Pair) {
	o.ServiceAttributes = v
}

// GetClientAttributes returns the ClientAttributes field value if set, zero value otherwise.
func (o *TokenResponse) GetClientAttributes() []Pair {
	if o == nil || IsNil(o.ClientAttributes) {
		var ret []Pair
		return ret
	}
	return o.ClientAttributes
}

// GetClientAttributesOk returns a tuple with the ClientAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetClientAttributesOk() ([]Pair, bool) {
	if o == nil || IsNil(o.ClientAttributes) {
		return nil, false
	}
	return o.ClientAttributes, true
}

// HasClientAttributes returns a boolean if a field has been set.
func (o *TokenResponse) HasClientAttributes() bool {
	if o != nil && !IsNil(o.ClientAttributes) {
		return true
	}

	return false
}

// SetClientAttributes gets a reference to the given []Pair and assigns it to the ClientAttributes field.
func (o *TokenResponse) SetClientAttributes(v []Pair) {
	o.ClientAttributes = v
}

// GetClientAuthMethod returns the ClientAuthMethod field value if set, zero value otherwise.
func (o *TokenResponse) GetClientAuthMethod() string {
	if o == nil || IsNil(o.ClientAuthMethod) {
		var ret string
		return ret
	}
	return *o.ClientAuthMethod
}

// GetClientAuthMethodOk returns a tuple with the ClientAuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetClientAuthMethodOk() (*string, bool) {
	if o == nil || IsNil(o.ClientAuthMethod) {
		return nil, false
	}
	return o.ClientAuthMethod, true
}

// HasClientAuthMethod returns a boolean if a field has been set.
func (o *TokenResponse) HasClientAuthMethod() bool {
	if o != nil && !IsNil(o.ClientAuthMethod) {
		return true
	}

	return false
}

// SetClientAuthMethod gets a reference to the given string and assigns it to the ClientAuthMethod field.
func (o *TokenResponse) SetClientAuthMethod(v string) {
	o.ClientAuthMethod = &v
}

// GetGrantId returns the GrantId field value if set, zero value otherwise.
func (o *TokenResponse) GetGrantId() string {
	if o == nil || IsNil(o.GrantId) {
		var ret string
		return ret
	}
	return *o.GrantId
}

// GetGrantIdOk returns a tuple with the GrantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetGrantIdOk() (*string, bool) {
	if o == nil || IsNil(o.GrantId) {
		return nil, false
	}
	return o.GrantId, true
}

// HasGrantId returns a boolean if a field has been set.
func (o *TokenResponse) HasGrantId() bool {
	if o != nil && !IsNil(o.GrantId) {
		return true
	}

	return false
}

// SetGrantId gets a reference to the given string and assigns it to the GrantId field.
func (o *TokenResponse) SetGrantId(v string) {
	o.GrantId = &v
}

// GetAudiences returns the Audiences field value if set, zero value otherwise.
func (o *TokenResponse) GetAudiences() []string {
	if o == nil || IsNil(o.Audiences) {
		var ret []string
		return ret
	}
	return o.Audiences
}

// GetAudiencesOk returns a tuple with the Audiences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetAudiencesOk() ([]string, bool) {
	if o == nil || IsNil(o.Audiences) {
		return nil, false
	}
	return o.Audiences, true
}

// HasAudiences returns a boolean if a field has been set.
func (o *TokenResponse) HasAudiences() bool {
	if o != nil && !IsNil(o.Audiences) {
		return true
	}

	return false
}

// SetAudiences gets a reference to the given []string and assigns it to the Audiences field.
func (o *TokenResponse) SetAudiences(v []string) {
	o.Audiences = v
}

// GetRequestedTokenType returns the RequestedTokenType field value if set, zero value otherwise.
func (o *TokenResponse) GetRequestedTokenType() TokenType {
	if o == nil || IsNil(o.RequestedTokenType) {
		var ret TokenType
		return ret
	}
	return *o.RequestedTokenType
}

// GetRequestedTokenTypeOk returns a tuple with the RequestedTokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetRequestedTokenTypeOk() (*TokenType, bool) {
	if o == nil || IsNil(o.RequestedTokenType) {
		return nil, false
	}
	return o.RequestedTokenType, true
}

// HasRequestedTokenType returns a boolean if a field has been set.
func (o *TokenResponse) HasRequestedTokenType() bool {
	if o != nil && !IsNil(o.RequestedTokenType) {
		return true
	}

	return false
}

// SetRequestedTokenType gets a reference to the given TokenType and assigns it to the RequestedTokenType field.
func (o *TokenResponse) SetRequestedTokenType(v TokenType) {
	o.RequestedTokenType = &v
}

// GetSubjectToken returns the SubjectToken field value if set, zero value otherwise.
func (o *TokenResponse) GetSubjectToken() string {
	if o == nil || IsNil(o.SubjectToken) {
		var ret string
		return ret
	}
	return *o.SubjectToken
}

// GetSubjectTokenOk returns a tuple with the SubjectToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetSubjectTokenOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectToken) {
		return nil, false
	}
	return o.SubjectToken, true
}

// HasSubjectToken returns a boolean if a field has been set.
func (o *TokenResponse) HasSubjectToken() bool {
	if o != nil && !IsNil(o.SubjectToken) {
		return true
	}

	return false
}

// SetSubjectToken gets a reference to the given string and assigns it to the SubjectToken field.
func (o *TokenResponse) SetSubjectToken(v string) {
	o.SubjectToken = &v
}

// GetSubjectTokenType returns the SubjectTokenType field value if set, zero value otherwise.
func (o *TokenResponse) GetSubjectTokenType() TokenType {
	if o == nil || IsNil(o.SubjectTokenType) {
		var ret TokenType
		return ret
	}
	return *o.SubjectTokenType
}

// GetSubjectTokenTypeOk returns a tuple with the SubjectTokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetSubjectTokenTypeOk() (*TokenType, bool) {
	if o == nil || IsNil(o.SubjectTokenType) {
		return nil, false
	}
	return o.SubjectTokenType, true
}

// HasSubjectTokenType returns a boolean if a field has been set.
func (o *TokenResponse) HasSubjectTokenType() bool {
	if o != nil && !IsNil(o.SubjectTokenType) {
		return true
	}

	return false
}

// SetSubjectTokenType gets a reference to the given TokenType and assigns it to the SubjectTokenType field.
func (o *TokenResponse) SetSubjectTokenType(v TokenType) {
	o.SubjectTokenType = &v
}

// GetSubjectTokenInfo returns the SubjectTokenInfo field value if set, zero value otherwise.
func (o *TokenResponse) GetSubjectTokenInfo() TokenInfo {
	if o == nil || IsNil(o.SubjectTokenInfo) {
		var ret TokenInfo
		return ret
	}
	return *o.SubjectTokenInfo
}

// GetSubjectTokenInfoOk returns a tuple with the SubjectTokenInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetSubjectTokenInfoOk() (*TokenInfo, bool) {
	if o == nil || IsNil(o.SubjectTokenInfo) {
		return nil, false
	}
	return o.SubjectTokenInfo, true
}

// HasSubjectTokenInfo returns a boolean if a field has been set.
func (o *TokenResponse) HasSubjectTokenInfo() bool {
	if o != nil && !IsNil(o.SubjectTokenInfo) {
		return true
	}

	return false
}

// SetSubjectTokenInfo gets a reference to the given TokenInfo and assigns it to the SubjectTokenInfo field.
func (o *TokenResponse) SetSubjectTokenInfo(v TokenInfo) {
	o.SubjectTokenInfo = &v
}

// GetActorToken returns the ActorToken field value if set, zero value otherwise.
func (o *TokenResponse) GetActorToken() string {
	if o == nil || IsNil(o.ActorToken) {
		var ret string
		return ret
	}
	return *o.ActorToken
}

// GetActorTokenOk returns a tuple with the ActorToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetActorTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ActorToken) {
		return nil, false
	}
	return o.ActorToken, true
}

// HasActorToken returns a boolean if a field has been set.
func (o *TokenResponse) HasActorToken() bool {
	if o != nil && !IsNil(o.ActorToken) {
		return true
	}

	return false
}

// SetActorToken gets a reference to the given string and assigns it to the ActorToken field.
func (o *TokenResponse) SetActorToken(v string) {
	o.ActorToken = &v
}

// GetActorTokenType returns the ActorTokenType field value if set, zero value otherwise.
func (o *TokenResponse) GetActorTokenType() TokenType {
	if o == nil || IsNil(o.ActorTokenType) {
		var ret TokenType
		return ret
	}
	return *o.ActorTokenType
}

// GetActorTokenTypeOk returns a tuple with the ActorTokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetActorTokenTypeOk() (*TokenType, bool) {
	if o == nil || IsNil(o.ActorTokenType) {
		return nil, false
	}
	return o.ActorTokenType, true
}

// HasActorTokenType returns a boolean if a field has been set.
func (o *TokenResponse) HasActorTokenType() bool {
	if o != nil && !IsNil(o.ActorTokenType) {
		return true
	}

	return false
}

// SetActorTokenType gets a reference to the given TokenType and assigns it to the ActorTokenType field.
func (o *TokenResponse) SetActorTokenType(v TokenType) {
	o.ActorTokenType = &v
}

// GetActorTokenInfo returns the ActorTokenInfo field value if set, zero value otherwise.
func (o *TokenResponse) GetActorTokenInfo() TokenInfo {
	if o == nil || IsNil(o.ActorTokenInfo) {
		var ret TokenInfo
		return ret
	}
	return *o.ActorTokenInfo
}

// GetActorTokenInfoOk returns a tuple with the ActorTokenInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetActorTokenInfoOk() (*TokenInfo, bool) {
	if o == nil || IsNil(o.ActorTokenInfo) {
		return nil, false
	}
	return o.ActorTokenInfo, true
}

// HasActorTokenInfo returns a boolean if a field has been set.
func (o *TokenResponse) HasActorTokenInfo() bool {
	if o != nil && !IsNil(o.ActorTokenInfo) {
		return true
	}

	return false
}

// SetActorTokenInfo gets a reference to the given TokenInfo and assigns it to the ActorTokenInfo field.
func (o *TokenResponse) SetActorTokenInfo(v TokenInfo) {
	o.ActorTokenInfo = &v
}

// GetAssertion returns the Assertion field value if set, zero value otherwise.
func (o *TokenResponse) GetAssertion() string {
	if o == nil || IsNil(o.Assertion) {
		var ret string
		return ret
	}
	return *o.Assertion
}

// GetAssertionOk returns a tuple with the Assertion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetAssertionOk() (*string, bool) {
	if o == nil || IsNil(o.Assertion) {
		return nil, false
	}
	return o.Assertion, true
}

// HasAssertion returns a boolean if a field has been set.
func (o *TokenResponse) HasAssertion() bool {
	if o != nil && !IsNil(o.Assertion) {
		return true
	}

	return false
}

// SetAssertion gets a reference to the given string and assigns it to the Assertion field.
func (o *TokenResponse) SetAssertion(v string) {
	o.Assertion = &v
}

// GetPreviousRefreshTokenUsed returns the PreviousRefreshTokenUsed field value if set, zero value otherwise.
func (o *TokenResponse) GetPreviousRefreshTokenUsed() bool {
	if o == nil || IsNil(o.PreviousRefreshTokenUsed) {
		var ret bool
		return ret
	}
	return *o.PreviousRefreshTokenUsed
}

// GetPreviousRefreshTokenUsedOk returns a tuple with the PreviousRefreshTokenUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetPreviousRefreshTokenUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.PreviousRefreshTokenUsed) {
		return nil, false
	}
	return o.PreviousRefreshTokenUsed, true
}

// HasPreviousRefreshTokenUsed returns a boolean if a field has been set.
func (o *TokenResponse) HasPreviousRefreshTokenUsed() bool {
	if o != nil && !IsNil(o.PreviousRefreshTokenUsed) {
		return true
	}

	return false
}

// SetPreviousRefreshTokenUsed gets a reference to the given bool and assigns it to the PreviousRefreshTokenUsed field.
func (o *TokenResponse) SetPreviousRefreshTokenUsed(v bool) {
	o.PreviousRefreshTokenUsed = &v
}

// GetClientEntityId returns the ClientEntityId field value if set, zero value otherwise.
func (o *TokenResponse) GetClientEntityId() string {
	if o == nil || IsNil(o.ClientEntityId) {
		var ret string
		return ret
	}
	return *o.ClientEntityId
}

// GetClientEntityIdOk returns a tuple with the ClientEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetClientEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientEntityId) {
		return nil, false
	}
	return o.ClientEntityId, true
}

// HasClientEntityId returns a boolean if a field has been set.
func (o *TokenResponse) HasClientEntityId() bool {
	if o != nil && !IsNil(o.ClientEntityId) {
		return true
	}

	return false
}

// SetClientEntityId gets a reference to the given string and assigns it to the ClientEntityId field.
func (o *TokenResponse) SetClientEntityId(v string) {
	o.ClientEntityId = &v
}

// GetClientEntityIdUsed returns the ClientEntityIdUsed field value if set, zero value otherwise.
func (o *TokenResponse) GetClientEntityIdUsed() bool {
	if o == nil || IsNil(o.ClientEntityIdUsed) {
		var ret bool
		return ret
	}
	return *o.ClientEntityIdUsed
}

// GetClientEntityIdUsedOk returns a tuple with the ClientEntityIdUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetClientEntityIdUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientEntityIdUsed) {
		return nil, false
	}
	return o.ClientEntityIdUsed, true
}

// HasClientEntityIdUsed returns a boolean if a field has been set.
func (o *TokenResponse) HasClientEntityIdUsed() bool {
	if o != nil && !IsNil(o.ClientEntityIdUsed) {
		return true
	}

	return false
}

// SetClientEntityIdUsed gets a reference to the given bool and assigns it to the ClientEntityIdUsed field.
func (o *TokenResponse) SetClientEntityIdUsed(v bool) {
	o.ClientEntityIdUsed = &v
}

// GetCnonceDuration returns the CnonceDuration field value if set, zero value otherwise.
func (o *TokenResponse) GetCnonceDuration() int64 {
	if o == nil || IsNil(o.CnonceDuration) {
		var ret int64
		return ret
	}
	return *o.CnonceDuration
}

// GetCnonceDurationOk returns a tuple with the CnonceDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetCnonceDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.CnonceDuration) {
		return nil, false
	}
	return o.CnonceDuration, true
}

// HasCnonceDuration returns a boolean if a field has been set.
func (o *TokenResponse) HasCnonceDuration() bool {
	if o != nil && !IsNil(o.CnonceDuration) {
		return true
	}

	return false
}

// SetCnonceDuration gets a reference to the given int64 and assigns it to the CnonceDuration field.
func (o *TokenResponse) SetCnonceDuration(v int64) {
	o.CnonceDuration = &v
}

// GetDpopNonce returns the DpopNonce field value if set, zero value otherwise.
func (o *TokenResponse) GetDpopNonce() string {
	if o == nil || IsNil(o.DpopNonce) {
		var ret string
		return ret
	}
	return *o.DpopNonce
}

// GetDpopNonceOk returns a tuple with the DpopNonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetDpopNonceOk() (*string, bool) {
	if o == nil || IsNil(o.DpopNonce) {
		return nil, false
	}
	return o.DpopNonce, true
}

// HasDpopNonce returns a boolean if a field has been set.
func (o *TokenResponse) HasDpopNonce() bool {
	if o != nil && !IsNil(o.DpopNonce) {
		return true
	}

	return false
}

// SetDpopNonce gets a reference to the given string and assigns it to the DpopNonce field.
func (o *TokenResponse) SetDpopNonce(v string) {
	o.DpopNonce = &v
}

// GetCnonce returns the Cnonce field value if set, zero value otherwise.
func (o *TokenResponse) GetCnonce() string {
	if o == nil || IsNil(o.Cnonce) {
		var ret string
		return ret
	}
	return *o.Cnonce
}

// GetCnonceOk returns a tuple with the Cnonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetCnonceOk() (*string, bool) {
	if o == nil || IsNil(o.Cnonce) {
		return nil, false
	}
	return o.Cnonce, true
}

// HasCnonce returns a boolean if a field has been set.
func (o *TokenResponse) HasCnonce() bool {
	if o != nil && !IsNil(o.Cnonce) {
		return true
	}

	return false
}

// SetCnonce gets a reference to the given string and assigns it to the Cnonce field.
func (o *TokenResponse) SetCnonce(v string) {
	o.Cnonce = &v
}

// GetCnonceExpiresAt returns the CnonceExpiresAt field value if set, zero value otherwise.
func (o *TokenResponse) GetCnonceExpiresAt() int64 {
	if o == nil || IsNil(o.CnonceExpiresAt) {
		var ret int64
		return ret
	}
	return *o.CnonceExpiresAt
}

// GetCnonceExpiresAtOk returns a tuple with the CnonceExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetCnonceExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CnonceExpiresAt) {
		return nil, false
	}
	return o.CnonceExpiresAt, true
}

// HasCnonceExpiresAt returns a boolean if a field has been set.
func (o *TokenResponse) HasCnonceExpiresAt() bool {
	if o != nil && !IsNil(o.CnonceExpiresAt) {
		return true
	}

	return false
}

// SetCnonceExpiresAt gets a reference to the given int64 and assigns it to the CnonceExpiresAt field.
func (o *TokenResponse) SetCnonceExpiresAt(v int64) {
	o.CnonceExpiresAt = &v
}

// GetRequestedIdTokenClaims returns the RequestedIdTokenClaims field value if set, zero value otherwise.
func (o *TokenResponse) GetRequestedIdTokenClaims() []string {
	if o == nil || IsNil(o.RequestedIdTokenClaims) {
		var ret []string
		return ret
	}
	return o.RequestedIdTokenClaims
}

// GetRequestedIdTokenClaimsOk returns a tuple with the RequestedIdTokenClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetRequestedIdTokenClaimsOk() ([]string, bool) {
	if o == nil || IsNil(o.RequestedIdTokenClaims) {
		return nil, false
	}
	return o.RequestedIdTokenClaims, true
}

// HasRequestedIdTokenClaims returns a boolean if a field has been set.
func (o *TokenResponse) HasRequestedIdTokenClaims() bool {
	if o != nil && !IsNil(o.RequestedIdTokenClaims) {
		return true
	}

	return false
}

// SetRequestedIdTokenClaims gets a reference to the given []string and assigns it to the RequestedIdTokenClaims field.
func (o *TokenResponse) SetRequestedIdTokenClaims(v []string) {
	o.RequestedIdTokenClaims = v
}

// GetRefreshTokenScopes returns the RefreshTokenScopes field value if set, zero value otherwise.
func (o *TokenResponse) GetRefreshTokenScopes() []string {
	if o == nil || IsNil(o.RefreshTokenScopes) {
		var ret []string
		return ret
	}
	return o.RefreshTokenScopes
}

// GetRefreshTokenScopesOk returns a tuple with the RefreshTokenScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetRefreshTokenScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.RefreshTokenScopes) {
		return nil, false
	}
	return o.RefreshTokenScopes, true
}

// HasRefreshTokenScopes returns a boolean if a field has been set.
func (o *TokenResponse) HasRefreshTokenScopes() bool {
	if o != nil && !IsNil(o.RefreshTokenScopes) {
		return true
	}

	return false
}

// SetRefreshTokenScopes gets a reference to the given []string and assigns it to the RefreshTokenScopes field.
func (o *TokenResponse) SetRefreshTokenScopes(v []string) {
	o.RefreshTokenScopes = v
}

func (o TokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Ticket) {
		toSerialize["ticket"] = o.Ticket
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
	if !IsNil(o.IdToken) {
		toSerialize["idToken"] = o.IdToken
	}
	if !IsNil(o.GrantType) {
		toSerialize["grantType"] = o.GrantType
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
	if !IsNil(o.ClientAuthMethod) {
		toSerialize["clientAuthMethod"] = o.ClientAuthMethod
	}
	if !IsNil(o.GrantId) {
		toSerialize["grantId"] = o.GrantId
	}
	if !IsNil(o.Audiences) {
		toSerialize["audiences"] = o.Audiences
	}
	if !IsNil(o.RequestedTokenType) {
		toSerialize["requestedTokenType"] = o.RequestedTokenType
	}
	if !IsNil(o.SubjectToken) {
		toSerialize["subjectToken"] = o.SubjectToken
	}
	if !IsNil(o.SubjectTokenType) {
		toSerialize["subjectTokenType"] = o.SubjectTokenType
	}
	if !IsNil(o.SubjectTokenInfo) {
		toSerialize["subjectTokenInfo"] = o.SubjectTokenInfo
	}
	if !IsNil(o.ActorToken) {
		toSerialize["actorToken"] = o.ActorToken
	}
	if !IsNil(o.ActorTokenType) {
		toSerialize["actorTokenType"] = o.ActorTokenType
	}
	if !IsNil(o.ActorTokenInfo) {
		toSerialize["actorTokenInfo"] = o.ActorTokenInfo
	}
	if !IsNil(o.Assertion) {
		toSerialize["assertion"] = o.Assertion
	}
	if !IsNil(o.PreviousRefreshTokenUsed) {
		toSerialize["previousRefreshTokenUsed"] = o.PreviousRefreshTokenUsed
	}
	if !IsNil(o.ClientEntityId) {
		toSerialize["clientEntityId"] = o.ClientEntityId
	}
	if !IsNil(o.ClientEntityIdUsed) {
		toSerialize["clientEntityIdUsed"] = o.ClientEntityIdUsed
	}
	if !IsNil(o.CnonceDuration) {
		toSerialize["cnonceDuration"] = o.CnonceDuration
	}
	if !IsNil(o.DpopNonce) {
		toSerialize["dpopNonce"] = o.DpopNonce
	}
	if !IsNil(o.Cnonce) {
		toSerialize["cnonce"] = o.Cnonce
	}
	if !IsNil(o.CnonceExpiresAt) {
		toSerialize["cnonceExpiresAt"] = o.CnonceExpiresAt
	}
	if !IsNil(o.RequestedIdTokenClaims) {
		toSerialize["requestedIdTokenClaims"] = o.RequestedIdTokenClaims
	}
	if !IsNil(o.RefreshTokenScopes) {
		toSerialize["refreshTokenScopes"] = o.RefreshTokenScopes
	}
	return toSerialize, nil
}

type NullableTokenResponse struct {
	value *TokenResponse
	isSet bool
}

func (v NullableTokenResponse) Get() *TokenResponse {
	return v.value
}

func (v *NullableTokenResponse) Set(val *TokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenResponse(val *TokenResponse) *NullableTokenResponse {
	return &NullableTokenResponse{value: val, isSet: true}
}

func (v NullableTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
