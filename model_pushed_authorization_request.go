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

// checks if the PushedAuthorizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PushedAuthorizationRequest{}

// PushedAuthorizationRequest struct for PushedAuthorizationRequest
type PushedAuthorizationRequest struct {
	// The pushed authorization request body received from the client application.  The value of parameters is the entire entity body (which is formatted in `application/x-www-form-urlencoded`) of the request from the client application. 
	Parameters string `json:"parameters"`
	// The client ID extracted from `Authorization` header of the pushed request from the client application. 
	ClientId *string `json:"clientId,omitempty"`
	// The client secret extracted from `Authorization` header of the pushed authorization request from the client application. 
	ClientSecret *string `json:"clientSecret,omitempty"`
	// The client certificate from the MTLS connection to pushed authorization endpoint from the client application.
	ClientCertificate *string `json:"clientCertificate,omitempty"`
	// The certificate path presented by the client during client authentication. These certificates are strings in PEM format. 
	ClientCertificatePath *string `json:"clientCertificatePath,omitempty"`
	// DPoP Header 
	Dpop *string `json:"dpop,omitempty"`
	// HTTP Method (for DPoP validation). 
	Htm *string `json:"htm,omitempty"`
	// HTTP URL base (for DPoP validation).
	Htu *string `json:"htu,omitempty"`
}

// NewPushedAuthorizationRequest instantiates a new PushedAuthorizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushedAuthorizationRequest(parameters string) *PushedAuthorizationRequest {
	this := PushedAuthorizationRequest{}
	this.Parameters = parameters
	return &this
}

// NewPushedAuthorizationRequestWithDefaults instantiates a new PushedAuthorizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushedAuthorizationRequestWithDefaults() *PushedAuthorizationRequest {
	this := PushedAuthorizationRequest{}
	return &this
}

// GetParameters returns the Parameters field value
func (o *PushedAuthorizationRequest) GetParameters() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationRequest) GetParametersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *PushedAuthorizationRequest) SetParameters(v string) {
	o.Parameters = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PushedAuthorizationRequest) GetClientId() string {
	if o == nil || isNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationRequest) GetClientIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PushedAuthorizationRequest) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PushedAuthorizationRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *PushedAuthorizationRequest) GetClientSecret() string {
	if o == nil || isNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || isNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *PushedAuthorizationRequest) HasClientSecret() bool {
	if o != nil && !isNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *PushedAuthorizationRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetClientCertificate returns the ClientCertificate field value if set, zero value otherwise.
func (o *PushedAuthorizationRequest) GetClientCertificate() string {
	if o == nil || isNil(o.ClientCertificate) {
		var ret string
		return ret
	}
	return *o.ClientCertificate
}

// GetClientCertificateOk returns a tuple with the ClientCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationRequest) GetClientCertificateOk() (*string, bool) {
	if o == nil || isNil(o.ClientCertificate) {
		return nil, false
	}
	return o.ClientCertificate, true
}

// HasClientCertificate returns a boolean if a field has been set.
func (o *PushedAuthorizationRequest) HasClientCertificate() bool {
	if o != nil && !isNil(o.ClientCertificate) {
		return true
	}

	return false
}

// SetClientCertificate gets a reference to the given string and assigns it to the ClientCertificate field.
func (o *PushedAuthorizationRequest) SetClientCertificate(v string) {
	o.ClientCertificate = &v
}

// GetClientCertificatePath returns the ClientCertificatePath field value if set, zero value otherwise.
func (o *PushedAuthorizationRequest) GetClientCertificatePath() string {
	if o == nil || isNil(o.ClientCertificatePath) {
		var ret string
		return ret
	}
	return *o.ClientCertificatePath
}

// GetClientCertificatePathOk returns a tuple with the ClientCertificatePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationRequest) GetClientCertificatePathOk() (*string, bool) {
	if o == nil || isNil(o.ClientCertificatePath) {
		return nil, false
	}
	return o.ClientCertificatePath, true
}

// HasClientCertificatePath returns a boolean if a field has been set.
func (o *PushedAuthorizationRequest) HasClientCertificatePath() bool {
	if o != nil && !isNil(o.ClientCertificatePath) {
		return true
	}

	return false
}

