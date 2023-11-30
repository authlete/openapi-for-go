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

// checks if the IntrospectionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntrospectionRequest{}

// IntrospectionRequest struct for IntrospectionRequest
type IntrospectionRequest struct {
	// An access token to introspect.
	Token string `json:"token"`
	// A string array listing names of scopes which the caller (= a protected resource endpoint of the service) requires. When the content type of the request from the service is `application/x-www-form-urlencoded`, the format of `scopes` is a space-separated list of scope names.  If this parameter is a non-empty array and if it contains a scope which is not covered by the access token,`action=FORBIDDEN` with `error=insufficient_scope` is returned from Authlete. 
	Scopes []string `json:"scopes,omitempty"`
	// A subject (= a user account managed by the service) whom the caller (= a protected resource endpoint of the service) requires.  If this parameter is not `null` and if the value does not match the subject who is associated with the access token, `action=FORBIDDEN` with `error=invalid_request` is returned from Authlete. 
	Subject *string `json:"subject,omitempty"`
	// Client certificate in PEM format, used to validate binding against access tokens using the TLS client certificate confirmation method. 
	ClientCertificate *string `json:"clientCertificate,omitempty"`
	// `DPoP` header presented by the client during the request to the resource server.  The header contains a signed JWT which includes the public key that is paired with the private key used to sign the JWT. See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details. 
	Dpop *string `json:"dpop,omitempty"`
	// HTTP method of the request from the client to the protected resource endpoint. This field is used to validate the `DPoP` header.  See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details. 
	Htm *string `json:"htm,omitempty"`
	// URL of the protected resource endpoint. This field is used to validate the `DPoP` header.  See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details. 
	Htu *string `json:"htu,omitempty"`
	// The resources specified by the `resource` request parameters in the token request. See \"Resource Indicators for OAuth 2.0\" for details. 
	Resources []string `json:"resources,omitempty"`
	// Authentication Context Class Reference values one of which the user authentication performed during the course  of issuing the access token must satisfy. 
	AcrValues []string `json:"acrValues,omitempty"`
	// The maximum authentication age which is the maximum allowable elapsed time since the user authentication  was performed during the course of issuing the access token. 
	MaxAge *int64 `json:"maxAge,omitempty"`
	// HTTP Message Components required to be in the signature. If absent, defaults to [ \"@method\", \"@target-uri\", \"authorization\" ]. 
	RequiredComponents []string `json:"requiredComponents,omitempty"`
	// The full URL of the userinfo endpoint. 
	Uri *string `json:"uri,omitempty"`
	// The HTTP message body of the request, if present. 
	Message *string `json:"message,omitempty"`
	// HTTP headers to be included in processing the signature. If this is a signed request, this must include the  Signature and Signature-Input headers, as well as any additional headers covered by the signature. 
	Headers []Pair `json:"headers,omitempty"`
}

// NewIntrospectionRequest instantiates a new IntrospectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntrospectionRequest(token string) *IntrospectionRequest {
	this := IntrospectionRequest{}
	this.Token = token
	return &this
}

// NewIntrospectionRequestWithDefaults instantiates a new IntrospectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntrospectionRequestWithDefaults() *IntrospectionRequest {
	this := IntrospectionRequest{}
	return &this
}

// GetToken returns the Token field value
func (o *IntrospectionRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *IntrospectionRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *IntrospectionRequest) SetToken(v string) {
	o.Token = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *IntrospectionRequest) GetScopes() []string {
	if o == nil || isNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionRequest) GetScopesOk() ([]string, bool) {
	if o == nil || isNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *IntrospectionRequest) HasScopes() bool {
	if o != nil && !isNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *IntrospectionRequest) SetScopes(v []string) {
	o.Scopes = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *IntrospectionRequest) GetSubject() string {
	if o == nil || isNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionRequest) GetSubjectOk() (*string, bool) {
	if o == nil || isNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *IntrospectionRequest) HasSubject() bool {
	if o != nil && !isNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *IntrospectionRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetClientCertificate returns the ClientCertificate field value if set, zero value otherwise.
func (o *IntrospectionRequest) GetClientCertificate() string {
	if o == nil || isNil(o.ClientCertificate) {
		var ret string
		return ret
	}
	return *o.ClientCertificate
}

// GetClientCertificateOk returns a tuple with the ClientCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionRequest) GetClientCertificateOk() (*string, bool) {
	if o == nil || isNil(o.ClientCertificate) {
		return nil, false
	}
	return o.ClientCertificate, true
}

