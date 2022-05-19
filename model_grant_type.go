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

// GrantType The grant type of the access token when the access token was created. 
type GrantType string

// List of grant_type
const (
	GRANTTYPE_AUTHORIZATION_CODE GrantType = "AUTHORIZATION_CODE"
	GRANTTYPE_IMPLICIT GrantType = "IMPLICIT"
	GRANTTYPE_PASSWORD GrantType = "PASSWORD"
	GRANTTYPE_CLIENT_CREDENTIALS GrantType = "CLIENT_CREDENTIALS"
	GRANTTYPE_REFRESH_TOKEN GrantType = "REFRESH_TOKEN"
	GRANTTYPE_CIBA GrantType = "CIBA"
	GRANTTYPE_DEVICE_CODE GrantType = "DEVICE_CODE"
)

// All allowed values of GrantType enum
var AllowedGrantTypeEnumValues = []GrantType{
	"AUTHORIZATION_CODE",
	"IMPLICIT",
	"PASSWORD",
	"CLIENT_CREDENTIALS",
	"REFRESH_TOKEN",
	"CIBA",
	"DEVICE_CODE",
}

func (v *GrantType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GrantType(value)
	for _, existing := range AllowedGrantTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GrantType", value)
}

// NewGrantTypeFromValue returns a pointer to a valid GrantType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGrantTypeFromValue(v string) (*GrantType, error) {
	ev := GrantType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GrantType: valid values are %v", v, AllowedGrantTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GrantType) IsValid() bool {
	for _, existing := range AllowedGrantTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to grant_type value
func (v GrantType) Ptr() *GrantType {
	return &v
}

type NullableGrantType struct {
	value *GrantType
	isSet bool
}

func (v NullableGrantType) Get() *GrantType {
	return v.value
}

func (v *NullableGrantType) Set(val *GrantType) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantType) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantType(val *GrantType) *NullableGrantType {
	return &NullableGrantType{value: val, isSet: true}
}

func (v NullableGrantType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

