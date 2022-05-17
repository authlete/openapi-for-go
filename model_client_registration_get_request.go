/*
Authlete API

Authlete API Document. 

API version: 2.2.19
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
)

// ClientRegistrationGetRequest struct for ClientRegistrationGetRequest
type ClientRegistrationGetRequest struct {
	// Client ID. 
	ClientId string `json:"clientId"`
	// Client registration access token.
	Token string `json:"token"`
}

// NewClientRegistrationGetRequest instantiates a new ClientRegistrationGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientRegistrationGetRequest(clientId string, token string) *ClientRegistrationGetRequest {
	this := ClientRegistrationGetRequest{}
	this.ClientId = clientId
	this.Token = token
	return &this
}

// NewClientRegistrationGetRequestWithDefaults instantiates a new ClientRegistrationGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientRegistrationGetRequestWithDefaults() *ClientRegistrationGetRequest {
	this := ClientRegistrationGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *ClientRegistrationGetRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *ClientRegistrationGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *ClientRegistrationGetRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetToken returns the Token field value
func (o *ClientRegistrationGetRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ClientRegistrationGetRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ClientRegistrationGetRequest) SetToken(v string) {
	o.Token = v
}

func (o ClientRegistrationGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["clientId"] = o.ClientId
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableClientRegistrationGetRequest struct {
	value *ClientRegistrationGetRequest
	isSet bool
}

func (v NullableClientRegistrationGetRequest) Get() *ClientRegistrationGetRequest {
	return v.value
}

func (v *NullableClientRegistrationGetRequest) Set(val *ClientRegistrationGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClientRegistrationGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClientRegistrationGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientRegistrationGetRequest(val *ClientRegistrationGetRequest) *NullableClientRegistrationGetRequest {
	return &NullableClientRegistrationGetRequest{value: val, isSet: true}
}

func (v NullableClientRegistrationGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientRegistrationGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


