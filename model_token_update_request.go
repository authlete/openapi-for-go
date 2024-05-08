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

// checks if the TokenUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenUpdateRequest{}

// TokenUpdateRequest struct for TokenUpdateRequest
type TokenUpdateRequest struct {
	// An access token. 
	AccessToken string `json:"accessToken"`
	// A new date at which the access token will expire in milliseconds since the Unix epoch (1970-01-01). If the `accessTokenExpiresAt` request parameter is not included in a request or its value is 0 (or negative), the expiration date of the access token is not changed. 
	AccessTokenExpiresAt *int64 `json:"accessTokenExpiresAt,omitempty"`
	// A new set of scopes assigned to the access token. Scopes that are not supported by the service and those that the client application associated with the access token is not allowed to request are ignored on the server side. If the `scopes` request parameter is not included in a request or its value is `null`, the scopes of the access token are not changed. Note that `properties` parameter is accepted only when `Content-Type` of the request is `application/json`, so don't use `application/x-www-form-urlencoded` if you want to specify `properties`. 
	Scopes []string `json:"scopes,omitempty"`
	// A new set of properties assigned to the access token. If the `properties` request parameter is not included in a request or its value is null, the properties of the access token are not changed. 
	Properties []Property `json:"properties,omitempty"`
	// A boolean request parameter which indicates whether the API attempts to update the expiration date of the access token when the scopes linked to the access token are changed by this request. 
	AccessTokenExpiresAtUpdatedOnScopeUpdate *bool `json:"accessTokenExpiresAtUpdatedOnScopeUpdate,omitempty"`
	// The hash of the access token value. Used when the hash of the token is known (perhaps from lookup) but the value of the token itself is not. The value of the `accessToken` parameter takes precedence. 
	AccessTokenHash *string `json:"accessTokenHash,omitempty"`
	// A boolean request parameter which indicates whether to update the value of the access token in the data store. If this parameter is set to `true` then a new access token value is generated by the server and returned in the response. 
	AccessTokenValueUpdated *bool `json:"accessTokenValueUpdated,omitempty"`
	// The flag which indicates whether the access token expires or not. By default, all access tokens expire after a period of time determined by their service. If this request parameter is `true` then the access token will not automatically expire and must be revoked or deleted manually at the service.  If this request parameter is `true`, the `accessTokenExpiresAt` request parameter is ignored. If this request parameter is `false`, the `accessTokenExpiresAt` request parameter is processed normally. 
	AccessTokenPersistent *bool `json:"accessTokenPersistent,omitempty"`
	// The thumbprint of the MTLS certificate bound to this token. If this property is set, a certificate with the corresponding value MUST be presented with the access token when it is used by a client. The value of this property must be a SHA256 certificate thumbprint, base64url encoded. 
	CertificateThumbprint *string `json:"certificateThumbprint,omitempty"`
	// The thumbprint of the public key used for DPoP presentation of this token. If this property is set, a DPoP proof signed with the corresponding private key MUST be presented with the access token when it is used by a client. Additionally, the token's `token_type` will be set to 'DPoP'. 
	DpopKeyThumbprint *string `json:"dpopKeyThumbprint,omitempty"`
	AuthorizationDetails *AuthzDetails `json:"authorizationDetails,omitempty"`
	// the flag which indicates whether the access token is for an external attachment. 
	ForExternalAttachment *bool `json:"forExternalAttachment,omitempty"`
}

// NewTokenUpdateRequest instantiates a new TokenUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenUpdateRequest(accessToken string) *TokenUpdateRequest {
	this := TokenUpdateRequest{}
	this.AccessToken = accessToken
	return &this
}

// NewTokenUpdateRequestWithDefaults instantiates a new TokenUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenUpdateRequestWithDefaults() *TokenUpdateRequest {
	this := TokenUpdateRequest{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *TokenUpdateRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TokenUpdateRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetAccessTokenExpiresAt returns the AccessTokenExpiresAt field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetAccessTokenExpiresAt() int64 {
	if o == nil || isNil(o.AccessTokenExpiresAt) {
		var ret int64
		return ret
	}
	return *o.AccessTokenExpiresAt
}

// GetAccessTokenExpiresAtOk returns a tuple with the AccessTokenExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetAccessTokenExpiresAtOk() (*int64, bool) {
	if o == nil || isNil(o.AccessTokenExpiresAt) {
		return nil, false
	}
	return o.AccessTokenExpiresAt, true
}

// HasAccessTokenExpiresAt returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasAccessTokenExpiresAt() bool {
	if o != nil && !isNil(o.AccessTokenExpiresAt) {
		return true
	}

	return false
}

