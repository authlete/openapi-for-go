/*
Authlete API

Authlete API Document. 

API version: 2.2.25
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
)

// AccessToken struct for AccessToken
type AccessToken struct {
	// The hash of the access token.
	AccessTokenHash *string `json:"accessTokenHash,omitempty"`
	// The timestamp at which the access token will expire.
	AccessTokenExpiresAt *int64 `json:"accessTokenExpiresAt,omitempty"`
	// The hash of the refresh token.
	RefreshTokenHash *string `json:"refreshTokenHash,omitempty"`
	// The timestamp at which the refresh token will expire.
	RefreshTokenExpiresAt *int64 `json:"refreshTokenExpiresAt,omitempty"`
	// The timestamp at which the access token was first created. 
	CreatedAt *int64 `json:"createdAt,omitempty"`
	// The timestamp at which the access token was last refreshed using the refresh token. 
	LastRefreshedAt *int64 `json:"lastRefreshedAt,omitempty"`
	// The ID of the client associated with the access token. 
	ClientId *int64 `json:"clientId,omitempty"`
	// The subject (= unique user ID) associated with the access token. 
	Subject *string `json:"subject,omitempty"`
	GrantType *GrantType `json:"grantType,omitempty"`
	// The scopes associated with the access token. 
	Scopes []string `json:"scopes,omitempty"`
	// The properties associated with the access token. 
	Properties []Property `json:"properties,omitempty"`
}

// NewAccessToken instantiates a new AccessToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessToken() *AccessToken {
	this := AccessToken{}
	return &this
}

// NewAccessTokenWithDefaults instantiates a new AccessToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenWithDefaults() *AccessToken {
	this := AccessToken{}
	return &this
}

// GetAccessTokenHash returns the AccessTokenHash field value if set, zero value otherwise.
func (o *AccessToken) GetAccessTokenHash() string {
	if o == nil || o.AccessTokenHash == nil {
		var ret string
		return ret
	}
	return *o.AccessTokenHash
}

// GetAccessTokenHashOk returns a tuple with the AccessTokenHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetAccessTokenHashOk() (*string, bool) {
	if o == nil || o.AccessTokenHash == nil {
		return nil, false
	}
	return o.AccessTokenHash, true
}

// HasAccessTokenHash returns a boolean if a field has been set.
func (o *AccessToken) HasAccessTokenHash() bool {
	if o != nil && o.AccessTokenHash != nil {
		return true
	}

	return false
}

// SetAccessTokenHash gets a reference to the given string and assigns it to the AccessTokenHash field.
func (o *AccessToken) SetAccessTokenHash(v string) {
	o.AccessTokenHash = &v
}

// GetAccessTokenExpiresAt returns the AccessTokenExpiresAt field value if set, zero value otherwise.
func (o *AccessToken) GetAccessTokenExpiresAt() int64 {
	if o == nil || o.AccessTokenExpiresAt == nil {
		var ret int64
		return ret
	}
	return *o.AccessTokenExpiresAt
}

// GetAccessTokenExpiresAtOk returns a tuple with the AccessTokenExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetAccessTokenExpiresAtOk() (*int64, bool) {
	if o == nil || o.AccessTokenExpiresAt == nil {
		return nil, false
	}
	return o.AccessTokenExpiresAt, true
}

// HasAccessTokenExpiresAt returns a boolean if a field has been set.
func (o *AccessToken) HasAccessTokenExpiresAt() bool {
	if o != nil && o.AccessTokenExpiresAt != nil {
		return true
	}

	return false
}

// SetAccessTokenExpiresAt gets a reference to the given int64 and assigns it to the AccessTokenExpiresAt field.
func (o *AccessToken) SetAccessTokenExpiresAt(v int64) {
	o.AccessTokenExpiresAt = &v
}

// GetRefreshTokenHash returns the RefreshTokenHash field value if set, zero value otherwise.
func (o *AccessToken) GetRefreshTokenHash() string {
	if o == nil || o.RefreshTokenHash == nil {
		var ret string
		return ret
	}
	return *o.RefreshTokenHash
}

// GetRefreshTokenHashOk returns a tuple with the RefreshTokenHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetRefreshTokenHashOk() (*string, bool) {
	if o == nil || o.RefreshTokenHash == nil {
		return nil, false
	}
	return o.RefreshTokenHash, true
}

// HasRefreshTokenHash returns a boolean if a field has been set.
func (o *AccessToken) HasRefreshTokenHash() bool {
	if o != nil && o.RefreshTokenHash != nil {
		return true
	}

	return false
}

// SetRefreshTokenHash gets a reference to the given string and assigns it to the RefreshTokenHash field.
func (o *AccessToken) SetRefreshTokenHash(v string) {
	o.RefreshTokenHash = &v
}

// GetRefreshTokenExpiresAt returns the RefreshTokenExpiresAt field value if set, zero value otherwise.
func (o *AccessToken) GetRefreshTokenExpiresAt() int64 {
	if o == nil || o.RefreshTokenExpiresAt == nil {
		var ret int64
		return ret
	}
	return *o.RefreshTokenExpiresAt
}

// GetRefreshTokenExpiresAtOk returns a tuple with the RefreshTokenExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetRefreshTokenExpiresAtOk() (*int64, bool) {
	if o == nil || o.RefreshTokenExpiresAt == nil {
		return nil, false
	}
	return o.RefreshTokenExpiresAt, true
}

// HasRefreshTokenExpiresAt returns a boolean if a field has been set.
func (o *AccessToken) HasRefreshTokenExpiresAt() bool {
	if o != nil && o.RefreshTokenExpiresAt != nil {
		return true
	}

	return false
}

// SetRefreshTokenExpiresAt gets a reference to the given int64 and assigns it to the RefreshTokenExpiresAt field.
func (o *AccessToken) SetRefreshTokenExpiresAt(v int64) {
	o.RefreshTokenExpiresAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AccessToken) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AccessToken) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *AccessToken) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetLastRefreshedAt returns the LastRefreshedAt field value if set, zero value otherwise.
func (o *AccessToken) GetLastRefreshedAt() int64 {
	if o == nil || o.LastRefreshedAt == nil {
		var ret int64
		return ret
	}
	return *o.LastRefreshedAt
}

// GetLastRefreshedAtOk returns a tuple with the LastRefreshedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetLastRefreshedAtOk() (*int64, bool) {
	if o == nil || o.LastRefreshedAt == nil {
		return nil, false
	}
	return o.LastRefreshedAt, true
}

// HasLastRefreshedAt returns a boolean if a field has been set.
func (o *AccessToken) HasLastRefreshedAt() bool {
	if o != nil && o.LastRefreshedAt != nil {
		return true
	}

	return false
}

// SetLastRefreshedAt gets a reference to the given int64 and assigns it to the LastRefreshedAt field.
func (o *AccessToken) SetLastRefreshedAt(v int64) {
	o.LastRefreshedAt = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AccessToken) GetClientId() int64 {
	if o == nil || o.ClientId == nil {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetClientIdOk() (*int64, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AccessToken) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int64 and assigns it to the ClientId field.
func (o *AccessToken) SetClientId(v int64) {
	o.ClientId = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *AccessToken) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *AccessToken) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *AccessToken) SetSubject(v string) {
	o.Subject = &v
}

// GetGrantType returns the GrantType field value if set, zero value otherwise.
func (o *AccessToken) GetGrantType() GrantType {
	if o == nil || o.GrantType == nil {
		var ret GrantType
		return ret
	}
	return *o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetGrantTypeOk() (*GrantType, bool) {
	if o == nil || o.GrantType == nil {
		return nil, false
	}
	return o.GrantType, true
}

// HasGrantType returns a boolean if a field has been set.
func (o *AccessToken) HasGrantType() bool {
	if o != nil && o.GrantType != nil {
		return true
	}

	return false
}

// SetGrantType gets a reference to the given GrantType and assigns it to the GrantType field.
func (o *AccessToken) SetGrantType(v GrantType) {
	o.GrantType = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *AccessToken) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetScopesOk() ([]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *AccessToken) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *AccessToken) SetScopes(v []string) {
	o.Scopes = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *AccessToken) GetProperties() []Property {
	if o == nil || o.Properties == nil {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetPropertiesOk() ([]Property, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *AccessToken) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *AccessToken) SetProperties(v []Property) {
	o.Properties = v
}

func (o AccessToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessTokenHash != nil {
		toSerialize["accessTokenHash"] = o.AccessTokenHash
	}
	if o.AccessTokenExpiresAt != nil {
		toSerialize["accessTokenExpiresAt"] = o.AccessTokenExpiresAt
	}
	if o.RefreshTokenHash != nil {
		toSerialize["refreshTokenHash"] = o.RefreshTokenHash
	}
	if o.RefreshTokenExpiresAt != nil {
		toSerialize["refreshTokenExpiresAt"] = o.RefreshTokenExpiresAt
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.LastRefreshedAt != nil {
		toSerialize["lastRefreshedAt"] = o.LastRefreshedAt
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.GrantType != nil {
		toSerialize["grantType"] = o.GrantType
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableAccessToken struct {
	value *AccessToken
	isSet bool
}

func (v NullableAccessToken) Get() *AccessToken {
	return v.value
}

func (v *NullableAccessToken) Set(val *AccessToken) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessToken) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessToken(val *AccessToken) *NullableAccessToken {
	return &NullableAccessToken{value: val, isSet: true}
}

func (v NullableAccessToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


