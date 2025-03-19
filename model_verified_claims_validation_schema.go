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

// VerifiedClaimsValidationSchema The verified claims validation schema set.
type VerifiedClaimsValidationSchema string

// List of verified_claims_validation_schema
const (
	VERIFIEDCLAIMSVALIDATIONSCHEMA_STANDARD            VerifiedClaimsValidationSchema = "standard"
	VERIFIEDCLAIMSVALIDATIONSCHEMA_STANDARDID_DOCUMENT VerifiedClaimsValidationSchema = "standard+id_document"
)

// All allowed values of VerifiedClaimsValidationSchema enum
var AllowedVerifiedClaimsValidationSchemaEnumValues = []VerifiedClaimsValidationSchema{
	"standard",
	"standard+id_document",
}

func (v *VerifiedClaimsValidationSchema) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VerifiedClaimsValidationSchema(value)
	for _, existing := range AllowedVerifiedClaimsValidationSchemaEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VerifiedClaimsValidationSchema", value)
}

// NewVerifiedClaimsValidationSchemaFromValue returns a pointer to a valid VerifiedClaimsValidationSchema
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVerifiedClaimsValidationSchemaFromValue(v string) (*VerifiedClaimsValidationSchema, error) {
	ev := VerifiedClaimsValidationSchema(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VerifiedClaimsValidationSchema: valid values are %v", v, AllowedVerifiedClaimsValidationSchemaEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VerifiedClaimsValidationSchema) IsValid() bool {
	for _, existing := range AllowedVerifiedClaimsValidationSchemaEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to verified_claims_validation_schema value
func (v VerifiedClaimsValidationSchema) Ptr() *VerifiedClaimsValidationSchema {
	return &v
}

type NullableVerifiedClaimsValidationSchema struct {
	value *VerifiedClaimsValidationSchema
	isSet bool
}

func (v NullableVerifiedClaimsValidationSchema) Get() *VerifiedClaimsValidationSchema {
	return v.value
}

func (v *NullableVerifiedClaimsValidationSchema) Set(val *VerifiedClaimsValidationSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifiedClaimsValidationSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifiedClaimsValidationSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifiedClaimsValidationSchema(val *VerifiedClaimsValidationSchema) *NullableVerifiedClaimsValidationSchema {
	return &NullableVerifiedClaimsValidationSchema{value: val, isSet: true}
}

func (v NullableVerifiedClaimsValidationSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifiedClaimsValidationSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