// SetAccessTokenExpiresAt gets a reference to the given int64 and assigns it to the AccessTokenExpiresAt field.
func (o *TokenUpdateRequest) SetAccessTokenExpiresAt(v int64) {
	o.AccessTokenExpiresAt = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetScopes() []string {
	if o == nil || isNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetScopesOk() ([]string, bool) {
	if o == nil || isNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasScopes() bool {
	if o != nil && !isNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *TokenUpdateRequest) SetScopes(v []string) {
	o.Scopes = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetProperties() []Property {
	if o == nil || isNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetPropertiesOk() ([]Property, bool) {
	if o == nil || isNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasProperties() bool {
	if o != nil && !isNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *TokenUpdateRequest) SetProperties(v []Property) {
	o.Properties = v
}

// GetAccessTokenExpiresAtUpdatedOnScopeUpdate returns the AccessTokenExpiresAtUpdatedOnScopeUpdate field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetAccessTokenExpiresAtUpdatedOnScopeUpdate() bool {
	if o == nil || isNil(o.AccessTokenExpiresAtUpdatedOnScopeUpdate) {
		var ret bool
		return ret
	}
	return *o.AccessTokenExpiresAtUpdatedOnScopeUpdate
}

// GetAccessTokenExpiresAtUpdatedOnScopeUpdateOk returns a tuple with the AccessTokenExpiresAtUpdatedOnScopeUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetAccessTokenExpiresAtUpdatedOnScopeUpdateOk() (*bool, bool) {
	if o == nil || isNil(o.AccessTokenExpiresAtUpdatedOnScopeUpdate) {
		return nil, false
	}
	return o.AccessTokenExpiresAtUpdatedOnScopeUpdate, true
}

// HasAccessTokenExpiresAtUpdatedOnScopeUpdate returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasAccessTokenExpiresAtUpdatedOnScopeUpdate() bool {
	if o != nil && !isNil(o.AccessTokenExpiresAtUpdatedOnScopeUpdate) {
		return true
	}

	return false
}

// SetAccessTokenExpiresAtUpdatedOnScopeUpdate gets a reference to the given bool and assigns it to the AccessTokenExpiresAtUpdatedOnScopeUpdate field.
func (o *TokenUpdateRequest) SetAccessTokenExpiresAtUpdatedOnScopeUpdate(v bool) {
	o.AccessTokenExpiresAtUpdatedOnScopeUpdate = &v
}

// GetAccessTokenHash returns the AccessTokenHash field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetAccessTokenHash() string {
	if o == nil || isNil(o.AccessTokenHash) {
		var ret string
		return ret
	}
	return *o.AccessTokenHash
}

// GetAccessTokenHashOk returns a tuple with the AccessTokenHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetAccessTokenHashOk() (*string, bool) {
	if o == nil || isNil(o.AccessTokenHash) {
		return nil, false
	}
	return o.AccessTokenHash, true
}

// HasAccessTokenHash returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasAccessTokenHash() bool {
	if o != nil && !isNil(o.AccessTokenHash) {
		return true
	}

	return false
}

// SetAccessTokenHash gets a reference to the given string and assigns it to the AccessTokenHash field.
func (o *TokenUpdateRequest) SetAccessTokenHash(v string) {
	o.AccessTokenHash = &v
}

// GetAccessTokenValueUpdated returns the AccessTokenValueUpdated field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetAccessTokenValueUpdated() bool {
	if o == nil || isNil(o.AccessTokenValueUpdated) {
		var ret bool
		return ret
	}
	return *o.AccessTokenValueUpdated
}

// GetAccessTokenValueUpdatedOk returns a tuple with the AccessTokenValueUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetAccessTokenValueUpdatedOk() (*bool, bool) {
	if o == nil || isNil(o.AccessTokenValueUpdated) {
		return nil, false
	}
	return o.AccessTokenValueUpdated, true
}

// HasAccessTokenValueUpdated returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasAccessTokenValueUpdated() bool {
	if o != nil && !isNil(o.AccessTokenValueUpdated) {
		return true
	}

	return false
}

// SetAccessTokenValueUpdated gets a reference to the given bool and assigns it to the AccessTokenValueUpdated field.
func (o *TokenUpdateRequest) SetAccessTokenValueUpdated(v bool) {
	o.AccessTokenValueUpdated = &v
}

// GetAccessTokenPersistent returns the AccessTokenPersistent field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetAccessTokenPersistent() bool {
	if o == nil || isNil(o.AccessTokenPersistent) {
		var ret bool
		return ret
	}
	return *o.AccessTokenPersistent
}

// GetAccessTokenPersistentOk returns a tuple with the AccessTokenPersistent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetAccessTokenPersistentOk() (*bool, bool) {
	if o == nil || isNil(o.AccessTokenPersistent) {
		return nil, false
	}
	return o.AccessTokenPersistent, true
}

// HasAccessTokenPersistent returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasAccessTokenPersistent() bool {
	if o != nil && !isNil(o.AccessTokenPersistent) {
		return true
	}

	return false
}

// SetAccessTokenPersistent gets a reference to the given bool and assigns it to the AccessTokenPersistent field.
func (o *TokenUpdateRequest) SetAccessTokenPersistent(v bool) {
	o.AccessTokenPersistent = &v
}

