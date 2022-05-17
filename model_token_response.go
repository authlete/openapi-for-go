/*
Authlete API

Authlete API Document. 

API version: 2.2.19
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
)

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
	AccessTokenResources []string `json:"accessTokenResources,omitempty"`
	AuthorizationDetails *AuthorizationDetails `json:"authorizationDetails,omitempty"`
	// The attributes of this service that the client application belongs to. 
	ServiceAttributes []Pair `json:"serviceAttributes,omitempty"`
	// The attributes of the client. 
	ClientAttributes []Pair `json:"clientAttributes,omitempty"`
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
	if o == nil || o.ResultCode == nil {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || o.ResultCode == nil {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *TokenResponse) HasResultCode() bool {
	if o != nil && o.ResultCode != nil {
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
	if o == nil || o.ResultMessage == nil {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || o.ResultMessage == nil {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *TokenResponse) HasResultMessage() bool {
	if o != nil && o.ResultMessage != nil {
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
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *TokenResponse) HasAction() bool {
	if o != nil && o.Action != nil {
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
	if o == nil || o.ResponseContent == nil {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || o.ResponseContent == nil {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *TokenResponse) HasResponseContent() bool {
	if o != nil && o.ResponseContent != nil {
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
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *TokenResponse) HasUsername() bool {
	if o != nil && o.Username != nil {
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
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *TokenResponse) HasPassword() bool {
	if o != nil && o.Password != nil {
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
	if o == nil || o.Ticket == nil {
		var ret string
		return ret
	}
	return *o.Ticket
}

// GetTicketOk returns a tuple with the Ticket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetTicketOk() (*string, bool) {
	if o == nil || o.Ticket == nil {
		return nil, false
	}
	return o.Ticket, true
}

// HasTicket returns a boolean if a field has been set.
func (o *TokenResponse) HasTicket() bool {
	if o != nil && o.Ticket != nil {
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
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *TokenResponse) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
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
	if o == nil || o.AccessTokenExpiresAt == nil {
		var ret int64
		return ret
	}
	return *o.AccessTokenExpiresAt
}

// GetAccessTokenExpiresAtOk returns a tuple with the AccessTokenExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetAccessTokenExpiresAtOk() (*int64, bool) {
	if o == nil || o.AccessTokenExpiresAt == nil {
		return nil, false
	}
	return o.AccessTokenExpiresAt, true
}

// HasAccessTokenExpiresAt returns a boolean if a field has been set.
func (o *TokenResponse) HasAccessTokenExpiresAt() bool {
	if o != nil && o.AccessTokenExpiresAt != nil {
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
	if o == nil || o.AccessTokenDuration == nil {
		var ret int64
		return ret
	}
	return *o.AccessTokenDuration
}

// GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetAccessTokenDurationOk() (*int64, bool) {
	if o == nil || o.AccessTokenDuration == nil {
		return nil, false
	}
	return o.AccessTokenDuration, true
}

// HasAccessTokenDuration returns a boolean if a field has been set.
func (o *TokenResponse) HasAccessTokenDuration() bool {
	if o != nil && o.AccessTokenDuration != nil {
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
	if o == nil || o.RefreshToken == nil {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil || o.RefreshToken == nil {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *TokenResponse) HasRefreshToken() bool {
	if o != nil && o.RefreshToken != nil {
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
	if o == nil || o.RefreshTokenExpiresAt == nil {
		var ret int64
		return ret
	}
	return *o.RefreshTokenExpiresAt
}

// GetRefreshTokenExpiresAtOk returns a tuple with the RefreshTokenExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetRefreshTokenExpiresAtOk() (*int64, bool) {
	if o == nil || o.RefreshTokenExpiresAt == nil {
		return nil, false
	}
	return o.RefreshTokenExpiresAt, true
}

// HasRefreshTokenExpiresAt returns a boolean if a field has been set.
func (o *TokenResponse) HasRefreshTokenExpiresAt() bool {
	if o != nil && o.RefreshTokenExpiresAt != nil {
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
	if o == nil || o.RefreshTokenDuration == nil {
		var ret int64
		return ret
	}
	return *o.RefreshTokenDuration
}

// GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetRefreshTokenDurationOk() (*int64, bool) {
	if o == nil || o.RefreshTokenDuration == nil {
		return nil, false
	}
	return o.RefreshTokenDuration, true
}

// HasRefreshTokenDuration returns a boolean if a field has been set.
func (o *TokenResponse) HasRefreshTokenDuration() bool {
	if o != nil && o.RefreshTokenDuration != nil {
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
	if o == nil || o.IdToken == nil {
		var ret string
		return ret
	}
	return *o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetIdTokenOk() (*string, bool) {
	if o == nil || o.IdToken == nil {
		return nil, false
	}
	return o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *TokenResponse) HasIdToken() bool {
	if o != nil && o.IdToken != nil {
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
	if o == nil || o.GrantType == nil {
		var ret string
		return ret
	}
	return *o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetGrantTypeOk() (*string, bool) {
	if o == nil || o.GrantType == nil {
		return nil, false
	}
	return o.GrantType, true
}

// HasGrantType returns a boolean if a field has been set.
func (o *TokenResponse) HasGrantType() bool {
	if o != nil && o.GrantType != nil {
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
	if o == nil || o.ClientId == nil {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetClientIdOk() (*int64, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TokenResponse) HasClientId() bool {
	if o != nil && o.ClientId != nil {
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
	if o == nil || o.ClientIdAlias == nil {
		var ret string
		return ret
	}
	return *o.ClientIdAlias
}

// GetClientIdAliasOk returns a tuple with the ClientIdAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetClientIdAliasOk() (*string, bool) {
	if o == nil || o.ClientIdAlias == nil {
		return nil, false
	}
	return o.ClientIdAlias, true
}

// HasClientIdAlias returns a boolean if a field has been set.
func (o *TokenResponse) HasClientIdAlias() bool {
	if o != nil && o.ClientIdAlias != nil {
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
	if o == nil || o.ClientIdAliasUsed == nil {
		var ret bool
		return ret
	}
	return *o.ClientIdAliasUsed
}

// GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetClientIdAliasUsedOk() (*bool, bool) {
	if o == nil || o.ClientIdAliasUsed == nil {
		return nil, false
	}
	return o.ClientIdAliasUsed, true
}

// HasClientIdAliasUsed returns a boolean if a field has been set.
func (o *TokenResponse) HasClientIdAliasUsed() bool {
	if o != nil && o.ClientIdAliasUsed != nil {
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
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *TokenResponse) HasSubject() bool {
	if o != nil && o.Subject != nil {
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
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetScopesOk() ([]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *TokenResponse) HasScopes() bool {
	if o != nil && o.Scopes != nil {
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
	if o == nil || o.Properties == nil {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetPropertiesOk() ([]Property, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *TokenResponse) HasProperties() bool {
	if o != nil && o.Properties != nil {
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
	if o == nil || o.JwtAccessToken == nil {
		var ret string
		return ret
	}
	return *o.JwtAccessToken
}

// GetJwtAccessTokenOk returns a tuple with the JwtAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetJwtAccessTokenOk() (*string, bool) {
	if o == nil || o.JwtAccessToken == nil {
		return nil, false
	}
	return o.JwtAccessToken, true
}

// HasJwtAccessToken returns a boolean if a field has been set.
func (o *TokenResponse) HasJwtAccessToken() bool {
	if o != nil && o.JwtAccessToken != nil {
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
	if o == nil || o.Resources == nil {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetResourcesOk() ([]string, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *TokenResponse) HasResources() bool {
	if o != nil && o.Resources != nil {
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
	if o == nil || o.AccessTokenResources == nil {
		var ret []string
		return ret
	}
	return o.AccessTokenResources
}

// GetAccessTokenResourcesOk returns a tuple with the AccessTokenResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetAccessTokenResourcesOk() ([]string, bool) {
	if o == nil || o.AccessTokenResources == nil {
		return nil, false
	}
	return o.AccessTokenResources, true
}

// HasAccessTokenResources returns a boolean if a field has been set.
func (o *TokenResponse) HasAccessTokenResources() bool {
	if o != nil && o.AccessTokenResources != nil {
		return true
	}

	return false
}

// SetAccessTokenResources gets a reference to the given []string and assigns it to the AccessTokenResources field.
func (o *TokenResponse) SetAccessTokenResources(v []string) {
	o.AccessTokenResources = v
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *TokenResponse) GetAuthorizationDetails() AuthorizationDetails {
	if o == nil || o.AuthorizationDetails == nil {
		var ret AuthorizationDetails
		return ret
	}
	return *o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetAuthorizationDetailsOk() (*AuthorizationDetails, bool) {
	if o == nil || o.AuthorizationDetails == nil {
		return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *TokenResponse) HasAuthorizationDetails() bool {
	if o != nil && o.AuthorizationDetails != nil {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given AuthorizationDetails and assigns it to the AuthorizationDetails field.
func (o *TokenResponse) SetAuthorizationDetails(v AuthorizationDetails) {
	o.AuthorizationDetails = &v
}

// GetServiceAttributes returns the ServiceAttributes field value if set, zero value otherwise.
func (o *TokenResponse) GetServiceAttributes() []Pair {
	if o == nil || o.ServiceAttributes == nil {
		var ret []Pair
		return ret
	}
	return o.ServiceAttributes
}

// GetServiceAttributesOk returns a tuple with the ServiceAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetServiceAttributesOk() ([]Pair, bool) {
	if o == nil || o.ServiceAttributes == nil {
		return nil, false
	}
	return o.ServiceAttributes, true
}

// HasServiceAttributes returns a boolean if a field has been set.
func (o *TokenResponse) HasServiceAttributes() bool {
	if o != nil && o.ServiceAttributes != nil {
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
	if o == nil || o.ClientAttributes == nil {
		var ret []Pair
		return ret
	}
	return o.ClientAttributes
}

// GetClientAttributesOk returns a tuple with the ClientAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetClientAttributesOk() ([]Pair, bool) {
	if o == nil || o.ClientAttributes == nil {
		return nil, false
	}
	return o.ClientAttributes, true
}

// HasClientAttributes returns a boolean if a field has been set.
func (o *TokenResponse) HasClientAttributes() bool {
	if o != nil && o.ClientAttributes != nil {
		return true
	}

	return false
}

// SetClientAttributes gets a reference to the given []Pair and assigns it to the ClientAttributes field.
func (o *TokenResponse) SetClientAttributes(v []Pair) {
	o.ClientAttributes = v
}

func (o TokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResultCode != nil {
		toSerialize["resultCode"] = o.ResultCode
	}
	if o.ResultMessage != nil {
		toSerialize["resultMessage"] = o.ResultMessage
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.ResponseContent != nil {
		toSerialize["responseContent"] = o.ResponseContent
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Ticket != nil {
		toSerialize["ticket"] = o.Ticket
	}
	if o.AccessToken != nil {
		toSerialize["accessToken"] = o.AccessToken
	}
	if o.AccessTokenExpiresAt != nil {
		toSerialize["accessTokenExpiresAt"] = o.AccessTokenExpiresAt
	}
	if o.AccessTokenDuration != nil {
		toSerialize["accessTokenDuration"] = o.AccessTokenDuration
	}
	if o.RefreshToken != nil {
		toSerialize["refreshToken"] = o.RefreshToken
	}
	if o.RefreshTokenExpiresAt != nil {
		toSerialize["refreshTokenExpiresAt"] = o.RefreshTokenExpiresAt
	}
	if o.RefreshTokenDuration != nil {
		toSerialize["refreshTokenDuration"] = o.RefreshTokenDuration
	}
	if o.IdToken != nil {
		toSerialize["idToken"] = o.IdToken
	}
	if o.GrantType != nil {
		toSerialize["grantType"] = o.GrantType
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.ClientIdAlias != nil {
		toSerialize["clientIdAlias"] = o.ClientIdAlias
	}
	if o.ClientIdAliasUsed != nil {
		toSerialize["clientIdAliasUsed"] = o.ClientIdAliasUsed
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.JwtAccessToken != nil {
		toSerialize["jwtAccessToken"] = o.JwtAccessToken
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.AccessTokenResources != nil {
		toSerialize["accessTokenResources"] = o.AccessTokenResources
	}
	if o.AuthorizationDetails != nil {
		toSerialize["authorizationDetails"] = o.AuthorizationDetails
	}
	if o.ServiceAttributes != nil {
		toSerialize["serviceAttributes"] = o.ServiceAttributes
	}
	if o.ClientAttributes != nil {
		toSerialize["clientAttributes"] = o.ClientAttributes
	}
	return json.Marshal(toSerialize)
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


