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

// ClientSecretUpdateRequest struct for ClientSecretUpdateRequest
type ClientSecretUpdateRequest struct {
	// The new value of the client secret. Valid characters for a client secret are `A-Z`, `a-z`, `0-9`, `-`, and `_`. The maximum length of a client secret is 86.
	ClientSecret string `json:"clientSecret"`
}

// NewClientSecretUpdateRequest instantiates a new ClientSecretUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientSecretUpdateRequest(clientSecret string) *ClientSecretUpdateRequest {
	this := ClientSecretUpdateRequest{}
	this.ClientSecret = clientSecret
	return &this
}

// NewClientSecretUpdateRequestWithDefaults instantiates a new ClientSecretUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientSecretUpdateRequestWithDefaults() *ClientSecretUpdateRequest {
	this := ClientSecretUpdateRequest{}
	return &this
}

// GetClientSecret returns the ClientSecret field value
func (o *ClientSecretUpdateRequest) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *ClientSecretUpdateRequest) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *ClientSecretUpdateRequest) SetClientSecret(v string) {
	o.ClientSecret = v
}

func (o ClientSecretUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	return json.Marshal(toSerialize)
}

type NullableClientSecretUpdateRequest struct {
	value *ClientSecretUpdateRequest
	isSet bool
}

func (v NullableClientSecretUpdateRequest) Get() *ClientSecretUpdateRequest {
	return v.value
}

func (v *NullableClientSecretUpdateRequest) Set(val *ClientSecretUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClientSecretUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClientSecretUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientSecretUpdateRequest(val *ClientSecretUpdateRequest) *NullableClientSecretUpdateRequest {
	return &NullableClientSecretUpdateRequest{value: val, isSet: true}
}

func (v NullableClientSecretUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientSecretUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


