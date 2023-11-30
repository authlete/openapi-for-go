/*
Authlete API

Authlete API Document. 

API version: 2.3.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
)

// checks if the AuthorizationIssueRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationIssueRequest{}

// AuthorizationIssueRequest struct for AuthorizationIssueRequest
type AuthorizationIssueRequest struct {
	// The ticket issued from Authlete `/auth/authorization` API. 
	Ticket string `json:"ticket"`
	// The subject (= a user account managed by the service) who has granted authorization to the client application. 
	Subject string `json:"subject"`
	// The time when the authentication of the end-user occurred. Its value is the number of seconds from `1970-01-01`. 
	AuthTime *string `json:"authTime,omitempty"`
	// The Authentication Context Class Reference performed for the end-user authentication.
	Acr *string `json:"acr,omitempty"`
	// The claims of the end-user (= pieces of information about the end-user) in JSON format. See [OpenID Connect Core 1.0, 5.1. Standard Claims](https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims) for details about the format. 
	Claims *string `json:"claims,omitempty"`
	// Extra properties to associate with an access token and/or an authorization code.
	Properties []Property `json:"properties,omitempty"`
	// Scopes to associate with an access token and/or an authorization code. If a non-empty string array is given, it replaces the scopes specified by the original authorization request. 
	Scopes []string `json:"scopes,omitempty"`
	// The value of the `sub` claim to embed in an ID token. If this request parameter is `null` or empty, the value of the `subject` request parameter is used as the value of the `sub` claim. 
	Sub *string `json:"sub,omitempty"`
	// JSON that represents additional JWS header parameters for ID tokens that may be issued based on the authorization request. 
	IdtHeaderParams *string `json:"idtHeaderParams,omitempty"`
	// Claim key-value pairs that are used to compute transformed claims. 
	ClaimsForTx *string `json:"claimsForTx,omitempty"`
	// the claims that the user has consented for the client application to know. 
	ConsentedClaims []string `json:"consentedClaims,omitempty"`
	AuthorizationDetails *AuthorizationDetails `json:"authorizationDetails,omitempty"`
	// Additional claims that are added to the payload part of the JWT access token. 
	JwtAtClaims *string `json:"jwtAtClaims,omitempty"`
	// The representation of an access token that may be issued as a result of the Authlete API call. 
	AccessToken *string `json:"accessToken,omitempty"`
}

// NewAuthorizationIssueRequest instantiates a new AuthorizationIssueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationIssueRequest(ticket string, subject string) *AuthorizationIssueRequest {
	this := AuthorizationIssueRequest{}
	this.Ticket = ticket
	this.Subject = subject
	return &this
}

// NewAuthorizationIssueRequestWithDefaults instantiates a new AuthorizationIssueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationIssueRequestWithDefaults() *AuthorizationIssueRequest {
	this := AuthorizationIssueRequest{}
	return &this
}

// GetTicket returns the Ticket field value
func (o *AuthorizationIssueRequest) GetTicket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ticket
}

// GetTicketOk returns a tuple with the Ticket field value
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetTicketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ticket, true
}

// SetTicket sets field value
func (o *AuthorizationIssueRequest) SetTicket(v string) {
	o.Ticket = v
}

// GetSubject returns the Subject field value
func (o *AuthorizationIssueRequest) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *AuthorizationIssueRequest) SetSubject(v string) {
	o.Subject = v
}

// GetAuthTime returns the AuthTime field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetAuthTime() string {
	if o == nil || isNil(o.AuthTime) {
		var ret string
		return ret
	}
	return *o.AuthTime
}

// GetAuthTimeOk returns a tuple with the AuthTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetAuthTimeOk() (*string, bool) {
	if o == nil || isNil(o.AuthTime) {
		return nil, false
	}
	return o.AuthTime, true
}

// HasAuthTime returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasAuthTime() bool {
	if o != nil && !isNil(o.AuthTime) {
		return true
	}

	return false
}

// SetAuthTime gets a reference to the given string and assigns it to the AuthTime field.
func (o *AuthorizationIssueRequest) SetAuthTime(v string) {
	o.AuthTime = &v
}

// GetAcr returns the Acr field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetAcr() string {
	if o == nil || isNil(o.Acr) {
		var ret string
		return ret
	}
	return *o.Acr
}

// GetAcrOk returns a tuple with the Acr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetAcrOk() (*string, bool) {
	if o == nil || isNil(o.Acr) {
		return nil, false
	}
	return o.Acr, true
}

// HasAcr returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasAcr() bool {
	if o != nil && !isNil(o.Acr) {
		return true
	}

	return false
}

// SetAcr gets a reference to the given string and assigns it to the Acr field.
func (o *AuthorizationIssueRequest) SetAcr(v string) {
	o.Acr = &v
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetClaims() string {
	if o == nil || isNil(o.Claims) {
		var ret string
		return ret
	}
	return *o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetClaimsOk() (*string, bool) {
	if o == nil || isNil(o.Claims) {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasClaims() bool {
	if o != nil && !isNil(o.Claims) {
		return true
	}

	return false
}

// SetClaims gets a reference to the given string and assigns it to the Claims field.
func (o *AuthorizationIssueRequest) SetClaims(v string) {
	o.Claims = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetProperties() []Property {
	if o == nil || isNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetPropertiesOk() ([]Property, bool) {
	if o == nil || isNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasProperties() bool {
	if o != nil && !isNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *AuthorizationIssueRequest) SetProperties(v []Property) {
	o.Properties = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetScopes() []string {
	if o == nil || isNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetScopesOk() ([]string, bool) {
	if o == nil || isNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasScopes() bool {
	if o != nil && !isNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *AuthorizationIssueRequest) SetScopes(v []string) {
	o.Scopes = v
}

// GetSub returns the Sub field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetSub() string {
	if o == nil || isNil(o.Sub) {
		var ret string
		return ret
	}
	return *o.Sub
}

// GetSubOk returns a tuple with the Sub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetSubOk() (*string, bool) {
	if o == nil || isNil(o.Sub) {
		return nil, false
	}
	return o.Sub, true
}

// HasSub returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasSub() bool {
	if o != nil && !isNil(o.Sub) {
		return true
	}

	return false
}

// SetSub gets a reference to the given string and assigns it to the Sub field.
func (o *AuthorizationIssueRequest) SetSub(v string) {
	o.Sub = &v
}

// GetIdtHeaderParams returns the IdtHeaderParams field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetIdtHeaderParams() string {
	if o == nil || isNil(o.IdtHeaderParams) {
		var ret string
		return ret
	}
	return *o.IdtHeaderParams
}

// GetIdtHeaderParamsOk returns a tuple with the IdtHeaderParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetIdtHeaderParamsOk() (*string, bool) {
	if o == nil || isNil(o.IdtHeaderParams) {
		return nil, false
	}
	return o.IdtHeaderParams, true
}

// HasIdtHeaderParams returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasIdtHeaderParams() bool {
	if o != nil && !isNil(o.IdtHeaderParams) {
		return true
	}

	return false
}

// SetIdtHeaderParams gets a reference to the given string and assigns it to the IdtHeaderParams field.
func (o *AuthorizationIssueRequest) SetIdtHeaderParams(v string) {
	o.IdtHeaderParams = &v
}

// GetClaimsForTx returns the ClaimsForTx field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetClaimsForTx() string {
	if o == nil || isNil(o.ClaimsForTx) {
		var ret string
		return ret
	}
	return *o.ClaimsForTx
}

// GetClaimsForTxOk returns a tuple with the ClaimsForTx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetClaimsForTxOk() (*string, bool) {
	if o == nil || isNil(o.ClaimsForTx) {
		return nil, false
	}
	return o.ClaimsForTx, true
}

// HasClaimsForTx returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasClaimsForTx() bool {
	if o != nil && !isNil(o.ClaimsForTx) {
		return true
	}

	return false
}

// SetClaimsForTx gets a reference to the given string and assigns it to the ClaimsForTx field.
func (o *AuthorizationIssueRequest) SetClaimsForTx(v string) {
	o.ClaimsForTx = &v
}

// GetConsentedClaims returns the ConsentedClaims field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetConsentedClaims() []string {
	if o == nil || isNil(o.ConsentedClaims) {
		var ret []string
		return ret
	}
	return o.ConsentedClaims
}

// GetConsentedClaimsOk returns a tuple with the ConsentedClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetConsentedClaimsOk() ([]string, bool) {
	if o == nil || isNil(o.ConsentedClaims) {
		return nil, false
	}
	return o.ConsentedClaims, true
}

// HasConsentedClaims returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasConsentedClaims() bool {
	if o != nil && !isNil(o.ConsentedClaims) {
		return true
	}

	return false
}

// SetConsentedClaims gets a reference to the given []string and assigns it to the ConsentedClaims field.
func (o *AuthorizationIssueRequest) SetConsentedClaims(v []string) {
	o.ConsentedClaims = v
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetAuthorizationDetails() AuthorizationDetails {
	if o == nil || isNil(o.AuthorizationDetails) {
		var ret AuthorizationDetails
		return ret
	}
	return *o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetAuthorizationDetailsOk() (*AuthorizationDetails, bool) {
	if o == nil || isNil(o.AuthorizationDetails) {
		return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasAuthorizationDetails() bool {
	if o != nil && !isNil(o.AuthorizationDetails) {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given AuthorizationDetails and assigns it to the AuthorizationDetails field.
func (o *AuthorizationIssueRequest) SetAuthorizationDetails(v AuthorizationDetails) {
	o.AuthorizationDetails = &v
}

// GetJwtAtClaims returns the JwtAtClaims field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetJwtAtClaims() string {
	if o == nil || isNil(o.JwtAtClaims) {
		var ret string
		return ret
	}
	return *o.JwtAtClaims
}

// GetJwtAtClaimsOk returns a tuple with the JwtAtClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetJwtAtClaimsOk() (*string, bool) {
	if o == nil || isNil(o.JwtAtClaims) {
		return nil, false
	}
	return o.JwtAtClaims, true
}

// HasJwtAtClaims returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasJwtAtClaims() bool {
	if o != nil && !isNil(o.JwtAtClaims) {
		return true
	}

	return false
}

// SetJwtAtClaims gets a reference to the given string and assigns it to the JwtAtClaims field.
func (o *AuthorizationIssueRequest) SetJwtAtClaims(v string) {
	o.JwtAtClaims = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetAccessToken() string {
	if o == nil || isNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil || isNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasAccessToken() bool {
	if o != nil && !isNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *AuthorizationIssueRequest) SetAccessToken(v string) {
	o.AccessToken = &v
}

func (o AuthorizationIssueRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationIssueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ticket"] = o.Ticket
	toSerialize["subject"] = o.Subject
	if !isNil(o.AuthTime) {
		toSerialize["authTime"] = o.AuthTime
	}
	if !isNil(o.Acr) {
		toSerialize["acr"] = o.Acr
	}
	if !isNil(o.Claims) {
		toSerialize["claims"] = o.Claims
	}
	if !isNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !isNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !isNil(o.Sub) {
		toSerialize["sub"] = o.Sub
	}
	if !isNil(o.IdtHeaderParams) {
		toSerialize["idtHeaderParams"] = o.IdtHeaderParams
	}
	if !isNil(o.ClaimsForTx) {
		toSerialize["claimsForTx"] = o.ClaimsForTx
	}
	if !isNil(o.ConsentedClaims) {
		toSerialize["consentedClaims"] = o.ConsentedClaims
	}
	if !isNil(o.AuthorizationDetails) {
		toSerialize["authorizationDetails"] = o.AuthorizationDetails
	}
	if !isNil(o.JwtAtClaims) {
		toSerialize["jwtAtClaims"] = o.JwtAtClaims
	}
	if !isNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	return toSerialize, nil
}

type NullableAuthorizationIssueRequest struct {
	value *AuthorizationIssueRequest
	isSet bool
}

func (v NullableAuthorizationIssueRequest) Get() *AuthorizationIssueRequest {
	return v.value
}

func (v *NullableAuthorizationIssueRequest) Set(val *AuthorizationIssueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationIssueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationIssueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationIssueRequest(val *AuthorizationIssueRequest) *NullableAuthorizationIssueRequest {
	return &NullableAuthorizationIssueRequest{value: val, isSet: true}
}

func (v NullableAuthorizationIssueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationIssueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


