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

// JweAlg this is the 'alg' header value for encrypted JWT tokens. Depending upon the context, this refers to key transport scheme to be used by the client and by the server. For instance: - as `authorizationEncryptionAlg` value, it refers to the encoding algorithm used by server for transporting they keys on JARM objects - as `requestEncryptionAlg` value, it refers to the expected key transport encoding algorithm that server expect from client when encrypting a Request Object - as `idTokenEncryptionAlg` value, it refers to the algorithm used by the server to key transport of id_tokens  **Please note that some of the algorithms are more secure than others, some are not supported very well cross platforms and some (like RSA1_5) is known to be weak**. 
type JweAlg string

// List of jwe_alg
const (
	JWEALG_RSA1_5 JweAlg = "RSA1_5"
	JWEALG_RSA_OAEP JweAlg = "RSA_OAEP"
	JWEALG_RSA_OAEP_256 JweAlg = "RSA_OAEP_256"
	JWEALG_A128_KW JweAlg = "A128KW"
	JWEALG_A192_KW JweAlg = "A192KW"
	JWEALG_A256_KW JweAlg = "A256KW"
	JWEALG_DIR JweAlg = "DIR"
	JWEALG_ECDH_ES JweAlg = "ECDH_ES"
	JWEALG_ECDH_ES_A128_KW JweAlg = "ECDH_ES_A128KW"
	JWEALG_ECDH_ES_A192_KW JweAlg = "ECDH_ES_A192KW"
	JWEALG_ECDH_ES_A256_KW JweAlg = "ECDH_ES_A256KW"
	JWEALG_A128_GCMKW JweAlg = "A128GCMKW"
	JWEALG_A192_GCMKW JweAlg = "A192GCMKW"
	JWEALG_A256_GCMKW JweAlg = "A256GCMKW"
	JWEALG_PBES2_HS256_A128_KW JweAlg = "PBES2_HS256_A128KW"
	JWEALG_PBES2_HS384_A192_KW JweAlg = "PBES2_HS384_A192KW"
	JWEALG_PBES2_HS512_A256_KW JweAlg = "PBES2_HS512_A256KW"
)

// All allowed values of JweAlg enum
var AllowedJweAlgEnumValues = []JweAlg{
	"RSA1_5",
	"RSA_OAEP",
	"RSA_OAEP_256",
	"A128KW",
	"A192KW",
	"A256KW",
	"DIR",
	"ECDH_ES",
	"ECDH_ES_A128KW",
	"ECDH_ES_A192KW",
	"ECDH_ES_A256KW",
	"A128GCMKW",
	"A192GCMKW",
	"A256GCMKW",
	"PBES2_HS256_A128KW",
	"PBES2_HS384_A192KW",
	"PBES2_HS512_A256KW",
}

func (v *JweAlg) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JweAlg(value)
	for _, existing := range AllowedJweAlgEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JweAlg", value)
}

// NewJweAlgFromValue returns a pointer to a valid JweAlg
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJweAlgFromValue(v string) (*JweAlg, error) {
	ev := JweAlg(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JweAlg: valid values are %v", v, AllowedJweAlgEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JweAlg) IsValid() bool {
	for _, existing := range AllowedJweAlgEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to jwe_alg value
func (v JweAlg) Ptr() *JweAlg {
	return &v
}

type NullableJweAlg struct {
	value *JweAlg
	isSet bool
}

func (v NullableJweAlg) Get() *JweAlg {
	return v.value
}

func (v *NullableJweAlg) Set(val *JweAlg) {
	v.value = val
	v.isSet = true
}

func (v NullableJweAlg) IsSet() bool {
	return v.isSet
}

func (v *NullableJweAlg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJweAlg(val *JweAlg) *NullableJweAlg {
	return &NullableJweAlg{value: val, isSet: true}
}

func (v NullableJweAlg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJweAlg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
