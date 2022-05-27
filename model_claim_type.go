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

// ClaimType the model 'ClaimType'
type ClaimType string

// List of claim_type
const (
	CLAIMTYPE_NORMAL ClaimType = "NORMAL"
	CLAIMTYPE_AGGREGATED ClaimType = "AGGREGATED"
	CLAIMTYPE_DISTRIBUTED ClaimType = "DISTRIBUTED"
)

// All allowed values of ClaimType enum
var AllowedClaimTypeEnumValues = []ClaimType{
	"NORMAL",
	"AGGREGATED",
	"DISTRIBUTED",
}

func (v *ClaimType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClaimType(value)
	for _, existing := range AllowedClaimTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClaimType", value)
}

// NewClaimTypeFromValue returns a pointer to a valid ClaimType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClaimTypeFromValue(v string) (*ClaimType, error) {
	ev := ClaimType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClaimType: valid values are %v", v, AllowedClaimTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClaimType) IsValid() bool {
	for _, existing := range AllowedClaimTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to claim_type value
func (v ClaimType) Ptr() *ClaimType {
	return &v
}

type NullableClaimType struct {
	value *ClaimType
	isSet bool
}

func (v NullableClaimType) Get() *ClaimType {
	return v.value
}

func (v *NullableClaimType) Set(val *ClaimType) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimType) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimType(val *ClaimType) *NullableClaimType {
	return &NullableClaimType{value: val, isSet: true}
}

func (v NullableClaimType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
