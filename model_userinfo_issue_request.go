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

// checks if the UserinfoIssueRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserinfoIssueRequest{}

// UserinfoIssueRequest struct for UserinfoIssueRequest
type UserinfoIssueRequest struct {
	// The access token that has been passed to the userinfo endpoint by the client application. In other words, the access token which was contained in the userinfo request. 
	Token string `json:"token"`
	// Claims in JSON format. As for the format, see [OpenID Connect Core 1.0, 5.1. Standard Claims](https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims). 
	Claims *string `json:"claims,omitempty"`
	// The value of the `sub` claim. If the value of this request parameter is not empty, it is used as the value of the `sub` claim. Otherwise, the value of the subject associated with the access token is used. 
	Sub *string `json:"sub,omitempty"`
	// Claim key-value pairs that are used to compute transformed claims. 
	ClaimsForTx *string `json:"claimsForTx,omitempty"`
	// The Signature header value from the request. 
	RequestSignature *string `json:"requestSignature,omitempty"`
	// HTTP headers to be included in processing the signature. If this is a signed request, this must include the  Signature and Signature-Input headers, as well as any additional headers covered by the signature.
	Headers []Pair `json:"headers,omitempty"`
}

// NewUserinfoIssueRequest instantiates a new UserinfoIssueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserinfoIssueRequest(token string) *UserinfoIssueRequest {
	this := UserinfoIssueRequest{}
	this.Token = token
	return &this
}

// NewUserinfoIssueRequestWithDefaults instantiates a new UserinfoIssueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserinfoIssueRequestWithDefaults() *UserinfoIssueRequest {
	this := UserinfoIssueRequest{}
	return &this
}

// GetToken returns the Token field value
func (o *UserinfoIssueRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UserinfoIssueRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UserinfoIssueRequest) SetToken(v string) {
	o.Token = v
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *UserinfoIssueRequest) GetClaims() string {
	if o == nil || isNil(o.Claims) {
		var ret string
		return ret
	}
	return *o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueRequest) GetClaimsOk() (*string, bool) {
	if o == nil || isNil(o.Claims) {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *UserinfoIssueRequest) HasClaims() bool {
	if o != nil && !isNil(o.Claims) {
		return true
	}

	return false
}

// SetClaims gets a reference to the given string and assigns it to the Claims field.
func (o *UserinfoIssueRequest) SetClaims(v string) {
	o.Claims = &v
}

// GetSub returns the Sub field value if set, zero value otherwise.
func (o *UserinfoIssueRequest) GetSub() string {
	if o == nil || isNil(o.Sub) {
		var ret string
		return ret
	}
	return *o.Sub
}

// GetSubOk returns a tuple with the Sub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueRequest) GetSubOk() (*string, bool) {
	if o == nil || isNil(o.Sub) {
		return nil, false
	}
	return o.Sub, true
}

// HasSub returns a boolean if a field has been set.
func (o *UserinfoIssueRequest) HasSub() bool {
	if o != nil && !isNil(o.Sub) {
		return true
	}

	return false
}

// SetSub gets a reference to the given string and assigns it to the Sub field.
func (o *UserinfoIssueRequest) SetSub(v string) {
	o.Sub = &v
}

// GetClaimsForTx returns the ClaimsForTx field value if set, zero value otherwise.
func (o *UserinfoIssueRequest) GetClaimsForTx() string {
	if o == nil || isNil(o.ClaimsForTx) {
		var ret string
		return ret
	}
	return *o.ClaimsForTx
}

// GetClaimsForTxOk returns a tuple with the ClaimsForTx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueRequest) GetClaimsForTxOk() (*string, bool) {
	if o == nil || isNil(o.ClaimsForTx) {
		return nil, false
	}
	return o.ClaimsForTx, true
}

// HasClaimsForTx returns a boolean if a field has been set.
func (o *UserinfoIssueRequest) HasClaimsForTx() bool {
	if o != nil && !isNil(o.ClaimsForTx) {
		return true
	}

	return false
}

// SetClaimsForTx gets a reference to the given string and assigns it to the ClaimsForTx field.
func (o *UserinfoIssueRequest) SetClaimsForTx(v string) {
	o.ClaimsForTx = &v
}

// GetRequestSignature returns the RequestSignature field value if set, zero value otherwise.
func (o *UserinfoIssueRequest) GetRequestSignature() string {
	if o == nil || isNil(o.RequestSignature) {
		var ret string
		return ret
	}
	return *o.RequestSignature
}

// GetRequestSignatureOk returns a tuple with the RequestSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueRequest) GetRequestSignatureOk() (*string, bool) {
	if o == nil || isNil(o.RequestSignature) {
		return nil, false
	}
	return o.RequestSignature, true
}

// HasRequestSignature returns a boolean if a field has been set.
func (o *UserinfoIssueRequest) HasRequestSignature() bool {
	if o != nil && !isNil(o.RequestSignature) {
		return true
	}

	return false
}

// SetRequestSignature gets a reference to the given string and assigns it to the RequestSignature field.
func (o *UserinfoIssueRequest) SetRequestSignature(v string) {
	o.RequestSignature = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *UserinfoIssueRequest) GetHeaders() []Pair {
	if o == nil || isNil(o.Headers) {
		var ret []Pair
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueRequest) GetHeadersOk() ([]Pair, bool) {
	if o == nil || isNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *UserinfoIssueRequest) HasHeaders() bool {
	if o != nil && !isNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []Pair and assigns it to the Headers field.
func (o *UserinfoIssueRequest) SetHeaders(v []Pair) {
	o.Headers = v
}

func (o UserinfoIssueRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserinfoIssueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	if !isNil(o.Claims) {
		toSerialize["claims"] = o.Claims
	}
	if !isNil(o.Sub) {
		toSerialize["sub"] = o.Sub
	}
	if !isNil(o.ClaimsForTx) {
		toSerialize["claimsForTx"] = o.ClaimsForTx
	}
	if !isNil(o.RequestSignature) {
		toSerialize["requestSignature"] = o.RequestSignature
	}
	if !isNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	return toSerialize, nil
}

type NullableUserinfoIssueRequest struct {
	value *UserinfoIssueRequest
	isSet bool
}

func (v NullableUserinfoIssueRequest) Get() *UserinfoIssueRequest {
	return v.value
}

func (v *NullableUserinfoIssueRequest) Set(val *UserinfoIssueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserinfoIssueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserinfoIssueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserinfoIssueRequest(val *UserinfoIssueRequest) *NullableUserinfoIssueRequest {
	return &NullableUserinfoIssueRequest{value: val, isSet: true}
}

func (v NullableUserinfoIssueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserinfoIssueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


