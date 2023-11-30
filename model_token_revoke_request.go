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

// checks if the TokenRevokeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenRevokeRequest{}

// TokenRevokeRequest struct for TokenRevokeRequest
type TokenRevokeRequest struct {
	// The identifier of an access token to revoke  The hash of an access token is recognized as an identifier as well as the access token itself. 
	AccessTokenIdentifier string `json:"accessTokenIdentifier"`
	// The identifier of a refresh token to revoke.  The hash of a refresh token is recognized as an identifier as well as the refresh token itself. 
	RefreshTokenIdentifier *string `json:"refreshTokenIdentifier,omitempty"`
	// The client ID of the access token to be revoked.  Both the numeric client ID and the alias are recognized as an identifier of a client. 
	ClientIdentifier *string `json:"clientIdentifier,omitempty"`
	// The subject of a resource owner. 
	Subject *string `json:"subject,omitempty"`
}

// NewTokenRevokeRequest instantiates a new TokenRevokeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRevokeRequest(accessTokenIdentifier string) *TokenRevokeRequest {
	this := TokenRevokeRequest{}
	this.AccessTokenIdentifier = accessTokenIdentifier
	return &this
}

// NewTokenRevokeRequestWithDefaults instantiates a new TokenRevokeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRevokeRequestWithDefaults() *TokenRevokeRequest {
	this := TokenRevokeRequest{}
	return &this
}

// GetAccessTokenIdentifier returns the AccessTokenIdentifier field value
func (o *TokenRevokeRequest) GetAccessTokenIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessTokenIdentifier
}

// GetAccessTokenIdentifierOk returns a tuple with the AccessTokenIdentifier field value
// and a boolean to check if the value has been set.
func (o *TokenRevokeRequest) GetAccessTokenIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessTokenIdentifier, true
}

// SetAccessTokenIdentifier sets field value
func (o *TokenRevokeRequest) SetAccessTokenIdentifier(v string) {
	o.AccessTokenIdentifier = v
}

// GetRefreshTokenIdentifier returns the RefreshTokenIdentifier field value if set, zero value otherwise.
func (o *TokenRevokeRequest) GetRefreshTokenIdentifier() string {
	if o == nil || isNil(o.RefreshTokenIdentifier) {
		var ret string
		return ret
	}
	return *o.RefreshTokenIdentifier
}

// GetRefreshTokenIdentifierOk returns a tuple with the RefreshTokenIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevokeRequest) GetRefreshTokenIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.RefreshTokenIdentifier) {
		return nil, false
	}
	return o.RefreshTokenIdentifier, true
}

// HasRefreshTokenIdentifier returns a boolean if a field has been set.
func (o *TokenRevokeRequest) HasRefreshTokenIdentifier() bool {
	if o != nil && !isNil(o.RefreshTokenIdentifier) {
		return true
	}

	return false
}

// SetRefreshTokenIdentifier gets a reference to the given string and assigns it to the RefreshTokenIdentifier field.
func (o *TokenRevokeRequest) SetRefreshTokenIdentifier(v string) {
	o.RefreshTokenIdentifier = &v
}

// GetClientIdentifier returns the ClientIdentifier field value if set, zero value otherwise.
func (o *TokenRevokeRequest) GetClientIdentifier() string {
	if o == nil || isNil(o.ClientIdentifier) {
		var ret string
		return ret
	}
	return *o.ClientIdentifier
}

// GetClientIdentifierOk returns a tuple with the ClientIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevokeRequest) GetClientIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.ClientIdentifier) {
		return nil, false
	}
	return o.ClientIdentifier, true
}

// HasClientIdentifier returns a boolean if a field has been set.
func (o *TokenRevokeRequest) HasClientIdentifier() bool {
	if o != nil && !isNil(o.ClientIdentifier) {
		return true
	}

	return false
}

// SetClientIdentifier gets a reference to the given string and assigns it to the ClientIdentifier field.
func (o *TokenRevokeRequest) SetClientIdentifier(v string) {
	o.ClientIdentifier = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *TokenRevokeRequest) GetSubject() string {
	if o == nil || isNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevokeRequest) GetSubjectOk() (*string, bool) {
	if o == nil || isNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *TokenRevokeRequest) HasSubject() bool {
	if o != nil && !isNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *TokenRevokeRequest) SetSubject(v string) {
	o.Subject = &v
}

func (o TokenRevokeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenRevokeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accessTokenIdentifier"] = o.AccessTokenIdentifier
	if !isNil(o.RefreshTokenIdentifier) {
		toSerialize["refreshTokenIdentifier"] = o.RefreshTokenIdentifier
	}
	if !isNil(o.ClientIdentifier) {
		toSerialize["clientIdentifier"] = o.ClientIdentifier
	}
	if !isNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	return toSerialize, nil
}

type NullableTokenRevokeRequest struct {
	value *TokenRevokeRequest
	isSet bool
}

func (v NullableTokenRevokeRequest) Get() *TokenRevokeRequest {
	return v.value
}

func (v *NullableTokenRevokeRequest) Set(val *TokenRevokeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRevokeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRevokeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRevokeRequest(val *TokenRevokeRequest) *NullableTokenRevokeRequest {
	return &NullableTokenRevokeRequest{value: val, isSet: true}
}

func (v NullableTokenRevokeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRevokeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


