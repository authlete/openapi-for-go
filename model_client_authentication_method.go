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

// ClientAuthenticationMethod The client authentication method that the client application declares that it uses at the token endpoint. This property corresponds to `token_endpoint_auth_method` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).
type ClientAuthenticationMethod string

// List of ClientAuthenticationMethod
const (
	CLIENTAUTHENTICATIONMETHOD_NONE                        ClientAuthenticationMethod = "NONE"
	CLIENTAUTHENTICATIONMETHOD_CLIENT_SECRET_BASIC         ClientAuthenticationMethod = "CLIENT_SECRET_BASIC"
	CLIENTAUTHENTICATIONMETHOD_CLIENT_SECRET_POST          ClientAuthenticationMethod = "CLIENT_SECRET_POST"
	CLIENTAUTHENTICATIONMETHOD_CLIENT_SECRET_JWT           ClientAuthenticationMethod = "CLIENT_SECRET_JWT"
	CLIENTAUTHENTICATIONMETHOD_PRIVATE_KEY_JWT             ClientAuthenticationMethod = "PRIVATE_KEY_JWT"
	CLIENTAUTHENTICATIONMETHOD_TLS_CLIENT_AUTH             ClientAuthenticationMethod = "TLS_CLIENT_AUTH"
	CLIENTAUTHENTICATIONMETHOD_SELF_SIGNED_TLS_CLIENT_AUTH ClientAuthenticationMethod = "SELF_SIGNED_TLS_CLIENT_AUTH"
)

// All allowed values of ClientAuthenticationMethod enum
var AllowedClientAuthenticationMethodEnumValues = []ClientAuthenticationMethod{
	"NONE",
	"CLIENT_SECRET_BASIC",
	"CLIENT_SECRET_POST",
	"CLIENT_SECRET_JWT",
	"PRIVATE_KEY_JWT",
	"TLS_CLIENT_AUTH",
	"SELF_SIGNED_TLS_CLIENT_AUTH",
}

func (v *ClientAuthenticationMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClientAuthenticationMethod(value)
	for _, existing := range AllowedClientAuthenticationMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClientAuthenticationMethod", value)
}

// NewClientAuthenticationMethodFromValue returns a pointer to a valid ClientAuthenticationMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClientAuthenticationMethodFromValue(v string) (*ClientAuthenticationMethod, error) {
	ev := ClientAuthenticationMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClientAuthenticationMethod: valid values are %v", v, AllowedClientAuthenticationMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClientAuthenticationMethod) IsValid() bool {
	for _, existing := range AllowedClientAuthenticationMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClientAuthenticationMethod value
func (v ClientAuthenticationMethod) Ptr() *ClientAuthenticationMethod {
	return &v
}

type NullableClientAuthenticationMethod struct {
	value *ClientAuthenticationMethod
	isSet bool
}

func (v NullableClientAuthenticationMethod) Get() *ClientAuthenticationMethod {
	return v.value
}

func (v *NullableClientAuthenticationMethod) Set(val *ClientAuthenticationMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableClientAuthenticationMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableClientAuthenticationMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientAuthenticationMethod(val *ClientAuthenticationMethod) *NullableClientAuthenticationMethod {
	return &NullableClientAuthenticationMethod{value: val, isSet: true}
}

func (v NullableClientAuthenticationMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientAuthenticationMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
