/*
Authlete API

Authlete API Document. 

API version: 2.2.25
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
	"fmt"
)

// JwsAlg The signature algorithm for JWT. This value is represented on 'alg' attribute of the header of JWT.  it's semantics depends upon where is this defined, for instance:   - as service accessTokenSignAlg value, it defines that access token are JWT and the algorithm used to sign it. Check your [KB article](https://kb.authlete.com/en/s/oauth-and-openid-connect/a/jwt-based-access-token).   - as client authorizationSignAlg value, it represents the signature algorithm used when [creating a JARM response](https://kb.authlete.com/en/s/oauth-and-openid-connect/a/enabling-jarm).   - or as client requestSignAlg value, it specifies which is the expected signature used by [client on a Request Object](https://kb.authlete.com/en/s/oauth-and-openid-connect/a/request-objects). 
type JwsAlg string

// List of jws_alg
const (
	JWSALG_NONE JwsAlg = "NONE"
	JWSALG_HS256 JwsAlg = "HS256"
	JWSALG_HS384 JwsAlg = "HS384"
	JWSALG_HS512 JwsAlg = "HS512"
	JWSALG_RS256 JwsAlg = "RS256"
	JWSALG_RS384 JwsAlg = "RS384"
	JWSALG_RS512 JwsAlg = "RS512"
	JWSALG_ES256 JwsAlg = "ES256"
	JWSALG_ES384 JwsAlg = "ES384"
	JWSALG_ES512 JwsAlg = "ES512"
	JWSALG_PS256 JwsAlg = "PS256"
	JWSALG_PS384 JwsAlg = "PS384"
	JWSALG_PS512 JwsAlg = "PS512"
	JWSALG_ES256_K JwsAlg = "ES256K"
	JWSALG_ED_DSA JwsAlg = "EdDSA"
)

// All allowed values of JwsAlg enum
var AllowedJwsAlgEnumValues = []JwsAlg{
	"NONE",
	"HS256",
	"HS384",
	"HS512",
	"RS256",
	"RS384",
	"RS512",
	"ES256",
	"ES384",
	"ES512",
	"PS256",
	"PS384",
	"PS512",
	"ES256K",
	"EdDSA",
}

func (v *JwsAlg) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JwsAlg(value)
	for _, existing := range AllowedJwsAlgEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JwsAlg", value)
}

// NewJwsAlgFromValue returns a pointer to a valid JwsAlg
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJwsAlgFromValue(v string) (*JwsAlg, error) {
	ev := JwsAlg(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JwsAlg: valid values are %v", v, AllowedJwsAlgEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JwsAlg) IsValid() bool {
	for _, existing := range AllowedJwsAlgEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to jws_alg value
func (v JwsAlg) Ptr() *JwsAlg {
	return &v
}

type NullableJwsAlg struct {
	value *JwsAlg
	isSet bool
}

func (v NullableJwsAlg) Get() *JwsAlg {
	return v.value
}

func (v *NullableJwsAlg) Set(val *JwsAlg) {
	v.value = val
	v.isSet = true
}

func (v NullableJwsAlg) IsSet() bool {
	return v.isSet
}

func (v *NullableJwsAlg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJwsAlg(val *JwsAlg) *NullableJwsAlg {
	return &NullableJwsAlg{value: val, isSet: true}
}

func (v NullableJwsAlg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJwsAlg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

