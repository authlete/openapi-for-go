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

// checks if the AuthorizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationRequest{}

// AuthorizationRequest struct for AuthorizationRequest
type AuthorizationRequest struct {
	// OAuth 2.0 authorization request parameters which are the request parameters that the OAuth 2.0 authorization endpoint of the authorization server implementation received from the client application.  The value of parameters is either (1) the entire query string when the HTTP method of the request from the client application is `GET` or (2) the entire entity body (which is formatted in `application/x-www-form-urlencoded`) when the HTTP method of the request from the client application is `POST`.
	Parameters string `json:"parameters"`
}

// NewAuthorizationRequest instantiates a new AuthorizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationRequest(parameters string) *AuthorizationRequest {
	this := AuthorizationRequest{}
	this.Parameters = parameters
	return &this
}

// NewAuthorizationRequestWithDefaults instantiates a new AuthorizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationRequestWithDefaults() *AuthorizationRequest {
	this := AuthorizationRequest{}
	return &this
}

// GetParameters returns the Parameters field value
func (o *AuthorizationRequest) GetParameters() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *AuthorizationRequest) GetParametersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *AuthorizationRequest) SetParameters(v string) {
	o.Parameters = v
}

func (o AuthorizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parameters"] = o.Parameters
	return toSerialize, nil
}

type NullableAuthorizationRequest struct {
	value *AuthorizationRequest
	isSet bool
}

func (v NullableAuthorizationRequest) Get() *AuthorizationRequest {
	return v.value
}

func (v *NullableAuthorizationRequest) Set(val *AuthorizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationRequest(val *AuthorizationRequest) *NullableAuthorizationRequest {
	return &NullableAuthorizationRequest{value: val, isSet: true}
}

func (v NullableAuthorizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


