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

// checks if the TokenCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenCreateResponse{}

// TokenCreateResponse struct for TokenCreateResponse
type TokenCreateResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take.
	Action *string `json:"action,omitempty"`
	// The newly issued access token.
	AccessToken *string `json:"accessToken,omitempty"`
	// The ID of the client application associated with the access token.
	ClientId *int64 `json:"clientId,omitempty"`
	// The time at which the access token expires.
	ExpiresAt *int64 `json:"expiresAt,omitempty"`
	// The duration of the newly issued access token in seconds.
	ExpiresIn *int64 `json:"expiresIn,omitempty"`
	// The grant type for the newly issued access token.
	GrantType *string `json:"grantType,omitempty"`
	// The extra properties associated with the access token.
	Properties []Property `json:"properties,omitempty"`
	// The newly issued refresh token.
	RefreshToken *string `json:"refreshToken,omitempty"`
	// Scopes which are associated with the access token.
	Scopes []string `json:"scopes,omitempty"`
	// The subject (= unique identifier) of the user associated with the newly issued access token.
	Subject *string `json:"subject,omitempty"`
	// The token type of the access token.
	TokenType *string `json:"tokenType,omitempty"`
	// If the authorization server is configured to issue JWT-based access tokens (= if `Service.accessTokenSignAlg` is set to a `non-null` value), a JWT-based access token is issued along with the original random-string one.
	JwtAccessToken       *string       `json:"jwtAccessToken,omitempty"`
	AuthorizationDetails *AuthzDetails `json:"authorizationDetails,omitempty"`
	// the flag which indicates whether the access token is for an external attachment.
	ForExternalAttachment *bool `json:"forExternalAttachment,omitempty"`
	// Set the unique token identifier.
	TokenId *string `json:"tokenId,omitempty"`
	// The scopes associated with the refresh token. May be null.
	RefreshTokenScopes []string `json:"refreshTokenScopes,omitempty"`
}

// NewTokenCreateResponse instantiates a new TokenCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenCreateResponse() *TokenCreateResponse {
	this := TokenCreateResponse{}
	return &this
}

// NewTokenCreateResponseWithDefaults instantiates a new TokenCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenCreateResponseWithDefaults() *TokenCreateResponse {
	this := TokenCreateResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *TokenCreateResponse) GetResultCode() string {
	if o == nil || IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *TokenCreateResponse) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *TokenCreateResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *TokenCreateResponse) GetResultMessage() string {
	if o == nil || IsNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *TokenCreateResponse) HasResultMessage() bool {
	if o != nil && !IsNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *TokenCreateResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *TokenCreateResponse) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateResponse) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *TokenCreateResponse) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *TokenCreateResponse) SetAction(v string) {
	o.Action = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *TokenCreateResponse) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *TokenCreateResponse) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *TokenCreateResponse) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TokenCreateResponse) GetClientId() int64 {
	if o == nil || IsNil(o.ClientId) {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateResponse) GetClientIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TokenCreateResponse) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int64 and assigns it to the ClientId field.
func (o *TokenCreateResponse) SetClientId(v int64) {
	o.ClientId = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *TokenCreateResponse) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateResponse) GetExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *TokenCreateResponse) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *TokenCreateResponse) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *TokenCreateResponse) GetExpiresIn() int64 {
	if o == nil || IsNil(o.ExpiresIn) {
		var ret int64
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateResponse) GetExpiresInOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *TokenCreateResponse) HasExpiresIn() bool {
	if o != nil && !IsNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int64 and assigns it to the ExpiresIn field.
func (o *TokenCreateResponse) SetExpiresIn(v int64) {
	o.ExpiresIn = &v
}

// GetGrantType returns the GrantType field value if set, zero value otherwise.
func (o *TokenCreateResponse) GetGrantType() string {
	if o == nil || IsNil(o.GrantType) {
		var ret string
		return ret
	}
	return *o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateResponse) GetGrantTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GrantType) {
		return nil, false
	}
	return o.GrantType, true
}

// HasGrantType returns a boolean if a field has been set.
func (o *TokenCreateResponse) HasGrantType() bool {
	if o != nil && !IsNil(o.GrantType) {
		return true
	}

	return false
}

// SetGrantType gets a reference to the given string and assigns it to the GrantType field.
func (o *TokenCreateResponse) SetGrantType(v string) {
	o.GrantType = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *TokenCreateResponse) GetProperties() []Property {
	if o == nil || IsNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateResponse) GetPropertiesOk() ([]Property, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *TokenCreateResponse) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *TokenCreateResponse) SetProperties(v []Property) {
	o.Properties = v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *TokenCreateResponse) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *TokenCreateResponse) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *TokenCreateResponse) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *TokenCreateResponse) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateResponse) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *TokenCreateResponse) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *TokenCreateResponse) SetScopes(v []string) {
	o.Scopes = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *TokenCreateResponse) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateResponse) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *TokenCreateResponse) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *TokenCreateResponse) SetSubject(v string) {
	o.Subject = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *TokenCreateResponse) GetTokenType() string {
	if o == nil || IsNil(o.TokenType) {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TokenType) {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *TokenCreateResponse) HasTokenType() bool {
	if o != nil && !IsNil(o.TokenType) {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *TokenCreateResponse) SetTokenType(v string) {
	o.TokenType = &v
}

// GetJwtAccessToken returns the JwtAccessToken field value if set, zero value otherwise.
func (o *TokenCreateResponse) GetJwtAccessToken() string {
	if o == nil || IsNil(o.JwtAccessToken) {
		var ret string
		return ret
	}
	return *o.JwtAccessToken
}

// GetJwtAccessTokenOk returns a tuple with the JwtAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateResponse) GetJwtAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.JwtAccessToken) {
		return nil, false
	}
	return o.JwtAccessToken, true
}

