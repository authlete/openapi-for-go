/*
Authlete API

Authlete API Document. 

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
)

// checks if the TokenCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenCreateRequest{}

// TokenCreateRequest struct for TokenCreateRequest
type TokenCreateRequest struct {
	GrantType GrantType `json:"grantType"`
	// The ID of the client application which will be associated with a newly created access token. 
	ClientId int64 `json:"clientId"`
	// The subject (= unique identifier) of the user who will be associated with a newly created access token. This parameter is required unless the grant type is `CLIENT_CREDENTIALS`. The value must consist of only ASCII characters and its length must not exceed 100. 
	Subject *string `json:"subject,omitempty"`
	// The scopes which will be associated with a newly created access token. Scopes that are not supported by the service cannot be specified and requesting them will cause an error. 
	Scopes []string `json:"scopes,omitempty"`
	// The duration of a newly created access token in seconds. If the value is 0, the duration is determined according to the settings of the service. 
	AccessTokenDuration *int64 `json:"accessTokenDuration,omitempty"`
	// The duration of a newly created refresh token in seconds. If the value is 0, the duration is determined according to the settings of the service.  A refresh token is not created (1) if the service does not support `REFRESH_TOKEN`, or (2) if the specified grant type is either `IMPLICIT`or `CLIENT_CREDENTIALS`. 
	RefreshTokenDuration *int64 `json:"refreshTokenDuration,omitempty"`
	// Extra properties to associate with a newly created access token. Note that properties parameter is accepted only when the HTTP method of the request is POST and Content-Type of the request is `application/json`, so don't use `GET` method or `application/x-www-form-urlencoded` if you want to specify properties. 
	Properties []Property `json:"properties,omitempty"`
	// A boolean request parameter which indicates whether to emulate that the client ID alias is used instead of the original numeric client ID when a new access token is created.  This has an effect only on the value of the aud claim in a response from [UserInfo endpoint](https://openid.net/specs/openid-connect-core-1_0.html#UserInfo). When you access the UserInfo endpoint (which is expected to be implemented using Authlete's `/api/auth/userinfo` API and `/api/auth/userinfo/issue` API) with an access token which has been created using Authlete's `/api/auth/token/create` API with this property (`clientIdAliasUsed`) `true`, the client ID alias is used as the value of the aud claim in a response from the UserInfo endpoint.  Note that if a client ID alias is not assigned to the client when Authlete's `/api/auth/token/create` API is called, this property (`clientIdAliasUsed`) has no effect (it is always regarded as `false`). 
	ClientIdAliasUsed *bool `json:"clientIdAliasUsed,omitempty"`
	// The value of the new access token.  The `/api/auth/token/create` API generates an access token. Therefore, callers of the API do not have to specify values of newly created access tokens. However, in some cases, for example, if you want to migrate existing access tokens from an old system to Authlete, you may want to specify values of access tokens. In such a case, you can specify the value of a newly created access token by passing a non-null value as the value of accessToken request parameter. The implementation of the `/api/auth/token/create` uses the value of the accessToken request parameter instead of generating a new value when the request parameter holds a non-null value.  Note that if the hash value of the specified access token already exists in Authlete's database, the access token cannot be inserted and the `/api/auth/token/create` API will report an error. 
	AccessToken *string `json:"accessToken,omitempty"`
	// The value of the new refresh token.  The `/api/auth/token/create` API may generate a refresh token. Therefore, callers of the API do not have to specify values of newly created refresh tokens. However, in some cases, for example, if you want to migrate existing refresh tokens from an old system to Authlete, you may want to specify values of refresh tokens. In such a case, you can specify the value of a newly created refresh token by passing a non-null value as the value of refreshToken request parameter. The implementation of the `/api/auth/token/create` uses the value of the refreshToken request parameter instead of generating a new value when the request parameter holds a non-null value.  Note that if the hash value of the specified refresh token already exists in Authlete's database, the refresh token cannot be inserted and the `/api/auth/token/create` API will report an error. 
	RefreshToken *string `json:"refreshToken,omitempty"`
	// Get whether the access token expires or not. By default, all access tokens expire after a period of time determined by their service.  If this request parameter is `true`, then the access token will not automatically expire and must be revoked or deleted manually at the service. If this request parameter is true, the `accessTokenDuration` request parameter is ignored. 
	AccessTokenPersistent *bool `json:"accessTokenPersistent,omitempty"`
	// The thumbprint of the MTLS certificate bound to this token. If this property is set, a certificate with the corresponding value MUST be presented with the access token when it is used by a client. The value of this property must be a SHA256 certificate thumbprint, base64url encoded. 
	CertificateThumbprint *string `json:"certificateThumbprint,omitempty"`
	// The thumbprint of the public key used for DPoP presentation of this token. If this property is set, a DPoP proof signed with the corresponding private key MUST be presented with the access token when it is used by a client. Additionally, the token's `token_type` will be set to 'DPoP'. 
	DpopKeyThumbprint *string `json:"dpopKeyThumbprint,omitempty"`
	AuthorizationDetails *AuthzDetails `json:"authorizationDetails,omitempty"`
	// The value of the resources to associate with the token. This property represents the value of one or more `resource` request parameters which is defined in \"RFC8707 Resource Indicators for OAuth 2.0\". 
	Resources []string `json:"resources,omitempty"`
	// the flag which indicates whether the access token is for an external attachment. 
	ForExternalAttachment *bool `json:"forExternalAttachment,omitempty"`
	// Additional claims that are added to the payload part of the JWT access token. 
	JwtAtClaims *string `json:"jwtAtClaims,omitempty"`
	// The Authentication Context Class Reference of the user authentication that the authorization server performed  during the course of issuing the access token. 
	Acr *string `json:"acr,omitempty"`
	// The time when the user authentication was performed during the course of issuing the access token. 
	AuthTime *int64 `json:"authTime,omitempty"`
	// Flag which indicates whether the entity ID of the client was used when the request for the access token was made.
	ClientEntityIdUsed *bool `json:"clientEntityIdUsed,omitempty"`
}

// NewTokenCreateRequest instantiates a new TokenCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenCreateRequest(grantType GrantType, clientId int64) *TokenCreateRequest {
	this := TokenCreateRequest{}
	this.GrantType = grantType
	this.ClientId = clientId
	return &this
}

// NewTokenCreateRequestWithDefaults instantiates a new TokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenCreateRequestWithDefaults() *TokenCreateRequest {
	this := TokenCreateRequest{}
	return &this
}

// GetGrantType returns the GrantType field value
func (o *TokenCreateRequest) GetGrantType() GrantType {
	if o == nil {
		var ret GrantType
		return ret
	}

	return o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetGrantTypeOk() (*GrantType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantType, true
}

// SetGrantType sets field value
func (o *TokenCreateRequest) SetGrantType(v GrantType) {
	o.GrantType = v
}

// GetClientId returns the ClientId field value
func (o *TokenCreateRequest) GetClientId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetClientIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *TokenCreateRequest) SetClientId(v int64) {
	o.ClientId = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *TokenCreateRequest) GetSubject() string {
	if o == nil || isNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetSubjectOk() (*string, bool) {
	if o == nil || isNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *TokenCreateRequest) HasSubject() bool {
	if o != nil && !isNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *TokenCreateRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *TokenCreateRequest) GetScopes() []string {
	if o == nil || isNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetScopesOk() ([]string, bool) {
	if o == nil || isNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *TokenCreateRequest) HasScopes() bool {
	if o != nil && !isNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *TokenCreateRequest) SetScopes(v []string) {
	o.Scopes = v
}

// GetAccessTokenDuration returns the AccessTokenDuration field value if set, zero value otherwise.
func (o *TokenCreateRequest) GetAccessTokenDuration() int64 {
	if o == nil || isNil(o.AccessTokenDuration) {
		var ret int64
		return ret
	}
	return *o.AccessTokenDuration
}

// GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetAccessTokenDurationOk() (*int64, bool) {
	if o == nil || isNil(o.AccessTokenDuration) {
		return nil, false
	}
	return o.AccessTokenDuration, true
}

// HasAccessTokenDuration returns a boolean if a field has been set.
func (o *TokenCreateRequest) HasAccessTokenDuration() bool {
	if o != nil && !isNil(o.AccessTokenDuration) {
		return true
	}

	return false
}

// SetAccessTokenDuration gets a reference to the given int64 and assigns it to the AccessTokenDuration field.
func (o *TokenCreateRequest) SetAccessTokenDuration(v int64) {
	o.AccessTokenDuration = &v
}

// GetRefreshTokenDuration returns the RefreshTokenDuration field value if set, zero value otherwise.
func (o *TokenCreateRequest) GetRefreshTokenDuration() int64 {
	if o == nil || isNil(o.RefreshTokenDuration) {
		var ret int64
		return ret
	}
	return *o.RefreshTokenDuration
}

// GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetRefreshTokenDurationOk() (*int64, bool) {
	if o == nil || isNil(o.RefreshTokenDuration) {
		return nil, false
	}
	return o.RefreshTokenDuration, true
}

// HasRefreshTokenDuration returns a boolean if a field has been set.
func (o *TokenCreateRequest) HasRefreshTokenDuration() bool {
	if o != nil && !isNil(o.RefreshTokenDuration) {
		return true
	}

	return false
}

// SetRefreshTokenDuration gets a reference to the given int64 and assigns it to the RefreshTokenDuration field.
func (o *TokenCreateRequest) SetRefreshTokenDuration(v int64) {
	o.RefreshTokenDuration = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *TokenCreateRequest) GetProperties() []Property {
	if o == nil || isNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetPropertiesOk() ([]Property, bool) {
	if o == nil || isNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *TokenCreateRequest) HasProperties() bool {
	if o != nil && !isNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *TokenCreateRequest) SetProperties(v []Property) {
	o.Properties = v
}

// GetClientIdAliasUsed returns the ClientIdAliasUsed field value if set, zero value otherwise.
func (o *TokenCreateRequest) GetClientIdAliasUsed() bool {
	if o == nil || isNil(o.ClientIdAliasUsed) {
		var ret bool
		return ret
	}
	return *o.ClientIdAliasUsed
}

// GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetClientIdAliasUsedOk() (*bool, bool) {
	if o == nil || isNil(o.ClientIdAliasUsed) {
		return nil, false
	}
	return o.ClientIdAliasUsed, true
}

// HasClientIdAliasUsed returns a boolean if a field has been set.
func (o *TokenCreateRequest) HasClientIdAliasUsed() bool {
	if o != nil && !isNil(o.ClientIdAliasUsed) {
		return true
	}

	return false
}

// SetClientIdAliasUsed gets a reference to the given bool and assigns it to the ClientIdAliasUsed field.
func (o *TokenCreateRequest) SetClientIdAliasUsed(v bool) {
	o.ClientIdAliasUsed = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *TokenCreateRequest) GetAccessToken() string {
	if o == nil || isNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil || isNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *TokenCreateRequest) HasAccessToken() bool {
	if o != nil && !isNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *TokenCreateRequest) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *TokenCreateRequest) GetRefreshToken() string {
	if o == nil || isNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetRefreshTokenOk() (*string, bool) {
	if o == nil || isNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *TokenCreateRequest) HasRefreshToken() bool {
	if o != nil && !isNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *TokenCreateRequest) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetAccessTokenPersistent returns the AccessTokenPersistent field value if set, zero value otherwise.
func (o *TokenCreateRequest) GetAccessTokenPersistent() bool {
	if o == nil || isNil(o.AccessTokenPersistent) {
		var ret bool
		return ret
	}
	return *o.AccessTokenPersistent
}

// GetAccessTokenPersistentOk returns a tuple with the AccessTokenPersistent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetAccessTokenPersistentOk() (*bool, bool) {
	if o == nil || isNil(o.AccessTokenPersistent) {
		return nil, false
	}
	return o.AccessTokenPersistent, true
}

// HasAccessTokenPersistent returns a boolean if a field has been set.
func (o *TokenCreateRequest) HasAccessTokenPersistent() bool {
	if o != nil && !isNil(o.AccessTokenPersistent) {
		return true
	}

	return false
}

// SetAccessTokenPersistent gets a reference to the given bool and assigns it to the AccessTokenPersistent field.
func (o *TokenCreateRequest) SetAccessTokenPersistent(v bool) {
	o.AccessTokenPersistent = &v
}

// GetCertificateThumbprint returns the CertificateThumbprint field value if set, zero value otherwise.
func (o *TokenCreateRequest) GetCertificateThumbprint() string {
	if o == nil || isNil(o.CertificateThumbprint) {
		var ret string
		return ret
	}
	return *o.CertificateThumbprint
}

// GetCertificateThumbprintOk returns a tuple with the CertificateThumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetCertificateThumbprintOk() (*string, bool) {
	if o == nil || isNil(o.CertificateThumbprint) {
		return nil, false
	}
	return o.CertificateThumbprint, true
}

// HasCertificateThumbprint returns a boolean if a field has been set.
func (o *TokenCreateRequest) HasCertificateThumbprint() bool {
	if o != nil && !isNil(o.CertificateThumbprint) {
		return true
	}

	return false
}

// SetCertificateThumbprint gets a reference to the given string and assigns it to the CertificateThumbprint field.
func (o *TokenCreateRequest) SetCertificateThumbprint(v string) {
	o.CertificateThumbprint = &v
}

// GetDpopKeyThumbprint returns the DpopKeyThumbprint field value if set, zero value otherwise.
func (o *TokenCreateRequest) GetDpopKeyThumbprint() string {
	if o == nil || isNil(o.DpopKeyThumbprint) {
		var ret string
		return ret
	}
	return *o.DpopKeyThumbprint
}

// GetDpopKeyThumbprintOk returns a tuple with the DpopKeyThumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetDpopKeyThumbprintOk() (*string, bool) {
	if o == nil || isNil(o.DpopKeyThumbprint) {
		return nil, false
	}
	return o.DpopKeyThumbprint, true
}

// HasDpopKeyThumbprint returns a boolean if a field has been set.
func (o *TokenCreateRequest) HasDpopKeyThumbprint() bool {
	if o != nil && !isNil(o.DpopKeyThumbprint) {
		return true
	}

	return false
}

// SetDpopKeyThumbprint gets a reference to the given string and assigns it to the DpopKeyThumbprint field.
func (o *TokenCreateRequest) SetDpopKeyThumbprint(v string) {
	o.DpopKeyThumbprint = &v
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *TokenCreateRequest) GetAuthorizationDetails() AuthzDetails {
	if o == nil || isNil(o.AuthorizationDetails) {
		var ret AuthzDetails
		return ret
	}
	return *o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetAuthorizationDetailsOk() (*AuthzDetails, bool) {
	if o == nil || isNil(o.AuthorizationDetails) {
		return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *TokenCreateRequest) HasAuthorizationDetails() bool {
	if o != nil && !isNil(o.AuthorizationDetails) {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given AuthzDetails and assigns it to the AuthorizationDetails field.
func (o *TokenCreateRequest) SetAuthorizationDetails(v AuthzDetails) {
	o.AuthorizationDetails = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *TokenCreateRequest) GetResources() []string {
	if o == nil || isNil(o.Resources) {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetResourcesOk() ([]string, bool) {
	if o == nil || isNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *TokenCreateRequest) HasResources() bool {
	if o != nil && !isNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *TokenCreateRequest) SetResources(v []string) {
	o.Resources = v
}

// GetForExternalAttachment returns the ForExternalAttachment field value if set, zero value otherwise.
func (o *TokenCreateRequest) GetForExternalAttachment() bool {
	if o == nil || isNil(o.ForExternalAttachment) {
		var ret bool
		return ret
	}
	return *o.ForExternalAttachment
}

// GetForExternalAttachmentOk returns a tuple with the ForExternalAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetForExternalAttachmentOk() (*bool, bool) {
	if o == nil || isNil(o.ForExternalAttachment) {
		return nil, false
	}
	return o.ForExternalAttachment, true
}

// HasForExternalAttachment returns a boolean if a field has been set.
func (o *TokenCreateRequest) HasForExternalAttachment() bool {
	if o != nil && !isNil(o.ForExternalAttachment) {
		return true
	}

	return false
}

// SetForExternalAttachment gets a reference to the given bool and assigns it to the ForExternalAttachment field.
func (o *TokenCreateRequest) SetForExternalAttachment(v bool) {
	o.ForExternalAttachment = &v
}

// GetJwtAtClaims returns the JwtAtClaims field value if set, zero value otherwise.
func (o *TokenCreateRequest) GetJwtAtClaims() string {
	if o == nil || isNil(o.JwtAtClaims) {
		var ret string
		return ret
	}
	return *o.JwtAtClaims
}

// GetJwtAtClaimsOk returns a tuple with the JwtAtClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetJwtAtClaimsOk() (*string, bool) {
	if o == nil || isNil(o.JwtAtClaims) {
		return nil, false
	}
	return o.JwtAtClaims, true
}

// HasJwtAtClaims returns a boolean if a field has been set.
func (o *TokenCreateRequest) HasJwtAtClaims() bool {
	if o != nil && !isNil(o.JwtAtClaims) {
		return true
	}

	return false
}

// SetJwtAtClaims gets a reference to the given string and assigns it to the JwtAtClaims field.
func (o *TokenCreateRequest) SetJwtAtClaims(v string) {
	o.JwtAtClaims = &v
}

// GetAcr returns the Acr field value if set, zero value otherwise.
func (o *TokenCreateRequest) GetAcr() string {
	if o == nil || isNil(o.Acr) {
		var ret string
		return ret
	}
	return *o.Acr
}

// GetAcrOk returns a tuple with the Acr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetAcrOk() (*string, bool) {
	if o == nil || isNil(o.Acr) {
		return nil, false
	}
	return o.Acr, true
}

// HasAcr returns a boolean if a field has been set.
func (o *TokenCreateRequest) HasAcr() bool {
	if o != nil && !isNil(o.Acr) {
		return true
	}

	return false
}

// SetAcr gets a reference to the given string and assigns it to the Acr field.
func (o *TokenCreateRequest) SetAcr(v string) {
	o.Acr = &v
}

// GetAuthTime returns the AuthTime field value if set, zero value otherwise.
func (o *TokenCreateRequest) GetAuthTime() int64 {
	if o == nil || isNil(o.AuthTime) {
		var ret int64
		return ret
	}
	return *o.AuthTime
}

// GetAuthTimeOk returns a tuple with the AuthTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetAuthTimeOk() (*int64, bool) {
	if o == nil || isNil(o.AuthTime) {
		return nil, false
	}
	return o.AuthTime, true
}

// HasAuthTime returns a boolean if a field has been set.
func (o *TokenCreateRequest) HasAuthTime() bool {
	if o != nil && !isNil(o.AuthTime) {
		return true
	}

	return false
}

// SetAuthTime gets a reference to the given int64 and assigns it to the AuthTime field.
func (o *TokenCreateRequest) SetAuthTime(v int64) {
	o.AuthTime = &v
}

// GetClientEntityIdUsed returns the ClientEntityIdUsed field value if set, zero value otherwise.
func (o *TokenCreateRequest) GetClientEntityIdUsed() bool {
	if o == nil || isNil(o.ClientEntityIdUsed) {
		var ret bool
		return ret
	}
	return *o.ClientEntityIdUsed
}

// GetClientEntityIdUsedOk returns a tuple with the ClientEntityIdUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreateRequest) GetClientEntityIdUsedOk() (*bool, bool) {
	if o == nil || isNil(o.ClientEntityIdUsed) {
		return nil, false
	}
	return o.ClientEntityIdUsed, true
}

// HasClientEntityIdUsed returns a boolean if a field has been set.
func (o *TokenCreateRequest) HasClientEntityIdUsed() bool {
	if o != nil && !isNil(o.ClientEntityIdUsed) {
		return true
	}

	return false
}

// SetClientEntityIdUsed gets a reference to the given bool and assigns it to the ClientEntityIdUsed field.
func (o *TokenCreateRequest) SetClientEntityIdUsed(v bool) {
	o.ClientEntityIdUsed = &v
}

func (o TokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["grantType"] = o.GrantType
	toSerialize["clientId"] = o.ClientId
	if !isNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !isNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !isNil(o.AccessTokenDuration) {
		toSerialize["accessTokenDuration"] = o.AccessTokenDuration
	}
	if !isNil(o.RefreshTokenDuration) {
		toSerialize["refreshTokenDuration"] = o.RefreshTokenDuration
	}
	if !isNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !isNil(o.ClientIdAliasUsed) {
		toSerialize["clientIdAliasUsed"] = o.ClientIdAliasUsed
	}
	if !isNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	if !isNil(o.RefreshToken) {
		toSerialize["refreshToken"] = o.RefreshToken
	}
	if !isNil(o.AccessTokenPersistent) {
		toSerialize["accessTokenPersistent"] = o.AccessTokenPersistent
	}
	if !isNil(o.CertificateThumbprint) {
		toSerialize["certificateThumbprint"] = o.CertificateThumbprint
	}
	if !isNil(o.DpopKeyThumbprint) {
		toSerialize["dpopKeyThumbprint"] = o.DpopKeyThumbprint
	}
	if !isNil(o.AuthorizationDetails) {
		toSerialize["authorizationDetails"] = o.AuthorizationDetails
	}
	if !isNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !isNil(o.ForExternalAttachment) {
		toSerialize["forExternalAttachment"] = o.ForExternalAttachment
	}
	if !isNil(o.JwtAtClaims) {
		toSerialize["jwtAtClaims"] = o.JwtAtClaims
	}
	if !isNil(o.Acr) {
		toSerialize["acr"] = o.Acr
	}
	if !isNil(o.AuthTime) {
		toSerialize["authTime"] = o.AuthTime
	}
	if !isNil(o.ClientEntityIdUsed) {
		toSerialize["clientEntityIdUsed"] = o.ClientEntityIdUsed
	}
	return toSerialize, nil
}

type NullableTokenCreateRequest struct {
	value *TokenCreateRequest
	isSet bool
}

func (v NullableTokenCreateRequest) Get() *TokenCreateRequest {
	return v.value
}

func (v *NullableTokenCreateRequest) Set(val *TokenCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenCreateRequest(val *TokenCreateRequest) *NullableTokenCreateRequest {
	return &NullableTokenCreateRequest{value: val, isSet: true}
}

func (v NullableTokenCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