// HasClientCertificate returns a boolean if a field has been set.
func (o *IntrospectionRequest) HasClientCertificate() bool {
	if o != nil && !isNil(o.ClientCertificate) {
		return true
	}

	return false
}

// SetClientCertificate gets a reference to the given string and assigns it to the ClientCertificate field.
func (o *IntrospectionRequest) SetClientCertificate(v string) {
	o.ClientCertificate = &v
}

// GetDpop returns the Dpop field value if set, zero value otherwise.
func (o *IntrospectionRequest) GetDpop() string {
	if o == nil || isNil(o.Dpop) {
		var ret string
		return ret
	}
	return *o.Dpop
}

// GetDpopOk returns a tuple with the Dpop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionRequest) GetDpopOk() (*string, bool) {
	if o == nil || isNil(o.Dpop) {
		return nil, false
	}
	return o.Dpop, true
}

// HasDpop returns a boolean if a field has been set.
func (o *IntrospectionRequest) HasDpop() bool {
	if o != nil && !isNil(o.Dpop) {
		return true
	}

	return false
}

// SetDpop gets a reference to the given string and assigns it to the Dpop field.
func (o *IntrospectionRequest) SetDpop(v string) {
	o.Dpop = &v
}

// GetHtm returns the Htm field value if set, zero value otherwise.
func (o *IntrospectionRequest) GetHtm() string {
	if o == nil || isNil(o.Htm) {
		var ret string
		return ret
	}
	return *o.Htm
}

// GetHtmOk returns a tuple with the Htm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionRequest) GetHtmOk() (*string, bool) {
	if o == nil || isNil(o.Htm) {
		return nil, false
	}
	return o.Htm, true
}

// HasHtm returns a boolean if a field has been set.
func (o *IntrospectionRequest) HasHtm() bool {
	if o != nil && !isNil(o.Htm) {
		return true
	}

	return false
}

// SetHtm gets a reference to the given string and assigns it to the Htm field.
func (o *IntrospectionRequest) SetHtm(v string) {
	o.Htm = &v
}

// GetHtu returns the Htu field value if set, zero value otherwise.
func (o *IntrospectionRequest) GetHtu() string {
	if o == nil || isNil(o.Htu) {
		var ret string
		return ret
	}
	return *o.Htu
}

// GetHtuOk returns a tuple with the Htu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionRequest) GetHtuOk() (*string, bool) {
	if o == nil || isNil(o.Htu) {
		return nil, false
	}
	return o.Htu, true
}

// HasHtu returns a boolean if a field has been set.
func (o *IntrospectionRequest) HasHtu() bool {
	if o != nil && !isNil(o.Htu) {
		return true
	}

	return false
}

// SetHtu gets a reference to the given string and assigns it to the Htu field.
func (o *IntrospectionRequest) SetHtu(v string) {
	o.Htu = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *IntrospectionRequest) GetResources() []string {
	if o == nil || isNil(o.Resources) {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionRequest) GetResourcesOk() ([]string, bool) {
	if o == nil || isNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *IntrospectionRequest) HasResources() bool {
	if o != nil && !isNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *IntrospectionRequest) SetResources(v []string) {
	o.Resources = v
}

// GetAcrValues returns the AcrValues field value if set, zero value otherwise.
func (o *IntrospectionRequest) GetAcrValues() []string {
	if o == nil || isNil(o.AcrValues) {
		var ret []string
		return ret
	}
	return o.AcrValues
}

// GetAcrValuesOk returns a tuple with the AcrValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionRequest) GetAcrValuesOk() ([]string, bool) {
	if o == nil || isNil(o.AcrValues) {
		return nil, false
	}
	return o.AcrValues, true
}