// HasJwtAccessToken returns a boolean if a field has been set.
func (o *TokenCreateResponse) HasJwtAccessToken() bool {
	if o != nil && !IsNil(o.JwtAccessToken) {
		return true
	}

	return false
}

// SetJwtAccessToken gets a reference to the given string and assigns it to the JwtAccessToken field.
func (o *TokenCreateResponse) SetJwtAccessToken(v string) {
	o.JwtAccessToken = &v
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *TokenCreateResponse) GetAuthorizationDetails() AuthzDetails {
	if o == nil || IsNil(o.AuthorizationDetails) {
		var ret AuthzDetails
		return ret
	}
	return *o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateResponse) GetAuthorizationDetailsOk() (*AuthzDetails, bool) {
	if o == nil || IsNil(o.AuthorizationDetails) {
		return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *TokenCreateResponse) HasAuthorizationDetails() bool {
	if o != nil && !IsNil(o.AuthorizationDetails) {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given AuthzDetails and assigns it to the AuthorizationDetails field.
func (o *TokenCreateResponse) SetAuthorizationDetails(v AuthzDetails) {
	o.AuthorizationDetails = &v
}

// GetForExternalAttachment returns the ForExternalAttachment field value if set, zero value otherwise.
func (o *TokenCreateResponse) GetForExternalAttachment() bool {
	if o == nil || IsNil(o.ForExternalAttachment) {
		var ret bool
		return ret
	}
	return *o.ForExternalAttachment
}

// GetForExternalAttachmentOk returns a tuple with the ForExternalAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateResponse) GetForExternalAttachmentOk() (*bool, bool) {
	if o == nil || IsNil(o.ForExternalAttachment) {
		return nil, false
	}
	return o.ForExternalAttachment, true
}

// HasForExternalAttachment returns a boolean if a field has been set.
func (o *TokenCreateResponse) HasForExternalAttachment() bool {
	if o != nil && !IsNil(o.ForExternalAttachment) {
		return true
	}

	return false
}

// SetForExternalAttachment gets a reference to the given bool and assigns it to the ForExternalAttachment field.
func (o *TokenCreateResponse) SetForExternalAttachment(v bool) {
	o.ForExternalAttachment = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *TokenCreateResponse) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateResponse) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *TokenCreateResponse) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *TokenCreateResponse) SetTokenId(v string) {
	o.TokenId = &v
}

// GetRefreshTokenScopes returns the RefreshTokenScopes field value if set, zero value otherwise.
func (o *TokenCreateResponse) GetRefreshTokenScopes() []string {
	if o == nil || IsNil(o.RefreshTokenScopes) {
		var ret []string
		return ret
	}
	return o.RefreshTokenScopes
}

// GetRefreshTokenScopesOk returns a tuple with the RefreshTokenScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateResponse) GetRefreshTokenScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.RefreshTokenScopes) {
		return nil, false
	}
	return o.RefreshTokenScopes, true
}

// HasRefreshTokenScopes returns a boolean if a field has been set.
func (o *TokenCreateResponse) HasRefreshTokenScopes() bool {
	if o != nil && !IsNil(o.RefreshTokenScopes) {
		return true
	}

	return false
}

// SetRefreshTokenScopes gets a reference to the given []string and assigns it to the RefreshTokenScopes field.
func (o *TokenCreateResponse) SetRefreshTokenScopes(v []string) {
	o.RefreshTokenScopes = v
}

func (o TokenCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenCreateResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !IsNil(o.ExpiresIn) {
		toSerialize["expiresIn"] = o.ExpiresIn
	}
	if !IsNil(o.GrantType) {
		toSerialize["grantType"] = o.GrantType
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refreshToken"] = o.RefreshToken
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.TokenType) {
		toSerialize["tokenType"] = o.TokenType
	}
	if !IsNil(o.JwtAccessToken) {
		toSerialize["jwtAccessToken"] = o.JwtAccessToken
	}
	if !IsNil(o.AuthorizationDetails) {
		toSerialize["authorizationDetails"] = o.AuthorizationDetails
	}
	if !IsNil(o.ForExternalAttachment) {
		toSerialize["forExternalAttachment"] = o.ForExternalAttachment
	}
	if !IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}
	if !IsNil(o.RefreshTokenScopes) {
		toSerialize["refreshTokenScopes"] = o.RefreshTokenScopes
	}
	return toSerialize, nil
}

type NullableTokenCreateResponse struct {
	value *TokenCreateResponse
	isSet bool
}

func (v NullableTokenCreateResponse) Get() *TokenCreateResponse {
	return v.value
}

func (v *NullableTokenCreateResponse) Set(val *TokenCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenCreateResponse(val *TokenCreateResponse) *NullableTokenCreateResponse {
	return &NullableTokenCreateResponse{value: val, isSet: true}
}

func (v NullableTokenCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