// SetClientCertificatePath gets a reference to the given string and assigns it to the ClientCertificatePath field.
func (o *PushedAuthorizationRequest) SetClientCertificatePath(v string) {
	o.ClientCertificatePath = &v
}

// GetDpop returns the Dpop field value if set, zero value otherwise.
func (o *PushedAuthorizationRequest) GetDpop() string {
	if o == nil || isNil(o.Dpop) {
		var ret string
		return ret
	}
	return *o.Dpop
}

// GetDpopOk returns a tuple with the Dpop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationRequest) GetDpopOk() (*string, bool) {
	if o == nil || isNil(o.Dpop) {
		return nil, false
	}
	return o.Dpop, true
}

// HasDpop returns a boolean if a field has been set.
func (o *PushedAuthorizationRequest) HasDpop() bool {
	if o != nil && !isNil(o.Dpop) {
		return true
	}

	return false
}

// SetDpop gets a reference to the given string and assigns it to the Dpop field.
func (o *PushedAuthorizationRequest) SetDpop(v string) {
	o.Dpop = &v
}

// GetHtm returns the Htm field value if set, zero value otherwise.
func (o *PushedAuthorizationRequest) GetHtm() string {
	if o == nil || isNil(o.Htm) {
		var ret string
		return ret
	}
	return *o.Htm
}

// GetHtmOk returns a tuple with the Htm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationRequest) GetHtmOk() (*string, bool) {
	if o == nil || isNil(o.Htm) {
		return nil, false
	}
	return o.Htm, true
}

// HasHtm returns a boolean if a field has been set.
func (o *PushedAuthorizationRequest) HasHtm() bool {
	if o != nil && !isNil(o.Htm) {
		return true
	}

	return false
}

// SetHtm gets a reference to the given string and assigns it to the Htm field.
func (o *PushedAuthorizationRequest) SetHtm(v string) {
	o.Htm = &v
}

// GetHtu returns the Htu field value if set, zero value otherwise.
func (o *PushedAuthorizationRequest) GetHtu() string {
	if o == nil || isNil(o.Htu) {
		var ret string
		return ret
	}
	return *o.Htu
}

// GetHtuOk returns a tuple with the Htu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationRequest) GetHtuOk() (*string, bool) {
	if o == nil || isNil(o.Htu) {
		return nil, false
	}
	return o.Htu, true
}

// HasHtu returns a boolean if a field has been set.
func (o *PushedAuthorizationRequest) HasHtu() bool {
	if o != nil && !isNil(o.Htu) {
		return true
	}

	return false
}

// SetHtu gets a reference to the given string and assigns it to the Htu field.
func (o *PushedAuthorizationRequest) SetHtu(v string) {
	o.Htu = &v
}

func (o PushedAuthorizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PushedAuthorizationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parameters"] = o.Parameters
	if !isNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !isNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if !isNil(o.ClientCertificate) {
		toSerialize["clientCertificate"] = o.ClientCertificate
	}
	if !isNil(o.ClientCertificatePath) {
		toSerialize["clientCertificatePath"] = o.ClientCertificatePath
	}
	if !isNil(o.Dpop) {
		toSerialize["dpop"] = o.Dpop
	}
	if !isNil(o.Htm) {
		toSerialize["htm"] = o.Htm
	}
	if !isNil(o.Htu) {
		toSerialize["htu"] = o.Htu
	}
	return toSerialize, nil
}

type NullablePushedAuthorizationRequest struct {
	value *PushedAuthorizationRequest
	isSet bool
}

func (v NullablePushedAuthorizationRequest) Get() *PushedAuthorizationRequest {
	return v.value
}

func (v *NullablePushedAuthorizationRequest) Set(val *PushedAuthorizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePushedAuthorizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePushedAuthorizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushedAuthorizationRequest(val *PushedAuthorizationRequest) *NullablePushedAuthorizationRequest {
	return &NullablePushedAuthorizationRequest{value: val, isSet: true}
}

func (v NullablePushedAuthorizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushedAuthorizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