// HasAcrValues returns a boolean if a field has been set.
func (o *IntrospectionRequest) HasAcrValues() bool {
	if o != nil && !isNil(o.AcrValues) {
		return true
	}

	return false
}

// SetAcrValues gets a reference to the given []string and assigns it to the AcrValues field.
func (o *IntrospectionRequest) SetAcrValues(v []string) {
	o.AcrValues = v
}

// GetMaxAge returns the MaxAge field value if set, zero value otherwise.
func (o *IntrospectionRequest) GetMaxAge() int64 {
	if o == nil || isNil(o.MaxAge) {
		var ret int64
		return ret
	}
	return *o.MaxAge
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionRequest) GetMaxAgeOk() (*int64, bool) {
	if o == nil || isNil(o.MaxAge) {
		return nil, false
	}
	return o.MaxAge, true
}

// HasMaxAge returns a boolean if a field has been set.
func (o *IntrospectionRequest) HasMaxAge() bool {
	if o != nil && !isNil(o.MaxAge) {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given int64 and assigns it to the MaxAge field.
func (o *IntrospectionRequest) SetMaxAge(v int64) {
	o.MaxAge = &v
}

// GetRequiredComponents returns the RequiredComponents field value if set, zero value otherwise.
func (o *IntrospectionRequest) GetRequiredComponents() []string {
	if o == nil || isNil(o.RequiredComponents) {
		var ret []string
		return ret
	}
	return o.RequiredComponents
}

// GetRequiredComponentsOk returns a tuple with the RequiredComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionRequest) GetRequiredComponentsOk() ([]string, bool) {
	if o == nil || isNil(o.RequiredComponents) {
		return nil, false
	}
	return o.RequiredComponents, true
}

// HasRequiredComponents returns a boolean if a field has been set.
func (o *IntrospectionRequest) HasRequiredComponents() bool {
	if o != nil && !isNil(o.RequiredComponents) {
		return true
	}

	return false
}

// SetRequiredComponents gets a reference to the given []string and assigns it to the RequiredComponents field.
func (o *IntrospectionRequest) SetRequiredComponents(v []string) {
	o.RequiredComponents = v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *IntrospectionRequest) GetUri() string {
	if o == nil || isNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionRequest) GetUriOk() (*string, bool) {
	if o == nil || isNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *IntrospectionRequest) HasUri() bool {
	if o != nil && !isNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *IntrospectionRequest) SetUri(v string) {
	o.Uri = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *IntrospectionRequest) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionRequest) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *IntrospectionRequest) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *IntrospectionRequest) SetMessage(v string) {
	o.Message = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *IntrospectionRequest) GetHeaders() []Pair {
	if o == nil || isNil(o.Headers) {
		var ret []Pair
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectionRequest) GetHeadersOk() ([]Pair, bool) {
	if o == nil || isNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *IntrospectionRequest) HasHeaders() bool {
	if o != nil && !isNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []Pair and assigns it to the Headers field.
func (o *IntrospectionRequest) SetHeaders(v []Pair) {
	o.Headers = v
}

func (o IntrospectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntrospectionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	if !isNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !isNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
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
	if !isNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !isNil(o.AcrValues) {
		toSerialize["acrValues"] = o.AcrValues
	}
	if !isNil(o.MaxAge) {
		toSerialize["maxAge"] = o.MaxAge
	}
	if !isNil(o.RequiredComponents) {
		toSerialize["requiredComponents"] = o.RequiredComponents
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

type NullableIntrospectionRequest struct {
	value *IntrospectionRequest
	isSet bool
}

func (v NullableIntrospectionRequest) Get() *IntrospectionRequest {
	return v.value
}

func (v *NullableIntrospectionRequest) Set(val *IntrospectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIntrospectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIntrospectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntrospectionRequest(val *IntrospectionRequest) *NullableIntrospectionRequest {
	return &NullableIntrospectionRequest{value: val, isSet: true}
}

func (v NullableIntrospectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntrospectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


