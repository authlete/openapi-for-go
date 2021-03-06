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

// TokenRequest struct for TokenRequest
type TokenRequest struct {
	// OAuth 2.0 token request parameters which are the request parameters that the OAuth 2.0 token endpoint of the authorization server implementation received from the client application.  The value of parameters is the entire entity body (which is formatted in `application/x-www-form-urlencoded`) of the request from the client application. 
	Parameters string `json:"parameters"`
	// The client ID extracted from `Authorization` header of the token request from the client application.  If the token endpoint of the authorization server implementation supports basic authentication as a means of client authentication, and the request from the client application contained its client ID in `Authorization` header, the value should be extracted and set to this parameter. 
	ClientId *string `json:"clientId,omitempty"`
	// The client secret extracted from `Authorization` header of the token request from the client application.  If the token endpoint of the authorization server implementation supports basic authentication as a means of client authentication, and the request from the client application contained its client secret in `Authorization` header, the value should be extracted and set to this parameter. 
	ClientSecret *string `json:"clientSecret,omitempty"`
	// The client certificate from the MTLS of the token request from the client application.
	ClientCertificate *string `json:"clientCertificate,omitempty"`
	// The certificate path presented by the client during client authentication. These certificates are strings in PEM format. 
	ClientCertificatePath *string `json:"clientCertificatePath,omitempty"`
	// Extra properties to associate with an access token. See [Extra Properties](https://www.authlete.com/developers/definitive_guide/extra_properties/) for details. 
	Properties *string `json:"properties,omitempty"`
	// `DPoP` header presented by the client during the request to the token endpoint.  The header contains a signed JWT which includes the public key that is paired with the private key used to sign the JWT. See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details. 
	Dpop *string `json:"dpop,omitempty"`
	// HTTP method of the token request. This field is used to validate the `DPoP` header.  In normal cases, the value is `POST`. When this parameter is omitted, `POST` is used as the default value. See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details. 
	Htm *string `json:"htm,omitempty"`
	// URL of the token endpoint. This field is used to validate the `DPoP` header.  If this parameter is omitted, the `tokenEndpoint` property of the Service is used as the default value. See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details.
	Htu *string `json:"htu,omitempty"`
}

// NewTokenRequest instantiates a new TokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRequest(parameters string) *TokenRequest {
	this := TokenRequest{}
	this.Parameters = parameters
	return &this
}

// NewTokenRequestWithDefaults instantiates a new TokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRequestWithDefaults() *TokenRequest {
	this := TokenRequest{}
	return &this
}

// GetParameters returns the Parameters field value
func (o *TokenRequest) GetParameters() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetParametersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *TokenRequest) SetParameters(v string) {
	o.Parameters = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TokenRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TokenRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TokenRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *TokenRequest) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *TokenRequest) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *TokenRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetClientCertificate returns the ClientCertificate field value if set, zero value otherwise.
func (o *TokenRequest) GetClientCertificate() string {
	if o == nil || o.ClientCertificate == nil {
		var ret string
		return ret
	}
	return *o.ClientCertificate
}

// GetClientCertificateOk returns a tuple with the ClientCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetClientCertificateOk() (*string, bool) {
	if o == nil || o.ClientCertificate == nil {
		return nil, false
	}
	return o.ClientCertificate, true
}

// HasClientCertificate returns a boolean if a field has been set.
func (o *TokenRequest) HasClientCertificate() bool {
	if o != nil && o.ClientCertificate != nil {
		return true
	}

	return false
}

// SetClientCertificate gets a reference to the given string and assigns it to the ClientCertificate field.
func (o *TokenRequest) SetClientCertificate(v string) {
	o.ClientCertificate = &v
}

// GetClientCertificatePath returns the ClientCertificatePath field value if set, zero value otherwise.
func (o *TokenRequest) GetClientCertificatePath() string {
	if o == nil || o.ClientCertificatePath == nil {
		var ret string
		return ret
	}
	return *o.ClientCertificatePath
}

