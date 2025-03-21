/*
Authlete API

Authlete API Document.

API version: 2.3.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
	"fmt"
)

// ClientAuthMethod The client authentication method that the client application declares that it uses at the token endpoint. This property corresponds to `token_endpoint_auth_method` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).
type ClientAuthMethod string

// List of client_auth_method
const (
	CLIENTAUTHMETHOD_NONE                        ClientAuthMethod = "NONE"
	CLIENTAUTHMETHOD_CLIENT_SECRET_BASIC         ClientAuthMethod = "CLIENT_SECRET_BASIC"
	CLIENTAUTHMETHOD_CLIENT_SECRET_POST          ClientAuthMethod = "CLIENT_SECRET_POST"
	CLIENTAUTHMETHOD_CLIENT_SECRET_JWT           ClientAuthMethod = "CLIENT_SECRET_JWT"
	CLIENTAUTHMETHOD_PRIVATE_KEY_JWT             ClientAuthMethod = "PRIVATE_KEY_JWT"
	CLIENTAUTHMETHOD_TLS_CLIENT_AUTH             ClientAuthMethod = "TLS_CLIENT_AUTH"
	CLIENTAUTHMETHOD_SELF_SIGNED_TLS_CLIENT_AUTH ClientAuthMethod = "SELF_SIGNED_TLS_CLIENT_AUTH"
)

// All allowed values of ClientAuthMethod enum
var AllowedClientAuthMethodEnumValues = []ClientAuthMethod{
	"NONE",
	"CLIENT_SECRET_BASIC",
	"CLIENT_SECRET_POST",
	"CLIENT_SECRET_JWT",
	"PRIVATE_KEY_JWT",
	"TLS_CLIENT_AUTH",
	"SELF_SIGNED_TLS_CLIENT_AUTH",
}

func (v *ClientAuthMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClientAuthMethod(value)
	for _, existing := range AllowedClientAuthMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClientAuthMethod", value)
}

// NewClientAuthMethodFromValue returns a pointer to a valid ClientAuthMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClientAuthMethodFromValue(v string) (*ClientAuthMethod, error) {
	ev := ClientAuthMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClientAuthMethod: valid values are %v", v, AllowedClientAuthMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClientAuthMethod) IsValid() bool {
	for _, existing := range AllowedClientAuthMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to client_auth_method value
func (v ClientAuthMethod) Ptr() *ClientAuthMethod {
	return &v
}

type NullableClientAuthMethod struct {
	value *ClientAuthMethod
	isSet bool
}

func (v NullableClientAuthMethod) Get() *ClientAuthMethod {
	return v.value
}

func (v *NullableClientAuthMethod) Set(val *ClientAuthMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableClientAuthMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableClientAuthMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientAuthMethod(val *ClientAuthMethod) *NullableClientAuthMethod {
	return &NullableClientAuthMethod{value: val, isSet: true}
}

func (v NullableClientAuthMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientAuthMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
