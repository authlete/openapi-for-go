/*
Authlete API

Authlete API Document. 

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
	"fmt"
)

// SubjectType The subject type that the client application requests. Details about the subject type are described in [OpenID Connect Core 1.0, 8. Subjct Identifier Types](https://openid.net/specs/openid-connect-core-1_0.html#SubjectIDTypes).  This property corresponds to `subject_type` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata). 
type SubjectType string

// List of subject_type
const (
	SUBJECTTYPE_PUBLIC SubjectType = "PUBLIC"
	SUBJECTTYPE_PAIRWISE SubjectType = "PAIRWISE"
)

// All allowed values of SubjectType enum
var AllowedSubjectTypeEnumValues = []SubjectType{
	"PUBLIC",
	"PAIRWISE",
}

func (v *SubjectType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubjectType(value)
	for _, existing := range AllowedSubjectTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubjectType", value)
}

// NewSubjectTypeFromValue returns a pointer to a valid SubjectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubjectTypeFromValue(v string) (*SubjectType, error) {
	ev := SubjectType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubjectType: valid values are %v", v, AllowedSubjectTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubjectType) IsValid() bool {
	for _, existing := range AllowedSubjectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to subject_type value
func (v SubjectType) Ptr() *SubjectType {
	return &v
}

type NullableSubjectType struct {
	value *SubjectType
	isSet bool
}

func (v NullableSubjectType) Get() *SubjectType {
	return v.value
}

func (v *NullableSubjectType) Set(val *SubjectType) {
	v.value = val
	v.isSet = true
}

func (v NullableSubjectType) IsSet() bool {
	return v.isSet
}

func (v *NullableSubjectType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubjectType(val *SubjectType) *NullableSubjectType {
	return &NullableSubjectType{value: val, isSet: true}
}

func (v NullableSubjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubjectType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