// GetClientCertificatePathOk returns a tuple with the ClientCertificatePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetClientCertificatePathOk() (*string, bool) {
	if o == nil || o.ClientCertificatePath == nil {
		return nil, false
	}
	return o.ClientCertificatePath, true
}

// HasClientCertificatePath returns a boolean if a field has been set.
func (o *TokenRequest) HasClientCertificatePath() bool {
	if o != nil && o.ClientCertificatePath != nil {
		return true
	}

	return false
}

// SetClientCertificatePath gets a reference to the given string and assigns it to the ClientCertificatePath field.
func (o *TokenRequest) SetClientCertificatePath(v string) {
	o.ClientCertificatePath = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *TokenRequest) GetProperties() string {
	if o == nil || o.Properties == nil {
		var ret string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetPropertiesOk() (*string, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *TokenRequest) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given string and assigns it to the Properties field.
func (o *TokenRequest) SetProperties(v string) {
	o.Properties = &v
}

// GetDpop returns the Dpop field value if set, zero value otherwise.
func (o *TokenRequest) GetDpop() string {
	if o == nil || o.Dpop == nil {
		var ret string
		return ret
	}
	return *o.Dpop
}

// GetDpopOk returns a tuple with the Dpop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetDpopOk() (*string, bool) {
	if o == nil || o.Dpop == nil {
		return nil, false
	}
	return o.Dpop, true
}

// HasDpop returns a boolean if a field has been set.
func (o *TokenRequest) HasDpop() bool {
	if o != nil && o.Dpop != nil {
		return true
	}

	return false
}

// SetDpop gets a reference to the given string and assigns it to the Dpop field.
func (o *TokenRequest) SetDpop(v string) {
	o.Dpop = &v
}

// GetHtm returns the Htm field value if set, zero value otherwise.
func (o *TokenRequest) GetHtm() string {
	if o == nil || o.Htm == nil {
		var ret string
		return ret
	}
	return *o.Htm
}

// GetHtmOk returns a tuple with the Htm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetHtmOk() (*string, bool) {
	if o == nil || o.Htm == nil {
		return nil, false
	}
	return o.Htm, true
}

// HasHtm returns a boolean if a field has been set.
func (o *TokenRequest) HasHtm() bool {
	if o != nil && o.Htm != nil {
		return true
	}

	return false
}

// SetHtm gets a reference to the given string and assigns it to the Htm field.
func (o *TokenRequest) SetHtm(v string) {
	o.Htm = &v
}

// GetHtu returns the Htu field value if set, zero value otherwise.
func (o *TokenRequest) GetHtu() string {
	if o == nil || o.Htu == nil {
		var ret string
		return ret
	}
	return *o.Htu
}

// GetHtuOk returns a tuple with the Htu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetHtuOk() (*string, bool) {
	if o == nil || o.Htu == nil {
		return nil, false
	}
	return o.Htu, true
}

// HasHtu returns a boolean if a field has been set.
func (o *TokenRequest) HasHtu() bool {
	if o != nil && o.Htu != nil {
		return true
	}

	return false
}

// SetHtu gets a reference to the given string and assigns it to the Htu field.
func (o *TokenRequest) SetHtu(v string) {
	o.Htu = &v
}

func (o TokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["parameters"] = o.Parameters
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if o.ClientCertificate != nil {
		toSerialize["clientCertificate"] = o.ClientCertificate
	}
	if o.ClientCertificatePath != nil {
		toSerialize["clientCertificatePath"] = o.ClientCertificatePath
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Dpop != nil {
		toSerialize["dpop"] = o.Dpop
	}
	if o.Htm != nil {
		toSerialize["htm"] = o.Htm
	}
	if o.Htu != nil {
		toSerialize["htu"] = o.Htu
	}
	return json.Marshal(toSerialize)
}

type NullableTokenRequest struct {
	value *TokenRequest
	isSet bool
}

func (v NullableTokenRequest) Get() *TokenRequest {
	return v.value
}

func (v *NullableTokenRequest) Set(val *TokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRequest(val *TokenRequest) *NullableTokenRequest {
	return &NullableTokenRequest{value: val, isSet: true}
}

func (v NullableTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


