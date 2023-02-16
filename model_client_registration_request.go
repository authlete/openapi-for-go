/*
Authlete API

Authlete API Document. 

API version: 2.2.30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
)

// checks if the ClientRegistrationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientRegistrationRequest{}

// ClientRegistrationRequest struct for ClientRegistrationRequest
type ClientRegistrationRequest struct {
	// Client metadata in JSON format that complies with [RFC 7591](https://datatracker.ietf.org/doc/html/rfc7591) (OAuth 2.0 Dynamic Client Registration Protocol). 
	Json string `json:"json"`
	// The client registration access token. Used only for GET, UPDATE, and DELETE requests. 
	Token *string `json:"token,omitempty"`
	// The client's identifier. Used for GET, UPDATE, and DELETE requests 
	ClientId *string `json:"clientId,omitempty"`
}

// NewClientRegistrationRequest instantiates a new ClientRegistrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientRegistrationRequest(json string) *ClientRegistrationRequest {
	this := ClientRegistrationRequest{}
	this.Json = json
	return &this
}

// NewClientRegistrationRequestWithDefaults instantiates a new ClientRegistrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientRegistrationRequestWithDefaults() *ClientRegistrationRequest {
	this := ClientRegistrationRequest{}
	return &this
}

// GetJson returns the Json field value
func (o *ClientRegistrationRequest) GetJson() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Json
}

// GetJsonOk returns a tuple with the Json field value
// and a boolean to check if the value has been set.
func (o *ClientRegistrationRequest) GetJsonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Json, true
}

// SetJson sets field value
func (o *ClientRegistrationRequest) SetJson(v string) {
	o.Json = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ClientRegistrationRequest) GetToken() string {
	if o == nil || isNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRegistrationRequest) GetTokenOk() (*string, bool) {
	if o == nil || isNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ClientRegistrationRequest) HasToken() bool {
	if o != nil && !isNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ClientRegistrationRequest) SetToken(v string) {
	o.Token = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ClientRegistrationRequest) GetClientId() string {
	if o == nil || isNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRegistrationRequest) GetClientIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ClientRegistrationRequest) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ClientRegistrationRequest) SetClientId(v string) {
	o.ClientId = &v
}

func (o ClientRegistrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientRegistrationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["json"] = o.Json
	if !isNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !isNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	return toSerialize, nil
}

type NullableClientRegistrationRequest struct {
	value *ClientRegistrationRequest
	isSet bool
}

func (v NullableClientRegistrationRequest) Get() *ClientRegistrationRequest {
	return v.value
}

func (v *NullableClientRegistrationRequest) Set(val *ClientRegistrationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClientRegistrationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClientRegistrationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientRegistrationRequest(val *ClientRegistrationRequest) *NullableClientRegistrationRequest {
	return &NullableClientRegistrationRequest{value: val, isSet: true}
}

func (v NullableClientRegistrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientRegistrationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


