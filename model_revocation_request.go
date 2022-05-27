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

// RevocationRequest struct for RevocationRequest
type RevocationRequest struct {
	// OAuth 2.0 token revocation request parameters which are the request parameters that the OAuth 2.0 token revocation endpoint ([RFC 7009](https://datatracker.ietf.org/doc/html/rfc7009)) of the authorization server implementation received from the client application.  The value of parameters is the entire entity body (which is formatted in `application/x-www-form-urlencoded`) of the request from the client application. 
	Parameters string `json:"parameters"`
	// The client ID extracted from `Authorization` header of the revocation request from the client application.  If the revocation endpoint of the authorization server implementation supports Basic Authentication as a means of client authentication, and the request from the client application contains its client ID in `Authorization` header, the value should be extracted and set to this parameter. 
	ClientId *string `json:"clientId,omitempty"`
	// The client secret extracted from `Authorization` header of the revocation request from the client application.  If the revocation endpoint of the authorization server implementation supports basic authentication as a means of client authentication, and the request from the client application contained its client secret in `Authorization` header, the value should be extracted and set to this parameter. 
	ClientSecret *string `json:"clientSecret,omitempty"`
	// The client certificate used in the TLS connection between the client application and the revocation endpoint. 
	ClientCertificate *string `json:"clientCertificate,omitempty"`
	// The certificate path presented by the client during client authentication.
	ClientCertificatePath *string `json:"clientCertificatePath,omitempty"`
}

// NewRevocationRequest instantiates a new RevocationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevocationRequest(parameters string) *RevocationRequest {
	this := RevocationRequest{}
	this.Parameters = parameters
	return &this
}

// NewRevocationRequestWithDefaults instantiates a new RevocationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevocationRequestWithDefaults() *RevocationRequest {
	this := RevocationRequest{}
	return &this
}

// GetParameters returns the Parameters field value
func (o *RevocationRequest) GetParameters() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *RevocationRequest) GetParametersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *RevocationRequest) SetParameters(v string) {
	o.Parameters = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *RevocationRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevocationRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *RevocationRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *RevocationRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *RevocationRequest) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevocationRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *RevocationRequest) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *RevocationRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetClientCertificate returns the ClientCertificate field value if set, zero value otherwise.
func (o *RevocationRequest) GetClientCertificate() string {
	if o == nil || o.ClientCertificate == nil {
		var ret string
		return ret
	}
	return *o.ClientCertificate
}

// GetClientCertificateOk returns a tuple with the ClientCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevocationRequest) GetClientCertificateOk() (*string, bool) {
	if o == nil || o.ClientCertificate == nil {
		return nil, false
	}
	return o.ClientCertificate, true
}

// HasClientCertificate returns a boolean if a field has been set.
func (o *RevocationRequest) HasClientCertificate() bool {
	if o != nil && o.ClientCertificate != nil {
		return true
	}

	return false
}

// SetClientCertificate gets a reference to the given string and assigns it to the ClientCertificate field.
func (o *RevocationRequest) SetClientCertificate(v string) {
	o.ClientCertificate = &v
}

// GetClientCertificatePath returns the ClientCertificatePath field value if set, zero value otherwise.
func (o *RevocationRequest) GetClientCertificatePath() string {
	if o == nil || o.ClientCertificatePath == nil {
		var ret string
		return ret
	}
	return *o.ClientCertificatePath
}

// GetClientCertificatePathOk returns a tuple with the ClientCertificatePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevocationRequest) GetClientCertificatePathOk() (*string, bool) {
	if o == nil || o.ClientCertificatePath == nil {
		return nil, false
	}
	return o.ClientCertificatePath, true
}

// HasClientCertificatePath returns a boolean if a field has been set.
func (o *RevocationRequest) HasClientCertificatePath() bool {
	if o != nil && o.ClientCertificatePath != nil {
		return true
	}

	return false
}

// SetClientCertificatePath gets a reference to the given string and assigns it to the ClientCertificatePath field.
func (o *RevocationRequest) SetClientCertificatePath(v string) {
	o.ClientCertificatePath = &v
}

func (o RevocationRequest) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableRevocationRequest struct {
	value *RevocationRequest
	isSet bool
}

func (v NullableRevocationRequest) Get() *RevocationRequest {
	return v.value
}

func (v *NullableRevocationRequest) Set(val *RevocationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRevocationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRevocationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevocationRequest(val *RevocationRequest) *NullableRevocationRequest {
	return &NullableRevocationRequest{value: val, isSet: true}
}

func (v NullableRevocationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevocationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