// GetCertificateThumbprint returns the CertificateThumbprint field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetCertificateThumbprint() string {
	if o == nil || isNil(o.CertificateThumbprint) {
		var ret string
		return ret
	}
	return *o.CertificateThumbprint
}

// GetCertificateThumbprintOk returns a tuple with the CertificateThumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetCertificateThumbprintOk() (*string, bool) {
	if o == nil || isNil(o.CertificateThumbprint) {
		return nil, false
	}
	return o.CertificateThumbprint, true
}

// HasCertificateThumbprint returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasCertificateThumbprint() bool {
	if o != nil && !isNil(o.CertificateThumbprint) {
		return true
	}

	return false
}

// SetCertificateThumbprint gets a reference to the given string and assigns it to the CertificateThumbprint field.
func (o *TokenUpdateRequest) SetCertificateThumbprint(v string) {
	o.CertificateThumbprint = &v
}

// GetDpopKeyThumbprint returns the DpopKeyThumbprint field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetDpopKeyThumbprint() string {
	if o == nil || isNil(o.DpopKeyThumbprint) {
		var ret string
		return ret
	}
	return *o.DpopKeyThumbprint
}

// GetDpopKeyThumbprintOk returns a tuple with the DpopKeyThumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetDpopKeyThumbprintOk() (*string, bool) {
	if o == nil || isNil(o.DpopKeyThumbprint) {
		return nil, false
	}
	return o.DpopKeyThumbprint, true
}

// HasDpopKeyThumbprint returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasDpopKeyThumbprint() bool {
	if o != nil && !isNil(o.DpopKeyThumbprint) {
		return true
	}

	return false
}

// SetDpopKeyThumbprint gets a reference to the given string and assigns it to the DpopKeyThumbprint field.
func (o *TokenUpdateRequest) SetDpopKeyThumbprint(v string) {
	o.DpopKeyThumbprint = &v
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetAuthorizationDetails() AuthzDetails {
	if o == nil || isNil(o.AuthorizationDetails) {
		var ret AuthzDetails
		return ret
	}
	return *o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetAuthorizationDetailsOk() (*AuthzDetails, bool) {
	if o == nil || isNil(o.AuthorizationDetails) {
		return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasAuthorizationDetails() bool {
	if o != nil && !isNil(o.AuthorizationDetails) {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given AuthzDetails and assigns it to the AuthorizationDetails field.
func (o *TokenUpdateRequest) SetAuthorizationDetails(v AuthzDetails) {
	o.AuthorizationDetails = &v
}

// GetForExternalAttachment returns the ForExternalAttachment field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetForExternalAttachment() bool {
	if o == nil || isNil(o.ForExternalAttachment) {
		var ret bool
		return ret
	}
	return *o.ForExternalAttachment
}

// GetForExternalAttachmentOk returns a tuple with the ForExternalAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetForExternalAttachmentOk() (*bool, bool) {
	if o == nil || isNil(o.ForExternalAttachment) {
		return nil, false
	}
	return o.ForExternalAttachment, true
}

// HasForExternalAttachment returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasForExternalAttachment() bool {
	if o != nil && !isNil(o.ForExternalAttachment) {
		return true
	}

	return false
}

// SetForExternalAttachment gets a reference to the given bool and assigns it to the ForExternalAttachment field.
func (o *TokenUpdateRequest) SetForExternalAttachment(v bool) {
	o.ForExternalAttachment = &v
}

func (o TokenUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accessToken"] = o.AccessToken
	if !isNil(o.AccessTokenExpiresAt) {
		toSerialize["accessTokenExpiresAt"] = o.AccessTokenExpiresAt
	}
	if !isNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !isNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !isNil(o.AccessTokenExpiresAtUpdatedOnScopeUpdate) {
		toSerialize["accessTokenExpiresAtUpdatedOnScopeUpdate"] = o.AccessTokenExpiresAtUpdatedOnScopeUpdate
	}
	if !isNil(o.AccessTokenHash) {
		toSerialize["accessTokenHash"] = o.AccessTokenHash
	}
	if !isNil(o.AccessTokenValueUpdated) {
		toSerialize["accessTokenValueUpdated"] = o.AccessTokenValueUpdated
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
	if !isNil(o.ForExternalAttachment) {
		toSerialize["forExternalAttachment"] = o.ForExternalAttachment
	}
	return toSerialize, nil
}

type NullableTokenUpdateRequest struct {
	value *TokenUpdateRequest
	isSet bool
}

func (v NullableTokenUpdateRequest) Get() *TokenUpdateRequest {
	return v.value
}

func (v *NullableTokenUpdateRequest) Set(val *TokenUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenUpdateRequest(val *TokenUpdateRequest) *NullableTokenUpdateRequest {
	return &NullableTokenUpdateRequest{value: val, isSet: true}
}

func (v NullableTokenUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


