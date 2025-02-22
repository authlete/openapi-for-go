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

// checks if the UserinfoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserinfoRequest{}

// UserinfoRequest struct for UserinfoRequest
type UserinfoRequest struct {
	// An access token.
	Token string `json:"token"`
	// Client certificate used in the TLS connection established between the client application and the userinfo endpoint.  The value of this request parameter is referred to when the access token given to the userinfo endpoint was bound to a client certificate when it was issued. See [OAuth 2.0 Mutual TLS Client Authentication and Certificate-Bound Access Tokens] (https://datatracker.ietf.org/doc/rfc8705/) for details about the specification of certificate-bound access tokens.
	ClientCertificate *string `json:"clientCertificate,omitempty"`
	// `DPoP` header presented by the client during the request to the user info endpoint.  The header contains a signed JWT which includes the public key that is paired with the private key used to sign the JWT. See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details.
	Dpop *string `json:"dpop,omitempty"`
	// HTTP method of the user info request. This field is used to validate the DPoP header. In normal cases, the value is either `GET` or `POST`.
	Htm *string `json:"htm,omitempty"`
	// URL of the user info endpoint. This field is used to validate the DPoP header.  If this parameter is omitted, the `userInfoEndpoint` property of the service is used as the default value. See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details.
	Htu *string `json:"htu,omitempty"`
	// The full URL of the userinfo endpoint.
	Uri *string `json:"uri,omitempty"`
	// The HTTP message body of the request, if present.
	Message *string `json:"message,omitempty"`
	// HTTP headers to be included in processing the signature. If this is a signed request, this must include the  Signature and Signature-Input headers, as well as any additional headers covered by the signature.
	Headers []Pair `json:"headers,omitempty"`
}

// NewUserinfoRequest instantiates a new UserinfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserinfoRequest(token string) *UserinfoRequest {
	this := UserinfoRequest{}
	this.Token = token
	return &this
}

// NewUserinfoRequestWithDefaults instantiates a new UserinfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserinfoRequestWithDefaults() *UserinfoRequest {
	this := UserinfoRequest{}
	return &this
}

// GetToken returns the Token field value
func (o *UserinfoRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UserinfoRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UserinfoRequest) SetToken(v string) {
	o.Token = v
}

// GetClientCertificate returns the ClientCertificate field value if set, zero value otherwise.
func (o *UserinfoRequest) GetClientCertificate() string {
	if o == nil || isNil(o.ClientCertificate) {
		var ret string
		return ret
	}
	return *o.ClientCertificate
}

// GetClientCertificateOk returns a tuple with the ClientCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoRequest) GetClientCertificateOk() (*string, bool) {
	if o == nil || isNil(o.ClientCertificate) {
		return nil, false
	}
	return o.ClientCertificate, true
}

// HasClientCertificate returns a boolean if a field has been set.
func (o *UserinfoRequest) HasClientCertificate() bool {
	if o != nil && !isNil(o.ClientCertificate) {
		return true
	}

	return false
}

// SetClientCertificate gets a reference to the given string and assigns it to the ClientCertificate field.
func (o *UserinfoRequest) SetClientCertificate(v string) {
	o.ClientCertificate = &v
}

// GetDpop returns the Dpop field value if set, zero value otherwise.
func (o *UserinfoRequest) GetDpop() string {
	if o == nil || isNil(o.Dpop) {
		var ret string
		return ret
	}
	return *o.Dpop
}

// GetDpopOk returns a tuple with the Dpop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoRequest) GetDpopOk() (*string, bool) {
	if o == nil || isNil(o.Dpop) {
		return nil, false
	}
	return o.Dpop, true
}

// HasDpop returns a boolean if a field has been set.
func (o *UserinfoRequest) HasDpop() bool {
	if o != nil && !isNil(o.Dpop) {
		return true
	}

	return false
}

// SetDpop gets a reference to the given string and assigns it to the Dpop field.
func (o *UserinfoRequest) SetDpop(v string) {
	o.Dpop = &v
}

// GetHtm returns the Htm field value if set, zero value otherwise.
func (o *UserinfoRequest) GetHtm() string {
	if o == nil || isNil(o.Htm) {
		var ret string
		return ret
	}
	return *o.Htm
}

// GetHtmOk returns a tuple with the Htm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoRequest) GetHtmOk() (*string, bool) {
	if o == nil || isNil(o.Htm) {
		return nil, false
	}
	return o.Htm, true
}

// HasHtm returns a boolean if a field has been set.
func (o *UserinfoRequest) HasHtm() bool {
	if o != nil && !isNil(o.Htm) {
		return true
	}

	return false
}

// SetHtm gets a reference to the given string and assigns it to the Htm field.
func (o *UserinfoRequest) SetHtm(v string) {
	o.Htm = &v
}

// GetHtu returns the Htu field value if set, zero value otherwise.
func (o *UserinfoRequest) GetHtu() string {
	if o == nil || isNil(o.Htu) {
		var ret string
		return ret
	}
	return *o.Htu
}

// GetHtuOk returns a tuple with the Htu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoRequest) GetHtuOk() (*string, bool) {
	if o == nil || isNil(o.Htu) {
		return nil, false
	}
	return o.Htu, true
}

// HasHtu returns a boolean if a field has been set.
func (o *UserinfoRequest) HasHtu() bool {
	if o != nil && !isNil(o.Htu) {
		return true
	}

	return false
}

// SetHtu gets a reference to the given string and assigns it to the Htu field.
func (o *UserinfoRequest) SetHtu(v string) {
	o.Htu = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *UserinfoRequest) GetUri() string {
	if o == nil || isNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoRequest) GetUriOk() (*string, bool) {
	if o == nil || isNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *UserinfoRequest) HasUri() bool {
	if o != nil && !isNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *UserinfoRequest) SetUri(v string) {
	o.Uri = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UserinfoRequest) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoRequest) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UserinfoRequest) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UserinfoRequest) SetMessage(v string) {
	o.Message = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *UserinfoRequest) GetHeaders() []Pair {
	if o == nil || isNil(o.Headers) {
		var ret []Pair
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoRequest) GetHeadersOk() ([]Pair, bool) {
	if o == nil || isNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *UserinfoRequest) HasHeaders() bool {
	if o != nil && !isNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []Pair and assigns it to the Headers field.
func (o *UserinfoRequest) SetHeaders(v []Pair) {
	o.Headers = v
}

func (o UserinfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserinfoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	if !isNil(o.ClientCertificate) {
		toSerialize["clientCertificate"] = o.ClientCertificate
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
	if !isNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	return toSerialize, nil
}

type NullableUserinfoRequest struct {
	value *UserinfoRequest
	isSet bool
}

func (v NullableUserinfoRequest) Get() *UserinfoRequest {
	return v.value
}

func (v *NullableUserinfoRequest) Set(val *UserinfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserinfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserinfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserinfoRequest(val *UserinfoRequest) *NullableUserinfoRequest {
	return &NullableUserinfoRequest{value: val, isSet: true}
}

func (v NullableUserinfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserinfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
